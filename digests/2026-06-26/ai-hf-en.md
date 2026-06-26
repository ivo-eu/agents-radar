# Hugging Face Trending Models Digest 2026-06-26

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-06-26 10:38 UTC

---

# Hugging Face Trending Models Digest — June 26, 2026

## 🚀 Today's Highlights

This week’s trending models showcase a powerful surge in quantized, fine-tuned, and MoE-based language models, with the GLM-5.2 (zai-org) leading at 2.5K likes and its GGUF variant by unsloth also in the top ranks. Multimodal models continue to dominate — Google’s gemma-4-12B-it (any-to-any) and MiniMax-M3 are heavily downloaded, while specialized vision tools like nvidia/LocateAnything-3B and baidu/Unlimited-OCR see strong community interest for practical applications. The ecosystem is marked by aggressive community quantization (GGUF, NVFP4) and abliterated uncensored variants, indicating a demand for both performant and flexible model access across consumer hardware.

## 🧠 Language Models

**zai-org/GLM-5.2**  
Author: zai-org | Likes: 2,535 | Downloads: 83,589  
A large MoE language model with DSA architecture, trending due to strong reasoning benchmarks and widespread GGUF support.

**WeiboAI/VibeThinker-3B**  
Author: WeiboAI | Likes: 721 | Downloads: 54,638  
A compact 3B Qwen2-based model fine-tuned for mathematical reasoning, popular for efficient math tutoring on limited hardware.

**Qwen/Qwen-AgentWorld-35B-A3B**  
Author: Qwen | Likes: 272 | Downloads: 13,186  
A 35B MoE agent‑oriented model with 3B active parameters, gaining traction as a lightweight backbone for autonomous agent pipelines.

**microsoft/FastContext-1.0-4B-SFT**  
Author: microsoft | Likes: 349 | Downloads: 5,735  
A 4B SFT model optimized for long‑context “Explorer SubAgent” tasks, notable for Microsoft’s open‑weight release in the growing agent space.

**LiquidAI/LFM2.5-230M**  
Author: LiquidAI | Likes: 91 | Downloads: 8,286  
A tiny 230M‑parameter foundation model from Liquid AI, interesting for its liquid neural network lineage and extreme efficiency.

## 🎨 Multimodal & Generation

**baidu/Unlimited-OCR**  
Author: baidu | Likes: 957 | Downloads: 134,146  
A universal OCR model that converts images to text, trending as a production‑ready solution for document digitization.

**empero-ai/Qwythos-9B-Claude-Mythos-5-1M**  
Author: empero-ai | Likes: 408 | Downloads: 20,346  
The base multimodal (image‑text‑to‑text) version of the Qwythos series, blending Qwen3.5 with Claude‑style synthetic data to boost reasoning.

**krea/Krea-2-Turbo**  
Author: krea | Likes: 262 | Downloads: 8,721  
A fast distilled text‑to‑image model based on Krea-2-Raw, popular for real‑time image generation in creative workflows.

**krea/Krea-2-Raw**  
Author: krea | Likes: 198 | Downloads: 10,408  
The foundational diffusion model behind the Krea‑2 ecosystem, offering high‑quality text‑to‑image generation under an open license.

**nvidia/LocateAnything-3B**  
Author: nvidia | Likes: 2,372 | Downloads: 494,756  
A 3B visual grounding model that can locate any object in an image by text description, widely adopted for robotics and image editing pipelines.

**datalab-to/lift**  
Author: datalab-to | Likes: 153 | Downloads: 6,054  
A Qwen3.5‑based model specialized in PDF understanding and extraction, trending for enterprise document AI use cases.

**MiniMaxAI/MiniMax-M3**  
Author: MiniMaxAI | Likes: 1,241 | Downloads: 169,951  
A large multimodal VL model from MiniMax, competing with proprietary alternatives and attracting attention for its image‑text reasoning quality.

**moonshotai/Kimi-K2.7-Code**  
Author: moonshotai | Likes: 994 | Downloads: 540,180  
A compressed, code‑focused vision‑language model from Kimi, trending for its strong performance on visual coding tasks at reduced size.

**owensong/Inflect-Nano-v1**  
Author: owensong | Likes: 203 | Downloads: 0  
An ultra‑small text‑to‑speech model designed for edge devices, gaining interest for its single‑file PyTorch implementation.

**nvidia/nemotron-3.5-asr-streaming-0.6b**  
Author: nvidia | Likes: 699 | Downloads: 56,434  
A cache‑aware streaming ASR model with 0.6B parameters, notable for real‑time speech recognition with low latency.

**google/gemma-4-12B-it**  
Author: google | Likes: 1,185 | Downloads: 2,309,976  
Google’s flagship any‑to‑any multimodal model (text, image, audio), receiving massive downloads for its unified architecture and broad capability.

**Comfy-Org/Krea-2**  
Author: Comfy-Org | Likes: 127 | Downloads: 10  
A ComfyUI workflow bundle for the Krea‑2 model family, easing diffusion deployment for the community.

**Boogu/Boogu-Image-0.1-Edit**  
Author: Boogu | Likes: 127 | Downloads: 903  
An image editing diffusion model supporting both English and Chinese, fostering accessible image manipulation tools.

