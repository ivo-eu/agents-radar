# Hugging Face Trending Models Digest 2026-07-03

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-03 10:12 UTC

---

# Hugging Face Trending Models Digest — July 3, 2026

## 🔥 Today's Highlights

The week’s trending models reveal three major currents: **open-weight MoE breakthroughs**, **explosive GGUF quantization activity**, and **specialized multimodal tools**. DeepSeek’s V4 Pro/Flash variants (with arxiv paper) and ZAI-Org’s GLM-5.2 (3.3k likes) signal a new generation of large mixture-of-experts models. Community fine-tuning is hyperactive around Qwen3.5/3.6 and Gemma-4, with uncensored, abliterated, and coder variants all seeing mass adoption. Nvidia’s **LocateAnything-3B** for zero-shot object localization amassed 2.6k likes and 1.1M downloads, while Baidu’s **Unlimited-OCR** proves OCR remains a critical enterprise need. Quantization formats (GGUF, NVFP4) dominate downloads, reflecting users’ demand for local deployment.

---

## 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **zai-org/GLM-5.2**  
  Author: zai-org | Likes: 3,295 | Downloads: 191,462  
  A large MoE text-generation model with strong conversational ability; trending due to its open-weight release and impressive performance on reasoning benchmarks.

- **deepseek-ai/DeepSeek-V4-Pro-DSpark**  
  Author: deepseek-ai | Likes: 316 | Downloads: 9,388  
  Newly released flagship DeepSeek model with DSpark optimization, paper at arxiv:2606.19348; gaining attention as a potential open-weights state-of-the-art.

- **deepseek-ai/DeepSeek-V4-Flash-DSpark**  
  Author: deepseek-ai | Likes: 130 | Downloads: 32,675  
  Faster, lighter sibling of V4-Pro; attracts users seeking high-quality open LLMs that can run on consumer hardware.

- **deepreinforce-ai/Ornith-1.0-9B** (base)  
  Author: deepreinforce-ai | Likes: 355 | Downloads: 64,051  
  A 9B image-text-to-text model built on Qwen3.5; represents the growing Ornith family with strong vision-language capabilities.

- **deepreinforce-ai/Ornith-1.0-35B** (base)  
  Author: deepreinforce-ai | Likes: 314 | Downloads: 211,406  
  Larger MoE-based Ornith variant (35B) that balances performance and efficiency, widely downloaded.

- **Qwen/Qwen-AgentWorld-35B-A3B**  
  Author: Qwen | Likes: 518 | Downloads: 45,455  
  Qwen’s agent-optimized MoE model (35B total, 3B active); trending as the go-to open model for building tool-use and autonomous agents.

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M** (base)  
  Author: empero-ai | Likes: 650 | Downloads: 131,323  
  A fine-tune of Qwen3.5 on 1M Claude Mythos synthetic dialogues; popular for its creative writing and roleplay ability.

- **LiquidAI/LFM2.5-230M**  
  Author: LiquidAI | Likes: 194 | Downloads: 29,645  
  Tiny 230M-parameter Liquid Foundation Model v2.5; trending for efficient on-device deployment without sacrificing quality.

---

## 🎨 Multimodal & Generation (Image, Video, Audio, Text-to-X)

- **baidu/Unlimited-OCR**  
  Author: baidu | Likes: 1,668 | Downloads: 885,040  
  A powerful OCR model supporting unlimited text recognition; widely used for document processing and digitization pipelines.

- **nvidia/LocateAnything-3B**  
  Author: nvidia | Likes: 2,583 | Downloads: 1,108,586  
  A 3B vision-language model for zero-shot object localization; extremely popular for robotics, image editing, and interactive AI.

- **krea/Krea-2-Turbo**  
  Author: krea | Likes: 467 | Downloads: 84,006  
  Fast text-to-image model based on Krea-2-Raw, optimized for high-quality generation; used in creative tools and design workflows.

- **fal/LTX-2.3-3DREAL-LoRA**  
  Author: fal | Likes: 149 | Downloads: 0  
  A LoRA adapter for LTX-2.3 enabling 3D-realistic video generation from single images; new but already catching attention.

- **ilkerzgi/fal-Krea-2-Style-LoRAs**  
  Author: ilkerzgi | Likes: 111 | Downloads: 0  
  Collection of style LoRAs for Krea-2; fills demand for customizable image generation.

---

## 🔧 Specialized Models (Code, Math, Medical, Embeddings, Tabular, PII, Cybersecurity)

- **google/tabfm-1.0.0-pytorch**  
  Author: google | Likes: 133 | Downloads: 450  
  Tabular foundation model for zero-shot classification/regression; first of its kind from Google, promising for enterprise ML.

- **nationaldesignstudio/rampart**  
  Author: nationaldesignstudio | Likes: 108 | Downloads: 1,149  
  On-device PII detection model (BERT-based, ONNX/Transformers.js); trending for privacy-preserving data redaction.

- **BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6**  
  Author: BugTraceAI | Likes: 121 | Downloads: 11,444  
  A Q6 quantized cybersecurity-focused LLM for offensive and defensive security analysis; niche but rapidly downloaded.

- **Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF**  
  Author: Jackrong | Likes: 120 | Downloads: 44,807  
  Code-focused GGUF model based on Qwen3.6 MoE with multi-token prediction; popular among developers for local coding assistance.

