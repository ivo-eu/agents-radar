# AI Open Source Trends 2026-06-10

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-10 05:08 UTC

---

# AI Open Source Trends Report  
**Date:** 2026-06-10 | **Data Sources:** GitHub Trending (16 repos) + AI Topic Search (81 repos, deduplicated)

---

## 1. Today's Highlights

Today’s GitHub AI ecosystem is dominated by **agent skill marketplaces** and **local-first LLM tooling**. The top trending project, `mvanhorn/last30days-skill`, is an AI agent skill that synthesizes grounded summaries from social media and news – accumulating **+3,191 stars in a single day**. Another standout is `RyanCodrai/turbovec` (+1,801 stars), a Rust-based vector index that demonstrates the community’s hunger for high-performance, embeddable retrieval. The rise of `addyosmani/agent-skills` and `phuryn/pm-skills` signals a shift toward modular, production-grade agent capabilities that can be dropped into any coding or automation workflow. Meanwhile, `Andyyyy64/whichllm` (+633 stars) offers a practical solution for running LLMs locally by benchmarking performance on a user’s own hardware, reflecting the broader trend of **personalized, on-device AI**.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, Dev Tools)

| Project | Stars (Total / Today) | What it is & Why It Matters Today |
|--------|----------------------|-----------------------------------|
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | 82,369 / — | High-throughput inference engine for LLMs; now a standard for serving open-weight models at scale. |
| [ollama/ollama](https://github.com/ollama/ollama) | 173,726 / — | Simplest way to run LLMs locally (Kimi-K2.6, DeepSeek, Qwen, etc.); today’s update includes new model support. |
| [huggingface/transformers](https://github.com/huggingface/transformers) | 161,467 / — | The de facto model-definition framework – essential for training, fine-tuning, and inference across all modalities. |
| [langchain4j/langchain4j](https://github.com/langchain4j/langchain4j) | 12,270 / — | The Java ecosystem’s answer to LangChain, unifying LLM providers and vector stores for enterprise JVM apps. |
| [maziyarpanahi/openmed](https://github.com/maziyarpanahi/openmed) | — / +191 today | Open-source healthcare AI framework; gaining traction as a vertical‑focused infrastructure layer. |
| [openai/plugins](https://github.com/openai/plugins) | — / +284 today | Official OpenAI plugins repository – foundational for extending GPT models with third-party tools. |
| [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow) | 195,625 / — | Classic ML framework; still relevant for production pipelines and AI-powered observability. |
| [pytorch/pytorch](https://github.com/pytorch/pytorch) | 100,626 / — | Dominant deep learning framework; powers the majority of open-source LLM research and training. |

### 🤖 AI Agents / Workflows (Agent Frameworks, Automation, Multi-Agent Systems)

| Project | Stars (Total / Today) | What it is & Why It Matters Today |
|--------|----------------------|-----------------------------------|
| [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) | 189,061 / — | Leading open-source agent harness that “grows with you” – heavily starred, actively developed. |
| [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | 184,861 / — | Vision of accessible AI agents; remains a benchmark for autonomous task execution. |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | 47,143 / — | AI productivity studio with 300+ assistants and autonomous agents – a one‑stop agent workspace. |
| [aaif-goose/goose](https://github.com/aaif-goose/goose) | — / +489 today | Extensible AI agent (Rust) that installs, executes, edits, and tests code with any LLM – trending for its CLI‑first approach. |
| [addyosmani/agent-skills](https://github.com/addyosmani/agent-skills) | — / +443 today | Production‑grade engineering skills for AI coding agents – the “npm” of agent capabilities. |
| [phuryn/pm-skills](https://github.com/phuryn/pm-skills) | — / +806 today | Marketplace of 100+ agentic skills for discovery, strategy, execution – transforming PM workflows. |
| [mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill) | — / +3,191 today | AI agent skill that researches and synthesizes grounded summaries from Reddit, X, YouTube, HN, etc. **Today’s highest‑starred project.** |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | 70,848 / — | Long‑horizon SuperAgent that researches, codes, and creates – combines sandbox, memory, and sub‑agents. |
| [OpenHands/OpenHands](https://github.com/OpenHands/OpenHands) | 76,347 / — | AI‑Driven Development environment – popular for agent‑assisted coding. |
| [shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code) | 65,739 / — | A nano Claude‑code–like agent harness built from scratch – educational and immediately usable. |
| [Panniantong/Agent-Reach](https://github.com/Panniantong/Agent-Reach) | 25,696 / — | Gives AI agents eyes to see the entire internet (Twitter, Reddit, GitHub, Bilibili, etc.) – zero API fees. |

### 📦 AI Applications (Specific Apps, Vertical Solutions)

| Project | Stars (Total / Today) | What it is & Why It Matters Today |
|--------|----------------------|-----------------------------------|
| [santifer/career-ops](https://github.com/santifer/career-ops) | 51,906 / +1,110 today | AI‑powered job search system built on Claude Code – 14 skill modes, PDF generation, batch processing. |
| [yikart/AiToEarn](https://github.com/yikart/AiToEarn) | — / +402 today | Leverages AI for earning – a speculative but fast‑growing vertical (AI + crypto / automation). |
| [hugohe3/ppt-master](https://github.com/hugohe3/ppt-master) | 25,707 / — | AI generates editable PowerPoint from any document – shapes, animations, speaker notes. |
| [maziyarpanahi/openmed](https://github.com/maziyarpanahi/openmed) | — / +191 today | Open‑source healthcare AI – specific domain application gaining community attention. |
| [francescopace/espectre](https://github.com/francescopace/espectre) | — / +134 today | (Filtered out earlier as non-AI, but included here for completeness – motion detection via Wi‑Fi CSI, not ML‑centric.) |

### 🧠 LLMs / Training (Model Weights, Training Frameworks, Fine-Tuning Tools)

| Project | Stars (Total / Today) | What it is & Why It Matters Today |
|--------|----------------------|-----------------------------------|
| [hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory) | 72,038 / — | Unified efficient fine‑tuning of 100+ LLMs and VLMs – a must‑use for custom model training. |
| [open-compass/opencompass](https://github.com/open-compass/opencompass) | 7,075 / — | Comprehensive LLM evaluation platform supporting 100+ datasets – crucial for model selection. |
| [skyzh/tiny-llm](https://github.com/skyzh/tiny-llm) | 4,265 / — | Course‑style project building a tiny vLLM + Qwen on Apple Silicon – educational yet practical. |
| [galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining) | 252 / — | Reliable library for pretraining foundation and world models – emerging direction. |
| [RyanLiu112/Awesome-Process-Reward-Models](https://github.com/RyanLiu112/Awesome-Process-Reward-Models) | 165 / — | Curated list of process reward models – important for RLHF and reasoning enhancement. |
| [llm-jp/awesome-japanese-llm](https://github.com/llm-jp/awesome-japanese-llm) | 1,409 / — | Overview of Japanese LLMs – reflects growing regional language model activity. |

### 🔍 RAG / Knowledge (Vector Databases, Retrieval-Augmented Generation, Knowledge Management)

| Project | Stars (Total / Today) | What it is & Why It Matters Today |
|--------|----------------------|-----------------------------------|
| [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) | 61,340 / — | Local‑first AI agent platform with built‑in vector DB – the “own your intelligence” movement. |
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | 82,342 / — | Leading open‑source RAG engine that fuses RAG with agent capabilities – production‑grade. |
| [run-llama/llama_index](https://github.com/run-llama/llama_index) | 50,052 / — | Document agent and OCR platform – cornerstone of modern RAG pipelines. |
| [NirDiamant/RAG_Techniques](https://github.com/NirDiamant/RAG_Techniques) | 27,806 / — | Notebook tutorials for advanced RAG techniques – educational resource driving adoption. |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) | 81,503 / — | Persistent context across sessions for any agent – compresses and injects relevant history. |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | 58,220 / — | Universal memory layer for AI agents – key for long‑term agent coherence. |
| [RyanCodrai/turbovec](https://github.com/RyanCodrai/turbovec) | — / +1,801 today | Rust vector index built on TurboQuant – high‑performance, embeddable retrieval for Python. |
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | 44,707 / — | Cloud‑native vector database – the standard for large‑scale ANN search. |
| [lancedb/lancedb](https://github.com/lancedb/lancedb) | 10,559 / — | Embeddable retrieval library for multimodal AI – developer‑friendly, search‑first. |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | 64,362 / — | Turn any code, schema, or document into a queryable knowledge graph – bridges code and docs. |
| [topoteretes/cognee](https://github.com/topoteretes/cognee) | 17,750 / — | Open‑source AI memory platform for agents – knowledge graph engine for persistent memory. |
| [VectifyAI/PageIndex](https://github.com/VectifyAI/PageIndex) | 32,820 / — | Document index for vectorless, reasoning‑based RAG – novel approach to retrieval. |
| [zilliztech/claude-context](https://github.com/zilliztech/claude-context) | 11,810 / — | Code search MCP for Claude Code – makes entire codebase context for coding agents. |

---

## 3. Trend Signal Analysis

The most explosive community attention today is concentrated in **agent skill ecosystems** and **local LLM performance benchmarking**. The surge of repos like `mvanhorn/last30days-skill` (+3,191 stars) and `phuryn/pm-skills` (+806 stars) indicates that developers are moving beyond monolithic agents toward **modular, composable skill libraries** that can be mixed and matched across different agent harnesses (Claude Code, Codex, OpenCode, etc.). This is a natural evolution: as coding agents become commoditized, the value shifts to the skills they wield.

A second strong signal is the **local‑first, hardware‑aware AI tooling**. `Andyyyy64/whichllm` (+633 stars) directly addresses a pain point: which LLM actually runs well on *my* machine? Combined with `turbovec` (+1,801 stars) – a high‑performance vector index in Rust with Python bindings – we see a growing demand for **efficient, embeddable infrastructure** that runs on laptops or edge devices, not just cloud GPUs. This aligns with the broader trend toward privacy, cost reduction, and offline capability.

Newly emerging directions include **agent‑specific memory systems** (`mem0ai/mem0`, `thedotmack/claude-mem`) that treat memory as a first‑class service, and **graph‑based knowledge representation** (`safishamsi/graphify`) that replaces flat vector stores with richer relational models. These innovations are directly responding to the limitations of naive RAG (context window, retrieval precision) and are likely to be integrated into next‑generation agent frameworks.

Finally, the explosion of **“system prompts” collections** (`x1xhlol/system-prompts-and-models-of-ai-tools`, +79 today) reflects the community’s desire to reverse‑engineer and clone the capabilities of closed‑source AI tools (Cursor, Devin, Replit). This open‑sourcing of agent internals accelerates the commoditization of coding assistants and democratises access to advanced agent behaviours.

---

## 4. Community Hot Spots

- **Agent Skill Marketplaces (addyosmani/agent-skills, phuryn/pm-skills)**  
  The concept of “skills as building blocks” is exploding. Developers should watch how these marketplaces standardise interfaces (MCP, function calling) and enable rapid composition of custom agents.

- **Local LLM Benchmarking (Andyyyy64/whichllm)**  
  With hardware diversity (Apple Silicon, CUDA, NPUs), tools that auto‑select the best performing LLM for a given device are critical. This project could become a standard pre‑flight check for local AI deployment.

- **Persistent Agent Memory (mem0ai/mem0, thedotmack/claude-mem)**  
  Long‑term memory is the missing piece for production agents. These projects offer drop‑in solutions for session persistence and compression – essential for any agent operating over days or weeks.

- **Vector‑First RAG Infra (turbovec, lancedb, PageIndex)**  
  The shift from cloud‑scale vector DBs to embeddable, high‑performance alternatives (especially in Rust) indicates that the next wave of RAG will be local and real‑time. Developers should experiment with these lighter alternatives.

- **Open‑Sourced Agent Internals (x1xhlol/system-prompts-and-models-of-ai-tools)**  
  The reverse‑engineering of proprietary AI tools is lowering the barrier to building advanced coding agents. This repository is a treasure trove for understanding what makes tools like Cursor and Devin tick.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*