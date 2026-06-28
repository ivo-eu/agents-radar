# Hugging Face Trending Models Digest 2026-06-28

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-28 10:09 UTC

---

# Hugging Face Trending Models Digest — 2026-06-28

## Today's Highlights

This week's trending models are dominated by the **Qwen 3.5/3.6 MoE ecosystem**, with uncensored fine-tunes and quantized variants from community creators like HauhauCS drawing massive download counts (Qwen3.6-35B-A3B-Uncensored alone surpassed 3.2M downloads). **Google's Gemma-4** continues strong momentum, particularly in coding and agentic fine-tunes (yuxinlu1's Gemma-4-12B-coder hit 2.4K weekly likes), while **Baidu's Unlimited-OCR** signals renewed interest in vision-language OCR at scale. NVIDIA makes a notable appearance with both LocateAnything-3B (2.4K likes) and their Model Optimizer quantizations (NVFP4), alongside new ASR streaming models. The **Ornith-1.0** family from deepreinforce-ai spans 9B to 397B parameters, suggesting a growing appetite for diverse-size open-weight MoE architectures.

---

## Trending Models

### 🧠 Language Models (LLMs, chat, instruction-tuned)

- **zai-org/GLM-5.2** — zai-org | 2,716 likes | 118,651 downloads  
  The most-liked model this week: a Mixture-of-Experts conversational model for text generation building on the GLM architecture.

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** — HauhauCS | 2,285 likes | 3,248,724 downloads  
  An uncensored, vision-capable MoE fine-tune of Qwen3.6 that has gone viral due to aggressive alignment removal and massive community demand.

- **deepreinforce-ai/Ornith-1.0-397B** — deepreinforce-ai | 130 likes | 1,116 downloads  
  The largest Ornith variant, a 397B-parameter MoE model for text generation pushing the boundaries of scale in open-weight LLMs.

- **deepseek-ai/DeepSeek-V4-Pro-DSpark** — deepseek-ai | 153 likes | 373 downloads  
  The official Pro variant of DeepSeek-V4 with agentic and distributed capabilities (DSpark), building on their v4 research.

- **WeiboAI/VibeThinker-3B** — WeiboAI | 742 likes | 59,337 downloads  
  A compact 3B Qwen2-based model specialized for mathematical reasoning — notable for its size-to-performance ratio.

- **LiquidAI/LFM2.5-230M** — LiquidAI | 132 likes | 12,384 downloads  
  A tiny 230M liquid foundation model exploring small-scale efficiency for text generation tasks.

- **microsoft/FastContext-1.0-4B-SFT** — microsoft | 368 likes | 6,779 downloads  
  A 4B SFT-tuned model designed for long-context applications, part of Microsoft's Explorer SubAgent initiative.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **baidu/Unlimited-OCR** — baidu | 1,164 likes | 295,064 downloads  
  A cutting-edge OCR model from Baidu handling image-to-text extraction at scale — trending due to practical utility and strong performance.

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M** — empero-ai | 506 likes | 52,492 downloads  
  A Qwen3.5-based image-text-to-text model blending Claude-style reasoning with vision capabilities.

- **krea/Krea-2-Turbo** — krea | 329 likes | 27,631 downloads  
  The turbo variant of Krea-2 image generation model, optimized for faster inference while maintaining quality.

- **nvidia/LocateAnything-3B** — nvidia | 2,418 likes | 646,451 downloads  
  An image-feature-extraction model from NVIDIA for general-purpose object localization — extremely popular for its versatility and Nvidia's reputation.

- **MiniMaxAI/MiniMax-M3** — MiniMaxAI | 1,255 likes | 188,314 downloads  
  A multimodal vision-language model supporting image-text-to-text, competing in the crowded VL space.

- **nvidia/nemotron-3.5-asr-streaming-0.6b** — nvidia | 725 likes | 67,419 downloads  
  A compact 0.6B streaming ASR model optimized for real-time speech recognition — signals Nvidia's push into efficient audio.

### 🔧 Specialized Models (code, math, vision/speech, embeddings)

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF** — yuxinlu1 | 2,439 likes | 549,926 downloads  
  The most-liked specialized model: a Gemma-4-based GGUF fine-tune optimized for coding and reasoning with composer2.5 architecture.

- **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF** — yuxinlu1 | 747 likes | 225,822 downloads  
  An agentic-on-terminal variant of the same Gemma-4 family, fine-tuned for autonomous coding and tool use.

- **huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated** — huihui-ai | 138 likes | 7,210 downloads  
  An "abliterated" (uncensored) version of the Gemma-4 coder, removing safety alignments for unrestricted coding use.

- **WeiboAI/VibeThinker-3B** — WeiboAI (also listed under LLMs) — specialized for math reasoning.

- **Chunjiang-Intelligence/DeepSeek-v4-Fable** — Chunjiang-Intelligence | 118 likes | 1,409 downloads  
  A cybersecurity-focused DeepSeek-V4 fine-tune, applying the "Fable" variant to security tasks.

### 📦 Fine-tunes & Quantizations (GGUF, NVFP4, community adapts)

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF** — empero-ai | 699 likes | 831,529 downloads  
  GGUF quantized version of Qwythos-9B for llama.cpp inference — extremely popular due to easy local deployment.

- **unsloth/GLM-5.2-GGUF** — unsloth | 429 likes | 146,023 downloads  
  Unsloth's GGUF quantization of GLM-5.2, enabling efficient inference on consumer hardware.

- **HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced** — HauhauCS | 98 likes | 40,820 downloads  
  A balanced, uncensored QAT quantization of Gemma-4-12B for multimodal vision-text tasks.

- **deepreinforce-ai/Ornith-1.0-35B-GGUF** — deepreinforce-ai | 350 likes | 79,630 downloads  
  GGUF quant of the 35B Ornith model, offering a middle ground between performance and accessibility.

- **nvidia/Qwen3.6-35B-A3B-NVFP4** — nvidia | 367 likes | 5,235,413 downloads  
  NVIDIA's FP4 quantization of Qwen3.6, achieving extreme compression with minimal quality loss — the most-downloaded model this week.

- **nvidia/GLM-5.2-NVFP4** — nvidia | 143 likes | 45,762 downloads  
  NVIDIA's FP4 quant of GLM-5.2, demonstrating cross-architecture quantization capabilities.

- **unsloth/Qwen-AgentWorld-35B-A3B-GGUF** — unsloth | 84 likes | 79,503 downloads  
  GGUF quant of the AgentWorld model, optimized for world-modeling and agent-based tasks.

---

## Ecosystem Signal

The **Qwen 3.5/3.6 MoE family** has become the dominant open-weight ecosystem this week, with representation from original authors (Qwen), community fine-tuners (HauhauCS, empero-ai), and quantization partners (NVIDIA, unsloth). The sheer download volume (5.2M for NVFP4 quant, 3.2M for uncensored fine-tune) indicates these are the go-to base models for experimentation and deployment.

**Gemma-4** is the clear runner-up, particularly in coding and agentic niches. The "fable5-composer2.5" fine-tune family (yuxinlu1, huihui-ai) shows a maturing ecosystem where community creators systematically push base models toward specific use cases (coder, agentic, abliterated). Notably, all three major open Coder models this week are Gemma-4 based.

**Quantization formats are consolidating**: GGUF (via unsloth, llama.cpp) remains the standard for local inference, while NVIDIA's NVFP4 gains traction as an enterprise-grade compressed format. The success of NVIDIA's Qwen3.6 NVFP4 quant (5.2M downloads) suggests hardware-optimized quantizations are becoming a distribution channel of their own.

**Uncensored/abliterated models** (HauhauCS, huihui-ai) continue to attract outsize engagement despite smaller author followings — the community appetite for unrestricted models remains strong, particularly for multimodal and coding domains.

**Ornith-1.0** (deepreinforce-ai) launches a full spectrum from 9B to 397B, a rare case of a single family spanning all compute tiers. Their MIT licensing and endpoint compatibility tags suggest commercial deployment intent.

**Open-weight vs. proprietary**: All 30 trending models are open-weights. NVIDIA (nemotron-3.5 ASR, LocateAnything) and Microsoft (FastContext) are the only "big tech" contributors, but their models are fully open — indicating the center of gravity has decisively shifted to the open ecosystem.

---

## Worth Exploring

**1. nvidia/LocateAnything-3B** — A standout for its 2.4K likes and 646K downloads: a general-purpose localization model that works across diverse visual domains. Its image-feature-extraction pipeline and Nvidia backing make it ideal for both production deployment and research into vision-based grounding.

**2. Qwen/Qwen-AgentWorld-35B-A3B** — The base model behind the AgentWorld family, specifically designed for world-modeling and agentic reasoning. With only 369 likes so far but a dedicated GGUF quant from unsloth already available, this represents an early-stage opportunity to explore next-gen interactive and simulation-based models.

**3. LiquidAI/LFM2.5-230M** — At just 230M parameters, this liquid foundation model offers a unique opportunity to study scaling dynamics and efficiency trade-offs in extremely small LLMs. Its popularity (132 likes) suggests growing interest in tiny models for edge deployment and rapid iteration — a contrarian signal in a MoE-giant week.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*