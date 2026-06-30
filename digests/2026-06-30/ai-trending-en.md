# AI Open Source Trends 2026-06-30

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-30 10:45 UTC

---

# AI Open Source Trends Report — 2026-06-30

## 1. Today’s Highlights

The AI open-source ecosystem today is dominated by **multi-agent orchestration and agent harness systems**, with two projects — `agency-agents` and `council-of-high-intelligence` — exploding past 300+ stars today, signaling intense developer interest in composable, cross-provider agent swarms. A new wave of **domain-specific AI agents** is emerging: `VulnClaw` applies agent workflows to cybersecurity penetration testing, while `ai-berkshire` brings multi-agent reasoning to value investing research. Meanwhile, the **RAG/Knowledge layer continues to mature**, with `graphify` (queryable knowledge graphs from code/document folders) and `mem0` (universal memory for agents) gaining substantial traction. Notably, `browser-use/video-use` extends agent capabilities into **video editing**, representing a new vertical for browser-based automation tools.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, CLI Tools)

| Project | Stars (Total / Today) | Description |
|---|---|---|
| [ollama/ollama](https://github.com/ollama/ollama) | 175,176 / — | Go-based LLM runner now supporting Kimi-K2.6, GLM-5.1, DeepSeek, gpt-oss — local model inference infrastructure. |
| [huggingface/transformers](https://github.com/huggingface/transformers) | 162,049 / — | The de-facto model-definition framework for text, vision, audio, and multimodal models. |
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | 84,889 / — | High-throughput LLM inference engine; essential for production serving. |
| [langchain-ai/langchain](https://github.com/langchain-ai/langchain) | 140,566 / — | The agent engineering platform — remains foundational for building LLM-powered applications. |
| [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) | 141,804 / — | API to search, scrape, and interact with the web at scale for AI agents. |
| [cupy/cupy](https://github.com/cupy/cupy) | — / +352 today | NumPy/SciPy GPU acceleration — critical infrastructure for AI/ML numerical computing. |
| [samchon/nestia](https://github.com/samchon/nestia) | 2,160 / — | NestJS helper + AI chatbot development, bridging TypeScript backend with LLM integration. |

### 🤖 AI Agents / Workflows (Agent Frameworks, Multi-Agent, Automation)

| Project | Stars (Total / Today) | Description |
|---|---|---|
| [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents) | — / +1,425 today | A complete AI agency: specialized agents for frontend, Reddit, discord, research — multi-agent system with defined personas and deliverables. |
| [0xNyk/council-of-high-intelligence](https://github.com/0xNyk/council-of-high-intelligence) | — / +331 today | 18 AI personas (Aristotle, Feynman, Kahneman) deliberate decisions across multiple LLM providers via multi-round structured deliberation. |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | 75,550 / — | Open-source long-horizon SuperAgent harness using sandboxes, memory, tools, sub-agents — tasks from minutes to hours. |
| [OpenHands/OpenHands](https://github.com/OpenHands/OpenHands) | 78,781 / — | AI-driven development agent framework. |
| [langgenius/dify](https://github.com/langgenius/dify) | 147,097 / — | Production-ready platform for agentic workflow development. |
| [CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit) | 35,651 / — | Frontend stack for agents & generative UI — React, Angular, Mobile, Slack. |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | 47,991 / — | AI productivity studio with smart chat, autonomous agents, and 300+ assistants across frontier LLMs. |

### 📦 AI Applications (Specific Vertical Solutions)

| Project | Stars (Total / Today) | Description |
|---|---|---|
| [xbtlin/ai-berkshire](https://github.com/xbtlin/ai-berkshire) | — / +1,386 today | Value investing research framework: multi-agent adversarial analysis using Claude Code / Codex — applies methodologies of Buffett, Munger, Duan Yongping, Li Lu. |
| [browser-use/video-use](https://github.com/browser-use/video-use) | — / +967 today | Edit videos using coding agents — extends browser-use paradigm into video editing. |
| [Unclecheng-li/VulnClaw](https://github.com/Unclecheng-li/VulnClaw) | — / +129 today | AI Agent + MCP toolchain for automated penetration testing: natural language → information gathering → vulnerability discovery → exploitation → report generation. |
| [HKUDS/Vibe-Trading](https://github.com/HKUDS/Vibe-Trading) | — / +839 today | Personal trading agent — uses AI to analyze markets and execute trades. |
| [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) | 52,159 / — | LLM-powered multi-market stock analysis with real-time news, dashboards, and automated notifications. |
| [altic-dev/FluidVoice](https://github.com/altic-dev/FluidVoice) | — / +830 today | Fastest macOS offline dictation app — voice-to-text fully local, Swift-native. |
| [santifer/career-ops](https://github.com/santifer/career-ops) | 56,739 / — | AI-powered job search system built on Claude Code — 14 skill modes, PDF generation, batch processing. |

### 🧠 LLMs / Training (Model Weights, Training Frameworks, Fine-Tuning)

| Project | Stars (Total / Today) | Description |
|---|---|---|
| [hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory) | 72,829 / — | Unified efficient fine-tuning of 100+ LLMs & VLMs (ACL 2024). |
| [jingyaogong/minimind](https://github.com/jingyaogong/minimind) | 52,375 / — | Train a 64M-parameter LLM from scratch in 2 hours — democratizing LLM training education. |
| [open-compass/opencompass](https://github.com/open-compass/opencompass) | 7,137 / — | LLM evaluation platform supporting 100+ datasets and models (Llama, Mistral, GPT-4, Qwen, GLM, etc.). |
| [ScrapeGraphAI/Scrapegraph-ai](https://github.com/ScrapeGraphAI/Scrapegraph-ai) | 27,858 / — | Python scraper based on AI, using LLM models for intelligent web scraping. |
| [acon96/home-llm](https://github.com/acon96/home-llm) | 1,371 / — | Home Assistant integration to control smart home using a local LLM. |

### 🔍 RAG / Knowledge (Vector Databases, Retrieval-Augmented Generation, Knowledge Management)

| Project | Stars (Total / Today) | Description |
|---|---|---|
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | 83,932 / — | Leading open-source RAG engine fusing retrieval-augmented generation with agent capabilities. |
| [run-llama/llama_index](https://github.com/run-llama/llama_index) | 50,532 / — | Document agent and OCR platform for RAG pipelines. |
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | 59,755 / — | Universal memory layer for AI agents — persistent, queryable, cross-session. |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | 74,612 / — | Turn any folder (code, docs, images, videos) into a queryable knowledge graph for AI agents. |
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | 45,025 / — | High-performance, cloud-native vector database for scalable ANN search. |
| [StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN) | 12,615 / — | [MLsys2026] — 97% storage savings in private RAG applications via storage-efficient neural retrieval. |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) | 85,136 / — | Persistent context across sessions for agents — captures, compresses, and injects relevant context using AI. |
| [pathwaycom/llm-app](https://github.com/pathwaycom/llm-app) | 59,179 / — | Ready-to-run cloud templates for RAG with live data sync — Sharepoint, Google Drive, Kafka, PostgreSQL. |

---

## 3. Trend Signal Analysis

**Dominant theme: Agent orchestration is graduating from novelty to production.** Projects like `agency-agents` (+1,425 today) and `council-of-high-intelligence` (+331 today) represent a clear shift: developers no longer want single-purpose agents; they want **swarms of specialized agents** that collaborate, debate, and coordinate across multiple LLM providers. The appeal lies in genuine model diversity — mixing philosophical perspectives (Aristotle, Feynman) or functional roles (frontend wizard, Reddit ninja) to produce higher-quality outputs.

**Vertical specialization is accelerating.** Today’s top gainers are not generic agent frameworks but **domain-specific agent systems**: `VulnClaw` (cybersecurity), `ai-berkshire` (value investing), `Vibe-Trading` (stock trading), `video-use` (video editing), and `santifer/career-ops` (job search). This signals maturity — the community is moving beyond "hello world" agent demos to building production-caliber tools for specific professional workflows.

**New tech stack components entering the mainstream:** The term **"agent harness"** appears repeatedly (shareAI-lab, deer-flow, Grammarly-io, CowAgent), suggesting a new architectural primitive is solidifying — a lightweight runtime that manages agent lifecycle, tools, memory, and model switching. Meanwhile, **"universal memory"** projects (mem0, claude-mem, cognee) treat memory as a separate, pluggable infrastructure layer, not just a feature.

**Connection to industry events:** The proliferation of projects designed for "Claude Code / Codex / OpenClaw / Gemini CLI" (ai-berkshire, claude-mem, grahify, OpenCLI) indicates that **LLM-native CLI tools** have become the standard interface for agent development. This aligns with recent releases from Anthropic and Google making their code agents available in CLI form, creating an ecosystem where a single "agent harness" can work across all major model providers.

**RAG is becoming an infrastructure commodity.** The RAG category shows less explosive growth but deeper integration — `LEANN` (97% storage savings), `pathwaycom/llm-app` (real-time data sync), and `graphify` (multi-modal knowledge graphs) all suggest RAG is moving from experimental to enterprise-ready, with focus on cost efficiency, live data, and multi-format coverage.

---

## 4. Community Hot Spots

- **🦾 Multi-Agent Orchestration Frameworks** — `agency-agents` and `council-of-high-intelligence` are the fastest-growing projects today. The community is converging on the idea that the best AI results come from agent teams with specialized roles, not monolithic models. Expect more projects offering "persona packs" and "agent assembly kits."

- **📈 Financial AI Agents** — `ai-berkshire` (+1,386 today), `Vibe-Trading` (+839), and `ZhuLinsen/daily_stock_analysis` form a cluster of high-interest financial agents. The trend toward AI-powered quantitative research and trading is real — developers are building sophisticated multi-agent frameworks for investment decisions.

- **🔐 AI + Cybersecurity Convergence** — `VulnClaw` shows an entire penetration testing pipeline automated by AI agents (information gathering → vulnerability discovery → exploitation → report generation). This is a high-value vertical where AI agents can dramatically reduce security analyst workload.

- **🎨 Creative Tools go Agent-Native** — `browser-use/video-use` (+967 today) signals that browser-based agent frameworks are expanding into multimedia (video editing). Combined with `FluidVoice` (offline dictation) and `CopilotKit` (generative UI), creative tools are becoming the next frontier for agent automation.

- **🧠 Memory as a Service Layer** — Projects like `mem0` (59,755 stars), `claude-mem` (85,136), and `headroomlabs-ai/headroom` (54,288) treat agent memory as a separate, optimized infrastructure component. The community recognizes that persistent, compressed, and context-aware memory is the key differentiator between toy demos and production agents.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*