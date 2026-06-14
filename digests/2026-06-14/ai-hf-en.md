# Hugging Face Trending Models Digest 2026-06-14

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-14 11:10 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-06-14

## 1. Today's Highlights

**DeepSeek-V4-Pro** dominates the leaderboard with 4,818 weekly likes and over 3M downloads, cementing DeepSeek's position as a primary open-weight LLM contender. The **Gemma 4 family** from Google continues its explosive ecosystem growth — the base 12B instruct model alone has crossed 1M downloads, while community fine-tunes, abliterated variants, and GGUF quantizations (led by **unsloth**) proliferate rapidly. A notable shift is the rise of **DiffusionGemma**, a 26B parameter MoE model combining diffusion and language modeling for image-text-to-text tasks, signaling convergence between generative and language architectures. The uncensored and "abliterated" fine-tune category is also surging, with models like *Qwen3.6-35B-A3B-Uncensored* and *supergemma4-26b-uncensored* collectively generating millions of downloads.

---

## 2. Trending Models by Category

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — deepseek-ai | ❤️ 4,818 | 📥 3,075,369  
  The week's most-liked model: a next-generation conversational LLM from DeepSeek, setting new benchmarks for open-weight reasoning and chat quality.

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — CohereLabs | ❤️ 361 | 📥 9,932  
  A compact MoE code-generation model from Cohere, leveraging the cohere2_moe architecture for efficient instruction-following and programming tasks.

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — nex-agi | ❤️ 242 | 📥 3,396  
  A Qwen3.5-MoE-based text-generation model optimized for reasoning and agentic workflows, trending for its strong performance-per-parameter ratio.

- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** — nex-agi | ❤️ 198 | 📥 7,010  
  The smaller sibling of Nex-N2-Pro, offering a lightweight yet capable MoE language model for edge deployment.

- **[prefeitura-rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)** — prefeitura-rio | ❤️ 184 | 📥 112,371  
  A massive 397B MoE model based on Qwen3.5, released under an open license by the municipality of Rio de Janeiro — notable for its institutional backing and public-sector AI initiative.

- **[XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash](https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash)** — XiaomiMiMo | ❤️ 108 | 📥 4,108  
  Xiaomi's latest MiMo variant featuring FP4 quantization and DFlash optimization, targeting on-device agent use cases.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — google | ❤️ 738 | 📥 198,912  
  Google's pioneering diffusion-language hybrid: a 26B MoE model that processes image-text-to-text via a diffusion backbone, blending generative image understanding with conversational capabilities.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — moonshotai | ❤️ 562 | 📥 15,145  
  A code-specialized vision-language model from Moonshot AI, trained with compressed-tensor techniques for efficient multimodal code reasoning.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — nvidia | ❤️ 1,977 | 📥 75,201  
  NVIDIA's compact image-feature extraction model for general object localization — trending for its versatility in visual grounding tasks.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — MiniMaxAI | ❤️ 444 | 📥 6,643  
  MiniMax's third-generation multimodal VL model, offering strong vision-language understanding in a compact package.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — google | ❤️ 1,001 | 📥 1,084,405  
  The most-downloaded Gemma 4 instruct variant, supporting any-to-any modality (text, image, interleaved), widely adopted as a base for fine-tuning and quantization.

