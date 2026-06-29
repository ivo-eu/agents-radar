# Hugging Face Trending Models Digest 2026-06-29

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-29 14:39 UTC

---

# Hugging Face Trending Models Digest — 2026-06-29

## Today's Highlights

This week’s trending models show a strong shift toward **mixture-of-experts (MoE) architectures** and **multimodal vision-language models**, with the GLM-5.2 family (zai-org, unsloth, NVIDIA) receiving massive community traction. **DeepSeek-V4** variants (Pro-DSpark, Flash-DSpark, Fable) signal a new generation of frontier-scale open-weight models, while **NVIDIA’s LocateAnything-3B** dominates vision-language localization tasks. Quantization remains a dominant theme: **GGUF and NVFP4 variants** account for over a third of the top 30, with uncensored fine-tunes from HauhauCS and deepreinforce’s Ornith series seeing explosive download rates. Finally, **AgentWorld-35B-A3B** from Qwen and **FastContext-4B** from Microsoft hint at the growing importance of agentic and long-context capabilities.

---

## Trending Models

### 🧠 Language Models (LLMs, chat, instruction-tuned)

- **[GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — *zai-org* | likes: 2,891 | downloads: 133,350  
  A 5.2-scale MoE conversational model that has become the week’s most-liked language model, driving broad community fine-tuning and quantization.

- **[DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)** — *deepseek-ai* | likes: 204 | downloads: 5,460  
  The latest frontier open-weight release from DeepSeek, featuring dynamic sparse activation; early adoption despite modest downloads.

- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)** — *LiquidAI* | likes: 147 | downloads: 15,463  
  A tiny (230M) liquid foundation model designed for edge deployment, gaining interest for its efficiency in resource-constrained settings.

- **[DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)** — *deepseek-ai* | likes: 89 | downloads: 2,239  
  A faster, lighter variant of DeepSeek-V4 optimized for high-throughput inference, accompanying the Pro version.

---

### 🎨 Multimodal & Generation (image, video, audio, text-to-X)

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** — *baidu* | likes: 1,328 | downloads: 362,945  
  A versatile OCR model that handles unlimited-length text in images; trending due to its robust real-world document processing capabilities.

- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)** — *Qwen* | likes: 423 | downloads: 26,223  
  A vision-language MoE model trained on agentic world-model tasks; rising interest as a backbone for autonomous agents.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** — *empero-ai* | likes: 545 | downloads: 79,540  
  A Qwen3.5-based vision-language model fused with Claude-style reasoning data, popular for creative and role-playing multimodal use.

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** — *krea* | likes: 384 | downloads: 38,454  
  A fine-tuned text-to-image model derived from Krea-2-Raw, delivering faster generation while maintaining image quality.

- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)** — *deepreinforce-ai* | likes: 259 | downloads: 19,170  
  The base 9B variant of the Ornith vision-language MoE series, attracting attention for its MIT license and endpoint compatibility.

- **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)** — *krea* | likes: 241 | downloads: 27,464  
  The raw (unfinetuned) text-to-image model underlying the Krea-2 ecosystem; a reference point for community fine-tuning.

- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)** — *deepreinforce-ai* | likes: 225 | downloads: 38,857  
  The 35B MoE variant of Ornith, balancing scale and inference cost for vision-language tasks.

- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)** — *deepreinforce-ai* | likes: 159 | downloads: 1,622  
  The largest Ornith variant (397B parameters), pushing state-of-the-art in multimodal understanding.

- **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)** — *fal* | likes: 109 | downloads: 0  
  A LoRA for LTX video models enabling photorealistic 3D-aware video generation; early release with high potential.

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — *nvidia* | likes: 739 | downloads: 76,154  
  A streaming automatic speech recognition model (0.6B parameters) optimized for low-latency on-device use; widely adopted for real-time audio.

---

### 🔧 Specialized Models (code, math, embeddings, OCR, tools, cybersecurity)

- **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)** — *Comfy-Org* | likes: 185 | downloads: 10  
  A ComfyUI integration node for Krea-2 models, enabling seamless text-to-image workflows in the popular UI framework; trending despite low downloads due to its utility.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — *nvidia* | likes: 2,460 | downloads: 728,320  
  A 3B vision model for object localization and feature extraction, racking up downloads as a go‑to solution for spatial AI tasks.

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — *WeiboAI* | likes: 746 | downloads: 63,449  
  A Qwen2-based math reasoning model that combines chain‑of‑thought with latent “vibe” embeddings; popular for educational and STEM applications.

