# Hugging Face Trending Models Digest 2026-07-31

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-31 00:15 UTC

---

# 🚀 Hugging Face Trending Models Digest – 2026-07-31

## 1. Today's Highlights

The week is dominated by **multimodal LLMs**, with **Kimi-K3** (8,993 likes) from moonshotai leading the pack as a compressed-tensor vision-language model. **Qwen3.6-35B-A3B** continues its meteoric rise, now with 6.1M downloads, while its uncensored fine-tune by HauhauCS gains 3.2K likes, highlighting strong demand for "freedom" variants. **GLM-5.2** (4.7K likes) from zai-org emerges as a powerful MoE conversational model, and **baidu/Unlimited-OCR** (3.6K likes, 2.6M downloads) signals a boom in high-accuracy OCR systems. Extreme quantization is also trending: **Ternary-Bonsai-27B** (2-bit, 1.1K likes) proves community interest in ultra-lightweight LLMs.

## 2. Trending Models

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** – zai-org, 4.7K likes, 1.5M downloads. A large MoE conversational model with strong Chinese/English alignment, trending for its DSA-based architecture and open-weight release.

- **[upstage/Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B)** – upstage, 702 likes, 12K downloads. A 250B-parameter open-weight LLM (pipeline text-generation), gaining traction for competitive performance in the frontier model space.

- **[Nanbeige/Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B)** – Nanbeige, 582 likes, 25K downloads. A compact 3B LLM optimized for efficiency, popular for edge deployment and fine-tuning experiments.

