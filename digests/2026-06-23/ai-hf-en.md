# Hugging Face Trending Models Digest 2026-06-23

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-23 10:50 UTC

---

# Hugging Face Trending Models Digest – June 23, 2026

## Today's Highlights

This week’s Hugging Face trending list is dominated by **DeepSeek-V4-Pro** (5,021 likes, 2.2M downloads), solidifying the DeepSeek lineage as a community favorite. **NVIDIA’s LocateAnything-3B** (2,303 likes) and **zai-org/GLM-5.2** (2,097 likes) introduce novel MoE and vision‑grounding capabilities, while **Google’s gemma-4-12B-it** (any‑to‑any pipeline) signals a push toward unified multimodal models. Small reasoning models (e.g., VibeThinker-3B) and aggressive quantization activity (GGUF, FP8) continue to drive accessibility. Notably, the **GLM-5.2** family alone has three entries (base, GGUF, FP8) in the top 30, reflecting strong community engagement.

## Trending Models

### 🧠 Language Models (LLMs, chat, reasoning)

- [**zai-org/GLM-5.2**](https://huggingface.co/zai-org/GLM-5.2) — zai-org | 2,097 likes, 40k downloads  
  A new MoE‑based conversational LLM with 5.2B active parameters, trending for its DSA‑MoE architecture and impressive generation quality.

- [**WeiboAI/VibeThinker-3B**](https://huggingface.co/WeiboAI/VibeThinker-3B) — WeiboAI | 632 likes, 41k downloads  
  A lightweight Qwen2‑based reasoning model specialized for math tasks, gaining traction for its strong performance at only 3B parameters.

- [**microsoft/FastContext-1.0-4B-SFT**](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT) — microsoft | 301 likes, 4.4k downloads  
  A 4B model fine‑tuned for ultra‑long context handling (up to 256K tokens), featuring an Explorer SubAgent for efficient retrieval.

- [**empero-ai/Qwythos-9B-Claude-Mythos-5-1M**](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M) — empero-ai | 149 likes, 1.8k downloads  
  A Qwen3.5‑based conversational model fine‑tuned on synthetic Claude‑style data, popular for its creative reasoning.

- [**Mia-AiLab/Qwable-3.6-27b**](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b) — Mia-AiLab | 127 likes, 25k downloads  
  A 27B Qwen3.6 variant (both transformers and GGUF formats), part of the growing Qwable family blending Qwen with custom tuning.

- [**lordx64/Qwable-v1**](https://huggingface.co/lordx64/Qwable-v1) — lordx64 | 168 likes, 4.5k downloads  
  A Qwen3.5 MoE model fine‑tuned for vision‑language tasks, bridging LLM and multimodal capabilities.

- [**deepseek-ai/DeepSeek-V4-Pro**](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) — deepseek-ai | 5,021 likes, 2.2M downloads  
  The top‑trending model this week: a large conversational MoE LLM with state‑of‑the‑art reasoning and open‑weight availability.

- [**poolside/Laguna-M.1**](https://huggingface.co/poolside/Laguna-M.1) — poolside | 93 likes, 2.8k downloads  
  A code‑focused LLM from Poolside, optimized for vLLM and SGLang deployment, trending among enterprise developers.

- [**CohereLabs/North-Mini-Code-1.0**](https://huggingface.co/CohereLabs/North-Mini-Code-1.0) — CohereLabs | 484 likes, 21.6k downloads  
  A compact MoE code model from Cohere, designed for efficient code generation and reasoning.

- [**MoonshotAI/Kimi-K2.7-Code**](https://huggingface.co/moonshotai/Kimi-K2.7-Code) — moonshotai | 966 likes, 447k downloads  
  A compressed, multimodal code model (image‑text‑to‑text) that combines vision and code reasoning; heavily downloaded.

### 🎨 Multimodal & Generation (vision, audio, text-to-X)

- [**MiniMaxAI/MiniMax-M3**](https://huggingface.co/MiniMaxAI/MiniMax-M3) — MiniMaxAI | 1,215 likes, 131k downloads  
  A powerful image‑text‑to‑text model (M3 VL) excelling at visual understanding and generation tasks.

- [**nvidia/LocateAnything-3B**](https://huggingface.co/nvidia/LocateAnything-3B) — nvidia | 2,303 likes, 274k downloads  
  NVIDIA’s 3B vision‑grounding model that can locate and segment any object in images, trending for its zero‑shot accuracy.

- [**HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) — HauhauCS | 2,129 likes, 3.9M downloads  
  An uncensored, MoE vision‑language model (35B total, 3B active) with aggressive fine‑tuning; the highest downloads on this list.

- [**google/diffusiongemma-26B-A4B-it**](https://huggingface.co/google/diffusiongemma-26B-A4B-it) — google | 1,053 likes, 949k downloads  
  Google’s 26B diffusion‑based multimodal model (4B active) for image‑text‑to‑text generation, combining Gemma with diffusion.

- [**google/gemma-4-12B-it**](https://huggingface.co/google/gemma-4-12B-it) — google | 1,147 likes, 1.9M downloads  
  A unified any‑to‑any model supporting text, image, and audio inputs/outputs; a flagship for multimodal generalization.

- [**datalab-to/lift**](https://huggingface.co/datalab-to/lift) — datalab-to | 128 likes, 3.2k downloads  
  A Qwen3.5‑based model fine‑tuned for PDF and document image understanding, filling a niche in enterprise document AI.

- [**Comfy-Org/Boogu-Image**](https://huggingface.co/Comfy-Org/Boogu-Image) — Comfy-Org | 80 likes, 0 downloads  
  A ComfyUI‑compatible image generation model (likely based on Diffusers), ready for local creative workflows.

### 🔧 Specialized Models (code, math, OCR, ASR, embeddings, image editing)

- [**baidu/Unlimited-OCR**](https://huggingface.co/baidu/Unlimited-OCR) — baidu | 314 likes, 8.4k downloads  
  Baidu’s state‑of‑the‑art OCR model for unlimited character recognition in images, trending for its speed and accuracy.

- [**nvidia/nemotron-3.5-asr-streaming-0.6b**](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b) — nvidia | 644 likes, 41k downloads  
  A lightweight 0.6B streaming ASR model with cache‑aware decoding, enabling real‑time speech transcription.

- [**owensong/Inflect-Nano-v1**](https://huggingface.co/owensong/Inflect-Nano-v1) — owensong | 171 likes, 0 downloads  
  An ultra‑small PyTorch TTS model (under 100 MB) for on‑device speech synthesis, niche but innovative.

- [**LiquidAI/LFM2.5-Embedding-350M**](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M) — LiquidAI | 105 likes, 10k downloads  
  A 350M sentence‑embedding model from Liquid AI, optimized for semantic similarity and retrieval tasks.

- [**LiquidAI/LFM2.5-ColBERT-350M**](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M) — LiquidAI | 83 likes, 2.5k downloads  
  A ColBERT variant of the same embedding model, designed for efficient multi‑vector retrieval.

- [**Boogu/Boogu-Image-0.1-Edit**](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit) — Boogu | 105 likes, 592 downloads  
  A Diffusers‑based image editing model (Apache‑2.0) that supports text‑guided edits, popular for its simplicity.

- [**ostris/ideogram_4_turbotime_lora**](https://huggingface.co/ostris/ideogram_4_turbotime_lora) — ostris | 106 likes, 3.7k downloads  
  A LoRA adapter for Ideogram‑4‑FP8 that accelerates generation (TurboTime), widely used by the image‑generation community.

### 📦 Fine-tunes & Quantizations (GGUF, AWQ, FP8, community adapters)

- [**yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) — yuxinlu1 | 2,199 likes, 456k downloads  
  A GGUF quantized version of a Gemma‑4 coder fine‑tune, offering efficient local code generation.

- [**yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF) — yuxinlu1 | 407 likes, 96k downloads  
  Another GGUF quant of a Gemma‑4 agentic (terminal‑use) model, trending for autonomous coding workflows.

- [**unsloth/GLM-5.2-GGUF**](https://huggingface.co/unsloth/GLM-5.2-GGUF) — unsloth | 267 likes, 56k downloads  
  The official UnsloTH quantization of GLM‑5.2, enabling low‑resource inference on consumer hardware.

- [**zai-org/GLM-5.2-FP8**](https://huggingface.co/zai-org/GLM-5.2-FP8) — zai-org | 139 likes, 395k downloads  
  The FP8 (8‑bit float) variant of GLM‑5.2, offering a good speed‑accuracy trade‑off.

- [**empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF**](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF) — empero-ai | 153 likes, 27k downloads  
  GGUF quantization of the Qwythos 9B reasoning model, popular for local use.

- [**bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF**](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF) — bytkim | 110 likes, 66k downloads  
  A GGUF variant of Qwen3.6‑27B with multi‑token prediction tuning, optimizing throughput.

## Ecosystem Signal

The current landscape reflects a **strong shift toward Mixture‑of‑Experts (MoE) architectures** across multiple model families: GLM‑5.2, DeepSeek‑V4, Qwen3.6 MoE, and Gemma‑4 all leverage MoE to balance performance and efficiency. **Open‑weight models continue to dominate** — every top‑10 entry by likes is fully open‑weight, reinforcing Hugging Face’s role as the hub for reproducible AI.  

**Quantization and fine‑tuning activity is explosive:** GGUF and FP8 variants of every popular model appear within days of release. Community creators like `yuxinlu1` and `empero-ai` are churning out specialized fine‑tunes (coder, agentic, uncensored) and quantized packages, making state‑of‑the‑art models accessible on consumer GPUs and edge devices.  

**Multimodality is becoming the default** — over half of the trending models have an image‑text‑to‑text pipeline, and Google’s `gemma-4-12B-it` (any‑to‑any) signals the next frontier. Small, specialized models (3B‑9B) for reasoning, code, and OCR are gaining disproportionate attention, as the community values task‑specific efficiency over raw scale.

## Worth Exploring

- [**deepseek-ai/DeepSeek-V4-Pro**](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) — The most‑liked model this week, offering a rare blend of MoE efficiency and top‑tier reasoning. Download it to benchmark against other flagship LLMs or to build high‑quality conversational agents.

- [**nvidia/LocateAnything-3B**](https://huggingface.co/nvidia/LocateAnything-3B) — A standout in vision grounding: 3B parameters, zero‑shot segmentation. Ideal for robotics, visual QA, and any application requiring precise object localization without training.

- [**microsoft/FastContext-1.0-4B-SFT**](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT) — Although newer and lesser‑downloaded, this 4B model with 256K context is a fascinating research artifact. It explores how long‑context fine‑tuning can avoid costly RAG pipelines — worth studying for developers working on document‑heavy applications.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*