# AI 开源趋势日报 2026-07-05

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-07-05 09:32 UTC

---

好的，这是为您准备的《AI 开源趋势日报》（2026-07-05）。

---

## **AI 开源趋势日报 | 2026-07-05**

### **第一步：AI 相关性过滤**
已从完整数据集中筛除与 AI/ML 无关的项目，如通用工具（`immich-app/immich`）、音乐播放器（`chthollyphile/folia-major`）、游戏管理器（`rommapp/romm`）等。

### **第二步：项目分类**
已按照五大核心维度对过滤后的项目进行归类。

### **第三步：完整报告**

---

#### **1. 今日速览**

今日 AI 开源生态的核心关键词是 **“Agent 技能工程 (Agent Skills)”** 与 **“极致 Token 优化”**。以 `caveman` 项目为代表的“石器时代”语言风格提示词工程爆火，单日增长 1089 星，成为当日现象级项目。同时，围绕 Claude Code、Codex 等编码 Agent 的“技能 (Skill)”与“系统提示词”生态迎来爆发式增长，`mattpocock/skills`、`agentskills/agentskills` 等项目定义了全新的开发范式。此外，AI 安全测试工具 `strix` 与跨平台 AI 会议助手 `meetily` 也凭借其专业性和本地化特性，获得了社区的强烈关注。

#### **2. 各维度热门项目**

**🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）**

- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐175,492
  - 本地大模型运行的事实标准。今日更新后支持 Kimi-K2.6、GLM-5.1 等最新模型，持续巩固其作为本地模型代理的基石地位。
- **[JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman)** ⭐0 (+1089 today)
  - 一个 Claude Code 的“技能包”，通过让 AI 使用原始、简短的“穴居人”风格输出，可削减 65% 的 Token 消耗。火爆全网，代表了提示词工程向极端效率优化的新方向。
