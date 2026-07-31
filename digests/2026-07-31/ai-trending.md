# AI 开源趋势日报 2026-07-31

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-07-31 00:15 UTC

---

# AI 开源趋势日报 | 2026-07-31

## 1. 今日速览

- **语音 Agent 方向爆发**：Hugging Face 推出 `speech-to-speech` 项目，一日斩获 628 星，本地语音代理成为新热点。
- **Agent 工具链持续繁荣**：`different-ai/openwork`（+915 stars）作为 Claude Cowork 的开源替代登顶，`affaan-m/ECC`（+804 stars）和 `mvanhorn/last30days-skill`（+378 stars）进一步丰富了 agent 技能生态。
- **RAG + 向量数据库热度不减**：`infiniflow/ragflow`、`mem0ai/mem0` 等成熟项目稳定增长，同时 `StarTrail-org/LEANN` 以 97% 存储压缩技术切入边缘设备 RAG 场景。
- **中文社区呼声高**：`shareAI-lab/learn-claude-code` 用纯 Bash 实现轻量 agent，`ZhuLinsen/daily_stock_analysis` 基于 LLM 的股票分析系统均获大量关注。
- **基础框架与训练工具持续迭代**：`huggingface/transformers`、`langchain-ai/langchain` 等老牌项目 stars 突破 16 万/14 万，`The-Pocket/PocketFlow` 以 100 行代码的 LLM 框架引发新讨论。

## 2. 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、开发工具）

- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐163,181  
  🤗 统一的模型定义框架，支持文本/视觉/音频/多模态的推理与训练，是 LLM 生态的核心基础设施。

- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐177,332  
  一键运行 Kimi、DeepSeek、Qwen 等最新模型的本地推理引擎，快速迭代支持国产模型。

- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** ⭐158,355  
  面向 AI Agent 的大规模网页搜索/抓取/交互 API，成为 agent 获取实时数据的标准工具。

