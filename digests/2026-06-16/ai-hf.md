# Hugging Face 热门模型日报 2026-06-16

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-16 05:20 UTC

---

# 🤗 Hugging Face 热门模型日报 (2026-06-16)

---

## 今日速览

今日榜单亮点纷呈：**DeepSeek-V4-Pro** 以近 5000 赞和 293 万下载量强势登顶，成为社区最受关注的通用大语言模型；Google 的 **DiffusionGemma-26B-A4B-it** 多模态 MoE 模型紧随其后，下载量超 31 万，且其 GGUF 量化版本同时上榜；NVIDIA 推出的 **LocateAnything-3B** 定位模型以 2062 赞成为黑马，展示了特定视觉任务的吸引力；社区微调与量化活动异常活跃，多个基于 Qwen 3.6、Gemma 4 的无审查/代码优化版本（如 HauhauCS、DavidAU）下载量极高，反映开发者对可部署、高性能模型的迫切需求。

---

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

| 模型 | 作者 | 👍点赞 | 📥下载 | 一句话说明 |
|------|------|--------|--------|------------|
| [**DeepSeek-V4-Pro**](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) | deepseek-ai | 4,867 | 2,934,763 | DeepSeek 最新一代旗舰对话模型，性能与规模均达新高度，本周绝对焦点。 |
| [**Nex-N2-Pro**](https://huggingface.co/nex-agi/Nex-N2-Pro) | nex-agi | 288 | 3,681 | 基于 Qwen3.5-MoE 的增强版对话模型，定位专业场景推理。 |
| [**Nex-N2-mini**](https://huggingface.co/nex-agi/Nex-N2-mini) | nex-agi | 220 | 8,260 | 同系列轻量版，兼顾效率与能力，适合资源受限部署。 |
| [**microsoft/FastContext-1.0-4B-SFT**](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT) | microsoft | 113 | 13 | 微软推出的超长上下文模型（基于 Qwen3），擅长处理海量文本任务。 |

### 🎨 多模态与生成（图像、视频、音频、文本到X）

| 模型 | 作者 | 👍点赞 | 📥下载 | 一句话说明 |
|------|------|--------|--------|------------|
| [**google/diffusiongemma-26B-A4B-it**](https://huggingface.co/google/diffusiongemma-26B-A4B-it) | google | 892 | 311,788 | 结合扩散模型与 Gemma 的多模态对话 MoE，支持图文混合输入输出。 |
| [**MiniMaxAI/MiniMax-M3**](https://huggingface.co/MiniMaxAI/MiniMax-M3) | MiniMaxAI | 854 | 14,312 | MiniMax 新一代多模态模型，在视觉-语言理解与生成上表现突出。 |
| [**nvidia/LocateAnything-3B**](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | 2,062 | 86,968 | NVIDIA 精确定位模型，可识别并定位图像中任意目标，应用广泛。 |
| [**prefeitura-rio/Rio-3.5-Open-397B**](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B) | prefeitura-rio | 304 | 188,723 | 巴西政府推出的超大 MoE 多模态模型，开源权重且下载量惊人。 |
| [**google/gemma-4-12B-it**](https://huggingface.co/google/gemma-4-12B-it) | google | 1,038 | 1,160,435 | Gemma 4 指令微调版，统一处理文本、图像、音频等任意输入输出。 |
| [**google/gemma-4-12B**](https://huggingface.co/google/gemma-4-12B) | google | 553 | 250,498 | 同系列基座模型，适合下游微调与研究。 |
| [**ideogram-ai/ideogram-4-fp8**](https://huggingface.co/ideogram-ai/ideogram-4-fp8) | ideogram-ai | 548 | 10,748 | Ideogram 第四代文生图模型 FP8 量化版，兼顾质量与速度。 |
| [**ideogram-ai/ideogram-4-nf4**](https://huggingface.co/ideogram-ai/ideogram-4-nf4) | ideogram-ai | 345 | 4,224 | 更极致的 NF4 量化版本，进一步降低显存需求。 |
| [**zai-org/SCAIL-2**](https://huggingface.co/zai-org/SCAIL-2) | zai-org | 191 | 0 | 基于扩散模型的图像到视频生成工具，支持角色动作驱动。 |
| [**Zyphra/ZONOS2**](https://huggingface.co/Zyphra/ZONOS2) | Zyphra | 91 | 414 | 开源文本转语音模型，Apache-2.0 许可，语音合成新选择。 |
| [**bosonai/higgs-audio-v3-tts-4b**](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b) | bosonai | 445 | 38,429 | 基于 Qwen3 的高质量 TTS 模型，支持多语言语音生成。 |

### 🔧 专用模型（代码、数学、医疗、嵌入）

| 模型 | 作者 | 👍点赞 | 📥下载 | 一句话说明 |
|------|------|--------|--------|------------|
| [**moonshotai/Kimi-K2.7-Code**](https://huggingface.co/moonshotai/Kimi-K2.7-Code) | moonshotai | 753 | 56,750 | 月之暗面推出的代码生成模型，采用压缩张量技术，效率高。 |
| [**yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) | yuxinlu1 | 718 | 20,207 | 基于 Gemma 4 的代码增强版 GGUF 量化模型，推理性能出色。 |
| [**CohereLabs/North-Mini-Code-1.0**](https://huggingface.co/CohereLabs/North-Mini-Code-1.0) | CohereLabs | 392 | 11,145 | Cohere 针对代码场景优化的 MoE 轻量模型，适合编程任务。 |
| [**Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF**](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF) | Jackrong | 202 | 62,469 | 基于 Qwen3.6 的代码专家模型 GGUF 版，支持 MTP 推理加速。 |
| [**Jackrong/Qwopus3.6-27B-v2-MTP-GGUF**](https://huggingface.co/Jackrong/Qwopus3.6-27B-v2-MTP-GGUF) | Jackrong | 312 | 184,446 | 同一作者的迭代版，进一步优化代码推理能力。 |
| [**nvidia/nemotron-3.5-asr-streaming-0.6b**](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b) | nvidia | 424 | 5,200 | NVIDIA 流式语音识别模型，支持缓存感知实时转录。 |
| [**microsoft/FastContext-1.0-4B-SFT**](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT) | microsoft | 113 | 13 | 长上下文专用模型，适合文档分析与代理系统。 |

### 📦 微调与量化（社区微调、GGUF、AWQ）

| 模型 | 作者 | 👍点赞 | 📥下载 | 一句话说明 |
|------|------|--------|--------|------------|
| [**HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | 1,859 | 2,697,882 | 基于 Qwen3.6 的无审查激进调教版，下载量极高，社区反响强烈。 |
| [**DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF**](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF) | DavidAU | 359 | 369,526 | 超大混合微调模型，集成多来源知识，GGUF 量化便于本地部署。 |
| [**unsloth/diffusiongemma-26B-A4B-it-GGUF**](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF) | unsloth | 276 | 107,243 | DiffusionGemma 官方 GGUF 量化版，降低门槛便于本地运行。 |
| [**unsloth/gemma-4-12b-it-GGUF**](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF) | unsloth | 618 | 980,781 | Gemma 4 指令版 GGUF，社区最热门的量化版本之一，下载近百万。 |
| [**unsloth/gemma-4-12B-it-qat-GGUF**](https://huggingface.co/unsloth/gemma-4-12B-it-qat-GGUF) | unsloth | 243 | 288,390 | 基于 QAT 训练后的量化版，精度更优。 |
| [**OBLITERATUS/Gemma-4-12B-OBLITERATED**](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED) | OBLITERATUS | 325 | 70,732 | 社区对 Gemma 4 的“粉碎性”微调版本，搭配 GGUF 双格式发布。 |
| [**unsloth/Kimi-K2.7-Code-GGUF**](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF) | unsloth | 103 | 9,327 | Kimi 代码模型的官方 GGUF 版，方便 CPU/低显存运行。 |
| [**unsloth/MiniMax-M3-GGUF**](https://huggingface.co/unsloth/MiniMax-M3-GGUF) | unsloth | 84 | 14,799 | MiniMax M3 多模态模型的 GGUF 量化版本。 |

---

## 生态信号

1. **模型家族竞争加剧**：DeepSeek V4 以绝对优势领跑通用 LLM 赛道；Google 的 **Gemma 4** 与 **DiffusionGemma** 系列在多模态和量化版本上布局密集（已有 5 个及以上衍生物上榜）；**Qwen 3.6** 生态极其活跃，成为社区微调的首选基座（约 1/3 的热门模型基于它）。
2. **开源权重成为主流**：榜单前 10 名全部为开源模型（DeepSeek V4、Gemma 4、DiffusionGemma、LocateAnything 等），说明社区对可自部署、可定制权重的偏好压倒闭源 API。
3. **量化与微调活动空前活跃**：Unsloth 几乎为每个热门模型提供 GGUF 版本，下载量常超原版；社区无审查微调（HauhauCS、DavidAU、OBLITERATUS）下载量达百万级别，反映开发者在安全合规之外对“自由”模式的强烈需求。
4. **垂直任务模型崛起**：NVIDIA 的定位模型、Ideogram 图像生成、Higgs Audio TTS 等专用模型点赞数不输通用模型，说明社区正从“大而全”向“精而专”扩散。

---

## 值得探索

1. **[DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
   当前最强开源对话模型，支持超长上下文与复杂推理，可作为通用 Agent 或研究基座。其 293 万下载量已验证社区信赖度。

2. **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**  
   融合扩散模型与 LLM 的 MoE 多模态架构，在图像理解、生成与对话中表现惊艳，是探索多模态 Agent 的理想起点。

3. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   仅有 3B 参数却可实现高精度目标定位，轻量高效，适合集成到机器人、辅助视觉等实际场景，值得关注其开放权重带来的应用潜力。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*