- **[ChromeDevTools/chrome-devtools-mcp](https://github.com/ChromeDevTools/chrome-devtools-mcp)** ⭐0 (+304 today)
  - 官方出品的 Chrome DevTools MCP 服务，使任何支持 MCP 协议的 AI 编码代理（如 Claude Code）能够直接调试浏览器。标志着 AI 开发工具链与浏览器核心工具的深度整合。
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐85,399
  - 高性能 LLM 推理与服务引擎。作为生产级部署的标配，其持续的更新支撑着下游无数 AI 应用的运行。
- **[0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig)** ⭐7,829
  - 使用 Rust 构建模块化、可扩展的 LLM 应用框架。代表 Rust 语言在 AI 基础设施领域，尤其是追求性能和安全性的 Agent 开发中的崛起。
- **[crynta/terax-ai](https://github.com/crynta/terax-ai)** ⭐0 (+62 today)
  - 仅 7MB 的轻量级终端优先 AI 原生开发工作空间。为追求极致性能与极简环境的开发者提供了新的选择。

**🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）**

- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐209,385
  - 一个与开发者共同成长的 Agent 框架。其高星数证明了社区对长期、可演化 Agent 的巨大需求。
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐185,366
  - 自主 AI Agent 的先驱。持续迭代，是探索通用任务自动化能力的重要参考。
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** ⭐79,463
  - AI 驱动的软件开发代理。致力于让 AI 参与完整的开发流程，代表了 AI 辅助编程从“补全”到“全栈 Agent”的演进。
- **[alibaba/page-agent](https://github.com/alibaba/page-agent)** ⭐0 (+742 today)
  - 阿里巴巴开源的 JavaScript 页面内 GUI 代理，允许用自然语言控制网页界面。在浏览器自动化和 RPA 领域开辟了全新的轻量级路径。
- **[usestrix/strix](https://github.com/usestrix/strix)** ⭐0 (+1904 today)
  - 开源的 AI 渗透测试工具，用于自动发现和修复应用漏洞。将 AI Agent 应用于网络安全，是 AI 安全领域的新兴力量。
- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** ⭐76,128
  - 字节跳动开源的长时域超级 Agent 框架，能够处理需要数小时完成的复杂任务。展示了顶尖大厂在高级 Agent 研究上的成果。
- **[ogulcancelik/herdr](https://github.com/ogulcancelik/herdr)** ⭐0 (+707 today)
  - 运行在终端里的 Agent 复用器。旨在管理、切换和编排多个不同 Agent，解决“Agent 孤岛”问题，是一个新兴的“Agent 中台”概念应用。

**📦 AI 应用（具体应用产品、垂直场景解决方案）**

- **[open-webui/open-webui](https://github.com/open-webui/open-webui)** ⭐144,246
  - 用户友好的 AI 交互界面，支持多种后端。是搭建私人 AI 助理/知识库必备的前端项目。
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐48,173
  - AI 生产力工作室，集成智能对话、自主 Agent、多模型支持。正成为个人工作流程中心的全能型选手。
- **[Zackriya-Solutions/meetily](https://github.com/Zackriya-Solutions/meetily)** ⭐0 (+718 today)
  - 本地优先、隐私至上的 AI 会议助手。支持实时转录、说话人分离和本地大模型摘要，满足了用户对数据安全与功能并重的强需求。
- **[santifer/career-ops](https://github.com/santifer/career-ops)** ⭐58,600
  - AI 驱动的求职系统，利用 Claude Code 实现简历优化、职位搜索等。代表了 AI Agent 在垂直职业场景中的深度应用。
- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** ⭐144,683
  - 为 AI Agent 设计的网页抓取和搜索 API。随着 AI 需要从互联网获取实时信息，此类工具的重要性与日俱增。
- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** ⭐36,737
  - 从文档自动生成可编辑的 PPT，且保留原生图表和动画。在办公自动化领域展示了 AI 的强大能力。
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** ⭐35,765
  - 用于构建 Agent 和生成式 UI 的前端框架。其提出 AG-UI 协议，推动 AI 与前端交互的标准化。

**🧠 大模型/训练（模型权重、训练框架、微调工具）**

- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐162,254
  - ML 模型定义与训练的基石框架。持续更新以支持最新的模型架构，是整个生态的底盘。
- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** ⭐101,500
  - 动态神经网络计算框架。GPU 加速和灵活的设计使其成为 AI 研究首选。
- **[tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)** ⭐196,044
  - 老牌的机器学习框架，在工业界和移动端部署仍有广泛应用。
- **[ultralytics/ultralytics](https://github.com/ultralytics/ultralytics)** ⭐59,123
  - 计算机视觉领域的事实标准，提供从 YOLOv5 到 YOLO26 的全系列目标检测模型及训练工具。
- **[microsoft/ML-For-Beginners](https://github.com/microsoft/ML-For-Beginners)** ⭐87,707
  - 微软出品的机器学习入门教程。是社区新人学习经典 ML 知识的最佳路径之一。
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7,155
  - 大模型评测平台。随着模型数量激增，公平、全面的评测平台成为行业刚需。

**🔍 RAG/知识库（向量数据库、检索增强、知识管理）**

- **[langgenius/dify](https://github.com/langgenius/dify)** ⭐147,721
  - 生产级的 Agent 工作流开发平台。其强大的 RAG 和 Agent 编排能力，使其成为构建企业级 AI 应用的首选之一。
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐84,311
  - 领先的开源 RAG 引擎。将高级 RAG 技术与 Agent 能力结合，为 LLM 创建卓越的上下文层。
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐45,074
  - 高性能云原生向量数据库。是大规模语义搜索和 RAG 系统的核心基础设施。
- **[Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify)** ⭐77,819
  - 将代码、文档等转化为可查询知识图谱的 AI 技能。代表 RAG 技术向更结构化的知识表示演进。
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐60,114
  - AI Agent 的通用记忆层。解决了 Agent 长期记忆和上下文持久化这一核心痛点。
- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** ⭐12,635
  - 实现 97% 存储节省的 RAG 系统。其高性能和隐私性，为在个人设备上运行大规模 RAG 应用提供了可能。
- **[alibaba/zvec](https://github.com/alibaba/zvec)** ⭐12,787
  - 阿里开源的轻量级进程内向量数据库。为追求极致低延迟的应用场景提供了新方案。

#### **3. 趋势信号分析**

今日榜单最核心的信号是 **“Agent 技能”** 生态的爆发。这不是简单的提示词分享，而是一种新的“代码 + 提示词”混合编程模式。`caveman` 的走红，表面上是娱乐化，实质上是社区对 **Token 成本极致优化** 的焦虑与探索。同时，`mattpocock/skills`、`agentskills/agentskills` 和 `alirezarezvani/claude-skills` 等项目，正在试图为这些“技能”建立标准和分发平台，这将深刻改变 AI 编码 Agent 的使用方式，使其从一个通用工具转变为可定制、可扩展的开发环境。

其次，**AI 安全与测试** 成为新热点。`strix` 的单日近 2000 星，反映出随着 AI Agent 被赋予越来越多的代码执行和系统操作权限，社区对安全性的关注度已从理论探讨转向实际工具需求。这预示着“红蓝对抗”的想法正在被引入 AI 领域。

最后，**MCP (Model Context Protocol)** 协议正快速渗透。从 Chrome DevTools 到 Unity 编辑器，各大平台纷纷推出 MCP 支持，这标志着 AI Agent 与现有软件生态的“互操作层”正在成为基础设施，其重要性不亚于 RESTful API。

#### **4. 社区关注热点**

- **Agent 技能（Skills）与系统提示词泄露**： `caveman` 和 `system_prompts_leaks` 项目“一正一邪”，一个教如何优化 Agent，一个“泄露”主流模型的系统提示词。开发者应密切关注此领域，这直接关系到 Agent 的效能与安全性。
- **AI 驱动的安全测试工具 `strix`**：单日 1904 星。它代表了 AI 从“攻击目标”转变为“攻击工具”的趋势。这是一个全新且高潜力的赛道，对 DevOps 和安全工程影响深远。
- **本地化、隐私优先的 AI 应用 `meetily`**：在云服务盛行的今天，`meetily` 通过 100% 本地处理获得大量星数，显示了用户对数据隐私的高度重视。这是 AI 应用设计的重要参考。
- **“Agent 复用器”（Agent Multiplexer）概念**：`herdr` 项目尝试解决多 Agent 管理协调问题。随着 Agent 数量增多，这类“元工具”的需求会非常强烈，是值得关注的前沿方向。
- **RAG 技术的“轻量化”与“结构化”**：`LEANN`（存储节省 97%）和 `Graphify`（知识图谱化）代表了 RAG 的两个细分方向：一是让 RAG 更“便宜”以便于在个人设备运行，二是让 RAG 更“聪明”以提供更精确的推理支持。两者都指向了更实用的 RAG 落地路径。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*