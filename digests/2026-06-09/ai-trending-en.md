# AI Open Source Trends 2026-06-09

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-09 04:18 UTC

---

# AI Open Source Trends Report — 2026-06-09

## 1. Today's Highlights

Today's GitHub trending data reveals a decisive shift toward **agent skill marketplaces and personal AI infrastructure**. The most explosive growth belongs to `mvanhorn/last30days-skill` (+3,558 stars), a single skill that synthesizes research across Reddit, X, YouTube, and Polmarket — signaling that the ecosystem is maturing from building agent *frameworks* to trading agent *skills* as composable modules. Meanwhile, `roboflow/supervision` (+1,288 stars) reminds us that computer vision tooling remains a bedrock need, and `aaif-goose/goose` (+699 stars) continues to popularize the "extensible agent harness" pattern pioneered by Claude Code. Notably, Google entered the agent-skill space with `google/skills`, and OpenAI released its `plugins` repo — both suggesting that major players are standardizing on skill-based architectures. On the infrastructure side, `turbovec` (+1,729 stars) introduces a Rust-based vector index with Python bindings, targeting performance-critical retrieval, while `MemPalace` (+170 stars) claims best-in-class open-source AI memory — a category seeing explosive interest.

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, Dev Tools)

- **[turbovec](https://github.com/RyanCodrai/turbovec)** ⭐0 (+1,729 today) — A Rust vector index with Python bindings using TurboQuant quantization; positions itself as a high-speed alternative to FAISS for production retrieval.
- **[roboflow/supervision](https://github.com/roboflow/supervision)** ⭐0 (+1,288 today) — Reusable computer vision tools (annotation, tracking, visualization) that abstract away boilerplate, now a critical dependency for CV pipelines.
- **[aaif-goose/goose](https://github.com/aaif-goose/goose)** [Rust] ⭐0 (+699 today) — An extensible AI agent harness that goes beyond code suggestions: installs, executes, edits, and tests with any LLM.
- **[MetaAI/whichllm](https://github.com/Andyyyy64/whichllm)** [Python] ⭐0 (+143 today) — One-command tool that benchmarks local LLMs on your hardware and ranks them by recency-aware performance, not parameter count.
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** ⭐76,279 — The leading open-source AI-driven development platform; continues to dominate the agent-IDE space.
- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐173,636 — The de facto local LLM runner; now supports Kimi-K2.6, GLM-5.1, and other frontier models out of the box.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐82,270 — High-throughput LLM inference engine; remains the backbone for production serving of open models.

### 🤖 AI Agents / Workflows (Agent Frameworks, Automation, Multi-Agent)

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** [Python] ⭐0 (+3,558 today) — An AI agent skill that researches any topic across Reddit, X, YouTube, HN, and Polymarket, then produces a grounded summary. Explosive growth suggests skill-marketplace demand is real.
- **[Panniantong/Agent-Reach](https://github.com/Panniantong/Agent-Reach)** [Python] ⭐0 (+679 today) — Gives agents "eyes" to read and search Twitter, Reddit, YouTube, GitHub, Bilibili, XiaoHongShu via CLI — zero API fees.
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** [TypeScript] ⭐0 (+378 today) — Frontend stack for agents and generative UI, supporting React, Angular, Mobile, Slack. Makers of the AG-UI Protocol.
- **[santifer/career-ops](https://github.com/santifer/career-ops)** [JavaScript] ⭐0 (+308 today) — AI-powered job search system built on Claude Code with 14 skill modes, Go dashboard, and batch PDF generation.
- **[phuryn/pm-skills](https://github.com/phuryn/pm-skills)** ⭐0 (+164 today) — A marketplace of 100+ agentic skills for product management (discovery, strategy, execution, launch, growth).
- **[google/skills](https://github.com/google/skills)** [Python] ⭐0 (+461 today) — Official Google repository of agent skills for Google products and technologies.
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐184,850 — The original autonomous agent project; still the largest community in the agent space.
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐187,671 — An agent framework designed to grow with the user, emphasizing long-term memory and self-evolution.

### 📦 AI Applications (Specific Apps, Vertical Solutions)

- **[santifer/career-ops](https://github.com/santifer/career-ops)** [JavaScript] ⭐0 (+308 today) — Full-stack job search automation with LLM-driven resume tailoring and application tracking.
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐47,086 — AI productivity studio unifying smart chat, autonomous agents, and 300+ assistants with unified LLM access.
- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** ⭐25,342 — Generates editable PowerPoint decks from any document using AI, with native shapes, animations, and audio narration.
- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** ⭐41,422 — LLM-powered stock analysis for A/H/US markets with multi-source data, decision dashboards, and zero-cost deployment.

### 🧠 LLMs / Training (Model Weights, Training Frameworks, Fine-Tuning)

- **[huggingface/transformers](https://github.com/huggingface/transformers)** ⭐161,424 — The universal model-definition framework for text, vision, audio, and multimodal models.
- **[hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)** ⭐72,006 — Unified fine-tuning for 100+ LLMs and VLMs (ACL 2024 recipient); remains the go-to for custom model training.
- **[galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining)** ⭐251 — Minimal, scalable library for pretraining foundation and world models; a newcomer worth watching for reproducibility.
- **[RyanLiu112/Awesome-Process-Reward-Models](https://github.com/RyanLiu112/Awesome-Process-Reward-Models)** ⭐163 — Curated collection of process reward model papers and resources, signaling growing interest in step-level supervision.

### 🔍 RAG / Knowledge (Vector Databases, Retrieval-Augmented Generation, Knowledge Management)

- **[MemPalace/mempalace](https://github.com/MemPalace/mempalace)** [Python] ⭐0 (+170 today) — Claims best-benchmarked open-source AI memory system; free and designed for persistent agent context.
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐82,244 — Leading open-source RAG engine fusing retrieval with agent capabilities for enterprise context layers.
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐44,689 — Cloud-native vector database for scalable ANN search; cornerstone of production RAG stacks.
- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** ⭐31,940 — High-performance vector database with Rust core; growing adoption in latency-sensitive applications.
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐58,097 — Universal memory layer for AI agents; persistent, cross-session context management.
- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** ⭐81,337 — Captures agent session context, compresses it, and injects relevant memory into future sessions — works with Claude Code, OpenClaw, and more.

## 3. Trend Signal Analysis

Today's data reveals **three accelerating movements** in the AI open-source ecosystem:

**First, the rise of the "Skill Marketplace" as a product category.** `mvanhorn/last30days-skill` (+3,558 stars), `phuryn/pm-skills` (+164 stars), and `google/skills` (+461 stars) all share a common premise: rather than building monolithic agents, the community is now packaging individual, composable *skills* — modular capabilities that agents can discover, download, and chain. This mirrors the early plugin economy of WordPress and the MCP (Model Context Protocol) ecosystem, but at a higher level of abstraction. The simultaneous release of OpenAI's `plugins` repo and Google's `skills` repo suggests that the industry's largest players are converging on this architecture.

**Second, "AI Memory" is emerging as a distinct infrastructure layer.** `MemPalace`, `mem0ai/mem0` (58k stars), `thedotmack/claude-mem` (81k stars), and `topoteretes/cognee` (17.7k stars) all tackle the same problem: agents lack persistent, compressing, cross-session memory. The volume of stars across these projects indicates that short-term context windows are no longer sufficient — developers want agents that remember user preferences, past tasks, and evolving knowledge. This is a direct response to the growing depth of agent workflows (hours-long tasks, multi-step research).

**Third, the Claude Code ecosystem is becoming a meta-platform.** `santifer/career-ops`, `phuryn/pm-skills`, `CopilotKit`, `thedotmack/claude-mem`, `aaif-goose/goose` — all either build on, extend, or interoperate with Claude Code's agent harness pattern. This suggests Anthropic's release strategy (an open, extensible CLI agent) has successfully created a *platform effect* where third-party developers build value on top of a shared runtime.

## 4. Community Hot Spots

- **Agent Skill Marketplaces** — Projects like `last30days-skill` and `pm-skills` are defining how agents discover and use capabilities. Developers should watch for standardization efforts (e.g., shared skill registries, skill-to-agent compatibility layers).

- **Local Hardware Benchmarking** — `whichllm` (+143 stars today) solves a real pain point: "which model runs best *on my machine*?" As local AI usage grows, benchmarking tools will become essential infrastructure.

- **Memory-as-a-Service for Agents** — `MemPalace`, `mem0`, `claude-mem`, and `cognee` all deliver persistent memory. The lack of a clear winner means standardization opportunities exist — and integration with vector databases (Milvus, Qdrant) is the next frontier.

- **Rust in AI Infrastructure** — `turbovec` (Rust + Python bindings), `goose` (Rust agent harness), `qdrant` (Rust vector DB), and `rig` (Rust LLM framework) show Rust gaining traction as the performance layer for AI tooling. Python remains the interface, but Rust is the engine.

- **Vision + Agent Convergence** — `roboflow/supervision` (+1,288 today) and `PaddleOCR` (81.5k stars) highlight that computer vision is not separate from the agent revolution — agents need to "see" the web (Agent-Reach), documents (OCR), and video feeds (Supervision). The integration of vision tools into agent workflows is a major growth vector.