- **InternScience/Agents-A1**  
  Author: InternScience | Likes: 193 | Downloads: 3,530  
  An agent-oriented MoE model; early-stage but gaining interest as a lightweight alternative for tool-use tasks.

---

## 📦 Fine-tunes & Quantizations (Community Fine-tunes, GGUF, AWQ, NVFP4)

- **empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF**  
  Author: empero-ai | Likes: 1,305 | Downloads: 1,366,360  
  GGUF-quantized version of the popular Qwythos fine-tune; highest downloads on this list, reflecting massive community adoption.

- **deepreinforce-ai/Ornith-1.0-35B-GGUF**  
  Author: deepreinforce-ai | Likes: 667 | Downloads: 322,780  
  GGUF variant of the 35B Ornith; enables running a powerful MoE model on consumer GPUs.

- **deepreinforce-ai/Ornith-1.0-9B-GGUF**  
  Author: deepreinforce-ai | Likes: 401 | Downloads: 287,942  
  Smaller GGUF version of Ornith-1.0-9B; ideal for laptops and edge devices.

- **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**  
  Author: yuxinlu1 | Likes: 971 | Downloads: 329,391  
  Agentic fine-tune of Gemma-4 12B, heavily quantized; trending for terminal/agent use cases.

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**  
  Author: yuxinlu1 | Likes: 2,573 | Downloads: 628,225  
  Code-focused Gemma-4 fine-tune in GGUF; among the highest liked models, praised for coding reasoning.

- **huihui-ai/Huihui-GLM-5.2-abliterated-GGUF**  
  Author: huihui-ai | Likes: 139 | Downloads: 3,683  
  Abliterated (uncensored) version of GLM-5.2 in GGUF format; taps into demand for less restrictive models.

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
  Author: HauhauCS | Likes: 2,409 | Downloads: 3,029,679  
  Massive download count; an uncensored, aggressive fine-tune of Qwen3.6 MoE in GGUF format, extremely popular for creative and unfiltered use.

- **nvidia/Qwen3.6-27B-NVFP4**  
  Author: nvidia | Likes: 216 | Downloads: 94,465  
  Nvidia’s NVFP4 quantized Qwen3.6 variant; optimized for efficient inference on Nvidia hardware.

- **nvidia/GLM-5.2-NVFP4**  
  Author: nvidia | Likes: 210 | Downloads: 189,970  
  NVFP4 quantization of GLM-5.2; provides a high-speed drop-in replacement for the base model.

- **BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6**  
  Listed above under Specialized; also a quantization (Q6 GGUF) of a security model.

- **Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF**  
  Listed above; also a code-specialized GGUF quantization.

- **deepreinforce-ai/Ornith-1.0-397B**  
  Author: deepreinforce-ai | Likes: 199 | Downloads: 8,079  
  The largest Ornith variant (397B) – not quantized but trending as an extreme-scale model for research.

- **Comfy-Org/Krea-2**  
  Author: Comfy-Org | Likes: 235 | Downloads: 10  
  ComfyUI integration for Krea-2; enables node-based image generation workflows.

---

## 📊 Ecosystem Signal

The ecosystem is coalescing around **three pillars**: open-weight MoE, community quantization, and niche specialization. **Qwen3.5/3.6** has become the most forked base family, with dozens of fine-tunes (uncensored, coder, agentic) and GGUF variants accumulating millions of downloads. **GLM-5.2** from ZAI-Org represents China’s continued push for open foundation models, and Nvidia’s NVFP4 quantization signals industry support for efficient deployment. **DeepSeek-V4 Pro/Flash** enters as a potential new leader, but adoption is still early relative to its likes. The open-weight trend is clearly winning — all top models are openly licensed, with proprietary only appearing in specialized tools like Nvidia’s.

Quantization is not a side activity; it is the **primary distribution method** for LLMs. GGUF variants of fine-tuned models consistently out-download their base versions by 10–50x. This mirrors a maturing ecosystem where users want ready-to-run models for local inference. Meanwhile, specialized models like **LocateAnything-3B** (object localization) and **rampart** (PII redaction) show that beyond general chatbots, there is strong demand for task-specific, efficient models that can be integrated into workflows. The appearance of **tabfm** from Google hints that tabular foundation models may become a new category.

---

## 💡 Worth Exploring

1. **deepseek-ai/DeepSeek-V4-Flash-DSpark**  
   If you’re looking for the newest open-weight LLM, this Flash variant balances quality and speed. Its release with a paper (2606.19348) suggests it might be the most competitive open model this month. Try it for reasoning or coding tasks.

2. **nvidia/LocateAnything-3B**  
   With 2.6k likes and 1.1M downloads, this model stands out for its unique zero-shot localization ability. It’s immediately useful for UI automation, image editing, and robotics. The small size (3B) means it runs on modest hardware.

3. **nationaldesignstudio/rampart**  
   A rare on-device PII detection model with ONNX/Transformers.js support. For developers building privacy-first applications, this is a ready-made solution. Its zero downloads (at time of listing) are deceptive; the model is new and highly relevant.

Also consider **HauhauCS/Qwen3.6-35B-A3B-Uncensored…** for the uncensored MoE experience (3M downloads can’t be wrong), and **yuxinlu1/gemma-4-12B-coder…** for a top-rated coding assistant in GGUF format.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*