# Hugging Face Trending Models Digest 2026-06-23

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-22 17:18 UTC

---

# Hugging Face Trending Models Digest — 2026-06-23

## Today's Highlights

DeepSeek dominates this week with **DeepSeek-V4-Pro** amassing over 5,000 weekly likes and 2.4M downloads, signaling strong community appetite for next-generation MoE architectures. NVIDIA's **LocateAnything-3B** (2,283 likes) is a standout multimodal entry, offering precise visual grounding in a compact 3B parameter package. Google is doubling down on unified any-to-any models with **gemma-4-12B-it** (1,134 likes) and the diffusion-transformer hybrid **diffusiongemma-26B-A4B-it** (1,046 likes), reflecting a clear industry push toward cross-modal reasoning. Meanwhile, the GGUF ecosystem is thriving: community quantizations of GLM-5.2, Qwen3.6, and Gemma-4 code variants are driving massive download counts, indicating strong grassroots demand for locally deployable models.

---

## Trending Models by Category

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — deepseek-ai · 5,006 likes · 2.4M downloads  
  DeepSeek's flagship MoE model with strong conversational and reasoning capabilities, leading the week in both community engagement and practical adoption.

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — zai-org · 1,973 likes · 33.6K downloads  
  Zhipu AI's latest MoE chat model with DSA optimization, gaining traction as a top open-weight alternative in the Chinese AI ecosystem.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — microsoft · 284 likes · 3.5K downloads  
  A compact 4B parameter model built on Qwen3, optimized for efficient long-context reasoning with sub-agent capabilities.

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — CohereLabs · 478 likes · 21.1K downloads  
  Cohere's MoE code-focused model, offering strong conversational code generation in a mid-size architecture.

- **[poolside/Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1)** — poolside · 90 likes · 2.7K downloads  
  A vLLM-compatible code generation model from poolside, optimized for deployment with sglang and production inference pipelines.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — nvidia · 2,283 likes · 247.5K downloads  
  NVIDIA's compact vision-language model for visual grounding and object localization, excelling at referring expression comprehension with only 3B parameters.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — MiniMaxAI · 1,198 likes · 120K downloads  
  A multimodal vision-language model from MiniMax, balancing strong image-text understanding with efficient inference.

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — google · 1,046 likes · 874.4K downloads  
  Google's hybrid diffusion-transformer model for image generation and understanding, combining Gemma's language backbone with diffusion decoding.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — google · 1,134 likes · 1.9M downloads  
  A unified any-to-any model (text, image, audio, video) built on the Gemma 4 architecture, Google's most versatile open-weight release.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — moonshotai · 957 likes · 412.8K downloads  
  A compressed, image-text-to-text model optimized for code understanding with visual context, from the Moonshot AI team.

