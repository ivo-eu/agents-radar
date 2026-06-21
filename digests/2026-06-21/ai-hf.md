# Hugging Face 热门模型日报 2026-06-21

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-21 11:26 UTC

---

# Hugging Face 热门模型日报（2026-06-21）

## 📌 今日速览

本周 Hugging Face 社区呈现出 **“大模型军备竞赛”** 的态势：DeepSeek V4 Pro 以近 5000 点赞登顶，成为最受关注的开源对话模型；多模态模型全面开花，Google 的 Gemma-4-12B-it（any-to-any）、MiniMax-M3、NVIDIA LocateAnything 均跻身前列；量化与社区微调活动异常活跃，多个 GGUF 变体（如 GLM-5.2-GGUF、Kimi-K2.7-Code-GGUF）下载量激增，反映出开发者对本地部署的强烈需求。此外，GLM 5.2 系列（MoE+DSA）凭借稀疏激活架构获得关注，预示着下一代高效推理的竞争加剧。

## 🔥 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) | deepseek-ai | **4,991** / 2,611,991 | 目前开源社区综合性能最强的对话模型，在推理、代码、数学等多项基准上接近闭源水平，成为本周绝对焦点。 |
| [GLM-5.2](https://huggingface.co/zai-org/GLM-5.2) | zai-org | **1,736** / 27,413 | 智谱新架构 GLM 5.2（MoE+DSA），采用动态稀疏注意力机制，以更低计算成本实现长上下文与高效推理，广受学术界关注。 |
| [Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B) | prefeitura-rio | **327** / 190,742 | 基于 Qwen 3.5 MoE 的 397B 超大开放对话模型，虽参数规模庞大但下载量高，说明社区对大模型本地尝试仍有兴趣。 |
| [Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1) | poolside | **77** / 2,580 | poolside 最新代码生成模型，支持 vLLM/SGLang 部署，专为工业级代码编程场景优化。 |

### 🎨 多模态与生成（图像、视频、音频、文本到X）

| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it) | google | **1,111** / 1,815,370 | 首个“any-to-any”统一多模态模型，可处理文本、图像、音频输入并输出任意模态，代表 Google 多模态前沿。 |
| [MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3) | MiniMaxAI | **1,167** / 104,076 | MiniMax 第三代多模态大模型，在视觉-语言理解与生成任务上表现突出，下载量超10万。 |
| [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | **2,227** / 241,845 | NVIDIA 推出的3B参数目标定位模型，能根据自然语言指令在图像中精准定位物体，适用机器人/自动驾驶场景。 |
| [moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code) | moonshotai | **938** / 363,308 | Kimi 新模型，支持图像与代码混合理解，压缩张量技术使其推理效率更高，专攻多模态代码场景。 |
| [google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it) | google | **1,025** / 762,861 | 融合扩散模型与 LLM 的对话式图像生成模型，26B 主参数 + 4B 激活参数，在生成质量与推理速度间取得平衡。 |
| [bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b) | bosonai | **504** / 76,120 | 基于 Qwen 3 的4B参数文本转语音模型，支持情感与韵律控制，开源 TTS 领域的新标杆。 |
| [Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit) | Boogu | **75** / 374 | 基于 Diffusers 的图像编辑 LoRA，专注于中文场景下的图像局部编辑，轻量易用。 |

### 🔧 专用模型（代码、数学、医疗、嵌入）

| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B) | WeiboAI | **532** / 20,277 | 微博AI推出的3B数学推理模型，基于 Qwen 2 微调，在数学任务上媲美更大模型，适合资源受限场景。 |
| [microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT) | microsoft | **250** / 2,593 | 微软发布的4B长上下文模型，支持 Explorer SubAgent 模式，专为需要快速检索与推理的代理任务设计。 |
| [nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b) | nvidia | **597** / 27,275 | NVIDIA 流式语音识别模型，0.6B 参数支持实时缓存感知推理，适合边缘端部署。 |
| [CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0) | CohereLabs | **470** / 19,551 | Cohere 最新代码生成模型，MoE 架构，在编程竞赛与代码补全任务上表现出众。 |
| [LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M) | LiquidAI | **83** / 7,726 | Liquid 第二代表征模型，350M 参数，在句子相似度与检索任务上对标更大嵌入模型，适合 RAG 流水线。 |

### 📦 微调与量化（社区微调、GGUF、AWQ）

| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [HauhauCS/Qwen3.6-35B-A3B-Uncensored-...](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | **2,057** / 3,966,691 | 社区基于 Qwen 3.6 的 35B MoE 去审查量化版，下载量近400万，反映用户对“无限制”模型的狂热需求。 |
| [yuxinlu1/gemma-4-12B-coder-fable5-...](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) | yuxinlu1 | **2,020** / 358,677 | 将 Gemma 4 12B 代码模型量化到 GGUF，在编程任务上保持高质量，适合本地运行。 |
| [zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8) | zai-org | **117** / 217,361 | 官方 FP8 量化版 GLM-5.2，在保持精度前提下显著降低显存占用，下载量极高。 |
| [unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF) | unsloth | **213** / 32,260 | Unsloth 社区优化的 GLM-5.2 GGUF 版，专为 llama.cpp 推理优化，加速本地部署。 |
| [yuxinlu1/gemma-4-12B-agentic-...](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF) | yuxinlu1 | **219** / 21,730 | 另一款 Gemma 4 代码模型量化版，侧重 Agent 与终端交互能力，适合自动化任务。 |
| [Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF) | Jackrong | **270** / 190,993 | 将 Qwen 3.6 27B 代码模型量化并加入 MTP（多任务 prompting），下载量近20万。 |
| [bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF) | bytkim | **99** / 36,421 | Qwen 3.6 的微调量化版，引入 pi-tune 策略，在参数效率与推理速度上更进一步。 |
| [unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF) | unsloth | **147** / 42,837 | Kimi K2.7 代码模型的 GGUF 版，由 Unsloth 提供，支持压缩张量高效推理。 |

## 🌍 生态信号

- **模型家族群雄逐鹿**：本周 **DeepSeek V4** 以绝对优势领跑，但 **Qwen 3.6/3.5** 系列以社区微调版本（Uncensored、MTP、Coder）形成强大矩阵，成为最活跃的底座模型。**Gemma 4**（Google）、**GLM 5.2**（智谱）和 **Kimi K2.7**（Moonshot）均有多款衍生变体，竞争激烈。
- **开源权重持续领先**：所有上榜模型均提供完整权重（safetensors/GGUF），开源社区对“可本地运行的大模型”偏好明显，尤其量化版本下载量远超官方版（如 HauhauCS/Qwen3.6 下载近 400 万 vs 原版 DeepSeek 260 万）。
- **量化与微调成为生态核心**：GGUF 格式占据量化模型半壁江山，Unsloth、yuxinlu1 等社区贡献者将最新模型迅速适配本地推理，推动“AI 民主化”。同时，专属微调（代码、数学、Agent）模型涌现，细分领域专业化趋势加速。

## 🧪 值得探索

1. **DeepSeek-V4-Pro** — 本周最亮眼的开源 LLM，性能和社区讨论热度均为峰值。若希望体验顶级开源对话能力，不容错过。
2. **google/gemma-4-12B-it** — Google 首个 any-to-any 模型，支持多模态统一处理，架构创新性极强，适合研究多模态融合与原生跨模态能力。
3. **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** — 下载量近 400 万的现象级社区量化模型，证明了用户对“低门槛 + 高自由度”模型的巨大需求，适合研究安全对齐边界与模型蒸馏技术。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*