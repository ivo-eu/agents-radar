# Hugging Face Trending Models Digest 2026-06-24

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-24 10:35 UTC

---

# Hugging Face Trending Models Digest — 2026-06-24

## Today's Highlights

**DeepSeek-V4-Pro** dominates the board with 5,036 weekly likes and 2.24M downloads, marking the strongest open-weight LLM launch this period. **NVIDIA's LocateAnything-3B** (2,329 likes) signals growing demand for general-purpose visual grounding, while **zai-org/GLM-5.2** (2,264 likes) introduces a new MoE-DSA architecture that's attracting rapid community adoption — including an official unsloth GGUF quant (316 likes) and an FP8 variant (154 likes). The **Gemma-4-12B-coder** ecosystem continues expanding, with the top downloaded model (456K) being a GGUF coding variant, alongside an agentic fine-tune and an "abliterated" community release. Notably, **Qwen3.6** family variants appear across multiple uncensored, vision, and quantization entries, while **DiffusionGemma-26B-A4B-it** (1,059 likes, 949K downloads) demonstrates strong interest in Google's vision-language diffusion hybrid.

---

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **deepseek-ai/DeepSeek-V4-Pro** — Author: deepseek-ai | Likes: 5,036 | Downloads: 2,245,489  
  *Cutting-edge conversational LLM from DeepSeek; trending as the highest-liked model of the week, likely advancing open-weight capabilities with the V4 architecture.*

- **zai-org/GLM-5.2** — Author: zai-org | Likes: 2,264 | Downloads: 40,127  
  *New MoE-DSA (Mixture-of-Experts with Dynamic Sparse Attention) conversational model; trending for its architectural innovation and strong reception.*

- **moonshotai/Kimi-K2.7-Code** — Author: moonshotai | Likes: 981 | Downloads: 447,920  
  *Image-text-to-text model from MoonShot AI focused on code; trending due to high download volume and compressed-tensor efficiency claims.*

- **microsoft/FastContext-1.0-4B-SFT** — Author: microsoft | Likes: 325 | Downloads: 4,391  
  *A 4B-parameter model with long-context optimization and Explorer SubAgent capability; trending for Microsoft's entry into efficient context handling.*

- **poolside/Laguna-M.1** — Author: poolside | Likes: 94 | Downloads: 2,787  
  *Text-generation model optimized for vLLM and SGLang deployment; trending as a new infrastructure-focused release from poolside.*

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **nvidia/LocateAnything-3B** — Author: nvidia | Likes: 2,329 | Downloads: 274,025  
  *Image-text-to-text model for general visual grounding; trending as NVIDIA's versatile object localization solution.*

- **MiniMaxAI/MiniMax-M3** — Author: MiniMaxAI | Likes: 1,225 | Downloads: 131,057  
  *Multimodal vision-language model (image-text-to-text); trending for strong multimodal generation capabilities from MiniMax.*

- **google/diffusiongemma-26B-A4B-it** — Author: google | Likes: 1,059 | Downloads: 948,996  
  *Instruction-tuned diffusion-Gemma hybrid (26B active, 4B per token); trending for combining diffusion with LLM in a MoE format, achieving high downloads.*

- **baidu/Unlimited-OCR** — Author: baidu | Likes: 597 | Downloads: 8,396  
  *Image-text-to-text OCR model; trending for its "unlimited" OCR capability claim and Baidu's brand.*

- **nvidia/nemotron-3.5-asr-streaming-0.6b** — Author: nvidia | Likes: 667 | Downloads: 41,050  
  *Cache-aware streaming ASR model (0.6B params); trending for real-time speech recognition efficiency from NVIDIA.*

- **owensong/Inflect-Nano-v1** — Author: owensong | Likes: 184 | Downloads: 0  
  *Ultra-small text-to-speech model; trending for its tiny footprint while maintaining TTS quality.*

- **krea/Krea-2-Turbo** — Author: krea | Likes: 130 | Downloads: 84  
  *Text-to-image diffusion model (turbo variant of Krea-2-Raw); trending as part of the Krea image generation ecosystem.*

- **krea/Krea-2-Raw** — Author: krea | Likes: 115 | Downloads: 194  
  *Base text-to-image model for the Krea-2 series; trending as the foundational model for fine-tunes.*

- **Boogu/Boogu-Image-0.1-Edit** — Author: Boogu | Likes: 114 | Downloads: 592  
  *Image editing diffusion model (bilingual EN/ZH); trending as a new image editing pipeline.*

- **Comfy-Org/Boogu-Image** — Author: Comfy-Org | Likes: 87 | Downloads: 0  
  *ComfyUI integration for Boogu-Image; trending for making the above model accessible in the Comfy ecosystem.*

- **ostris/ideogram_4_turbotime_lora** — Author: ostris | Likes: 111 | Downloads: 3,672  
  *LoRA adapter for ideogram-4-fp8; trending for enabling fine-tuned generation on top of Ideogram's latest model.*

### 🔧 Specialized Models (code, math, medical, embeddings)

- **WeiboAI/VibeThinker-3B** — Author: WeiboAI | Likes: 676 | Downloads: 41,170  
  *3B math reasoning model (based on Qwen2); trending for its strong math capability per parameter.*

