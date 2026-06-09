# AI 开源趋势日报 2026-06-09

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-06-09 04:18 UTC

---

# AI 开源趋势日报（2026-06-09）

## 一、今日速览

今日 GitHub 热门趋势中，**AI Agent 生态全面爆发**：跨平台信息聚合技能 `last30days-skill` 单日斩获 3558 stars，向量索引工具 `turbovec` 日增 1729 stars，视觉工具 `roboflow/supervision` 日增 1288 stars，开源 Agent `goose` 日增 699 stars。Google 和 OpenAI 官方出手，分别发布 `google/skills` 和 `plugins`，加速 Agent 标准化。同时，本地 LLM 评测工具 `whichllm` 和 AI 记忆系统 `MemPalace` 的走红，反映出社区对**可落地、可本地运行**的 AI 基础设施的强烈需求。

## 二、各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、CLI）

- **[turbovec](https://github.com/RyanCodrai/turbovec)** – ⭐0 (今日 +1729) · 基于 TurboQuant 的向量索引，Rust 编写 + Python 绑定，高性能向量检索新选择。
- **[whichllm](https://github.com/Andyyyy64/whichllm)** – ⭐0 (今日 +143) · 一键评测本地 LLM 实际性能，按基准排序，无需数参数。
- **[roboflow/supervision](https://github.com/roboflow/supervision)** – ⭐0 (今日 +1288) · 可复用的计算机视觉工具库，专注目标检测、跟踪等工程化能力。
- **[openai/plugins](https://github.com/openai/plugins)** – ⭐0 (今日 +296) · OpenAI 官方插件框架，提供 Agent 可调用的标准化接口。
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** – ⭐82,270 · 高吞吐量 LLM 推理引擎，生产级性能。
- **[ollama/ollama](https://github.com/ollama/ollama)** – ⭐173,636 · 一键运行本地 LLM 的经典工具，支持主流开源模型。

### 🤖 AI 智能体 / 工作流（Agent 框架、自动化、多智能体）

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** – ⭐0 (今日 +3558) · AI Agent 技能：自动调研 Reddit、X、YouTube 等多源信息并生成总结摘要。
- **[aaif-goose/goose](https://github.com/aaif-goose/goose)** – ⭐0 (今日 +699) · 开源可扩展 Agent，超越代码建议，可安装、执行、编辑、测试，兼容任意 LLM。
- **[Panniantong/Agent-Reach](https://github.com/Panniantong/Agent-Reach)** – ⭐0 (今日 +679) · 让 AI Agent 无 API 费用地读取 Twitter、Reddit、Bilibili 等全网内容。
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** – ⭐0 (今日 +378) · Agent 与生成式 UI 的前端框架，支持 React、Angular、Mobile 等。
- **[google/skills](https://github.com/google/skills)** – ⭐0 (今日 +461) · Google 官方推出的 Agent 技能，覆盖 Google 产品和技术栈。
- **[danielmiessler/Personal_AI_Infrastructure](https://github.com/danielmiessler/Personal_AI_Infrastructure)** – ⭐0 (今日 +62) · 面向个人能力的 Agentic AI 基础设施，聚焦增强人类。
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** – ⭐76,279 · AI 驱动的软件开发 Agent，社区长期热门。

### 📦 AI 应用（具体场景、垂直解决方案）

- **[santifer/career-ops](https://github.com/santifer/career-ops)** – ⭐0 (今日 +308) · AI 驱动的求职系统：14 种技能模式、Go 仪表盘、PDF 简历生成。
- **[phuryn/pm-skills](https://github.com/phuryn/pm-skills)** – ⭐0 (今日 +164) · 产品经理技能市场：100+ Agent 技能、命令和插件，覆盖产品全周期。
- **[luongnv89/claude-howto](https://github.com/luongnv89/claude-howto)** – ⭐0 (今日 +312) · Claude Code 的可视化指南，从基础到高级 Agent，附带可直接使用的模板。
- **[Panniantong/Agent-Reach](https://github.com/Panniantong/Agent-Reach)** （已在上文智能体类出现，此处不重复）
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** – ⭐47,086 · AI 生产力工作室，集成聊天、自主 Agent 和 300+ 助手。

### 🧠 大模型 / 训练（模型、微调、评估）

- **[hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)** – ⭐72,006 · 统一高效微调框架，支持 100+ LLM 和 VLM（ACL 2024）。
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** – ⭐7,069 · LLM 评估平台，支持 Llama3、Qwen、GPT-4 等多模型，数据集 100+。
- **[skyzh/tiny-llm](https://github.com/skyzh/tiny-llm)** – ⭐4,259 · 面向系统工程师的 LLM 推理服务实战课程（基于 Apple Silicon）。
- **[galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining)** – ⭐251 · 可靠、可扩展的基础模型与世界模型预训练库。

### 🔍 RAG / 知识库（向量数据库、检索增强、知识管理）

- **[MemPalace/mempalace](https://github.com/MemPalace/mempalace)** – ⭐0 (今日 +170) · 开源 AI 记忆系统，基准测试最佳，免费使用。
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** – ⭐58,097 · 通用 AI Agent 记忆层，支持长期上下文保持。
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** – ⭐82,244 · 领先的开源 RAG 引擎，融合 Agent 能力与知识检索上下文。
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** – ⭐44,689 · 高性能云原生向量数据库，专为可扩展向量 ANN 搜索构建。
- **[Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm)** – ⭐61,271 · 本地优先的 Agent 体验平台，支持多种 LLM 与 RAG。

## 三、趋势信号分析

今日热榜呈现出 **“Agent 技能 + 高性能基础设施”双轮驱动**的特征：

1. **Agent 技能市场爆发**：`last30days-skill`（3558 stars）和 `pm-skills` 等“技能市场”类项目走红，说明社区不再满足于单一 Agent 框架，而开始追求**可组合、可交易**的 Agent 技能生态。这表明 Agent 正从“玩具”走向“工具”。
2. **本地化与性能优化突显**：`turbovec`（向量索引）和 `whichllm`（本地 LLM 评测）的兴起，反映出开发者对**低成本、可私有化部署**的 AI 组件的急切需求；`MemPalace` 强调“免费”和“最佳基准”也呼应了这一趋势。
3. **大厂加速 Agent 标准化**：Google 发布 `google/skills`，OpenAI 更新 `plugins`，官方以“技能”和“插件”形式定义 Agent 能力边界，可能主导未来 Agent 生态的互操作标准。
4. **计算机视觉持续回热**：`roboflow/supervision` 单日 1288 stars，表明 CV 工程化工具在 Agent 需要“看世界”的背景下重新获得关注。

## 四、社区关注热点

- 🚀 **`last30days-skill`** – 跨平台信息聚合 Agent，适合研究、监测类场景，日增 Stars 断层第一，值得深入体验其效果与可扩展性。
- 🛠️ **`goose` (aaif-goose)** – 开源可扩展 Agent，超越代码建议，支持安装/执行/测试，兼容任意 LLM，可作为构建自定义 Agent 的首选底座。
- 🧠 **`MemPalace`** – 开源 AI 记忆系统，基准最佳，解决了 Agent 长期记忆这个核心痛点，值得关注其与 RAG 的区别与融合。
- 📐 **`whichllm`** – 一键评测本地 LLM 实际性能，为硬件选型和模型对比提供客观依据，实用性强。
- 🔌 **`google/skills`** – 官方 Agent 技能，可能成为未来 Agent 技能的标准模板，建议开发者提前了解其设计与 API 规范。