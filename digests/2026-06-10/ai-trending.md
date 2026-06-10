# AI 开源趋势日报 2026-06-10

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-10 05:08 UTC

---

# AI 开源趋势日报（2026-06-10）

## 今日速览

- **Agent 技能生态爆发**：Trending 榜出现 4 个与“AI Agent 技能/工具”相关的新项目（`last30days-skill`、`pm-skills`、`agent-skills`、`system-prompts-and-models`），总新增 stars 超 5000，开发者正疯狂为 Claude Code、Codex 等 CLI Agent 构建可复用技能库。
- **本地 LLM 评测工具首次登榜**：`whichllm` 以 +633 stars 冲入 Trending，反映社区对“在自有硬件上跑通最强模型”的迫切需求，与近期大量小型高性能模型（如 GLM-5.1、Kimi K2.6）发布直接相关。
- **AI 驱动的求职/职业自动化成为新热点**：`career-ops`（+1110 stars）将 AI Agent 技能模式扩展到求职场景，14 种技能模式 + Go 仪表盘，凸显“AI 工作流 + 垂直场景”的爆发潜力。
- **向量数据库赛道持续火热**：Milvus、Qdrant、Weaviate、LanceDB 等明星项目均保持活跃，同时 `turbovec`（Rust 编写向量索引）以 +1801 stars 成为今日最大黑马，暗示混合量化方案在嵌入式场景的兴起。

## 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、CLI 等）

