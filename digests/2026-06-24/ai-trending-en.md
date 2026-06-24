# AI Open Source Trends 2026-06-24

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-24 10:35 UTC

---

# AI Open Source Trends Report
**Date:** 2026-06-24  
**Analyst:** Technical AI Open-Source Analyst

---

## 1. Today's Highlights

The AI open-source ecosystem is experiencing an explosive surge in **agent-harness tooling**, with multiple high-star projects converging around standardized agent skill frameworks and cross-platform compatibility. **Claude Code** continues to dominate as the preferred agent runtime, spawning a massive ecosystem of plugins, best-practice guides, and optimization tools. Notably, **NVIDIA's ECC** (agent harness optimization) and **NousResearch's Hermes Agent** have emerged as dual pillars, each exceeding 200K stars and driving a new paradigm of "agentic engineering" over "vibe coding." A distinct trend toward **specialized vertical agents** is visible — from stock analysis and cybersecurity to video production and job search — indicating the ecosystem is maturing beyond general-purpose chatbots. The rise of **MCP (Model Context Protocol)** servers for code intelligence and memory systems suggests a structural shift toward persistent, context-aware agent architectures.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure

| Project | Stars | Description |
|---------|-------|-------------|
| [vllm-project/vllm](https://github.com/vllm-project/vllm) | 83,704 total | High-throughput LLM inference engine; the de facto standard for serving open-source models in production |
| [DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp) | 1,300 today | High-performance code intelligence MCP server — indexes entire codebases into persistent knowledge graphs in milliseconds, 99% fewer tokens |
| [shareAI-lab/learn-claude-code](https://github.com/shareAI-lab/learn-claude-code) | 68,206 total | A minimal "agent harness" built from zero, demonstrating how to build Claude Code–like functionality with pure bash |
| [CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit) | 35,453 total | The frontend stack for agents and Generative UI — React, Angular, Mobile, Slack — creators of the AG-UI Protocol |

### 🤖 AI Agents / Workflows

| Project | Stars | Description |
|---------|-------|-------------|
| [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) | 201,493 total (+936 today) | "The agent that grows with you" — agentic framework from Nous Research; self-evolution and adaptive skills |
| [affaan-m/ECC](https://github.com/affaan-m/ECC) | 220,870 total (+593 today) | Agent harness performance optimization system: skills, instincts, memory, security — works with Claude Code, Codex, Cursor and beyond |
| [bytedance/deer-flow](https://github.com/bytedance/deer-flow) | 74,267 total (+739 today) | Long-horizon SuperAgent harness from ByteDance that researches, codes, and creates with sandboxes, tools, subagents, and memory |
| [garrytan/gstack](https://github.com/garrytan/gstack) | 1,011 today | Garry Tan's opinionated Claude Code setup: 23 tools acting as CEO, Designer, Eng Manager, Release Manager, Doc Engineer, and QA |
| [mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills) | 1,041 today | 817 structured cybersecurity skills mapped to 6 frameworks (MITRE ATT&CK, NIST CSF 2.0, etc.) — standard agentskills.io format |
| [santifer/career-ops](https://github.com/santifer/career-ops) | 55,477 total | AI-powered job search system built on Claude Code with 14 skill modes, Go dashboard, and PDF generation |
| [TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents) | 88,284 total | Multi-agent LLM financial trading framework — agents collaborate on market analysis and execution |
| [revfactory/harness](https://github.com/revfactory/harness) | 128 today | Meta-skill framework that designs domain-specific agent teams and generates the skills they use |

### 📦 AI Applications

| Project | Stars | Description |
|---------|-------|-------------|
| [calesthio/OpenMontage](https://github.com/calesthio/OpenMontage) | 3,592 today | World's first open-source agentic video production system: 12 pipelines, 52 tools, 500+ agent skills |
| [palmier-io/palmier-pro](https://github.com/palmier-io/palmier-pro) | 1,630 today | macOS video editor built specifically for AI agents — native Swift app |
| [jamiepine/voicebox](https://github.com/jamiepine/voicebox) | 1,045 today | Open-source AI voice studio: clone, dictate, create voices — full voice cloning pipeline |
| [ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis) | 47,881 total (+1,119 today) | LLM-powered multi-market stock analysis with real-time news, dashboards, and automated notifications |
| [CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio) | 47,743 total | AI productivity studio with smart chat, autonomous agents, and 300+ assistants — unified LLM access |
| [koala73/worldmonitor](https://github.com/koala73/worldmonitor) | 294 today | Real-time global intelligence dashboard: AI-powered news aggregation, geopolitical monitoring, and infrastructure tracking |
| [Panniantong/Agent-Reach](https://github.com/Panniantong/Agent-Reach) | 39,199 total | "Eyes for AI agents" — one CLI to read and search Twitter, Reddit, YouTube, GitHub, and more, zero API fees |

### 🧠 LLMs / Training

| Project | Stars | Description |
|---------|-------|-------------|
| [ollama/ollama](https://github.com/ollama/ollama) | 174,828 total | Local LLM runner — now supports Kimi-K2.6, GLM-5.1, DeepSeek, Qwen, Gemma and more |
| [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) | 185,141 total | The original autonomous agent vision — now focused on accessible AI tools for everyone |
| [OpenHands/OpenHands](https://github.com/OpenHands/OpenHands) | 78,207 total | AI-driven development environment — agents that code, debug, and deploy autonomously |
| [galilai-group/stable-pretraining](https://github.com/galilai-group/stable-pretraining) | 267 total | Reliable, minimal, and scalable pretraining library for foundation and world models |
| [zjunlp/LightThinker](https://github.com/zjunlp/LightThinker) | 164 total | [EMNLP 2025] Step-by-step compression for chain-of-thought reasoning — reduces token costs |

### 🔍 RAG / Knowledge

| Project | Stars | Description |
|---------|-------|-------------|
| [mem0ai/mem0](https://github.com/mem0ai/mem0) | 59,321 total | Universal memory layer for AI agents — cross-session persistent context |
| [safishamsi/graphify](https://github.com/safishamsi/graphify) | 71,391 total | Turns code, SQL schemas, docs, images, or videos into queryable knowledge graphs — works with all major coding agents |
| [infiniflow/ragflow](https://github.com/infiniflow/ragflow) | 83,517 total | Leading open-source RAG engine fusing retrieval with agent capabilities for superior context layer |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) | 84,036 total | Persistent context across sessions for every agent — captures, compresses, and injects relevant context |
| [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) | 62,012 total | Local-first RAG platform — everything needed for a private agent experience |
| [topoteretes/cognee](https://github.com/topoteretes/cognee) | 20,945 total | Self-hosted knowledge graph engine giving AI agents persistent long-term memory across sessions |

---

## 3. Trend Signal Analysis

**Agent-Harness Standardization is the Dominant Theme.** Today's data reveals a clear inflection point: the ecosystem is moving from ad-hoc agent scripts to standardized, modular agent harnesses. Projects like **ECC** (220K stars) and **learn-claude-code** (68K stars) are defining what an "agent harness" should be — complete with skills, instincts, memory, and security layers. The **agentskills.io** standard (exemplified by the cybersecurity skills repo at 1,041 today's stars) is gaining traction as a cross-platform format for reusable agent capabilities.

**Multi-Agent Orchestration Goes Mainstream.** ByteDance's **deer-flow** (739 today's stars) and **revfactory/harness** demonstrate a new paradigm: agents that design and manage teams of sub-agents. This mirrors the enterprise move from single assistants to agentic workflows where agents act as CEO, engineer, and QA simultaneously (as in **gstack**).

**Code Intelligence MCP is the New Infrastructure Layer.** The explosive growth of **codebase-memory-mcp** (1,300 today's stars) — a single binary that indexes any repo into a knowledge graph in milliseconds — signals that persistent code understanding is becoming a standard service for coding agents. This parallels the rise of **claude-mem** (84K stars) for general agent memory.

**Video and Voice are the Next Agent Frontiers.** **OpenMontage** (3,592 today's stars) and **palmier-pro** (1,630 today's stars) represent the first wave of agent-native creative tools — not just wrappers around existing APIs, but systems designed from the ground up for AI agent workflows. **voicebox** (1,045 today's stars) extends this to voice cloning and generation.

**Financial Agents Show Strong Vertical Adoption.** **TradingAgents** (88K total), **daily_stock_analysis** (1,119 today), and **OpenBB** (69K total) indicate finance remains the leading vertical for AI agent deployment, with specialized tooling for multi-source data ingestion and decision dashboards.

---

## 4. Community Hot Spots

- **🐙 [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) (201K stars, +936 today)** — Self-evolving agent architecture; closely watched as the open-source counterpart to proprietary agent systems. The "grows with you" philosophy is resonating deeply with developers who want adaptive, not static, agents.

- **🧩 [affaan-m/ECC](https://github.com/affaan-m/ECC) (220K stars)** — The "optimization system" for agent harnesses; its universal compatibility with Claude Code, Codex, Cursor, and OpenCode makes it the most cross-platform agent enhancement tool on the market. Essential for anyone building production agent workflows.

- **📹 [calesthio/OpenMontage](https://github.com/calesthio/OpenMontage) (3,592 today)** — First-mover in agentic video production; 500+ agent skills for video pipeline automation. This could define the standard for how agents handle creative production tasks.

- **🧠 [DeusData/codebase-memory-mcp](https://github.com/DeusData/codebase-memory-mcp) (1,300 today)** — High-performance MCP server for code intelligence; "99% fewer tokens" is a compelling efficiency claim. Watch for this becoming the default code indexing layer for agent IDEs.

- **🔐 [mukul975/Anthropic-Cybersecurity-Skills](https://github.com/mukul975/Anthropic-Cybersecurity-Skills) (1,041 today)** — 817 structured cybersecurity skills mapped to 6 industry frameworks. The agentskills.io standardization effort is critical for enterprise adoption of security agents.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*