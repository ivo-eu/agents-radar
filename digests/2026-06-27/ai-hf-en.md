# Hugging Face Trending Models Digest 2026-06-27

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-27 09:15 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-06-27

## 1. Today’s Highlights

This week’s trend is dominated by two forces: **MoE‑based architectures** and **aggressive quantization**. The top‑liked model is zai-org/GLM-5.2 (2,624 likes), a MoE language model from Baidu’s GLM family, while NVIDIA’s LocateAnything‑3B (2,391 likes) proves that specialist vision models can compete with generalist LLMs. Community fine‑tunes of Gemma‑4‑12B and Qwen3.6 continue to surge, with HauhauCS’s uncensored Qwen3.6 variant amassing over 3.3M downloads. GGUF quantized versions of almost every major release are equally popular, reflecting a strong preference for local deployment. Finally, Baidu’s Unlimited‑OCR, Krea‑2 Turbo (text‑to‑image), and the LiquidAI 230M parameter model show that the ecosystem values both extreme specialization and extreme efficiency.

## 2. Trending Models by Category

### 🧠 Language Models

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — zai-org · 2,624 likes · 98,994 downloads  
  The week’s most‑liked model, a 5.2B‑scale MoE transformer optimized for conversational and instruction‑following tasks.

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — yuxinlu1 · 2,412 likes · 536,130 downloads  
  A GGUF‑quantized Gemma‑4‑12B fine‑tune focused on code generation and reasoning, one of the highest‑downloaded models this week.

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — WeiboAI · 736 likes · 57,521 downloads  
  A compact 3B‑parameter model specialized for math reasoning, built on Qwen2 and drawing attention for its strong inference‑to‑size ratio.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — HauhauCS · 2,268 likes · 3,331,475 downloads  
  An uncensored, “aggressive” fine‑tune of Qwen3.6‑35B‑A3B (MoE), extremely popular for its minimal safety filtering and vision capabilities.

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** — deepreinforce‑ai · 258 likes · 20,266 downloads  
  A GGUF version of the Qwen3.5‑based Ornith model, part of a family ranging from 9B to 397B parameters, all trending under the same author.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — microsoft · 359 likes · 6,447 downloads  
  A 4B‑parameter model built on Qwen3 and designed for ultra‑long context processing (up to ~1M tokens), notable for its “Explorer SubAgent” architecture.

- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)** — LiquidAI · 120 likes · 9,791 downloads  
  A **230M**‑parameter language model from Liquid AI, proving that small models can still capture the community’s curiosity for local deployment.

- **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)** — Chunjiang‑Intelligence · 108 likes · 1,328 downloads  
  A cybersecurity‑focused decoder based on DeepSeek v4, reflecting growing demand for domain‑specific safety models.

### 🎨 Multimodal & Generation

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** — baidu · 1,077 likes · 212,760 downloads  
  A powerful image‑to‑text OCR model released by Baidu, trending for its ability to handle diverse and complex document layouts.

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** — krea · 294 likes · 17,445 downloads  
  A text‑to‑image diffusion model that promises faster inference over its predecessor (Krea‑2‑Raw), with a focus on high‑quality generation.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — nvidia · 2,391 likes · 570,466 downloads  
  A 3B‑parameter vision model that can locate arbitrary objects in images without predefined class lists — a major step in open‑world image grounding.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — MiniMaxAI · 1,248 likes · 182,714 downloads  
  A multimodal VL model (M3) that handles both image and text inputs, gaining traction for its strong cross‑modal understanding.

- **[datalab-to/lift](https://huggingface.co/datalab-to/lift)** — datalab‑to · 159 likes · 6,676 downloads  
  A Qwen3.5‑based model specialized in PDF understanding and document extraction, filling a niche business need.

### 🔧 Specialized Models

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — nvidia · 712 likes · 61,857 downloads  
  A streaming automatic speech recognition model with only 0.6B parameters, optimized for low‑latency real‑time transcription.

- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)** — nvidia · 363 likes · 5,022,254 downloads  
  An NVFP4‑quantized version of Qwen3.6 MoE, achieving huge download numbers thanks to NVIDIA’s hardware‑aware optimization.

- **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)** — HauhauCS · 93 likes · 32,222 downloads  
  An uncensored, QAT‑quantized (quantization‑aware training) Gemma‑4‑12B variant, catering to users who prefer minimal content filtering.

### 📦 Fine‑tunes & Quantizations

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** — empero‑ai · 622 likes · 712,627 downloads  
  A GGUF‑quantized variant of the Qwythos‑9B model (image‑text‑to‑text), among the most downloaded GGUF files this week.

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — unsloth · 420 likes · 125,230 downloads  
  The official GGUF conversion of GLM‑5.2 by unsloth, making the popular MoE model accessible on consumer hardware.

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** — yuxinlu1 · 699 likes · 206,828 downloads  
  Another Gemma‑4‑12B GGUF, this time tuned for “agentic” and terminal‑based coding workflows.

- **[Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF)** — Jackrong · 95 likes · 49,935 downloads  
  A GGUF of a Qwen3.6‑based 27B coder model with multi‑token prediction (MTP) support, valued for its compatibility with llama.cpp.

## 3. Ecosystem Signal

Several macro trends are shaping the Hugging Face ecosystem this week:

- **MoE is mainstream**: Models like GLM‑5.2, Qwen3.6‑35B‑A3B, and Ornith‑1.0‑397B all use Mixture‑of‑Experts architectures, allowing large effective parameter counts while keeping inference costs manageable. The community is voting with downloads, especially when combined with quantization.

- **Open‑weight dominance continues**: Almost every trending model (GLM, Qwen, Gemma) has open weights; no proprietary API‑only model appears in the top 30. This signals strong preference for self‑hosted and fine‑tunable solutions.

- **GGUF and NVFP4 quantization wars**: GGUF remains the most popular format for local inference (13 of 30 models are GGUF), but NVIDIA’s NVFP4 is gaining traction with hardware‑tuned quantizations (e.g., Qwen3.6‑NVFP4, GLM‑5.2‑NVFP4). The community is clearly optimizing for speed and memory on both CPU/Apple Silicon (GGUF) and NVIDIA GPUs (NVFP4).

- **Uncensored and “abliterated” fine‑tunes**: Many high‑download models explicitly remove safety filters (HauhauCS, huihui‑ai), reflecting a sub‑community that prioritizes maximal freedom and research use over alignment. This is a persistent and growing niche.

- **Specialist models (OCR, ASR, object grounding) see strong engagement**: Baidu’s Unlimited‑OCR and NVIDIA’s LocateAnything prove that task‑specific models can compete with general‑purpose LLMs in likes and downloads, especially when they solve a clear pain point (documents, computer vision).

- **Small models are not neglected**: The 230M LiquidAI LFM2.5 and the 0.6B Nemotron ASR show that ultra‑efficient models still attract attention, likely because they run on edge devices or serve as building blocks in larger pipelines.

## 4. Worth Exploring

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — A breakthrough in open‑world object localization: given any text description, it can point to regions in an image without requiring class labels. With nearly 2.4k likes and 570k downloads in a single week, it’s a must‑try for anyone working on vision‑language tasks.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — A versatile multimodal VL model that handles both images and text in a unified transformer. With 1,248 likes and steady downloads, it’s an excellent choice for those exploring image captioning, VQA, or document understanding outside the OCR niche.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — A 4B‑parameter model that pushes context windows to about 1M tokens without quadratic memory blowups. If you work on long document analysis, codebases, or multi‑turn conversational memory, this model is worth studying both as a practical tool and as a research artifact.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*