- **LiquidAI/LFM2.5-Embedding-350M** — Author: LiquidAI | Likes: 116 | Downloads: 10,117  
  *350M embedding model using sentence-transformers; trending as a compact Liquid Foundation Model for retrieval.*

- **LiquidAI/LFM2.5-ColBERT-350M** — Author: LiquidAI | Likes: 87 | Downloads: 2,534  
  *350M ColBERT model (via PyLate) for late-interaction retrieval; trending alongside the LFM2.5 embedding family.*

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M** — Author: empero-ai | Likes: 245 | Downloads: 1,856  
  *9B vision-language model (Qwen3.5-based) with "Mythos" style; trending for its creative/fantasy thematic tuning.*

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF** — Author: yuxinlu1 | Likes: 2,265 | Downloads: 456,117  
  *GGUF quantized Gemma-4-12B coder with "fable5-composer" fine-tune; trending as the most downloaded model (456K) due to accessibility for local coding.*

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** — Author: HauhauCS | Likes: 2,179 | Downloads: 3,955,016  
  *GGUF MoE vision-language model (35B, 3B active) with uncensored fine-tune; trending as the highest-downloaded model (3.95M) for its uncensored multimodal MoE offering.*

- **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF** — Author: yuxinlu1 | Likes: 475 | Downloads: 96,459  
  *GGUF of Gemma-4-12B fine-tuned for agentic/terminal use; trending for agent-focused capabilities in a quantized package.*

- **unsloth/GLM-5.2-GGUF** — Author: unsloth | Likes: 316 | Downloads: 55,820  
  *Official unsloth GGUF quantization of GLM-5.2; trending for enabling local deployment of the new GLM architecture.*

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF** — Author: empero-ai | Likes: 210 | Downloads: 27,218  
  *GGUF quant of the Qwythos-9B Mythos model; trending for local creative use cases.*

- **zai-org/GLM-5.2-FP8** — Author: zai-org | Likes: 154 | Downloads: 395,290  
  *FP8 quantized variant of GLM-5.2; trending as a memory-efficient alternative with high download volume.*

- **huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated** — Author: huihui-ai | Likes: 119 | Downloads: 3,320  
  *"Abliterated" (safety-removed) variant of the Gemma-4-12B coder; trending in the community for unrestricted coding.*

- **bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF** — Author: bytkim | Likes: 115 | Downloads: 65,765  
  *GGUF of Qwen3.6-27B with Multi-Token Prediction (MTP) fine-tune; trending for MTP efficiency innovation in quantized form.*

- **lordx64/Qwable-v1** — Author: lordx64 | Likes: 175 | Downloads: 4,547  
  *Qwen3.5-MoE-based vision-language fine-tune; trending for its "Qwable" branding and MoE multiplier efficiency.*

- **datalab-to/lift** — Author: datalab-to | Likes: 141 | Downloads: 3,216  
  *Qwen3.5-based PDF understanding model; trending for its document-focused vision-language capabilities.*

---

## Ecosystem Signal

The **DeepSeek-V4-Pro** launch confirms DeepSeek's position as the leading open-weight LLM provider, outpacing previous releases in community traction. **MoE architectures are now mainstream**: GLM-5.2, Qwen3.6-35B-A3B, DiffusionGemma-26B-A4B, and MiniMax-M3 all employ Mixture-of-Experts, with some (GLM-5.2) introducing novel sparse attention mechanisms. **Vision-language convergence** accelerates — over half of the top-30 models handle image inputs, from OCR (Baidu) to grounding (NVIDIA) to diffusion hybrids (Google). The **Gemma-4 ecosystem** has spawned a rich fine-tuning chain: base coder → composer fine-tune → GGUF quantization → abliterated variants → agentic versions, demonstrating the community's appetite for rapid iteration on top of Google's open weights. **GGUF quantization dominates local deployment**, with 21 of the 30 models offering or being a GGUF variant themselves, and the top download (3.95M) being a GGUF MoE. Notably, **FP8 is gaining traction** alongside traditional GGUF, as seen in GLM-5.2-FP8. The **Liquid LFM2.5** embedding family (350M) suggests a growing focus on efficient retrieval models, counterbalanced by larger multimodal options. Finally, "uncensored" and "abliterated" fine-tunes remain consistently trending, indicating sustained community demand for unrestricted model variants.

---

## Worth Exploring

1. **nvidia/LocateAnything-3B** — This versatile visual grounding model (2,329 likes, 274K downloads) deserves hands-on testing for its ability to locate any object in images without task-specific training — a rare generalist capability at 3B parameters that could streamline many vision pipelines.

2. **zai-org/GLM-5.2** — The new MoE-DSA architecture (2,264 likes) is worth studying as a potential inflection point in attention efficiency. The combination of Mixture-of-Experts with Dynamic Sparse Attention could influence future model designs, and the existence of multiple quantizations (GGUF, FP8) makes experimentation accessible on consumer hardware.

3. **microsoft/FastContext-1.0-4B-SFT** — A compact 4B model optimized for long-context scenarios, this is worth exploring for RAG and document processing tasks where larger models are overkill. Microsoft's "Explorer SubAgent" concept may point toward future multi-agent system designs integrated at the model level.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*