# AI CLI 工具社区动态日报 2026-06-18

> 生成时间: 2026-06-18 03:18 UTC | 覆盖工具: 9 个

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

好的，各位技术决策者、开发者，以下是根据今日（2026-06-18）各主流 AI CLI 工具的社区动态，生成的横向对比分析报告。

---

# AI CLI 工具生态横向对比报告 | 2026-06-18

## 1. 生态全景

当前 AI CLI 工具生态正处于 **“功能趋同，体验分化”** 的成熟阶段。各工具在 Agent 自动化、多模型支持、IDE 集成等核心能力上已无明显代差，竞争焦点正转向**计费模型、Agent 可靠性、平台兼容性及企业级安全**等“用户体验与信任”维度。社区反馈显示，传统功能请求（如支持新模型）热度下降，取而代之的是对**可预测性、透明度和成本控制**的刚性需求，标志着市场正从早期采用者向主流开发者过渡。

## 2. 各工具活跃度对比

| 工具名称 | 今日 Issues | 今日 PRs | 版本发布 | 社区活跃度评级 |
| :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 10 (高关注) | 7 | ✅ v2.1.181 | ★★★★★ |
| **OpenAI Codex** | 10 (高关注) | 10 | ✅ 3个Alpha版本 | ★★★★★ |
| **Gemini CLI** | 10 (高关注) | 10 | ✅ v0.47.0 & v0.48.0-preview.0 | ★★★★☆ |
| **GitHub Copilot CLI** | 10 (高关注) | 0 | ❌ 无 | ★★★★☆ |
| **Kimi Code CLI** | 2 | 0 | ❌ 无 | ★☆☆☆☆ |
| **OpenCode** | 10 (高关注) | 10 | ✅ v1.17.8 | ★★★★☆ |
| **Pi** | 10 (高关注) | 10 | ❌ 无 | ★★★★★ |
| **Qwen Code** | 6 | 10 | ✅ v0.18.3 | ★★★★☆ |
| **DeepSeek TUI** | 10 (高关注) | 10 | ❌ 无 (v0.9.0开发中) | ★★★★★ |

**解读**：
- **Claude Code、OpenAI Codex、Pi、DeepSeek TUI** 社区活跃度最高，且Issue复杂度高，覆盖核心功能与架构性讨论。
- **OpenCode、Gemini CLI、Qwen Code** 迭代节奏稳定，PR与版本发布密集，处于快速功能开发期。
- **Kimi Code CLI** 活跃度极低，尚处早期，社区反馈微弱。

## 3. 共同关注的功能方向

多个工具社区不约而同地聚焦于以下议题，反映出开发者群体的普遍诉求：

