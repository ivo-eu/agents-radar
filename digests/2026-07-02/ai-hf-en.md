# Hugging Face Trending Models Digest 2026-07-02

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-02 10:17 UTC

---

# Hugging Face Trending Models Digest – 2026-07-02

## Today's Highlights
This week's top trends are dominated by **Mixture-of-Experts (MoE) architectures** and **massive community quantization efforts**. The **GLM-5.2** family leads in likes (3,202), with NVIDIA and unsloth already shipping optimized versions. **NVIDIA's LocateAnything-3B** (2,557 likes, 1M+ downloads) shows strong demand for specialized vision transformers. **DeepSeek-V4-Pro/Flash-DSpark** and the **Ornith-1.0** series (from 9B to 397B) signal a rush to scale MoE models while keeping inference efficient. Quantized GGUF variants of nearly every major release (Gemma 4, Qwen3.6, GLM-5.2) indicate the community’s hunger for local, fast inference.

## Trending Models

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)
- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** – Author: zai-org | Likes: 3,202 | Downloads: 176,154  
  A large MoE conversational model from Zhipu AI, trending for its strong general reasoning and the rapid ecosystem of fine-tunes and quantizations around it.
- **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)** – Author: deepseek-ai | Likes: 285 | Downloads: 8,184  
  The latest generation of DeepSeek’s flagship LLM; interest is high despite modest downloads, likely due to its public research paper (arXiv:2606.19348).
- **[deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)** – Author: deepseek-ai | Likes: 127 | Downloads: 23,939  
  A faster variant of DeepSeek-V4 optimized for low-latency inference; draws attention as a potential alternative to proprietary models.
- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)** – Author: LiquidAI | Likes: 185 | Downloads: 26,357  
  A tiny 230M parameter language model from Liquid AI, trending for its efficient design and surprisingly capable performance at a very small scale.

### 🎨 Multimodal & Generation (Image, Video, Text-to-Image, Vision-Language)
- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** – Author: krea | Likes: 446 | Downloads: 69,788  
  High-quality text-to-image model based on Krea-2-Raw, optimized for speed; popularity driven by the creative community and Comfy-Org integrations.
- **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)** – Author: fal | Likes: 138 | Downloads: 0  
  A LoRA adapter for LTX-2.3 enabling realistic 3D-style image-to-video generation; marks the growing interest in controllable video synthesis.
- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)** – Author: deepreinforce-ai | Likes: 337 | Downloads: 58,385  
  Base MoE vision-language model (9B active) from the Ornith family, gaining traction for its efficient multimodal reasoning and full-spectrum quantization support.
- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)** – Author: deepreinforce-ai | Likes: 297 | Downloads: 185,633  
  The 35B variant of Ornith, pushing multimodal MoE performance while remaining deployable; strong download numbers reflect community trust.
- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)** – Author: deepreinforce-ai | Likes: 192 | Downloads: 7,358  
  The largest Ornith model (397B parameters, MoE), a frontier-scale multimodal model for research and high-end inference.
- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)** – Author: Qwen | Likes: 501 | Downloads: 39,448  
  A 35B MoE vision-language model specialized for agentic tasks and world-modeling; part of Qwen’s push toward embodied AI.
- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** – Author: empero-ai | Likes: 626 | Downloads: 124,909  
  A fine-tuned Qwen3.5-based vision-language model with 1M+ context window; trending for its combination of extended context and Claude-Mythos-style personality tuning.
