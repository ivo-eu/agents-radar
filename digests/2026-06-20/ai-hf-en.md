# Hugging Face Trending Models Digest 2026-06-20

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-20 10:17 UTC

---

# Hugging Face Trending Models Digest — 2026-06-20

## Today's Highlights

The biggest release this week is **DeepSeek-V4-Pro**, which dominates with nearly 5,000 weekly likes and 2.8 million downloads, signaling continued strong demand for frontier open-weight reasoning models. The **Qwen3.6 family** is experiencing explosive community activity, particularly through aggressive fine-tunes and GGUF quantizations — the top-downloaded model overall is a Qwen3.6 uncensored variant with 3.8 million downloads. Multimodal models remain the dominant pipeline type in the top 30, with image-text-to-text models from Google, NVIDIA, and MiniMax capturing significant attention. The ecosystem also shows a clear shift toward **Mixture-of-Experts (MoE)** architectures, with models like GLM-5.2, Qwen3.6 MoE variants, and Cohere’s North-Mini-Code appearing prominently. Finally, **Gemma-4** continues to generate a large fine-tuning and quantization ecosystem, anchored by Google’s official Gemma-4-12B-it model and multiple community derivatives.

---

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — Author: deepseek-ai | Likes: 4,973 | Downloads: 2,797,050  
  *The week's most-liked model, a state-of-the-art conversational LLM driving massive adoption in the open-weight space.*

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — Author: zai-org | Likes: 1,577 | Downloads: 19,683  
  *A new MoE-driven conversational model from Zhipu AI, gaining traction as a strong open-weight alternative with DSA (dynamic sparse attention).*

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — Author: WeiboAI | Likes: 477 | Downloads: 16,270  
  *A compact 3B Qwen2-based reasoning model optimized for math, trending for its strong performance-per-parameter ratio.*

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — Author: microsoft | Likes: 235 | Downloads: 1,998  
  *Microsoft's new 4B model designed for efficient long-context processing, notable for its Explorer SubAgent architecture.*

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — Author: CohereLabs | Likes: 461 | Downloads: 18,783  
  *A MoE-based code-focused model from Cohere, competing in the specialized coding LLM space with strong community interest.*

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — Author: nex-agi | Likes: 337 | Downloads: 7,724  
  *A Qwen3.5 MoE-based multimodal-capable text generation model gaining attention for its blend of efficiency and versatility.*

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — Author: OBLITERATUS | Likes: 358 | Downloads: 110,450  
  *An aggressively fine-tuned community variant of Gemma-4-12B-it, popular among users seeking uncensored or stylized outputs.*

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — Author: MiniMaxAI | Likes: 1,141 | Downloads: 85,771  
  *A cutting-edge vision-language model from MiniMax, trending for its strong multimodal understanding and competitive open-weight availability.*

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — Author: moonshotai | Likes: 915 | Downloads: 317,963  
  *Moonshot AI's code-focused multimodal model with compressed-tensor optimization, seeing high download volume for vision-coding tasks.*

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — Author: google | Likes: 1,015 | Downloads: 673,464  
  *A 26B MoE diffusion language model from Google with 4B active parameters, combining diffusion-based generation with conversational capabilities.*

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — Author: nvidia | Likes: 2,203 | Downloads: 235,606  
  *NVIDIA's 3B image understanding model for object localisation and feature extraction, extremely well-liked for its targeted utility.*

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — Author: nvidia | Likes: 568 | Downloads: 21,426  
  *NVIDIA's lightweight streaming ASR model (0.6B) using cache-aware architecture, trending as a compact speech recognition solution.*

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — Author: google | Likes: 1,100 | Downloads: 1,696,240  
  *Google's flagship any-to-any model with unified architecture handling text, image, and multimodal inputs, one of the most downloaded models this week.*

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — Author: zai-org | Likes: 235 | Downloads: 0  
  *A new pose-driven character animation diffusion model for image-to-video generation, notable for its zero-download launch with strong early likes.*

