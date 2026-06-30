# AI 开源趋势日报 2026-06-30

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-30 10:45 UTC

---

# AI 开源趋势日报（2026-06-30）

---

## 1. 今日速览

今日 GitHub AI 开源生态呈现 **多智能体爆发 + 垂直场景深耕** 两大主线。`agency-agents` 以 +1425 今日 stars 领涨，将多角色 agent 编排成“AI 代理机构”；`council-of-high-intelligence` 则另辟蹊径，让 18 位 AI 名人（亚里士多德、费曼、卡尼曼等）跨模型多轮辩论决策。金融领域出现两个高关注项目：`ai-berkshire`（+1386）将价值投资大师方法论封装为多 agent 研究框架，`Vibe-Trading`（+839）聚焦个人交易 Agent。视频编辑场景的 `browser-use/video-use`（+967）让 coding agent 直接操作视频，开辟新交互范式。基础工具方面，本地推理引擎 `ollama` 已支持 Kimi、GLM、DeepSeek 等最新模型，RAG 生态继续扩张，`RAGFlow`、`Milvus` 等保持高活跃。

---

## 2. 各维度热门项目

### 🔧 AI 基础工具（框架、推理引擎、开发 CLI）

- **ollama/ollama** ⭐175,176  
  [https://github.com/ollama/ollama](https://github.com/ollama/ollama)  
  一键运行 Kimi、GLM、DeepSeek、Qwen 等数十种 LLM 的本地推理工具，今日新增支持 K2.6、GLM-5.1 等最新模型，社区基础设施地位稳固。

- **vllm-project/vllm** ⭐84,889  
  [https://github.com/vllm-project/vllm](https://github.com/vllm-project/vllm)  
  高吞吐、低内存的 LLM 推理与服务引擎，生产级部署标配，持续优化 PagedAttention 和前缀缓存。

- **huggingface/transformers** ⭐162,049  
  [https://github.com/huggingface/transformers](https://github.com/huggingface/transformers)  
  多模态模型定义、训练与推理的统一框架，支撑社区 90% 以上的模型生态。

- **langchain-ai/langchain** ⭐140,566  
  [https://github.com/langchain-ai/langchain](https://github.com/langchain-ai/langchain)  
  Agent 工程核心平台，提供链、工具、记忆、多模型编排等标准组件，今日趋势中未见新爆发但基底热度不变。

- **cupy/cupy** ⭐0（今日+352）  
  [https://github.com/cupy/cupy](https://github.com/cupy/cupy)  
  NumPy/SciPy 的 GPU 加速替代，在数据预处理、科学计算领域为 AI 工作流提供底层性能支持。

- **firecrawl/firecrawl** ⭐141,804  
  [https://github.com/firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)  
  大规模网页搜索、抓取与交互 API，为 AI Agent 提供实时数据输入，今日依然保持在 llm 分类头部。

- **open-webui/open-webui** ⭐143,525  
  [https://github.com/open-webui/open-webui](https://github.com/open-webui/open-webui)  
  用户友好的 AI 聊天界面，支持 Ollama、OpenAI API，自部署首选，社区活跃度极高。

---

### 🤖 AI 智能体 / 工作流（Agent 框架、自动化、多智能体）

- **msitarzewski/agency-agents** ⭐0（今日+1425）  
  [https://github.com/msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents)  
  一套完整的“AI 代理机构”——从前端设计师到 Reddit 社区运营，每个 agent 具备独立人格与交付能力，今日最火爆新品。

- **0xNyk/council-of-high-intelligence** ⭐0（今日+331）  
  [https://github.com/0xNyk/council-of-high-intelligence](https://github.com/0xNyk/council-of-high-intelligence)  
  18 个 AI 身份（亚里士多德、费曼、托瓦兹等）跨多 LLM 提供商进行结构化多轮辩论，一键 `/council` 启动决策，创意与工程结合。

- **browser-use/video-use** ⭐0（今日+967）  
  [https://github.com/browser-use/video-use](https://github.com/browser-use/video-use)  
  让 coding agent 直接编辑视频，将浏览器自动化能力扩展到多媒体领域，开辟 Agent 新场景。

- **Unclecheng-li/VulnClaw** ⭐0（今日+129）  
  [https://github.com/Unclecheng-li/VulnClaw](https://github.com/Unclecheng-li/VulnClaw)  
  基于 AI Agent + MCP 工具链的全自动渗透测试系统，自然语言输入 → 信息收集 → 漏洞利用 → 报告生成全链路。

- **Significant-Gravitas/AutoGPT** ⭐185,227  
  [https://github.com/Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)  
  自主 Agent 鼻祖，持续迭代，今日仍是 llm 分类锚点项目。

- **langgenius/dify** ⭐147,097  
  [https://github.com/langgenius/dify](https://github.com/langgenius/dify)  
  生产级 Agentic Workflow 开发平台，支持可视化编排、工具调用、知识库集成，企业部署热度不减。

- **bytedance/deer-flow** ⭐75,550  
  [https://github.com/bytedance/deer-flow](https://github.com/bytedance/deer-flow)  
  字节跳动开源的长时间跨度 SuperAgent，集成沙箱、记忆、工具、子 Agent，可处理数小时级复杂任务。

- **HKUDS/nanobot** ⭐44,886  
  [https://github.com/HKUDS/nanobot](https://github.com/HKUDS/nanobot)  
  轻量级开源 AI Agent，支持工具调用、聊天与工作流，强调可扩展性与低门槛。

---

### 📦 AI 应用（具体产品、垂直场景）

- **xbtlin/ai-berkshire** ⭐0（今日+1386）  
  [https://github.com/xbtlin/ai-berkshire](https://github.com/xbtlin/ai-berkshire)  
  “AI 时代的伯克希尔”——集成巴菲特、芒格、段永平、李录四大价值投资方法论，多 Agent 对抗分析，金融 AI 又一个明星。

- **HKUDS/Vibe-Trading** ⭐0（今日+839）  
  [https://github.com/HKUDS/Vibe-Trading](https://github.com/HKUDS/Vibe-Trading)  
  个人交易 Agent，提供市场分析、策略执行与监控，面向量化爱好者。

- **altic-dev/FluidVoice** ⭐0（今日+830）  
  [https://github.com/altic-dev/FluidVoice](https://github.com/altic-dev/FluidVoice)  
  macOS 上最快的离线语音转文字应用，全部本地处理，隐私优先，实用性强。

- **commaai/openpilot** ⭐0（今日+458）  
  [https://github.com/commaai/openpilot](https://github.com/commaai/openpilot)  
  开源机器人操作系统，升级 300+ 车型的辅助驾驶系统，AI 在真实世界落地的标杆。

- **ZhuLinsen/daily_stock_analysis** ⭐52,159  
  [https://github.com/ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)  
  LLM 驱动的多市场股票智能分析系统，集成行情、新闻、决策看板、自动推送，零成本定时运行。

- **hugohe3/ppt-master** ⭐34,763  
  [https://github.com/hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)  
  AI 根据文档生成原生可编辑 PowerPoint（含动画、语音旁白），支持自定义模板，办公自动化利器。

- **CherryHQ/cherry-studio** ⭐47,991  
  [https://github.com/CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)  
  AI 生产力工作室，集成智能聊天、自主 Agent、300+ 助手插件，统一接入主流 LLM。

---

### 🧠 大模型 / 训练（训练框架、微调、评估）

- **hiyouga/LlamaFactory** ⭐72,829  
  [https://github.com/hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)  
  统一高效微调框架，支持 100+ LLM & VLM，ACL 2024 论文项目，社区微调首选。

- **jingyaogong/minimind** ⭐52,375  
  [https://github.com/jingyaogong/minimind](https://github.com/jingyaogong/minimind)  
  从零训练 64M 参数小 LLM 的完整教程，2 小时可复现，适合学习与实验，对教育社区价值巨大。

- **open-compass/opencompass** ⭐7,137  
  [https://github.com/open-compass/opencompass](https://github.com/open-compass/opencompass)  
  大模型评估平台，支持 100+ 数据集、多种主流模型，是模型选型与量化评测的基础设施。

- **pytorch/pytorch** ⭐101,115  
  [https://github.com/pytorch/pytorch](https://github.com/pytorch/pytorch)  
  动态神经网络与 GPU 加速框架，AI 训练的事实标准，持续集成新硬件与优化。

- **tensorflow/tensorflow** ⭐195,981  
  [https://github.com/tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)  
  老牌 ML 框架，在工业部署和生产环境仍占重要份额。

- **0xPlaygrounds/rig** ⭐7,788  
  [https://github.com/0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig)  
  Rust 语言构建的 LLM 应用框架，模块化、高性能，适合对效率有极致要求的场景。

---

### 🔍 RAG / 知识库（向量数据库、检索增强、知识管理）

- **infiniflow/ragflow** ⭐83,932  
  [https://github.com/infiniflow/ragflow](https://github.com/infiniflow/ragflow)  
  领先的开源 RAG 引擎，融合 Agent 能力，为 LLM 提供优质上下文层。

- **milvus-io/milvus** ⭐45,025  
  [https://github.com/milvus-io/milvus](https://github.com/milvus-io/milvus)  
  云原生高性能向量数据库，支撑大规模 ANN 搜索，生产级 RAG 基础设施。

- **qdrant/qdrant** ⭐32,816  
  [https://github.com/qdrant/qdrant](https://github.com/qdrant/qdrant)  
  高性能向量搜索引擎，Rust 编写，支持过滤、聚合，AI 应用首选之一。

- **mem0ai/mem0** ⭐59,755  
  [https://github.com/mem0ai/mem0](https://github.com/mem0ai/mem0)  
  通用 AI Agent 记忆层，跨会话持久化上下文，与 Claude Code、Codex 等深度集成。

- **Safishamsi/graphify** ⭐74,612  
  [https://github.com/safishamsi/graphify](https://github.com/safishamsi/graphify)  
  将代码、SQL、文档、图片等任意文件夹转为可查询知识图谱，支持多种 AI 助手，RAG 新范式。

- **Mintplex-Labs/anything-llm** ⭐62,334  
  [https://github.com/Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm)  
  本地优先的全功能 RAG 应用，支持多种文档、模型、记忆，一句命令即可运行。

- **weaviate/weaviate** ⭐16,472  
  [https://github.com/weaviate/weaviate](https://github.com/weaviate/weaviate)  
  云原生向量数据库，支持对象+向量混合搜索，具有原生容错与扩展能力。

---

## 3. 趋势信号分析

**1. 多智能体协作成为社区爆发点。**  
`agency-agents` 单日 +1425 stars 位列 Trending 第二，其“代理机构”概念——将多个专业化 Agent 组合成一个团队——引发广泛共鸣。`council-of-high-intelligence` 则通过让 18 个不同性格/专业的 AI 角色跨模型辩论来决策，折射出社区对 **Agent 多样性、角色化、共识机制** 的浓厚兴趣。

**2. 垂直领域 Agent 应用密集涌现。**  
金融领域尤为突出：`ai-berkshire`（价值投资多 Agent）、`Vibe-Trading`（个人交易 Agent）、`daily_stock_analysis`（LLM 驱动股票分析）同日登榜，暗示 AI 在量化投资和个人理财中的渗透加速。安全领域的 `VulnClaw`（渗透测试自动化），视频编辑领域的 `video-use` 也显示 Agent 正从通用对话向 **行业专用工具** 转化。

**3. 新兴技术栈出现：MCP 协议 + Agent 编排。**  
`VulnClaw` 明确使用 MCP（Model Context Protocol）工具链，`council-of-high-intelligence` 强调跨模型多样性，`agency-agents` 内置 Reddit、前端等技能插件。**MCP + 技能编排** 正成为 Agent 的事实标准架构，这与近期 Anthropic 推广的 MCP 协议高度吻合。

**4. 本地推理与 RAG 生态持续壮大。**  
`ollama` 已同步支持最新开源模型（Kimi-K2.6、GLM-5.1 等），`RAGFlow`、`anything-llm`、`mem0` 等 RAG 工具不断降低知识管理门槛，**“本地大模型 + 私有知识库”** 依然是开发者最稳固的基本盘。

---

## 4. 社区关注热点

- 🔥 **多 Agent 协作框架**：`agency-agents` 和 `council-of-high-intelligence` 代表了两种不同的 Agent 协作范式（专业化分工 vs 多角色辩论），值得跟踪后续迭代和生态发展。

- 💰 **金融 AI 工具化**：`ai-berkshire` 和 `Vibe-Trading` 展示了如何将复杂投资方法论封装成可复用的 Agent 框架，金融开发者可借鉴其多 Agent 对抗分析设计模式。

- 🎬 **Agent 多媒体操作**：`browser-use/video-use` 将编程 Agent 的能力扩展到视频编辑领域，未来 Agent 可能操控一切界面（Web、桌面、多媒体），这是人机交互的下一跳。

- 🛡️ **AI 安全自动化**：`VulnClaw` 结合 MCP 实现全链路渗透测试，安全从业者可关注如何用自然语言驱动安全工具链，提升效率。

- 📚 **RAG 记忆层**：`mem0` 和 `graphify` 分别从会话记忆和知识图谱两条路径解决 Agent 长期记忆问题，后者尤其适合代码库和文档知识的持续管理。

---

*报告基于 2026-06-30 GitHub Trending 及 AI 主题搜索数据生成，仅反映当日社区动态。*

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*