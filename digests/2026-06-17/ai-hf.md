# Hugging Face 热门模型日报 2026-06-17

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-17 03:58 UTC

---

好的，作为 AI 模型生态分析师，以下是基于 2026-06-17 数据生成的《Hugging Face 热门模型日报》。

---

### **Hugging Face 热门模型日报 (2026-06-17)**

#### **今日速览**

本周 Hugging Face 生态呈现三大核心趋势：**多模态 MoE 模型**成为绝对主流，不仅语言模型，连语音和视频生成模型也普遍采用混合专家架构（MoE）。**DeepSeek-V4-Pro** 以绝对优势登顶热度榜，展示了闭源级性能在开源社区的强大影响力。同时，**Unsloth** 主导的 **GGUF 量化生态**极其活跃，几乎与所有热门新模型同步推出量化版本，极大降低了部署门槛。此外，**代码推理**和**无审查**（Uncensored）模型仍保持强劲需求，社区微调版本迭代迅速。

#### **热门模型**

##### 🧠 语言模型（LLM、对话模型、指令微调）

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**
  - 作者: deepseek-ai | 点赞: 4,898 | 下载: 2,829,747
  - 一句话：本周最热模型，基于 DeepSeek V4 架构的推理增强版本，凭借强大的性能登顶热度榜，下载量接近 300 万。

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**
  - 作者: google | 点赞: 1,055 | 下载: 1,223,383
  - 一句话：Google 原生的 12B unified 模型，支持文本、图像、音频等任意输入输出（Any-to-Any），是统一多模态的重要探索。

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**
  - 作者: zai-org | 点赞: 540 | 下载: 0
  - 一句话：最新一代 GLM MoE 模型，下载量虽为 0（可能新发布），但点赞数高，说明社区关注度极高。

##### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)**
  - 作者: Qwen | 点赞: 2,139 | 下载: 3,360,615
  - 一句话：Qwen 官方 3.6 版本，35B 参数仅激活 3B，MoE 架构，是目前社区下载量最高的视觉多模态模型，被誉为“高性价比之王”。

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**
  - 作者: nvidia | 点赞: 2,107 | 下载: 98,698
  - 一句话：NVIDIA 发布的交互式目标定位模型，支持通过文本或点击精准定位图像中任意物体，应用前景广阔。

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**
  - 作者: MiniMaxAI | 点赞: 1,023 | 下载: 25,064
  - 一句话：MiniMax 新一代多模态 MoE 模型，在视觉理解与智能体任务上表现突出，是国产多模态模型的重要力量。

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**
  - 作者: google | 点赞: 951 | 下载: 375,974
  - 一句话：Gemma 4 的扩散模型变体，结合了扩散生成与自回归 LLM，专为多轮图像生成对话设计。

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)**
  - 作者: ideogram-ai | 点赞: 560 | 下载: 12,466
  - 一句话：顶级文生图模型 Ideogram 的第四代 FP8 量化版本，保持了极好的图像质量与排版能力。

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)**
  - 作者: bosonai | 点赞: 467 | 下载: 43,361
  - 一句话：4B 参数的 TTS 模型，基于 Qwen3 架构，主打高质量、多语种语音合成，是音频生成领域的亮眼新星。

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**
  - 作者: nvidia | 点赞: 483 | 下载: 5,777
  - 一句话：NVIDIA 流式语音识别模型，支持缓存感知，0.6B 小参数即可实现高精度实时 ASR。

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)**
  - 作者: zai-org | 点赞: 209 | 下载: 0
  - 一句话：基于扩散模型的角色动画生成器，支持姿态驱动视频生成，是视频生成方向的新尝试。

