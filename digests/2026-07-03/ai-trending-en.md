# AI Open Source Trends 2026-07-03

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-07-03 10:12 UTC

---

# AI Open Source Trends Report – 2026-07-03

## 1. Today's Highlights

The open-source AI ecosystem is seeing explosive growth in agent harnesses and skill-based optimization tools. **msitarzewski/agency-agents** (+3,032 stars today) leads the trending list with a full multi-agent "agency" framework, while **usestrix/strix** (+2,137 stars) brings AI-powered penetration testing into the mainstream. Token economy innovation is another hot topic: **JuliusBrussee/caveman** (+926 stars) demonstrates a Claude Code skill that slashes token usage by 65% by simulating caveman speech, and **headroomlabs-ai/headroom** (56k total) offers compression of tool outputs before they reach the LLM. The rise of cross-platform agent skills (e.g., **openai/codex-plugin-cc**, **affaan-m/ECC**) and the emergence of dedicated agent specification projects like **agentskills/agentskills** signal a maturing of the agent ecosystem toward standardisation and interoperability.

## 2. Top Projects by Category

### 🔧 AI Infrastructure (frameworks, SDKs, inference engines, dev tools, CLI)

- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** – ⭐101,336 total, +65 today – The foundational deep learning framework, still actively developed and essential for nearly all AI pipelines.
- **[langflow-ai/langflow](https://github.com/langflow-ai/langflow)** – +117 today – Low-code visual platform for building AI agents and workflows, bridging the gap between prototyping and production.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** – ⭐85,252 – High-throughput LLM inference engine, critical for serving large models efficiently.
- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** – ⭐143,455 – API for web search, scraping, and interaction at scale; a backbone for agent data collection.
- **[ChromeDevTools/chrome-devtools-mcp](https://github.com/ChromeDevTools/chrome-devtools-mcp)** – +104 today – MCP server for Chrome DevTools, enabling coding agents to debug and inspect web pages programmatically.
- **[agentskills/agentskills](https://github.com/agentskills/agentskills)** – +86 today – Specification and documentation for Agent Skills, a step toward standardised agent capabilities.

### 🤖 AI Agents / Workflows (agent frameworks, automation, multi-agent systems)

- **[msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents)** – +3,032 today – A complete, opinionated multi-agent "agency" with specialised roles (frontend wizards, community managers, etc.), ready to deploy.
- **[affaan-m/ECC](https://github.com/affaan-m/ECC)** – ⭐225,451 total, +486 today – Agent harness performance optimisation system with skills, instincts, memory, and security; supports Claude Code, Codex, Cursor, and more.
- **[obra/superpowers](https://github.com/obra/superpowers)** – +897 today – An agentic skills framework and software development methodology that promises repeatable, high-quality agent output.
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** – ⭐208,369 – "The agent that grows with you" – a popular long-running autonomous agent framework.
- **[langgenius/dify](https://github.com/langgenius/dify)** – ⭐147,517 – Production-ready platform for agentic workflow development, now a staple in the RAG and agent space.
- **[langchain-ai/langchain](https://github.com/langchain-ai/langchain)** – ⭐140,816 – The most widely used agent engineering platform, constantly evolving with new integrations and abstraction layers.
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** – ⭐35,730 – Frontend stack for agents and generative UI (React, Angular, Mobile, Slack); makers of the AG-UI Protocol.

### 📦 AI Applications (specific apps, vertical solutions)

- **[usestrix/strix](https://github.com/usestrix/strix)** – +2,137 today – Open-source AI penetration testing tool that finds and fixes app vulnerabilities automatically – a security application with immediate practical impact.
- **[HKUDS/Vibe-Trading](https://github.com/HKUDS/Vibe-Trading)** – +939 today – Personal trading agent built on LLMs, targeting retail investors with automated market analysis.
- **[santifer/career-ops](https://github.com/santifer/career-ops)** – ⭐58,283 total, +372 today – AI-powered job search system with 14 skill modes, Go dashboard, and batch processing – a concrete use case for Claude Code skills.
- **[browser-use/video-use](https://github.com/browser-use/video-use)** – +554 today – Edit videos by delegating tasks to coding agents – a novel application bridging media production and AI orchestration.
- **[ZhuLinsen/daily_stock_analysis](https://github.com/ZhuLinsen/daily_stock_analysis)** – ⭐53,721 – LLM-driven multi-market stock analysis system with real-time news and automated notifications.
- **[hugohe3/ppt-master](https://github.com/hugohe3/ppt-master)** – ⭐36,341 – AI generates real, editable PowerPoint files from any document, including speaker notes and audio narration.
- **[hasaneyldrm/exercises-dataset](https://github.com/hasaneyldrm/exercises-dataset)** – +938 today – A comprehensive fitness dataset (433 exercises) with thumbnails and animations, useful for training pose estimation or recommendation models.

### 🧠 LLMs / Training (model weights, training frameworks, fine-tuning tools)

- **[ollama/ollama](https://github.com/ollama/ollama)** – ⭐175,353 – The go-to tool for running LLMs locally; now supports Kimi-K2.6, GLM-5.1, DeepSeek, and many more models.
- **[bytedance/deer-flow](https://github.com/bytedance/deer-flow)** – ⭐75,988 – Open-source long-horizon SuperAgent harness that researches, codes, and creates, with sandboxes, memories, and subagents.
- **[jingyaogong/minimind](https://github.com/jingyaogong/minimind)** – ⭐52,512 – Train a 64M-parameter LLM from scratch in just 2 hours – lowers the barrier for LLM training experimentation.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** – ⭐85,252 – Also listed in Infrastructure; key for inference but also used in fine-tuning pipelines via efficient serving.
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** – ⭐7,148 – Comprehensive LLM evaluation platform supporting 100+ datasets and multiple models – critical for benchmarking.
- **[JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman)** – ⭐82,213 total, +926 today – While not a training tool, this Claude Code skill dramatically compresses token usage, impacting inference cost and context window management.

### 🔍 RAG / Knowledge (vector databases, retrieval-augmented generation, knowledge management)

- **[open-webui/open-webui](https://github.com/open-webui/open-webui)** – ⭐143,978 – User-friendly AI interface with built-in RAG abilities, supporting Ollama, OpenAI, and more – a community favourite.
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** – ⭐84,206 – Leading open-source RAG engine combining retrieval with agent capabilities for a powerful context layer.
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** – ⭐60,010 – Universal memory layer for AI agents, enabling persistent context across sessions.
- **[headroomlabs-ai/headroom](https://github.com/headroomlabs-ai/headroom)** – ⭐56,027 – Compress tool outputs, logs, and RAG chunks by 60–95% before they reach the LLM – a complementary approach to token optimisation.
- **[FlowiseAI/Flowise](https://github.com/FlowiseAI/Flowise)** – ⭐54,238 – Visual builder for AI agents and RAG pipelines, lowering the barrier for non-developers.
- **[run-llama/llama_index](https://github.com/run-llama/llama_index)** – ⭐50,622 – Leading document agent and OCR platform, a key building block for enterprise RAG.
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** – ⭐45,059 – High-performance, cloud-native vector database essential for large-scale similarity search in RAG systems.

## 3. Trend Signal Analysis

The most explosive community attention today is around **agent skill optimisation and specification**. The simultaneous rise of *ECC*, *caveman*, *headroom*, and *agentskills* indicates that the community is moving beyond building general-purpose agents to fine-tuning their efficiency and reusability. Token compression (caveman’s 65% reduction) and output compression (headroom’s 60–95%) are addressing the real-world cost and context window bottlenecks that limit agent deployment. This suggests a maturation phase: developers are now optimising agent economics rather than just chasing new capabilities.

A new tech stack direction is the **cross-platform agent harness** – tools like *ECC* and *openai/codex-plugin-cc* explicitly support multiple agent CLIs (Claude Code, Codex, Cursor, Gemini). This interoperability push, combined with the *agentskills* specification effort, hints at an emerging “agent skill standard” that could allow skills to be shared across platforms. We may be witnessing the early days of a package ecosystem for AI agents.

The rise of **specialised vertical agent applications** (pen testing, trading, job search, video editing, PowerPoint generation) shows that agents are moving from demos to real-world productivity. The success of *career-ops* (58k total) and *ppt-master* (36k total) proves there is strong demand for domain-specific agent skills.

Finally, the **RAG and vector database** category remains the largest and most stable, with established players (Langchain, Dify, Milvus, Qdrant) continuing to grow. However, new entrants like *StarTrail-org/LEANN* (97% storage savings) and *headroom* are innovating on the efficiency frontier, suggesting that the next phase of RAG competition will be about performance and cost, not just feature breadth.

## 4. Community Hot Spots

- **Agent skill standardisation** – Watch *agentskills/agentskills* and *obra/superpowers*; they could become the de facto way to package and share agent capabilities across Claude Code, Codex, and others.
- **Token & output compression** – Projects like *caveman* and *headroom* are solving a painful real-world problem; similar techniques will likely be integrated into mainstream agent frameworks.
- **Cross-platform agent harnesses** – *ECC* and *openai/codex-plugin-cc* are blurring boundaries between AI ecosystems; developers should evaluate whether their favourite agent CLI will integrate these harnesses.
- **AI-powered security tools** – *strix*’s explosive growth signals strong demand for automated vulnerability scanning; expect more security-focused agent applications.
- **Local RAG with storage savings** – *StarTrail-org/LEANN* and *lancedb* are making RAG feasible on personal devices; valuable for privacy-conscious users and edge deployment.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*