| 共同方向 | 涉及工具 | 具体诉求 |
| :--- | :--- | :--- |
| **计费透明性与成本控制** | **Claude Code** (#16157)、**OpenAI Codex** (#28823) | 用户对订阅后仍被限额、额度消耗异常敏感，急需清晰的计费策略和实时的成本仪表盘。 |
| **Agent 行为可控性与可靠性** | **Claude Code** (#34255)、**Gemini CLI** (#21409)、**DeepSeek TUI** (#3275/3279) | Agent 挂起、自问自答、模式切换混乱、任务状态误报等，是普遍痛点。用户要求更可控、可预测的行为模式。 |
| **认证与安全** | **OpenAI Codex** (#25749)、**Gemini CLI** (#26525)、**OpenCode** (#1973) | 强迫使用遗留手机号验证、子代理任务状态误报、MCP工具权限管理等，表明安全与合规成为企业级采用的门槛。 |
| **MCP/插件生态成熟化** | **Claude Code** (#26094)、**OpenAI Codex** (#27132)、**GitHub Copilot CLI** (#3812)、**OpenCode** (#32751) | 社区期望 MCP 工具调用更可靠、权限管理更精细、身份追踪更清晰，以支持复杂自动化。 |
| **多平台与跨设备体验** | **Claude Code** (#39636)、**OpenAI Codex** (#25178)、**OpenCode** (#3541) | ARM Windows、Wayland Linux、老版本Windows 10等非主流平台的基础功能稳定性和兼容性是普遍需求。 |

## 4. 差异化定位分析

| 工具名称 | 核心定位 | 目标用户 | 技术路线与优势 |
| :--- | :--- | :--- | :--- |
| **Claude Code** | **深度Agent与多任务管理** | 高级个体开发者、技术主管 | 强调Agent的复杂任务编排能力、`/config`动态调整、远程控制。社区最关注其**计费透明度**和**MCP生态成熟度**。 |
| **OpenAI Codex** | **平台集成与多模态扩展** | 全栈开发者、企业用户 | 紧密集成 OpenAI 生态，探索实时语音、Computer Use、Rust 内核重写。痛点集中在**认证安全**与**平台稳定性**。 |
| **Gemini CLI** | **云端Agent与安全优先** | Google Cloud用户、注重隐私的开发者 | 强调 Google 级的安全审计（CI/CD守卫）、组件级评测。社区聚焦**Agent稳定性**和**子代理行为控制**。 |
| **GitHub Copilot CLI** | **企业级安全与插件系统** | GitHub 企业用户、大型团队 | 优势在于 GitHub 生态的深度集成和插件扩展能力。核心痛点是**企业自定义模型接入**和**工具权限管理**。 |
| **Kimi Code CLI** | **轻量、场景化（潜在）** | 早期探索者 | 极其早期，社区反馈几乎空白。从仅有的两个Issue看，未来可能致力于**运行态模式切换**和**企业环境适配**。 |
| **OpenCode** | **开源、开放的 Agent 终端** | 社区驱动的开发者 | 开源是其最大标签，社区贡献活跃（多代理、模糊编辑）。关注点是**多代理编排**、**平台兼容性**和**运行时权限**。 |
| **Pi** | **多模型终端与极致UX** | 模型研究者、重度终端用户 | 支持极多模型（Azure, Anthropic OAuth等），关注**底层架构优化**（依赖管理）和**交互细节打磨**（流式渲染）。 |
| **Qwen Code** | **国际化与本地化并重** | 全球开发者，尤其关注东亚市场 | 积极适配国内外模型和通讯平台（QQ/Telegram），社区活跃。痛点集中在**模型选择持久化**和**SSH下TUI稳定性**。 |
| **DeepSeek TUI** | **从本地工具到协作平台** | 团队协作开发者 | 定位转型明确：通过“Workrooms”实现协作。当前社区中心矛盾是**Agent行为可控性**与**新模式带来的不确定性**。 |

## 5. 社区热度与成熟度

- **成熟社区（高热度、高复杂性）**：**Claude Code** 和 **OpenAI Codex** 的社区最为成熟，讨论内容已超越基础功能，深入到计费、安全、架构等高级议题，Issue 质量高，代表了行业顶尖用户的诉求。
- **快速迭代期（高热度、高产出）**：**OpenCode**、**Pi**、**DeepSeek TUI** 处于快速迭代期，社区反馈积极且直接驱动了 PR 和版本发布，体现了开源或开放社区的强大生命力。
- **稳定发展期（中等热度、技术导向）**：**Gemini CLI** 和 **Qwen Code** 社区活跃度稳定，讨论偏向技术实现和功能优化，路线清晰。
- **早期探索期（低热度）**：**Kimi Code CLI** 社区尚处早期，用户反馈有限，功能与定位仍在摸索中。

## 6. 值得关注的趋势信号

1.  **从“增强协作”到“强化平台”**：以 **DeepSeek TUI** 的“Workrooms”和 **Claude Code** 的“Remote Control”为代表，AI CLI 工具正从单用户终端向**多用户协作平台**演进。开发者需关注这类工具的团队协作能力和数据管理方案。

2.  **“成本焦虑”倒逼产品透明化**：**Claude Code** 和 **OpenAI Codex** 的高赞计费Issue表明，开发者对成本的敏感度已远超预期。为AI工具建立清晰的“成本仪表盘”和“额度预警机制”将成为**必备功能**，而非锦上添花。

3.  **Agent 协作模式遭遇“信任危机”**：几乎所有工具的社区都对Agent“自作主张”的行为表达了不满。这提示我们，未来的AI CLI需要提供更**明确的上下文管理、权限“白名单”、和“撤销/回退”能力**，以重建用户信任。**“安全默认”原则必须回归。**

4.  **MCP/插件生态成为“必争之地”**：从API兼容性到权限粒度，各工具都在积极构建自己的生态护城河。对于开发者而言，选择生态更开放、集成更顺滑的AI CLI，未来将拥有更强的扩展能力和更低的迁移成本。

5.  **TUI（终端界面）体验进入“精耕细作”期**：多种工具的Issue关注流式渲染卡顿、ANSI乱码、导航截断等细节，说明TUI体验已成为影响用户日常效率的关键因素。单纯的“能工作”已无法满足需求，**“优雅且高性能”的终端交互**是下一阶段竞争的重点。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills 社区热点报告（数据截至 2026-06-18）

本报告基于 anthropics/skills 官方仓库的 Pull Requests（Top 20 按评论数排序）及 Issues（全部 13 条）分析生成，聚焦社区关注度最高的 Skills 动态、需求趋势及生态洞察。

## 1. 热门 Skills 排行

以下列出评论数最高的 5~8 个新增/改进 Skills，涵盖功能、社区讨论焦点及当前状态（均为 **Open**）。

| 排名 | Skill/PR | 功能简介 | 社区讨论热点 | 状态 |
|------|----------|----------|--------------|------|
| 1 | **document-typography**<br>[PR #514](https://github.com/anthropics/skills/pull/514) | 防止 AI 生成文档中的孤字换行、孤立段落、编号错位等排版问题 | 用户普遍反映生成文档排版粗糙，该技能直接解决高频痛点；社区期待集成到所有文档技能中 | Open |
| 2 | **ODT**<br>[PR #486](https://github.com/anthropics/skills/pull/486) | 创建、填充、读取、转换 OpenDocument 格式（.odt/.ods），支持 LibreOffice | 填补了开源文档格式支持空白；讨论集中在 ISO 标准兼容性和与 DOCX 技能的差异 | Open |
| 3 | **frontend-design**（改进）<br>[PR #210](https://github.com/anthropics/skills/pull/210) | 重构前端设计技能，提升指令清晰度与可执行性 | 关注如何让 Claude 在单次对话中可靠遵循设计指引；强调指令具体化而非泛泛描述 | Open |
| 4 | **skill-quality-analyzer + skill-security-analyzer**<br>[PR #83](https://github.com/anthropics/skills/pull/83) | 元技能：对 Claude Skills 进行质量评估（结构/文档/可测试性等）和安全分析 | 社区对技能质量缺乏客观标准；安全分析引起对恶意技能防护的讨论 | Open |
| 5 | **SAP-RPT-1-OSS predictor**<br>[PR #181](https://github.com/anthropics/skills/pull/181) | 基于 SAP 开源表格基础模型进行预测分析，处理 SAP 业务数据 | 企业级用户关注 ERP 数据集成；讨论模型调用方式和权限要求 | Open |
| 6 | **testing-patterns**<br>[PR #723](https://github.com/anthropics/skills/pull/723) | 覆盖测试哲学、单元测试、React 组件测试、E2E 测试等完整测试栈指导 | 测试是开发者通用需求；社区讨论 AAA 模式、测试 Trophy 模型以及不测试反模式 | Open |
| 7 | **shodh-memory**<br>[PR #154](https://github.com/anthropics/skills/pull/154) | 持久记忆系统，让 AI 代理跨对话保持上下文 | 聚焦记忆结构设计（按实体/关系/时间戳）；讨论隐私与上下文窗口占用 | Open |
| 8 | **ServiceNow**<br>[PR #568](https://github.com/anthropics/skills/pull/568) | 完整 ServiceNow 平台技能，覆盖 ITSM、ITOM、SecOps、ITAM、CSDM 等 | 企业复杂性高，社区关注脚本生成与 CMDB 查询的可靠性；讨论与现有 CMDB 技能的冲突 | Open |

## 2. 社区需求趋势（来自 Issues）

Issues 中社区最期待的新 Skill 方向或改进诉求：

- **组织级技能共享与分发**（[#228](https://github.com/anthropics/skills/issues/228)，👍7）—— 用户希望直接在 Claude.ai 内共享技能，避免手动导出/导入。此 Issue 评论最多，且获得高赞，反映企业部署强需求。
- **评估工具稳定性**（[#556](https://github.com/anthropics/skills/issues/556) & [#1169](https://github.com/anthropics/skills/issues/1169)）—— `run_eval.py` 普遍报告 recall=0%，导致优化循环失效。社区强烈要求修复技能触发检测。
- **安全信任边界**（[#492](https://github.com/anthropics/skills/issues/492)）—— 社区技能混入 `anthropic/` 命名空间引发信任风险，呼吁命名空间隔离或官方审核。
- **Agent 治理与安全模式**（[#412](https://github.com/anthropics/skills/issues/412)）—— 社区提案“agent-governance”技能，涵盖策略执行、威胁检测、审计追踪，但尚未实现。
- **跨平台兼容性**（[#1061](https://github.com/anthropics/skills/issues/1061) & [#29](https://github.com/anthropics/skills/issues/29)）—— Windows 用户报告 subprocess、编码错误；AWS Bedrock 用户询问技能兼容性。平台适配是持续痛点。
- **技能质量标准化**（[#202](https://github.com/anthropics/skills/issues/202)）—— skill-creator 被批评过于冗长，不符合最佳实践，社区希望得到更简洁、可执行的模板。
- **重复技能与命名冲突**（[#189](https://github.com/anthropics/skills/issues/189)）—— `document-skills` 和 `example-skills` 安装相同内容，导致重复占用上下文窗口。

## 3. 高潜力待合并 Skills

以下 PR 评论活跃（按排序靠前）且尚未合并，近期可能落地：

| Skill PR | 亮点 | 落地潜力 |
|----------|------|----------|
| [document-typography #514](https://github.com/anthropics/skills/pull/514) | 解决 AI 文档生成最高频排版缺陷，逻辑清晰，社区呼声高 | ⭐⭐⭐⭐⭐ |
| [skill-quality-analyzer #83](https://github.com/anthropics/skills/pull/83) | 元技能，可提升整个生态质量，符合 Anthropic 治理需求 | ⭐⭐⭐⭐ |
| [testing-patterns #723](https://github.com/anthropics/skills/pull/723) | 覆盖完整测试栈，开发者通用需求，内容详实 | ⭐⭐⭐⭐ |
| [ODT #486](https://github.com/anthropics/skills/pull/486) | 填补 ISO 文档格式空缺，与 LibreOffice 生态良好对接 | ⭐⭐⭐ |
| [SAP predictor #181](https://github.com/anthropics/skills/pull/181) | 企业级特色鲜明，有明确的 SAP 用户群体 | ⭐⭐⭐ |
| [AURELION suite #444](https://github.com/anthropics/skills/pull/444) | 结构化认知框架，整合 kernel/advisor/agent/memory 四个子技能，体系完善 | ⭐⭐⭐ |

## 4. Skills 生态洞察

**一句话总结：** 当前社区最集中的诉求是**提升技能质量与工具链可靠性**——包括评估脚本修复、安全命名空间隔离、跨平台兼容性，并迫切期待组织级技能共享机制的落地，以支撑企业规模化采用。

> 数据来源：[anthropics/skills](https://github.com/anthropics/skills)

---

好的，各位开发者，以下是基于 GitHub 上 Anthropics/Claude-Code 仓库公开数据生成的 2026 年 6 月 18 日 Claude Code 社区动态日报。

---

# Claude Code 社区动态日报 | 2026-06-18

## 今日速览

- **v2.1.181 发布**：新增 `/config` 命令语法，允许直接在对话中修改设置（如 `thinking` 模式），并添加了 `sandbox.allowAppleEvents` 配置项。
- **社区热议：计费与定价问题**：`#16157`（Max 订阅立即达使用限制）和 `#17432`（印度区专属定价）两个 Issue 保持极高关注度，反映了用户对成本透明度和区域化定价的迫切需求。
- **远程控制可靠性成焦点**：Issue `#34255` 引发关于 Remote Control 自动重连失败的广泛讨论，成为社区高频反馈的技术痛点。

---

## 版本发布

### v2.1.181

**链接**: [Release v2.1.181](https://github.com/anthropics/claude-code/releases/tag/v2.1.181)

**更新亮点**：

- **`/config` 命令增强**: 现在可以使用 `/config key=value` 的语法直接在提示中修改任何设置（例如 `/config thinking=false`），该功能支持交互模式、`-p` 参数模式以及 Remote Control 模式。这大幅提升了开发者在工作流中动态调整 Claude 行为的灵活性。
- **Apple Events 支持**: 新增 `sandbox.allowAppleEvents` 设置项，允许在沙盒环境下的命令发送 Apple Events（仅限 macOS）。这对于需要与 macOS 原生应用交互的自动化脚本非常重要。
- **新增环境变量**: 加入了 `CLAUDE_CLIENT_P` 环境变量（详情待补充）。

---

## 社区热点 Issues

1.  **`[BUG] Instantly hitting usage limits with Max subscription`** (`#16157`)
    - **链接**: [Issue #16157](https://github.com/anthropics/claude-code/issues/16157)
    - **重要性**: **极高**。评论数 (1475) 和点赞数 (691) 均为社区最高，这是一个持续半年的严重计费问题。用户反馈在购买 Max 订阅后，仍会立即遇到使用限制，严重影响了付费用户的体验。社区对此情绪强烈，是 Anthropic 当前面临的最大信任危机之一。
    - **社区反应**: 评论区有大量用户提供日志和复现步骤，要求官方澄清计费策略并修复此 Bug。

2.  **`[FEATURE] India-Specific Pricing Plans (INR) for Claude & Claude Code`** (`#17432`)
    - **链接**: [Issue #17432](https://github.com/anthropics/claude-code/issues/17432)
    - **重要性**: **高**。评论数 (198) 和点赞数 (444) 都表明这是一个跨区域的普遍需求。用户指出 OpenAI 和 Google 均已提供本地化定价，Anthropic 的纯美元定价将许多新兴市场的开发者拒之门外。
    - **社区反应**: 印度开发者社区积极反馈，并提供了大量数据支持本地化定价的合理性。

3.  **`[BUG] Remote Control: automatic reconnection doesn't work`** (`#34255`)
    - **链接**: [Issue #34255](https://github.com/anthropics/claude-code/issues/34255)
    - **重要性**: **高**。Remote Control 是跨设备协作的核心功能，自动重连失败（连接静默断开且无法恢复）严重影响了用户体验，尤其是在移动办公或网络不稳定的场景下。
    - **社区反应**: 用户正在积极提供日志和网络环境信息，以协助定位根因。

4.  **`[FEATURE] Message queue mode`** (`#50246`)
    - **链接**: [Issue #50246](https://github.com/anthropics/claude-code/issues/50246)
    - **重要性**: **高**。点赞数高达 99，反映了开发者在多任务并发场景下的真实痛点。当前 Claude 执行任务时，任何新输入都会中断当前工作，用户急需一个消息队列来暂存想法。
    - **社区反应**: 社区对此设计方向高度认可，并补充了关于队列优先级和排序的讨论。

5.  **`[BUG] Cowork VM guest kernel never boots on Snapdragon X Plus (ARM64)`** (`#39636`)
    - **链接**: [Issue #39636](https://github.com/anthropics/claude-code/issues/39636)
    - **重要性**: **高**。随着 ARM 架构 Windows 设备（如 Surface Pro X）的普及，Cowork 功能在此类平台上的完全不可用是一个重大的平台兼容性缺陷。
    - **社区反应**: 用户已提供详细的硬件信息和日志，但问题尚未解决，这影响了 ARM Windows 用户对 Cowork 功能的信任。

6.  **`[BUG] Drag and drop not working in VS Code extension chat panel`** (`#25128`)
    - **链接**: [Issue #25128](https://github.com/anthropics/claude-code/issues/25128)
    - **重要性**: **高**。这是一个影响广泛的基础功能回归问题。VS Code 是最大的 IDE 用户群体，拖拽功能的失效（从 v2.1.6 开始）严重降低了 IDE 插件的可用性。
    - **社区反应**: 用户确认该问题在多个版本后仍未修复，并提供了详细的对比测试结果。

7.  **`[BUG] Bash tool calls emitted as raw <invoke> text`** (`#63870`)
    - **链接**: [Issue #63870](https://github.com/anthropics/claude-code/issues/63870)
    - **重要性**: **高**。这是一个严重的核心功能 Bug。Bash 工具不执行而是输出原始文本，意味着代理化能力完全失效。用户提供了详细的 JSONL 日志，包含 23 次失败调用，证据确凿。
    - **社区反应**: 用户指出该问题与 `#61122` 等已知问题相似但更严重，期待官方给予更高优先级。

8.  **`[BUG] Claude Desktop 1.1.3189 (Cowork) serializes MCP object parameters as strings`** (`#26094`)
    - **链接**: [Issue #26094](https://github.com/anthropics/claude-code/issues/26094)
    - **重要性**: **高**。MCP (Model Context Protocol) 是 Claude 生态的核心扩展方式。此 Bug 导致 MCP 服务端接受到的对象参数变为字符串，会直接破坏依赖于复杂参数的 MCP 服务（如 Notion 集成）。
    - **社区反应**: 用户跟踪该问题已超过 4 个月，并在 Cowork 和 Desktop 客户端均验证了此 Bug。

9.  **`No visibility into background subagent activity`** (`#67485`)
    - **链接**: [Issue #67485](https://github.com/anthropics/claude-code/issues/67485)
    - **重要性**: **中**。这是一个关于“代理化”和“多代理协作”的功能性问题。当主代理在后台调度子代理时，UI 完全无反馈，用户无法感知任务进度，造成困惑。
    - **社区反应**: 用户描述了在复杂的迁移任务中使用此模式时遇到的具体问题，并提出了 UI 改进方案。

10. **`[BUG] UI diff/file view stays pinned to original harness worktree`** (`#65767`)
    - **链接**: [Issue #65767](https://github.com/anthropics/claude-code/issues/65767)
    - **重要性**: **中**。该问题影响了使用 `EnterWorktree` 工具进行多工作树开发的用户。UI 视图无法跟随上下文切换，可能导致开发者基于过时或错误的信息做出判断。
    - **社区反应**: 用户详细描述了 UI 与底层状态不一致的场景，该问题对复杂的 CI/CD 工作流影响较大。

---

## 重要 PR 进展

1.  **`[PR] Open Source Claude Code (Open)`** (`#41447`)
    - **链接**: [PR #41447](https://github.com/anthropics/claude-code/pull/41447)
    - **重要性**: 社区发起的开源倡议。试图关闭多项标签为“开源”的 Issue，虽未合入，但反映了社区对 Claude Code 开源的强烈期待。

2.  **`[PR] fix(code-review): allow re-reviews when new commits are pushed (Open)`** (`#19867`)
    - **链接**: [PR #19867](https://github.com/anthropics/claude-code/pull/19867)
    - **重要性**: 改进了 `code-review` 插件的核心逻辑。自动跳过重复审查，并在有新提交时允许重新审查，这对于 CICD 流程中的插件非常实用。

3.  **`[PR] Update frontend-design skill (Closed)`** (`#69226`)
    - **链接**: [PR #69226](https://github.com/anthropics/claude-code/pull/69226)
    - **重要性**: 对 `frontend-design` 技能进行了改进，并自动升级插件版本，表明官方正在持续迭代和优化内置技能包。

4.  **`[PR] fix: Update Dockerfile to use native installer (Open)`** (`#33443`)
    - **链接**: [PR #33443](https://github.com/anthropics/claude-code/pull/33443)
    - **重要性**: 针对 DevContainer 安装的优化，移除了已废弃的 npm 安装方式，改用原生安装器，保持了环境的一致性和现代化。

5.  **`[PR] add the missing source to claude code (Open)`** (`#41611`)
    - **链接**: [PR #41611](https://github.com/anthropics/claude-code/pull/41611)
    - **重要性**: 与 `#41447` 类似，是社区对开源贡献的一部分，意图补充遗漏的源代码。

6.  **`[PR] docs: use standard GitHub capitalization in README (Closed)`** (`#60427`)
    - **链接**: [PR #60427](https://github.com/anthropics/claude-code/pull/60427)
    - **重要性**: 提升文档质量的常规贡献，体现了社区对项目细节的关注。

7.  **`[PR] docs: polish plugins README wording (Closed)`** (`#60732`)
    - **链接**: [PR #60732](https://github.com/anthropics/claude-code/pull/60732)
    - **重要性**: 对插件生态的描述进行了润色，使文档更易读、更专业。

---

## 功能需求趋势

1.  **代理与多任务管理**: 社区对 Claude 的“代理”执行模式提出了更高要求。代表 Issue `#50246`（消息队列模式）和 `#68998`（队列命令）均获得了大量关注。用户希望在不中断当前任务的前提下，能够提前规划或暂存后续指令，形成任务流水线。
2.  **IDE 与平台集成深化**: 问题 `#25128`（VS Code 拖拽）和 `#69241`（JetBrains 自动接受编辑）显示出用户不再满足于基础集成，而是追求与 IDE 更深度、更流畅的交互体验。他们期望 Claude Code 的行为能更贴近原生的 IDE 操作模式。
3.  **MCP 生态与外部系统连接**: 用户希望提升 MCP 的可靠性和易用性。Issue `#26094`（MCP 参数序列化）的长期存在影响了其信任度。同时，`#60327`（查询外部设计系统）和 `#69205`（远程 MCP OAuth）则指向更复杂的异构系统集成场景。
4.  **后台活动可见性**: `#67485` 是一个重要的信号，表明随着“多代理”模式的推广，用户希望能够像看“仪表盘”一样，实时了解后台子代理的状态、进度和结果，而不仅仅是看到主代理在“思考”。
5.  **平台兼容性（非主流架构）**: 从 `#39636` (ARM64 Windows) 和 `#5277` (SSH 远程开发) 可以看出，Claude Code 的用户群体正在扩展至使用非主流平台和远程开发环境的开发者，他们对这些场景下的基础功能（如图像粘贴）有强烈需求。

---

## 开发者关注点

- **计费透明性**: `#16157` 是当前压倒性的第一关注点。付费用户对订阅后仍被限额感到非常沮丧，这直接侵蚀了对产品的信任。官方急需清晰的定义和修复。
- **连接与稳定性**: `#34255`（远程控制重连失败）和 `#69062`（代理任务无法启动）表明，交互的可靠性是用户的核心底线。开发者在工作中最怕进程“无响应”或“静默失败”。
- **主机与工作流一致性**: `#65767` 指出 UI 与底层工作树状态不同步的问题，这不仅是体验问题，更可能导致误操作。开发者在自动化工作流中要求更严格的上下文管理。
- **成本预测**: `#64701` 显示用户会主动监控工具使用的成本。当 CLI 内置的价格地图与实际 API 计费不一致时，他们会立刻发现并报告。精确的成本计算是开发者的一个隐性需求。
- **性能焦虑**: `#68931` 指出在空闲状态下 CPU 占用 100% 的问题，这表明用户对资源消耗非常敏感。一个高效的开发工具应该在空闲时保持静默，避免无谓的性能浪费。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报｜2026-06-18

---

## 📋 今日速览

- **连续发布三个 Rust Alpha 版本**：0.141.0-alpha.5/6/7，显示核心组件加速迭代。
- **认证与数据库问题成社区焦点**：多起账号被强制要求遗留手机号验证、更新后数据库损坏导致应用无法启动，用户抱怨强烈。
- **性能与资源占用回归报告增多**：macOS 上 `syspolicyd` 与 `trustd` 进程 CPU/内存异常、Windows 下 Crashpad 无限制生成 dump 文件（日增 5GB+），引发开发者关注。

---

## 🚀 版本发布

过去 24 小时内，Codex 发布了三个 Rust 核心库的 Alpha 版本：

- **rust-v0.141.0-alpha.5**
- **rust-v0.141.0-alpha.6**
- **rust-v0.141.0-alpha.7**

目前没有附带详细发布说明，但从连续版本号推断，团队正在进行密集的 bug 修复或底层重构。建议关注后续的 Release Notes。

---

## 🔥 社区热点 Issues（精选 10 条）

### 1. [#23794 Codex Desktop 不再显示上下文/Token 用量指示器（已关闭）](https://github.com/openai/codex/issues/23794)
- **评论 170 | 👍 168**
- **重要性**：社区高度关注的核心 UX 问题。更新后桌面应用彻底丢失了 Token 使用量可视化，用户无法判断上下文是否即将耗尽，严重影响 coding 流程。大量用户要求紧急回滚或修复。

### 2. [#25749 无法访问的遗留手机号导致账号被锁（OPEN）](https://github.com/openai/codex/issues/25749)
- **评论 49 | 👍 30**
- **重要性**：即使已通过 Google OAuth 和 MFA 登录，Codex 仍要求对已废弃的旧手机号进行短信验证，且无任何恢复或更换途径。该问题影响大量已绑定旧号码的用户，属于严重的认证设计缺陷。

### 3. [#25719 macOS 上 Codex Desktop 反复触发 syspolicyd / trustd CPU 内存暴增（OPEN）](https://github.com/openai/codex/issues/25719)
- **评论 31 | 👍 39**
- **重要性**：应用持续触发系统安全服务占用大量 CPU 和内存，导致笔记本发热、风扇狂转。macOS 用户群体大，该问题具有广泛代表性。

### 4. [#17827 可自定义的状态栏（OPEN，Feature Request）](https://github.com/openai/codex/issues/17827)
- **评论 16 | 👍 71**
- **重要性**：社区呼声最高的功能请求之一。用户希望像 Claude Code 一样，在 TUI 底部展示自定义的 Token 用量、模型名称、速率限制、Git 分支等实时信息。👍 71 反映强烈需求。

### 5. [#21211 线程导航/加载因元数据膨胀与历史数据预加载变慢（OPEN）](https://github.com/openai/codex/issues/21211)
- **评论 12 | 👍 2**
- **重要性**：社区专家指出 SQLite 中无限制的线程元数据导致导航性能急剧下降。这是影响老用户多线程工作流的潜在性能陷阱。

### 6. [#24006 macOS 更新后 Codex 无法访问本地数据库（OPEN）](https://github.com/openai/codex/issues/24006)
- **评论 11 | 👍 9**
- **重要性**：类似错误频繁出现（参见 #24030、#28666），应用更新后数据库损坏或路径丢失导致完全无法启动，严重影响日常工作。

### 7. [#25737 Codex CLI 登录强制要求 SMS OTP，忽略硬件密钥（OPEN）](https://github.com/openai/codex/issues/25737)
- **评论 11 | 👍 6**
- **重要性**：使用 FIDO2 安全密钥完成主认证后，CLI 仍跳转到手机 OTP 页面，而浏览器登录可正常使用 AAS（高级账户安全）。迫使注重安全的用户只能使用 Web 端，降低 CLI 可用性。

### 8. [#25178 Windows Computer Use 截图在 Windows 10 22H2 上失败（OPEN）](https://github.com/openai/codex/issues/25178)
- **评论 11 | 👍 4**
- **重要性**：核心 Computer Use 功能无法在 Windows 10 上获取截图，报错“SetIsBorderRequired failed: 不支持此接口”。影响大量仍使用 Win10 的企业用户。

### 9. [#25921 Codex Desktop 无限制生成 Crashpad dump 文件（日增 5GB+）（OPEN）](https://github.com/openai/codex/issues/25921)
- **评论 9 | 👍 2**
- **重要性**：Crashpad 的 `pending` 目录在一天内产生 5.4 万文件、4.9GB 数据。极度浪费磁盘空间并可能触发 macOS 的磁盘配额报警，属于严重资源管理 bug。

### 10. [#88823 5 小时使用额度消耗速度远快于历史同期（OPEN，今日新增）](https://github.com/openai/codex/issues/28823)
- **评论 4 | 👍 0**
- **重要性**：用户发现本地 telemetry 与实际使用额度计数器严重不符，同一个会话消耗比例异常高，可能意味着配额计算出现了回归。今日刚创建，值得关注后续讨论。

---

## 📌 重要 PR 进展（精选 10 条）

### 1. [#28838 支持 Codex Home 指令目录（OPEN）](https://github.com/openai/codex/pull/28838)
- **功能**：允许在 `~/.codex/instructions/` 中放置 `*.md` 文件作为全局指令，按排序顺序追加，保持与 `AGENTS.md` 的继承关系。
- **价值**：极大简化团队/个人跨项目配置的复用，属于开发者体验的显著提升。

### 2. [#28836 支持助手实时追加文本（OPEN）](https://github.com/openai/codex/pull/28836)
- **功能**：让 App Server 的 `thread/realtime/appendText` API 能够携带 `assistant` 角色，支持前端实时语音连续性场景。
- **价值**：完善实时语音对话的上下文拼接，是 AI Voice 交互的重要基建。

### 3. [#28835 添加 App-Server 当前时间提供者（OPEN）](https://github.com/openai/codex/pull/28835)
- **功能**：新增 `currentTime/read` 协议，支持会话配置 `current_time_reminder` 特性，实现时间感知的上下文注入。
- **价值**：为后续“变量延迟（varlatency）”特性铺路，模型能感知当前时间以做出更准确的回答。

### 4. [#28813 在 Esc 中断前暂停活动目标（OPEN）](https://github.com/openai/codex/pull/28813)
- **功能**：修复当用户按 Esc 中断当前回合时，活跃的 `/goal` 未进入暂停状态的问题。
- **价值**：确保 /goal 状态机在中断后正确持久化，避免目标状态混乱。

### 5. [#28814 记录历史时分配响应项 ID（OPEN）](https://github.com/openai/codex/pull/28814)
- **功能**：为客户端创建的响应项在写入历史时赋予唯一 ID，同时保留服务端返回的 ID 不变。
- **价值**：解决因缺少 ID 导致历史回放与恢复时数据丢失的问题，提升对话连续性。

### 6. [#28824 基于系统时钟的时间提醒实现（OPEN）](https://github.com/openai/codex/pull/28824)
- **功能**：在核心层添加可注入的当前时间提供者（内置系统实现），在模型请求前将 UTC 时间开发者提醒写入历史。
- **价值**：进一步推进“变量延迟”特性，让模型能够基于当前时间调整响应策略。

### 7. [#27986 控制实时语音自动转交的前缀（已关闭，已合并）](https://github.com/openai/codex/pull/27986)
- **功能**：允许在 `thread/realtime/start` 中配置 `codexResponseHandoffPrefix`，仅对自动 V1 评论添加前缀，最终回答保持原样。
- **价值**：细粒度控制实时语音对话风格，满足个性化需求。

### 8. [#27132 在工具调用项中发射受信任的 MCP 应用身份（OPEN）](https://github.com/openai/codex/pull/27132)
- **功能**：向 MCP 工具调用项添加可选的 `appContext`，包含 `connectorId`、`linkId`、`mcpAppResourceUri` 等元数据，并在历史与恢复时保持。
- **价值**：为 MCP 生态提供标准化身份追踪，利于安全审计与上下文回溯。

### 9. [#28826 为实时路由回合使用唯一 ID（已关闭，已合并）](https://github.com/openai/codex/pull/28826)
- **功能**：修正实时语音交接时使用会话本地计数器作为 turn ID 的 bug，改为全局唯一 ID，避免恢复后 ID 重复。
- **价值**：修复分布式场景下回话 ID 冲突，保障语音录制与回放的正确性。

### 10. [#28784 修复 install.sh 在旧版 awk 上的校验和解析（已关闭，已合并）](https://github.com/openai/codex/pull/28784)
- **功能**：将 `mawk` 不支持的 interval 表达式替换为兼容语法，使 CLI 安装程序能在 Debian 等老系统上正常工作。
- **价值**：解决 Linux 用户因系统 awk 版本过新而无法安装的问题，提升跨平台兼容性。

---

## 📊 功能需求趋势

从 Issues 与 PR 中可总结出社区当前最关注的功能方向：

1. **上下文与用量可视化增强**（#17827、#23794）  
   用户强烈要求恢复/提供 Token 用量、模型名称、速率限制等实时指标，且支持自定义布局。

2. **认证流程优化**（#25749、#25737、#28672）  
   申诉被迫使用遗留手机号验证、CLI 应支持硬件安全密钥（Passkey/FIDO2）、Business 账号 401 频繁退出。

3. **性能与资源管理**（#25719、#25921、#21211）  
   macOS 下系统进程异常、Crashpad 无限制 dump 生成、线程列表加载缓慢，成为开发者普遍痛点。

4. **计算机使用（Computer Use）稳定性**（#25178、#24207、#28834）  
   Windows 10 截图失败、Intel Mac 锁屏功能不可用、更新后插件初始化失败，影响自动化流程。

5. **时间感知与变量延迟**（#28835、#28824）  
   允许模型感知当前时间，逐步接近“上下文能自动带时间戳”的需求，是提升模型回答准确性的新方向。

6. **MCP 生态成熟化**（#27132、#27500）  
   标准化 MCP 工具调用的身份、支持更多扩展表单类型，表明 Codex 正积极拥抱开放工具协议。

7. **CLI/TUI 体验优化**（#17827、#28532、#28482）  
   自定义状态行、深色主题兼容性、/goal 中断与暂停状态正确性，开发者期望 TUI 能对标行业最佳实践。

---

## 🔧 开发者关注点

- **数据库损坏与启动失败**（#24006、#24030、#28666）  
  macOS 更新后 `state_5.sqlite` 损坏、Windows 上 SQLite “database is locked”，导致应用完全无法启动。开发者呼吁 OpenAI 提供修复脚本或自动容错机制。

- **配额与使用计量不一致**（#28823、#28837、#28827）  
  多个用户报告 5 小时使用额度消耗速度异常快，且 CLI 缺少额度重置功能。Pro 用户担心计费逻辑回归。

- **认证强迫依赖手机短信**（#25749、#25737）  
  已启用 MFA 或 Passkey 的账号仍被要求 SMS 验证，且无法更换或跳过。开发者指出这是安全透明度和自动化 CI/CD 场景的重大障碍。

- **频繁使用的“一次性”错误**（#8190 #23794）  
  上下文窗口溢出提示生硬、缺少 Token 用量指示器，导致用户在复杂任务中频繁中断。长期存在的 #8190 仍无解决。

- **跨平台兼容性问题**（#25178 #24207 #28262 #28532）  
  Intel Mac、Windows 10 22H2、含韩文用户名的 Windows 环境、SSH 到 Windows 的 Light 主题渲染均出现独立 bug。开发者希望加强 QA 覆盖。

---

> **今日提醒**：如果你遇到了数据库损坏或启动失败，可以尝试清理 `~/Library/Application Support/com.openai.codex/` 下的缓存（注意备份），但具体解决方案仍需等待官方修复。认证问题可暂时通过浏览器端使用 Codex Web 或 ChatGPT 绕过。

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI 社区动态日报 | 2026-06-18

## 📌 今日速览

- 连续发布 **v0.47.0 稳定版** 和 **v0.48.0-preview.0 预览版**，主要包含后端定义优化及 npm 依赖更新冷却机制。
- 社区持续聚焦 **Agent 稳定性**：Generalist 代理挂死、子代理成功率误报、浏览器代理 Wayland 兼容等问题仍是热门。
- 安全方面取得重要进展：多个 CI/CD 权限提升风险被识别并修复，同时引入 `trustedFolders.json` 列表格式支持。

---

## 📦 版本发布

### [v0.47.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.47.0)
- 更新日志自动生成，版本号里程碑。
- 核心变更：`Respect backend def` —— 优化后端定义识别逻辑，提升配置匹配准确性。

### [v0.48.0-preview.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.48.0-preview.0)
- 引入 **dependabot 冷却期**，npm 包更新将遵循 14 天冷却策略，减少频繁升级带来的风险。
- 包含 `ref` 等基础调整，为正式版做准备。

---

## 🔥 社区热点 Issues

### 1. [#24353 Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353) ⭐ 7 评论
- **重要性**：组件级评测 EPIC，已生成 76 个 behavioral eval 测试，覆盖 6 个 Gemini 模型。是未来质量保证的基石。
- **社区反应**：持续跟踪中，多名贡献者参与讨论测试框架扩展。

### 2. [#21409 Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409) ⭐ 8 👍 / 7 评论
- **重要性**：Generalist 代理创建文件夹等简单操作即导致永久挂起，是用户反馈最多的稳定性问题之一。
- **社区反应**：8 个 👍 表明影响面广；临时解决方案是手动禁止使用子代理。

### 3. [#22745 Assess AST-aware file reads, search, mapping](https://github.com/google-gemini/gemini-cli/issues/22745) ⭐ 1 👍 / 7 评论
- **重要性**：探索 AST 感知的文件读取/搜索/代码映射，可减少 token 冗余、提升工具调用精度。
- **社区反应**：已启动多项调研子 issue，属于长期性能改进方向。

### 4. [#22323 Subagent recovery after MAX_TURNS reported as GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323) ⭐ 2 👍 / 6 评论
- **重要性**：子代理在达到最大轮次后仍报告成功状态，导致用户误以为任务完成，实为中断。
- **社区反应**：需重新设计终止逻辑，避免状态欺骗。

### 5. [#25166 Shell command execution stuck with “Waiting input”](https://github.com/google-gemini/gemini-cli/issues/25166) ⭐ 3 👍 / 4 评论
- **重要性**：执行简单命令行后（如 `ls`）仍显示“等待输入”，导致会话卡死。
- **社区反应**：标记为 effort/medium，多位用户复现。

### 6. [#21983 browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983) ⭐ 1 👍 / 4 评论
- **重要性**：Wayland 下浏览器代理无法启动或失败，Linux 用户受影响。
- **社区反应**：需要适配 Wayland 显示服务，等待重现测试。

### 7. [#26525 Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525) ⭐ 0 👍 / 5 评论
- **重要性**：Auto Memory 在读取本地对话时可能暴露秘钥到模型上下文，安全性隐患。
- **社区反应**：安全团队主导，要求前置确定性脱敏。

### 8. [#20079 Symlink agents not recognized](https://github.com/google-gemini/gemini-cli/issues/20079) ⭐ 0 👍 / 4 评论
- **重要性**：`~/.gemini/agents/filename.md` 如果是指向其他位置的软链接，会被忽略。
- **社区反应**：简单但烦人的配置问题，期望支持 symlink。

### 9. [#22672 Agent should stop/discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672) ⭐ 1 👍 / 3 评论
- **重要性**：模型在 git 操作中使用 `--force` 等危险命令，需要增加保护机制。
- **社区反应**：社区呼吁更保守的默认行为，避免数据丢失。

### 10. [#21968 Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968) ⭐ 0 👍 / 6 评论
- **重要性**：自定义技能和子代理被模型主动调用的频率过低，用户需手动提示才能触发。
- **社区反应**：核心 agent 决策问题，影响扩展性。

---

## 🚀 重要 PR 进展

### 1. [#28000 fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file](https://github.com/google-gemini/gemini-cli/pull/28000) (OPEN)
- **功能**：修复 `write_file` 写入 `.ipynb` 和普通 JSON 文件时破坏格式的严重 bug，恢复 Colab/JupyterLab 中的兼容性。
- **影响**：高频使用的工具，修复及时。

### 2. [#27996 fix(core): decode response body using charset from Content-Type header](https://github.com/google-gemini/gemini-cli/pull/27996) (OPEN)
- **功能**：`web-fetch` 现在根据 HTTP 响应头的 charset 解码，解决中文、日文等站点乱码问题。
- **社区价值**：提升多语言抓取准确性。

### 3. [#27994 fix(core): insert skill/agent content literally in system prompt substitutions](https://github.com/google-gemini/gemini-cli/pull/27994) (OPEN)
- **功能**：解决 `applySubstitutions()` 中使用 `String.prototype.replace` 字符串形式导致的特殊字符被误解释的问题。现在技能/子代理内容将原样嵌入。
- **重要性**：提高提示词可靠性。

### 4. [#28002 Changelog for v0.47.0](https://github.com/google-gemini/gemini-cli/pull/28002) (OPEN)
- **功能**：自动生成的发布日志，请审核合并。

### 5. [#27648 feat(core): support list format in trustedFolders.json](https://github.com/google-gemini/gemini-cli/pull/27648) (CLOSED)
- **功能**：允许 `trustedFolders.json` 使用 JSON 数组格式，简化手动维护信任目录列表。
- **影响**：降低用户配置门槛。

### 6. [#27990 test(core-tools): resolve macOS symlink path mismatches in tests](https://github.com/google-gemini/gemini-cli/pull/27990) (OPEN)
- **功能**：修复 macOS 上 `/var` -> `/private/var` 软链接导致的测试失败，提升跨平台稳定性。

### 7. [#27987 fix(cli): throw FatalConfigError instead of process.exit in parseArguments](https://github.com/google-gemini/gemini-cli/pull/27987) (OPEN)
- **功能**：重构参数解析，使用异常替代 `process.exit(1)`，改善测试和错误处理一致性。

### 8. [#27997 docs: remove references to deprecated consumer and free tiers](https://github.com/google-gemini/gemini-cli/pull/27997) (OPEN)
- **功能**：文档清理，移除已下线免费层级的引用，更新服务条款。

### 9. [#27780 security: gate chained E2E on same-repository checkout for workflow_run](https://github.com/google-gemini/gemini-cli/pull/27780) (OPEN)
- **功能**：防止 fork PR 通过 artifact 元数据注入恶意仓库路径，避免 `GEMINI_API_KEY` 泄露。
- **安全影响**：属于 CI/CD 安全加固，阻止供应链攻击。

### 10. [#27783 security: gate PRT label workflows on same-repository pull_request_target](https://github.com/google-gemini/gemini-cli/pull/27783) (OPEN)
- **功能**：为 `pr-size-labeler` 和 `pr-rate-limiter` 添加同仓库守卫，避免外部 fork 滥用标签修改权限。

---

## 📊 功能需求趋势

从本期 Issues 可归纳出社区最关注的四个方向：

1. **Agent 稳定性和智能水平**  
   - 子代理状态误报（#22323）、代理不主动使用技能（#21968）、危险操作防护（#22672）等。
   - 社区期望更少的人工干预，更可靠的自主决策。

2. **评估与质量体系**  
   - 组件级评测（#24353）、behavioral eval 扩展、AST 感知工具评估（#22745）。
   - 强调可重复、可衡量的回归测试，确保每次改版不退步。

3. **安全与隐私**  
   - Auto Memory 脱敏（#26525）、CI/CD 权限加固（#27780、#27783）、symlink 信任问题（#20079）。
   - 用户对模型操作敏感数据（如 shell、文件）的担忧日益增长。

4. **编辑器/终端体验**  
   - 终端 resize 闪烁（#21924）、外部编辑器退出后画面错乱（#24935）。
   - 核心体验优化，影响日常使用效率。

---

## 🧩 开发者关注点

| 痛点/需求 | 相关 Issue/PR | 说明 |
|-----------|--------------|------|
| **代理挂死** | #21409、#25166、#22465 | Generalist 代理、shell 命令、交互式提示均可能无限等待 |
| **子代理行为不可控** | #22323、#22093、#21983 | 状态误报、越权执行、Wayland 兼容 |
| **文件写入损坏** | #28000 | .ipynb/JSON 文件被破坏，影响数据持久化 |
| **配置不生效** | #20079、#22267 | symlink 代理文件不被识别、browser agent 忽略 settings.json |
| **依赖更新风险** | #27788、#27948 | 频繁更新导致测试不稳定，社区推动 14 天冷却期 |
| **安全敏感操作缺乏防护** | #22672、#26525 | 模型使用 `--force`、内存日志泄露秘钥 |
| **测试与 CI 稳定性** | #28000 测试修复、#27780/27783 安全守卫 | 跨平台（macOS 软链接）和 fork PR 安全 |

---

*以上动态基于 `google-gemini/gemini-cli` 公开仓库数据，统计截止 2026-06-18 24:00 UTC。*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-06-18

## 📌 今日速览
- 今日无新版本发布，但社区围绕 **6月16日 Copilot 服务中断** 引发的模型封锁问题集中反馈，相关 Issue #3832 获得 13 个 👍。
- **插件与 MCP** 成为焦点：preToolUse 钩子无法静默批准指令、子代理丢失 MCP 工具访问等问题持续发酵。
- 企业级用户对 **自定义模型支持**（#3730）和 **上下文窗口限制**（#3355）的需求日益强烈。

---

## 📦 版本发布
无

---

## 🔥 社区热点 Issues（Top 10）

### 1. [#2643] preToolUse 钩子：即使 permissionDecision: allow 仍弹出确认对话框  
- **作者**: jeziellopes  
- **标签**: `area:plugins`  
- **评论**: 10 | 👍: 1  
- **核心**: 插件开发者在 `preToolUse` 中通过 `updatedInput` 静默改写指令，并设置 `permissionDecision: allow`，但 CLI v1.0.24 依然每次弹出确认框，导致无法实现完全静默的指令重写。  
- **链接**: [Issue #2643](https://github.com/github/copilot-cli/issues/2643)

### 2. [#1973] 功能请求：交互模式下设置工具白名单  
- **作者**: Dicer-J  
- **标签**: `area:permissions`, `area:configuration`  
- **评论**: 10 | 👍: 20  
- **核心**: 交互模式中每次工具调用（如 grep、cat）都需要手动批准，而 `/allow-all` 又太危险。用户希望为只读工具设置白名单，减少审批频繁度。  
- **链接**: [Issue #1973](https://github.com/github/copilot-cli/issues/1973)

### 3. [#254] copilot-cli 反复要求重新登录  
- **作者**: yurivict  
- **标签**: `more-info-needed`, `unable-to-reproduce`  
- **评论**: 9 | 👍: 4  
- **核心**: 长期存在的登录态丢失问题，即使已成功登录，新会话依然提示需要再次登录。影响 GitHub Business 账户用户。  
- **链接**: [Issue #254](https://github.com/github/copilot-cli/issues/254)

### 4. [#3560] 执行失败：CAPIError 400 – 重复 item id  
- **作者**: lgphp  
- **标签**: `area:models`, `area:tools`  
- **评论**: 5 | 👍: 1  
- **核心**: 同一工作流在傍晚突然出现 `Duplicate item found with id` 错误，仅发生在工具函数调用后的下一轮对话，普通聊天正常。可能为服务端临时 bug。  
- **链接**: [Issue #3560](https://github.com/github/copilot-cli/issues/3560)

### 5. [#3832] [已关闭] 所有模型显示为“Blocked/Disabled” – 6月16日中断后遗症  
- **作者**: yzeng58  
- **标签**: `bug`  
- **评论**: 5 | 👍: 13  
- **核心**: 6月16日 Copilot 服务中断后，模型选择界面所有模型被标记为“阻塞/禁用”，用户无法选择任何模型启动新会话。已关闭，但社区高度关注此事件的后续影响。  
- **链接**: [Issue #3832](https://github.com/github/copilot-cli/issues/3832)

### 6. [#3355] 允许配置 Claude Opus 4.6 的上下文窗口（当前仅 200K，模型支持 1M）  
- **作者**: aksingh  
- **标签**: `area:context-memory`, `area:models`  
- **评论**: 3 | 👍: 4  
- **核心**: CLI 对 Claude Opus 4.6 仅开放 200K tokens，而模型原生支持 1M，导致深度技术会话中频繁触发自动摘要，影响连贯性。  
- **链接**: [Issue #3355](https://github.com/github/copilot-cli/issues/3355)

### 7. [#3730] 支持企业管理的自定义模型（Enterprise-Managed Custom Models）  
- **作者**: sebdanielsson  
- **标签**: `area:enterprise`, `area:models`  
- **评论**: 2 | 👍: 4  
- **核心**: 企业管理员在 Copilot 管理面板配置的自定义模型在 VS Code 中可用，但 Copilot CLI 无法识别，期望 CLI 也能使用企业自定义端点。  
- **链接**: [Issue #3730](https://github.com/github/copilot-cli/issues/3730)

### 8. [#3754] `copilot --resume "含空格名称"` 静默失败，退出码 1  
- **作者**: JoeBrockhaus  
- **标签**: `area:sessions`  
- **评论**: 2 | 👍: 1  
- **核心**: 使用 `--resume` 恢复含空格的会话名称时，CLI 直接退出且无任何输出，与文档描述矛盾。  
- **链接**: [Issue #3754](https://github.com/github/copilot-cli/issues/3754)

### 9. [#3812] 子代理无法访问 MCP 工具  
- **作者**: SpeedJack  
- **标签**: `area:agents`, `area:mcp`  
- **评论**: 2 | 👍: 0  
- **核心**: 自定义子代理无法感知和使用 MCP 工具，尽管顶层代理可以。推测与 MCP 工具的延迟加载机制有关。  
- **链接**: [Issue #3812](https://github.com/github/copilot-cli/issues/3812)

### 10. [#3839] Ollama Cloud 不支持 Copilot CLI 发送的 `custom_tool_call` 载荷  
- **作者**: weweaaa  
- **标签**: `triage`  
- **评论**: 1 | 👍: 7  
- **核心**: 在 Fleet Mode 中使用 BYOK 模型并通过 Ollama Cloud 路由时，因 `custom_tool_call` 字段不被 Ollama 解析导致 400 错误。  
- **链接**: [Issue #3839](https://github.com/github/copilot-cli/issues/3839)

---

## 🚧 重要 PR 进展
今日无新 Pull Request 更新。

---

## 🔭 功能需求趋势
从近期 Issue 中提炼出社区最关注的五个功能方向：

1. **插件系统增强**  
   - 静默 command rewrite（#2643）  
   - 文档匹配器（matcher）支持（#3820）  
   - 一键更新所有插件（#3830）

2. **模型与上下文扩展**  
   - 允许配置更大的上下文窗口（#3355）  
   - 企业自定义模型接入（#3730）  
   - 更灵活的 `/effort` 快速切换推理力度（#3074）

3. **权限与安全**  
   - 工具白名单替代全量审批（#1973）  
   - 内容排除策略误用于 CLI（#3841）

4. **MCP 与子代理**  
   - 子代理访问 MCP 工具（#3812）  
   - 预加载 MCP 工具而非延迟加载（#3787）  
   - 技能文件声明额外 MCP 服务器（#3292）

5. **自定义命令与别名**  
   - 用户可配置别名和自定义 `/` 命令（#3844）  
   - 持久化 `/instructions` 设置（#3840）

---

## 👨‍💻 开发者关注点
- **稳定与兼容性**：6月16日中断导致的模型封锁、CAPI 400 重复 ID 错误、Ollama Cloud 荷不兼容等问题，显示服务端更新可能对客户端产生冲击。  
- **子代理透明度**：子代理默认使用与主会话不同的模型（#3824），且用户无感知，开发者期望模型选择透明化。  
- **文件系统问题**：`plugin install` 因 git fsmonitor 套接字文件失败（#3842）；`ContentExclusionFilter` 崩溃（#3828）。  
- **登录与恢复**：重复登录（#254）和含空格会话恢复失败（#3754）持续困扰用户。  
- **体验细节**：鼠标交互无法禁用（#3804）、恢复会话时不显示原工作目录（#3837）以及已加密附件污染会话（#3791）等问题反映用户体验打磨仍有空间。

> 以上分析基于 GitHub Copilot CLI 公开仓库 2026-06-18 数据，由 AI 自动生成。

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 | 2026-06-18

---

## 今日速览

过去24小时内项目无新版本发布，社区活跃度较低，主要围绕两项新提交的 feature request 展开：支持会话运行时动态切换执行模式（Agent ↔ 集群），以及增加忽略 SSL 证书的选项。后者反映了企业安全软件带来的实际痛点，值得关注。

---

## 版本发布

无

---

## 社区热点 Issues

**1. #2459 | [Feature Request] 支持会话运行中切换执行模式（Agent ↔ 集群）**  
- **作者**: PresentXoX  
- **状态**: OPEN | 评论: 0 | 👍: 0  
- **链接**: [MoonshotAI/kimi-cli Issue #2459](https://github.com/MoonshotAI/kimi-cli/issues/2459)  
- **重要性与分析**:  
  该需求提出在同一个会话中，无需中断即可在 **Agent 模式** 与 **集群模式** 之间切换。当前用户必须在启动时决定运行模式，而实际工作流中常常需要根据任务复杂度动态调整并行能力或资源分配。若能实现，将显著提升使用灵活性，尤其适合长时间运行的分析流水线。目前尚未引起社区讨论，但方向击中多模式轮换的痛点，值得项目组优先评估。

**2. #2458 | [enhancement] Add option to ignore ssl certificate**  
- **作者**: dmorsin  
- **状态**: OPEN | 评论: 0 | 👍: 0  
- **链接**: [MoonshotAI/kimi-cli Issue #2458](https://github.com/MoonshotAI/kimi-cli/issues/2458)  
- **重要性与分析**:  
  用户在受企业管控的 PC 上使用杀毒软件，该软件通过 **中间人（MiTM）** 方式注入自己的证书，导致 Kimi CLI 在登录时因证书不匹配而失败。当前 CLI 没有提供 `--insecure` 或 `--ignore-ssl` 等选项。这一需求直接影响了企业环境用户的可用性，且涉及安全与合规权衡，需要官方提供可控的绕过机制（如环境变量或命令行标志）。社区虽暂无讨论，但此类问题在组织级部署中非常常见，建议尽快支持。

---

## 重要 PR 进展

无

---

## 功能需求趋势

基于今日仅有的两个 Issues，社区最关注的 functional 方向如下：

| 方向 | 说明 |
|------|------|
| **会话/运行态模式切换** | 用户希望在同一个 CLI 会话中动态切换 Agent 模式与集群模式，避免重启或新建会话带来的上下文丢失。 |
| **企业网络/安全适配** | 增加 `--insecure` 标志或类似机制，允许在受企业反病毒/MiTM 环境中跳过 SSL 验证，解决登录失败问题。 |

两个需求分别指向 **运行态灵活性** 和 **部署环境兼容性**，表明社区用户已从单纯的功能使用进入更精细的场景适配阶段。

---

## 开发者关注点

- **企业环境适配成为强需求**：来自 `dmorsin` 的 SSL 证书问题并非个例，企业级用户的终端上常安装有管控软件（如杀毒、DLP），这类软件以 MITM 方式劫持流量，导致命令行工具无法正常认证。开发者应优先考虑提供 `--insecure` 或 `NODE_TLS_REJECT_UNAUTHORIZED=0` 之类的方式，并做好安全警告。
- **模式切换的成本感知**：`#2459` 虽然今天只有 1 人提出，但“在运行中切换模式”是许多多步骤任务（例如先用 Agent 快速原型，再切到集群做大规模处理）的自然诉求。社区可能希望 CLI 能像 Kubernetes 的 `kubectl config use-context` 一样轻量，而非重启一个全新的过程。

---

*数据来源：GitHub MoonshotAI/kimi-cli，截至 2026-06-18 08:00 UTC*  
*日报由 AI 自动生成，仅供参考。*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-06-18

## 今日速览
- 发布 v1.17.8，主要修复 OpenAI 兼容提供商的 MCP 工具 schema 校验问题，并优化会话时间线加载性能。
- 社区围绕多代理协调（#17994）、TPS 显示（#6096）和 Windows 终端兼容性（#3541、#24817）讨论激烈；多个 PR 带来模糊编辑、模型自动发现等新功能。
- 开发者对 Windows 下的 ANSI 乱码、崩溃后终端残留以及本地模型（Ollama）支持不足的抱怨较为集中。

## 版本发布
### v1.17.8
- **Core 改进**：会话时间线加载速度大幅提升，消除闪烁和滚动跳跃。
- **Bug 修复**：
  - OpenAI 兼容提供商现在能够正确接受之前验证失败的 MCP 工具 schema（感谢 @jquense）。
  - Cloudflare AI Gateway 现在能正确接收配置的 API key（感谢 @keefetang）。

## 社区热点 Issues（按评论数排序）

1. **#17994 [FEATURE]：支持隔离工作区中的多代理编排**  
   评论 21 | 👍 2  
   希望内置运行“团队”代码代理的能力，类似于其他工具的隔离沙箱。社区对此功能期待已久，但讨论热度与点赞数有差距，可能需更多设计细节。  
   [https://github.com/anomalyco/opencode/issues/17994](https://github.com/anomalyco/opencode/issues/17994)

2. **#6096 [FEATURE]：增加实验性的每秒 Token 计算与显示 (TPS)**  
   评论 18 | 👍 55  
   呼声最高的功能请求之一，用户希望在每条消息回复中看到 TPS，以评估模型性能。社区普遍支持，有 55 个赞。  
   [https://github.com/anomalyco/opencode/issues/6096](https://github.com/anomalyco/opencode/issues/6096)

3. **#23566 文档称 LSP 默认启用，但与实际行为不符**  
   评论 10 | 👍 20  
   文档指出“自动安装 Kotlin 项目的 LSP”，但实际 LSP 被默认禁用。用户认为文档误导，要求同步或修复默认行为。  
   [https://github.com/anomalyco/opencode/issues/23566](https://github.com/anomalyco/opencode/issues/23566)

4. **#32172 [FEATURE]：为 Z.AI 提供商添加 GLM-5.2 模型支持**  
   评论 10 | 👍 0  
   Z.AI 发布了最新推理模型 GLM-5.2，用户请求原生支持。虽然赞数少，但评论数多说明有兴趣。  
   [https://github.com/anomalyco/opencode/issues/32172](https://github.com/anomalyco/opencode/issues/32172)

5. **#3541 [Bug, Windows]：新界面导致屏幕错乱**  
   评论 9 | 👍 0  
   用户在 Windows 下通过 WSL 使用，遇到频繁屏幕闪烁/乱码问题。虽然赞少，但代表了 Windows 用户的典型痛点。  
   [https://github.com/anomalyco/opencode/issues/3541](https://github.com/anomalyco/opencode/issues/3541)

6. **#8538 Session 查找失败：从非 Git 目录启动 PTY 时抛出 NotFoundError**  
   评论 7 | 👍 0  
   在嵌套 PTY 中执行操作时 session 查找失败，影响使用多会话的高级用户。  
   [https://github.com/anomalyco/opencode/issues/8538](https://github.com/anomalyco/opencode/issues/8538)

7. **#31204 [BUG]：agent 切换时 session_message.seq 违反 NOT NULL 约束**  
   评论 7 | 👍 3  
   最近迁移后，任何触发代理切换的会话都会崩溃。严重阻碍多代理工作流。  
   [https://github.com/anomalyco/opencode/issues/31204](https://github.com/anomalyco/opencode/issues/31204)

8. **#24817 [Linux] Ctrl+Z 关闭/挂起 OpenCode 而非撤销输入**  
   评论 5 | 👍 2  
   Linux 上 Ctrl+Z 发送 SIGTSTP 而不是 undo，与常见终端行为冲突。  
   [https://github.com/anomalyco/opencode/issues/24817](https://github.com/anomalyco/opencode/issues/24817)

9. **#7928 [Feature]：运行时权限模式切换（类似 Claude Code 的 Shift+Tab）**  
   评论 5 | 👍 17  
   当前默认 auto-edit 过于激进，用户希望能在运行中动态切换权限模式。17 个赞表明广泛支持。  
   [https://github.com/anomalyco/opencode/issues/7928](https://github.com/anomalyco/opencode/issues/7928)

10. **#32745 桌面版连接 OpenRouter 时卡在“Authorization in progress...”**  
    评论 4 | 👍 0  
    授权流程阻塞，网络正常却无法完成连接。新近问题，影响新用户入门。  
    [https://github.com/anomalyco/opencode/issues/32745](https://github.com/anomalyco/opencode/issues/32745)

## 重要 PR 进展

1. **#32731 feat(opencode): 自动发现 OpenAI 兼容提供商的模型**  
   无需手动列出模型，openode 调用 GET /models 接口自动填充。解决长期以来的配置痛点（Closes #6231）。  
   [https://github.com/anomalyco/opencode/pull/32731](https://github.com/anomalyco/opencode/pull/32731)

2. **#32761 feat(core): 将 V1 模糊编辑策略移植到 V2 核心编辑工具**  
   移植 9 种模糊替换策略 + Levenshtein 距离，提高 LLM 编辑成功率。大幅减少因空白/缩进差异导致的“oldString not found”。  
   [https://github.com/anomalyco/opencode/pull/32761](https://github.com/anomalyco/opencode/pull/32761)

3. **#32767 fix(tui): 恢复委托子代理会话的 ESC 中断功能**  
   修复回归，ESC 键再次可用于中断子代理任务（Closes #3699, #4073）。  
   [https://github.com/anomalyco/opencode/pull/32767](https://github.com/anomalyco/opencode/pull/32767)

4. **#32766 feat(core): 在公共 API 层接受显式存储**  
   提取数据库层，允许测试和嵌入注入临时存储。增强可测试性。  
   [https://github.com/anomalyco/opencode/pull/32766](https://github.com/anomalyco/opencode/pull/32766)

5. **#32751 fix(acp): 在权限对话框标题中显示命令**  
   使 ACP 模式下的 shell 命令在权限请求对话框中可见，提升透明度。  
   [https://github.com/anomalyco/opencode/pull/32751](https://github.com/anomalyco/opencode/pull/32751)

6. **#32752 feat(opencode): 添加 `session select` 交互式选择器**  
   使用 @clack/prompts 列出项目范围的根会话，类似 fzf 体验。  
   [https://github.com/anomalyco/opencode/pull/32752](https://github.com/anomalyco/opencode/pull/32752)

7. **#32758 fix(opencode): 重新读取 plugin.trigger 输出以支持数组替换**  
   修复插件修改 `output.messages` 被静默丢弃的 bug。  
   [https://github.com/anomalyco/opencode/pull/32758](https://github.com/anomalyco/opencode/pull/32758)

8. **#32750 feat: 添加全局会话列表范围切换（local → project → global）**  
   新增 Ctrl+G 快捷键，在会话对话框内循环切换范围。  
   [https://github.com/anomalyco/opencode/pull/32750](https://github.com/anomalyco/opencode/pull/32750)

9. **#27554 feat(opencode): 本地 LAN 提供商发现 + 自动发现模型**  
   结合 mDNS、SSDP 等协议自动发现局域网内的 OpenAI 兼容服务。  
   [https://github.com/anomalyco/opencode/pull/27554](https://github.com/anomalyco/opencode/pull/27554)

10. **#32612 fix(codex): 从 ChatGPT 账户模型列表中排除 `-pro` 模型**  
    修复 `gpt-5.5-pro` 在 OAuth 账户下选择后请求全失败的问题。  
    [https://github.com/anomalyco/opencode/pull/32612](https://github.com/anomalyco/opencode/pull/32612)

## 功能需求趋势
- **多代理与编排**：隔离工作区内的代理团队协作需求突出（#17994），并伴随 Agent 切换时的数据库约束问题（#31204）。
- **模型与提供商扩展**：持续请求支持新模型（GLM-5.2, glm-5.2:cloud），以及自动发现本地/私有提供商（#32731, #27554）。
- **终端与平台兼容性**：Windows 和 Linux 上的 ANSI 转义码乱码、Ctrl+Z 行为异常成为高频投诉。
- **权限与安全**：运行时权限模式切换（#7928）和命令可见性（#32751）体现用户对自动编辑安全性的关注。
- **开发体验改进**：TPS 显示（#6096）、LSP 文档同步（#23566）、文件附件限制（#32732）等局部优化。

## 开发者关注点
- **Windows 环境痛点**：屏幕闪烁、ANSI 转义码乱码、崩溃后终端残留、PowerShell 识别问题（#3541, #21277, #16675, #32754）。
- **性能与稳定性**：v1.17.8 更新后部分用户反馈严重卡顿和冻结（#32746），以及 HTTP 崩溃（#32694）。
- **配置与模型连接**：OpenRouter 授权卡住（#32745）、本地 Ollama 支持不明确（#32756）、模型自动发现缺失（#32731）。
- **文档与实际行为不符**：LSP 默认启用声称与实际关闭的矛盾（#23566），影响用户信任。
- **编辑可靠性**：LLM 生成的 oldString 因空白差异导致编辑失败，推动模糊匹配改进（#32761）。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

好的，这是为您生成的 Pi 社区动态日报。

---

# Pi 社区动态日报 | 2026-06-18

## 今日速览

今日社区活跃度极高，主要集中在 **Bug 修复与基础设施优化** 上。核心修复包括：流式 Markdown 强制滚动问题得到解决，以及 Anthropic 缓存计费逻辑的修正。同时，社区对 **新提供商支持 (Azure AI Foundry)** 和 **自适应推理模型 (Thinkng Level "max")** 的需求非常强烈，相关 PR 已合并或处于开放状态。

## 社区热点 Issues

1.  **#5825: [Bug] 流式 Markdown 强制滚动至底部**
    -   **重要性**: 这是一个影响阅读体验的关键 Bug。当 AI 快速输出 Markdown 时，用户滚动查看历史内容，但页面被强制拉回底部，导致无法正常阅读，且与 `clear on shrink` 设置有关。
    -   **社区反应**: 评论数高达 12 条，是今日讨论热度最高的话题，开发者已标记为 `inprogress`。
    -   **链接**: [earendil-works/pi Issue #5825](https://github.com/earendil-works/pi/issues/5825)

2.  **#5653: [Progress] 移除对 Shrinkwrap 的依赖**
    -   **重要性**: 这是一个深层依赖管理问题。由于 `pi-ai` 和 `pi-coding-agent` 重复打包了 `pi-ai`，导致 API 提供者注册表（模块级 Map）出现隔离，引发未知行为。修复此问题对包管理稳定性和性能至关重要。
    -   **社区反应**: 11 条评论，开发者已标记为 `inprogress`，表明团队正在积极解决。
    -   **链接**: [earendil-works/pi Issue #5653](https://github.com/earendil-works/pi/issues/5653)

3.  **#3715: [Bug] 本地 LLM 流在 5 分钟后因 `bodyTimeout` 中断**
    -   **重要性**: 影响使用本地模型（如 vLLM）进行长时间推理的用户。即使设置了 `retry.provider.timeoutMs`，也无法覆盖底层 HTTP 客户端 (undici) 的 `bodyTimeout` 默认值。
    -   **社区反应**: 评论数 11，获 4 个赞，说明该问题对本地部署或长任务用户影响甚广。虽已关闭，但修复思路值得关注。
    -   **链接**: [earendil-works/pi Issue #3715](https://github.com/earendil-works/pi/issues/3715)

4.  **#534: [Bug] Linux 上配置文件位置不符合规范**
    -   **重要性**: 这是一个低优先级但影响面广的 UX 问题。Pi 在 Linux 上将配置文件直接放在 `$HOME` 而非遵循 XDG 规范，可能导致与其它工具的不兼容和目录混乱。
    -   **社区反应**: 虽然只有 9 条评论，但获得了 **20 个赞**，是今日点赞数最高的 Issue，反映了社区对遵守平台规范的强烈期望。
    -   **链接**: [earendil-works/pi Issue #534](https://github.com/earendil-works/pi/issues/534)

5.  **#5654: [功能请求] 为自定义消息添加 `excludeFromContext` 标志**
    -   **重要性**: 用户希望通过 `/status` 等命令添加的展示信息**不**被发送给 LLM，以避免浪费 Token 或引入噪声。这触及了 Agent 交互中信息展示与上下文管理的核心矛盾。
    -   **社区反应**: 7 条评论，表明开发者们对如何更精细地控制上下文内容有普遍需求。
    -   **链接**: [earendil-works/pi Issue #5654](https://github.com/earendil-works/pi/issues/5654)

6.  **#5821: 支持 Anthropic OAuth 订阅在 Agent SDK 应用中生效**
    -   **重要性**: 用户希望已经在为 Claude 订阅付费的用户，在 Pi 中通过 Agent SDK 使用时无需额外充值或配置，直接使用现有订阅。
    -   **社区反应**: 7 条评论，Anthropic 官方已确认支持此方式，社区用户正推动 Pi 实现相应集成。
    -   **链接**: [earendil-works/pi Issue #5821](https://github.com/earendil-works/pi/issues/5821)

7.  **#5830: [Bug] 树形导航器截断长条目且无法阅读**
    -   **重要性**: 严重影响重度用户的导航和操作体验。当聊天历史包含长文件名或复杂路径时，`tree` 导航器直接截断，用户无法查看完整内容，必须通过其他繁琐方式确认。
    -   **社区反应**: 4 条评论，问题清晰，被快速关闭（可能是通过 PR 修复或重新分类），体现了团队对高频 UX 问题的重视。
    -   **链接**: [earendil-works/pi Issue #5830](https://github.com/earendil-works/pi/issues/5830)

8.  **#5810: [功能请求] RPC: 暴露会话条目和树结构**
    -   **重要性**: 标志着社区开始探索**通过外部程序驱动 Pi**。用户希望通过 RPC 接口获取会话的完整条目和树形结构，以便构建自定义 UI 或自动化工作流。
    -   **社区反应**: 3 条评论，属于前瞻性的 API 扩展需求，为未来 IDE 或高级插件集成铺平道路。
    -   **链接**: [earendil-works/pi Issue #5810](https://github.com/earendil-works/pi/issues/5810)

9.  **#5862: [Bug] Codex 订阅错误：“超出配额”但 Codex CLI 正常工作**
    -   **重要性**: 对集成 OpenAI Codex 功能的用户造成直接阻塞。在 Pi 中 OAuth 认证成功后，发起对话却得到“超出配额”的错误提示，而官方 Codex CLI 却工作正常，说明 Pi 的请求或认证处理存在兼容性问题。
    -   **社区反应**: 2 条评论，属于新上报的、影响核心功能的 Bug，需紧急排查。
    -   **链接**: [earendil-works/pi Issue #5862](https://github.com/earendil-works/pi/issues/5862)

10. **#5848: [功能请求] 压缩：支持基于上下文窗口百分比的触发机制**
    -   **重要性**: 当前压缩策略基于绝对的 `reserveTokens` 值，导致在不同上下文窗口的模型上表现不一致。用户希望实现百分比触发，使压缩行为更具普适性和智能化。
    -   **社区反应**: 1 条评论，虽然评论不多，但这是一个聪明的优化提案，能显著提升所有用户的上下文管理体验，尤其对于使用不同大小模型的用户。
    -   **链接**: [earendil-works/pi Issue #5848](https://github.com/earendil-works/pi/issues/5848)

## 重要 PR 进展

1.  **#5701: [已合并] 修复(ai): 调整 minimax-m3 上下文大小**
    -   **内容**: 将 OpenRouter 上 Minimax-M3 模型的上下文大小从 1M 修正为 524288，以修复因超限导致的错误。
    -   **链接**: [earendil-works/pi PR #5701](https://github.com/earendil-works/pi/pull/5701)

2.  **#5738: [已合并] 修复(ai): 对 Anthropic 1小时缓存写入定价为 2x 输入**
    -   **内容**: 修复了缓存计费 Bug。之前所有缓存写入都被统一按 5 分钟速率计费，导致 1 小时缓存写入被低估。此 PR 正确识别并按照 2 倍基础输入价格计算 1 小时缓存写入费用。
    -   **链接**: [earendil-works/pi PR #5738](https://github.com/earendil-works/pi/pull/5738)

3.  **#631: [已合并] 修复(ai): Google 思维链检测及移除不支持的 ID 字段**
    -   **内容**: 修正了 Google 提供商的思维链检测逻辑，错误地将包含 `thoughtSignature` 的内容识别为思维链。同时移除了不支持的 ID 字段，解决了与 Google API 的兼容性问题。
    -   **链接**: [earendil-works/pi PR #631](https://github.com/earendil-works/pi/pull/631)

4.  **#5846: [开放] 修复(tui): 稳定流式代码块渲染**
    -   **内容**: 直接对应热点 Issue #5825，通过修复流式渲染逻辑来防止强制滚动到底部的问题。
    -   **链接**: [earendil-works/pi PR #5846](https://github.com/earendil-works/pi/pull/5846)

5.  **#5849: [已合并] 功能(ai): 为 Anthropic Claude 增加 Azure AI Foundry 提供商**
    -   **内容**: 增加了对 Azure AI Foundry 上托管 Anthropic Claude 模型的一流支持，包括正确的 URL 格式、请求头和 Entra ID 认证。
    -   **链接**: [earendil-works/pi PR #5849](https://github.com/earendil-works/pi/pull/5849)

6.  **#5801: [已关闭] Nix 化 Pi**
    -   **内容**: 增加了 Nix Flake 支持，使得使用 Nix 包管理器的用户可以方便地构建、运行和配置 Pi。
    -   **链接**: [earendil-works/pi PR #5801](https://github.com/earendil-works/pi/pull/5801)

7.  **#5850: [已合并] chore(deps): 升级 vitest 并覆盖 esbuild**
    -   **内容**: 机械性的依赖升级，修复了 `npm audit` 中报告的 5 个高危漏洞（来自开发依赖），提升了开发环境的安全性。
    -   **链接**: [earendil-works/pi PR #5850](https://github.com/earendil-works/pi/pull/5850)

8.  **#5829: [已合并] 功能: 为自适应推理模型添加 “max” 思考级别**
    -   **内容**: 适配了 Anthropic 新模型的最高推理强度。Pi 原来的 `ThinkingLevel` 止步于 `xhigh`，此 PR 增加了 `max` 级别，以支持如 `claude-opus-4.8` 等模型的最高配置。
    -   **链接**: [earendil-works/pi PR #5829](https://github.com/earendil-works/pi/pull/5829)

9.  **#5832: [开放] 修复(ai): 展示提供商 HTTP 错误体而非模糊的 SDK 消息**
    -   **内容**: 改善了错误处理的透明度。当 API 返回错误时，Pi 将展示 HTTP 响应的具体错误信息，而不是 SDK 给出的模糊通用消息，极大方便了开发者调试。
    -   **链接**: [earendil-works/pi PR #5832](https://github.com/earendil-works/pi/pull/5832)

10. **#5859: [开放] 修复(ai): 将 Responses API 的提示词作为 instructions 发送**
    -   **内容**: 针对 OpenAI 新的 Responses API 进行了适配，将系统提示词（system prompt）放置在顶层的 `instructions` 字段中，而非作为普通的 `input` 消息，以符合 API 规范，确保行为正确。
    -   **链接**: [earendil-works/pi PR #5859](https://github.com/earendil-works/pi/pull/5859)

## 功能需求趋势

-   **云提供商集成**: 社区对特定云平台的原生支持需求强烈，如 **Azure AI Foundry** (PR #5849) 和 **Anthropic OAuth 订阅** (Issue #5821)。这表明用户群体正从通用 API Key 转向更复杂的云基础设施和计费模式。
-   **模型支持与配置细化**: 对新模型（如 **GLM-5.2**, **Minimax M3** 等）的支持是持续需求。同时，需求正从简单的“能不能用”转向“好不好用”，如 **支持 1M 上下文窗口** (Issue #5692)、**配置思考级别 (包括 'max')** (PR #5829) 以及 **细化压缩策略 (百分比触发)** (Issue #5848)。
-   **扩展性与自动化 API**: **RPC 接口扩展** (Issue #5810) 和 **Extension API 增强** (Issue #5781) 的提出，标志着社区正积极构建更强大的外部工具链和自动化工作流，Pi 作为 Agent 平台的开放性日益重要。
-   **核心 UX 打磨**: 修复 **流式渲染** (Issue #5825)、**树形导航截断** (Issue #5830) 等 bug，以及提出 **排除上下文的自定义消息** (Issue #5654)，都显示出在功能快速迭代后，社区将重心转向了交互细节和用户体验的精细化打磨。

## 开发者关注点

-   **错误信息不透明**: 多个 Issue(如 #5832) 聚焦于一个核心痛点：当 API 调用失败时，Pi 显示的抽象错误信息（如 `UnknownError`, 空 body）让开发者无法定位问题。社区强烈要求**透传原始 HTTP 错误体**以加速调试。
-   **依赖项管理和冗余**: Issue #5653 指出了包依赖项冗余导致的模块隔离问题，这是一个高级性能与稳定性问题，开发者社区对此的关注度很高，期待通过优化包打包策略来解决。
-   **本地与远程模型的一致性问题**: Issue #3715 (5分钟超时) 和 Issue #5845 (压缩机制问题在本地模型上更易暴露) 表明，本地部署的开发者面临与远程 API 不同的挑战，对 Pi 的底层网络和协议处理有更高的容错和自定义要求。
-   **状态一致性与文件锁**: Issue #5844 暴露了会话文件在并发写入时的 `EEXIST` 错误，导致实际 Agent 错误被掩盖。这反映了在异步和持久化场景下，文件系统状态管理是一个需要更稳健处理的技术点。
-   **平台规范遵循度**: Issue #534 关于 Linux XDG 规范的诉求获得了广泛赞誉，表明开发者不仅要求功能使用，还期望软件能优雅地融入其所在操作系统的生态环境，遵守社区最佳实践。

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

## Qwen Code 社区动态日报 | 2026-06-18

### 1. 今日速览

- **发布 v0.18.3 稳定版**，主要修复了 `sed` 编辑在文件历史中的追踪问题，并合并了 v0.18.2 的 CLI 取消问答流程修复。
- **模型选择持久化问题终结**：社区长期反馈的“多 provider 同名模型”无法记忆选择的问题，已由 #5241 / #5179 合并解决。
- **TUI 阻塞漏洞曝光**：SSH 下使用 `login1.inhibit-block-sleep` 功能时，认证弹窗导致界面无响应，已有 #5281 紧急报告。

### 2. 版本发布

过去 24 小时内发布了以下版本：

| 版本 | 类型 | 关键变更 |
|------|------|----------|
| v0.18.3-nightly.20260618 | 每日构建 | 基于 v0.18.3，修复 `sed` 编辑历史追踪（PR #5255） |
| **v0.18.3** | **稳定版** | 合并 v0.18.2 的 `ask_user_question` 取消修复（PR #5243） |
| v0.18.3-preview.0 | 预览版 | 与 v0.18.3 相同，标记为预览 |
| v0.18.2 | 稳定版 | 警告超大上下文指令（PR #5073）；文档更新（PR #5074） |
| v0.18.1-preview.1 | 预览版 | 同上修复，预览标记 |

**重点说明**：`v0.18.3` 是今天的正式稳定版，推荐升级。

---

### 3. 社区热点 Issues（共 6 条，全部收录）

| ID | 状态 | 标题 | 重要性 | 链接 |
|----|------|------|--------|------|
| #5173 | ✅ CLOSED | **多 provider 共享 model id 时选择无法持久化** | 影响多 API 网关用户，每日反馈较多；已由 PR #5241 修复 | [查看](https://github.com/QwenLM/qwen-code/issues/5173) |
| #5280 | 🔄 OPEN | **重新启用长命令搜索建议测试覆盖** | 测试回归，行为已可用但覆盖率丢失；社区贡献者 @tt-a1i 主动认领 | [查看](https://github.com/QwenLM/qwen-code/issues/5280) |
| #5277 | 🔄 OPEN | **重新启用 TableRenderer 前景色重置测试** | 底层渲染稳定性，相同贡献者跟进 | [查看](https://github.com/QwenLM/qwen-code/issues/5277) |
| #5275 | 🔄 OPEN | **重新启用 BaseSelectionList 滚动覆盖测试** | 组件回归测试，需处理 `act` 警告 | [查看](https://github.com/QwenLM/qwen-code/issues/5275) |
| #2845 | ✅ CLOSED | **支持按内容识别文件类型（如 `.dat` 内 PHP 代码）** | 被长期搁置，今日由 PR #5256 修复，社区期待已久 | [查看](https://github.com/QwenLM/qwen-code/issues/2845) |
| #5281 | 🔄 OPEN | **TUI 因 `login1.inhibit-block-sleep` 认证弹窗而阻塞** | **P2 级紧急 bug**；SSH 用户无法操作，需快速定位 | [查看](https://github.com/QwenLM/qwen-code/issues/5281) |

> 注：因仅 6 条 Issue 更新，此处全部列出。社区对测试覆盖和 TUI 稳定性的关注度显著上升。

---

### 4. 重要 PR 进展（精选 10 条）

| PR | 状态 | 摘要 | 亮点 | 链接 |
|----|------|------|------|------|
| #5241 | ✅ merged | **fix(model): 通过 baseUrl 消除 provider 歧义** | 直接修复 #5173，选择持久化；已合入 v0.18.3 后续版本 | [查看](https://github.com/QwenLM/qwen-code/pull/5241) |
| #5179 | ✅ merged | **fix(model): 记住选择的 provider（#5173）** | 与 #5241 协同，社区贡献 @doudouOUC | [查看](https://github.com/QwenLM/qwen-code/pull/5179) |
| #5256 | ✅ merged | **fix(core): 按内容检测 `.dat` 文件而非仅凭扩展名** | 修复 #2845，让隐藏 PHP 代码的 `.dat` 文件正常识别 | [查看](https://github.com/QwenLM/qwen-code/pull/5256) |
| #5181 | 🔄 in-review | **fix(core): 防止 `/quit` 时自动内存提取 OOM** | 解决 `buildTranscriptMessages()` 导致堆内存溢出，P1 优先 | [查看](https://github.com/QwenLM/qwen-code/pull/5181) |
| #5279 | 🔄 OPEN | **fix(core): 添加工具调用断路器防止无限循环** | 针对 #5234 的聚焦修复，解决 AI 死循环 | [查看](https://github.com/QwenLM/qwen-code/pull/5279) |
| #5231 | 🔄 OPEN | **feat(core,cli): 工作流工具 token 预算 + 运行时表面** | 新增 per-run 输出 token 预算，提升资源控制 | [查看](https://github.com/QwenLM/qwen-code/pull/5231) |
| #5202 | 🔄 OPEN | **feat(channel): 新增 QQ 机器人通道适配器** | 补齐国内 IM 支持，已有 Telegram/微信/钉钉/飞书；QQ 用户关注 | [查看](https://github.com/QwenLM/qwen-code/pull/5202) |
| #5220 | 🔄 OPEN | **feat(i18n): 工具显示名称本地化** | TUI/WebShell 中工具标签翻译，提升中文体验 | [查看](https://github.com/QwenLM/qwen-code/pull/5220) |
| #5145 | 🔄 OPEN | **feat(cli): 在输入占位符中显示跟进建议** | 模型回复后提示下一个可能的提问，交互优化 | [查看](https://github.com/QwenLM/qwen-code/pull/5145) |
| #5283 | 🔄 OPEN | **test(cli): 启用命令搜索长建议测试** | 回归测试覆盖，配合 #5280 提升质量 | [查看](https://github.com/QwenLM/qwen-code/pull/5283) |

---

### 5. 功能需求趋势

从 Issues 和 PR 中可提炼出社区当前最关注的方向：

- **多 Provider 与持久化配置**：修复 #5173 表明用户对多个模型网关（Token Plan、IdeaLab、BFF 等）的切换持久性要求强烈。
- **文件识别智能化**：#2845 / #5256 体现社区希望工具按内容而不是扩展名决定文件类型，尤其针对 `.dat`、`.log` 等无特定扩展名的文本文件。
- **国际化与本地化**：#5220 工具标签翻译、#2993 中文翻译修复，说明非英语用户对完整 UI 本地化有稳定需求。
- **跨平台稳定性**：#5281 TUI 在 Linux SSH + 睡眠抑制场景下崩溃，暴露了系统交互层的兼容性问题。
- **工作流资源管理**：#5231 token 预算功能代表用户希望精细控制每次调用的 token 消耗，避免意外超支。
- **通道适配扩展**：#5202 QQ 机器人通道的添加，意味着社区期望 Qwen Code 能接入更多国内即时通讯平台。

---

### 6. 开发者关注点

- **痛点 - 选择不可持久化**：在 `modelProviders` 中多个相同 model id 但不同 `baseUrl` 时，重启后丢失选择（已修复）。
- **痛点 - SSH 下 TUI 无响应**：#5281 阻止用户正常操作，且与睡眠抑制功能冲突，评审中需优先处理。
- **高频需求 - 测试覆盖回归**：多个 Issue 请求重新启用被跳过的测试（#5280 / #5277 / #5275），说明社区开发者对测试质量敏感，愿主动贡献。
- **稳定性 - OOM**：记忆提取时堆内存溢出（#5181）是严重性能问题，P1 优先级。
- **交互 - 工具循环**：#5279 断路器针对 AI 无限调用工具的 bug，影响用户体验和资源消耗。
- **平台 - macOS 26+ Liquid Glass**：#5284 显示社区紧跟 Apple 最新设计语言，希望桌面版图标跟上。

---

*以上数据来源于 [github.com/QwenLM/qwen-code](https://github.com/QwenLM/qwen-code) 2026-06-18 的公开动态。*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，各位开发者，以下是 2026年6月18日的 DeepSeek TUI (CodeWhale) 社区动态日报。

---

# DeepSeek TUI (CodeWhale) 社区动态日报 | 2026-06-18

## 今日速览

今日社区焦点集中在 **v0.9.0 版本的重磅功能“Workrooms”** 上，其 Phase 1 的PR已提交，旨在将CodeWhale从一个本地终端工具转变为支持协作的聊天原生平台。与此同时，多个关于 **Agent 行为失控** 的Bug报告（如自问自答、模式切换混乱）引发了社区的广泛讨论，开发者也积极响应并提交了多个修复补丁。

## 社区热点 Issues

以下为过去24小时内最值得关注的10个Issue：

1.  **[bug] Agent 过度介入，自问自答，偏离用户意图**
    *   **链接:** [Hmbown/CodeWhale Issue #3275](https://github.com/Hmbown/CodeWhale/issues/3275)
    *   **重要性:** 这是一个用户体验的核心问题。用户报告Agent会进入一个自我驱动的循环，在没有用户确认的情况下自行提出、回答并执行任务，严重偏离了原始指令。该问题被认为是#3061的回归，反映了模型行为控制的复杂性。社区用户对此非常关注，期望能快速修复。

2.  **[bug] Plan/Agent 模式切换不一致 & 工具权限混乱**
    *   **链接:** [Hmbown/CodeWhale Issue #3279](https://github.com/Hmbown/CodeWhale/issues/3279)
    *   **重要性:** 该问题详细描述了Plan模式切换到Agent模式后的权限混乱现象：`write_file`和`exec_shell`等工具被持续拒绝，但UI显示已处于Agent模式。修复后又出现越权执行。这暴露了模式切换状态管理的缺陷，直接影响用户的工作流。

3.  **[enhancement] v0.9.0 EPIC: 聊天原生 CodeWhale Workrooms**
    *   **链接:** [Hmbown/CodeWhale Issue #3209](https://github.com/Hmbown/CodeWhale/issues/3209)
    *   **重要性:** 这是未来版本的旗舰特性。旨在将CodeWhale从终端APP进化为支持多线程、可分享、支持GitHub上下文和移动端访问的“聊天原生工作间”，这代表了项目向协作平台演进的方向。

4.  **[bug] v0.8.61 自动生成多个 Agent 后 UI 卡死**
    *   **链接:** [Hmbown/CodeWhale Issue #3289](https://github.com/Hmbown/CodeWhale/issues/3289)
    *   **重要性:** 一个严重的稳定性问题。用户在Plan模式下多次输入改进计划后，UI直接卡死，无法操作。该问题直接影响用户使用体验，是当前版本的“拦路虎”。

5.  **[bug] [moonshot] v0.8.61 对某些 Schema 形状的修复不完整**
    *   **链接:** [Hmbown/CodeWhale Issue #3281](https://github.com/Hmbown/CodeWhale/issues/3281)
    *   **重要性:** 针对Kimi/Moonshot模型的参数修复不彻底。当Schema根节点包含`$ref`、`anyOf`等高级特性时，仍会因缺少`type:object`而导致400错误。这是一个对特定模型用户影响较大的兼容性问题。

6.  **[enhancement] 支持 OpenCode Go/Zen**
    *   **链接:** [Hmbown/CodeWhale Issue #1481](https://github.com/Hmbown/CodeWhale/issues/1481)
    *   **重要性:** 这是一个持续已久的呼声。社区用户希望CodeWhale能够支持OpenCode Go/Zen作为DeepSeek模型的提供商，因为其提供了DeepSeek-V4且价格便宜。这表明社区对模型供应商的多样性和成本效益有强烈需求。

7.  **[bug] `pre_tool_snapshot` 未遵循 `snapshots.enabled=false` 配置**
    *   **链接:** [Hmbown/CodeWhale Issue #3292](https://github.com/Hmbown/CodeWhale/issues/3292)
    *   **重要性:** 引人关注的配置失效问题。用户明确关闭了快照功能，但Git仓库仍被完整复制，占用了数GB的磁盘空间。这属于数据管理和资源使用的严重Bug，会极大地降低用户信任。

8.  **[enhancement] `config.toml` 文件内容改进**
    *   **链接:** [Hmbown/CodeWhale Issue #3282](https://github.com/Hmbown/CodeWhale/issues/3282)
    *   **重要性:** 一个细节但重要的质量改进需求。用户在TUI中修改配置时，所有注释内容会被自动清除。这阻碍了用户通过注释来管理配置，反映了社区对配置文件“人性化”管理的追求。

9.  **[documentation, v0.9.0, cleanup, tui] EPIC: 阶段式命令边界重构**
    *   **链接:** [Hmbown/CodeWhale Issue #2870](https://github.com/Hmbown/CodeWhale/issues/2870)
    *   **重要性:** 一个长期的、底层的架构改造EPIC。旨在将命令边界进行分段重构，以实现更小、更可合并的PR。这表明项目在追求新功能的同时，也在进行重要的内部架构清理和优化。

10. **[enhancement, question] 支持非交互模式下的会话连续性**
    *   **链接:** [Hmbown/CodeWhale Issue #1530](https://github.com/Hmbown/CodeWhale/issues/1530)
    *   **重要性:** 这是一个面向自动化和CI/CD的功能请求。用户希望在`exec`和`--prompt`等非交互模式下，能够通过`--resume`或`--session-id`等方式继续之前的对话，这对于构建多轮对话工作流至关重要。

## 重要 PR 进展

以下是过去24小时内提交的10个重要PR：

1.  **[feat] 实现 Workrooms Phase 1 — 数据模型、端点和文档**
    *   **链接:** [Hmbown/CodeWhale PR #3277](https://github.com/Hmbown/CodeWhale/pull/3277)
    *   **重要性:** **今日最重磅的PR**。为v0.9.0的Workroom功能打下了坚实的地基。包含了完整的设计RFC、数据模型、REST API端点、数据库迁移和CLI工具，是CodeWhale迈向协作平台的第一步。

2.  **[fix] 修复 Plan/Agent 模式切换 —— `approval_mode` 恢复 + 自动执行防护**
    *   **链接:** [Hmbown/CodeWhale PR #3283](https://github.com/Hmbown/CodeWhale/pull/3283)
    *   **重要性:** 直接回应了社区热点Issue #3279。修复了模式切换时`approval_mode`（审批模式）未恢复和从YOLO模式切回后自动执行计划的两个关键Bug。

3.  **[fix] TUI 渲染性能：对思考流去抖渲染**
    *   **链接:** [Hmbown/CodeWhale PR #3284](https://github.com/Hmbown/CodeWhale/pull/3284)
    *   **重要性:** 解决了#1620中提到的推理（thinking）块渲染卡顿的问题。通过对每次推理Delta进行去抖处理，大幅减少了UI刷新次数，从而显著提升了思考流的渲染速度。

4.  **[fix] 在 stall/cancel 恢复前持久化 session，确保 --continue 保留历史**
    *   **链接:** [Hmbown/CodeWhale PR #3285](https://github.com/Hmbown/CodeWhale/pull/3285)
    *   **重要性:** 修复了#2739中提到的数据丢失问题。当agent卡住或被取消后，使用`--continue`会丢失整个未完成轮次的内容。PR通过在清理前持久化session来确保历史完整。

5.  **[fix] 确保 Kimi 参数根节点对所有 Schema 形状都有 `type:object`**
    *   **链接:** [Hmbown/CodeWhale PR #3286](https://github.com/Hmbown/CodeWhale/pull/3286)
    *   **重要性:** 直接修复了社区热点Issue #3281。通过修改`sanitize_for_kimi_parameters`函数，确保即使是使用了`$ref`、`allOf`等高级特性的Schema根节点也会被注入`type:object`，从而避免Kimi/Moonshot API返回400错误。

6.  **[fix] 提议增加 `scope_discipline` 规则以防止自我问答循环**
    *   **链接:** [Hmbown/CodeWhale PR #3290](https://github.com/Hmbown/CodeWhale/pull/3290)
    *   **重要性:** 社区成员yekern主动提交的PR，尝试从`constitution.md`层面解决Issue #3273（与#3275类似）描述的Agent自问自答问题。通过在系统提示中加入“范围纪律”规则来约束Agent行为。

7.  **[fix] 遵守 `snapshots.enabled` 配置，禁止 per-tool 快照**
    *   **链接:** [Hmbown/CodeWhale PR #3293](https://github.com/Hmbown/CodeWhale/pull/3293)
    *   **重要性:** 直接修复了热点Issue #3292。PR确保了在`snapshots.enabled = false`的配置下，每个`write_file`/`edit_file`工具调用前不会创建快照，解决了用户磁盘空间被意外占满的问题。

8.  **[fix] 保留 config 文件中的注释**
    *   **链接:** [Hmbown/CodeWhale PR #3291](https://github.com/Hmbown/CodeWhale/pull/3291)
    *   **重要性:** 直接回应了社区需求Issue #3282。PR修改了所有重写配置文件（`config.toml`等）的代码路径，使用`toml_edit`库合并新旧文件，从而保留了用户的注释和临时禁用的配置项。

9.  **[fix] 将composer历史写入 `.codewhale` 目录，而非遗留的 `.deepseek`**
    *   **链接:** [Hmbown/CodeWhale PR #3294](https://github.com/Hmbown/CodeWhale/pull/3294)
    *   **重要性:** 修复了在新安装中仍会创建遗留的`~/.deepseek/`目录的问题。PR将`composer_history.txt`的默认读写路径改为`~/.codewhale/`，使新安装的目录结构更干净。

10. **[feat] TUI 端：运行时加载并执行 `ask permission` 规则**
    *   **链接:** [Hmbown/CodeWhale PR #3295](https://github.com/Hmbown/CodeWhale/pull/3295)
    *   **重要性:** 实现了在TUI运行时加载`permissions.toml`中的“询问”权限规则。这个PR直接提升了工具的权限控制和安全性，允许用户更精细地控制模型调用`exec_shell`等危险操作。

## 功能需求趋势

综合过去24小时的动态，社区最关注的功能方向可以提炼为以下几点：

1.  **协作与平台化 (Collaboration & Platformization):** 从 Workrooms EPIC 和大量相关PR可以看出，社区对 CodeWhale 的期望已超越一个终端工具，转向支持团队协作、会话持久化和远程访问的平台化产品。
2.  **Agent 行为可控性 (Agent Behavior Control):** 这是当前最大的痛点。大量Issue和PR集中在解决Agent的自问自答、模式切换混乱和权限管理问题。社区迫切需要一个更“听话”、行为可预测的Agent。
3.  **配置体验与数据管理 (Configuration & Data Hygiene):** 社区对配置文件的“人性化”管理（保留注释）和数据存储的“干净”程度（不遗留旧文件夹、快照可控）提出了更高的要求。
4.  **LLM 提供商兼容性 (LLM Provider Compatibility):** 对支持更多提供商（如OpenCode Go/Zen）以及修复特定提供商API的Bug（如Kimi/Moonshot）持续保持关注。

## 开发者关注点

开发者反馈中反复出现的痛点和需求包括：

-   **Agent 行为“失控”**：Agent 不按照用户指示工作，自行其是（自问自答、超范围执行），这是对开发者信任度的最大挑战。
-   **模式切换“混乱”**：Plan模式和Agent模式切换后，权限设置和行为逻辑不一致，修复后出现新问题，导致工作流程中断。
-   **性能问题**：特别是在处理长思考流（thinking stream）时，UI渲染卡顿，严重影响观察和交互体验。
-   **配置管理“反直觉”**：用户在TUI中修改配置，不知不觉丢失了所有注释信息；关闭快照功能后磁盘空间依然被占用，这些都不符合用户的预期。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*