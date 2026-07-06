# AI Open Source Trends 2026-07-06

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-07-06 13:05 UTC

---

# AI Open Source Trends Report — 2026-07-06

## 1. Today's Highlights

The open-source AI ecosystem is experiencing an extraordinary wave of **agent skill commoditization** and **system prompt exposure**. The trending list is dominated by repositories that package engineering-grade skills, prompts, and plugins for AI coding agents (Claude Code, Codex, Gemini CLI, Cursor). The single largest signal is the explosive popularity of `system_prompts_leaks` (1,386 stars today), which publicly catalogs exact system prompts from Anthropic, OpenAI, Google, and xAI — indicating a community-wide hunger to understand and optimize agent behavior. Simultaneously, Rust-based local AI applications (`meetily`, `herdr`) are gaining traction, signaling a shift toward privacy-first, high-performance AI tooling that runs entirely on-device. The vector database space remains vibrant, with Alibaba's new lightweight in-process database `zvec` (355 stars today) competing in the embedded AI storage layer.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure (frameworks, SDKs, inference engines, dev tools, CLI)

- **[firecrawl/firecrawl](https://github.com/firecrawl/firecrawl)** ⭐145,674 (+834 today) — The leading open-source API for web scraping and data ingestion at scale, now a critical data pipeline component for RAG and agent workflows.
- **[alibaba/zvec](https://github.com/alibaba/zvec)** ⭐13,186 (+355 today) — Alibaba's new lightweight, lightning-fast in-process vector database written in C++, optimized for embedded and edge deployment scenarios.
- **[steipete/CodexBar](https://github.com/steipete/CodexBar)** ⭐0 (+598 today) — A macOS menubar utility that displays real-time usage statistics for OpenAI Codex and Claude Code API consumption — developer tooling meets cost observability.
- **[ogulcancelik/herdr](https://github.com/ogulcancelik/herdr)** ⭐0 (+783 today) — Terminal-based agent multiplexer that routes requests across multiple AI coding agents, enabling parallel execution and redundancy.
- **[vllm-project/vllm](https://github.com/vllm-project/vllm)** ⭐85,479 — High-throughput LLM inference engine that continues to be the standard for serving open models in production environments.
- **[rig](https://github.com/0xPlaygrounds/rig)** ⭐7,840 — Modular Rust framework for building LLM applications, gaining traction as Rust becomes more prominent in the AI toolchain.

### 🤖 AI Agents / Workflows (agent frameworks, automation, multi-agent systems)

- **[alirezarezvani/claude-skills](https://github.com/alirezarezvani/claude-skills)** ⭐0 (+611 today) — A massive collection of 337 pre-built skills and plugins for Claude Code and 8 other coding agents — the most comprehensive agent skill marketplace seen to date.
- **[addyosmani/agent-skills](https://github.com/addyosmani/agent-skills)** ⭐0 (+1,114 today) — Production-grade engineering skill templates — think of it as an agent skill standard library — authored by Google Chrome engineering leader Addy Osmani.
- **[openai/codex-plugin-cc](https://github.com/openai/codex-plugin-cc)** ⭐0 (+910 today) — OpenAI's official plugin enabling Codex to be invoked from Claude Code — a rare cross-platform interoperability move between competing agent ecosystems.
- **[gastownhall/gastown](https://github.com/gastownhall/gastown)** ⭐0 (+293 today) — Multi-agent workspace manager written in Go, designed for orchestrating teams of specialized agents working on shared tasks.
- **[NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent)** ⭐210,095 — The agent framework "that grows with you" — one of the most starred repositories on GitHub, representing the community's favorite open-source agent harness.
- **[Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT)** ⭐185,402 — The original autonomous agent framework continues to evolve, now at a scale that makes it a foundational reference for all agent projects.

### 📦 AI Applications (specific apps, vertical solutions)

- **[Zackriya-Solutions/meetily](https://github.com/Zackriya-Solutions/meetily)** ⭐0 (+2,493 today) — **Today's #1 trending project.** Privacy-first AI meeting assistant using Parakeet/Whisper for 4x faster live transcription and local Ollama-based summarization — 100% offline on macOS/Windows.
- **[asgeirtj/system_prompts_leaks](https://github.com/asgeirtj/system_prompts_leaks)** ⭐0 (+1,386 today) — The most comprehensive public archive of proprietary AI agent system prompts (Claude, ChatGPT, Gemini, Grok, Copilot, Codex, Perplexity) — invaluable for reverse engineering and competitive analysis.
- **[bradautomates/claude-video](https://github.com/bradautomates/claude-video)** ⭐0 (+368 today) — Gives Claude the ability to watch video files by downloading, frame extraction, transcription, and bulk analysis — bridging the video understanding gap for text-native agents.
- **[karakeep-app/karakeep](https://github.com/karakeep-app/karakeep)** ⭐0 (+178 today) — Self-hostable bookmark manager with AI-powered automatic tagging and full-text search — a practical application of local ML classification.
- **[Leonxlnx/taste-skill](https://github.com/Leonxlnx/taste-skill)** ⭐0 (+1,453 today) — An agent skill that injects "taste" into AI outputs, reducing generic and boring responses — directly addressing the homogenization problem in LLM-generated text.
- **[cherry-studio](https://github.com/CherryHQ/cherry-studio)** ⭐48,215 — AI productivity studio with 300+ assistants and autonomous agents, unifying access to frontier LLMs in a single desktop application.

### 🧠 LLMs / Training (model weights, training frameworks, fine-tuning tools)

- **[ollama/ollama](https://github.com/ollama/ollama)** ⭐175,587 — The simplest way to run local LLMs, now supporting Kimi-K2.6, GLM-5.1, MiniMax, DeepSeek, Qwen, and Gemma — the universal on-device model runtime.
- **[open-webui/open-webui](https://github.com/open-webui/open-webui)** ⭐144,410 — The most popular user-friendly interface for local LLMs (Ollama + OpenAI API), now effectively the de facto open-source ChatGPT alternative.
- **[JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman)** ⭐85,449 — A Claude Code skill that cuts token usage by 65% using a "caveman speak" compression technique — hilarious in concept but serious in its cost-saving implications for heavy API users.
- **[AarambhDevHub/aarambh-ai](https://github.com/AarambhDevHub/aarambh-ai)** ⭐9 — A decoder-only LLM built entirely in Rust using Candle, supporting INT4/GGUF quantization and LoRA/QLoRA fine-tuning — a showcase of the Rust AI ecosystem maturing.

### 🔍 RAG / Knowledge (vector databases, retrieval-augmented generation, knowledge management)

- **[mvanhorn/last30days-skill](https://github.com/mvanhorn/last30days-skill)** ⭐0 (+237 today) — An agent skill that researches any topic across Reddit, X, YouTube, HN, Polymarket, and the web, then synthesizes grounded summaries — RAG for social and news data.
- **[mem0ai/mem0](https://github.com/mem0ai/mem0)** ⭐60,207 — Universal memory layer for AI agents, providing persistent context across sessions — a critical infrastructure piece for long-running agent workflows.
- **[Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify)** ⭐78,463 — Transforms any folder of code, SQL schemas, docs, or images into a queryable knowledge graph — bridging RAG with graph-based AI reasoning.
- **[milvus-io/milvus](https://github.com/milvus-io/milvus)** ⭐45,093 — The leading cloud-native vector database for scalable ANN search, now a mature standard in production RAG stacks.
- **[infiniflow/ragflow](https://github.com/infiniflow/ragflow)** ⭐84,406 — Open-source RAG engine combining retrieval-augmented generation with agent capabilities, creating a superior LLM context layer.

---

## 3. Trend Signal Analysis

**The dominant signal is the emergence of an "Agent Skill Economy."** Today's trending list reveals a structural shift: developers are no longer just building agent _frameworks_ — they are building and trading agent _skills_ as composable, reusable modules. The success of `claude-skills` (337 skills), `agent-skills`, and `taste-skill` demonstrates that the community has reached a consensus that agents need standardized "skill packs" analogous to VS Code extensions or npm packages. The explosive popularity of `system_prompts_leaks` (1,386 stars) confirms that prompt engineering has become a form of competitive intelligence — developers want to know exactly how Anthropic, OpenAI, and Google configure their agents.

**Rust is making a serious push into AI tooling.** Three Rust projects hit the trending list today: `meetily` (transcription), `herdr` (agent multiplexer), and `zvec` (vector database). This represents a departure from Python's dominance in AI and suggests a growing demand for performance-critical, privacy-preserving, local-first AI tools. The 2,493 stars for `meetily` — a fully offline meeting assistant — signals that users are actively seeking alternatives to cloud-dependent AI services.

**Cross-platform agent interoperability is becoming explicit.** OpenAI's release of `codex-plugin-cc` — allowing Codex to be called from Claude Code — is unprecedented. This is a tacit admission that the future is multi-agent: developers will mix and match agents from different providers. Tools like `herdr` (agent multiplexer) and `gastown` (multi-agent workspace manager) are infrastructure being built to support this reality.

**Token optimization is a pressing community concern.** The `caveman` skill (85K stars) — a joke that compresses token usage by 65% — is a dead-serious indicator that API costs are driving innovation. Combined with `headroom` (57K stars, token compression library) and the broader skill ecosystem, the community is actively seeking ways to extract more value per token.

---

## 4. Community Hot Spots

- **[asgeirtj/system_prompts_leaks](https://github.com/asgeirtj/system_prompts_leaks)** — The single most important resource for understanding how frontier AI labs configure their agents. Every agent developer should study these prompts as they represent the "operating system" of modern AI. **Focus for: competitive analysis, prompt engineering research.**

- **[alirezarezvani/claude-skills](https://github.com/alirezarezvani/claude-skills)** — At 337 skills with 30+ agents supported, this is the most comprehensive skill marketplace in existence. It represents where the ecosystem is standardizing. **Focus for: developers building agent workflows who want to avoid reinventing common patterns.**

- **[Zackriya-Solutions/meetily](https://github.com/Zackriya-Solutions/meetily)** — The #1 trending project today. It demonstrates that privacy-first, local AI can outperform cloud solutions in latency-sensitive tasks like real-time transcription. Written in Rust, it's a blueprint for on-device AI application architecture. **Focus for: privacy engineers, Rust developers, meeting tool builders.**

- **[openai/codex-plugin-cc](https://github.com/openai/codex-plugin-cc)** — OpenAI's cross-platform plugin for Claude Code is a major strategic signal. It suggests a future where agent interoperability is a product feature, not an accident. **Focus for: anyone building multi-agent systems, API integration architects.**

- **[JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman)** — Despite its comedic premise, this skill (85K+ stars) directly addresses the most expensive problem in LLM application development: token waste. The compression technique — forcing concise, primitive output — is a viable production strategy for cost-sensitive deployments. **Focus for: cost optimization, prompt compression research, production API users.**

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*