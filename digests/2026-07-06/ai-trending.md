# AI 开源趋势日报 2026-07-06

> 数据来源: GitHub Trending + GitHub Search API | 生成时间: 2026-07-06 13:05 UTC

---

# AI 开源趋势日报｜2026-07-06

## 1. 今日速览
- **智能体技能生态大爆发**：多个新项目聚焦为 Claude Code、Codex 等编码代理提供可复用的技能（Skills），今日新增 stars 包揽榜单前列，标志着 AI 编程进入“技能市集”时代。
- **本地优先的隐私 AI 应用持续走热**：Rust 编写的本地会议助手 Meetily 以 +2493 stars 登顶今日增长王，完全离线运行，强调隐私与效率。
- **系统提示泄露引起广泛好奇**：`system_prompts_leaks` 项目今日 +1386 stars，收集各大模型厂商的隐藏系统提示，成为研究模型行为的独特资源。
- **向量数据库赛道再迎新选手**：阿里巴巴开源的 `zvec`（轻量级进程内向量数据库）进入 Trending，兼顾高性能与低资源消耗，适合边缘场景。
- **RAG 与知识图谱深度融合**：多个项目（如 Graphify、Cognee）将知识图谱引入 RAG，提升长期记忆与推理能力，成为本周社区热议方向。

---

## 2. 各维度热门项目

### 🔧 AI 基础工具（框架、SDK、推理引擎、CLI）

