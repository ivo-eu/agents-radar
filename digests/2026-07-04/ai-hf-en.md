# Hugging Face Trending Models Digest 2026-07-04

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-04 09:06 UTC

---

# Hugging Face Trending Models Digest – 2026-07-04

## Today’s Highlights
This week’s trending charts are dominated by **MoE (Mixture-of-Experts) architectures** and **quantized GGUF variants**, reflecting the community’s push for efficient, deployable large models. The **Ornith family** (9B, 35B, 397B) from deepreinforce-ai and **Qwen 3.5/3.6 derivatives** (AgentWorld, uncensored variants) underscore the rapid adoption of sparsely activated transformers. NVIDIA contributes both hardware-optimized formats (NVFP4) and novel vision-language tools like **LocateAnything-3B**, while Google introduces a tabular foundation model. The “abliterated” and uncensored fine-tunes of GLM-5.2 and Qwen 3.6 signal ongoing demand for unrestricted conversational models.

---

## Trending Models

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** – Author: zai-org | Likes: 3,359 | Downloads: 208,920  
  A large MoE conversational model with strong multilingual performance, trending due to its high quality and the active ecosystem of fine-tuned variants.

- **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)** – Author: deepseek-ai | Likes: 353 | Downloads: 10,306  
  The latest flagship from DeepSeek, featuring advanced reasoning and sparse activation; accompanied by a technical report (arxiv:2606.19348).

- **[deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)** – Author: deepseek-ai | Likes: 146 | Downloads: 40,271  
  A faster, inference-optimized sibling of DeepSeek-V4-Pro, gaining traction for production deployments.

- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)** – Author: deepreinforce-ai | Likes: 369 | Downloads: 69,837  
  A compact MoE model (9B activated) that balances performance and efficiency, serving as the base for many GGUF quantizations.

- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)** – Author: deepreinforce-ai | Likes: 325 | Downloads: 218,657  
  A mid-size MoE model (35B total, 3B active) optimized for multimodal tasks and text generation.

- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)** – Author: deepreinforce-ai | Likes: 203 | Downloads: 33,268  
  The monstrous 397B MoE model pushing the limits of sparse activation, appealing to research labs with heavy compute.

- **[InternScience/Agents-A1](https://huggingface.co/InternScience/Agents-A1)** – Author: InternScience | Likes: 216 | Downloads: 5,456  
  A Qwen-3.5-MoE based agentic model designed for tool use and autonomous workflows.

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** – Author: empero-ai | Likes: 663 | Downloads: 135,665  
  A fine-tuned 9B model blending Qwen 3.5 with Claude-like “mythos” style, popular for creative writing and roleplay.

- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)** – Author: Qwen | Likes: 529 | Downloads: 50,188  
  An MoE agentic model (35B total, 3B active) designed for interactive environments and planning tasks.

- **[nvidia/Nemotron-Labs-TwoTower-30B-A3B-Base-BF16](https://huggingface.co/nvidia/Nemotron-Labs-TwoTower-30B-A3B-Base-BF16)** – Author: nvidia | Likes: 116 | Downloads: 10,479  
  A novel two-tower MoE architecture from NVIDIA, optimized for retrieval-augmented generation and efficient inference.

---

### 🎨 Multimodal & Generation (Image, Video, Text-to-X)

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** – Author: baidu | Likes: 1,698 | Downloads: 988,379  
  A powerful end-to-end OCR model capable of handling unlimited text in images, widely used for document digitization.

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** – Author: nvidia | Likes: 2,592 | Downloads: 1,194,542  
  A 3B vision-language model that can locate and ground arbitrary objects in images, trending for its zero-shot detection capabilities.

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** – Author: krea | Likes: 488 | Downloads: 89,384  
  A faster variant of the Krea-2 text-to-image model, offering high-quality generation with improved inference speed.

- **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)** – Author: fal | Likes: 152 | Downloads: 0  
  A LoRA adapter for LTX-2.3 video models that adds realistic 3D style, gaining attention for video-to-3D workflows.

- **[ilkerzgi/fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs)** – Author: ilkerzgi | Likes: 118 | Downloads: 0  
  A collection of style LoRAs for Krea-2, enabling community-driven aesthetic customization.

- **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)** – Author: Comfy-Org | Likes: 244 | Downloads: 10  
  A ComfyUI workflow package for the Krea-2 model, simplifying image generation integration.

---

### 🔧 Specialized Models (Code, Security, Tabular, PII)

