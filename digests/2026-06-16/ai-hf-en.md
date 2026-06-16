# Hugging Face Trending Models Digest 2026-06-16

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-16 05:20 UTC

---

# Hugging Face Trending Models Digest — 2026-06-16

## Today's Highlights

This week’s trending models showcase the rapid convergence of multimodal and MoE architectures, with **DeepSeek-V4-Pro** (4.9k likes, 2.9M downloads) leading the pack as a pure text-generation powerhouse. **NVIDIA’s LocateAnything-3B** (2k likes) introduces a novel spatial understanding capability, while **HauhauCS/Qwen3.6-35B-A3B-Uncensored** (1.9k likes, 2.7M downloads) reflects sustained community interest in uncensored, fine-tuned MoE models. Google’s **Gemma-4 family** appears in multiple forms (base, instruction-tuned, quantized, coder, and “obliterated” fine-tunes), indicating a thriving ecosystem around this unified any-to-any architecture. Audio and video generation also see notable entries from **ideogram-ai** (text-to-image), **zai-org** (character animation), and **bosonai** (TTS).

## Trending Models

### 🧠 Language Models
- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** by deepseek-ai (Likes: 4,867, Downloads: 2,934,763) — A large 4th-generation conversational LLM that dominates the top spot with exceptional reasoning and chat performance.
- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** by microsoft (Likes: 113, Downloads: 13) — A small 4B parameter model fine-tuned for ultra-long context windows, ideal for retrieval and agent tasks.
- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** by CohereLabs (Likes: 392, Downloads: 11,145) — A compact MoE code LLM optimized for instruction-following and conversational coding.