- [ollama/ollama](https://github.com/ollama/ollama) ⭐175k  
  目前最流行的本地大模型运行器，已支持 Kimi、DeepSeek、Qwen、Gemma 等主流模型。今日更新模型列表，覆盖最新发布。
- [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) ⭐145k（今日 +834）  
  网页抓取与交互 API，专为 AI 代理和 RAG 系统设计，支持大规模数据采集。近期新增浏览器互动功能，成为 Agent 数据入口标配。
- [vllm-project/vllm](https://github.com/vllm-project/vllm) ⭐85k  
  LLM 推理引擎标杆，吞吐量高、显存优化好，被大量生产环境采用。今日社区关注其在多模态模型上的扩展。
- [langchain-ai/langchain](https://github.com/langchain-ai/langchain) ⭐141k  
  智能体工程平台，本月发布的 tool calling + MCP 支持大幅降低多工具编排复杂度。
- [steipete/CodexBar](https://github.com/steipete/CodexBar) （今日 +598）  
  macOS 菜单栏小工具，无需登录即可查看 OpenAI Codex 和 Claude Code 的使用统计，对开发者极其实用。

### 🤖 AI 智能体 / 工作流（Agent 框架、自动化、多 Agent）

- [addyosmani/agent-skills](https://github.com/addyosmani/agent-skills) （今日 +1114）  
  面向 AI 编程代理的生产级技能集合，覆盖工程、测试、文档等场景。开发者可直接复用，大幅提升 Coding Agent 效率。
- [alirezarezvani/claude-skills](https://github.com/alirezarezvani/claude-skills) （今日 +611）  
  含 330+ 条 Claude Code 技能、30+ 个代理、70+ 自定义命令，堪称“技能超市”。支持 Claude Code、Codex、Cursor 等 8 个代理。
- [Zackriya-Solutions/meetily](https://github.com/Zackriya-Solutions/meetily) （今日 +2493）  
  完全本地的 AI 会议助手（macOS / Windows），4 倍速 Whisper 实时转录 + 说话人分离 + Ollama 总结。今日增长最高，隐私焦虑下的替代品。
- [openai/codex-plugin-cc](https://github.com/openai/codex-plugin-cc) （今日 +910）  
  OpenAI 官方插件：在 Claude Code 中直接调用 Codex 进行代码审查或任务委托，跨平台协作的示范。
- [mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill) （今日 +237）  
  智能体技能：自动搜索 Reddit、X、YouTube、HN、Polymarket 等平台，生成主题综合摘要。适合舆情追踪与内容策展。
- [gastownhall/gastown](https://github.com/gastownhall/gastown) （今日 +293）  
  多智能体工作空间管理器（Go 语言），可同时运行、调度、通信多个 AI 代理，类似“Agent IDE”。

### 📦 AI 应用（具体产品 / 垂直场景）

- [asgeirtj/system_prompts_leaks](https://github.com/asgeirtj/system_prompts_leaks) （今日 +1386）  
  收集 Anthropic、OpenAI、Google、xAI 等公司最新模型（Claude Fable 5、GPT-5.5、Gemini 3.5 等）的系统提示，是逆向工程和行为分析的重要资料。
- [Leonxlnx/taste-skill](https://github.com/Leonxlnx/taste-skill) （今日 +1453）  
  给 AI 注入“品味”的 JavaScript 技能，阻止模型生成千篇一律的废话。对内容创作者和产品设计者有启发。
- [bradautomates/claude-video](https://github.com/bradautomates/claude-video) （今日 +368）  
  “让 Claude 看懂任何视频”：自动下载视频、提取帧、转录语音，将多模态信息送入 Claude 分析。打开视频理解新范式。
- [karakeep-app/karakeep](https://github.com/karakeep-app/karakeep) （今日 +178）  
  自托管书签管理工具，基于 AI 自动打标签和全文搜索。隐私优先，支持链接、笔记和图片。
- [ruvnet/RuView](https://github.com/ruvnet/RuView) （今日 +471）  
  利用 WiFi 信号实现实时空间感知、生命体征监测和存在检测，无需摄像头。探索“非视觉”AI 感知新路径。

### 🧠 大模型 / 训练（框架、微调、评估）

- [ultralytics/ultralytics](https://github.com/ultralytics/ultralytics) ⭐59k  
  YOLOv8/v11/v26 的统一训练推理框架，社区活跃度极高，持续集成目标检测、实例分割、姿态估计等。
- [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) ⭐210k  
  以“与用户共同成长”为理念的 AI 代理，强调持续学习和自我改进。虽被标记为 Agent，但其底层语言模型迭代是核心。
- [open-compass/opencompass](https://github.com/open-compass/opencompass) ⭐7k  
  多模型评估平台，支持 Llama3、GPT-4、Qwen、Claude 等 100+ 模型的评测。今日社区关注其对长上下文推理的测评结果。
- [AarambhDevHub/aarambh-ai](https://github.com/AarambhDevHub/aarambh-ai) ⭐9  
  纯 Rust 实现的解码器 LLM（25M~1.3B），使用 Candle 框架，支持 INT4/GGUF 量化、LoRA/QLoRA 微调、GRPO 对齐，适合在边缘设备上学习训练流程。

### 🔍 RAG / 知识库（向量数据库、检索增强、知识管理）

- [alibaba/zvec](https://github.com/alibaba/zvec) （今日 +355）⭐13k  
  轻量级进程内向量数据库，C++ 实现，主打闪电速度与无外部依赖，适合嵌入式场景和移动端。
- [infiniflow/ragflow](https://github.com/infiniflow/ragflow) ⭐84k  
  领先的开源 RAG 引擎，融合智能体与知识图谱，为企业级 LLM 提供上下文层。本周新版本支持多模态检索。
- [Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify) ⭐78k  
  将任意代码、文档、数据库结构自动转化为可查询知识图谱，深度绑定 Claude Code、Codex 等编程代理，提升代码理解质量。
- [mem0ai/mem0](https://github.com/mem0ai/mem0) ⭐60k  
  通用 AI 代理记忆层，为对话型 Agent 提供长期记忆。今日关注其对多会话上下文压缩的改进。
- [topoteretes/cognee](https://github.com/topoteretes/cognee) ⭐27k  
  开源 AI 记忆平台，支持知识图谱驱动的持久化记忆，让 Agent 跨会话保持上下文。RAG 与 Agent 记忆的热点结合。
- [StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN) ⭐12k  
  在个人设备上实现 97% 存储压缩的 RAG 应用，兼顾速度、准确性与隐私，适合本地部署。

---

## 3. 趋势信号分析

**（1）智能体技能（Agent Skills）成为今日最热赛道**  
Trending 榜单中至少 5 个项目直接围绕“AI Coding Agent Skills”（如 agent-skills、claude-skills、taste-skill、last30days-skill）。这与 Claude Code、Codex、OpenCode 等代理工具的流行直接相关。开发者不再满足于单纯使用代理，而是开始为其定制可复用的“插件式”能力模块，类似现代 IDE 的插件市场。这一趋势预示着 AI 编程将进入“技能栈”竞争阶段。

**（2）本地优先 + 隐私保护持续崛起**  
Meetily（今日 +2493）、ollama 本地推理、zvec 本地向量数据库等项目均强调“无需云服务”。用户对数据隐私的重视、边缘硬件能力的提升正在推动一股“去中心化 AI”浪潮。注意 Meetily 使用 Rust 编写，性能与安全性并重。

**（3）系统提示泄漏作为研究资源价值凸显**  
`system_prompts_leaks` 在几小时内暴涨，说明开发者对模型“隐藏行为”的好奇心与日俱增。这些泄露的系统提示有助于社区理解模型的防护机制、价值观对齐以及潜规则，也可能是未来 Prompt 安全研究的起点。

**（4）多模态 Agent 能力加速封装**  
`claude-video`、`Agent-Reach`（扫描社交媒体）等项目将非文本输入（视频、网页）转化为 Agent 可理解的信号，标志着 Agent 正在从纯文本交互向多模态感知过渡。这与近期多模态模型的成熟（如 GPT-5.5、Gemini 3.5）形成呼应。

**（5）向量数据库“小而精”方向受青睐**  
与 Milvus、Qdrant 等重量级产品不同，`zvec` 主打进程内嵌入式（无需独立服务），适合 IoT、移动设备、浏览器扩展等场景。这表明向量数据库正在分化出超轻量级分支，与 RAG 本地化趋势匹配。

---

## 4. 社区关注热点（开发者必读）

- **🔥 Agent Skills 生态 —— [addyosmani/agent-skills](https://github.com/addyosmani/agent-skills) 与 [alirezarezvani/claude-skills](https://github.com/alirezarezvani/claude-skills)**  
  理由：两个项目提供了数百条可直接投入生产的 Agent 技能，几乎覆盖所有常见的编码、运维、管理任务。开发者只需一条命令即可安装，是提高 AI 编码效率的捷径。

- **🚀 本地 AI 会议助手 —— [Zackriya-Solutions/meetily](https://github.com/Zackriya-Solutions/meetily)**  
  理由：今日增长王（+2493 stars），完全离线、隐私安全、性能超越云端方案。基于 Rust 的实时转录 + Ollama 总结，展示本地硬件能够承载的高质量 AI 工作流。

- **🕵️ 系统提示资源库 —— [asgeirtj/system_prompts_leaks](https://github.com/asgeirtj/system_prompts_leaks)**  
  理由：唯一的跨厂商系统提示集合，直接暴露各大模型的隐藏配置。对安全研究人员、Prompt 工程师、模型行为观察者是不可多得的一手资料。

- **🧠 轻量级向量数据库 —— [alibaba/zvec](https://github.com/alibaba/zvec)**  
  理由：阿里巴巴开源，进程内运行，无容器开销，极低延迟。适合嵌入到桌面应用或边缘设备中做本地 RAG，与 Meta 的 FAISS 形成对比，值得关注其后续生态。

- **📹 视频理解代理 —— [bradautomates/claude-video](https://github.com/bradautomates/claude-video)**  
  理由：将视频的多模态信息送入 Claude，标志着 Agent 能力从文本到视频的跨越。对视频内容分析、自动化剪辑、教学辅助等场景有直接应用价值。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*