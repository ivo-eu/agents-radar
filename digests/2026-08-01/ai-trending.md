# AI 开源趋势日报 2026-08-01

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-08-01 00:12 UTC

---

# AI 开源趋势日报 · 2026-08-01

## 筛选说明

- **Trending 榜单**：12 个仓库中，与 AI/ML 明确相关保留 7 个；剔除 chatwoot、kaneo、tuicr、ESP32-Bit-Pirate、paperswithbacktest 等通用/非 AI 项目。
- **主题搜索**：79 个候选仓库中，剔除 netdata、julia、airflow、siyuan 等“可服务 AI 但不是 AI 核心”的通用基础设施/工具。
- 合并去重后，本次追踪 AI 相关项目约 **82 个**；以下按主要类别展示最值得关注的项目。

---

## 1. 今日速览

- **Agent“Skill/Harness”化成为绝对主线**：Trending 近半项目与 AI 编程客户端的可插拔技能、轻量 harness 相关，`openwork` 今日 +806。
- **教育类 AI 内容吸星力强**：`microsoft/AI-For-Beginners` 今日 +1592，入门 AI 的全球需求仍在爆发。
- **头部 Agent 项目继续领跑**：`ECC`、`Hermes-Agent`、`AutoGPT` 等保持超十万级 star，Agent 基础设施仍是最大热点。
- **RAG 与向量库加速走向“轻量、私域、低存储”**：`LEANN`、`PageIndex`、`zvec`、`lancedb` 等代表新方向。
- **AI 安全 Agent 开始独立成赛道**：`reverse-skill` 将 AI 路由能力用于渗透测试与安全研究，今日 +335。

---

## 2. 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）

