# AI 开源趋势日报 2026-06-15

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-15 03:43 UTC

---

## 2026-06-15 AI 开源趋势日报

### 今日速览
- **AI Agent 安全成为新焦点**：NVIDIA 发布的 SkillSpector 今日暴增 964 stars，专门扫描 AI Agent 技能中的安全漏洞，反映出社区对 Agent 安全性的高度关注。
- **金融领域大模型项目登榜**：`shiyu-coder/Kronos` 作为金融市场的基座模型，今日 +244 stars，显示 AI 在垂直行业（金融）的落地加速。
- **统一 AI 接口工具受追捧**：`andrewyng/aisuite` 今日 +291 stars，为多生成式 AI 提供商提供统一接口，降低了开发者切换成本。
- **开源 Agent 生态持续爆发**：`NousResearch/hermes-agent`（193k stars）、`langgenius/dify`（145k stars）等项目持续领跑，Agent 框架和平台仍是最大热点。
- **RAG 技术栈趋于成熟**：`infiniflow/ragflow`（82k stars）、`mem0ai/mem0`（58k stars）等项目迭代加速，RAG 正从工具走向平台化。

---

### 各维度热门项目

#### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）
- **[andrewyng/aisuite](https://github.com/andrewyng/aisuite)** ⭐0 (+291 today)  
  统一接口访问多个 Generative AI 提供商（如 OpenAI、Anthropic 等），降低切换成本，适合多模型应用开发。
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐82,865  
  高性能 LLM 推理引擎，支持 PagedAttention 等优化，是本地部署和云端推理的首选。
- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐174,183  
  一键运行本地大模型（Kimi、DeepSeek、Qwen 等），极大简化了自托管推理流程。
- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐161,591  
  最流行的模型定义与训练框架，支持文本、视觉、音频等多模态模型。
- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** ⭐100,760  
  深度学习主力框架，AI 研究者和工程师的基础工具。
- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** ⭐132,837  
  大规模网页搜索、抓取和数据交互 API，专为 AI Agent 设计，让模型能实时获取网络信息。
- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** ⭐98,847  
  让 AI Agent 能够自动化操作网页，是浏览器自动化领域的新星，与 Puppeteer 形成互补。

---

#### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐193,664  
  成长型智能体框架，支持技能学习与长期记忆，社区活跃度极高。
- **[langgenius/dify](https://github.com/langgenius/dify)** ⭐145,229  
  生产级 Agent 工作流开发平台，支持可视化编排、工具调用、RAG 集成，已获大量企业采用。
- **[langchain-ai/langchain](https://github.com/langchain-ai/langchain)** ⭐139,303  
  Agent 工程老牌框架，提供链式调用、工具集成、多模交互能力，生态丰富。
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐184,943  
  自主 Agent 先驱，实现任务分解、自我反思、长期记忆，至今仍是 Agent 开发标杆。
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** ⭐77,089  
  AI 驱动的开发工具，让 Agent 直接参与代码编写、调试和项目管理。
- **[FlowiseAI/Flowise](https://github.com/FlowiseAI/Flowise)** ⭐53,581  
  可视化构建 AI Agent 和 RAG 流水线，无需编码即可搭建复杂的多步骤推理流程。
- **[NVIDIA/SkillSpector](https://github.com/NVIDIA/SkillSpector)** ⭐0 (+964 today)  
  专注于 AI Agent 技能的安全扫描器，可检测恶意模式、漏洞和风险，是今日最受关注的新项目。

---

#### 📦 AI 应用（具体应用产品、垂直场景解决方案）
- **[shiyu-coder/Kronos](https://github.com/shiyu-coder/Kronos)** ⭐0 (+244 today)  
  金融领域的基础模型，专门处理金融语言（市场数据、研报、新闻），今日登榜显示行业大模型需求旺盛。
- **[TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents)** ⭐86,210  
  多智能体金融交易框架，结合 LLM 进行市场分析和自动交易，是 AI+金融的明星项目。
- **[PaddlePaddle/PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)** ⭐82,221  
  支持 100+ 语言的 OCR 工具，可将 PDF/图片转为结构化数据，是 RAG 和文档自动化的重要组件。
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐47,336  
  集智能对话、自主 Agent、300+ 助手于一体的 AI 生产力工作室，统一访问前沿 LLM。
- **[santifer/career-ops](https://github.com/santifer/career-ops)** ⭐53,754  
  AI 驱动的求职系统，支持 14 种技能模式、PDF 生成和批量处理，是垂直场景 AI 应用的典型。
- **[ScrapeGraphAI/Scrapegraph-ai](https://github.com/ScrapeGraphAI/Scrapegraph-ai)** ⭐27,207  
  基于 AI 的智能爬虫，可根据用户意图自主提取网页信息，替代传统规则爬虫。

---

#### 🧠 大模型/训练（模型权重、训练框架、微调工具）
- **[tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)** ⭐195,659  
  老牌机器学习框架，仍被大量企业用于模型训练和生产部署。
- **[ultralytics/ultralytics](https://github.com/ultralytics/ultralytics)** ⭐58,389  
  YOLO 系列目标检测框架，持续迭代，是计算机视觉领域的训练和推理首选。
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐82,865  
  除推理外，也支持微调（如 LoRA），是训练-推理一体化工具。
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7,084  
  LLM 评估平台，支持 100+ 数据集，帮助开发者选择合适的模型权重。
- **[acon96/home-llm](https://github.com/acon96/home-llm)** ⭐1,358  
  在 Home Assistant 中集成本地 LLM 进行智能家居控制，体现中小型模型的嵌入式应用。

---

#### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐82,738  
  领先的开源 RAG 引擎，融合 Agent 能力，提供高质量上下文层，支持企业级部署。
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐58,572  
  AI Agent 的通用记忆层，实现跨会话的长期记忆，是 RAG 与 Agent 结合的关键组件。
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐44,778  
  高性能云原生向量数据库，支持十亿级向量 ANN 搜索，是 RAG 的标配存储。
- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** ⭐32,278  
  高可扩展向量搜索引擎，强调高性能和云原生，深受 AI 应用开发者欢迎。
- **[run-llama/llama_index](https://github.com/run-llama/llama_index)** ⭐50,127  
  LlamaIndex 是文档 Agent 和 OCR 平台，提供数据索引、检索和 Agent 集成功能。
- **[siyuan-note/siyuan](https://github.com/siyuan-note/siyuan)** ⭐44,453  
  隐私优先的自托管知识管理软件，支持 AI 辅助笔记和知识图谱，适合个人或团队打造私有 RAG 知识库。
- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** ⭐82,306  
  为每个 Agent 提供持久上下文，自动压缩历史对话并注入相关上下文，实现跨 session 记忆。

---

### 趋势信号分析

今日开源社区呈现出三个明显趋势：

1. **AI Agent 安全从幕后走向前台**：NVIDIA 的 `SkillSpector` 单日新增 964 stars，表明随着 Agent 部署量激增，安全审计和恶意技能检测成为刚需。这标志着 Agent 生态从“能跑就行”进入“安全可靠”的新阶段。

2. **垂直领域大模型加速落地**：`Kronos`（金融）和 `TradingAgents`（金融交易）的登榜，说明社区不再满足于通用大模型，而是追求在特定行业（如金融、医疗、法律）中真正产生业务价值的专用模型。结合 `Career-Ops`（求职）等项目，AI 正在渗透到具体职业场景。

3. **工具链走向统一与轻量化**：`andrewyng/aisuite` 提供统一接口，`ollama` 让本地推理一键化，`lancedb` 强调嵌入式向量库。开发者越来越青睐“开箱即用”且不锁定供应商的工具。同时，Rust 生态的 `rig`、`lancedb` 等项目持续涌现，性能优先成为新选择。

**与近期行业事件的关联**：今日数据未显示直接与某次重要模型发布相关，但 `hermes-agent` 和 `dify` 的持续高热，印证了 2026 年 Q2“Agent 即应用”的主流思路。安全扫描器 `SkillSpector` 的出现，可能受近期多起 AI 供应链攻击事件的推动。

---

### 社区关注热点

- **🥇 NVIDIA SkillSpector**（[链接](https://github.com/NVIDIA/SkillSpector)）—— 既然 Agent 越来越强，谁能保证它不被恶意利用？此项目为 Agent 技能提供类似“杀毒软件”的能力，是每个 Agent 开发者都该关注的安全基础设施。
- **🥇 shiyu-coder/Kronos**（[链接](https://github.com/shiyu-coder/Kronos)）—— 金融垂直领域的“基石模型”，如果能跑通，将开启行业预训练模型的新赛道（如法律、医疗、制造）。
- **🔥 mem0ai/mem0**（[链接](https://github.com/mem0ai/mem0)）—— 记忆是 Agent 智能的关键瓶颈。Mem0 将长期记忆模块化，可与任何 Agent 框架集成，是解决“Agent 金鱼脑”的核心组件。
- **🔥 opencompass/opencompass**（[链接](https://github.com/open-compass/opencompass)）—— 模型选型越来越难，OpenCompass 提供超 100 个维度的评测基准，帮助开发者找到最适合自己场景的模型。
- **💡 Introduction-to-Autonomous-Robots**（[链接](https://github.com/Introduction-to-Autonomous-Robots/Introduction-to-Autonomous-Robots)）—— 今日 +293 stars，虽为教材，但反映了“机器人 + LLM”交叉学科的升温，是从事具身智能或机器人控制的入门首选。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*