##### 🔧 专用模型（代码、数学、医疗、嵌入）

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**
  - 作者: moonshotai | 点赞: 808 | 下载: 102,206
  - 一句话：月之暗面推出的代码专用多模态模型，采用压缩张量技术，在代码生成与理解任务上表现优异。

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**
  - 作者: yuxinlu1 | 点赞: 1,202 | 下载: 60,921
  - 一句话：基于 Gemma 4 Coder 的社区微调版本，结合 Fable 5 和 Composer 2.5 技术，主打代码推理能力。

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)**
  - 作者: CohereLabs | 点赞: 412 | 下载: 12,129
  - 一句话：Cohere 推出的 MoE 代码模型，结合 North 系列优势，面向代码生成与推理场景。

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**
  - 作者: microsoft | 点赞: 163 | 下载: 192
  - 一句话：微软推出的长上下文模型，基于 Qwen3 微调，专为“探索代理”（Explorer SubAgent）等长文本任务优化。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**
  - 作者: WeiboAI | 点赞: 199 | 下载: 0
  - 一句话：3B 参数的数学推理模型，基于 Qwen2 架构，专注提升小模型的逻辑与数学能力。

##### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**
  - 作者: HauhauCS | 点赞: 1,897 | 下载: 2,716,651
  - 一句话：Qwen3.6 的无审查激进微调版，下载量极高，说明社区对“无限制”模型有巨大需求。

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)**
  - 作者: DavidAU | 点赞: 371 | 下载: 366,279
  - 一句话：名字超长、功能超全的社区微调版，融合了多模型特性（Claude 4.6, Opus等）与无审查思想，是社区“缝合”技术的代表。

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)**
  - 作者: OBLITERATUS | 点赞: 336 | 下载: 76,044
  - 一句话：Gemma 4 的无审查微调版本，紧随官方发布，体现了社区对 Gemini 开源模型的快速“定制化”传统。

- **[unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)**, **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)**, **[unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF)** 等
  - 一句话：Unsloth 主导的 GGUF 量化生态，几乎覆盖了本周所有热门模型，极大简化了本地部署和消费级显卡推理。

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)**
  - 作者: Jackrong | 点赞: 220 | 下载: 79,157
  - 一句话：结合 Qwen3.6 与 Opus 特性的代码模型 GGUF 版本，支持视觉，是多模态代码工具的重要补充。

#### **生态信号**

**1. 模型家族“三国杀”：DeepSeek V4、Qwen 3.6 与 Gemma 4 三分天下。** DeepSeek-V4-Pro 凭借极致的单轮推理性能登顶；Qwen3.6 系列（特别是 35B-A3B）以高性价比和多模态能力稳居下载量榜首；Google Gemma 4 则依靠 Any-to-Any 架构和强大的 Diffusion 变体（DiffusionGemma）开辟新战场，成为开源多模态的标杆。

**2. 开源权重性能逼近闭源，但闭源模型仍受追捧。** DeepSeek-V4-Pro 和 Kimi-K2.7 等模型的开源权重表现已接近甚至超越部分闭源 API，社区对其的认可度极高（高点赞与高下载）。这预示着“开源性能天花板”正在被不断推高。

**3. 量化生态与“无审查”微调是社区核心驱动力。** **Unsloth** 几乎成为热门的标准基础设置，GGUF 版本的同步发布是模型快速传播的关键。同时，“Uncensored”标签高频出现，表明开发者社区对于模型自由度、个性化定制的需求极为旺盛，已成为微调活动的主要方向。

#### **值得探索**

1.  **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**
  - **理由**：作为本周最热模型，其是体验顶尖开源语言模型性能的首选，尤其适合复杂推理、代码生成和长文写作任务。下载量近 300 万，社区验证充分。

2.  **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**
  - **理由**：探索“LLM + Diffusion”统一架构的前沿之作。它打破了传统自回归模型与扩散模型的界限，为构建真正原生多模态（图像生成、理解、对话）的智能体提供了新范式。

3.  **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**
  - **理由**：NVIDIA 出品，精准解决“视觉定位”这一痛点问题。模型仅 3B，轻量高效，可广泛应用于自动驾驶、工业质检、遥感分析等需要目标检测与定位的实际场景，实用价值极高。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*