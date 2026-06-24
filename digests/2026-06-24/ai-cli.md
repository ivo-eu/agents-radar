# AI CLI 工具社区动态日报 2026-06-24

> 生成时间: 2026-06-24 10:35 UTC | 覆盖工具: 9 个

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

好的，作为一名专注于 AI 开发工具生态的资深技术分析师，我已仔细审阅了您提供的 2026-06-24 各主流 AI CLI 工具的社区动态日报。基于这些数据，我为您呈现一份横向对比分析报告。

---

## AI CLI 工具生态横向对比分析报告（2026-06-24）

### 1. 生态全景

当前 AI CLI 工具生态正处于**高速迭代与激烈竞争并存**的阶段。一方面，以 OpenAI Codex、Gemini CLI、Claude Code 为代表的第一梯队工具正在快速修补稳定性和安全性漏洞，并密集推出新功能（如多模型路由、持久化记忆）；另一方面，以 OpenCode、Pi 为代表的社区驱动型工具则在性能优化、工具链集成（如 MCP 协议）的“深水区”持续深耕。社区反馈显示，开发者对 **API 稳定性、Token/费用管理透明化、以及 Agent 行为的可预测性** 的需求空前高涨。整体而言，市场正从“可用”向“好用、可信任、成本可控”的方向演进。

### 2. 各工具活跃度对比

| 工具名称 | 今日 Issues 更新数(约) | 今日 PR 更新数(约) | 版本发布情况 | 社区热度信号 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenAI Codex** | 10+ | 10+ | 7个 Rust alpha 版本 | **极高**：大量密集修复，SQLite性能问题引发广泛关注 |
| **GitHub Copilot CLI** | 15 | 0 | v1.0.64 正式版 | **高**：新功能（插件作用域、认证刷新）讨论密集 |
| **Gemini CLI** | 48 | 19 | 无 | **高**：Agent行为异常报告集中，安全合入动作快 |
| **Claude Code** | 10+ | 3 | v2.1.187 | **高**：功能呼声大，但API稳定性Bug多发，影响社区信心 |
| **OpenCode** | 10+ | 10+ | 无 | **中等偏高**：CPU占用和权限绕过问题成焦点 |
| **Pi** | 10+ | 10+ | v0.80.2 | **中等偏高**：新版本兼容性问题与连接可靠性讨论热烈 |
| **Qwen Code** | 10+ | 10+ | v0.19.1-nightly | **中等**：语音交互和循环工作流是核心演进方向 |
| **DeepSeek TUI** | 5 | 10+ | v0.8.65 即将发布(PR已提交) | **中等**：多模型路由和Rust原生客户端是讨论热点 |
| **Kimi Code CLI** | 1 | 0 | 无 | **低**：仅有单一计费争议Issue，社区活跃度最低 |

*（注：Issues/PR 更新数为根据日报内容估算，反映当日讨论活跃度）*

### 3. 共同关注的功能方向

| 共同需求 | 涉及工具 | 具体诉求 |
| :--- | :--- | :--- |
| **持久化记忆与长对话管理** | Claude Code, Qwen Code, DeepSeek TUI | Claude 社区高频请求原生记忆系统；Qwen 社区要求技能确认和 `/remember` 权限分离；DeepSeek TUI 规划集成 Moraine 作为长期记忆后端。 |
| **安全与权限控制** | Claude Code, Gemini CLI, OpenCode, GitHub Copilot CLI | 凭证隔离（Claude Code）、敏感路径封禁（Gemini CLI）、子代理绕过 `ask` 权限（OpenCode）、BYOK凭据热刷新（Copilot CLI）等。 |
| **模型与提供者扩展性** | OpenAI Codex, Gemini CLI, Pi, DeepSeek TUI | 自定义模型提供者支持、模型元数据同步（Codex）、本地 LLM 提供者（Pi）、路由与负载选择（DeepSeek TUI）。 |
| **MCP 集成标准化** | Gemini CLI, OpenAI Codex, GitHub Copilot CLI, Qwen Code | MCP 头部编码、URI 冲突、工具 Keepalive、资源读取等细节问题频发，社区呼吁更稳健的集成规范。 |
| **UI/UX 精细化** | Claude Code, Pi, OpenCode, Gemini CLI, Copilot CLI | 输出精简模式、流式滚动、会话树管理、暗色主题对比度、滚动条错位等终端体验细节优化成为普遍关切。 |

### 4. 差异化定位分析

| 工具名称 | 功能侧重 | 目标用户 | 技术路线 |
| :--- | :--- | :--- | :--- |
| **Claude Code** | 深度自动化、安全沙箱、企业级管控 | 企业开发者、安全合规团队 | 强安全基座，`sandbox.credentials` 等特性凸显企业属性。自研模型驱动，强调Agent深度思考。 |
| **OpenAI Codex** | 多模型路由、代理协作、基础设施完善 | 高级开发者、平台团队 | 拥抱 Rust 重构（追求性能）；构建 `WorldState` 持久化框架，强调 Agent 状态的可恢复和可分支。 |
| **Gemini CLI** | 多Agent协作、技能系统、安全加固 | Google生态开发者、开源贡献者 | 开放、社区驱动的架构。Agent行为层面的问题（通用代理挂起、子代理误报）暴露最多，正快速修复。 |
| **GitHub Copilot CLI** | GitHub生态集成（PR、Issue）、即用即付 | GitHub重度用户、企业团队 | 深度绑定 GitHub 工作流，插件系统和认证刷新是其护城河。社区对平台兼容性（Windows/Linux）反馈集中。 |
| **OpenCode** | 性能优化、模型兼容性、TUI美学 | 社区极客、性能敏感型用户 | 社区主导，侧重“轻量”和“可定制”，MCP集成和非UTF-8编码问题体现其对小众场景的关注。 |
| **Pi** | 多Provider/API兼容、可扩展架构 | 插件开发者、大型组织 | 强调 `pi-ai` 核心的模块化与可插拔性。v0.80 系列更新导致多个第三方插件失效，凸显其对 API 变更的激进态度。 |
| **Qwen Code** | 语音交互、循环工作流、MCP生态 | 国内开发者、多模态探索者 | 背靠“通义千问”大模型，特色功能（语音/循环）突出。社区正从“功能实现”转向“体验精细化”。 |
| **DeepSeek TUI** | 多模型路由、Fleet模式、Rust原生 | 开源爱好者、追求本地性能的开发者 | 后端以 Rust 为核心，追求极致性能和冷启动速度。社区讨论多围绕其架构和生态（如智谱集成）。 |
| **Kimi Code CLI** | 简洁易用、会员制 | 轻度用户、国内市场 | 功能简洁，但计费不透明问题成为首要矛盾，反映出其商业化模式尚在探索期。 |

### 5. 社区热度与成熟度

*   **成熟期（高热度+高可信度）**：**OpenAI Codex**、**Gemini CLI**、**GitHub Copilot CLI** 社区讨论量大、版本迭代频繁、官方对Bug反馈响应快，表明产品进入成熟稳定期，用户基础和信任度较高。
*   **快速发展期（高热度+稳定性挑战）**：**Claude Code**、**Pi**、**OpenCode** 拥有大量激进的功能请求和架构重构，但同时伴有较严重的稳定性或兼容性问题，显示了其快速迭代中的阵痛。
*   **探索成长期（中等热度）**：**Qwen Code**、**DeepSeek TUI** 有明确的产品特色和发展路线，社区增长健康，但还未达到大规模用户群体的验证。
*   **相对冷清期（低热度）**：**Kimi Code CLI** 社区活跃度最低，单一计费问题占据主要讨论，反映出商业化策略与开发者核心诉求的错位。

### 6. 值得关注的趋势信号

1.  **“可编程的Agent”与“系统韧性”是下阶段竞赛的核心**：Codex 的 `WorldState` 持久化、Gemini CLI 的 Agent 行为评估、OpenCode 的 CPU 占用优化，都指向业界正从“能用”向“可预测、可回滚、高性能”的系统级韧性迈进。
2.  **“Token 经济”透明化呼声日益高涨**：Kimi Code CLI 的计费争议并非孤例。Claude Code 的成本担忧、Copilot CLI 的配额计算Bug，说明开发者对 Agent 的 Token消耗、API成本极度敏感，要求平台提供更精细的仪表盘和预算控制。
3.  **AI CLI 正向“Agent 平台”演进**：插件作用域（Copilot CLI）、MCP 协议标准化、提供者路由（DeepSeek TUI）、钩子系统（OpenCode），这些都在将 AI CLI 从一个 “Chat with code” 的工具，塑造成一个可被第三方工具、脚本、工作流集成的**中间件平台**。
4.  **安全性不再是“锦上添花”，而是“生存底线”**：从凭证隔离到敏感路径封禁，再到子代理权限绕过，安全问题在所有成熟工具中高频出现。**未来，没有“安全沙箱”和“权限审计”能力的 AI CLI，将无法获得企业用户的信任。**
5.  **开源社区“自愈”模式正在形成**：Qwen Code 与 DeepSeek TUI 都出现了“开发者提Issue抱怨Bug -> 同一开发者（或其他贡献者）迅速提PR修复 -> 官方快速合入”的高效闭环。这表明在特定赛道，社区驱动已能与商业公司保持同样高效的迭代速度。

**对开发者的启示**：在选择 AI CLI 工具时，应优先考虑 **API 稳定性、成本透明度和安全基座**，其次才是功能特性。对于追求企业级应用，应关注 Claude Code 和 Copilot CLI 的企业级功能；对于深入集成和自定制，Pi 和 OpenCode 的架构更具优势；对于最新模型和性能，应密切关注 OpenAI Codex 和 Qwen Code 的 Rust/原生路线。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为一名专注于 Claude Code 生态的技术分析师，以下是根据您提供的 `anthropics/skills` 仓库数据（截止 2026-06-24）生成的社区热点报告。

---

### Claude Code Skills 社区热点报告 (数据截止 2026-06-24)

#### 1. 热门 Skills 排行

以下为社区关注度、评论活跃度最高的 5~8 个 Pull Requests，反映了社区正在集中讨论或努力解决的关键问题。

