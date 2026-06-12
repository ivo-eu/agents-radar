# Hugging Face 热门模型日报 2026-06-12

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-12 06:22 UTC

---

# Hugging Face 热门模型日报（2026-06-12）

## 今日速览
- **DeepSeek-V4-Pro 登顶**：周点赞4,785，下载超400万，成为本周最受关注的文本生成模型，延续 DeepSeek 家族强劲势头。
- **Gemma‑4 家族全面开花**：Google 发布多模态统一模型（gemma‑4‑12B‑it、gemma‑4‑12B），社区迅速跟进推出量化版（unsloth、GGUF）和微调版（abliterated、OBLITERATED），形成生态裂变。
- **多模态定位与视频生成成新热点**：Nvidia 的 LocateAnything‑3B（通用定位模型）、ByteDance 的 Bernini‑R（图像→视频）以及 zai‑org 的 SCAIL‑2（角色动画）均收获高赞，显示多模态从“生成”向“可控理解与动作”延伸。
- **量化与部署持续火爆**：HauhauCS 的 Qwen3.6‑Uncensored 下载量突破 300 万，unsloth 为 Gemma‑4 全系列提供 GGUF 量化，官方也推出 FP4/NVFP4 版本，低精度推理已成主流诉求。

---

## 热门模型

### 🧠 语言模型（LLM、对话、指令微调）

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
  作者：deepseek‑ai | 👍 4,785 | 📥 4,061,006  
  *DeepSeek 最新旗舰模型，本周绝对热度第一，展现出强大的通用对话与推理能力。*

- **[nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16)**  
  作者：nvidia | 👍 198 | 📥 59,066  
  *Nvidia 的 550B 参数 MoE 超大规模模型，BF16 精度版本，适合企业级部署。*

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)**  
  作者：nex‑agi | 👍 210 | 📥 1,185  
  *基于 Qwen3.5 MoE 架构的优化模型，在推理效率与性能间取得平衡。*

- **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)**  
  作者：nex‑agi | 👍 166 | 📥 1,222  
  *Nex-N2 系列的小型版本，适合资源受限场景下的文本生成。*

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
  作者：nvidia | 👍 1,886 | 📥 131,794  
  *通用图像定位模型，可在任意图像中通过文本指令定位目标物体，本周多模态明星。*

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**  
  作者：google | 👍 943 | 📥 675,936  
  *Google 最新统一多模态模型（any‑to‑any），支持图像、文本混合输入输出。*

- **[google/gemma-4-12B](https://huggingface.co/google/gemma-4-12B)**  
  作者：google | 👍 519 | 📥 140,221  
  *Gemma‑4 基础版（非对话微调），适合进一步微调或研究。*

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**  
  作者：google | 👍 517 | 📥 0  
  *结合扩散模型与语言模型的超大规模多模态生成模型，主攻图像+文本到文本任务。*

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)**  
  作者：ideogram‑ai | 👍 489 | 📥 7,170  
  *Ideogram 4 代图像生成模型，FP8 量化版，兼顾质量与部署效率。*

- **[ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4)**  
  作者：ideogram‑ai | 👍 319 | 📥 6,124  
  *Ideogram 4 的 NF4 量化版本，进一步降低显存需求。*

- **[stepfun-ai/Step-3.7-Flash](https://huggingface.co/stepfun-ai/Step-3.7-Flash)**  
  作者：stepfun‑ai | 👍 368 | 📥 50,187  
  *快速视觉‑语言模型（VLM），注重推理速度，适合实时应用。*

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)**  
  作者：bosonai | 👍 366 | 📥 19,948  
  *4B 参数文本转语音模型，基于 Qwen3 架构，合成效果自然。*

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
  作者：nvidia | 👍 374 | 📥 4,965  
  *流式自动语音识别模型，支持缓存感知，适合低延迟语音场景。*

- **[google/magenta-realtime-2](https://huggingface.co/google/magenta-realtime-2)**  
  作者：google | 👍 180 | 📥 19,806  
  *实时音频生成模型（文本→音频），TFLite 格式，适合移动端。*

- **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)**  
  作者：ByteDance | 👍 224 | 📥 305  
  *图像+文本到视频生成模型，可基于图像和提示生成连贯视频片段。*

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)**  
  作者：zai‑org | 👍 121 | 📥 0  
  *角色动画视频生成模型，基于姿态驱动，适合虚拟角色创作。*

- **[MisoLabs/MisoTTS](https://huggingface.co/MisoLabs/MisoTTS)**  
  作者：MisoLabs | 👍 195 | 📥 0  
  *高保真语音合成模型，支持多种语音风格。*

---

### 🔧 专用模型（代码、HR、领域特定）

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)**  
  作者：CohereLabs | 👍 310 | 📥 1,859  
  *专为代码生成优化的 MoE 模型，在编程任务上表现出色。*

