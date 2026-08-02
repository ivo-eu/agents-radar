# AI Open Source Trends 2026-08-02

> Sources: GitHub Trending + GitHub Search API | Generated: 2026-08-02 00:13 UTC

---

## Filtering Note

From the trending list, I excluded non-AI/ML repos: `usekaneo/kaneo`, `github/gh-stack`, `iv-org/invidious`, `ansible/ansible`, and `paperswithbacktest/awesome-systematic-trading` (not clearly AI-specific). All Topic Search results were already AI/ML-tagged, so most were retained for categorization.

---

## 1. Today's Highlights

The AI open-source ecosystem today is dominated by **agent infrastructure and skill ecosystems** rather than single models. The highest-starred trending repo is `reverse-skill`, an AI-powered security skill router for Claude Code/Cursor/Cline — a sign that AI coding clients are becoming extensible agent platforms. **Memory and context layers** are also crystallizing: TencentDB-Agent-Memory, `mem0`, `claude-mem`, and `cognee` all target persistent agent memory. **Voice AI** and **3D generation** drew notable attention via Hugging Face's `speech-to-speech` and Microsoft's `TRELLIS.2`. At the same time, Microsoft's beginner AI/GenAI curricula racked up ~1,050 combined today's stars, indicating strong onboarding demand.

---

## 2. Top Projects by Category

### 🔧 AI Infrastructure

