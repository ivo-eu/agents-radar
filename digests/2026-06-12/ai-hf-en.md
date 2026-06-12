# Hugging Face Trending Models Digest 2026-06-12

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-12 06:22 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-06-12

## 1. Today's Highlights

The Hugging Face hub is dominated by two major trends: **ultra-large MoE language models** and **multimodal any-to-any architectures**. DeepSeek-V4-Pro leads by a landslide with 4,785 weekly likes and 4M downloads, cementing the open‑weight frontier for 600B+ scale. Google’s Gemma‑4 family (12B, 26B-A4B) and its diffusion variant *diffusiongemma* continue their rapid community adoption, with quantized GGUF versions from `unsloth` seeing massive download numbers. Nvidia pushes both specialized multimodal (LocateAnything‑3B) and massive text‑generation (Nemotron‑3 Ultra 550B‑A55B) models, while fine‑tune and “abliterated” variants (e.g., Qwen3.6‑35B‑A3B uncensored, Gemma‑4‑12B‑OBLITERATED) signal a strong demand for aligned and unaligned community releases.

## 2. Trending Models by Category

### 🧠 Language Models (LLMs, chat, instruction‑tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — Author: deepseek-ai | Likes: 4,785 | Downloads: 4,061,006  
  A massive open‑weight MoE model pushing the frontier of conversational AI, trending as the most liked model this week.

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — Author: nex-agi | Likes: 210 | Downloads: 1,185  
  A Qwen‑3.5‑based MoE model optimized for general text generation and reasoning.

- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** — Author: nex-agi | Likes: 166 | Downloads: 1,222  
  The smaller sibling of Nex-N2-Pro, offering efficient performance for lightweight deployments.