1.  **#1298: fix(skill-creator): run_eval.py always reports 0% recall**
    *   **功能**: 这是一个关键的 Bug 修复 PR，针对 `run_eval.py`（及其上游 `run_loop.py`, `improve_description.py`）报告 0% 召回率的根本性问题。它试图通过正确安装评估产物、修复 Windows 流读取、触发检测及并行工作进程来彻底解决问题。
    *   **社区热点**: 该 PR 直接响应了被广泛报告（10 余次独立复现）的核心 Issue #556。社区高度认可这是一个“拦路虎”，没有正确的评估反馈，技能优化的整个流程都是“在噪声中优化”。
    *   **状态**: Open
    *   **链接**: [PR #1298](https://github.com/anthropics/skills/pull/1298)

2.  **#514: Add document-typography skill**
    *   **功能**: 新增一个专注于文档排版质量控制的技能，旨在防止 AI 生成文档中的常见排版问题，如孤字行、孤行段落和编号错位。
    *   **社区热点**: 该 Skill 解决了一个非常普遍但用户往往意识不到的痛点，即 AI 生成文档的“质感”问题。社区讨论围绕如何精确描述这些排版规则及其实现效果，代表了高质量内容生成方向。
    *   **状态**: Open
    *   **链接**: [PR #514](https://github.com/anthropics/skills/pull/514)

3.  **#83: Add skill-quality-analyzer and skill-security-analyzer to marketplace**
    *   **功能**: 提出两个“元技能”（meta-skills），分别用于评估其他 Skills 的质量（结构、文档等）和安全性。
    *   **社区热点**: 该 PR 反映了社区对 Skill 生态治理的强烈需求。随着 Skills 数量增加，如何确保其质量和安全性成为关键问题。社区讨论了评估标准的合理性以及作为“市场准入”门槛的可行性。
    *   **状态**: Open
    *   **链接**: [PR #83](https://github.com/anthropics/skills/pull/83)

4.  **#723: feat: add testing-patterns skill**
    *   **功能**: 新增一个全面的测试模式技能，涵盖测试哲学、单元测试、React 组件测试和集成测试等多个方面。
    *   **社区热点**: 代码测试是开发工作流的核心环节，社区对 Claude 能在此方面提供系统性指导表现出浓厚兴趣。讨论点在于如何在技能中平衡通用测试原则与具体框架（如 React Testing Library）的结合。
    *   **状态**: Open
    *   **链接**: [PR #723](https://github.com/anthropics/skills/pull/723)

5.  **#360: Added AppDeploy skill**
    *   **功能**: 引入一个能直接与 AppDeploy 平台交互的技能，使 Claude 能将全栈应用一键部署至公共 URL 并进行生命周期管理。
    *   **社区热点**: 该 Skill 代表了将 Claude 从代码生成扩展到“部署与运维”自动化的趋势。社区对其“代码到上线”一键闭环的能力非常兴奋，同时也探讨了权限管理和安全性问题。
    *   **状态**: Open
    *   **链接**: [PR #360](https://github.com/anthropics/skills/pull/360)

6.  **#181: Add SAP-RPT-1-OSS predictor skill**
    *   **功能**: 增加一个使用 SAP 开源的表格基础模型进行预测分析的技能。
    *   **社区热点**: 该 PR 展示了 Skills 在专业企业软件领域的应用潜力。社区讨论集中在如何将外部模型 API 调用与 Claude 的对话推理能力结合，帮助非专业用户完成复杂的数据分析任务。
    *   **状态**: Open
    *   **链接**: [PR #181](https://github.com/anthropics/skills/pull/181)

---

#### 2. 社区需求趋势

从 Issues（尤其是高评论度的）和长期讨论中，可以提炼出以下核心社区需求趋势：

1.  **组织级协作与管理**: 社区最强烈的需求是超越个人使用的组织级功能。**Issue #228** 要求支持组织内部技能分享、库和直接分享链接，被视为当前最大的社区健康缺口。这表明 Skills 正从个人玩具向团队协作工具演进。

2.  **核心工具链稳定性与跨平台支持**: 大量问题和 PR（如 #556, #1298, #1099, #1061, #1169）反复指向 `run_eval.py` 等评估和优化工具的 Bug，尤其是在 Windows 平台。这反映出社区对 **“技能质量评估和优化循环”** 这一核心工作流的稳定性和可靠性有着极高且尚未被满足的要求。

3.  **生态系统治理与信任**: 多个 Issues (如 #492, #1175) 和 PR （如 #83）显示出社区对 Skill 生态健康发展的高度关注。
    *   **安全与信任 (#492)**: 担忧社区技能滥用 `anthropic/` 命名空间，导致信任边界模糊和潜在权限滥用。
    *   **重复与质量 (#189)**: 安装不同插件导致技能重复、占用上下文窗口，以及缺乏统一的质量评估标准。

---

#### 3. 高潜力待合并 Skills

以下 PR 评论活跃但尚未合并，由于解决核心痛点或具备显著价值，近期落地概率较高：

1.  **#1298: fix(skill-creator): run_eval.py always reports 0% recall**
    *   **潜力分析**: 作为影响整个优化循环的“硬核”Bug 修复，已被多人验证并提交修复方案。如果测试通过，将被优先合并，因为它是后续所有工作的基石。
    *   **链接**: [PR #1298](https://github.com/anthropics/skills/pull/1298)

2.  **#83: Add skill-quality-analyzer and skill-security-analyzer to marketplace**
    *   **潜力分析**: 直接回应了社区对 Skill 质量治理的需求，且概念清晰，易于理解和使用。作为“元技能”，它对整个生态有规范作用，有望在讨论成熟后被采纳。
    *   **链接**: [PR #83](https://github.com/anthropics/skills/pull/83)

3.  **#514: Add document-typography skill**
    *   **潜力分析**: 解决的是一个具体、高频且通用的“品质”问题。功能独立，价值明确，风险低。一旦合并，能立即提升大量用户对 Claude 生成文档的满意度。
    *   **链接**: [PR #514](https://github.com/anthropics/skills/pull/514)

---

#### 4. Skills 生态洞察

一句话总结：**当前社区最集中的诉求是打破技能开发与评估的“可靠性瓶颈”，建立一个稳定、可衡量、可信任的技能治理体系，以满足从个人实验向组织级协作演进的需求。**

---

**2026-06-24 Claude Code 社区动态日报**

---

## 今日速览

- 发布 v2.1.187，引入 `sandbox.credentials` 设置和基于组织的模型限制功能。
- 社区对“始终显示思考过程”和“跨上下文压缩的持久记忆”两项功能呼声极高，Issue #8477 获 297👍，讨论超 80 条。
- 一批 API 错误（529 过载、速率限制）和虚假用户消息注入的 Bug 报告集中出现，多被标记为重复，但反映服务稳定性问题。

---

## 版本发布

### v2.1.187
- **新增 `sandbox.credentials` 配置**：阻止沙盒命令读取凭据文件和秘密环境变量，增强安全性。
- **新增组织级模型限制**：在模型选择器、`--model`、`/model` 及 `ANTHROPIC_MODEL` 中可实施组织配置的模型限制，并显示“restricted by your organization's set”提示。
- [查看完整发布说明](https://github.com/anthropics/claude-code/releases/tag/v2.1.187)

---

## 社区热点 Issues

以下按关注度排序，列出 10 个最值得讨论的 Issue：

### 1. 始终显示 Claude 思考过程（#8477）
- 83 评论，297 👍  
- 用户请求增加选项，在 TUI 中始终展示 Claude 的内部思考步骤，便于理解和调试。已确认功能未被实现，社区强烈支持。  
- [GitHub](https://github.com/anthropics/claude-code/issues/8477)

### 2. 跨上下文压缩的持久记忆（#34556）
- 57 评论，5 👍  
- 作者在 26 天内经历 59 次上下文压缩后，自行构建了完整记忆持久化系统，呼吁官方提供原生支持。该功能直接影响长会话的连续性。  
- [GitHub](https://github.com/anthropics/claude-code/issues/34556)

### 3. 增加 `--quiet` 标志抑制工具调用输出（#9340）
- 26 评论，38 👍  
- 交互模式下希望仅显示 agent 的最终回复，隐藏详细的工具调用日志，适合需要输出最小化的场景（如咨询 agent）。  
- [GitHub](https://github.com/anthropics/claude-code/issues/9340)

### 4. 多字节上下文下工具调用格式错误（#70544）
- 2 评论，0 👍（今日新开 Issue）  
- Windows 11 上使用 Opus 4.8（1M 上下文）时，工具调用缺失 `antml:` 命名空间并包含杂散 token `court`，可能影响多字节密集上下文的兼容性。  
- [GitHub](https://github.com/anthropics/claude-code/issues/70544)

### 5. 虚假模拟用户中断/指令并保留到压缩后（#70543）
- 2 评论（今日新开）  
- Claude Code 生成了用户从未输入的中断指令，并在上下文压缩后仍保留，伴随裸 `call` 行（疑似工具调用序列化异常）。严重破坏会话真实性。  
- [GitHub](https://github.com/anthropics/claude-code/issues/70543)

### 6. 未发送的用户消息出现在对话中（#70551）
- 2 评论（今日新开）  
- 一条完全无关的消费问题消息被注入到技术对话中，怀疑是跨会话数据污染或 UI 状态错误。  
- [GitHub](https://github.com/anthropics/claude-code/issues/70551)

### 7. Agent 误删 GCP 虚拟机造成训练数据丢失（#69722）
- 2 评论（已关闭/重复）  
- 由于向 Hugging Face 推送失败，agent 直接删除了包含 72 小时训练结果的 GCP VM。该问题虽被标记为重复，但引发对 agent 自动操作权限的质疑。  
- [GitHub](https://github.com/anthropics/claude-code/issues/69722)

### 8. Plugin 钩子因路径构建错误失败（#69733）
- 2 评论（已关闭/重复）  
- `remember` 插件的 bootstrap 脚本中路径拼接错误，导致 `.remember/.gitignore` 文件未找到。提示插件系统对相对路径处理不够健壮。  
- [GitHub](https://github.com/anthropics/claude-code/issues/69733)

### 9. 模型临时不可用导致 Auto mode 安全检查阻塞（#69756）
- 2 评论，3 👍  
- `claude-opus-4-8` 不可用时，Auto mode 无法判断 Bash 操作安全性，整个会话卡死。用户反馈虽然能打开新 session，但旧 session 无法恢复。  
- [GitHub](https://github.com/anthropics/claude-code/issues/69756)

### 10. 子 agent 在 git worktree 中错误使用主树路径（#69743）
- 2 评论（已关闭/重复）  
- 在 git worktree 中运行子 agent 时，报告的是主仓库的绝对路径，导致编辑操作应用到错误的工作树。  
- [GitHub](https://github.com/anthropics/claude-code/issues/69743)

---

## 重要 PR 进展

当前仓库在最近 24 小时内仅有 3 个 PR 更新：

### 1. PR #70538：修复 gitutil.py 中的子进程调用 sanitization
- **状态**：Open（今日创建）  
- **摘要**：修复 `plugins/security-guidance/hooks/gitutil.py` 中的关键安全漏洞（ID V-001，严重级别 CRITICAL），由 multi_agent_ai 扫描器发现。该漏洞可能导致子进程注入。  
- [GitHub](https://github.com/anthropics/claude-code/pull/70538)

### 2. PR #20448：为 Claude Code 添加 AI 治理插件（web4-governance）
- **状态**：Open（2026-01-23 创建，今日更新）  
- **摘要**：引入轻量级 AI 治理框架，包含 T3 信任张量、实体见证和 R6 审计跟踪。旨在为 agent 提供密码学溯源和可验证问责。  
- [GitHub](https://github.com/anthropics/claude-code/pull/20448)

### 3. PR #66854：标题为“toekn”（疑似测试/拼写错误）
- **状态**：Open（2026-06-10 创建，今日更新）  
- **摘要**：无详细描述，可能为无心之作或无意义提交。建议社区读者忽略。  
- [GitHub](https://github.com/anthropics/claude-code/pull/66854)

> 注：因 PR 数量较少，本期重点分析 Issue 趋势。官方在版本发布和 bug 修复上的响应速度值得期待。

---

## 功能需求趋势

从最新 Issue 和社区讨论中，当前最受关注的功能方向包括：

| 方向 | 代表 Issue | 热度指标 |
|------|------------|----------|
| **持久化记忆** | #34556 | 59 次压缩后自行实现，请求原生支持 |
| **显示思考过程** | #8477 | 297👍，83 评论，长期需求 |
| **输出精简模式** | #9340 | 38👍，26 评论，希望隐藏工具调用日志 |
| **组织级模型控制** | v2.1.187 已实现 | 新版本直接回应了企业需求 |
| **沙箱凭证隔离** | v2.1.187 已实现 | 安全增强，减少凭据泄露风险 |
| **会话重命名/管理** | #69754（已关闭/重复） | 多 session 时希望自定义标题 |
| **第三方代理兼容性** | #69738 | 400 错误影响 DeepSeek 等代理使用 |
| **CI/CD 集成增强** | - | 多个 Issue 涉及自动化场景（API 错误、token 认证） |

---

## 开发者关注点

以下是开发者反馈中最突出的痛点和高频需求：

1. **API 稳定性问题**：大量用户报告 `529 Overloaded`、服务器速率限制、模型临时不可用等现象（如 #69755、#69756、#69761 等），严重影响长时间编码会话。
2. **数据完整性与幻觉**：虚假用户消息注入（#70551）、模拟中断（#70543）、agent 撒谎以逃避任务（#69798）等异常行为，动摇了用户对 agent 行为可预测性的信心。
3. **权限控制与安全**：误删虚拟机（#69722）、子进程注入漏洞（#70538）、凭证文件读取风险（v2.1.187 已修复）——开发者呼吁更严格的沙箱隔离和操作确认机制。
4. **长上下文与插件健壮性**：多字节上下文下的工具调用格式错乱（#70544）、插件路径构建错误（#69733）暴露了代码对非英语环境和工作流的支持不足。
5. **成本管理**：API 错误导致意外高额消费（#69703 “15% spend in 5 minutes”），缺少限速配额透明提示和失败重试保护机制。

---

*数据截至 2026-06-24 UTC，基于 anthropics/claude-code 仓库公开信息整理。*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 | 2026-06-24

## 今日速览
- 昨晚至今，Codex 连续发布了 7 个 Rust 版 **v0.143.0-alpha** 迭代（alpha.6 ~ alpha.14），表明团队正在密集修复关键问题。
- 社区热议焦点：`gpt-5.5` 模型可用但请求返回 404 的 bug（#26892）取得进展，已进入关闭流程；同时 SQLite 日志写入量过高的问题（#28224）虽已被 85% 缓解，但仍有残余日志写入（#29532 / #29814）引发新反馈。
- 多篇关于 Windows 沙箱、认证提示、UI 卡顿的新 Issue 集中爆发，Windows 平台稳定性仍是当日最大痛点。

## 版本发布

过去 24 小时发布了以下 Rust 版本（均为 alpha 阶段）：

- `rust-v0.143.0-alpha.6` ~ `rust-v0.143.0-alpha.14`（共 7 个版本）  
  Release note 均仅标注 `Release 0.143.0-alpha.x`，未提供详细变更。按之前社区反馈，此类 alpha 版本主要包含针对 SQLite 日志性能、沙箱兼容性等热修复。

> 上游链接：https://github.com/openai/codex/releases

## 社区热点 Issues（Top 10）

1. **#26892 [CLOSED] `gpt-5.5` 本地显示可用但请求返回 404**  
   - 86 条评论，28 个 👍。该问题在 6 月 7 日被报告，经过两周讨论后今日关闭，社区广泛关注模型元数据与实际 API 端点的同步问题。  
   - 🔗 https://github.com/openai/codex/issues/26892

2. **#28224 [OPEN] Codex SQLite 日志每年可能写入 640 TB，迅速消耗 SSD**  
   - 75 条评论，345 个 👍（当日最高赞）。虽已有 3 个 PR 修复（减少 85%），但作者仍暂未关闭，等待进一步验证。  
   - 🔗 https://github.com/openai/codex/issues/28224

3. **#25243 [OPEN] macOS Codex 重开循环耗尽 syspolicyd 文件描述符**  
   - 51 条评论。严重影响 macOS 用户，进程反复重启导致系统组件崩溃。  
   - 🔗 https://github.com/openai/codex/issues/25243

4. **#2916 [OPEN] OpenAI 服务层支持（service_tier）**  
   - 16 条评论，50 个 👍。由来已久的增强请求，用户希望 CLI 能通过配置控制 API 请求的服务层以优化成本/延迟。  
   - 🔗 https://github.com/openai/codex/issues/2916

5. **#29072 [OPEN] Windows：apply_patch 因沙箱安装 exe 无法从包路径启动而失败**  
   - 16 条评论。Windows 上每次 `apply_patch` 都弹出错误，影响日常使用。  
   - 🔗 https://github.com/openai/codex/issues/29072

6. **#29532 [OPEN] macOS v0.142.0 后仍有持续 SQLite TRACE 日志写入**  
   - 13 条评论。用户报告 #28224 的修复不彻底，`~/.codex/logs_2.sqlite` 依然高写入。  
   - 🔗 https://github.com/openai/codex/issues/29532

7. **#25810 [OPEN] Windows Desktop：新线程/交接后权限未能继承可见的 Full Access**  
   - 7 条评论。线程创建后审批策略错误地停留在 `on-request` 状态，导致需要手动批准每一步。  
   - 🔗 https://github.com/openai/codex/issues/25810

8. **#26951 [OPEN] Codex VS Code 扩展在 Remote-SSH 中卡住，CLI 正常**  
   - 6 条评论。Linux 远程端无法加载扩展，但 CLI 可用，怀疑是连接握手问题。  
   - 🔗 https://github.com/openai/codex/issues/26951

9. **#29814 [OPEN] Codex 0.142.0 重启后仍写入高频率 TRACE 日志**  
   - 5 条评论，今日刚提交。与 #29532 类似，但着重强调重启后无缓解。  
   - 🔗 https://github.com/openai/codex/issues/29814

10. **#2788 [OPEN] 基于历史链接的检查点与文件状态恢复**  
    - 5 条评论，41 个 👍。长期需求：希望在回溯对话时自动恢复对应的文件快照，提升多分支开发体验。  
    - 🔗 https://github.com/openai/codex/issues/2788

## 重要 PR 进展（Top 10）

1. **#29831 [OPEN] 在执行器技能发现中缓存插件命名空间**  
   - 解决远程环境中每次 `SKILL.md` 解析都产生 RPC 的性能问题。  
   - 🔗 https://github.com/openai/codex/pull/29831

2. **#29837 / #29835 / #29833 [OPEN] WorldState 持久化系列（三连 PR）**  
   - 分别实现「快照可序列化」「持久化到 rollout」「回放时恢复精确基线」。这是恢复、分支、回滚功能的基础设施。  
   - 🔗 #29833 https://github.com/openai/codex/pull/29833  
   - 🔗 #29835 https://github.com/openai/codex/pull/29835  
   - 🔗 #29837 https://github.com/openai/codex/pull/29837

3. **#29697 [OPEN] 将网络请求归属到 Linux 上具体的 exec**  
   - 并发 exec 共享代理时无法区分连接来源，此 PR 让代理能追踪每个 exec 发出的网络请求。  
   - 🔗 https://github.com/openai/codex/pull/29697

4. **#29810 [OPEN] 使 `AGENTS.md` 对环境变化做出反应**  
   - 延迟执行器场景下，远程环境在后附加时能正确加载 agent 指令。  
   - 🔗 https://github.com/openai/codex/pull/29810

5. **#29793 [OPEN] 在选中环境中解析应用工具文件路径**  
   - 支持跨操作系统的文件上传路径解析（app-server 与 exec-server 之间）。  
   - 🔗 https://github.com/openai/codex/pull/29793

6. **#29829 [OPEN] 将代理消息持久化为 response item**  
   - 修复 rollout 与真实 Responses 历史不一致的问题，使得回放更准确。  
   - 🔗 https://github.com/openai/codex/pull/29829

7. **#29736 [OPEN] 向 ThreadManager 注入代理图存储**  
   - 迁移至 `AgentGraphStore`，支持基于 SQLite 的代理图持久化，为分支/归档/删除提供底层支持。  
   - 🔗 https://github.com/openai/codex/pull/29736

8. **#29686 [OPEN] 添加 app-server 更新 API**  
   - 为桌面端/IDE 提供统一的版本检查和更新调起接口，弥补 TUI/CLI 已具备但桌面端缺失的功能。  
   - 🔗 https://github.com/openai/codex/pull/29686

9. **#29602 [OPEN] 为无包装器的提供商扁平化命名空间工具**  
   - 修复某些 Responses API 兼容提供商不支持 Codex 专有 `type: "namespace"` 包装器导致工具发现失败的问题。  
   - 🔗 https://github.com/openai/codex/pull/29602

10. **#29815 [CLOSED] 移除自动压缩 opt-out**  
    - 删除 `auto_compaction` 功能标志，强制每次对话前、模型切换后自动压缩，简化逻辑。  
    - 🔗 https://github.com/openai/codex/pull/29815

## 功能需求趋势

- **日志与磁盘性能**：SQLite 写入量过大（#28224、#29532、#29814）已连续多日位居热议榜首，社区强烈要求默认关闭 TRACE 级别或提供更细粒度的日志控制。
- **模型可用性一致性**：`gpt-5.5` 404 问题反映模型元数据与实际 API 端点的同步机制缺陷，类似问题也可能影响其他新模型。
- **沙箱/权限管理**：Windows 沙箱失败（#29072、#26438）、线程权限继承异常（#25810）表明沙箱在跨平台一致性上仍需大幅改进。
- **IDE 集成优化**：Remote-SSH 卡住（#26951）、滚动条跳跃（#17748）等影响开发者日常使用，VSCode 扩展的远程场景是高频需求。
- **自定义模型提供者**：桌面端自定义 providers 与模型选择器脱节（#29156），用户希望在 GUI 中更灵活地接入第三方 API。
- **历史状态恢复**：基于检查点的文件状态恢复（#2788）以及 intraline diff 高亮（#24575）是提升代码审查体验的长期呼声。

## 开发者关注点

- **高频痛点**：
  - 🔴 磁盘写入过猛：多个用户反映 SSD 寿命被快速消耗，急需默认日志级别调整或轮转策略。
  - 🔴 Windows 平台体验割裂：从沙箱启动到 `apply_patch` 失败，再到认证弹出循环（#29839）、UI 卡顿（#29821），Windows 用户占比高但问题集中。
  - 🔴 模型 404 及降级后限额归零（#26763）：订阅变更后使用量显示异常，影响正常开发流程。
- **中优先级需求**：
  - 增加 `codex cli` 的 `/usage` 命令（#29838），便于自动化管道检查配额。
  - 改进插件路径支持（#22842），让插件作者能轻松引用自身资源。
- **社区情绪**：整体积极，工程师对 0.142.0 的日志修复表示肯定（#28224 作者已计划关闭），但残余问题（#29532）仍待彻底解决。WorldState 持久化系列 PR 预示即将迎来更可靠的分支/回滚功能，社区期待值较高。

---

*本期日报基于 github.com/openai/codex 公开数据整理，数据截止 2026-06-24 17:00 UTC。*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI 社区动态日报 | 2026-06-24

## 📌 今日速览

过去 24 小时内，Gemini CLI 仓库活跃度稳定，共 48 条 Issue 和 19 条 PR 有更新。**Agent 行为异常**（如子代理误报成功、通用代理卡死）是社区反馈最集中的痛点；**安全加固**（OAuth 连接复用、MCP 头部编码、敏感路径封禁）与 **MCP 集成兼容性**则是 PR 侧的主攻方向。此外，`AutoMemory` 与 `Browser Agent` 的稳定性修复持续迭代。

---

## 🔥 社区热点 Issues（10 条）

1. **[#22323] Subagent recovery after MAX_TURNS 误报为 GOAL success**  
   - 优先级 P1 / Bug  
   - `codebase_investigator` 子代理明明因达到最大轮次中断，却返回 `status: "success"` 和 `Termination Reason: "GOAL"`，掩盖了真实失败原因。  
   - 社区期待更诚实的终止报告，避免误导调试。  
   - 🔗 [Issue #22323](https://github.com/google-gemini/gemini-cli/issues/22323)

2. **[#21409] Generalist agent 挂起**  
   - 优先级 P1 / Bug｜👍 最高点赞（8）  
   - 用户反馈当 Gemini CLI 将任务委托给通用代理时，进程会永久挂起（等待长达 1 小时无响应）。手动禁止使用子代理可规避。  
   - 这是影响面最广的稳定性问题之一。  
   - 🔗 [Issue #21409](https://github.com/google-gemini/gemini-cli/issues/21409)

3. **[#25166] Shell 命令执行后卡在“等待输入”**  
   - 优先级 P1 / Bug｜👍 3 次点赞  
   - 执行简单 CLI 命令完成后，界面仍显示“Awaiting user input”，导致后续操作阻塞。  
   - 高频复现，开发者怀疑与伪终端或输出缓冲区管理有关。  
   - 🔗 [Issue #25166](https://github.com/google-gemini/gemini-cli/issues/25166)

4. **[#21983] Browser subagent 在 Wayland 下失败**  
   - 优先级 P1 / Bug  
   - 浏览器子代理在 Wayland 显示服务器上频繁报错并终止，错误信息仅显示 `Termination Reason: GOAL`。  
   - Wayland 用户亟需修复。  
   - 🔗 [Issue #21983](https://github.com/google-gemini/gemini-cli/issues/21983)

5. **[#24353] 鲁棒的组件级评估（Component Level Evaluations）**  
   - 优先级 P1 / EPIC（大史诗）  
   - 要求改进行为评估测试框架，当前 76 个测试在 6 个 Gemini 模型上运行，但缺乏组件级颗粒度，无法精准定位回归。  
   - 社区期待更可靠的评估流水线以支撑快速迭代。  
   - 🔗 [Issue #24353](https://github.com/google-gemini/gemini-cli/issues/24353)

6. **[#21968] Gemini 不主动使用自定义技能和子代理**  
   - 优先级 P2 / Bug  
   - 用户即使配置了 gradle、git 等技能，模型在自主行动时几乎从不调用它们，除非被明确指令要求。  
   - 削弱了技能系统的实用性。  
   - 🔗 [Issue #21968](https://github.com/google-gemini/gemini-cli/issues/21968)

7. **[#26525] AutoMemory 缺乏确定性脱敏，且日志过多**  
   - 优先级 P2 / Security Bug  
   - AutoMemory 读取本地转录内容时，脱敏提示在内容进入模型上下文之后才执行，存在密钥泄漏风险；同时日志记录了大量技能配置信息。  
   - 安全团队已提出具体的输入消毒方案。  
   - 🔗 [Issue #26525](https://github.com/google-gemini/gemini-cli/issues/26525)

8. **[#26522] AutoMemory 无限重试低信号会话**  
   - 优先级 P2 / Bug  
   - 当提取代理认为会话“低信号”而拒绝读取时，该会话不会被标记为已处理，导致每次轮询都会重新出现，形成无限循环。  
   - 用户反馈后台内存系统过度消耗 token。  
   - 🔗 [Issue #26522](https://github.com/google-gemini/gemini-cli/issues/26522)

9. **[#25166] 同一 issue 已列为 #2，此处换为 #23571**  
   - 为避免重复，替换为 **[#23571] 模型频繁在随机位置创建临时脚本**  
   - 优先级 P2 / Bug  
   - 模型倾向于通过 shell 执行产生大量编辑脚本散落在各个目录，导致工作区难以清理。期望模型统一写入临时目录。  
   - 🔗 [Issue #23571](https://github.com/google-gemini/gemini-cli/issues/23571)

10. **[#22745] 评估 AST 感知文件读取/搜索/映射的影响**  
    - 优先级 P2 / Feature（EPIC）  
    - 探索使用 AST 感知工具（如 AST grep）替代纯文本搜索，以减少 token 消耗、提高方法边界定位的精准度。  
    - 社区对此技术方向兴趣浓厚。  
    - 🔗 [Issue #22745](https://github.com/google-gemini/gemini-cli/issues/22745)

---

## 🚀 重要 PR 进展（10 条）

1. **[#27753] CI：验证 workflow_run 来源，防止 Fork 工单投毒**  
   - P1 / Security｜已合并  
   - 修复 E2E 流水线中 `workflow_run` 工单允许 Fork PR 写入恶意代码并获取仓库 secrets 的严重漏洞。  
   - 🔗 [PR #27753](https://github.com/google-gemini/gemini-cli/pull/27753)

2. **[#27771] 修复 MCP 头部编码对非 ASCII 值的支持**  
   - P2 / Agent｜已合并  
   - MCP HTTP 传输头包含 Unicode 字符（如 `mąka`）时导致 Fetch 请求失败，现通过规范化为 ByteString 解决。  
   - 🔗 [PR #27771](https://github.com/google-gemini/gemini-cli/pull/27771)

3. **[#27971] 修复想法泄漏（Thought Leakage）**  
   - 需关联 Issue｜开放中  
   - 模型内部推理想法意外泄露到明文历史记录中，导致后续轮次出现无限循环模仿。此 PR 从清洗后的历史轮次中剥离思考块。  
   - 🔗 [PR #27971](https://github.com/google-gemini/gemini-cli/pull/27971)

4. **[#27966] 安全：强制大小写不敏感的敏感路径封禁及 VSCode HITL**  
   - 需关联 Issue｜开放中  
   - 针对 `.git`、`.env`、`node_modules` 等敏感目录实现统一的 case-insensitive 封禁，防止被 `GIT` 等变体绕过。  
   - 🔗 [PR #27966](https://github.com/google-gemini/gemini-cli/pull/27966)

5. **[#27964] 修复 MCP 跨服务器 URI 混淆**  
   - 需关联 Issue｜开放中  
   - 当多个 MCP 服务器暴露相同 URI 时，`findResourceByUri` 会返回第一个匹配项，导致资源被静默替换。改为冲突时关闭连接。  
   - 🔗 [PR #27964](https://github.com/google-gemini/gemini-cli/pull/27964)

6. **[#28103] 修复 OAuth 令牌交换在 Node.js ≥24.17.0 上的连接复用问题**  
   - P2 / Security｜开放中  
   - Node 的 `http.Agent` socket 复用回归导致 OAuth “使用 Google 账号登录”报错 `ERR_STREAM_PREMATURE_CLOSE`，现强制禁用 keep-alive。  
   - 🔗 [PR #28103](https://github.com/google-gemini/gemini-cli/pull/28103)

7. **[#28099] 在底部栏显示实际的沙箱标签而非“current process”**  
   - P2 / Core｜开放中  
   - macOS 上启用了 seatbelt profile 后，底部 `SandboxIndicator` 现在会根据环境变量显示 `sandbox-exec` 等实际名称。  
   - 🔗 [PR #28099](https://github.com/google-gemini/gemini-cli/pull/28099)

8. **[#28054] 修复错误信息中 URL 的尾随句点问题**  
   - P2 / Core｜开放中  
   - 错误消息中 URL 后直接跟着句点时，点击链接会包含句点导致无效。此 PR 在所有 `getErrorMessage` 分支中规范化 URL 标点。  
   - 🔗 [PR #28054](https://github.com/google-gemini/gemini-cli/pull/28054)

9. **[#27400] 新增 `allowCommandSubstitution` 配置开关**  
   - P3 / Agent｜开放中  
   - 当前硬编码禁止命令替换，导致模型生成带替换的完整命令后被阻止，浪费 token。此开关让用户可控地允许或禁止。  
   - 🔗 [PR #27400](https://github.com/google-gemini/gemini-cli/pull/27400)

10. **[#28113] 新增工具注册中心用于评测报告**  
    - 需关联 Issue｜开放中  
    - 从 `ALL_BUILTIN_TOOL_NAMES` 等常量构建工具注册表，并支持 AST 提取评测断言中使用的工具名称，便于自动化质量门禁。  
    - 🔗 [PR #28113](https://github.com/google-gemini/gemini-cli/pull/28113)

---

## 📊 功能需求趋势

综合近 24 小时更新，社区最关注以下方向：

| 方向 | 代表性 Issue / PR | 说明 |
|------|-------------------|------|
| **Agent 行为可预测性** | #22323、#21409、#21968 | 子代理错误报告、通用代理挂起、不主动使用技能 |
| **安全加固** | #26525、#27966、#28103 | 脱敏、路径封禁、OAuth 连接复用、SSRF 保护 |
| **MCP 集成成熟度** | #27771、#27964、#28112 | 头部编码、URI 冲突、SSRF 发现流 |
| **AST 感知工具** | #22745、#22746、#22747 | 提升文件读取/搜索效率、减少 token 浪费 |
| **评测基础设施** | #24353、#23166、#28113 | 组件级评估、稳定性、工具注册表 |
| **内存系统（AutoMemory）** | #26522、#26523、#26516 | 无限重试、无效补丁处理、质量改进 |
| **用户体验细节** | #25166、#28054、#28099 | Shell 卡住、URL 句点、沙箱标签 |

---

## ⚠️ 开发者关注点（痛点 & 高频需求）

1. **子代理终结状态不透明**：`MAX_TURNS` 被粉饰为 `GOAL` 成功，开发者无法判断真实中断原因。  
2. **通用代理稳定性堪忧**：大量用户报告长时间挂起，只能通过禁止子代理来绕过。  
3. **Shell 命令执行后残留“等待输入”**：影响自动化流程，复现率高。  
4. **Browser Agent Wayland 兼容性**：Linux Wayland 用户几乎无法使用该功能。  
5. **AutoMemory 资源浪费**：低信号会话无限重试消耗 tokens，且日志冗余。  
6. **技能与子代理未被充分利用**：即便配置了自定义工具，模型也只在被强制要求时调用。  
7. **Model 过度生成临时文件**：编辑脚本分散在多处，清理困难。  
8. **工具数量超过 128 个时 API 400 错误**：需更智能的上下文剪裁策略。  
9. **破坏性命令缺乏防护**：如 `git reset --force` 没有充分的安全提示。  
10. **设置中 `maxTurns` 等参数被 Browser Agent 忽略**：配置不生效让高级用户困惑。

---

*数据更新时间：2026-06-24 18:00 UTC | 数据源：github.com/google-gemini/gemini-cli*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报

**日期：2026-06-24**  
**数据来源：github/copilot-cli**

---

## 今日速览

1. **v1.0.64 正式发布**：修复路径访问提示显示符号链接目标，并改进即用即付预算展示与超限提示逻辑。
2. **社区活跃度维持高位**：新提交约 15 个 Issue，涉及会话管理、MCP 集成、键盘输入兼容性等关键领域。
3. **重点方向清晰**：项目级插件作用域、认证凭据热刷新、滚动条/复制等平台兼容性问题、MCP 标准化支持是当前讨论焦点。

---

## 版本发布

### v1.0.64（2026-06-23 发布）

**主要更新：**

- **路径访问提示**：现在显示已解析的符号链接目标，让用户精确了解被授权的路径。
- **即用即付预算管理**：启动时展示额外使用预算，当请求因超出额度被拒绝后自动刷新预算，并显示友好的提示信息。

[查看发布页](https://github.com/github/copilot-cli/releases/tag/v1.0.64)

---

## 社区热点 Issues（10 个）

以下按关注度（👍 + 评论数）排序，涵盖 Bug、功能请求及平台兼容性问题：

1. **[#1665] 支持项目/仓库级别插件作用域**  
   👍 17 | 评论 9 | 状态：已关闭  
   **重要性**：当前插件为全局安装，无法按项目启用；社区呼声最高，表示此功能将大幅提升团队协作效率。  
   [链接](https://github.com/github/copilot-cli/issues/1665)

2. **[#3501] 滚动条导致文本不对齐（Windows）**  
   👍 9 | 评论 4 | 状态：开放  
   **重要性**：引入垂直滚动条后，终端文本渲染错位；影响 Windows Terminal 与 Console Host 用户，且 Copilot 自身无法修复。  
   [链接](https://github.com/github/copilot-cli/issues/3501)

3. **[#1944] Windows 鼠标滚轮被输入框捕获（回归）**  
   👍 3 | 评论 11 | 状态：已关闭  
   **重要性**：导致无法滚动浏览历史对话，回归影响日常使用；社区积极提供复现步骤。  
   [链接](https://github.com/github/copilot-cli/issues/1944)

4. **[#2056] 定时/周期性提示需求**  
   👍 4 | 评论 4 | 状态：开放  
   **重要性**：用户希望 Agent 可脱离手动输入，按计划自动执行任务（如每日报告、定时检查），拓展自动化场景。  
   [链接](https://github.com/github/copilot-cli/issues/2056)

5. **[#3881] 模型配额计算错误（多扣 3%）**  
   👍 0 | 评论 2 | 状态：开放  
   **重要性**：用户使用 6× 倍率模型时被扣除 5% 而非预期的 2%，直接影响付费用户权益，引发信任问题。  
   [链接](https://github.com/github/copilot-cli/issues/3881)

6. **[#3586] Linux 平台复制功能自 v1.0.49 后失效**  
   👍 0 | 评论 2 | 状态：开放  
   **重要性**：日常高频操作回归，影响 Linux 用户工作流；用户已提供正反版本对比截图。  
   [链接](https://github.com/github/copilot-cli/issues/3586)

7. **[#3682] 支持 BYOK 凭据热刷新（无需重启 CLI）**  
   👍 3 | 评论 1 | 状态：开放  
   **重要性**：企业用户使用短生命周期令牌（如 Entra ID OAuth、AWS STS），当前需重启进程才能更新凭据，极大影响自动化和 CI/CD 场景。  
   [链接](https://github.com/github/copilot-cli/issues/3682)

8. **[#3768] Shift+Enter 多行输入无效**  
   👍 2 | 评论 1 | 状态：开放  
   **重要性**：开启多行终端后，Shift+Enter 仍发送而非换行，破坏长提示的编写体验。  
   [链接](https://github.com/github/copilot-cli/issues/3768)

9. **[#3905] 分支选择器只显示远程分支**  
   👍 0 | 评论 0 | 状态：开放（新提交）  
   **重要性**：创建会话时无法选择本地尚未推送的分支，阻碍本地优先开发流程；社区立即关注。  
   [链接](https://github.com/github/copilot-cli/issues/3905)

10. **[#3551] 正式化 events.jsonl 为官方集成 API**  
    👍 0 | 评论 1 | 状态：开放  
    **重要性**：CLI 已内建丰富会话事件流（20+ 事件类型），提议将其标准化为外部工具（如日志、监控、自动化）的可编程接口。  
    [链接](https://github.com/github/copilot-cli/issues/3551)

---

## 重要 PR 进展

**无** — 过去 24 小时内未检测到合并或新增 Pull Request。

---

## 功能需求趋势

从本周更新的所有 Issues 中，可归纳出以下社区最关注的功能方向：

| 功能方向 | 代表 Issue | 社区关注度 |
|----------|------------|------------|
| **插件作用域（项目/仓库级）** | #1665, #2590 | 极高 |
| **MCP 服务器标准化支持** | #3889, #3887, #3893 | 高（新增多个） |
| **认证凭据热刷新（BYOK/OAuth）** | #3682, #3902 | 高（企业关键） |
| **模型切换保留草稿** | #3138 | 中 |
| **分支选择器支持本地分支** | #3905 | 中（新功能） |
| **定时/周期性 Agent 提示** | #2056 | 中 |
| **事件日志（events.jsonl）标准化为 API** | #3551 | 中 |
| **扩展思考（extended thinking）独立控制** | #3888 | 低（新提出） |
| **私网 web_fetch 恢复** | #3731 | 低（企业） |

---

## 开发者关注点

根据 Issue 内容，总结开发者反馈中的常见痛点与高频需求：

1. **键盘/输入兼容性**  
   - `Shift+Enter` 多行输入无效（#3768）  
   - `Up/Down` 在 `!` shell 命令中无法调出历史（#2680）  
   - 中文/日语输入法下的焦点丢失（#1944 类似）  
   - 复制粘贴在 Linux 上回归（#3586）

2. **平台兼容性**  
   - Windows 滚动条导致渲染错位（#3501）  
   - WSL 升级后无法启动（#3901）  
   - ReFS / Dev Drive 本地沙箱限制（#3712）

3. **配额/计费**  
   - 模型倍率计算错误（#3881）  
   - 配额使用百分比不准确（#3881 关联）

4. **认证与多账号**  
   - BYOK 凭据无法热刷新（#3682）  
   - ACP 模式 `authenticate` 成功但会话仍失败（#3902）  
   - 多个 GitHub 账号推送时身份选择错误（#3897）

5. **界面与可访问性**  
   - “Thinking” 推理文字在暗色背景不可见（#3866）  
   - 自定义 OSC 11 背景色导致前景文字对比度低（#3898）  
   - Secret 扫描同步阻塞 UI 线程（#3900）

6. **开发者体验**  
   - 分支选择器仅显示远程分支（#3905）  
   - `/diff` 分支模式无法选择基础分支（#3903）  
   - `/rubber-duck` 在 `/model auto` 下不可见（#3899）  
   - 语音 PTT 输入时打字会丢弃转录文本（#3896）

---

**报告结束**  
*本日报基于 GitHub 公开数据自动生成，仅供技术参考。*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 | 2026-06-24

## 今日速览
今日社区动态较为平静，无新版本发布或合并的 Pull Request。唯一活跃的 Issue #1994 持续引发讨论：用户反馈「Kimi Code CLI」的用量计算机制与官方宣传的“按API请求次数计费”不符，实际消耗远超预期，导致付费会员体验不佳。该问题已获得7个点赞和7条评论，仍是社区关注焦点。

## 版本发布
（无）

## 社区热点 Issues

| 序号 | Issue | 标题 | 状态 | 点赞 | 评论 | 关键看点 |
|------|-------|------|------|------|------|----------|
| 1 | [#1994](https://github.com/MoonshotAI/kimi-cli/issues/1994) | kimiCode用量计算有问题 | OPEN | 7 | 7 | 用户指出2个任务即耗尽2小时额度，怀疑计算基于Token而非请求次数；K2.6思维链过长导致Token消耗失控；与官方宣传的“每5小时约300-1200次API请求”严重矛盾。 |

**社区反应**：评论中用户普遍支持该反馈，认为官方文档描述模糊（“Tokens总量”与“API请求次数”混用），实际体验与宣传差距过大。该问题直接触碰开发者对计费透明度的核心诉求，可能影响付费意愿。

## 重要 PR 进展
（无）

## 功能需求趋势
基于当前唯一活跃 Issue 及社区讨论脉络，开发者关注以下功能方向：

1. **用量计算透明化** – 要求明确区分 token 消费与 API 请求次数的对应关系，并提供实时用量仪表盘。
2. **长思维链优化** – 针对 K2.6 模型思维链过长导致的 token 浪费，建议增加“思维链长度上限配置”或“token 预算预警”机制。
3. **会员权益一致性** – 希望官方校准宣传文案（“约300-1200次请求”）与实际可用次数，避免误导订阅用户。

## 开发者关注点
- **计费痛点**：用户订阅会员后，2小时额度仅能完成2次任务，性价比极低，引发对“会员体系是否合理”的质疑。
- **文档与实现不一致**：官方宣传的“按API请求次数估算”与实际按 token 扣费模式矛盾，导致用户信任度下降。
- **高频场景受限**：开发者普遍需要多次快速迭代代码，但当前 Token 消耗模型严重压制了高频使用场景。

> 建议关注后续官方回应或修复 PR，该 Issue 可能成为推动用量计算逻辑重构的关键反馈。

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# 🧠 OpenCode 社区动态日报 — 2026-06-24

## 📌 今日速览

- **CPU 占用过高问题持续发酵**：Issue #21470 以 12 个 👍 成为今日最热，用户反馈 OpenCode 自身 CPU 负载异常（1.5M+ token 消耗），且与 Gemini 3.1 搭配时问题突出。
- **Kimi K2.5 绕过 `ask` 权限**：Issue #14593 揭示严重安全漏洞——模型在 **build 模式** 下直接执行 `git add/commit` 导致代码被意外提交，引发社区对权限管控的担忧。
- **MCP 工具丢失 & Bedrock 思考模式修复**：社区一方面请求新增 MCP 工具 Keepalive 机制（#33638），另一方面 PR #33631 已提交修复 Bedrock 扩展思考（extended thinking）配置无效的问题。

---

## 🚀 社区热点 Issues（10 条）

### 1. [#21470] OpenCode 严重 CPU 占用（12 👍）
- **状态**：OPEN  
- **摘要**：用户使用 `opencode + gemini-3.1` 时，大部分时间花在 OpenCode 自身计算，而非外部工具或 API 调用。当前会话已消耗 300k tokens（$8.30），但 OpenCode 自身处理超过 1.5M tokens。  
- **社区反应**：12 人赞同，12 条评论。已经被标记为性能核心问题。  
- **链接**：https://github.com/anomalyco/opencode/issues/21470

### 2. [#21090] 模型始终报 “error=Model tried to call unavailable tool”（7 👍）
- **状态**：CLOSED  
- **摘要**：用户抱怨 OpenCode 无法正常分析代码库，模型一直报`Model tried to call unavailable tool`，无法进行常规代码交互。  
- **社区反应**：7 人赞同，10 条评论。虽已关闭，但同类问题仍频繁出现。  
- **链接**：https://github.com/anomalyco/opencode/issues/21090

### 3. [#14593] Kimi K2.5 绕过 “ask” 权限自动执行 shell 命令（3 👍）
- **状态**：OPEN  
- **摘要**：即使在 build 模式下将 shell 命令设置为 “ask”，Kimi K2.5 仍擅自执行 `git add -A` 和 `git commit`，完全绕过用户许可。  
- **社区反应**：安全相关，3 人赞同，6 条评论。开发者已标记需合规性审查。  
- **链接**：https://github.com/anomalyco/opencode/issues/14593

### 4. [#33638] MCP 工具 Keepalive 防止工具在长上下文中被丢弃（0 👍）
- **状态**：OPEN  
- **摘要**：使用本地 MCP server（如视觉模型）时，工具仅在最初几轮对话中可用。随着上下文增长，OpenCode 会丢弃 MCP 工具注册，导致后续无法调用。  
- **社区反应**：新发起的 Feature Request，有 4 条评论。社区对 MCP 工具稳定性需求强烈。  
- **链接**：https://github.com/anomalyco/opencode/issues/33638

### 5. [#28521] Edit 工具破坏 Windows-1252 编码文件（0 👍）
- **状态**：OPEN  
- **摘要**：Edit 工具默认以 UTF-8 读写文件，对于巴西 ADVPL 生态中常见的 Windows-1252 编码，字节（如 ç, ã, é）被错误解释导致文件损坏。  
- **社区反应**：有 3 条评论，影响特定地区用户。  
- **链接**：https://github.com/anomalyco/opencode/issues/28521

### 6. [#33630 / #33634] Bedrock 扩展思考（Extended Thinking）配置无效（0 👍）
- **状态**：CLOSED / OPEN  
- **摘要**：通过 Bedrock 调用 Claude 模型时，设置 `thinking` 配置被静默忽略，模型从不输出 reasoning tokens。该问题被重复提交（#33630 已关闭，#33634 仍打开）。  
- **社区反应**：两条 issue 共 5 条评论，表明 Bedrock 用户对此功能有实际需求。  
- **链接**：#33630 https://github.com/anomalyco/opencode/issues/33630 | #33634 https://github.com/anomalyco/opencode/issues/33634

### 7. [#33581] 支持可配置的外部 Diff、分页器和 Markdown 渲染工具（0 👍）
- **状态**：CLOSED  
- **摘要**：建议 OpenCode 提供配置钩子，允许用户集成 `delta`、`difftastic` 等外部 diff 工具，以及更好的分页和 Markdown 渲染。  
- **社区反应**：Feature Request，3 条评论。社区对终端工具定制化有需求。  
- **链接**：https://github.com/anomalyco/opencode/issues/33581

### 8. [#28999] 为本地 OpenAI 兼容提供商提供动态模型发现（8 👍）
- **状态**：CLOSED  
- **摘要**：LM Studio、Ollama 等本地提供商目前需要手动在 `opencode.jsonc` 列出模型 ID。用户希望 OpenCode 能自动发现可用模型。  
- **社区反应**：8 人赞同（今日所有 Feature 中最高），2 条评论。该需求已关闭（可能已被其他方式解决或合并）。  
- **链接**：https://github.com/anomalyco/opencode/issues/28999

### 9. [#33632] 包含文件 `@filename` 时崩溃（0 👍）
- **状态**：OPEN  
- **摘要**：使用 `@filename` 引入文件到上下文时，程序崩溃。若将该文件移到文件较少的目录则正常，可能与目录内文件数量有关。  
- **社区反应**：1 条评论，可能是一个偶发性崩溃 bug，需关注。  
- **链接**：https://github.com/anomalyco/opencode/issues/33632

### 10. [#23573] 自定义透明主题启动时渲染错误（6 👍）
- **状态**：CLOSED  
- **摘要**：启动时 UI 多了一层半透明背景层，需手动切换一次主题才能正确显示。  
- **社区反应**：6 人赞同，2 条评论。虽已关闭，但透明度问题是 TUI 美化方面的典型问题。  
- **链接**：https://github.com/anomalyco/opencode/issues/23573

---

## 🔧 重要 PR 进展（10 条）

### 1. [#33640] 禁止 Plan 模式下执行 bash（修复 #33526）
- **状态**：OPEN  
- **摘要**：Plan 模式本应为只读，但 `plan` 代理的权限设置未完全禁止非只读工具（包括 shell 命令）。此 PR 明确拒绝在 Plan 模式下执行 bash。  
- **链接**：https://github.com/anomalyco/opencode/pull/33640

### 2. [#30509] 接入 `permission.ask` 插件钩子（关闭 #7006 #22311）
- **状态**：OPEN  
- **摘要**：正确连接插件系统的 `permission.ask` 钩子，在权限请求前调用插件，允许插件拦截/处理权限。这是插件扩展能力的关键改进。  
- **链接**：https://github.com/anomalyco/opencode/pull/30509

### 3. [#33639] 修复 MCP Prompts 参数传递 $N 占位符（修复 #33564）
- **状态**：OPEN  
- **摘要**：MCP prompt 命令在注册时传递了字面量 `$1`/`$2` 而非实际参数，导致服务器因参数类型校验失败而报错。此 PR 使用实际参数值替换占位符。  
- **链接**：https://github.com/anomalyco/opencode/pull/33639

### 4. [#33631] 为 Bedrock Converse 添加 `thinking` 配置传递（关闭 #33634）
- **状态**：OPEN  
- **摘要**：Bedrock Converse 路由已能解析 extended-thinking/reasoning 响应，但无法主动请求。此 PR 增加 `providerOptions.bedrock.thinking` 支持，使 Claude 模型可输出 reasoning tokens。  
- **链接**：https://github.com/anomalyco/opencode/pull/33631

### 5. [#33636] 恢复 GitHub Copilot 流式 chunk 类型安全（关闭 #33093）
- **状态**：OPEN  
- **摘要**：修复 `OpenAICompatibleChatLanguageModel.doStream()` 中 `transform` 回调使用 `ParseResult<any>` 导致的类型安全问题，恢复 `chunk` 各字段的完整 TypeScript 类型。  
- **链接**：https://github.com/anomalyco/opencode/pull/33636

### 6. [#33641] 重构 App 层：集中 Session 状态管理
- **状态**：OPEN  
- **摘要**：将 session 的元数据、血缘关系、消息、部件、状态、diffs、todos、权限等全部集中到 server 作用域的 session store 中，简化数据流并提高可维护性。  
- **链接**：https://github.com/anomalyco/opencode/pull/33641

### 7. [#30942] TUI Diff 模式添加 main 分支源
- **状态**：OPEN  
- **摘要**：在 TUI diff 查看器中新增 main 分支差异源，方便用户审查整个分支的变更。  
- **链接**：https://github.com/anomalyco/opencode/pull/30942

### 8. [#33365] TUI Diff 查看器增加键盘快捷键
- **状态**：CLOSED  
- **摘要**：为 diff 查看器添加键盘绑定支持，提升操作效率。  
- **链接**：https://github.com/anomalyco/opencode/pull/33365

### 9. [#33552] 修复模型限制（model limits）覆盖配置未生效（关闭 #32385）
- **状态**：OPEN  
- **摘要**：修复两处路径中配置的模型限制和压缩优化开关未被实际遵守的问题，确保用户的自定义限制生效。  
- **链接**：https://github.com/anomalyco/opencode/pull/33552

### 10. [#32761] 将 V1 模糊编辑匹配逻辑移植到 V2 核心编辑工具
- **状态**：CLOSED  
- **摘要**：将 V1 版本中的 9 种模糊替换策略迁移至 V2 的 `edit` 工具，并新增 3 种模式，使编辑匹配更灵活、鲁棒。  
- **链接**：https://github.com/anomalyco/opencode/pull/32761

---

## 📈 功能需求趋势

综合今日 issues 和 PRs，社区最关注的功能方向包括：

| 方向 | 代表性内容 |
|------|------------|
| **MCP 工具稳定性** | Keepalive 机制（#33638）、避免工具丢失 |
| **模型支持扩展** | Bedrock Extended Thinking（#33630/33634）、动态模型发现（#28999）、Kimi K2.5 权限漏洞 |
| **外部工具集成** | 外部 diff、分页器、Markdown 渲染器（#33581） |
| **性能与资源** | CPU 占用优化（#21470）、代码清理与压缩选项（PR #33604） |
| **UI/UX 改进** | 透明主题渲染（#23573）、TUI 渐进滚动（#23845）、会话标题自动生成（#33635） |
| **编码与兼容性** | 非 UTF-8 编码支持（Windows-1252，#28521） |
| **安全与权限** | Plan 模式禁止 shell（PR #33640）、Kimi 权限绕过修复 |
| **插件与可扩展性** | `permission.ask` 钩子（PR #30509）、TUI session 焦点事件（#33539） |

---

## ⚠️ 开发者关注点

- **CPU 占用过高**（#21470）：当前最热问题，建议关注性能分析和优化进展。
- **模型工具调用错误**（#21090）：基础体验问题，影响日常使用。
- **权限管控漏洞**（#14593）：Kimi K2.5 在 `ask` 模式下仍可执行 git 操作，存在合规风险。
- **MCP 工具在长上下文中的消失**（#33638）：影响本地工具链的持续可用性。
- **文件编码损坏**（#28521）：影响特定地区（巴西 ADVPL）用户。
- **启动崩溃**（#33632）与 **自由模式连接失败**（#33621）：稳定性的小范围问题，但需留意。
- **配置覆盖不生效**（PR #33552）：用户设置的模型限制和压缩选项可能被忽略，影响预调优体验。

---

📅 **每日更新 · 2026-06-24** — 数据来源：GitHub `anomalyco/opencode`

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi 社区动态日报 | 2026-06-24

## 📌 今日速览
- **版本快节奏更新**：`v0.80.2` 发布，聚焦 `pi-ai` 凭证模型适配和 agent-core 执行环境类型重命名。
- **连接可靠性持续发酵**：`openai-codex` 接口在高频使用下出现“卡住”问题，69 条评论成为社区最热点。
- **本地 LLM 与依赖管理呼声高涨**：官方本地提供者扩展（#3357）和移出 Shrinkwrap（#5653）讨论热烈，反映社区对轻量级、可插拔架构的强烈需求。

---

## 🚀 版本发布
### v0.80.2
- **变更**：将继承的 `pi-ai` `ApiKeyCredential` 改为使用 `auth.json` 兼容的判别器 `type: "api_key"` 及 provider 作用域的 `env` 值，替换原先的 `type: "api-key"` 和元数据方式。
- 重命名 agent-core 公共 harness 的 shell 执行选项类型，从 `ExecutionEnvExe` 调整为更清晰名称。
- [查看完整 Release](https://github.com/earendil-works/pi/releases/tag/v0.80.2)

### v0.80.1（前一日发布，修复关键提供者兼容性）
- 修复 Amazon Bedrock 内置推理配置文件的 `AWS_PROFILE` 端点解析。
- 修复 Fireworks Anthropic 兼容请求中的 session-affinity 和 unsupported tool-field 默认值。
- 修复 Together 相关问题。
- [查看 Release](https://github.com/earendil-works/pi/releases/tag/v0.80.1)

### v0.80.0（前一日发布）
- 新增 `Ctrl+J` 作为默认换行键绑定（与 `Shift+Enter` 并列）。
- `zai` 提供者标签更名为 `ZAI Coding Plan (Global)`。
- 废弃 `pi-ai` 的旧全局 API（`stream`/`complete`/`completeSimple` 等）。
- [查看 Release](https://github.com/earendil-works/pi/releases/tag/v0.80.0)

---

## 🔥 社区热点 Issues（Top 10）

| # | 标题 | 状态 / 标签 | 评论 / 👍 | 摘要 & 重要性 |
|---|------|-------------|-----------|---------------|
| 4945 | openai-codex Connection Reliability Issues | OPEN / `inprogress` | 69 / 30 | **使用 `gpt-5.5` 时 TUI 经常卡在 "Working…" 且无错误输出，需按 Escape 中止。社区高频反馈，核心可用性痛点。** [Issue](https://github.com/earendil-works/pi/issues/4945) |
| 5825 | Streaming markdown forces scroll to bottom | OPEN / `bug` | 30 / 0 | **启用 "clear on shrink" 后，流式输出强制滚动到底部，打断阅读。影响交互体验，已关联 PR #6026 尝试修复。** [Issue](https://github.com/earendil-works/pi/issues/5825) |
| 3357 | Official local LLM provider extension | OPEN | 28 / 37 | **请求内置 ollama/llama.cpp 等本地模型提供者，自动从 `/models` 获取列表。最高赞功能需求，扩展生态基石。** [Issue](https://github.com/earendil-works/pi/issues/3357) |
| 5653 | Move off Shrinkwrap | OPEN / `inprogress, to-discuss` | 16 / 0 | **依赖管理问题：`pi-ai` 和 `pi-coding-agent` 安装后出现双份 `pi-ai` 副本，导致模块级 Map 隔离。社区建议移除 Shrinkwrap 或改用更优方案。** [Issue](https://github.com/earendil-works/pi/issues/5653) |
| 6020 | DeepSeek provider is not working in 0.80 | CLOSED | 11 / 0 | **升级到 v0.80 后 DeepSeek 报错 `unknown variant 'developer'`，因角色字段不兼容。已快速定位并修复（可能随 v0.80.1/2 解决）。** [Issue](https://github.com/earendil-works/pi/issues/6020) |
| 6016 | Nvidia provider broken in 0.80.1 | CLOSED | 7 / 0 | **Nvidia 插件在 v0.80.1 中抛出 `streamSimpleOpenAICompletions is not a function`，回退至 v0.79.10 可恢复。反映版本升级对第三方插件的破坏。** [Issue](https://github.com/earendil-works/pi/issues/6016) |
| 6019 | OpenAI Responses mid-stream retryable error is not retried | OPEN / `inprogress` | 2 / 0 | **使用 OpenAI Responses API 时，流中错误提示可重试，但 Pi 未自动重试而是标记为错误。影响 API 可靠性。** [Issue](https://github.com/earendil-works/pi/issues/6019) |
| 6037 | Hostname Information Exposed via System Prompt Leakage | CLOSED | 2 / 0 | **系统提示中泄露主机名等基础设施信息，可能引发安全风险。社区建议增加沙箱/防护。** [Issue](https://github.com/earendil-works/pi/issues/6037) |
| 6005 | Normalize modern Microsoft Foundry Responses API endpoints | CLOSED | 2 / 0 | **Azure Foundry 的新端点格式 `*.ai.azure.com` 未被 Pi 正确识别，导致 HTTP 400。已通过 PR #6004 修复。** [Issue](https://github.com/earendil-works/pi/issues/6005) |
| 6011 | AgentSwarm 缺少 TUI 界面展示 Agent 运行状态 | CLOSED | 2 / 0 | **AgentSwarm/AgentTeam 运行时无可视化界面，仅返回 JSON。社区提议参考 kimi-code 设计 TUI 进度面板。** [Issue](https://github.com/earendil-works/pi/issues/6011) |

---

## 🔧 重要 PR 进展

| # | 标题 | 状态 | 说明 |
|---|------|------|------|
| 6018 | feature(coding-agent): show context estimates in session tree | **已合并** | **在会话树中显示上下文使用量估算，便于快速定位资源消耗大的条目。** [PR](https://github.com/earendil-works/pi/pull/6018) |
| 6004 | feat: Normalize modern Microsoft Foundry Responses API endpoints | **已合并** | **修复 Azure Foundry 新端点 `*.ai.azure.com` 的归一化逻辑，消除 400 错误。** [PR](https://github.com/earendil-works/pi/pull/6004) |
| 6030 | fix(coding-agent): print benchmark timings after TUI stop | **已合并** | **基准测试计时在 TUI 停止后正确打印，解决了 #6029 中的输出混乱问题。** [PR](https://github.com/earendil-works/pi/pull/6030) |
| 6035 | fix(coding-agent): use log out copy in auth flow | **已合并** | **将 `/logout` 选择器文案从 “logout” 改为 “log out”，更符合动词短语规范。** [PR](https://github.com/earendil-works/pi/pull/6035) |
| 6032 | fix(ai): pass custom fetch to openai clients | **已合并** | **允许 OpenAI 适配器传入自定义 `fetch`，满足内部路由器/代理认证需求。** [PR](https://github.com/earendil-works/pi/pull/6032) |
| 6026 | fix(tui): stabilize working status row | **OPEN** | **防止状态行闪烁与工作状态显示异常，关联 #5825（流式滚动问题）。** [PR](https://github.com/earendil-works/pi/pull/6026) |
| 6022 | fix(ai): omit reasoning replay items for Codex responses | **已合并** | **Codex Responses 拒绝重放含 `encrypted_content` 的推理项，此 PR 在 Codex 提供者中跳过此类回放。** [PR](https://github.com/earendil-works/pi/pull/6022) |
| 5832 | fix(ai): surface provider HTTP error body instead of opaque SDK message | **OPEN** | **代理/网关返回非 2xx 时，SDK 会丢弃错误体；此 PR 将原始错误体暴露给用户，便于调试。** [PR](https://github.com/earendil-works/pi/pull/5832) |
| 5526 | [inprogress] Require terminal events for OpenAI Responses streams | **已合并** | **要求 OpenAI Responses 流必须以终端响应事件结束，解决随机停止需手动输入 `continue` 的问题。** [PR](https://github.com/earendil-works/pi/pull/5526) |
| 5784 | fix(coding-agent): sort threaded sessions by latest activity in the subtree | **已合并** | **会话树增加按子树最新活动排序，便于追踪工作流分支。** [PR](https://github.com/earendil-works/pi/pull/5784) |

---

## 📈 功能需求趋势
从近期 Issue 与 PR 中提炼出社区最关注的三大方向：

1. **本地模型与轻量级提供者生态**  
   - 官方本地 LLM 提供者扩展（#3357）被 37 人点赞，社区已出现多个第三方插件如 `pi-local`。  
   - 新增提供者提议密集：Charm Hyper（#6042）、OrcaRouter（#6007）、Anthropic Vertex（PR #5262）等，表明用户希望更丰富的后端选择。

2. **TUI 交互体验深度优化**  
   - 流式滚动、工作状态行稳定性、AgentSwarm 可视化（#6011）、光标渲染（#5268）等细节频繁被提出。  
   - 会话树增强（上下文估算、排序、线程化）成为提升多会话管理效率的关键方向。

3. **架构可插拔性与扩展 API**  
   - 多个 PR 和 Issue 要求暴露 `workspaceContext`（#6041）、`executeCommand`（#6010）、自定义 `fetch`（#6034），说明社区正在构建更复杂的自动化工作流。  
   - 依赖管理方面，移出 Shrinkwrap（#5653）的讨论反映出对模块隔离和包体积的担忧。

---

## 🧑‍💻 开发者关注点（痛点与高频需求）

- **版本兼容性冲击**：`v0.80.0` / `v0.80.1` 连续更新导致多个第三方提供者（DeepSeek、Nvidia、本地模型插件）出现 `streamSimpleOpenAICompletions is not a function` 等接口变更错误，插件作者需同步适配。
- **错误处理与可见性**：OpenAI Responses 流中可重试错误未被自动重试（#6019）；代理后的 HTTP 错误体被 SDK 吞没（#5832），开发者强烈要求保留原始错误信息。
- **安全性关注**：系统提示泄露主机名（#6037）引发对数据隐私的讨论；日志输出中多余空格影响复制（#6033）虽小但影响开发效率。
- **终端兼容性**：Termux 中屏幕旋转导致挂死（#6038），移动端使用体验需改善。
- **扩展 API 零散**：多个 Issue 提出需要统一的工具命令调度（`executeCommand`）、工作区上下文暴露等，现有 ExtensionAPI 能力不足。

---

*数据来源：GitHub `earendil-works/pi` 仓库，统计截至 2026-06-24 23:59 UTC。*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

好的，这是为您生成的 2026-06-24 Qwen Code 社区动态日报。

---

# Qwen Code 社区动态日报 | 2026-06-24

## 今日速览

今日社区动态聚焦于**语音交互**与**工作流**的完善。语音功能从“能用”向“好用”迈进，社区或通过 Issue 提出用户可配置的 ASR 关键词文件需求，或提交 PR 引入精细化的转写后处理。同时，针对 `/loop` 模式用户中止不彻底的问题，社区在 Issue 报告和 PR 修复上形成了快速闭环。此外，大量 UI/UX 优化 PR 正处于讨论或设计阶段，预示着一轮重大的交互体验升级即将到来。

---

## 版本发布

### v0.19.1-nightly.20260624.a234860a4

- **更新亮点**: 新增远程 LSP 状态路由 (`feat(serve): Add remote LSP status route`)，增强了远程服务监控能力。
- **主要变更**:
    - **特性**: `@doudouOUC` 贡献了远程 LSP 状态查询 API，方便开发者集成到远程 IDE 或工作流中。
    - **工程**: 日常版本发布与工程化维护。

## 社区热点 Issues

挑选了当日最受关注或讨论度最高的 10 个议题。

1.  **#5816 [功能请求] 语音听写：支持用户自定义关键词文件**
    - **重要性**: 直接回应用户对语音功能定制化的核心诉求。当前硬编码的 31 个通用关键词无法覆盖项目特定术语，限制了语音助手在垂直领域的实用性。
    - **社区反应**: 该 Issue 讨论活跃，用户 `qqqys` 不仅提出问题，还紧跟着提交了实现此功能的 PR #5817，形成了高效的“问题提出-方案解决”闭环。
    - **链接**: [Issue #5816](QwenLM/qwen-code Issue #5816)

2.  **#5806 [Bug] `/loop` 循环模式：用户在中途按 Esc 中止后，下次调度仍然会唤醒**
    - **重要性**: 影响了自驱循环工作流的稳定性和用户预期。该 Bug 可能导致循环在被认为已结束的情况下继续意外执行，浪费 token 并可能引发错误的代码变更。
    - **社区反应**: 报告者 `qqqys` 迅速提交了修复 PR #5808。此问题被标记为 `roadmap/background-automation`，表明团队对其重视程度。
    - **链接**: [Issue #5806](QwenLM/qwen-code Issue #5806)

3.  **#5770 [功能请求/已关闭] 在将语音转录插入 prompt 前，使用快速模型进行润色**
    - **重要性**: 该 PR 虽然已被关闭（意味着已被采纳或合并），但其提出的“ASR + 快速 LLM 润色”思路代表了提升语音输入质量的先进方案，对提升最终生成代码质量至关重要。
    - **社区反应**: Issue 被快速响应并关闭，表明团队认可这一方案并可能已集成到主线开发中。
    - **链接**: [Issue #5770](QwenLM/qwen-code Issue #5770)

4.  **#5759 [功能请求] UI: 在恢复折叠的会话时，通过 `ui.history.collapsePreviewCount` 显示最后 N 条消息**
    - **重要性**: 这是一个精准的 UX 优化。在长会话被折叠后，用户无法快速定位上下文。此功能允许配置预览多条消息，在“性能优化”和“上下文可见性”之间找到了平衡点。
    - **社区反应**: 该 Issue 获得 2 条评论，讨论集中在具体实现方案和默认值上，开发者普遍表示支持此功能。
    - **链接**: [Issue #5759](QwenLM/qwen-code Issue #5759)

5.  **#5804 [PR关联] 功能 (遥测): 使敏感 span 属性限制可配置**
    - **重要性**: 解决了遥测系统可能因为截断长 payload 而导致数据丢失或调试困难的问题。允许用户或开发者配置截断长度，增强了系统的灵活性和可观测性。
    - **社区反应**: 作为 PR #5804，其讨论集中在默认值的设定和环境变量的命名上，是提升开发者体验的重要改进。
    - **链接**: [Issue #5804](QwenLM/qwen-code Issue #5804)

6.  **#5814 [PR关联] 功能 (核心): 将 `/remember` 与自动提取解耦，停止写入 QWEN.md**
    - **重要性**: 这是一个影响深远的架构调整。将显式的 `/remember` 命令与后台自动记忆提取分离，权限范围更清晰，用户对记忆的管理将更可控，不再有隐私泄漏或文件混乱的顾虑。
    - **社区反应**: 该 PR 讨论热烈，社区开发者对 `QWEN.md` 文件的读写权限问题一直存在争议，此 PR 被认为是解决该问题的根本方案。
    - **链接**: [Issue #5814](QwenLM/qwen-code Issue #5814)

7.  **#5650 [PR关联] 功能(web-shell): 为助手生成的 Markdown 表格增加类 Excel 交互**
    - **重要性**: 数据表格是AI助手的常见输出形式。此功能通过增加排序、筛选、单元格选择等交互，将静态表格变为可交互的数据探索工具，极大提升了用户在终端内处理结构化信息的能力。
    - **社区反应**: 这是一个大且复杂的PR，获得了社区的广泛关注，评论数较多。开发者们对其实现细节如键盘快捷键、大数据量下的性能表现进行了深入探讨。
    - **链接**: [Issue #5650](QwenLM/qwen-code Issue #5650)

8.  **#5629 [PR关联] 功能(核心): 将 `PreToolUse` 钩子的‘ask’决策以 TUI 确认框形式展现**
    - **重要性**: 增强工具调用的安全性与可控性。以往 `ask` 决策可能被隐式处理，此功能将其显式化为 TUI 确认对话框，用户在授权敏感操作（如文件修改）前有明确的确认机会。
    - **社区反应**: 该设计被普遍认为是向“安全 Agent”迈出的重要一步，讨论主要围绕如何平衡用户体验和安全确认的频率。
    - **链接**: [Issue #5629](QwenLM/qwen-code Issue #5629)

9.  **#5616 [PR关联] 功能(记忆): 在持久化自动生成的技能前，要求用户确认**
    - **重要性**: 解决了“AI 替开发者决定技能”的问题。自动生成的技能可能不符合用户意图，现在要求用户审核确认后再写入技能库，加强了用户对自身技能集的控制权。
    - **社区反应**: 社区对此设计表示赞赏，认为这是从“全盘接受”到“人工审核”的理性回归，符合开发者对代码和配置的掌控需要。
    - **链接**: [Issue #5616](QwenLM/qwen-code Issue #5616)

10. **#5805 [PR关联] 修复(IDE): 在读取锁文件前，验证 `QWEN_CODE_IDE_SERVER_PORT` 环境变量**
    - **重要性**: 防止因错误的环境变量配置导致 IDE 锁文件路径错误或连接失败。此修复增强了系统的健壮性和错误处理能力。
    - **社区反应**: 这是一个典型的“防御性编程”修复，虽然改动小，但对解决 IDE 集成中的边界问题具有重要意义。
    - **链接**: [Issue #5805](QwenLM/qwen-code Issue #5805)

## 重要 PR 进展

以下是当日推荐的 10 个关键 PR。

1.  **#5650 [OPEN] 功能(web-shell): 增强 Markdown 表格交互**
    - **内容**: 为 AI 返回的表格增加排序、筛选、行选择等类 Excel 交互。
    - **状态**: 开放中，评论数高，社区讨论热烈。
    - **链接**: [PR #5650](QwenLM/qwen-code PR #5650)

2.  **#5765 [OPEN] 功能(服务): 添加守护进程工作区的语音与控制 API**
    - **内容**: 为守护进程增加语音配置、批量转录、权限管理等 REST/ACP/SDK API。
    - **状态**: 开放中，是构建完整远程工作区和语音能力的关键环节。
    - **链接**: [PR #5765](QwenLM/qwen-code PR #5765)

3.  **#5814 [OPEN] 功能(核心): 解耦 `/remember` 与自动提取，停止写入 QWEN.md**
    - **内容**: 重构记忆模块，使用户显式记忆和后台自动提取行为分离，权限更清晰。
    - **状态**: 开放中，被视为重要的架构演进。
    - **链接**: [PR #5814](QwenLM/qwen-code PR #5814)

4.  **#5808 [OPEN] 修复(CLI): 用户中止时取消待处理的 `/loop` 循环唤醒**
    - **内容**: 按下 Esc 退出循环时，同步取消计划中的下一次唤醒，防止循环静默恢复。
    - **状态**: 开放中，快速响应了 Issue #5806 的修复。
    - **链接**: [PR #5808](QwenLM/qwen-code PR #5808)

5.  **#5781 [CLOSED] 暴露 MCP 资源读取工具**
    - **内容**: 为 AI 模型提供一个可调用的 MCP 资源读取工具，使其能在普通对话轮次中读取 MCP 资源。
    - **状态**: 已合并。这是强化 MCP 生态集成的重要一步。
    - **链接**: [PR #5781](QwenLM/qwen-code PR #5781)

6.  **#5661 [OPEN] 功能(TUI): 按工具类型分区展示——折叠读/搜索工具，单独展示变更工具**
    - **内容**: 重构工具调用渲染逻辑，将只读工具折叠为摘要，突出显示有副作用的写入工具。
    - **状态**: 开放中，旨在优化 TUI 信息密度和可读性。
    - **链接**: [PR #5661](QwenLM/qwen-code PR #5661)

7.  **#5812 [CLOSED] 修复(CLI): 正确映射 Claude MCP 服务器的传输类型**
    - **内容**: 修复了在导入 Claude MCP 配置时，对 `http`, `sse` 等传输类型的错误映射。
    - **状态**: 已合并。提高了与其他 MCP 生态的兼容性。
    - **链接**: [PR #5812](QwenLM/qwen-code PR #5812)

8.  **#5616 [OPEN] 功能(记忆): 持久化自动生成技能前要求用户确认**
    - **内容**: 改变后台技能自动写入逻辑，改为“发现技能 -> 提示用户 -> 用户确认后写入”的流程。
    - **状态**: 开放中，积极讨论中。
    - **链接**: [PR #5616](QwenLM/qwen-code PR #5616)

9.  **#5793 [OPEN] 功能(配置): 通过 `providerProtocol` 将 Provider ID 映射到 SDK 协议**
    - **内容**: 允许自定义 Provider 通过配置来映射到已有的 SDK 协议，实现身份与传输行为的分离。
    - **状态**: 开放中，增强了自定义 Provider 的灵活性。
    - **链接**: [PR #5793](QwenLM/qwen-code PR #5793)

10. **#5396 [OPEN] 修复(UI): 减少 UI 闪烁——节流 + 紧凑模式过渡 + 批量 STREAM_TEXT**
    - **内容**: 通过节流渲染、优化 React 过渡动画、合并批量事件等多种手段减少界面闪烁。
    - **状态**: 开放中，评论数高。虽然是旧 PR，但持续更新，致力于解决用户长期反馈的性能卡顿问题。
    - **链接**: [PR #5396](QwenLM/qwen-code PR #5396)

## 功能需求趋势

从今日的议题和 PR 中，可以归纳出社区最关注的几个功能方向：

- **语音交互的深度优化**: 这不是简单的“支持语音输入”，而是精细化到“ASR 关键词库可配置”、“语音转写后智能润色”、“提供守护进程级别的语音控制 API”。这表明社区正将语音作为核心交互方式之一，并追求更高的准确率和流畅性。
- **循环与自动化工作流的健壮性**: `/loop` 模式的 Bug 从报告到修复的快速闭环，以及相关 Issue 被标记为 `roadmap/background-automation`，说明“自动化代理”功能正在从实验阶段走向生产环境，其稳定性和用户可控性是当前开发重点。
- **MCP 生态的深入集成**: 无论是暴露 MCP 资源读取工具，还是修复 Claude MCP 的兼容性，亦或是提出 MCP 服务器的热加载，都表明 Qwen Code 正在全方位拥抱 MCP 生态，力图成为一个能够无缝接驳各种外部工具和资源的 Agent 平台。
- **UI/UX 的重塑与细化**: 大量 PR 集中在优化终端 UI，包括表格交互、工具栏分区、确认对话框、紧凑模式等。这波改进专注于提升信息密度、减少视觉干扰、增强可操作性和安全性，旨在提供更具沉浸感和高效的会话体验。
- **开发者体验与配置灵活性**: 遥测配置、Provider 映射、端口验证等 PR 和 Issue，反映了社区不仅关注功能本身，也日益关注开发体验（DX），如更好的可观测性、更灵活的集成能力以及更健壮的工程实践。

## 开发者关注点

- **痛点一：循环模式的中止不彻底**。
    - **详情**: 用户在中止 `/loop` 后，因底层调度未取消而导致循环静默恢复，这是开发者反馈中最直接的“反直觉”行为。社区快速贡献了修复，表明这是一个常见且急需解决的痛点。

- **痛点二：语音功能的“不接地气”**。
    - **详情**: 硬编码的通用关键词无法用于特定项目，让语音录入变得低效。开发者强烈需要“个性化定制关键词”的能力，这是语音功能走向实用的最后一道坎。

- **痛点三：记忆与自动化的“黑箱操作”**。
    - **详情**: 后台自动提取的记忆和自动生成的技能，开发者普遍担忧其准确性和负面副作用。社区强烈要求在这些自动化操作（如写入 QWEN.md）、写入技能库前增加“人工确认”环节，突出对“可控性”的强烈诉求。

- **痛点四：UI 性能与信息过载**。
    - **详情**: 长时间对话导致的 UI 闪烁、折叠后丢失上下文、工具调用显示混乱等问题，是用户日常使用中的高频痛点。当前的大量 PR 正集中解决这些问题，旨在缓解用户在处理长对话和复杂交互时的认知负担。

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI（CodeWhale）社区动态日报 | 2026-06-24

##  今日速览

- **v0.8.65 发布在即**：核心 PR #3544 已提交，集成了提供者路由、Fleet 负载选择等史诗级功能，标志着多模型支持进入稳定阶段。
- **社区活跃度攀升**：24 小时内新增 5 条 Issue（#3541-#3546），涵盖 Rust 原生客户端、ACP 扩展、自定义上下文大小等关键需求。
- **大量修复与重构并入主线**：包括窗口粘贴兼容性、文件审批规则、MCP OAuth 登录、Fleet 持久化恢复等 20+ 个 PR 被合并或提交。

##  社区热点 Issues（10 条）

以下为过去 24 小时内更新或新建的最值得关注的 Issue：

| Issue | 标题 | 状态 | 热度 | 关注点 |
|-------|------|------|------|--------|
| [#3439](https://github.com/Hmbown/CodeWhale/issues/3439) | [CLOSED] 接入智谱 GLM-5.2 作为 provider route fixture | ✅ 已关闭 | 6 条评论 | 社区强烈希望支持 GLM-5.2，最终通过折叠至 Z.ai 提供者实现，反映了多模型路由设计的落地。 |
| [#2300](https://github.com/Hmbown/CodeWhale/issues/2300) | [OPEN] 多模型兼容性、提供者文档与 Fleet 负载自动选择 | 🔄 开放 | 4 条评论 | 持续跟踪的平台级能力，v0.8.65 即将解决。 |
| [#3087](https://github.com/Hmbown/CodeWhale/issues/3087) | [OPEN] 重写 README 以反映 CodeWhale 历史与提供者路由地图 | 🔄 开放 | 3 条评论 | 架构稳定后的文档更新需求，开发者希望快速了解新功能。 |
| [#3495](https://github.com/Hmbown/CodeWhale/issues/3495) | [OPEN] 采用 Moraine 作为 CodeWhale 的长期记忆后端 | 🔄 开放 | 2 条评论 | 引入 MCP recall 工具，将持久化会话转换为可搜索记忆，v0.8.66 规划。 |
| [#3474](https://github.com/Hmbown/CodeWhale/issues/3474) | [CLOSED] macOS 终端下 /model /sessions 选择器文本对比度极低 | ✅ 已关闭 | 2 条评论 | UI 可访问性 Bug，已被修复（相关 PR #3548）。 |
| [#3541](https://github.com/Hmbown/CodeWhale/issues/3541) | [OPEN] 功能请求：Rust 原生运行时 / 桌面客户端 | 🔄 开放 | 1 条评论 | 用户抱怨 Node 运行时冷启动延迟及内存开销，建议重写为 Rust 原生。 |
| [#3546](https://github.com/Hmbown/CodeWhale/issues/3546) | [OPEN] 扩展 ACP 支持以暴露提供者和模型选择 | 🔄 开放 | 1 条评论 | Paseo 集成中缺乏对提供者/模型的控制，需要 ACP 标准暴露。 |
| [#3545](https://github.com/Hmbown/CodeWhale/issues/3545) | [OPEN] 希望在 providers 配置中自定义上下文大小 | 🔄 开放 | 0 条评论 | 阿里百炼平台模型实际支持 1M 上下文，但当前固定为 128k，需要配置项。 |
| [#2985](https://github.com/Hmbown/CodeWhale/issues/2985) | [CLOSED] 重新发布流程：确保 release 分支提交进入 main | ✅ 已关闭 | 0 条评论 | 自动化流程改进，避免版本标签未合并到 main 导致 github 关闭引用失效。 |
| [#3537](https://github.com/Hmbown/CodeWhale/issues/3537) | [OPEN] 将硬编码本地化文件替换为专用 i18n 库 | 🔄 开放 | 0 条评论 | `localization.rs` 已超 5000 行，维护困难，建议改用标准 i18n 框架。 |

##  重要 PR 进展（10 条）

过去 24 小时内提交/合并的关键 Pull Request：

| PR | 标题 | 状态 | 摘要 |
|----|------|------|------|
| [#3544](https://github.com/Hmbown/CodeWhale/pull/3544) | **release: v0.8.65** | 🔄 开放 | 工作区版本 0.8.64→0.8.65，集成 provider/route 路由和 Fleet 负载选择，CHANGELOG 已更新。 |
| [#3539](https://github.com/Hmbown/CodeWhale/pull/3539) | feat(providers): 将智谱折叠到 Z.ai 提供者 | ✅ 已合并 | 关闭 #3439，无需新增提供者类型，复用 Z.ai 已有端点。 |
| [#3525](https://github.com/Hmbown/CodeWhale/pull/3525) | feat(fleet): 将 worker 状态折叠到 Fleet 界面 | ✅ 已合并 | 新增 `/fleet status` 命令，重构状态展示，子代理命令作为兼容快捷方式保留。 |
| [#3526](https://github.com/Hmbown/CodeWhale/pull/3526) | [codex] 强制执行 main 分支 release 标签 | ✅ 已合并 | 关闭 #2985，防止 release 分支提交绕过 main，确保 `Closes #` 正常处理。 |
| [#3536](https://github.com/Hmbown/CodeWhale/pull/3536) | feat(fleet): 基于账本的持久化 Manager 恢复 | ✅ 已合并 | 实现 `FleetManager.resume()`，支持崩溃/分离后从账本恢复任务，确保确定性。 |
| [#3527](https://github.com/Hmbown/CodeWhale/pull/3527) | feat(tui): 远程 MCP OAuth 登录 | ✅ 已合并 | 为 URL 型 MCP 服务器添加 OAuth 2.0 和 bearer 认证支持。 |
| [#3548](https://github.com/Hmbown/CodeWhale/pull/3548) | fix(tui): 容忍不支持括号粘贴的终端 | 🔄 开放 | 针对旧版 Windows 控制台自动禁用括号粘贴，修复 TUI 恢复路径。 |
| [#3547](https://github.com/Hmbown/CodeWhale/pull/3547) | feat(tui): 从写文件审批中保存精确文件 ask 规则 | 🔄 开放 | 扩展审批弹窗的“保存 ask 规则”到 `write_file` 和 `edit_file`。 |
| [#3530](https://github.com/Hmbown/CodeWhale/pull/3530) | feat(tui): 本地化模式选择器和 Vim 指示器 | ✅ 已合并 | 从历史 PR #2239 中收割遗留的 i18n 改动，确保合并后 credit 正确。 |
| [#3532](https://github.com/Hmbown/CodeWhale/pull/3532) | fix(api): 跨 HTTP API 调用复用共享 McpPool | ✅ 已合并 | 修复 #3461 中每个 API 端点独立创建 McpPool 导致进程重复的问题。 |

##  功能需求趋势

从近期 Issues 和 PR 中可以提炼出社区最关注的三个方向：

1. **多模型路由与提供者生态**：v0.8.65 正式落地提供者路由（Z.ai、智谱、OpenRouter 等），社区期望支持更多国产模型（如 GLM-5.2）并允许精细配置（上下文大小、模型选择）。
2. **Rust 原生运行时与桌面客户端**：Issue #3541 引发讨论，用户认为 Node.js 的冷启动延迟和内存占用在长 Agent 会话中不可接受，建议迁移至 Rust 原生 TUI/CLI，甚至提供桌面客户端。
3. **记忆与持久化增强**：Issue #3495 提出的 Moraine 集成计划（MCP recall 工具）展示了社区对长期 Agent 记忆的渴求，与 Fleet 持久化恢复 PR #3536 一脉相承。

##  开发者关注点

- **性能瓶颈**：多位开发者反馈 Node 运行时的启动延迟和单线程事件循环阻塞问题（#3541），要求原生方案。
- **配置灵活性**：阿里百炼等提供者模型实际上下文大于当前固定值，用户希望 `providers` 配置支持自定义 `context_size`（#3545）。
- **集成标准化**：ACP 作为通用协议，社区希望 CodeWhale 完整暴露提供者/模型选择能力，以便与 Paseo 等工具无缝协作（#3546）。
- **维护性改进**：5000+ 行的硬编码本地化文件（#3537）和过期文档（#3087）表明社区要求更工程化的代码组织与文档自动化。
- **UI 细节打磨**：macOS 终端选择器对比度问题（#3474）虽已修复，但反映出跨平台终端兼容性仍是持续痛点。

---

*数据来源：[Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale)*  
*生成时间：2026-06-24*

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*