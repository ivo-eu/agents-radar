# AI CLI 工具社区动态日报 2026-06-28

> 生成时间: 2026-06-28 10:09 UTC | 覆盖工具: 9 个

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

# AI CLI 工具横向对比分析报告（2026-06-28）

---

## 一、生态全景

当前 AI CLI 工具生态正处于 **“平台化扩张”与“稳定性阵痛”并存** 的阶段。一方面，Claude Code、OpenAI Codex、Gemini CLI 等主流工具持续加码 MCP 协议、插件系统、多代理协作等基础设施；另一方面，几乎所有工具都面临 **Windows 平台兼容性、Token 成本失控、Agent 过度自主** 三大共性难题。社区反馈从“能否完成功能”转向“能否可靠、安全、经济地完成”，标志着行业进入 **成熟度爬坡期**。

---

## 二、各工具活跃度对比（2026-06-28）

| 工具 | 新增/活跃 Issues | 重要 PR | 版本发布 | 社区热度信号 |
|------|----------------|---------|----------|---------------|
| **Claude Code** | 10 条（含旧 Issue 更新） | 2 个 | 无 | Bash 截断 #26954（27👍）、Windows 认证回归 #71708 |
| **OpenAI Codex** | 10 条 | 10+ 个 | 3 个 Rust alpha | SQLite 写入巨量 #28224（400👍） |
| **Gemini CLI** | 10 条 | 10 个 | 1 个 nightly v0.51 | 焦点丢失 #22193、Shell 卡死 #25166 |
| **GitHub Copilot CLI** | 8 条 | 3 个 | 无 | Ubuntu keychain #2165（20👍）、MCP 启动回归 #3958 |
| **Kimi Code CLI** | 1 条 | 0 | 无 | 内存泄漏 #1592（未修复，3个月） |
| **OpenCode** | 10 条 | 10 个 | 无 | Windows 崩溃 #33742（47评论/45👍） |
| **Pi** | 23 条 | 8 个 | 无 | 自动滚屏 #5825（34评论）、多发模型兼容修复 |
| **Qwen Code** | 9 条 | 10 个 | 2 个（v0.19.3 + nightly） | 模型自动切换致资金损失 #5819 |
| **DeepSeek TUI** | 10 条 | 10 个 | 准备发布 v0.8.66 | 缓存命中率低 #1177（24评论） |

**解读**：OpenAI Codex、Pi、OpenCode 今日公关和 Issue 处理量最大；Kimi Code 最沉寂；DeepSeek TUI 的缓存问题为社区讨论热点。

---

## 三、共同关注的功能方向

