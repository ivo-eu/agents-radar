# Hugging Face Trending Models Digest 2026-06-19

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-19 12:58 UTC

---

# Hugging Face Trending Models Digest — June 19, 2026

## Today's Highlights

This week's leaderboard is dominated by **DeepSeek-V4-Pro**, which tops the chart with nearly 5,000 weekly likes and over 3 million downloads, cementing its status as the most popular open-weight LLM. **NVIDIA** makes a strong showing with **LocateAnything-3B**, a specialized vision model gaining traction for its practical object-localization capabilities. The multimodal frontier is expanding rapidly, with **Google's diffusiongemma-26B-A4B-it** and the unified **Gemma-4-12B-it** (any-to-any pipeline) attracting significant attention. Meanwhile, community quantization efforts, particularly **unsloth** GGUF conversions of nearly every major release, continue to drive accessibility. A notable trend is the rise of **uncensored and "obliterated" fine-tunes**, especially around the Qwen3.6 and Gemma-4 families, reflecting strong demand for less-restrictive model variants.

---

## Trending Models

### 🧠 Language Models (LLMs, chat models, instruction-tuned)

- **DeepSeek-V4-Pro**  
  Author: deepseek-ai | Likes: 4,959 | Downloads: 3,015,772  
  The top trending LLM—a conversational MoE model with strong reasoning and coding capabilities, likely setting new standards for open-weight performance.

- **GLM-5.2**  
  Author: zai-org | Likes: 1,445 | Downloads: 11,871  
  A conversational, MoE-based language model from zai-org, drawing interest for its reported efficiency and advanced dialogue capabilities.

- **FastContext-1.0-4B-SFT**  
  Author: microsoft | Likes: 213 | Downloads: 1,437  
  A small (4B) SFT model from Microsoft, optimized for agent-like tasks and long-context reasoning—worth watching for lightweight deployment.

- **Nex-N2-Pro**  
  Author: nex-agi | Likes: 332 | Downloads: 7,507  
  A Qwen3.5-MoE-based text-generation model, gaining traction for its balanced performance and multimodal support.

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **nvidia/LocateAnything-3B**  
  Author: nvidia | Likes: 2,183 | Downloads: 228,669  
  A vision model specialized for object localization in images, trending due to its practical utility and NVIDIA's backing.

- **MiniMaxAI/MiniMax-M3**  
  Author: MiniMaxAI | Likes: 1,114 | Downloads: 67,836  
  A multimodal MoE model (image-text-to-text) from MiniMax, popular for its versatility and agent-like behavior.

- **google/diffusiongemma-26B-A4B-it**  
  Author: google | Likes: 1,007 | Downloads: 601,208  
  A diffusion-based multimodal model that processes images and text, notable for its innovative architecture and high download volume.

- **google/gemma-4-12B-it**  
  Author: google | Likes: 1,091 | Downloads: 1,590,882  
  Google's flagship "any-to-any" model, capable of handling multiple modalities (image, text, audio) in a single pipeline.

- **prefeitura-rio/Rio-3.5-Open-397B**  
  Author: prefeitura-rio | Likes: 325 | Downloads: 190,639  
  An enormous (397B) MoE multimodal model, likely the largest on the leaderboard, attracting curiosity for its scale.

- **zai-org/SCAIL-2**  
  Author: zai-org | Likes: 230 | Downloads: 0  
  An image-to-video diffusion model focused on character animation and pose-driven generation—niche but innovative.

- **Zyphra/ZONOS2**  
  Author: Zyphra | Likes: 116 | Downloads: 719  
  A text-to-speech model, one of the few non-LLM entries, gaining interest for its Apache-2.0 license.

- **bosonai/higgs-audio-v3-tts-4b**  
  Author: bosonai | Likes: 491 | Downloads: 69,143  
  A 4B TTS model with multimodal Qwen3 backbone, likely offering high-quality speech synthesis.

- **owensong/Inflect-Nano-v1**  
  Author: owensong | Likes: 110 | Downloads: 0  
  An ultra-small text-to-speech model (tagged "ultra-small"), appealing for edge deployment.

### 🔧 Specialized Models (code, math, medical, embeddings, ASR)

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**  
  Author: yuxinlu1 | Likes: 1,767 | Downloads: 268,102  
  A GGUF quantized coding-focused variant of Gemma-4, highly downloaded for local code assistant use.

- **moonshotai/Kimi-K2.7-Code**  
  Author: moonshotai | Likes: 893 | Downloads: 274,865  
  A multimodal code model (image-text-to-text) with compressed-tensor support, trending for its integrated vision-reasoning capabilities.

- **WeiboAI/VibeThinker-3B**  
  Author: WeiboAI | Likes: 426 | Downloads: 12,148  
  A math-specialized 3B model, gaining attention as a small but capable reasoning tool.

- **CohereLabs/North-Mini-Code-1.0**  
  Author: CohereLabs | Likes: 454 | Downloads: 17,693  
  A code-focused MoE model from Cohere, competing in the niche of smaller, efficient code assistants.

