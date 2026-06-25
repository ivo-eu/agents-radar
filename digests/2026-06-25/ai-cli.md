# AI CLI 工具社区动态日报 2026-06-25

> 生成时间: 2026-06-25 10:25 UTC | 覆盖工具: 9 个

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

好的，作为资深技术分析师，基于您提供的2026年6月25日各AI CLI工具社区动态日报，以下为您生成一份横向对比分析报告。

---

## AI CLI 工具社区动态横向分析报告 (2026-06-25)

### 1. 生态全景

当前AI CLI工具生态正处于从“能用”向“好用、可靠、安全”过渡的关键阶段，但普遍面临安全策略误报、配额/速率限制不透明、Agent行为不可预测三大核心痛点。各工具在MCP（模型上下文协议）集成、多模型供应商支持和跨平台兼容性上投入显著，体现出生态正快速向开放、可扩展方向演进。同时，开发者社区对工具的信任度高度依赖于稳定性与可解释性，安全误报和配额计算错误已成为影响付费用户留存的首要风险。

### 2. 各工具活跃度对比

| 工具名称 | 热点Issues | 重要PR | 版本发布 | 社区热点关键词 |
|---------|-----------|--------|---------|------------|
| **Claude Code** | 10 | 3 | 2个稳定版 | 安全误报/ClAudit、Agent Teams Bug、子代理内存 |
| **OpenAI Codex** | 10 | 10 | 2个稳定+3个alpha | 配额异常、模型容量满、Hooks功能、Windows沙箱 |
| **Gemini CLI** | 10 | 10 | 1个nightly | 子代理状态误报、路径遍历修复、Auto Memory脱敏 |
| **GitHub Copilot CLI** | 10 | 2 | 1个稳定版 | 会话恢复认证丢失、模型列表为空、技能子文件夹 |
| **Kimi Code CLI** | 4 | 2 | 无 | 循环读取bug、上下文压缩浪费Token、MCP子代理传播 |
| **OpenCode** | 10 | 10 | 1个稳定版 | Windows段错误、IME输入崩溃、MCP资源模板 |
| **Pi** | 10 | 9 | 无 | OpenAI Codex连接卡死、Amazon Bedrock Mantle、编排器 |
| **Qwen Code** | 8 | 10 | 3个版本（稳定+预览+nightly） | 流式超时、语音听写、上下文压缩网关超时、LSP路由 |
| **CodeWhale** | 5 | 10 | 无 | Fleet负载自动选择、ACP协议暴露、Windows环境变量继承 |

### 3. 共同关注的功能方向

- **安全策略优化（误报降噪）**：Claude Code、Gemini CLI、Copilot CLI 社区均有大量关于安全分类器/AUP策略误报的投诉，涵盖恶意软件分析、云IAM运维、键盘输入等场景。这表明当前基于规则或静态关键词的安全检测机制已严重干扰正常开发工作，社区普遍呼吁更智能、上下文感知的决策引擎。
- **配额/速率限制透明化与修复**：OpenAI Codex、Copilot CLI 用户强烈反映配额计算错误、模型容量满、5小时限制异常消耗等问题，直接导致付费用户无法正常使用。社区要求提供更清晰的用量仪表盘和日志审计能力。
- **Agent Teams与子代理行为可靠性**：Claude Code 的 Agent Teams 领导指针错误、Gemini CLI 的子代理Max Turns误报为成功、OpenCode 的子代理中断恢复等，均指向多Agent协作场景下的状态管理缺陷。社区期望更可预测、可中断的委托执行机制。
- **MCP（模型上下文协议）生态成熟化**：Claude Code、Kimi Code、OpenCode、Qwen Code、CodeWhale 均在本日有MCP相关的PR（资源模板、OAuth、子代理配置传播、默认搜索启用等）。MCP正成为事实上的插件标准，但配置一致性、路径解析、可选服务器非阻塞启动等细节仍需打磨。
- **Windows平台兼容性与稳定性**：OpenCode（Bun段错误）、Copilot CLI（沙箱COM+错误）、CodeWhale（环境变量继承）、Qwen Code（无直接提及但社区有反馈）都反映出Windows用户面临的独特痛点。随着CLI工具向企业桌面渗透，Windows支持成为不可忽视的硬性要求。
- **多模型/供应商路由与配置自动化**：Pi（新增Bedrock Mantle）、CodeWhale（Fleet Loadout Auto）、OpenAI Codex（模型容量问题间接推动）、Gemini CLI（128工具限制）等反映了社区对灵活切换后端模型、自动分配计算资源的强烈需求，尤其是在成本优化和容灾场景下。

### 4. 差异化定位分析

- **Claude Code**：安全导向最突出，拥有独立的ClAudit分类器和AUP策略模块，但也因此误报最多；Agent Teams协作模式独特但尚不稳定。目标用户为需要强安全合规的企业开发者。
- **OpenAI Codex**：生态最成熟，配额的精细化管理是其双刃剑——虽提供了多档订阅，但计算逻辑易引发争议。TUI体验持续优化（Markdown渲染、线程管理），适合追求稳定和丰富插件的重度用户。
- **Gemini CLI**：Google生态集成（ACP、A2A协议）和系统级安全保障（路径大小写敏感封锁、thought泄漏修复）是其亮点。Agent行为修复投入大，但子代理可靠性仍待提升。适合对安全与可审计性要求高的团队。
- **GitHub Copilot CLI**：深度绑定GitHub生态，技能管理（子文件夹、`preToolUse`钩子）和会话持久化是核心短板。企业代理支持（Kerberos）是差异化阵地。适合GitHub企业用户。
- **Kimi Code CLI**：体量最小但迭代快，MCP配置传播和上下文压缩Token浪费是其痛点，自定义模型端点稳定性饱受诟病。适合对成本敏感、愿意尝试新兴工具的个人开发者。
- **OpenCode**：基于Bun运行时的性能优势与段错误并存的矛盾体。MCP资源模板和`--mini`模式显示其轻量、易扩展定位，但Windows稳定性是当前最大软肋。适合跨平台在macOS/Linux上追求性能的用户。
- **Pi**：连接可靠性和流处理健壮性是其研发重点（空闲超时、错误重试），同时实验性编排器（orchestrator）和RPC接口透露出向守护进程/可编程后端演进的野心。适合需要作为后台服务嵌入工作流的团队。
- **Qwen Code**：语音听写和远程LSP路由是独特卖点，CI自动化修复工作流效率高。上下文压缩的流式化问题表明其对代理部署场景有深度考量。适合中英文混合场景、需要语音交互的用户。
- **CodeWhale**：Fleet负载自动选择与多模型路由是核心特色，ACP协议开放性和品牌重塑（去DeepSeek化）显示其试图成为独立于模型供应商的底层平台。适合搭建内部AI代理集群的组织。

### 5. 社区热度与成熟度

- **高热度、高成熟度**：**Claude Code**、**OpenAI Codex** 社区体量最大，Issue/PR讨论深入，版本号已进入2.x/0.14x，功能覆盖面广，但安全误报和配额问题对其信任度造成冲击。**GitHub Copilot CLI** 作为微软生态组件，用户基数大但参与度相对温和。
- **高热度、快速迭代**：**Gemini CLI**、**OpenCode**、**Pi** 日更版本或处于大规模重构期，PR数量多、Bug报告密集，社区活跃度极高，但也反映出稳定性不足。**Qwen Code** 发布节奏稳定（一日三版），语音和CI方向领先。
- **中等热度、追赶阶段**：**Kimi Code CLI**、**CodeWhale** 社区体量较小，但关键PR和Issue针对性较强，代表新兴玩家在特定领域（MCP、自动载荷）的差异化布局。

### 6. 值得关注的趋势信号

- **安全策略“自动化僵局”**：过多误报正在摧毁自动化模式（auto-mode）的信任基础。开发者开始主动关闭安全过滤器或降级到手动模式，这违背了AI CLI工具提升效率的初衷。未来需要引入可解释的决策树、用户反馈闭环和动态灵敏度校准。
- **Agent状态管理成为系统级挑战**：子代理“幽灵复活”、领导指针粘连、失败被误报为成功等现象说明当前Agent架构在设计上缺乏机器可读的「契约状态」和「中断恢复」原语。这将成为下一代Agent框架必须解决的基础问题。
- **MCP正从可选插件进化为核心协议**：多个工具在本日PR中不约而同地强化MCP的默认启用、子代理继承、OAuth支持、资源模板等能力。MCP很可能在2027年前成为AI CLI工具的“USB-C接口”——统一、通用、必要。
- **语音交互首次进入CLI生态**：Qwen Code的语音听写功能从CLI扩展到桌面和Web Shell，标志着AI CLI工具开始探索多模态输入。虽然目前仅限中文开发者，但预示着未来“说话编程”将从Demo走向实用。
- **跨实例编排与守护进程化**：Pi的编排器（orchestrator）和CodeWhale的Fleet负载自动选择，暗示AI CLI工具正从单次对话工具进化为可长期运行、可水平扩展的AI Agent基础设施。这与云原生和微服务架构的思维方式一脉相承。
- **成本透明化压力剧增**：OpenAI Codex、Copilot CLI的配额计算争议表明，用户对“看不见的消耗”容忍度极低。未来工具必须提供实时的Token消耗明细、按任务估算成本、以及可配置的配额预警，否则将面临流失风险。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills 社区热点报告（数据截止 2026-06-25）

## 1. 热门 Skills 排行

以下按评论数降序列出当前关注度最高的 PR（均为 Open 状态）：

