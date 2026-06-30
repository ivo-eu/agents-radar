# Hugging Face Trending Models Digest 2026-06-30

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-30 10:45 UTC

---

# Hugging Face Trending Models Digest — June 30, 2026

## Today’s Highlights

This week’s trending models are dominated by **large MoE architectures** (Qwen3.5/3.6, GLM-5.2) and their quantization variants, with **nvidia/Qwen3.6-35B-A3B-NVFP4** reaching a staggering **5.5M downloads** — the highest single-model traffic on the list. **Community fine-tuning** is exceptionally active: the *yuxinlu1/gemma-4-12B-coder* GGUF (2.5K likes, 575K downloads) and *HauhauCS/Qwen3.6-35B-A3B-Uncensored* (2.3K likes, 3M downloads) show strong demand for specialized code and uncensored variants. Meanwhile, **multimodal models** are gaining ground — *baidu/Unlimited-OCR* (1.4K likes) and *nvidia/LocateAnything-3B* (2.5K likes, 800K downloads) represent a shift toward practical vision-language applications. Finally, the resurgence of **DeepSeek V4** with DSpark optimization and cybersecurity fine-tunes signals that both performance and safety-focused derivatives are in high demand.

---

## Trending Models by Category

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **[GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — *zai-org* — 2,980 likes, 142,547 downloads. MoE-DSA conversational model with strong reasoning, the most liked model this week.  
- **[DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)** — *deepseek-ai* — 238 likes, 6,939 downloads. Latest DeepSeek V4 with DSpark optimization for enhanced reasoning and efficiency.  
- **[DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)** — *deepseek-ai* — 103 likes, 4,446 downloads. Faster sister variant of DeepSeek V4, also using DSpark.  
- **[LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)** — *LiquidAI* — 159 likes, 17,839 downloads. Compact 230M model from the LFM2 series, suited for lightweight deployment.

### 🔧 Specialized Models (Code, Math, Medical, Embeddings)

- **[VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — *WeiboAI* — 751 likes, 67,777 downloads. 3B math-reasoning model fine-tuned for problem-solving and arithmetic tasks.

### 🎨 Multimodal & Generation (Image, Video, Audio, Text-to-X)

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** — *baidu* — 1,427 likes, 429,056 downloads. Universal OCR model for text extraction from images, trending for its practical utility.  
- **[Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)** — *Qwen* — 448 likes, 28,480 downloads. MoE world model for agentic tasks, supporting multimodal inputs.  
- **[Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)** — *krea* — 246 likes, 32,008 downloads. Base text-to-image diffusion model for the Krea-2 ecosystem.  
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — *nvidia* — 2,493 likes, 800,597 downloads. 3B image localisation model for zero-shot object detection and segmentation.  
- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — *nvidia* — 749 likes, 87,115 downloads. Streaming speech recognition model with 0.6B parameters for real-time ASR.

### 📦 Fine-tunes & Quantizations (Community Fine-tunes, GGUF, AWQ, LoRAs)

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** — 1,002 likes, 970,663 downloads. Quantized GGUF version of the Qwythos fine-tune (Qwen3.5‑based) for reasoning.  
- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** — 571 likes, 99,359 downloads. Full‑precision multimodal fine-tune of Qwen3.5.  
- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** — 498 likes, 157,418 downloads. GGUF quant of the 35B MoE Ornith model.  
- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)** — 324 likes, 98,750 downloads. 9B GGUF variant of Ornith.  
- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)** — 293 likes, 26,151 downloads. Full‑precision 9B Ornith multimodal fine-tune.  
- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)** — 251 likes, 69,048 downloads. Full‑precision 35B Ornith MoE.  
- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)** — 172 likes, 2,564 downloads. Extreme‑scale 397B MoE variant of Ornith.  
- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** — 863 likes, 257,216 downloads. GGUF fine‑tune of Gemma 4 for agentic coding and terminal use.  
- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — 2,512 likes, 575,255 downloads. GGUF fine‑tune of Gemma 4 specialised for code generation and reasoning.  
- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** — 403 likes, 45,668 downloads. Faster fine‑tune of Krea-2-Raw for text-to-image generation.  
- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — 471 likes, 180,394 downloads. Unsloth’s GGUF quantisation of GLM-5.2.  
- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)** — 176 likes, 104,746 downloads. Nvidia’s NVFP4 4‑bit quantisation of GLM-5.2.  
- **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)** — 132 likes, 1,519 downloads. Cybersecurity‑focused fine‑tune of DeepSeek V4 for penetration testing.  
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — 2,343 likes, 3,017,678 downloads. Uncensored fine‑tune of Qwen3.6 MoE with an aggressive personality.  
- **[unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)** — 118 likes, 155,333 downloads. GGUF quantisation of Qwen AgentWorld.  
- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)** — 380 likes, 5,495,402 downloads. Nvidia’s NVFP4 4‑bit quantisation of Qwen3.6 MoE (highest downloads).  
- **[ilkerzgi/fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs)** — 82 likes, 0 downloads. LoRA collection for applying creative styles to Krea-2 generations.  
- **[BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)** — 86 likes, 253 downloads. Q6 quantised fine‑tune of Qwen3 for offensive security and bug hunting.  
- **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)** — 119 likes, 0 downloads. LoRA for LTX-2.3 video model enabling 3D‑realistic image-to-video generation.

---

## Ecosystem Signal

The week’s top trends coalesce around **Mixture-of-Experts (MoE) architectures** — the Qwen3.5/3.6 and GLM-5.2 families dominate both base releases and community derivatives. Quantisation is a major driver: nvidia’s **NVFP4** format and community **GGUF** variants account for over **6 million combined downloads**, indicating heavy demand for local deployment and inference on commodity hardware. Fine‑tuning activity is fractal — from massive 397B Ornith models to specialised cybersecurity and math models, the community is actively tailoring base models for niche use cases. **Open‑weight models** remain the norm; no proprietary‑only models appear on this list. Notably, **Krea‑2** and **LTX‑2.3** demonstrate a parallel surge in image/video generation tools, with associated LoRAs extending creative control. Finally, the presence of **Comfy‑Org/Krea‑2** (a ComfyUI node pack) hints at a growing ecosystem of model‑adjacent tooling that lowers the barrier for artists and developers.

---

## Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — With 2.5K likes and 800K downloads, this zero‑shot localisation model is a standout for computer vision tasks. Its small size (3B) and high accuracy make it immediately deployable in production pipelines.

2. **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — The most‑liked fine‑tune of Gemma 4 (2.5K likes), available as a convenient GGUF. It excels at code generation and reasoning, and its quantised form can be run locally with modest resources.

3. **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** — A practical, production‑ready OCR model with strong community traction (1.4K likes, 429K downloads). Its image‑text‑to‑text pipeline is ideal for document digitisation, translation, and accessibility use cases.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*