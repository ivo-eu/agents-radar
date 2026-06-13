# AI 开源趋势日报 2026-06-13

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-13 10:15 UTC

---

# AI 开源趋势日报（2026-06-13）

## 今日速览
- **AI Agent 生态持续爆发**：今日 Trending 上涌现多个“Agent Skills”项目，如 `addyosmani/agent-skills`（+2656 stars）、`obra/superpowers`（+1275 stars）、`phuryn/pm-skills`（+827 stars），表明社区正从单 Agent 走向技能市场化、可组合化。
- **KV Cache 加速成为 LLM 推理热点**：`LMCache/LMCache` 凭借“最快 KV Cache 层”获得关注，与近期 vLLM、开源推理优化趋势一致。
- **医疗与项目管理场景的 AI 专化**：`maziyarpanahi/openmed`（开放医疗 AI）和 `msitarzewski/agency-agents`（全功能 AI 代理机构）分别切入垂直行业，显示 Agent 应用正在细分。
- **RAG 与向量数据库持续演进**：`NirDiamant/RAG_Techniques` 等高级教程类项目热度不减，同时 `StarTrail‑org/LEANN` 提出 97% 存储节约的端侧 RAG 方案，推动隐私与效率平衡。

## 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）
- **[LMCache/LMCache](https://github.com/LMCache/LMCache)** ⭐0（+28 today）  
  专为 LLM 设计的最快 KV Cache 层，显著提升推理吞吐与显存效率，与 vLLM 等引擎互补。
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐82,749  
  高吞吐、内存高效的 LLM 推理与服务引擎，已成为社区标准之一。
- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐174,017  
  一键运行本地大模型的 CLI 工具，支持 Kimi、DeepSeek、Qwen 等最新模型。
- **[Picovoice/picollm](https://github.com/Picovoice/picollm)** ⭐312  
  端侧 LLM 推理库，采用 X-Bit 量化技术，适配边缘设备。
- **[skyzh/tiny-llm](https://github.com/skyzh/tiny-llm)** ⭐4,272  
  面向系统工程师的 LLM 推理 Serving 实战课程，基于 Apple Silicon 搭建迷你 vLLM + Qwen。
- **[scikit-learn/scikit-learn](https://github.com/scikit-learn/scikit-learn)** ⭐66,316  
  经典机器学习框架，今日因其稳定性和广泛生态仍被列入 ML 主题。

### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）
- **[addyosmani/agent-skills](https://github.com/addyosmani/agent-skills)** ⭐0（+2656 today）  
  面向 AI 编码代理的生产级技能集合（Shell 脚本），让 Agent 直接调用工程最佳实践。
- **[obra/superpowers](https://github.com/obra/superpowers)** ⭐0（+1275 today）  
  Agent 技能框架与软件开发方法论，强调“可工作”的代理协作流程。
- **[msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents)** ⭐0（+1026 today）  
  完整的 AI 代理“机构”，包含前端、Reddit、幽默注入等 10+ 专业 Agent，一键部署。
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐192,307  
  自我成长的通用 Agent 框架，热度极高，支持技能积累与工具调用。
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐184,923  
  开创性自主 Agent 项目，持续推动 AI 自动任务执行标准。
- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** ⭐98,570  
  让 AI Agent 能像人类一样操作浏览器，网页自动化任务首选工具。
- **[shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code)** ⭐66,352  
  从零实现类 Claude Code 的 Agent Harness，教学与实用兼具。

### 📦 AI 应用（具体应用产品、垂直场景解决方案）
- **[maziyarpanahi/openmed](https://github.com/maziyarpanahi/openmed)** ⭐0（+515 today）  
  开源医疗 AI 平台，汇聚诊断、文献检索、影像分析等模型，面向医疗垂直场景。
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐47,263  
  AI 生产力工作室：智能对话、自主 Agent、300+ 助手，统一调用前沿 LLM。
- **[zhayujie/CowAgent](https://github.com/zhayujie/CowAgent)** ⭐45,267  
  轻量级超级 AI 助手（前 chatgpt-on-wechat），支持任务规划、工具调用、记忆演化。
- **[OpenBB-finance/OpenBB](https://github.com/OpenBB-finance/OpenBB)** ⭐69,042  
  金融数据平台，为分析师和 AI Agent 提供统一接口，已集成多种数据源。
- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** ⭐42,368  
  LLM 驱动的 A/H/美股票智能分析系统，零成本定时运行，深受散户投资者欢迎。
- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** ⭐27,113  
  AI 从任意文档生成可编辑 PowerPoint（含动画、语音旁白、模板支持）。

### 🧠 大模型/训练（模型权重、训练框架、微调工具）
- **[tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)** ⭐195,631  
  开源机器学习框架，持续支持 TPU 与生产级部署。
- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** ⭐100,711  
  动态神经网络框架，GPU 加速，科研与工业界首选。
- **[hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)** ⭐72,127  
  统一高效微调 100+ LLM/VLM，ACL 2024 收录，成为社区微调标配。
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7,081  
  开源 LLM 评测平台，支持 Llama3、Mistral、GPT-4 等 100+ 数据集。
- **[LiberCoders/FeatureBench](https://github.com/LiberCoders/FeatureBench)** ⭐75（今天新上榜）  
  ICLR 2026 官方实现——复杂功能开发的 Agentic 编码基准测试。
- **[thinkwee/AwesomeOPD](https://github.com/thinkwee/AwesomeOPD)** ⭐622  
  On‑Policy Distillation 资源合集，关注 LLM 知识蒸馏最新进展。

### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）
- **[langgenius/dify](https://github.com/langgenius/dify)** ⭐145,045  
  生产级 Agentic 工作流开发平台，内置 RAG、工具调用、多模型支持。
- **[open-webui/open-webui](https://github.com/open-webui/open-webui)** ⭐141,313  
  用户友好的 AI 界面（支持 Ollama、OpenAI），成为本地部署首选前端。
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐82,616  
  领先的开源 RAG 引擎，融合 Agent 能力，提供 LLM 上下文层。
- **[PaddlePaddle/PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)** ⭐82,072  
  图片/PDF → 结构化数据 → LLM 的 OCR 工具包，支持 100+ 语言。
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐58,469  
  通用 AI Agent 记忆层，跨会话持久化上下文，增强 Agent 连续性。
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐44,759  
  高性能云原生向量数据库，大规模向量 ANN 搜索引擎。
- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** ⭐11,914  
  MLsys2026 论文实现：RAG 存储压缩 97%，在个人设备上实现快速、准确、隐私的端侧 RAG。
- **[zilliztech/claude-context](https://github.com/zilliztech/claude-context)** ⭐11,834  
  Claude Code 的 MCP 代码搜索工具，让整个代码库成为 Agent 上下文。

## 趋势信号分析
1. **Agent 技能市场化**：`addyosmani/agent-skills` 与 `obra/superpowers` 的爆发式增长（今日合计 +3931 stars）表明，社区正从“构建单个 Agent”转向“标准化、可复用的 Agent Skills”。类似 Docker Hub 的“Agent 技能市场”正在形成。
2. **垂直场景 Agent 成新增长点**：医疗（`openmed`）、金融（`daily_stock_analysis`）、项目管理（`pm-skills`）等领域的专门 Agent 项目密集出现，显示 AI 代理正从通用编码助手向行业解决方案渗透。
3. **推理加速与端侧化并行**：`LMCache` 专注于 KV Cache 优化，呼应了 vLLM 的流行；同时 `Picovoice/picollm`、`LEANN` 等端侧项目提示开发者关注低资源部署与隐私保护。
4. **RAG 成熟度提升**：从零散应用（如 `anything-llm`）走向系统化平台（`dify`、`ragflow`），且出现高级教程（`RAG_Techniques`）和存储优化方案（`LEANN`），RAG 技术栈已进入精细化竞争阶段。
5. **与行业事件关联**：今日 Trending 中 `agent-skills` 的爆火与 2026 年 6 月各大模型厂商（OpenAI、Anthropic）推出“Agent SDK”的浪潮高度相关，开源社区正积极围堵生态缺口。

## 社区关注热点
- **Agent Skills 标准化（addyosmani/agent-skills）**：今日 stars 增长 2656，展示了如何用 Shell 脚本封装代码审查、环境排查等工程技能，值得 Agent 开发者参考和贡献。
- **端侧 RAG 极致压缩（StarTrail-org/LEANN）**：97% 存储节约配合本地推理，为移动端、IoT 设备上的私有知识库提供了可行路径。
- **全栈 Agent 机构（msitarzewski/agency-agents）**：从文案、客服到市场推广的 Agent 组合，验证了“多 Agent 协同”的商业模式可行性。
- **复杂功能编码基准（LiberCoders/FeatureBench）**：ICLR 2026 新基准，专注评估 Agent 在真实软件功能开发中的表现，可能成为继 SWE‑bench 后的新标准。
- **LLM 记忆层（mem0ai/mem0）**：跨会话记忆是 Agent 实用化的关键短板，该项目提供统一解决方案，已与 Claude Code、Hermes Agent 等集成。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*