- **[fdtn-ai/antares-1b](https://huggingface.co/fdtn-ai/antares-1b)** – fdtn-ai, 240 likes, 9.8K downloads. A 1B security-focused LLM using GraniteMoEHybrid, tapping into the niche for safe, small models.

- **[amd/Instella-MoE-16B-A3B-Think](https://huggingface.co/amd/Instella-MoE-16B-A3B-Think)** – amd, 94 likes, 1.3K downloads. AMD’s Mixture-of-Experts reasoning model (16B total, 3B active) targeting thoughtful inference, implying growing AMD ecosystem presence.

### 🎨 Multimodal & Generation (Vision, Audio, Text-to-X)

- **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)** – moonshotai, 8.99K likes, 388K downloads. A state-of-the-art image-text-to-text model with compressed tensors, trending as the week’s most liked model overall.

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** – baidu, 3.6K likes, 2.6M downloads. High-accuracy OCR model (image-text-to-text) from Baidu, trending due to vast real-world utility and multilingual support.

- **[thinkingmachines/Inkling](https://huggingface.co/thinkingmachines/Inkling)** – thinkingmachines, 1.7K likes, 46K downloads. A new multimodal conversational model, likely based on novel architecture, attracting early adopters.

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)** – Qwen, 2.6K likes, 6.1M downloads. The flagship MoE multimodal model (35B total, 3B active) from Alibaba, the week’s most downloaded, cementing Qwen’s dominance.

- **[microsoft/Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B)** – microsoft, 222 likes, 2.3K downloads. A computer-use oriented multimodal model (image-text-to-text), indicating Microsoft’s push into agentic vision.

- **[microsoft/Mage-VL](https://huggingface.co/microsoft/Mage-VL)** – microsoft, 120 likes, 3K downloads. A dedicated multimodal vision-language model, likely supporting Comfy-Org pipelines (see Mage-Flow).

- **[ATH-MaaS/OvisOCR2](https://huggingface.co/ATH-MaaS/OvisOCR2)** – ATH-MaaS, 350 likes, 57K downloads. An OCR-specific model based on Qwen3.5, competing with baidu’s Unlimited-OCR in the document AI space.

- **[owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)** – owensong, 319 likes, 1.1K downloads. A lightweight local TTS model optimized for CPU/edge, trending as offline voice synthesis gains popularity.

- **[owensong/Inflect-Nano-v2](https://huggingface.co/owensong/Inflect-Nano-v2)** – owensong, 119 likes, 654 downloads. Even smaller TTS sibling of Inflect-Micro, aimed at ultra-lightweight deployment.

- **[Audio8/Audio8-TTS-Preview-0.6b](https://huggingface.co/Audio8/Audio8-TTS-Preview-0.6b)** – Audio8, 125 likes, 225 downloads. A new TTS model using ArkTTS, previewing high-quality speech synthesis.

- **[microsoft/VibeVoice-ASR-BitNet](https://huggingface.co/microsoft/VibeVoice-ASR-BitNet)** – microsoft, 120 likes, 3.9K downloads. A BitNet-based automatic speech recognition model (GGUF/GGML), pushing efficient ASR.

### 🔧 Specialized Models (Code, OCR, Security)

- **[Kwaipilot/KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev)** – Kwaipilot, 350 likes, 9.2K downloads. A code generation model (image-text-to-text pipeline) built on Qwen3.5 MoE, trending for developer tooling.

- **[poolside/Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1)** – poolside, 846 likes, 73K downloads. A text-generation model optimized for software engineering tasks (specialized code model), seeing strong adoption.

### 📦 Fine-tunes & Quantizations (Community Fine-tunes, GGUF, AWQ)

- **[DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF)** – DavidAU, 1.0K likes, 956K downloads. An uncensored, heavily fine-tuned GGUF variant of Qwen3.6-27B, trending among community users seeking “unfiltered” responses.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** – HauhauCS, 3.2K likes, 1.8M downloads. Another uncensored Qwen3.6 MoE fine-tune, extremely popular for its “aggressive” prompt style and high performance.

- **[prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)** – prism-ml, 1.1K likes, 698K downloads. A 2-bit ternary quantization of a 27B model (GGUF), trending as a proof-of-concept for extreme compression while retaining utility.

- **[unsloth/Kimi-K3-GGUF](https://huggingface.co/unsloth/Kimi-K3-GGUF)** – unsloth, 206 likes, 12K downloads. Official GGUF conversion of Kimi-K3, enabling local inference with llama.cpp.

- **[unsloth/Kimi-K3](https://huggingface.co/unsloth/Kimi-K3)** – unsloth, 175 likes, 766 downloads. The base Kimi-K3 weights re-uploaded by Unsloth for compatibility.

- **[unsloth/Laguna-S-2.1-GGUF](https://huggingface.co/unsloth/Laguna-S-2.1-GGUF)** – unsloth, 258 likes, 159K downloads. GGUF version of the Laguna code model, making it accessible for CPU inference.

- **[nota-ai/Solar-Open2-250B-Nota-NVFP4](https://guggingface.co/nota-ai/Solar-Open2-250B-Nota-NVFP4)** – nota-ai, 147 likes, 7.8K downloads. NVFP4 quantization of Solar-Open2-250B, a cutting-edge 4-bit floating-point format for high-quality compression.

- **[DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF)** – DavidAU, 156 likes, 248K downloads. Another uncensored GGUF fine-tune, this time of Qwen3.5-9B, with custom matrix merging.

- **[LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF)** – LuffyTheFox, 241 likes, 162K downloads. Yet another uncensored Qwen3.6 MoE variant, integrating Hermes-style alignment.

- **[EschaLabs/Qwen3.6-35B-A3B-Escha-W2](https://huggingface.co/EschaLabs/Qwen3.6-35B-A3B-Escha-W2)** – EschaLabs, 92 likes, 201 downloads. A new weight-2-bit quantized version of Qwen3.6 MoE, representing experimental ultra-low precision.

- **[Comfy-Org/Mage-Flow](https://huggingface.co/Comfy-Org/Mage-Flow)** – Comfy-Org, 96 likes, 45K downloads. A diffusion model finetune (image generation) tied to Microsoft’s Mage-VL, trending for ComfyUI integration.

## 3. Ecosystem Signal

**Model Family Momentum:** The **Qwen3.6** family (especially the 35B-A3B MoE variant) is the current ecosystem centerpiece, with a large number of uncensored fine-tunes, quantizations (GGUF, W2, NVFP4), and derivative models (by DavidAU, HauhauCS, EschaLabs, LuffyTheFox). **Kimi-K3** from moonshotai is the top performer by likes, signaling that new entrants from Chinese labs continue to command attention. **GLM-5.2** (zai-org) is a strong alternative in the MoE space, demonstrating that “GLM” branding retains pull. **Solar-Open2-250B** from upstage represents the open-weight frontier, but its quantization (NVFP4) is gaining more traction than the base model itself.

**Open-Weight vs. Proprietary:** All models on this list are open-weight, with many developers offering GGUF/GGML conversions (Unsloth, the community). No closed-source API-only models appear; the trend strongly favors community-accessible, locally runnable models. The “uncensored” niche (especially from DavidAU and HauhauCS) shows a persistent demand for unrestricted conversational and creative use.

**Quantization & Efficiency:** Extreme quantization is a key theme: **Ternary-Bonsai** (2-bit), **NVFP4** (from nota-ai), and **Escha-W2** push compression boundaries. Simultaneously, small models like **Inflect-Micro** (TTS) and **antares-1b** (LLM) cater to edge devices. **Unsloth** continues its role as the go-to converter for GGUF compatibility, reinforcing the importance of local inference.

## 4. Worth Exploring

1. **[prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)** – A 2-bit ternary quantization of a 27B model. It’s a case study in extreme compression – evaluate whether quality degradation is acceptable for memory-constrained applications.

2. **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)** – The week’s most liked model. Its compressed-tensor design and strong multimodal performance make it a benchmark for the next generation of vision-language models.

3. **[owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)** – A tiny CPU-friendly TTS model. Worth trying for offline, privacy-preserving voice synthesis without sacrificing naturalness – a sign of where edge AI is headed.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*