- **[sapientinc/HRM-Text-1B](https://huggingface.co/sapientinc/HRM-Text-1B)**  
  作者：sapientinc | 👍 751 | 📥 134,752  
  *人力资源领域专用文本生成模型，涵盖 HR 相关文档与流程自动化。*

---

### 📦 微调与量化（社区微调、GGUF、低精度）

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  作者：HauhauCS | 👍 1,685 | 📥 3,057,541  
  *Qwen3.6 的无审查（Uncensored）微调版，下载量突破 300 万，极受欢迎。*

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)**  
  作者：unsloth | 👍 564 | 📥 711,706  
  *unsloth 为 Gemma‑4 提供的 GGUF 量化版，部署门槛极低，社区首选。*

- **[unsloth/gemma-4-12B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF)**  
  作者：unsloth | 👍 202 | 📥 148,252  
  *基于 QAT（量化感知训练）的 Gemma‑4 量化版，兼顾精度与推理速度。*

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)**  
  作者：unsloth | 👍 185 | 📥 0  
  *DiffusionGemma 的 GGUF 量化版，首次将扩散模型纳入 GGUF 生态。*

- **[unsloth/gemma-4-26B-A4B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF)**  
  作者：unsloth | 👍 144 | 📥 129,110  
  *26B 量级 Gemma‑4 的 QAT + GGUF 组合，为更大模型推理提供方案。*

- **[google/gemma-4-12B-it-qat-q4_0-gguf](https://huggingface.co/google/gemma-4-12B-it-qat-q4_0-gguf)**  
  作者：google | 👍 130 | 📥 96,749  
  *Google 官方提供的 QAT 加 4‑bit GGUF 量化版本，质量与效率兼得。*

- **[nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4)**  
  作者：nvidia | 👍 169 | 📥 91,117  
  *Nvidia 官方对 550B 模型采用的 NVFP4 浮点量化，显存需求大幅降低。*

- **[XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash](https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash)**  
  作者：XiaomiMiMo | 👍 88 | 📥 660  
  *小米 MiMo 模型的 FP4 动态量化版，面向边缘设备 Agent。*

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)**  
  作者：OBLITERATUS | 👍 237 | 📥 14,838  
  *去除内容限制的 Gemma‑4 微调版，附带 GGUF 支持，社区“去审查”潮流代表。*

- **[huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated)**  
  作者：huihui‑ai | 👍 144 | 📥 6,400  
  *“Abliterated”微调版 Gemma‑4，移除模型内在偏见，提升输出自由度。*

- **[Comfy-Org/Ideogram-4](https://huggingface.co/Comfy-Org/Ideogram-4)**  
  作者：Comfy‑Org | 👍 136 | 📥 0  
  *ComfyUI 节点封装版 Ideogram‑4，便于在 ComfyUI 工作流中直接调用。*

---

## 生态信号
- **Gemma‑4 生态爆发**：Google 释放了基座、指令版、扩散版及官方量化，社区 unsloth、huihui‑ai、OBLITERATUS 等迅速跟进，形成“发布‑量化‑微调”完整链条，类似之前的 LLaMA 3 生态。
- **DeepSeek 持续领跑**：V4‑Pro 以 4785 赞和 400 万下载稳坐第一，与 Nemotron 3 Ultra 一起反映大型 MoE 模型的竞争力；闭源与开源权重并存，但主流趋势仍是开放权重。
- **量化部署成为标配**：从 GGUF、NF4、FP4 到 NVFP4，质量损失控制优秀，用户更关注“能跑多远”而非“参数多大”。HauhauCS 的 Uncesored 下载破 300 万，说明社区对无限制模型的需求旺盛。
- **多模态任务走向“理解+定位”**：不仅仅是生成，LocateAnything‑3B 的定位能力、Bernini‑R 的图像→视频、SCAIL‑2 的角色动画，反映多模态正从“告诉模型画什么”进化到“跟模型一起做应用”。

---

## 值得探索
1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — 结合视觉定位与自然语言指令，开箱即用，是今年最具实用潜力的多模态模型之一，推荐在机器人、视觉搜索场景中试用。
2. **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — 当前最强开源文本模型之一，适合作为 LLM 推理与知识问答的基线，同时关注其微调生态。
3. **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — 低门槛部署 Gemma‑4 的最佳路径，本地推理性能出色，推荐作为普通用户尝鲜多模态能力的首选量化版本。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*