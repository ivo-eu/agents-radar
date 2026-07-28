# AI CLI 工具社区动态日报 2026-07-28

> 生成时间: 2026-07-28 00:11 UTC | 覆盖工具: 9 个

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

好的，作为专注于 AI 开发工具生态的资深技术分析师，我已根据您提供的 2026-07-28 各工具社区动态，为您整理出以下横向对比分析报告。

---

## AI CLI 开发工具生态横向对比分析报告 (2026-07-28)

### 1. 生态全景

2026年7月下旬，AI CLI 开发工具生态进入了 **“体验精细化”与“稳定性攻坚”** 并行的阶段。一方面，以 Claude Code、OpenAI Codex 为代表的头部工具正面临社区对计费模型、系统级Bug及跨平台稳定性的集中质疑；另一方面，以 Gemini CLI、Pi、OpenCode 为代表的工具则更聚焦于 Agent 行为的可预测性、扩展能力及上下文管理。整体而言，社区不再满足于“能用”，而是对“可靠”、“可控”和“智能”提出了更高要求，企业级集成（如外部上下文、记忆）和用户体验细节（如 Git 工作流、语音交互）成为新突破口。

### 2. 各工具活跃度对比

| 工具名称 | Issues 活跃 (Top 10) | PR 活跃 (Top 10) | 版本发布 | 核心关注点 |
| :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 极高 (10条，评论/赞密集) | 高 (5条) | 无新版本 | 计费故障、模型访问、系统级Bug (剪贴板/崩溃) |
| **OpenAI Codex** | 高 (10条，评论多) | 极高 (10条，合并密集) | 2个 Alpha 版 | 计费/重置问题、Windows 崩溃、磁盘占用 |
| **Gemini CLI** | 高 (10条，评论/赞活跃) | 高 (10条) | 1个 Nightly 版 | Agent 挂起与误报、上下文加载、新模型支持 |
| **GitHub Copilot CLI** | 高 (7条有效) | 中 (10条，多为文档/配置) | 1个正式版 | Plan模式回归、交互显示异常、子进程管理 |
| **Kimi Code CLI** | 低 (4条新/更) | 中 (4条) | 无新版本 | VS Code 扩展稳定性、Windows 编码兼容性 |
| **OpenCode** | 高 (10条) | 高 (10条) | 2个小版本 | Agent 循环/恢复、模型兼容性、桌面交互 |
| **Pi** | 极高 (28条总量，10条精选) | 高 (10条) | 无新版本 | 扩展 API、会话管理、多模态 tokenizer |
| **DeepSeek TUI** | 低 (5条) | 极高 (10条，大规模合并) | 无 (v0.9.2候选集成) | TUI 视觉、会话持久化、Fleet 控制面板 |

**分析**:
-   **Claude Code** 和 **OpenAI Codex** 因其庞大用户基数，社区反馈量级最高，但也暴露出最多的“成长的烦恼”。
-   **OpenCode** 和 **Pi** 处于功能扩张与稳定性并重的阶段，社区互动非常活跃。
-   **DeepSeek TUI** 在这一天完成了大规模代码合并，处于发布候选冲刺阶段，动效惊人。

### 3. 共同关注的功能方向

| 共同方向 | 涉及工具 | 具体诉求体现 |
| :--- | :--- | :--- |
| **Agent 行为可靠性** | Gemini CLI, OpenCode, Pi, DeepSeek TUI | 子代理误报成功/挂起/失败后无法恢复；工具调用陷入无限循环；前台 shell 阻塞交互。 |
| **跨平台兼容性** | Claude Code, OpenAI Codex, GitHub Copilot CLI, Kimi Code, OpenCode | Windows GPU崩溃/编码问题；macOS 剪贴板/全屏Bug；Linux Wayland/WSL 兼容性问题。 |
| **计费与配额透明度** | Claude Code, OpenAI Codex, GitHub Copilot CLI | Fable 5 计费故障；重置按钮无效/计数错误；ACP模式下Token用量不透明。 |
| **企业级集成** | Gemini CLI, Qwen Code, Claude Code | 外部上下文/记忆集成；MCP集成稳定性、认证方式、角色权限管理。 |
| **IDE/扩展生态** | Kimi Code CLI, Pi, OpenCode | VS Code 扩展 UI Bug；提供Agent状态/模型的扩展API；文件链接可点击。 |
| **用户体验细节** | Gemini CLI, GitHub Copilot CLI, OpenCode, DeepSeek TUI | 配置项持久化；GEMINI.md文件被忽略；主题修改受限；思考块默认展开。 |

### 4. 差异化定位分析

-   **Claude Code & OpenAI Codex (头部玩家)**: 背靠顶级模型，功能全、用户多，但正面临“大公司病”——系统复杂导致Bug频发，计费模型成为核心矛盾。社区反馈显示，它们正从“功能领先”转向“稳定性补课”。
-   **GitHub Copilot CLI (平台绑定)**: 深度集成 GitHub 生态，侧重代码审查流程（Plan模式、Diff功能）。其最大优势是组织管理能力，但最大风险是受限于 GitHub 生态，灵活性不足。
-   **Gemini CLI (生态开放)**: 背靠 Google，侧重 Agent 与工具的智能化编排（多Agent、自定义技能）。社区更关注其**抽象能力**，即Agent是否能“聪明地”使用工具。
-   **Kimi Code CLI (本地化)**: 聚焦中文环境和中国开发者体验（Windows编码、中文翻译）。其核心挑战在于**扩展（VS Code）的稳定性**。
-   **OpenCode (激进创新)**: 定位为“开源实验场”，追求高度灵活性和模型兼容性（如硬编码 vs 动态上下文），社区讨论技术方案多于使用反馈。
-   **Pi (开发工具IDE化)**: 致力于将终端CLI打造成类似VS Code的 **IDE**，核心逻辑围绕扩展API、Markdown渲染、FTS5搜索等，是“扩展才是未来”理念的代表。
-   **DeepSeek TUI (极客社区)**: 侧重前沿技术探索，如 `Fleet` 多Agent调度架构，对社区协同开发流程极为成熟（大规模PR合并）。其代码质量与架构设计备受推崇。

### 5. 社区热度与成熟度

-   **成熟型**: **Claude Code** 和 **OpenAI Codex** 社区规模最大，但负面情绪占比高，正处于从“功能尝鲜”到“稳定可靠”的转型阵痛期。
-   **活跃迭代型**: **Pi**、**OpenCode** 和 **Gemini CLI** 社区活跃度极高，功能更新快，敢于实验新技术。其成熟度体现在社区能够提出并讨论深层技术问题（如上下文保真度、扩展API设计），而非仅停留在“如何配置”的层面。
-   **快速成长型**: **DeepSeek TUI** 代码质量和发展速度惊人，社区规模虽小但参与度极深。**Qwen Code** 处于功能增强的加速期，其 Web Shell 和语音功能是其亮点。

### 6. 值得关注的趋势信号

1.  **Agent 可靠性的“信任危机”**: 多个工具的社区都反馈了 Agent 行为不可预测的问题（挂起、谎报、死循环）。这已成为阻碍AI CLI工具从“玩具”走向“生产力工具”的核心瓶颈。**对开发者而言，选择 Agent 时，应优先考察其在失败场景下的恢复与报告机制，而不仅仅是功能多寡。**

2.  **“开放生态”成为竞争优势**: 企业级集成需求（外部上下文、记忆、MCP兼容性）在 Qwen Code、Gemini CLI 和 Pi 中呼声很高。**能提供灵活、标准化的集成接口，而非封闭的模型+服务，将成为未来工具的核心竞争力。**

3.  **本地化与国际化同样重要**: Kimi Code 对 Windows 编码的修复，DeepSeek TUI 对中文翻译的二次润色，表明**服务非英语/特殊平台用户**是扩大市场的关键。工具开发者不应忽视这些“细分”痛点。

4.  **自动化/无头场景短板突出**: Claude Code (`setup-token`)、GitHub Copilot CLI (ACP模式) 和 OpenAI Codex (子代理) 都暴露了在 CI/CD、远程或非交互式场景下的功能缺失或计费问题。**这表明大多数工具仍以交互式使用为基础构建，对自动化场景的支持是下一阶段需要补全的重要能力。**

5.  **从“功能竞争”到“体验竞争”**: 修复 macOS 剪贴板、Windows 终端渲染、`glob` 工具假阴性等“小而精”的Bug占据了很大讨论量。**这说明社区对工具的要求已从“能做什么”升级为“用得是否舒服”。任何反直觉或破坏系统级行为的Bug，都将极大损害用户忠诚度。**

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为一名专注于 Claude Code 生态的技术分析师，以下是基于 `anthropics/skills` 仓库数据（截至 2026-07-28）的社区热点报告。

---

### Claude Code Skills 社区热点报告 (2026-07-28)

#### 1. 热门 Skills 排行 (按 PR 活跃度与影响力)

基于评论数、关联 Issue 数量和问题严重性，以下是最受社区关注的几个 Skills (PR)：

