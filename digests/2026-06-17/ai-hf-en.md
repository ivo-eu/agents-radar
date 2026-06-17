# Hugging Face Trending Models Digest 2026-06-17

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-17 03:58 UTC

---

# Hugging Face Trending Models Digest — June 17, 2026

## Today's Highlights

DeepSeek's **DeepSeek-V4-Pro** dominates the trending chart with nearly 5,000 weekly likes and 2.8M downloads, signaling the community's insatiable appetite for cutting-edge open-weight reasoning models. The MoE architecture race intensifies: Qwen leads with their **Qwen3.6-35B-A3B** (2,139 likes, 3.3M downloads), while Google's experimental **DiffusionGemma-26B-A4B-it** introduces a novel diffusion-based multimodal approach. Community quantization activity is at fever pitch, with unsloth and individual contributors producing over 10 GGUF variants of top models, and uncensored fine-tunes (HauhauCS, DavidAU) drawing massive download numbers despite lower like counts. Coding-focused models and image-text-to-text pipelines now account for roughly two-thirds of trending entries, reflecting a platform shift toward multimodal reasoning.

---

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — Author: deepseek-ai | 4,898 likes | 2,829,747 downloads  
  The flagship open-weight reasoning model of the DeepSeek V4 lineage, trending for its exceptional reasoning performance and massive adoption across the community.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — Author: moonshotai | 808 likes | 102,206 downloads  
  A compressed code generation model from Kimi's K2.5 family, gaining traction for efficient coding workloads with reduced memory footprint via compressed-tensors.

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — Author: CohereLabs | 412 likes | 12,129 downloads  
  Cohere's compact MoE code model optimized for conversational coding tasks, drawing interest as a smaller alternative to larger proprietary code LMs.

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — Author: WeiboAI | 199 likes | 0 downloads  
  A 3B-parameter math-oriented Qwen2-based model focused on reasoning, newly released and attracting early attention from the math-AI community.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — Author: microsoft | 163 likes | 192 downloads  
  Microsoft's lightweight 4B parameter model with extended context handling, leveraging Qwen3 base and Explorer SubAgent capabilities for agentic workflows.

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — Author: zai-org | 540 likes | 0 downloads  
  ZAI's latest MoE conversational model with dynamic sparse attention (DSA), gathering strong initial community interest despite zero current downloads.

---

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — Author: google | 951 likes | 375,974 downloads  
  Google's revolutionary diffusion-based multimodal LLM that generates and understands images, video, and text in a unified architecture; a paradigm shift from autoregressive-only models.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — Author: MiniMaxAI | 1,023 likes | 25,064 downloads  
  MiniMax's third-generation multimodal MoE model supporting image-text-to-text tasks, standing out for its agent-friendly design and MoE efficiency.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — Author: nvidia | 2,107 likes | 98,698 downloads  
  NVIDIA's compact yet powerful grounding and object detection model, trending for its ability to locate any object in images at just 3B parameters.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — Author: google | 1,055 likes | 1,223,383 downloads  
  The instruction-tuned variant of Google's Gemma 4 unified model, operating in any-to-any modality (text, image, audio, video) and exceptionally popular for deployment.

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — Author: bosonai | 467 likes | 43,361 downloads  
  A 4B-parameter TTS model built on Higgs multimodal Qwen3 architecture, gaining momentum for high-quality speech synthesis with low latency.

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — Author: nvidia | 483 likes | 5,777 downloads  
  NVIDIA's streaming ASR model with cache-aware decoding, highly efficient at just 0.6B parameters for real-time speech recognition.

