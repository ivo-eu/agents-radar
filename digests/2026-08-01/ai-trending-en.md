# AI Open Source Trends 2026-08-01

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-08-01 00:12 UTC

---

# AI Open Source Trends Report — 2026-08-01

**Scope:** Filtered from GitHub Trending and AI topic-search data. Non-AI trending repos (e.g., Chatwoot, Kaneo, tuicr, ESP32-Bit-Pirate, awesome-systematic-trading) were excluded. For trending-only repos, GitHub's snapshot did not include baseline stars, so daily deltas are reported.

## 1. Today's Highlights

Today’s GitHub AI activity is dominated by the **agent-harness / agent-skill layer**. ECC, hermes-agent, AutoGPT and claude-mem continue to hold massive star bases, while new trending repos such as [openwork](https://github.com/different-ai/openwork) (+806), [last30days-skill](https://github.com/mvanhorn/last30days-skill) (+658), and [reverse-skill](https://github.com/zhaoxuya520/reverse-skill) (+335) show that the community is packaging agent capabilities as reusable **skills** for Claude Code, Cursor, Kiro, Cline, and similar CLIs. Microsoft’s [AI-For-Beginners](https://github.com/microsoft/AI-For-Beginners) jumped +1,592 today, underscoring resurgent demand for structured AI education. The day also saw open-source parity with commercial agent products: openwork is explicitly a Claude Cowork alternative powered by opencode, while GitHub’s [Copilot SDK](https://github.com/github/copilot-sdk) (+7) signals platform-level agent integration.

## 2. Top Projects by Category

### 🔧 AI Infrastructure

- [ollama](https://github.com/ollama/ollama) — ⭐177,455 — Local runtime now listing Kimi-K2.6, GLM-5.2, MiniMax, DeepSeek, gpt-oss, Qwen, Gemma; the default self-hosted model gateway.
- [huggingface/transformers](https://github.com/huggingface/transformers) — ⭐163,210 — Model-definition framework for SOTA text/vision/audio models; still the center of model adoption.
- [langchain](https://github.com/langchain-ai/langchain) — ⭐143,114 — The agent engineering platform; standard wiring for tools, RAG, and LLM orchestration.
- [firecrawl](https://github.com/firecrawl/firecrawl) — ⭐158,725 — Web search/scrape/interact API at scale; core data-acquisition infrastructure for agents.
- [pytorch](https://github.com/pytorch/pytorch) — ⭐102,091 — Foundational deep-learning framework for training and research.
- [github/copilot-sdk](https://github.com/github/copilot-sdk) — today +7 — Multi-platform SDK for embedding GitHub Copilot Agent into apps and services.
- [headroomlabs-ai/headroom](https://github.com/headroomlabs-ai/headroom) — ⭐63,572 — Compresses tool outputs, logs, and RAG chunks before they reach the LLM; 20–95% token savings.
- [1jehuang/jcode](https://github.com/1jehuang/jcode) — today +527 — RAM-efficient Rust harness for code/agent execution; lean execution layers are a trending niche.

### 🤖 AI Agents / Workflows

- [ECC](https://github.com/affaan-m/ECC) — ⭐236,638 — Agent harness performance optimization system providing skills, instincts, memory, and security for Claude Code, Codex, Cursor, and more.
- [hermes-agent](https://github.com/NousResearch/hermes-agent) — ⭐223,408 — “The agent that grows with you”; adaptive personal agent framework.
- [AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) — ⭐185,741 — The long-running vision of accessible autonomous agents.
- [dify](https://github.com/langgenius/dify) — ⭐150,930 — Collaborative workspace for building agentic workflows and RAG pipelines.
- [browser-use](https://github.com/browser-use/browser-use) — ⭐107,421 — Makes websites accessible to AI agents for browser automation.
- [claude-mem](https://github.com/thedotmack/claude-mem) — ⭐89,179 — Persistent cross-session context for agents via AI compression and re-injection.
- [openwork](https://github.com/different-ai/openwork) — today +806 — Open-source alternative to Claude Cowork, powered by opencode.
- [last30days-skill](https://github.com/mvanhorn/last30days-skill) — today +658 — Agent skill that researches Reddit, X, YouTube, HN, Polymarket and the web, then synthesizes grounded summaries.

Also notable: [reverse-skill](https://github.com/zhaoxuya520/reverse-skill) (+335 today) brings AI-driven routing to offensive-security and pentest workflows in coding agents.

### 📦 AI Applications

- [open-webui](https://github.com/open-webui/open-webui) — ⭐147,482 — User-friendly AI chat interface supporting Ollama, OpenAI API, and more.
- [MoneyPrinterTurbo](https://github.com/harry0703/MoneyPrinterTurbo) — ⭐100,813 — Generate HD short videos from a keyword via AI model + automated workflow.
- [career-ops](https://github.com/santifer/career-ops) — ⭐62,402 — AI job-search assistant: scans job portals, scores listings, tailors CVs, runs locally in AI coding CLIs.
- [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) — ⭐59,700 — LLM-powered multi-market stock analysis with real-time news and decision dashboards.
- [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) — ⭐49,214 — AI productivity studio with smart chat, autonomous agents, and 300+ assistants.
- [hugohe3/ppt-master](https://github.com/hugohe3/ppt-master) — ⭐42,195 — AI turns documents/topics into native PowerPoint decks with animations, narration, and data-backed charts.
- [deepfakes/faceswap](https://github.com/deepfakes/faceswap) — today +93 — Deepfakes software; still a reference in face synthesis and media manipulation.
- [microsoft/AI-For-Beginners](https://github.com/microsoft/AI-For-Beginners) — today +1,592 — 12-week, 24-lesson AI curriculum; huge education spike today.

### 🧠 LLMs / Training

- [rasbt/LLMs-from-scratch](https://github.com/rasbt/LLMs-from-scratch) — ⭐100,240 — Step-by-step implementation of a ChatGPT-like LLM in PyTorch from scratch.
- [open-compass/opencompass](https://github.com/open-compass/opencompass) — ⭐7,251 — LLM evaluation platform supporting 100+ datasets and major model families.
- [skyzh/tiny-llm](https://github.com/skyzh/tiny-llm) — ⭐4,427 — Learn LLM inference serving on Apple Silicon by building a tiny vLLM + Qwen.
- [genieincodebottle/generative-ai](https://github.com/genieincodebottle/generative-ai) — ⭐2,580 — Comprehensive GenAI roadmap, projects, use cases, and interview-prep resources.
- [thinkwee/AwesomeOPD](https://github.com/thinkwee/AwesomeOPD) — ⭐781 — Curated list of on-policy distillation research.
- [chrisliu298/awesome-llm-unlearning](https://github.com/chrisliu298/awesome-llm-unlearning) — ⭐616 — Resource hub for machine unlearning in LLMs.
- [AarambhDevHub/aarambh-studio](https://github.com/AarambhDevHub/aarambh-studio) — ⭐54 — Decoder-only LLM built from scratch in Rust/Candle: MoE, sparse attention, no Python/PyTorch.
- [ai-glimpse/toyllm](https://github.com/ai-glimpse/toyllm) — ⭐25 — Learning LLM construction from scratch.

### 🔍 RAG / Knowledge

- [Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify) — ⭐99,740 — Turns any codebase, docs, SQL schemas, configs, and PDFs into a queryable knowledge graph; no vector store.
- [infiniflow/ragflow](https://github.com/infiniflow/ragflow) — ⭐86,526 — Leading open-source RAG engine fused with agent capabilities.
- [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) — ⭐64,167 — Local-first agent experience with your own data; “own your intelligence.”
- [run-llama/llama_index](https://github.com/run-llama/llama_index) — ⭐51,262 — Document agent and OCR platform for RAG workflows.
- [milvus-io/milvus](https://github.com/milvus-io/milvus) — ⭐45,440 — Cloud-native vector database for scalable ANN search.
- [VectifyAI/PageIndex](https://github.com/VectifyAI/PageIndex) — ⭐34,940 — “Vectorless,” reasoning-based document index for RAG.
- [qdrant/qdrant](https://github.com/qdrant/qdrant) — ⭐33,697 — High-performance vector database and vector search engine.
- [topoteretes/cognee](https://github.com/topoteretes/cognee) — ⭐29,635 — Open-source AI memory platform for agents using self-hosted knowledge graphs.

## 3. Trend Signal Analysis

The dominant signal is the **consolidation of agent “skills” and memory into a reusable layer**. ECC, hermes-agent, and claude-mem hold huge star bases, while today’s new repositories — last30days-skill, reverse-skill — show developers shipping single-purpose skill packs that can be dropped into Claude Code, Cursor, Kiro, Cline, and other agent CLIs. This suggests the ecosystem is moving from monolithic agent frameworks toward **composable, skill-based behavior packages**.

A second clear trend is **open-source parity with proprietary agent products**. [openwork](https://github.com/different-ai/openwork) (+806) explicitly positions itself as an open Claude Cowork alternative built on opencode. GitHub’s [Copilot SDK](https://github.com/github/copilot-sdk) indicates platform vendors are opening agent integration points. Meanwhile, cost and resource efficiency are rising: [headroom](https://github.com/headroomlabs-ai/headroom) cuts tokens before they hit the LLM, [jcode](https://github.com/1jehuang/jcode) offers a RAM-efficient Rust harness, and [LEANN](https://github.com/StarTrail-org/LEANN) claims 97% storage savings for on-device RAG.

On the model-infrastructure side, [ollama](https://github.com/ollama/ollama) now lists Kimi-K2.6, GLM-5.2, MiniMax, DeepSeek, gpt-oss, Qwen, and Gemma — a signal that local deployment of multiple frontier models is becoming the default. “Vectorless RAG” is also emerging as a counter-movement: [Graphify](https://github.com/Graphify-Labs/graphify), [PageIndex](https://github.com/VectifyAI/PageIndex), and LEANN favor deterministic parsing and knowledge graphs over pure vector search. Finally, security is becoming a mainstream open-source AI theme, with [reverse-skill](https://github.com/zhaoxuya520/reverse-skill) for offensive security and [awesome-MLSecOps](https://github.com/RiccardoBiosas/awesome-MLSecOps) for securing ML systems.

## 4. Community Hot Spots

- **Agent memory and context tooling** — [ECC](https://github.com/affaan-m/ECC), [claude-mem](https://github.com/thedotmack/claude-mem), [mem0](https://github.com/mem0ai/mem0), [headroom](https://github.com/headroomlabs-ai/headroom). Context limits and token costs are the main bottlenecks; memory and compression give immediate ROI.
- **Local “cowork” UIs / agent desktops** — [openwork](https://github.com/different-ai/openwork), [AionUi](https://github.com/iOfficeAI/AionUi), [open-webui](https://github.com/open-webui/open-webui). Demand for private, open-source alternatives to Claude Cowork and hosted agent workspaces is high.
- **Skill packs for coding agents** — [last30days-skill](https://github.com/mvanhorn/last30days-skill), [reverse-skill](https://github.com/zhaoxuya520/reverse-skill), [learn-claude-code](https://github.com/shareAI-lab/learn-claude-code). Skills are becoming the plugin standard for extending agent CLIs.
- **RAG without vector DBs** — [Graphify](https://github.com/Graphify-Labs/graphify), [PageIndex](https://github.com/VectifyAI/PageIndex), [LEANN](https://github.com/StarTrail-org/LEANN). Deterministic knowledge-graph and reasoning-based retrieval offers storage savings and privacy advantages.
- **Vertical domain agents** — [career-ops](https://github.com/santifer/career-ops), [Vibe-Trading](https://github.com/HKUDS/Vibe-Trading), [daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis), [ppt-master](https://github.com/hugohe3/ppt-master). Finance, job search, and office productivity are proving grounds for near-term agent ROI.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*