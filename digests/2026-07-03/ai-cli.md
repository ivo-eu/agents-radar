# AI CLI 工具社区动态日报 2026-07-03

> 生成时间: 2026-07-03 10:12 UTC | 覆盖工具: 9 个

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

好的，作为一名专注于 AI 开发工具生态的资深技术分析师，我将基于您提供的 2026-07-03 社区动态摘要，为您生成一份横向对比分析报告。

---

## AI CLI 工具生态横向对比分析报告 (2026-07-03)

### 1. 生态全景

当前 AI CLI 工具生态正处于从“功能可用”向“生产可靠”快速演进的关键阶段。各主流工具在保持高频迭代的同时，社区反馈的焦点已从“能否完成代码任务”转向“能否稳定、安全、高效地完成复杂开发工作流”。**跨平台兼容性、Agent 任务可靠性、API 成本控制和会话上下文管理**成为所有工具共同面临的挑战。与此同时，工具间的差异化定位也愈发清晰，生态正形成以通用任务、企业集成、开源灵活、模型原生为核心的四大阵营。

### 2. 各工具活跃度对比

| 工具名称 | 24h Issue 更新数 | 24h PR 活跃数 | 版本发布 | 社区讨论热度 (基于评论/点赞) |
| :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 30 | 3 (有效) | ✅ **v2.1.199** | 高 (SSL、认证、子代理注入等话题关注度高) |
| **OpenAI Codex** | 10+ (精选) | 10+ (合并为主) | ✅ **Rust Alpha x2** | **非常高** (Linux桌面#11023 获683👍，GPT-5.5 Token聚类热议) |
| **Gemini CLI** | 10+ (精选) | 10 | ✅ **Nightly** | **高** (子代理误报成功、通用Agent挂起等问题影响大) |
| **GitHub Copilot CLI**| 31 | 2 (无效/低质量) | 无 (最新为v1.0.69-0) | 中等 (alt-screen、BYOK、代理支持等讨论热烈) |
| **Kimi Code CLI** | 1 | 1 | 无 | **低** (仓库整体活跃度不高) |
| **OpenCode** | 9 | 10 | 无 | 中等 (Go Provider限速、数据库迁移等阻塞性问题多) |
| **Pi** | 7 | 10 | 无 | 高 (`content is not iterable` 崩溃是核心痛点，新供应商集成活跃) |
| **Qwen Code** | 6 | 50 | ✅ **v0.19.5 + Nightly** | 中等 (Web Shell UI增强、空参数工具调用bug等) |
| **DeepSeek TUI** | 10 | 10 | 无 (但含大批量修复PR) | **高** (v0.8.67设置流程、Fleet内存溢出、技能系统安全问题汇聚) |

### 3. 共同关注的功能方向

| 核心方向 | 涉及工具 | 具体诉求 |
| :--- | :--- | :--- |
| **跨平台桌面与终端体验** | **Claude Code**, **Codex**, **Gemini CLI**, **Copilot CLI**, **Qwen Code**, **DeepSeek TUI** | - **Linux/Wayland原生支持** (Codex #11023, Gemini #21983) <br> - **Windows剪贴板/终端兼容性** (Claude #69781, Copilot #3501, Kimi #2481, OpenCode #22003) <br> - **macOS功耗优化** (Codex 用户提及) |
| **会话与上下文管理** | **Claude Code**, **Codex**, **Copilot CLI**, **Pi**, **DeepSeek TUI** | - **Session历史丢失/Fork困难** (Claude #72335) <br> - **长期对话健壮性** (Codex #18993, Pi 上下文溢出) <br> - **Clear vs New 命令混淆** (Copilot #3569) |
| **API稳定性与成本控制** | **Claude Code**, **Codex**, **Gemini CLI**, **OpenCode** | - **高频超时导致Token浪费** (Claude #72260) <br> - **速率配额异常消耗** (Codex #30943) <br> - **Agent挂起消耗轮次** (Gemini #21409) <br> - **Provider限速误报** (OpenCode #34884) |
| **模型行为透明与兼容** | **Codex**, **Pi**, **Qwen Code**, **Copilot CLI** | - **推理Token聚类影响性能** (Codex #30364) <br> - **新模型兼容性不佳** (Pi #6278, Copilot #3997) <br> - **空参数工具调用导致重试** (Qwen #6249) |
| **安全与准入机制** | **Claude Code**, **Codex**, **Gemini CLI**, **DeepSeek TUI** | - **提示注入风险** (Claude #72332) <br> - **账号/OAuth认证障碍** (Claude #51588, Qwen #6251) <br> - **沙箱安全性** (Codex #15310, DeepSeek #3920) |
| **MCP生态深化** | **Gemini CLI**, **Copilot CLI**, **OpenCode**, **DeepSeek TUI** | - **MCP工具暴露给Agent失败** (OpenCode #33027) <br> - **MCP资源跨服务器混淆** (Gemini #28143) <br> - **MCP服务器动态启动** (DeepSeek #3866) |
| **Agent自治性与决策能力** | **Claude Code**, **Gemini CLI**, **OpenCode** | - **子代理误报任务状态** (Gemini #22323) <br> - **不主动调用用户定义的Skill/子代理** (Gemini #21968) <br> - **扩展能力变更后模型未感知** (Qwen #6244, OpenCode) |

### 4. 差异化定位分析

*   **Claude Code (通用务实的体验派)**：强调**开箱即用的开发体验**。版本发布稳健，社区反馈集中在**TUI交互细节**(粘贴、快捷键)、**安全策略**(误拦ffmpeg、提示注入)和**认证问题**上。它更像是一个力求稳定、满足主流开发者日常诉求的“水桶型”工具。

*   **OpenAI Codex (平台霸权与模型主导)**：**社区体量最大、生态诉求最广**。其动态反映出它不仅是CLI工具，更是OpenAI平台能力的延伸。社区高票需求(**Linux桌面、独立安装包**)和围绕**GPT-5.5模型行为**的深度讨论，显示用户对其高频依赖，并期望它成为跨平台、跨生态的中心节点。**凭据路由、MITM代理等企业级基础设施的持续合入**，也凸显其向企业服务渗透的野心。

*   **Gemini CLI (Agent架构探索先锋)**：**聚焦Agent的复杂行为与可靠性**。社区讨论集中在**子代理状态误报、Agent挂起、Skill/子代理调用逻辑**等高级话题。其背后是谷歌对“多Agent协作”和“计划-执行”范式的深度探索。虽然用户对稳定性仍有微词，但这种架构的前瞻性也吸引了一批愿意尝鲜、关注复杂任务自动化的高阶开发者。

*   **GitHub Copilot CLI (深度绑定GitHub生态)**：定位是**开发者的“副驾驶”连接器**，而非独立平台。社区需求围绕**终端TUI自定义、企业代理支持、BYOK模型端点**，体现了其用户群（高度依赖GitHub生态的开发者）的明确边界。**社区PR贡献活跃度低**，表明其开源模式更像是托管仓库，而非共建社区。

*   **OpenCode / Pi / DeepSeek TUI / Qwen Code (开源创新与垂直深耕)**：这些工具呈现出极强的**技术灵活性和社区驱动**。
    *   **Pi**: 以“供应商扩展”见长，快速集成DeepInfra、GLM等新选择，解决兼容性**Bug**并增强功能。
    *   **OpenCode**: 重心放在 **MCP集成、code-mode实验性功能**和**性能/Stability修复**上，代表了向全功能本地开发平台演进的方向。
    *   **Qwen Code**: 围绕**多模态(CUA驱动)、Web Shell 和其扩展生态**做文章，并积极适配OpenAI兼容接口。
    *   **DeepSeek TUI**: 激进迭代 **V2架构**，一次性解决大量设置流程、技能系统和Fleet模块的**审计问题**，体现了对自身基础能力的大胆重构。

### 5. 社区热度与成熟度

*   **高活跃/较成熟**: **OpenAI Codex** 和 **Claude Code**。这两个工具拥有庞大的用户群，Issue数量和讨论深度均居前列。它们已度过早期功能开发期，进入了围绕稳定性、安全性和企业级功能进行“打磨”的阶段。Codex社区更侧重于**功能请求**和**模型深度讨论**；Claude Code社区则更聚焦于**具体Bug修复**和**用户体验优化**。

*   **高活跃/快速迭代**: **Gemini CLI**。社区活跃度很高，但问题类型多为架构级的Agent行为异常和功能缺失，表明其核心架构仍在快速演进中。用户既是使用者也充当着测试者，对创新功能接受度高，但同时对稳定性也有明确期待。

*   **中活跃/生态特色**: **Pi**, **OpenCode**, **DeepSeek TUI**。这些工具处于快速成长期，社区围绕其特色功能（Pi的供应商生态、OpenCode的MCP、DeepSeek的Agent架构）积极贡献，同时也揭示了大量早期使用者才能发现的深层次Bug和体验割裂。

*   **低活跃/相对冷清**: **Kimi Code CLI 和 GitHub Copilot CLI**。Kimi的活跃度极低，可能处于项目早期或资源投入有限阶段。Copilot CLI虽然Issue量大，但多为**功能请求**，且**核心PR贡献极少**，社区更像一个“用户反馈中心”，而非共建生态平台。

### 6. 值得关注的趋势信号

1.  **“空内容”崩溃成为高频痛点**：Pi、Qwen Code 等多个工具集中出现因模型返回 `null content` 或空参数导致的无限重试或无响应错误。这强烈提示：**当工具调用（Tool Use）与流式推理结合时，模型输出的不稳定性和极端情况处理，是当前最急需解决的基础设施问题。**

2.  **Sandbox与安全模型的“信任危机”**: Claude Code的安全过滤器误拦、Codex的沙箱静默回退、DeepSeek TUI的“虚假”安全声明，都指向一个问题：**当前的安全策略过于粗糙，无法在“保护”与“赋能”间取得平衡**。未来的安全模型必须是可配置、可审查、且行为透明的。用户希望自己控制边界，而非被武断阻止。

3.  **“多模型/多Provider”成为标配而非特性**：Pi快速集成 DeepInfra、GLM，Copilot CLI 和 Gemini CLI 对 BYOK 和自定义端点的强烈需求，以及 Qwen Code 对 OpenAI 兼容性的兜底修复，都表明**开发者不愿被锁定在单一模型生态中**。 “模型路由器”式的、能动态选择不同提供商和模型的CLI将是主流形态。

4.  **TUI交互从“可以用”到“必须好用”**: Copilot CLI 的 alt-screen 讨论、Claude Code的粘贴问题、OpenCode的会话多选删除，表明**终端用户界面（TUI）的交互逻辑正成为影响开发者满意度的决定性因素**。仅仅显示输出已不够，高效的键盘导航、自然的粘贴体验、符合直觉的会话管理，都是决定用户留存的关键。

5.  **Agent自主性提升带来“撒谎”难题**：Gemini CLI 的子代理误报“成功”以及 Claude Code 的工具调用静默失败，揭示了当前 **Agent 架构在状态管理和错误报告上的严重缺陷**。随着Agent自主执行任务的比例增高，如何确保其“诚实”地报告执行结果，而非通过“粉饰太平”来结束轮次，将成为Agent可用性的核心挑战。

6.  **MCP协议从“连通”走向“治理”**：多个工具的MCP问题不再仅仅是“连不上”，而是上升到**资源混淆、跨服务器错误、动态启动**等治理层面。这标志着MCP生态正进入深水区，**对元数据管理、命名空间隔离和权限控制**的需求日益凸显。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，以下是根据您提供的 `anthropics/skills` 仓库数据（截至 2026-07-03）生成的社区热点分析报告。

---

## Claude Code Skills 社区热点报告（截至 2026-07-03）

### 1. 热门 Skills 排行

以下是根据社区关注度（评论/关联讨论）选出的 5~8 个热门 Pull Requests，它们反映了社区最关心的 Skill 开发和修复方向。

-   **#1298: fix(skill-creator): run_eval.py always reports 0% recall**
    -   **功能**：核心修复。解决 `run_loop.py`、`improve_description.py` 等工具链中 `run_eval.py` 始终报告 `recall=0%` 的根本性问题。该问题导致描述优化循环完全失效，社区已有 10+ 次独立复现。修复涉及安装逻辑、Windows 兼容性、触发检测及并行处理等多个方面。
    -   **讨论热点**：社区对该 `skill-creator` 核心工具链的稳定性表现出极高关注。此 PR 被视为解决描述优化流程“基建”问题的关键，PR 本身也是多个相关修复的集大成者。
    -   **状态**：**Open** — 评论活跃，合并优先级极高。

    [GitHub PR #1298](https://github.com/anthropics/skills/pull/1298)

-   **#514: Add document-typography skill**
    -   **功能**：新增 `document-typography` Skill，专门用于修复 AI 生成文档中的排版问题，如孤行（orphan）、寡段（widow）和编号错位等。
    -   **讨论热点**：社区对 AI 生成文档的“细粒度”质量把控有强烈需求。所有 AI 生成用户都会面临此类问题，该 Skill 的实用性获得广泛认可。
    -   **状态**：**Open** — 讨论持续，等待合并。

    [GitHub PR #514](https://github.com/anthropics/skills/pull/514)

-   **#538: fix(pdf): correct case-sensitive file references in SKILL.md**
    -   **功能**：修复 `SKILL.md` 文件中 8 处大小写不一致的文件引用（如 `REFERENCE.md` → `reference.md`），解决了在大小写敏感系统（如 Linux/macOS）上 PDF Skill 无法正常工作的问题。
    -   **讨论热点**：这是一个典型的“跨平台兼容性”问题。虽然修复简单，但暴露了社区在维护高质量、跨平台 Skills 时遇到的普遍挑战。
    -   **状态**：**Open** — 讨论相对平稳，属于高质量、低争议的修复。

    [GitHub PR #538](https://github.com/anthropics/skills/pull/538)

-   **#486: Add ODT skill**
    -   **功能**：新增 `ODT` Skill，支持创建、填充、读取和转换 OpenDocument 格式文件（`.odt`, `.ods`）。覆盖了 LibreOffice 用户和需要开放标准文档的群体。
    -   **讨论热点**：社区对主流闭源格式之外的开源办公格式支持有明确期待。该 Skill 填补了生态缺口，回应了特定用户群体的长期呼吁。
    -   **状态**：**Open**— 讨论持续，是社区急需的实用型 Skill。

    [GitHub PR #486](https://github.com/anthropics/skills/pull/486)

-   **#210: Improve frontend-design skill clarity and actionability**
    -   **功能**：对现有 `frontend-design` Skill 进行全面修订，提升其指令的清晰性、可操作性和内部一致性，确保 Claude 能在一个对话中精确执行每一个指引。
    -   **讨论热点**：核心聚焦于“Skill 质量”。社区关注的重点已从“有没有该 Skill”转向“这个 Skill 好不好用”。高质量的 Skill 描述和引导才能发挥 Claude 的最佳能力。
    -   **状态**：**Open** — 社区持续关注已有 Skills 的迭代优化。

    [GitHub PR #210](https://github.com/anthropics/skills/pull/210)

-   **#83: Add skill-quality-analyzer and skill-security-analyzer**
    -   **功能**：新增两个“元技能”：
        - **skill-quality-analyzer**：从结构与文档、有效性、健壮性、安全性、可维护性五个维度评估 Skill 质量。
        - **skill-security-analyzer**：分析 Skill 潜在的安全风险。
    -   **讨论热点**：社区对创建高质量、高安全性技能的自检需求。这两个元技能的出现，标志着社区开始从“数量增长”转向“质量管理”。
    -   **状态**：**Open** — 对塑造健康社区生态有重要价值。

    [GitHub PR #83](https://github.com/anthropics/skills/pull/83)

-   **#1367: feat(skills): add self-audit**
    -   **功能**：新增 `self-audit` Skill，在 AI 输出交付前执行两步审核：1）机械文件级验证，2）按损害严重性优先级进行四维推理审计。
    -   **讨论热点**：社区对于 AI 生成结果的“最终质量把关”有强烈需求。这个 Skill 类似一个通用的“交付质检员”，可以显著提升 AI 生成物的可靠度。
    -   **状态**：**Open** — 非常新的 PR，但已展现高潜力。

    [GitHub PR #1367](https://github.com/anthropics/skills/pull/1367)

-   **#723: Add testing-patterns skill**
    -   **功能**：新增 `testing-patterns` Skill，覆盖测试哲学（Testing Trophy模型）、单元测试（AAA模式）、React组件测试（Testing Library）以及端到端测试等全面测试体系。
    -   **讨论热点**：社区对高质量、结构化的测试方案有巨大需求。该 Skill 试图标准化和自动化 Claude 在项目中的测试行为。
    -   **状态**：**Open** — 评论活跃，是技术类 Skill 的热门候选。

    [GitHub PR #723](https://github.com/anthropics/skills/pull/723)

### 2. 社区需求趋势

从 Issues 中可以提炼出以下社区最期待的新 Skill 方向：

-   **安全性 & 信任构建**：Issue **#492**（34条评论）引发了关于`anthropic/`命名空间下社区 Skill 信任问题的热烈讨论。社区强烈呼吁建立更清晰的来源标识和安全审查机制，防止信任边界滥用。这预示着一个关于“Skill 安全评估”或“官方认证标签”的解决方案将很快出现。

-   **协作与共享生态**：Issue **#228**（14条评论，7个👍）提出了在组织内直接分享 Skill 的需求，以替代目前“下载-发送-手动上传”的低效流程。这反映出社区正在从个人使用走向团队协作，对 Skill 的“流通效率”提出了更高要求。

-   **平台兼容性（尤其 Windows）**：Issue **#1061**（3条评论）和 Issue **#556**（12条评论）共同指向了 Windows 平台上的严重兼容性问题，特别是 `skill-creator` 工具链（如 `run_eval.py`）完全不可用。这是当前阻碍社区参与贡献的“硬伤”，解决该问题被认为是社区发展的优先事项。

-   **Agent 治理与效率**：Issues **#412**（Agent Governance）、**#1329**（compact-memory）的提出，表明社区对构建更复杂、更可控的 Agent 行为产生兴趣。需求集中在：
    - **Agent 治理**：为 Agent 系统建立策略执行、威胁检测和审计跟踪等安全模式。
    - **智能状态压缩**：用符号化表示法管理长时运行 Agent 的上下文，以减少 token 消耗和提高效率。

-   **开发者工具链稳定性**：Issues **#556**、**#202**、**#1169** 反复指向 `skill-creator` 工具链的可靠性问题（如始终报告0%召回率、最佳实践缺失等）。这表明社区对用于**开发 Skill 的元工具**的稳定性要求极高，其质量直接影响了整个生态的活力。

### 3. 高潜力待合并 Skills

以下 PR 评论活跃，功能实用，是短期内最有可能落地、为社区带来显著价值的高潜力候选：

-   **#1298: fix(skill-creator): run_eval.py always reports 0% recall** — 这是当前生态的“基建级”修复，直接关系到 `skill-creator` 工具链是否可用，合并优先级最高。
    [GitHub PR #1298](https://github.com/anthropics/skills/pull/1298)

-   **#1367: feat(skills): add self-audit** — 作为通用的“质量门禁”，能显著提升所有 AI 交付物的可靠性，满足了社区对“最终检查”的普遍需求。
    [GitHub PR #1367](https://github.com/anthropics/skills/pull/1367)

-   **#1302: Add color-expert skill** — 一个高度垂直、专业性极强的 Skill，填补了颜色科学领域的空白，适合设计师和相关专业人士。
    [GitHub PR #1302](https://github.com/anthropics/skills/pull/1302)

-   **#723: Add testing-patterns skill** — 作为开发流程的核心环节，结构化的测试 Skill 对提升代码质量和开发效率至关重要，是开发者社区的强需求。
    [GitHub PR #723](https://github.com/anthropics/skills/pull/723)

-   **#514: Add document-typography skill** — 解决了 AI 文档生成“最后一公里”的常见痛点和用户体验问题，实用性极高。
    [GitHub PR #514](https://github.com/anthropics/skills/pull/514)

-   **#806: feat: add sensory skill** — 为 macOS 用户提供了更原生的自动化方案，作为“计算机使用”功能的补充，有明确的用户场景和差异化价值。
    [GitHub PR #806](https://github.com/anthropics/skills/pull/806)

### 4. Skills 生态洞察

**当前社区最集中的诉求是解决 `skill-creator` 工具链的“系统级”稳定性问题和社区技能生态的“治理与信任”机制缺失。** 这意味着社区的核心矛盾已从“如何创造更多技能”转向“如何让技能创作过程可靠、结果可信，并在团队/组织中高效流通”。

---

好的，这是 2026 年 7 月 3 日的 Claude Code 社区动态日报。

---

# Claude Code 社区动态日报 | 2026-07-03

## 今日速览

- **新版发布**：发布 v2.1.199，修复了 SSL 代理环境下的 TLS 证书报错，并优化了堆叠 Slash 命令的执行逻辑。
- **社区活跃度回暖**：过去 24 小时内共有 30 个 Issue 获得更新，主要集中在前端交互（TUI）、API 稳定性及平台兼容性（WSL2/macOS）方面。
- **安全与准入问题突出**：多个 Issue 社区关注度高，涉及账号认证阻塞、子代理提示注入、以及 SSH 会话历史丢失等影响开发体验的严重 bug。

## 版本发布

### v2.1.199
- **发布说明**：这是一个针对开发体验的修复版本。
- **主要更新**：
    - **Slash 命令增强**：修复了堆叠调用（如 `/skill-a /skill-b do XYZ`）时，只会执行第一个技能的 bug，现在最多可同时加载前 5 个技能。
    - **SSL/TLS 错误处理优化**：修复了在 TLS 监控代理环境、`NODE_EXTRA_CA_CERTS` 缺失或证书过期时，系统会不断重试而没有任何有用提示的问题。现在会直接展示可操作的修复指导，大幅节省了开发者排查问题的时间。

## 社区热点 Issues (Top 10)

1. **[BUG] 账号恢复后认证系统仍封锁登陆 (#51588)**
    - **作者**: chrs-myrs | **评论**: 7 | **👍**: 9
    - **重要性**: 严重影响已解决问题的用户回归，属于严重工程缺陷。
    - **社区反应**: 超过 9 人点赞，表明该问题在 Trust & Safety 流程中较普遍。
    - **链接**: [Issue #51588](https://github.com/anthropics/claude-code/issues/51588)

2. **[BUG] macOS 粘贴图片报错 UTF-8 代理字符不合法 (#69781)**
    - **作者**: GuiAmaral85 | **评论**: 4 | **👍**: 1
    - **重要性**: 直接导致 macOS 用户在桌面端粘贴图片失败，影响日常使用。
    - **链接**: [Issue #69781](https://github.com/anthropics/claude-code/issues/69781)

3. **[BUG] 安全过滤器误拦截正常的 ffmpeg 操作 (#72355)**
    - **作者**: sworrl | **评论**: 4 | **👍**: 0
    - **重要性**: 尽管标记为重复，但揭示了安全过滤策略过于敏感，可能阻塞用户正常授权的工作流。
    - **链接**: [Issue #72355](https://github.com/anthropics/claude-code/issues/72355)

4. **[BUG] 工具调用标记偶尔被当作纯文本输出，命令静默不运行 (#72240)**
    - **作者**: monkeychen | **评论**: 3 | **👍**: 0
    - **重要性**: 在长时间会话中，工具调用失败可能导致开发任务静默中断。
    - **链接**: [Issue #72240](https://github.com/anthropics/claude-code/issues/72240)

5. **[BUG] 在 Vertex AI 上运行 Opus 4.8 时，showing 摘要模式返回空 thinking 块 (#72231)**
    - **作者**: BilalAtique | **评论**: 2 | **👍**: 0
    - **重要性**: 影响部署在 Vertex AI 上并依赖思考链摘要显示的开发者。
    - **链接**: [Issue #72231](https://github.com/anthropics/claude-code/issues/72231)

6. **[BUG] WSL2 (WezTerm) 中粘贴功能彻底失效 (#72235)**
    - **作者**: TsLouis | **评论**: 2 | **👍**: 0
    - **重要性**: WSL2 上的核心交互中断，严重影响该平台上用户的开发效率。
    - **链接**: [Issue #72235](https://github.com/anthropics/claude-code/issues/72235)

7. **[BUG] 约 50% 的 API 请求超时导致 Token 浪费 (#72260)**
    - **作者**: mirkokozmosz | **评论**: 2 | **👍**: 0
    - **重要性**: 严重浪费用户配额，尤其在小请求上超时却附带大量缓存上下文，成本急剧升高。
    - **链接**: [Issue #72260](https://github.com/anthropics/claude-code/issues/72260)

8. **[BUG] 子代理输出中存在提示注入，尝试写入伪造的安全发现 (#72332)**
    - **作者**: Mmckelve45 | **评论**: 2 | **👍**: 1
    - **重要性**: 潜在的安全风险，子代理被劫持可能影响终端用户决策。
    - **链接**: [Issue #72332](https://github.com/anthropics/claude-code/issues/72332)

9. **[BUG] Linux VPS 上 OAuth 登录失败，重定向 URI 不受支持 (#72344)**
    - **作者**: comercialvectahub-cmd | **评论**: 2 | **👍**: 1
    - **重要性**: 影响无头 Linux 服务器用户（如 VPS/CI/CD）的正常使用，阻碍远程开发。
    - **链接**: [Issue #72344](https://github.com/anthropics/claude-code/issues/72344)

10. **[BUG] SSH 会话历史空白，Fork 功能因直接连接原始 IP 而失败 (#72335)**
    - **作者**: mannmann2 | **评论**: 2 | **👍**: 0
    - **重要性**: 破坏 SSH 远程开发的核心会话管理功能，导致工作流中断。
    - **链接**: [Issue #72335](https://github.com/anthropics/claude-code/issues/72335)

## 重要 PR 进展 (Top 6)

1. **[CLOSED] 修复 devcontainer 防火墙脚本因重复 IP 而崩溃 (#42701)**
    - **作者**: michaelkonecny
    - **重要性**: 解决了 Devcontainer 启动时的硬错误，提升了开发环境搭建的健壮性。
    - **链接**: [PR #42701](https://github.com/anthropics/claude-code/pull/42701)

2. **[OPEN] 从防火墙白名单中移除已失效的 Hostname (#72451)**
    - **作者**: gmli-eu
    - **重要性**: 解决 Devcontainer 因 Hostname 无法解析导致启动失败的问题，属于基础设施修复。
    - **链接**: [PR #72451](https://github.com/anthropics/claude-code/pull/72451)

3. **[OPEN] 修复 README 中 GitHub 拼写错误 (#73476)**
    - **作者**: Elmahditcham
    - **重要性**: 文档维护，表明社区对项目文档质量的关注。
    - **链接**: [PR #73476](https://github.com/anthropics/claude-code/pull/73476)

*(注：其余 3 个 PR 为机器人生成或重复提交，在此略过)*

## 功能需求趋势

- **跨平台兼容性**：大量 Issue 集中在 macOS、WSL2、Linux 等平台上的交互问题，表明开发者希望在所有主流开发环境下获得一致、稳定的体验。
- **会话与上下文管理**：多个 Issue 建议增强 Fork 锚点、回顾历史输入以及跨摘要保留上下文的能力，说明社区对长时间会话的健壮性和可交互性有更高期望。
- **API 稳定性与成本控制**：频繁的超时和令牌浪费是突出痛点，开发者希望更高效的重试和上下文传输机制。
- **安全与准入**：账号登录、OAuth 认证、以及防止提示注入是持续热点，社区期待更完善的认证和安全护栏机制。

## 开发者关注点

- **TUI/交互 Bug 是最大痛点**：无论是快捷键失灵、粘贴失败还是 UI 元素点击异常，直接干扰了开发者的核心工作流。
- **API 可靠性成隐形成本**：部分开发者反馈超时导致配额浪费，表明除功能外，稳定性已成为评判工具效率的关键指标。
- **认证与登录不是小事**：从 VPS OAuth 失败到账号恢复后仍被封锁，这些“入门级”问题会严重阻碍潜在新用户的转化。
- **会话隔离有待加强**：后台任务输出错乱到 SSH 会话历史丢失，社区迫切需要更精准的会话管理和数据持久化机制。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 | 2026-07-03

---

## 今日速览

过去 24 小时内，Codex 发布了两个 Rust 工具的 alpha 版本，但社区焦点更多地集中在 **Linux 桌面应用** 和 **Windows 独立安装程序** 两大长期需求上（合计点赞超 840）。同时，**GPT-5.5 推理 token 聚类** 和 **速率配额异常消耗** 等严重 bug 引发广泛讨论。PR 方面，OpenAI 团队密集合并了关于 **凭据路由代理**、**命令行批准机制** 以及 **Guardian 策略更新** 的多项关键改进。

---

## 版本发布

- **rust-v0.143.0-alpha.35** | 2026-07-03  
  `0.143.0-alpha.35` 发布  
  [查看详情](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.35)

- **rust-v0.143.0-alpha.34** | 2026-07-03  
  `0.143.0-alpha.34` 发布  
  [查看详情](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.34)

> 两个 alpha 版本均无额外变更说明，推测为持续集成或依赖更新。

---

## 社区热点 Issues（精选 10 条）

### 1. Linux 桌面应用支持 — [#11023](https://github.com/openai/codex/issues/11023)
- **状态**: OPEN | 创建于 2026-02-07 | **140 条评论** | **683 👍**
- **重要性**: 社区呼声最高的功能请求。由于 macOS 上 #10432 导致功耗问题，大量用户亟需原生 Linux 桌面客户端。评论中用户积极讨论跨平台方案（Electron、Tauri）及性能对比。

### 2. 支持 Windows 独立安装程序 — [#13993](https://github.com/openai/codex/issues/13993)
- **状态**: OPEN | 创建于 2026-03-08 | **78 条评论** | **161 👍**
- **重要性**: 企业用户因系统策略或离线环境无法使用 Microsoft Store，需求强烈。官方已标记 `User Request` 和 `Feature`，但进展缓慢。

### 3. `X-OpenAI-Internal-Codex-Responses-Lite` 报错 — [#30224](https://github.com/openai/codex/issues/30224)
- **状态**: OPEN | 创建 2026-06-26 | **67 条评论** | **22 👍**
- **重要性**: 影响 Plus 订阅用户使用特定 API 头时触发“模型不支持”错误，疑似后端路由配置问题，被大星标记为 `bug` + `custom-model`。

### 4. GPT-5.5 推理 token 聚类导致性能下降 — [#30364](https://github.com/openai/codex/issues/30364)
- **状态**: OPEN | 创建 2026-06-27 | **34 条评论** | **49 👍**
- **重要性**: 用户研究发现 GPT-5.5 的 `reasoning_output_tokens` 异常固定于 516/1034/1552 三个数值，暗示模型存在 token 截断或量化舍入，可能影响复杂任务推理质量，社区反响热烈。

### 5. VS Code 扩展无法打开历史会话 — [#18993](https://github.com/openai/codex/issues/18993)
- **状态**: CLOSED | 创建 2026-04-22 | **39 条评论** | **53 👍**
- **重要性**: 影响 Plus 用户的日常使用，虽已关闭但表明会话持久化逻辑仍存在间歇性 regression。开发者在关闭前确认了临时修复，但长期方案仍在规划。

### 6. 桌面自动化静默回退至 workspace-write 沙箱 — [#15310](https://github.com/openai/codex/issues/15310)
- **状态**: OPEN | 创建 2026-03-20 | **17 条评论** | **14 👍**
- **重要性**: 即使配置了 `danger-full-access`，定时任务仍使用受限沙箱，直到用户手动进入聊天 UI 才修正。严重影响自动化工作流的安全性预期。

### 7. Windows 沙箱回归 (CLI 0.136.0) — [#26158](https://github.com/openai/codex/issues/26158)
- **状态**: CLOSED | 创建 2026-06-03 | **16 条评论** | **6 👍**
- **重要性**: 用户被迫回退至 0.132.0 以恢复 Windows 沙箱功能，问题涉及 `CreateProcessAsUserW` 失败。已关闭但未提及根因永久修复，建议升级用户留意。

### 8. 项目级插件市场与仓库级配置 — [#18115](https://github.com/openai/codex/issues/18115)
- **状态**: OPEN | 创建 2026-04-16 | **9 条评论** | **45 👍**
- **重要性**: 社区希望 `.codex/config.toml` 支持本地插件市场，使团队协作时共享 skill 配置，属于企业级需求。

### 9. Codex App 持续崩溃 (Windows 11) — [#30531](https://github.com/openai/codex/issues/30531)
- **状态**: OPEN | 创建 2026-06-29 | **4 条评论** | **0 👍**
- **重要性**: 用户反馈 App 频繁崩溃，无法完成工作。虽评论数少，但结合 #30824（macOS 崩溃）和 #30624（冻结），表明桌面应用稳定性仍是普遍痛点。

### 10. 速率配额异常耗尽（无活动） — [#30943](https://github.com/openai/codex/issues/30943)
- **状态**: OPEN | 创建 2026-07-02 | **3 条评论** | **0 👍**
- **重要性**: 用户证明 usage 显示 0% 但实际未发送任何消息，明确指出是后台配额记账 bug。鉴于类似投诉 (#30641) 持续出现，可能影响大量 Plus/Pro 用户权益。

---

## 重要 PR 进展（精选 10 条）

### 1. 传播明确的命令批准目的 — [#30976](https://github.com/openai/codex/pull/30976)
- **状态**: OPEN | 创建 2026-07-03
- **摘要**: 为命令批准回调添加 `purpose` 信息，防止 `execve` 拦截与沙箱重试混淆。修正生命周期转换错误。

### 2. 保留命令身份跨重复批准 — [#30969](https://github.com/openai/codex/pull/30969)
- **状态**: OPEN | 创建 2026-07-03
- **摘要**: 修复多次批准回调覆盖父命令元数据的问题，确保命令身份持久化。

### 3. 更新 Guardian 策略措辞 — [#29038](https://github.com/openai/codex/pull/29038)
- **状态**: MERGED | 创建 2026-06-19
- **摘要**: 阐明敏感数据移动、用户授权边界，保留受限本地、只读和仓库作用域操作。已合入主线。

### 4. 收紧实时确认指导 — [#29028](https://github.com/openai/codex/pull/29028)
- **状态**: MERGED | 创建 2026-06-19
- **摘要**: 减少实时前端提示中重复的开场白，鼓励简洁过渡。含默认提示的新增验收测试。

### 5. 避免重复的 ImageGen Markdown 输出 — [#28996](https://github.com/openai/codex/pull/28996)
- **状态**: MERGED | 创建 2026-06-18
- **摘要**: 修正图片生成时同时渲染 Markdown 和内嵌 nugget，导致三份相同结果的问题。

### 6. 凭据路由发现与 MITM 代理集成（系列 PR）
- **[#22675](https://github.com/openai/codex/pull/22675)** — 通过 MITM 代理路由凭据流量  
- **[#22673](https://github.com/openai/codex/pull/22673)** — 发现用于托管代理的凭据路由  
- **[#27503](https://github.com/openai/codex/pull/27503)** — 会话期间刷新凭据路由  
- **[#28981](https://github.com/openai/codex/pull/28981)** — 通过配置重载器重基实时代理状态  
- **[#28984](https://github.com/openai/codex/pull/28984)** — 添加凭据路由后端适配器  
- 全部已合并，共同构建了托管代理对凭据 HTTPS 端点的安全转发能力。

### 7. 可配置的 ChatGPT OAuth — [#28998](https://github.com/openai/codex/pull/28998)
- **状态**: MERGED | 创建 2026-06-18
- **摘要**: 允许通过配置显式指定 OAuth 端点，生产环境默认保留原值。支持浏览器/设备码/刷新等所有流程，并忽略项目级覆盖。

### 8. 跨进程云配置包缓存协调 — [#27971](https://github.com/openai/codex/pull/27971)
- **状态**: MERGED | 创建 2026-06-12
- **摘要**: 解决多个 Codex 进程共享 `CODEX_HOME` 时并发获取相同云配置包的问题，通过文件锁协调刷新归属，减少冗余网络请求。

### 9. 文件系统符号链接语义 — [#28930](https://github.com/openai/codex/pull/28930)
- **状态**: MERGED | 创建 2026-06-18
- **摘要**: 为 `fs/getMetadata` 添加 `followSymlinks` 参数，区分符号链接与其目标，提升文件系统操作准确性。

### 10. 安全 shell 命令分类需字面单词 — [#28961](https://github.com/openai/codex/pull/28961)
- **状态**: MERGED | 创建 2026-06-18
- **摘要**: 要求未加引号的 shell 单词保持字面才允许使用解析参数进行安全分类，避免引号内通配符被错误展开。

---

## 功能需求趋势

从过去 24 小时更新的 Issues 和 PRs 中，社区最关注的三大方向：

1. **跨平台桌面应用**  
   - **Linux 原生支持** (#11023，683 👍) 和 **Windows 独立安装程序** (#13993，161 👍) 是持续数月的最强需求。
   - 用户对 Electron/Tauri 讨论激烈，期望解决 macOS 功耗问题并适配不同操作系统限制。

2. **模型行为透明与可靠性**  
   - **GPT-5.5 token 聚类** (#30364) 暴露了模型输出的潜在硬编码问题，引发对推理质量下降的担忧。
   - 多个报告指出 **速率配额记账错误**（#30943, #30641），用户对透明度与公平性高度敏感。

3. **沙箱与安全机制增强**  
   - **桌面自动化沙箱回退** (#15310)、**Windows 沙箱 regression** (#26158) 表明沙箱实现仍不稳定。
   - 社区呼吁更细粒度的 **项目级插件/技能配置** (#18115) 和 **主题化记忆存储** (#19758)，以支持复杂工作流。

---

## 开发者关注点

- **桌面应用稳定性**：多平台崩溃报告（#30824 macOS、#30531 Windows）影响生产力，开发者期望更严格的内存与文件系统测试。
- **会话与历史管理**：VS Code 扩展无法打开历史会话 (#18993)、Windows App 线程乱序 (#29561)、无限滚动循环 (#28755) 等 bug 表明前端状态管理仍需打磨。
- **CLI/TUI 体验**：`Ctrl-Z` 后终端不可用 (#31013)、消息着色错误 (#30973)、`apply_patch` 上下文忽略 (#30946) 等细节问题影响 CLI 日常使用。
- **凭据与代理**：大量 PR 围绕凭据路由和 MITM 代理，反映出用户对内部/第三方 API 集成安全性的关注度提升。
- **升级兼容性**：社区对 alpha 版本变动保持警惕，如 #26158 用户明确选择停留在 0.132.0，提醒团队在 RC 阶段加强 regression 测试。

---

*数据更新时间：2026-07-03 14:00 UTC*  
*来源：GitHub – openai/codex*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，这是为您生成的 2026-07-03 Gemini CLI 社区动态日报。

---

# Gemini CLI 社区动态日报 | 2026-07-03

## 今日速览
今日社区动态聚焦于 **Agent 可靠性**与 **核心交互稳定性**。一个关于子代理在到达对话轮次上限后误报“成功”状态的 Bug（#22323）引发了广泛讨论，凸显了Agent状态报告的重要性。与此同时，针对 VSCode 插件焦点丢失、Shell 命令执行卡顿及终端响应问题的多个修复 PR 正在推进，开发者体验的打磨依然是社区关注的焦点。

## 版本发布
- **v0.51.0-nightly.20260703.gf7af4e518**: 发布了一个新的 Nightly 版本。核心变更为为 `caretaker` 功能新增了 Egress Cloud Run 服务的基础框架（skeleton）。【GitHub链接】

## 社区热点 Issues
1.  **[#22323] Subagent recovery after MAX_TURNS is reported as GOAL success** (P1/Bug)
    - **重要性**: 这是一个严重的逻辑 Bug。`codebase_investigator` 子代理在因达到最大对话轮次限制（`MAX_TURNS`）而中断时，其 `Termination Reason` 却报告为 “GOAL”，暗示任务成功完成。这会导致用户和主代理产生误解，掩盖了任务被强制中断的真相，并可能引发后续的决策错误。社区对此高度关注，评论数达到9条。【GitHub链接】

2.  **[#21409] Generalist agent hangs** (P1/Bug)
    - **重要性**: 用户报告通用代理在执行简单任务（如创建文件夹）时永久挂起，此问题已存在数月，且获得了8个👍，表明影响范围较广。问题的根本原因可能与子代理的调度机制有关，用户发现通过指令禁止代理启用子代理可以规避此问题。【GitHub链接】

3.  **[#25166] Shell command execution gets stuck with "Waiting input"** (P1/Bug)
    - **重要性**: 一个影响日常使用的高频 Bug。在 CLI 执行完简单的 Shell 命令后，Gemini CLI 会错误地显示“等待用户输入”并卡死，即使命令本身已经完成。此问题严重破坏了交互的流畅性。【GitHub链接】

4.  **[#19873] Leverage model's bash affinity via Zero-Dependency OS Sandboxing** (P2/Enhancement)
    - **重要性**: 这是一个更长期的架构性提案。旨在利用 Gemini 3 模型“原生”的 Bash 操作能力，通过零依赖的沙箱机制来执行命令，从而提升安全性并减少对自定义工具的依赖。该讨论仍在进行中，体现了社区对更安全、更高效的底层执行方式的追求。【GitHub链接】

5.  **[#21968] Gemini does not use skills and sub-agents enough** (P2/Bug)
    - **重要性**: 用户反馈 Gemini CLI 不会主动调用用户自定义的 Skills 和子代理，即使当前任务与 Skill 的描述高度相关。这暴露了代理在自主决策和任务规划方面的不足，影响了功能的实际效用和扩展性。【GitHub链接】

6.  **[#22745] Assess the impact of AST-aware file reads, search, and mapping** (P2/Feature)
    - **重要性**: 社区正在探究通过 AST（抽象语法树）感知的文件读取和搜索，来提升代码理解的精确度和效率。如果实现，有望显著减少因文件读取错位而导致的无效交互轮次，并降低 Token 消耗。【GitHub链接】

7.  **[#26525] Add deterministic redaction and reduce Auto Memory logging** (P2/Bug/Security)
    - **重要性**: 该问题聚焦于 Auto Memory 功能的隐私和安全。目前，提取提示要求模型进行脱敏，但在这个过程中，内容已经发送给模型，存在信息泄露风险。同时，系统日志可能记录 Skill 信息。社区呼吁提供确定性的、非模型参与的脱敏机制。【GitHub链接】

8.  **[#21983] browser subagent fails in wayland** (P1/Bug)
    - **重要性**: 在 Linux Wayland 显示服务器环境下，Browser 子代理无法正常工作。这限制了该功能在特定平台上的可用性，对部分开发者而言是一个关键的兼容性障碍。【GitHub链接】

9.  **[#24353] Robust component level evalutions** (P1/Bug/EPIC)
    - **重要性**: 这是一个史诗级（EPIC）问题。社区正在推动建立更健壮的组件级评估体系，以确保代码改动不会破坏核心功能，并提升测试的覆盖率和可靠性。【GitHub链接】

10. **[#21000] Experiment with using native file tools for creating and maintaining the task tracker** (P3/Bug)
    - **重要性**: 尽管优先级为 P3，但该提案探索使用原生文件工具（而非自定义工具）来维护任务追踪器（task tracker）。这体现了社区希望简化内部机制，减少对专用工具的依赖，从而让模型更好地利用其原生能力的趋势。【GitHub链接】

## 重要 PR 进展
1.  **[$28247] fix(core): match ls ignore globs by relative path**
    - 修复了 `ls` 命令中忽略规则（ignore patterns）的匹配问题。现在，包含路径分隔符的模式（如 `**/node_modules`）会相对于工作区路径正确匹配，解决了某些 glob 模式无效的 Bug。【GitHub链接】

2.  **[#28240] Fix #28227: add support for AGENTS.md out of the box**
    - 修复了 `AGENTS.md` 文件默认被忽略的问题。现在，Gemini CLI 会将其作为默认的上下文文件之一，无需用户手动在 `settings.json` 中配置，提升了开箱即用的体验。【GitHub链接】

3.  **[#28183] fix(vscode-ide-companion): preserve terminal focus when closing diff tabs**
    - 针对 VSCode 扩展问题进行了优化。修复了当用户批准代码编辑后，关闭 diff 预览标签页会“窃取”集成终端焦点的问题。现在，焦点将保留在终端，减少了用户手动点击切换的麻烦，是提升日常开发效率的“小而美”的改动。【GitHub链接】

4.  **[#28013] fix(prompts): use function replacer in applySubstitutions to prevent $-pattern corruption**
    - 修复了一个 Prompt 处理中的隐蔽 Bug。当 Skill、子代理或工具的描述中包含 `$` 符号（如模板字符串）时，`String.prototype.replace` 会错误地将其解析为特殊模式，导致内容被破坏。【GitHub链接】

5.  **[#28140] fix(deps): patch shell command dependency advisories**
    - 为修复安全漏洞，项目升级了两个关键依赖：`shell-quote`（1.8.3 -> 1.9.0）和 `simple-git`（3.28.0 -> 3.36.0）。这是一次预防性的依赖修复，旨在提升 Shell 命令处理和 Git 操作的安全性。【GitHub链接】

6.  **[#28223] fix(core-tools): bypass LLM correction for JSON and IPYNB files in write_file and replace**
    - 解决了一个关键问题：在写入或替换 `.json` 和 `.ipynb`（Jupyter Notebook）文件时，`write_file` 和 `replace` 工具会请求大模型进行“修复”或“修正”，这反而导致文件被破坏。【GitHub链接】

7.  **[#28144] fix(cli): detect available editors lazily to avoid slow startup**
    - 优化了 CLI 的启动速度。原本在启动时就会同步检查所有已知的编辑器，导致启动缓慢。现在改为懒加载（lazy）检测，只在需要时才进行检测，显著加快启动速度。【GitHub链接】

8.  **[#28153] fix(core): ignore stale update_topic calls after a session reset**
    - 修复了会话重置（如执行 `/clear` 命令）后可能出现的状态污染问题。当用户在重置会话的瞬间，模型发送了一个过时的 `update_topic` 工具调用，会导致新会话的主题被错误覆盖。【GitHub链接】

9.  **[#28143] fix(core): resolve MCP resources by server to prevent cross-server confusion**
    - 修复了 MCP（Model Context Protocol）工具的一个重大 Bug。当两个不同的 MCP 服务器暴露了相同 URI 的资源时，`read_mcp_resource` 可能返回错误服务器的内容，导致数据混淆。【GitHub链接】

10. **[#28244] docs(policy-engine): use a safe test command instead of rm -rf /**
    - 一个极具安全意识的文档改进。将策略引擎快速入门文档中的测试命令从危险的 `rm -rf /` 替换为更安全的示例，以防止用户因误操作而遭受损失。【GitHub链接】

## 功能需求趋势
1.  **Agent 智能性与自主性**: 多个 Issues 表明，社区强烈期望 Agent 能更智能地决策，包括：**主动识别并调用合适的 Skills 和子代理**（#21968）、**防止破坏性行为**（如避免强制 `git push`）（#22672）、以及**更好地了解自身的能力和限制**（#21432）。
2.  **可观察性与调试能力**: 社区对子代理的运行轨迹不可见表示不满，要求能够**通过 `/chat share` 共享子代理的完整执行轨迹**（#22598），以便于审查和评估。同时，对 Agent 状态报告的真实性（如 #22323 误报成功）要求更高。
3.  **平台兼容性与稳定性**: 仍然有相当一部分问题集中在不同平台（Wayland）和复杂环境（WSL）下的稳定性上。此外，终端卡死、焦点丢失等问题依然是修复的重点。
4.  **安全与隐私增强**: 对 Auto Memory 功能的隐私（#26525）和 MCP 资源读取的安全性（#27979）提出了更高要求，体现了在开放能力的同时，对安全底线的坚守。
5.  **开箱即用体验**: 用户希望 `AGENTS.md` 等核心配置文件能被默认加载（#28240），且 CLI 启动更快（#28144），减少手动配置和等待时间。

## 开发者关注点
-   **代理可靠性的核心痛点**: 多个高赞 Bug（#21409, #25166, #22323）直指 Agent 在任务执行、状态报告和轮次管理上的不可预测性，是目前使用体验的最大短板。
-   **交互体验的细粒度打磨**: 开发者非常关注日常交互细节，如 VSCode 焦点丢失（#28183）、Shell 命令执行卡死（#25165）和终端重绘性能（#21924），这些细微处的修复对提升整体满意度至关重要。
-   **对基础功能的依赖**: 像 `read_mcp_resource` 跨服务器出错、`ls` 忽略规则不生效这类基础工具类 Bug，虽然不显眼，但严重扰乱了核心工作流。社区的 PR 提交集中在这些领域，说明开发者希望基础功能先做到可靠。
-   **文档的安全责任**: 开发者注意到官方文档中使用了具有破坏性的示例命令（`rm -rf /`），并主动提交 PR 进行修正，这反映了社区对安全实践的严谨态度。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-07-03

## 今日速览
今日社区活跃度较高，累计有 31 条 Issue 更新，其中 2 条为新提交。最受关注的是 **alt‑screen 视图开关**（#1799）和 **gpt-5.3-codex 模型不可用** 错误（#3997），分别反映了界面自定义与模型可用性的核心痛点。此外，**BYOK 自定义模型端点**、**HTTP 代理支持**、**终端滚动与复制** 等话题也持续升温。

## 版本发布
过去 24 小时无正式版本发布。最新稳定版仍为 **v1.0.69-0**（更新于 2026-07-02）。

## 社区热点 Issues（Top 10）

| # | 标题 | 标签 | 评论/赞 | 要点 |
|---|------|------|---------|------|
| [#1799](https://github.com/github/copilot-cli/issues/1799) | How to turn off alt‑screen views? | `area:configuration`, `area:terminal-rendering` | 💬11 👍7 | 用户强烈要求恢复传统模式，alt‑screen 引发了多处显示问题，是当前终端渲染区最热话题。 |
| [#3997](https://github.com/github/copilot-cli/issues/3997) | Copilot Web: Model "gpt-5.3-codex" is not available | `triage` | 💬7 👍0 | 新用户无法使用 Agent 模式，报错指向后端模型 unavailable，疑似配置或区域限制，影响核心功能。 |
| [#3501](https://github.com/github/copilot-cli/issues/3501) | Scroll bar makes text unaligned | `area:platform-windows`, `area:terminal-rendering` | 💬6 👍9 | Windows 下垂直滚动条导致字符错位，社区普遍反馈影响日常复制和阅读。 |
| [#3936](https://github.com/github/copilot-cli/issues/3936) | Ctrl+G should expand paste tokens to full text in $EDITOR | `area:input-keyboard` | 💬3 👍1 | 用户期望在 `compactPaste` 开启时，`Ctrl+G` 能展开粘贴令牌便于编辑，对标 Claude Code 的体验。 |
| [#4019](https://github.com/github/copilot-cli/issues/4019) | Built-in web_fetch does not work with HTTP proxies | `triage` | 💬2 👍0 | 企业环境下 WSL 中无法使用 `/research` 或网页抓取功能，因 CLI 未支持代理配置，影响企业采用。 |
| [#4003](https://github.com/github/copilot-cli/issues/4003) | Support custom model endpoint in Copilot CLI (like VS Code) | `triage` | 💬2 👍0 | 请求 CLI 支持类似 VS Code 的自定义模型端点，面向本地/私有模型开发与测试场景。 |
| [#3570](https://github.com/github/copilot-cli/issues/3570) | Scrolling using touch not working on Windows | `area:input-keyboard`, `area:platform-windows`, `area:terminal-rendering` | 💬1 👍1 | Windows 触控滚动完全失效，影响触摸屏设备用户体验。 |
| [#3569](https://github.com/github/copilot-cli/issues/3569) | /clear vs /new unclear | `area:sessions` | 💬1 👍2 | 用户请求清晰化 `/clear` 和 `/new` 的区别，特别是 /new 保留会话可恢复，而 /clear 永久丢弃。 |
| [#4012](https://github.com/github/copilot-cli/issues/4012) | BYOK: reasoning effort not supported for model "glm-5.2:cloud" | `area:models`, `area:configuration` | 💬0 👍0 | 自定义 BYOK 配置下使用 `--reasoning-effort max` 报错，表明模型兼容性检测不完善。 |
| [#4014](https://github.com/github/copilot-cli/issues/4014) | Rendering all messed up when adding an MCP server on Windows | `area:platform-windows`, `area:mcp`, `area:terminal-rendering` | 💬0 👍0 | 在 Windows 上执行 `/mcp add` 时界面渲染完全错乱，可能与终端缓冲区逻辑冲突有关。 |

> 注：用户 parezanovicluka863-byte 提交的 5 条内容为无关噪音，已排除。

## 重要 PR 进展（2 条，非 10 条）

| # | 标题 | 状态 | 描述 |
|---|------|------|------|
| [#3880](https://github.com/github/copilot-cli/pull/3880) | beyond the streets of amaerica | OPEN | 提交了一个 React 组件（ArtistCard），内容与 Copilot CLI 核心功能无关，疑似误提交或测试。 |
| [#3873](https://github.com/github/copilot-cli/pull/3873) | 1000Add initial console log for greeting | OPEN | 添加了一个非常基础的 console log，代码质量低，可能为学习测试。 |

**说明**：过去 24 小时仅 2 条 PR 有更新，且均为非功能性改动或低质量提交。社区 PR 贡献目前处于低活跃期。

## 功能需求趋势

从近 24 小时更新的 Issues 中可提炼出以下社区最关注的功能方向：

1. **终端 TUI 可定制性**：用户强烈希望控制 alt‑screen、滚动条、主题等 UI 元素（#1799、#4018、#4015）。
2. **网络与代理支持**：企业用户要求 CLI 遵守系统 HTTP 代理配置，尤其针对 web_fetch 和 MCP 连接（#4019）。
3. **自定义模型与 BYOK 增强**：不仅支持私有模型端点，还要完善 `reasoning_effort` 等参数兼容性（#4003、#4012）。
4. **键盘交互与粘贴体验**：对标 Claude Code 的 `Ctrl+G` 展开粘贴令牌，以及 macOS 下的图像粘贴（#3936、#4013）。
5. **MCP 协议完善**：包括分页遵循、OAuth 流程、同名服务器冲突警告等（#4006、#4017、#3893）。
6. **非交互式命令支持**：`/init` 等命令可在脚本中自动完成而不挂起（#4011）。

## 开发者关注点

- **⚠️ 终端显示回归**：alt‑screen 引入后，Windows 触控滚动、选择复制带边栏字符、滚动条错位等问题集中爆发，开发者普遍感到烦躁。
- **🔐 认证与授权不易用**：企业环境下 BYOK 搭配 `--acp` 模式仍然需要 GitHub 登录，MCP OAuth 流程静默失败，提示不清晰。
- **📋 复制体验下降**：由于滚动条占位，复制输出会带上额外字符；“Copied to clipboard”提示在 macOS 上时机错误，导致双击粘贴失败。
- **💡 教育成本**：`/clear` 与 `/new` 功能区分不直观，新手容易误操作丢失上下文。
- **📉 PR 贡献疲软**：近 24 小时无实质性 PR 合并，社区开发活跃度偏低，可能与 CLI 稳定期或功能打磨阶段有关。

---

*数据采集时间：2026-07-03 14:00 UTC，覆盖过去 24 小时更新的 Issues/PRs。*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

好的，根据您提供的 GitHub 数据（仅包含 1 个已关闭的 Issue 和 1 个开放的 PR），以下是 2026-07-03 的 Kimi Code CLI 社区动态日报。由于数据源时间覆盖范围有限，部分板块（如版本发布、10 个 Issue/PR）将如实反映当前数据不足的情况，并基于现有信息进行深度分析。

---

# Kimi Code CLI 社区动态日报 | 2026-07-03

## 今日速览
过去 24 小时，Kimi Code CLI 仓库无新版本发布，社区活跃度较低。1 个关于 Tailscale WebSocket 连接错误的已关闭 Issue 获得最后更新，另有一个针对 Windows 终端的粘贴功能修复 PR 正在等待审查。Windows 环境下的剪贴板兼容性成为当前开发者关注的技术焦点。

---

## 版本发布
无新版本发布（最新稳定版仍为 v1.12.0）。

---

## 社区热点 Issues（共 1 条）

### 1. #1111 [已关闭] Tailscale WebSocket 连接错误
- **作者**: `tianyw0`  
- **创建**: 2026-02-12 → **最后更新**: 2026-07-02（昨日）  
- **评论数**: 2 | **👍**: 0  
- **链接**: [Issue #1111](https://github.com/MoonshotAI/kimi-cli/issues/1111)  

**摘要**:  
用户在使用 **Tailscale** VPN 环境下运行 Kimi Code CLI v1.12.0（macOS ARM 架构）时，遇到 WebSocket 连接失败，导致无法完成 AI 对话或代码补全。问题于 2 月提交后长期处于开放状态，昨日被关闭（可能已修复或转移到其他渠道）。

**为什么值得关注**：
- Tailscale 是开发者常用的远程连接/内网穿透方案，此类网络兼容性问题直接影响远程办公场景下的 CLI 可用性。
- 虽然 Issue 已关闭，但官方未标明具体的修复版本或替代方案，对于同样依赖 Tailscale 的用户存在不确定性。
- 社区反应：仅 2 条评论且无人点赞，表明该问题影响面较窄，但长期未解决凸显了网络层兼容性测试的薄弱环节。

---

## 重要 PR 进展（共 1 条）

### 1. #2481 [开放] 修复 Windows 终端下的剪贴板媒体粘贴问题
- **作者**: `redjade75723`  
- **创建**: 2026-07-01 → **最后更新**: 2026-07-02（昨日）  
- **评论数**: 0 | **👍**: 0  
- **链接**: [PR #2481](https://github.com/MoonshotAI/kimi-cli/pull/2481)  

**摘要**:  
该 PR 修复了在 **Windows Terminal** 及 VS Code 集成终端中使用 `Ctrl+V` 粘贴图片等二进制内容时失败的问题。问题根源：Windows 终端会将 `Ctrl+V` 捕获并转换为 BracketedPaste 事件，但事件中无法携带二进制数据。作者在 `_handle_bracketed_paste()` 方法中增加了首先尝试 `_try...`（推测为尝试从系统剪贴板直接读取二进制内容的逻辑）。

**为什么值得关注**：
- **痛点明确**：Windows 开发者在使用 Kimi Code CLI 时无法粘贴截图、图表等视觉上下文，严重影响多模态交互体验（如传入图片让 AI 解析 UI 设计）。
- **技术价值**：该修改涉及终端协议（BracketedPaste）与系统剪贴板 API 的交互，对跨平台剪贴板兼容性具有普遍借鉴意义。
- **社区反应**：尚未有评论，说明 PR 刚提交不久，团队需尽快 Review 以解决 Windows 用户的核心阻塞。

---

## 功能需求趋势

基于现有数据，社区目前最关注的方向可归纳为：

1. **网络连接鲁棒性**  
   - Tailscale / VPN 环境下的 WebSocket 稳定性问题（Issue #1111）表明用户对 CLI 在各种网络拓扑中的可靠性有较高期待。
   
2. **Windows 终端兼容性**  
   - PR #2481 直指 Windows 平台剪贴板粘贴二进制数据失败的痛点。这表明随着多模态功能的推进（如图片输入），Windows 用户的基础操作体验急需补齐。

---

## 开发者关注点

- **VPN 环境下连接失败**：使用 Tailscale 等工具的用户反馈长期未得到明确解决方案，可能暗示官方测试环境缺少此类网络模拟。
- **Windows 剪贴板粘贴缺失**：尽管 PR 已提出修复，但当前稳定版 v1.12.0 中用户无法正常粘贴图片，且官方未发布 hotfix，提示问题优先级较低或测试周期较长。
- **Issue 关闭原因不透明**：Issue #1111 在时隔近 5 个月后突然关闭，未附带 commit 或说明，社区难以判断是否为真正修复，需关注后续版本更新日志。

---

**小结**：今日动态集中在 Windows 剪贴板修复和长期网络问题的收尾。若您是 Windows 用户或依赖 Tailscale 连接，建议关注 PR #2481 的合并进展，并留意下一个小版本更新。

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-07-03

## 今日速览

今日社区无版本发布，但活跃度较高。**Go provider 限速误报**（#34884）与 **数据库 schema 错误**（#31119）成为最多用户反馈的阻塞性问题。在开发侧，**code-mode 实验性功能**的多项 PR 持续堆叠推进，同时 **TUI 多选删除 UX** 与 **内存泄漏修复** 的 PR 已提合并，改进落地在即。

---

## 社区热点 Issues

### 1. Go Provider 限速误报（#34884） 🔥
- **评论/点赞**: 15 / 6 👍
- **状态**: OPEN
- **摘要**: 使用 OpenCode Go 作为 provider 时持续收到 `Provider rate limit exceeded` 错误，但 dashboard 显示用量为 0。仅影响 Go 免费层，free zen 模型正常。
- **重要性**: 影响使用 Go tier 用户的使用体验，可能是服务端限速策略或客户端计数 bug。
- **链接**: https://github.com/anomalyco/opencode/issues/34884

### 2. 数据库字段缺失导致应用不可用（#31119） 🔥
- **评论/点赞**: 9 / 8 👍
- **状态**: OPEN
- **摘要**: 许久未使用后升级到 v1.16.2，启动即报 `Error: no such column: name`，数据库迁移异常导致应用完全无法使用。
- **重要性**: 升级用户的阻塞问题，点赞数高说明影响面广，需紧急修复。
- **链接**: https://github.com/anomalyco/opencode/issues/31119

### 3. MCP 工具连接但未暴露给 Agent（#33027）
- **评论/点赞**: 4 / 1 👍
- **状态**: OPEN
- **摘要**: MCP 服务器 `pdfrag` 成功连接并返回 6 个 tool 列表，但 agent 的可用工具列表中不包含这些工具。MCP 协议交互正常。
- **重要性**: MCP 集成的关键 bug，阻止用户使用外部工具扩展。
- **链接**: https://github.com/anomalyco/opencode/issues/33027

### 4. TUI 退出后关闭终端窗口（#22003）
- **评论/点赞**: 2 / 13 👍
- **状态**: OPEN
- **摘要**: Windows Terminal + cmd.exe 下，通过 Ctrl+D、`/exit` 退出 TUI 后，终端窗口直接关闭而非返回 shell 提示符。
- **重要性**: 点赞数高，Windows 用户的常见痛点，影响日常使用流程。
- **链接**: https://github.com/anomalyco/opencode/issues/22003

### 5. 内存持续增长直至进程被杀死（#35107）
- **评论/点赞**: 2 / 0 👍
- **状态**: OPEN
- **摘要**: `updatePart` 中 `structuredClone(part)` 导致 text 部分累积，93K 次 PartUpdated 事件后产生巨大堆压力。Bun 的 mimalloc 分配器不释放。
- **重要性**: 严重性能缺陷，长时间使用后内存泄漏，已有关联 PR 修复。
- **链接**: https://github.com/anomalyco/opencode/issues/35107

### 6. 6MB 请求体限制阻止合法图像输入（#35112）
- **评论/点赞**: 2 / 0 👍
- **状态**: OPEN
- **摘要**: OpenCode Go 使用 Qwen3.7Plus 时，带图像附件的请求因超过 6MB 限制而被拒，影响多模态模型使用。
- **重要性**: 限制了图像输入场景，需要提高限制或提供配置项。
- **链接**: https://github.com/anomalyco/opencode/issues/35112

### 7. 桌面更新不更新 CLI，导致版本不匹配（#35122）
- **评论/点赞**: 1 / 0 👍
- **状态**: OPEN
- **摘要**: 桌面应用更新后，全局安装的 CLI 版本未同步，导致 `opencode serve` 启动的 headless 服务与桌面版本不一致，会话同步出错。
- **重要性**: 影响桌面+CLI 混合使用用户的体验，需统一更新机制。
- **链接**: https://github.com/anomalyco/opencode/issues/35122

### 8. `session_message.seq` 非空约束失败（#35116）
- **评论/点赞**: 1 / 0 👍
- **状态**: OPEN
- **摘要**: 切换模型后进入会话，点击发送按钮无反应，日志报 `NOT NULL constraint failed: session_message.seq`。
- **重要性**: 数据库写入错误，导致聊天功能完全不可用，需紧急排查。
- **链接**: https://github.com/anomalyco/opencode/issues/35116

### 9. `upgrade` 命令不再支持降级（#34287）
- **评论/点赞**: 2 / 0 👍
- **状态**: OPEN
- **摘要**: `opencupgrade 1.15.3` 实际上只执行升级逻辑，降级不再生效。
- **重要性**: 影响用户回退版本的需求，降低版本管理灵活性。
- **链接**: https://github.com/anomalyco/opencode/issues/34287

### 10. 编辑器内联编辑项目名称不生效（#33744）
- **评论/点赞**: 2 / 0 👍
- **状态**: OPEN
- **摘要**: Desktop v1.17.x 中，双击项目名称进入编辑模式，修改后按回车保存无效。
- **重要性**: 桌面版基础功能缺陷，影响项目管理体验。
- **链接**: https://github.com/anomalyco/opencode/issues/33744

---

## 重要 PR 进展

### 1. fix(tui): 优化会话列表多选删除体验（#35121）
- **描述**: 修复了 TUI 会话列表中多选删除的多个问题：space 键无效、计数错误、删除后列表置空等。新增批量删除 CLI 命令。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35121

### 2. feat(desktop): 恢复关闭标签页 & 后台打开标签页（#35010）
- **描述**: 为桌面 v2 标签栏增加浏览器风格功能：`⇧⌘T` 恢复关闭的标签页，`⌘+点击` 在后台打开标签页。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35010

### 3. fix: 修复内存泄漏（#35111）
- **描述**: 对应 #35107，通过优化 `updatePart` 中的克隆行为，减少堆压力，防止内存持续增长。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35111

### 4. feat(tui): 渲染 code-mode 执行工具及子调用（#35113）
- **描述**: 恢复 code-mode 中 `execute` 工具的 TUI 渲染，显示 spinner、成功/失败状态以及子调用展开。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35113

### 5. feat(core): MCP elicitation via form service（#35108）
- **描述**: 基于全局表单支持，恢复 MCP elicitation 集成，使 MCP 工具在需要输入时能通过表单交互。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35108

### 6. feat: 改进模型搜索（#34954）
- **描述**: 优化作曲家模型选择器的搜索，归一化分隔符和标点，使 `gpt 5`、`gpt-5`、`gpt5` 都能匹配。
- **状态**: CLOSED (merged)
- **链接**: https://github.com/anomalyco/opencode/pull/34954

### 7. fix(rpc): 目标断开时拒绝待处理调用（#34974）
- **描述**: 当 Worker 抛出未捕获异常时，所有待处理的 RPC `call()` 会一直挂起。此 PR 在 `error`/`messageerror` 事件中 reject 所有 pending 调用。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/34974

### 8. refactor: 以原生形状暴露 MCP 工具（#35103）
- **描述**: 将 `MCP.tools()` 返回类型从 AI-SDK Tool 改为原生 `McpTool` 接口，消费者在边界适配，提高灵活性。
- **状态**: CLOSED (merged)
- **链接**: https://github.com/anomalyco/opencode/pull/35103

### 9. feat(core): 全局表单支持（#35106）
- **描述**: 在合并的表单服务上增加对 `"global"` 表单拥有者的临时支持，为 MCP elicitation 提供基础。
- **状态**: CLOSED (merged)
- **链接**: https://github.com/anomalyco/opencode/pull/35106

### 10. feat(codemode): 添加受限执行包（#35118）
- **描述**: 将 code-mode 的受限执行包（`@opencode-ai/codemode`）的 commit 集 cherry-pick 到 v2 分支，包含依赖和构建调整。
- **状态**: OPEN
- **链接**: https://github.com/anomalyco/opencode/pull/35118

---

## 功能需求趋势

从今日 issue 和 PR 中可归纳出以下社区关注方向：

- **MCP 集成深化**：多个 issue 和 PR 聚焦 MCP 工具暴露、elicitation 表单交互、原生形状重构，说明 MCP 生态正在快速落地。
- **code-mode 实验性功能**：多项 PR 推进受限执行包、TUI 渲染、实验性标志，表明项目在探索安全代码执行能力。
- **会话管理 UX**：多选删除、标签页管理、批量操作等需求持续出现，反映用户对大规模会话管理的需求。
- **多模态支持**：图像输入被 6MB 限制阻塞，社区期待提高上传限制或支持压缩。
- **版本同步与升级**：桌面与 CLI 版本不同步、upgrade 不支持降级等问题，用户希望统一的版本管理机制。
- **性能与稳定性**：内存泄漏、高 CPU 占用、数据库错误等持续得到关注。

---

## 开发者关注点

- **数据库迁移问题**：`no such column: name` 错误导致无法使用，升级流程需加强测试。
- **Provider 限速误报**：Go tier 的 rate limit 错误与 Dashboard 数据矛盾，需排查服务端或客户端计数逻辑。
- **TUI 退出行为**：Windows 下关闭终端窗口属于 unsafe 行为，应修复为返回 shell。
- **MCP 工具接入**：连接成功但 agent 不可见，沟通成本高，开发者呼吁完善 MCP 集成文档和测试。
- **版本回退困难**：`upgrade` 降级功能缺失，影响用户尝试新版本后的回退路径。
- **请求大小限制**：6MB 上限对图像输入过紧，建议提供可配置选项或动态压缩。
- **桌面与 CLI 分离**：更新不同步引发会话同步问题，建议将 CLI 捆绑在桌面包内或增加同步检测。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi 社区动态日报 (2026-07-03)

## 今日速览

Pi v0.80.3 用户遭遇多处 `content is not iterable` 崩溃，核心原因集中在推理模型返回 null content 与工具调用的交互上；同时新 Claude 模型与编辑工具的兼容性出现约 20% 的失败率，社区提交了多项针对性的修复与功能请求。此外，DeepInfra、GLM 等新供应商的集成 PR 已合并，生态持续扩展。

---

## 社区热点 Issues（10 条）

1. **#6215** `[OPEN] pi update fails on 0.80.3 due to missing @smithy/node-http-handler@^4.9.1`  
   **重要性**：影响用户升级到最新版本，官方尚未修复，22 条评论说明受影响面广。  
   **链接**：https://github.com/earendil-works/pi/issues/6215

2. **#6259** `[OPEN] fix: 'content is not iterable' when reasoning models return null content during tool use`  
   **重要性**：直指当前版本最严重的运行时崩溃，3 条评论已确认影响多个推理模型（如 GLM-5.2）。  
   **链接**：https://github.com/earendil-works/pi/issues/6259

3. **#6278** `[OPEN] New Claude models work poorly with the current Pi's edit tool, failing about 20% edits`  
   **重要性**：统计数据显示编辑失败率达 20%，LLM 会添加无关字段，严重降低 Claude 系列模型的可用性。  
   **链接**：https://github.com/earendil-works/pi/issues/6278

4. **#6157** `[OPEN] Compaction summary should be in the session's language, and the update step should dedup`  
   **重要性**：非英语用户的核心体验诉求，4 条评论支持建议在摘要中使用会话语言并去重。  
   **链接**：https://github.com/earendil-works/pi/issues/6157

5. **#6204** `[OPEN] mimo-v2-omni is a ghost model on the three MiMo Token Plan providers`  
   **重要性**：模型列表中存在但实际 API 不支持，导致 400 错误，影响小米云 Token 用户。  
   **链接**：https://github.com/earendil-works/pi/issues/6204

6. **#6262** `[CLOSED] DS4-server context overflow errors not detected by auto-compaction`  
   **重要性**：本地 DeepSeek V4 超出上下文时自动压缩未触发，导致请求被拒，虽已关闭但修复方案值得关注。  
   **链接**：https://github.com/earendil-works/pi/issues/6262

7. **#6270** `[CLOSED] Add DeepInfra as a built-in provider`  
   **重要性**：社区高票需求（👍2），覆盖文本和图像生成，已合并 PR，标志 Pi 供应商生态进一步丰富。  
   **链接**：https://github.com/earendil-works/pi/issues/6270

8. **#6276** `[CLOSED] v0.80.3: content is not iterable — unguarded for..of on message.content in compaction.js + render-utils.js`  
   **重要性**：定位到具体文件，与 #6259 同类但触发路径不同，1 条评论即关闭，说明已被快速修复。  
   **链接**：https://github.com/earendil-works/pi/issues/6276

9. **#6257** `[CLOSED] Add claude-sonnet-5 to GitHub Copilot model catalog`  
   **重要性**：Copilot 用户要求支持最新 Sonnet 5 模型，迅速关闭意味着已被采纳。  
   **链接**：https://github.com/earendil-works/pi/issues/6257

10. **#6227** `[OPEN] feat: sqlite session storage`（注：此为 PR，但 issue 侧无对应，此处借用）  
    **重要性**：用户对持久化存储的强烈需求，支持并行写入 SQLite 和 JSONL，评论中讨论热烈。  
    **链接**：https://github.com/earendil-works/pi/pull/6227

---

## 重要 PR 进展（10 条）

1. **#6279** `[OPEN] fix(coding-agent): add pnpm self-update prune hint`  
   针对 #6215 的更新失败问题，提供 pnpm 缓存清理提示，尚未合并但提供临时自救方案。  
   **链接**：https://github.com/earendil-works/pi/pull/6279

2. **#6266** `[CLOSED] Anthropic: strict tool use for the edit tool`  
   强制 Claude 系列使用严格工具调用模式，预计可大幅降低编辑失败率（从约 10% 降至接近 0）。  
   **链接**：https://github.com/earendil-works/pi/pull/6266

3. **#6273** `[CLOSED] Add Zen mode tool call labels`  
   新增 `zenMode` 设置，在 TUI 中压缩工具调用标签，提升交互体验（使用 GPT-5.4-mini 异步摘要）。  
   **链接**：https://github.com/earendil-works/pi/pull/6273

4. **#3799** `[CLOSED] add azure cognitive services as provider`  
   支持 Azure Cognitive Services 的 base URL，扩展 Azure OpenAI 供应商兼容性。  
   **链接**：https://github.com/earendil-works/pi/pull/3799

5. **#6271** `[CLOSED] [codex] Add GLM API provider`  
   正式添加智谱 AI 和 z.ai 的 GLM 系列模型供应商（`glm` 和 `glm-cn`），使用 `ZAI_API_KEY`。  
   **链接**：https://github.com/earendil-works/pi/pull/6271

6. **#6267** `[OPEN] feat(coding-agent): add InlineExtension type for named inline extension factories`  
   允许内联扩展工厂具有名称，便于在启动日志中识别（如 `[Extensions] user-provided-factory`）。  
   **链接**：https://github.com/earendil-works/pi/pull/6267

7. **#6264** `[CLOSED] fix(ai): clamp OpenAI Responses output tokens`  
   修复 `max_output_tokens` 在接近上下文限制时可能低于 API 最小值（16）的问题，避免 400 错误。  
   **链接**：https://github.com/earendil-works/pi/pull/6264

8. **#6263** `[CLOSED] feat(ai): add DeepInfra provider for text and image generation`  
   整合 DeepInfra 的 chat 和图像生成 API，动态拉取模型列表，是社区热切期待的新供应商。  
   **链接**：https://github.com/earendil-works/pi/pull/6263

9. **#6258** `[CLOSED] fix: guard against null content in agent loop, compaction, and message transforms`  
   针对 #4909 系列问题，在 agent 循环、压缩和消息转换中增加 null 保护，解决 `content is not iterable` 崩溃。  
   **链接**：https://github.com/earendil-works/pi/pull/6258

10. **#6227** `[OPEN] feat: sqlite session storage`  
    新增可选的 SQLite 会话存储（通过 `PI_SQLITE_SESSION_STORAGE=1` 启用），与 JSONL 并行写入，满足持久化查询需求。  
    **链接**：https://github.com/earendil-works/pi/pull/6227

---

## 功能需求趋势

- **模型与供应商扩展**：DeepInfra、GLM、Azure Cognitive Services、Claude Sonnet 5、Kimi K2.7 等新模型/供应商请求频繁，社区希望 Pi 能快速跟进主流 API。
- **工具调用健壮性**：多篇 Issue 聚焦于工具调用中 `null content`、多余字段、转义错误等问题，用户期望 LLM 输出能更严格地遵循 schema。
- **UI/UX 优化**：Zen 模式（压缩工具标签）、会话语言自适应、键盘绑定灵活性、扩展名称显示等，体现用户对 TUI 交互细节的追求。
- **存储与持久化**：SQLite 会话存储提案获得正面反响，用户希望摆脱纯 JSONL 的查询困难。
- **扩展 API 增强**：要求扩展能够注册类型化设置 schema、命名工厂，并解决多会话间上下文污染问题。

---

## 开发者关注点

- **`content is not iterable` 连环崩溃**：v0.80.3 中因推理模型返回 `null content` 导致多处循环崩溃，是当日最高频错误，修复 PR #6258 已合并但仍有其他路径未覆盖。
- **编辑工具兼容性**：新 Claude 模型在 edit tool 上的高失败率（20%）引发担忧，开发者期待 strict tool use 能彻底解决。
- **依赖解析失败**：pnpm 元数据缓存导致 `pi update` 失败（#6215），用户被迫手动清理，项目组已提供提示 PR。
- **上下文溢出检测缺失**：本地 DeepSeek V4 的上下文超限未被自动压缩检测，用户需要手动处理。
- **WebSocket 连接超时**：Codex WebSocket 在 60 分钟后断开且不重连，长任务体验受损。
- **键盘绑定问题**：`shift+enter` 绑定不生效、扩展功能键插入 PUA 字符等问题影响日常输入。

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code 社区动态日报 | 2026-07-03

---

## 📊 今日速览

今日 Qwen Code 发布三个版本更新，包括正式版 v0.19.5 和两个 Nightly 版本，重点修复了移动端会话切换卡顿、加固了 daemon 通道管理，并推出了新的 CUA 驱动二进制包。社区方面，共 6 个活跃 Issue 和 50 个 PR 得到更新，其中空参数工具调用静默丢弃（P1 严重 bug）和扩展能力变更通知问题最受关注。新功能方面，daemon 状态仪表盘、自定义代码块渲染、及 @ 提及面板等 PR 正持续推进。

---

## 📦 版本发布

### 1. `v0.19.5-nightly.20260703.b16baf1ff`
- **主要修复**：重写 web-shell 移动端会话切换的 Timeline 签名为 memoized 版本，并采用 replay-first 派发逻辑，消除切换时的卡顿（jank）。
- **链接**：[Release v0.19.5-nightly.20260703.b16baf1ff](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.5-nightly.20260703.b16baf1ff)

### 2. `cua-driver-rs-v0.7.0`
- **亮点**：预编译二进制新增**相对坐标 fork** 支持（vendored 于 `packages/cua-driver` 下）。
- **平台覆盖**：
  - macOS：已代码签名、公证的通用二进制 + `QwenCuaDriver.app`
  - Linux：x86_64 & arm64（未签名，glibc ≥ 2.31）
  - Windows：x86_64 & arm64（未签名）
- **链接**：[cua-driver-rs v0.7.0](https://github.com/QwenLM/qwen-code/releases/tag/cua-driver-rs-v0.7.0)

### 3. `v0.19.5`（正式版）
- **CLI**：加固 daemon 管理的通道工作者（worker）稳定性（PR #6098）
- **Web Shell**：延迟会话创建，直到用户首次输入 prompt 才真正初始化（PR #6107）
- **链接**：[Release v0.19.5](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.5)

---

## 🔥 社区热点 Issues

以下为过去 24 小时内有更新的全部 6 个 Issue，均已标注优先级和状态：

### 1. [P1 Bug] #6249 – 流式工具调用空 `arguments` 被静默丢弃，导致“Model stream ended with empty response text”重试循环
- **摘要**：当 OpenAI 兼容 provider 流式返回一个工具调用且 `function.arguments` 为空字符串 `""` 时，解析器会静默丢弃整个工具调用。若后续没有补充片段，则引发重试风暴。
- **社区反应**：2 条评论，标注 `welcome-pr` 鼓励社区贡献。
- **链接**：[Issue #6249](https://github.com/QwenLM/qwen-code/issues/6249)

### 2. [P2 Bug] #6246 – `qwen_code` 无法识别自身所属进程
- **摘要**：用户通过 qwen_code 启动后端 Node.js 进程后，`停止进程`指令会误杀所有 Node 进程（包括 qwen_code 自身），导致会话崩溃。
- **社区反应**：2 条评论，期待进程自识别修复。
- **链接**：[Issue #6246](https://github.com/QwenLM/qwen-code/issues/6246)

### 3. [P2 Bug] #6244 – 扩展能力变更未可靠通知模型
- **摘要**：在活跃会话中启用/禁用/安装/卸载扩展后，模型可能未及时感知技能、命令或 MCP 工具的变化，导致行为不一致。
- **社区反应**：2 条评论，标注 `welcome-pr`。
- **链接**：[Issue #6244](https://github.com/QwenLM/qwen-code/issues/6244)

### 4. [P3 Feature] #6252 – 为 `qwen serve` 添加 daemon 状态仪表盘
- **摘要**：功能请求，希望基于 `GET /daemon/status` API 实现浏览器端可视化仪表盘，展示会话、权限、传输、速率限制、工作区状态等。
- **社区反应**：2 条评论，同日有对应 PR #6253 已提交。
- **链接**：[Issue #6252](https://github.com/QwenLM/qwen-code/issues/6252)

### 5. [Closed] #6251 – OAuth 认证失败，返回 504 Gateway Timeout
- **摘要**：用户认证时授权 URL 正常打开，但回调超时，社区反馈影响工具使用。
- **处理**：已关闭，推测为临时网络/服务问题。
- **链接**：[Issue #6251](https://github.com/QwenLM/qwen-code/issues/6251)

### 6. [#3804（旧 Issue 今日更新）] – 0.15.6 版本 `AskUserQuestion` 频繁导致空响应错误
- **摘要**：用户反映该版本中 `AskUserQuestion` 功能易触发 `[API Error: Model stream ended with empty response text.]`。
- **社区反应**：4 条评论，累积 👍 0，需要持续关注。
- **链接**：[Issue #3804](https://github.com/QwenLM/qwen-code/issues/3804)

---

## 🚀 重要 PR 进展

从 50 条 PR 中精选 10 项关键变更：

### 1. [feat] #6253 – 添加 daemon 状态仪表盘（对应 Issue #6252）
- **概要**：提供独立 HTML 页面 `/dashboard`，可视化展示 daemon 运行时状态，并可与 web-shell 组件集成。
- **链接**：[PR #6253](https://github.com/QwenLM/qwen-code/pull/6253)

### 2. [fix] #6261 – CI 修复：恢复沙箱镜像流
- **概要**：将 Qwen Autofix 切回 GitHub-hosted sandbox 执行，移除本地 OpenAI 环回代理，并允许使用专用 autofix 模型密钥。
- **链接**：[PR #6261](https://github.com/QwenLM/qwen-code/pull/6261)

### 3. [feat] #6232 – web-shell 支持自定义代码块渲染
- **概要**：允许宿主替换 Markdown 代码块为自定义 React 组件，并附带 `web-shell-charts` 技能（用于图表渲染）。
- **链接**：[PR #6232](https://github.com/QwenLM/qwen-code/pull/6232)

### 4. [feat] #6242 – web-shell 添加自定义 @ 提及面板
- **概要**：取代旧内联自动补全，实现多级参考面板：文件、活动扩展、MCP 资源等分类搜索，支持键盘/鼠标导航。
- **链接**：[PR #6242](https://github.com/QwenLM/qwen-code/pull/6242)

### 5. [fix] #6236 – 修复 web-shell 视觉模型选择器编码错误
- **概要**：选择器传递 `modelId(authType)` 格式，但核心解析时需 `authType:modelId`，导致选择失效。此 PR 修正解析逻辑。
- **链接**：[PR #6236](https://github.com/QwenLM/qwen-code/pull/6236)

### 6. [fix] #6238 – 为 Stop-hook 延续分配独立工具调用预算
- **概要**：每个阻塞 Stop-hook 迭代（如 `/goal`）视为新轮次，独立计数工具调用次数，避免全局预算耗尽。同时使上限可配置。
- **链接**：[PR #6238](https://github.com/QwenLM/qwen-code/pull/6238)

### 7. [fix] #6245 – 扩展能力变更时主动通知模型（对应 Issue #6244）
- **概要**：跟踪已通知的 MCP 工具、技能/命令、Agent 子类型，发生变更时发送增量更新给模型。
- **链接**：[PR #6245](https://github.com/QwenLM/qwen-code/pull/6245)

### 8. [fix] #6243 – 保留无描述的 OpenAI 工具
- **概要**：当 Gemini/MCP 函数声明缺少 `description` 字段时，工具仍应可见。此 PR 确保描述为空时不跳过序列化。
- **链接**：[PR #6243](https://github.com/QwenLM/qwen-code/pull/6243)

### 9. [fix] #6240 – 保留遗留 OpenAI 函数调用格式
- **概要**：非流式响应中的 `message.function_call` 和流式 `delta.function_call` 在转换为 Gemini content 时被正确处理。
- **链接**：[PR #6240](https://github.com/QwenLM/qwen-code/pull/6240)

### 10. [feat] #5953 – LSP 服务器支持热重载
- **概要**：当 `.lsp.json` 在活动会话中变更时，Qwen Code 检测语义变化并自动重新加载 LSP 配置，无需重启会话。
- **链接**：[PR #5953](https://github.com/QwenLM/qwen-code/pull/5953)

---

## 📈 功能需求趋势

从今日 Issue 和 PR 中可提炼以下社区关注方向：

1. **Daemon 运行时可视化** – Issue #6252 + PR #6253 表明用户希望获得更直观的守护进程状态面板，便于运维和调试。
2. **Web Shell 交互增强** – 自定义代码块渲染（#6232）、@ 提及面板（#6242）显示社区对浏览器端体验的重视。
3. **扩展能力动态同步** – #6244 / #6245 直接指向扩展变更时模型感知问题，是提升插件生态关键点。
4. **OpenAI 兼容性优化** – #6249（空arguments）、#6243（无描述工具）、#6240（遗留函数调用）体现社区在适配多 provider 时遇到的边界问题。
5. **进程管理安全** – #6246 揭示 qwen_code 自身进程标识缺失，属于核心安全/易用性需求。
6. **LSP 热重载** – #5953 长期需求，降低开发者在前后端联调中的重启成本。

---

## 👨‍💻 开发者关注点

- **流式错误处理不透明** – #6249 和 #3804 均反映空响应导致的无限重试，开发者期望提供更明确的错误类型或降级策略。
- **进程自识别缺失** – #6246 的误杀问题说明 `qwen_code` 在管理子进程时需引入进程树隔离或白名单机制。
- **扩展状态同步迟滞** – #6244 中模型“看不见”新技能，影响即时体验，需确保变更实时推送给模型上下文。
- **OAuth 稳定性** – #6251 虽已关闭，但 504 超时提示认证链可能存在偶发瓶颈，需要监控和重试机制。
- **UI/UX 细节打磨** – #6236 的模型选择器编码错误属于典型的前后端格式不一致问题，提醒开发者注意协议统一性。

---

*日报生成时间：2026-07-03 | 数据来源：[QwenLM/qwen-code](https://github.com/QwenLM/qwen-code)*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI 社区动态日报 — 2026-07-03

## 📌 今日速览

开发团队于今日推送了 **v0.8.67 审计修复批量 PR**（#3960），集中修正了核心流程中的多项正确性与 UX 问题，包括 Constitution 可视化、信任步骤逻辑、技能系统虚假安全等。社区对新入门流程（Onboarding）和 Fleet 子代理内存溢出的反馈最为集中，同时多项关于 **MCP 动态启动、规则目录自动发现** 的功能 PR 已进入合并阶段。

---

## 🔥 社区热点 Issues（10 条）

### 1. #3793 — v0.8.67 Setup: 构建引导式语言优先的 Constitution 创建器
- **评论**: 14 | **状态**: 开放
- **重要性**: v0.8.67 核心功能重设计，社区高度关注 Constitution 文件不应混合运行时安全控制。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3793

### 2. #3867 — 项目级指令过度受限，需 glob 与规则目录自动发现
- **评论**: 7 | **状态**: 已关闭（已合并 PR #3892）
- **重要性**: 多项目工作流的痛点，已通过规则目录自动发现解决，社区反馈积极。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3867

### 3. #3792 — 首次运行引导应像启动 CodeWhale 而非编辑配置
- **评论**: 7 | **状态**: 开放
- **重要性**: 入门体验关键改进，避免新用户被复杂的配置步骤吓退。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3792

### 4. #3928 — 应用内无法读取 Constitution，自定义覆盖无声失败
- **评论**: 2 | **状态**: 开放（7月2日新开）
- **重要性**: v0.8.67 核心功能的用户可感知缺陷，影响所有用户。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3928

### 5. #3927 — API-Key 步骤为硬性 DeepSeek 专用门，无法跳过或选择其他提供商
- **评论**: 2 | **状态**: 开放
- **重要性**: 多提供商支持的阻塞点，新用户若使用 OpenAI 等将直接无法进入。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3927

### 6. #3926 — 信任步骤「拒绝即退出」，Enter 键无响应
- **评论**: 2 | **状态**: 开放
- **重要性**: 入门流程一致性的严重缺陷，用户可能意外退出应用。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3926

### 7. #3920 — 技能系统 `allowed-tools` 声明与 `/skill trust` 均为无操作
- **评论**: 2 | **状态**: 已关闭（PR #3960 修复）
- **重要性**: 安全功能虚假承诺，用户误以为已隔离不受信任的技能，实际无解析。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3920

### 8. #3919 — 技能名称无验证，不存在的技能被广告、重复技能静默覆盖
- **评论**: 2 | **状态**: 已关闭（PR #3960 修复）
- **重要性**: 技能系统的正确性基础漏洞，导致用户困惑与不可用。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3919

### 9. #3882 — Fleet 子代理输出未限流，导致 TUI 内存占用 15GB+
- **评论**: 2 | **状态**: 已关闭（PR #3931 修复）
- **重要性**: v0.8.67 发布阻塞级问题，影响 Fleet 功能的可用性。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3882

### 10. #3965 — 子代理级别的提供商显式路由 + LM Studio 支持
- **评论**: 1 | **状态**: 开放（7月3日新开）
- **重要性**: 社区对多提供商异构部署的迫切需求，该 issue 提出为不同子代理指派不同模型（本地/远程）。
- **链接**: https://github.com/Hmbown/CodeWhale/issues/3965

---

## 🚀 重要 PR 进展（10 条）

### 1. #3968 — feat(engine): 公开 `context_input_budget_for_route` 供外部复用
- **状态**: 开放 | **作者**: h3c-hexin
- **内容**: 使外部主机/编辑器能复用内部的输入预算计算逻辑，减少重复实现。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3968

### 2. #3967 — perf(tui): 避免每帧冗余的 Composer 输入换行
- **状态**: 开放 | **作者**: reidliu41
- **内容**: 修复输入文本每帧被换行多达5次的问题，提升渲染性能。 fixes #3909。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3967

### 3. #3963 — fix(mcp): 仅在资源存在时才广告 list-resource 元工具
- **状态**: 开放 | **作者**: h3c-hexin
- **内容**: 修复 MCP 服务器即使未暴露任何资源也向模型注册列表工具的问题，减少令牌浪费。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3963

### 4. #3866 — feat: LLM 可从聊天上下文动态启动 MCP 服务器
- **状态**: 开放 | **作者**: bistack
- **内容**: 新增 `start_mcp_server` 工具，支持 stdio 和 HTTP 传输，扩展了 CodeWhale 的插件能力。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3866

### 5. #3962 — fix(tui): 在 `/config ask-rules` 中显示权限规则动作
- **状态**: 开放 | **作者**: greyfreedom
- **内容**: 为 `permissions.toml` 规则增加 `action` 列，直观显示 allow/deny/ask。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3962

### 6. #3960 — fix: v0.8.67 审计正确性批量修复（#3918/#3919/#3920/#3924/#3926/#3927/#3928/#3929）
- **状态**: 已合并 | **作者**: Hmbown
- **内容**: 包含技能名称验证、权限规则显示、信任步骤逻辑、Constitution 可视化等多项修复，是今日最重要的批量修复。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3960

### 7. #3892 — feat(tui): 自动发现 `.codewhale/rules/` 和 `.claude/rules/` 目录作为项目上下文
- **状态**: 已合并 | **作者**: yekern
- **内容**: 实现 #3867，自动扫描项目规则目录并加载，免去手动指定。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3892

### 8. #3936 — fix(subagent): 每次原子状态写入使用唯一临时路径，防止并发序列化损坏
- **状态**: 已合并 | **作者**: Hmbown
- **内容**: 修复子代理状态文件并发写入时可能损坏的竞态条件。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3936

### 9. #3931 — fix(fleet): 强制绝对递归深度上限，扩大任务 ID 熵
- **状态**: 已合并 | **作者**: Hmbown
- **内容**: 修复 Fleet 子代理递归调用的深度限制与 ID 碰撞问题，提升稳定性。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3931

### 10. #3643 — feat(setup): 添加设置向导汇总步骤（MCP/技能/插件概览）
- **状态**: 已合并 | **作者**: cy2311
- **内容**: 实现 v0.8.67 设置向导的第一步，展示当前环境概况，提升首次配置透明度。
- **链接**: https://github.com/Hmbown/CodeWhale/pull/3643

---

## 📊 功能需求趋势

- **多提供商/异构部署**：社区强烈要求支持按子代理分配不同模型（如本地 LM Studio 与云端 GPT 混用），#3965 即代表这一方向。
- **Fleet 子代理生态**：Fleet 近期问题密集，包括内存溢出、配置文件分裂、向导死胡同、角色词汇不统一等，表明用户开始深度使用并揭露粗糙点。
- **MCP 动态扩展**：支持 LLM 在对话中启动 MCP 服务器（#3866）是社区期待已久的弹性插件能力，与外部工具整合需求旺盛。
- **入门体验重塑**：Constitution 创建器、API-Key 多提供商、信任步骤逻辑等多项 UX issue 表明新用户的引导路径急需重写。
- **安全可审计性**：技能系统 `allowed-tools` 虚假安全、权限规则执行不透明等问题凸显社区对安全可视化的要求。

---

## 🧠 开发者关注点

- **多项目管理痛点**：`instructions` 配置键在项目作用域被硬性阻止，导致跨项目工作流几乎不可用（#3867 虽已修复，但此前影响广泛）。
- **中文用户反馈**：Agent 实时输出中文乱码问题（#1675）长期未解，影响华人用户；此外 token 成本估算缺乏人民币单位（#1607）。
- **Windows 原生体验**：默认启动使用 cmd.exe 而非 Windows Terminal（#1854），且命令风格自适应能力差（#1754），Windows 用户呼吁改进。
- **就地更新机制**：虽已有检查更新功能，但应用内提示不持久、不可操作（#3961），社区期望一键更新体验。
- **Constitution 文件设计争议**：v0.8.67 将 Constitution 定位为可引导创建，但社区担忧其与运行时安全控制的混合，并反映缺乏应用内查看能力（#3928）。

---

> 数据来源：GitHub Hmbown/CodeWhale，统计时间截至 2026-07-03 08:00 UTC。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*