# Hugging Face 热门模型日报 2026-08-01

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-08-01 00:12 UTC

---

# Hugging Face 热门模型日报（2026-08-01）

## 今日速览

- 🔥 月之暗面 **Kimi-K3** 以 9,273 次周点赞断层式登顶，成为本周最受关注的多模态模型。
- 🚀 **DeepSeek-V4-Flash** 正式版下载量逼近 300 万，0731 新快照版本同步发布，开源社区热度持续攀升。
- 🧬 **Qwen3.5/3.6** 生态全面爆发，多个 Uncensored 微调版与 GGUF 量化版霸榜下载前列。
- 🇨🇳 **GLM-5.2、Unlimited-OCR** 等国产模型表现强势，中文开源势力进一步壮大。
- ⚡ 极限量化成为新趋势：三值（Ternary）/2-bit、NVFP4、W2 等方案涌现，本地部署门槛大幅降低。

---

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

| 模型（作者 / 链接） | 点赞 | 下载 | 一句话说明 |
|---|---|---|---|
| deepseek-ai / [DeepSeek-V4-Flash](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash) | 1,923 | 2,923,499 | 本周下载量近 300 万的爆款开放权重 LLM，主打高性能文本生成与对话。 |
| deepseek-ai / [DeepSeek-V4-Flash-0731](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731) | 967 | 0 | V4-Flash 的 7 月 31 日新快照版，刚上线、下载量尚未积累，值得跟踪的技术迭代。 |
| zai-org / [GLM-5.2](https://huggingface.co/zai-org/GLM-5.2) | 4,706 | 1,651,533 | 智谱新一代 MoE-DSA 架构旗舰模型，纯语言模型热度第一，下载超 165 万。 |
| poolside / [Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1) | 862 | 76,212 | poolside 的企业级文本生成模型新版本，面向软件开发与商业场景。 |
| upstage / [Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B) | 714 | 12,911 | Upstage 的 250B 级开源权重大模型，主打多语言文本生成能力。 |
| Nanbeige / [Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B) | 595 | 26,928 | 北冥实验室的 3B 轻量 LLM，小参数易部署，适合资源受限场景。 |
| XYZAILab / [XYZ-Aquila-mini](https://huggingface.co/XYZAILab/XYZ-Aquila-mini) | 352 | 579 | 基于 Qwen3.5 MoE 架构的小型模型，主打轻量高效。 |
| XYZAILab / [XYZ-Aquila-pro](https://huggingface.co/XYZAILab/XYZ-Aquila-pro) | 326 | 869 | Aquila 系列专业版，集成 agentic-search（智能体搜索）能力。 |

### 🎨 多模态与生成（图像、视频、音频、文本到 X）

| 模型（作者 / 链接） | 点赞 | 下载 | 一句话说明 |
|---|---|---|---|
| moonshotai / [Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3) | 9,273 | 493,481 | 月之暗面新一代多模态模型，本周热度断层第一，支持图文理解与特征提取。 |
| thinkingmachines / [Inkling](https://huggingface.co/thinkingmachines/Inkling) | 1,663 | 57,259 | Thinking Machines 的视觉-语言多模态对话模型，社区关注度持续走高。 |
| thinkingmachines / [Inkling-Small](https://huggingface.co/thinkingmachines/Inkling-Small) | 194 | 2,971 | Inkling 的小尺寸版本，适合多模态应用的轻量部署。 |
| microsoft / [Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B) | 234 | 2,726 | 微软的 computer-use 多模态模型，可理解屏幕并执行 GUI 操作任务。 |
| microsoft / [Mage-VL](https://huggingface.co/microsoft/Mage-VL) | 149 | 5,650 | 微软的视觉语言理解模型，面向通用多模态推理。 |
| owensong / [Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2) | 348 | 1,449 | 面向 CPU/边缘设备的本地 TTS 模型，主打低延迟语音合成。 |
| owensong / [Inflect-Nano-v2](https://huggingface.co/owensong/Inflect-Nano-v2) | 121 | 802 | Inflect 系列极致轻量版 TTS，进一步压缩体积、适配弱算力设备。 |
| Audio8 / [Audio8-TTS-Preview-0.6b](https://huggingface.co/Audio8/Audio8-TTS-Preview-0.6b) | 151 | 2,481 | Audio8 的 0.6B 语音合成预览版，基于 ArkTTS 架构。 |
| Comfy-Org / [Mage-Flow](https://huggingface.co/Comfy-Org/Mage-Flow) | 106 | 60,162 | ComfyUI 生态的 diffusion 单文件图像生成模型，配套微软 Mage 系列。 |

### 🔧 专用模型（代码、数学、医疗、嵌入）

| 模型（作者 / 链接） | 点赞 | 下载 | 一句话说明 |
|---|---|---|---|
| baidu / [Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR) | 3,660 | 2,513,603 | 百度开源的超大规模 OCR 模型，251 万下载，支持开放类别文本识别。 |
| Kwaipilot / [KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev) | 369 | 10,241 | 基于 Qwen3.5 MoE 的代码生成开发版模型，专注编程场景。 |
| microsoft / [VibeVoice-ASR-BitNet](https://huggingface.co/microsoft/VibeVoice-ASR-BitNet) | 133 | 5,464 | 微软采用 BitNet 架构的语音识别模型，支持 GGUF/GGML 格式。 |

### 📦 微调与量化（社区微调、GGUF、AWQ）

| 模型（作者 / 链接） | 点赞 | 下载 | 一句话说明 |
|---|---|---|---|
| DavidAU / [Qwen3.6-27B-Fable-Fusion-711-…-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF) | 1,142 | 1,119,057 | 社区对 Qwen3.6-27B 的深度微调，主打 Uncensored 风格，GGUF 格式下载超 111 万。 |
| HauhauCS / [Qwen3.6-35B-A3B-…-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | 3,205 | 1,835,931 | Qwen3.6-35B MoE 热门社区微调版，184 万下载，风格激进、支持视觉输入。 |
| prism-ml / [Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf) | 1,124 | 712,835 | 三值（Ternary）/2-bit 极限量化 27B 模型，71 万下载，探索超低比特推理。 |
| LuffyTheFox / [Qwen3.6-35B-A3B-…-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF) | 270 | 212,426 | Qwen3.6-35B-A3B MoE 的非审查微调 GGUF 版，社区热度颇高。 |
| DavidAU / [Qwen3.5-9B-…-GGUF](https://huggingface.co/DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF) | 172 | 261,856 | Qwen3.5-9B 社区微调 GGUF，集成 NEO Imatrix 与 MTP 优化，26 万下载。 |
| unsloth / [Kimi-K3-GGUF](https://huggingface.co/unsloth/Kimi-K3-GGUF) | 226 | 36,180 | unsloth 出品的 Kimi-K3 GGUF 量化版，方便本地一键部署。 |
| unsloth / [Kimi-K3](https://huggingface.co/unsloth/Kimi-K3) | 215 | 1,044 | unsloth 优化适配的 Kimi-K3 版本，便于社区微调与二次开发。 |
| nota-ai / [Solar-Open2-250B-Nota-NVFP4](https://huggingface.co/nota-ai/Solar-Open2-250B-Nota-NVFP4) | 151 | 18,531 | Nota AI 对 Solar-Open2-250B 的 NVFP4 量化版，适配 vLLM 高效推理。 |
| unsloth / [DeepSeek-V4-Flash-0731-GGUF](https://huggingface.co/unsloth/DeepSeek-V4-Flash-0731-GGUF) | 173 | 0 | DeepSeek-V4-Flash-0731 的 GGUF 量化版，刚发布、待社区下载。 |
| EschaLabs / [Qwen3.6-35B-A3B-Escha-W2](https://huggingface.co/EschaLabs/Qwen3.6-35B-A3B-Escha-W2) | 106 | 599 | Qwen3.6-35B-A3B 的 W2（2-bit）量化实验版，追求极致压缩的 MoE 路线。 |

---

## 生态信号

模型生态呈现"基座集中、周边繁荣"的特征：**DeepSeek-V4、Kimi-K3、GLM-5.2、Solar-Open2** 四大新基座同台竞技，而 **Qwen3.6** 凭借开放的权重策略成为社区微调与量化最活跃的家族。开源权重路线持续占据主导，且中方实验室（月之暗面、DeepSeek、智谱、百度）在开放程度与社区热度上全面领先。量化侧，**GGUF 已成本地部署的事实标准**，unsloth 深度参与头部模型格式适配；Ternary-Bonsai（三值/2-bit）、NVFP4、W2 等极限压缩方案涌现，预示大模型在小内存设备上的普及将显著加速。

---

## 值得探索

1. **moonshotai/Kimi-K3**（[链接](https://huggingface.co/moonshotai/Kimi-K3)）— 本周热度断层第一，将压缩张量（compressed-tensors）与多模态结合，技术路线极具研究价值。
2. **deepseek-ai/DeepSeek-V4-Flash**（[链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash)）— 下载量近 300 万的现象级开放权重模型，可同时关注其 0731 快照版（[链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731)）以追踪最新迭代。
3. **prism-ml/Ternary-Bonsai-27B-gguf**（[链接](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)）— 三值/2-bit 极限量化代表作，是低成本本地部署与量化算法研究的最佳样本。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*