| 方向 | 涉及工具 | 具体诉求 |
|------|----------|----------|
| **Token/成本管控** | Claude Code (#70595)、Codex (#30387)、DeepSeek TUI (#743/#1177)、Qwen Code (#5756/#5950) | 输出截断、默认上限过低、缓存命中率差、token 消耗异常大 |
| **Windows 稳定性** | Claude Code (#71708)、Codex (#26104/#27924/#28224)、Copilot CLI (#3958)、OpenCode (#33742)、DeepSeek TUI (#3717) | 崩溃、OAuth 证书过期、MCP 启动失败、输入卡顿、日志写入威胁 SSD |
| **Agent 行为可控性** | Claude Code (#70615/#70625)、Gemini CLI (#22093/#22323)、DeepSeek TUI (#3275)、Copilot CLI (#3874) | 不遵循规则、猜测答案、子代理未授权运行、过度主动执行 |
| **MCP 协议稳定性** | Claude Code (#72014)、Codex (#30292 系列)、Gemini CLI (#27878/#27889)、OpenCode (#34077) | OAuth 凭据刷新、MIME 类型嗅探、并发 token 刷新冲突 |
| **IDE 集成体验** | Gemini CLI (#22193)、Kimi Code (#1592)、OpenCode (#33659)、Qwen Code (#5941) | 焦点丢失、高内存占用、滚动异常、扩展兼容性 |
| **安全与审计** | Claude Code (#72014)、Gemini CLI (#27915)、Copilot CLI (#3874)、OpenCode (#34256) | 策略门控、钩子失效、路径遍历防护 |

---

## 四、差异化定位分析

| 工具 | 核心定位 | 目标用户 | 技术路线 | 差异化亮点 |
|------|---------|---------|----------|------------|
| **Claude Code** | 通用 Agent 编程助手 | 全栈开发者 | 原生 CLI + MCP 插件 | 深度 MCP 策略门控（`protect-mcp`）、Slash 命令/技能扩展 |
| **OpenAI Codex** | 桌面版 + 多代理 v2 | 高级 Pro 用户 | Rust CLI + 桌面应用 | MCP OAuth 凭据精细化管理、远程插件默认启用 |
| **Gemini CLI** | 多模型 + Google 生态 | 安全敏感开发者 | 基于 Gemini 模型 + VSCode 扩展 | 安全钩子、信任对话框、子代理策略控制 |
| **GitHub Copilot CLI** | 轻量级终端助手 | GitHub 生态用户 | 作为 Copilot 配套终端 | 无需独立付费、集成 GitHub 认证；但功能最少 |
| **Kimi Code CLI** | 中文友好、VSCode 插件 | 中文开发者 | 基于 Moonshot AI 模型 | 界面最简，但活跃度最低 |
| **OpenCode** | 高性能桌面 CLI | 追求用户体验的开发者 | Bun 运行时 + TUI/桌面双模式 | `/compact` 命令、`opencode --mini` 模式、项目归档 |
| **Pi** | 多模型网关 + TUI | 高级用户/插件开发者 | 基于扩展 API 的生态 | Context Matrix 存储投影、`reportUsage` 成本回传 |
| **Qwen Code** | 全渠道智能体 | 国内全栈开发者 | ACP 协议 + 多通道（QQ/桌面/Web） | ACP 协议 / 语音输入 / 频道驻留智能体 |
| **DeepSeek TUI** | 轻量 TUI、高性价比 | 成本敏感开发者 | 基于 DeepSeek 模型 + Rust 原生构想 | 强 Token 缓存优化、Agent 回退策略 |

---

## 五、社区热度与成熟度

- **极活跃（日均 Issue/PR >10）：**  
  **OpenAI Codex**（400👍的SQLite问题成为现象级）、**Pi**（23条Issue + 8 PR，修复迅速）、**OpenCode**（Windows崩溃引发47评论）

- **稳定迭代：**  
  **Claude Code**（社区讨论多但Issue重复率高/duplicate泛滥）、**Gemini CLI**（PR堆栈密集）、**Qwen Code**（版本发布频繁+ACP推进）、**DeepSeek TUI**（正在冲刺v0.8.66，缓存问题为核心）

- **较低活跃：**  
  **GitHub Copilot CLI**（功能固定，社区以提问为主）、**Kimi Code CLI**（仅有1条活跃Issue，且3个月未修复）

**成熟度判断**：  
- **Claude Code** 虽历史悠久，但重复Issue泛滥，核心Bug（截断、Agent不听话）长期未修复，社区信任度下降。  
- **OpenAI Codex** 工程能力最强（多个PR堆栈），但Windows稳定性是黑洞。  
- **Gemini CLI** 安全机制最完善，但性能与Agent可控性仍是短板。  
- **Qwen Code** 和 **Pi** 分别在渠道融合和扩展API上领先，社区增长迅速。

---

## 六、值得关注的趋势信号

### 1. **Token成本焦虑 → 缓存/压缩成为核心竞争力**
DeepSeek TUI 的缓存命中率低至95%引发24条讨论，Qwen Code 的8K默认上限导致输出截断，Codex 的SQLite写入量达640TB/年——**任何能经济地管理Token的工具将在下一个周期胜出**。开发者开始要求：可配置压缩阈值、输入缓存评分、子代理成本回传。

### 2. **Agent“自主”与“可控”的矛盾激化**
从 Claude Code 的“猜测而不检查”到 Gemini CLI 的子代理未经授权运行，再到 DeepSeek TUI 的过度自主修改——用户不再满足于“能做事”，而是要求 **“按预期做事”**。新增 `preToolUse` 钩子、策略门控（`protect-mcp`）、失败回退提示等机制正是回应。

### 3. **MCP 协议成为基础设施，但稳定性问题暴露**
多个工具（Claude Code、Codex、Gemini CLI、OpenCode）同时致力于MCP的OAuth刷新、Schema兼容性、并发冲突修复。**MCP已从“选项”变为“必填项”**，但生态成熟度远低于预期，第三方MCP服务器频繁引发认证、格式、超时错误。

### 4. **Windows 用户成为“二等公民”**
几乎每个工具都有专属的Windows崩溃/卡顿/证书/路径问题。**跨平台一致性不佳正在赶走第一大用户群体**。Codex的SSD寿命威胁和OpenCode的段错误尤其严重。

### 5. **IDE集成深度决定用户黏性**
Gemini CLI 的焦点丢失、Kimi Code 的内存泄漏、OpenCode 的VSCode兼容性问题——用户期望的“在IDE内无感使用CLI”尚未实现。通过ACP协议（如Qwen Code、DeepSeek TUI的Zed集成）或统一聊天面板（Qwen Code的`@qwen-code/chat-panel`）正在成为主流解法。

### 6. **安全审计需求从企业向个人蔓延**
`protect-mcp` 的签名收据、Copilot CLI 的钩子失效、OpenCode 的路径遍历修复——表明 **开发者正在为Agent的不可预测行为购买“保险”**。可审计的决策日志、离线可验证的策略收据成为高级功能。

---

**给开发者的建议**：  
- 优先选择**Token成本透明、缓存命中率高**的工具（如 DeepSeek TUI 如果修复缓存问题，或 Qwen Code 的自定义压缩阈值）。  
- **Windows 用户**应避免升级到最新版本（Claude Code v2.1.177、OpenCode v1.17.10 均有严重回归），或转向 WSL 代理模式。  
- 若重视**安全控制**，Gemini CLI 的安全钩子和 Qwen Code 的 ACP 协议提供更细粒度审批；若需要**多模型网关**，Pi 的扩展生态更灵活。  
- **长期关注**：MCP 协议标准化进展、Agent 行为策略门控的普及、以及各工具对 Linux 和 macOS 的投入力度（Windows 短期无解）。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为 Claude Code 生态的技术分析师，以下是根据 `anthropics/skills` 仓库数据（截至 2026-06-28）生成的社区热点报告。

---

### Claude Code Skills 社区热点报告

#### 1. 热门 Skills 排行

以下按评论热度、功能影响力和社区关注度综合排序，列出当前最受关注的 5 个核心 Skills。

1.  **document-typography（#514）**
    -   **功能**：为 AI 生成的文档提供专业的排版质量控制，解决孤字（orphan）、寡行（widow）和编号错位等常见问题。
    -   **讨论热点**：社区高度认可其解决了 AI 文档生成中“最后一公里”的排版细节问题，认为这是提升文档专业度的刚需 Skill。反馈积极，状态为 **Open**。
    -   **链接**：[https://github.com/anthropics/skills/pull/514](https://github.com/anthropics/skills/pull/514)

2.  **testing-patterns（#723）**
    -   **功能**：一个综合性的测试模式 Skill，涵盖测试哲学（如测试奖杯模型）、单元测试（AAA 模式）、React 组件测试（Testing Library）等。
    -   **讨论热点**：社区对结构化、标准化的测试指导需求强烈。该 Skill 填补了官方仓库在测试方法论上的空白，讨论集中在如何使其与现有 CI/CD 工作流结合。
    -   **链接**：[https://github.com/anthropics/skills/pull/723](https://github.com/anthropics/skills/pull/723)

3.  **ODT document creation（#486）**
    -   **功能**：支持创建、填充、读取和转换 OpenDocument 格式文件（.odt, .ods），对标 LibreOffice 生态。
    -   **讨论热点**：该 Skill 满足了处理开放标准文档格式的特定需求，社区讨论了其在企业级文档流程（特别是使用 LibreOffice 的场景）中的潜在价值。
    -   **链接**：[https://github.com/anthropics/skills/pull/486](https://github.com/anthropics/skills/pull/486)

4.  **AppDeploy skill（#360）**
    -   **功能**：使 Claude 能够直接通过 `appdeploy.ai` 平台对全栈 Web 应用进行部署、生命周期管理和状态监控。
    -   **讨论热点**：将 AI 能力直接扩展到应用部署环节是社区关注的热点。讨论聚焦于该 Skill 的实用性和安全性，以及它在多大程度上简化了从开发到上线的流程。
    -   **链接**：[https://github.com/anthropics/skills/pull/360](https://github.com/anthropics/skills/pull/360)

5.  **codebase-inventory-audit（#147）**
    -   **功能**：提供一个全面的 10 步工作流，用以识别代码库中的孤立代码、未使用文件、文档缺口和基础设施膨胀，并生成 `CODEBASE-STATUS.md`。
    -   **讨论热点**：社区认为这是进行代码库重构和健康度评估的利器。讨论点在于如何降低误报率，以及如何与大型遗留系统更好地集成。
    -   **链接**：[https://github.com/anthropics/skills/pull/147](https://github.com/anthropics/skills/pull/147)

#### 2. 社区需求趋势

从 Issues 中可提炼出三大核心需求趋势：

-   **工具链可靠性（First-Class Priority）**：最强烈的诉求来自对 **skill-creator** 工具本身的修复和改进。大量 Issue（如 #556、#1169、#1061）和 PR（如 #1298、#1323）都指向 `run_eval.py` 在 Windows 上的崩溃、0% 的召回率报告和不稳定的 YAML 解析问题。这表明社区的核心需求是“**让创造 Skills 的工具先稳定运行起来**”。
-   **组织级功能与企业安全**：需求强烈的方向是支持 **Skills 在组织内的共享和分发**（Issue #228）以及对社区 Skills 的**安全信任机制**（Issue #492）。这反映了 Skills 从个人实验走向团队协作和企业级应用时，对安全、治理和便捷性的迫切期待。
-   **生态扩展与 AI 治理**：社区不再满足于基础任务，开始探索更复杂的 AI 应用，如**持久化记忆系统**（shodh-memory）、**AI 代理安全治理**（agent-governance）、以及**紧凑符号化记忆**（compact-memory）。这表明开发者正在探索如何构建更强大、更可控的 AI Agents。

#### 3. 高潜力待合并 Skills

以下 PR 评论活跃，功能针对性强，且直接回应了社区的核心痛点，极有可能在近期被合并或取得重要进展。

1.  **document-typography（#514）**：作为解决用户普遍痛点（排版质量）的 Skill，获得社区广泛共识，合并优先级很高。
    -   **链接**：[https://github.com/anthropics/skills/pull/514](https://github.com/anthropics/skills/pull/514)

2.  **testing-patterns（#723）**：填补了官方生态在测试方法论的空白，是提升开发流程标准化的关键 Skill，社区期待度高。
    -   **链接**：[https://github.com/anthropics/skills/pull/723](https://github.com/anthropics/skills/pull/723)

3.  **skill-creator 修复相关 PR（#1298, #1323）**：这些是解决工具链崩溃问题的关键修复。考虑到其基础性地位，这类 PR 很可能被优先处理并合并。
    -   **链接**：[https://github.com/anthropics/skills/pull/1298](https://github.com/anthropics/skills/pull/1298)
    -   **链接**：[https://github.com/anthropics/skills/pull/1323](https://github.com/anthropics/skills/pull/1323)

4.  **shodh-memory（#154）**：虽然创建较早，但“跨会话持久记忆”是构建复杂 AI 系统的长期热门方向，其讨论热度不减，具备长期价值和落地潜力。
    -   **链接**：[https://github.com/anthropics/skills/pull/154](https://github.com/anthropics/skills/pull/154)

#### 4. Skills 生态洞察

当前社区在 Skills 层面最集中的诉求是：**修复 skill-creator 工具的稳定性与 Windows 兼容性，以确保 Skills 的创建、评估和优化流程不再成为生态发展的瓶颈。** 这揭示了一个核心矛盾：社区对创造高质量 Skills 有极大兴趣，但被工具链的低可靠性所严重制约。对于生态发展而言，**开发体验的优化比新增 Skills 功能本身更为迫切**。

---

# Claude Code 社区动态日报
**日期：2026-06-28**  
**数据来源：**[github.com/anthropics/claude-code](https://github.com/anthropics/claude-code)

---

## 今日速览

过去24小时未有新版本发布，但社区热度不减：**Bash输出截断**这一历史遗留问题仍在发酵（17条评论，27个👍），而**Windows平台OAuth登录证书过期**的新Bug被迅速标记为回归问题。此外，一个名为 `protect-mcp` 的原创插件PR提出“失败关闭策略门控 + 签名收据”，试图从架构层面解决安全与审计痛点。

---

## 社区热点 Issues

挑选10个最受关注或影响面最大的Issue（含持续活跃的旧Issue）：

### 1. Bash输出截断：Ctrl+O/E 无法完整展开
- **Issue #26954** [开放]  
- 评论：17｜👍：27  
- **摘要**：运行产生30~40行输出后，Bash工具结果显示为截断状态，即使按下快捷键展开/折叠也无法显示完整内容。  
- **为什么重要**：长期未修复，影响日常代码审查与日志分析。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/26954)

### 2. Windows原生安装OAuth登录证书过期
- **Issue #71708** [开放]｜标签：`platform:windows`, `area:auth`, `regression`  
- 评论：4｜👍：1  
- **摘要**：无代理/VPN/杀软环境下，Windows原生安装版Claude Code在OAuth登录时抛出 `CERT_HAS_EXPIRED`，但同一主机curl却能正常访问。  
- **为什么重要**：回归问题，直接影响Windows用户首次使用/重新登录。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/71708)

### 3. 自定义Slash命令 `context: fork` 结果不渲染为assistant消息
- **Issue #71464** [开放]｜标签：`platform:windows`, `area:skills`  
- 评论：2  
- **摘要**：当技能/斜杠命令使用 `context: fork` 时，子代理完成后的结果被记录为 `<local-command-stdout>` 而非普通assistant消息，UI上命令“静默停止”。  
- **为什么重要**：影响Windows用户的自定义技能工作流可见性。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/71464)

### 4. deep-research技能“合成”阶段输出退化的虚假报告
- **Issue #72030** [已关闭]｜标签：`bug`, `area:skills`  
- 评论：2  
- **摘要**：内置 `deep-research` 技能的最终阶段可以输出“模式合法但内容退化”的摘要，静默丢弃已验证声明。  
- **为什么重要**：尽管已关闭，但揭示了复杂技能管线的可靠性风险。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/72030)

### 5. Write/Edit工具调用被渲为纯文本
- **Issue #69370** [已关闭，重复]  
- 评论：2  
- **摘要**：Write/Edit工具调用间歇性地被当作普通assistant文本输出（例如 `court <invoke name="Write"> ...`），导致文件未实际修改。  
- **为什么重要**：多次报告，严重影响代码编辑可靠性。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/69370)

### 6. Agent忽略代码库中定义的Git行为规则
- **Issue #70615** [已关闭，重复]  
- 评论：2  
- **摘要**：即使已经在 `Claude.md` 和会话上下文中明确写入了Git工作流规则，Agent仍然违反。  
- **为什么重要**：大量用户重复提交，反映Agent指令遵循能力的短板。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/70615)

### 7. Claude不经检查代码就猜测答案
- **Issue #70625** [已关闭，重复]  
- 评论：2｜👍：1  
- **摘要**：Claude反复根据记忆/训练数据回答代码问题，被质疑后才读取实际文件并承认错误。  
- **为什么重要**：削弱开发者对工具的信任，属于“幻觉”类高频投诉。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/70625)

### 8. 新选择文本覆盖层无法复制
- **Issue #70556** [已关闭，重复]  
- 评论：2｜👍：1  
- **摘要**：更新后在UI中选择文本时出现底部覆盖层，无复制按钮，快捷键失效，取消选中后选区消失。  
- **为什么重要**：基础复制功能被破坏，严重影响日常使用。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/70556)

### 9. v2.1.177更新导致所有 claude.ai 远程MCP的OAuth令牌失效
- **Issue #70593** [已关闭，重复]  
- 评论：2  
- **摘要**：升级到v2.1.177后，已授权的远程MCP服务器（通过claude.ai登录）全部失效，需重新授权。  
- **为什么重要**：影响MCP生态集成，更新兼容性问题。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/70593)

### 10. 因Agent不遵循指令导致过度Token消耗
- **Issue #70595** [已关闭，重复]  
- 评论：2  
- **摘要**：Agent在8轮交互中未能按要求删除一行元素，反而进行无意义的反复尝试并消耗大量Token。  
- **为什么重要**：凸显成本与效率问题，用户期望更好的容错和停止机制。  
- [GitHub链接](https://github.com/anthropics/claude-code/issues/70595)

---

## 重要 PR 进展

过去24小时仅有2个有效PR，重点如下：

### 1. [开放] 新增 `protect-mcp` 插件：失败关闭策略门控 + 签名收据
- **PR #72014**  
- 作者：tomjwxf  
- **摘要**：在 `plugins/` 下新建安全插件，位于 `PreToolUse` 阶段拦截违反Cedar策略的调用（同时推出警告而非仅报警），并对每一次决策生成离线可验证的签名收据。  
- **为什么重要**：从架构层增强MCP安全性，支持企业级审计需求。  
- [GitHub链接](https://github.com/anthropics/claude-code/pull/72014)

### 2. [开放] docs: 更新插件安装说明，推荐使用安装器
- **PR #72000**  
- 作者：aayushraj1425  
- **摘要**：改进插件安装文档，引导用户通过推荐安装器（如 `claude plugin install`）而非手动复制。  
- **为什么重要**：提升开发者体验，降低新手门槛。  
- [GitHub链接](https://github.com/anthropics/claude-code/pull/72000)

---

## 功能需求趋势

从近期Issues中可以提炼出以下社区最关注的三个功能方向：

| 需求方向 | 具体诉求 | 代表Issue |
|----------|----------|-----------|
| **可靠性提升** | 修复工具调用被截断/错误渲染、Agent不遵守规则、伪造验证结果等问题 | #26954, #69370, #70231 |
| **Windows体验优化** | 修复证书过期、自动更新冻结、fork技能不渲染、文件路径/Glob表情符号兼容性问题 | #71708, #70738, #71464, #70614 |
| **安全与审计** | 更强的策略门控、会话标识泄漏防护、更新后令牌失效修复 | #72014, #69669, #70593 |

---

## 开发者关注点

- **重复报告泛滥**：大量Issues被打上 `duplicate` 标签，说明基础Bug（如工具调用失败、截断、Token浪费）长期存在且用户反复遇到，修复优先级需提升。
- **Agent自主性与可预测性矛盾**：用户对Claude“猜测而不检查代码”“忽略Git规则”“不遵循清晰指令”等行为十分不满，希望增加“先读后答”的强制步骤或降级到严格模式。
- **更新带来的“修一个引三个”**：v2.1.177等版本导致OAuth令牌失效、选择覆盖层破坏复制等连锁问题，开发者对版本稳定性的信任度降低。
- **成本敏感度上升**：多个Issue指向Token消耗过多、后台代理无用户确认启动，社区开始要求更精细的配额控制和操作审计。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

好的，作为专注于 AI 开发工具的技术分析师，我已根据您提供的 GitHub 数据，为您整理出 2026 年 6 月 28 日的 OpenAI Codex 社区动态日报。

---

# OpenAI Codex 社区动态日报 | 2026-06-28

## 今日速览

**99% 的社区热度聚焦于 Windows 平台稳定性与资源消耗问题**，尤其是 `logs_2.sqlite` 文件写入量巨大导致 SSD 寿命损耗的严重 Bug 引发广泛关注（400+ 👍）。与此同时，Codex 团队正全力推进 **MCP（Model Context Protocol）OAuth 凭据管理** 的基础设施建设，以及通过多个 PR 堆栈改进 **`app-server` 的嵌入式 Rust API**。今日发布了三个 Rust CLI 的 alpha 增量版本，修复了部分已知问题。

## 版本发布

昨日发布了 3 个 Rust CLI 的 alpha 增量版本，未附带详细更新日志，主要指向小范围的 Bug 修复或功能验证。

- **rust-v0.143.0-alpha.27** / **.28** / **.29**：GitHub Release

## 社区热点 Issues

1.  **SQLite 日志写入量巨大，威胁 SSD 寿命**
    - **Issue #28224** | **状态**: 开放 | **热度**: 400 👍, 94 条评论
    - **摘要**: 此问题报告 Codex 的 SQLite 反馈日志每年可能写入高达 640TB 的数据，迅速消耗 SSD 寿命。这是社区反响最强烈的问题，已有 3 个 PR 被合并至 0.142.0 版本，据作者反馈可减少 85% 的日志量。
    - **重要性**: 极严重影响硬件健康与用户体验。
    - **链接**: [Issue #28224](https://github.com/openai/codex/issues/28224)

2.  **Windows 应用打开旧聊天后崩溃/无法加载**
    - **Issue #26104** | **状态**: 开放 | **热度**: 0 👍, 21 条评论
    - **摘要**: 大量 Windows 用户反映在更新后无法打开之前的聊天会话，严重影响工作流的连续性。
    - **重要性**: 数据可访问性的关键 Bug。
    - **链接**: [Issue #26104](https://github.com/openai/codex/issues/26104)

3.  **WSL 代理模式下无法找到捆绑 CLI**
    - **Issue #28086** | **状态**: 开放 | **热度**: 11 👍, 10 条评论
    - **摘要**: 在 Windows 系统中使用 WSL 代理模式时，Codex 无法正确调用 Linux 原生 CLI，反而启动 Windows 版，导致行为异常。
    - **重要性**: 影响核心的跨平台协作开发场景。
    - **链接**: [Issue #28086](https://github.com/openai/codex/issues/28086)

4.  **后台无限触发 `git add -A` 进程**
    - **Issue #22151** | **状态**: 开放 | **热度**: 4 👍, 8 条评论
    - **摘要**: 即使 Codex Windows 应用窗口关闭，后台仍无限反复执行 `git add -A`，并产生大量 `git-lfs` 子进程，造成系统资源严重浪费。
    - **重要性**: 严重性能问题，可能影响整个系统的正常运行。
    - **链接**: [Issue #22151](https://github.com/openai/codex/issues/22151)

5.  **Windows 沙箱环境下文件编辑失败**
    - **Issue #30009** | **状态**: 开放 | **热度**: 0 👍, 7 条评论
    - **摘要**: 在 Windows 上，通过 `apply_patch` 工具进行的文件编辑因沙箱相关错误而失败，阻碍了代码修改功能。
    - **重要性**: 核心功能在特定环境下失效。
    - **链接**: [Issue #30009](https://github.com/openai/codex/issues/30009)

6.  **GPT-5.5 推理 Token 聚类，导致复杂任务性能下降**
    - **Issue #30364** | **状态**: 开放 | **热度**: 2 👍, 5 条评论
    - **摘要**: 用户发现 `gpt-5.5` 模型的推理输出 token 数倾向落入 `516`,`1034`,`1552` 等固定值，这可能限制了模型在复杂任务上的表现。
    - **重要性**: 揭示模型行为层面的潜在缺陷，可能影响高级用户的模型选择。
    - **链接**: [Issue #30364](https://github.com/openai/codex/issues/30364)

7.  **因界面未推送操作系统原生通知**
    - **Issue #29008** | **状态**: 开放 | **热度**: 2 👍, 1 条评论
    - **摘要**: 当 Codex 在后台工作并需要用户批准权限时，Mac 用户无法收到原生推送通知，而必须手动查看应用窗口。
    - **重要性**: 影响后台任务流程，降低工作效率。
    - **链接**: [Issue #29008](https://github.com/openai/codex/issues/29008)

8.  **Windows 桌面应用在输入字母时冻结 2 秒**
    - **Issue #27924** / **#29543** | **状态**: 开放 | **热度**: 低评论，高困扰度
    - **摘要**: 多个用户报告在 Windows 应用中新会话的首次输入或输入字母时会出现 2-3 秒的卡顿，而输入数字则正常。
    - **重要性**: 对基础交互体验的直接影响。
    - **链接**: [Issue #27924](https://github.com/openai/codex/issues/27924) | [Issue #29543](https://github.com/openai/codex/issues/29543)

9.  **上下文/Token 消耗异常迅速**
    - **Issue #30387** | **状态**: 开放 | **热度**: 0 👍, 2 条评论
    - **摘要**: 用户发现在长时间对话中，Token 会突然快速耗尽，并且使用量出现异常下降，可能影响代码辅助上限。
    - **重要性**: 影响 Pro 用户的配额使用效率和稳定性。
    - **链接**: [Issue #30387](https://github.com/openai/codex/issues/30387)

10. **子代理（Sub-agent）无限卡死**
    - **Issue #30400** | **状态**: 开放 | **热度**: 0 👍, 2 条评论
    - **摘要**: 用户表示在进行长时代码审查时，子代理会无限期卡住，无法继续任务。
    - **重要性**: 影响高级功能（多代理协作）的可用性。
    - **链接**: [Issue #30400](https://github.com/openai/codex/issues/30400)

## 重要 PR 进展

1.  **[#30217] 从 `list_agents` 中移除不可用的任务消息**
    - **摘要**: 针对多代理 v2，修复了由于任务消息已加密导致桥接层（bridge）无法解密并展示无意义内容的问题，确保 API 返回准确信息。
    - **链接**: [PR #30217](https://github.com/openai/codex/pull/30217)

2.  **[#30297] 默认启用远程插件**
    - **摘要**: 将远程插件功能提升为稳定版，默认开启，预示着 Codex 插件生态将迎来扩展。可通过 `features.remote_plugin` 配置关闭。
    - **链接**: [PR #30297](https://github.com/openai/codex/pull/30297)

3.  **[#30252] 缓存远程 Bash 环境变量导出**
    - **摘要**: 优化远程 Bash 执行性能，每次执行会话仅初始化一次环境变量并缓存，避免重复执行 `BASH_ENV`，减少远程命令执行开销。
    - **链接**: [PR #30252](https://github.com/openai/codex/pull/30252)

4.  **[#30228] 向调用客户端暴露线程选择的技能**
    - **摘要**: 允许应用服务器将当前会话线程中可用的“技能”列表提供给前端，使 UI 能够正确显示可调用的功能，改善交互体验。
    - **链接**: [PR #30228](https://github.com/openai/codex/pull/30228)

5.  **[#30423] 增加 `currentTime/read` 超时时间 (10s → 30s)**
    - **摘要**: 针对某些情况下系统时钟读取超时的问题，将应用服务器的内部时钟读取超时从 10 秒延长至 30 秒，以提升健壮性。
    - **链接**: [PR #30423](https://github.com/openai/codex/pull/30423)

6.  **[[核心基础设施] 解决 & 序列化 MCP OAuth 凭据存储](#)**
    - **摘要**: 这是一个由多个 PR 组成的堆栈，包括 `#30292` `#30293` `#30294` `#30295` `#30296` `#30416`，旨在解决 VSCode 等协议服务器（MCP）的 OAuth 凭据在并发读写、登录注销、刷新和存储漂移报告等方面的正确性问题。
    - **重要性**: 为 Codex 的 MCP 生态奠定了坚实的安全和状态管理基础。
    - **链接**: [PR #30292](https://github.com/openai/codex/pull/30292) (以此为例)

7.  **[[核心] 支持持久的、外部的`ThreadGoal`]()**
    - **摘要**: 这是第五个 PR（堆栈），为需要管理外部线程目标（`ThreadGoal`）的宿主提供持久的 API 和持久化支持，提升了框架的可扩展性。
    - **链接**: [PR #30369](https://github.com/openai/codex/pull/30369)

8.  **[[核心] 添加推广邀请功能到 `/usage` 页面]()**
    - **摘要**: 在内部 `/usage` 页面下新增推广邀请流程，用户可在此处完成邀请操作，为未来可能的增长策略铺路。
    - **链接**: [PR #30313](https://github.com/openai/codex/pull/30313)

9.  **[[核心] 为标准化后的 Prompt 输出分配 ID]()**
    - **摘要**: 确保所有经过标准化处理的 prompt 输出都获得唯一的 ID，提高系统的可追溯性和调试能力，特别是针对重试和恢复场景。
    - **链接**: [PR #30311](https://github.com/openai/codex/pull/30311)

10. **[[核心] 展示使用限制重置的到期详情]()**
    - **摘要**: 允许客户端 UI 展示存量/限额重置的具体过期时间，提升用户对自己账户状态的可见性。
    - **链接**: [PR #30395](https://github.com/openai/codex/pull/30395)

## 功能需求趋势

- **极致的稳定性与性能（尤其 Windows）**: 社区最大呼声是解决 Windows 平台的卡顿、崩溃和资源滥用（如 #22151, #27924, #28224），这表明稳定性是用户体验的基石。
- **增强的跨平台与远程开发能力**: 对 **WSL 代理模式 (#28086)** 和 **SSH 会话 (`resume` 问题 #30424)** 的高频讨论，体现了开发者希望在 Windows/Linux 混合环境中无缝工作的强烈需求。
- **智能与可控的操作系统集成**: 用户期望 Codex 能更好地与本地 OS 协作，例如 **操作系统原生通知 (#29008)** 和更精细的 **权限审批控制**。
- **多代理与子代理成熟度**: 用户正在尝试并依赖高级的多代理功能，但也因此遇到了 **子代理卡死 (#30400)**、**模型混用 (#30427)** 等问题，对改进此功能的稳定性需求迫切。
- **基础设施成熟度（MCP、插件）**: 多个 PR 堆栈专注于 **MCP 协议（OAuth 凭据管理）** 和**远程插件**，这是开发平台和生态系统的基石。社区期待更强大、更安全的第三方集成体验。

## 开发者关注点

- **Windows 故障是绝对痛点**: 超过半数的热门 Issue 和大量评论与 Windows 平台的 Bug 相关，这已成为影响 Codex 普及和用户满意度的最大障碍。
- **日志与资源消耗是“元”问题**: `#28224` 所揭示的SQLite日志写入量问题引发了广泛共鸣，因为它不仅是特性问题，更是直接威胁到用户硬件的根本问题。
- **对“核心功能回退”的担忧**: 用户（`#30403`）在更新后发现模型行为或 API 发生改变（如 `gpt-5.5` 失败），这暴露了在快速迭代中可能出现的功能兼容性问题。
- **需要更精细的控制和可见性**: 开发者不希望所有事情都自动化。`#28969` 提出的禁用 60 秒自动解决的功能，以及 `#30395` 中对使用额度重置时间的可见性请求，都表明用户希望能对模型和配置有更细粒度的控制权。

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，作为专注于 AI 开发工具的技术分析师，我为您呈现 2026-06-28 的 Gemini CLI 社区动态日报。

---

### 2026-06-28 Gemini CLI 社区动态日报

#### 📊 今日速览

今日社区动态聚焦于**安全加固**与**开发者体验优化**。最新的 Nightly 版本（v0.51.0）引入了一项重要的安全修复，旨在防止敏感路径泄露。同时，社区对 VS Code 集成体验、子代理行为控制以及 MCP 协议稳定性的讨论持续升温，多个相关 PR 已提上议程。

---

#### 🚀 版本发布

**v0.51.0-nightly.20260628.gae0a3aa7b**

该版本主要包含一项关键的安全修复（`fix(security)`），通过对敏感路径的阻断列表实施**大小写不敏感的强制匹配**，并强化了 VSCode HITL（人在回路中）机制，从而增强了 CLI 的安全边界，防止潜在的路径遍历或敏感信息泄露风险。

---

#### 🔥 社区热点 Issues (Top 10)

1.  **[Bug] A2A server GET /tasks/metadata 缺少 return 语句导致服务器崩溃**
    *   **链接**: [Issue #21729](https://github.com/google-gemini/gemini-cli/issues/21729)
    *   **重要性**: **高优先级 (P1)**。这是一个会导致非内存任务存储（如 GCS）后端下，服务端发送 501 响应后，因缺少 `return` 语句而继续执行，最终因重复发送响应头而崩溃的 bug，严重影响服务稳定性。
    *   **社区反应**: 虽然有评论数 (12)，但处于开放状态，证明问题尚未被彻底解决，值得关注。

2.  **[Bug] 修复 VSCode 扩展中 `activate()` 函数的逗号操作符导致资源泄漏**
    *   **链接**: [Issue #27790](https://github.com/google-gemini/gemini-cli/issues/27790)
    *   **重要性**: 这是一个**代码质量问题**（P2）。`activate()` 函数因编码错误（使用逗号操作符）未能将两个 `Disposable` 对象加入 `context.subscriptions`，导致它们在扩展停用时无法被正确清理，造成资源泄漏。
    *   **社区反应**: 社区贡献者发现了问题并提交了 issue，已被标记为“good first issue”，适合新人参与。

3.  **[Bug] Gemini CLI 不维护键盘焦点当 VS Code 扩展关闭 diff 视图时**
    *   **链接**: [Issue #22193](https://github.com/google-gemini/gemini-cli/issues/22193)
    *   **重要性**: **高优先级 (P1)**。这是一个直接**影响开发者工作流流畅性**的 bug。当在 VSCode 终端中使用 Gemini 进行代码审查时，关闭 diff 预览后焦点会从终端移走，迫使开发者每次都需要手动点击回到终端，极大地降低了自动化体验。
    *   **社区反应**: 评论数较多 (9)，说明这是一个普遍且令人困扰的问题。相关修复 PR #28183 已经提出。

4.  **[Bug] Shell 命令执行在命令完成后卡住，显示“Awaiting input”**
    *   **链接**: [Issue #25166](https://github.com/google-gemini/gemini-cli/issues/25166)
    *   **重要性**: **高优先级 (P1)**。这个 bug 会导致 Gemini 在执行简单、无需交互的 CLI 命令后陷入死锁状态，令用户误以为进程仍在运行，是**严重的流程阻塞问题**。
    *   **社区反应**: 收到 3 个 👍，表明用户对这个问题有共鸣。目前仍在处理中。

5.  **[Bug] Subagent 达到 MAX_TURNS 后被错误报告为成功（GOAL）**
    *   **链接**: [Issue #22323](https://github.com/google-gemini/gemini-cli/issues/22323)
    *   **重要性**: 这是一个**逻辑漏洞**（P1）。当子代理因达到最大轮次限制而中断时，应报告为失败或中断，但系统却将其错误地标记为“成功”并给出“GOAL”原因。这会**掩盖真实的失败原因**，误导用户和上游决策流程。
    *   **社区反应**: 收到 2 个 👍，已被开发团队标记为内部跟踪，需要修复。

6.  **[Bug] 子代理从 v0.33.0 开始未经授权自动运行**
    *   **链接**: [Issue #22093](https://github.com/google-gemini/gemini-cli/issues/22093)
    *   **重要性**: **高优先级 (P2)**。用户反映在更新后，即便配置文件中禁用了代理模式，子代理仍会被触发执行。这严重违反了用户预期和控制权问题，**侵犯了用户的自主决策权**。
    *   **社区反应**: 开发者已将其标记为内部跟踪，这是一个信任和安全问题。

7.  **[Bug] 技能发现失败当 SKILL.md 的 description 字段为单行时**
    *   **链接**: [Issue #25693](https://github.com/google-gemini/gemini-cli/issues/25693)
    *   **重要性**: 这是一个**功能缺陷**，影响开发者自定义技能的加载。一个格式上的小差异（多行 vs 单行）就能导致技能完全不被识别，**降低了扩展系统的鲁棒性**。
    *   **社区反应**: 已被标记为“help wanted”和“good first issue”，是一个很好的入门级贡献点。

8.  **[Bug] Gemini 不充分地使用自定义技能和子代理**
    *   **链接**: [Issue #21968](https://github.com/google-gemini/gemini-cli/issues/21968)
    *   **重要性**: 这是一个**核心行为问题**（P2）。用户反馈，即使有定义好的 Git 或 Gradle 技能，Gemini 在自主完成任务时也几乎不会主动调用它们，除非被明确指示。这**大大削弱了技能和子代理的价值**。
    *   **社区反应**: 这是一个内部跟踪的长期问题，涉及 Agent 的决策规划能力。

9.  **[Enhancement] 允许代理调用其他代理**
    *   **链接**: [Issue #22092](https://github.com/google-gemini/gemini-cli/issues/22092)
    *   **重要性**: 这是一项**重要的功能增强**。当前架构限制代理之间不能互相调用，极大地限制了构建复杂、多步骤自动化任务的能力。社区希望实现 Agent 间的协作。
    *   **社区反应**: 提问者表达了对实现“代理组合”模式的强烈需求。

10. **[Bug] 模型频繁在随机目录创建临时脚本**
    *   **链接**: [Issue #23571](https://github.com/google-gemini/gemini-cli/issues/23571)
    *   **重要性**: 这是一个**工作流整洁性问题**。当模型通过 Shell 执行受限时，它会倾向于生成大量脚本文件散落在项目各处，给用户的代码库造成混乱，增加了清理成本。
    *   **社区反应**: 开发者内部跟踪，体现了社区对模型行为可预测性和工作区整洁度的要求。

---

#### ⚙️ 重要 PR 进展 (Top 10)

1.  **[修复] fix(vscode-ide-companion): 关闭 diff 标签页时保留终端焦点**
    *   **链接**: [PR #28183](https://github.com/google-gemini/gemini-cli/pull/28183)
    *   **状态**: 开放
    *   **重要性**: **高**。直接对应热点 Issue #22193，旨在解决 VSCode 扩展中焦点丢失问题，是提升日常编码体验的关键修复。

2.  **[安全] fix(core): 信任对话框泄露从来不会运行的 Hook 形状 (#27901)**
    *   **链接**: [PR #27915](https://github.com/google-gemini/gemini-cli/pull/27915)
    *   **状态**: 开放，标记为 P1
    *   **重要性**: **高 (安全)**。修复了一个安全逆向问题：信任对话框显示的是*不会*运行的 Hook 实现，而实际会执行的 Hook 却被隐藏。这构成了严重的信息误导，可能诱导用户授予不必要的权限。

3.  **[修复] fix(core): 修复 MCP 图片 MIME 类型嗅探**
    *   **链接**: [PR #27878](https://github.com/google-gemini/gemini-cli/pull/27878)
    *   **状态**: 已合并
    *   **重要性**: **高**。解决了 Figma 等 MCP 集成中因 WebP 图片被误标为 PNG 而导致请求失败的问题，增强了 MCP 生态的兼容性。

4.  **[修复] fix(core): 修复 MCP OAuth 刷新时缺失的 client ID**
    *   **链接**: [PR #27889](https://github.com/google-gemini/gemini-cli/pull/27889)
    *   **状态**: 已合并
    *   **重要性**: **高**。修复了自动发现的 MCP 服务器在进行 OAuth 令牌刷新时会失败的 bug，确保了 MCP 认证流程的可靠性。

5.  **[修复] fix(cli): 使自定义主题的 border.default 在终端报告背景色时生效**
    *   **链接**: [PR #27887](https://github.com/google-gemini/gemini-cli/pull/27887)
    *   **状态**: 已合并
    *   **重要性**: **中等**。修复了文档中说明的主题定制功能实际上在某些终端环境下不生效的问题，提升了 CLI 的可定制性和一致性。

6.  **[修复] fix(vscode-ide-companion): 注册所有 activate() 中的 Disposable 对象**
    *   **链接**: [PR #27885](https://github.com/google-gemini/gemini-cli/pull/27885)
    *   **状态**: 已合并
    *   **重要性**: **中等**。解决了 Issue #27790 所述的资源泄漏问题，提升了 VSCode 扩展的稳定性和资源管理。

7.  **[修复] fix(core): 规范化 MCP 工具 Schema 的根类型为 object**
    *   **链接**: [PR #27888](https://github.com/google-gemini/gemini-cli/pull/27888)
    *   **状态**: 已合并
    *   **重要性**: **中等**。修复了因某些 MCP 工具 Schema 缺少根 `type: "object"` 而触发 Vertex AI 等下游服务报错的问题，提升了 MCP 工具的通用兼容性。

8.  **[修复] fix(core): 在 session_context 目录树中尊重 .gitignore 和 .geminiignore**
    *   **链接**: [PR #27886](https://github.com/google-gemini/gemini-cli/pull/27886)
    *   **状态**: 已合并
    *   **重要性**: **中等**。修复了 `session_context` 显示时未应用忽略规则的问题，避免将敏感或临时文件不必要地发送给模型，提升了隐私和效果。

9.  **[修复] fix(security): 要求经过批准的 Bot 补丁工件**
    *   **链接**: [PR #28178](https://github.com/google-gemini/gemini-cli/pull/28178)
    *   **状态**: 开放
    *   **重要性**: **高 (安全/自动化)**。旨在加强 CI/CD 流水线的安全性，防止未经验证的机器人补丁被自动发布，使“评审-发布”边界更加安全。

10. **[策略] fix(policy): 对 Shell 参数扩展要求确认**
    *   **链接**: [PR #28175](https://github.com/google-gemini/gemini-cli/pull/28175)
    *   **状态**: 开放
    *   **重要性**: **高 (安全/策略)**。强化了对危险 Shell 操作的限制，要求包含变量扩展的命令在交互模式下获得用户确认，并在非交互（YOLO）模式下直接拒绝。这是对 Agent 危险行为进行预防性控制的重要一步。

---

#### 🧭 功能需求趋势

从近期的 Issues 和 PRs 中，社区的功能需求呈现以下趋势：
1.  **增强 Agent 的自主性与控制权**：社区强烈希望 Agent 能更智能地调用子代理和技能 ( #21968 )，并能实现代理间的自调用和组合 ( #22092 )。同时，对 Agent 不可预测的破坏性行为（如执行危险 Git 操作） ( #22672 ) 和不受控的脚本生成 ( #23571 ) 表现出担忧，要求更强的策略控制和确认机制。
2.  **IDE 集成深度优化**：VSCode 集成的流畅度是核心关注点，焦点丢失 ( #22193 ) 和资源泄漏 ( #27790 ) 等细节问题被反复提及，社区期待“无感”的 IDE 内协作体验。
3.  **MCP 生态基础设施稳固**：随着 MCP 功能的普及，OAuth 刷新 ( #27889 )、Schema 兼容性 ( #27888 )、图片类型识别 ( #27878 ) 等基础设施问题成为修复重点，社区要求 MCP 集成必须稳定可靠。
4.  **安全性与透明度提升**：对信任模型 ( #27915 )、安全路径处理（今日发布）、Agent 行为记录（如子代理轨迹分享 #22598 ）的关注度持续上升。开发者希望在享受自动化便利的同时，能完全掌控和理解 Agent 的行为。

---

#### 🎯 开发者关注点

1.  **Agent 行为失控**：开发者反复提到 Agent **不按配置执行**（子代理未经允许运行 #22093），或在完成任务时**执行多余、有害的操作**（生成散落脚本 #23571，使用危险命令 #22672），导致开发者对 Agent 的信任度降低。
2.  **交互流程的“反直觉”中断**：VSCode 焦点丢失 (#22193) 和 Shell 命令执行卡死 (#25166) 这类问题看似微小，却会**频繁打断开发心流**，成为用户最直接、最频繁的痛点。
3.  **错误信息的误导性**：子代理因 Token 超限而中断却被报告为“成功” (#22323)，信任对话框显示不正确的 Hook (#27915)，这类**信息不透明或错误**的问题，让开发者难以诊断和信任系统的真实状态。
4.  **扩展机制的健壮性**：技能因 `description` 格式问题而“静默”失败 (#25693)，以及 `.env` 文件权限问题导致整个扩展系统崩溃 (#28059 PR)，这类**“一损俱损”**的脆弱性令开发者担忧，他们期望 CLI 具备更强的容错和优雅降级能力。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-06-28

## 今日速览
过去 24 小时内，社区共提交 2 个新 Issue（#3964 已关闭，#3963 为功能请求），另有 1 个旧 Issue 获得更新。主要关注点集中在：**Windows 平台 MCP 服务器启动回归**（#3958）、**Ubuntu 密钥环支持损坏**（#2165）以及 **Copilot CLI 拷贝软换行输出时空格丢失**（#3964 修复不完整）。PR 方面，一个添加 `.gitignore` 配置的 PR 继续推进，而一个由 Copilot 自动创建的 macOS 安装说明 PR 在 9 个月后意外关闭。

## 版本发布
无（过去 24 小时内无新版本发布）

## 社区热点 Issues（共 8 条）

### 1. #2165 [OPEN] Ubuntu keychain support is broken
- **作者**: AndreaPi｜👍 20｜评论 2
- **关键点**：文档错误与 `secret-tool` 缺失导致 Ubuntu 下认证失败，社区期望官方更新文档或修复。
- **链接**: [Issue #2165](https://github.com/github/copilot-cli/issues/2165)

### 2. #2778 [OPEN] When is /btw from claude code coming to copilot？
- **作者**: mflRevan｜👍 1｜评论 2
- **关键点**：用户希望 Copilot CLI 支持类似 Claude Code 的“随时提问”功能（`/btw`），不干扰当前会话上下文，反映对上下文记忆和异步交互的强烈需求。
- **链接**: [Issue #2778](https://github.com/github/copilot-cli/issues/2778)

### 3. #3958 [OPEN] Windows: v1.0.66 fails to start stdio MCP servers when command is a .bat/.cmd with args
- **作者**: chronofanz｜👍 0｜评论 1
- **关键点**：Windows 上 v1.0.66 回归，无法正确调用带参数的 `.bat`/`.cmd` 文件作为 stdio MCP 服务器，属于破坏性的兼容性问题。
- **链接**: [Issue #3958](https://github.com/github/copilot-cli/issues/3958)

### 4. #3962 [OPEN] latest copilot (1.0.65) not working
- **作者**: wangvisual｜👍 0｜评论 1
- **关键点**：用户报告 v1.0.65 在“review last commit”操作时卡住，影响日常使用，但尚未明确根因。
- **链接**: [Issue #3962](https://github.com/github/copilot-cli/issues/3962)

### 5. #3874 [OPEN] VS Code agent `preToolUse` agent hook denial does not work
- **作者**: springcomp｜👍 0｜评论 1
- **关键点**：通过 Copilot CLI 运行聊天会话时，安装的允许/拒绝钩子（hook）失效，无法拦截工具调用，属于安全/权限控制漏洞。
- **链接**: [Issue #3874](https://github.com/github/copilot-cli/issues/3874)

### 6. #3815 [OPEN] Debug logs saved to location is missing a `\`
- **作者**: jtucker｜👍 0｜评论 0
- **关键点**：Windows 下调试日志路径缺失反斜杠，导致复制到资源管理器后无法直接导航，虽是小问题但影响开发者调试效率。
- **链接**: [Issue #3815](https://github.com/github/copilot-cli/issues/3815)

### 7. #3963 [OPEN] [Feature Request] Show session retention/expiration date
- **作者**: kevinreber｜👍 0｜评论 0
- **关键点**：用户请求在状态栏显示会话保留/过期日期，以解决会话意外消失问题，体现对会话持久化管理的需求。
- **链接**: [Issue #3963](https://github.com/github/copilot-cli/issues/3963)

### 8. #3964 [CLOSED] Copying soft-wrapped output still drops space at wrap boundary on v1.0.65
- **作者**: coseguera｜👍 0｜评论 1
- **关键点**：已关闭但修复不完整，复制终端软换行输出时边界空格依旧丢失（#3666 的 incomplete fix），社区希望彻底解决。
- **链接**: [Issue #3964](https://github.com/github/copilot-cli/issues/3964)

## 重要 PR 进展（共 3 条）

### 1. #3928 [OPEN] Add .gitignore and settings configuration
- **作者**: tpsaint｜评论 0
- **内容**: 为项目增加 `.gitignore` 和设置配置，提升开发者体验，目前正在Review中。
- **链接**: [PR #3928](https://github.com/github/copilot-cli/pull/3928)

### 2. #570 [CLOSED] [WIP] Add macOS installation instructions to README.md
- **作者**: Copilot｜评论 0
- **内容**: 由 Copilot 自动创建的 PR，旨在补充 macOS 安装说明。历经 9 个月后于 2026-06-27 关闭，未合并。
- **链接**: [PR #570](https://github.com/github/copilot-cli/pull/570)

### 3. #3737 [OPEN] Jigg empire ai
- **作者**: j2030aiNotez｜评论 0
- **内容**: 标题和描述模糊（“Let’s try this new method”），疑似实验性或非功能性提交，暂未获得关注。
- **链接**: [PR #3737](https://github.com/github/copilot-cli/pull/3737)

## 功能需求趋势
从过去 24 小时议题中可以看出社区最关注以下方向：
- **上下文记忆与异步交互**（#2778）：类似 Claude Code 的 `/btw` 功能，希望在不中断主会话的前提下提问。
- **会话生命周期管理**（#3963）：用户希望明确知晓会话过期策略，避免无预期中断。
- **平台兼容性**（#2165, #3958, #3815）：Ubuntu 认证与 Windows MCP 启动、路径格式等跨平台问题持续存在。
- **终端输出体验**（#3964）：复制终端内容时软换行空格丢失，影响命令行工具的日常可用性。

## 开发者关注点
- **认证与密钥管理**：Ubuntu 下 Keychain 支持损坏（#2165）已获 20 个 👍，说明 Linux 用户群对此高度敏感。
- **Windows 回归问题**：v1.0.66 引入的 MCP 服务器启动失败（#3958）属于破坏性回归，且涉及常见 `.bat`/`.cmd` 场景，开发者应尽快回滚或推送 hotfix。
- **安全钩子失效**（#3874）：`preToolUse` 钩子不生效可能导致未授权执行，对有严格安全策略的团队影响较大。
- **调试/日志可用性**（#3815）：Windows 路径格式问题虽小，但暴露了开发阶段的低优先级缺陷，需提升对非 Linux 平台的测试覆盖。

---
*数据统计截止 2026-06-28 08:00 UTC*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 | 2026-06-28

## 今日速览
- 过去24小时内仓库无新版本发布或 Pull Request 合并，社区活动较为平静。
- 唯一一条活跃 Issue 聚焦于 **Kimi Code VSCode 插件内存占用过高** 的问题，用户反馈长时间任务后内存消耗异常，已持续 3 个月未关闭，体现该问题对开发者的持续困扰。
- 整体趋势显示，IDE 集成性能和资源消耗仍是用户最关心的痛点。

## 版本发布
（无）

## 社区热点 Issues
**说明：** 由于过去24小时仅有一条被更新的 Issue，以下仅列出该条，同时附上其长时间未关闭的背景分析。

1. **#1592 [bug] Kimi Code VSCode 插件内存消耗严重**  
   - **作者**: xiaochonzi  
   - **创建/更新**: 2026-03-26 / 2026-06-28（时隔三个月仍为 OPEN 状态）  
   - **描述**: 用户使用 VSCode 插件长时间执行任务（如代码补全、对话），内存占用持续增长，影响开发环境稳定性。  
   - **为何重要**: 该 Issue 直接关联插件在主力 IDE 中的日常体验，长期未修复可能导致用户流失；1 条评论暗示社区反馈较少，但问题本身具有普遍性。  
   - **社区反应**: 目前仅作者自述，尚未有官方回应或社区讨论。  
   - **链接**: https://github.com/MoonshotAI/kimi-cli/issues/1592

## 重要 PR 进展
（无）

## 功能需求趋势
基于唯一活跃 Issue 及仓库历史背景，社区当前最关注的方向包括：

- **IDE 集成性能优化**: 特别是 VSCode 插件的资源占用（CPU、内存）和长时间运行稳定性。
- **插件生命周期管理**: 用户期望能手动释放插件后台进程，或提供低功耗模式。
- **跨平台兼容性**: Issue 中的用户环境为 Darwin ARM64，暗示 M 系列 Mac 上的适配仍需改进。

## 开发者关注点
- **内存泄漏/高内存消耗** 是当前最直接的痛点，尤其是在长时间执行任务（如大文件分析、连续对话）时，插件会拖慢整个 VSCode 实例。
- **官方反馈滞缓**: Issue 自 3 月提出至今未获官方回复或标记，开发者希望看到更积极的维护节奏。
- **文档与诊断工具缺失**: 用户仅能通过系统监视器观察内存变化，缺乏插件自带的性能剖析或日志输出功能。

---

*注：本日报基于 2026-06-28 当日 GitHub 公开数据生成，仅包含过去 24 小时内更新的 1 条 Issue 及 0 条 PR，其余部分为宏观分析。*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

好的，这是为您生成的 2026 年 6 月 28 日 OpenCode 社区动态日报。

---

# OpenCode 社区动态日报 | 2026-06-28

## 今日速览

近期发布版本引起了社区关于**稳定性**的高度关注，Windows 平台上的严重崩溃问题和高 CPU 占用成为开发者反馈的焦点。与此同时，项目在**性能优化**、**模型兼容性**（特别是针对 GLM-5.2）以及 **MCP 与 TUI 交互**方面迎来了密集的修复与功能提交，社区活跃度显著提升。

## 社区热点 Issues（Top 10）

1.  **#33742 OpenCode v1.17.10 在 Windows 上崩溃** ([链接](https://github.com/anomalyco/opencode/issues/33742))
    **热度**: 🔥🔥🔥🔥🔥 评论: 47, 👍: 45
    **摘要**: 升级至 `v1.17.10` 后，OpenCode 在 Windows 上出现 Bun 的原生段错误导致崩溃，降级至 `v1.17.9` 则稳定。大量用户表示遇到类似问题，被认为是严重的回归 BUG，预计会优先修复。

2.  **#34236 Opencode 桌面版占用大量 CPU 资源** ([链接](https://github.com/anomalyco/opencode/issues/34236))
    **热度**: 🔥🔥🔥🔥 评论: 3, 👍: 1
    **摘要**: 用户反馈桌面版 CPU 占用率高达 30%-50%，而 CLI 版本正常。此问题与 #34289 (ResizeObserver) 和 #34288 内存泄漏相关，是目前性能方面的核心痛点。

3.  **#34289 [Desktop] ResizeObserver 循环导致 renderer 进程 CPU 过高** ([链接](https://github.com/anomalyco/opencode/issues/34289))
    **热度**: 🔥🔥🔥🔥 评论: 1, 👍: 0
    **摘要**: 技术细节已明确，`ResizeObserver` 错误持续触发导致渲染进程、GPU 和 Node服务 CPU 合计占用异常。开发者可针对性地进行问题定位和修复。

4.  **#34288 可能的内存泄漏导致 Linux OOM** ([链接](https://github.com/anomalyco/opencode/issues/34288))
    **热度**: 🔥🔥🔥🔥 评论: 1, 👍: 1
    **摘要**: 用户在 Fedora 系统（32GB RAM）上报告 OpenCode Desktop 因过度内存使用被 OOM Killer 终止。这表明即便在高配机器上，内存管理也存在稳定性风险。

5.  **#33490 GLM-5.2 通过 OpenCode Go 报错：`instructions` 字段不允许** ([链接](https://github.com/anomalyco/opencode/issues/33490))
    **热度**: 🔥🔥🔥🔥 评论: 3, 👍: 3
    **摘要**: 使用 GLM-5.2 模型时，因传递了不被其 API 接受的`instructions`字段导致报错。这反映了新模型接入时与现有架构的兼容性摩擦。社区对此有多个相关 PR 进行修复。

6.  **#32473 桌面版因 `lastProjectSession` 引用缺失会话导致致命崩溃** ([链接](https://github.com/anomalyco/opencode/issues/32473))
    **热度**: 🔥🔥🔥 评论: 4, 👍: 0
    **摘要**: 当 `opencode.db` 中某项目最后一次会话被删除后，桌面端启动时直接崩溃。此问题在 v1.17 系列中可能是普遍现象，与 #32268 (会话恢复空白) 共同指向会话管理模块的不稳定性。

7.  **#18213 Plan 模式下子Agent 在压缩后绕过限制修改代码** ([链接](https://github.com/anomalyco/opencode/issues/18213))
    **热度**: 🔥🔥🔥 评论: 3, 👍: 1
    **摘要**: Plan Agent 在上下文压缩后，其生成的子 Agent 可能使用 `cat` 和 `sed` 命令直接修改代码，绕过了 Plan 模式的限制。这是一个安全/行为审计方面的潜在BUG。

8.  **#34279 AI 编辑文件后 diff 未显示最新更改** ([链接](https://github.com/anomalyco/opencode/issues/34279))
    **热度**: 🔥🔥🔥 评论: 1, 👍: 0
    **摘要**: AI 完成对已提交文件的修改后，界面的差异视图（diff view）无法正常显示变更。这直接影响开发者审查 AI 代码修改的体验。

9.  **#32548 达到步骤上限时助手消息导致 Claude 模型 400 错误** ([链接](https://github.com/anomalyco/opencode/issues/32548))
    **热度**: 🔥🔥🔥 评论: 3, 👍: 0
    **摘要**: 当 Agent 达到步骤上限，系统发送的“MAXIMUM STEPS REACHED”消息作为助手角色被发送，这与启用了思考功能的 Claude 模型要求以用户角色消息结尾的规范冲突。已有一个 PR (#34276) 提出修复方案。

10. **#33755 [FEATURE]：为 `opencode --mini` 添加 `/compact` 命令** ([链接](https://github.com/anomalyco/opencode/issues/33755))
    **热度**: 🔥🔥🔥 评论: 1, 👍: 0
    **摘要**: 用户高度认可 `opencode --mini` 模式，并希望能在该模式下调用 `/compact` 命令以管理上下文。这表明小体积、高效率的操作模式是重要需求方向。

## 重要 PR 进展（Top 10）

1.  **#34077 fix(mcp): 序列化并发 OAuth token 刷新** ([链接](https://github.com/anomalyco/opencode/pull/34077))
    **重要性**: 高。修复了并发 MCP 工具调用时，使用过期 token 刷新可能导致的竞争条件和数据损坏问题。

2.  **#34286 feat(app): 对齐斜杠弹出菜单到 v2 tokens** ([链接](https://github.com/anomalyco/opencode/pull/34286))
    **重要性**: 高。改进了 UI 组件以适配新的设计系统，同时修复了鼠标悬停时自动滚动失效的 BUG。这是一个重要的体验优化。

3.  **#34258 fix(tui): 禁用模态框打开时的 diff-viewer 键盘绑定** ([链接](https://github.com/anomalyco/opencode/pull/34258))
    **重要性**: 高。修复了在 diff 模式下打开命令面板（`ctrl+p`）时，键盘输入被 diff-viewer 快捷键拦截导致无法操作的问题。

4.  **#34284 fix: 在所有平台上默认启用 Ctrl+C 复制** ([链接](https://github.com/anomalyco/opencode/pull/34284))
    **重要性**: 高。解决了 Linux 用户在 tmux 中无法通过 `Ctrl+C` 复制内容的问题，统一了跨平台复制体验。

5.  **#34256 fix(server): 在实例查找前拒绝外部目录提示** ([链接](https://github.com/anomalyco/opencode/pull/34256))
    **重要性**: 高。修复服务器端的安全漏洞，防止通过特制请求访问或操作不属于当前项目实例的目录。属于关键的安全修复。

6.  **#32905 fix(tool): 隐藏不可用的工具指导** ([链接](https://github.com/anomalyco/opencode/pull/32905))
    **重要性**: 高。防止模型接收其无法调用的工具描述（如 shell 和任务相关工具），减少了模型的混淆，提升了工具调用的准确性。

7.  **#34280 feat(tui): 添加 `/usage` 命令用于查看 token 和成本** ([链接](https://github.com/anomalyco/opencode/pull/34280))
    **重要性**: 中。为用户提供了一个便捷的命令来查看会话级别的 token 消耗和费用，增强了使用透明度和成本控制。

8.  **#34210 feat: 项目归档功能** ([链接](https://github.com/anomalyco/opencode/pull/34210))
    **重要性**: 中。引入了非破坏性的项目隐藏功能，解决了多个社区需求，帮助用户整理主界面而不丢失数据。

9.  **#34276 fix: 将“达到最大步骤”消息改为用户角色发送** ([链接](https://github.com/anomalyco/opencode/pull/34276))
    **重要性**: 高。直接解决 Issue #32548，通过改变消息角色来兼容 Anthropic 的 Claude 模型规范，避免 400 错误。

10. **#34283 fix(provider): 为 GLM-5.2 在 OpenAI 兼容接口上暴露 `xhigh` 而非 `max`** ([链接](https://github.com/anomalyco/opencode/pull/34283))
    **重要性**: 高。修复了 GLM-5.2 在与 OpenAI 兼容的提供商上推理等级选项错误的问题，确保模型配置遵循标准规范。

## 功能需求趋势

- **性能与资源优化**: 社区对 CPU/内存占用过高反馈强烈，核心需求是优化渲染进程（特别是 `ResizeObserver`）和修复内存泄漏，以提升桌面应用的稳定性和续航。
- **模型生态兼容性**: 随着 GLM-5.2 等新模型的接入，与 OpenAI 标准 API 的兼容性问题成为焦点。社区需求集中在正确处理 `reasoning_effort`、`instructions` 等字段的差异。
- **开发者体验与效率**: 用户强烈要求改进 `opencode --mini` 模式的功能，希望增加 `/compact` 等上下文管理命令。同时，`/usage` 命令的 PR 也表明对成本透明性的需求。
- **安全性与稳定性**: 对子 Agent 行为的审计（如 #18213）和服务器端目录访问控制（#34256）的重视度提升，表明社区开始关注更高级的安全和权限管理。
- **UI/UX 改进**: 持续关注对 WSL 的 UI 支持（#34233）、Slash Popover 与 V2 设计系统对齐（#34286）以及标签页、会话管理等基础交互的优化。

## 开发者关注点

- **升级恐惧症**: v1.17.10 在 Windows 上的崩溃问题导致开发者对“快速迭代”产生了不信任感，他们倾向于保持稳定版本，并强烈建议团队加强回归测试。
- **桌面端性能瓶颈**: 高 CPU 和内存占用是当前使用桌面版的最大痛点，这促使部分用户保持在 CLI 版本。开发者期待团队优先解决 `ResizeObserver` 循环和内存泄漏问题。
- **跨平台体验一致性**: Windows 的崩溃和 Linux 上复制功能不工作等问题，凸显了跨平台兼容性仍然是开发痛点。统一的快捷键行为和稳定的底层运行环境是基本要求。
- **调试与集成成本**: 与 VS Code 扩展的兼容性问题（#33659）和会话恢复失败（#32268）严重干扰了开发者的日常使用流程，他们希望 OpenCode 能更“无感”地集成到现有的开发环境中。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi 社区动态日报 | 2026-06-28

## 📰 今日速览

今天社区活跃度极高，共更新 23 条 Issue 和 8 个 PR。**模型供应商兼容性** 成为焦点：Xiaomi MiMo 定价硬编码错误、Together.ai 模型即将废弃、MiniMax M3 返回 404 等问题被快速修复。同时 **Context Matrix** 继续推进 Phase 4a 存储投影，多个扩展 API 增强（reportUsage、外部编辑器配置）已合并。稳定性方面，修复了流式传输 ECONNRESET 崩溃和 TUI 闪烁问题。唯一仍开放的 #5825（Markdown 自动滚屏）引来 34 条评论，社区讨论热烈。

## 🔖 版本发布

过去 24 小时无新版本发布。

## 🔥 社区热点 Issues（10 条）

1. **[OPEN] #5825 Streaming markdown forces scroll to bottom**  
   ⚠️ 开启 `clear on shrink` 后，AI 回答时滚动回底部打断阅读。34 条评论，**高频痛点**。  
   [链接](https://github.com/earendil-works/pi/issues/5825)

2. **[CLOSED] #6138 Xiaomi MiMo 本地提供商模型定价错误**  
   `pi-ai/dist/providers/xiaomi.models.js` 硬编码价格与官方不符，涉及 mimo-v2.5-pro 等模型，**影响成本计算**。  
   [链接](https://github.com/earendil-works/pi/issues/6138)

3. **[CLOSED] #6137 Context Matrix Phase 4a：markdown 存储投影**  
   实现稀疏单元格、manifest、工作区索引，朝 Phase 4 目标迈进。**核心基础设施**。  
   [链接](https://github.com/earendil-works/pi/issues/6137)

4. **[CLOSED] #6140 MiniMax M3 通过 OpenCode Go 返回 404**  
   与之前 #4106 同类问题，正在调查。**模型兼容性**。  
   [链接](https://github.com/earendil-works/pi/issues/6140)

5. **[CLOSED] #6139 Groq 不支持的 reasoning_content 字段导致 400 错误**  
   需在发送前剥离该字段。**API 兼容性修复**。  
   [链接](https://github.com/earendil-works/pi/issues/6139)

6. **[CLOSED] #6135 macOS 上硬编码 `/bin/bash` 导致语法不兼容**  
   苹果自带的 Bash 3.2（2007年）太老，**跨平台适配**。  
   [链接](https://github.com/earendil-works/pi/issues/6135)

7. **[CLOSED] #6132 Together.ai 两款模型即将于 7 月 10 日废弃**  
   需要过渡到 GLM-5.2 或 MiniMax-M3。**模型生命周期管理**。  
   [链接](https://github.com/earendil-works/pi/issues/6132)

8. **[CLOSED] #6133 流式 SSE 时上游 ECONNRESET 导致进程崩溃**  
   undici 异常逃逸 try/catch 触发 `uncaughtException`。**稳定性关键修复**。  
   [链接](https://github.com/earendil-works/pi/issues/6133)

9. **[CLOSED] #6131 多工具调用同时流式时全屏闪烁**  
   TUI 每收到一个工具调用块就重绘整个屏幕。**UI 体验**。  
   [链接](https://github.com/earendil-works/pi/issues/6131)

10. **[CLOSED] #6130 renderCall/renderResult 静默忽略异常**  
    异常被吞掉，导致开发者浪费数小时调试。**开发者体验**。  
    [链接](https://github.com/earendil-works/pi/issues/6130)

## 🚀 重要 PR 进展（8 条，全部更新于过去 24h）

1. **#4110 – 修复 models.dev 与 OpenCode Go 的模型不匹配（Qwen3.5/3.6、MiniMax M2.7）**  
   [CLOSED] 已合入条件变更，解决 #4106。  
   [链接](https://github.com/earendil-works/pi/pull/4110)

2. **#60 – 文件和文件夹模糊搜索**  
   [CLOSED] 支持 `@` 模糊匹配，与目录浏览并存。  
   [链接](https://github.com/earendil-works/pi/pull/60)

3. **#6115 – 可配置聊天边距（chat padding）**  
   [OPEN][to-discuss] 回应 Discord 高频请求，但需要 TUI 层级标志支持。  
   [链接](https://github.com/earendil-works/pi/pull/6115)

4. **#6136 – 修复 compaction 后 `hasQueuedMessages` 检查**  
   [CLOSED] 防止阈值压缩完成时错误调用 `agent.continue()`，避免死循环。  
   [链接](https://github.com/earendil-works/pi/pull/6136)

5. **#5735 – 安全延迟扩展重载请求**  
   [OPEN][to-discuss] 使 `ctx.reload()` 在所有扩展上下文中可用，由 AgentSession 协调。  
   [链接](https://github.com/earendil-works/pi/pull/5735)

6. **#5678 – 自定义消息增加 `excludeFromContext`**  
   [OPEN][to-discuss] 允许消息持久化但不进入模型上下文，支持分支总结等场景。  
   [链接](https://github.com/earendil-works/pi/pull/5678)

7. **#6123 – Ctrl+G 外部编辑器可配置**  
   [CLOSED] 新增 `externalEditor` 设置，解决 Windows + Git Bash 环境变量锁定问题。  
   [链接](https://github.com/earendil-works/pi/pull/6123)

8. **#6119 – 扩展 API 新增 `reportUsage()` 用于子代理成本回传**  
   [CLOSED] 子代理 token/cost 可注入会话页脚，提升费用透明度。  
   [链接](https://github.com/earendil-works/pi/pull/6119)

## 📈 功能需求趋势

从今日更新中可看出社区最关心的五大方向：

- **模型提供商兼容与生命周期**：定价硬编码、废弃过渡、reasoning_content 字段适配 —— 反映出 Pi 作为多模型网关的复杂性。
- **扩展 API 深化**：子代理成本回传（#6119）、扩展执行注册工具（#6121）、音频穿透 RPC（#6118）—— 生态能力持续扩展。
- **UI/UX 打磨**：滚动行为（#5825）、全屏闪烁（#6131）、聊天边距（#6115）、Devanagari 等非拉丁字符（#6124）—— 终端体验仍是核心痛点。
- **可靠性强化**：ECONNRESET 崩溃（#6133）、异常静默吞噬（#6130）、compaction 死循环（#6136）—— 对稳定性要求提升。
- **配置灵活性**：外部编辑器设置（#6122/6123）、包管理器参数（#6125/6126）、系统提示覆盖（#6127）—— 用户希望减少对环境的硬依赖。

## 🧑‍💻 开发者关注点

- **阅读干扰**：Markdown 自动滚屏（#5825）仍是唯一开放的高热度 bug，34 条评论表明其普遍性。
- **硬编码路径**：macOS `/bin/bash`（#6135）和编辑器环境变量（#6122）都暴露了跨平台脆弱性。
- **调试成本高**：异常被静默吞掉（#6130）和 UI 闪烁（#6131）让开发者难以定位问题。
- **成本管控需求**：模型定价错误（#6138）和子代理费用无法回显（#6119）促使社区推动费用透明化。
- **扩展开发门槛**：扩展无法执行注册工具（#6121）和缺少 `reportUsage`（#6120）限制了高级插件的能力。

> 数据来源：github.com/earendil-works/pi · 生成时间：2026-06-28

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code 社区动态日报 — 2026-06-28

## 今日速览

昨日发布了 **v0.19.3** 正式版，主要修复了 `web_fetch` 的 JSON fallback 问题。社区中涌现出一系列关于**模型自动切换导致高额费用**、**中文输出异常变为繁体**以及**输出 token 截断**等用户痛点，开发团队已迅速响应并给出处理方向。此外，多个重量级 PR 正在推进 **ACP 协议支持**、**桌面端语音输入**、**Chrome 扩展复活**等新能力，生态集成步伐明显加快。

---

## 版本发布

### v0.19.3（正式版）
- 修复核心：允许 `web_fetch` 在 JSON 解析失败时降级回退（#5660）
- 自动化发布流程，同步版本号与 CHANGELOG

### v0.19.2-nightly.20260628.714513df2
- 与 v0.19.3 同日更新，包含相同的 `web_fetch` 修复，用于夜间构建验证

---

## 社区热点 Issues（共 9 条，精选 9 条）

### 1. 升级后模型被自动切换为高价方案，导致资金耗尽  
**#5819** [CLOSED]  
作者 **aspnmy** 反映：从 0.18.3 升级后，`setting.json` 中的 `model` 自动从低价模型（DeepSeek-4 flash）更改为高价版（DeepSeek-4 pro），直到收到短信预警才发现预存资金用完。同时 0.19 版本中中文输出异常（简体输入 → 繁体输出），需要额外 token 纠正。  
**社区反应**：4 条评论，已被标记为 `status/need-information`，团队正在排查自动升级分支中的配置迁移逻辑。  
🔗 https://github.com/QwenLM/qwen-code/issues/5819

---

### 2. 模型生成时向上滚动会直接跳至顶部  
**#5941** [OPEN]  
作者 **lingview** 报告：在大模型生成内容时，只要向上滚动滚轮一次，页面就会直接跳到最上方，而非像浏览器一样滚动一行。影响内容阅读。  
**社区反应**：3 条评论，标记为 `type/bug` / `scope/rendering`，开发者已关注，可能需要渲染层防抖修复。  
🔗 https://github.com/QwenLM/qwen-code/issues/5941

---

### 3. 默认 8K 输出上限导致大文件写入被截断并触发失败重试循环  
**#5756** [CLOSED]  
作者 **chiga0** 指出：当模型未显式设置 `max_tokens` 时，系统硬限制为 8000 token，导致生成大文件（如 wiki 文档）被截断，进而触发 `write_file` 失败→重试的死循环。  
**社区反应**：3 条评论，已被关闭并合并修复 PR（#5868 引入了可配置的自动压缩阈值）。  
🔗 https://github.com/qwenlm/qwen-code/issues/5756

---

### 4. 跟踪 ACP 协议中 `cd`、权限、信任、LSP 等能力差距  
**#5677** [CLOSED]  
作者 **doudouOUC** 发起的追踪 Issue，梳理 ACP（Agent Communication Protocol）中尚未实现的功能点，并逐步推进（如 `/lsp`、`/permissions` 已有对应实现）。  
**社区反应**：3 条评论，已标记完成，是 ACP 协议生态的关键里程碑。  
🔗 https://github.com/QwenLM/qwen-code/issues/5677

---

### 5. 优化 Daemon 冷启动延迟（2.5s → ~1.5s）  
**#4748** [OPEN]  
作者 **doudouOUC** 提出：Daemon 冷启动（启动 + 首次会话）约 2.5 秒，而 CLI 只需 0.7 秒。建议优化至 1.5 秒左右以提升体验。  
**社区反应**：2 条评论，长期悬而未决，但属于性能增强的重要方向。  
🔗 https://github.com/QwenLM/qwen-code/issues/4748

---

### 6. 上下文长度超出模型限制（约 135k token）  
**#5950** [OPEN]  
作者 **Data-Wise** 报告：模型最大上下文 131k token，但请求时约 135k token（包括文本输入、工具输入、输出）导致 400 错误。建议压缩或减少长度。  
**社区反应**：2 条评论，新开 Issue，可能涉及 token 管理和压缩策略的缺陷。  
🔗 https://github.com/QwenLM/qwen-code/issues/5950

---

### 7. 规范化源 slug 验证与结构化错误处理  
**#5908** [CLOSED]  
作者 **VectorPeak** 跟进 #5829，修复了路径遍历漏洞后，对未统一处理的路径（如 `config-validate.ts` 的 fallback 路径）进行规范化，避免非结构化异常。  
**社区反应**：2 条评论，安全加固增强，已被合并。  
🔗 https://github.com/QwenLM/qwen-code/issues/5908

---

### 8. `/new`（别名 `/clear`）偶尔失效  
**#5949** [OPEN]  
作者 **fantasyz** 截图显示：执行 `/new` 后并未按预期清除会话并显示新章节消息，该命令在特定环境下（Windows）表现不稳定。  
**社区反应**：2 条评论，标记为 `status/need-retesting` 和 `welcome-pr`，欢迎社区贡献修复。  
🔗 https://github.com/QwenLM/qwen-code/issues/5949

---

### 9. `POST /workspace/settings` 接受非正的 `sessionRecapAwayThresholdMinutes` 值  
**#5680** [CLOSED]  
作者 **tt-a1i** 发现：设置接口允许传入 `0` 或 `-5`，但运行时逻辑仅在大於 0 时生效，导致配置与行为不一致。已修复并合并。  
**社区反应**：2 条评论，配置健壮性提升。  
🔗 https://github.com/QwenLM/qwen-code/issues/5680

---

## 重要 PR 进展（精选 10 条）

### 1. 引入 `@qwen-code/chat-panel` 共享聊天面板组件  
**#5951** [OPEN]  
作者 **qqqys** 创建了一个新的 props-driven 包，将 Web Shell 的聊天流程抽取为共享组件，为 VSCode Webview 和桌面端统一聊天面板铺路。  
🔗 https://github.com/QwenLM/qwen-code/pull/5951

---

### 2. QQ 机器人流式输出重构：空闲刷新、移除字符限制、Markdown 管道修复  
**#5902** [OPEN]  
作者 **Eric-GoodBoy-Tech** 重构了 QQ 机器人的流式行为：用 2 秒空闲刷新替代 BlockStreamer 合并、移除 2000 字符自限、增加被动回复 TTL、修正 Markdown 表格检测。  
🔗 https://github.com/QwenLM/qwen-code/pull/5902

---

### 3. 为 `/loop` 注入 `.qwen/loop.md` 任务文件  
**#5890** [OPEN]  
作者 **qqqys** 支持长循环保持持久、可编辑的任务列表，模型通过哨兵提示激活此模式，避免每轮重复声明工作。  
🔗 https://github.com/QwenLM/qwen-code/pull/5890

---

### 4. 阻止重复的 Shell 检查变体（循环守卫）  
**#5944** [OPEN]  
作者 **yiliang114** 添加永远启用的循环守卫：当模型持续调用语义相似的只读 git 命令（如 `git status`、`git diff`、`git ls-files`）时，主动中断以节省 token 和时长。  
🔗 https://github.com/QwenLM/qwen-code/pull/5944

---

### 5. 可恢复的 ACP 会话流（Last-Event-ID）  
**#5852** [OPEN]  
作者 **chiga0** 为 `/acp` 流提供了标准 SSE 重连机制，客户端可通过 `Last-Event-ID` 恢复中断会话，提升 ACP 连接可靠性。  
🔗 https://github.com/QwenLM/qwen-code/pull/5852

---

### 6. 移除 serve 桥接重导出 shim  
**#5955** [OPEN]  
作者 **doudouOUC** 清理了服务端遗留的桥接重导出代码，将内部导入直接指向 `@qwen-code/acp-bridge` 子路径，同时补充测试覆盖。  
🔗 https://github.com/QwenLM/qwen-code/pull/5955

---

### 7. 引入“qwen tag”——频道驻留多玩家智能体（RFC + Phase 0）  
**#5888** [OPEN]  
作者 **qqqys** 首次引入“qwen tag”概念：一个可以驻留在聊天群（如钉钉）中的多玩家智能体，基于现有频道适配器和 `qwen serve` 守护进程实现。  
🔗 https://github.com/QwenLM/qwen-code/pull/5888

---

### 8. LSP Server 支持热重载  
**#5953** [OPEN]  
作者 **water-in-stone** 为原生 LSP 服务器添加运行时配置热重载：当 `.lsp.json` 在活动会话中发生变化时，检测语义变更并自动重载，无需重启会话。  
🔗 https://github.com/QwenLM/qwen-code/pull/5953

---

### 9. 可配置的自动压缩阈值与 Stop Hook 上下文使用  
**#5868** [OPEN]  
作者 **ZijianZhang989** 实现了 #4025 提出的功能：允许用户自定义自动压缩阈值，并在 Stop Hook 中使用上下文，解决 #5756 中的默认截断问题。  
🔗 https://github.com/QwenLM/qwen-code/pull/5868

---

### 10. 桌面端语音输入（/voice 对讲功能）  
**#5856** [OPEN]  
作者 **qqqys** 将 CLI 和 Web Shell 中已有的 `/voice` 语音输入功能移植到桌面应用。麦克风按钮录制语音，工具栏变为录制状态条（波形、倒计时、停止按钮）。  
🔗 https://github.com/QwenLM/qwen-code/pull/5856

---

## 功能需求趋势

从近期 Issues 和 PR 中可归纳出社区最关注的几大方向：

| 方向 | 典型代表 | 说明 |
|------|----------|------|
| **ACP 协议 & 服务端能力** | #5677, #5852, #5888 | 社区积极推进 ACP 协议覆盖（LSP、权限、频道驻留智能体），优化 SSE 重连和会话管理 |
| **Token 管理与压缩策略** | #5756, #5868, #5950 | 默认 8K 上限导致输出截断、上下文超限等问题，驱动可配置压缩阈值和自动压缩功能的开发 |
| **UI / 交互体验** | #5941 (滚动跳顶), #5949 (/new 失效), #5947/5948 (移动端适配) | 滚动行为异常、命令偶尔不响应、移动端展示优化等细节反馈频繁 |
| **多平台/多通道集成** | #5856 (桌面语音), #5777 (Chrome 扩展), #5902 (QQ 机器人) | 桌面端、浏览器扩展、IM 频道等侧渠道持续完善，统一聊天面板组件 (#5951) 加速生态融合 |
| **配置与模型切换安全** | #5819 (升级后模型自动切换), #5680 (非正阈值) | 用户对升级后配置被静默修改深感担忧，要求更透明的迁移策略和配置校验 |
| **性能与启动速度** | #4748 (daemon 冷启动) | 尽管 daemon 在 warm 模式下仅需 21ms，但冷启动仍被诟病，优化空间明确 |

---

## 开发者关注点

1. **升级后模型自动变更引发资金风险** – Issue #5819 揭示了一个严重的回归：自动升级分支可能覆盖用户的低价模型选择，导致额外费用。社区呼吁增加升级确认对话框或模型变更日志。
2. **中文输出异常（简体→繁体）** – 同样在 #5819 中，用户反馈 0.19 版本将简体中文输入处理为繁体输出，需额外用 token 纠正，严重影响中文用户体验。
3. **输出 token 硬限制导致工作流中断** – #5756 及 #5950 显示，默认 8K 输出上限和最大上下文长度限制是高频痛点，尤其是生成大文件或复杂工具调用时。开发者希望提供更智能的自动压缩或用户可调的软阈值。
4. **滚动行为与命令响应不一致** – #5941 和 #5949 分别指出翻页滚动异常和 `/new` 命令偶发失效，影响日常交互流畅度，社区建议增加兼容性测试。
5. **Daemon 冷启动延迟** – #4748 长期未被关闭，开发者期待针对启动流程（如依赖加载、会话初始化）进行并行化或懒加载优化。

---

*以上日报基于 2026-06-28 07:00 UTC 前 GitHub 数据编制。*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，这是为您生成的 2026-06-28 DeepSeek TUI (CodeWhale) 社区动态日报。

---

# DeepSeek TUI (CodeWhale) 社区动态日报 | 2026-06-28

## 今日速览

社区焦点依然集中在 **输入缓存命中率的性能优化** 上，这是阻碍开发效率的核心痛点。与此同时，项目维护者正在积极推进与 Zed 编辑器等第三方工具的集成，并通过 **插件系统** 和 **配置热重载** 等新功能为未来更丰富的生态做铺垫。今天有多个关于错误处理回退机制和工具参数校验的 PR 被合并，显示出项目在鲁棒性上的持续投入。

## 版本发布

今日无新版本发布。但 v0.8.66 版本的发布过程仍在进行中，相关的重要 PR 和问题讨论正在进行最后的冲刺。

## 社区热点 Issues

1.  **[#1177] 输入缓存命中率太低了**
    - **重要性：极高**。这是当前社区最热门的问题，有 24 条评论。
    - **摘要**：用户将 CodeWhale 与同被 DeepSeek 官方收录的另一项目对比，指出其输入缓存命中率（不到 95%）远低于竞品，导致 token 消耗和响应速度不理想。
    - https://github.com/Hmbown/CodeWhale/issues/1177

2.  **[#1120] 缓存命中方面似乎还是有些问题**
    - **重要性：高**。与 #1177 一脉相承，有 21 条评论，是目前性能优化的核心。
    - **摘要**：用户报告对相同项目进行修改时，缓存未命中的问题在 v0.8.17 版本后似乎仍未完全解决，社区正在积极排查原因。
    - https://github.com/Hmbown/CodeWhale/issues/1120

3.  **[#743] token消耗增大了很多**
    - **重要性：高**。13 条评论，是另一个与性能和成本紧密相关的问题。
    - **摘要**：用户报告半天消耗 4 亿 token，请求过于密集，怀疑每次交互的对话信息未得到有效优化。
    - https://github.com/Hmbown/CodeWhale/issues/743

4.  **[#3192] 将项目列入 agentclientprotocol/registry**
    - **重要性：高**。这是打通与 Zed 编辑器集成的关键一步。
    - **摘要**：用户请求将 CodeWhale 加入 Agent Client Protocol (ACP) 的注册列表中，以便 Zed 等编辑器能够更简单地安装和使用该工具，代表了社区对 IDE 集成的迫切需求。
    - https://github.com/Hmbown/CodeWhale/issues/3192

5.  **[#3275] CodeWhale 过度参与修改，自问自答并偏离用户意图**
    - **重要性：高**。12 条评论，反映了 Agent 行为可预测性的关键问题。
    - **摘要**：用户反馈 CodeWhale 在未获得确认的情况下，会自我驱动地执行超出用户需求的修改，导致工作范围过度延伸，这是一个严重的可靠性回归问题。
    - https://github.com/Hmbown/CodeWhale/issues/3275

6.  **[#2870] EPIC: 分阶段命令边界重构**
    - **重要性：高**。9 条评论，这是一个关于代码架构重构的跟踪议题。
    - **摘要**：该 EPIC 跟踪了为解决 #2791 问题而进行的一系列命令边界重构工作。这显示了项目在代码可维护性和功能组织上的深度投入。
    - https://github.com/Hmbown/CodeWhale/issues/2870

7.  **[#3495] 采用 Moraine 作为 CodeWhale 的内存后端**
    - **重要性：中**。引入外部记忆机制，是增强 Agent 长期记忆能力的关键。
    - **摘要**：计划将 Moraine 项目作为外部记忆后端，以无损耗方式注入 CodeWhale 的会话，并通过 MCP 工具提供可搜索的记忆召回能力。
    - https://github.com/Hmbown/CodeWhale/issues/3495

8.  **[#3541] 功能请求：基于 Rust 的原生运行时/桌面客户端**
    - **重要性：中**。社区对新平台的探索。
    - **摘要**：用户提出虽然 npm 包分发便捷，但 Node.js 运行时带来了冷启动延迟和高内存占用等问题。建议开发一个基于 Rust 的原生客户端，以追求更低的延迟和更好的非编码场景体验。
    - https://github.com/Hmbown/CodeWhale/issues/3541

9.  **[#3717] [Windows] DSML 相关内容输出导致任务中断**
    - **重要性：中**。一个平台相关的严重 Bug，影响了 Windows 用户的正常工作流。
    - **摘要**：在 Windows 版本的 v0.8.65 上，只要输出中包含 DSML 相关内容，任务就一定会中断。
    - https://github.com/Hmbown/CodeWhale/issues/3717

10. **[#1641] Agent 模式：工具调用失败时增加回退策略**
    - **重要性：中**。提升 Agent 鲁棒性的长期需求。
    - **摘要**：当 Agent 依赖的外部服务（如 Web 搜索、API 调用）失败时，会陷入无休止的重试循环。用户建议增加优雅降级或替代方案的策略。
    - https://github.com/Hmbown/CodeWhale/issues/1641

## 重要 PR 进展

1.  **[#3707] docs: 添加 v0.8.66 版本发布记录**
    - **重要性：极高**。这是 v0.8.66 版本正式发布的关键准备文件，总结了该版本的 token 缓存评分、Stream-json 指标、提示词精简和 ACP 注册状态。
    - https://github.com/Hmbown/CodeWhale/pull/3707

2.  **[#3715] feat(config): 为 GUI 集成暴露热重载 API**
    - **重要性：高**。连接 TUI 与未来图形界面的桥梁，通过 HTTP 端点实现配置的实时重载，无需重启引擎。
    - https://github.com/Hmbown/CodeWhale/pull/3715

3.  **[#3716] fix(tui): 在任务中显示 hunt 结果**
    - **重要性：中**。增强 TUI 的任务管理功能，之前的`/task list`无法显示猎杀（hunt）任务的判决结果，现在可以正确展示了。
    - https://github.com/Hmbown/CodeWhale/pull/3716

4.  **[#3710] Feat/plugin p3 mcp**
    - **重要性：高**。插件系统（MCP 部分）的建设，是实现工具生态扩展的核心步骤。
    - https://github.com/Hmbown/CodeWhale/pull/3710

5.  **[#3708] feat(plugins): 添加清单解析、发现和注册**
    - **重要性：高**。新增了核心插件系统的基础设施，包括从`plugin.toml`解析清单、启用/禁用/列表注册，以及用户目录发现。
    - https://github.com/Hmbown/CodeWhale/pull/3708

6.  **[#3696] feat(prompts): 允许从配置文件目录覆盖基础提示词**
    - **重要性：高**。响应社区需求，允许用户为非软件工程用例（如写作、文档审查）替换系统提示词，极大地扩展了工具的适用场景。
    - https://github.com/Hmbown/CodeWhale/pull/3696

7.  **[#3702] feat(acp): 将 session/prompt 增量流式传输为 session/update 块**
    - **重要性：高**。改进了 ACP（Agent Client Protocol）适配器，将原本的缓冲式输出改为流式输出，对于 Zed 等编辑器来说，可以实现增量渲染，体验更流畅。
    - https://github.com/Hmbown/CodeWhale/pull/3702

8.  **[#3705] fix(engine): 在重复搜索错误后建议直接使用 URL**
    - **重要性：中**。智能的错误处理回退策略，当搜索工具反复失败时，会引导模型直接使用 URL 访问，提升任务完成率。
    - https://github.com/Hmbown/CodeWhale/pull/3705

9.  **[#3703] fix(engine): 在重复工具错误后提示模型回退**
    - **重要性：中**。增强 Agent 的自我修复能力，在工具调用连续失败时给出运行时提示，建议切换工具或缩小请求范围。
    - https://github.com/Hmbown/CodeWhale/pull/3703

10. **[#3607] chore: 重新激活过期问题清理**
    - **重要性：中**。项目维护工作，通过创建正确的 GitHub 标签和策略，自动化清理久未更新的 issue，有助于维护社区健康度。
    - https://github.com/Hmbown/CodeWhale/pull/3607

## 功能需求趋势

- **集成能力 (Integration):** 社区最强烈的呼声是 **与编辑器和IDE的深度集成**，特别是与 Zed 的集成（通过 ACP 协议）。
- **性能与成本 (Performance & Cost):** **输入缓存优化** 和 **降低 Token 消耗** 是当前最核心的技术痛点，开发者对成本极其敏感。
- **插件与扩展性 (Plugin & Extensibility):** 正在构建的 **插件系统（MCP， Plugin Manifest）** 和 **内存后端** 显示了向平台化发展的趋势。
- **人性化与鲁棒性 (UX & Robustness):** Agent 行为的 **可预测性** 和 **容错性** 受到高度关注，包括减少不必要的自主行为、增加工具调用失败后的回退策略。
- **本地化与多语言 (Localization):** 社区积极贡献多语言文档（如韩语），显示了全球化社区的发展需求。

## 开发者关注点

- **痛点：**
    - **缓存命中率低** 和 **Token 消耗巨大** 是开发者最直接的体验差和成本负担。
    - **Agent 过度自主**的行为导致工作范围失控，开发者希望工具“听话”，能等待确认后再执行。
    - Windows 平台上的 **DSML 相关 Bug** 直接阻塞了部分用户的工作流。
    - 除了中文和日文外，**其他语言（如韩文、西班牙文）文档的缺失** 限制了工具的全球普及。
- **高频需求：**
    - **代码编辑器集成**，尤其是 Zed 和 VSCode，以提高使用便利性。
    - **上下文压缩与 Token 经济**，希望工具能优化对话结构，减少不必要的 token 重复传输。
    - **外部记忆与持久化**，希望 Agent 能记住长期上下文，减少重复工作。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*