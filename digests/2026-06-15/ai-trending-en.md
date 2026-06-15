# AI Open Source Trends 2026-06-15

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-06-15 03:43 UTC

---

# AI Open Source Trends Report
**Date**: 2026-06-15
**Data Sources**: GitHub Trending, AI Topic Search (7-day active projects with topic tags)

---

## 1. Today's Highlights

Today's AI open-source landscape is dominated by three converging themes: **AI agent security** hitting mainstream awareness, **vertical AI applications** (finance, robotics) attracting explosive interest, and the continued consolidation of **agent infrastructure and orchestration** tools. NVIDIA's SkillSpector (964 stars today) signals the market's growing recognition that agent safety is no longer optional—it's a prerequisite for production deployment. Meanwhile, the financial AI vertical sees two major entries: Kronos, a foundation model for financial markets (244 stars), and TradingAgents (86K total), demonstrating that domain-specific agents are moving from experiments to serious tools. The emergence of "model-agnostic" and "provider-agnostic" platforms like andrewyng/aisuite (291 stars today) and Cherry Studio (47K total) reflects a maturing ecosystem where developers demand flexibility and vendor independence.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure (Frameworks, SDKs, Inference Engines, Dev Tools)

- **[andrewyng/aisuite](https://github.com/andrewyng/aisuite)** – ⭐291 today. A simple, unified interface to multiple Generative AI providers, reducing provider lock-in and simplifying multi-model development.
- **[langchain-ai/langchain](https://github.com/langchain-ai/langchain)** – ⭐139K total. The leading agent engineering platform for Python, now deeply integrated with MCP protocol and multi-agent orchestration.
- **[langchain4j/langchain4j](https://github.com/langchain4j/langchain4j)** – ⭐12.3K total. The JVM-native LangChain counterpart, enabling enterprise Java/Spring Boot teams to build LLM-powered applications with tool calling and RAG.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** – ⭐82.9K total. High-throughput inference engine for LLMs, now supporting Kimi-K2.6, DeepSeek, and other latest models (as reflected in Ollama's updated model list).
- **[0xPlaygrounds/rig](https://github.com/0xPlaygrounds/rig)** – ⭐7.6K total. A Rust-native framework for building modular, scalable LLM applications—gaining traction in systems programming communities.
- **[Picovoice/picollm](https://github.com/Picovoice/picollm)** – ⭐312 total. On-device LLM inference with X-bit quantization, enabling edge deployment scenarios.

### 🤖 AI Agents / Workflows (Agent Frameworks, Automation, Multi-Agent)

- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** – ⭐193.7K total. "The agent that grows with you"—a highly popular agent framework emphasizing memory, self-evolution, and multi-model support.
- **[CherryHQ/cherry-studio](https://github.com/CherryHQ/cherry-studio)** – ⭐47.3K total. AI productivity studio unifiying smart chat, autonomous agents, and 300+ skill assistants—a "all-in-one" agent workspace.
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** – ⭐184.9K total. The original autonomous agent project, still actively maintained and evolving with new planning and tool-use capabilities.
- **[OpenHands/OpenHands](https://github.com/OpenHands/OpenHands)** – ⭐77.1K total. AI-driven development agent that writes, debugs, and deploys code autonomously.
- **[browser-use/browser-use](https://github.com/browser-use/browser-use)** – ⭐98.8K total. Makes websites accessible for AI agents by automating browser interactions—critical for web-based agent workflows.
- **[CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit)** – ⭐35.1K total. Frontend stack for agents & generative UI (React, Angular, Mobile, Slack)—makers of the AG-UI Protocol.
- **[FlowiseAI/Flowise](https://github.com/FlowiseAI/Flowise)** – ⭐53.6K total. Low-code visual builder for AI agents, making agent development accessible to non-engineers.

### 📦 AI Applications (Specific Apps, Vertical Solutions)

- **[shiyu-coder/Kronos](https://github.com/shiyu-coder/Kronos)** – ⭐244 today. A foundation model for financial markets, representing a new class of domain-specialized LLMs beyond general-purpose models.
- **[TauricResearch/TradingAgents](https://github.com/TauricResearch/TradingAgents)** – ⭐86.2K total. Multi-agent LLM financial trading framework with sophisticated portfolio management and risk analysis.
- **[OpenBB-finance/OpenBB](https://github.com/OpenBB-finance/OpenBB)** – ⭐69.2K total. Financial data platform for analysts, quants and AI agents—integrating multiple data sources with agent capabilities.
- **[acon96/home-llm](https://github.com/acon96/home-llm)** – ⭐1.4K total. Local LLM integration for Home Assistant, enabling privacy-preserving smart home automation.
- **[NVIDIA/SkillSpector](https://github.com/NVIDIA/SkillSpector)** – ⭐964 today. Security scanner for AI agent skills—detects vulnerabilities, malicious patterns, and security risks in agent toolchains.
- **[Introduction-to-Autonomous-Robots/Introduction-to-Autonomous-Robots](https://github.com/Introduction-to-Autonomous-Robots/Introduction-to-Autonomous-Robots)** – ⭐293 today. Educational resource for autonomous robotics bridging LLM capabilities with physical world interaction.

### 🧠 LLMs / Training (Models, Weights, Training Frameworks, Fine-tuning)

- **[huggingface/transformers](https://github.com/huggingface/transformers)** – ⭐161.6K total. The de facto standard framework for model definition and training across modalities (text, vision, audio, multimodal).
- **[ollama/ollama](https://github.com/ollama/ollama)** – ⭐174.2K total. Local LLM runner now supporting Kimi-K2.6, GLM-5.1, MiniMax, DeepSeek, Qwen, Gemma—becoming the universal model runner for developers.
- **[pytorch/pytorch](https://github.com/pytorch/pytorch)** – ⭐100.8K total. Foundational ML framework underpinning virtually all modern LLM training and fine-tuning.
- **[tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)** – ⭐195.7K total. Still widely used for production ML deployments, especially in enterprise settings.
- **[open-compass/opencompass](https://github.com/open-compass/opencompass)** – ⭐7.1K total. Comprehensive LLM evaluation platform supporting 100+ datasets and all major models—critical for model selection decisions.
- **[skyzh/tiny-llm](https://github.com/skyzh/tiny-llm)** – ⭐4.3K total. Educational project teaching LLM inference serving (build a tiny vLLM + Qwen)—a sign of growing interest in understanding internals.

### 🔍 RAG / Knowledge (Vector Databases, Retrieval-Augmented Generation, Knowledge Management)

- **[langgenius/dify](https://github.com/langgenius/dify)** – ⭐145.2K total. Production-ready RAG platform with agentic workflow development, combining document retrieval with multi-step agent reasoning.
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** – ⭐82.7K total. Leading open-source RAG engine fusing cutting-edge retrieval with agent capabilities for superior context layer.
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** – ⭐44.8K total. High-performance cloud-native vector database for scalable ANN search—the most deployed vector DB in production.
- **[qdrant/qdrant](https://github.com/qdrant/qdrant)** – ⭐32.3K total. Rust-based vector database with high performance and massive scale, also available as managed cloud.
- **[NirDiamant/RAG_Techniques](https://github.com/NirDiamant/RAG_Techniques)** – ⭐27.9K total. Comprehensive notebook tutorials covering advanced RAG techniques—a go-to learning resource.
- **[topoteretes/cognee](https://github.com/topoteretes/cognee)** – ⭐17.8K total. Open-source AI memory platform for agents with persistent long-term memory via self-hosted knowledge graph engine.
- **[StarTrail-org/LEANN](https://github.com/StarTrail-org/LEANN)** – ⭐11.9K total. [MLsys2026] Enables RAG on everything with 97% storage savings while maintaining accurate, private on-device retrieval.
- **[neuml/txtai](https://github.com/neuml/txtai)** – ⭐12.7K total. All-in-one AI framework for semantic search, LLM orchestration and language model workflows.

---

## 3. Trend Signal Analysis

**Agent Security Becomes Non-Negotiable**: NVIDIA's SkillSpector (+964 today) represents a critical inflection point. As autonomous agents gain tool-calling capabilities (browser-use, file access, API calls), the community is recognizing that every skill is a potential attack vector. This trend is reinforced by the proliferation of skill marketplaces and the "agent-as-a-service" model—without security scanning, malicious skills could exfiltrate data, execute unauthorized commands, or poison agent memory. Expect this to spawn a new category of "Agent Security Posture Management" (ASPM) tools.

**Financial AI Vertical Reaches Escape Velocity**: The simultaneous traction of Kronos (+244), TradingAgents (86K total), and OpenBB (69K) signals that finance is becoming the killer vertical for AI agents. Unlike general coding agents, financial agents require multi-modal data processing (news, price feeds, SEC filings), real-time decision-making, and portfolio-level reasoning—pushing agent architectures toward specialized, multi-agent systems. The "LLM + financial data APIs + risk management" stack is emerging as a standard blueprint.

**Model-Agnostic Infrastructure Becomes Mainstream**: Aisuite (+291 today) joins a growing trend toward abstraction layers that decouple applications from specific providers. This is reinforced by Cherry Studio's support for 300+ assistants across multiple LLMs, Ollama's model-agnostic local runtime, and the MCP protocol's widespread adoption. The community is voting against lock-in—whether OpenAI, Anthropic, or open-source models—by building tools that can swap between providers while maintaining consistent behavior.

**Memory and Knowledge Persistence Are Infrastructure, Not Features**: Projects like cognee (17.8K), mem0 (58.6K), and tarl (via claude-mem's 82.3K) are redefining "memory" as a first-class infrastructure layer for agents. The emerging stack separates short-term context windows from persistent long-term memory stored in knowledge graphs or vector stores, with compression and retrieval mechanisms built in. This mirrors how databases evolved from application-specific features to standalone infrastructure.

**Edge and On-Device Inference Gains Steam**: Tiny-LLM (4.3K), picollm (312), and the continued popularity of ollama (174K) indicate strong developer interest in running models locally—for privacy, latency, and cost reasons. The "on-device LLM" trend is no longer just about chatbots; it's about powering autonomous robots, home automation, and privacy-sensitive financial applications without cloud round-trips.

---

## 4. Community Hot Spots

- **🛡️ Agent Security (NVIDIA/SkillSpector, ECC)**: Explosive growth in agent adoption demands security tooling. SkillSpector's 964 daily stars show developers are actively seeking ways to audit and secure agent skill chains. Expect this to become as essential as code scanning in CI/CD pipelines.

- **💰 Financial AI Agents (Kronos, TradingAgents, OpenBB)**: Financial multi-agent systems are the most active vertical today. The combination of real-time data APIs, LLM reasoning, and portfolio-level optimization is creating a new class of "quant agents" that could reshape algorithmic trading.

- **🧠 Long-Term Memory Infrastructure (mem0, cognee, claude-mem)**: Persistent, queryable memory is the missing piece for production agents. These projects are laying the foundation for agents that learn, remember, and improve over time—transforming them from stateless tools to stateful collaborators.

- **🌐 Web-Native Agent Interaction (browser-use, firecrawl, OpenCLI)**: The ability for agents to interact with any website—not just those with APIs—is unlocking vast automation possibilities. browser-use's 98K total stars and firecrawl's 132K reflect massive enthusiasm for web-native agent capabilities.

- **⚡ Multi-Model Unified Interfaces (aisuite, Cherry Studio, Ollama)**: As the model landscape fragments (Kimi, GLM, DeepSeek, Qwen, Gemma…), tools that provide a single interface across providers with zero code changes are becoming indispensable. This "write once, run anywhere" approach for LLMs is a strong bet for long-term adoption.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*