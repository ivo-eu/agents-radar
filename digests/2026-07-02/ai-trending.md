# AI 开源趋势日报 2026-07-02

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-07-02 10:17 UTC

---

# 🤖 AI 开源趋势日报 | 2026-07-02

## 一、今日速览

今日 GitHub Trending 上 AI 项目占比高达 75%，**多智能体协作**与**AI 安全工具**成为最亮眼的两个主题。`agency-agents` 以 2114 颗日增星登顶，展示了 “一站式 AI 代理工厂” 的吸引力；`strix` 将 AI 注入渗透测试，刷新了安全工具的开源热度。与此同时，视频编辑智能体 `video-use`、终端 agent 多路复用器 `herdr` 等新兴方向集中涌现，表明社区正从单一 LLM 调用转向**复杂 Agent 编排与垂直应用**的爆发期。RAG 生态持续成熟，向量数据库和记忆层工具链已进入生产级竞争阶段。

## 二、各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具、CLI）

| 项目 | Stars | 一句话说明 |
|------|-------|------------|
| [ollama/ollama](https://github.com/ollama/ollama) | ⭐175,273 | 一键运行 Kimi、DeepSeek、Qwen 等主流 LLM 的本地推理引擎，社区最流行的 “模型超市” |
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | ⭐85,125 | 高吞吐、低显存占用的 LLM 推理服务引擎，生产部署标配 |
| [langchain-ai/langchain](https://github.com/langchain-ai/langchain) | ⭐140,723 | 智能体工程平台，提供从链式调用到工具编排的全套 SDK |
| [usestrix/strix](https://github.com/usestrix/strix) | ⭐0 (+1211 today) | 开源 AI 渗透测试工具，自动发现并修复应用漏洞，今日爆款 |
| [diegosouzapw/OmniRoute](https://github.com/diegosouzapw/OmniRoute) | ⭐0 (+1010 today) | 免费 AI 网关：单端点接入 231+ 模型提供商，支持 tokens 压缩与自动回退 |
| [TencentCloud/CubeSandbox](https://github.com/TencentCloud/CubeSandbox) | ⭐0 (+79 today) | 为 AI Agent 设计的即时、并发、安全轻量级沙箱 |

### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）

| 项目 | Stars | 一句话说明 |
|------|-------|------------|
| [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents) | ⭐0 (+2114 today) | 从前端开发到 Reddit 社区运营，一键部署完整 AI 代理工厂 |
| [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | ⭐185,252 | 经典通用 AI 代理框架，支持自主任务规划与执行 |
| [browser-use/browser-use](https://github.com/browser-use/browser-use) | ⭐102,114 | 让 AI 智能体像人类一样操作浏览器，自动化在线任务 |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | ⭐75,851 | 字节开源的长时序超级 Agent，支持研究、编程、创建等数小时级复杂任务 |
| [browser-use/video-use](https://github.com/browser-use/video-use) | ⭐0 (+693 today) | 用代码代理自动编辑视频，开辟 AI+视频创作新路径 |
| [ogulcancelik/herdr](https://github.com/ogulcancelik/herdr) | ⭐0 (+609 today) | 终端智能体多路复用器，并行管理多个 CLI 代理 |
| [0xNyk/council-of-high-intelligence](https://github.com/0xNyk/council-of-high-intelligence) | ⭐0 (+161 today) | 汇集亚里士多德、费曼等 18 个 AI 人格进行多轮结构化讨论 |

### 📦 AI 应用（具体产品、垂直场景解决方案）

| 项目 | Stars | 一句话说明 |
|------|-------|------------|
| [open-webui/open-webui](https://github.com/open-webui/open-webui) | ⭐143,781 | 用户友好的 AI 聊天界面，支持 Ollama 和 OpenAI API，部署即用 |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | ⭐48,049 | 集合智能聊天、自主代理和 300+ 助手的多模型生产力工作室 |
| [altic-dev/FluidVoice](https://github.com/altic-dev/FluidVoice) | ⭐0 (+572 today) | macOS 本地听写应用，AI 增强模型 + 完全离线，对标 Wispr Flow |
| [HKUDS/Vibe-Trading](https://github.com/HKUDS/Vibe-Trading) | ⭐0 (+694 today) | 个人交易智能体，基于市场情绪和量化策略自动执行 |
| [microsoft/AI-For-Beginners](https://github.com/microsoft/AI-For-Beginners) | ⭐0 (+1096 today) | 微软出品的 12 周 AI 入门课程，今天突然爆火 |
| [Unclecheng-li/VulnClaw](https://github.com/Unclecheng-li/VulnClaw) | ⭐0 (+132 today) | 基于 AI Agent + MCP 的全自动渗透测试，自然语言驱动完整攻击链 |

### 🧠 大模型/训练（模型权重、训练框架、微调工具）

| 项目 | Stars | 一句话说明 |
|------|-------|------------|
| [huggingface/transformers](https://github.com/huggingface/transformers) | ⭐162,115 | 业界标准模型定义与推理框架，支持文本、视觉、音频等多种模态 |
| [pytorch/pytorch](https://github.com/pytorch/pytorch) | ⭐101,048 | 动态神经网络框架，GPU 加速，深度学习生态基石 |
| [hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory) | ⭐72,897 | 统一高效微调 100+ LLM/VLM 的工具，ACL 2024 论文方案 |
| [jingyaogong/minimind](https://github.com/jingyaogong/minimind) | ⭐52,460 | 2 小时从零训练 64M 参数小模型，极简入门大模型训练 |
| [open-compass/opencompass](https://github.com/open-compass/opencompass) | ⭐7,143 | 全面 LLM 评测平台，支持 100+ 数据集和多模型横向对比 |

### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）

| 项目 | Stars | 一句话说明 |
|------|-------|------------|
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | ⭐84,119 | 领先的开源 RAG 引擎，融合 Agent 能力，为 LLM 提供上下文层 |
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | ⭐45,043 | 高性能云原生向量数据库，支持十亿级向量 ANN 搜索 |
| [qdrant/qdrant](https://github.com/qdrant/qdrant) | ⭐32,884 | 面向下一代 AI 的高性能向量数据库，支持云端与自建 |
| [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) | ⭐62,429 | 本地优先的 RAG 桌面应用，让你拥有自己的 AI 知识库 |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | ⭐59,927 | AI Agent 的通用记忆层，跨会话持久化上下文 |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) | ⭐85,452 | 为 Claude 等 Agent 提供持久上下文，自动压缩并注入历史会话 |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | ⭐75,581 | 将代码、文档、图片等转为可查询的知识图谱，赋能代码智能体 |

## 三、趋势信号分析

**1. Agent 框架从“单一”走向“多智能体生态”**：今日 Trending 榜单中，`agency-agents`、`council-of-high-intelligence`、`herdr` 等均涉及多代理或多角色协作，社区不再满足于单 Agent 对话，而是追求**分工、辩论、编排**等复杂行为。这与 Anthropic 近期发布的 Agent 最佳实践和 MCP 协议的推广高度契合。

**2. AI+安全赛道异军突起**：`strix`（1211 日增星）和 `VulnClaw`（132 日增星）同时出现，表明开发者正在将 LLM 的理解与规划能力注入渗透测试、漏洞利用等传统安全工具。这一方向此前主要由商业产品覆盖，今日开源社区开始发力，可能催生新的安全工具栈。

**3. 视频编辑成为 Agent 新接口**：`video-use`（693 日增星）是 `browser-use` 团队的最新尝试，将代码代理的能力从浏览器操作扩展到视频剪辑。这是多模态 Agent 落地的自然延伸——Agent 不仅能“看”网页，还能“剪”视频。类似项目反映了社区对**富媒体生成与编辑**的强烈需求。

**4. 低门槛 AI 基础设施持续升温**：`OmniRoute`（1010 日增星）提供免费的多模型网关和 tokens 压缩技术，`FluidVoice`（572 日增星）则是本地离线语音识别的成熟方案。这些工具正在降低 AI 应用开发与使用的门槛，吸引更多非专业用户参与。

## 四、社区关注热点

- 🏆 **`msitarzewski/agency-agents`** —— 今日星王，堪称“Agent 界的 WordPress”，一键部署泛化能力极强，值得深入其多 Agent 设计模式。
- 🔒 **AI 渗透测试：`usestrix/strix` 与 `Unclecheng-li/VulnClaw`** —— AI+安全融合的新范式，建议关注其工具链与 MCP 协议结合的可能性。
- 🎬 **视频 Agent：`browser-use/video-use`** —— 视频编辑自动化方向刚刚起步，适合开发者探索如何将 Agent 扩展到富媒体领域。
- 🧠 **记忆层 `mem0ai/mem0` 和 `thedotmack/claude-mem`** —— 解决 Agent 长期记忆的核心痛点，这两个项目正在快速迭代，是构建持久化 Agent 应用的基础设施。
- 🌐 **免费 AI 网关 `OmniRoute`** —— 支持 50+ 免费模型和 tokens 压缩，对独立开发者和小团队极大降低 LLM 调用成本，值得关注其后续兼容性。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*