- [ollama/ollama](https://github.com/ollama/ollama) — ⭐177,455  
  本地 LLM 推理事实标准，支持 Kimi、GLM-5.2、DeepSeek、gpt-oss、Qwen、Gemma 等最新开源权重；端侧部署带动整个 Agent 生态。

- [huggingface/transformers](https://github.com/huggingface/transformers) — ⭐163,210  
  模型定义/推理/训练统一框架，是开源模型落地的基础依赖。

- [github/copilot-sdk](https://github.com/github/copilot-sdk) — 今日 +7  
  GitHub 官方 SDK，可把 Copilot Agent 集成进应用和服务，是 Agent 原生应用开发的重要基础设施。

- [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) — ⭐158,725  
  为 Agent 提供搜索、抓取、Web 交互的 API；Agent 数据入口类工具需求持续走高。

- [langchain4j/langchain4j](https://github.com/langchain4j/langchain4j) — ⭐12,757  
  JVM 上的 LLM 应用开发库，支持 MCP、RAG、tool calling，企业 Java 技术栈接入 AI 的首选桥。

- [0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig) — ⭐8,113  
  Rust 生态模块化 LLM 应用框架，适合高性能、低资源服务端场景。

- [headroomlabs-ai/headroom](https://github.com/headroomlabs-ai/headroom) — ⭐63,572  
  压缩工具输出/日志/RAG chunk 的 token，编码 Agent 平均省 20% token，是当前“Agent 降本”热门工具。

- [open-compass/opencompass](https://github.com/open-compass/opencompass) — ⭐7,251  
  覆盖 100+ 数据集的 LLM 评测平台，模型选择越来越依赖系统化评测。

### 🤖 AI 智能体/工作流

- [affaan-m/ECC](https://github.com/affaan-m/ECC) — ⭐236,638  
  Agent harness 性能优化系统，整合 skills、instincts、memory、security，是当前 Agent 基础设施的现象级项目。

- [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) — ⭐223,408  
  “陪你成长的 Agent”，主打长期学习与自我进化，代表 Agent 从工具向伙伴演进。

- [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) — ⭐185,741  
  老牌通用 Agent 项目，仍是社区理解 autonomous agent 的基准。

- [langchain-ai/langchain](https://github.com/langchain-ai/langchain) — ⭐143,114  
  Agent 工程平台，工具调用、记忆、编排的事实标准之一。

- [browser-use/browser-use](https://github.com/browser-use/browser-use) — ⭐107,421  
  让 Agent 像人一样操作浏览器，是“Agent 上网”链路的核心组件。

- [langgenius/dify](https://github.com/langgenius/dify) — ⭐150,930  
  可视化构建 Agentic workflow 与 RAG 流水线，企业级“Agent 中台”首选。

- [different-ai/openwork](https://github.com/different-ai/openwork) — 今日 +806  
  Claude Cowork 的开源替代，由 opencode 驱动；今日最受关注的协作式 Agent 项目。

- [mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill) — 今日 +658  
  可跨 Reddit、X、YouTube、HN、Polymarket 做主题调研并输出总结的 Agent Skill，展示“技能包”正在成为 Agent 能力交付新形态。

### 📦 AI 应用

- [open-webui/open-webui](https://github.com/open-webui/open-webui) — ⭐147,482  
  开源 AI 对话/管理界面，对接 Ollama、OpenAI API，几乎是本地 LLM 标配前端。

- [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) — ⭐64,167  
  Local-first 的 Agent 工作台，私有部署、自带 RAG，适合个人与团队。

- [harry0703/MoneyPrinterTurbo](https://github.com/harry0703/MoneyPrinterTurbo) — ⭐100,813  
  利用 AI 大模型和自动化工作流一键生成短视频，AI 内容生产领域现象级应用。

- [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) — ⭐49,214  
  聚合前沿 LLM 的 AI 生产力工作室，支持 300+ 助手与自主 Agent。

- [santifer/career-ops](https://github.com/santifer/career-ops) — ⭐62,402  
  开源 AI 求职助手：扫描职位、A–F 评分、定制简历并跟踪申请，是垂直 Agent 应用代表。

- [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) — ⭐59,700  
  LLM 驱动的多市场股票分析系统，集成行情、新闻、看板与自动推送。

- [hugohe3/ppt-master](https://github.com/hugohe3/ppt-master) — ⭐42,195  
  将文档或主题转为原生 PPT，支持动画、图表、配音，办公自动化 AI 应用新贵。

- [deepfakes/faceswap](https://github.com/deepfakes/faceswap) — 今日 +93  
  经典人脸合成/Deepfake 工具，持续受关注，也持续提示 AI 内容安全议题。

### 🧠 大模型/训练

- [pytorch/pytorch](https://github.com/pytorch/pytorch) — ⭐102,091  
  深度学习训练/研究主框架，所有开源模型生态的“地基”。

- [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow) — ⭐196,633  
  老牌 ML 框架，生产环境仍广泛应用。

- [keras-team/keras](https://github.com/keras-team/keras) — ⭐64,190  
  高复用深度学习 API，适合快速原型与教学。

- [rasbt/LLMs-from-scratch](https://github.com/rasbt/LLMs-from-scratch) — ⭐100,240  
  手把手用 PyTorch 从零实现 ChatGPT 类 LLM，是“造模型”路线最重要的学习资源。

- [skyzh/tiny-llm](https://github.com/skyzh/tiny-llm) — ⭐4,427  
  在 Apple Silicon 上从零构建 tiny vLLM + Qwen，推理系统工程师实战课程。

- [genieincodebottle/generative-ai](https://github.com/genieincodebottle/generative-ai) — ⭐2,580  
  生成式 AI 全栈资源，含路线图、项目、面试与编码准备。

- [AarambhDevHub/aarambh-studio](https://github.com/AarambhDevHub/aarambh-studio) — ⭐54  
  纯 Rust + Candle 从零编写的 decoder-only LLM，无 Python/PyTorch，追求极致轻量。

- [ai-glimpse/toyllm](https://github.com/ai-glimpse/toyllm) — ⭐25  
  “Learning LLM from Scratch” 的低门槛实验项目。

### 🔍 RAG/知识库

- [milvus-io/milvus](https://github.com/milvus-io/milvus) — ⭐45,440  
  云原生向量数据库，生产级 RAG 检索层主力。

- [infiniflow/ragflow](https://github.com/infiniflow/ragflow) — ⭐86,526  
  深度 RAG 引擎，融合 Agent 能力，提供面向 LLM 的上下文层。

- [qdrant/qdrant](https://github.com/qdrant/qdrant) — ⭐33,697  
  高性能向量数据库与检索引擎，AI 应用基础设施。

- [run-llama/llama_index](https://github.com/run-llama/llama_index) — ⭐51,262  
  文档 Agent 与 OCR 平台，RAG 应用开发核心框架。

- [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) — ⭐89,179  
  跨会话持久记忆层，压缩并注入 Agent 上下文，解决“金鱼记忆”问题。

- [mem0ai/mem0](https://github.com/mem0ai/mem0) — ⭐62,216  
  Agent 通用记忆层，是长期记忆赛道的代表性项目。

- [topoteretes/cognee](https://github.com/topoteretes/cognee) — ⭐29,635  
  开源 AI 记忆平台，用知识图谱给 Agent 持久长期记忆。

- [StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN) — ⭐12,753  
  “RAG on Everything”，在个人设备上运行 RAG，宣称节省 97% 存储，是轻量私域 RAG 新方向。

---

## 3. 趋势信号分析

今日社区爆发点非常集中：**Agent 技术栈正从“框架”下沉为“可插拔技能/记忆/harness”**。Trending 上的 `openwork`、`last30days-skill`、`reverse-skill`、`jcode` 均属这一波，共同特征是与 Claude Code、Cursor、Codex 等 AI 编程客户端深度绑定，以低耦合方式增强 Agent 能力。

同时，“**轻量/私域/低成本**”成为新信号：`LEANN`、`PageIndex`、`zvec`、`lancedb` 等把 RAG 压到个人设备，`headroom` 则用压缩减少 token 开支。Rust 正在加速进入 AI 应用层（`rig`、`aarambh-studio`、`jcode`），显示性能敏感型 Agent 基础设施开始选择系统级语言。

与近期开源模型密集发布（Kimi、GLM-5.2、DeepSeek、gpt-oss、Qwen 等进入 Ollama 生态）联动，本地推理 + Agent Skill 生态正进入快速裂变期。

---

## 4. 社区关注热点

- **Agent Skill 化生态**：关注 `reverse-skill`、`last30days-skill`、`Graphify-Labs/graphify`、`JuliusBrussee/caveman`。开发者不需要构建完整 Agent 框架，只需一个 Skill 包就能扩展现有 AI 编程客户端能力。
- **记忆与上下文层**：`thedotmack/claude-mem`、`mem0ai/mem0`、`topoteretes/cognee` 解决 Agent 跨会话记忆问题，是 Agent 走向长期自治的关键基础设施。
- **无向量/轻量 RAG**：`VectifyAI/PageIndex`、`StarTrail-org/LEANN`、`headroomlabs-ai/headroom` 正在降低 RAG 的存储与 token 成本，值得跟踪。
- **本地/终端推理**：`ollama/ollama` 持续聚合最新开源模型，`skyzh/tiny-llm`、`ai-glimpse/toyllm`、`alibaba/zvec` 代表从模型到向量库的“端侧化”趋势。
- **AI 安全与治理**：`RiccardoBiosas/awesome-MLSecOps`、`chrisliu298/awesome-llm-unlearning` 与今日登榜的 `reverse-skill` 一起，说明 Agent 越强大，红队、越权、遗忘与安全评测越重要。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*