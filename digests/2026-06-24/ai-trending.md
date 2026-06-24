# AI 开源趋势日报 2026-06-24

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-24 10:35 UTC

---

## AI 开源趋势日报 | 2026-06-24

---

### 1. 今日速览

- **智能体（Agent）赛道持续井喷**：字节跳动开源 **deer-flow**（长时SuperAgent）获得 739 今日新增，Nous Research 的 **hermes-agent** 总量突破 20 万星，成为社区焦点。
- **视频/语音 AI 应用崛起**：**OpenMontage**（开源视频生产系统）单日暴增 3592 星，**voicebox**（AI 语音工作室）新增 1045 星，视频与语音生成正成为开源新热点。
- **Claude 生态工具链加速成熟**：官方插件目录 **claude-plugins-official**、最佳实践 **claude-code-best-practice** 以及 **gstack**（Garry Tan 的 Claude Code 配置）等大量涌现，Agent 工作流标准化趋势明显。
- **代码智能 MCP 服务器爆发**：**codebase-memory-mcp**（高性能代码知识图谱）单日新增 1300 星，展示了以知识图谱压缩上下文、降低 token 消耗的新范式。

---

### 2. 各维度热门项目

#### 🔧 AI 基础工具（框架、SDK、推理引擎、CLI）

| 项目 | Stars | 一句话说明 |
|------|-------|-----------|
| [ollama](https://github.com/ollama/ollama) | 174,828 | 本地运行大模型的一站式工具，支持 Kimi、GLM、DeepSeek 等多模型。 |
| [vllm](https://github.com/vllm-project/vllm) | 83,704 | 高吞吐、低内存的 LLM 推理引擎，生产环境标配。 |
| [firecrawl](https://github.com/firecrawl/firecrawl) | 138,403 | 为 AI Agent 提供网页搜索/爬取/数据交互的 API，🔥 级流行。 |
| [langchain](https://github.com/langchain-ai/langchain) | 140,076 | 最流行的 Agent 工程框架，集成了工具调用、RAG、MCP 支持。 |
| [transformers](https://github.com/huggingface/transformers) | 161,861 | HuggingFace 核心库，支持文本/视觉/语音/多模态模型推理与训练。 |
| [rig](https://github.com/0xPlaygrounds/rig) | 7,732 | Rust 生态的模块化 LLM 应用框架，面向高性能场景。 |

#### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）

| 项目 | Stars | 今日新增（Trending） | 一句话说明 |
|------|-------|----------------------|-----------|
| [hermes-agent](https://github.com/NousResearch/hermes-agent) | 201,493 | +936 | 与你一起成长的通用 Agent，支持技能、记忆、工具。 |
| [AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | 185,141 | — | 最具影响力的自主 Agent 项目，持续推动 Agent 民主化。 |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | 74,267 | +739 | 字节开源的长时 SuperAgent 框架，集成沙箱、记忆、子代理等。 |
| [OpenHands](https://github.com/OpenHands/OpenHands) | 78,207 | — | AI 驱动的软件开发助手，自动完成编码任务。 |
| [CopilotKit](https://github.com/CopilotKit/CopilotKit) | 35,453 | — | 前端 Agent UI 栈，支持 React/Angular/Slack 等多平台。 |
| [codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp) | 0 (新) | +1,300 | 高性能 MCP 服务，将代码库索引为知识图谱，毫秒级查询，减少 99% token 消耗。 |
| [garrytan/gstack](https://github.com/garrytan/gstack) | 0 (新) | +1,011 | Garry Tan 的 Claude Code 工作流配置，含 23 个角色工具（CEO/设计师/QA 等）。 |
| [affaan-m/ECC](https://github.com/affaan-m/ECC) | 220,870 | +593 | Agent 性能优化系统，支持技能、直觉、记忆、安全，可接入 Claude Code/Codex 等。 |

#### 📦 AI 应用（具体产品、垂直场景）

| 项目 | Stars | 今日新增 | 一句话说明 |
|------|-------|----------|-----------|
| [calesthio/OpenMontage](https://github.com/calesthio/OpenMontage) | 0 (新) | +3,592 | 世界首个开源 Agent 视频生产系统，12 条管线、52 个工具、500+ 技能。 |
| [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) | 47,881 | +1,119 | LLM 驱动的多市场股票智能分析，支持行情、新闻、决策仪表盘。 |
| [jamiepine/voicebox](https://github.com/jamiepine/voicebox) | 0 (新) | +1,045 | 开源 AI 语音工作室，支持克隆、听写、创作。 |
| [palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro) | 0 (新) | +1,630 | macOS 原生 AI 视频编辑器。 |
| [koala73/worldmonitor](https://github.com/koala73/worldmonitor) | 0 (新) | +294 | AI 驱动的全球实时情报仪表盘，聚合新闻、地缘政治追踪。 |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | 47,743 | — | AI 生产力工作室，支持智能对话、自主 Agent、300+ 助手。 |
| [OpenBB-finance/OpenBB](https://github.com/OpenBB-finance/OpenBB) | 69,609 | — | 面向分析师和 AI Agent 的金融数据平台。 |

#### 🧠 大模型/训练（模型权重、训练框架、微调工具）

| 项目 | Stars | 一句话说明 |
|------|-------|-----------|
| [pytorch/pytorch](https://github.com/pytorch/pytorch) | 101,029 | 深度学习框架，动态图与 GPU 加速的行业标准。 |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow) | 195,913 | 谷歌开源的 ML 框架，覆盖从研究到生产。 |
| [ultralytics/ultralytics](https://github.com/ultralytics/ultralytics) | 58,755 | YOLO 系列目标检测最新实现，支持训练/推理。 |
| [open-compass/opencompass](https://github.com/open-compass/opencompass) | 7,117 | 开源 LLM 评测平台，支持 100+ 数据集、多模型对比。 |
| [galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining) | 267 | 可靠、可扩展的预训练库，支持 Foundation 模型与 World 模型。 |
| [zjunlp/LightThinker](https://github.com/zjunlp/LightThinker) | 164 | [EMNLP2025] 轻量思考：逐步压缩推理步骤，减少 token 消耗。 |

#### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）

| 项目 | Stars | 一句话说明 |
|------|-------|-----------|
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | 44,928 | 高性能云原生向量数据库，支持大规模近似搜索。 |
| [qdrant/qdrant](https://github.com/qdrant/qdrant) | 32,602 | 高可扩展的向量数据库与搜索引擎，适用于下一代 AI。 |
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | 83,517 | 领先的开源 RAG 引擎，融合 Agent 能力构建上下文层。 |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | 59,321 | 通用 AI Agent 记忆层，跨会话持久化知识。 |
| [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) | 62,012 | 本地优先的 Agent 体验，支持 RAG、多模型、文档管理。 |
| [DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp) | 0 (新) | 代码知识图谱 MCP，158 种语言支持，毫秒级查询，99% 更少 token（同时归入此类别）。 |

---

### 3. 趋势信号分析

- **智能体（Agent）进入“工作流标准化”阶段**：今日热榜中，**gstack**、**ECC**、**claude-plugins-official** 等项目并非 Agent 框架本身，而是针对特定 Agent（如 Claude Code）的配置、插件、性能优化系统。这表明社区不再满足于“能跑”，而是追求**高效率、可复用的工作流模板**。**codebase-memory-mcp** 则将代码库转化为知识图谱，直击 Agent 上下文窗口瓶颈，是解决“长程任务”的关键基础设施。
- **视频/语音 AI 生成首次大规模登榜**：**OpenMontage**（+3592）和**palmier-pro**（+1630）是榜单前两名，均与视频生产相关。结合**voicebox**（语音克隆）的爆发，可以判断**多模态生成（特别是视频）正在从闭源走向开源**。这背后可能受到近期 Sora、可灵等模型开源影响，开发者开始构建上游工具链。
- **Agent 工具链与“AI 原生编程”深度绑定**：大量项目围绕 Claude Code、Codex、Cursor 等 AI 编码助手开发（如 ECC、gstack、DeusData 的 MCP）。AI 编程助手已成为 Agent 生态的“入口”，其插件/技能市场正在快速形成，类似早期 VS Code 扩展生态。
- **RAG 与向量数据库趋于成熟**：上榜项目多为稳定头部项目（milvus、qdrant、ragflow），但**mem0**（通用记忆层）和**anything-llm**（本地优先）仍保持高热度，说明开发者更关注“记忆持久化”与“本地部署”两大方向。

---

### 4. 社区关注热点

- 🔥 **OpenMontage（+3592）** —— 若你想搭建自己的 AI 视频生产线，这是目前最完整的开源方案，值得研究其 12 条管线架构。
- 🤖 **hermes-agent（201k stars + 今日新增 936）** —— Nous Research 出品，强调“与你一起成长”，其记忆与技能生态或成为下一代 Agent 范式。
- 🧠 **codebase-memory-mcp（+1300）** —— 代码智能的“瑞士军刀”，158 种语言支持、毫秒级查询。如果你在用 Claude Code 维护大型代码库，这个项目能显著降低上下文费用。
- 📦 **gstack（+1011）** —— 由硅谷知名投资人 Garry Tan 亲自维护的 Claude Code 配置，包含 CEO、设计师、工程经理等 23 个代理角色。对于团队引入 Agent 工作流极具参考价值。
- 🎤 **voicebox（+1045）** —— 开源 AI 语音工作室，支持克隆与创建。语音交互类应用（播客、有声书、虚拟助手）的开发者应重点关注。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*