# AI 开源趋势日报 2026-06-23

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-23 10:50 UTC

---

# AI 开源趋势日报 (2026-06-23)

## 1. 今日速览

今日 GitHub AI 开源社区呈现三大热点：**AI 视频生成/编辑工具**爆发式增长，`OpenMontage` 与 `palmier-pro` 单日分别获得近 3000 和 2500 星；**AI Agent 技能包与 MCP 生态**持续发酵，`mattpocock/skills`（+2051）和 `codebase-memory-mcp`（+1185）深受开发者青睐；**大模型推理轻量化**方向继续走热，`airllm` 以单 4GB GPU 运行 70B 模型引发关注。主题搜索中，`ollama`、`AutoGPT`、`langchain` 等经典项目依然保持极高活跃度，而 `browser-use`、`TradingAgents` 等新型 Agent 框架也快速积累星数。

## 2. 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）

- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐174,774  
  本地运行大模型的极简工具，现已支持 Kimi、GLM、MiniMax 等最新国产模型，是个人开发者上手 LLM 的首选。

- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐83,619  
  高性能 LLM 推理与服务引擎，PagedAttention 技术极大提升吞吐量，成为生产部署的主流选择。

- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** ⭐137,846 (+615 today)  
  面向 AI Agent 的网页搜索与抓取 API，支持大规模结构化数据获取，今日热榜持续上升。

