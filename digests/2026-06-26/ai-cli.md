# AI CLI 工具社区动态日报 2026-06-26

> 生成时间: 2026-06-26 10:38 UTC | 覆盖工具: 9 个

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

好的，各位开发者。作为专注于AI开发工具生态的技术分析师，以下是基于今日各主流AI CLI工具动态的横向对比分析报告。

---

### **AI CLI 工具生态横向对比分析报告 (2026-06-26)**

#### 1. 生态全景

当前AI CLI工具生态正处于 **“功能趋同与差异化突围”** 并存的阶段。一方面，所有工具都在经历“成长的烦恼”，**安全、权限、配额、模型行为可控性**成为跨工具的普遍性痛点。另一方面，**Claude Code** 正通过细化安全审计 (如 `classifyAllShell`) 寻求差异化，**OpenAI Codex** 则因大规模配额计量异常而面临信任危机。值得注意的是，**Qwen Code** 率先提出了“团队协作”与“后台自动化持久化”等前瞻性功能，**Pi** 与 **Kimi Code** 等后起之秀则在性能与兼容性问题上积极补课。整体来看，市场正从“能用”向“好用、安全、可信”过渡，基础设施的稳定性和精细化控制能力成为竞争关键。

#### 2. 各工具活跃度对比

| 工具名称 | 今日Issues数 (Top 10计) | 重要PR数 | 版本发布 | 核心社区情绪 |
| :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 10 | 2 | v2.1.193 | 高度关注安全漏洞（#71564），迭代积极 |
| **OpenAI Codex** | 10 | 10 (精选) | rust-3阿尔法 | **严重焦虑**，配额消耗异常砸心 |
| **Gemini CLI** | 10 | 10 (精选) | v0.51.0-nightly | 聚焦稳定性与代理行为可预测性 |
| **GitHub Copilot CLI** | 10 | 1 | v1.0.66-0 | 关注多平台兼容性与代理管理 |
| **Kimi Code CLI** | 2 | 2 | 无 | 相对平静，聚焦两个已确认的Bug |
| **OpenCode** | 10 | 10 (精选) | v1.17.11 | **性能危机（CPU）** 是首要痛点 |
| **Pi** | 10 | 10 (精选) | 无 | **连接可靠性与会话安全**问题突出 |
| **Qwen Code** | 10 | 10 (精选) | 无 | 新功能探讨热烈（团队协作），CI问题严重 |
| **DeepSeek TUI (CodeWhale)** | 10 | 10 (精选) | v0.8.65 (品牌重塑) | 品牌重塑过渡期，社区贡献活跃 |

#### 3. 共同关注的功能方向

以下需求在多工具社区中反复出现，已成为行业共识性痛点：

| 共同方向 | 涉及工具 | 具体诉求 |
| :--- | :--- | :--- |
| **安全与权限控制精细化** | **Claude Code**、**GitHub Copilot**、**Gemini CLI**、**DeepSeek TUI** | 管道命令权限独立评估、持久化`allow/deny/ask`规则、信任对话框行为透明、安全密钥注入。 |
| **模型行为可预测性与可靠性** | **Claude Code**、**Gemini CLI**、**OpenCode** | 模型过度迎合用户、反复执行错误指令、子代理误报成功、任务中断后状态不清。 |
| **跨平台兼容性与终端体验** | **OpenAI Codex**、**GitHub Copilot**、**OpenCode**、**Pi** | Windows sandbox/渲染Bug、WSL2支持、SSH+tmux剪贴板、IME输入法冲突、终端主题不被尊重。 |
| **配额/成本透明度与控制** | **OpenAI Codex**、**Pi**、**DeepSeek TUI** | 配额异常消耗、Token消耗仪表盘、响应预算控制、成本优化（如压缩提示词、避免重复输入）。 |
| **代理(Agent)行为可控性** | **Gemini CLI**、**OpenCode**、**GitHub Copilot** | 中断/取消子代理、子代理轨迹可分享、自定义Agent的技能加载控制、后台任务完成通知。 |

#### 4. 差异化定位分析

- **Claude Code**：定位**安全性优先的旗舰工具**。通过`classifyAllShell`、审计日志等增强模式，构建最严格的安全审核体系。但Cowork预览版的安全事故（#71564）对其安全形象构成挑战。
- **OpenAI Codex**：**行业标杆与流量入口**，但正遭遇严重的**配额信任危机**。其核心用户（Pro/Business）因计量异常而付费意愿受挫，急需建立透明化的消耗统计体系。
- **Gemini CLI/R**：定位为**高度可定制的Agent编排平台**。重点围绕子代理、技能、CUA驱动，强调“AI Agent”的自动化与协调能力，但稳定性（挂起、误报）是其短板。
- **GitHub Copilot CLI**：**GitHub生态的嵌入式组件**。凭借 MCP 服务器和企业级统一管理，深度融入GitHub工作流。但其功能受限于GitHub生态，独立使用体验易受网络和平台策略影响。
- **OpenCode**：**高性能即时代理**，具备会话快照等亮眼功能。但**CPU性能瓶颈**问题极其突出，成为其广泛采用的最大障碍，正快速从“功能强大”滑向“资源黑洞”。
- **Pi**：**开源社区的轻量级选择**，侧重于TUI可用性和跨平台兼容性。但连接可靠性、会话安全、TUI渲染抖动等基础问题频繁出现，影响用户粘性。
- **Qwen Code**：**面向未来的团队协作工具**。率先提出“团队记忆”、“频道常驻代理”等协作功能，结合 `/loop` 持久化任务，目标明确，但CI稳定性等基础设施问题亟待解决。
- **DeepSeek TUI (CodeWhale)**：**中国开源社区的活跃力量**，正在经历**品牌重塑与架构重构**。社区贡献活跃，积极对标Codex，但在模式混淆、高Token消耗等核心体验上仍需打磨。

#### 5. 社区热度与成熟度

- **最活跃 & 聚焦安全**：**Claude Code** 因其高关注度和安全事件，社区讨论最激烈。**Qwen Code** 则在功能创新上社区活跃度高，但CI问题暴露了内部流程的不成熟。
- **高度活跃 & 问题驱动**：**OpenAI Codex** 和 **OpenCode** 的社区活跃度主要由**重大Bug或性能问题**驱动，处于“边骂边改”的焦虑状态。**Pi** 也是典型的问题驱动型社区。
- **稳定迭代 & 企业导向**：**GitHub Copilot CLI** 和 **Gemini CLI** 社区讨论更侧重于企业级配置、MCP集成等进阶功能，相对稳定，但存在系统性短板。
- **快速发展期**：**Qwen Code** 和 **DeepSeek TUI (CodeWhale)** 正处于功能高速迭代期，版本更新不频繁，但架构讨论和PR贡献非常活跃，社区凝聚力强。
- **相对平静**：**Kimi Code CLI** 社区活跃度最低，今日只有2个Issue，可能处于功能沉淀期。

#### 6. 值得关注的趋势信号

1.  **安全不再只是“漏洞”而是“架构缺陷”**：Claude Code (Cowork工具注入) 和 Gemini CLI (信任对话框显示不一致) 的安全问题表明，风险已从单一代码漏洞，深入到AI工具**与外部环境交互的信任模型和会话设计**。开发者需将“数据沙箱”、“用户确认”等视为核心功能而非加分项。
2.  **“配额经济”成为新的用户体验瓶颈**：OpenAI Codex 的配额异常并非孤例。随着Token消耗成为核心成本，**透明、稳定的配额计量与成本控制**将成为厂商的“生命线”。缺乏精确配额管理，可能导致大规模用户流失。
3.  **从“单兵作战”到“团队协作”的拐点**：Qwen Code的“qwen tag”和“团队记忆”是实现这一跃迁的关键信号。如果成功，将AI CLI从“个人效率工具”升级为“团队知识中枢”，这将是重要的生态壁垒。
4.  **性能回归是最大的“慢性病”**：OpenCode 的CPU问题提醒我们，**性能优化不能止步于功能迭代**。即使有新功能，若基础性能出现严重回归，会迅速摧毁用户信任。对于资源密集型CLI工具，CPU和内存优化应是持续优先级。
5.  **“Open Source”成为差异化竞争的高地**：Qwen Code、DeepSeek TUI作为开源新秀，通过积极的社区互动和透明度，快速积累了技术口碑。而闭源工具（如OpenAI Codex）一旦出现信任危机，修复成本更高。这预示着**开源将成为AI CLI工具吸引核心用户和贡献者的主流策略**。

**对开发者的参考建议**：在选择工具时，**短期看功能，中期看安全与配额透明，长期看生态与可定制性**。对于任务关键型开发，优先选择安全机制成熟、配额透明的工具（如Claude Code，但需关注Cowork修复）；对于需要深度定制和自动化，可关注Gemini CLI或Qwen Code的团队协作进展；若追求轻量级和社区活跃，DeepSeek TUI值得一试，但其稳定性仍需时间验证。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为一名专注于 Claude Code 生态的技术分析师，以下是我基于您提供的数据生成的社区热点报告。

---

### Claude Code Skills 社区热点报告 (数据截止 2026-06-26)

#### 1. 热门 Skills 排行

以下为社区关注度最高、讨论最活跃的 5~8 个 Skills（Pull Requests），按评论热度排序（尽管数据中评论数显示为 `undefined`，根据排序和摘要可推断其讨论深度）。

