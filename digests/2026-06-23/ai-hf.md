# Hugging Face 热门模型日报 2026-06-23

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-23 10:50 UTC

---

# 🧠 Hugging Face 热门模型日报 —— 2026-06-23

## 📰 今日速览

本周 Hugging Face 热点由 **DeepSeek-V4-Pro** 领跑（5,021 赞），成为社区最受关注的通用语言模型。多模态与生成模型强势崛起，**Google DiffusionGemma**、**Gemma-4-12B-it** 以及 **NVIDIA LocateAnything-3B** 均获得高赞，标志着“统一感知+生成”架构走向主流。代码与数学专用模型持续细分，**Gemma-4-12B-Coder** 量化版本下载量破 45 万，**Kimi-K2.7-Code** 也跻身热门。社区量化活动异常活跃，GGUF 版本包揽近三分之一榜单，GLM-5.2、Qwen3.6、Qwythos 等均有多个量化分支。

---

## 🔥 热门模型

### 🧠 语言模型（LLM、对话、指令微调）

- **[DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
  作者：deepseek-ai | 👍 5,021 | 📥 2,245,489  
  新一代旗舰对话模型，本周点赞冠军，延续 DeepSeek 在推理与生成上的前沿突破。

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  作者：zai-org | 👍 2,097 | 📥 40,127  
  智谱 GLM 系列最新 MoE 架构版本，支持对话与文本生成，社区关注度极高。

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
  作者：microsoft | 👍 301 | 📥 4,391  
  微软推出的快速上下文扩展模型，专为长序列任务优化，标签含 Explorer SubAgent。

- **[poolside/Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1)**  
  作者：poolside | 👍 93 | 📥 2,787  
  新晋开源语言模型，支持 vLLM 和 SGLang，旨在探索高效推理。

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**  
  作者：MiniMaxAI | 👍 1,215 | 📥 131,057  
  多模态视觉-语言模型，支持图像理解与文本输出，MiniMax 最新旗舰。

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
  作者：nvidia | 👍 2,303 | 📥 274,025  
  NVIDIA 发布的通用目标定位模型，可基于自然语言在图像中定位任意物体。

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**  
  作者：google | 👍 1,053 | 📥 948,996  
  结合扩散模型与 Gemma 架构的生成式多模态模型，擅长图像与文本混合任务。

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**  
  作者：google | 👍 1,147 | 📥 1,991,703  
  任何到任何（any-to-any）的统一多模态模型，Gemma 4 系列最受瞩目的版本。

- **[datalab-to/lift](https://huggingface.co/datalab-to/lift)**  
  作者：datalab-to | 👍 128 | 📥 3,216  
  基于 Qwen3.5 的文档理解模型，支持 PDF 与图像上下文推理。

- **[owensong/Inflect-Nano-v1](https://huggingface.co/owensong/Inflect-Nano-v1)**  
  作者：owensong | 👍 171 | 📥 0  
  超轻量级文本转语音（TTS）模型，适合边缘设备部署。

- **[Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit)**  
  作者：Boogu | 👍 105 | 📥 592  
  基于 Diffusers 的图像编辑模型，支持中英文提示。

- **[ostris/ideogram_4_turbotime_lora](https://huggingface.co/ostris/ideogram_4_turbotime_lora)**  
  作者：ostris | 👍 106 | 📥 3,672  
  Ideogram 4 的 LoRA 微调版，加速文生图推理。

- **[Comfy-Org/Boogu-Image](https://huggingface.co/Comfy-Org/Boogu-Image)**  
  作者：Comfy-Org | 👍 80 | 📥 0  
  ComfyUI 工作流形式的 Boogu 图像模型，便于节点化部署。

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  作者：HauhauCS | 👍 2,129 | 📥 3,955,016  
  社区去审查版 Qwen3.6 MoE 多模态模型，GGUF 量化，下载量极高。

- **[lordx64/Qwable-v1](https://huggingface.co/lordx64/Qwable-v1)**  
  作者：lordx64 | 👍 168 | 📥 4,547  
  基于 Qwen3.5 MoE 的社区微调多模态模型，支持图像理解。

---

### 🔧 专用模型（代码、数学、医疗、嵌入、OCR、语音）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  作者：baidu | 👍 314 | 📥 8,396  
  百度推出的通用 OCR 模型，支持不限场景的文本识别。

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
  作者：nvidia | 👍 644 | 📥 41,050  
  流式自动语音识别模型，支持缓存感知推理，适合实时场景。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
  作者：WeiboAI | 👍 632 | 📥 41,170  
  专注于数学推理的 3B 模型，基于 Qwen2 微调，展示小模型在数学任务上的潜力。

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
  作者：moonshotai | 👍 966 | 📥 447,920  
  Kimi 系列代码专用版，支持多模态代码推理，压缩后体积更小。

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)**  
  作者：CohereLabs | 👍 484 | 📥 21,634  
  Cohere 推出的轻量代码生成模型，MoE 结构，注重推理效率。

- **[LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)**  
  作者：LiquidAI | 👍 105 | 📥 10,117  
  LFM2.5 系列的句子嵌入模型，适合语义检索和聚类。

- **[LiquidAI/LFM2.5-ColBERT-350M](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M)**  
  作者：LiquidAI | 👍 83 | 📥 2,534  
  基于 ColBERT 架构的稠密检索模型，兼顾效率与精度。

---

### 📦 微调与量化（GGUF、FP8、社区微调）

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
  作者：unsloth | 👍 267 | 📥 55,820  
  GLM-5.2 的官方 GGUF 量化版，借助 unsloth 工具链优化推理速度。

- **[zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8)**  
  作者：zai-org | 👍 139 | 📥 395,290  
  原厂自行量化的 FP8 版本，兼顾精度与显存占用。

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
  作者：yuxinlu1 | 👍 2,199 | 📥 456,117  
  Gemma-4 代码模型的 GGUF 量化，下载量极高，社区首选代码专用量化版。

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
  作者：yuxinlu1 | 👍 407 | 📥 96,459  
  Gemma-4 Agent 功能的精细量化版，专攻终端与代理场景。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
  作者：empero-ai | 👍 153 | 📥 27,218  
  基于 Qwen3.5 的推理增强模型，GGUF 量化后更易本地部署。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
  作者：empero-ai | 👍 149 | 📥 1,856  
  同一模型的原始 safetensors 版，供需要完整精度的用户选用。

- **[bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF)**  
  作者：bytkim | 👍 110 | 📥 65,765  
  使用 pi-tune 技术微调 Qwen3.6 并量化为 GGUF，主打 MTP（multi-turn prompting）。

- **[Mia-AiLab/Qwable-3.6-27b](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b)**  
  作者：Mia-AiLab | 👍 127 | 📥 24,935  
  社区基于 Qwen3.6 微调的 27B 模型，同时提供 GGUF 与 safetensors 版本。

---

## 🌐 生态信号

**模型家族竞争白热化**：DeepSeek-V4 以绝对优势领跑，但 **Qwen3.6** 与 **Gemma-4** 凭借丰富的社区衍生版本（代码、Agent、去审查）形成强大生态矩阵。**GLM-5.2** 的 MoE-DSA 架构也开始获得量化团队支持。**开源权重大幅碾压闭源**：本周 Top10 无一闭源模型，权重公开与合理许可（如 Apache-2.0）依然是社区爆款的基础。**量化活动创纪录**：GGUF 版本占据 1/3 榜单，且下载量往往超越原版——用户更愿直接使用“即开即用”的量化包。社区微调（Qwable、Qwythos）聚焦于**去除审查**和**推理增强**，表明定制化需求旺盛。

---

## 🔭 值得探索

1. **[DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** —— 本周最强通用语言模型，适合作为对话与推理基准，值得深入评估其长上下文与指令遵循能力。

2. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** —— 开创性的“万物定位”范式，将自然语言与视觉 grounding 结合，对机器人、质检等场景有巨大潜力。

3. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** —— 尽管是去审查版，但它是 Qwen3.6 MoE 多模态能力的最佳量化演示，下载量接近 400 万，可体验大规模稀疏激活带来的效率红利。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*