- **[lyogavin/airllm](https://github.com/lyogavin/airllm)** ⭐0 (+193 today)  
  在单张 4GB GPU 上运行 70B 大模型，通过内存优化和量化实现极低资源推理，重新定义边缘 AI 边界。

- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐47,702  
  集成智能聊天、自治 Agent 和 300+ 助手的跨模型 AI 生产力套件，支持统一接入前沿 LLM。

- **[langchain-ai/langchain](https://github.com/langchain-ai/langchain)** ⭐139,963  
  最流行的 LLM 应用开发框架，提供链式编排、Agent、工具调用等抽象，生态极为丰富。

### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）

- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐185,104  
  自主 AI Agent 先驱，支持任务规划、互联网访问、代码执行，持续迭代至全功能 Agent 平台。

- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** ⭐73,649 (+738 today)  
  字节跳动开源的长期任务 SuperAgent，内置沙箱、记忆、工具和子 Agent 系统，可处理数分钟至数小时的复杂任务。

- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** ⭐78,085  
  AI 驱动的软件开发 Agent，能自主阅读代码、编写测试、提交 PR，大幅提升开发效率。

- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** ⭐100,229  
  让 AI Agent 像人一样操作浏览器，自动化在线任务，是 Web 自动化 Agent 领域的标杆项目。

- **[TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents)** ⭐88,108  
  基于多 Agent 的 LLM 金融交易框架，将大模型推理能力应用于量化策略与风险控制。

- **[mattpocock/skills](https://github.com/mattpocock/skills)** ⭐0 (+2051 today)  
  来自资深工程师的 Claude Code 技能包，涵盖 CEO、设计师、工程经理等 23 种角色工具，一键复用高效开发流程。

### 📦 AI 应用（具体应用产品、垂直场景解决方案）

- **[calesthio/OpenMontage](https://github.com/calesthio/OpenMontage)** ⭐0 (+2938 today)  
  全球首个开源 Agentic 视频制作系统，集成 12 条管线、52 种工具、500+ 技能，可将 AI 编程助手转变为完整视频工作室。

- **[palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro)** ⭐0 (+2463 today)  
  专为 AI 打造的 macOS 视频编辑器，内置智能剪辑、特效生成等 AI 能力，定位专业创作者。

- **[jamiepine/voicebox](https://github.com/jamiepine/voicebox)** ⭐0 (+529 today)  
  开源 AI 语音工作室，支持声音克隆、听写、创造，为内容创作者提供全栈语音工具。

- **[heygen-com/hyperframes](https://github.com/heygen-com/hyperframes)** ⭐0 (+395 today)  
  写 HTML 即可渲染视频，专为 Agent 设计的视频生成工具，降低视频制作门槛。

- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** ⭐46,447 (+1557 today)  
  LLM 驱动的多市场股票分析系统，整合行情、新闻、决策看板与自动推送，支持零成本定时运行。

- **[garrytan/gstack](https://github.com/garrytan/gstack)** ⭐0 (+573 today)  
  直接使用著名投资人 Garry Tan 的 Claude Code 配置，包含 23 个经过实战检验的角色工具。

- **[mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills)** ⭐0 (+956 today)  
  817 个结构化网络安全技能，映射至 6 个主流框架，可被 Claude Code、Copilot、Cursor 等 20+ 平台调用，是 Agent 安全技能的标准化资源。

### 🧠 大模型/训练（模型权重、训练框架、微调工具）

- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐161,832  
  🤗 深度学习模型标准库，支持文本、视觉、音频等模态的推理与训练，生态核心。

- **[tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)** ⭐195,844  
  全球最广泛使用的机器学习框架，持续迭代，支持端侧与云端大规模训练。

- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** ⭐100,973  
  动态神经网络的标杆框架，AI 研究的首选平台，与 HuggingFace 深度集成。

- **[ultralytics/ultralytics](https://github.com/ultralytics/ultralytics)** ⭐58,714  
  YOLOv8 及其后续版本，提供训练、推理、部署一站式解决方案，计算机视觉领域的事实标准。

- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7,114  
  大模型评测平台，支持 100+ 数据集和主流模型，为模型选型提供客观基准。

- **[galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining)** ⭐266  
  可靠、最小化的基座模型预训练库，专注于训练稳定性和可扩展性。

- **[zjunlp/LightThinker](https://github.com/zjunlp/LightThinker)** ⭐164  
  [EMNLP 2025] 轻量级推理压缩方法，逐步压缩思维链，降低推理成本，是高效推理的前沿探索。

### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）

- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐83,426  
  领先的 RAG 引擎，融合 Agent 能力与上下文层，支持多种数据源和深度检索。

- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐44,910  
  高性能云原生向量数据库，专为规模化的向量 ANN 搜索设计，是 AI 知识库基石。

- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** ⭐32,573  
  高扩展性向量数据库，支持百万级向量近实时搜索，提供云端和自托管部署。

- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐59,203  
  通用 AI 记忆层，为 Agent 提供跨会话持久化记忆，支持语义搜索和自动压缩。

- **[DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp)** ⭐0 (+1185 today)  
  高性能代码智能 MCP 服务器，将代码库索引为持久知识图谱，毫秒级查询，支持 158 种语言，零依赖静态二进制。

- **[safishamsi/graphify](https://github.com/safishamsi/graphify)** ⭐70,937  
  将代码、SQL 模式、文档等任意文件夹转化为可查询的知识图谱，兼容多种 AI 编码助手，是代码理解的强力工具。

- **[siyuan-note/siyuan](https://github.com/siyuan-note/siyuan)** ⭐44,569  
  隐私优先、自托管的开源知识管理软件，融合 AI 搜索与笔记，适合个人和团队构建知识库。

## 3. 趋势信号分析

从今日热榜和主题搜索中可以提炼出以下关键趋势：

1. **AI 视频与多模态创作工具井喷**：`OpenMontage` 和 `palmier-pro` 单日分别获近 3000 和 2500 星，加上 `hyperframes` 和 `voicebox`，表明社区对“用 AI 生成/编辑视频”的需求已从概念验证转向实用工具。这类工具往往与 Agent 结合（如 OpenMontage 定位为 Agent 视频工厂），暗示“AI Agent + 视频”将成为下一波内容生产范式。

2. **Agent 技能包与 MCP 协议生态成熟**：`mattpocock/skills` (+2051)、`Anthropic-Cybersecurity-Skills` (+956)、`codebase-memory-mcp` (+1185) 等项目的爆红，说明开发者不再满足于单个 Agent 框架，而是追求**可复用的技能模块**和**标准化的工具集成协议**。MCP（Model Context Protocol）正在成为连接 Agent 与现实工具的“USB 接口”。

3. **大模型推理持续向轻量化和边缘端倾斜**：`airllm` 以单 4GB GPU 运行 70B 模型，同时 `LightThinker` 提出压缩思维链技术，反映出在模型参数不断膨胀的背景下，社区对**低成本推理**的刚需依然强劲。这一方向可能催生更多量化、蒸馏、投机解码等创新工具。

4. **金融 AI 应用热度回升**：`TradingAgents` 星数突破 88k，`daily_stock_analysis` 单日新增 1557 星，说明 LLM 在金融领域（量化交易、市场分析）的应用想象空间正在被更多开发者 validate。

## 4. 社区关注热点

- **OpenMontage**：首个开源 Agentic 视频制作系统，今天暴涨 2938 星。如果你关注 AI 生成视频的未来，这是必研究的项目。它展示了如何将多个 Agent 技能（脚本、分镜、渲染）编排成完整管线。

- **mattpocock/skills**：来自知名 TypeScript 教育家，提供了 23 个可直接放入 `.claude` 目录的工具角色。对希望快速提升 Claude Code 工作效率的开发者而言，这是现成的“最佳实践”。

- **codebase-memory-mcp**：C 语言编写的高性能 MCP 服务器，将代码库索引为知识图谱，毫秒级查询，且零依赖。它代表 MCP 工具链向“极致性能”演进的方向，适合大型代码库的 Agent 辅助开发。

- **daily_stock_analysis**：LLM 驱动的多市场股票分析系统，单日 +1557 星。其“零成本定时运行”的设计理念降低了小型散户使用 AI 分析的门槛，值得关注其背后与 LLM 交互的 Prompt 工程技巧。

- **airllm**：单 4GB GPU 运行 70B 模型，这对于资源受限的开发者和边缘设备意义重大。建议研究其使用的内存卸载和量化策略，可能启发更多类似优化方案。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*