### 🎨 Multimodal & Generation
- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** by google (Likes: 892, Downloads: 311,788) — A 26B-parameter diffusion-based multimodal model for image-text-to-text tasks, blending generative and conversational capabilities.
- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** by MiniMaxAI (Likes: 854, Downloads: 14,312) — A multimodal MoE model supporting image, text, and agentic workflows.
- **[prefeitura-rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)** by prefeitura-rio (Likes: 304, Downloads: 188,723) — An open-weight 397B MoE vision-language model based on Qwen3.5, designed for conversational multimodal tasks.
- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** by google (Likes: 1,038, Downloads: 1,160,435) — Instruction-tuned version of Gemma-4-12B, supporting any-to-any (image, text, audio) understanding and generation.
- **[google/gemma-4-12B](https://huggingface.co/google/gemma-4-12B)** by google (Likes: 553, Downloads: 250,498) — Base any-to-any model of the Gemma-4 family, serving as the foundation for many fine-tunes.
- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** by nex-agi (Likes: 288, Downloads: 3,681) — A MoE Qwen3.5-based text-generation model with multimodal tags, optimized for reasoning and agent use.
- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** by nex-agi (Likes: 220, Downloads: 8,260) — Smaller version of Nex-N2-Pro, balancing performance and efficiency for lightweight deployment.
- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** by ideogram-ai (Likes: 548, Downloads: 10,748) — FP8 quantized version of Ideogram 4, a state-of-the-art text-to-image diffusion model.
- **[ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4)** by ideogram-ai (Likes: 345, Downloads: 4,224) — NF4 quantized variant of Ideogram 4, offering additional memory savings for image generation.
- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** by zai-org (Likes: 191, Downloads: 0) — A pose-driven character animation model that generates videos from images, leveraging diffusion.
- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** by bosonai (Likes: 445, Downloads: 38,429) — A 4B parameter multimodal TTS model built on Qwen3, capable of both text generation and high-quality speech synthesis.
- **[Zyphra/ZONOS2](https://huggingface.co/Zyphra/ZONOS2)** by Zyphra (Likes: 91, Downloads: 414) — A lightweight Apache-2.0 TTS model for efficient text-to-speech generation.
- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** by nvidia (Likes: 424, Downloads: 5,200) — A streaming ASR model with cache-aware architecture for low-latency speech recognition.

### 🔧 Specialized Models
- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** by moonshotai (Likes: 753, Downloads: 56,750) — A compressed image-feature-extraction model specialized for code understanding and generation.
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** by nvidia (Likes: 2,062, Downloads: 86,968) — A 3B parameter model for spatial localization and feature extraction in images, trending for its novel zero-shot pointing capability.

### 📦 Fine-tunes & Quantizations
- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** by yuxinlu1 (Likes: 718, Downloads: 20,207) — A GGUF-quantized Gemma-4 fine-tune focused on coding and reasoning with composer integration.
- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** by unsloth (Likes: 276, Downloads: 107,243) — Unsloth’s GGUF conversion of the DiffusionGemma multimodal model for efficient local inference.
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** by HauhauCS (Likes: 1,859, Downloads: 2,697,882) — A heavily downloaded GGUF fine-tune of Qwen3.6 MoE, uncensored and aggression-tuned, popular in the community.
- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** by OBLITERATUS (Likes: 325, Downloads: 70,732) — An “obliterated” uncensored fine-tune of Gemma-4-12B, pushing safety limits.
- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** by Jackrong (Likes: 202, Downloads: 62,469) — A GGUF-quantized Qwopus3.6 vision-language coder model with multi-task prediction.
- **[Jackrong/Qwopus3.6-27B-v2-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-v2-MTP-GGUF)** by Jackrong (Likes: 312, Downloads: 184,446) — Updated version of the Qwopus3.6 coder GGUF, with improved performance.
- **[unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF)** by unsloth (Likes: 103, Downloads: 9,327) — Unsloth’s GGUF quantization of Kimi-K2.7-Code for local code generation.
- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** by unsloth (Likes: 618, Downloads: 980,781) — The most popular Gemma-4 GGUF, enabling local deployment of the instruction-tuned model.
- **[unsloth/gemma-4-12B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF)** by unsloth (Likes: 243, Downloads: 288,390) — A quantization-aware training version of Gemma-4-12B-it GGUF for higher accuracy at low bitwidths.
- **[unsloth/MiniMax-M3-GGUF](https://huggingface.co/unsloth/MiniMax-M3-GGUF)** by unsloth (Likes: 84, Downloads: 14,799) — GGUF conversion of MiniMax-M3 multimodal MoE for agentic workflows.
- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** by DavidAU (Likes: 359, Downloads: 369,526) — An extreme community fine-tune merging multiple base models into an uncensored, code-optimized GGUF.
- **[RazzzHF/Realism_Engine_Ideogram_4](https://huggingface.co/RazzzHF/Realism_Engine_Ideogram_4)** by RazzzHF (Likes: 99, Downloads: 0) — A fine-tune of Ideogram 4 focused on photorealism, yet to accumulate downloads.

## Ecosystem Signal

The current landscape is defined by three converging trends: **MoE architectures, multimodal unification, and aggressive quantization.** MoE models like DeepSeek-V4-Pro, Qwen3.6, Gemma-4, and North-Mini-Code dominate both pure language and vision-language categories, enabling high capacity with efficient inference. The line between “language model” and “multimodal model” continues to blur — over half of the top 30 models support at least image understanding, and Google’s Gemma-4 introduces an any-to-any pipeline covering text, image, and audio.

Open-weight momentum remains strong, with Google, DeepSeek, NVIDIA, and Cohere releasing permissively licensed models. On the community side, **Unsloth** is the primary quantization engine, converting most popular releases into GGUF format for local deployment. Fine-tuning for uncensored or “aggressive” behavior (e.g., HauhauCS, DavidAU, OBLITERATUS) remains a vibrant subculture, drawing massive download numbers. Meanwhile, emerging modalities like video generation (SCAIL-2) and streaming ASR (nemotron) signal growing interest beyond text and images.

## Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — This 3B model offers a unique spatial reasoning capability (pointing and localization) rarely seen in general-purpose vision models. It’s compact, well-received, and opens new use cases in robotics, UI automation, and visual grounding.

2. **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — One of the few pose-driven image-to-video models trending, SCAIL-2 represents the cutting edge of character animation diffusion. Ideal for researchers exploring controllable video generation.

3. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — Despite low downloads, this small 4B model from Microsoft is purpose-built for long-context and agent sub-tasks. Its efficiency and novel training recipe make it a promising baseline for resource-constrained deployments.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*