- **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)** — *Chunjiang-Intelligence* | likes: 126 | downloads: 1,463  
  A fine‑tune of DeepSeek‑v4 specialized for cybersecurity threat detection and response; an example of domain‑specific trust‑safety fine‑tuning.

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — *microsoft* | likes: 369 | downloads: 7,027  
  A 4B model optimized for long‑context agentic tasks (Explorer SubAgent), reflecting the industry push toward context‑efficient reasoning.

---

### 📦 Fine-tunes & Quantizations (community fine-tunes, GGUF, NVFP4)

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** — *empero-ai* | likes: 879 | downloads: 907,682  
  GGUF‑quantized version of the Qwythos vision-language model; extremely high downloads thanks to CPU/edge compatibility.

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** — *yuxinlu1* | likes: 828 | downloads: 241,409  
  A GGUF‑quantized Gemma‑4 agentic model that fuses multiple fine‑tunes (Fable5, Composer2.5); popular for terminal‑based coding agents.

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** — *deepreinforce-ai* | likes: 445 | downloads: 123,598  
  GGUF quant of the 35B Ornith model, enabling vision‑language inference on consumer hardware.

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — *yuxinlu1* | likes: 2,493 | downloads: 561,577  
  A GGUF‑quantized Gemma‑4 fine‑tune specialized for code generation and reasoning; the second most‑liked model this week.

- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)** — *deepreinforce-ai* | likes: 292 | downloads: 68,667  
  Smallest Ornith GGUF variant, ideal for on‑device multimodal apps.

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — *unsloth* | likes: 456 | downloads: 164,180  
  Unsloth’s GGUF conversion of GLM‑5.2, making the MoE model accessible to Llama.cpp users.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — *HauhauCS* | likes: 2,326 | downloads: 3,089,944  
  An uncensored, aggressively fine‑tuned GGUF of Qwen3.6‑35B‑A3B; the highest download count on the list, driven by demand for unrestricted vision‑language chat.

- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)** — *nvidia* | likes: 161 | downloads: 81,944  
  NVIDIA’s ModelOpt‑quantized GLM‑5.2 in NVFP4 format, offering high efficiency on Hopper‑architecture GPUs.

- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)** — *nvidia* | likes: 375 | downloads: 5,392,518  
  NVFP4 quantization of Qwen3.6‑35B‑A3B; the second most‑downloaded model overall, reflecting NVIDIA’s push into efficient MoE inference.

- **[unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)** — *unsloth* | likes: 102 | downloads: 116,693  
  GGUF version of the AgentWorld model, lowering the barrier for running world‑model agents locally.

- **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)** — *HauhauCS* | likes: 104 | downloads: 46,053  
  A quantization‑aware fine‑tune (QAT) of Gemma‑4 with uncensored outputs, balancing safety removal and performance.

---

## Ecosystem Signal

**MoE architectures dominate the top 30**, with GLM‑5.2, Qwen‑AgentWorld, Ornith, and DeepSeek‑V4 all leveraging mixture‑of‑experts designs. This trend signals industry‑wide convergence on sparse activation as the preferred path to scaling without proportional compute cost.

**Open‑weight momentum is accelerating.** DeepSeek‑V4 (Pro/Flash), GLM‑5.2, and the Ornith series are all open‑weight, while NVIDIA and Microsoft release models under permissive licenses. The contrast with proprietary API-only models is sharp: community adoption (downloads, fine‑tunes, quantizations) is heavily skewed toward open models.

**Quantization and fine‑tuning have become the primary consumption channel.** GGUF, NVFP4, and QAT variants account for 11 of the 30 trending models, and their combined downloads exceed 10 million. The rise of “uncensored” fine‑tunes from HauhauCS reflects a growing demand for models without safety guardrails, especially in role‑playing and creative generation.

**Agentic and world‑model capabilities are gaining ground.** Models like Qwen‑AgentWorld, FastContext, and the agentic Gemma‑4 fine‑tunes indicate that the next frontier is building models that can reason, plan, and act in simulated environments.

**Smaller, efficient models (LFM2.5‑230M, LocateAnything‑3B) prove that specialization and low‑footprint design can still attract significant attention**, especially when paired with clear use cases (edge deployment, vision localization).

---

## Worth Exploring

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — With 728k downloads and high likes, this model is a standout for anyone building vision‑based agents, robotics, or spatial AI. Its combination of small size (3B) and robust localization performance makes it a practical research and production asset.

2. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — As long‑context and agentic reasoning become critical, this 4B model’s “Explorer SubAgent” architecture offers a lightweight yet powerful approach. Worth studying for its trade‑offs between context length, compute, and accuracy.

3. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — The most‑downloaded model on the list (over 3 million) represents the extreme end of community customization. Exploring it (cautiously) can reveal insights into user expectations for model behavior, the limits of fine‑tuning, and the ethics of uncensored release.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*