# Hugging Face Trending Models Digest 2026-07-28

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-28 00:11 UTC

---

# Hugging Face Trending Models Digest — 2026-07-28

## Today's Highlights

This week's trending list is dominated by **multimodal and MoE (Mixture-of-Experts) architectures**, led by moonshotai's **Kimi-K3** (6,046 likes) and the **Qwen3.6-35B-A3B** family, which spawned multiple community fine-tunes and quantizations. **GLM-5.2** from zai-org continues to gain traction with over 1M downloads, while **extreme quantization** (1-bit, ternary, 2-bit models like *Bonsai-27B-gguf* and *Ternary-Bonsai-27B-gguf*) signals growing demand for local inference. A major trend is the convergence of **image-text-to-text pipelines** with vision capabilities, now appearing in code models, OCR tools, and computer-use agents. The community fine-tuning ecosystem remains highly active, particularly around "uncensored" variants of Qwen and Laguna models.

---

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[poolside/Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1)**  
  *poolside* | 754 likes | 63,605 downloads  
  A large text-generation model built for code and general reasoning, trending as the base for multiple quantization variants.

- **[upstage/Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B)**  
  *upstage* | 628 likes | 3,761 downloads  
  A massive 250B parameter open-weight LLM designed for frontier-level language tasks, gaining attention for its scale and performance.

- **[Nanbeige/Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B)**  
  *Nanbeige* | 492 likes | 16,518 downloads  
  A compact 3B parameter LLM optimized for efficient local deployment, popular for its strong performance-to-size ratio.

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  *zai-org* | 4,547 likes | 1,003,547 downloads  
  A powerful MoE-based conversational model with high download counts, anchoring a rapidly growing ecosystem of fine-tunes and quantized variants.

