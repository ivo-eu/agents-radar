# Hugging Face Trending Models Digest 2026-06-13

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-13 10:15 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-06-13

## Today's Highlights

**DeepSeek-V4-Pro** dominates the leaderboard this week with 4,803 weekly likes and over 3.2 million downloads, signaling the community's hunger for powerful open-weight reasoning models. **Nvidia's LocateAnything-3B** (1,935 likes) and the uncensored **Qwen3.6-35B-A3B** variant (1,739 likes) both saw explosive growth, reflecting strong interest in specialized vision-language and unconstrained conversational models. Google's **Gemma 4 family** remains a central ecosystem force, with multiple fine-tunes, quantized variants (GGUF), and abliterated versions appearing across the top 30. The trend toward **Mixture-of-Experts (MoE)** architectures is unmistakable, with nearly half the list featuring MoE-based models, and multimodal pipelines (image-text-to-text, any-to-any) now dominate over pure text-only models.

---

## Trending Models

### 🧠 Language Models (LLMs, Chat Models, Instruction-Tuned)

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — *deepseek-ai* · ❤️ 4,803 · ⬇️ 3,250,404  
  The week's runaway hit: a next-generation conversational LLM with strong reasoning capabilities, driving massive adoption.

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — *moonshotai* · ❤️ 418 · ⬇️ 1,689  
  A compressed MoE code model from Kimi's lineage, optimized for coding tasks with reduced parameter footprint.

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — *CohereLabs* · ❤️ 339 · ⬇️ 6,533  
  Cohere's lightweight MoE code model, positioned for efficient code generation and conversational coding assistance.

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — *nex-agi* · ❤️ 230 · ⬇️ 3,092  
  A Qwen3.5-based MoE model for text generation with vision support, blending code and multimodal capabilities.

- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** — *nex-agi* · ❤️ 182 · ⬇️ 3,760  
  Smaller sibling of Nex-N2-Pro, offering a compact MoE alternative with similar architecture.

### 🎨 Multimodal & Generation (Image, Video, Audio, Text-to-X)

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — *google* · ❤️ 646 · ⬇️ 92,080  
  Google's massive diffusion-Gemma hybrid (26B, 4B active) for image-text-to-text, a flagship multimodal MoE.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — *nvidia* · ❤️ 1,935 · ⬇️ 69,443  
  Nvidia's focused vision model for locating and identifying objects in images—trending rapidly for its 3B efficiency.

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — *google* · ❤️ 981 · ⬇️ 1,005,883  
  The instruction-tuned variant of Gemma 4, an any-to-any multimodal model with massive download volume.

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — *MiniMaxAI* · ❤️ 334 · ⬇️ 1,031  
  A multimodal vision-language model from MiniMax, targeting general image-text-to-text tasks.

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — *ideogram-ai* · ❤️ 506 · ⬇️ 6,535  
  Ideogram's latest text-to-image diffusion model in FP8 precision, balancing quality and efficiency.

- **[ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4)** — *ideogram-ai* · ❤️ 329 · ⬇️ 3,276  
  NF4 quantized version of Ideogram-4 for lower memory usage on consumer hardware.

