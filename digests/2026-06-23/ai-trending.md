# AI 开源趋势日报 2026-06-23

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-22 17:18 UTC

---

## AI 开源趋势日报（2026-06-23）

### 今日速览
- **AI 视频创作工具集中爆发**：`OpenMontage`（首个开源智能体视频制作系统）和 `palmier-pro`（AI 原生 macOS 编辑器）分别以 **+2,935** 和 **+2,462** 的今日增长领跑，标志 AI 视频生成进入开源深水区。
- **Agent 技能库和工具链成为新热点**：`Anthropic-Cybersecurity-Skills`（817 个结构化网络安全技能）、`mattpocock/skills`（工程师 Claude Code 技能集）、`garrytan/gstack`（CEO 工具集）均获大量关注，社区正为 AI Agent 注入“专业能力”。
- **MCP 生态加速扩张**：`codebase-memory-mcp`（高性能代码知识图谱 MCP 服务器）首日即获 **+1,186** star，反映开发者对模型上下文协议的热情。
- **长时任务 SuperAgent 框架受追捧**：字节开源的 `deer-flow` 再增 **+736** star，其“研究、编码、创造”一体的能力引发讨论。
- **金融 AI 应用持续升温**：`daily_stock_analysis`（LLM 驱动股票分析）单日新增 **+1,560** star，结合实时新闻与决策看板的模式受到散户与量化群体关注。

---

### 🔧 AI 基础工具（框架、SDK、推理引擎、CLI）