- **[ostris/ideogram_4_turbotime_lora](https://huggingface.co/ostris/ideogram_4_turbotime_lora)** — ostris · 100 likes · 3.2K downloads  
  A LoRA adapter for Ideogram 4 FP8, enabling faster image generation with style transfer capabilities.

- **[Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit)** — Boogu · 98 likes · 473 downloads  
  An early-stage image editing diffusion model with bilingual (EN/ZH) support, licensed under Apache 2.0.

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — nvidia · 624 likes · 34.9K downloads  
  NVIDIA's streaming automatic speech recognition model with cache-aware architecture, optimized for real-time transcription.

### 🔧 Specialized Models (code, math, medical, embeddings)

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — yuxinlu1 · 2,147 likes · 414.7K downloads  
  A highly popular GGUF-quantized Gemma-4 code variant with enhanced reasoning, driving massive downloads for local coding assistants.

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — WeiboAI · 599 likes · 32.4K downloads  
  A 3B math reasoning model based on Qwen2, optimized for step-by-step problem solving with a "vibe" chain-of-thought approach.

- **[LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)** — LiquidAI · 100 likes · 8.8K downloads  
  A compact 350M parameter sentence embedding model from Liquid AI, ideal for retrieval and semantic search tasks.

- **[LiquidAI/LFM2.5-ColBERT-350M](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M)** — LiquidAI · 76 likes · 2.2K downloads  
  A ColBERT-style late interaction embedding model for efficient multi-vector retrieval, built on the LFM2.5 architecture.

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** — baidu · 102 likes · 47 downloads  
  Baidu's vision-language model for unlimited OCR scenarios, targeting document understanding and text extraction across diverse formats.

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — unsloth · 243 likes · 41.8K downloads  
  Unsloth's GGUF quantization of GLM-5.2, enabling efficient local inference of Zhipu's MoE model.

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** — yuxinlu1 · 354 likes · 50.3K downloads  
  A GGUF-quantized agentic variant of Gemma-4, fine-tuned for terminal and tool-use tasks with scaling factor adjustments.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — HauhauCS · 2,101 likes · 4.1M downloads  
  A massive GGUF-quantized MoE vision model based on Qwen3.6, uncensored and aggressively fine-tuned for maximum creativity, with by far the most downloads this week.

- **[zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8)** — zai-org · 131 likes · 334.7K downloads  
  Official FP8 quantization of GLM-5.2 for memory-efficient deployment, maintaining quality while reducing footprint.

- **[lordx64/Qwable-v1](https://huggingface.co/lordx64/Qwable-v1)** — lordx64 · 160 likes · 3.7K downloads  
  A Qwen3.5 MoE fine-tune with image-text-to-text capabilities, showcasing community experimentation with Qwen vision models.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** — empero-ai · 118 likes · 6.6K downloads  
  A GGUF-quantized reasoning model built on Qwen3.5, inspired by Claude-style mythos storytelling and logic.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** — empero-ai · 110 likes · 842 downloads  
  The non-quantized base version of the Qwythos reasoning fine-tune, using safetensors for direct Transformers integration.

- **[bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF)** — bytkim · 104 likes · 52.8K downloads  
  A GGUF-quantized Qwen3.6 variant with multi-token prediction (MTP) fine-tuning, optimized for clustered inference.

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — Jackrong · 280 likes · 214.6K downloads  
  A popular vision-language coder quantization based on Qwen3.6, combining coding ability with image input for multimodal programming.

- **[Mia-AiLab/Qwable-3.6-27b](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b)** — Mia-AiLab · 124 likes · 24K downloads  
  A GGUF quantization of a Qwen3.6 fine-tune, representing the growing ecosystem of community-level Qwen adaptations.

- **[owensong/Inflect-Nano-v1](https://huggingface.co/owensong/Inflect-Nano-v1)** — owensong · 165 likes · 0 downloads  
  An ultra-small text-to-speech model (likely sub-100M params) from the training-only release, notable for its extreme efficiency target.

---

## Ecosystem Signal

**Model Family Momentum**: The Qwen ecosystem (3.5, 3.6, MoE variants) is showing explosive community activity — nearly a third of the trending list are Qwen derivatives. Gemma-4 is the second most active family, with Google releasing official checkpoints while the community rapidly builds GGUF quantizations and domain fine-tunes (coding, agentic). DeepSeek-V4-Pro represents the "superstar" open-weight release, dominating likes and downloads as the largest MoE model in the list.

**Open-Weight vs Proprietary**: The trend is overwhelmingly open-weight. All 30 models are either fully open (Apache 2.0 or permissive) or available as community fine-tunes. No strictly proprietary API-only models appear. This suggests the HuggingFace community values inspectability and local deployment, especially for coding and multimodal tasks.

**Quantization & Fine-Tuning Activity**: GGUF quantizations dominate downloads — community quantizers (unsloth, yuxinlu1, various individuals) are riding the wave of new base releases. The "uncensored" and "aggressive" fine-tune category (e.g., HauhauCS's 4M-download Qwen3.6 variant) signals ongoing demand for less filtered models. Multi-token prediction (MTP) fine-tunes are an emerging technical niche (bytkim, Jackrong), suggesting researchers are exploring alternative training objectives.

**Notable Gaps**: Speech and audio remain underrepresented (only one ASR and one TTS model). Embedding models are niche but present (LiquidAI's two entries). No dedicated video generation or 3D models appear in the top 30, indicating those modalities may be trending separately or slower in the open-weight space.

---

## Worth Exploring

**1. [google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**  
This is a genuinely novel architecture — a diffusion-process decoder fused with a language model backbone. For researchers interested in generative AI architectures beyond autoregressive transformers, this model represents Google's bet on hybrid diffusion-LLM generation. Its 874K downloads suggest the community is actively testing it for image generation and multimodal tasks.

**2. [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
At only 3B parameters, this model achieves strong visual grounding performance, making it one of the most practical multimodal models for edge deployment. Its combination of high likes (2,283) and downloads (247K) indicates both enthusiasm and real-world usage. For developers building vision-based agents, OCR pipelines, or robotics applications, the small footprint and strong localization capability make this a must-test.

**3. [deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
The undisputed top model of the week with 5,006 likes. DeepSeek's MoE architecture has consistently pushed open-weight boundaries. V4-Pro likely brings improvements in reasoning, context handling, and efficiency. For anyone evaluating state-of-the-art open LLMs for production or research, this is the baseline to beat. Its 2.4M downloads also suggest strong ecosystem support and community validation.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*