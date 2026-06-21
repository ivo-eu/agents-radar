# AI Open Source Trends 2026-06-21

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-21 11:26 UTC

---

# AI Open Source Trends Report — 2026-06-21

## 1. Today's Highlights

Today's GitHub trending reveals a clear surge in **agent memory and efficiency tooling**, with `headroom` (+3,795 stars) and `codebase-memory-mcp` (+1,271 stars) dominating the charts. The ecosystem is shifting from "can we build agents?" to "how do we make them cheaper, faster, and more persistent?" — evidenced by the explosive adoption of token compression and knowledge graph-based memory systems. Meanwhile, `deer-flow` from ByteDance (72,213 total stars) confirms that long-horizon agentic workflows are entering mainstream production. A notable surprise is the emergence of system prompt leak collections (`system_prompts_leaks`) and skills-sharing repositories (`mattpocock/skills`), suggesting the community is actively reverse-engineering and standardizing prompt engineering patterns.

## 2. Top Projects by Category

### 🔧 AI Infrastructure

- **[headroom](https://github.com/chopratejas/headroom)** — ⭐0 (+3,795 today)  
  Token compression library that reduces LLM input by 60-95% with no loss in answer quality. The #1 trending AI repo today.

- **[codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp)** — ⭐0 (+1,271 today)  
  High-performance MCP server indexing codebases into a persistent knowledge graph in milliseconds. 158 languages supported, sub-ms queries.

- **[tursodatabase/turso](https://github.com/tursodatabase/turso)** — ⭐0 (+801 today)  
  In-process SQLite-compatible database gaining traction for AI agent state management and local-first applications.

- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** — ⭐83,457 (topic:llm)  
  High-throughput LLM inference engine, still the backbone for self-hosted model serving.

- **[ollama/ollama](https://github.com/ollama/ollama)** — ⭐174,631 (topic:llm)  
  Local LLM runner now supporting Kimi-K2.6, GLM-5.1, and other recent models — the default choice for local AI development.

### 🤖 AI Agents / Workflows

- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** — ⭐72,213 (+415 today, topic:llm)  
  Long-horizon SuperAgent that researches, codes, and creates using sandboxes, memory, and sub-agents. ByteDance's production-grade agent framework.

- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** — ⭐77,884 (topic:llm)  
  AI-driven development agent, a staple in the AI coding assistant space.

- **[langgenius/dify](https://github.com/langgenius/dify)** — ⭐146,019 (topic:llm)  
  Production-ready agentic workflow development platform, now the standard for visual agent pipelines.

- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** — ⭐198,618 (topic:llm)  
  Self-growing agent framework — the agent that evolves with user interactions.

- **[mattpocock/skills](https://github.com/mattpocock/skills)** — ⭐0 (+1,395 today)  
  Skills for real engineers, straight from the author's `.claude` directory — catalyzing the "shareable agent skills" movement.

- **[shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code)** — ⭐67,630 (topic:ai-agent)  
  Nano agent harness built from scratch — a minimalist educational take on Claude Code-like functionality.

- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** — ⭐47,605 (topic:ai-agent)  
  AI productivity studio with 300+ assistants and autonomous agents, aggregating frontier LLMs.

### 📦 AI Applications

- **[calesthio/OpenMontage](https://github.com/calesthio/OpenMontage)** — ⭐0 (+677 today)  
  World's first open-source agentic video production system with 12 pipelines, 52 tools, and 500+ agent skills.

- **[palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro)** — ⭐0 (+902 today)  
  macOS video editor built for AI-native workflows.

- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** — ⭐0 (+519 today, topic:ai-agent)  
  LLM-powered multi-market stock analysis system with real-time news, decision dashboards, and auto-notifications.

- **[koala73/worldmonitor](https://github.com/koala73/worldmonitor)** — ⭐0 (+633 today)  
  Real-time global intelligence dashboard: AI-powered news aggregation and geopolitical monitoring.

- **[mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills)** — ⭐0 (+343 today)  
  754 pre-structured cybersecurity skills mapped to MITRE ATT&CK, NIST CSF, and other frameworks — ready for AI agents.

- **[mikumifa/biliTickerBuy](https://github.com/mikumifa/biliTickerBuy)** — ⭐0 (+164 today)  
  Bilibili ticket purchase assistant — an example of AI agent applied to e-commerce automation.

### 🧠 LLMs / Training

- **[huggingface/transformers](https://github.com/huggingface/transformers)** — ⭐161,767 (topic:llm)  
  The de facto model-definition framework for text, vision, audio, and multimodal models.

- **[hiyouga/LlamaFactory](https://github.com/hiyouga/LlamaFactory)** — ⭐72,318 (topic:llm)  
  Unified fine-tuning framework for 100+ LLMs and VLMs — the go-to tool for model customization.

- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** — ⭐7,108 (topic:llm-model)  
  Comprehensive LLM evaluation platform supporting 100+ datasets and major models.

- **[testtimescaling/testtimescaling.github.io](https://github.com/testtimescaling/testtimescaling.github.io)** — ⭐104 (topic:llm-model)  
  Survey on test-time scaling in LLMs — a frontier research direction gaining community traction.

- **[galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining)** — ⭐266 (topic:llm-model)  
  Minimal, scalable pretraining library for foundation and world models.

### 🔍 RAG / Knowledge

- **[topoteretes/cognee](https://github.com/topoteretes/cognee)** — ⭐0 (+361 today, topic:vector-db)  
  Open-source AI memory platform giving agents persistent long-term memory via self-hosted knowledge graphs.

- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** — ⭐83,269 (topic:rag)  
  Leading open-source RAG engine combining cutting-edge retrieval with agent capabilities.

- **[thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)** — ⭐83,478 (topic:rag)  
  Persistent context across sessions — compresses agent session data and injects relevant context into future sessions.

- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** — ⭐59,022 (topic:rag)  
  Universal memory layer for AI agents — the standard for cross-session state management.

- **[safishamsi/graphify](https://github.com/safishamsi/graphify)** — ⭐70,087 (topic:rag)  
  Knowledge graph creator from code, schemas, docs, images — queryable by Claude Code, Codex, and other coding agents.

- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** — ⭐12,464 (topic:vector-db)  
  Storage-efficient RAG achieving 97% savings while maintaining fast, accurate, private retrieval.

- **[VectifyAI/PageIndex](https://github.com/VectifyAI/PageIndex)** — ⭐33,262 (topic:vector-db)  
  Vectorless, reasoning-based document indexing for RAG — an alternative to traditional vector databases.

## 3. Trend Signal Analysis

The most explosive category today is **agent cost optimization and memory persistence**. `headroom` (3,795 stars) is the standout — a token compression library that cuts LLM input by 60-95% without degrading answer quality. Its rapid adoption signals a maturing market where organizations are no longer just experimenting with agents but actively seeking to reduce operational costs. This is a direct response to the scaling challenges of production agent deployments: long context windows, multi-step reasoning chains, and repeated tool calls are burning tokens at an unsustainable rate.

**Code intelligence infrastructure** is the second major trend. `codebase-memory-mcp` (+1,271 stars) and `headroom` both address the same fundamental bottleneck: LLMs need structured, compressed access to existing knowledge. The MCP (Model Context Protocol) ecosystem is becoming the connective tissue between codebases and agents, with persistent knowledge graphs emerging as the default architecture for agent memory.

A **new direction** is the commoditization of system prompts and agent skills. `system_prompts_leaks` (352 stars) and `mattpocock/skills` (+1,395) represent a shift toward prompt engineering as a shareable, open-source practice. Developers are reverse-engineering prompts from Claude, ChatGPT, Gemini, and others, then redistributing them as reusable skills.

The **multimodal agent frontier** is advancing with `OpenMontage` (agentic video production) and `palmier-pro` (AI video editor). These signal that the video generation hype is converging with agentic workflows, moving from one-shot generation to multi-stage production pipelines.

## 4. Community Hot Spots

- **Token compression and cost reduction** — `headroom` is the project to watch. Any developer deploying agents in production should evaluate it immediately for immediate 2-10x cost savings.
- **Codebase indexing via knowledge graphs** — `codebase-memory-mcp` and `safishamsi/graphify` define a new pattern: give agents persistent, queryable access to your full codebase. Expect this to become table stakes for developer tools.
- **System prompt engineering as open-source practice** — `system_prompts_leaks` and `mattpocock/skills` are creating a library of battle-tested prompts. Developers should contribute their own and watch for standardization efforts.
- **Agent memory unification** — `cognee`, `claude-mem`, and `mem0` are competing to become the universal memory layer. The winner will define how agents maintain state across sessions and tasks.
- **Cybersecurity skills for agents** — `Anthropic-Cybersecurity-Skills` (754 structured skills across 5 frameworks) represents a new category: domain-specific skill sets for AI agents. This pattern is likely to spread to legal, medical, and financial domains.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*