- **[google/magenta-realtime-2](https://huggingface.co/google/magenta-realtime-2)** — *google* · ❤️ 186 · ⬇️ 7,331  
  Google's real-time text-to-audio model, optimized for TFLite deployment and low-latency music/audio generation.

- **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — *ByteDance* · ❤️ 231 · ⬇️ 426  
  ByteDance's image-text-to-video model with a renderer approach, pushing video generation from static inputs.

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — *bosonai* · ❤️ 395 · ⬇️ 32,162  
  A 4B TTS model based on Qwen3 architecture, enabling high-quality speech synthesis.

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — *zai-org* · ❤️ 146 · ⬇️ 0  
  A diffusion-based character animation model for pose-driven image-to-video generation.

- **[XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash](https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash)** — *XiaomiMiMo* · ❤️ 101 · ⬇️ 3,590  
  Xiaomi's agent-focused text generation model with FP4 and DFlash optimization.

### 🔧 Specialized Models (Code, Math, Medical, Embeddings, Speech)

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — *nvidia* · ❤️ 395 · ⬇️ 3,975  
  Nvidia's cache-aware streaming ASR model (0.6B) for real-time speech recognition.

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — *Jackrong* · ❤️ 129 · ⬇️ 11,291  
  A 27B vision-language coder model in GGUF format, weighted around multi-task programming.

- **[MisoLabs/MisoTTS](https://huggingface.co/MisoLabs/MisoTTS)** — *MisoLabs* · ❤️ 195 · ⬇️ 0  
  A newly trending TTS model for speech synthesis with PyTorch/safetensors, zero downloads yet strong interest.

### 📦 Fine-Tunes & Quantizations (Community Fine-Tunes, GGUF, AWQ)

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — *HauhauCS* · ❤️ 1,739 · ⬇️ 2,411,202  
  An aggressively uncensored MoE fine-tune of Qwen3.6 (35B, 3B active) with vision, extremely popular for unrestricted use.

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — *unsloth* · ❤️ 224 · ⬇️ 42,885  
  Unsloth's GGUF conversion of Google's DiffusionGemma, bringing efficient local inference to this massive multimodal model.

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — *unsloth* · ❤️ 570 · ⬇️ 872,895  
  The most downloaded Gemma 4 quantized variant; Unsloth's optimized GGUF for Gemma-4-12B-it.

- **[unsloth/gemma-4-12B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF)** — *unsloth* · ❤️ 207 · ⬇️ 227,830  
  QAT-optimized GGUF variant of Gemma-4-12B-it for improved accuracy at lower bit depths.

- **[unsloth/gemma-4-26B-A4B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF)** — *unsloth* · ❤️ 148 · ⬇️ 260,757  
  QAT GGUF of the larger 26B Gemma-4 variant, balancing size and performance.

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — *OBLITERATUS* · ❤️ 260 · ⬇️ 50,289  
  An abliterated/uncensored fine-tune of Gemma-4-12B, popular for removing safety filters.

- **[huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated)** — *huihui-ai* · ❤️ 149 · ⬇️ 8,270  
  Another abliterated Gemma 4 variant from huihui-ai, contributing to the uncensoring trend.

- **[google/gemma-4-12B-it-qat-q4_0-gguf](https://huggingface.co/google/gemma-4-12B-it-qat-q4_0-gguf)** — *google* · ❤️ 136 · ⬇️ 210,048  
  Google's own official QAT quantization to q4_0 GGUF, providing a first-party efficient deployment option.

- **[Comfy-Org/Ideogram-4](https://huggingface.co/Comfy-Org/Ideogram-4)** — *Comfy-Org* · ❤️ 142 · ⬇️ 0  
  ComfyUI-optimized version of Ideogram-4 for node-based image generation workflows.

- **[RazzzHF/Realism_Engine_Ideogram_4](https://huggingface.co/RazzzHF/Realism_Engine_Ideogram_4)** — *RazzzHF* · ❤️ 87 · ⬇️ 0  
  A realist fine-tune of Ideogram-4, targeting photorealistic outputs.

---

## Ecosystem Signal

**MoE is the new standard.** Over 40% of trending models this week use Mixture-of-Experts architectures—from Google's DiffusionGemma and Gemma-4 series to Qwen-based variants and Cohere's North-Mini. The community is increasingly choosing sparse activation over dense models. **Google's Gemma 4 ecosystem dominates** with the widest array of fine-tunes, quantizations (by Unsloth, OBLITERATUS, huihui-ai), and official releases, signaling a platform-like strategy. **DeepSeek-V4-Pro** proves that the open-weight community still craves frontier-class LLMs, especially when paired with direct developer access. **Uncensored/abliterated variants continue to thrive**, reflecting user demand for unrestricted model behavior. Quantization is now table stakes: nearly every major release has a GGUF variant within days, led by Unsloth's polished conversions. The shift toward **multimodal inputs** (image-text-to-text, any-to-any) is definitive—pure text-only models are increasingly rare among high-interest releases. Emerging video generation (Bernini-R, SCAIL-2) and real-time audio (Magenta Realtime 2, Nemotron ASR) point to the next frontier.

---

## Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — A remarkably efficient 3B vision model that achieved nearly 2K likes in a single week. It's a strong candidate for deployment on edge devices while still offering robust feature extraction and localization. Worth trying for anyone building multimodal search or object detection pipelines.

2. **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — A cutting-edge image-text-to-video model with a renderer-based approach. With only 426 downloads but 231 likes, it's early-stage but promising. Researchers interested in controllable video generation from static images should study its architecture.

3. **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — The most downloaded multimodal model this week (1M+), and the foundation for numerous fine-tunes. Testing the base model gives you a reference point for evaluating the many community variants, and its any-to-any pipeline makes it uniquely versatile for prototyping.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*