| 排序 | PR | 功能概要 | 社区关注焦点 | 状态 |
|------|----|----------|--------------|------|
| 1 | [#1298 fix(skill-creator): run_eval.py always reports 0% recall](https://github.com/anthropics/skills/pull/1298) | 彻底修复评估脚本（窗口流、触发检测、并行 worker），使描述优化循环能真正生效。 | 关联多个独立复现的 bug（#556、#1169），严重阻碍技能创建者优化描述，是当前最热的修复 PR。 | Open |
| 2 | [#514 Add document-typography skill](https://github.com/anthropics/skills/pull/514) | 新增文档排版质量技能，解决孤儿词、寡妇段落、编号对齐等 AI 生成文档的常见问题。 | 用户反馈几乎所有 AI 文档都存在这些排版缺陷，需求直接且普遍。 | Open |
| 3 | [#538 fix(pdf) case-sensitive file references](https://github.com/anthropics/skills/pull/538) | 修复 PDF 技能 SKILL.md 中 8 处引用大小写不匹配问题，保证在大小写敏感文件系统上正常工作。 | 影响 Linux/macOS 用户，暴露了技能元数据的一致性问题。 | Open |
| 4 | [#486 Add ODT skill](https://github.com/anthropics/skills/pull/486) | 新增 OpenDocument 文本创建、模板填充、ODT→HTML 转换技能，覆盖 LibreOffice/ISO 标准格式。 | 社区对跨平台、开源办公文档格式的支持有明确期待，与 #514 形成文档技能系列。 | Open |
| 5 | [#210 Improve frontend-design skill clarity](https://github.com/anthropics/skills/pull/210) | 重写前端设计技能，确保每条指令在单次对话中可执行，避免模糊指导。 | 代表社区对**技能可操作性**的追求，而非仅仅是知识说明。 | Open |
| 6 | [#83 Add skill-quality-analyzer & skill-security-analyzer](https://github.com/anthropics/skills/pull/83) | 添加两个元技能：质量分析器（结构、文档、示例、资源、测试）和安全分析器（潜在风险检查）。 | 社区渴望自检工具来评估技能质量，特别是在官方命名空间下发表的第三方技能。 | Open |
| 7 | [#539 fix(skill-creator) warn on unquoted YAML special chars](https://github.com/anthropics/skills/pull/539) | 在 quick_validate.py 中添加预解析校验，检测 description 字段中未引号的特殊 YAML 字符（如 `:`）。 | 避免静默解析失败，提升技能创建者工具的健壮性。 | Open |
| 8 | [#541 fix(docx) tracked change w:id collision](https://github.com/anthropics/skills/pull/541) | 修复 DOCX 技能在添加修订时因硬编码 ID 与已有书签冲突导致文档损坏。 | 深度 OOXML 兼容性问题，体现了用户对**生产级文档处理**的严格要求。 | Open |

## 2. 社区需求趋势

从 Issues 中可提炼出以下主要方向：

- **安全与信任（#492）**：社区成员发现社区技能冒用 `anthropic/` 命名空间，造成信任边界滥用，强烈要求隔离或官方认证机制。相关讨论持续 19 条评论，2 个 👍。
- **组织级协作（#228）**：用户希望 Skills 能在组织内直接共享，而非通过文件下载 / Slack 传播。该 issue 获得 7 个 👍，是目前兴趣最高的功能请求。
- **核心工具修复（#556、#1169、#1061）**：`run_eval.py` 在 Windows 和 Linux 上均存在 0% 触发率、子进程读取、编码等严重 bug，直接导致技能创建者的优化流程完全失效。相关 issue 合计 15+ 条评论，暴露出官方工具的成熟度不足。
- **新技能类型提案**：
  - [#412] **Agent 治理技能** — 为 AI 代理系统添加策略执行、威胁检测、审计追踪。
  - [#1329] **compact-memory 技能** — 用符号化记号压缩代理长期记忆，减少上下文消耗。
  - [#1175] **SharePoint Online 文档处理** — 结合权限控制与上下文窗口优化。
- **标准化与平台兼容**：包括将 Skills 暴露为 MCP（#16）、与 AWS Bedrock 集成（#29）、以及文档技能去重问题（#189）。

**趋势总结**：社区重心从“创造新技能”逐渐转向“使技能生态安全、可共享、工具链可靠”。

## 3. 高潜力待合并 Skills

以下 PR 评论活跃、功能明确，且尚未合并，有望近期落地：

| PR | 技能 | 价值 | 链接 |
|----|------|------|------|
| [#514](https://github.com/anthropics/skills/pull/514) | document-typography | 直接解决广泛存在的排版缺陷，落地即能提升所有 Claude 输出质量。 | 查看 |
| [#486](https://github.com/anthropics/skills/pull/486) | ODT 读写转换 | 填补 LibreOffice / OpenDocument 生态空白，与 PDF 技能互补。 | 查看 |
| [#83](https://github.com/anthropics/skills/pull/83) | skill-quality-analyzer + skill-security-analyzer | 元技能赋能社区自行审核技能合规性与安全性，呼应 #492 安全需求。 | 查看 |
| [#723](https://github.com/anthropics/skills/pull/723) | testing-patterns | 覆盖测试金字塔（单元、React、集成、E2E），是工程团队的刚需。 | 查看 |
| [#147](https://github.com/anthropics/skills/pull/147) | codebase-inventory-audit | 提供 10 步清理工作流，产出统一状态清单，适合技术债务管理。 | 查看 |
| [#154](https://github.com/anthropics/skills/pull/154) | shodh-memory | 跨对话持久记忆系统，增强 Agent 长期任务能力。 | 查看 |

## 4. Skills 生态洞察

> **当前社区最集中的诉求是：在修复核心技能创建工具链（尤其是评估脚本 0% 触发率与 Windows 兼容性）的基础上，围绕安全性、组织协作和跨平台实用场景，扩展并规范化第三方的技能生态。**

---

好的，这是根据您提供的 GitHub 数据生成的 2026-06-25 Claude Code 社区动态日报。

---

# Claude Code 社区动态日报 | 2026-06-25

## 今日速览

今日发布两个小版本更新，重点修复了 `/rewind` 功能和流式响应中的滚动问题。社区焦点集中在 **ClAudit 安全分类器的大规模误报问题**上，大量用户报告在合法安全运维工作中被错误拦截，已形成多个高热度 Issue 和重复投诉。此外，关于 **Agent Teams（团队协作）** 的模式缺陷以及 **Workflow 子代理（Subagent）并发导致内存溢出** 的问题也引起了广泛关注。

## 版本发布

### v2.1.191
-   **新增 `/rewind` 功能**：支持从执行 `/clear` 命令之前的对话点恢复会话。
-   **修复流式响应滚动问题**：修复了在流式输出期间，用户阅读较早输出内容时，滚动位置自动跳到底部的问题。
-   **修复后台 Agent 停止后复活问题**：修复了从任务面板停止后台 Agent 后，该 Agent 会“复活”并继续运行的 Bug。

### v2.1.190
-   **常规更新**：包含多项错误修复和可靠性改进。

## 社区热点 Issues

1.  **[BUG] Telegram 插件：MCP 通知未注入对话** (Issue #42138)
    -   **重要性**：严重影响 Telegram 插件用户的体验，导致自动化通知无法正常工作。
    -   **社区反应**：评论数最高 (7)，说明该问题影响范围较广，用户反馈积极。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/42138)

2.  **[BUG] Agent Teams：团队领导“指针”错误指向队友** (Issue #64550)
    -   **重要性**：这是一个影响 Agent Teams 核心协作模式的严重 Bug。团队领导会错误地以队友身份执行指令，导致子 Agent 创建失败，完全打乱团队工作流。
    -   **社区反应**：多条评论，开发者正积极排查。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/64550)

3.  **[BUG][harness] ClAudit 自动模式分类器误报（1）** (Issue #70801)
    -   **重要性**：用户 sworrl 报告 ClAudit 将安装合法自动启动程序的请求错误分类并拒绝。
    -   **社区反应**：用户 sworrl 今日集中提交了大量类似 Issue，此条评论数最高 (4)。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70801)

4.  **[BUG][harness] ClAudit 自动模式分类器误报（2）** (Issue #70792)
    -   **重要性**：与前一条类似，用户在进行自动化工具编辑时被错误拦截。这表明 ClAudit 在“自动启动”、“持久化”等行为上判定过于严格。
    -   **社区反应**：用户 sworrl 今日高频提交的系列问题之一。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70792)

5.  **[BUG][cyber] 安全过滤器误报：阻止恶意软件分析** (Issue #70821)
    -   **重要性**：这是一个非常典型的安全误报案例，Claude Code 的安全过滤器错误地将对恶意软件样本的分析（一项标准的防御性工作）识别为攻击行为并阻止。这直接阻碍了网络安全分析师的工作。
    -   **社区反应**：被认为是“假阳性”，用户呼吁降低分类器灵敏度。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70821)

6.  **[BUG][cyber] 安全过滤器误报：阻止云IAM运维** (Issue #70796)
    -   **重要性**：用户在执行云 IAM（身份和访问管理）场景下的脚本时被错误拦截。这暴露了安全分类器对特定技术术语的过度匹配问题。
    -   **社区反应**：多个 Issue 指向同一问题，表明该误报模式具有普遍性。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70796)

7.  **[BUG][aup] AUP / 使用策略误报：阻止合法操作** (Issue #70790)
    -   **重要性**：用户在执行常规的、已授权的云 IAM 工作时，被“使用政策”模块错误拦截。此类误报会严重打击开发者对 Claude Code 自动模式的信任。
    -   **社区反应**：用户明确指出这是“正常、授权范围内的请求”。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70790)

8.  **[Enhancement] 工作流子代理：内存感知节流** (Issue #69033)
    -   **重要性**：用户反馈在执行大量子代理并发任务（如深度研究）时，内存不足导致终端崩溃。当前并发限制基于 CPU 核心数，而非内存。这是一个重要的可靠性增强需求。
    -   **社区反应**：获得点赞，表明复杂任务的用户对此有切实需求。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/69033)

9.  **[BUG][harness] ClAudit 误报：阻止读取自身服务配置** (Issue #70809)
    -   **重要性**：用户在调试部署问题时，通过 SSH 读取自身服务的配置文件被拦截。这暴露了分类器在“授权行为”和“敏感数据”界定上的模糊性。
    -   **社区反应**：用户强调这是“授权范围内的加固工作”。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70809)

10. **[BUG][cyber] 安全过滤器误报：阻止无意义的键盘输入** (Issue #70812)
    -   **重要性**：最令人哭笑不得的误报。用户输入了随机键盘占位符（如“ASDF”），结果被网络安全过滤器拦截。这说明分类器对任何“可疑”内容都非常敏感，甚至是无意义的字符串。
    -   **社区反应**：用户将其标记为荒谬，凸显了安全策略的过度严格。
    -   [查看详情](https://github.com/anthropics/claude-code/issues/70812)

## 重要 PR 进展

由于当日只有 3 个 PR 更新，以下列出全部。

-   **[PR #70634] fix: 处理正常使用中的服务器速率限制** (OPEN)
    -   **功能**：修复了在正常使用场景下服务器返回速率限制错误时的处理逻辑。
    -   **重要性**：提升了 Claude Code 的鲁棒性，减少因 API 限流导致的意外中断。
    -   [查看详情](https://github.com/anthropics/claude-code/pull/70634)

-   **[PR #70633] fix: 处理 Anthropic API 的速率限制头部** (OPEN)
    -   **功能**：增加了对 Anthropic API 返回的特定速率限制头部的处理。
    -   **重要性**：与 #70634 协同工作，提供更细粒度的速率控制。
    -   [查看详情](https://github.com/anthropics/claude-code/pull/70633)

-   **[PR #70582] fix: 修复 llm.py 中接受用户控制 URL 的安全问题** (OPEN)
    -   **功能**：修复了 `plugins/security-guidance/hooks/llm.py` 文件中存在的严重安全漏洞（CRITICAL 级别），该漏洞可能允许攻击者控制 URL 进行注入攻击。
    -   **重要性**：直接修复了高危安全漏洞，对安全插件用户至关重要。
    -   [查看详情](https://github.com/anthropics/claude-code/pull/70582)

## 功能需求趋势

-   **安全分类器优化（首要趋势）**：社区压倒性地关注 **ClAudit / 安全过滤器的误报率**。大量 Issue 表明当前的分类器在网络安全（Cyber）、使用策略（AUP）和自动模式（Auto-mode）三个维度上都过于严格，严重阻碍了合法、授权的安全运维和开发工作。用户急需一种更精准、上下文感知的分类机制。
-   **团队协作与 Agent 行为**：对 **Agent Teams** 模式的反馈持续存在，主要关注内部状态管理（如领导指针错误）。同时，**Workflow 子代理（Subagent）** 的内存管理也成为一个新的明确需求。
-   **平台支持与兼容性**：从 Issue 标签看，社区对 **Linux 平台的支持**和 **Windows 平台**的特定问题都有持续关注。

## 开发者关注点

-   **安全策略误报的挫败感**：今天是开发者痛点非常集中的一天。用户 sworrl 提交了十几个关于 ClAudit 误报的 Issue，涵盖了从恶意软件分析、云 IAM 运维到无意义键盘输入的广泛场景。这表明当前的安全策略不仅不智能，而且产生了严重的噪音，降低了开发者的工作效率和对工具的信任。这可能是 Claude Code 团队近期最需要优先解决的高优问题。
-   **隐式状态管理的困惑**：Agent Teams 中领导指针“粘”在队友身上的 Bug，以及后台 Agent 停止后“复活”的问题，都指向了 Agent 生命周期和状态管理上的隐式、不可预测的行为。开发者需要更透明、可控的状态机制。
-   **复杂工作负载的资源瓶颈**：深层次研究等需要大规模子代理并发任务的场景，遇到了实际的内存瓶颈。开发者希望团队能基于内存（而非仅基于CPU核心数）来实现更智能的并发控制。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 — 2026-06-25

---

## 1. 今日速览

- 过去 24 小时发布了 **rust-v0.142.2** 和 **rust-v0.142.1** 两个稳定补丁，重点优化了 MCP 工具搜索默认开启以及 macOS/Windows 系统代理支持；同时多个 alpha 版本继续迭代。  
- 社区最强烈的反馈集中在 **配额/速率限制异常**：多用户报告模型容量满、“Selected model is at capacity”、Pro 20x 配额被当做 Plus 级别消耗，甚至 5 小时限制在 41 分钟内耗尽。  
- 开发者在 **Windows 沙箱兼容性、插件系统 OAuth 改进、hooks 完整度** 等方面提交了大量 PR，平台兼容性与扩展能力正在加速完善。

---

## 2. 版本发布

### 🎯 rust-v0.142.2  
- **新特性**  
  - MCP 工具现在默认使用工具搜索（tool search），在保持向后兼容的同时提升工具发现能力。（[#29486](https://github.com/openai/codex/issues/29486)）  
  - macOS 认证客户端在启用 `respect_system_proxy` 时，可遵守系统代理、PAC 和 WPAD 设置。（[#26709](https://github.com/openai/codex/issues/26709)）  

### 🎯 rust-v0.142.1  
- **新特性**  
  - 新增可选的 Windows 系统代理支持（PAC、WPAD、静态代理及绕过规则），用于认证流程。（[#26708](https://github.com/openai/codex/pull/26708)）  

### 🔬 Alpha 迭代  
- `rust-v0.143.0-alpha.21` / `.16` / `.15`：持续推动 0.143 主线，包含底层重构与新功能验证。

> 完整变更日志：[Compare 0.142.0...0.142.1](https://github.com/openai/codex/compare/rust-v0.142.0...rust-v0.142.1)

---

## 3. 社区热点 Issues（Top 10）

1. **#29760 - “Selected model is at capacity. Please try a different model.”**  
   - 出现频率极高，Pro 用户使用 gpt-5.4 high 模型时频繁遇到容量已满。社区反应：20 条评论，仅 1 👍，反映大量用户受影响。  
   - 链接：[#29760](https://github.com/openai/codex/issues/29760)

2. **#30008 - 同样问题（CLI + App 均出现）**  
   - 用户同时使用 Codex App 和 CLI 时均提示“容量已满”，Pro 20 用户。16 条评论，0 👍，说明是普遍性故障。  
   - 链接：[#30008](https://github.com/openai/codex/issues/30008)

3. **#29968 - Pro20x 订阅配额被降级为 Plus 级别**  
   - 用户反映订阅用量被错误地按照 Plus 规则计算，导致无法正常使用。10 👍，讨论激烈。  
   - 链接：[#29968](https://github.com/openai/codex/issues/29968)

4. **#29955 - 配额瞬间耗尽：1 条消息消耗 100 积分，5h 限制重置后立即为 0%**  
   - 典型配额 bug，用户刚过重置点即被限制，严重影响日常使用。15 条评论。  
   - 链接：[#29955](https://github.com/openai/codex/issues/29955)

5. **#30002 - 服务器端配额会计错误：5h 限制在 ~41 分钟内重新触发**  
   - 用户提供详细数据：同一账号在 41 分钟内只实际消耗 ~1.35M tokens 就被限制，而正常窗口能用到 ~156M tokens。12 条评论，3 👍。  
   - 链接：[#30002](https://github.com/openai/codex/issues/30002)

6. **#21753 - [增强] 实现与 Claude Code 完全一致的 Hooks 功能**  
   - 长期跟踪 issue，社区渴望更完整的自动化表面（生命周期、事件名等）。20 条评论，18 👍。  
   - 链接：[#21753](https://github.com/openai/codex/issues/21753)

7. **#16817 - [Bug] Mac 桌面 App 重启后无法加载已有线程**  
   - 老问题仍未被完全解决（4 月提出），用户重启后所有已打开的线程消失。13 条评论，7 👍。  
   - 链接：[#16817](https://github.com/openai/codex/issues/16817)

8. **#29782 - Windows 沙箱每次都弹出 COM+ 注册表错误**  
   - 使用 `apply_patch` 时循环弹窗，影响 Windows 用户的核心体验。7 条评论，0 👍，但属于严重的平台 bug。  
   - 链接：[#29782](https://github.com/openai/codex/issues/29782)

9. **#29963 - 配额消耗有严重 bug（Pro 20x 用户）**  
   - 用户提供截图表明限额异常消耗，并指出在 macOs、Ubuntu 上均有出现。5 条评论，5 👍，说明严重影响付费用户。  
   - 链接：[#29963](https://github.com/openai/codex/issues/29963)

10. **#30007 - [增强] 改进 TUI 中 Markdown 渲染（宽混合语言、数学公式、表格）**  
    - 功能请求，用户希望终端内能更好显示多语言代码块、数学公式和表格。2 条评论，0 👍，但反映了对 CLI 体验的持续需求。  
    - 链接：[#30007](https://github.com/openai/codex/issues/30007)

---

## 4. 重要 PR 进展（Top 10）

1. **#29923 - 支持外部时钟休眠（sleep/wake）**  
   - 新增实验性 `currentTime/sleep` 通知与 `currentTime/wake` 请求，使外部时钟可精确控制 Codex 的等待行为。  
   - 链接：[#29923](https://github.com/openai/codex/pull/29923)

2. **#29845 - 为 Windows 启动器注入显式应用路径**  
   - 重构 Windows 统一执行的可执行文件解析，引入 `WindowsProcessLaunch` 类型携带路径，为后续功能铺路。  
   - 链接：[#29845](https://github.com/openai/codex/pull/29845)

3. **#30017 - 核心：从步骤上下文刷新 turn diff 根目录**  
   - 修复当延迟环境（deferred executor）附加后，diff 跟踪器仍使用旧环境根目录的问题，确保文件路径正确。  
   - 链接：[#30017](https://github.com/openai/codex/pull/30017)

4. **#30016 - 核心：子代理继承当前步骤的环境**  
   - 使由后续采样请求产生的子代理能正确继承当前环境（而不是保存的旧快照），解决延迟 executor 场景下的环境错乱。  
   - 链接：[#30016](https://github.com/openai/codex/pull/30016)

5. **#30020 - 核心：将步骤环境传给 turn 输入扩展**  
   - 允许 turn 输入扩展获取最新的步骤环境，避免因 `TurnContext` 冻结而错过延迟环境更新。  
   - 链接：[#30020](https://github.com/openai/codex/pull/30020)

6. **#29917 - exec-server：并发处理后初始化请求**  
   - 允许独立 RPC 在长轮询期间并行处理，同时保持 `initialize` / `initialized` 的顺序约束，提升吞吐量。  
   - 链接：[#29917](https://github.com/openai/codex/pull/29917)

7. **#29935 - 按线程来源分析 app-server 事件**  
   - 避免 Work 线程活动被错误归因到 Desktop 连接，为管理面板提供准确的线程级分析数据。  
   - 链接：[#29935](https://github.com/openai/codex/pull/29935)

8. **#29890 - 添加插件列表强制刷新功能**  
   - 当外部 app 重新认证后，允许用户手动刷新插件缓存，无需重启整个 Codex。  
   - 链接：[#29890](https://github.com/openai/codex/pull/29890)

9. **#28406 - 复用回合内的工具路由器**  
   - 每个 turn 内延迟构建一次 `ToolRouter`，后续采样请求复用同一 Arc，保持工具 schema 一致，同时允许 MCP 变更在下一 turn 生效。  
   - 链接：[#28406](https://github.com/openai/codex/pull/28406)

10. **#28407 - 避免阻塞在可选 MCP 启动上**  
    - 对于可选 MCP 服务器，在启动时忽略它们（若无缓存快照），不阻塞工具列表返回；必要服务器仍保持阻塞契约。  
    - 链接：[#28407](https://github.com/openai/codex/pull/28407)

---

## 5. 功能需求趋势

从过去 24 小时的 Issue 和 PR 中可以提炼出以下社区最关注的方向：

- **配额与速率限制的透明性与修复**  
  大量 Issue 指向配额计算错误、模型容量不足、5h 限制异常快速消耗，是当前最大的痛点。社区要求开源配额机制或提供更清晰的仪表盘。

- **MCP / 插件系统的成熟度**  
  多个 PR 完善了 MCP 的 OAuth 支持、工具路由复用、可选服务器非阻塞启动；Issue 则要求更好的插件安装/授权本地化、插件列表强制刷新。社区期待插件系统成为“一等公民”。

- **Windows 平台兼容性**  
  Windows 用户报告沙箱 COM+ 错误、Git 文件夹创建循环、executable PATH 访问被拒、更新失败等问题。Windows 支持正成为开发者的重点诉求。

- **Hooks 自动化表面**  
  Issue #21753 持续获得高赞，社区希望 Codex 提供与 Claude Code 完全对等的 hooks（生命周期事件、自定义操作），以构建自动化工作流。

- **TUI / CLI 体验改进**  
  包括 Markdown 渲染增强（数学公式、表格）、宽混合语言支持、系统消息级别（info/notice）等。开发者追求更高效的终端交互。

- **i18n / 本地化**  
  Issue #30025 提出中文简体等语言支持，反映全球用户对非英语界面的需求。

---

## 6. 开发者关注点

综合今日反馈，开发者最集中反映的痛点包括：

- **配额异常频繁**：Pro 用户无法正常使用，5h 重置后 41 分钟即被限；部分用户积分瞬间归零，严重影响信任度。
- **模型容量不足**：gpt-5.4 high 等热门模型长期显示“at capacity”，迫使开发者降级或重试。
- **Windows 沙箱的稳定性**：COM+ 错误、CreateProcessAsUser 失败、PATH 隔离等问题导致常用命令无法执行。
- **线程/会话状态丢失**：Mac 桌面 App 重启后线程消失；Windows 上线程切换缓慢；上下文窗口溢出后线程不可恢复。
- **插件管理不便**：安装的插件在 TUI 中不显示、远程插件列表覆盖本地、OAuth 流程复杂。
- **更新失败**：Windows 原生安装的 CLI 无法通过 `codex update` 完成更新。

这些高频问题直接关系到付费用户的留存率，OpenAI 应优先处理配额计算和模型容量分配。

---

*日报基于 GitHub 数据自动生成，数据截止时间 2026-06-25 23:59 UTC。*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，作为专注于 AI 开发工具的技术分析师，以下是为您生成的 2026 年 6 月 25 日 Gemini CLI 社区动态日报。

---

# Gemini CLI 社区动态日报 | 2026-06-25

## 今日速览
今日社区动态聚焦于安全性与 Agent 核心行为的修复。最新的夜间版本重点修复了一个严重的路径遍历漏洞。社区反馈中，关于 Agent（特别是子 Agent）的执行逻辑问题依然是讨论热点，同时围绕内存系统（Auto Memory）的优化和安全审计也正在进行中。此外，一个关于文档更新的 PR 指出了即将到来的服务变更，值得所有用户关注。

## 版本发布

### v0.49.0-nightly.20260625.gd845bc5d4
- **关键修复**：修复了在安装技能（Skill）时可能存在的路径遍历漏洞，提升了 CLI 的安全性。
- **功能修复**：修复了待处理工具（pending tools）和信任覆盖（trust overrides）相关的问题。
- **CI 优化**：对持续集成流程进行了优化。
- **发布链接**: [查看发布详情](https://github.com/google-gemini/gemini-cli/releases/tag/v0.49.0-nightly.20260625.gd845bc5d4)

## 社区热点 Issues

1.  **[#22323] Subagent recovery after MAX_TURNS is reported as GOAL success, hiding interruption**
    - **重要性**: **高。** 这是一个严重的 Agent 行为 Bug。当子代理因达到最大轮次（MAX_TURNS）而被中断时，系统错误地报告为“成功”，导致用户无法发现任务实际上是失败的。这直接影响了用户对 Agent 工作流的信任。
    - **社区反应**: 8条评论，社区成员 `matei-anghel` 详细报告了该问题，目前被标记为需要重新测试。
    - **链接**: [Issue #22323](https://github.com/google-gemini/gemini-cli/issues/22323)

2.  **[#24353] Robust component level evaluations**
    - **重要性**: **高。** 这是一个“史诗级”议题，旨在建立更强大的组件级评估体系。这对于提升 Agent 的稳定性和可预测性至关重要，是确保 Agent 质量的基础设施建设。
    - **社区反应**: 7条评论，讨论涉及如何设计和执行更细粒度的测试。被标记为客户问题。
    - **链接**: [Issue #24353](https://github.com/google-gemini/gemini-cli/issues/24353)

3.  **[#22745] Assess the impact of AST-aware file reads, search, and mapping**
    - **重要性**: **中高。** 该议题探讨利用抽象语法树（AST）来优化文件读取、搜索和代码映射。如果能实现，将极大提升 Agent 对代码的理解效率和准确性，减少不必要的 Token 消耗和工具调用。
    - **社区反应**: 7条评论，讨论了 AST 感知工具对 Agent 能力提升的潜力。
    - **链接**: [Issue #22745](https://github.com/google-gemini/gemini-cli/issues/22745)

4.  **[#21409] Generalist agent hangs**
    - **重要性**: **高。** 一个经典且负反馈较多的问题。当 CLI 将任务委托给通用 Agent 时，会无限期挂起，直到用户取消。这严重影响了用户体验，尤其是对于需要复杂推理的任务。
    - **社区反应**: 7条评论，获得了 8 个赞，说明受该问题困扰的用户很多。解决方法之一是明确指示模型不要使用子代理。
    - **链接**: [Issue #21409](https://github.com/google-gemini/gemini-cli/issues/21409)

5.  **[#26525] Add deterministic redaction and reduce Auto Memory logging**
    - **重要性**: **高。** 这是一个与安全和隐私相关的关键议题。Auto Memory 功能在处理本地对话记录时，可能存在敏感信息泄露的风险。提议增加确定性的信息脱敏机制，并减少日志记录，是提升安全性的重要举措。
    - **社区反应**: 5条评论，讨论集中在如何更安全地处理转录内容。
    - **链接**: [Issue #26525](https://github.com/google-gemini/gemini-cli/issues/26525)

6.  **[#25166] Shell command execution gets stuck with "Waiting input" after command completes**
    - **重要性**: **高。** 一个非常影响日常使用的 Bug。在 Shell 命令执行完成后，CLI 卡在“等待输入”状态，导致流程中断。这对于自动化工作流是灾难性的。
    - **社区反应**: 4条评论，获得了 3 个赞，用户 `rnett` 报告该问题频繁出现。
    - **链接**: [Issue #25166](https://github.com/google-gemini/gemini-cli/issues/25166)

7.  **[#21983] browser subagent fails in wayland**
    - **重要性**: **中。** Wayland 用户遇到的特定问题。浏览器子代理在 Wayland 显示服务器下会失败，限制了部分用户的可用性。
    - **社区反应**: 4条评论，该问题已被标记为需要重新测试。
    - **链接**: [Issue #21983](https://github.com/google-gemini/gemini-cli/issues/21983)

8.  **[#26522] Stop Auto Memory from retrying low-signal sessions indefinitely**
    - **重要性**: **中。** 优化 Auto Memory 功能。当前，对于“低价值”的对话，系统会无限重试，浪费计算资源和 Token。提议优化重试逻辑，提升系统效率。
    - **社区反应**: 5条评论，讨论如何定义“低信号”会话并停止无意义的重试。
    - **链接**: [Issue #26522](https://github.com/google-gemini/gemini-cli/issues/26522)

9.  **[#20079] `~/.gemini/agents/filename.md` is not recognized as an agent if filename.md is a symlink**
    - **重要性**: **中。** 用户工作流相关的问题。无法使用符号链接来管理自定义 Agent 配置文件，这限制了用户的知识管理灵活性。
    - **社区反应**: 4条评论，该问题被标记为需要更多信息。
    - **链接**: [Issue #20079](https://github.com/google-gemini/gemini-cli/issues/20079)

10. **[#24246] Gemini CLI encounters 400 error with > 128 tools**
    - **重要性**: **中。** 限制了 Agent 能力的可扩展性。当可用工具超过128个时（例如，用户自定义技能很多或集成了大量MCP服务器），CLI会报错。社区期望 Agent 能更智能地处理大量工具集。
    - **社区反应**: 3条评论，该问题被标记为需要更多信息。
    - **链接**: [Issue #24246](https://github.com/google-gemini/gemini-cli/issues/24246)

## 重要 PR 进展

1.  **[#27971] fix(core): strip thoughts from scrubbed history turns and resolve thought leakage**
    - **重要性**: **高。** 修复了“思维泄漏”问题。此问题导致模型的内部推理过程（Thought）被错误地泄露到对话历史中，进而混淆模型引发无限循环。该修复对于提升 Agent 的长期对话稳定性至关重要。
    - **链接**: [PR #27971](https://github.com/google-gemini/gemini-cli/pull/27971)

2.  **[#28053] fix(core-tools): resolve defensive path resolution for at-reference files and fix macOS tests**
    - **重要性**: **高。** 修复了一个生产环境中的关键 Bug。当模型传递以 `@` 开头的文件路径时（如 `@policies/new-policies.txt`），文件系统工具会报错“文件未找到”。此 PR 提供了完整的路径解析防御方案。
    - **链接**: [PR #28053](https://github.com/google-gemini/gemini-cli/pull/28053)

3.  **[#27966] fix(security): enforce case-insensitive sensitive path blocklist and vscode hitl**
    - **重要性**: **高。** 安全相关。修复了对敏感目录（如 `.git`, `.env`）路径大小写不敏感的绕过漏洞，并增强了与 VS Code 的集成。这是对安全性的重要加固。
    - **链接**: [PR #27966](https://github.com/google-gemini/gemini-cli/pull/27966)

4.  **[#27996] fix(core): decode response body using charset from Content-Type header**
    - **重要性**: **中高。** 修复了 `web-fetch` 工具总是使用 UTF-8 解码响应体的问题，导致部分非 UTF-8 编码（如中文GBK）的网页内容出现乱码。此修复提升了 CLI 在全球化场景下的可用性。
    - **链接**: [PR #27996](https://github.com/google-gemini/gemini-cli/pull/27996)

5.  **[#27994] fix(core): insert skill/agent content literally in system prompt substitutions**
    - **重要性**: **中高。** 修复了系统提示词中技能（Skill）和子代理（Subagent）内容替换时的字符串问题，避免了内容中的特殊字符被错误解析。确保了自定义技能能正确注入到提示词中。
    - **链接**: [PR #27994](https://github.com/google-gemini/gemini-cli/pull/27994)

6.  **[#27986] feat(acp): report cached and thought tokens in PromptResponse.usage**
    - **重要性**: **中。** 新功能。当 Gemini CLI 作为 ACP 服务器运行时，现在可以在 Token 使用报告中上报缓存（cached）和推理（thought）Token。这对于需要精确计算成本的 ACP 客户端来说非常有价值。
    - **链接**: [PR #27986](https://github.com/google-gemini/gemini-cli/pull/27986)

7.  **[#28094] fix(a2a-server): deep-merge user and workspace settings**
    - **重要性**: **中。** 修复了 A2A 服务器中用户设置与工作区设置合并的问题。原先的浅拷贝会导致嵌套配置（如工具、遥测设置）被意外覆盖。改为深拷贝后，可以更灵活地隔离和管理不同层级的配置。
    - **链接**: [PR #28094](https://github.com/google-gemini/gemini-cli/pull/28094)

8.  **[#27979] Wrap read_mcp_resource output with wrapUntrusted() for consistency with mcp-tool**
    - **重要性**: **中。** 安全一致性修复。确保从 MCP 服务器读取的资源（`read_mcp_resource`）与调用 MCP 工具（`mcp-tool`）一样，返回的内容都被标记为“不可信”。防止潜在的提示注入风险。
    - **链接**: [PR #27979](https://github.com/google-gemini/gemini-cli/pull/27979)

9.  **[#28015] feat(caretaker): implement Cloud Run webhook ingestion service**
    - **重要性**: **中。** 新功能。为实现“看护者代理”（Caretaker Agent）而构建的云服务。该服务可以接收 GitHub Webhook，验证签名，并将 issue 元数据存储起来，为更强大的自动化运维自动化打下基础。
    - **链接**: [PR #28015](https://github.com/google-gemini/gemini-cli/pull/28015)

10. **[#27636] perf: optimize VirtualizedList and fix click handling**
    - **重要性**: **中。** 性能优化。优化了终端的虚拟列表渲染性能，并修复了点击处理问题。这直接关系到 CLI 界面在处理大量对话历史或日志时的流畅度和响应速度。
    - **链接**: [PR #27636](https://github.com/google-gemini/gemini-cli/pull/27636)

## 功能需求趋势
- **Agent 行为可靠性与可预测性**: 社区强烈需求 Agent 能更稳定地执行任务，包括正确处理失败（如 `#22323`）、避免死锁（如 `#25166`）和更智能地管理可用工具（如 `#24246`）。
- **安全与隐私加固**: 随着 Agent 能力增强，安全问题愈发突出。社区关注点包括路径遍历防护（已修复）、敏感信息脱敏（如 `#26525`）、信任覆盖机制以及防止提示注入。
- **系统集成与兼容性**: 用户期望 CLI 能与各种系统顺畅集成，包括不同操作系统（如 Wayland 支持 `#21983`）、终端模拟器以及外部编辑器。
- **代码理解能力提升**: 通过引入 AST 感知的文件读写和代码搜索来提升 Agent 对代码库的理解深度和效率，是一个持续关注的趋势。
- **评估与质量体系**: 社区和开发者都希望建立更系统化的评估标准（如 `#24353`），以确保 Agent 的各项功能表现符合预期。

## 开发者关注点
- **Autonomous Agent 执行错误**: 开发者普遍反馈 Agent 在执行复杂任务时（尤其是涉及多个子代理）行为异常，如错误报告成功（#22323）、无响应挂起（#21409）和卡在 Shell 命令（#25166）。这是最突出的痛点。
- **文件和路径处理问题**: 许多 Bug 都与文件系统交互有关，包括符号链接不被识别（#20079）、路径大小写绕过安全限制（#27966）、`@` 路径解析失败（#28053）等，表明这部分功能需要进一步加固。
- **系统集成兼容性**: 在特定环境下（如 Wayland、macOS）的测试失败和功能失效是常见的开发者痛点。
- **安全与权限控制**: 开发者对内省和记录功能（如 Auto Memory）可能带来的安全风险非常敏感，希望有更严格的控制和脱敏机制。
- **性能与内存优化**: 在处理长对话或大数据量时，终端的渲染性能（如 `#27636` 的 VirtualizedList 优化）和内存日志管理是持续的关注点。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-06-25

---

## 今日速览

- **v1.0.65 发布**：修复了带斜杠前缀字符串参数误触文件系统权限提示的问题，`/cd` 命令现在会持久化工作目录。
- **模型与会话问题持续发酵**：多个用户反馈恢复会话后模型选择列表为空或显示“未认证”，社区讨论热烈（#3596、#3913、#3680）。
- **插件功能迎来改进诉求**：用户强烈要求支持技能子文件夹（#1632，获21赞）以及`preToolUse`钩子的静默重写能力（#2643）。

---

## 版本发布

### v1.0.65（2026-06-24）

主要变更：

- `/cd` 命令现在会持久化工作目录，恢复会话时会自动回到之前所在的目录，并在新目录中发现自定义 agent。
- 修复：带斜杠前缀的字符串参数（如 `--body "/azp run"`）不再触发虚假的文件系统权限询问。
- 全屏时间线保持功能（`Fullscreen timeline stays anc`，原文截断，推测为保持锚点位置）。

---

## 社区热点 Issues（10 条）

### 1. [Feature] 提供列出所有支持模型的方式  
**#700** | 创建: 2025-12-02 | 评论: 14 | 👍: 4  
用户希望有类似 `copilot --list-models` 的命令，能一次查看所有可用模型及其倍率信息。这是基础功能需求，缺乏官方列表让模型选择变得困难。  
🔗 [github/copilot-cli Issue #700](https://github.com/github/copilot-cli/issues/700)

### 2. [Bug] `preToolUse` 静默重写命令时仍弹出确认对话框  
**#2643** | 创建: 2026-04-11 | 评论: 12 | 👍: 2  
插件开发者在使用 `updatedInput` + `permissionDecision: allow` 后，每次重写命令仍会被交互式确认打断，无法实现静默重写。社区希望提供真正的“允许”行为。  
🔗 [github/copilot-cli Issue #2643](https://github.com/github/copilot-cli/issues/2643)

### 3. [Feature] 支持技能子文件夹以更好组织  
**#1632** | 创建: 2026-02-23 | 评论: 9 | 👍: 21  
用户拥有超过10个技能，扁平结构难以管理。当前系统不允许将技能放入子文件夹，这是目前点赞最高的功能请求之一。  
🔗 [github/copilot-cli Issue #1632](https://github.com/github/copilot-cli/issues/1632)

### 4. [Bug] 恢复会话后无法加载模型列表：未认证  
**#3596** | 创建: 2026-05-31 | 评论: 7 | 👍: 11  
当通过 `--resume` 恢复特定会话时，`/model` 命令报错 “Not authenticated”，但新会话正常。大量用户受影响，且与 #3680、#3913 高度关联。  
🔗 [github/copilot-cli Issue #3596](https://github.com/github/copilot-cli/issues/3596)

### 5. [Bug] 6月16日服务中断后所有模型显示为“Blocked/Disabled”  
**#3832** | 创建: 2026-06-17 | 评论: 6 | 👍: 13  
6月16日 GitHub Copilot 短暂宕机后，模型选择界面中所有模型均被标记为不可用，无法选择或开始新会话。这暴露了状态恢复逻辑缺陷。  
🔗 [github/copilot-cli Issue #3832](https://github.com/github/copilot-cli/issues/3832)

### 6. [Bug] 模型倍率计算错误：请求消耗 5% 而非 2%  
**#3881** | 创建: 2026-06-22 | 评论: 3 | 👍: 0  
用户选择 Claude Sonnet 4.5（6倍消耗），本应扣除 2% 配额，实际扣除了 5%。用户要求返还 3% 的配额，并怀疑计费逻辑有误。  
🔗 [github/copilot-cli Issue #3881](https://github.com/github/copilot-cli/issues/3881)

### 7. [Bug] 恢复会话后模型选择列表为空  
**#3913** | 创建: 2026-06-24 | 评论: 3 | 👍: 1  
与 #3596 类似，但发生在 v1.0.64，所有可用模型出现在“封锁列表”中，无法选择。此问题在 v1.0.65 是否修复尚待确认。  
🔗 [github/copilot-cli Issue #3913](https://github.com/github/copilot-cli/issues/3913)

### 8. [Bug] Escape 应取消当前任务并聚焦已排队的提示，而非丢弃  
**#3692** | 创建: 2026-06-05 | 评论: 2 | 👍: 1  
当前 Escape 会取消所有排队消息，导致已输入的后续提示丢失。用户期望它只取消当前运行中的任务，然后自动执行队列中的下一条消息。  
🔗 [github/copilot-cli Issue #3692](https://github.com/github/copilot-cli/issues/3692)

### 9. [Feature] 可配置快捷键（支持快速切换模型）  
**#2419** | 创建: 2026-03-31 | 评论: 2 | 👍: 5  
用户希望将 F1/F2/F3 等键绑定到斜杠命令，尤其用于快速切换模型，避免每次都要输入 `/model` 和记忆模型 ID。同类型的 #1729 也有 5 赞。  
🔗 [github/copilot-cli Issue #2419](https://github.com/github/copilot-cli/issues/2419)

### 10. [Bug] 技能 `argument-hint` 格式验证过于严格  
**#3929** | 创建: 2026-06-25 | 评论: 1 | 👍: 0  
用户在技能 md 文件中使用 `argument-hint: [regression directory]`，CLI 报错要求必须是字符串。但根据 VS Code 规范，数组格式本应合法。  
🔗 [github/copilot-cli Issue #3929](https://github.com/github/copilot-cli/issues/3929)

---

## 重要 PR 进展（仅 2 条，全部列出）

### 1. [Open] 添加 .gitignore 和设置配置  
**#3928** | 作者: tpsaint | 创建: 2026-06-25 | 状态: Open  
旨在为项目仓库添加缺失的 `.gitignore` 和配置文件，提升开发规范性。  
🔗 [github/copilot-cli PR #3928](https://github.com/github/copilot-cli/pull/3928)

### 2. [Closed] 使用 GitHub Agentic Workflows 自动化 Issue 分类  
**#2587** | 作者: andyfeller | 创建: 2026-04-08 | 合并: 2026-06-24 | 状态: Closed  
引入 AI 驱动的自动标签工作流，在 issue 创建/重新打开时自动应用 `area:` 和 `triage` 标签，提升社区治理效率。  
🔗 [github/copilot-cli PR #2587](https://github.com/github/copilot-cli/pull/2587)

---

## 功能需求趋势

| 需求方向 | 相关 Issue 示例 | 热度 |
|---------|---------------|------|
| **模型管理** | 列出模型（#700）、快速切换（#2419）、倍率显示 | ⭐⭐⭐⭐⭐ |
| **会话持久化改进** | 恢复后认证状态丢失（#3596、#3913、#3680） | ⭐⭐⭐⭐⭐ |
| **插件/技能系统** | 子文件夹支持（#1632）、静默重写（#2643）、参数格式验证（#3929） | ⭐⭐⭐⭐ |
| **可配置键绑定** | 全局快捷键（#2419、#1729）、Escape 语义（#3692） | ⭐⭐⭐ |
| **企业 & 代理** | 企业服务器配置（#3909）、Kerberos 代理（#523）、SDK 代理（#2978） | ⭐⭐⭐ |
| **主题/可访问性** | 细粒度元素主题（#2123） | ⭐⭐ |

---

## 开发者关注点（痛点与高频需求）

1. **会话恢复后的认证混乱**：多个 issue 指出恢复旧会话后 `/model` 操作失败，而新会话正常。用户期望恢复会话后能保持认证状态，无需重新登录。
2. **模型选择界面不透明**：缺少 `--list-models` 命令，且当前 `/model` 在恢复会话时空列表或封锁列表，严重影响体验。
3. **配额计算不可信**：倍率模型的实际消耗与预期不一致（#3881），用户希望有官方解释或日志来验证。
4. **插件静默能力不足**：`preToolUse` 钩子即使显式设置 `permissionDecision: allow` 仍会弹出确认；技能子文件夹的组织诉求强烈。
5. **Linux 兼容性**：AppImage 泄漏 `LD_LIBRARY_PATH` 导致 git HTTPS 失败（#3925），影响企业用户的部署。
6. **快捷键与交互细节**：Windows 上 Ctrl+Enter 不执行队列而是换行（#3760）；Escape 行为不合直觉；Markdown 渲染器将连字符错误当作删除线（#3920）。

---

*数据时间范围：2026-06-24 ~ 2026-06-25 | 数据来源：[github.com/github/copilot-cli](https://github.com/github/copilot-cli)*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 | 2026-06-25

**数据来源**: [MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)  
**统计时段**: 2026-06-24 UTC 更新内容

---

## 今日速览

- 昨日无新版本发布，社区活跃度集中在 **4 个 Issue** 和 **2 个已合并 PR**。  
- 长期遗留的 **循环读取 bug（#640）** 在更新后仍有 14 条评论，社区继续关注稳定性问题。  
- 新提出的 **上下文压缩浪费 Token 的增强建议（#2472）** 以及已修复的 **MCP 子代理配置传播问题（#1942）** 成为开发者关注重点。

---

## 版本发布

（无新版本发布）

---

## 社区热点 Issues

### 1. [#640] [bug] Kimi CLI stuck in reading one file again and again and stuck in a loop  
- **作者**: isbafatima90-arch | **创建**: 2026-01-19 | **更新**: 2026-06-24 | **评论**: 14 | 👍: 1  
- **摘要**: 用户在 Linux 上使用自定义 Anthropic endpoint 和 `mimo-v2-flash` 模型时，Kimi CLI 反复读取同一个文件并进入死循环。  
- **为什么重要**: 该问题已存在近 6 个月，严重阻塞自定义模型用户的工作流程，社区期望官方给出根本性修复。  
- **链接**: [Issue #640](https://github.com/MoonshotAI/kimi-cli/issues/640)

### 2. [#2473] [CLOSED] [bug] web bug  
- **作者**: DCY501 | **创建**: 2026-06-24 | **更新**: 2026-06-24 | **评论**: 0 | 👍: 0  
- **摘要**: 使用 `/web` 指令时出现报错，影响 Kimi Code CLI 的 Web 模式使用。  
- **为什么重要**: 尽管已关闭（可能快速热修复），但反映 `web` 指令仍有可靠性问题，尤其是在 `kimi-for-coding,k2.7` 模型上。  
- **链接**: [Issue #2473](https://github.com/MoonshotAI/kimi-cli/issues/2473)

### 3. [#2469] [CLOSED] [bug] `kimi web` starts MCP servers from the CLI installation directory, breaking workspace-relative MCP tools  
- **作者**: Zehee | **创建**: 2026-06-22 | **更新**: 2026-06-24 | **评论**: 0 | 👍: 0  
- **摘要**: 启动 `kimi web` 时，MCP 服务器会从 CLI 安装目录启动，而非工作区目录，导致工作区相对路径的 MCP 工具失效。  
- **为什么重要**: 该问题影响多项目开发场景，已修复并关闭，是 MCP 集成质量的重要改进。  
- **链接**: [Issue #2469](https://github.com/MoonshotAI/kimi-cli/issues/2469)

### 4. [#2472] [OPEN] [enhancement] Context compaction reloads system prompt and project instructions, wasting ~20k tokens  
- **作者**: 865x44 | **创建**: 2026-06-24 | **更新**: 2026-06-24 | **评论**: 0 | 👍: 0  
- **摘要**: 上下文压缩（Context compaction）触发后，系统提示和项目级指令（如 `AGENTS.md`、技能、环境上下文）会从零重新加载，每次浪费约 20k Token。  
- **为什么重要**: Token 浪费直接影响用户成本和上下文窗口利用率，该增强建议获得开发者共鸣，可能需要缓存或增量更新机制。  
- **链接**: [Issue #2472](https://github.com/MoonshotAI/kimi-cli/issues/2472)

---

## 重要 PR 进展

### 1. [#1942] [CLOSED] fix(mcp): propagate MCP configs to subagents and resume immediately  
- **作者**: msenol | **创建**: 2026-04-20 | **更新**: 2026-06-24 | **评论**: - | 👍: 0  
- **摘要**: 修复两个连带问题：子代理（explore、coder、plan）从未收到 MCP 配置导致工具不可用；恢复会话时 MCP 配置丢失。  
- **意义**: 该 PR 解除了多代理协作场景下的 MCP 功能瓶颈，提升 MCP 生态的可靠性。  
- **链接**: [PR #1942](https://github.com/MoonshotAI/kimi-cli/pull/1942)

### 2. [#1377] [CLOSED] feat: add vim-style j/k keyboard navigation for approval and question…  
- **作者**: IAMLEIzZ | **创建**: 2026-03-09 | **更新**: 2026-06-24 | **评论**: - | 👍: 0  
- **摘要**: 为交互式审批和问答界面添加 Vim 风格的 `j`/`k` 键盘导航。  
- **意义**: 提升 CLI 操作效率，满足 Vim 用户习惯，增强终端体验。  
- **链接**: [PR #1377](https://github.com/MoonshotAI/kimi-cli/pull/1377)

---

## 功能需求趋势

综合当日活跃的 4 个 Issue 和 2 个 PR，社区最关注的功能方向如下：

| 方向 | 表现 |
|------|------|
| **稳定性和 Bug 修复** | 循环读取（#640）、web 指令报错（#2473）持续被反馈，用户对自定义模型端支持要求高 |
| **Token 成本与上下文优化** | #2472 提出上下文压缩浪费 Token，建议实现增量/缓存机制，降低重复加载 |
| **MCP 生态完善** | #2469（路径问题已修复）、#1942（子代理和恢复支持）表明社区对 MCP 工具链的可靠性寄予厚望 |
| **终端交互改进** | #1377 合并的 vim 风格导航，反映 CLI 用户对高效键盘操作的偏好 |

---

## 开发者关注点

- **自定义模型端点稳定性**（#640）：用户使用 `config.toml` 配置非官方端点时，频繁遇到死循环，官方需要提供更健壮的错误处理和重试逻辑。
- **Token 浪费的透明化**（#2472）：上下文压缩机制应提供日志或可配置选项，避免重复加载系统提示；社区期待性能报告或缓存策略。
- **MCP 配置一致性**（#2469、#1942）：工作区相对路径、子代理继承、会话恢复三个方面均出现过问题，开发者希望 MCP 配置成为核心功能并得到充分测试。
- **Web 模式的可用性**（#2473）：尽管已关闭，但快速修复后仍需验证多模型兼容性以及 `kimi web` 与本地 MCP 的协作。

---

*本文档基于 GitHub 公开数据自动生成，仅供参考。*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-06-25

## 今日速览
昨日发布的 **v1.17.10** 引入了 MCP 增强和 `--mini` CLI 模式，但社区迅速发现该版本在 Windows 上因底层 Bun 运行时导致大量段错误（segfault），用户被迫降级至 v1.17.9，并暴露了 SQLite 模式兼容问题。与此同时，社区提交了大量修复 PR，涵盖 IME 输入崩溃、空斜杠命令、子代理中断等多个痛点。

---

## 版本发布

### v1.17.10
- **核心改进**  
  - 添加 MCP Server 指令到会话上下文（@Arcadi4）  
  - 支持 OpenCode 托管提供商集成  
  - 新增 MCP 资源模板列表与资源读取工具  
  - 新增 `--mini` CLI 模式（轻量交互）  
- **Bug 修复**  
  - 隐藏 MCP 资源模板工具（某些条件下）

> ⚠️ 该版本因 Windows 段错误问题引发大量反馈，详情见下方社区热点。

---

## 社区热点 Issues（10 条）

### 1. #33742 – OpenCode v1.17.10 在 Windows 上因 Bun 段错误崩溃；v1.17.9 稳定
- **评论 19 | 👍 26**  
- **摘要**：升级到 v1.17.10 后 Windows 上频繁出现 Bun 原生段错误，降级后恢复正常，怀疑回归。  
- **链接**：https://github.com/anomalyco/opencode/issues/33742

### 2. #33773 – v1.17.10 Bun segfault + 降级后 SQLite 模式损坏（缺少 replacement_seq 列）
- **评论 7 | 👍 15**  
- **摘要**：Windows 用户同时遇到段错误和降级后数据库 schema 不兼容问题，`replacement_seq` 列缺失导致工具无法启动。  
- **链接**：https://github.com/anomalyco/opencode/issues/33773

### 3. #19081 – 回放时 assistant 消息中的 reasoning_content 被剥离，导致本地推理 KV 缓存失效
- **评论 15 | 👍 21**  
- **摘要**：助手回复中的思考块（thinking tokens）在后续回放时被静默移除，破坏本地推理的 KV 缓存机制，影响长对话性能。  
- **链接**：https://github.com/anomalyco/opencode/issues/19081

### 4. #33767 – Windows 11 终端中使用 OpenCode 时 Bun 段错误
- **评论 12 | 👍 9**  
- **摘要**：终端中突发段错误，乱码循环打印，需强制关闭终端。  
- **链接**：https://github.com/anomalyco/opencode/issues/33767

### 5. #33752 – Agent 回复选项时 Bun 崩溃
- **评论 12 | 👍 0**  
- **摘要**：Agent 提供选择项时频繁触发错误，用户无法定位修复。  
- **链接**：https://github.com/anomalyco/opencode/issues/33752

### 6. #8785 – Bun has crashed（旧 Issue 仍被引用）
- **评论 40 | 👍 7**  
- **摘要**：跨版本的经典 Windows 崩溃报告，本次因 v1.17.10 再次被引用为同类问题。  
- **链接**：https://github.com/anomalyco/opencode/issues/8785

### 7. #33865 – v1.17.10 在 Windows x64 上反复段错误
- **评论 2 | 👍 4**  
- **摘要**：Windows 11 环境，npm 全局安装后反复崩溃，错误与 Bun v1.3.14 相关。  
- **链接**：https://github.com/anomalyco/opencode/issues/33865

### 8. #26890 – 从配置中移除插件后 OpenCode 段错误，永久不可用
- **评论 2 | 👍 3**  
- **摘要**：移除 `opencode.jsonc` 中的插件条目后立即崩溃，重启无法恢复。  
- **链接**：https://github.com/anomalyco/opencode/issues/26890

### 9. #33328 – [WEB] 切换项目时草稿输入丢失
- **评论 3 | 👍 0**  
- **摘要**：Web UI 中未发送的草稿在切回项目时不复存在，疑似状态管理缺陷。  
- **链接**：https://github.com/anomalyco/opencode/issues/33328

### 10. #33863 – 多实例 + 大量 git.exe 僵尸进程导致系统资源耗尽
- **评论 1 | 👍 0**  
- **摘要**：单实例 OpenCode 衍生多个 `git.exe` 及 `conhost.exe`，内存/磁盘占用飙升，需重启系统。  
- **链接**：https://github.com/anomalyco/opencode/issues/33863

---

## 重要 PR 进展（10 条）

### 1. #33871 – 修复：在 Windows 上禁用 Kitty 键盘以避免 IME 崩溃
- **描述**：中文/日文输入法导致 TUI 崩溃，根因为 Kitty 键盘协议拦截 IME 按键并错误编码，结合非 UTF-8 控制台引发未处理异常。  
- **链接**：https://github.com/anomalyco/opencode/pull/33871

### 2. #33870 – 修复：忽略空斜杠命令
- **描述**：单独的 `/` 被解析为空命令，导致 TUI 卡死。关闭 #33867。  
- **链接**：https://github.com/anomalyco/opencode/pull/33870

### 3. #32767 – 修复：恢复子代理会话的 ESC 中断功能
- **描述**：修复长期存在的回归问题，允许用户按 Escape 中断委托给子 agent 的会话。关闭 #3699/#4073/#23534。  
- **链接**：https://github.com/anomalyco/opencode/pull/32767

### 4. #33868 – 修复：在引用别名下搜索文件
- **描述**：`@docs` 正常匹配但 `@docs/...` 路径搜索失败，修正别名匹配逻辑。关闭 #33573。  
- **链接**：https://github.com/anomalyco/opencode/pull/33868

### 5. #33864 – 修复：桌面版 Issue 链接指向真实表单
- **描述**：帮助菜单中链接的 `feature_request.yml` 和 `bug_report.yml` 在仓库中实际不存在，改为正确的表单地址。关闭 #33855。  
- **链接**：https://github.com/anomalyco/opencode/pull/33864

### 6. #33861 – 修复：MCP prompt 命令使用真实参数渲染
- **描述**：之前列出命令时使用 `$1`、`$2` 等占位符，导致带复杂参数服务器的命令格式错误。关闭 #33564。  
- **链接**：https://github.com/anomalyco/opencode/pull/33861

### 7. #33822 – 修复 CI：Windows 上使用 Bun 1.4 canary
- **描述**：Bun 1.3.14 在 Windows 上大量段错误，Rust 重写版（Bun 1.4 canary）更稳定。  
- **链接**：https://github.com/anomalyco/opencode/pull/33822

### 8. #33856 – 修复 CLI：忽略 mini 模式中的陈旧回合事件
- **描述**：在 `--mini` 模式下，旧回合的剩余事件会干扰新 prompt，导致错误状态。关闭 #33761。  
- **链接**：https://github.com/anomalyco/opencode/pull/33856

### 9. #33854 – 修复配置：`{env:VAR}` 缺失时静默解析为空字符串
- **描述**：之前缺失环境变量不报错，导致 MCP 服务器用空 header 启动。现改为抛出 `ConfigInvalidError`。关闭 #33853。  
- **链接**：https://github.com/anomalyco/opencode/pull/33854

### 10. #33839 – 新增：TUI 中增加 `/session-id` 和 `/session-info` 斜杠命令
- **描述**：允许用户快速复制会话 ID 或信息到剪贴板，无需 LLM 往返或插件。关闭 #11937。  
- **链接**：https://github.com/anomalyco/opencode/pull/33839

---

## 功能需求趋势

- **MCP 集成深化**：v1.17.10 添加了 MCP 资源模板和读取工具，社区同时提交了渲染修复和缺失变量错误处理，表明 MCP 正成为核心扩展机制。
- **Windows 稳定性成为第一优先级**：大量段错误报告驱动了 CI 切换到 Bun 1.4 canary、IME 键盘协议禁用等紧急修复。
- **Skills 配置语法变更**：`skills` 从数组改为对象，用户反馈旧语法兼容性和版本重下载机制不足，需改进迁移体验。
- **会话与草稿管理**：切换项目丢失输入、reasoning_content 剥离、子代理中断等问题显示会话状态持久化仍是薄弱环节。
- **桌面版 UI 增强**：新增 Chrome 标签切换快捷键、重新设计会话 UI 样式（v2 tokens），表明团队正加速桌面版完善。

---

## 开发者关注点

1. **Windows 崩溃头痛**  
   - v1.17.10 捆绑的 Bun 1.3.14 在 Windows x64 上引发大量段错误，用户被迫降级并忍受 SQLite schema 不兼容。社区急切期待官方修复（PR #33822 已在 CI 中使用 Canary 版本测试）。
2. **配置迁移成本**  
   - `skills` 从数组变对象未提供自动迁移，旧配置文件直接报错或静默失效，用户需手动重写。
3. **推理过程透明性**  
   - `reasoning_content` 被剥离不仅影响缓存，也让用户无法审查模型的思考过程，期待官方回放逻辑保留所有内容。
4. **资源泄漏问题**  
   - 多实例和 git 进程泄漏导致系统卡顿，虽非首次报告（#33863），但影响生产环境使用。
5. **环境变量安全**  
   - `{env:VAR}` 缺省不报错可能泄露敏感信息（如 MCP 认证头），PR #33854 及时修正了此隐患。

> 日报基于 `anomalyco/opencode` 仓库 2026-06-25 的公开数据生成，仅反映过去 24 小时社区动态。建议关注 #33742 和 #33822 以获取 Windows 稳定性修复进展。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi 社区动态日报 | 2026-06-25

数据来源：[github.com/badlogic/pi-mono](https://github.com/badlogic/pi-mono)（项目别名 `earendil-works/pi`）

---

## 📊 今日速览

今天社区活跃度较高，共更新 **29 条 Issue** 和 **9 个 PR**。最值得关注的动态包括：针对 OpenAI Codex 和 Anthropic 的连接卡死问题有了修复 PR（#6051）；Amazon Bedrock Mantle 新供应商的支持进入最后阶段（#5509）；实验性的 `pi-orchestrator` 包被提出（#6064），为多实例管理铺路。此外，关于 TUI 崩溃、会话恢复后输入历史丢失等体验问题也在今天得到关闭。

---

## 🚀 版本发布

过去 24 小时无新版本发布。

---

## 🔥 社区热点 Issues（Top 10）

### 1. `#4945` openai-codex 连接可靠性问题
- **状态**：OPEN · `[inprogress]` · 70 条评论 · 👍 30
- **摘要**：`openai-codex` / `gpt-5.5` 交互式 TUI 经常卡在 `Working…`，无输出也无错误，只能按 Escape 中止。过去几天反复出现。
- **重要原因**：影响核心用户体验，社区反馈强烈，是当前最受关注的 Issue。
- **[链接](https://earendil-works/pi/issue/4945)**

### 2. `#5363` 新增 Amazon Bedrock Mantle 供应商（OpenAI 兼容模型）
- **状态**：OPEN · 14 条评论 · 👍 4
- **摘要**：现有 `amazon-bedrock` 使用 Converse API，而 Bedrock Mantle 模型使用 OpenAI 兼容的 API，需新增供应商。
- **重要原因**：AWS 用户刚需，对应 PR #5509 已开放，预计很快合并。
- **[链接](https://earendil-works/pi/issue/5363)**

### 3. `#5291` 使用 Anthropic 订阅时会话卡在 “Working”
- **状态**：CLOSED · 7 条评论 · 👍 2
- **摘要**：新 Anthropic Enterprise 订阅用户发现会话间歇性卡死，中断有时无效，等待更久后才恢复。
- **重要原因**：与 #4945 问题类似，已被 PR #6051 修复（添加流空闲超时和重试机制）。
- **[链接](https://earendil-works/pi/issue/5291)**

### 4. `#5671` ~/.pi 与项目内 cwd/.pi 重叠
- **状态**：CLOSED · 6 条评论 · 👍 5
- **摘要**：`$HOME/.pi` 既存放全局设置也存放项目本地设置，导致在 `$HOME` 目录下运行时混淆。虽不严重，但引起配置命名讨论。
- **重要原因**：涉及多个用户的工作空间管理，社区关注度高。
- **[链接](https://earendil-works/pi/issue/5671)**

### 5. `#5595` openai-completions 的 maxTokens 未正确透传
- **状态**：OPEN · `[inprogress, to-discuss]` · 5 条评论 · 👍 2
- **摘要**：使用 Together.ai 等 OpenAI 兼容供应商的推理模型（如 DeepSeek v4 pro）时，输出 token 被截断，不受设置控制。
- **重要原因**：直接影响模型输出质量，需要底层修复。
- **[链接](https://earendil-works/pi/issue/5595)**

### 6. `#5810` RPC 暴露会话条目和树（`get_entries`、`get_tree`）
- **状态**：CLOSED · 5 条评论
- **摘要**：新增两个只读 RPC 命令，允许外部工具异步读取会话内容，支持游标分页。
- **重要原因**：为 headless 集成和工具生态打通提供基础设施。
- **[链接](https://earendil-works/pi/issue/5810)**

### 7. `#6050` TUI 全重绘清除终端回滚缓冲区
- **状态**：CLOSED · 5 条评论
- **摘要**：交互模式下，TUI 频繁重绘导致终端滚动条跳回聊天开头，影响阅读。
- **重要原因**：直接影响 TUI 体验，已修复。
- **[链接](https://earendil-works/pi/issue/6050)**

### 8. `#6019` OpenAI Responses 流中断后未重试
- **状态**：CLOSED · 4 条评论
- **摘要**：当 OpenAI 响应流开始后出现可重试错误，Pi 却以 `stopReason: "error"` 结束，未尝试重试。
- **重要原因**：浪费 API 调用，降低成功率，已修复。
- **[链接](https://earendil-works/pi/issue/6019)**

### 9. `#6009` OpenAI Responses 乱序输出导致推理状态丢失
- **状态**：CLOSED · 3 条评论
- **摘要**：当 Responses API 的 output items 乱序完成时，推理块的 `thinkingSignature` 被丢弃，后续响应无法利用先前推理。
- **重要原因**：影响推理模型（如 GPT-5.5）的连贯性。
- **[链接](https://earendil-works/pi/issue/6009)**

### 10. `#6060` TUI footer 渲染工具调用消息时报错 `content is not iterable`
- **状态**：CLOSED · 3 条评论
- **摘要**：当助手消息仅包含工具调用（无文本内容）时，统计 token 的函数报错导致 TUI 崩溃。
- **重要原因**：属于新出现的崩溃 bug，今天立即修复。
- **[链接](https://earendil-works/pi/issue/6060)**

---

## 📌 重要 PR 进展（全部 9 条）

### 1. `#5832` [OPEN] 修复：暴露供应商 HTTP 错误正文，替代模糊的 SDK 消息
- **摘要**：当代理/网关返回非 2xx 响应时，Pi 之前丢弃了错误正文，显示 `Unknown: UnknownError`。此 PR 让 HTTP 错误正文透传。
- **状态**：OPEN，已关联 Issue #5763。
- **[链接](https://earendil-works/pi/pull/5832)**

### 2. `#6063` [CLOSED] 扩展加载计时功能
- **摘要**：实现 `PI_STARTUP_BENCHMARK=1` 和 `PI_TIMING=1` 环境变量，可单独统计每个扩展的加载时间。同时修复退出时输出 OSC 垃圾字符的问题。
- **状态**：CLOSED，关联 Issue #6062。
- **[链接](https://earendil-works/pi/pull/6063)**

### 3. `#6064` [OPEN] 实验性：Pi 编排器（orchestrator）
- **摘要**：新增 `@earendil-works/pi-orchestrator` 包，运行本地编排守护进程并通过 Unix Socket 提供 IPC，支持启动、列出、停止多个 Pi 实例。
- **状态**：OPEN，刚提交。
- **[链接](https://earendil-works/pi/pull/6064)**

### 4. `#6056` [CLOSED] 简化子代理配置，添加默认 agent，使用 MiniMax 模型
- **摘要**：将示例 agent 配置（planner, reviewer, scout, worker）切换到 MiniMax-M2.7 模型，简化输出格式，添加 `default.md` 通用 agent。
- **状态**：CLOSED。
- **[链接](https://earendil-works/pi/pull/6056)**

### 5. `#6055` [CLOSED] 同上（#6056 的重复提交，已关闭）
- **摘要**：与 #6056 内容相同，已关闭。
- **[链接](https://earendil-works/pi/pull/6055)**

### 6. `#6054` [CLOSED] 新增 `runParallelAgentTasks` + 并行批处理系统提示
- **摘要**：允许 agent 在 `toolExecution: "parallel"` 之外，真正并行运行独立的 agent 循环（每个循环一个 LLM 往返）。同时向系统提示添加批处理指南。
- **状态**：CLOSED，关联 Issue #6053。
- **[链接](https://earendil-works/pi/pull/6054)**

### 7. `#5509` [OPEN] 新增 Amazon Bedrock Mantle OpenAI Responses 供应商
- **摘要**：基于 Azure OpenAI 响应适配器，支持 Bedrock Mantle API（当前仅支持 GPT 5.5 和 5.4）。此 PR 已更新近三周，今天有进展。
- **状态**：OPEN，关联 Issue #5363。
- **[链接](https://earendil-works/pi/pull/5509)**

### 8. `#6051` [CLOSED] 修复：从挂起的流中恢复并重试未建模的 Bedrock 错误
- **摘要**：添加流空闲超时（默认 240s）、连接超时（默认 120s）以及错误重试策略（包括 403 限速错误）。修复了 Bedrock 和 Anthropic 连接卡死问题。
- **状态**：CLOSED，修复 #4945 和 #5291。
- **[链接](https://earendil-works/pi/pull/6051)**

### 9. `#6048` [CLOSED] 修复：恢复会话时资源显示在消息之前
- **摘要**：修正了 `2417adb` 引入的回归问题：恢复或重载会话后，资源（Context, Skills, Prompts, Extensions）应显示在已恢复的消息上方，而非下方。
- **状态**：CLOSED。
- **[链接](https://earendil-works/pi/pull/6048)**

---

## 💡 功能需求趋势

从今日热点的 Issues 和 PR 中，可以提炼出社区最关注的四个方向：

1. **多模型供应商扩展**  
   - 新增 Amazon Bedrock Mantle（OpenAI 兼容）是明确的呼声，相关 Issue #5363 和 PR #5509 活跃度很高。  
   - 同时 MiniMax 模型（#6061）和 DeepSeek（#5595）也在社区讨论圈中。

2. **连接可靠性 & 流处理健壮性**  
   - 多个 issue 反映 `Working…` 卡死问题（#4945, #5291），已由 PR #6051（空闲超时、重试）和 #6019（流中断重试）尝试解决。  
   - 社区期待更优雅的超时和错误恢复机制。

3. **远程控制与可编程性**  
   - RPC 暴露会话条目（#5810）和实验性的 `pi-orchestrator`（PR #6064）表明开发者希望在 CI/CD、IDE 插件等场景中驱动 Pi。  
   - 跨实例生命周期管理成为新趋势。

4. **TUI 用户体验优化**  
   - 滚动回滚清除（#6050）、行截断崩溃（#6058）、输入历史丢失（#6066）、斜杠技能选择器（#6059）等问题表明社区对终端交互的细节要求越来越高。  
   - 扩展性能计时（#6062）也反映了对启动速度的关注。

---

## 🔧 开发者关注点（痛点 & 高频需求）

- **连接断续导致任务中断**：多个用户反馈会话意外卡死（#4945, #5291），必须手动干预。尽管 PR #6051 已修复，但仍需持续监控。  
- **token 限制设置无效**：openai-completions 供应商（#5595）的 maxTokens 未被透传，影响 DeepSeek 等推理模型的效果。  
- **文件被静默截断**：`SessionManager.open()` 在传入非会话文件时直接截断文件（#6002），属于高危 bug，虽未合并但社区反响强烈。  
- **扩展加载时间增加**：多家用户反映扩展加载变慢，催生了 #6062 计时功能，但根因仍需定位。  
- **跨项目 Node.js 版本依赖**：个别项目使用旧版 Node 导致 Pi 无法运行，社区期待单文件可执行二进制（#6065）以避免 nvm 切换。  
- **安全包审核**：一个下载量 20 万/月的包 `@hypabolic/pi-hypa` 被标记为“恶意或危险行为”（#6052），社区请求官方介入检查。此事虽仅 2 条评论，但提醒扩展生态的安全审查需要加强。

---

**总结**：今日社区重点在于修复连接稳定性、丰富模型供应商选择，以及优化终端用户交互体验。`pi-orchestrator` 的提出预示 Pi 正朝向可作为后台守护进程的远程执行工具进化。开发者对包安全、进程隔离和启动性能的关切值得持续跟踪。

*数据截止 2026-06-25 23:59 UTC。*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code 社区动态日报 | 2026-06-25

**今日速览**  
Qwen Code 今日发布 v0.19.2 正式版，新增远程 LSP 状态路由；社区围绕会话管理、语音听写与 CI 流程优化提交了多个高热度 PR。同时，新的 Issue 反映了流式超时、上下文压缩网关超时等运行稳定性问题，开发团队正通过自动化解锁候选 Issue 以提升修复效率。

---

## 版本发布

### v0.19.2（正式版）
- **新增** `feat(serve): Add remote LSP status route`（PR #5762 by @doudouOUC），为远程 LSP 服务提供状态查询端点。
- **发布说明**：[GitHub Release](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2)

### v0.19.2-preview.0
- 包含与 v0.19.2 相同的 LSP 路由功能，作为预览通道先行验证。
- [GitHub Release](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2-preview.0)

### v0.19.2-nightly.20260625.b2f11b735
- **修复** `fix(core): allow web_fetch JSON fallback` by @tt-a1i（PR #5660），增强 `web_fetch` 工具在 JSON 解析失败时的降级处理。
- [GitHub Release](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2-nightly.20260625.b2f11b735)

---

## 社区热点 Issues（共 8 条，全部收录）

1. **#401 [BUG] API Error: Streaming setup timeout after 6s**  
   - **作者** @adi0900 | **更新** 2026-06-25 | **评论数 8**  
   - 用户安装 CLI 后频繁遇到流式超时，提示减少输入长度或增加超时配置。该问题已存在近一年，今日有更新，说明社区仍未彻底解决。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/401)

2. **#5863 [FEATURE] enrich GET /session/:id/status with live turn-phase / active tools / pending permissions**  
   - **作者** @samuelhsin | **评论数 2**  
   - 提议扩展会话状态接口，返回当前轮次阶段、活跃工具和待定权限信息，是会话管理精细化的重要诉求。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5863)

3. **#5861 [BUG] Context compression request should use stream=true to avoid gateway timeout**  
   - **作者** @doudouOUC | **评论数 2**  
   - 上下文压缩服务使用非流式请求，导致大模型推理完成前 HTTP 无法响应，在代理部署下容易触发网关超时。**欢迎 PR 贡献**。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5861)

4. **#5855 [FEATURE] query a single session's status by id**  
   - **作者** @samuelhsin | **评论数 2**  
   - 新增 daemon 端 HTTP 接口 `GET /session/:id/status`，返回单会话的实时摘要（如工作目录、客户端数量、是否正在执行 prompt）。**标记为 ready-for-agent**。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5855)

5. **#5789 [FEATURE] Enable built-in status line preset by default for new users**  
   - **作者** @pomelo-nwu | **已关闭** | **评论数 3**  
   - 建议新用户首次启动时默认显示状态行（当前模型、Git 分支、上下文使用量、工作目录），降低上手门槛。社区认可度高，已合并关闭。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5789)

6. **#5782 [ENHANCEMENT] WebFetch should reject URLs containing userinfo**  
   - **作者** @VectorPeak | **已关闭** | **评论数 2**  
   - 拒绝 `http(s)://user:pass@host` 格式的 URL 请求，防止敏感信息在用户界面或诊断日志中泄露。**欢迎 PR 贡献**。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5782)

7. **#5816 [FEATURE] Voice dictation: support a user-configurable keyterms file for ASR biasing**  
   - **作者** @qqqys | **已关闭** | **评论数 2**  
   - 语音听写（#5755）目前使用硬编码开发者术语列表，用户无法扩展。提议支持自定义关键词文件，提升领域识别准确率。已关闭，期望后续实现。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5816)

8. **#5742 [FEATURE] Improve voice package distribution for mirror registry installs**  
   - **作者** @qqqys | **已关闭** | **评论数 1**  
   - 从镜像/私有仓库安装时，原生语音捕获包 `@qwen-code/audio-capture` 可能被静默跳过。建议将其捆绑到公共包中。  
   - [链接](https://github.com/QwenLM/qwen-code/issues/5742)

---

## 重要 PR 进展（精选 10 条）

1. **#5847 feat(serve): add runtime context injection for per-turn system-reminders**  
   - **作者** @callmeYe | **状态 OPEN**  
   - 新增按会话的运行时上下文存储，外部调用者（daemon API、SDK）可注入身份/规则等动态上下文，注入为每轮系统提醒。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5847)

2. **#5856 feat(desktop): voice dictation in the desktop app**  
   - **作者** @qqqys | **状态 OPEN**  
   - 将语音听写功能移植到桌面应用，工具栏增加麦克风按钮，支持录音条与暂停按钮，匹配 CLI 和 Web Shell 体验。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5856)

3. **#5738 fix(cli): default to virtualized terminal history**  
   - **作者** @ZevGit | **状态 OPEN**  
   - 默认启用虚拟化历史视图，新用户直接获得应用内可滚动历史面板；原有使用终端回滚的用户可设置 `ui.useTerminalBuffer` 恢复。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5738)

4. **#5778 feat(cli): add /model --vision for a fallback vision model**  
   - **作者** @yiliang114 | **状态 OPEN**  
   - 新增 `/model --vision` 命令及交互选择器，当主模型不支持图片时，自动切换到配置的视觉模型。完善了多模态场景。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5778)

5. **#5829 fix(desktop): reject unsafe source slugs before deletion**  
   - **作者** @VectorPeak | **状态 OPEN**  
   - 防止通过路径式 slug（如 `../sessions`）将删除目标移到工作区外的目录，增强文件操作安全性。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5829)

6. **#5550 Add a Secret Disclosure mandate to prevent secret exposure on broad file tasks**  
   - **作者** @warmjademe | **状态 OPEN**  
   - 在宽泛文件任务（如复制/同步全部文件）中，添加秘密泄露强制检测，防止私钥、`.env` 等被误拷贝到公开目标。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5550)

7. **#5661 feat(tui): partition tool display by type — collapse read/search, show mutation tools individually**  
   - **作者** @chiga0 | **状态 OPEN**  
   - 按工具类型分区显示：读类/搜索类工具折叠为摘要行，写类工具逐条展示，提升 UI 可读性与操作透明度。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5661)

8. **#5848 feat(ui): add ui.history.collapsePreviewCount to show last N turns when resuming collapsed sessions**  
   - **作者** @mvanhorn | **状态 OPEN**  
   - 新增 `ui.history.collapsePreviewCount` 设置，恢复折叠会话时保留最近 N 轮用户轮次可见，平衡信息保留与界面简洁。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5848)

9. **#5811 fix(cli): improve token speed accounting**  
   - **作者** @seekskyworld | **状态 OPEN**  
   - 优化 token/s 显示：统计思考文本、使用墙钟时间、暂停工具执行期间计时，使实时速率更准确。  
   - [链接](https://github.com/QwenLM/qwen-code/pull/5811)

10. **#5849 feat(cli): support @extension mention in input autocomplete**  
    - **作者** @callmeYe | **状态 OPEN**  
    - 在 CLI 输入中支持 `@` 唤起扩展名、描述与“扩展”徽章的自动补全，提升扩展发现与调用体验。  
    - [链接](https://github.com/QwenLM/qwen-code/pull/5849)

---

## 功能需求趋势

- **会话管理深入细化**：多个 Issue 和 PR 聚焦会话状态查询（单会话、多字段丰富）、上下文压缩流式化、断连恢复（#5030）等，反映社区对 daemon 模式下的会话控制要求越来越高。
- **语音功能全面铺开**：从 CLI 到 Web Shell 再到桌面应用，语音听写覆盖所有前端；同时用户希望自定义关键词文件，说明语音功能已进入实用化迭代阶段。
- **安全审计自动化**：WebFetch URL 用户信息过滤、删除路径净化、秘密泄露强制报告等 PR 显示社区对敏感信息保护有强烈需求，且正积极贡献。
- **CI/CD 流程优化**：多个 CI PR（#5854、#5859、#5860、#5862）旨在缓解 runner 饱和、将任务迁移到自托管或托管 runner、自动修复 Issue 候选过滤等，表明项目规模增长后运维压力增大。
- **UI 可配置性与信息密度**：状态行默认开启、折叠历史预览数量、工具类型分区显示等，说明用户期望更智能、可定制的界面体验。

---

## 开发者关注点

- **流式超时与上下文压缩瓶颈**：Issue #401 和 #5861 是高频痛点，尤其在代理部署和长上下文场景中，请求卡死或网关超时严重影响体验。多位社区成员建议增加可配置超时、默认启用流式。
- **安装与分发可靠性**：Issue #5742 指出语音包在镜像注册表下的静默缺失，提示需要改进包分发策略，确保功能完整性。
- **安全合规细节**：WebFetch 接受含凭据 URL（#5782）的漏洞虽已关闭，但类似问题在其他文件操作中仍需警惕；PR #5550 和 #5829 代表社区正主动加固安全边界。
- **自动化修复效率**：PR #5860 揭示自动修复工作流因筛选过严而找不到候选 Issue，反映了自动化运维需要更精细的规则调优。

---

*日报自动生成，数据截至 2026-06-25 23:59 UTC。如有疏漏，欢迎指正。*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，以下是基于您提供的 GitHub 数据生成的 2026-06-25 DeepSeek TUI (CodeWhale) 社区动态日报。

**说明**: 根据您提供的数据，近期社区动态主要围绕项目 **CodeWhale** 展开，这可能是 DeepSeek TUI 的迭代项目或相关工具的代号。本报告将基于该数据进行分析。

---

# 🐋 CodeWhale 社区日报 - 2026-06-25

## 1. 今日速览

今日社区活跃，核心进展集中在 `v0.8.65` 版本的收尾工作。关键内容包括：**ACP 协议适配**获得实质性突破，现已开放 provider/model 选择能力；一系列**关键 Bug 被修复**，如 Windows 环境变量继承、JS 执行代理等问题；同时，**“Fleet loadout auto”** 和 **上下文自定义**成为社区呼声最高的新需求。此外，一个 **Plan 与 Agent 模式混淆**的顽固 Bug 再次被用户提及，引发关注。

## 2. 版本发布

*过去24小时内无新版本发布。*

## 3. 社区热点 Issues

本期共 10 条活跃 Issue，以下为最值得关注的 5 条：

*   **#3205: v0.8.65: Fleet模型类、自动载荷与语义路由角色**
    *   **重要性**: ⭐⭐⭐⭐⭐ 核心特性。这是 `v0.8.65` 版本的旗舰级需求，旨在为 TUI、CLI、Fleet 等所有子系统构建统一的模型/载荷选择器。其中的 **“Fleet loadout auto”** 自动模式是解决 Fleet 计算资源分配的关键。
    *   **社区反应**: 讨论热烈（10 条评论），但暂无用户点赞，可能因功能尚在开发中，用户更关注实现细节。
    *   **链接**: [Hmbown/CodeWhale Issue #3205](https://github.com/Hmbown/CodeWhale/issues/3205)

*   **#2300: v0.8.65: 多模型兼容性、Provider 文档与自动Fleet载荷选择**
    *   **重要性**: ⭐⭐⭐⭐ 用户需求。作为 #3205 的用户侧验收清单，此 Issue 集中了关于 **provider 路由、多模型支持**的改进诉求，特别是对 vLLM 和 OpenAI 兼容端点差异的文档化需求。
    *   **社区反应**: 有 7 条评论，表明用户对 provider 配置的复杂性和文档清晰度有明确期望。
    *   **链接**: [Hmbown/CodeWhale Issue #2300](https://github.com/Hmbown/CodeWhale/issues/2300)

*   **#3572: [Bug] Windows用户环境变量未被CodeWhale shell继承**
    *   **重要性**: ⭐⭐⭐⭐ 影响广泛的 Bug。影响了在 Windows 上使用 `exec_shell` 命令的所有开发者，导致构建工具等依赖用户级环境变量的工作流失效。此 Issue 已有对应的 PR (#3578) 修复。
    *   **社区反应**: 用户报告清晰，已快速得到解决（已关闭）。
    *   **链接**: [Hmbown/CodeWhale Issue #3572](https://github.com/Hmbown/CodeWhale/issues/3572)

*   **#3546: [已关闭] 扩展 ACP 支持，暴露 Provider 和模型选择**
    *   **重要性**: ⭐⭐⭐⭐ 协议集成。允许外部 ACP 客户端（如 Paseo）发现并切换 CodeWhale 的 provider 和模型。这极大增强了 CodeWhale 作为 AI 后端的互操作性。
    *   **社区反应**: 从 Issue 提出到关闭仅用一天，并有对应 PR (#3576) 实现，表明开发效率极高，是今日关键进展之一。
    *   **链接**: [Hmbown/CodeWhale Issue #3546](https://github.com/Hmbown/CodeWhale/issues/3546)

*   **#3568: [Bug] Plan 和 Agent 模式混合的问题似乎仍然存在**
    *   **重要性**: ⭐⭐⭐ 顽固 Bug 复现。用户报告在 Plan 模式下 AI 仍错误地执行了 Agent 的“尝试-修改”流程，说明模式切换逻辑存在核心问题，需要重点关注。
    *   **社区反应**: 用户提供了详细的聊天导出文件作为证据，有助于开发者快速定位问题。
    *   **链接**: [Hmbown/CodeWhale Issue #3568](https://github.com/Hmbown/CodeWhale/issues/3568)

## 4. 重要 PR 进展

*   **#3578: [已合并] 保留Windows SDK环境变量根**
    *   **功能**: 修复了 `#3572`，确保 Windows 子进程能正确获取用户级环境变量（如 SDK、工具链路径）。
    *   **链接**: [Hmbown/CodeWhale PR #3578](https://github.com/Hmbown/CodeWhale/pull/3578)

*   **#3576: [已合并] 在ACP stdio适配器上暴露Provider和模型选择**
    *   **功能**: 实现了 `#3546` 需求，允许外部工具通过 ACP 协议动态切换 CodeWhale 的后端模型。
    *   **链接**: [Hmbown/CodeWhale PR #3576](https://github.com/Hmbown/CodeWhale/pull/3576)

*   **#3577: [已合并] 修复JS执行工具代理环境变量问题**
    *   **功能**: 修复了 `js_execution` 工具无法感知系统代理设置 (`HTTP_PROXY`) 的问题，解决了特定网络环境下的超时故障。
    *   **链接**: [Hmbown/CodeWhale PR #3577](https://github.com/Hmbown/CodeWhale/pull/3577)

*   **#3573: [已合并] 添加Provider上下文窗口覆盖功能**
    *   **功能**: 允许用户为不同 provider 自定义 `context_window` 大小，并贯通整个流程（TUI、CLI、压缩等）。
    *   **链接**: [Hmbown/CodeWhale PR #3573](https://github.com/Hmbown/CodeWhale/pull/3573)

*   **#3569: [已合并] TUI显示已配置的ask rules**
    *   **功能**: 添加 `/config ask-rules` 只读界面，让用户清晰看到当前生效的权限规则详情。
    *   **链接**: [Hmbown/CodeWhale PR #3569](https://github.com/Hmbown/CodeWhale/pull/3569)

*   **#3566: [已合并] 澄清精简工具转录行**
    *   **功能**: 改进了工具调用记录的显示，在合并显示时保留关键信息（如 `git_log` 等），让输出更清晰。
    *   **链接**: [Hmbown/CodeWhale PR #3566](https://github.com/Hmbown/CodeWhale/pull/3566)

*   **#3571: [开放中] 为OHOS和工具链做清理**
    *   **功能**: 尝试将 Rust 工具链从锁定版本 `1.88` 改为 `stable`，并清理无用配置，让项目能自动使用最新的稳定版 Rust 编译器。
    *   **链接**: [Hmbown/CodeWhale PR #3571](https://github.com/Hmbown/CodeWhale/pull/3571)

*   **#3526: [已合并] 强制执行main分支支持的发布标签**
    *   **功能**: 加强了发布流程安全性，确保 release 版本只能从已合并到 `main` 分支的 commit 构建，防止“分支发布”导致的问题。
    *   **链接**: [Hmbown/CodeWhale PR #3526](https://github.com/Hmbown/CodeWhale/pull/3526)

*   **#3197: [已合并] 将DeepSeek蓝色消费者重命名为Whale色调**
    *   **功能**: 视觉层面的品牌标识迁移，将代码中的颜色常量从 “DEEPSEEK_BLUE” 迁移至 “WHALE_ACCENT”，逐步去化原有命名。
    *   **链接**: [Hmbown/CodeWhale PR #3197](https://github.com/Hmbown/CodeWhale/pull/3197)

*   **#3551: [已合并] 澄清工具详情快捷键提示**
    *   **功能**: 优化 UI 提示，将模糊的快捷键提示（如 `Alt+V/v`）改为更明确的操作描述（如 `Alt+V opens details`），提升用户体验。
    *   **链接**: [Hmbown/CodeWhale PR #3551](https://github.com/Hmbown/CodeWhale/pull/3551)

## 5. 功能需求趋势

*   **模型/Provider 路由与配置的深度优化**: 社区强烈要求一个更智能、更自动化的模型选择和管理系统（如 Fleet Loadout Auto），并希望获得对 Provider 端点、模型 ID、认证方式以及上下文大小的全面自定义能力。这反映出用户对“多模型策略”和“灵活配置”的刚性需求。
*   **协议互操作性增强**: 以 ACP 协议集成 (Issue #3546) 为代表，社区正在探索将工具作为更开放的 AI 后端。满足外部 IDE 或编排工具的集成需求，正成为一个新的功能增长点。
*   **内存与召回机制的现代化**: PR #3575 提出集成 `moraine-mcp` 作为全新的召回工具，并增加配置门控来废弃旧有的 push/inject 机制，表明工具的内存管理正从被动注入向主动、模块化的 MCP 服务演进。

## 6. 开发者关注点

*   **跨平台兼容性是持续痛点**: Windows 用户遇到了环境变量无法继承的问题 (#3572)，凸显了在多 OS 环境下保持一致的开发体验仍是亟待解决的挑战。
*   **模式混淆问题影响核心体验**: Plan/Agent 模式混用的问题 (#3568) 再次出现，这关乎 AI 代理行为是否符合用户预期，是决定工具“可控性”的核心 Bug，需要彻底修复。
*   **编译环境与工具链的“自由”诉求**: 有开发者提交 PR (#3570, #3571) 希望将 Rust 工具链从锁定版本改为 `stable`，反映出社区希望项目能更灵活地适应新编译器，减少对特定工具链版本的依赖。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*