| 项目 | Stars | 说明 |
|------|-------|------|
| [ollama/ollama](https://github.com/ollama/ollama) | ⭐174,729 | 一键运行 Kimi、DeepSeek、Qwen 等主流模型，本地推理首选 |
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | ⭐83,559 | 高吞吐、低延迟的 LLM 推理与服务引擎，生产环境标配 |
| [huggingface/transformers](https://github.com/huggingface/transformers) | ⭐161,805 | 模型生态霸主，支持文本、视觉、音频等多模态 |
| [langchain-ai/langchain](https://github.com/langchain-ai/langchain) | ⭐139,886 | Agent 工程平台，本周 Agent 化趋势进一步加强 |
| [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) | ⭐136,930 (+736 today) | AI Agent 的网页数据管道，搜索/抓取/交互一体 |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | ⭐73,091 (+736 today) | 开源长时 SuperAgent 框架，集成沙箱、记忆、子 Agent |
| [lyogavin/airllm](https://github.com/lyogavin/airllm) | ⭐4.5k (+453 today) | 单张 4GB GPU 运行 70B 模型，边缘推理新突破 |
| [DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp) | ⭐1,186 today | 极速代码知识图谱 MCP 服务器，毫秒级索引 158 种语言 |

---

### 🤖 AI 智能体 / 工作流（Agent 框架、自动化、多智能体）

| 项目 | Stars | 说明 |
|------|-------|------|
| [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | ⭐185,077 | 自主 Agent 鼻祖，持续迭代多任务规划能力 |
| [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) | ⭐199,771 | “陪你成长的 Agent”，强调记忆与持续进化 |
| [mattpocock/skills](https://github.com/mattpocock/skills) | ⭐2,051 today | 知名 TS 工程师出品的 Claude Code 技能集，一键导入 |
| [garrytan/gstack](https://github.com/garrytan/gstack) | ⭐649 today | Garry Tan 的 23 个 Claude Code 工具，模拟 CEO/Designer/QA 角色 |
| [mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills) | ⭐957 today | 817 个标准化网络安全技能，覆盖 MITRE ATT&CK、NIST 等六大框架 |
| [CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit) | ⭐35,399 | React/RN/Slack 等前端 Agent UI 组件库，AG-UI 协议发起者 |
| [shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code) | ⭐67,870 | 从零构建类 Claude Code Agent，教程与轻量实现 |

---

### 📦 AI 应用（具体产品、垂直场景解决方案）

| 项目 | Stars | 说明 |
|------|-------|------|
| [calesthio/OpenMontage](https://github.com/calesthio/OpenMontage) | ⭐2,935 today | 世界首个开源智能体视频制作系统：12 流水线、52 工具、500+ 技能 |
| [palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro) | ⭐2,462 today | macOS 原生 AI 视频编辑器，释放剪辑生产力 |
| [jamiepine/voicebox](https://github.com/jamiepine/voicebox) | ⭐508 today | 开源 AI 语音工作室：克隆、听写、创作 |
| [heygen-com/hyperframes](https://github.com/heygen-com/hyperframes) | ⭐369 today | 写 HTML 渲染视频，专为 AI Agent 设计的视频生成 |
| [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) | ⭐45,656 (+1,560 today) | LLM 驱动多市场股票分析：实时行情、新闻、决策看板 |
| [TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents) | ⭐87,978 | 多智能体金融交易框架，融合 LLM 与量化策略 |
| [FlowiseAI/Flowise](https://github.com/FlowiseAI/Flowise) | ⭐53,902 | 可视化构建 AI Agent 和 RAG 应用，低代码首选 |
| [jeecgboot/JeecgBoot](https://github.com/jeecgboot/JeecgBoot) | ⭐46,835 | AI 低代码平台，一句话生成页面/流程/表单，集成知识库与 MCP |

---

### 🧠 大模型 / 训练（模型权重、训练框架、微调工具）

| 项目 | Stars | 说明 |
|------|-------|------|
| [open-compass/opencompass](https://github.com/open-compass/opencompass) | ⭐7,112 | LLM 评测平台，支持 100+ 数据集，模型能力对比标准 |
| [galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining) | ⭐266 | 可靠、可缩放的基础模型预训练库，专注稳定训练 |
| [zjunlp/LightThinker](https://github.com/zjunlp/LightThinker) | ⭐164 | EMNLP 2025 论文：逐步推理压缩，降低思维链 token 开销 |
| [0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig) | ⭐7,709 | Rust 生态 LLM 应用框架，模块化、高性能 |
| [Mirrowel/LLM-API-Key-Proxy](https://github.com/Mirrowel/LLM-API-Key-Proxy) | ⭐511 | 通用 LLM 网关：单 API 对接多个模型提供商，负载均衡 |

---

### 🔍 RAG / 知识库（向量数据库、检索增强、知识管理）

| 项目 | Stars | 说明 |
|------|-------|------|
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | ⭐83,361 | 领先开源 RAG 引擎，融合 Agent 能力构建上下文层 |
| [PaddlePaddle/PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR) | ⭐83,290 | 将 PDF/图片转为结构化数据，支持 100+ 语言，RAG 数据预处理利器 |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | ⭐59,135 | AI Agent 的通用记忆层，跨会话持久化 |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | ⭐70,619 | 将代码/文档/图像等任何文件夹转化为可查询知识图谱 |
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | ⭐44,890 | 高性能云原生向量数据库，大规模 ANN 搜索标准 |
| [qdrant/qdrant](https://github.com/qdrant/qdrant) | ⭐32,557 | Rust 写的高性能向量搜索引擎，支持过滤与云服务 |
| [siyuan-note/siyuan](https://github.com/siyuan-note/siyuan) | ⭐44,555 | 隐私优先的知识管理工具，支持 AI 辅助笔记与知识图谱 |

---

### 趋势信号分析

今天最强烈的信号是 **AI 视频创作全面开源化**。`OpenMontage` 将 500+ Agent 技能与 52 个工具编排成 12 条生产流水线，首次让 AI Coding Assistant 直接输出完整视频；`palmier-pro` 和 `hyperframes` 则从不同角度（macOS 原生应用、Agent 可调用的 HTML→视频引擎）补充了生态。这背后与近期多模态大模型的成熟（如视频理解/生成模型开源）高度相关。

第二个趋势是 **Agent 技能标准化与复用**。`Anthropic-Cybersecurity-Skills` 首次将 800 多个安全技能映射到 6 个行业框架，`mattpocock/skills` 和 `garrytan/gstack` 则展示了顶级工程师如何打包自己的“角色工具集”。这表明社区开始从“造 Agent 框架”转向“造 Agent 能力库”。

第三个值得注意的信号是 **MCP 协议生态加速**。`codebase-memory-mcp` 作为纯 C 语言实现的 MCP 服务器，首日即达千星，与 `claude-mem`、`zilliztech/claude-context` 等项目共同推动“代码智能”成为 Agent 标配特性。字节开源的 `deer-flow` 也内置了消息网关，支持 MCP 协作。

整体来看，今日热点与 **Anthropic Claude Code 生态**、**MCP 标准**、**长时 Agent** 紧密绑定。开发者的兴趣正从“让 Agent 能对话”转向“让 Agent 能交付专业成果”。

---

### 社区关注热点（值得重点跟进的 5 个方向）

- **[codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp)**：纯 C 语言、零依赖、158 语言支持、毫秒级代码知识图谱。如果你在开发 AI 代码助手或 MCP 服务，这个项目值得深入分析。
- **[mattpocock/skills](https://github.com/mattpocock/skills)**：TypeScript 权威人物公开自己的 `.claude` 配置，包含完整工程师技能集。可作为团队标准化 Agent 技能的最佳参考。
- **[OpenMontage](https://github.com/calesthio/OpenMontage)**：开源视频制作智能体，12 条流水线设计高度模块化，可用于快速搭建自动化视频内容生产管线。
- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)**：字节跳动开源的 SuperAgent 框架，内置沙箱、记忆、工具、子 Agent，适合研究/开发需要长时间持续执行的复杂任务。
- **[Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills)**：首个大规模结构化 Agent 安全技能库，跨 20+ 平台，安全工程师可直接将其注入 Claude Code、Copilot 等工具。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*