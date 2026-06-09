# Hugging Face Trending Models Digest 2026-06-09

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-09 04:18 UTC

---

# Hugging Face Trending Models Digest — June 9, 2026

## 1. Today’s Highlights

DeepSeek’s V4 family dominates the leaderboard with **DeepSeek-V4-Pro** (4,724 likes, 5.4M downloads) and **V4-Flash** (1,449 likes, 3.3M downloads) driving massive community engagement. Google’s **Gemma 4** series is the highest-volume new release — both the base `any-to-any` 12B model and multiple quantized variants (by Unsloth and Google itself) are climbing fast. Meanwhile, **NVIDIA** is active on multiple fronts: the locate-anything vision model, streaming ASR, the ultra-large Nemotron-3 MoE (550B, 55B active), and the Cosmos3 omni-world-model family. A notable dark horse is **SulphurAI/Sulphur-2-base** (1,602 likes), a text-to-video model derived from Lightricks’ LTX-2.3 that has already accumulated 1.7M downloads.

## 2. Trending Models by Category

### 🧠 Language Models (LLMs, chat, instruction-tuned)

- **deepseek-ai/DeepSeek-V4-Pro** — *deepseek-ai* | Likes: 4,724 | Downloads: 5,399,597  
  The flagship reasoning LLM from DeepSeek, featuring ultra-long context and state-of-the-art performance; trending as the most popular model overall.

- **deepseek-ai/DeepSeek-V4-Flash** — *deepseek-ai* | Likes: 1,449 | Downloads: 3,262,529  
  A smaller, faster variant of V4 optimized for speed, gaining traction for deployment and fine-tuning.

- **sapientinc/HRM-Text-1B** — *sapientinc* | Likes: 728 | Downloads: 163,953  
  A compact text-generation model (1B) designed for human resource management tasks, trending for its niche enterprise focus.

- **LiquidAI/LFM2.5-8B-A1B** — *LiquidAI* | Likes: 551 | Downloads: 135,131  
  An 8B-parameter MoE model with 1B active parameters, balancing efficiency and quality; part of Liquid’s LFM2.5 series.

- **nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16** — *nvidia* | Likes: 167 | Downloads: 55,910  
  NVIDIA’s massive 550B-parameter MoE model (55B active), pushing the frontier of scale and serving as a base for further quantization.

- **JetBrains/Mellum2-12B-A2.5B-Thinking** — *JetBrains* | Likes: 260 | Downloads: 17,448  
  A 12B MoE model (2.5B active) fine-tuned for reasoning and code tasks, notable for its name-brand developer tooling background.

- **nex-agi/Nex-N2-Pro** — *nex-agi* | Likes: 128 | Downloads: 716  
  A Qwen3.5-derived MoE model for text generation, early in its lifecycle but gaining interest as a community MoE.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **nvidia/LocateAnything-3B** — *nvidia* | Likes: 1,636 | Downloads: 121,594  
  A 3B vision-language model for open-vocabulary object detection and segmentation; trending for its “locate anything” capability.

- **google/gemma-4-12B-it** — *google* | Likes: 756 | Downloads: 554,173  
  Google’s latest multimodal (any-to-any) instruction-tuned model, supporting text, images, and audio inputs/outputs.

- **unsloth/gemma-4-12b-it-GGUF** — *unsloth* | Likes: 504 | Downloads: 645,263  
  The most popular Gemma-4 quantization (GGUF), enabling local inference even on consumer hardware; highest downloads among Gemma-4 variants.

- **ideogram-ai/ideogram-4-fp8** — *ideogram-ai* | Likes: 400 | Downloads: 5,495  
  Ideogram’s latest text-to-image model in FP8 precision, offering high-quality generation while reducing memory footprint.

- **stepfun-ai/Step-3.7-Flash** — *stepfun-ai* | Likes: 352 | Downloads: 45,535  
  A vision-language model (Step 3.7 series) optimized for fast multimodal inference.

- **ByteDance/Bernini-R** — *ByteDance* | Liked: 186 | Downloads: 278  
  An image-text-to-video model from ByteDance, using a renderer architecture (arxiv:2605.22344); early stage but attracting research attention.

- **SulphurAI/Sulphur-2-base** — *SulphurAI* | Likes: 1,602 | Downloads: 1,707,062  
  A text-to-video model derived from LTX-2.3, fine-tuned on a large dataset; extremely popular due to its quality and easy deployment via Diffusers.

- **joyaiecho/JoyAI-Echo** — *jdopensource* | Likes: 106 | Downloads: 4,053  
  A text-to-video + audio-video generation model based on LTX-Video, supporting simultaneous audio-video output.

- **bosonai/higgs-audio-v3-tts-4b** — *bosonai* | Likes: 254 | Downloads: 15,005  
  A 4B-parameter TTS model built on Qwen3 backbone, offering high-quality speech synthesis.

- **MisoLabs/MisoTTS** — *MisoLabs* | Likes: 157 | Downloads: 0  
  A PyTorch-based TTS model focused on speech synthesis and voice cloning; early release but no downloads yet (likely just announced).