- **[Zyphra/ZONOS2](https://huggingface.co/Zyphra/ZONOS2)** — Author: Zyphra | Likes: 116 | Downloads: 763  
  *An Apache-2.0 licensed text-to-speech model gaining attention as a fully open TTS alternative.*

### 🔧 Specialized Models (code, math, medical, embeddings)

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — Author: yuxinlu1 | Likes: 1,887 | Downloads: 312,332  
  *A Gemma-4-based coding model fine-tuned with the Fable5 and Composer frameworks, extremely popular as a GGUF quant for local code generation.*

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** — Author: yuxinlu1 | Likes: 126 | Downloads: 6,307  
  *An agentic-focused variant of the Gemma-4 coder, optimized for terminal and autonomous agent workflows.*

- **[owensong/Inflect-Nano-v1](https://huggingface.co/owensong/Inflect-Nano-v1)** — Author: owensong | Likes: 129 | Downloads: 0  
  *An ultra-small text-to-speech model (Nano class) for low-resource speech synthesis, trending as a size-efficiency benchmark.*

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — Author: bosonai | Likes: 493 | Downloads: 72,225  
  *A 4B multimodal TTS model built on Qwen3, offering high-quality speech synthesis with multimodal grounding.*

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — Author: HauhauCS | Likes: 2,013 | Downloads: 3,812,636  
  *The most downloaded model this week — an uncensored, aggressively fine-tuned Qwen3.6 MoE variant in GGUF format for local deployment.*

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — Author: unsloth | Likes: 193 | Downloads: 22,586  
  *Unsloth's GGUF quantization of GLM-5.2, enabling efficient local inference of this MoE conversational model.*

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — Author: Jackrong | Likes: 263 | Downloads: 168,502  
  *A code-focused GGUF fine-tune of Qwen3.6 with Multi-Turn Prediction (MTP), popular among local coding users.*

- **[lordx64/Qwable-v1](https://huggingface.co/lordx64/Qwable-v1)** — Author: lordx64 | Likes: 133 | Downloads: 2,769  
  *A community fine-tune of Qwen3.5 MoE, demonstrating the rapid iteration cycle around the Qwen family.*

- **[unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF)** — Author: unsloth | Likes: 143 | Downloads: 37,260  
  *Unsloth's GGUF version of the Kimi-K2.7-Code model, providing compressed-tensor quantizations for efficient local code assistance.*

- **[zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8)** — Author: zai-org | Likes: 108 | Downloads: 138,174  
  *The official FP8 quantization of GLM-5.2, enabling efficient deployment on modern hardware.*

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — Author: unsloth | Likes: 318 | Downloads: 216,396  
  *Unsloth's GGUF quantization of Google's diffusiongemma-26B, making this large MoE model more accessible for local use.*

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** — Author: DavidAU | Likes: 406 | Downloads: 587,521  
  *An extreme community fine-tune of Qwen3.6 combining multiple style merges and uncensored thinking, demonstrating the "frankenmodel" trend.*

- **[bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF)** — Author: bytkim | Likes: 89 | Downloads: 20,465  
  *A GGUF fine-tune of Qwen3.6 with Multi-Turn Prediction and Pi-tune optimizations for improved reasoning chains.*

- **[Mia-AiLab/Qwable-3.6-27b](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b)** — Author: Mia-AiLab | Likes: 108 | Downloads: 17,311  
  *A GGUF community variant of Qwen3.6-27B, part of the rapid fine-tuning ecosystem around this model family.*

---

## Ecosystem Signal

The ecosystem this week is defined by a clear **convergence on MoE architectures** as the dominant paradigm for both frontier and community models. Models like GLM-5.2, Qwen3.6, Cohere's North-Mini-Code, and Google's diffusiongemma all use MoE routing, suggesting the field has decisively moved beyond dense transformers for high-efficiency deployment.

**Qwen3.6 has emerged as the most active base model for community experimentation**, outpacing even Gemma-4 in fine-tune and quantization volume. The top-downloaded model is a Qwen3.6 uncensored variant, and at least six Qwen3.6 derivatives appear in the top 30. This indicates Qwen3.6 offers a favorable combination of performance, open license, and quantization-friendliness for community builders.

**Multimodal models now dominate the trending list** — 16 of 30 models are image-text-to-text or any-to-any. This reflects a market shift where pure text models are increasingly commoditized, and differentiation comes from multimodal capabilities.

**The GGUF quantization ecosystem continues to expand aggressively**, with Unsloth leading the charge across GLM-5.2, Kimi-K2.7-Code, and diffusiongemma. Community fine-tunes are becoming more extreme — names include "Heretic," "Obliterated," and "Aggressive" — reflecting a growing subculture of users seeking unfiltered or highly stylized model behavior.

**Open-weight models are winning decisively** over proprietary alternatives in community visibility. DeepSeek-V4-Pro, GLM-5.2, Gemma-4, and Qwen3.6 all demonstrate that open release strategies drive ecosystem adoption and developer mindshare.

---

## Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — With the second-highest likes count (2,203) and a focused 3B parameter size, this model represents a sweet spot for efficient visual grounding tasks. It's worth studying as a template for how specialized, task-specific models can achieve outsized community impact compared to general-purpose alternatives.

2. **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — This is a genuinely novel architectural approach: a diffusion-based language model with MoE sparsity. With 673k downloads and 1k+ likes, it's worth exploring as an indicator of where Google is pushing the generation paradigm beyond standard autoregressive decoding.

3. **[Zyphra/ZONOS2](https://huggingface.co/Zyphra/ZONOS2)** — While it has relatively low numbers, its Apache-2.0 license and fully open TTS positioning make it a potential inflection point for open-weight speech synthesis. With the new TTS models from bosonai and owensong also trending, ZONOS2 is worth monitoring as part of the open TTS resurgence.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*