- **[google/gemma-4-12B](https://huggingface.co/google/gemma-4-12B)** — google | ❤️ 537 | 📥 213,502  
  The base (non-instruct) Gemma 4 12B model, popular for continued pre-training and domain adaptation.

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — ideogram-ai | ❤️ 525 | 📥 8,263  
  FP8 quantized version of Ideogram 4, a leading text-to-image diffusion model, enabling faster inference with minimal quality loss.

- **[ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4)** — ideogram-ai | ❤️ 334 | 📥 3,763  
  NF4 variant of Ideogram 4, offering even more aggressive quantization for memory-constrained image generation.

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — zai-org | ❤️ 164 | 📥 0  
  A character animation model for pose-driven image-to-video generation, leveraging diffusion for controllable video synthesis.

- **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — ByteDance | ❤️ 235 | 📥 459  
  ByteDance's image-text-to-video renderer (accompanied by a paper on arXiv), capable of generating temporally consistent video from multimodal inputs.

### 🔧 Specialized Models (code, speech, ASR, TTS)

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — nvidia | ❤️ 406 | 📥 4,505  
  NVIDIA's compact streaming ASR model (600M parameters) with cache-aware architecture for real-time speech recognition, trending for low-latency deployment.

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — bosonai | ❤️ 418 | 📥 35,122  
  A 4B parameter text-to-speech model built on the Higgs multimodal Qwen3 backbone, offering high-quality voice synthesis.

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — unsloth | ❤️ 253 | 📥 80,118  
  Community GGUF quantization of DiffusionGemma by unsloth, enabling local inference on consumer hardware.

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — OBLITERATUS | ❤️ 287 | 📥 60,949  
  An "obliterated" (abliterated) variant of Gemma 4-12B, removing safety alignment for unrestricted use cases.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — HauhauCS | ❤️ 1,775 | 📥 2,516,709  
  A heavily modified Qwen3.6 MoE fine-tune with aggressive uncensoring and behavioral modifications — remarkable for 2.5M+ downloads despite its niche appeal.

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — unsloth | ❤️ 584 | 📥 926,372  
  The most popular Gemma 4 GGUF quantization, enabling efficient local deployment on CPU and GPU via llama.cpp.

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — Jackrong | ❤️ 168 | 📥 33,720  
  A quantized Qwen3.6 code model with multi-token prediction (MTP) capabilities, optimized for Code LLM inference.

- **[huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated)** — huihui-ai | ❤️ 152 | 📥 8,717  
  An abliterated (de-aligned) version of Gemma 4-12B-it, part of the growing uncensored fine-tune ecosystem.

- **[Jiunsong/supergemma4-26b-uncensored-gguf-v2](https://huggingface.co/Jiunsong/supergemma4-26b-uncensored-gguf-v2)** — Jiunsong | ❤️ 822 | 📥 96,437  
  A GGUF-quantized, uncensored fine-tune of the 26B Gemma 4 model, noted for "fast" inference claims.

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** — DavidAU | ❤️ 327 | 📥 375,966  
  An extreme community fine-tune combining Qwen3.6 with Claude-style thinking prompts, uncensored, with IMatrix quantization — exemplifies the "maximalist" fine-tuning trend.

- **[unsloth/gemma-4-26B-A4B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF)** — unsloth | ❤️ 152 | 📥 317,261  
  QAT (quantization-aware training) GGUF of Gemma 4-26B, demonstrating advanced quantization methodology from the unsloth team.

---

## 3. Ecosystem Signal

**Mixture-of-Experts (MoE)** has become the de facto architecture for state-of-the-art open models — virtually every new flagship entry (DeepSeek-V4-Pro, DiffusionGemma, Gemma 4, Qwen3.6 variants, Rio-3.5-Open) uses MoE gating to achieve high parameter counts with manageable compute. The **Gemma 4 family** is the week's most dynamic ecosystem: Google released the base and instruct models, while community tooling (unsloth GGUF quantizations), fine-tunes (huihui-ai's abliterated, OBLITERATUS's uncensored, Jiunsong's supergemma), and derivative works have already generated over 3M combined downloads. This signals a healthy open-weight ecosystem where a single model family can spawn dozens of community adaptations within days.

The **uncensored / abliterated** sub-niche continues to grow aggressively. Models like *HauhauCS/Qwen3.6-35B-A3B-Uncensored* and *supergemma4-26b-uncensored* are achieving viral download numbers (2.5M+ and 96K+ respectively), indicating strong demand for models with reduced safety guardrails — likely driven by creative writing, roleplay, and local experimentation use cases. Meanwhile, **quantization specialization** is accelerating: unsloth has emerged as the dominant GGUF provider across Gemma 4 and DiffusionGemma, while fp8 and nf4 formats from Ideogram and FP4 from Xiaomi highlight a growing appetite for memory-efficient deployment. Interestingly, the **image-text-to-text** pipeline dominates 30% of the top-30 list, reflecting an industry-wide pivot toward multimodal understanding as a baseline expectation for new models.

---

## 4. Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — With 1,977 likes (second-most-liked this week) and only 3B parameters, this model punches far above its weight. It fills a practical niche for object localization without needing a full VLM, making it ideal for robotics, document parsing, and image analysis pipelines. Its small size and NVIDIA engineering pedigree make it a strong candidate for production integration.

2. **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — A genuinely novel architecture that marries diffusion models with language model transformers for multimodal understanding. While its 198K downloads show strong adoptions, its "image-text-to-text" pipeline is still rare — studying its behavior and fine-tuning potential could yield insights into the next generation of hybrid generative models.

3. **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — Image-to-video generation remains one of the most challenging and demanded capabilities. Bernini-R, backed by an arXiv paper (2605.22344), represents state-of-the-art in controllable video synthesis from multimodal inputs. For researchers and developers interested in video generation, this is a must-study reference implementation.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*