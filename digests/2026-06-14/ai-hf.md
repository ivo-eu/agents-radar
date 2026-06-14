# Hugging Face 热门模型日报 2026-06-14

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-14 11:10 UTC

---

# 🗞️ Hugging Face 热门模型日报 | 2026-06-14

## 今日速览

本周 Hugging Face 社区迎来两股强势力：**DeepSeek-V4-Pro** 以 4,818 点赞和 307 万下载强势登顶，成为最受关注的文本生成模型；同时 **Google Gemma-4 系列**全面开花，其 any-to-any 统一架构搭配社区量化版本（unsloth GGUF）占据了趋势榜近三分之一席位。多模态赛道持续升温，**NVIDIA LocateAnything-3B** 凭借“定位任意物体”的独特能力拿下近 2,000 点赞，而 **Ideogram-4** 系列在图像生成领域稳居焦点。此外，**Qwen3.6 系列**的社区魔改版（如 uncensored、代码优化）产生了大量衍生模型，反映开源生态正在快速向“基座+微调+量化”的全链路外溢。

---

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) | deepseek-ai | 4,818 | 3,075,369 | DeepSeek 第四代旗舰，对话与推理能力极强，社区热度断层第一。 |
| [CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0) | CohereLabs | 361 | 9,932 | Cohere 面向代码的 MoE 小模型，轻量高效，适合本地编码助手。 |
| [nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro) | nex-agi | 242 | 3,396 | 基于 Qwen3.5 MoE 架构的指令模型，强调多模态与文本生成平衡。 |
| [nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini) | nex-agi | 198 | 7,010 | 相同系列的小型版本，适合资源受限场景。 |
| [XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash](https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro-FP4-DFlash) | XiaomiMiMo | 108 | 4,108 | 小米自研 MoE 模型，带 Agent 能力，采用 FP4 极低精度量化。 |

