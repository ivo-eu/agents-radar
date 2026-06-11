# AI Open Source Trends 2026-06-11

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-11 03:32 UTC

---

# AI Open Source Trends Report – June 11, 2026

## Today’s Highlights

The open‑source AI landscape continues to be dominated by **agent infrastructure and skill ecosystems**. Today’s trending list is nearly saturated with projects that package reusable “skills” for AI coding agents – from Google’s official `google/skills` to community‑built marketplaces like `phuryn/pm-skills` and the explosive `mvanhorn/last30days-skill` ( +2 535 stars today). At the same time, large‑scale training and fine‑tuning efforts remain strong: `FareedKhan-dev/train-llm-from-scratch` and `galilai-group/stable-pretraining` signal continued appetite for in‑house LLM development. Computer vision also sees a major update with `roboflow/supervision` (+695 today), while healthcare AI (`openmed`) and edge inference (`picollm`) demonstrate expanding vertical applications. The “agent skills” pattern – where individual capabilities are isolated, cataloged, and composed – appears to be the defining architectural trend of the moment.

## Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, Dev Tools, CLI)

- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐82 476 (total)  
  High‑throughput, memory‑efficient inference and serving engine for LLMs – the backbone for many production deployments.

- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐173 809  
  Run local LLMs (Kimi‑K2.6, GLM‑5.1, DeepSeek, etc.) with a single command; now the de facto standard for private model execution.

