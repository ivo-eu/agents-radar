# AI Open Source Trends 2026-06-14

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-14 11:10 UTC

---

# AI Open Source Trends Report — 2026-06-14

## 1. Today's Highlights

**Security-first AI infrastructure is surging.** NVIDIA's **SkillSpector** (⭐804 today) — a security scanner for AI agent skills — exploded on the trending chart, reflecting urgent community demand for safety tooling as agent adoption accelerates. **Financial AI is a breakout vertical**, with two major projects: **shiyu-coder/Kronos** (⭐238 today, a foundation model for financial markets) and **TauricResearch/TradingAgents** (⭐85,996 total, multi-agent trading framework) signaling a gold rush in LLM-powered quantitative finance. **Unified AI API layers are consolidating**, led by **andrewyng/aisuite** (⭐127 today) and supported by growing interest in "one SDK to rule them all" patterns. Meanwhile, the **RAG and vector database ecosystem continues its explosive scale**, with **mem0ai/mem0** (⭐58,521) and **thedotmack/claude-mem** (⭐82,191) pushing persistent memory layers for agents into the mainstream.

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, CLI Tools)

- [NVIDIA/SkillSpector](https://github.com/NVIDIA/SkillSpector) — ⭐804 today · Python
  Security scanner for AI agent skills; detects vulnerabilities and malicious patterns in agent code.
- [andrewyng/aisuite](https://github.com/andrewyng/aisuite) — ⭐127 today · Python
  Unified interface to multiple Generative AI providers; reduces provider-switching friction for developers.
- [vllm-project/vllm](https://github.com/vllm-project/vllm) — ⭐82,818 total · Python
  High-throughput, memory-efficient LLM inference and serving engine; the de facto standard for production deployment.
- [0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig) — ⭐7,614 total · Rust
  Modular, scalable LLM application framework in Rust; gaining traction for performance-critical agent pipelines.
- [ollama/ollama](https://github.com/ollama/ollama) — ⭐174,116 total · Go
  Local LLM runtime now supporting Kimi-K2.6, GLM-5.1, MiniMax, DeepSeek, and more; the entry point for self-hosted AI.

### 🤖 AI Agents / Workflows (Agent Frameworks, Automation, Multi-Agent Systems)

- [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) — ⭐184,932 total · Python
  The original autonomous agent vision; continues to evolve with new skill systems and security layers.
- [langchain-ai/langchain](https://github.com/langchain-ai/langchain) — ⭐139,239 total · Python
  The agent engineering platform; now the standard orchestration layer for production agent workflows.
- [OpenHands/OpenHands](https://github.com/OpenHands/OpenHands) — ⭐76,984 total · Python
  AI-driven development agent; competes with Cursor/Claude Code as an open-source coding assistant.
- [Eigenwise/atomic-agents](https://github.com/Eigenwise/atomic-agents) — ⭐5,979 total · Python
  "Building AI agents, atomically" — a composable, modular approach to agent construction.
- [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) — ⭐193,091 total · Python
  The agent that grows with you; integrates memory, skills, and self-evolution.

### 📦 AI Applications (Specific Apps, Vertical Solutions)

- [shiyu-coder/Kronos](https://github.com/shiyu-coder/Kronos) — ⭐238 today · Python
  Foundation model for the language of financial markets; a specialized vertical LLM for trading signals.
- [TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents) — ⭐85,996 total · Python
  Multi-agent LLM financial trading framework; the most-starred finance-specific AI project.
- [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) — ⭐61,556 total · JavaScript
  Local-first agent experience; everything needed for running private AI agents on personal hardware.
- [acon96/home-llm](https://github.com/acon96/home-llm) — ⭐1,358 total · Python
  Home Assistant integration for local LLM control of smart home devices; practical edge AI.
- [music-assistant/server](https://github.com/music-assistant/server) — ⭐270 today · Python
  AI-powered media library manager; connects streaming services to open-source hardware.

### 🧠 LLMs / Training (Model Weights, Training Frameworks, Fine-Tuning)

- [huggingface/transformers](https://github.com/huggingface/transformers) — ⭐161,576 total · Python
  The foundational model-definition framework; still the hub for all state-of-the-art models.
- [open-compass/opencompass](https://github.com/open-compass/opencompass) — ⭐7,082 total · Python
  LLM evaluation platform supporting 100+ datasets across Llama3, Mistral, GPT-4, Qwen, and more.
- [thinkwee/AwesomeOPD](https://github.com/thinkwee/AwesomeOPD) — ⭐625 total
  Curated list for On-Policy Distillation; a niche but growing area in LLM optimization.
- [chrisliu298/awesome-llm-unlearning](https://github.com/chrisliu298/awesome-llm-unlearning) — ⭐598 total
  Resource repository for machine unlearning in LLMs; critical for compliance and safety research.

### 🔍 RAG / Knowledge (Vector Databases, Retrieval-Augmented Generation, Knowledge Management)

- [infiniflow/ragflow](https://github.com/infiniflow/ragflow) — ⭐82,682 total · Python
  Leading open-source RAG engine fusing retrieval with agent capabilities for LLM context layering.
- [mem0ai/mem0](https://github.com/mem0ai/mem0) — ⭐58,521 total · Python
  Universal memory layer for AI agents; persistent context across sessions.
- [run-llama/llama_index](https://github.com/run-llama/llama_index) — ⭐50,116 total · Python
  Document agent and OCR platform; the leading framework for structured data extraction and indexing.
- [milvus-io/milvus](https://github.com/milvus-io/milvus) — ⭐44,772 total · Go
  Cloud-native vector database for scalable ANN search; enterprise-grade retrieval infrastructure.
- [NirDiamant/RAG_Techniques](https://github.com/NirDiamant/RAG_Techniques) — ⭐27,928 total · Jupyter Notebook
  Comprehensive tutorial collection of advanced RAG techniques; the go-to learning resource.
- [StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN) — ⭐11,918 total · Python
  RAG on everything with 97% storage savings; a breakthrough in efficient private RAG.
- [topoteretes/cognee](https://github.com/topoteretes/cognee) — ⭐17,819 total · Python
  Open-source AI memory platform for agents with knowledge graph engine.

## 3. Trend Signal Analysis

**🌪️ Explosive Attention: AI Agent Security.**

NVIDIA/SkillSpector's 804-stars-in-one-day debut signals a **paradigm shift** in the agent ecosystem. As autonomous agents proliferate (AutoGPT at 184K, OpenHands at 76K, langchain at 139K), the security layer — previously an afterthought — has become the #1 bottleneck. SkillSpector is the first major open-source tool specifically targeting **agent skill vulnerabilities**, and its immediate traction suggests that every agent framework will need integrated security scanning within 6 months.

**📈 New Tech Stacks Emerging: Financial AI as a First-Class Citizen.**

Two concurrent developments — **Kronos** (a foundation model for financial markets) and **TradingAgents** (multi-agent trading framework) — point to **vertical foundation models** becoming a real trend. Unlike generic LLMs fine-tuned for finance, Kronos is purpose-built from scratch for the "language" of markets (price sequences, order books, news). This mirrors the transition from GPT-3 → Codex for coding, but now applied to finance. Expect similar "domain-native" foundation models for healthcare, legal, and scientific research.

**🔄 Connection to Recent Events: The "Unified API" War Heats Up.**

The simultaneous growth of **aisuite** (andrewyng), **rig** (Rust), **langchain4j** (Java), and **open-webui** suggests the ecosystem is consolidating around **a single abstraction layer**. This mirrors the "Kubernetes of AI" thesis — developers don't want to learn 15 provider SDKs. The battle is now between Python-centric (aisuite), Rust-performant (rig), and JVM-enterprise (langchain4j) approaches. The winner will define how all future AI applications are built.

**📊 Multi-Agent Architectures Go Mainstream.**

Projects like **TradingAgents**, **Cognee**, **atomic-agents**, and **mem0** all share a common architectural pattern: agents that **remember, collaborate, and specialize**. The "single agent" paradigm (AutoGPT v1) is giving way to **swarm architectures** where agents have distinct roles (trader, analyst, risk manager). This is reflected in the 58K stars for mem0 (memory layer) and the explosive growth of MCP (Model Context Protocol) compatible tools.

**💾 Loca-first + RAG = The New Default.**

The combination of **anything-llm** (61K, local-first), **ollama** (174K), **LEANN** (11K, 97% storage savings), and **RAG_Techniques** (27K) reinforces a clear direction: **AI is moving to the edge**. Users want private, efficient, local RAG without cloud dependency. StarTrail-org/LEANN's 97% storage compression is a breakthrough that makes on-device RAG practical for phones and laptops.

## 4. Community Hot Spots

- 🔥 **Agent Security Scanning** (NVIDIA/SkillSpector, ⭐804/day) — Every agent developer should watch this space. Security is the unaddressed gap in the agent boom, and SkillSpector is filling it fast.

- 📈 **Financial AI Foundations** (shiyu-coder/Kronos, TauricResearch/TradingAgents) — The combination of a purpose-built financial foundation model + multi-agent trading framework creates a complete stack. Watch for derivative projects in other verticals.

- 🧩 **Unified API Abstractions** (andrewyng/aisuite, rig, langchain4j) — The "one SDK" movement is real. Developers building production apps should evaluate **aisuite** for Python, **rig** for Rust performance, and **langchain4j** for Java enterprise stacks.

- 🧠 **Persistent Agent Memory** (mem0ai/mem0, thedotmack/claude-mem, topoteretes/cognee) — Memory is the missing piece for agents that work across sessions. The 58K+ stars on mem0 signal that this is now table stakes for any serious agent framework.

- 🪨 **Prompt Optimization & Token Reduction** (JuliusBrussee/caveman, ⭐72K) — "Why use many token when few token do trick" — caveman's 65% token reduction with minimal quality loss is a wild, pragmatic hack that could save millions in inference costs. Watch for similar "token frugality" techniques entering production.

---

*Report generated: 2026-06-14 | Data sources: GitHub Trending + AI Topic Search*

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*