### 🎨 多模态与生成（图像、视频、音频、文本到X）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | 1,977 | 75,201 | 可定位图像中任意物体的多模态模型，交互式视觉感知新标杆。 |
| [google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it) | google | 738 | 198,912 | Google 融合扩散与 LLM 的 26B 多模态模型，支持图像理解与生成。 |
| [moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code) | moonshotai | 562 | 15,145 | Kimi 的代码增强版，注重图像特征提取与压缩推理。 |
| [google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it) | google | 1,001 | 1,084,405 | Gemma-4 统一架构的指令版本，支持任何模态输入输出。 |
| [MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3) | MiniMaxAI | 444 | 6,643 | MiniMax 第三代多模态大模型，文本图像双向理解。 |
| [ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8) | ideogram-ai | 525 | 8,263 | Ideogram 最新图像生成模型，FP8 量化版，效率与质量兼顾。 |
| [ideogram-ai/ideogram-4-nf4](https://huggingface.co/ideogram-ai/ideogram-4-nf4) | ideogram-ai | 334 | 3,763 | 同上，NF4 量化版本，进一步降低显存需求。 |
| [bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b) | bosonai | 418 | 35,122 | 4B 参数 TTS 模型，基于 Qwen3 架构，文本转语音清晰自然。 |
| [prefeitura-rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B) | prefeitura-rio | 184 | 112,371 | 397B 巨量多模态模型（Qwen3.5 MoE 变体），开放权重引关注。 |
| [zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2) | zai-org | 164 | 0 | 基于姿态驱动的角色动画生成模型（图像→视频），但无下载量。 |
| [ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R) | ByteDance | 235 | 459 | 字节跳动图像→视频生成模型，新架构 arxiv 论文同步发表。 |
| [HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | 1,775 | 2,516,709 | Qwen3.6 的无审查 MoE 变体，下载量极高，社区“去限制”需求旺盛。 |
| [DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF) | DavidAU | 327 | 375,966 | 超长名的社区魔改版，融合多个模型风格，追求极致代码与无限制。 |
| [Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF) | Jackrong | 168 | 33,720 | Qwen3.6 的代码优化版，GGUF 格式适合本地部署。 |

### 🔧 专用模型（代码、数学、医疗、嵌入、语音）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b) | nvidia | 406 | 4,505 | NVIDIA 的流式语音识别模型（0.6B），支持缓存感知实时 ASR。 |

（注：部分代码模型已归入语言模型分类，此表只保留纯专用任务模型。）

### 📦 微调与量化（社区微调、GGUF、AWQ）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF) | unsloth | 253 | 80,118 | unsloth 提供的 diffusiongemma 量化版，方便本地运行。 |
| [unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF) | unsloth | 584 | 926,372 | Gemma-4 最流行的 GGUF 量化版，下载量近百万。 |
| [unsloth/gemma-4-12B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF) | unsloth | 220 | 255,424 | Gemma-4 的 QAT（训练后量化）GGUF 版本。 |
| [unsloth/gemma-4-26B-A4B-it-qat-GGUF](https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF) | unsloth | 152 | 317,261 | Gemma-4 26B 的 QAT 量化版，兼顾大小与性能。 |
| [OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED) | OBLITERATUS | 287 | 60,949 | 社区“去除限制”的 Gemma-4 魔改版，同时支持 GGUF。 |
| [huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated) | huihui-ai | 152 | 8,717 | 另一款“abliterated”（移除安全+过滤）的 Gemma-4 变体。 |
| [Jiunsong/supergemma4-26b-uncensored-gguf-v2](https://huggingface.co/Jiunsong/supergemma4-26b-uncensored-gguf-v2) | Jiunsong | 822 | 96,437 | 26B 无审查 GGUF 版本，声称速度快且权重开放。 |
| [Comfy-Org/Ideogram-4](https://huggingface.co/Comfy-Org/Ideogram-4) | Comfy-Org | 146 | 0 | ComfyUI 工作流适配的 Ideogram-4 封装（无下载）。 |
| [RazzzHF/Realism_Engine_Ideogram_4](https://huggingface.co/RazzzHF/Realism_Engine_Ideogram_4) | RazzzHF | 91 | 0 | 基于 Ideogram-4 的写实风格引擎，暂未开放下载。 |

---

## 生态信号

1. **模型家族势头**：Google Gemma-4 全面爆发——从原始权重到指令微调、any-to-any 统一架构，再到 unsloth 批量推出的 GGUF 量化版本，Gemma-4 已成为当前最活跃的广义基座家族。DeepSeek-V4-Pro 则以绝对点赞数巩固了中国厂商在开源社区的影响力。Qwen3.6 系列通过社区“无审查”和代码特化魔改，催生了大量衍生产品，显示其架构可塑性极强。
2. **开源 vs 闭源**：本周热门模型几乎全部为开源权重（含社区微调），仅少量链接暂时无下载。Meta/Google/DeepSeek 等头部玩家的开放策略继续挤压闭源模型空间，但社区“abliterated”（移除安全限制）现象也反映出用户对开放式使用的强烈需求。
3. **量化与微调活动**：unsloth 一家贡献了 5 个 GGUF 量化模型，下载量普遍较高，说明本地部署仍是硬需求。此外，多个模型以 “uncensored” 或 “heretic” 为卖点，这类微调版本下载量甚至超过原始模型（如 HauhauCS 的 Qwen3.6 变体下载达 250 万），构成独特的“后训练生态”。

---

## 值得探索

- **🔍 NVIDIA LocateAnything-3B**：它不只是一个多模态模型，而是首个“通用定位器”——可给定描述在图像中高亮任意物体，对机器人、自动驾驶、视觉问答等下游任务意义重大，下载量仅 7.5 万仍有巨大潜力。
- **🧠 DeepSeek-V4-Pro**：作为当前 Hugging Face 最具人气的模型，其 307 万下载量说明已在实际应用中被广泛测试，适合需要高智商对话与推理的开发者。
- **🎬 ByteDance/Bernini-R**：字节最新图像→视频生成模型，配合论文 (arxiv:2605.22344) 可深入理解其架构创新，是追踪视频生成前沿的好起点。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*