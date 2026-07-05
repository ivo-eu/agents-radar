# Hugging Face Trending Models Digest 2026-07-05

> Source: [Hugging Face Hub](https://huggingface.co/) | 30 models | Generated: 2026-07-05 09:32 UTC

---

# 🤗 Hugging Face Trending Models Digest — 2026-07-05

## Today’s Highlights

This week’s trending leaderboard is dominated by **Qwen 3.5/3.6 MoE derivatives** and a wave of **GGUF quantizations** for local deployment, with several models surpassing 1 M downloads. Notably, **GLM‑5.2** (3,418 likes) from ZAI Org and **NVIDIA’s LocateAnything‑3B** (2,607 likes) both signal strong interest in efficient MoE architectures and multimodal grounding. The **Ornith 1.0** family from deepreinforce‑ai covers 9B, 35B, and 397B scales, with GGUF variants accounting for hundreds of thousands of downloads. Meanwhile, **DeepSeek V4** (Pro & Flash) and **Gemma 4‑based coder/agentic** fine‑tunes continue the trend of large‑scale reasoning models optimized for local inference.

## Trending Models

### 🧠 Language Models (LLMs, Chat, Instruction-Tuned)

- **[zai‑org/GLM‑5.2](https://huggingface.co/zai‑org/GLM‑5.2)**  
  *Author: zai‑org · Likes: 3,418 · Downloads: 220,379*  
  A 5.2‑scale MoE conversational model that combines strong reasoning with a dense‑sparse mixture architecture; trending for its balance of quality and parameter efficiency.

- **[deepseek‑ai/DeepSeek‑V4‑Pro‑DSpark](https://huggingface.co/deepseek‑ai/DeepSeek‑V4‑Pro‑DSpark)**  
  *Author: deepseek‑ai · Likes: 374 · Downloads: 12,580*  
  The latest Pro variant of DeepSeek V4, featuring a “DSpark” inference acceleration; draws attention for its arxiv paper and competitive small‑model performance.

- **[deepseek‑ai/DeepSeek‑V4‑Flash‑DSpark](https://huggingface.co/deepseek‑ai/DeepSeek‑V4‑Flash‑DSpark)**  
  *Author: deepseek‑ai · Likes: 158 · Downloads: 48,696*  
  A lighter Flash version of DeepSeek V4 for faster generation; popular for on‑device and low‑latency use cases.

- **[deepreinforce‑ai/Ornith‑1.0‑35B](https://huggingface.co/deepreinforce‑ai/Ornith‑1.0‑35B)**  
  *Author: deepreinforce‑ai · Likes: 335 · Downloads: 224,641*  
  A 35B MoE language model built on Qwen 3.5, with vision capabilities; part of a family covering 9B–397B scales.

- **[deepreinforce‑ai/Ornith‑1.0‑397B](https://huggingface.co/deepreinforce‑ai/Ornith‑1.0‑397B)**  
  *Author: deepreinforce‑ai · Likes: 210 · Downloads: 70,585*  
  The largest Ornith variant – a 397B MoE – pushing the frontier of open‑weight reasoning models.

- **[Qwen/Qwen‑AgentWorld‑35B‑A3B](https://huggingface.co/Qwen/Qwen‑AgentWorld‑35B‑A3B)**  
  *Author: Qwen · Likes: 538 · Downloads: 55,113*  
  A Qwen 3.5 MoE fine‑tuned for agentic tasks; reflects the growing “agent foundation model” trend.

- **[mistralai/Leanstral‑1.5‑119B‑A6B](https://huggingface.co/mistralai/Leanstral‑1.5‑119B‑A6B)**  
  *Author: mistralai · Likes: 104 · Downloads: 26*  
  A sparse 119B model (6B active) from Mistral, built on Leanstral‑2603; early but promising for efficient large‑scale inference.

- **[nvidia/Nemotron‑Labs‑TwoTower‑30B‑A3B‑Base‑BF16](https://huggingface.co/nvidia/Nemotron‑Labs‑TwoTower‑30B‑A3B‑Base‑BF16)**  
  *Author: nvidia · Likes: 122 · Downloads: 10,696*  
  A two‑tower MoE design from NVIDIA for bilingual or multi‑task setups; explored for improved specialist‑generalist trade‑offs.

---

### 🎨 Multimodal & Generation (Image, Video, Audio, Text‑to‑X)

- **[nvidia/LocateAnything‑3B](https://huggingface.co/nvidia/LocateAnything‑3B)**  
  *Author: nvidia · Likes: 2,607 · Downloads: 1,247,265*  
  A 3B vision‑language model for fine‑grained object grounding and localization; trending for its impressive zero‑shot detection and high download count.

- **[baidu/Unlimited‑OCR](https://huggingface.co/baidu/Unlimited‑OCR)**  
  *Author: baidu · Likes: 1,719 · Downloads: 1,044,217*  
  An image‑text‑to‑text model capable of unlimited‑style OCR (scene text, documents, handwriting); widely used for multilingual document processing.

- **[Krea/Krea‑2‑Turbo](https://huggingface.co/krea/Krea‑2‑Turbo)**  
  *Author: krea · Likes: 502 · Downloads: 99,049*  
  A distilled text‑to‑image model from Krea’s Raw base; popular among creative users for fast, high‑quality generation.

- **[InternScience/Agents‑A1](https://huggingface.co/InternScience/Agents‑A1)**  
  *Author: InternScience · Likes: 252 · Downloads: 7,010*  
  A Qwen 3.5 MoE model fine‑tuned for vision‑language agent tasks; demonstrates the convergence of MoE and multimodal agents.

- **[Jackrong/Qwopus3.6‑35B‑A3B‑Coder‑MTP‑GGUF](https://huggingface.co/Jackrong/Qwopus3.6‑35B‑A3B‑Coder‑MTP‑GGUF)**  
  *Author: Jackrong · Likes: 134 · Downloads: 84,951*  
  A vision‑language coder GGUF based on Qwen 3.6 MoE; combines code generation with image understanding.

---

### 🔧 Specialized Models (Code, Security, Tabular, PII)

- **[yuxinlu1/gemma‑4‑12B‑coder‑fable5‑composer2.5‑v1‑GGUF](https://huggingface.co/yuxinlu1/gemma‑4‑12B‑coder‑fable5‑composer2.5‑v1‑GGUF)**  
  *Author: yuxinlu1 · Likes: 2,598 · Downloads: 651,758*  
  A GGUF‑quantized Gemma 4 fine‑tune for code generation and reasoning; extremely popular for local coding assistants.

- **[google/tabfm‑1.0.0‑pytorch](https://huggingface.co/google/tabfm‑1.0.0‑pytorch)**  
  *Author: google · Likes: 205 · Downloads: 2,670*  
  A tabular foundation model for zero‑shot classification and regression; marks Google’s entry into foundation models for structured data.

- **[BugTraceAI/BugTraceAI‑CORE‑Ultra‑27B‑Q6](https://huggingface.co/BugTraceAI/BugTraceAI‑CORE‑Ultra‑27B‑Q6)**  
  *Author: BugTraceAI · Likes: 132 · Downloads: 12,196*  
  A 27B Q6‑quantized model specialized in cybersecurity and offensive‑security reasoning; appeals to red‑team practitioners.

- **[nationaldesignstudio/rampart](https://huggingface.co/nationaldesignstudio/rampart)**  
  *Author: nationaldesignstudio · Likes: 125 · Downloads: 2,783*  
  An ONNX‑compatible BERT model for PII (personally identifiable information) detection; used in privacy‑aware pipelines.

---

### 📦 Fine‑tunes & Quantizations (Community Fine‑tunes, GGUF, NVFP4)

- **[empero‑ai/Qwythos‑9B‑Claude‑Mythos‑5‑1M‑GGUF](https://huggingface.co/empero‑ai/Qwythos‑9B‑Claude‑Mythos‑5‑1M‑GGUF)**  
  *Author: empero‑ai · Likes: 1,486 · Downloads: 1,533,844*  
  A GGUF quantized Qwen 3.5 vision‑language model fine‑tuned on synthetic Claude‑style data; extremely high download count for local deployment.

- **[deepreinforce‑ai/Ornith‑1.0‑9B‑GGUF](https://huggingface.co/deepreinforce‑ai/Ornith‑1.0‑9B‑GGUF)**  
  *Author: deepreinforce‑ai · Likes: 427 · Downloads: 352,002*  
  GGUF version of the popular Ornith‑1.0‑9B; trending for accessibility on consumer hardware.

- **[deepreinforce‑ai/Ornith‑1.0‑35B‑GGUF](https://huggingface.co/deepreinforce‑ai/Ornith‑1.0‑35B‑GGUF)**  
  *Author: deepreinforce‑ai · Likes: 716 · Downloads: 394,164*  
  The 35B Ornith quantized; balances performance and resource usage, driving high adoption.

- **[yuxinlu1/gemma‑4‑12B‑agentic‑fable5‑composer2.5‑v2‑3.5x‑tau2‑GGUF](https://huggingface.co/yuxinlu1/gemma‑4‑12B‑agentic‑fable5‑composer2.5‑v2‑3.5x‑tau2‑GGUF)**  
  *Author: yuxinlu1 · Likes: 1,015 · Downloads: 355,871*  
  An agentic‑oriented Gemma 4 GGUF fine‑tune with composer‑style instruction tuning; top for goal‑driven coding tasks.

- **[HauhauCS/Qwen3.6‑35B‑A3B‑Uncensored‑HauhauCS‑Aggressive](https://huggingface.co/HauhauCS/Qwen3.6‑35B‑A3B‑Uncensored‑HauhauCS‑Aggressive)**  
  *Author: HauhauCS · Likes: 2,463 · Downloads: 3,018,257*  
  A heavily uncensored, “aggressive” Qwen 3.6 MoE GGUF; massive downloads reflect community demand for less‑restricted models.

- **[nvidia/Qwen3.6‑27B‑NVFP4](https://huggingface.co/nvidia/Qwen3.6‑27B‑NVFP4)**  
  *Author: nvidia · Likes: 261 · Downloads: 297,130*  
  NVIDIA’s NVFP4 4‑bit quantization of Qwen 3.6; demonstrates hardware‑aware quantization for efficient inference on NVIDIA GPUs.

- **[nvidia/GLM‑5.2‑NVFP4](https://huggingface.co/nvidia/GLM‑5.2‑NVFP4)**  
  *Author: nvidia · Likes: 229 · Downloads: 280,087*  
  An NVFP4 quantization of GLM‑5.2; similar trend of enabling large MoE on consumer GPUs.

- **[huihui‑ai/Huihui‑GLM‑5.2‑abliterated‑GGUF](https://huggingface.co/huihui‑ai/Huihui‑GLM‑5.2‑abliterated‑GGUF)**  
  *Author: huihui‑ai · Likes: 163 · Downloads: 5,609*  
  An “abliterated” (safety‑removed) GGUF of GLM‑5.2; niche but notable for alignment research.

- **[DavidAU/Qwen3.5‑9B‑Claude‑4.6‑HighIQ‑THINKING‑HERETIC‑UNCENSORED](https://huggingface.co/DavidAU/Qwen3.5‑9B‑Claude‑4.6‑HighIQ‑THINKING‑HERETIC‑UNCENSORED)**  
  *Author: DavidAU · Likes: 145 · Downloads: 53,962*  
  An uncensored, reasoning‑focused fine‑tune of Qwen 3.5; part of the “thinking” model trend.

- **[empero‑ai/Qwythos‑9B‑Claude‑Mythos‑5‑1M](https://huggingface.co/empero‑ai/Qwythos‑9B‑Claude‑Mythos‑5‑1M)**  
  *Author: empero‑ai · Likes: 672 · Downloads: 144,933*  
  The base version (non‑quantized) of the Qwythos model; remains popular for fine‑tuning and evaluation.

---

## Ecosystem Signal

**Qwen 3.5/3.6 MoE** is the dominant base family this week, appearing in 10+ models ranging from 9B to 397B. The combination of MoE efficiency, vision capabilities, and strong community fine‑tuning (uncensored, agentic, coder) has made it the “Llama of 2026” for many developers. **DeepSeek V4** and **GLM‑5.2** represent the second tier of MoE contenders, both receiving official NVIDIA quantizations (NVFP4) that lower the barrier for local deployment. **Gemma 4** continues to attract focused fine‑tunes for coding and agentic workflows, with GGUF versions exceeding 650k downloads.

Open‑weight models are clearly winning the popularity contest – proprietary models are absent from trending. Quantization (especially GGUF and NVFP4) is the primary driver of adoption, as users prioritize running large models on consumer hardware. **Uncensored/abliterated variants** remain a persistent niche, often with disproportionately high likes and downloads. The emergence of **tabfm** from Google signals a new frontier: foundation models for non‑natural‑language data (tabular). **NVIDIA’s LocateAnything‑3B** highlights that small, specialized multimodal models can achieve broad appeal when they solve a clear use case (object grounding). The ecosystem is maturing: multiple families, extensive quantization support, and task‑specific fine‑tuning are now the norm.

## Worth Exploring

1. **[nvidia/LocateAnything‑3B](https://huggingface.co/nvidia/LocateAnything‑3B)**  
   *Why explore:* A compact yet powerful grounding model that achieves high zero‑shot detection; ideal for robotics, visual QA, and document understanding – and one of the most liked models this week.

2. **[google/tabfm‑1.0.0‑pytorch](https://huggingface.co/google/tabfm‑1.0.0‑pytorch)**  
   *Why explore:* Google’s first tabular foundation model, offering zero‑shot classification/regression without task‑specific training – a potential game‑changer for data‑science workflows and a model worth studying for its architecture.

3. **[deepreinforce‑ai/Ornith‑1.0‑397B](https://huggingface.co/deepreinforce‑ai/Ornith‑1.0‑397B)**  
   *Why explore:* At 397B total parameters (active not specified), this is among the largest openly‑available MoE models. It pushes the frontier of open‑weight capabilities and can serve as a benchmark for scaling laws and inference optimization.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*