- **[Zyphra/ZONOS2](https://huggingface.co/Zyphra/ZONOS2)** — Author: Zyphra | 100 likes | 539 downloads  
  An Apache-2.0 licensed TTS model from Zyphra, notable for its open-release strategy in a typically proprietary speech space.

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — Author: ideogram-ai | 560 likes | 12,466 downloads  
  FP8 quantized version of Ideogram's text-to-image diffusion model, drawing attention for high-quality image generation with reduced memory requirements.

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — Author: zai-org | 209 likes | 0 downloads  
  A pose-driven character animation and video generation diffusion model, pushing boundaries in controlled video synthesis.

- **[prefeitura-rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)** — Author: prefeitura-rio | 316 likes | 189,744 downloads  
  A massive 397B-parameter open MoE model based on Qwen3.5 architecture for image-text-to-text, surprisingly popular for its sheer scale and open-weight release.

---

### 🔧 Specialized Models (code, math, medical, embeddings)

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — Author: yuxinlu1 | 1,202 likes | 60,921 downloads  
  A heavily fine-tuned Gemma-4-based coding model combining multiple Composer and Fable5 techniques, trending as the top coding GGUF variant.

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — Author: Jackrong | 220 likes | 79,157 downloads  
  A 27B parameter code-focused variant of the Qwopus3.6 MoE family, optimized with Multi-Token Prediction (MTP) and GGUF for local deployment.

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — Author: nex-agi | 310 likes | 4,957 downloads  
  nex-agi's premium MoE model based on Qwen3.5 architecture, targeting professional agentic and reasoning workflows.

- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** — Author: nex-agi | 229 likes | 9,161 downloads  
  A smaller, more efficient sibling of Nex-N2-Pro, offering similar capabilities at reduced compute cost.

---

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — Author: HauhauCS | 1,897 likes | 2,716,651 downloads  
  An uncensored, aggressive fine-tune of Qwen3.6-35B-A3B; by far the most downloaded GGUF variant, reflecting strong demand for unfiltered local models.

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** — Author: DavidAU | 371 likes | 366,279 downloads  
  An extreme uncensored fine-tune of Qwen3.6 with numerous technique mashups (Heretic, NEO-CODE, Di-IMatrix), highly downloaded despite its lengthy name.

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — Author: OBLITERATUS | 336 likes | 76,044 downloads  
  Popular fine-tune of Google's Gemma 4 that removes safety alignment, trending due to the uncensored model movement and high download volume.

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — Author: unsloth | 290 likes | 120,435 downloads  
  Unsloth's GGUF quantization of Google's DiffusionGemma, making this novel multimodal architecture accessible to consumer hardware.

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — Author: unsloth | 634 likes | 1,009,602 downloads  
  The most downloaded GGUF version of Gemma 4, showcasing unsloth's dominant role in community quantization and the model's strong appeal.

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)** — Author: Qwen | 2,139 likes | 3,360,615 downloads  
  The official base MoE model from Qwen, with 35B total parameters and 3B active, setting the standard for efficient multimodal reasoning and driving the MoE wave.

---

## Ecosystem Signal

**MoE models are the dominant architecture of 2026.** Every major release in this digest — DeepSeek-V4-Pro, Qwen3.6-35B-A3B, GLM-5.2, Rio-3.5-Open-397B, North-Mini-Code-1.0, and MiniMax-M3 — uses mixture-of-experts. The community has decisively embraced sparse activation for its superior throughput-to-quality ratio, with 3B–4B active parameters from 27B–400B total being the sweet spot.

**Open-weight releases continue to outpace proprietary alternatives.** Google's Gemma 4 and DiffusionGemma, DeepSeek's V4 Pro, and Qwen's 3.6 series are all fully open-weight under permissive licenses. Only Ideogram's text-to-image model retains a more restrictive release. The trend is clear: open-weight models are winning mindshare and downloads by large margins.

**Quantization and uncensored fine-tuning are the two biggest community activities.** Unsloth alone accounts for 6 of the 30 trending models, and GGUF variants drive the vast majority of downloads for flagship models. Simultaneously, uncensored fine-tunes (from OBLITERATUS, HauhauCS, DavidAU) consistently achieve high download counts, revealing a persistent community demand for models without safety alignment layers — a trend that raises both use-case and ethical considerations.

**Multimodal is no longer optional.** Image-text-to-text pipelines now dominate the top of the trending list, with code models and speech models closely following. Google's DiffusionGemma, which unifies generation and understanding across all modalities, may represent the template for future model architectures.

---

## Worth Exploring

1. **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — This is the most architecturally novel model on the list, using diffusion within a language model for multimodal understanding and generation. Its any-to-any capability could redefine how we think about foundation models. Worth studying for its architecture and testing for multimodal reasoning tasks.

2. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — At 3B parameters with 2,107 likes, this is one of the most liked models per parameter. It demonstrates that focused, task-specific models can achieve both high community engagement and practical utility. Ideal for robotics, visual grounding, and retrieval-augmented generation pipelines.

3. **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — A fresh release with zero downloads yet steady likes. As a 3B math-focused model, it represents the growing niche of small, highly specialized reasoning models. Worth monitoring for emerging small-model RL or process-supervised training techniques that might generalize to other domains.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*