- **[sapientinc/HRM-Text-1B](https://huggingface.co/sapientinc/HRM-Text-1B)** — Author: sapientinc | Likes: 751 | Downloads: 134,752  
  A 1B‑parameter text‑generation model gaining traction for resource‑constrained applications.

- **[XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash](https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash)** — Author: XiaomiMiMo | Likes: 88 | Downloads: 660  
  Xiaomi’s latest MiMo agent‑oriented model with FP4 quantization for efficiency.

- **[nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16)** — Author: nvidia | Likes: 198 | Downloads: 59,066  
  A 550B‑parameter MoE model from Nvidia, targeting high‑end reasoning and enterprise use.

- **[nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4)** — Author: nvidia | Likes: 169 | Downloads: 91,117  
  The NVFP4‑quantized version of the Nemotron‑3 Ultra, balancing precision and memory footprint.

### 🎨 Multimodal & Generation (image, video, audio, text‑to‑X)

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — Author: nvidia | Likes: 1,886 | Downloads: 131,794  
  A 3B image‑text‑to‑text model for object localization and visual reasoning, trending due to its general‑purpose utility.

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — Author: google | Likes: 517 | Downloads: 0  
  Google’s first diffusion‑based vision‑language model (26B‑A4B MoE), merging generation with understanding.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — Author: google | Likes: 943 | Downloads: 675,936  
  The instruction‑tuned flagship of the Gemma‑4 family, supporting any‑to‑any multimodal tasks.

- **[google/gemma-4-12B](https://huggingface.co/google/gemma-4-12B)** — Author: google | Likes: 519 | Downloads: 140,221  
  The base version of Gemma‑4‑12B, widely adopted for fine‑tuning and customization.

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — Author: ideogram-ai | Likes: 489 | Downloads: 7,170  
  State‑of‑the‑art text‑to‑image model in FP8 quantization, gaining popularity for high‑quality generation.

- **[ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4)** — Author: ideogram-ai | Likes: 319 | Downloads: 6,124  
  The NF4‑quantized variant of Ideogram‑4, offering a smaller footprint while maintaining quality.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — Author: HauhauCS | Likes: 1,685 | Downloads: 3,057,541  
  A highly controversial uncensored MoE vision‑language model, trending due to massive community interest and downloads.

- **[stepfun-ai/Step-3.7-Flash](https://huggingface.co/stepfun-ai/Step-3.7-Flash)** — Author: stepfun-ai | Likes: 368 | Downloads: 50,187  
  A fast vision‑language model optimized for real‑time multimodal reasoning.

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — Author: bosonai | Likes: 366 | Downloads: 19,948  
  A 4B text‑to‑speech model built on Qwen‑3 multimodal backbone, trending for high‑quality voice synthesis.

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — Author: nvidia | Likes: 374 | Downloads: 4,965  
  A lightweight streaming ASR model with cache‑aware architecture, ideal for real‑time speech recognition.

- **[google/magenta-realtime-2](https://huggingface.co/google/magenta-realtime-2)** — Author: google | Likes: 180 | Downloads: 19,806  
  A text‑to‑audio model for real‑time music and sound generation, backed by recent Google research.

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — Author: zai-org | Likes: 121 | Downloads: 0  
  A pose‑driven image‑to‑video model for character animation, showcasing advances in controlled video generation.

- **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — Author: ByteDance | Likes: 224 | Downloads: 305  
  A text‑to‑video renderer with strong visual consistency, released under Apache 2.0.

- **[MisoLabs/MisoTTS](https://huggingface.co/MisoLabs/MisoTTS)** — Author: MisoLabs | Likes: 195 | Downloads: 0  
  A new TTS model focused on natural speech synthesis, still in early adoption.

- **[huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated)** — Author: huihui-ai | Likes: 144 | Downloads: 6,400  
  An “abliterated” (uncensored) fine‑tune of Gemma‑4‑12B‑it, catering to the uncensoring trend.

### 🔧 Specialized Models (code, math, medical, embeddings)

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — Author: CohereLabs | Likes: 310 | Downloads: 1,859  
  A compact MoE model specialized for code generation and conversational coding, using Cohere’s second‑generation architecture.

### 📦 Fine‑tunes & Quantizations (community fine‑tunes, GGUF, AWQ)

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — Author: unsloth | Likes: 564 | Downloads: 711,706  
  The most popular quantization of Gemma‑4‑12B‑it, enabling local CPU inference with unmodified quality.

- **[unsloth/gemma-4-12B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF)** — Author: unsloth | Likes: 202 | Downloads: 148,252  
  A QAT‑trained GGUF variant for even better quantization accuracy.

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — Author: unsloth | Likes: 185 | Downloads: 0  
  The first GGUF quantization of diffusiongemma, making diffusion‑based VLMs accessible.

- **[unsloth/gemma-4-26B-A4B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF)** — Author: unsloth | Likes: 144 | Downloads: 129,110  
  QAT GGUF of the larger Gemma‑4 26B MoE, popular for balancing quality and size.

- **[google/gemma-4-12B-it-qat-q4_0-gguf](https://huggingface.co/google/gemma-4-12B-it-qat-q4_0-gguf)** — Author: google | Likes: 130 | Downloads: 96,749  
  Official Google‑released 4‑bit GGUF quantization of Gemma‑4‑12B‑it, ensuring trust and performance.

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — Author: OBLITERATUS | Likes: 237 | Downloads: 14,838  
  A heavily modified/deprecated fine‑tune of Gemma‑4, popular among the “abliterated” community.

- **[Comfy-Org/Ideogram-4](https://huggingface.co/Comfy-Org/Ideogram-4)** — Author: Comfy-Org | Likes: 136 | Downloads: 0  
  A ComfyUI wrapper for Ideogram‑4, enabling seamless integration into diffusion pipelines.

## 3. Ecosystem Signal

The current landscape reveals several structural shifts:

- **MoE dominance**: Almost every large model this week uses Mixture‑of‑Experts (e.g., DeepSeek‑V4, Gemma‑4‑26B‑A4B, Nemotron‑3 Ultra 550B‑A55B, Qwen3.6‑35B‑A3B). MoE allows scaling to 550B+ parameters while keeping inference costs manageable, and the community is adopting MoE quantizations (GGUF) at scale.

- **Open‑weight advancement**: DeepSeek‑V4‑Pro and Nvidia’s Nemotron‑3 Ultra (both 500B+ MoE) are fully open‑weight, signaling a move away from gated releases. Google’s Gemma‑4 family also remains open, while diffusiongemma represents a new diffusion‑VLM approach.

- **Quantization as a first‑class citizen**: `unsloth` continues to lead GGUF conversions, making large models runnable on consumer hardware. The sheer download numbers (711k for Gemma‑4‑12B‑GGUF) show that inference efficiency is the primary barrier to adoption.

- **Fine‑tuning for alignment and unalignment**: “Abliterated” (uncensored) fine‑tunes of Gemma‑4 and Qwen3.6 are highly active, reflecting ongoing demand for models without safety guardrails. Simultaneously, task‑specific fine‑tunes like HRM‑Text‑1B and North‑Mini‑Code target professional niches.

- **Multimodal convergence**: The line between LLMs and vision/audio models is blurring. Gemma‑4‑it, Qwen3.6, and diffusiongemma all handle any‑to‑any modalities. Text‑to‑audio (magenta, Higgs) and image‑to‑video (SCAIL‑2, Bernini‑R) are also maturing quickly.

## 4. Worth Exploring

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — The most liked model this week represents the current state‑of‑the‑art in open‑weight LLMs. Its massive scale (MoE) and strong benchmark performance make it essential for anyone pushing the frontier of conversational AI or agent systems.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — With 1,886 likes and 131k downloads, this 3B multimodal model punches above its weight for visual grounding and object localization. It’s a great entry point for vision‑language applications without requiring massive compute.

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — An experimental fusion of diffusion models with LLM reasoning. Though still fresh (0 downloads from GGUF), the architecture is innovative and could define a new class of generative VLMs. Worth studying for researchers and early adopters.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*