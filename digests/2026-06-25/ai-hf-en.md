# Hugging Face Trending Models Digest 2026-06-25

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-25 10:25 UTC

---

# Hugging Face Trending Models Digest — 2026-06-25

## Today's Highlights

This week’s trending models are dominated by **MoE (Mixture-of-Experts) architectures**, with **DeepSeek-V4-Pro** (5,054 likes) and **GLM-5.2** (2,400 likes) leading the language model charge. **Gemma-4-12B** variants appear in multiple finetunes and quantized forms, reflecting strong community interest in Google’s latest open-weight release. **Multimodal models** are surging — **NVIDIA’s LocateAnything-3B** (2,354 likes) and **MiniMax-M3** (1,229 likes) showcase vision-language capabilities, while **Baidu’s Unlimited-OCR** (804 likes) tackles document understanding. Quantization remains a major driver, with **GGUF** and **NVFP4** formats enabling local deployment of these large models.

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
  Author: deepseek-ai | Likes: 5,054 | Downloads: 1,878,217  
  The most-liked model this week, a powerful MoE conversational LLM from DeepSeek that builds on their V4 lineage with improved reasoning and context handling.

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  Author: zai-org | Likes: 2,400 | Downloads: 67,107  
  A MoE-DSA text-generation model from Tsinghua’s GLM team, trending for its strong performance in chat and instruction-following tasks.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**  
  Author: google | Likes: 1,166 | Downloads: 2,187,644  
  Google’s latest open-weight “any-to-any” multimodal LLM, supporting text, image, and audio inputs; widely downloaded due to its versatility.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
  Author: moonshotai | Likes: 988 | Downloads: 502,106  
  A compressed image-text-to-text model optimized for code generation and feature extraction, drawing attention for its efficiency.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
  Author: microsoft | Likes: 339 | Downloads: 5,276  
  A lightweight 4B model built for agentic sub‑task exploration, designed to handle long contexts efficiently.

- **[poolside/Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1)**  
  Author: poolside | Likes: 96 | Downloads: 2,956  
  A vLLM/SGLang‑compatible LLM focused on specialized conversational tasks, gaining traction in production deployment circles.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  Author: baidu | Likes: 804 | Downloads: 70,743  
  An image-text-to-text pipeline that performs universal OCR with high accuracy, trending for its broad applicability in document processing.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
  Author: nvidia | Likes: 2,354 | Downloads: 407,838  
  A 3B vision-language model for fine-grained object localization in images, riding the wave of demand for spatial understanding.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**  
  Author: MiniMaxAI | Likes: 1,229 | Downloads: 154,350  
  A multimodal VL model from MiniMax, supporting image-text-to-text with state-of-the-art visual reasoning capabilities.

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** (and **[Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)**)  
  Author: krea | Likes: 216 / 177 | Downloads: 2,996 / 5,113  
  Text-to-image diffusion models (Turbo variant for speed, Raw for quality) that are gaining popularity for design and creative workflows.

