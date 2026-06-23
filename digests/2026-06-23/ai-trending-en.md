# AI Open Source Trends 2026-06-23

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-23 10:50 UTC

---

# AI Open Source Trends Report — 2026-06-23

## 1. Today's Highlights

The open-source AI ecosystem is experiencing an explosive shift toward **production-grade agent harnesses** and **persistent agent memory**, with three dominant themes emerging. ByteDance's **deer-flow** (73.6K total, +738 today) represents a new class of long-horizon "SuperAgent" that combines sandboxed execution, memory graphs, and subagent orchestration — signaling the maturation of agent frameworks beyond toy demos. Meanwhile, **OpenMontage** (+2,938 today) and **hyperframes** (+395 today) are pioneering the intersection of AI agents and video production, turning coding assistants into full studio pipelines. The MCP (Model Context Protocol) ecosystem continues to expand rapidly, with **codebase-memory-mcp** (+1,185 today) delivering sub-millisecond code intelligence and **claude-context** (MCP server for code search) gaining traction. Notably, the **agent skills** movement has gone mainstream — `garrytan/gstack` (+573 today), `mattpocock/skills` (+2,051 today), and `mukul975/Anthropic-Cybersecurity-Skills` (+956 today) represent a new paradigm of sharing opinionated, reusable agent configurations across 20+ platforms.

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, Inference, Dev Tools, MCP)

- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** — ⭐137.8K (+615 today) — The de facto API for web data extraction for AI agents, now integrating with agent workflows at scale.
- **[ollama/ollama](https://github.com/ollama/ollama)** — ⭐174.8K — Local LLM runner now supporting Kimi-K2.6, GLM-5.1, and other frontier models; the go-to for on-device inference.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** — ⭐83.6K — High-throughput LLM inference engine, essential for production deployments.
- **[huggingface/transformers](https://github.com/huggingface/transformers)** — ⭐161.8K — The universal model-definition framework; continues to be the bedrock of open-source AI.
- **[DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp)** — ⭐1.2K (+1,185 today) — High-performance MCP server that indexes codebases into persistent knowledge graphs; 158 languages, sub-ms queries, zero dependencies — a standout for developer tooling.
- **[zilliztech/claude-context](https://github.com/zilliztech/claude-context)** — ⭐11.9K — Code search MCP for Claude Code, making entire codebase context available to coding agents.
- **[ai-website-cloner-template](https://github.com/JCodesMore/ai-website-cloner-template)** ⭐100 (+100 today) — Demonstrates the emerging pattern of using AI coding agents for website cloning with a single command.

### 🤖 AI Agents / Workflows

- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** — ⭐73.6K (+738 today) — An open-source long-horizon SuperAgent harness that researches, codes, and creates using sandboxes, memories, tools, subagents, and message gateways — handles minute-to-hour tasks.
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** — ⭐35.4K — Frontend stack for agents & Generative UI supporting React, Angular, Mobile, and Slack; makers of the AG-UI Protocol.
- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** — ⭐100.2K — Makes websites accessible for AI agents; the standard for browser-based task automation.
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** — ⭐78.1K — AI-driven development platform; established leader for autonomous software engineering agents.
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** — ⭐185.1K — The original accessible AI agent vision; continues to evolve with new capabilities.
- **[zhayujie/CowAgent](https://github.com/zhayujie/CowAgent)** — ⭐45.6K — Open-source super AI assistant & agent harness with plans, tools, skills, and self-evolving memory — lightweight and extensible.
- **[shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code)** — ⭐68.0K — A nano "agent harness" built from scratch; essential educational resource for understanding agent internals.
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** — ⭐47.7K — AI productivity studio with smart chat, autonomous agents, and 300+ assistants — unified access to frontier LLMs.

### 📦 AI Applications (Vertical Solutions)

- **[calesthio/OpenMontage](https://github.com/calesthio/OpenMontage)** — ⭐2.9K (+2,938 today) — World's first open-source agentic video production system: 12 pipelines, 52 tools, 500+ agent skills — turns AI coding assistants into full video production studios. **Today's most explosive growth.**
- **[palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro)** — ⭐2.5K (+2,463 today) — macOS video editor built for AI; complements the video-generation surge.
- **[heygen-com/hyperframes](https://github.com/heygen-com/hyperframes)** — ⭐395 (+395 today) — Write HTML, render video — built specifically for agents; signals a new paradigm for programmatic video generation.
- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** — ⭐46.4K (+1,557 today) — LLM-powered multi-market stock analysis system with real-time news, decision dashboards, and automated notifications — zero-cost scheduled runs.
- **[TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents)** — ⭐88.1K — Multi-agent LLM financial trading framework; one of the highest-starred finance-AI projects.
- **[jamiepine/voicebox](https://github.com/jamiepine/voicebox)** — ⭐0.5K (+529 today) — Open-source AI voice studio: clone, dictate, and create — tapping into the voice AI wave.
- **[mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills)** — ⭐1.0K (+956 today) — 817 structured cybersecurity skills for AI agents mapped to 6 frameworks — works with 20+ platforms.
- **[OpenBB-finance/OpenBB](https://github.com/OpenBB-finance/OpenBB)** — ⭐69.6K — Financial data platform for analysts, quants, and AI agents — the infrastructure layer for AI trading.

### 🧠 LLMs / Training

- **[airllm](https://github.com/lyogavin/airllm)** — ⭐193 (+193 today) — 70B model inference on a single 4GB GPU; democratizes LLM access for resource-constrained setups.
- **[huggingface/transformers](https://github.com/huggingface/transformers)** — ⭐161.8K — Also the primary training and inference framework for the entire open-source LLM ecosystem.
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** — ⭐200.4K — "The agent that grows with you" — among the highest-starred agent + model projects.
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** — ⭐7.1K — LLM evaluation platform supporting 100+ datasets and all major model families — critical for model benchmarking.
- **[Eigenwise/atomic-agents](https://github.com/Eigenwise/atomic-agents)** — ⭐6.0K — Building AI agents, atomically — a modular approach to agent construction.

### 🔍 RAG / Knowledge (Vector Databases, Retrieval, Knowledge Management)

- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** — ⭐59.2K — Universal memory layer for AI agents; the leading solution for persistent agent context.
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** — ⭐83.4K — Open-source RAG engine fusing retrieval-augmented generation with Agent capabilities — a leading context layer for LLMs.
- **[Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm)** — ⭐62.0K — Local-first agent experience with full RAG capabilities — "stop renting your intelligence."
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** — ⭐44.9K — High-performance cloud-native vector database for scalable ANN search; the industry standard.
- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** — ⭐32.6K — High-performance, massive-scale vector database; also available as a cloud service.
- **[PaddlePaddle/PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)** — ⭐83.4K — Turns any PDF or image into structured data for AI — bridges images/PDFs and LLMs with 100+ language support.
- **[safishamsi/graphify](https://github.com/safishamsi/graphify)** — ⭐70.9K — AI coding assistant skill that turns folders of code, schemas, scripts, docs, and images into queryable knowledge graphs.
- **[headroomlabs-ai/headroom](https://github.com/headroomlabs-ai/headroom)** — ⭐47.7K — Compresses tool outputs, logs, files, and RAG chunks by 60-95% before reaching the LLM — a novel token optimization approach.

## 3. Trend Signal Analysis

**Agent Harnesses Are Going Mainstream.** The single most explosive trend visible today is the rise of **structured, shareable agent configurations** — what the community is calling "agent harnesses" or "agent skills." Projects like `deer-flow`, `shareAI-lab/learn-claude-code`, `mattpocock/skills`, and `garrytan/gstack` represent a paradigm shift: instead of building agents from scratch, developers are now importing opinionated, production-tested setups that include tool definitions, memory configurations, subagent topologies, and skill libraries. The `.claude` directory is becoming a new standard for sharing agent expertise, akin to how `.github` workflows standardized CI/CD.

**Long-Horizon Autonomy Is the New Frontier.** ByteDance's `deer-flow` and the `Anthropic-Cybersecurity-Skills` repo (817 structured skills) signal that the community is moving beyond single-turn prompting and simple tool-calling toward **persistent, multi-hour agentic tasks**. The combination of sandboxed execution, persistent memory graphs, subagent delegation, and skill libraries creates a new category of "SuperAgent" capable of research, coding, and creative production. This aligns with the industry-wide push toward agentic workflows that can plan, execute, and self-correct over extended periods.

**MCP Protocol Maturation.** The appearance of multiple high-quality MCP (Model Context Protocol) servers — `codebase-memory-mcp`, `claude-context`, and `headroom` — confirms that MCP is becoming the standard interface for connecting LLMs to external data sources. Codebase-memory-mcp's claim of "99% fewer tokens" through indexed knowledge graphs is particularly noteworthy, addressing the fundamental cost and context-window challenges of agent-based development.

**Video Generation Meets Agentic Pipelines.** `OpenMontage`'s explosive debut (+2,938 stars in one day) and `hyperframes`' "Write HTML. Render video. Built for agents" approach signal a new vertical: **agentic video production**. Rather than using standalone AI video tools, developers are now wiring LLM agents into multi-tool video creation pipelines — a pattern likely to spread to other creative domains.

**The Agent Skills Economy.** The rapid adoption of `skills` directories (mattpocock's 2K+ stars, the cybersecurity skills repo, career-ops with 55K stars) points to an emerging **agent skills marketplace**. Developers are packaging domain expertise (cybersecurity, career ops, finance) as reusable agent skill sets, mapped to industry frameworks. This could evolve into a decentralized skill registry, similar to npm packages but for agent capabilities.

## 4. Community Hot Spots

- 🔥 **Long-Horizon Agent Harnesses** — ByteDance's [`deer-flow`](https://github.com/bytedance/deer-flow) and the agent harness movement (`shareAI-lab/learn-claude-code`, `zhayujie/CowAgent`) — the most active area for developers building production agent systems that handle minute-to-hour tasks.
- 🔥 **Code Intelligence MCP Servers** — [`codebase-memory-mcp`](https://github.com/DeusData/codebase-memory-mcp) and [`zilliztech/claude-context`](https://github.com/zilliztech/claude-context) — solving the "codebase context" problem for coding agents; performance and token efficiency are the key differentiators.
- 🔥 **Agent Skills / .claude Ecosystem** — [`mattpocock/skills`](https://github.com/mattpocock/skills), [`garrytan/gstack`](https://github.com/garrytan/gstack), and [`mukul975/Anthropic-Cybersecurity-Skills`](https://github.com/mukul975/Anthropic-Cybersecurity-Skills) — the emerging standard for sharing opinionated agent configurations; watch for registry/template systems.
- 🔥 **Agentic Video Production** — [`OpenMontage`](https://github.com/calesthio/OpenMontage) and [`hyperframes`](https://github.com/heygen-com/hyperframes) — the intersection of AI agents and creative tooling; video is just the first vertical.
- 🔥 **Persistent Agent Memory** — [`mem0ai/mem0`](https://github.com/mem0ai/mem0) and [`thedotmack/claude-mem`](https://github.com/thedotmack/claude-mem) — solving the "session reset" problem with AI-driven compression and context injection; critical for all agentic workflows.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*