| 项目 | Stars 总量 | 今日新增 | 一句话说明 |
|------|-----------|----------|-----------|
| 🥇 [RyanCodrai/turbovec](https://github.com/RyanCodrai/turbovec) | 0 | +1801 | 基于 Rust 的 TurboQuant 向量索引库，高性能 Python 绑定，定义下一代嵌入式向量搜索。 |
| 🥈 [opencv/opencv](https://github.com/opencv/opencv) | 78k+ | +102 | 经典计算机视觉库，今日小幅更新，仍是 CV 基础设施的首选。 |
| 🥉 [roboflow/supervision](https://github.com/roboflow/supervision) | 20k+ | +733 | 可复用的计算机视觉工具库，提供标注、追踪、可视化等模块，降低 CV 应用开发门槛。 |
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | 82,369 | — | 高吞吐 LLM 推理引擎，持续被社区选为生产级部署首选。 |
| [huggingface/transformers](https://github.com/huggingface/transformers) | 161,467 | — | 模型定义与推理框架，社区生态的基石，支持最新发布的 GLM-5.1 等。 |
| [ollama/ollama](https://github.com/ollama/ollama) | 173,726 | — | 一键运行本地大模型，已支持 Kimi-K2.6、DeepSeek 等最新模型。 |
| [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) | 130,814 | — | 大规模网页搜索与抓取 API，为 LLM 提供实时互联网数据。 |
| [browser-use/browser-use](https://github.com/browser-use/browser-use) | 98,013 | — | 让 AI Agent 能自动操作浏览器，自动化在线任务。 |

### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）

| 项目 | Stars 总量 | 今日新增 | 一句话说明 |
|------|-----------|----------|-----------|
| 🥇 [mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill) | 0 | +3191 | AI Agent 技能：自动研究 Reddit/X/YouTube 等平台话题并生成综合摘要，今日最火 Trending 项目。 |
| 🥈 [santifer/career-ops](https://github.com/santifer/career-ops) | 51,906 | +1110 | AI 驱动的求职系统，14 种 Agent 技能模式（简历、面试、网络），Go 仪表盘 + PDF 生成。 |
| 🥉 [aaif-goose/goose](https://github.com/aaif-goose/goose) | 1.3k+ | +489 | 开源可扩展 AI Agent，超越代码补全，支持安装、执行、编辑、测试，配合任意 LLM。 |
| [addyosmani/agent-skills](https://github.com/addyosmani/agent-skills) | 最新 | +443 | 为 AI 编码 Agent 提供的生产级工程技能集合（Shell 脚本），教 Agent 写测试、部署、监控。 |
| [phuryn/pm-skills](https://github.com/phuryn/pm-skills) | 0 | +806 | 产品经理的 100+ 种 AI Agent 技能，从需求发现到产品启动，覆盖全生命周期。 |
| [x1xhlol/system-prompts-and-models-of-ai-tools](https://github.com/x1xhlol/system-prompts-and-models-of-ai-tools) | 最新 | +79 | 收集 Claude Code、Cursor、Windsurf 等 30+ 工具的系统提示与模型信息，Agent 开发者必看。 |
| [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) | 189,061 | — | 重磅 Agent 框架，“与你一起成长的智能体”，支持插件、记忆、多模态。 |
| [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | 184,861 | — | 经典自主 AI Agent 项目，持续迭代为通用智能体平台。 |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | 70,848 | — | 字节跳动开源的长时域 SuperAgent，支持研究、编码、创作，集成沙箱与记忆。 |

### 📦 AI 应用（具体产品、垂直场景解决方案）

| 项目 | Stars 总量 | 今日新增 | 一句话说明 |
|------|-----------|----------|-----------|
| 🥇 [maziyarpanahi/openmed](https://github.com/maziyarpanahi/openmed) | 0 | +191 | 开源医疗 AI 平台，瞄准医疗场景的数据分析与辅助诊断。 |
| 🥈 [hugohe3/ppt-master](https://github.com/hugohe3/ppt-master) | 25,707 | — | AI 从文档自动生成可编辑 PPT，支持原生形状、动画、配音，遵循模板。 |
| 🥉 [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) | 41,565 | — | LLM 驱动的中美港股票分析系统，零成本定时运行，多数据源+决策仪表盘。 |
| [openai/plugins](https://github.com/openai/plugins) | 最新 | +284 | OpenAI 官方插件库，Agent 生态扩展的基础设施，今日可能更新了插件标准。 |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | 47,143 | — | AI 生产力工作室：智能聊天+自主Agent+300+助手，统一接入前沿 LLM。 |

### 🧠 大模型/训练（模型权重、训练框架、微调工具、评测）

| 项目 | Stars 总量 | 今日新增 | 一句话说明 |
|------|-----------|----------|-----------|
| 🥇 [Andyyyy64/whichllm](https://github.com/Andyyyy64/whichllm) | 0 | +633 | 一行命令找出本地最佳 LLM：基于真实、时效性基准排名，而非参数数量。 |
| 🥈 [hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory) | 72,038 | — | 统一高效微调 100+ 模型，ACL 2024 论文实现，仍是微调首选。 |
| 🥉 [open-compass/opencompass](https://github.com/open-compass/opencompass) | 7,075 | — | 全面 LLM 评测平台，支持 100+ 数据集和主流模型。 |
| [galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining) | 252 | — | 可靠、可扩展的基础模型预训练库，面向世界模型。 |
| [skyzh/tiny-llm](https://github.com/skyzh/tiny-llm) | 4,265 | — | 在 Apple Silicon 上构建 tiny vLLM + Qwen 的学习课程，适合系统工程师。 |

### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）

| 项目 | Stars 总量 | 今日新增 | 一句话说明 |
|------|-----------|----------|-----------|
| 🥇 [milvus-io/milvus](https://github.com/milvus-io/milvus) | 44,707 | — | 高性能云原生向量数据库，支撑大规模 ANN 搜索，Agent 记忆层首选。 |
| 🥈 [qdrant/qdrant](https://github.com/qdrant/qdrant) | 31,990 | — | Rust 编写的高性能向量数据库，支持云原生部署。 |
| 🥉 [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | 82,342 | — | 领先开源 RAG 引擎，融合 Agent 能力，提供 LLM 上下文层。 |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | 58,220 | — | AI Agent 的通用记忆层，跨会话持久化上下文。 |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) | 81,503 | — | 为 Claude Code 等 Agent 提供跨会话上下文压缩与注入。 |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | 64,362 | — | 将代码、文档、SQL 等转为可查询的知识图谱，Agent 技能型 RAG。 |
| [NirDiamant/RAG_Techniques](https://github.com/NirDiamant/RAG_Techniques) | 27,806 | — | 系统化 RAG 技术教程，覆盖高级检索、混合搜索、评估等。 |

## 趋势信号分析

今日社区爆发性关注集中在 **AI Agent 技能生态** 和 **本地模型评测工具** 两个方向。Trending 前两名 `last30days-skill`（+3191）和 `turbovec`（+1801）分别代表“Agent 能做什么”和“Agent 用什么工具”。前者是纯应用层技能包，后者是底层向量索引基础设施，两者共同构建 Agent 的“感知-记忆-行动”闭环。

**新兴技术栈信号**：`whichllm` 首次登榜，结合近一周 Kimi-K2.6、GLM-5.1 等多款模型密集发布，开发者需要工具快速筛选“自己硬件上跑得最快最好的模型”，反映边缘计算和本地部署的持续热潮。同时 `turbovec` 采用 Rust + Python 绑定，暗示混合量化（TurboQuant）在嵌入式场景的成熟度提升。

**行业事件关联**：今日 OpenAI 更新了其插件库（`openai/plugins`+284），可能是为了配合即将发布的 Agent SDK 升级；而 `system-prompts-and-models-of-ai-tools` 项目收集了 30+ 工具的提示词架构，说明 Agent 领域正在形成“提示词工程公共知识库”需求。此外，`pm-skills` 和 `career-ops` 将 Agent 技能从工程领域扩展到产品管理、求职等业务场景，预示 Agent 技能市场即将爆发。

## 社区关注热点（开发者重点追踪）

- 🚀 **Agent 技能市场（Skill Marketplace）**：`last30days-skill` 和 `phuryn/pm-skills` 表明，为 Agent 编写可复用、可组合的技能包正成为新开发范式。建议关注 `addyosmani/agent-skills` 的生产级工程技能。
- 🧪 **本地模型选择工具**：`whichllm` 解决“我用哪个模型”的痛点，配合 Ollama 和 vLLM，未来可能演变为模型评测标准层。
- 🔩 **Rust 向量索引**：`turbovec` 证明 Rust 在 AI 基础设施中的潜力，与 Qdrant、LanceDB 形成 Rust 向量生态。
- 💼 **AI 职业自动化**：`career-ops` 展示了 Agent 技能如何改造传统求职流程，后续可能催生“Agent 即服务”新商业模式。
- 🩺 **开源医疗 AI**：`openmed` 虽然在早期，但 +191 stars 反映医疗垂直领域对开源方案的需求，值得关注其技术栈。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*