- **[datalab-to/lift](https://huggingface.co/datalab-to/lift)**  
  Author: datalab-to | Likes: 148 | Downloads: 5,189  
  An image-text-to-text model specialized for PDF parsing and document understanding, filling a niche in enterprise document AI.

- **[Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit)**  
  Author: Boogu | Likes: 122 | Downloads: 824  
  A bilingual (EN/ZH) image editing diffusion model, modestly trending among creators needing localized editing tools.

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
  Author: nvidia | Likes: 683 | Downloads: 50,553  
  A cache‑aware streaming ASR model from NVIDIA, notable for its real‑time speech recognition capabilities.

- **[owensong/Inflect-Nano-v1](https://huggingface.co/owensong/Inflect-Nano-v1)**  
  Author: owensong | Likes: 198 | Downloads: 0  
  An ultra‑small text-to-speech model (PyTorch), intriguing for potential on‑device TTS deployment despite zero downloads (new release).

### 🔧 Specialized Models (code, math, medical, embeddings)

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
  Author: yuxinlu1 | Likes: 2,317 | Downloads: 495,813  
  A code‑specialized GGUF‑quantized finetune of Gemma-4-12B, heavily downloaded for its strong reasoning and coding benchmarks.

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
  Author: yuxinlu1 | Likes: 563 | Downloads: 165,187  
  Another Gemma-4‑based GGUF, optimized for agentic and terminal‑based interactions, popular among automation enthusiasts.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** (and **[GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**)  
  Author: empero-ai | Likes: 401 / 342 | Downloads: 134,294 / 10,160  
  A 9B reasoning model influenced by the “Claude Mythos” series, offered in both full precision and GGUF format.

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
  Author: WeiboAI | Likes: 703 | Downloads: 51,717  
  A 3B math‑focused language model (based on Qwen2), trending for its lightweight yet capable mathematical reasoning.

- **[LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)** (and **[LFM2.5-ColBERT-350M](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M)**)  
  Author: LiquidAI | Likes: 120 / 88 | Downloads: 11,833 / 3,600  
  Sentence‑similarity and ColBERT embeddings from LiquidAI’s LFM2.5 family, gaining traction in retrieval-augmented generation pipelines.

- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)**  
  Author: Qwen | Likes: 188 | Downloads: 3,389  
  A 35B MoE variant (with 3B active parameters) designed for agentic environments, part of the expanding Qwen ecosystem.

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
  Author: unsloth | Likes: 365 | Downloads: 88,915  
  The unsloth‑quantized GGUF version of GLM-5.2, making the powerful MoE model accessible for CPU/consumer GPU inference.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  Author: HauhauCS | Likes: 2,215 | Downloads: 3,520,206  
  An aggressively uncensored GGUF finetune of Qwen3.6 MoE, extremely high download count reflecting demand for creative unfiltered models.

- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)**  
  Author: nvidia | Likes: 317 | Downloads: 4,602,255  
  NVIDIA’s Model‑Optimized quantization (NVFP4) of the Qwen3.6 MoE, the most downloaded model this week, enabling efficient GPU inference.

- **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)**  
  Author: HauhauCS | Likes: 77 | Downloads: 15,128  
  A balanced, uncensored GGUF variant of Gemma-4-12B, part of a broader trend of “abliterated” or uncensored community finetunes.

- **[huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated)**  
  Author: huihui-ai | Likes: 126 | Downloads: 4,874  
  An “abliterated” (refusal‑removed) Gemma-4-12B coder finetune, reflecting community desire for more permissive model behavior.

- **[Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF)**  
  Author: Jackrong | Likes: 83 | Downloads: 19,382  
  A GGUF of a Qwen3.6‑derivative coder model with Multi‑Token Prediction (MTP) support, showing innovative quantization targets.

## Ecosystem Signal

The June 2026 Hugging Face landscape is defined by **Mixture-of-Experts architectures becoming mainstream**. Models like GLM-5.2, DeepSeek-V4-Pro, and Qwen3.6-35B-A3B dominate leaderboards and downloads, proving that MoE’s ability to activate only a fraction of parameters is the preferred path to scaling without prohibitive compute costs. **Open-weight models from major labs** (Google, NVIDIA, DeepSeek, Alibaba/Qwen, Baidu) continue to push the envelope, while proprietary‑style models from MiniMax and Moonshot also release openly, blurring the line between open and closed.

**Quantization is no longer an afterthought — it is a primary release format.** NVIDIA’s NVFP4 scheme and the widespread GGUF conversions (often by unsloth, yuxinlu1, and community developers) make models like Qwen3.6 and Gemma-4 viable on consumer hardware. Indeed, the highest download counts belong to quantized variants (4.6M for Qwen3.6-NVFP4, 3.5M for the uncensored GGUF), indicating that **practical deployability drives adoption** far more than raw capability alone.

Another clear trend is **unified multimodal models**: Google’s “any-to-any” Gemma-4-12B-it, NVIDIA’s LocateAnything, and MiniMax-M3 all blur the line between language and vision. At the same time, specialized models (OCR, streaming ASR, nano‑TTS, embeddings) are thriving, filling niches that general‑purpose models cannot yet cover efficiently. **Community finetuning** around safety‑removal (“uncensored”, “abliterated”) remains a durable sub‑trend, especially for coding and creative use cases.

## Worth Exploring

1. **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
   With the highest likes and strong reasoning benchmarks, this MoE model is the definitive release of the week. Studying its architecture and mixture routing can inform next‑generation scaling strategies.

2. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   A compact (3B) model that performs detailed object localization without requiring a full object‑detection pipeline. Its high likes and downloads suggest a genuine unmet need in vision‑language interaction.

3. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
   A rare example of a model explicitly tuned for “Explorer SubAgent” tasks — useful for anyone building multi‑agent systems or long‑context retrieval. Its small footprint (4B) is ideal for prototyping agentic architectures.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*