1.  **`fix(skill-creator)` (PR #1298, #1099, #1050, #1323)**:
    -   **功能**: 修复 `skill-creator` 工具链的核心问题，特别是 `run_eval.py` 在 Windows 平台上的崩溃和 0% 召回率问题。这是整个 Skills 生态的开发与评估基础工具。
    -   **社区热点**: **此为当前社区最集中的痛点**。多个 PR 和 Issues (如 #556, #1169, #1061) 反复提及在 Windows 上运行评估脚本时，所有测试查询均显示“未触发”，导致描述优化循环失效。社区高度关注跨平台兼容性、子进程处理、编码（cp1252）及文件描述符问题。
    -   **状态**: `OPEN` (多个相关 PR 构成修复集群).
    -   **链接**: [PR #1298](https://github.com/anthropics/skills/pull/1298) | [PR #1099](https://github.com/anthropics/skills/pull/1099) | [PR #1050](https://github.com/anthropics/skills/pull/1050) | [PR #1323](https://github.com/anthropics/skills/pull/1323) | [Issue #556](https://github.com/anthropics/skills/issues/556)

2.  **`Add document-typography skill` (PR #514)**:
    -   **功能**: 解决 AI 生成文档中常见的排版问题，如孤行、寡段和编号错位。这是一项提升输出专业度的实用 Skill。
    -   **社区热点**: 讨论集中在 AI 输出中普遍存在的排版瑕疵，社区认为这是提升文档“交付出手感”的关键一步，但用户可能不会明确提出需求（未明确说出问题）。
    -   **状态**: `OPEN`.
    -   **链接**: [PR #514](https://github.com/anthropics/skills/pull/514)

3.  **`Add self-audit skill` (PR #1367)**:
    -   **功能**: 引入“推理质量门”机制，在交付前对 AI 输出进行机械检查（文件完整性）和四维度推理审计，旨在提升最终交付物的可靠性和逻辑自洽性。
    -   **社区热点**: 该 Skill 回应了社区对 AI 生成内容质量控制的长期需求。Issues #1385 进一步讨论了将其扩展为“预任务校准-对抗性审查-交付验证”的三门管线方案，体现了社区对系统化质量保证的追求。
    -   **状态**: `OPEN` (近期提交，活跃度高).
    -   **链接**: [PR #1367](https://github.com/anthropics/skills/pull/1367) | [Issue #1385](https://github.com/anthropics/skills/issues/1385)

4.  **`Add testing-patterns skill` (PR #723)**:
    -   **功能**: 提供覆盖单元测试、React 组件测试、E2E 测试的综合性测试模式指南，旨在规范 Claude 的测试代码生成行为。
    -   **社区热点**: 社区对提高 AI 生成测试代码的质量和一致性有强烈需求。该 Skill 完整覆盖测试堆栈，符合行业最佳实践，被认为能显著提升开发效率。
    -   **状态**: `OPEN`.
    -   **链接**: [PR #723](https://github.com/anthropics/skills/pull/723)

5.  **`Add pyxel skill` (PR #525)**:
    -   **功能**: 将 Pyxel 复古游戏引擎与 Claude Code 连接，使用户能通过自然语言描述创建像素风格游戏。
    -   **社区热点**: 代表社区在**创意编码**领域的探索。尽管受众相对垂直，但该 Skill 展示了 Skills 生态与技术生态（MCP 服务器）的整合能力，是社区创新的一个标志性案例。
    -   **状态**: `OPEN` (长期未合并，但持续受关注).
    -   **链接**: [PR #525](https://github.com/anthropics/skills/pull/525)

6.  **`Add SAP-RPT-1-OSS predictor skill` (PR #181)**:
    -   **功能**: 允许 Claude 直接使用 SAP 的开源表格基础模型进行预测分析。
    -   **社区热点**: 代表了 Skills 生态向**垂直行业和企业级应用**的拓展。社区关注点不仅在于功能本身，更在于其作为连接 Claude 与复杂企业 AI 模型的桥梁价值。
    -   **状态**: `OPEN`.
    -   **链接**: [PR #181](https://github.com/anthropics/skills/pull/181)

7.  **`Add plan-file-hygiene skill` (PR #1479)**:
    -   **功能**: 提出“规划文件卫生”概念，旨在管理长期运行 Agent 中积累的规划文件，避免其无节制增长，成为“上下文窗口杀手”。
    -   **社区热点**: 该提案回应了 #1417 号 Issue 中社区对**长期会话和 Agent 状态管理**的痛点。其核心理念（文件生命周期管理）被认为是 Agent 工程中一个被忽视但至关重要的领域。
    -   **状态**: `OPEN` (最新 PR，社区响应积极).
    -   **链接**: [PR #1479](https://github.com/anthropics/skills/pull/1479) | [Issue #1417](https://github.com/anthropics/skills/issues/1417)

#### 2. 社区需求趋势 (从 Issues 提炼)

1.  **安全与信任边界**: 社区对 Skill 命名空间滥用（#492）和权限控制（#1175）高度关注。核心诉求是 **明确官方与社区 Skills 的边界，并建立可审计的安全机制**。
2.  **企业级功能与治理**: 强烈需求组织级的 Skill 共享（#228）、代理系统安全模式（#412），以及对 AI 输出的系统化质量审核（#1385）。社区正从个人使用转向团队协作和企业应用，对治理和流程有更高要求。
3.  **跨平台与工具链稳定性**: `skill-creator` 在 Windows 上的`run_eval.py` 0% 召回率问题（#556, #1169, #1061）是当前最大的社区技术债务。这表明**开发者工具的可靠性与跨平台兼容性是生态繁荣的基石**。
4.  **上下文窗口优化与管理**: 社区持续关注上下文窗口消耗问题，如 `claude-api` Skill 注入过多 token ( #1487 )，以及 Agent 长期运行中的状态膨胀（#1329, #1479）。**“上下文窗口卫生”已成为一个核心设计原则**。
5.  **明确的功能定位与避免重复**: 社区指出官方 `document-skills` 和 `example-skills` 存在内容重复（#189），并期望 `skill-creator` 本身能成为更符合最佳实践的 Skill（#202）。社区希望官方仓库能提供更清晰、无冗余的模块划分。

#### 3. 高潜力待合并 Skills

1.  **`skill-creator` 修复集群**：以 PR #1298 为核心的 Windows 兼容性修复集群，合并价值最高。它们是 Skills 生态的开发者基础设施，其阻塞问题直接影响所有社区贡献者的开发体验。
2.  **`self-audit` (PR #1367)**：该 Skill 主题新颖、功能完整，且与社区对质量保证的关注高度契合。它提出了一个明确的、可集成到工作流中的“质量控制门”方案，合并后将成为一项基础能力。
3.  **`plan-file-hygiene` (PR #1479)**：直击 Agent 开发的痛点，是一个经过社区讨论的“面向问题”的解决方案。由于其解决的是通用性问题，合并后对提升长期 Agent 会话的稳定性有显著帮助。
4.  **`testing-patterns` (PR #723)**：内容完整且成熟，能显著提升 Claude 在编码任务中的可用性。一旦合并，它将成为开发者最常调用的 Skills 之一。
5.  **`document-typography` (PR #514)**：功能独立且价值明确——解决一个无人提及但又无处不在的痛点。合并后对用户体验的提升是立竿见影的。

#### 4. Skills 生态洞察

**当前社区最集中的总体诉求是：围绕 `skill-creator` 工具链的可靠性、跨平台可用性，以及一个能够安全、高效管理 Agent 长期任务的生命周期治理体系。**

这个诉求体现为两个并行但密切相关的主题：一是向下夯实基础，修复开发工具 Bugs，提升开发者体验；二是向上构建框架，为 Agent 的长期运行、输出质量和企业级应用提供结构和安全保障。

---

好的，作为专注于 AI 开发工具的技术分析师，我将根据您提供的 GitHub 数据，为您生成 2026 年 7 月 28 日的 Claude Code 社区动态日报。

---

## Claude Code 社区动态日报 | 2026-07-28

### 今日速览

1.  **Fable 5 计费故障持续发酵**：尽管 Fable 5 已成为 Max 套餐的标准配置，但大量 Max 用户报告在 CLI 和桌面端无法使用该模型，被错误地要求购买“使用额度”（usage credits），社区反应激烈。
2.  **GitHub Connector 回归性故障**：一个影响广泛的 Bug 导致 Claude Code 无法读取任何仓库内容（包括公开仓库），已持续一个月仍未修复，严重影响了开发体验。
3.  **多平台稳定性问题涌现**：Windows 桌面版 GPU 进程崩溃、macOS 全屏模式破坏系统剪贴板等问题被密集报告，显示出跨平台兼容性仍是主要痛点。

### 社区热点 Issues

1.  **[BUG] Fable 5 在 Max 套餐中被锁定，要求使用额度** [#79337](https://github.com/anthropics/claude-code/issues/79337)
    -   **重要性**：⭐⭐⭐⭐⭐
    -   **摘要**：Fable 5 被列为 Max 套餐的一部分后，用户在 `Max` 套餐中却无法使用，会话被静默降级到 Opus 4.8，并提示需要购买使用额度。该问题引发了大量用户在 `Pro` 套餐下相关 Issue (#71542) 下的讨论，是当前社区最核心的计费与功能冲突Bug。
    -   **社区反应**：47 条评论，16 个赞。用户普遍感到困惑和不满，认为这是严重的计费逻辑或模型路由错误。

2.  **[BUG] GitHub Connector 链接成功但无法访问任何仓库内容** [#71542](https://github.com/anthropics/claude-code/issues/71542)
    -   **重要性**：⭐⭐⭐⭐⭐
    -   **摘要**：一个持续一个月的回归性问题。用户能够成功授权并连接 GitHub，但 Claude Code 无法读取任何仓库（包括公开和私有）的内容，导致核心的代码理解和分析功能完全失效。
    -   **社区反应**：43 条评论，37 个赞。开发者普遍将这个问题视为使用体验中的“硬伤”，严重影响了日常开发工作。

3.  **[Bug] Fable 5 在交互式模型选择器中仍然需要使用额度（针对使用 setup-token 的 Max 账户）** [#79597](https://github.com/anthropics/claude-code/issues/79597)
    -   **重要性**：⭐⭐⭐⭐
    -   **摘要**：与 #79337 问题类似，但此问题更具体地发生在使用 `CLAUDE_CODE_OAUTH_TOKEN` 进行认证的自动化和无头（headless）环境中，进一步暴露了认证和计费逻辑的复杂性。
    -   **社区反应**：8 条评论，9 个赞。表明此Bug不仅影响交互式用户，也影响使用 CI/CD 流程的开发者。

4.  **[ENHANCEMENT] 提供禁用 Cowork 后台服务的方法（Windows）** [#57371](https://github.com/anthropics/claude-code/issues/57371)
    -   **重要性**：⭐⭐⭐⭐
    -   **摘要**：Windows 版本的 Claude Desktop 会强制安装并运行一个名为 `CoworkVMService` 的服务，对于不需要此功能的用户来说是一个资源浪费和潜在的性能干扰。
    -   **社区反应**：15 条评论，39 个赞。这是目前社区点赞数最高的功能请求，反映了用户对软件自主控制权的强烈诉求。

5.  **[BUG] 全屏渲染器破坏 macOS 系统剪贴板** [#72455](https://github.com/anthropics/claude-code/issues/72455)
    -   **重要性**：⭐⭐⭐⭐
    -   **摘要**：在 macOS 的 Terminal.app 中，当 Claude Code 进入全屏模式后，复制粘贴功能会在系统中所有应用程序中失效，问题严重性为“系统级”。
    -   **社区反应**：5 条评论，5 个赞。用户反馈这是一个高风险、影响面大的 Bug，会严重中断工作流。

6.  **[BUG] 7月17日大规模计费事故：使用额度被错误扣费** [#81703](https://github.com/anthropics/claude-code/issues/81703)
    -   **重要性**：⭐⭐⭐⭐
    -   **摘要**：Anthropic 已承认的 7 月 17 日计费事故。订阅用户的用量被错误地计入有偿的使用额度，导致部分用户产生了高额账单（如 $704.71）。
    -   **社区反应**：5 条评论。用户要求官方对已发生的错误收费进行全额退还。

7.  **[BUG] Windows 桌面版打开内置浏览器页面导致应用崩溃** [#81275](https://github.com/anthropics/claude-code/issues/81275)
    -   **重要性**：⭐⭐⭐
    -   **摘要**：Claude Desktop for Windows 的 MSIX 版本，在打开内置的 `Browser pane` 功能时，GPU 进程会因一个确定性错误码（0x60C201E）而崩溃，导致整个应用闪退。
    -   **社区反应**：5 条评论。这是一个导致核心功能（Cowork Browser）完全不可用的严重Bug。

8.  **[BUG] 现有 Max 订阅用户被锁定，无法登录（Web、桌面、CLI）** [#70115](https://github.com/anthropics/claude-code/issues/70115)
    -   **重要性**：⭐⭐⭐
    -   **摘要**：一个长期存在的后端认证路由错误，导致部分现有 Max 用户无法登录，登录流程会错误地跳转到“创建账户”页面。
    -   **社区反应**：2 条评论。此问题链接了多个相似Issue (#36797, #39788等)，说明是一个未彻底解决的顽固性Bug。

9.  **[Feature Request] 为远程控制界面增加“接受、清空上下文并进入自动模式”的审批选项** [#81393](https://github.com/anthropics/claude-code/issues/81393)
    -   **重要性**：⭐⭐⭐
    -   **摘要**：用户希望在远程或非交互式控制中也能使用 `showClearContextOnPlanAccept` 功能，在执行计划前自动清空对话上下文以避免Token浪费和混淆。
    -   **社区反应**：1 条评论，是新提出的功能请求，代表了高级用户对工作流优化的需求。

10. **[Feature Request] 稳定项目身份识别与跨机器同步** [#81391](https://github.com/anthropics/claude-code/issues/81391)
    -   **重要性**：⭐⭐⭐
    -   **摘要**：Claude Code 的自动记忆（auto memory）功能基于项目路径的哈希值，导致在不同操作系统或路径布局下，无法跨机器正确同步同一项目的记忆。
    -   **社区反应**：1 条评论。这是一个面向专业开发者的高级需求，旨在解决多设备协作时的体验一致性问题。

### 重要 PR 进展

1.  **[fix] devcontainer: 当可选域名解析失败时，不中断防火墙设置** [#81673](https://github.com/anthropics/claude-code/pull/81673)
    -   **内容**：修复了 DevContainer 初始化脚本因 `statsig.anthropic.com` 等可选域名无法解析而导致整个网络规则配置失败的问题。
    -   **重要性**：提升容器化开发环境的稳定性和可靠性。

2.  **[fix] hookify: 使包导入不依赖安装目录名** [#81672](https://github.com/anthropics/claude-code/pull/81672)
    -   **内容**：修复了 `hookify` 插件因市场安装目录名不是预期的 `hookify` 而导致的导入失败问题。
    -   **重要性**：解决了插件市场的兼容性问题，使得第三方插件生态更加健壮。

3.  **[fix] plugins: 引用和路径修复** [#81670](https://github.com/anthropics/claude-code/pull/81670)
    -   **内容**：修复了两个导致 `hookify` 插件设置后失效的bug：1) 未引用的 `${CLAUDE_PLUGIN_ROOT}` 变量导致路径包含空格的 hook 失败；2) 清除了历史遗留的 `hookify` 配置示例。
    -   **重要性**：直接解决了插件用户在特定配置下的安装和使用问题。

4.  **[feature] Add web4-governance plugin** [#20448](https://github.com/anthropics/claude-code/pull/20448)
    -   **内容**：一个提议很久的新插件，旨在为 Claude Code 引入基于“信任张量”和“实体见证”的 AI 治理框架。
    -   **重要性**：虽然进展缓慢，但代表了社区在 AI 工具责任和可追溯性方面的探索方向。

5.  **[docs] fix security-guidance plugin entry in plugins/README.md** [#81576](https://github.com/anthropics/claude-code/pull/81576)
    -   **内容**：修正了 `plugins/README.md` 中对 `security-guidance` 插件功能的错误描述（错误地声称其拥有单个 Hook 和 9 个模式，实际有 3 个 Hook 和 25 个模式）。
    -   **重要性**：确保文档准确，帮助用户正确理解和使用插件。

### 功能需求趋势

1.  **Fable 5 模型计费与访问策略修复**：这是当前压倒性的需求。核心诉求是让 Fable 5 模型在 Max 套餐中按预期工作，解决 CLI、桌面端、setup-token认证等多种场景下的“支付了套餐费用却仍需付费”的计费故障。
2.  **跨平台与跨机器状态同步**：用户不仅关注单机的体验，开始要求跨机器的配置、自动记忆和项目身份同步。这表明 Claude Code 正被用于更正规、多环境的开发流程中。
3.  **用户体验精细化与重载**：从“禁用 Cowork 服务”、“优化画布/终端交互”到“在远程审批中加入清空上下文选项”，社区关注点正从“能用”转向“好用”，追求减少干扰、提升效率的精细化控制。
4.  **安全性与权限管理**：`Browser tool` 读取操作与 `launchPreviewAllowedOrigins` 设定不一致的问题，表明用户对安全边界和权限提示的“可预测性”有更高要求，希望信任配置能被严格执行。
5.  **工具加载性能优化**：`Workflow` 工具因体积大且默认加载而浪费 Token 的问题，反映出社区对 Token 开销和使用效率的敏感度，希望引入更智能的懒加载机制。

### 开发者关注点

1.  **计费和额度问题 (Billing/Usage Credits)**：这是目前开发者社区最大的痛点和关注焦点。Fable 5 的计费故障以及 7 月 17 日的计费事故，直接影响了用户对服务公平性和可靠性的信任。
2.  **用户体验中“反直觉”的 Bug**：如“全屏模式破坏系统剪贴板”、“窗口缩放到 1 行高会清空会话”、“`--bare` 模式丢失认证信息”等Bug，虽然不如计费问题严重，但因其违反直觉且破坏性强，也引发了大量讨论。
3.  **集成与扩展的可靠性**：GitHub Connector 的长时故障以及 Windows 桌面版浏览器崩溃问题，凸显了与外部系统和平台集成的稳定性是开发者极为看重的，任何回归性 Bug 都会立即遭到社区反馈的“暴击”。
4.  **对无头/自动化环境支持不足**：使用 `CLAUDE_CODE_OAUTH_TOKEN` 的用户面临 Fable 5 无法使用、`--bare` 模式认证丢失等多重问题，说明自动化场景下的测试和优化仍存在不足。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 | 2026-07-28

## 今日速览
过去24小时，OpenAI Codex 发布了两个 Rust 版 alpha 小版本（0.146.0-alpha.12/13），修复集中在 Windows 平台稳定性与 CI 工具链对齐。社区反馈中 **重置计数浪费**（#31606）与 **Codex Diff 崩溃**（#35058）仍是最高热度问题，Windows 端 GPU 进程崩溃（#34133、#35352）和 **子代理磁盘占用飙升**（#34061）成为开发者新痛点。PR 侧 `copyberry[bot]` 密集合入多项基础设施修复（Windows 非 TTY 中断、MCP 并发准备、OpenTelemetry 幂等关闭等），整体偏向稳健性加固。

---

## 版本发布
- **rust-v0.146.0-alpha.13**：`Codex CLI` Rust 组件 0.146.0-alpha.13 发布。
- **rust-v0.146.0-alpha.12**：`Codex CLI` Rust 组件 0.146.0-alpha.12 发布。

两个版本均为增量迭代，未附带详细变更日志，推测主要同步了底层依赖与锁文件更新。

---

## 社区热点 Issues（Top 10）

### 1. [OP] Reset failed, did not apply and 1 reset is wasted (#31606)
**评论: 52 | 👍: 61**  
Pro 用户反映点击重置按钮后未生效，但重置计数从 2 降为 1。社区大量 +1 并出现多平台复现报告。OpenAI 尚未回复。  
https://github.com/openai/codex/issues/31606

### 2. [OP] [Windows] Codex App crashes in CrBrowserMain when Browser Use opens a page (#32683)
**评论: 27 | 👍: 8**  
Windows 桌面端在调用浏览器时因 `chrome.dll+0x2e08f46` 触发 `0xC0000005` 访问冲突崩溃。影响 Pro 20x 用户。  
https://github.com/openai/codex/issues/32683

### 3. [OP] [Windows] Page.captureScreenshot crashes GPU process after Code Integrity Event 3033 rejects bundled vk_swiftshader.dll (#34133)
**评论: 24 | 👍: 0**  
Windows 10 下截图导致 GPU 进程崩溃，根源是代码完整性检查拒绝捆绑的 `vk_swiftshader.dll`。用户反馈应用变慢、卡死甚至无法重启。  
https://github.com/openai/codex/issues/34133

### 4. [OP] Codex Diff crashes with “Oops, an error has occurred” in VS Code on macOS (#35058)
**评论: 20 | 👍: 48**  
VS Code 扩展中 Codex Diff 标签页完全不可用，每次打开都报错。在 Apple Silicon 及最新 VS Code 1.128.0 上复现，社区呼声极高。  
https://github.com/openai/codex/issues/35058

### 5. [OP] Insane Codex Disk Usage from Subagents (#34061)
**评论: 14 | 👍: 1**  
Codex CLI 子代理在 macOS 上产生巨大磁盘占用（报告显示数百 GB），涉及会话缓存与临时文件，影响 Pro 用户。  
https://github.com/openai/codex/issues/34061

### 6. [OP] Codex Desktop exits when the embedded browser GPU process crashes and unsigned SwiftShader fallback is blocked (#35352)
**评论: 12 | 👍: 0**  
与 #34133 同源，桌面版在嵌入式浏览器 GPU 崩溃后因 `swiftshader` 回退被阻止而直接退出。Windows Pro 用户。  
https://github.com/openai/codex/issues/35352

### 7. [OP] Codex Desktop Windows+WSL resolves bundled plugin cache as invalid C:\mnt\c path (#24268)
**评论: 10 | 👍: 3**  
WSL 工作区中插件缓存路径被错误解析为 `C:\mnt\c`，导致插件加载失败。Windows Store 版本，持续开放逾两个月。  
https://github.com/openai/codex/issues/24268

### 8. [OP] Codex subagents drain full week quota overnight - usage counting broken (#35463)
**评论: 3 | 👍: 0**  
Pro 20x 用户发现子代理一夜之间耗尽整周配额，认为使用量计数存在 bug。CLI 0.145.0, gpt-5.6-sol。  
https://github.com/openai/codex/issues/35463

### 9. [OP] Incomplete residual fidelity across capture, model-visible, and durable state (#35528)
**评论: 4 | 👍: 2**  
指出工具输出被截断、合并或上下文压缩后，Codex 未携带“忠实残差”，导致 agent 后续步骤丢失必要上下文。设计缺陷级别。  
https://github.com/openai/codex/issues/35528

### 10. [OP] Code mode can report completion with live nested exec sessions but no model-visible handles (#35613)
**评论: 3 | 👍: 0**  
刚提交的新 issue：代码模式在仍有活跃嵌套执行会话时过早报告完成，且这些会话没有模型可见句柄，导致用户误以为任务结束。  
https://github.com/openai/codex/issues/35613

---

## 重要 PR 进展（Top 10）

### 1. Point crossterm patch to the OpenAI OSS fork (#35688)
**已合并**  
将 `crossterm` Cargo 补丁指向 OpenAI OSS 分支，同步更新锁文件和 allowlist。确保内部构建一致性。  
https://github.com/openai/codex/pull/35688

### 2. Load cloud-managed profiles for `codex sandbox` (#35685)
**已合并**  
支持在 sandbox 启动时从云端加载管理配置，通过 `--include-managed-config` 参数传递权限配置。  
https://github.com/openai/codex/pull/35685

### 3. Preserve paginated thread metadata across resumes (#35678)
**已合并**  
修复分页历史中线程预览、标题等元数据在恢复时丢失的问题，改为从 SQLite 保留的状态中还原。  
https://github.com/openai/codex/pull/35678

### 4. Prepare MCP and plugin recommendations concurrently (#35675)
**已合并**  
将 MCP 发现与端点插件推荐并行化，减少串行等待时间。  
https://github.com/openai/codex/pull/35675

### 5. Route curated plugins by authentication mode (#35671)
**已合并**  
根据当前认证模式（ChatGPT/remote/API）选择正确的策展插件列表，避免账户切换后插件错误。  
https://github.com/openai/codex/pull/35671

### 6. Raise the Windows exec yield floor to 10 seconds (#35670)
**已合并**  
将 Windows 上 `exec_command` 的初始让步时间下限从默认值提升至 10 秒，缓解高延迟场景下的过早超时。  
https://github.com/openai/codex/pull/35670

### 7. Terminate Windows non-TTY processes on interrupt (#35655)
**已合并**  
支持在 Windows 非 TTY 进程中通过 `write_stdin` 发送 Ctrl-C 中断，此前报告为“不支持”。  
https://github.com/openai/codex/pull/35655

### 8. Preserve TUI input when terminal focus returns (#35649)
**已合并**  
修复 TUI 界面在终端焦点回归时丢失键盘输入的问题，通过缓存启动时探测的调色板避免阻塞输入循环。  
https://github.com/openai/codex/pull/35649

### 9. Preserve thread metadata when rollout files are missing (#35644)
**已合并**  
当回滚文件丢失时跳过对应的线程，但仍将其元数据保留在状态数据库中，避免 UI 列表空白。  
https://github.com/openai/codex/pull/35644

### 10. Make OpenTelemetry provider shutdown idempotent (#35642)
**已合并**  
修复 OpenTelemetry 提供者在显式关闭后因重复调用 `Drop` 而 panic 的问题，增加幂等性保护。  
https://github.com/openai/codex/pull/35642

---

## 功能需求趋势

从本期 Issues 与 PR 中可提炼出以下社区最关心的功能方向：

1. **MCP (Model Context Protocol) 生态完善**  
   - 多个 issue 要求 OAuth 生命周期可靠、SSO 支持（#35006）
   - PR 中涉及 MCP 与插件推荐的并发准备、认证模式路由

2. **多 Agent 模型兼容性与版本管理**  
   - `gpt-5.6-luna` 被标记为 `MultiAgent V1` 导致 V2 拒绝（#35097）
   - `gpt-5.6-sol` 不可用（#34027）、模型选择器重复显示（#35493）

3. **上下文保真度与残差机制**  
   - #35528 系统性地提出多平面（捕获、模型可见、持久状态）下残差丢失问题
   - 社区期望工具输出截断后能保留“可恢复的摘要”

4. **移动端 / 远程语音交互**  
   - #35687 请求为 Codex Mobile Remote 增加语音对话能力，与远端 Mac 实时结合

5. **Windows 原生体验强化**  
   - 大量 Windows 崩溃、GPU、沙箱、输入延迟问题（#34133、#35352、#34450、#33732），社区要求优先修复

---

## 开发者关注点

- **Windows 平台稳定性成最大痛点**：GPU 进程崩溃（`vk_swiftshader.dll` 被拒）、非 TTY 中断不支持、沙箱挂起、输入延迟等问题密集出现。开发者呼吁 OpenAI 在 Windows 端投入更多测试资源。
- **配额与重置机制不透明**：重置按钮不生效且消耗计数（#31606）、子代理一夜耗光整周配额（#35463）、使用量计数疑似不准。Pro/Plus 用户对成本敏感，反馈激烈。
- **Codex Diff 功能不可用**：macOS 下 VS Code 中 100% 崩溃（#35058，48👍），严重阻碍代码审查流程。该功能属于核心工作流，修复优先级应调高。
- **多 Agent 配置混乱**：模型名重复、版本标识冲突、代理模式不兼容导致任务失败。开发者希望统一模型元数据与版本策略。
- **磁盘与性能泄漏**：子代理磁盘占用狂飙、`node_repl` 常驻进程不释放（#35582）、桌面端发热卡顿。长期运行场景下资源管理亟需优化。

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，各位开发者，大家好！欢迎查阅 **2026 年 7 月 28 日** 的 Gemini CLI 社区动态日报。在过去的 24 小时里，社区围绕 Agent 行为的可预测性、核心工具的稳定性以及新模型支持展开了积极讨论。其中，关于**子代理在达到最大轮次后谎报成功的 Bug** 和 **通用代理在特定情况下挂起**的问题是社区关注的焦点，同时，开发者对 **AST 感知工具** 和 **新模型 (gemini-3.5-flash) 的支持**也表现出浓厚兴趣。

---

### 版本发布

- **[v0.54.0-nightly.20260727.g3818efbbf](https://github.com/google-gemini/gemini-cli/compare/v0.54.0-nightly.20260726.g3818efbbf...v0.54.0-nightly.20260727.g3818efbbf)**：新的每日构建版本发布。此版本为常规夜间构建，社区尚未报告具体的新增功能或显著变化。

### 社区热点 Issues（过去 24 小时更新）

1. **[#22323] 子代理在达到 MAX_TURNS 后误报为 GOAL 成功**  
   一个严重的 Bug，`codebase_investigator` 子代理在达到最大轮次限制后，并未如实报告中断，而是错误地报告任务“成功”。这可能导致用户对任务状态产生严重误判。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/22323) | 评论: 12

2. **[#21409] 通用代理 (Generalist agent) 在执行任务时挂起**  
   这是社区反映强烈的老问题，当 `gemini-cli` 将任务委托给通用代理时，会永久挂起（即使是简单的创建文件夹操作）。用户需要通过指令禁止使用子代理才能绕过。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/21409) | 评论: 8 | 👍: 8

3. **[#11799] `gemini-cli` 似乎忽略了 `GEMINI.md` 上下文文件**  
   一个影响开发工作流的关键问题：模型似乎完全忽略本地的 `GEMINI.md` 上下文文件，即使 `/memory show` 显示文件已加载。用户必须手动使用 `@GEMINI.md` 指令才能生效。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/11799) | 评论: 5 | 👍: 4

4. **[#25166] Shell 命令执行完成后卡在“等待输入”状态**  
   一个影响日常使用的高频 Bug：在 Gemini 执行完简单的 CLI 命令后，界面会卡住，显示仍在“等待用户输入”，阻碍后续操作。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/25166) | 评论: 4 | 👍: 3

5. **[#22745] 评估 AST 感知文件读取、搜索和代码映射的影响**  
   一个重要的功能探索议题，旨在通过引入 AST 感知工具，提高代码定位和读取的精度，减少 Token 消耗和与模型交互的次数，从而提升 Agent 的性能。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/22745) | 评论: 7

6. **[#21968] Gemini 未能充分利用自定义技能 (Skills) 和子代理**  
   社区用户反馈，即使已经明确定义了自定义技能（如 `gradle`、`git`），模型在执行相关任务时仍倾向于自行处理，而不是主动调用这些工具，导致效率低下。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/21968) | 评论: 6

7. **[#24353] 稳健的组件级评估**  
   这是一项内部计划，旨在推进针对不同组件的自动化评估体系，以确保 Gemini CLI 在各种模型和场景下的行为可靠性。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/24353) | 评论: 7

8. **[#22232] 增强浏览器 Agent（browser_agent）的韧性**  
   用户提出，`browser_agent` 在处理浏览器配置文件锁定时采取“快速失败”策略，导致任务中断。建议增加会话接管和锁恢复机制，以提高其在复杂场景下的稳定性。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/22232) | 评论: 4

9. **[#24246] Gemini CLI 在超过 128 个工具时遇到 400 错误**  
   当环境中启用的工具超过一定数量时，Gemini CLI 会因请求过大而报错。社区希望 Agent 能更智能地根据当前场景筛选可用工具，而不是一次性发送所有工具声明。  
   [链接](https://github.com/google-gemini/gemini-cli/issues/24246) | 评论: 3

10. **[#21983] 浏览器子代理在 Wayland 环境中失败**  
    报告显示，在 Wayland 显示服务器协议下，浏览器子代理无法正常工作，限制了 Linux 用户的使用体验。  
    [链接](https://github.com/google-gemini/gemini-cli/issues/21983) | 评论: 4

### 重要 PR 进展（过去 24 小时更新）

1. **[#28551] 修复 macOS 沙盒模式下因缺少 Seatbelt 配置文件而崩溃**  
   修复了在 macOS 上使用 `-s` 沙盒模式启动时，因未找到静态的 `.sb` 配置文件而立即崩溃的问题。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28551) | 状态: OPEN

2. **[#28481] 修复 MCP OAuth 令牌刷新失败问题**  
   修复了通过 OAuth 发现 + 动态客户端注册配置的 MCP 服务器，在刷新令牌时因本地校验失败而清空凭据，导致用户每 60 分钟需重新认证的问题。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28481) | 状态: OPEN | 优先级: P1

3. **[#28485] 为所有用户添加 gemini-3.5-flash 模型选择器**  
   解决了 v0.51.0 版本用户无法通过模型选择器使用 `gemini-3.5-flash` 或 `gemini-3.6-flash` 的问题，确保新模型能够被正确发现和选择。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28485) | 状态: OPEN | 优先级: P2

4. **[#28546] 修复使用 GEMINI_API_KEY 认证时，未清理 Authorization 头导致冲突**  
   解决了当优先使用 `GEMINI_API_KEY` 认证时，仍会发送遗留在环境变量中的 `Authorization` 头的问题，防止向 Google API 发送错误请求。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28546) | 状态: OPEN | 优先级: P1

5. **[#28364] 修复模型配置深度合并问题**  
   修复了用户自定义的模型配置无法正确覆盖默认配置的问题。之前使用浅拷贝导致部分配置被覆盖，PR 实现了深度合并以确保用户设置生效。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28364) | 状态: CLOSED | 优先级: P2

6. **[#28363] 防止 ShellExecutionService 中的 AbortSignal 监听器泄露**  
   修复了 Shell 执行服务中的内存泄漏问题。确保进程结束后，事件监听器被正确移除，避免长时间会话导致内存占用不断增长。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28363) | 状态: CLOSED | 优先级: P2

7. **[#28369] 新增本地评估报告命令和开发者文档**  
   为开发者新增了 `npm run eval:report` 命令，可以汇总本地测试结果并生成报告，方便进行 Behavioral Eval 的社区协作。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28369) | 状态: CLOSED

8. **[#28531] 修复 Windows 上 CRLF 行结束符导致侧边差异视图无效**  
   修复了在 Windows 系统中，由于 `a2a-server` 包生成的内容使用 CRLF 行结束符，导致 Gemini Code Assist 的差异视图无法高亮代码变更的问题。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28531) | 状态: CLOSED

9. **[#28549] 修复 MCP 计划模式（Plan Mode）安全状态**  
   修复了 MCP 工具在计划模式下的权限问题，明确告知用户工具只读状态是由 MCP 服务器声明的，而不是由 Gemini CLI 验证的，提升了安全性。  
   [链接](https://github.com/google-gemini/gemini-cli/pull/28549) | 状态: OPEN

10. **[#28447] 添加 Windows PowerShell 故障排除文档**  
    一个非常实用的文档更新，针对 Windows 用户安装后无法执行 `gemini` 命令的问题，提供了清晰的安装指导和故障排查步骤。  
    [链接](https://github.com/google-gemini/gemini-cli/pull/28447) | 状态: OPEN | 优先级: P2

### 功能需求趋势

从上榜的 Issues 和 PRs 中，可以提炼出社区目前最关注的几个功能方向：

1. **Agent 行为的智能化与可预测性**：社区不再满足于 Agent 能“做什么”，而是关注它“何时做”以及“如何报告”。这体现在对 **子代理误报成功、不主动使用自定义技能** 的批评上。开发者希望 Agent 能更智能地判断何时调用子模块，并能诚实、清晰地报告自己的执行状态。
2. **核心稳定性与可靠性**：**通用代理挂起、Shell 执行后卡死** 这类问题直接影响了工具的最基础可用性。开发者首要关注的是 CLI 工具能稳定地完成基本任务，而不是因各种意外状态中断工作流。
3. **对新模型和配置的及时支持**：**gemini-3.5-flash 模型选择器问题** 以及 **MCP OAuth 认证问题** 表明，开发者对新模型和新协议的支持非常敏感。他们希望 Gemini CLI 能够快速跟进并稳定支持上游技术栈的更新。
4. **安全与权限的清晰化**：**“计划模式”的权限边界** 和 **API Key 认证冲突** 等问题表明，随着 Agent 能力增强，开发者对安全边界的意识也在提升。他们需要工具能明确、诚实地展示其行为边界，尤其是在访问凭证和执行高危操作时。

### 开发者关注点

- **核心痛点**：“Agent 不可靠”是开发者反馈的最强音。无论是子代理的“伪成功”汇报，还是通用代理的“永久挂起”，都极大地损害了用户对自动化的信任。这或许是当前阶段最需要优先解决的信任危机。
- **高频需求**：
    - **修复 Agent 挂起和卡死**：这是影响体验的最大障碍，用户期待一个“用完即走”的流畅交互。
    - **提升 Agent 的“工具利用”能力**：用户花费时间配置了技能和子代理，但模型不主动调用是一种资源浪费。社区期待 Agent 能变得更“聪明”，能根据上下文主动调用合适的工具。
    - **持续优化上下文处理**：`GEMINI.md` 被忽略的问题再次提醒开发者，上下文窗口的利用是否精准、可靠，直接决定了 AI 编程助手的上限。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-07-28

---

## 📌 今日速览

- **新版 v1.0.76-0 发布**：MCP 工具加载速度提升，新增 Autopilot 模式持续保持选项；同时修复了会话启动时的早期警告。
- **社区讨论活跃**：Plan 模式下 shell 命令被阻止的回归问题 (#4188) 引发高度关注；Windows Terminal 下交互模式显示异常 (#4159、#4263) 成为多用户痛点。
- **功能需求聚焦**：持久化 Autopilot 模式、模型自动切换、ACP 协议对 token 用量/计费属性的暴露，以及非 Git 仓库的 Rewind 支持呼声最高。

---

## 🚀 版本发布

### v1.0.76-0 (2026-07-27)

**改进**
- MCP 工具从 definition-scoped 快照中加载更快，支持进程级和每服务器缓存可关闭。
- Autopilot 模式在 `task_complete` 后默认保持选中状态；若需每次任务后返回交互模式，可设置 `stayInAutopilot: false`。

**修复**
- 恢复了在 `un` 时的早期警告（原文截断，推测为 `unhandled error` 或类似场景）。

➡️ 详情：https://github.com/github/copilot-cli/releases

---

## 🔥 社区热点 Issues（Top 10）

### 1. Plan 模式下 shell 命令被阻止 — 回归问题
- **#4188** [OPEN] [area:permissions, area:tools] Regression on plan-mode  
- **评论数**: 6 | 👍: 3  
- 用户在最新版本中发现 Plan 模式禁止执行 shell 命令（如 `gh` CLI），导致计划无法丰富上下文。社区视为严重回归。  
- https://github.com/github/copilot-cli/issues/4188

### 2. 模型自动切换（规划 vs 执行）
- **#2792** [CLOSED] [area:agents, area:models] Automatic switching between model for planning and execution  
- **评论数**: 5 | 👍: 16  
- 用户期望为规划步骤和执行步骤分别配置不同模型，以提升效率。该功能获得大量点赞，目前虽已关闭但仍是社区高频请求。  
- https://github.com/github/copilot-cli/issues/2792

### 3. 子进程僵尸累积（Linux）
- **#4163** [CLOSED] [area:platform-linux, area:tools] copilot CLI 1.0.71 does not reap child processes  
- **评论数**: 5 | 👍: 3  
- Copilot 进程未正确回收子进程，导致僵尸进程持续累积（约 2/min），影响系统稳定性。已修复但值得关注。  
- https://github.com/github/copilot-cli/issues/4163

### 4. CAPI 5MB 请求体限制导致会话中断
- **#4183** [CLOSED] [area:context-memory, area:models] Auto-compaction does not prevent CAPI 5 MB failure  
- **评论数**: 4 | 👍: 10  
- 即使上下文 Token 未超限，序列化后的 CAPI Responses 请求体达到 5MB 限制时会话永久卡死。自动压缩未能避免。  
- https://github.com/github/copilot-cli/issues/4183

### 5. Rewind 不支持非 Git 仓库
- **#1381** [OPEN] [area:sessions] "Rewind is not available because you're not in a git repository."  
- **评论数**: 3 | 👍: 9  
- 用户使用其他 VCS（如 `jj`）时无法使用 Rewind 功能，VS Code 中的 Copilot 无此限制。社区要求增加非 Git 支持。  
- https://github.com/github/copilot-cli/issues/1381

### 6. ACP 模式缺乏 token 与 credits 用量通知
- **#4233** [OPEN] [area:non-interactive] [ACP] Emit `usage_update` in `--acp` mode for parity with interactive statusline  
- **评论数**: 2 | 👍: 2  
- ACP 客户端（如 Zed）无法显示上下文窗口或 AI 信用额度指示器，因为 CLI 未在 ACP 会话中发送 `usage_update` 事件。  
- https://github.com/github/copilot-cli/issues/4233

### 7. Autopilot 模式下 `task_complete` 工具再次不可用
- **#4161** [OPEN] [area:agents, area:tools] task_complete tool unavailable after switching back to autopilot mode  
- **评论数**: 2 | 👍: 3  
- 该问题曾在 v1.0.4 修复，但最新版本又出现回归，用户需手动重新进入 Autopilot 才能调用 `task_complete`。  
- https://github.com/github/copilot-cli/issues/4161

### 8. Windows Terminal 中交互模式内容消失
- **#4159** [OPEN] [area:platform-windows, area:terminal-rendering] Copilot CLI interactive mode turns blank after submitting a prompt  
- **评论数**: 1 | 👍: 3  
- 交互模式下首次提交后 UI 变空白，非交互模式正常。多用户报告类似问题（#4263）。  
- https://github.com/github/copilot-cli/issues/4159

### 9. macOS 钥匙串重复弹窗（签名分区冲突）
- **#4273** [OPEN] [triage] macOS: keychain prompts on every launch when GitHub-signed and Microsoft-signed copilot binaries share login-keychain items  
- **评论数**: 0  
- 同一 CLI 以两个不同开发者身份签名时，钥匙串项的分区列表冲突，导致每次启动都弹窗要求权限。  
- https://github.com/github/copilot-cli/issues/4273

### 10. `glob` 工具对多级路径模式假阴性
- **#4271** [OPEN] [triage] `glob` tool false-negatives on any multi-segment pattern unless prefixed with `**/`  
- **评论数**: 0  
- 任何包含路径分隔符的模式（如 `2026/07/*.md`）在未加 `**/` 前缀时都返回“无匹配文件”，严重影响文件搜索。  
- https://github.com/github/copilot-cli/issues/4271

---

## 📦 重要 PR 进展（Top 10）

| PR | 内容 | 状态 | 链接 |
|----|------|------|------|
| #1609 | 更新 PAT 权限添加说明，明确 `Copilot Requests` 位于 Account 选项卡下 | OPEN | https://github.com/github/copilot-cli/pull/1609 |
| #1598 | 为 `install.sh` 添加 `trap` 清理临时目录，防止意外退出时泄漏 | OPEN | https://github.com/github/copilot-cli/pull/1598 |
| #1333 | 修正 README 中的语法和 Markdown 格式问题（添加冠词、移除多余空行） | OPEN | https://github.com/github/copilot-cli/pull/1333 |
| #1116 | 修正文档误导：0x 模型不会消耗配额，需在 README 中明确说明 | OPEN | https://github.com/github/copilot-cli/pull/1116 |
| #988 | 修正 README 中 `brew` 安装命令缺少前缀的问题 (`brew install copilot-cli` → `brew install github/copilot-cli/copilot-cli`) | OPEN | https://github.com/github/copilot-cli/pull/988 |
| #4030 | 新增 GitHub Actions 工作流，自动构建和部署 Jekyll 网站到 GitHub Pages | OPEN | https://github.com/github/copilot-cli/pull/4030 |
| #3928 | 添加 `.gitignore` 和设置配置文件 | OPEN | https://github.com/github/copilot-cli/pull/3928 |
| #2800 | 添加初始 devcontainer 配置 (0352859567) | OPEN | https://github.com/github/copilot-cli/pull/2800 |
| #3873 | 添加初始控制台问候日志 (1000) | OPEN | https://github.com/github/copilot-cli/pull/3873 |
| #4057 | 安装相关更新 | OPEN | https://github.com/github/copilot-cli/pull/4057 |

> 注意：部分 PR 内容较简略或可能存在测试性质，但仍被社区维护者保留。

---

## 🔍 功能需求趋势

1. **Autopilot 模式持久化**：用户希望能够在命令行启动时即指定持久的 Autopilot 模式，而不仅限于初始模式 (#3977、#4161)。
2. **模型自适应调度**：规划步骤与执行步骤使用不同模型（如高推理模型做计划，高速模型执行）以优化成本和速度 (#2792)。
3. **ACP 协议完善**：要求在 ACP 模式下暴露 token 用量、AI 信用额度、上下文层级等元数据，以支持第三方 IDE 集成 (#4233、#4174、#4275)。
4. **上下文窗口管理**：自动压缩未能解决 CAPI 5MB 上限，社区期待更智能的上下文裁剪或分页机制 (#4183)。
5. **非 Git VCS 支持 Rewind**：用户希望 Rewind 功能不依赖于 Git，以兼容 `jj`、Mercurial 等版本控制工具 (#1381)。
6. **插件与钩子机制**：期待完善 `sessionStart` 等钩子的触发时机与文档 (#1730)，以及 symlink 配置文件的正式支持 (#3264)。
7. **Windows Terminal 兼容性**：多个 Windows 用户报告渲染问题，期望在交互模式下保持内容稳定 (#4159、#4263)。
8. **安全与认证**：macOS 钥匙串弹窗问题表明需要统一签名策略或使用更安全的凭证存储方式 (#4273)。

---

## 💡 开发者关注点（痛点 & 高频求助）

- **Plan 模式回归影响工作流**：`gh` 等 shell 命令被阻止后，开发者无法在计划阶段动态读取 Issue 或创建内容，严重打断自动化流程 (#4188)。
- **Windows 场景频繁报错**：交互模式变白 / 内容消失、键盘缓存的左右箭头不停止、tmux 下剪贴板失效等问题，说明 Windows 端的测试覆盖仍需加强 (#4159、#4263、#4191、#4274)。
- **模型选择被组织策略屏蔽**：新模型显示“被组织策略禁用”但管理员设置中无法找到开启选项，表明组织策略管理界面存在漏洞 (#4272)。
- **glob 工具不直观**：路径通配符必须加 `**/` 前缀才能生效，与开发者常见用法不符，降低文件搜索效率 (#4271)。
- **BYOK/custom provider 启动提示被忽略**：通过 `-i` 传入的启动提示在自定义 provider 下不被自动提交，导致交互体验不一致 (#4258)。
- **子代理计费不可见**：OTel span 中缺失 `github.copilot.nano_aiu` 和 `github.copilot.cost` 属性，导致外部成本核算失真 (#4224)。
- **MCP 配置灵活性不足**：用户期望通过文档明确 symlink、hook 配置文件的跨平台行为，以便更好地共享设置 (#3264)。

---

📅 **数据来源**：github.com/github/copilot-cli | 更新截止 2026-07-27 23:59 UTC  
📝 **日报生成时间**：2026-07-28

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 — 2026-07-28

## 今日速览

- 社区主要聚焦于 **VS Code 扩展的稳定性**，两个新提交的 bug 报告（#2563、#2317）分别指向审批弹窗不渲染和文件路径不可点击等交互问题。  
- **Windows 用户编码兼容性**成为修复热点，两个 PR（#2561、#2560）同时针对非 UTF-8 环境下的启动崩溃问题提出解决方案。  
- 后处理 Hook 任务被 GC 意外回收（#2564）以及 prompt cache 控制项（#2562）等底层修复也在同步推进中。

---

## 版本发布

过去24小时内无新版本发布。

---

## 社区热点 Issues

### #1070 [已关闭] 登录失败：无法连接到 auth.kimi.com:443 网络不可达  
- **作者**：notedit ｜ **创建**：2026-02-09 ｜ **更新**：2026-07-27 ｜ **评论**：8  
- **重要性**：该问题虽已关闭，但近期再次被更新，说明用户可能仍遇到类似网络环境下的登录异常。  
- **社区反应**：8 条评论，无点赞，讨论可能集中在特定网络策略或代理配置上。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/issues/1070

### #2317 [开放] [BUG] [VS Code 扩展] Plan 模式下的文件路径在聊天 Webview 中不可点击  
- **作者**：vlad-at-work ｜ **创建**：2026-05-17 ｜ **更新**：2026-07-27 ｜ **评论**：3  
- **重要性**：影响 VS Code 扩展日常使用体验，路径不可点击降低了编辑效率。  
- **社区反应**：讨论中，等待修复。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/issues/2317

### #2564 [开放] fix(hooks): PostToolUse / PostToolUseFailure 任务在完成前被 GC 回收  
- **作者**：belenov-maker ｜ **创建**：2026-07-27 ｜ **更新**：2026-07-27 ｜ **评论**：0  
- **重要性**：**非确定性丢 Hook** 问题，影响自定义工具后处理逻辑的可靠性，可能造成任务异常中断。  
- **社区反应**：新提 bug，暂无回复，但根因分析已给出（`config.toml` 中的 hooks 被 GC 收割）。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/issues/2564

### #2563 [开放] [BUG] VS Code 扩展：审批提示（ExitPlanMode / 工具权限）间歇性不渲染，导致无限等待或 600 秒超时  
- **作者**：edpa2019 ｜ **创建**：2026-07-27 ｜ **更新**：2026-07-27 ｜ **评论**：0  
- **重要性**：**高频阻塞问题**，用户使用 kimi-k3 模型时扩展 UI 完全卡死，严重干扰工作流。  
- **社区反应**：刚提交，尚无回复，但问题描述详尽（macOS、扩展 0.6.4）。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/issues/2563

---

## 重要 PR 进展

### #2539 [开放] fix(mcp): 为 Moonshot API 规范化工具名称  
- **作者**：lihailong00 ｜ **创建**：2026-07-23 ｜ **更新**：2026-07-27  
- **内容**：为 MCP 工具生成兼容 Moonshot 的稳定别名，保留原始名称用于路由；修复 schema 中 object 类型缺失和 `anyOf`/`required` 形状问题。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/pull/2539

### #2562 [开放] fix(llm): 允许禁用 prompt cache key  
- **作者**：lihailong00 ｜ **创建**：2026-07-27 ｜ **更新**：2026-07-27  
- **内容**：在 `kimi` provider 配置中增加 `prompt_cache_key` 布尔开关，设为 `false` 时可跳过 session 派生的缓存键，兼容非托管 Kimi 提供者场景。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/pull/2562

### #2561 [开放] 修复启动时 stdio 使用非 UTF-8 编码时的 UnicodeEncodeError  
- **作者**：LHMQ878 ｜ **创建**：2026-07-27 ｜ **更新**：2026-07-27  
- **内容**：解决 Windows Git Bash 下 `'gbk' codec can't encode character '▐'` 崩溃（源自欢迎横幅 logo 中的特殊字符）。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/pull/2561

### #2560 [开放] 修复 web 模式 banner 在 stdout 非 UTF-8 时的 UnicodeEncodeError  
- **作者**：LHMQ878 ｜ **创建**：2026-07-27 ｜ **更新**：2026-07-27  
- **内容**：解决 `kimi web` 在 Windows 中文 locale（codepage 936/GBK）下输出重定向时，因箭头符号 `➜` 导致的崩溃。  
- **链接**：https://github.com/MoonshotAI/kimi-cli/pull/2560

---

## 功能需求趋势

从最近一日的社区讨论和修复来看，最突出的需求方向包括：

- **VS Code 扩展稳定性**：拖放路径点击、审批弹窗渲染等交互问题是用户高频反馈的痛点，表明扩展 UI 层的可靠性仍是当前短板。
- **Windows 字符编码兼容性**：连续两个 PR 修复非 UTF-8 编码下的启动崩溃，说明 Windows 多语言环境支持是重要的兼容性需求。
- **MCP 工具集成标准化**：PR #2539 对 Moonshot API 的 MCP 工具名称规范化，显示社区正推动更稳定的第三方工具链对接。
- **Hooks 生命周期可靠性**：Issue #2564 揭示的 GC 回收问题暴露出后台任务管理的潜在缺陷，用户期望更确定的后处理执行保证。
- **配置灵活性**：PR #2562 引入 `prompt_cache_key` 开关，反映用户希望对缓存行为有更细粒度的控制，而非完全依赖平台默认。

---

## 开发者关注点

- **Windows 环境兼容性**：多位 Windows 用户反馈 `kimi` 命令在 Git Bash 或中文 locale 下直接崩溃，编码 Bug 导致无法正常使用 CLI 或 Web 模式。建议非 Windows 用户提前关注修复进度。
- **IDE 集成阻塞**：VS Code 扩展中的审批弹窗不渲染、Plan 模式路径无法点击等问题已持续数月（Issue #2317 创建于 2026-05-17），严重影响日常工作效率。扩展团队需优先处理 UI 事件循环或通信机制的缺陷。
- **非确定性行为**：Hook 任务偶尔被 GC 丢弃（#2564）让高级用户难以依赖 `PostToolUse` 做自动化工作流，期望添加更健壮的引用保护或同步机制。
- **网络与认证稳定性**：旧 Issue #1070 虽已关闭，但登录时 SSL 连接失败的可能场景（如代理、CN 网络限制）仍需文档指引或重试策略优化。

---

*数据来源：MoonshotAI/kimi-cli GitHub 仓库，统计窗口为 2026-07-27 UTC 时间。*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报（2026-07-28）

## 今日速览

今日发布了两个小版本修复：v1.18.7 针对 macOS 全屏标题栏、命令面板闪回等桌面 Bug 进行修补；v1.18.6 修复了分支缓存冲突和旧版 MCP 兼容性问题。社区 Issue 活跃度较高，主要集中在模型无限循环、子 agent 故障恢复、桌面主题修改受限、上下文窗口硬编码等痛点。PR 方面，多个修复和功能贡献正在推进，包括编辑工具输出优化、Watcher 生命周期管理、文件链接点击支持等。

## 版本发布

### 🐞 v1.18.7
- **Desktop Bugfixes**：
  - 移除 macOS 全屏模式下多余的标题栏内边距。
  - 修复当影子命令被移除时，命令面板条目错误重现的问题。
  - 项目选择器下拉列表在项目过多时支持滚动（感谢 @david1gp）。
- 社区贡献者：1 位。

### 🐞 v1.18.6
- **Core Bugfixes**：
  - 修复分支特定仓库缓存：刷新一个引用不再错误移动另一个分支的 checkout。
- **Desktop 改进**：
  - 提升与较新版客户端 API 在目录、项目、会话、终端流程的兼容性。
- **Desktop Bugfixes**：
  - 修复旧版 MCP（模型上下文协议）的兼容性问题。

## 社区热点 Issues（10 条）

1. **#28596 – [BUG] 反复工具调用导致无限循环**  
   [链接](https://github.com/anomalyco/opencode/issues/28596)  
   模型有时会无限重复相同的工具/执行调用（参数完全相同），必须手动中断。社区提议 agent 应具备预防机制。该 Issue 创建于 5 月但至今未解决，评论 5 条，是长期存在的核心稳定性问题。

2. **#38384 – 启动时持续报错“Missing required parameter”**  
   [链接](https://github.com/anomalyco/opencode/issues/38384)  
   每次启动 opencode 都会在 TUI 中弹出该错误，虽不影响使用但令人困惑。评论 2 条，用户期望官方明确原因或修复。

3. **#39196 – 前台子 agent 失败后无 task_id，父 agent 无法恢复**  
   [链接](https://github.com/anomalyco/opencode/issues/39196)  
   当子 agent 失败或被取消时，task 工具仅返回错误字符串而无 task_id，导致父 agent 无法恢复会话，子 agent 的部分工作丢失。这是任务编排的关键缺陷，社区高度关注。

4. **#39210 – [needs:compliance] BUG：无任何响应**  
   [链接](https://github.com/anomalyco/opencode/issues/39210)  
   用户输入 prompt 后无任何反馈，换 chat 和换免费 API 均无效。疑似合规性或 API 兼容问题，目前无评论。

5. **#39205 – [needs:compliance] 桌面端无法修改主题**  
   [链接](https://github.com/anomalyco/opencode/issues/39205)  
   设置中“Theme”选项只能修改一次，再次修改需要重新打开设置页面。界面交互 Bug，影响用户体验。

6. **#39204 – deepseek-v4-flash-free 在每次工具调用后停止 agent 循环**  
   [链接](https://github.com/anomalyco/opencode/issues/39204)  
   使用该免费模型时，Build agent 频繁在单次工具调用后停止，即使模型表示仍要继续。用户需手动输入 “continue”。社区认为需要检测模型输出并自动重试。

7. **#39202 – [FEATURE] 翻译修改：agent 应译为“智能体”**  
   [链接](https://github.com/anomalyco/opencode/issues/39202)  
   中文文档中将 “agent” 译为“代理”不符合 AI 领域习惯，建议改为“智能体”。社区反馈唯一，属于文档本地化建议。

8. **#35863 – [BUG] 上下文窗口硬编码为 200k，未从模型元数据动态获取**  
   [链接](https://github.com/anomalyco/opencode/issues/35863)  
   OpenCode 的上下文窗口处理依赖硬编码值，导致自动压缩、溢出检测过早触发。点赞 1 个，评论 1 条，是影响多模型兼容性的架构问题。

9. **#39212 – [FEATURE] 明确 task_id 的来源和恢复指南**  
   [链接](https://github.com/anomalyco/opencode/issues/39212)  
   Task 工具描述称“输出中包含可复用的 task_id”，但实际输出从未显示该字段。要求文档澄清。与 #39196 相关。

10. **#39208 – [bug, 2.0] glob/grep 工具无默认超时，一次调用挂起 21 分钟**  
    [链接](https://github.com/anomalyco/opencode/issues/39208)  
    V2 的 glob 和 grep 工具缺少墙钟超时，导致低匹配模式在大搜索根下无限运行。社区建议添加默认超时和用户可配置选项。

## 重要 PR 进展（10 条）

1. **#39211 – feat(core): 改进编辑工具输出**  
   [链接](https://github.com/anomalyco/opencode/pull/39211)  
   将合成的新旧差异预览替换为简洁的替换计数输出；报告实际匹配数；在无匹配时包含目标路径；返回具体的文件缺失/目录错误。已合并，将提升编辑操作的可用性。

2. **#39213 – docs: 明确 task_id 来源和子 agent 恢复指南**  
   [链接](https://github.com/anomalyco/opencode/pull/39213)  
   针对 #39212 的文档改进，无代码变更。说明 task_id 仅在成功时返回，并给出恢复建议。已由社区贡献者提交。

3. **#33453 – fix(provider): 默认自定义模型支持图像输入**  
   [链接](https://github.com/anomalyco/opencode/pull/33453)  
   修复自定义模型默认只支持文本的问题，新模型默认同时支持文本和图像输入，保留 V2 能力模态。自动清理 PR，已合并。

4. **#39203 – refactor(core): 使用 RcMap 管理 Watcher 生命周期**  
   [链接](https://github.com/anomalyco/opencode/pull/39203)  
   使 Watcher 获取操作可中断（中断安全），避免原生 Parcel 订阅占用导致消费者卡死 10 秒。已合并，提升稳定性。

5. **#39209 – fix(desktop): 本地运行时使用本地数据库**  
   [链接](https://github.com/anomalyco/opencode/pull/39209)  
   停止在未打包的桌面运行中禁用 channel database，让桌面开发时使用本地数据库。有助于开发调试。

6. **#39206 – fix(desktop): 使 file:// 聊天链接可点击**  
   [链接](https://github.com/anomalyco/opencode/pull/39206)  
   修复桌面应用中 `file://` 链接和绝对路径看似可点击但无反应的问题。原因：DOMPurify 清除和缺少事件处理。社区贡献。

7. **#29831 – fix(core): 解决 Windows 后台进程挂起问题**  
   [链接](https://github.com/anomalyco/opencode/pull/29831)  
   修复 shell 命令启动后台进程后挂起：通过监听进程退出而非仅 close 事件，并在退出后继续读取输出 500ms 静默后结束。长期未合并但持续活跃。

8. **#38534 – feat(tui): 添加 toast 挂载事件**  
   [链接](https://github.com/anomalyco/opencode/pull/38534)  
   新增 `tui.toast.mount` 生命周期事件供服务端插件使用，TUI 在 toast 挂载后发送 POST 请求。新功能 PR，社区贡献。

9. **#37625 – fix(provider): 标准化 Kimi 工具架构以兼容 MCP**  
   [链接](https://github.com/anomalyco/opencode/pull/37625)  
   通过模型无关兼容层投影 Kimi 工具架构，防止自定义或 MCP 工具拒绝整个提示。保持显式类型权威，回退兼容。

10. **#38060 – fix(opencode): 排除被拒绝的 MCP 工具**  
    [链接](https://github.com/anomalyco/opencode/pull/38060)  
    确保全局 `tools` 配置中 `{ "mymcp_*": false }` 能正确禁用整个 MCP 服务而非忽略。修复配置语义 Bug。

## 功能需求趋势

从近期的 Issue 和 PR 可看出社区最关注的方向：

- **Agent 稳定性与自我恢复**：多处反馈模型进入无限循环、子 agent 失败无法恢复、上下文窗口硬编码等问题，社区渴望更强的容错和自愈能力。
- **桌面端交互体验**：主题修改受限、文件链接不可点击、全屏标题栏问题等，表明桌面应用细节打磨需求迫切。
- **模型兼容性与动态配置**：硬编码上下文窗口、自定义模型默认能力（如图像输入）、MCP 工具冲突等，社区希望 OpenCode 能更智能地适配不同 Provider。
- **工具箱超时与安全**：glob/grep 无超时导致会话挂起，社区呼吁加入默认超时机制。
- **文档本地化**：术语翻译（如 agent→智能体）和 task_id 使用说明等，体现非英语用户对文档质量的关注。

## 开发者关注点

- **模型循环与中断问题**：几乎所有用户在使用 agent 时都可能遇到无限重复调用，严重影响开发效率，急需官方排查。
- **子 agent 失败恢复困难**：#39196 暴露的任务编排缺陷让开发者对复杂工作流失去信心，当前无临时规避方案。
- **上下文窗口不准确**：#35863 指出硬编码导致模型过早压缩，影响长上下文任务质量，开发者期待动态元数据支持。
- **免费模型体验差**：#39204 中 deepseek-v4-flash-free 频繁停止，使得低成本尝试 OpenCode 的用户受挫。
- **桌面设置 Bug 高频**：主题修改只能一次、设置页刷新等低级别 Bug 虽不大但反复出现，说明桌面端测试覆盖不足。
- **OAuth 登录失败**：#39207 显示 GitHub OAuth 因邮箱为空导致 SQL 错误，影响用户注册流程，需紧急修复。

*日报生成基于 GitHub 仓库 anomalyco/opencode 公开数据，数据截止 2026-07-28 UTC。*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# 2026-07-28 Pi 社区动态日报

## 📊 今日速览

昨日社区异常活跃：共处理 **28 个 Issue**（含 18 个已关闭）和 **26 个 PR**，核心聚焦于扩展 API 完善、多模态 tokenizer 崩溃修复以及会话模型作用域暴露。**#5263**（会话内模型默认临时化）以 10 👍 成为社区最关注的需求；同时一项恶意包报告（#7186）引发安全警示。性能方面，**#7194** 揭露远程沙箱下全量重绘的 1 秒周期问题，已被标记为 `bug`。

---

## 🔥 社区热点 Issues

### 1. [#5263] Make in-session model and thinking-level changes ephemeral by default  
**状态：OPEN | 评论：10 | 👍：10**  
提议将会话内对模型和思考级别的修改默认设为临时（仅当前会话有效），并在 `/settings` 新增“默认模型”入口管理全局配置。社区高度期待这一行为一致性改进。  
[链接](https://github.com/earendil-works/pi/issues/5263)

### 2. [#5023] terminal scrolls to beginning without reason  
**状态：CLOSED | 评论：10 | 👍：0**  
反馈终端在模型运行时随机跳到会话开头再快速滚动到底部，影响编辑体验。虽已关闭但根因未完全公开，需关注后续类似报告。  
[链接](https://github.com/earendil-works/pi/issues/5023)

### 3. [#6747] An API for enhancing agent message markdown  
**状态：OPEN | 评论：8 | 👍：2**  
希望允许扩展在展示层修改 agent 消息的 markdown 渲染，而不影响发送给 LLM 的原始内容。该需求推动扩展能力边界，社区讨论活跃。  
[链接](https://github.com/earendil-works/pi/issues/6747)

### 4. [#7157] OpenCode Go provider displays as "OpenCode Zen Go"  
**状态：OPEN | 评论：5 | 👍：0**  
显示名称错误导致 `pi --list-models` 输出不一致，属于低认知负担但影响用户信任的 UI 缺陷，已有 PR 修复。  
[链接](https://github.com/earendil-works/pi/issues/7157)

### 5. [#7161] anthropic-messages never sends x-client-request-id  
**状态：OPEN | 评论：4 | 👍：0**  
Claude 路径缺少 session 亲和性头，导致网关无法正确路由用户会话，对使用多账户代理的用户影响明显。  
[链接](https://github.com/earendil-works/pi/issues/7161)

### 6. [#7128] New default PI_* guideline in system prompt over-encourages unnecessary bash calls  
**状态：CLOSED | 评论：3 | 👍：0**  
系统提示新增的“检查 PI_* 环境变量”引导导致 agent 频繁执行 env 查询，浪费 token 与延迟。社区认为应优化提示词设计。  
[链接](https://github.com/earendil-works/pi/issues/7128)

### 7. [#7171] Dedupe byte-identical context files in the cwd->root walk  
**状态：CLOSED | 评论：3 | 👍：0**  
当前上下文文件加载仅按路径去重，未考虑字节内容完全相同的文件（如 worktree 与根目录的 `AGENTS.md`），建议增加内容哈希去重。  
[链接](https://github.com/earendil-works/pi/issues/7171)

### 8. [#7192] Expose session scoped models to extensions (ctx.scopedModels)  
**状态：CLOSED | 评论：2 | 👍：0**  
扩展无法读取当前会话的 `ScopedModel` 列表，限制自定义模型选择器的实现。已通过 PR #7191 合入。  
[链接](https://github.com/earendil-works/pi/issues/7192)

### 9. [#7190] setCustomEditorComponent copies stale borderColor from defaultEditor instead of the active editor  
**状态：CLOSED | 评论：2 | 👍：0**  
自定义编辑器边框颜色被错误的默认编辑器覆盖，影响 UI 一致性。已关闭并合入修复。  
[链接](https://github.com/earendil-works/pi/issues/7190)

### 10. [#7170] Support for aws (bedrock) credential_process  
**状态：CLOSED | 评论：2 | 👍：0**  
`~/.aws/credentials` 中使用 `credential_process` 外部凭证方式的用户在 Pi 中无法使用 Bedrock，已通过配置优先级修复（参见 PR #7176）。  
[链接](https://github.com/earendil-works/pi/issues/7170)

---

## 🔧 重要 PR 进展

### 1. [#7022] WIP: fix(coding-agent): guard tree navigation during responses  
**状态：OPEN**  
尝试修复在 agent 流式响应时使用 `/tree` 导航引发的怪异行为。目前为 PoC，社区可提前测试。  
[链接](https://github.com/earendil-works/pi/pull/7022)

### 2. [#7163] feat: search index sqlite  
**状态：OPEN**  
为 SQLite 存储增加 `SessionRepo.search()` 方法与 FTS5 全文搜索支持，大幅提升会话检索能力。JSONL 和内存暂仍为全量加载。  
[链接](https://github.com/earendil-works/pi/pull/7163)

### 3. [#7191] feat(extensions): expose ctx.scopedModels to extensions  
**状态：CLOSED**  
解决扩展无法获取会话已启用模型列表的问题，对构建第三方模型选择器至关重要。  
[链接](https://github.com/earendil-works/pi/pull/7191)

### 4. [#7081] feat(ai): support Claude Opus 5 on Bedrock  
**状态：CLOSED**  
配置 Claude Opus 5 使用自适应思考（必需），同时修复 Bedrock 错误消息详细泄露问题。  
[链接](https://github.com/earendil-works/pi/pull/7081)

### 5. [#7184] fix(ai): strip multimodal media markers from tool results to prevent tokenizer crashes  
**状态：CLOSED**  
工具结果中残留 `|image|` 等标记但无实际图片时，多模态 tokenizer 崩溃。移除标记以解决此高风险问题。  
[链接](https://github.com/earendil-works/pi/pull/7184)

### 6. [#7183] test(settings): add regression tests for autocompleteMaxVisible persistence  
**状态：CLOSED**  
为 #7179 报告的值未持久化问题添加回归测试，确保 `SettingsManager` 正确读写。  
[链接](https://github.com/earendil-works/pi/pull/7183)

### 7. [#7178] feat(coding-agent): show status when toggling tool-output expansion  
**状态：CLOSED**  
为 Ctrl+O 切换工具输出展开增加状态行提示，与思考块切换保持一致，提升操作反馈。  
[链接](https://github.com/earendil-works/pi/pull/7178)

### 8. [#7168] feat: auth print  
**状态：CLOSED**  
新增 `auth print-api-key` 与 `print-bearer-token` 命令，方便调试认证配置。  
[链接](https://github.com/earendil-works/pi/pull/7168)

### 9. [#7176] fix(ai): prefer configured Bedrock profile over ambient AWS keys  
**状态：OPEN**  
修复当环境变量 `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` 存在时，Pi 配置的 Bedrock 配置文件被忽略的问题。  
[链接](https://github.com/earendil-works/pi/pull/7176)

### 10. [#6881] feat(ai): use provider-reported cost when responses include it  
**状态：OPEN**  
当返回结果包含实际费用字段时，优先使用提供者报告的成本而非目录费率，目前支持 OpenAI 兼容路径。  
[链接](https://github.com/earendil-works/pi/pull/6881)

---

## 📈 功能需求趋势

通过分析昨日所有 Issue，社区主要关注以下方向：

- **扩展 API 完善**：要求暴露更多运行时数据（作用域模型、终端颜色方案、事件总线生命周期），并允许扩展修改消息展示（#6747、#7192、#7197、#7193）。
- **多模态处理增强**：tokenizer 对无数据标记的脆弱性引发两次修复（#7184、#7181），社区期望更健壮的媒体标记处理。
- **性能优化**：大缓存环境下的 `visibleWidth` 竞争（#7196）、远程沙箱全量重绘（#7194）成为热点。
- **新模型与提供商支持**：Claude Opus 5（#7081）、Z.AI 的 `max_tokens` 兼容（#7174）、Merge Gateway 内置请求（#5986 虽已关闭但仍有价值）。
- **配置与持久化改进**：`autocompleteMaxVisible` 值重启后丢失（#7179）、`autoload` 配置文件路径透传问题（#7179），社区期待更可靠的持久化机制。

---

## 👨‍💻 开发者关注点

### 高频痛点

1. **终端行为异常**  
   - 随机滚动到顶部（#5023）  
   - Shift+Enter 在 Windows Terminal 下提交而非换行（#7175）  
   - 键盘协商后无法确定可交互状态（#7177）

2. **扩展安装与加载问题**  
   - 扩展目录为符号链接时不被识别（#7195）  
   - git 安装失败后污染目录，后续重试无效（#7189）  
   - `pi install git:` 误装 `peerDependencies` 而 npm 来源不会（#7182）

3. **网关与认证兼容性**  
   - Anthropic 路径缺少 `x-client-request-id` 导致多账户代理无法会话保持（#7161）  
   - AWS Bedrock `credential_process` 不支持（#7170）  
   - 恶意包报告（#7186）警示第三方包安全性

4. **性能开销**  
   - 远程沙箱全量重绘周期 1s（#7194）  
   - 大缓存下 `visibleWidth` 计算成为 CPU 热点（#7196）  
   - 系统提示引导 agent 执行多余 env 命令（#7128）

### 社区情绪

- **积极**：扩展 API 逐渐开放（`scopedModels`、颜色方案）获得开发者好评；多模态 tokenizer 崩溃的快速修复表明团队响应及时。  
- **期待**：希望暴露更多内部状态（如事件总线生命周期、当前会话阶段），以便构建更智能的第三方工具。  
- **谨慎**：恶意包事件提醒加强包审核；部分 bug 长期存在（终端滚动、重绘）需要在稳定性上投入更多。

---

*数据来源：GitHub earendil-works/pi，统计时间 2026-07-27 00:00 - 23:59 UTC*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

好的，没问题。作为专注于 AI 开发工具的技术分析师，我将根据您提供的 GitHub 数据，为您生成一份结构清晰、内容专业的 Qwen Code 社区动态日报。

---

# Qwen Code 社区动态日报 (2026-07-28)

## 今日速览

今日社区动态活跃，核心聚焦于三项进展：**核心 CLI 的本地化修复与仓库治理（如三年前 PR 失败检测）**、**Web Shell 的 Git 工作流与语音交互功能增强**，以及两项关于 **企业级外部上下文与记忆集成的功能提案** 引发社区讨论。此外，昨日发布的两个 `dsw-manual-poc` 基准测试预发布版本显示，在 SWE-bench Verified 上取得了 **376/500 的解决率**，展现了出色的代码修复能力。

## 版本发布

- **[v0.21.0-nightly.20260727.c003e1718]**
  - **主要内容**: 这是一个每日构建的 Nightly 版本。主要修复了 CLI 中关于“洞察”天数/小时度量在本地时间显示不一致的问题。
  - **标签**: `fix`, `cli`
  - **链接**: [查看详情](https://github.com/QwenLM/qwen-code/releases/tag/v0.21.0-nightly.20260727.c003e1718)

## 社区热点 Issues (Top 10)

1.  **#7585: 提议：添加直接外部上下文提供者配置**
    - **重要性**: 该提案旨在为 Qwen Code 添加一个官方扩展，允许 CLI 进程从管理员绑定的外部知识服务获取仓库级共享上下文，而无需修改核心。这是 **企业级知识管理与集成** 的关键需求。
    - **社区反应**: 9条评论，处于 `need-discussion` 状态，表明社区正在积极探讨其实现细节和与现有方案的兼容性。
    - **链接**: https://github.com/QwenLM/qwen-code/issues/7585

2.  **#7449: 提议：定义企业外部记忆集成配置**
    - **重要性**: 与 #7585 相辅相成，聚焦于定义供应商无关的企业外部记忆（Memory）集成标准。这是一个**文档优先、向后兼容**的提案，为 Qwen Code 接入外部记忆系统（如向量数据库、知识图谱）奠定基础。
    - **社区反应**: 6条评论，社区关注如何在不改核心 API 的前提下实现兼容性。
    - **链接**: https://github.com/QwenLM/qwen-code/issues/7449

3.  **#7861: 主分支 CI 失败：E2E 测试**
    - **重要性**: **主分支（main）CI 流水线失败是一个严重信号**，直接影响所有后续开发的质量保障。此 Issue 由机器人自动创建，标记为 `status/ready-for-agent`，表明应优先处理。
    - **社区反应**: 2条评论，开发者正关注解决方案。
    - **链接**: https://github.com/QwenLM/qwen-code/issues/7861

4.  **#7167: 集群 Shepherd 仪表盘**
    - **重要性**: 由机器人自动维护的 CI/CD 状态仪表盘。它能直观展示当前 PR 的工作流程状态（如空闲、同步中），是团队**监控 CI 健康度**和管理合并队列的窗口。
    - **社区反应**: 4条评论，自动化维护中。
    - **链接**: https://github.com/QwenLM/qwen-code/issues/7167

## 重要 PR 进展 (Top 10)

1.  **#7731: 为 Web Shell 添加 Git 分支选择器、提交对话框和创建 PR 流程**
    - **内容**: 模拟 IntelliJ 风格的 Git 面板。添加了分支切换、新建、检出功能，以及完整的提交和创建 PR (Pull Request) 流程，极大提升了 Web Shell 的 IDE 体验。
    - **标签**: `feat`, `web-shell`, `git`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7731

2.  **#7842: 核心修复：对永久性配额用尽的429错误快速失败，而非静默重试**
    - **内容**: 识别 API 返回的 `429` 状态码中表示“永久性配额用尽”的错误，并立即返回友好提示，避免无意义的静默重试，提升用户体验和资源效率。
    - **标签**: `fix`, `core`, `self-reported`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7842

3.  **#7859: 为 Web Shell 添加原生实时语音功能（macOS）**
    - **内容**: 引入原生 Live Voice 体验。用户可在任何应用中通过快捷键（Cmd+Cmd）唤起对话，支持开始/切换、静音等操作，是**语音交互体验上的一大步**。
    - **标签**: `feat`, `web-shell`, `voice`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7859

4.  **#7484: 核心修复：为纯文本模型桥接工具返回的图像结果**
    - **内容**: 允许纯文本的基础模型“理解”工具执行过程中发现的图片。这使得文本模型也能处理来自内置工具、MCP 和扩展工具的图像结果，增强了模型的工具处理能力。
    - **标签**: `fix`, `core`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7484

5.  **#7414: 功能：基于回滚历史的高风险路径检测**
    - **内容**: 分析仓库111条回滚提交和46个被回滚的PR，引入一个数据驱动的审查门禁，自动检测**高风险变更路径**，辅助代码审查。
    - **标签**: `feat`, `triage`, `self-reported`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7414

6.  **#7863: 核心修复：使用 `-e` 参数传递 Grep 模式，避免前导破折号被误解为选项**
    - **内容**: 修复 Grep 工具，当搜索模式以破折号开头（如 `-foo`）时，会因被解析为命令行选项而失败的问题。
    - **标签**: `fix`, `core`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7863

7.  **#7827: 修复安全模式：保留调用方指定的顶级 MCP 服务器**
    - **内容**: 修复在 `--safe-mode` 下，所有 MCP 服务被无条件丢弃的问题。现在会保留用户显式指定的“顶级”MCP 服务，仅过滤本地/环境配置的。
    - **标签**: `fix`, `safe-mode`, `mcp`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7827

8.  **#7877: 功能：为外部上下文添加“自动回忆”功能**
    - **内容**: 在现有的私有直接外部上下文集成中，增加一个管理员可安装的、确定性的“自动回忆”（Auto Recall）配置项。它作为 Hook 运行，在用户提交提示时自动触发，与按需 MCP 配置互斥。
    - **标签**: `feat`, `external-context`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7877

9.  **#7856: 功能：为 Web Shell 添加 Composer 底部渲染器**
    - **内容**: 提供一个自定义渲染点，允许宿主在聊天组件的 Composer（输入框）后立即显示上下文内容，增强了界面的可扩展性。
    - **标签**: `feat`, `web-shell`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7856

10. **#7820: 修复测试：恢复首次输出基准测量的有效性并修正其产物模式**
    - **内容**: 作为已合入 PR #7761 的后续修复，处理了其合并前未被采纳的十条审查意见，涉及 opt-in 测量和测试基础设施。
    - **标签**: `fix`, `test`, `infrastructure`
    - **链接**: https://github.com/QwenLM/qwen-code/pull/7820

## 功能需求趋势

从今日的 Issues 和 PR 中，可以清晰看到社区与开发团队关注的三大功能方向：

1.  **企业级集成与扩展性**: 社区强烈关注如何将 Qwen Code 无缝集成到企业工作流中。这表现为对 **外部上下文提供商（#7585）** 和 **外部记忆集成（#7449）** 的提案，以及增强 **安全模式（#7827）** 和 **MCP 服务（#7847）** 的灵活控制。核心在于“可插拔”和“可管理”。

2.  **沉浸式开发体验 (IDE 化)**: 尽管是 CLI 工具，但团队正大力增强其用户体验。这包括为 **Web Shell 添加完整的 Git 工作流（#7731）**、**原生实时语音交互（#7859）** 以及 **原生文件夹选取器（#7849）**，旨在提供媲美桌面 IDE 的交互体验。

3.  **健壮性与可观测性**: 代码质量和系统稳定性是永恒的主题。趋势集中在：**更智能的限流/错误处理（#7842）**、**基于数据分析的风险审查（#7414）**、**修复核心工具（如 Grep）的边界问题（#7863）**，以及保障 **CI/CD 的可靠性（#7861）**。

## 开发者关注点

从 Bug 修复和社区反馈中，可以提炼出开发者最敏感的痛点和高频需求：

-   **CLI 本地化问题**: 曾有开发者反馈“洞察”数据的度量（时长）在**本地时间显示不一致**。虽然已在 v0.21.0-nightly 修复，但这反映了开发者对 **CLI 输出准确性与区域设置适配** 的细颗粒度要求。
-   **API 限流的透明处理**: 开发者不希望 API 返回的 `429` 状态码被“静默”处理。PR #7842 的提出，说明开发者需要**明确、友好的错误反馈**，以便快速了解资源配额状态并采取行动。
-   **Git 工作流的原生支持**: 大量关于 Web Shell 增强的 PR（尤其是 #7731）表明，CLI 用户期望获得一套**完整、流畅的 Git 工作流支持**，这几乎是所有开发者进行版本协作的基础需求。
-   **工具链的边界情况**: 修复 Grep 工具（#7863）等PR表明，开发者对**命令行工具的健壮性**有极高要求，任何微小的参数异常（如前导破折号）都不应被忽视。

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI 社区动态日报 | 2026-07-28

## 今日速览
过去 24 小时无新版本发布，但社区围绕 **v0.9.2 候选版本** 进行了大规模集成，合并了十余项关键功能，涵盖 TUI 视觉改进、Fleet 精确保存、会话持久化、Lane 控制平面等。一个关于前台 shell 抢占输入的热门 Bug 被报告（ #4930），成为当前开放讨论焦点。此外，**中文翻译二次润色**（ #4908）和 **OpenCode Zen 提供商支持**（ #4467）也已顺利合入。

## 社区热点 Issues
（共 5 条，全部列出）

**#4042** `[已关闭]`  
**feat: Environment-level tool sandboxing for sub-agents**  
作者：JayBeest | 更新：2026-07-27 | 评论：20  
社区反馈热烈，讨论子代理环境级工具沙箱的运行时强制机制，涵盖 session、sub‑agent、Fleet worker 和 MCP 服务器。已确认 `--disallowed-tools` 参数生效。  
[🔗 查看 Issue](https://github.com/Hmbown/CodeWhale/issues/4042)

**#4925** `[已关闭]`  
**Add thinking_default_expanded setting for always-expanded reasoning blocks**  
作者：M-Maciej | 更新：2026-07-27 | 评论：1  
针对 SSH/tmux 用户空格键被终端层捕获的问题，新增配置项使思考块默认展开，无需手动按空格。已通过 PR #4928 实现。  
[🔗 查看 Issue](https://github.com/Hmbown/CodeWhale/issues/4925)

**#4907** `[已关闭]`  
**ci(web): main push always fails because deploy trigger contradicts manual-only preflight**  
作者：Hmbown | 更新：2026-07-27 | 评论：1  
Web 前端 CI 在 `main` 分支上因部署触发器与手动检查冲突而稳定失败，已修复。  
[🔗 查看 Issue](https://github.com/Hmbown/CodeWhale/issues/4907)

**#4751** `[已关闭]`  
**Settings IA rework: Fleet/Models section boundaries**  
作者：Hmbown | 更新：2026-07-27 | 评论：1  
根据用户截图反馈，重设设置信息架构：将误放在 Fleet 段落的 Goal‑command/Workflow 开关移至正确位置，移除已废弃的 “Legacy fallback model” 行。  
[🔗 查看 Issue](https://github.com/Hmbown/CodeWhale/issues/4751)

**#4930** `[开放中]`  
**Enter during foreground shell should detach it before steering**  
作者：M-Maciej | 更新：2026-07-27 | 评论：1  
当代理处于前台 Bash 命令（如 `sleep 30`）阻塞状态时，用户自然输入并回车无效。期望行为：自动分离前台 shell 后再处理用户指令。  
[🔗 查看 Issue](https://github.com/Hmbown/CodeWhale/issues/4930)

## 重要 PR 进展
（精选 10 条，按功能分类）

**#4911** `[已合并]`  
**v0.9.2 release candidate integration (umbrella)**  
作者：Hmbown  
v0.9.2 集成主 PR，包含 82 个 commits（超前 main），集成了 Fleet、Lane、会话、账单、TUI 可见性等全部候选功能。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4911)

**#4928** `[已合并]`  
**feat(tui): add thinking_default_expanded setting**  
作者：M-Maciej  
新增 `thinking_default_expanded` 配置，默认展开思考块。解决了 SSH/tmux 下空格键被捕获的问题，同时保持折叠/展开切换能力。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4928)

**#4920** `[已合并]`  
**fix: kimi-k3 selection — sticky model memory, lying resolve, missing catalog ids**  
作者：Hmbown  
根因修复三个缺陷：`provider_models` 会话内存无条件覆盖命令行 `--model` 标志；Kimi K3 模型路由错误导向旧版本；目录中缺少 K3 编号。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4920)

**#4927** `[已合并]`  
**fix(billing): dispatch-receipt classification, Moonshot/MiniMax product truth, honest ceilings**  
作者：Hmbown  
账单修复：一笔完成的轮次应从当时的分发收据结算，而非重读实时配置；Moonshot 拆分直连与平台计费；严格设置费率上限。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4927)

**#4919** `[已合并]`  
**feat: lane control-plane contract, nonblocking /lane interrupt, CLI/TUI fleet parity**  
作者：Hmbown  
实现 3250 行控制平面合约，定义了稳定的 `<domain>.<verb>` ID、读写权限、限额收据、别名及 `fenced TerminalTransition` 停止机制。达到 CLI 与 TUI 的 Fleet 操作一致性。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4919)

**#4924** `[已合并]`  
**feat(fleet): saved exact Fleets + reasoning Router — two-phase admission, verified ceilings**  
作者：Hmbown  
重写 Saved Fleet 与 Router：冻结精确 (provider, model) 路由，加入权限/shell 上限、角色别名规范化（oracle/advisor→consultant）、碰撞检测。Router 引入两阶段准入和收据验证。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4924)

**#4922** `[已合并]`  
**feat(sessions): persistent rail, opt-in auto-resume, dashboard peek**  
作者：Hmbown  
实现持久化会话轨道，支持归档、自动恢复（需用户确认）、会话侧面板及 `/sessions` 命令。自动恢复的决策类型明确区分显式标志与默认行为。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4922)

**#4467** `[已合并]`  
**Feat/opencode zen provider**  
作者：snail-vs  
新增 OpenCode Zen 作为模型感知提供商，支持 Zen 模型的路由（Responses/Anthropic Messages/Chat Completions），并添加 Zen 专用 API 密钥缺失提示及 Claude 认证修复（使用 `x-api-key`）。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4467)

**#4908** `[已合并]`  
**I18n(zh-Hans): update simplified-Chinese translations to match latest en.json**  
作者：SparkofSpike  
在 #4805 基础上进行第二轮简中翻译质量提升，对抗性审查全部 1134 个键，由专用 reviewer 子代理独立验证。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4908)

**#4931** `[开放中]`  
**Migrate QA PTY test harness from vt100 to rio-vt**  
作者：raphamorim  
将 TUI 测试工具从 `vt100` 切换为 `rio-vt`（Rio 终端引擎），提升 PTY 输出解析精度和渲染测试覆盖率。  
[🔗 查看 PR](https://github.com/Hmbown/CodeWhale/pull/4931)

## 功能需求趋势
从本周 Issues 和 PR 中可归纳出社区最关注的方向：

- **TUI 用户体验优化**：思考块默认展开（ #4925）、前台 shell 分离（ #4930）、视觉程序切片（ #4923）。
- **子代理与沙箱安全**：环境级工具沙箱（ #4042）、Fleet 精确权限限制（ #4924）。
- **多提供商与模型路由**：OpenCode Zen 提供商（ #4467）、StepFun 账单设置（ #4921）、Auto 模型路由范围限制（ #4917）。
- **会话与工作流持久化**：持久会话轨道与自动恢复（ #4922）、Fleet 精确保存（ #4924）。
- **国际化与本地化**：继英文后，简中翻译获得第二轮深度审查（ #4908）。
- **CI/CD 可靠性**：Web 部署触发器修复（ #4907）、Clippy 严格 lint 维护（ #4932）。

## 开发者关注点

- **SSH/tmux 兼容性**：空格键被终端层捕获导致无法展开思考块（ #4925），社区强烈要求默认展开或提供配置选项（已通过 4928 解决）。
- **前台 shell 阻塞时交互**：用户输入被忽略，期望自动分离或中断（ #4930），该 Issue 仍开放，需设计方案。
- **模型选择“粘性”错误**：`--model` 参数被会话内存覆盖（ #4920），开发者反馈“明明指定了 kimi-k3 却跑到 k2.7”，影响信任。
- **设置信息架构混乱**：Fleet 段落放置不相关的控制项（ #4751），用户截图反馈已通过 IA 重做修复。
- **账单逻辑模糊**：中途中切换提供商导致重读实时配置而非收据（ #4927），社区要求明确结算时间点。

> 数据来源：GitHub [Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale)（DeepSeek TUI 社区）。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*