- **[ChromeDevTools/chrome-devtools-mcp](https://github.com/ChromeDevTools/chrome-devtools-mcp)** ⭐0 (+80 today)  
  为编码 agent 提供 Chrome DevTools 的 MCP 接口，让 agent 能直接操控浏览器进行自动化测试与调试。

- **[0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig)** ⭐8,105  
  用 Rust 构建模块化 LLM 应用的高性能框架，适合追求极致推理速度的团队。

### 🤖 AI 智能体/工作流（Agent 框架、自动化、多智能体）

- **[langgenius/dify](https://github.com/langgenius/dify)** ⭐150,839  
  集成 Agentic Workflow、RAG 管线和多模型支持的协作平台，从原型到生产零重构。

- **[AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐185,757  
  面向所有人的可访问 AI 愿景，赋予 agent 自主规划与执行任务的能力。

- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** ⭐107,334  
  让 AI Agent 能像人一样操作浏览器，自动化完成在线任务，是 agent 落地互联网的关键组件。

- **[different-ai/openwork](https://github.com/different-ai/openwork)** [TypeScript] ⭐0 (+915 today)  
  开源替代 Claude Cowork，支持 24/7 后台运行的 Agent 协作工具，兼容多种 CLI 代理。

- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐222,874  
  与用户共同成长的 agent 框架，强调持续学习和记忆能力。

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** ⭐0 (+378 today)  
  AI Agent 技能，可跨 Reddit、X、YouTube 等平台研究任意主题并生成综合摘要。

### 📦 AI 应用（具体产品、垂直场景解决方案）

- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐49,168  
  集成智能聊天、自主 agent 和 300+ 助手的 AI 生产力工作室，统一接入前沿 LLM。

- **[harry0703/MoneyPrinterTurbo](https://github.com/harry0703/MoneyPrinterTurbo)** ⭐100,660  
  一键生成高清短视频的 AI 工作流，支持主题/关键词驱动，适合内容创作。

- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** ⭐59,615  
  LLM 驱动的多市场股票智能分析系统，集成行情、新闻、决策看板与自动推送。

- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** ⭐42,019  
  AI 将文档或主题转化为原生 PowerPoint 文件，支持动画、图表、语音旁白。

- **[santifer/career-ops](https://github.com/santifer/career-ops)** ⭐62,317  
  开源 AI 求职工具，扫描招聘平台、智能评分简历、定制化投递，本地运行在 CLI agent 中。

### 🧠 大模型/训练（模型权重、训练框架、微调工具）

- **[rasbt/LLMs-from-scratch](https://github.com/rasbt/LLMs-from-scratch)** ⭐100,180  
  从零实现类 ChatGPT 的 LLM，用 PyTorch 逐步构建，是深度学习工程师的经典教程。

- **[skyzh/tiny-llm](https://github.com/skyzh/tiny-llm)** ⭐4,427  
  面向系统工程师的 LLM 推理服务课程，在 Apple Silicon 上构建微型 vLLM + Qwen。

- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7,248  
  支持 100+ 数据集、多种主流模型的 LLM 评测平台，模型选型的必备工具。

- **[AarambhDevHub/aarambh-studio](https://github.com/AarambhDevHub/aarambh-studio)** ⭐51  
  纯 Rust + Candle 实现的 decoder-only LLM，集成 Gated DeltaNet、MoE、量化训练，新晋实验项目。

- **[The-Pocket/PocketFlow](https://github.com/The-Pocket/PocketFlow)** ⭐11,072  
  仅 100 行代码的 LLM 框架，完全透明可定制，适合学习与快速原型。

### 🔍 RAG/知识库（向量数据库、检索增强、知识管理）

- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐86,443  
  领先的 RAG 引擎，融合 Agent 能力，为 LLM 提供高质量上下文层。

- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐62,145  
  为 AI Agent 提供通用记忆层，实现跨会话持久化上下文。

- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐45,433  
  高性能云原生向量数据库，支持大规模 ANN 搜索，是 RAG 系统的存储基石。

- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** ⭐89,081  
  Agent 跨会话记忆工具，自动压缩历史并注入未来会话，兼容多种 CLI agent。

- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** ⭐12,748  
  [MLsys2026] 实现 97% 存储节省的 RAG 系统，可在个人设备上完全私有运行，代表边缘 RAG 趋势。

## 3. 趋势信号分析

- **Agent 技能生态爆发式增长**：`affaan-m/ECC`、`mvanhorn/last30days-skill` 等专注于为 Claude Code、Codex 等 CLI agent 提供可插拔技能（记忆、搜索、安全），反映出社区从“构建 agent 框架”转向“为 agent 提供即用能力”的下游生态成熟。
- **语音交互 Agent 首次登榜**：`huggingface/speech-to-speech` 今日新增 628 星，表明本地化、开源的语音代理成为新方向，配合端侧模型（如 Ollama 支持的模型）有望推动离线语音助手落地。
- **“开源替换 Claude Cowork”成新热点**：`different-ai/openwork` 获 915 星，说明开发者对闭源 AI 助手（如 Claude、GitHub Copilot）的完整开源替代需求强烈，且强调 24/7 后台运行和兼容多种 CLI。
- **RAG 进入“极致压缩”与“私有化”阶段**：`LEANN` 和 `headroomlabs-ai/headroom` 分别从存储压缩和 token 压缩入手，降低 RAG 成本；`mem0` 和 `thedotmack/claude-mem` 推动记忆持久化，RAG 正从检索增强向“主动知识管理”演进。
- **中文项目持续活跃**：`shareAI-lab/learn-claude-code`、`ZhuLinsen/daily_stock_analysis` 等中文社区项目在 agent 和金融场景涌现，反映国内开发者对自主可控 AI 工具的热情。

## 4. 社区关注热点

- **🌟 本地语音 Agent 入门**：关注 `huggingface/speech-to-speech`，它展示了如何用纯开源模型构建本地语音代理，适合智能家居、语音助手等场景。
- **🧩 Agent 技能市场雏形**：`ECC` 和 `last30days-skill` 开创了“为 CLI agent 编写技能”的新模式，类似于插件生态，值得跟踪其标准化进展。
- **⚡ 极致轻量框架 PocketFlow**：`The-Pocket/PocketFlow` 仅 100 行代码实现 LLM 框架，对于想快速构建自定义 agent 的开发者是理想起点。
- **🔐 私有化 RAG 方案 LEANN**：`StarTrail-org/LEANN` 在边缘设备上实现 97% 存储压缩，完全离线运行，适合对数据隐私要求严苛的企业用户。
- **🦀 Rust 在 AI 基础设施中的渗透**：`rig`、`lancedb`、`qdrant` 等项目均用 Rust 开发，表明性能优先的 AI 组件正转向 Rust 生态，值得开发者关注。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*