- **[google/tabfm-1.0.0-pytorch](https://huggingface.co/google/tabfm-1.0.0-pytorch)** – Author: google | Likes: 156 | Downloads: 1,177  
  A tabular foundation model supporting zero-shot classification and regression on arbitrary tabular data, representing a breakthrough in structured data AI.

- **[BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)** – Author: BugTraceAI | Likes: 127 | Downloads: 12,001  
  A quantized 27B model fine-tuned for cybersecurity and offensive security tasks, including vulnerability discovery.

- **[nationaldesignstudio/rampart](https://huggingface.co/nationaldesignstudio/rampart)** – Author: nationaldesignstudio | Likes: 118 | Downloads: 1,881  
  A BERT-based token classifier for detecting personally identifiable information (PII), optimized for privacy-preserving pipelines.

---

### 📦 Fine-tunes & Quantizations (GGUF, NVFP4, Abliterated)

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** – Author: empero-ai | Likes: 1,399 | Downloads: 1,464,047  
  The GGUF quantized version of the Qwythos creative model, extremely popular for local inference on llama.cpp.

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** – Author: deepreinforce-ai | Likes: 692 | Downloads: 359,659  
  GGUF quantization of the 35B Ornith model, enabling MoE efficiency on consumer hardware.

- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)** – Author: deepreinforce-ai | Likes: 417 | Downloads: 320,660  
  The compact GGUF variant of Ornith-9B, widely downloaded for lightweight local chat.

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** – Author: yuxinlu1 | Likes: 997 | Downloads: 342,752  
  A heavily fine-tuned Gemma 4 12B for agentic coding and terminal use, quantized via GGUF for fast execution.

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** – Author: yuxinlu1 | Likes: 2,589 | Downloads: 641,260  
  Another Gemma-4 coder variant with strong reasoning capabilities, one of the top trending GGUF models.

- **[huihui-ai/Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF)** – Author: huihui-ai | Likes: 149 | Downloads: 4,701  
  An “abliterated” (uncensored) version of GLM-5.2 in GGUF format, appealing to the unrestricted conversation community.

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** – Author: HauhauCS | Likes: 2,442 | Downloads: 2,993,053  
  A massively popular uncensored Qwen 3.6 MoE model in GGUF, driving the highest download count on this list.

- **[Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF)** – Author: Jackrong | Likes: 127 | Downloads: 59,971  
  A specialized coding MoE model based on Qwen 3.6, quantized with MTP (Multi-Token Prediction) support.

- **[nvidia/Qwen3.6-27B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4)** – Author: nvidia | Likes: 234 | Downloads: 184,521  
  A 4-bit floating point quantization of Qwen 3.6 using NVIDIA’s NVFP4 format, optimized for Blackwell GPUs.

- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)** – Author: nvidia | Likes: 216 | Downloads: 236,501  
  The NVFP4 quantized version of GLM-5.2, offering high throughput on NVIDIA hardware.

- **[unsloth/Qwen3.6-27B-MTP-GGUF](https://huggingface.co/unsloth/Qwen3.6-27B-MTP-GGUF)** – Author: unsloth | Likes: 940 | Downloads: 2,752,390  
  A GGUF quantized variant of Qwen 3.6 27B with multi-token prediction, from the popular Unsloth optimization library.

---

## Ecosystem Signal

The current trending landscape reveals a clear **shift toward Mixture-of-Experts (MoE) models** as the default architecture for both proprietary and open-weight releases. Qwen 3.5/3.6 derivatives dominate the list, alongside the Ornith family, GLM-5.2, and DeepSeek-V4, all leveraging sparse activation to scale up total parameters while keeping inference costs manageable. The **GGUF quantization ecosystem** remains the primary vehicle for community adoption, with nearly half of the trending models being quantized variants—highlighting the demand for efficient local deployment on consumer hardware (e.g., llama.cpp). NVIDIA is actively shaping the quantization landscape with its proprietary NVFP4 format, aiming to tie model optimization to its GPU ecosystem. 

The **“uncensored” and “abliterated” fine-tune trend** continues to grow, as seen with the HauhauCS Qwen 3.6 variant (2.9M downloads) and the huihui-ai GLM-5.2 variant, catering to users seeking fewer content restrictions. On the multimodal front, vision-language models like LocateAnything-3B and Unlimited-OCR are gaining traction for practical applications in detection and document processing. Google’s **tabfm-1.0.0** marks a rare entry into tabular foundation models, signaling expansion beyond text and vision. Meanwhile, the Krea-2 image generation ecosystem is building momentum with style LoRAs and ComfyUI integration, reflecting a maturing open-source diffusion pipeline.

---

## Worth Exploring

1. **nvidia/LocateAnything-3B** – With 2.5K+ likes and 1.1M downloads, this compact vision-language model excels at zero-shot object grounding. It is an excellent study case for efficient, deployable multimodal AI and demonstrates NVIDIA’s push into open-weight models beyond LLMs.

2. **Qwen/Qwen-AgentWorld-35B-A3B** – This MoE model is purpose-built for agentic tasks (tool use, planning) and represents a state-of-the-art balance between capacity and efficiency. Its design—35B total, 3B active—makes it a strong candidate for research on autonomous agents and interactive AI.

3. **google/tabfm-1.0.0-pytorch** – As one of the first tabular foundation models with zero-shot capabilities, this model could reshape how structured data tasks are approached. It is worth studying for anyone working in finance, healthcare, or any domain reliant on tabular prediction, even though its adoption is still early.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*