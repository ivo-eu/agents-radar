# Hugging Face Trending Models Digest 2026-06-21

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-21 11:26 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-06-21

## 1. Today's Highlights

DeepSeek-V4-Pro leads the board with nearly 5,000 weekly likes and 2.6M downloads, signaling continued dominance of reasoning-focused MoE architectures. Google expands its Gemini-inspired family with **gemma-4-12B-it** (any-to-any multimodal) and **diffusiongemma-26B-A4B-it**, both climbing fast. NVIDIA’s **LocateAnything-3B** (2,227 likes) highlights the growing demand for specialized vision-language grounding models. Community GGUF fine-tunes (e.g., HauhauCS's uncensored Qwen3.6 MoE with 3.9M downloads) underscore the vibrant open‑source quantization ecosystem. Meanwhile, new entrants in streaming ASR (NVIDIA Nemotron 0.6B) and ultra‑small TTS (Inflect-Nano-v1) point to a broadening of modalities.

## 2. Trending Models by Category

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **deepseek-ai/DeepSeek-V4-Pro**  
  *Author: deepseek-ai | Likes: 4,991 | Downloads: 2,611,991*  
  A large MoE conversational model with advanced reasoning and code capabilities, trending due to its state‑of‑the‑art performance and open‑weight release.

- **zai-org/GLM-5.2**  
  *Author: zai-org | Likes: 1,736 | Downloads: 27,413*  
  A conversational MoE-DSA model from Zhipu AI, gaining traction for its efficient mixture‑of‑experts design and strong Chinese‑English bilingual performance.

- **CohereLabs/North-Mini-Code-1.0**  
  *Author: CohereLabs | Likes: 470 | Downloads: 19,551*  
  A compact MoE code‑focused LLM (Cohere2 architecture), trending as a lightweight alternative for code generation and completion.

- **microsoft/FastContext-1.0-4B-SFT**  
  *Author: microsoft | Likes: 250 | Downloads: 2,593*  
  A 4B parameter model optimized for ultra‑long context (>1M tokens) using a novel “Explorer SubAgent” attention mechanism, drawing research interest.

- **nex-agi/Nex-N2-Pro**  
  *Author: nex-agi | Likes: 340 | Downloads: 7,872*  
  A Qwen3.5‑based MoE conversational model, popular for its balance of performance and efficiency in production settings.

- **poolside/Laguna-M.1**  
  *Author: poolside | Likes: 77 | Downloads: 2,580*  
  A vLLM‑compatible code generation model from poolside, optimized for low‑latency inference with SGLang support.

### 🎨 Multimodal & Generation (Image, Video, Audio, Text-to-X)

- **google/gemma-4-12B-it**  
  *Author: google | Likes: 1,111 | Downloads: 1,815,370*  
  A flagship any‑to‑any multimodal model (text, image, audio, video) from Google, setting a new bar for unified generative AI.

- **google/diffusiongemma-26B-A4B-it**  
  *Author: google | Likes: 1,025 | Downloads: 762,861*  
  A diffusion‑based multimodal model (26B active 4B per token) that excels in image‑text understanding and conversational generation.

- **nvidia/LocateAnything-3B**  
  *Author: nvidia | Likes: 2,227 | Downloads: 241,845*  
  A vision‑language grounding model that can precisely locate objects in images given free‑form queries, trending for robotics and AR/VR applications.

- **MiniMaxAI/MiniMax-M3**  
  *Author: MiniMaxAI | Likes: 1,167 | Downloads: 104,076*  
  A multimodal VL model from MiniMax with strong image‑to‑text reasoning, gaining attention for its competitive benchmark results.

- **moonshotai/Kimi-K2.7-Code**  
  *Author: moonshotai | Likes: 938 | Downloads: 363,308*  
  A compressed‑tensor multimodal model specialized in code understanding with vision (e.g., diagram‑to‑code), popular among developers.

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
  *Author: HauhauCS | Likes: 2,057 | Downloads: 3,966,691*  
  An uncensored, aggressively fine-tuned Qwen3.6 MoE vision model (GGUF), trending for its high download volume and unrestricted outputs.

- **prefeitura-rio/Rio-3.5-Open-397B**  
  *Author: prefeitura-rio | Likes: 327 | Downloads: 190,742*  
  A massive 397B MoE multimodal model (Qwen3.5 base) fine‑tuned by a Brazilian municipality, showcasing public‑sector open‑source AI.

- **nvidia/nemotron-3.5-asr-streaming-0.6b**  
  *Author: nvidia | Likes: 597 | Downloads: 27,275*  
  A streaming automatic speech recognition model (0.6B) with cache‑aware optimization, driving real‑time ASR on edge devices.

- **bosonai/higgs-audio-v3-tts-4b**  
  *Author: bosonai | Likes: 504 | Downloads: 76,120*  
  A 4B text‑to‑speech model built on Qwen3, delivering natural‑sounding speech synthesis with multilingual support.

- **owensong/Inflect-Nano-v1**  
  *Author: owensong | Likes: 146 | Downloads: 0*  
  An ultra‑small (sub‑100M) TTS model designed for resource‑constrained environments, early but generating buzz for on‑device deployment.

- **ostris/ideogram_4_turbotime_lora**  
  *Author: ostris | Likes: 84 | Downloads: 2,452*  
  A LoRA adapter for Ideogram 4 FP8 that accelerates text‑to‑image generation, popular among diffusion enthusiasts.

- **Boogu/Boogu-Image-0.1-Edit**  
  *Author: Boogu | Likes: 75 | Downloads: 374*  
  A diffusers‑based image editing model, newly released with bilingual (EN/ZH) support and Apache‑2.0 license.

- **lordx64/Qwable-v1**  
  *Author: lordx64 | Likes: 143 | Downloads: 3,351*  
  A multimodal Qwen3.5‑based model with vision capabilities, gaining steam as a versatile open‑source alternative.

- **Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF**  
  *Author: Jackrong | Likes: 270 | Downloads: 190,993*  
  A Qwen3.6‑based vision + code model quantized to GGUF, popular for running on consumer hardware.

- **datalab-to/lift**  
  *Author: datalab-to | Likes: 97 | Downloads: 516*  
  A Qwen3.5‑based multimodal model optimized for PDF understanding and document extraction.

### 🔧 Specialized Models (Code, Math, Embeddings, Long Context)

- **WeiboAI/VibeThinker-3B**  
  *Author: WeiboAI | Likes: 532 | Downloads: 20,277*  
  A 3B math‑reasoning model fine‑tuned on Qwen2, trending for efficient step‑by‑step problem solving.

- **LiquidAI/LFM2.5-Embedding-350M**  
  *Author: LiquidAI | Likes: 83 | Downloads: 7,726*  
  A 350M embedding model (sentence‑transformers) from LiquidAI, attracting attention for high‑quality dense retrieval with small footprint.

- **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**  
  *Author: yuxinlu1 | Likes: 219 | Downloads: 21,730*  
  An agentic coding model (Gemma‑4 base) fine‑tuned for terminal‑based autonomous code execution, offered in GGUF.

- **MoonshotAI/Kimi-K2.7-Code** (also in multimodal) – listed above; its code specialization qualifies here as well.

### 📦 Fine‑Tunes & Quantizations (Community Fine‑Tunes, GGUF, AWQ)

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**  
  *Author: yuxinlu1 | Likes: 2,020 | Downloads: 358,677*  
  A community merge of Gemma‑4 with two coding fine‑tunes, quantized to GGUF – one of the most‑downloaded code models this week.

- **unsloth/GLM-5.2-GGUF**  
  *Author: unsloth | Likes: 213 | Downloads: 32,260*  
  A GGUF quantization of GLM‑5.2 optimized for CPU/edge inference, popular due to unsloth’s efficient kernel support.

- **zai-org/GLM-5.2-FP8**  
  *Author: zai-org | Likes: 117 | Downloads: 217,361*  
  Official FP8 quantization of GLM‑5.2, enabling memory‑efficient deployment on modern GPUs.

- **unsloth/Kimi-K2.7-Code-GGUF**  
  *Author: unsloth | Likes: 147 | Downloads: 42,837*  
  GGUF version of Kimi‑K2.7‑Code, making the multimodal code model accessible to consumer hardware.

- **bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF**  
  *Author: bytkim | Likes: 99 | Downloads: 36,421*  
  A pi‑tuned version of Qwen3.6‑27B with multi‑token prediction (MTP), quantized to GGUF for local experimentation.

- **Mia-AiLab/Qwable-3.6-27b**  
  *Author: Mia-AiLab | Likes: 115 | Downloads: 22,879*  
  A GGUF quantization of a Qwen3.6‑based multimodal model, filling the gap for mid‑size vision‑language models on CPU.

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** (already in multimodal, but also a fine‑tune/quantization – note its GGUF tag and uncensored nature).

- **Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF** (already in multimodal as a quantization).

## 3. Ecosystem Signal

**MoE is the new norm.** Over half of this week’s trending models use mixture‑of‑experts architectures (GLM‑5.2, DeepSeek‑V4‑Pro, Qwen3.6 MoE variants, Cohere North), balancing scale and inference cost. The success of large‑scale MoEs (Rio‑397B, DiffusionGemma‑26B A4B) suggests the community is moving beyond dense models.

**Open‑weight dominance continues.** Major labs (Google, DeepSeek, NVIDIA, Microsoft, Cohere, LiquidAI) release fully open‑weight models, often with permissive licenses (Apache‑2.0, cc‑by‑nc‑4.0). Proprietary model weight access is no longer a differentiator; instead, the race is on for dataset quality, training innovations, and specialized fine‑tuning.

**GGUF quantization is the backbone of community adoption.** Nearly 40% of the listed models are GGUF‑quantized. Community fine‑tuners (yuxinlu1, HauhauCS, bytkim) are building on flagship models (Gemma‑4, Qwen3.6, Kimi) and pushing them to consumer hardware – a testament to the LLM “modding” culture.

**Multimodal and code domains are converging.** Code models increasingly incorporate vision (Kimi‑K2.7‑Code, Qwopus‑Coder), while multimodal models gain code‑specific fine‑tunes. Agentic coding (Gemma‑4‑agentic) and long‑context (FastContext) represent two frontier areas attracting rapid iteration.

**Small models gain traction for specialized tasks.** 3B‑4B models (VibeThinker, LocateAnything, FastContext) serve high‑value niches – math, grounding, long context – without the overhead of larger MoEs.

## 4. Worth Exploring

1. **microsoft/FastContext-1.0-4B-SFT**  
   *Why:* It introduces “Explorer SubAgent” attention for handling contexts beyond 1M tokens in a 4B parameter model – a potential breakthrough for document‑scale retrieval and reasoning.

2. **nvidia/LocateAnything-3B**  
   *Why:* As a 3B vision grounding model, it enables precise object localization from natural language, making it a strong candidate for robotics, GUI automation, and visual question answering.

3. **LiquidAI/LFM2.5-Embedding-350M**  
   *Why:* Its small footprint (350M) paired with state‑of‑the‑art embedding quality suggests it could become the go‑to model for semantic search and RAG pipelines, especially on‑device.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*