- **[InternScience/Agents-A1](https://huggingface.co/InternScience/Agents-A1)** – Author: InternScience | Likes: 162 | Downloads: 1,533  
  An agent-focused MoE vision-language model from InternScience; small but growing interest in autonomous agent use cases.

### 🔧 Specialized Models (Code, OCR, Localization, Security, Evaluation)
- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** – Author: baidu | Likes: 1,615 | Downloads: 758,489  
  A high-performance OCR model supporting unlimited-length text recognition; widely adopted for document digitization and screen capture tasks.
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** – Author: nvidia | Likes: 2,557 | Downloads: 1,006,831  
  A universal object localization model (3B params) that can identify and segment any object in an image; trending due to its zero-shot capability and NVIDIA’s ecosystem support.
- **[meituan-longcat/LongCat-2.0](https://huggingface.co/meituan-longcat/LongCat-2.0)** – Author: meituan-longcat | Likes: 152 | Downloads: 0  
  An evaluation results collection for long-context models; important for the research community despite zero downloads (likely a metric-only release).

### 📦 Fine-tunes & Quantizations (GGUF, NVFP4, LoRA, Community Adaptations)
- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** – Author: yuxinlu1 | Likes: 2,557 | Downloads: 614,069  
  A highly-rated GGUF quantization of a Gemma 4 fine-tune for coding and reasoning; one of the most downloaded coding models this week.
- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** – Author: yuxinlu1 | Likes: 931 | Downloads: 314,374  
  A more advanced Gemma 4 variant fine-tuned for agentic coding and terminal use; popularity shows demand for local code agents.
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** – Author: HauhauCS | Likes: 2,388 | Downloads: 3,078,904  
  An uncensored, aggressively tuned GGUF version of Qwen3.6-35B-MoE with vision; explosive download count indicates strong interest in unaligned multimodal models.
- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** – Author: empero-ai | Likes: 1,193 | Downloads: 1,250,562  
  GGUF quantization of the extended-context Qwythos model; extremely popular for running Claude-style models locally.
- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)** – Author: deepreinforce-ai | Likes: 380 | Downloads: 255,123  
  GGUF variant of the Ornith 9B model; one of many Ornith quantizations seeing heavy adoption.
- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** – Author: deepreinforce-ai | Likes: 627 | Downloads: 284,585  
  The 35B GGUF version, balancing quality and size; very popular for on-premises deployment.
- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)** – Author: nvidia | Likes: 202 | Downloads: 159,698  
  NVIDIA’s 4-bit floating point quantization of GLM-5.2 using Model Optimizer; shows enterprise interest in efficient inference on NVIDIA hardware.
- **[nvidia/Qwen3.6-27B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4)** – Author: nvidia | Likes: 186 | Downloads: 27,249  
  Similarly optimized Qwen3.6 variant for NVIDIA GPUs; part of the growing NVFP4 quantization ecosystem.
- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** – Author: unsloth | Likes: 491 | Downloads: 239,283  
  Unsloth’s efficient GGUF conversion of GLM-5.2; popular for fast local inference on consumer hardware.
- **[unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)** – Author: unsloth | Likes: 131 | Downloads: 227,332  
  GGUF version of the agentic Qwen model; high downloads reflect desire to run agent models locally.
- **[huihui-ai/Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF)** – Author: huihui-ai | Likes: 121 | Downloads: 2,592  
  Abliterated (de-censored) version of GLM-5.2 in GGUF; niche but notable for the “uncensoring” trend.
- **[Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF)** – Author: Jackrong | Likes: 107 | Downloads: 29,589  
  Coding-focused GGUF quantization of a Qwen3.6 MoE model with vision support; fills a niche for local coding + VLMs.
- **[BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)** – Author: BugTraceAI | Likes: 113 | Downloads: 8,037  
  6-bit GGUF quantization of a cybersecurity-focused fine-tune; emerging interest in specialized local security models.
- **[ilkerzgi/fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs)** – Author: ilkerzgi | Likes: 102 | Downloads: 0  
  Collection of style LoRAs for Krea-2; shows the community adaptor ecosystem around popular base models.

## Ecosystem Signal
The current landscape is defined by **MoE dominance and quantization saturation**. GLM-5.2, Qwen3.5/3.6 MoE, and Ornith families all leverage Mixture-of-Experts to deliver strong performance at lower active parameter counts — a trend that enables larger models to run on consumer GPUs. The **quantization ecosystem is more vibrant than ever**: nearly every major release immediately receives GGUF variants (often multiple from different quantizers), and NVIDIA’s NVFP4 format is carving out a hardware-specific lane. The **“uncensored” and “abliterated” sub-trend** (see HauhauCS, huihui-ai) signals a persistent community demand for models without safety filters, despite potential misuse. On the *open-weight vs. proprietary* front, DeepSeek-V4 and GLM-5.2 are fully open, while NVIDIA is contributing optimized versions of other’s open models — a positive sign for the ecosystem. **Gemma 4 fine-tunes** (yuxinlu1) are especially notable for their coding and agentic capabilities, suggesting that Google’s base model is becoming a strong canvas for community tuning. Finally, **small models** (LFM2.5-230M) and **specialized vision models** (LocateAnything-3B, Unlimited-OCR) show that efficiency and task-specific utility continue to drive downloads.

## Worth Exploring
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** – With 1M+ downloads and 2.5K likes in one week, this model is a must-try for anyone working on object detection, segmentation, or visual grounding. Its zero-shot localization capability could replace many task-specific pipelines.
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** – An extreme community experiment with massive adoption (3M+ downloads). Using this model provides insight into user preferences for uncensored multimodal models and the performance of aggressive fine-tuning on Qwen3.6 MoE.
- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** – The top coding model this week. Its combination of Gemma 4 base, composer fine-tuning, and GGUF quantization makes it an excellent candidate for local coding assistants. Worth testing against GPT-4 class models for code generation and reasoning tasks.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*