- [github/copilot-sdk](https://github.com/github/copilot-sdk) — Today +142 · Official multi-platform SDK for integrating GitHub Copilot Agent into apps/services; a major signal that coding agents are opening up as platforms.
- [ollama/ollama](https://github.com/ollama/ollama) — ⭐177,524 · Local LLM runner now highlighting Kimi-K2.6, GLM-5.2, MiniMax, DeepSeek, Qwen, Gemma; the default for local model serving.
- [huggingface/transformers](https://github.com/huggingface/transformers) — ⭐163,226 · The model-definition framework for SOTA transformer models across text, vision, audio, and multimodal tasks.
- [vllm-project/vllm](https://github.com/vllm-project/vllm) — ⭐87,884 · High-throughput, memory-efficient LLM inference and serving engine.
- [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) — ⭐159,096 · Web search/scrape/interact API purpose-built for LLM and agent data ingestion.
- [headroomlabs-ai/headroom](https://github.com/headroomlabs-ai/headroom) — ⭐63,874 · Compresses tool outputs, logs, files, and RAG chunks before they reach the LLM — a new token-efficiency layer for agents.
- [CopilotKit/CopilotKit](https://github.com/CopilotKit/CopilotKit) — ⭐36,400 · Frontend stack for building agentic UI and generative interfaces in React/Angular/Mobile.
- [TencentCloud/TencentDB-Agent-Memory](https://github.com/TencentCloud/TencentDB-Agent-Memory) — Today +227 · Team-level memory hub for AI agents: chat memory, skills, LLM-Wiki, and code-graph assets shared across agents.

### 🤖 AI Agents / Workflows

- [langchain-ai/langchain](https://github.com/langchain-ai/langchain) — ⭐143,185 · The leading agent engineering platform for building LLM-powered workflows.
- [Significant-Gravitas/AutoGPT](https://github.com/Significant-Gravitas/AutoGPT) — ⭐185,751 · The foundational open-source autonomous agent platform.
- [NousResearch/hermes-agent](https://github.com/NousResearch/hermes-agent) — ⭐223,835 · "The agent that grows with you" — a self-evolving agent with memory and tool use.
- [bytedance/deer-flow](https://github.com/bytedance/deer-flow) — Today +209 · Open-source long-horizon SuperAgent harness with sandboxes, memories, tools, skills, subagents, and a message gateway.
- [FlowiseAI/Flowise](https://github.com/FlowiseAI/Flowise) — ⭐55,090 · Visual builder for AI agents and workflows, now a standard low-code agent platform.
- [affaan-m/ECC](https://github.com/affaan-m/ECC) — ⭐236,839 · Agent harness performance optimization system: skills, instincts, memory, security, and research-first development for Claude Code/Codex/Cursor.
- [zhaoxuya520/reverse-skill](https://github.com/zhaoxuya520/reverse-skill) — Today +1,320 · AI-powered security skill router for authorized pentesting/reverse engineering; top trending repo today.
- [NomaDamas/k-skill](https://github.com/NomaDamas/k-skill) — Today +53 · A Korean-language skill collection that makes AI coding agents "Korean natives" — a localization trend for agent skills.

### 📦 AI Applications

- [open-webui/open-webui](https://github.com/open-webui/open-webui) — ⭐147,544 · User-friendly, self-hosted AI interface supporting Ollama and OpenAI-compatible APIs.
- [Mintplex-Labs/anything-llm](https://github.com/Mintplex-Labs/anything-llm) — ⭐64,209 · Local-first AI agent experience with full RAG and "own your intelligence" philosophy.
- [huggingface/speech-to-speech](https://github.com/huggingface/speech-to-speech) — Today +442 · Build local voice agents using open-source models; a step-change in practical voice AI.
- [abus-aikorea/voice-pro](https://github.com/abus-aikorea/voice-pro) — Today +58 · Gradio WebUI for TTS, zero-shot voice cloning, Whisper transcription, YouTube download, and vocal isolation.
- [microsoft/TRELLIS.2](https://github.com/microsoft/TRELLIS.2) — Today +107 · Native and compact structured latents for 3D generation; generative AI expands beyond text/images.
- [harry0703/MoneyPrinterTurbo](https://github.com/harry0703/MoneyPrinterTurbo) — ⭐101,007 · AI-driven automated short-video generation from a topic or keyword.
- [santifer/career-ops](https://github.com/santifer/career-ops) — ⭐62,474 · Open-source AI job search agent that scans listings, scores roles, tailors CVs, and runs locally in AI coding CLIs.
- [OpenBB-finance/OpenBB](https://github.com/OpenBB-finance/OpenBB) — ⭐71,264 · Open data platform for analysts, quants, and AI agents in financial workflows.

### 🧠 LLMs / Training

- [rasbt/LLMs-from-scratch](https://github.com/rasbt/LLMs-from-scratch) — ⭐100,310 · Step-by-step implementation of a ChatGPT-like LLM in PyTorch — the reference learning repo for LLM internals.
- [open-compass/opencompass](https://github.com/open-compass/opencompass) — ⭐7,255 · Comprehensive LLM evaluation platform supporting 100+ datasets and major model families.
- [skyzh/tiny-llm](https://github.com/skyzh/tiny-llm) — ⭐4,428 · A systems-engineering course that builds a tiny vLLM + Qwen on Apple Silicon.
- [AarambhDevHub/aarambh-studio](https://github.com/AarambhDevHub/aarambh-studio) — ⭐56 · Decoder-only LLM from scratch in pure Rust using Candle — no Python/PyTorch — with MoE and sparse attention.
- [microsoft/AI-For-Beginners](https://github.com/microsoft/AI-For-Beginners) — Today +949 · 12-week, 24-lesson AI curriculum; among the fastest-gaining repos today.
- [microsoft/generative-ai-for-beginners](https://github.com/microsoft/generative-ai-for-beginners) — Today +108 · 21-lesson GenAI course covering foundations and building practices.

### 🔍 RAG / Knowledge

- [infiniflow/ragflow](https://github.com/infiniflow/ragflow) — ⭐86,574 · Leading open-source RAG engine combining retrieval-augmented generation with agent capabilities.
- [run-llama/llama_index](https://github.com/run-llama/llama_index) — ⭐51,279 · The leading document agent and OCR platform for RAG pipelines.
- [milvus-io/milvus](https://github.com/milvus-io/milvus) — ⭐45,455 · High-performance, cloud-native vector database built for scalable ANN search.
- [qdrant/qdrant](https://github.com/qdrant/qdrant) — ⭐33,712 · High-performance vector database/search engine for next-generation AI.
- [mem0ai/mem0](https://github.com/mem0ai/mem0) — ⭐62,274 · Universal memory layer for AI agents — one of the most important memory projects right now.
- [topoteretes/cognee](https://github.com/topoteretes/cognee) — ⭐29,676 · Open-source AI memory platform using self-hosted knowledge graphs for persistent agent memory.
- [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem) — ⭐89,260 · Captures agent session activity, compresses it with AI, and injects relevant context into future sessions.
- [Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify) — ⭐100,254 · Turns any codebase, docs, SQL schemas, and PDFs into a queryable knowledge graph for coding agents.

---

## 3. Trend Signal Analysis

Today's explosion of attention is around **agent skill ecosystems**, not raw models. The top trending repo, `reverse-skill`, is a skill router for security testing that extends Claude Code/Cursor/Cline — evidence that AI coding clients have become the new "runtime" for developer-facing agents. Alongside `k-skill` and `ECC`, this suggests the community is packaging specialized capabilities (security, localization, career tasks) as installable agent skills.

**Memory is rapidly becoming first-class AI infrastructure.** TencentDB-Agent-Memory in the trending list, plus `mem0`, `claude-mem`, and `cognee` in topic search, all address the same bottleneck: agents lose context. Persistent team/session memory is emerging as a foundational layer for long-horizon work.

Several new technical directions are appearing: **token compression** for agent tool outputs (`headroom`), **vectorless/reasoning-based RAG** (`PageIndex`, `Graphify`), **Rust-native LLM training** (`aarambh-studio`), and **voice-agent frameworks** (`speech-to-speech`). These are not just demos — they target performance, memory, and multimodal practical use.

This wave is connected to the rapid pace of local model releases: Ollama now lists Kimi-K2.6, GLM-5.2, MiniMax, DeepSeek, gpt-oss, Qwen, and Gemma. Cheaper and stronger local models make agent/skill/memory layers viable on commodity hardware, which in turn drives the ecosystem toward autonomous, multimodal, locally executed agent infrastructure.

---

## 4. Community Hot Spots

- **Agent skill packs and routers for AI coding CLIs** — [reverse-skill](https://github.com/zhaoxuya520/reverse-skill), [k-skill](https://github.com/NomaDamas/k-skill), [ECC](https://github.com/affaan-m/ECC). These turn editors into extensible agent platforms and are the fastest-growing direction today.
- **Memory and context layers for agents** — [TencentDB-Agent-Memory](https://github.com/TencentCloud/TencentDB-Agent-Memory), [mem0](https://github.com/mem0ai/mem0), [claude-mem](https://github.com/thedotmack/claude-mem), [cognee](https://github.com/topoteretes/cognee). Persistent memory is the key blocker for real-world agent autonomy.
- **Voice AI goes local** — [huggingface/speech-to-speech](https://github.com/huggingface/speech-to-speech), [voice-pro](https://github.com/abus-aikorea/voice-pro). Local voice agents and voice cloning are now practical with open-source models.
- **RAG modernization** — [Graphify](https://github.com/Graphify-Labs/graphify), [PageIndex](https://github.com/VectifyAI/PageIndex), [headroom](https://github.com/headroomlabs-ai/headroom). Knowledge graphs, vectorless retrieval, and token compression are redefining traditional vector DB RAG.
- **Local LLM serving and model diversity** — [ollama](https://github.com/ollama/ollama), [vllm](https://github.com/vllm-project/vllm). Rapidly expanding model families keep the local-first ecosystem hot and power all of the above.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*