- **google/magenta-realtime-2** — *google* | Likes: 153 | Downloads: 17,531  
  Google’s real-time text-to-audio model (TF-Lite), capable of generating music and sound effects interactively.

- **nvidia/nemotron-3.5-asr-streaming-0.6b** — *nvidia* | Likes: 298 | Downloads: 3,957  
  A lightweight streaming ASR model (0.6B) with cache-aware optimization, trending for real-time transcription use cases.

- **PaddlePaddle/PaddleOCR-VL-1.6** — *PaddlePaddle* | Likes: 278 | Downloads: 9,924  
  A vision-language OCR model built on ERNIE 4.5, providing end-to-end OCR with natural language understanding.

- **nvidia/Cosmos3-Nano** and **nvidia/Cosmos3-Super** — *nvidia* | Likes: 207 / 159 | Downloads: 34,104 / 27,548  
  Two members of NVIDIA’s Cosmos3 family (omni-world models), with Nano for lightweight applications and Super for high fidelity.

### 🔧 Specialized Models (code, math, medical, embeddings, etc.)

- *No strictly specialized models in the top 30, but **PaddleOCR-VL** and **HRM-Text-1B** could be considered domain-specific.*

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** — *HauhauCS* | Likes: 1,558 | Downloads: 3,036,465  
  A heavily modified Qwen3.6 MoE (35B total, 3B active) with uncensored fine-tuning and aggressive prompt processing; extremely popular for roleplay and unrestricted use.

- **unsloth/Qwen3.6-27B-MTP-GGUF** — *unsloth* | Likes: 697 | Downloads: 1,186,648  
  Unsloth’s GGUF quantization of Qwen3.6-27B with MTP (multi-turn prompt) tuning, enabling efficient local inference.

- **unsloth/gemma-4-12B-it-qat-GGUF** — *unsloth* | Likes: 150 | Downloads: 121,399  
  A quantized (QAT) GGUF version of the Gemma-4-12B-it, further reducing model size while maintaining quality.

- **unsloth/gemma-4-26B-A4B-it-qat-GGUF** — *unsloth* | Likes: 105 | Downloads: 87,455  
  The 26B (4B active) MoE variant of Gemma-4-it, quantized for consumer hardware; part of the broader Gemma-4 quantization wave.

- **google/gemma-4-12B-it-qat-q4_0-gguf** — *google* | Likes: 102 | Downloads: 52,386  
  Google’s own official QAT GGUF quantization (q4_0) of Gemma-4-12B-it, competing with Unsloth’s offerings.

- **nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4** — *nvidia* | Likes: 145 | Downloads: 66,219  
  NVIDIA’s official NVFP4 (4-bit floating point) quantization of the 550B Nemotron model, drastically reducing memory requirements.

- **ideogram-ai/ideogram-4-nf4** — *ideogram-ai* | Likes: 266 | Downloads: 4,963  
  A 4-bit normal float quantization of Ideogram-4, enabling high-quality image generation on lower VRAM GPUs.

## 3. Ecosystem Signal

Three macro-trends stand out this week. **Mixture-of-Experts (MoE) continues its dominance** — the top three LLM releases (Nemotron-3 Ultra 550B-A55B, LFM2.5-8B-A1B, Mellum2-12B-A2.5B) and the top multimodal model (Qwen3.6-35B-A3B) all use MoE architectures, reflecting the industry’s pivot toward parameter efficiency. **Open-weight releases from major labs are accelerating**: DeepSeek (V4 Pro/Flash), Google (Gemma-4), NVIDIA (Nemotron-3, Cosmos3, LocateAnything), and ByteDance (Bernini-R) have all dropped fully open models in a single week, signaling a strong open-weight momentum. **Quantization is mainstream** — Unsloth’s GGUF pipeline for Gemma-4 and Qwen3.6, plus NVIDIA’s own NVFP4 and Google’s QAT GGUF, show that model publishers now treat quantized variants as first-class artifacts. On the multimodal side, **text-to-video is exploding** (Sulphur-2, JoyAI-Echo) alongside **any-to-any models** (Gemma-4 unified pipeline), suggesting the community is ready for interactive, cross-modal agents. The high download counts for uncensored MoE variants (3M downloads for HauhauCS) highlight a sustained demand for unrestricted creative and roleplay use.

## 4. Worth Exploring

1. **nvidia/LocateAnything-3B** — This model turns image-text-to-text into a precise segmentation tool. It’s trending for good reason: open-vocabulary grounding is a key building block for robotics, VQA, and data labeling. At only 3B parameters, it’s lightweight enough to experiment with on a single GPU.

2. **JetBrains/Mellum2-12B-A2.5B-Thinking** — With only 2.5B active parameters, this 12B total MoE offers reasoning performance rivaling much larger dense models. JetBrains’ involvement hints at strong code-related capabilities. It’s an excellent candidate for studying MoE reasoning efficiency.

3. **SulphurAI/Sulphur-2-base** — For anyone interested in text-to-video generation, this model (based on Lightricks LTX-2.3) is already battle-tested with 1.7M downloads. It’s the most accessible high-quality video generation model today, and its Diffusers integration makes it simple to deploy and fine-tune.

---

*All model links and metadata as of 2026-06-09 from Hugging Face Hub.*