- **[langchain-ai/langchain](https://github.com/langchain-ai/langchain)** ⭐138 997  
  The agent engineering platform – still the most widely used framework for building LLM‑powered chains and agents.

- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐161 487  
  Universal model‑definition library for text, vision, audio, and multimodal models, critical for both research and production.

- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐47 185  
  AI productivity studio with smart chat, autonomous agents, and 300+ built‑in assistants, unifying access to frontier LLMs.

- **[0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig)** ⭐7 583  
  Build modular, scalable LLM applications in Rust – a growing ecosystem for performance‑sensitive AI workloads.

- **[Picovoice/picollm](https://github.com/Picovoice/picollm)** ⭐312  
  On‑device LLM inference powered by X‑bit quantization, enabling edge deployment of large models.

### 🤖 AI Agents / Workflows (Agent Frameworks, Automation, Multi‑Agent Systems)

- **[addyosmani/agent-skills](https://github.com/addyosmani/agent-skills)** ⭐821 (today)  
  Production‑grade engineering skills for AI coding agents – the leading example of reusable agent capabilities.

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** ⭐2 535 (today)  
  AI agent skill that researches any topic across Reddit, X, YouTube, HN, and more, then synthesizes a grounded summary – the highest‑starred trending repo today.

- **[google/skills](https://github.com/google/skills)** ⭐211 (today)  
  Official Agent Skills for Google products and technologies, bringing first‑party support to the agent skill ecosystem.

- **[obra/superpowers](https://github.com/obra/superpowers)** ⭐1 104 (today)  
  An agentic skills framework and software development methodology that works – practical patterns for composing agent behaviors.

- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐190 125  
  “The agent that grows with you” – a highly starred, extensible agent harness for personal AI assistants.

- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐184 888  
  The original autonomous agent framework, continuously updated and still the most popular end‑to‑end agent project.

- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** ⭐70 923  
  Open‑source long‑horizon SuperAgent harness that researches, codes, and creates – handles multi‑hour tasks with sandboxes, memories, and subagents.

### 📦 AI Applications (Specific Apps, Vertical Solutions)

- **[harry0703/MoneyPrinterTurbo](https://github.com/harry0703/MoneyPrinterTurbo)** ⭐1 389 (today)  
  One‑click AI‑generated short videos from text – a viral content‑creation tool powered by LLMs.

- **[roboflow/supervision](https://github.com/roboflow/supervision)** ⭐695 (today)  
  Reusable computer vision tools, making it easy to build detection, segmentation, and tracking pipelines.

- **[maziyarpanahi/openmed](https://github.com/maziyarpanahi/openmed)** ⭐527  
  Open‑source healthcare AI – a platform for clinical‑grade language models and medical data processing.

- **[ruvnet/RuView](https://github.com/ruvnet/RuView)** ⭐420 (today)  
  Turns commodity WiFi signals into real‑time spatial intelligence and vital sign monitoring – novel AI application in IoT.

- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** ⭐26 208  
  AI generates real, editable PowerPoint presentations from any document, preserving native shapes and animations.

- **[zhayujie/CowAgent](https://github.com/zhayujie/CowAgent)** ⭐45 210  
  Open‑source super AI assistant and agent harness with planning, tools, memory, and multi‑model support – a full user‑facing application.

### 🧠 LLMs / Training (Model Weights, Training Frameworks, Fine‑Tuning Tools)

- **[FareedKhan-dev/train-llm-from-scratch](https://github.com/FareedKhan-dev/train-llm-from-scratch)** ⭐247 (today)  
  Straightforward method for training your own LLM from data to generation – ideal for hands‑on learning.

- **[hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)** ⭐72 056  
  Unified efficient fine‑tuning of 100+ LLMs and VLMs – the go‑to toolkit for adapting pretrained models.

- **[galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining)** ⭐254  
  Reliable, minimal, scalable library for pretraining foundation and world models – a new contender in training infrastructure.

- **[skyzh/tiny-llm](https://github.com/skyzh/tiny-llm)** ⭐4 267  
  Course on building a tiny vLLM + Qwen for Apple Silicon – educational resource for LLM inference serving.

- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** ⭐7 080  
  Comprehensive LLM evaluation platform supporting 100+ datasets and models – essential for benchmarking.

- **[RyanLiu112/Awesome-Process-Reward-Models](https://github.com/RyanLiu112/Awesome-Process-Reward-Models)** ⭐167  
  Curated collection of process reward models – a rapidly evolving area for RL‑based LLM alignment.

### 🔍 RAG / Knowledge (Vector Databases, Retrieval‑Augmented Generation, Knowledge Management)

- **[langgenius/dify](https://github.com/langgenius/dify)** ⭐144 771  
  Production‑ready platform for agentic workflow development with built‑in RAG capabilities.

- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐82 423  
  Leading open‑source RAG engine combining cutting‑edge retrieval with agent capabilities – a top choice for context‑augmented LLMs.

- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐44 724  
  High‑performance, cloud‑native vector database for scalable ANN search – the backbone of many RAG pipelines.

- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐58 293  
  Universal memory layer for AI agents – enables persistent, queryable context across sessions.

- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** ⭐81 668  
  Captures everything an agent does, compresses with AI, and injects relevant context into future sessions – a “memory” skill for coding agents.

- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** ⭐32 016  
  High‑performance vector database and search engine, now with a dedicated cloud offering.

- **[lancedb/lancedb](https://github.com/lancedb/lancedb)** ⭐10 570  
  Developer‑friendly embedded retrieval library for multimodal AI – lightweight alternative for local RAG.

- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** ⭐11 903  
  “RAG on Everything” with 97% storage savings – a novel approach to private, on‑device RAG.

## Trend Signal Analysis

The **most explosive community attention** today is on **agent skills** – small, composable capabilities that can be attached to coding agents like Claude Code, Codex, or Copilot. Projects like `addyosmani/agent-skills`, `phuryn/pm-skills`, and `mvanhorn/last30days-skill` are gaining thousands of stars per day, and Google’s official `google/skills` signals industry‑level endorsement. This pattern represents a shift from monolithic agent frameworks toward **skill marketplaces** where developers can package domain‑specific expertise (project management, research, code review) and share them as reusable modules. The “last30days” skill, which synthesizes multi‑platform research in a single call, exemplifies the user‑facing value of this approach.

A new direction appearing for the first time is **agent memory and context injection** as a standalone product category. `claude-mem` (81k stars) and `zilliztech/claude-context` (11k stars) treat context as a persistent, searchable layer – a concept that is rapidly becoming essential for long‑running agent tasks. This connects directly to the recent proliferation of coding agents (Claude Code, Codex, Gemini CLI) that need to maintain state across sessions. The skill ecosystem and the memory layer are two sides of the same coin: both aim to make agents **stateful and extensible** without rewriting the core harness.

On the training side, `stable-pretraining` and `train-llm-from-scratch` show continued interest in “from first principles” LLM development, likely driven by the availability of powerful open‑source models (Kimi, Qwen, DeepSeek) and the desire for greater control. The **vector‑database space remains hot** but mature, with established players like Milvus, Qdrant, and LanceDB competing on performance, while new entrants like `LEANN` push the frontier of storage efficiency.

Finally, the **computer vision** community is still highly active: `roboflow/supervision` and `ultralytics` projects indicate that CV pipelines are becoming as modular and skill‑driven as LLM workflows. The rise of “agent skills” may soon extend to vision – allowing agents to pick and choose CV tools on demand.

## Community Hot Spots

- **[addyosmani/agent-skills](https://github.com/addyosmani/agent-skills)** – A must‑watch for anyone building or using AI coding agents. The collection of production‑grade skills sets a standard for quality and reusability.

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** – The highest‑starred trending repo today; its approach of multi‑platform research synthesis is a practical killer app for agent skills.

- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** – With 81k stars, this project is redefining how agents handle context. Developers working on long‑running agents should experiment with its persistent memory layer.

- **[roboflow/supervision](https://github.com/roboflow/supervision)** – Continues to be the most productive way to build computer vision applications. Recent updates align with the “skill” trend, making CV reusable across agents.

- **[FareedKhan-dev/train-llm-from-scratch](https://github.com/FareedKhan-dev/train-llm-from-scratch)** – A hands‑on resource for understanding the full training pipeline. As foundation models proliferate, this skill is increasingly valuable for builders who want to customize their own LLMs.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*