- **nvidia/nemotron-3.5-asr-streaming-0.6b**  
  Author: nvidia | Likes: 556 | Downloads: 18,809  
  A streaming ASR model (0.6B) with cache-aware architecture, important for real-time speech recognition.

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, AWQ)

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
  Author: HauhauCS | Likes: 1,985 | Downloads: 3,730,978  
  An uncensored, MoE vision model (Qwen3.6-based) with massive download count, reflecting strong demand for less-restricted models.

- **unsloth/GLM-5.2-GGUF**  
  Author: unsloth | Likes: 150 | Downloads: 8,392  
  Unsloth's GGUF quantization of GLM-5.2, enabling local inference on consumer hardware.

- **unsloth/Kimi-K2.7-Code-GGUF**  
  Author: unsloth | Likes: 137 | Downloads: 33,667  
  GGUF version of Kimi-K2.7-Code, making multimodal coding accessible offline.

- **unsloth/diffusiongemma-26B-A4B-it-GGUF**  
  Author: unsloth | Likes: 315 | Downloads: 202,867  
  Quantized diffusiongemma, highly popular for running Google's diffusion multimodal model on modest GPUs.

- **unsloth/MiniMax-M3-GGUF**  
  Author: unsloth | Likes: 104 | Downloads: 24,354  
  GGUF quantization of MiniMax-M3, part of unsloth's systematic quantization coverage.

- **unsloth/gemma-4-12b-it-GGUF**  
  Author: unsloth | Likes: 657 | Downloads: 1,150,270  
  The most downloaded GGUF of Gemma-4-12B-it, enabling any-to-any capabilities on local machines.

- **zai-org/GLM-5.2-FP8**  
  Author: zai-org | Likes: 96 | Downloads: 93,927  
  FP8 quantized version of GLM-5.2, offering memory-efficiency for high-performance inference.

- **DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF**  
  Author: DavidAU | Likes: 403 | Downloads: 588,753  
  An extremely long-named, uncensored, GGUF fine-tune of Qwen3.6 hybridized with Claude-style reasoning, showing the community's appetite for creative merges.

- **Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF**  
  Author: Jackrong | Likes: 256 | Downloads: 148,525  
  A GGUF quantization of a Qwen3.6 coding variant, popular for vision-language coding tasks.

- **Mia-AiLab/Qwable-3.6-27b**  
  Author: Mia-AiLab | Likes: 96 | Downloads: 16,105  
  A GGUF quantized Qwen3.6 model, likely a fine-tune for general use.

- **OBLITERATUS/Gemma-4-12B-OBLITERATED**  
  Author: OBLITERATUS | Likes: 351 | Downloads: 106,885  
  An "obliterated" (uncensored) fine-tune of Gemma-4-12B-it, following the trend of removing safety filters.

---

## Ecosystem Signal

The most powerful signal this week is the **dominance of DeepSeek-V4-Pro**, which surpasses second place by a factor of 2.5× in likes and 10× in downloads among non-quantized models—a clear indicator that the open-weight community is coalescing around DeepSeek's MoE architecture. **Google's Gemma-4 family** is the second major force, with three entries (diffusiongemma, gemma-4-12B-it, and their GGUF versions) all in the top 30, signaling that Google is successfully competing with open releases. 

**Multimodal models** are now the norm rather than the exception: over half of the top 30 models use image-text-to-text or any-to-any pipelines. NVIDIA's LocateAnything demonstrates that specialized vision models (object localization) can also break into the mainstream. 

**Quantization** continues to be a key driver of downloads. Unsloth's systematic GGUF conversions of nearly every trending model (GLM-5.2, Kimi-K2.7-Code, diffusiongemma, MiniMax-M3, gemma-4) receive thousands of downloads each, proving that the community prioritizes local, affordable inference. **Uncensored fine-tunes** (Qwen3.6 variants, Gemma-4-OBLITERATED) are also surging, indicating a persistent demand for models without content restrictions, even as base model providers tighten safety layers. 

**MoE architecture** is pervasive—models like DeepSeek-V4, GLM-5.2, Rio-3.5, Qwen3.5/3.6, and Cohere's North-Mini all leverage mixture-of-experts for efficiency. This suggests the field is converging on MoE as the default for scaling performance without exploding compute costs.

---

## Worth Exploring

1. **DeepSeek-V4-Pro** → [Link](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)  
   *Why:* With the highest likes and downloads, this model is the current pinnacle of open-weight LLMs. Studying its architecture (likely MoE with advanced reasoning) and its community reception can reveal where the ecosystem is headed. It's a must-benchmark for anyone evaluating open alternatives to proprietary models.

2. **nvidia/LocateAnything-3B** → [Link](https://huggingface.co/nvidia/LocateAnything-3B)  
   *Why:* NVIDIA's specialized vision model is a rare example of a non-LLM breaking into the top 10. It showcases the growing demand for **grounded, task-specific models** (object localization) that can be easily integrated into applications. Its small 3B size and practical utility make it immediately deployable.

3. **unsloth/gemma-4-12b-it-GGUF** → [Link](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)  
   *Why:* This GGUF quantization of Google's unified multimodal model is the most downloaded of its kind. It represents the convergence of three trends: any-to-any multimodal, consumer-hardware accessibility (GGUF), and the popularity of the Gemma-4 family. Trying this model gives a realistic view of local multimodal inference performance.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*