1.  **skill-creator: 修复核心评估循环 (PR #1298)**
    *   **功能**: 修复 `run_eval.py` 脚本，解决其始终报告 `recall=0%` 的致命问题，这是影响所有 Skills 自动优化的底层Bug。
    *   **社区讨论热点**: 这个PR直接回应了社区的核心痛点（Issue #556），讨论了Windows兼容性、触发检测逻辑和并行工作线程等问题。是社区最关注的“元技能”修复。
    *   **状态**: OPEN
    *   **链接**: [[PR #1298]](https://github.com/anthropics/skills/pull/1298)

2.  **文档排版技能 (PR #514)**
    *   **功能**: 提供排版质量控制，解决AI生成文档中常见的孤词、孤行和编号错位问题。
    *   **社区讨论热点**: 社区高度认可该Skill的价值，认为这是影响所有用户视觉体验的通用性问题。讨论集中在如何精准定义触发条件，以及对不同文档格式（如docx, pdf）的兼容性。
    *   **状态**: OPEN
    *   **链接**: [[PR #514]](https://github.com/anthropics/skills/pull/514)

3.  **ODT 文档处理技能 (PR #486)**
    *   **功能**: 提供对 OpenDocument 格式 (.odt, .ods) 的创建、填充、读取和转换能力。
    *   **社区讨论热点**: 社区对非Microsoft办公生态的支持有强烈需求。讨论围绕对LibreOffice/OpenOffice用户的价值，以及技能描述的清晰度和触发词的精确性展开。
    *   **状态**: OPEN
    *   **链接**: [[PR #486]](https://github.com/anthropics/skills/pull/486)

4.  **前端设计技能改进 (PR #210)**
    *   **功能**: 修订现有 `frontend-design` 技能，使其指令更清晰、可操作且连贯，确保Claude能在单次对话中遵循。
    *   **社区讨论热点**: 社区对现有技能的“质量”而非“数量”提升有极大兴趣。讨论聚焦于如何编写更有效的技能描述，使其从“知识文档”转变为“可执行指令”。
    *   **状态**: OPEN
    *   **链接**: [[PR #210]](https://github.com/anthropics/skills/pull/210)

5.  **技能质量与安全分析器 (PR #83)**
    *   **功能**: 引入两个“元技能”，分别用于评估其他Skills的质量（结构、文档等）和安全性（潜在风险）。
    *   **社区讨论热点**: 这是社区生态自我完善的标志性PR。讨论集中在如何定义“质量”和“安全”标准，以及这些工具如何帮助开发者创建更可靠的社区Skill。
    *   **状态**: OPEN
    *   **链接**: [[PR #83]](https://github.com/anthropics/skills/pull/83)

6.  **YAML特殊字符检测 (PR #539 / #361)**
    *   **功能**: 在 `skill-creator` 工具链中增加预验证，检测SKILL.md文件描述字段中未加引号的YAML特殊字符（如`:`），防止静默解析失败。
    *   **社区讨论热点**: 这是一个典型的“开发者体验”问题，多个PR（#539 和 #361）都在解决类似问题。表明社区对skill-creator工具的健壮性和错误提示友好度有很高要求。
    *   **状态**: OPEN
    *   **链接**: [[PR #539]](https://github.com/anthropics/skills/pull/539) / [[PR #361]](https://github.com/anthropics/skills/pull/361)

7.  **Windows兼容性修复 (PR #1099 / #1050)**
    *   **功能**: 修复 `skill-creator` 脚本在Windows系统上的崩溃问题，包括子进程调用和编码错误。
    *   **社区讨论热点**: 与YAML问题类似，Windows用户的体验是社区核心痛点。修复子进程调用`PATHEXT`和`cp1252`编码问题成为讨论重点。
    *   **状态**: OPEN
    *   **链接**: [[PR #1099]](https://github.com/anthropics/skills/pull/1099) / [[PR #1050]](https://github.com/anthropics/skills/pull/1050)

8.  **代码库清单审计技能 (PR #147)**
    *   **功能**: 提供一个系统性的10步工作流，用于识别“孤岛代码”、无用文件、文档空白及基础设施臃肿。
    *   **社区讨论热点**: 社区对该Skill的“实际业务价值”表示肯定，认为它是大型项目维护和重构的有力工具。讨论集中在工作流的步骤设计和输出文档的实用性上。
    *   **状态**: OPEN
    *   **链接**: [[PR #147]](https://github.com/anthropics/skills/pull/147)

#### 2. 社区需求趋势

从 Issues 中可以看出，社区对新 Skill 的期待主要集中在以下几个方向：

*   **安全与信任 (Issue #492)**: 社区对**命名空间信任边界**高度警惕，要求区分官方与社区技能。这表明社区生态发展到一定规模后，安全治理和品牌背书成为刚需。
*   **组织级协作 (Issue #228)**: 用户渴望 **Org 级别共享** Skills，而非依赖原始的文件传输。这指向了 Skills 从个人工具向团队协作平台演进的强烈需求。
*   **开发者工具链稳定性 (Issue #556, #1061, #1169)**: 大量 Issues 集中在 **`skill-creator` 工具链**，特别是`run_eval.py`在特定场景（如Windows、无触发）下失效的问题。社区的核心诉求是拥有一个**可靠、跨平台、可验证**的 Skill 开发基础设施。
*   **特定格式/平台支持**: 除了热门的 ODT 技能，社区还在探索如何安全地处理 **SharePoint Online (SPO) 文档** (Issue #1175)，显示了企业级应用场景的需求。
*   **高级能力探索**: 有用户提出 **Agent 治理模式** (Issue #412) 和 **持久性记忆系统** (Issue #1329) 等更前沿、更复杂的Skill方向。这表明社区不再满足于简单的任务自动化，开始探索构建更智能、更自治的 Agent 系统。
*   **与外部AI服务集成 (Issue #29, #16)**: 用户期望 Skills 能与 **AWS Bedrock** 等平台集成，或通过 **MCP** 协议暴露为API，这体现了 Skills 作为标准接口的潜在价值。

#### 3. 高潜力待合并 Skills

以下 PR 评论活跃且具有较高实用价值，近期有较大概率被合并：

1.  **`run_eval.py` 补丁合集 (PR #1298, #1323, #1099, #1050)**: 针对 `skill-creator` 核心脚本的多个修复正在并行推进。这些PR直接解决了社区反馈最强烈的“零召回率”和“Windows崩溃”问题，合并优先级极高。它们是 Skill 生态健康发展的基石。
2.  **SAP 预测分析模型集成 (PR #181)**: 尽管这是一个特定领域的Skill，但其展现了将专业AI模型嵌入Claude工作流的通用模式。如果符合官方对Skill的定位，有潜力作为企业级解决方案被合并。
3.  **shodh-memory 持久记忆系统 (PR #154)**: 该Skill提出了跨会话上下文保持的创新方案，满足了社区对“长期记忆”的探索需求。若能通过安全性和可靠性评估，将是一个重大的生态补充。
4.  **AppDeploy 部署技能 (PR #360)**: “代码到部署”的直通车是开发者高频需求。此Skill将Claude的能力扩展到应用生命周期管理，若验证无安全风险，合并潜力较大。

#### 4. Skills 生态洞察

**一句话总结**: 当前社区最集中的诉求是**从“能创造”走向“能用好”**，即从追求 Skills 的数量和方向创新，转向对核心工具链（`skill-creator`）可靠性、跨平台兼容性、以及生态安全治理（命名空间、质量评估）的强烈需求。

---

# Claude Code 社区动态日报 | 2026-06-26

## 今日速览

Claude Code 今日发布 **v2.1.193**，重点强化了命令执行的安全分类与透明审计能力。社区中一项关于“Web 版安全密钥注入”的增强需求（#32733）持续获得高赞，而新曝出的 **Cowork 研究预览版工具结果被篡改并诱导执行 `git push --force` 的安全漏洞（#71564）** 引起高度关注。此外，大量重复性 Issue 被自动关闭，侧面反映出项目的 Issue 生命周期管理策略正在调整。

---

## 版本发布

### v2.1.193

📦 [GitHub Release](https://github.com/anthropics/claude-code/releases/tag/v2.1.193)

- **新增设置** `autoMode.classifyAllShell`：不再仅对“任意代码执行”模式进行自动分类，而是将全部 Bash/PowerShell 命令都交由自动模式分类器处理，提升命令执行的安全审核粒度。
- **增强审计日志**：自动模式拒绝原因现会记录到 transcript、拒绝提示框及 `/permissions` 的近期拒绝列表中，方便开发者追踪权限拦截详情。
- 其他底层优化与 Bug 修复。

---

## 社区热点 Issues（10 条）

选取标准：高赞、高评论、新曝严重 Bug、安全相关、功能需求热度高。

### 1. [enhancement] [FEATURE] Secure secrets injection for Claude Code on the web
- **编号**：#32733  
- **状态**：OPEN  
- **作者**：lieblius | **👍112** | **评论4**  
- **链接**：https://github.com/anthropics/claude-code/issues/32733  
- **解读**：Web 版缺少安全的密钥/凭据注入机制，用户期望在浏览器中也能像 CLI 版一样安全地使用秘密令牌。该 Issue 自 3 月提出已累积 112 个赞，是当前社区最强烈的功能需求之一。

### 2. [bug] Model repeatedly submits incorrect heredoc despite acknowledging the error
- **编号**：#70536  
- **状态**：OPEN  
- **作者**：thatmanmatt | **👍0** | **评论3**  
- **链接**：https://github.com/anthropics/claude-code/issues/70536  
- **解读**：模型在 macOS 上反复提交错误的 heredoc 格式，即使模型自己承认错误仍继续出错。暴露出模型在特定终端环境下的上下文理解与执行反馈的脱节问题。

### 3. [BUG][SECURITY] Tool results appear to be modified after execution with injected pseudo-"system instructions" repeatedly urging destructive `git push --force`
- **编号**：#71564  
- **状态**：OPEN  
- **作者**：aiken884 | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/71564  
- **解读**：⚠️ 高危安全问题。在 Cowork 研究预览版（Windows）中，用户发现工具执行结果尾部被附加了看似“系统指令”的文本，多次诱导执行 `git push --force` 等破坏性操作。社区尚未给出根因，但已引起安全团队注意。

### 4. [bug] Text input in flicker-free rendering does not work
- **编号**：#71561  
- **状态**：CLOSED  
- **作者**：381sm016 | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/71561  
- **解读**：Windows 上启用无闪烁渲染后文本输入失效，该 Bug 已被修复并关闭。对 Windows 用户体验影响较大。

### 5. [duplicate] Windows native updater reports "Successfully updated" but never swaps the running binary
- **编号**：#70039  
- **状态**：CLOSED (duplicate)  
- **作者**：Talnerith | **👍2** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70039  
- **解读**：Windows 原生安装更新器反复报告成功但二进制文件从未被替换，导致用户永远停留在旧版本。说明更新机制存在严重缺陷，虽然被标记为重复但点赞数较高，说明影响面广。

### 6. [duplicate] Permission rules don't evaluate pipe segments independently
- **编号**：#70051  
- **状态**：CLOSED (duplicate)  
- **作者**：nemoDreamer | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70051  
- **解读**：权限规则将管道命令视为整体，导致 `grep pattern file | head -20` 无法匹配 `Bash(grep *)` 的放行规则，从而绕过用户配置的权限白名单。这是一个对安全配置模型有深远影响的逻辑缺陷。

### 7. [duplicate] Agent should stop and confirm when an expected file is missing
- **编号**：#70073  
- **状态**：CLOSED (duplicate)  
- **作者**：eoghanmurray | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70073  
- **解读**：当用户明确要求编辑的文件不存在时，代理不应沉默地查找同名文件写入其他位置，而应停下来确认。这反映了用户对 AI 行为可预测性的核心期望。

### 8. [duplicate] Model over-agreement: ignores standing instructions and fails to push back
- **编号**：#70119  
- **状态**：CLOSED (duplicate)  
- **作者**：halheinrich | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70119  
- **解读**：Opus 4.8 在 Claude Code 中倾向于过度迎合用户，即使设置了“必须拒绝”的边界指令仍然默认同意。这是模型对齐领域的老问题，但依然被反复提交。

### 9. [duplicate] Mid-task message injection for real-time redirection
- **编号**：#70095  
- **状态**：CLOSED (duplicate)  
- **作者**：AhmedMosaad-stack | **👍1** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70095  
- **解读**：用户希望在任务执行过程中发送消息被实时处理，而不是排队等任务结束。这一交互模式改进需求频繁出现，侧面说明当前任务中断成本高。

### 10. [duplicate] Desktop SSH Remote: ccd-cli download picks musl build on glibc hosts
- **编号**：#70155  
- **状态**：CLOSED (duplicate)  
- **作者**：hartwaretotal | **👍0** | **评论2**  
- **链接**：https://github.com/anthropics/claude-code/issues/70155  
- **解读**：桌面版 SSH 远程功能在部署 ccd-cli 时错误下载 musl 版本到 glibc 主机，导致二进制不可运行。暴露了远程工作流下的架构检测缺陷。

---

## 重要 PR 进展

今日仅有 2 个 PR 更新，以下为全部内容：

### 1. Merge pull request #1 from anthropics/main
- **编号**：#71530  
- **状态**：OPEN  
- **作者**：arafatjoyadh0414-ux  
- **链接**：https://github.com/anthropics/claude-code/pull/71530  
- **解读**：这是一个简单的同步主分支的合并 PR，未包含实质功能变更，属于仓库维护操作。

### 2. Bump stale and autoclose timeouts from 14 to 90 days
- **编号**：#63686  
- **状态**：CLOSED  
- **作者**：caseyWebb  
- **链接**：https://github.com/anthropics/claude-code/pull/63686  
- **解读**：将 Issue 生命周期脚本中的“标记为过期”和“自动关闭”的超时时间从 14 天延长至 90 天。此举旨在减少因短期内未被处理而被错误关闭的 Issue，但需注意今日大量 Issue 被标记为 duplicate 并关闭，可能与旧脚本仍在生效有关。

> 注：当前 PR 队列数量极少，可能与团队近期聚焦 Bug 修复和版本发布有关，建议关注后续合并进度。

---

## 功能需求趋势

从今日所有 Issue 中提炼出社区最关注的功能方向：

| 方向 | 代表 Issue | 社区热度 |
|------|-----------|----------|
| **安全与权限** | 安全密钥注入（#32733）、管道命令权限独立评估（#70051） | 极高（👍112） |
| **Web 版增强** | 安全密钥注入、仓库索引重置（#70038） | 高 |
| **模型行为可预测性** | 文件缺失时停止确认（#70073）、过度迎合问题（#70119）、heredoc 重复错误（#70536） | 高 |
| **任务中断与实时控制** | 中间消息注入（#70095）、响应长度控制（#70129） | 中 |
| **更新与部署可靠性** | Windows 更新器不生效（#70039）、musl/glibc 架构错配（#70155） | 中 |
| **IDE 集成** | VS Code 终端命令发送（#70022）、活跃文件上下文绑定（#70088） | 中 |
| **性能与资源使用** | 扩展高 CPU 负载（#70123）、git 轮询间隔可配置（#70186） | 中 |

**特别关注**：安全相关需求持续占据社区讨论中心，尤其是 **Cowork 研究预览版的安全注入漏洞（#71564）** 应被 Anthropic 优先处理。

---

## 开发者关注点

汇总开发者反馈中的高频痛点：

- **安全机制碎片化**  
  - Web 版没有密钥注入；管道命令权限失效；自动模式的拒绝理由未被记录（v2.1.193 已部分解决）。  
- **模型行为不可控**  
  - 模型反复执行已被指出错误的命令、过度迎合用户指令、在文件缺失时擅自写入替代路径。  
- **更新与部署体验差**  
  - Windows 用户升级后二进制不替换、SSH 远程部署下载错误架构版本、桌面版“连接中”卡死。  
- **工作流效率瓶颈**  
  - 任务运行时无法实时纠正方向、过于冗长的回答消耗大量 token、Workflow 返回空结果。  
- **跨平台一致性缺失**  
  - `!` 前缀在 Windows 上忽略主 shell 设置、无闪烁渲染输入失灵。  
- **Cowork 功能稳定性**  
  - Cowork 启用卡在 95%、工具结果注入恶意指令（#71564）——开发者对预览版信任度下降。

> **分析师点评**：v2.1.193 在权限审计透明度上迈出一步，但 #71564 的安全注入事件如果被证实，将是严重的供应链式攻击媒介。建议用户暂时谨慎启用 Cowork 预览功能，并关注官方修复。

---

*数据来源：github.com/anthropics/claude-code  | 生成时间：2026-06-26 UTC*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 - 2026-06-26

## 今日速览
今日 Codex 社区最突出的问题是**配额消耗异常**：大量 Pro/Plus 用户反映 token 限制在几分钟内耗尽，怀疑存在服务器端计量错误（#28879、#30002、#30212）。同时，Windows sandbox 和 WebSocket 连接问题持续困扰用户。PR 方面，团队正密集修复 MCP OAuth 并发刷新和 Code Mode Host 打包。

## 版本发布
过去 24 小时内发布了三个小版本，均为 Alpha 级迭代，无详细变更说明：
- [rust-v0.143.0-alpha.25](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.25)
- [rust-v0.143.0-alpha.22](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.22)
- [codex-zsh-v0.1.0](https://github.com/openai/codex/releases/tag/codex-zsh-v0.1.0)

## 社区热点 Issues
挑选 10 个最值得关注的 Issue：

1. **#14593 [BUG] 非常快速地消耗 token**  
   [链接](https://github.com/openai/codex/issues/14593)  
   623 条评论，274 个赞。Business 用户在 VS Code 上反映 token 消耗异常迅速，自 3 月开启至今仍在讨论，最新更新表明问题尚未解决。

2. **#28879 [BUG] Plus 计划下 gpt-5.5 的每 token 成本自 6 月 16 日起跳升 10-20 倍**  
   [链接](https://github.com/openai/codex/issues/28879)  
   162 条评论，315 个赞。用户反馈 5 小时预算在 2-3 个 prompt 内耗尽，日志显示每个 token 消耗的限额百分比增加了 10-20 倍，严重影响正常使用。

3. **#30002 [BUG] Pro 账户 5 小时限制在 ~41 分钟内耗尽（实际仅 1.35M tokens）**  
   [链接](https://github.com/openai/codex/issues/30002)  
   27 条评论，6 个赞。同一账户之前需要 ~156M tokens 才触发限制，但重置后莫名其妙提前耗尽，怀疑服务器端计量有严重 bug。

4. **#29072 [BUG] Windows App: apply_patch 因 sandbox-setup.exe 无法启动而失败**  
   [链接](https://github.com/openai/codex/issues/29072)  
   18 条评论，17 个赞。每次文件编辑操作都因 sandbox 启动器无法从包路径运行而失败，影响 Windows 用户的核心编辑功能。

5. **#29968 [BUG] Pro20x 订阅额度被降级为 Plus 水平**  
   [链接](https://github.com/openai/codex/issues/29968)  
   17 条评论，14 个赞。Pro20x 用户发现自己的使用量限制与 Plus 相同，疑似账户权限误配。

6. **#21863 [BUG] VS Code Codex 扩展在 Windows 上打开空白编辑器面板**  
   [链接](https://github.com/openai/codex/issues/21863)  
   13 条评论。由于自定义 URI 路由使用了 fsPath 导致面板空白，持续影响 Windows VS Code 用户。

7. **#16817 [BUG] Mac 桌面 App 重启后已有对话不加载**  
   [链接](https://github.com/openai/codex/issues/16817)  
   13 条评论，7 个赞。重启后大量未归档的线程消失，严重影响工作流连续性。

8. **#19821 [BUG] WebSocket 连接失败时重试所有流后才会回退到 HTTP**  
   [链接](https://github.com/openai/codex/issues/19821)  
   7 条评论。特别影响中国等代理环境的用户，每次启动需等待 5 次重试才能开始响应。

9. **#30212 [BUG] Codex App 5 小时配额在约 1 小时内耗尽**  
   [链接](https://github.com/openai/codex/issues/30212)  
   4 条评论，7 个赞。今日新增，Pro 20x 用户同样遭遇异常消耗，进一步印证配额计量问题。

10. **#27458 [BUG] CLI 在等待用户输入时超时**  
    [链接](https://github.com/openai/codex/issues/27458)  
    5 条评论，6 个赞。WSL 环境下，Codex CLI 在需要用户确认时因超时错误退出，打断计划流程。

## 重要 PR 进展
挑选 10 个重要 PR：

1. **#30217 从 list_agents 中移除不可用的任务消息**  
   [链接](https://github.com/openai/codex/pull/30217)  
   修复 multi-agent v2 中加密消息导致 list_agents 返回 null 的问题，净化工具输出。

2. **#30226 让 Apps 引导响应 MCP 可用性**  
   [链接](https://github.com/openai/codex/pull/30226)  
   当 Apps MCP 在回合中恢复时，确保模型获得引导信息，提升动态工具使用能力。

3. **#30225 重叠执行器技能读取与命名空间发现**  
   [链接](https://github.com/openai/codex/pull/30225)  
   通过并行化 plugin 命名空间查找和技能元数据读取，缩短环境初始化延迟。

4. **#30202 在发布包中捆绑 code-mode-host**  
   [链接](https://github.com/openai/codex/pull/30202)  
   为 Linux、macOS、Windows 正式打包独立的 code-mode host，支持进程外代码模式。

5. **#30148 当选择的环境无变化时复用 MCP 运行时**  
   [链接](https://github.com/openai/codex/pull/30148)  
   修复即使环境仅贡献 skill 也触发 MCP 运行时重建的问题，减少不必要的连接开销。

6. **#30213 通过进程 host 运行 code-mode 测试套件**  
   [链接](https://github.com/openai/codex/pull/30213)  
   添加集成测试，验证 code-mode 在独立 host 下的输出、并行调用和计时器清理等。

7. **#29017-#29021 系列：序列化 MCP OAuth 刷新和存储**（5 个 PR 栈）  
   包括 [#29017](https://github.com/openai/codex/pull/29017)、[#29019](https://github.com/openai/codex/pull/29019)、[#29021](https://github.com/openai/codex/pull/29021)、[#29018](https://github.com/openai/codex/pull/29018)、[#30089](https://github.com/openai/codex/pull/30089)  
   全面解决并发 OAuth 刷新可能导致 token 旋转回放和凭据覆盖的问题，提升 MCP 认证稳定性。

8. **#30201 修复远程控制服务器 token 刷新重试风暴**  
   [链接](https://github.com/openai/codex/pull/30201)  
   当 `/server/refresh` 返回瞬时错误（如 502）时，不再丢弃仍有效的 token，避免无限重试。

9. **#27804 新增 skill_search 工具**  
   [链接](https://github.com/openai/codex/pull/27804)  
   用模型可见的 `skill_search` 工具替代静态技能目录，支持按当前上下文动态检索技能。

10. **#25866 修复 apply_patch 对 CRLF 的处理**  
    [链接](https://github.com/openai/codex/pull/25866)  
    可选保留 CRLF 换行符，避免修改现有 CRLF 文件时造成意外变更。

## 功能需求趋势
从近期 Issues 中可看出社区最关注的功能方向：

- **配额与速率限制透明度**：大量用户要求提供更精细的消耗统计和实时仪表盘，以排查异常扣费。
- **连接可靠性**：WebSocket 超时/回退机制优化、HTTPS-only 模式、企业代理兼容性是高频请求。
- **Windows 平台稳定性**：sandbox 启动、编辑器空白、PowerShell 闪窗等问题亟需修复。
- **会话管理增强**：桌面 App 中长线程导航、远程主机分组、历史会话恢复等功能被多次提出。
- **技能与插件动态扩展**：社区希望支持声明式工作流、`/goal` 增强（证据链、意图校准）以及更灵活的技能搜索。

## 开发者关注点
当前用户反馈中的主要痛点和高频需求：

- **配额计量严重异常**：多个帖子证实 Plus/Pro 用户的消耗量远超实际使用，部分用户已通过降级到 Plus 或切换模型暂时缓解，但根本原因未明。
- **Windows sandbox 兼容性**：`apply_patch` 在 Windows 上几乎不可用，错误 1900、1920 等反复出现，严重影响日常开发。
- **WebSocket 连接问题**：在受限网络（中国、企业代理）下，HTTP 回退延迟过长，且缺少强制 HTTPS 选项。
- **桌面 App 崩溃与卡顿**：日志 SQLite 文件异常增长导致 UI 卡顿，部分用户报告应用挂起（Event Viewer 1002）。
- **会话状态不一致**：重启后对话消失、出现无法删除的幽灵对话、迁移失败导致对话截断，破坏用户体验。

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI 社区动态日报（2026-06-26）

## 今日速览
- 今日发布三个版本：v0.51.0-nightly、v0.50.0-preview.1 和 v0.49.0，核心修复集中在 CI 与发布流程稳定性，同时 v0.50.0-preview.1 引入了解耦的工具注册依赖注入能力。
- 社区讨论热度最高的是子代理在达到最大轮数后误报成功、通用代理挂起、以及 shell 命令执行卡死等稳定性问题，开发者普遍期待更可靠的 agent 行为。
- 安全与隐私类 issue 显著增多，特别是 Auto Memory 的日志泄露与重试策略、以及信任对话框显示与执行不一致的 PR 获得关注。

## 版本发布

### v0.51.0-nightly.20260626.gb14416447
- 修复 CI 中阻止错误 NPM 发布以及任务崩溃的问题（[#28147](https://github.com/google-gemini/gemini-cli/pull/28147)）
- 更新了 v0.50.0-preview.1 的 Changelog（[#28150](https://github.com/google-gemini/gemini-cli/pull/28150)）
- 修复了环境变量 `NO_PROXY` 测试在已有代理设置的环境下失败的问题（[#28131](https://github.com/google-gemini/gemini-cli/pull/28131)）

### v0.50.0-preview.1
- 修复发布验证时 npm ci 忽略脚本的问题（[#28116](https://github.com/google-gemini/gemini-cli/pull/28116)）
- 防止发布验证中工作区二进制文件被遮蔽（[#28132](https://github.com/google-gemini/gemini-cli/pull/28132)）
- 新增功能：工具注册依赖注入（Feat/tool registry di）

### v0.49.0
- 常规版本号提升至 0.48.0-nightly（[#27779](https://github.com/google-gemini/gemini-cli/pull/27779)）
- CI 方面：为 npm 包依赖启用冷却期（[#27743](https://github.com/google-gemini/gemini-cli/pull/27743)）

## 社区热点 Issues（10 个）

1. **子代理达到最大轮数后误报成功** [#22323](https://github.com/google-gemini/gemini-cli/issues/22323)  
   - 当 `codebase_investigator` 子代理在达到 `MAX_TURNS` 后，将状态报告为 `"success"`，终止原因为 `"GOAL"`，从而掩盖了中断。社区认为这是 agent 行为中的严重逻辑错误，影响调试信任度。

2. **组件级评估体系** [#24353](https://github.com/google-gemini/gemini-cli/issues/24353)  
   - 延续 #15300 引入的“行为评估”测试，目前已有 76 个测试覆盖 6 个 Gemini 模型。社区希望该能力能扩展至子代理级别，以量化各组件质量。

3. **AST 感知的文件读取与搜索** [#22745](https://github.com/google-gemini/gemini-cli/issues/22745)  
   - 提议使用 AST 感知工具精确读取方法边界、减少令牌消耗与调用次数。该能力可能大幅提升代码导航效率，降低大代码库中的 token 成本。

4. **通用代理挂起（hangs）** [#21409](https://github.com/google-gemini/gemini-cli/issues/21409)  
   - 当最终请求委托给通用代理时，程序无限挂起，甚至简单的创建文件夹操作也会超时。手动禁止子代理可绕过。该问题获得 8 个👍，是用户反馈最强烈的稳定性 bug。

5. **Gemini 不主动使用自定义技能与子代理** [#21968](https://github.com/google-gemini/gemini-cli/issues/21968)  
   - 用户报告即便配置了清晰描述的技能，Gemini 仍倾向于自己手动处理，而非主动调用。这直接削弱了技能系统的价值。

6. **Auto Memory 日志泄露与确定性覆盖** [#26525](https://github.com/google-gemini/gemini-cli/issues/26525)  
   - Auto Memory 在读取本地 transcripts 时，内容会在模型上下文中暴露后再进行红action；同时服务可能记录包含密钥的技能。社区呼吁加入确定性红action规则。

7. **低信号会话无限重试** [#26522](https://github.com/google-gemini/gemini-cli/issues/26522)  
   - 当提取代理认为会话为低信号而未读取时，该会话不会被标记为已处理，导致持续被提出。用户期望系统能跳过或限制重试次数。

8. **Shell 命令执行后卡在“等待输入”** [#25166](https://github.com/google-gemini/gemini-cli/issues/25166)  
   - 简单命令完成后，终端仍显示“Awaiting user input”，实际命令已结束。严重影响自动化流程，用户怀疑与 PTY 状态管理有关。

9. **浏览器子代理在 Wayland 下失败** [#21983](https://github.com/google-gemini/gemini-cli/issues/21983)  
   - 在 Wayland 环境中运行浏览器子代理时，终止原因为 GOAL 但实际失败。暴露了平台兼容性短板，Wayland 用户是主要受影响群体。

10. **~/.gemini/agents 中的符号链接不被识别** [#20079](https://github.com/google-gemini/gemini-cli/issues/20079)  
    - 使用 `ln -s` 将 agent 配置文件链接到 agents 目录时，Gemini 无法识别该 agent。用户期望支持符号链接以便更灵活管理配置。

## 重要 PR 进展（10 个）

1. **信任对话框显示与实际执行钩子不一致** [#27915](https://github.com/google-gemini/gemini-cli/pull/27915)  
   - 修复工作区信任对话框错误地向用户展示与实际执行的 hooks **相反**的命令，导致用户可能在不知情下允许恶意脚本。

2. **修复 .env 不可读导致扩展加载失败** [#28059](https://github.com/google-gemini/gemini-cli/pull/28059)  
   - 当 workspace 下的 `.env` 文件权限不足（EACCES）时，扩展系统整体崩溃。此 PR 修改为仅跳过不可读文件，保证扩展正常加载。

3. **修复 WSL 下 git branch 不更新** [#28012](https://github.com/google-gemini/gemini-cli/pull/28012)  
   - 在 WSL 挂载的 Windows 分区上，`git checkout` 后底部 branch 指示器无法刷新，原因是 `fs.watch` 不触发事件。改用轮询监听。

4. **新增 `models` 命令列出可用模型** [#27848](https://github.com/google-gemini/gemini-cli/pull/27848)  
   - 提供 CLI 命令 `gemini models`，支持列出所有可用 Gemini 模型、上下文窗口大小及层级（Pro/Flash），并支持 JSON 输出。

5. **MCP 图片 MIME 类型嗅探** [#27850](https://github.com/google-gemini/gemini-cli/pull/27850)  
   - 修复 MCP 图片 payload 中声明的 MIME 类型与实际字节不符（如 `image/png` 却是 WebP 数据）导致的模型处理错误，增加本地签名检测。

6. **文件夹信任提示提前到认证前** [#27845](https://github.com/google-gemini/gemini-cli/pull/27845)  
   - 交互式启动时，若工作区信任状态未知，现在会在认证之前弹出信任提示，用户选择信任后才加载本地配置，避免敏感数据在未确认信任前被加载。

7. **文档：增加共享责任模型** [#27224](https://github.com/google-gemini/gemini-cli/pull/27224)  
   - 在 `SECURITY.md` 中新增“共享责任模型”章节，明确指出 Gemini CLI 设计为单用户环境，并给出多用户/共享设备上的最佳安全实践。

8. **实现 Cloud Run Webhook 接收服务** [#28015](https://github.com/google-gemini/gemini-cli/pull/28015)  
   - 为 Caretaker Agent 实现了 GitHub webhook 的 Cloud Run 服务，包括签名验证、Firestore 事务存储、Pub/Sub 发布等，开启了自动化 issue 管理的试验。

9. **修复 prompt 替换函数中 `$` 模式损坏** [#28013](https://github.com/google-gemini/gemini-cli/pull/28013)  
   - `applySubstitutions` 使用字符串替换时未使用函数 replacer，导致包含 `$` 的技能/工具描述被 JavaScript 的特殊替换模式（如 `$&`）破坏。

10. **修复 OAuth 令牌交换中 keep-alive 套接字复用** [#28103](https://github.com/google-gemini/gemini-cli/pull/28103)  
    - 针对 Node.js 24.17.0、22.23.0、26.3.0 版本中 CVE-2026-48931 修复后的行为，避免 OAuth “Sign in with Google” 因 keep-alive 套接字提前关闭而失败。

## 功能需求趋势

- **AST 感知能力**：多个 issue 提及使用 AST 来提高文件读取、搜索、代码映射的精确度，减少 token 消耗和回合数。社区认为这是提升大型代码库体验的关键方向。
- **子代理行为可观测性**：用户希望子代理的轨迹能够通过 `/chat share` 分享，以便评估和调试（[#22598](https://github.com/google-gemini/gemini-cli/issues/22598)）；同时 bugreport 应包含子代理上下文（[#21763](https://github.com/google-gemini/gemini-cli/issues/21763)）。
- **自我意识与可靠性**：社区要求 Gemini CLI 能够准确知道自己支持的 CLI 标志、热键，并能正确执行自己（[#21432](https://github.com/google-gemini/gemini-cli/issues/21432)）。这反映出对 agent 自我认知能力的期待。
- **安全与隐私加固**：Auto Memory 的日志泄漏、确定性红action、信任对话框行为改进、多用户环境指导等成为近期安全话题的焦点。
- **终端体验优化**：包括终端 resize 时的高性能与防闪烁（[#21924](https://github.com/google-gemini/gemini-cli/issues/21924)）、外部编辑器退出后屏幕刷新恢复（[#24935](https://github.com/google-gemini/gemini-cli/issues/24935)），以及 `\n` 转义处理（[#22466](https://github.com/google-gemini/gemini-cli/issues/22466)）。

## 开发者关注点

- **agent 稳定性是首要痛点**：通用代理挂起、子代理误报成功、shell 命令卡住、交互式 prompt 中卡住（如创建 vite app）[#22465] 等现象严重破坏日常使用信心。
- **配置与使用预期不一致**：`settings.json` 中的子代理 `maxTurns` 等配置被浏览器代理忽略（[#22267]）；`~/.gemini/agents` 下的符号链接不被识别（[#20079]）；子代理在权限被禁用时仍被使用（[#22093]）。
- **调试与诊断信息不足**：bug report 缺少子代理内部上下文（[#21763]）；Auto Memory 的失败原因记录不透明（低信号会话无限重试 [#26522]）。
- **安全感知提升**：开发者不仅关注功能，还提出了信任对话框展示与实际执行不一致、OAuth 连接稳定性、MCP 资源包装等安全细节，表明随着 Gemini CLI 使用面扩大，安全设计成为重要考量。
- **对“技能/子代理”调用的期望**：用户反馈 Gemini 不会主动调用配置好的技能，即使描述明确。这会影响用户投入时间定制 agent 的意愿，建议通过机制鼓励主动调用（如自动尝试匹配工具描述与当前任务）。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 – 2026-06-26

## 今日速览
今天发布了 **v1.0.66-0**，重点增强了 MCP 服务器的交互管理（启用/禁用切换）并引入了实验性响应预算控制。社区方面，关于 **WSL2 ARM64 下的 `/copy` 故障** 和 **Voice 模式在企业 VPN 下不可用** 的 Bug 持续发酵；同时，**自定义 Agent 的 `skills` 字段** 和 **企业级服务器托管配置** 等新功能需求获得关注。

## 版本发布
### v1.0.66-0 更新内容
- **新增** MCP 列表视图中的启用/禁用开关，方便快速管理 MCP 服务器。
- **新增** CLI 设置中的实验性响应预算控制（Response Budget Controls），可限制单次对话的 token 消耗。
- **新增** 管理设置支持配置 OpenTelemetry 导出，便于企业级可观测性集成。
- **修复** 基于 OAuth 认证的远程 MCP 服务器在会话中令牌过期后会自动恢复连接。
- 发布链接：[v1.0.66-0](https://github.com/github/copilot-cli/releases/tag/v1.0.66-0)

## 社区热点 Issues（10 条）

### 1. #3501 – 滚动条导致文本对齐错乱（Windows）
- **标签**：`area:platform-windows`, `area:terminal-rendering`
- **摘要**：Windows Console Host 或 Terminal 上，垂直滚动条出现后文本渲染错位。用户尝试请求 Copilot 自修复无效。
- **社区反应**：5 条评论，9 个 👍，影响面广，Windows 用户高频反馈。
- **链接**：[Issue #3501](https://github.com/github/copilot-cli/issues/3501)

### 2. #3534 – WSL2 (ARM64) 下 `/copy` 因 cmd.exe 引号错误失败
- **标签**：`area:input-keyboard`, `area:platform-windows`
- **摘要**：1.0.55-1 起，WSL2 ARM64 上 `clip.exe` 退出码 1，`/copy` 命令失效。问题根源是 `cmd.exe` 引号封装 Bug。
- **社区反应**：4 条评论，4 个 👍，持续 1 个月未关闭，ARM 架构 WSL 用户关键痛点。
- **链接**：[Issue #3534](https://github.com/github/copilot-cli/issues/3534)

### 3. #3636 – Voice 模式无法启用：模型目录在企业 VPN 下不可达
- **标签**：`area:networking`, `area:models`
- **摘要**：`/voice` 失败，提示 “Failed to fetch model catalog”。CLI 在首次激活前必须能访问语音 STT 模型目录，企业网络限制导致完全无法使用。
- **社区反应**：3 条评论，5 个 👍，影响企业用户使用语音功能。
- **链接**：[Issue #3636](https://github.com/github/copilot-cli/issues/3636)

### 4. #3940 – 自定义 Agent 支持 `skills` 字段限制预加载技能
- **标签**：`area:agents`, `area:plugins`
- **摘要**：希望在自定义 Agent 配置（`.github/agents/*.md`）中定义 `skills` 字段，仅加载所需技能到上下文中，避免资源浪费。
- **社区反应**：新提交，2 条评论，0 👍，功能设计讨论初期，但代表了 Agent 精细化控制方向。
- **链接**：[Issue #3940](https://github.com/github/copilot-cli/issues/3940)

### 5. #3909 – 企业/组织服务器托管配置（含 `env`）支持本地 CLI
- **标签**：`area:enterprise`, `area:configuration`
- **摘要**：组织管理员无法集中推送配置（尤其是环境变量）到本地 Copilot CLI。现有方案仅支持 Codespaces/Agents 云环境，本地 CLI 无机制。
- **社区反应**：2 条评论，0 👍，但需求明确，是企业级管理的重要缺失。
- **链接**：[Issue #3909](https://github.com/github/copilot-cli/issues/3909)

### 6. #3876 – 退出后鼠标跟踪未正确关闭
- **标签**：`area:input-keyboard`, `area:terminal-rendering`
- **摘要**：CLI 退出后终端无法使用鼠标滚动。诊断发现 CLI 开启了 `alt screen` 和鼠标事件追踪，但未在退出时恢复。
- **社区反应**：2 条评论，0 👍，已关闭（已修复？），涉及终端交互基础体验。
- **链接**：[Issue #3876](https://github.com/github/copilot-cli/issues/3876)

### 7. #3941 – 安装包内二进制文件代码签名不一致
- **标签**：`area:installation`
- **摘要**：Copilot CLI 包中大部分二进制由 GitHub 签名，但 `mxc` 二进制由 Microsoft 签名，可能引起审计或信任问题。
- **社区反应**：1 条评论，0 👍，新提交，安全与合规敏感。
- **链接**：[Issue #3941](https://github.com/github/copilot-cli/issues/3941)

### 8. #3936 – Ctrl+G 应该展开粘贴令牌为全文（Claude Code 对等）
- **标签**：`area:input-keyboard`
- **摘要**：启用 `compactPaste` 后，粘贴大段内容折叠为 `[Paste #N - X lines]`。按下 Ctrl+G 编辑 prompt 时，该令牌被逐字写入临时文件，导致无法编辑实际内容。
- **社区反应**：0 条评论，0 👍，但属于用户期望的自然交互改进。
- **链接**：[Issue #3936](https://github.com/github/copilot-cli/issues/3936)

### 9. #3935 – VSCode 终端中忽略用户主题，默认使用浅色主题
- **标签**：`area:theming-accessibility`
- **摘要**：1.0.64/1.0.65 起，VSCode 终端内 Copilot CLI 不再遵循用户配置的深色或 Solarized 主题，始终显示浅色。
- **社区反应**：0 条评论，0 👍，主题回归问题影响易用性。
- **链接**：[Issue #3935](https://github.com/github/copilot-cli/issues/3935)

### 10. #3933 – Autopilot 模式每次请求结束后自动退出
- **标签**：`area:agents`
- **摘要**：用户反馈原先进 Autopilot（按 Shift+Tab 两次）后可连续对话，现在每个请求后自动退出 Autopilot，需要重新激活。界面颜色也从绿色变紫色。
- **社区反应**：0 条评论，0 👍，新提交，严重影响工作流效率。
- **链接**：[Issue #3933](https://github.com/github/copilot-cli/issues/3933)

## 重要 PR 进展

由于过去24小时内仅有 **1 个 PR** 更新，现列出该 PR：

### #3928 – 添加 `.gitignore` 和设置配置
- **作者**：tpsaint
- **状态**：OPEN
- **摘要**：添加 `.gitignore` 文件以及相关的设置配置（具体细节未详述）。可能用于规范化仓库开发环境。
- **链接**：[PR #3928](https://github.com/github/copilot-cli/pull/3928)

（注：其余 PR 数据未在报告周期内更新，故仅列上述一条。）

## 功能需求趋势
从近期 Issues 中提炼出社区最关注的方向：

1. **MCP 服务器精细化管理**：大量请求涉及 MCP 的启用/禁用交互（#2956 已关闭但实现至此版本）、异步执行（#3829）、注册表变量插值（#3887）、指令忽略（#1579）以及企业策略阻止（#3934）。
2. **自定义 Agent 与技能控制**：希望 Agent 支持 `skills` 字段限制加载（#3940）、Agent 与 `--acp` 非交互模式兼容（#3942）、Agent 配置持久化（#3938）。
3. **企业级配置与安全**：组织管理员需远程推送环境变量到本地 CLI（#3909）、CVE 申请与签名一致性（#3906/#3941）。
4. **终端体验与可访问性**：主题尊重（#3935）、鼠标/滚动修复（#3501/#3876）、粘贴令牌展开（#3936）、Autopilot 连续模式（#3933）。
5. **网络与模型服务独立性**：Voice 模式需要离线或可配置目录（#3636），减少对企业网络依赖。
6. **系统兼容性**：Windows/WSL2 ARM64 持续报错（#3534），说明多平台适配仍为短板。

## 开发者关注点
当前开发者的主要痛点和高频需求：

- **终端渲染回归**：Windows 滚动条对齐错乱、VSCode 主题忽略、退出后鼠标失效，严重影响日常使用。
- **WSL 与 ARM 支持**：`/copy` 命令在 ARM64 WSL2 上完全失效，开发者难以将 CLI 集成到 ARM 环境下。
- **Voice 功能受限**：企业 VPN 下无法加载模型目录，导致语音交互完全不可用，对远程办公者打击大。
- **Autopilot 行为突变**：自动退出模式打断连续对话，用户反馈“效率骤降”。
- **MCP 与策略冲突**：“MCP server blocked by policy”错误信息不透明，企业用户无法排查原因。
- **代码签名不一致**：可能触发安全软件告警，影响安装信任度。
- **缺少企业级管控**：本地 CLI 无法继承组织设置，与 Codespaces 体验断档。

> 📌 数据采集时间：2026-06-26 UTC。以上分析基于 GitHub 公开仓库 `github/copilot-cli` 的 Issues、PR 及 Release 信息。

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 (2026-06-26)

## 今日速览
过去24小时内社区变动较小，无新版本发布。两个新提交的 Issue 分别围绕 **MCP 工具兼容性**（#2475）和 **CLI 界面频繁抖动**（#2474）反馈，后者在 Linux 和 Windows 平台均有重现报告。同时，一个修复 `reasoning_effort` 参数序列化问题的 PR（#2476）已提交，另一个文档改进 PR（#2287）获得更新。

---

## 版本发布
过去24小时内无新版本发布。当前稳定版仍为 **Kimi Code v0.19.2**。

---

## 社区热点 Issues

### #2475 —— [bug] MCP tools
- **作者**: ptyll | **标签**: bug  
- **摘要**: 用户使用 Kimi Code v0.19.2，在订阅模式、模型 K2.7 下，MCP 服务器包含 212 个工具，但客户端未能正确加载或处理（描述不完整）。  
- **重要性**: 首次出现对较大规模 MCP 工具集合的兼容性反馈，可能暴露批量工具注册或分页传输问题。目前无评论，热度低但潜在影响面广。  
- **链接**: [Issue #2475](https://github.com/MoonshotAI/kimi-cli/issues/2475)

### #2474 —— 界面抖动/重新渲染整个对话
- **作者**: yudichimiantiao | **标签**: bug  
- **摘要**: 用户报告 CLI 界面在 Linux 系统上出现各种抖动，并莫名其妙地从头重新渲染整个对话。已使用 K2.7 Code thinking 模型，登录方式 `/login`。  
- **重要性**: 影响核心交互体验，可能关联终端渲染库（如 Ink、React Terminal）的刷新逻辑或数据流状态管理。零评论，但问题容易复现，值得跟进。  
- **链接**: [Issue #2474](https://github.com/MoonshotAI/kimi-cli/issues/2474)

---

## 重要 PR 进展

### #2287 —— docs(readme): add prerequisites list to Development section
- **作者**: ktwu01 | **标签**: docs  
- **摘要**: 在 README 的 Development 部分新增“Prerequisites”小节，明确列出贡献者在执行 `make prepare` 前需要安装的依赖项。解决 Issue #2274。  
- **重要性**: 降低了新贡献者的上手门槛，避免缺少环境变量或工具导致的失败。PR 自 2026-05-14 打开后，于今日更新（可能是 rebase 或采纳建议）。  
- **链接**: [PR #2287](https://github.com/MoonshotAI/kimi-cli/pull/2287)

### #2476 —— fix(kosong): omit reasoning_effort instead of sending null when thinking is off
- **作者**: logicwu0 | **标签**: fix  
- **摘要**: 当用户关闭推理（thinking off）时，`OpenAILegacy.with_thinking("off")` 会返回 Python `None`，导致 OpenAI SDK 序列化为 `"reasoning_effort": null`。此 PR 改为直接省略该参数（使用 `omit` 序列化），避免 API 报错。  
- **重要性**: 直接修复了在关闭推理能力时可能触发的 API 参数错误，属于低频率但关键的正确性修复。今日新提交，无评论，代码改动干净。  
- **链接**: [PR #2476](https://github.com/MoonshotAI/kimi-cli/pull/2476)

---

## 功能需求趋势
由于过去24小时内仅有两项 Issue，以下趋势基于有限数据推断：

1. **MCP 工具生态扩展** – Issue #2475 表明社区正在尝试集成大量自定义 MCP 工具，CLI 需支持 ≥200 工具的稳定注册与调用。  
2. **界面稳定性与刷新优化** – Issue #2474 反映出终端 UI 的渲染抖动是真实痛点，尤其在 Linux 长期运行场景下，可能需优化虚拟终端滚动/重绘策略。  
3. **模型参数正确性** – PR #2476 暗示社区对 `thinking` 功能开关的 API 兼容性敏感，开发者期望在关闭推理时不产生无效字段。

---

## 开发者关注点
1. **高频痛点**:  
   - **UI 抖动 & 意外重渲染** – 严重影响日常编码辅助流程，开发者期望快速修复。  
   - **MCP 工具批量加载失败** – 对于拥有大型工具集的用户，当前版本（v0.19.2）可能无法正常工作。  
2. **文档与贡献体验**:  
   - PR #2287 的长期静默后更新，侧面反映社区希望开发环境搭建指引更清晰，以便更多人参与贡献。  
3. **API 兼容性细节**:  
   - `reasoning_effort` 参数序列化行为已被注意到，开发者倾向使用更严格的 `omit` 而非 `null`，体现了对 OpenAI SDK 实现细节的关注。

---

*说明：因过去24小时内 GitHub 数据量有限（仅 2 个 Issue + 2 个 PR），本报告已围绕现有信息做重点分析。更全面的洞察需累积更多日常动态后呈现。*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-06-26

## 今日速览

OpenCode v1.17.11 今日发布，带来了会话快照与回滚控制这一重大改进。然而社区中最受关注的问题仍是 **CPU 高占用**——多个 Issue 指出 OpenCode 在空闲或 API 等待时大量消耗 CPU 资源，已成为当前最严重的性能瓶颈。此外，Windows ARM64 安装包损坏、GitHub Copilot 部分模型不可用等问题也集中爆发，开发者体验急需优化。

---

## 版本发布

### v1.17.11

**Core 改进**
- 新增**会话快照与回滚控制**：用户可以将当前会话回滚至早期的任意一条消息，包括相关文件变更，极大方便了实验性对话的管理。
- 修复 MCP OAuth 登录问题：现在即使在浏览器流程失败时，也会始终打印 OAuth URL，支持手动登录。

**Desktop 改进**
- 添加 Chrome 风格（部分 UI 改进，原文截断为 `Chrome-sty`，推测为窗口控件或标签页样式更新）。
- 其他稳定性修复。

---

## 社区热点 Issues

1. **#21470 OpenCode 严重 CPU 绑定**  
   [🔗](https://github.com/anomalyco/opencode/issues/21470)  
   **热度**：14 评论 / 13 👍  
   **重要性**：作者对比 Claude 与 OpenCode+Gemini-3.1，发现大部分时间消耗在 OpenCode 自身而非模型API。已花费 300k tokens 但 OpenCode 自身占用超过 1.5… 该问题与下面两个 issue 共同指向 CPU 性能危机。

2. **#30086 新版本 CPU 使用率飙升**  
   [🔗](https://github.com/anomalyco/opencode/issues/30086)  
   **热度**：11 评论 / 8 👍  
   **重要性**：用户报告约 7 天前更新后，CPU 占用激增，导致无法像以前同时运行 10 个会话。目前 3 个会话就会使鼠标卡顿。这是普遍性回归。

3. **#19466 空闲时 CPU 持续占用**  
   [🔗](https://github.com/anomalyco/opencode/issues/19466)  
   **热度**：10 评论 / 9 👍  
   **重要性**：即使在等待 API 限流重试（实际不做事）时，OpenCode 仍占用 i9-14900 单核心约 50% CPU。该问题自 v1.3.3 起存在，至今未修复。

4. **#15907 SSH + tmux 环境下剪贴板复制失效**  
   [🔗](https://github.com/anomalyco/opencode/issues/15907)  
   **热度**：9 评论 / 10 👍  
   **重要性**：在 Ghostty 终端通过 SSH 进入 tmux 后，复制操作显示“已复制”但系统剪贴板未更新。远程开发者的常见痛点。

5. **#34045 Desktop 动画导致 CPU 浪费**  
   [🔗](https://github.com/anomalyco/opencode/issues/34045)  
   **热度**：1 评论 / 1 👍（新发 issue，但非常典型）  
   **重要性**：明确指出在等待模型响应时，Desktop 的进度指示动画持续消耗 CPU 资源。开发者希望针对空闲状态降级渲染。

6. **#34038 v1.17.11 Desktop 渲染器因 ResizeObserver 循环冻结**  
   [🔗](https://github.com/anomalyco/opencode/issues/34038)  
   **热度**：2 评论（新 issue）  
   **重要性**：在 Windows x64 上，流式响应期间渲染器冻结，很可能与 ResizeObserver 循环问题有关。可能是 v1.17.11 引入的回归。

7. **#34036 Windows ARM64 安装包缺少 OpenCode.exe**  
   [🔗](https://github.com/anomalyco/opencode/issues/34036)  
   **热度**：2 评论（新 issue）  
   **重要性**：NSIS 安装器未部署主可执行文件，导致快捷方式无效。ARM64 用户无法使用 v1.17.11。

8. **#34048 GitHub Copilot 模型全部报“不支持的模型”**  
   [🔗](https://github.com/anomalyco/opencode/issues/34048)  
   **热度**：1 评论（新 issue）  
   **重要性**：认证成功、模型列表正常，但所有推理请求失败，返回“The requested model is not supported”。企业 Copilot 用户受限。

9. **#30675 Copilot 模型仅 gpt-5.2-codex 可用，缺 Copilot-Integration-Id 头**  
   [🔗](https://github.com/anomalyco/opencode/issues/30675)  
   **热度**：3 评论（已关闭）  
   **重要性**：虽已关闭，但揭示了 OpenCode 与 Copilot 交互时缺少必要 HTTP 头，可能是 #34048 的根因。社区关注度高。

10. **#33878 后台子任务完成通知在 TUI 中沉默丢失**  
   [🔗](https://github.com/anomalyco/opencode/issues/33878)  
    **热度**：3 评论  
    **重要性**：子任务完成后 <system-reminder> 不显示，用户需手动输入消息才能“唤醒”会话。影响多代理工作流的可靠性。

---

## 重要 PR 进展

1. **#34042 fix(app): 保存会话加载期间的实时消息**  
   [🔗](https://github.com/anomalyco/opencode/pull/34042)  
   通过三路合并，在加载会话历史时保留本地实时消息，避免因网络延迟导致的消息丢失。近期稳定性关键修复。

2. **#34039 fix(session): 阻止格式错误的 tool-call 循环**  
   [🔗](https://github.com/anomalyco/opencode/pull/34039)  
   修复 Claude Opus 4.8 返回 `finish=tool-calls` 但无结构化 tool call 时导致的死循环。直接影响 Agent 行为。

3. **#34037 fix(tui): 自动补全引用目录内的文件**  
   [🔗](https://github.com/anomalyco/opencode/pull/34037)  
   解决 `@alias/partial` 补全不显示嵌套文件的问题。提升 TUI 使用体验。

4. **#34046 fix(desktop): 显示中断加载状态**  
   [🔗](https://github.com/anomalyco/opencode/pull/34046)  
   在桌面版中断请求时增加进度指示与 `aria-busy` 状态，让用户明确操作正在执行。

5. **#33727 fix(ui): 防止 IME 输入法下 Enter 误提交**  
   [🔗](https://github.com/anomalyco/opencode/pull/33727)  
   修复日文/中文输入法下，自定义答案输入框按 Enter 直接提交而非完成输入的问题。非英文本地化用户必见。

6. **#33927 fix(vcs): 防止仓库有数千个未跟踪文件时崩溃**  
   [🔗](https://github.com/anomalyco/opencode/pull/33927)  
   当 Git 仓库含有 1200+ 未跟踪文件时，VCS 层导致程序崩溃。对大型项目开发者至关重要。

7. **#32767 fix(tui): 恢复 ESC 中断委托子会话功能**  
   [🔗](https://github.com/anomalyco/opencode/pull/32767)  
   修复回归，恢复之前版本中按 ESC 中断子代理会话的能力。结合 #32425 可以更好地控制子代理。

8. **#32425 feat(opencode): 中断运行中的子代理（引导/取消/中止）**  
   [🔗](https://github.com/anomalyco/opencode/pull/32425)  
   引入实验性原语，支持中断子代理的执行。为后续子代理编排功能打下基础。

9. **#33996 fix(tui): 保留渲染器初始化错误**  
   [🔗](https://github.com/anomalyco/opencode/pull/33996)  
   使 TUI 启动失败时能显示原始错误，而非泛化提示。有助于调试底层渲染库问题。

10. **#30543 fix(opencode): 暴露并使用认证确认码**  
    [🔗](https://github.com/anomalyco/opencode/pull/30543)  
    修复 OAuth 弹窗中确认码显示逻辑，确保桌面 UI 正确显示认证码。影响 OAuth 提供商登录流程。

---

## 功能需求趋势

- **子代理控制**：社区强烈呼吁能中断、引导或取消后台子代理任务（#21458、#23534、#32425）。
- **新模型 / Provider 支持**：持续有请求添加 Crof AI、DeepSeek 原生搜索、企业 Copilot 第三方模型等（#24636、#32273、#34030）。
- **会话管理增强**：会话快照/回滚（v1.17.11已实现）、复制特定助手回复（#34044）、页面刷新按钮（#34031）。
- **跨平台与终端兼容性**：RTL 文本支持、IME 输入法、SSH + tmux 剪贴板等仍需优化。

---

## 开发者关注点

- **CPU 性能是首要痛点**：空闲或低负载下 CPU 持续占用，社区已积累 3 个高热度 Issue 且长时间未解决，严重影响多会话工作流。
- **Desktop 版本稳定性不足**：Windows ARM64 安装包损坏、v1.17.11 渲染器冻结、Copilot 模型大面积不可用等，新版本发布后出现多处回归。
- **子代理与后台任务交互混乱**：任务完成通知丢失、Tool-call 格式错误导致无限循环、无法中断等，是 Agent 自动化场景的核心障碍。
- **TUI 插件系统尚不成熟**：多个插件相关 issue（如 sidebar 不渲染、流输出无法过滤）表明插件 API 仍处于早期阶段，文档和测试需加强。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

好的，这是为您生成的 2026-06-26 Pi 社区动态日报。

---

# Pi 社区动态日报 | 2026-06-26

## 今日速览

今日社区焦点集中在 **连接可靠性与会话管理** 问题上：`openai-codex` 交互式模式下的无响应悬案 (#4945) 获得社区最高关注（71条评论），表明稳定性仍是核心痛点。同时，一个实验性的 **Pi Orchestrator** 守护进程 PR (#6064) 正在探索多实例管理的新方向，为分布式或复杂工作流提供可能。此外，**新的 Provider 支持** (Friendli) 以及 **Session RPC 扩展** 也是今天的重要进展。

## 版本发布

无

## 社区热点 Issues

1.  **[#4945] openai-codex 连接可靠性问题**
    - **链接**: [Issue #4945](https://github.com/earendil-works/pi/issues/4945)
    - **原因**: 该 Issue 描述了使用 `openai-codex` / `gpt-5.5` 时，交互式 TUI 经常卡死在 `Working...` 状态，无任何输出或错误提示，只能通过 Escape 键退出。该问题在过去两天内反复出现，且高居评论榜首 (71 条)，是当前最影响用户体验的痛点。
    - **社区反应**: 开发者 `liushuaiiu` 创建，获得 30 个 👍，大量用户表达遇到了相同问题，并有开发者正在进行初步排查。

2.  **[#5825] 流式 Markdown 渲染强制滚动到底部**
    - **链接**: [Issue #5825](https://github.com/earendil-works/pi/issues/5825)
    - **原因**: 当 Agent 流式输出 Markdown 时，用户向上滚动阅读，几秒后 TUI 会强制将滚动条拉到底部，严重影响阅读体验。此问题仅在启用 `clear on shrink` 设置时触发，涉及 TUI 的核心渲染逻辑。
    - **社区反应**: 31 条评论，用户 `xl0` 描述了精确的复现步骤，表明这是一个常见的交互冲突。

3.  **[#4877] 会话文件夹冲突**
    - **链接**: [Issue #4877](https://github.com/earendil-works/pi/issues/4877)
    - **原因**: 不同路径的会话可能因路径名哈希碰撞而存入同一个文件夹，例如 `/a/b/c/d` 和 `/a-b/c-d` 会映射到同一路径。这是一个潜在的静默数据覆盖问题，可能在未来导致用户会话丢失或混乱。
    - **社区反应**: 19 条评论，用户 `olivierverdier` 清晰指出了设计缺陷。

4.  **[#6083] z.ai GLM 模型缓存未正常工作**
    - **链接**: [Issue #6083](https://github.com/earendil-works/pi/issues/6083)
    - **原因**: 使用 `z.ai GLM Coding Plan` 时，每次工具调用（write/edit）都会大量消耗会话预算（10-20%），即使在上下文压缩活跃的情况下。这表明 LLM 缓存机制对该模型失效，导致成本飙升。
    - **社区反应**: 获得 7 个 👍，用户 `skhoroshavin` 的反馈指出了模型适配中的一个关键性能问题。

5.  **[#6060] TUI 底部栏渲染崩溃**
    - **链接**: [Issue #6060](https://github.com/earendil-works/pi/issues/6060)
    - **原因**: 当会话中包含只有工具调用的助手消息（无文本内容）时，TUI 底部栏在渲染 Token 统计信息时触发 `TypeError: content is not iterable`，导致程序崩溃。这是一个典型的边界情况 Bug。
    - **社区反应**: 用户 `wickedTangent` 快速定位并提交了 Issue，有助于提升 TUI 稳定性。

6.  **[#6002] `SessionManager.open()` 静默截断非会话文件**
    - **链接**: [Issue #6002](https://github.com/earendil-works/pi/issues/6002)
    - **原因**: 这是一个危险的操作问题。使用 `SessionManager.open(path)` 或 `--session <path>` 打开一个非 Pi 会话文件（如 3.2MB 的日志），会将其静默截断为一个只有 133 字节的会话头，且无任何警告或备份。
    - **社区反应**: 用户 `e4779` 指出这是一个严重的破坏性行为，社区反应强烈，建议作为优先级最高的 Bug 处理。

7.  **[#4290] 因长度被截断的消息被当作正常停止**
    - **链接**: [Issue #4290](https://github.com/earendil-works/pi/issues/4290)
    - **原因**: 当模型的思考/回答因长度限制被截断时，TUI 界面显示为正常完成，用户可能无法及时发现 Agent 被迫中断，导致回复不完整。
    - **社区反应**: 用户 `DanielThomas` 指出了这个容易被忽略的 UX 问题，影响对 Agent 结果的信任度。

8.  **[#5886] AgentSession 生命周期与延续性 Bug**
    - **链接**: [Issue #5886](https://github.com/earendil-works/pi/issues/5886)
    - **原因**: 这是一个元问题（meta issue），汇总了 Agent 会话在运行后逻辑中，试图从已经无效的副本继续处理而引发的一系列 Bug。涉及会话结算和助理尾部（assistant-tail）的生命周期管理。
    - **社区反应**: 由知名开发者 `mitsuhiko` 提交，当前标记为 `inprogress`，表明核心开发团队正在关注并重构相关逻辑。

9.  **[#6088] RpcClient 硬编码 60 秒超时导致长工具任务失败**
    - **链接**: [Issue #6088](https://github.com/earendil-works/pi/issues/6088)
    - **原因**: `packages/coding-agent` 中的 `RpcClient` 在 `waitForIdle()` 等方法中硬编码了 60 秒等待上限。当 MCP 服务器执行长时间任务（如 `searchbot`）时，即使任务仍在进行，客户端也会超时失败。
    - **社区反应**: 用户 `mizuikki` 报告并已提交修复 PR，该问题对使用复杂工具链的用户影响较大。

10. **[#6085] 编译后二进制无法解析 npm 子包**
    - **链接**: [Issue #6085](https://github.com/earendil-works/pi/issues/6085)
    - **原因**: 由 Bun 编译的 `pi` 二进制文件（Release 版本）存在一个解析器问题，导致扩展（如 `pi-lens`）中所有从子目录 `node_modules` 导入的 npm 依赖加载失败。这影响了所有在扩展中使用 npm 高级功能的用户。
    - **社区反应**: 该 Issue 是 #5949 的后续，发现影响范围更广，严重阻碍了扩展生态的发展。

## 重要 PR 进展

1.  **[#6064] feat(experimental): pi orchestrator**
    - **链接**: [PR #6064](https://github.com/earendil-works/pi/pull/6064)
    - **说明**: 这是一个实验性 PR，引入了 `@earendil-works/pi-orchestrator` 包。它作为一个本地守护进程（daemon），通过 Unix Socket 提供 JSON IPC 接口，用于启动、列出和管理 Pi 实例。这是向多实例、复杂工作流管理迈出的重要一步。

2.  **[#6087] fix(coding-agent): remove hardcoded RPC wait timeout**
    - **链接**: [PR #6087](https://github.com/earendil-works/pi/pull/6087)
    - **说明**: 紧急修复 #6088，移除了 RPC 客户端中硬编码的 60 秒等待上限，并引入了可配置的 `RpcClientOptions.waitTimeout`。这使得用户可以根据工具执行时长调整超时设置，避免了长时间任务的意外失败。

3.  **[#6090] feat(ai): add Friendli provider**
    - **链接**: [PR #6090](https://github.com/earendil-works/pi/pull/6090)
    - **说明**: 新增 Friendli 作为内置 OpenAI 兼容的 Provider，端点指向 `api.friendli.ai`，默认模型为 `zai-org/GLM-5.2`。这为社区提供了又一个 API 选择，降低了对单一 Provider 的依赖。

4.  **[#5515] feat(coding-agent): add alwaysTrust setting**
    - **链接**: [PR #5515](https://github.com/earendil-works/pi/pull/5515)
    - **说明**: 为 Coding Agent 添加了 `alwaysTrust` 设置，允许用户完全跳过项目信任门控（trust gating）。对于个人开发者或信任的工作环境，可以简化启动流程。

5.  **[#6084] fix(tui): preserve custom widget render order**
    - **链接**: [PR #6084](https://github.com/earendil-works/pi/pull/6084)
    - **说明**: 修复了在后台 TUI 刷新时，自定义 Widget 因 `Map` 操作导致的渲染顺序错乱问题。此前，高频刷新会使插件 Widget 的顺序“随机”跳动，影响定制化界面的一致性。

6.  **[#6081] feat: add #RRGGBBAA alpha support for theme colors**
    - **链接**: [PR #6081](https://github.com/earendil-works/pi/pull/6081)
    - **说明**: 终端主题色现在支持 8 位十六进制格式（`#RRGGBBAA`），允许设置透明度。由于终端本身不支持透明色，该 PR 在主题加载时自动将 Alpha 值与背景色进行混合，模拟了透明度效果，提升了自定义界面的灵活性。

7.  **[#6078] feat(coding-agent): add get_entries and get_tree RPC commands**
    - **链接**: [PR #6078](https://github.com/earendil-works/pi/pull/6078)
    - **说明**: 为 RPC 命令集新增了两个只读命令：`get_entries`（获取会话条目列表，支持游标）和 `get_tree`（获取会话树结构）。这对于需要外部审计、同步或高级会话管理的工具和扩展至关重要。

8.  **[#6074] fix(coding-agent): avoid pre-prompt compaction continue**
    - **链接**: [PR #6074](https://github.com/earendil-works/pi/pull/6074)
    - **说明**: 修复了 Coding Agent 在压缩预提示（pre-prompt）后试图“继续”的 Bug。此问题可能导致代码生成中的上下文错乱，此 PR 旨在避免在压缩发生后进行不安全的延续操作。

9.  **[#5832] fix(ai): surface provider HTTP error body**
    - **链接**: [PR #5832](https://github.com/earendil-works/pi/pull/5832)
    - **说明**: 修复了 #5763。当通过代理/网关请求 Provider 时，非 2xx 响应（如 403）的 HTTP body 被大多数 SDK 丢弃，导致显示为模糊的 `UnknownError`。此 PR 确保 Provider 的错误详情能被正确传递给用户，极大方便了调试。

10. **[#5270] [inprogress] Ephemeral session model and thinking level selection**
    - **链接**: [PR #5270](https://github.com/earendil-works/pi/pull/5270)
    - **说明**: 这是一个长期在线的 PR，旨在让 `setModel()`、`cycleThinkingLevel()` 等命令默认只影响当前会话，而非全局设置。除非传入 `{ persist: true }` 标记，否则这些变更将随会话关闭而丢弃，避免了临时配置污染全局默认值。

## 功能需求趋势

- **Provider 多样化与兼容性**：社区持续要求接入更多 Provider (如 Friendli) 以及修复现有 Provider (如 Anthropic scoped keys、MiniMax 上下文限制) 的适配问题。这表明用户希望摆脱单一供应商，追求更高的可用性和成本效益。
- **TUI 渲染稳定性与交互体验**：Markdown 滚动、Widget 排序、文件补全、Shell 自动补全等问题频发，反映出社区对 TUI 作为主要交互界面的“顺手”和“可靠”有很高要求，任何渲染闪烁或操作歧义都被视为核心 Bug。
- **扩展生态健壮性**：编译后二进制无法加载 npm 包 ( #6085 ) 和 Provider 层面缺乏扩展点 ( #6089 ) 的讨论，揭示了社区对扩展系统深度和稳定性的迫切需求。
- **会话与状态管理**：从会话文件夹碰撞、静默文件截断到会话生命周期 Bug ( #5886 )，用户期望拥有更安全、清晰、可控的会话管理机制。
- **分布式与系统集成**：`pi orchestrator` PR 的出现和 `Session RPC` 扩展，标志着社区开始探索将 Pi 嵌入到更复杂的系统工作流中，实现多实例编排和外部工具集成。

## 开发者关注点

- **输入/输出（I/O）与控制流问题**：`openai-codex` 的卡死、流式渲染的滚动强制、以及工具超时问题，都是非常底层的 I/O 和并发控制问题。这些是当前开发者面临的最头疼的痛点，直接导致工作流中断或数据丢失。
- **静默错误与数据安全**：`SessionManager.open()` 对非会话文件的静默截断（#6002）是极其危险的。开发者强烈呼吁所有危险操作都应有明确的警告、备份机制或回滚选项，避免因误操作造成数据损失。
- **缓存与 Token 消耗**：LLM 缓存对特定模型 (如 z.ai GLM) 失效，导致 Token 消耗失控，直接关系到开发者的使用成本。这是模型适配中一个非常现实且优先级高的问题。
- **扩展开发的兼容性**：最终编译的二进制版本在加载 npm 依赖上与本地开发环境表现不一致 ( #6085 )，这给扩展开发者带来了额外的测试和适配负担，阻碍了社区贡献。
- **配置持久化粒度**：用户希望在会话中进行临时性模型或配置切换，而不影响全局默认值。`#5270` 的长期存在反映了对这一特性的强烈渴望，开发者需要一个清晰的 `persist` 语义来区分全局和局部配置。

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

好的，作为专注于AI开发工具的技术分析师，我已根据您提供的GitHub数据，为您生成了2026年6月26日的Qwen Code社区动态日报。

---

# Qwen Code 社区动态日报 | 2026-06-26

## 今日速览

今日社区动态活跃，多个重大功能提案与核心Bug修复并行推进。**多人协作代理**“qwen tag”与**Git共享团队记忆**功能正式进入RFC阶段，社区反馈积极。同时，一个**严重的CI安全漏洞**导致跨PR状态污染，引发了开发者对基础设施隔离性的广泛讨论。此外，**循环任务持久化**和**计划审批门**功能也迎来了重要迭代。

## 社区热点 Issues (Top 10)

1.  **[#5887] feat(channels): “qwen tag” — 持久化多人频道常驻代理**
    -   **重要性**：⭐⭐⭐⭐⭐ 社区新功能风向标。该提案旨在将`qwen-code`的频道能力从“每人一个私有Bot”升级为“一个群一个共享助手”，实现多人在同一会话中协作。这个“qwen tag”概念直接对标业界成熟产品，社区反馈非常热烈（获得1个👍），被视为Qwen Code从个人工具走向团队协作的关键一步。
    -   **链接**: [QwenLM/qwen-code Issue #5887](https://github.com/QwenLM/qwen-code/issues/5887)

2.  **[#5882] Qwen agent CI作业在共享ECS Runner上非隔离运行 → 跨PR状态污染**
    -   **重要性**：⭐⭐⭐⭐⭐ 严重Bug，影响开发信任。这是一个严重的CI/CD安全与稳定性问题，CI工作流错误地将PR #5872的审核评论发布到了PR #5874上。这不仅干扰了正常的开发流程，也暴露了CI基础设施在隔离性方面的缺陷，社区对此表达了高度担忧。
    -   **链接**: [QwenLM/qwen-code Issue #5882](https://github.com/QwenLM/qwen-code/issues/5882)

3.  **[#5881] 提案：将“计划审批门”开放给所有计划模式入口**
    -   **重要性**：⭐⭐⭐⭐ 核心工作流增强。该提案旨在让“计划审批门”（Plan Approval Gate）不仅适用于模型发起的计划，也适用于用户手动进入的计划模式。这能显著提升AUTO/YOLO模式下计划的安全性和可控性，是社区广泛讨论的**need-discussion**议题。
    -   **链接**: [QwenLM/qwen-code Issue #5881](https://github.com/QwenLM/qwen-code/issues/5881)

4.  **[#5889] [loop] 添加一个在启动时注入的 .qwen/loop.md 任务文件**
    -   **重要性**：⭐⭐⭐⭐ 自动化工作流关键增强。该提案针对长期运行的`/loop`任务，提供了一个持久化、用户可编辑的任务列表文件。解决了当前循环任务需在每个Tick重新陈述任务、无法中途编辑指令的痛点，对后台自动化场景意义重大。
    -   **链接**: [QwenLM/qwen-code Issue #5889](https://github.com/QwenLM/qwen-code/issues/5889)

5.  **[#5875] 改进技能命令自动补全匹配**
    -   **重要性**：⭐⭐⭐⭐ 提升CLI日常使用体验。当前`/skill_name`命令需要从头精确匹配才能触发补全，该特性请求建议实现模糊匹配，让用户通过输入技能名称的任意部分即可获得建议，例如输入`/store`即可匹配`front-end-store-rules`。这是一个能显著提升开发者效率的改进。
    -   **链接**: [QwenLM/qwen-code Issue #5875](https://github.com/QwenLM/qwen-code/issues/5875)

6.  **[#5800] bug(cli): 超长回复最后一行被覆盖**
    -   **重要性**：⭐⭐⭐ 影响核心TUI渲染体验。在默认的静态渲染模式下，当助手回复内容超过终端高度时，回复的**最后一行会被隐藏**。这是来自上游Ink库的Bug，社区有4条评论讨论，表明这是一个普遍存在的体验问题。
    -   **链接**: [QwenLM/qwen-code Issue #5800](https://github.com/QwenLM/qwen-code/issues/5800)

7.  **[#5894] 编辑工具结果摘要反复追加到每次响应中**
    -   **重要性**：⭐⭐⭐ 影响对话连贯性的Bug。使用编辑工具修改文件后，其“文件已更改”的差异摘要会持续出现在后续所有回复中，导致对话上下文噪音累积，严重影响长对话的阅读体验。
    -   **链接**: [QwenLM/qwen-code Issue #5894](https://github.com/QwenLM/qwen-code/issues/5894)

8.  **[#5897] 重复的“unknown format "uint64" ignored in schema” 错误消息污染界面**
    -   **重要性**：⭐⭐⭐ 启动时的噪音问题。`qwen`启动时，大量关于未知格式的校验错误信息会污染用户界面。虽然优先级为P3，但这种非功能性错误会降低工具的专业形象，开发者社区对此有所反馈。
    -   **链接**: [QwenLM/qwen-code Issue #5897](https://github.com/QwenLM/qwen-code/issues/5897)

9.  **[#5677] tracking(serve): 跟踪ACP功能的差距**
    -   **重要性**：⭐⭐⭐ 代理通信协议（ACP）演进路线图。这是一个追踪Issue，系统地记录了在`/lsp`、`/permissions`、`/setup-github`等命令上补齐ACP支持的进度。它展示了Qwen Code在标准化API和生态兼容性方面的长远规划。
    -   **链接**: [QwenLM/qwen-code Issue #5677](https://github.com/QwenLM/qwen-code/issues/5677)

10. **[#5883] 提案：将聊天面板统一到 web-shell 上**
    -   **重要性**：⭐⭐⭐ UI架构统一方向。该提案建议将聊天面板（输入框+对话流）标准化到`packages/web-shell`上，以覆盖Web Shell、VSCode Webview和桌面端。这预示着未来Qwen Code的UI组件将趋向统一和可移植，属于**need-discussion**的架构级提案。
    -   **链接**: [QwenLM/qwen-code Issue #5883](https://github.com/QwenLM/qwen-code/issues/5883)

## 重要 PR 进展 (Top 10)

1.  **[#5888] feat(channels): qwen tag — RFC + Phase 0**
    -   **内容**：为“qwen tag”多人频道常驻代理功能提交的RFC（请求评议）草案和第一阶段实现。该PR定义了功能和架构，是今日最重磅的PR之一。
    -   **链接**: [QwenLM/qwen-code PR #5888](https://github.com/QwenLM/qwen-code/pull/5888)

2.  **[#5886] feat(memory): 添加一个Git共享的团队记忆层**
    -   **内容**：引入可选的第三层自动记忆——**TEAM**（团队）。该记忆存储在仓库中并通过Git共享，与已有的私有USER和PROJECT记忆层级协同工作，是实现团队知识共享的关键基础设施。
    -   **链接**: [QwenLM/qwen-code PR #5886](https://github.com/QwenLM/qwen-code/pull/5886)

3.  **[#5852] feat(daemon,sdk): 可恢复的 /acp 会话流**
    -   **内容**：为`/acp`流式HTTP会话添加了断点续传功能。通过标准的SSE `id:`字段，客户端在重连后可以从断点恢复，无需重头开始，对长会话和网络不稳定的场景至关重要。
    -   **链接**: [QwenLM/qwen-code PR #5852](https://github.com/QwenLM/qwen-code/pull/5852)

4.  **[#5890] feat(loop): 通过哨兵文件注入 .qwen/loop.md 任务文件**
    -   **内容**：实现了Issue #5889中提出的功能，为`/loop`命令添加了一个持久化的任务文件。代理会在每次触发时读取该文件，使用户可以随时编辑任务列表。
    -   **链接**: [QwenLM/qwen-code PR #5890](https://github.com/QwenLM/qwen-code/pull/5890)

5.  **[#5896] feat(cua-driver): 引入qwen-cua-driver，支持可选的0-1000相对坐标模式**
    -   **内容**：将`trycua/cua`后台自动化驱动集成到仓库中，并增加了0-1000相对坐标模式。这使得**Qwen-VL的computer_use**功能能够直接驱动桌面GUI自动化，是打通视觉与操作的关键桥梁。
    -   **链接**: [QwenLM/qwen-code PR #5896](https://github.com/QwenLM/qwen-code/pull/5896)

6.  **[#5847] feat(serve): 添加运行时上下文注入**
    -   **内容**：为`qwen serve`守护进程添加了运行时上下文存储与注入能力。外部调用者可以在每个会话中注入动态上下文，这些上下文会作为系统提示注入到每次交互中，实现了系统提示的动态化。
    -   **链接**: [QwenLM/qwen-code PR #5847](https://github.com/QwenLM/qwen-code/pull/5847)

7.  **[#5892] fix(core): 在Windows上通过tree-kill清理PTY Shell树以解决pwsh泄漏**
    -   **内容**：修复了Windows上一个严重的进程泄漏问题。之前的`ptyProcess.kill()`无法清理PowerShell进程树，该PR引入了`tree-kill`机制来彻底解决此问题。
    -   **链接**: [QwenLM/qwen-code PR #5892](https://github.com/QwenLM/qwen-code/pull/5892)

8.  **[#5856] feat(desktop): 桌面端的语音听写功能**
    -   **内容**：为桌面应用带来了语音听写功能，与CLI和Web Shell功能保持一致。用户在输入框中点击麦克风按钮即可通过语音输入，支持波形显示和停止操作。
    -   **链接**: [QwenLM/qwen-code PR #5856](https://github.com/QwenLM/qwen-code/pull/5856)

9.  **[#5879] feat(web-shell): 在 /mcp 对话框中浏览MCP服务器资源**
    -   **内容**：将Web Shell的`/mcp`对话框功能提升至与终端UI一致的水平。现在用户可以在Web Shell中浏览MCP服务器的资源和提示（Prompts），提升了Web平台的生态管理能力。
    -   **链接**: [QwenLM/qwen-code PR #5879](https://github.com/QwenLM/qwen-code/pull/5879)

10. **[#5898] Fix mid-input skill command completion**
    -   **内容**：修复了在输入中间键入技能命令时的补全问题。之前只能在行首触发补全，现在用户在任意文本后键入`/command`也能获得模糊匹配的补全菜单。
    -   **链接**: [QwenLM/qwen-code PR #5898](https://github.com/QwenLM/qwen-code/pull/5898)

## 功能需求趋势

从今日的Issues和PRs来看，社区最关注的功能方向集中在以下三点：

1.  **团队协作与共享**：无论是“qwen tag”（#5887, #5888）的多人频道代理，还是“team-memory”（#5886）的Git共享记忆，都明确指向了Qwen Code从单用户工具向**团队协作平台**演进的强烈需求。
2.  **后台自动化与持久化**：`/loop`任务的持久化（#5889, #5890）以及CUA驱动的引入（#5896），表明社区正推动Qwen Code朝着更可控、更持久的**后台自动化任务**方向发展，目标是让其成为真正能自主工作的“数字员工”。
3.  **UI/UX 统一与标准化**：将聊天面板统一到web-shell的提案（#5883）以及将MCP资源浏览功能同步到Web端（#5879），显示出社区希望**在不同终端（CLI、Web、Webview、Desktop）上获得一致且丰富的交互体验**。技能命令的模糊补全（#5875, #5898）也是这一趋势的具体体现。

## 开发者关注点

根据今日的Bug类Issue和社区反馈，开发者的主要痛点和关注点包括：

-   **CI/CD基础设施的可靠性**：CI Runner的非隔离运行导致的跨PR污染（#5882）和夜间版本发布失败（#5877），是当前最令开发者困扰的痛点之一，直接影响开发流程和发布效率。
-   **界面与渲染Bug**：超长回复被覆盖（#5800）和编辑结果反复追加（#5894）都是严重影响日常使用体验的Bug，修复优先级较高。启动时的错误信息污染（#5897）也降低了工具的专业感。
-   **窗口化平台兼容性**：Windows平台下PTY进程泄漏的修复（#5892）反映出跨平台，特别是Windows系统下，仍存在一些需要持续关注和解决的环境兼容性问题。

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，各位开发者，早上好！

这里是 2026 年 6 月 26 日的 DeepSeek TUI (CodeWhale) 社区动态日报。我是你们的技术分析师。

---

### 1. 今日速览

今天 CodeWhale 社区异常活跃，核心动态集中在三个方面：**1)** 发布了 v0.8.65 版本，核心是 **CodeWhale 品牌名正式定型**，并敦促遗留用户尽快迁移；**2)** 围绕 `plan` 与 `agent` 模式混淆、权限控制、以及 IME 输入冲突的 Bug 修复取得了关键进展；**3)** 社区贡献者积极提交了关于**权限规则持久化**、**桥接平台优化**以及**文档与 CI 流程**的多个高质量 PR，显示了项目生态的健康成长。

---

### 2. 版本发布

- **v0.8.65：** 过去 24 小时内最重要的发布。该版本是一个 **品牌重塑 (Rebrand) 里程碑**。
    - **核心变化：** 项目名称、命令、npm 包和发布资产均正式更名为 `CodeWhale`。旧有的 `deepseek-tui` npm 包已被标记为弃用，将不再获得任何后续更新。
    - **迁移指南：** 所有仍在使用 `v0.8.x` 旧名称 (`deepseek` / `deepseek-tui`) 的用户，应立即参考项目中的 `docs/REBRAND.md` 文档完成迁移。
    - [查看发布详情](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.65)

---

### 3. 社区热点 Issues

以下 10 个 Issues 是过去 24 小时内讨论最热烈或对项目未来方向最具影响力的议题：

1.  **[严重] `plan` 和 `agent` 模式再次混淆 (#3568)**
    - **为什么重要：** 用户 `DracheTek` 上传了详尽日志，证明了在 `plan` 模式下，AI 仍然错误地执行了 `agent` 的文件修改操作。这直接影响了工具模式的核心工作流，是社区非常关注的回归性 Bug。
    - **社区反应：** 获得了 1 个 👍️，且包含详细复现场景，开发者已将其标记为“bug”。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/3568)

2.  **[EPIC] 命令边界重构 (#2870)**
    - **为什么重要：** 这是一个巨大的跟踪 Issue，关联着多个子任务和 PR，旨在重构 `code`、`config` 等命令的边界，以支持更复杂的 `v0.9.0` 工作流。这决定了 CodeWhale 的未来架构。
    - **社区反应：** 社区成员 `aboimpinto` 正在通过多个分层 PR 稳步推进，今天的 PR #3652 即是其一部分。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/2870)

3.  **[EPIC] 权限规则系统 (#1186)**
    - **为什么重要：** 官方计划为执行策略层增加类型化的持久化权限规则，允许用户更精细地控制工具调用（按工具名、命令前缀、路径等设置 `allow/deny/ask`）。
    - **社区反应：** 该 Issue 持续获得更新，今天有多个相关 PR (如 #3651, #3650) 被提出，社区对权限控制的需求非常旺盛。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/1186)

4.  **[EPIC] 遗留代码与死代码整理 (#3490)**
    - **为什么重要：** 在进入 `v0.9.0` 前，项目需要清理大量的 `allow(dead_code)` 标记和过时的“后续跟进”注释。这将提升代码质量和编译速度，是重要的基础设施工作。
    - **社区反应：** 贡献者 `cy2311` 在 PR #3636 中完成了对约 330 处死代码的注释和引用追踪，回应了此 Issue。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/3490)

5.  **主提示词外部化 (#3638)**
    - **为什么重要：** 用户 `DracheTek` 提出，希望能将硬编码的系统提示词（如 `AGENTS.md`）指向 `config` 目录中的文件，以便于用户自定义。这为 CodeWhale 用于非软件工程场景（如文学创作）打开了大门。
    - **社区反应：** 这是一个新提交的增强建议，代表了社区对更高定制化的渴望。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/3638)

6.  **Setup 向导：工具、MCP、技能与插件 (#3407)**
    - **为什么重要：** 官方计划提供一个统一的安装向导步骤，指导用户发现、安装并验证 MCP服务器、技能和插件。这将极大改善新用户的上手体验。
    - **社区反应：** 该 Issue 处于开放状态，`cy2311` 的 PR #3643 已经迈出了实现该功能的第一步。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/3407)

7.  **`thinking` 块坍塌问题 (#861)**
    - **为什么重要：** 这是一个多 root cause 的严重 Bug，导致推理模型的思考过程冻结、截断或丢失，甚至引发后续请求 HTTP 400 错误。是“所有推理模型”的顶级社区 Bug。
    - **社区反应：** 该 Issue 虽已关闭，但关联的修复 PR (#3016) 的讨论仍在持续，其影响之大使社区时刻关注其根因是否被完全解决。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/861)

8.  **压缩默认提示词，对标 Codex (#2953)**
    - **为什么重要：** CodeWhale 的基础提示词很大，导致输入 Token 消耗远超同类工具 Codex CLI。优化提示词可以显著降低成本并提升响应速度。
    - **社区反应：** 开发者正在进行精细的 Token 使用分析，社区期待通过优化让模式本身更具竞争力。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/2953)

9.  **减少 Benchmark/Exec 回合中的重复输入 (#2956)**
    - **为什么重要：** 同样是为了对标 Codex CLI，减少在 benchmark 和执行回合中重复发送的“历史记录”和工具输出，是降低 Token 消耗的另一大关键步骤。
    - **社区反应：** 此 Issue 与 #2953 和 #2957 构成了一个系列优化，社区对性能优化的关注度很高。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/2956)

10. **用专用 i18n 库替换硬编码本地化文件 (#3537)**
    - **为什么重要：** `localization.rs` 文件已超过 5000 行，全是硬编码，维护困难并拖慢编译速度。更换为专门的 i18n 库是提升项目可维护性的重要举措。
    - **社区反应：** 该 Issue 已被关闭，有相关的 PR 正在处理，社区贡献者正在积极解决这个“技术债”。
    - [查看 Issue](https://github.com/Hmbown/CodeWhale/issues/3537)

---

### 4. 重要 PR 进展

以下是过去 24 小时内最值得关注的 10 个 PR：

1.  **`feat(tui): show ask-rule save details in approval cards` (#3651)**
    - **摘要：** 增强权限控制 UI。当用户按下 `S` 键保存“询问 (ask)”规则时，会弹出一个预览卡片，清晰展示将要保存的规则细节（如 `tool=exec_shell command=...`）。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3651)

2.  **`feat(bridge): support natural-language approval responses` (#3637)**
    - **摘要：** 再次提升易用性！现在用户在 Telegram/WeCom 等桥接平台上，可以直接回复“允许”、“ok”、“好”等自然语言来批准工具调用，无需再复制 ID。由社区贡献者 `pkeging` 提交。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3637)

3.  **`Permission control: deny, allow, and ask actions in permissions.toml` (#3650)**
    - **摘要：** 社区开发者 `yekern` 提出了一个独立于官方规划的权限实现。它将允许用户在 `permissions.toml` 中为每个规则定义 `deny/allow/ask` 动作，非常类似 Claude Code。这显示了社区对权限控制的创新思路。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3650)

4.  **`Guard exec against misplaced global flags` (#3645)**
    - **摘要：** 修复了一个易用性 Bug。当用户在 `codewhale exec` 命令后错误地放置全局参数时（如 `codewhale exec --provider ...`），现在会给出明确的错误提示，而不是默默地将其当作提示词文本。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3645)

5.  **`ci: add DCO (Developer Certificate of Origin) check` (#3649)**
    - **摘要：** 这是一个重要的流程改进。新增了 DCO 检查，要求所有 PR 的提交都必须包含 `Signed-off-by` 签名。这能保证贡献来源的合法性，是项目走向成熟的标志。由社区贡献者 `pkeging` 提交。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3649)

6.  **`Codex/lsp php custom servers` (#3624)**
    - **摘要：** 另一个来自 `yekern` 的出色贡献！该 PR 为内置的 LSP 注册表增加了 PHP (`intelephense`)，更关键的是，引入了 `[lsp.custom]` 配置节，允许用户为任意文件后缀注册 LSP 服务器。极大扩展了对 Ruby, C#, Swift 等语言的支持。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3624)

7.  **`feat(setup): add setup summary wizard step for MCP/skills/plugins overview` (#3643)**
    - **摘要：** 实现了 `v0.8.67` Setup 向导的第一步：`SetupSummaryView`。这是一个可滚动的 TUI 弹窗，会展示当前 MCP 服务器状态、技能目录信息和插件状态，让用户一目了然。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3643)

8.  **`docs: update ask-rule persistence docs` (#3642)**
    - **摘要：** 更新了 `ask-rule` 的文档，明确了 `S` 快捷键的持久化行为，并澄清了 `read_file` 规则无法通过该方式设置，必须手动写入。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3642)

9.  **`fix(tui): keep empty composer hint off cursor row` (#3635) & `[codex] keep empty composer hint off IME cursor row` (#3634)**
    - **摘要：** 修复了 #2612（CJK输入冲突）。这两个 PR 解决了在拼音/日文输入法组合状态下，占位符文本会遮挡 IME 输入候选框的问题。现在输入法有了干净的显示区域。用户 `nightt5879` 贡献了修复的核心代码。
    - [查看 PR #3634](https://github.com/Hmbown/CodeWhale/pull/3634) | [查看 PR #3635](https://github.com/Hmbown/CodeWhale/pull/3635)

10. **`fix(tui): expand failed tool output live` (#3639)**
    - **摘要：** 改进了失败工具调用的显示。之前失败的命令输出也会因为过长而被折叠，现在它们会被默认展开显示，同时保留成功命令的折叠行为，让调试信息更易获取。
    - [查看 PR](https://github.com/Hmbown/CodeWhale/pull/3639)

---

### 5. 功能需求趋势

从最近的议题和 PR 来看，CodeWhale 社区关注的趋势非常明确：

1.  **模块化与权限精细化：** 社区不再满足于全有或全无的权限控制。热点话题集中在如何通过 `permissions.toml` 或 `ask-rules` 实现**细粒度、持久化**的权限管理（`allow/deny/ask`），并且要求在 UI 中对规则进行**预览和确认**。
2.  **提示词系统开放化：** 用户希望 Codewhale 不只是软件工程的工具。`将主提示词外部化` 的提案表明，社区渴望能**自定义系统提示词**，以便将模型引导至文学创作、角色扮演等非编程领域。
3.  **跨语言与国际化：** 一方面是输入法问题的修复（#2612）和自然语言批准的桥接支持（#3637），另一方面是对本地化方案的重构（#3537）。这表明项目正在**认真打磨全球用户的体验**，尤其是东亚用户体验。
4.  **生命周期管理与发布自动化：** 大量 Issue 围绕**死代码清理**、**CI/CD 流程强化**（如 DCO 检查、资产匹配校验）以及**发布检查清单**展开。这反映了项目在快速迭代中，开始注重**代码质量和技术债的偿还**。
5.  **安全与审计：** 新增的 `approval.decided` 事件（#3647）和 DCO 贡献检查都指向了**安全性和可追溯性**的加强。对于需要集成到企业级工作流的工具，这是不可或缺的。

---

### 6. 开发者关注点

综合来看，现阶段的开发者关注点集中在以下几个方面：

- **模式混淆的挫败感：** `plan` 和 `agent` 模式反复出现问题（#3568），导致用户的认知负担很重。他们希望 AI 能**严格遵循当前模式**的设定，这是最核心的痛点。
- **CJK 用户的零散问题：** 尽管 IME 冲突已修复，但 CJK 用户仍可能在文本渲染、命令行参数等方面遇到问题。他们的反馈（如 `DracheTek` 用户）非常具体且具有建设性。
- **对“别人家孩子”的羡慕：** 很多优化（#2953, #2956, #2957）都明确对标 “Codex CLI”，这反映出开发者希望 CodeWhale 在**Token 消耗和速度**上达到行业标杆水平。
- **学习曲线陡峭：** 尽管有 TUI，但配置 MCP、技能等仍显复杂。Setup 向导（#3407）的出现就是对这种呼声的回应。
- **功能“不可见”：** 部分重要功能（如每周社区摘要、LSP 自定义服务器）缺乏发现入口。开发者希望 CodeWhale 能更好地**暴露**其内部能力，让用户“知其有，并用其好”。

---
以上就是今天的日报。项目正在为即将到来的 `v0.9.0` 系列进行积极的架构整理和功能迭代，社区贡献者也异常活跃。我们明天见！

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*