- **[Motif-Technologies/Motif-3-Beta](https://huggingface.co/Motif-Technologies/Motif-3-Beta)**  
  *Motif-Technologies* | 199 likes | 2,532 downloads  
  A feature-extraction and generation model noted for its unique architectural approach, attracting early adopters.

- **[fdtn-ai/antares-1b](https://huggingface.co/fdtn-ai/antares-1b)**  
  *fdtn-ai* | 207 likes | 6,421 downloads  
  A 1B parameter security-focused LLM with a hybrid MoE architecture, trending for its niche in safe code generation.

---

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)**  
  *moonshotai* | 6,046 likes | 2,850 downloads  
  A new multimodal flagship model supporting image-text-to-text with compressed-tensor optimizations, the top-trending model this week.

- **[thinkingmachines/Inkling](https://huggingface.co/thinkingmachines/Inkling)**  
  *thinkingmachines* | 1,602 likes | 36,196 downloads  
  A conversational multimodal model designed for rich image-text interactions, gaining traction for its chat-oriented design.

- **[microsoft/Mage-Flow](https://huggingface.co/microsoft/Mage-Flow)**  
  *microsoft* | 386 likes | 1,691 downloads  
  A text-to-image generation and editing model using diffusion techniques, notable for Microsoft's push into creative AI.

- **[baseten/GLM-5.2-Vision-NVFP4](https://huggingface.co/baseten/GLM-5.2-Vision-NVFP4)**  
  *baseten* | 124 likes | 2,276 downloads  
  An NVFP4-quantized vision-language variant of GLM-5.2, reflecting the trend toward efficient multimodal deployment.

- **[nvidia/Cosmos3-Edge](https://huggingface.co/nvidia/Cosmos3-Edge)**  
  *nvidia* | 133 likes | 33,127 downloads  
  A diffusion-based edge model for video/image generation, signaling Nvidia's investment in on-device media generation.

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)**  
  *Qwen* | 2,546 likes | 6,187,853 downloads  
  The base vision‑MoE model driving the week's most active fine-tuning ecosystem, with massive community adoption.

- **[microsoft/Mage-Flow-Edit-Turbo](https://huggingface.co/microsoft/Mage-Flow-Edit-Turbo)**  
  *microsoft* | 100 likes | 1,115 downloads  
  A fast instruction-based image editing model, an optimized sibling of Mage-Flow for real-time editing.

---

### 🔧 Specialized Models (code, math, medical, embeddings, OCR, TTS)

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  *baidu* | 3,329 likes | 2,645,773 downloads  
  A universal OCR model capable of recognizing text in unlimited contexts (documents, scenes, multilingual), extremely popular for its versatility.

- **[Kwaipilot/KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev)**  
  *Kwaipilot* | 240 likes | 5,312 downloads  
  A MoE-based code generation model with vision input, targeting developer workflows with multimodal code understanding.

- **[owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)**  
  *owensong* | 223 likes | 483 downloads  
  A lightweight text-to-speech model optimized for CPU and edge AI, trending for local voice synthesis use cases.

- **[microsoft/Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B)**  
  *microsoft* | 138 likes | 1,406 downloads  
  A 27B computer-use agent model that grounds language in GUI actions, part of Microsoft's push toward autonomous desktop agents.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
  *moonshotai* | 1,323 likes | 695,744 downloads  
  A code-specialized variant of the Kimi family with compressed tensors, trending for its efficient code reasoning.

- **[ATH-MaaS/OvisOCR2](https://huggingface.co/ATH-MaaS/OvisOCR2)**  
  *ATH-MaaS* | 327 likes | 42,152 downloads  
  A Qwen3.5-based OCR model with enhanced vision-language capabilities, addressing structured document extraction.

- **[conradlocke/krea2-identity-edit](https://huggingface.co/conradlocke/krea2-identity-edit)**  
  *conradlocke* | 555 likes | 0 downloads  
  A LoRA adapter for identity-preserving image editing with Krea-2, gaining interest despite zero downloads (likely newly published).

- **[owensong/Inflect-Nano-v2](https://huggingface.co/owensong/Inflect-Nano-v2)**  
  *owensong* | 91 likes | 349 downloads  
  A smaller, faster TTS sibling of Inflect-Micro, focused on ultra-constrained edge devices.

---

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ, NVFP4)

- **[DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF)**  
  *DavidAU* | 748 likes | 634,146 downloads  
  An extensively merged and uncensored GGUF fine-tune of Qwen3.6, popular for its "no-restrictions" approach and high download count.

- **[unsloth/Laguna-S-2.1-GGUF](https://huggingface.co/unsloth/Laguna-S-2.1-GGUF)**  
  *unsloth* | 217 likes | 117,456 downloads  
  The premier GGUF quantization of Laguna-S-2.1 from the Unsloth team, optimized for vLLM and local inference.

- **[prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)**  
  *prism-ml* | 1,069 likes | 648,938 downloads  
  A 2-bit ternary quantized 27B model pushing the limits of compression, trending for extreme memory savings.

- **[poolside/Laguna-S-2.1-GGUF](https://huggingface.co/poolside/Laguna-S-2.1-GGUF)**  
  *poolside* | 154 likes | 85,554 downloads  
  The official GGUF release of Laguna-S-2.1, a baseline for the model's quantization ecosystem.

- **[poolside/Laguna-S-2.1-NVFP4](https://huggingface.co/poolside/Laguna-S-2.1-NVFP4)**  
  *poolside* | 148 likes | 158,308 downloads  
  A 4-bit floating point (NVFP4) quantization of Laguna-S-2.1, leveraging NVIDIA's format for efficient inference on Ada GPUs.

- **[prism-ml/Bonsai-27B-gguf](https://huggingface.co/prism-ml/Bonsai-27B-gguf)**  
  *prism-ml* | 659 likes | 2,257,928 downloads  
  A 1-bit quantized 27B model (the most extreme compression this week), achieving the highest download count on the list.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  *HauhauCS* | 3,133 likes | 1,894,395 downloads  
  An aggressive, uncensored GGUF fine-tune of Qwen3.6-35B-A3B with vision capabilities, among the top downloads.

- **[LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V5-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V5-GGUF)**  
  *LuffyTheFox* | 187 likes | 83,658 downloads  
  A Hermes-style uncensored fine-tune of Qwen3.6 MoE, blending roleplay and reasoning in a quantized form.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
  *empero-ai* | 2,487 likes | 1,336,263 downloads  
  A quantized 9B reasoning model merging Claude-style instruction data with Mythos-style creativity, extremely popular for creative writing tasks.

---

## Ecosystem Signal

**Dominant model families** this week are **Qwen3.6 (MoE, especially the 35B-A3B variant)** and **GLM-5.2**, each spawning extensive quantization and fine-tuning ecosystems. The Qwen3.6 family alone accounts for at least six models on the list. The **Laguna-S-2.1** line from poolside is also growing, backed by multiple quantization formats (GGUF, NVFP4). Notably, **Kimi-K3** and **Kimi-K2.7-Code** from moonshotai signal a new multimodal-code family gaining momentum.

**Open-weight models dominate** the trending list. No proprietary-only models appear; even large releases like Solar-Open2-250B and GLM-5.2 are fully open. This reflects a continued shift toward openness in the foundation model ecosystem.

**Extreme quantization is a clear trend**: 1-bit (Bonsai-27B-gguf), ternary 2-bit (Ternary-Bonsai-27B-gguf), and NVFP4 quants demonstrate the community's drive to run large models on limited hardware. **Uncensored fine-tunes** remain a popular subculture, with multiple variants of Qwen3.6 and other models emphasizing content freedom.

**Infrastructure note**: GGUF continues to be the dominant quantization format, supported by tools like llama.cpp and Unsloth. The emergence of **NVFP4** for Ada GPUs and **compressed tensors** in Kimi models suggests new optimization frontiers.

---

## Worth Exploring

1. **[prism-ml/Bonsai-27B-gguf](https://huggingface.co/prism-ml/Bonsai-27B-gguf)** — At 1-bit quantization, this 27B model achieves the highest download count on the list (2.2M). It's a benchmark for how far LLM compression can go while maintaining usability, making it essential study material for deployment engineers.

2. **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)** — The week's top-trending model with 6k likes. As a multimodal compressed-tensor model from a major lab, it likely sets new efficiency standards for image-text-to-text tasks. Worth evaluating for multimodal workflows requiring low latency.

3. **[ATH-MaaS/OvisOCR2](https://huggingface.co/ATH-MaaS/OvisOCR2)** — Tied to the Qwen3.5 ecosystem and the OCR boom (see Unlimited-OCR's 2.6M downloads). This model represents the convergence of VLMs and document parsing, a rapidly growing enterprise use case.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*