## 🔧 Specialized Models

**Chunjiang-Intelligence/DeepSeek-v4-Fable**  
Author: Chunjiang-Intelligence | Likes: 102 | Downloads: 1,103  
A cybersecurity‑focused variant of DeepSeek v4, specialized for threat analysis and penetration testing assistance.

## 📦 Fine-tunes & Quantizations

**unsloth/GLM-5.2-GGUF**  
Author: unsloth | Likes: 396 | Downloads: 107,553  
GGUF quantization of GLM-5.2, enabling efficient local inference and highly compatible with llama.cpp.

**yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**  
Author: yuxinlu1 | Likes: 642 | Downloads: 186,663  
A GGUF‑packed agentic fine‑tune of Gemma‑4‑12B, designed for terminal and coding agent scenarios with improved temperature scaling.

**yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**  
Author: yuxinlu1 | Likes: 2,376 | Downloads: 516,333  
The most downloaded GGUF this week — a coding‑focused, reasoned Gemma‑4 variant, extremely popular among developers for local code generation.

**empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF**  
Author: empero-ai | Likes: 520 | Downloads: 486,810  
Quantized version of Qwythos-9B, bringing Claude‑style reasoning to llama.cpp users at 4‑bit.

**HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
Author: HauhauCS | Likes: 2,249 | Downloads: 3,453,492  
An uncensored, aggressive‑style GGUF fine‑tune of Qwen3.6‑35B‑A3B MoE, massively downloaded for unrestricted role‑play and content generation.

**deepreinforce-ai/Ornith-1.0-35B-GGUF**  
Author: deepreinforce-ai | Likes: 140 | Downloads: 3,002  
A 35B GGUF variant of the Ornith model, aimed at high‑quality text generation with MIT license compatibility.

**deepreinforce-ai/Ornith-1.0-9B-GGUF**  
Author: deepreinforce-ai | Likes: 89 | Downloads: 1,779  
The 9B sibling of Ornith, offering a lighter alternative for consumer GPUs with endpoint‑compatible quantization.

**Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF**  
Author: Jackrong | Likes: 90 | Downloads: 35,027  
A GGUF‑quantized vision‑code model (27B) using Multi‑Token Prediction, designed for coding tasks that require visual context.

**HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced**  
Author: HauhauCS | Likes: 87 | Downloads: 23,772  
A balanced, uncensored QAT‑fine‑tuned Gemma‑4‑12B GGUF, expanding the uncensored ecosystem for multimodal applications.

**nvidia/Qwen3.6-35B-A3B-NVFP4**  
Author: nvidia | Likes: 350 | Downloads: 4,812,629  
NVIDIA’s 4‑bit NVFP4 quantization of Qwen3.6 MoE, the most downloaded model this week — optimized for TensorRT LLM and high‑throughput deployment.

**huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated**  
Author: huihui-ai | Likes: 130 | Downloads: 5,445  
An “abliterated” (safety‑filter removed) version of the Gemma‑4‑12B coder fine‑tune, providing unrestricted coding assistance.

---

## 🌐 Ecosystem Signal

The dominant trend this week is **community quantization of large MoE models** — GLM-5.2, Qwen3.6‑35B‑A3B, and Gemma‑4‑12B all have GGUF or NVFP4 variants among the top downloads, reflecting a mature ecosystem where users prioritize local inference over API calls. The rise of **uncensored and abliterated variants** (HauhauCS, huihui-ai) signals a growing demand for unrestricted model behavior, especially in creative and role‑play domains. **Multimodal models are converging** — Google’s Gemma‑4‑12B‑it (any‑to‑any), MiniMax‑M3, and nvidia/LocateAnything all blur the lines between vision, language, and other modalities, indicating that unified multimodal architectures are becoming the norm rather than the exception. **Small, efficient models** (VibeThinker‑3B, LFM2.5‑230M, Inflect‑Nano) also gain traction for edge deployment. Open‑weight releases from major labs (Microsoft, NVIDIA, Google, Qwen) reinforce the shift toward open‑ecosystem dominance, while proprietary players like MoonshotAI (Kimi‑K2.7) and MiniMax release strong alternatives under permissive licenses.

---

## 🔍 Worth Exploring

1. **nvidia/Qwen3.6-35B-A3B-NVFP4**  
   With 4.8 million downloads in a single week, this NVIDIA‑quantized MoE is essential for anyone deploying large models on NVIDIA GPUs. It demonstrates how optimized inference (NVFP4) can make a 35B model accessible on consumer hardware while retaining high quality.

2. **nvidia/LocateAnything-3B**  
   At 3B parameters, it provides fast, precise visual grounding without the overhead of larger VLMs. Its 2.3K likes and near‑half‑million downloads highlight a clear need for lightweight, production‑ready segmentation and localization.

3. **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
   Despite controversy around uncensored models, this GGUF is the third most downloaded overall, reflecting the community’s interest in unfiltered MoE for narrative and conversational applications. Studying its fine‑tuning approach (aggressive style, Qwen3.6 base) can inform future customization workflows.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*