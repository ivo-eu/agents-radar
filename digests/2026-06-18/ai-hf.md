# Hugging Face 热门模型日报 2026-06-18

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-18 03:18 UTC

---

好的，这是根据 2026-06-18 Hugging Face Hub 热门模型数据生成的日报。

---

# Hugging Face 热门模型日报 — 2026-06-18

## 🧭 今日速览

本周 Hugging Face 生态持续爆发，**DeepSeek-V4-Pro** 以近 5000 点赞登顶，成为社区最关注的纯文本开源 LLM；视觉语言模型（VLM）竞争白热化，**Qwen3.6-35B-A3B** 下载量突破 368 万，**Google Gemma-4-12B-it** 的 any-to-any 架构引发大量讨论。MoE 架构成为多模态模型的主流选择，同时 **GGUF 量化版** 和 **Uncensored 微调版** 的社区二创异常活跃，反映出开发者对本地部署和指令越狱的强烈需求。此外，**NVIDIA LocateAnything-3B** 在视觉定位细分领域获得了 2141 点赞，表明专用小模型仍有巨大潜力。

## 🔥 热门模型分类整理

### 🧠 语言模型（LLM / 对话 / 指令）

- **[DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — deepseek-ai | ★ 4,926 | 📥 2,804,646  
  DeepSeek 最新旗舰开源对话模型，本周热度最高，参数规模大、推理能力强。

- **[GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — zai-org | ★ 1,052 | 📥 666  
  MoE 架构的对话模型，基于 GLM 系列改进，下载量虽少但获高质量社区关注。

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — CohereLabs | ★ 422 | 📥 13,449  
  Cohere 推出的代码专用小型 MoE 模型，适合代码生成与理解。

- **[Microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — microsoft | ★ 186 | 📥 537  
  微软推出的 4B 参数长上下文 SFT 模型，专注高效上下文处理。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — WeiboAI | ★ 316 | 📥 1,950  
  数学推理专用小模型，基于 Qwen2 微调，适合数学解题与逻辑推理。

### 🎨 多模态与生成（图像、视频、音频、文本到 X）

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)** — Qwen | ★ 2,157 | 📥 3,683,883  
  Qwen 最新的 MoE 视觉语言模型，355B 总参、3B 激活，下载量本周最高。

- **[NVIDIA/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — nvidia | ★ 2,141 | 📥 130,389  
  NVIDIA 推出的视觉定位模型，可基于文本/图像指代定位目标物体。

- **[Google/Gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** — google | ★ 1,072 | 📥 922,952  
  Gemma 4 的指令微调版，支持 any-to-any 多模态输入输出，定位通才。

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** — MiniMaxAI | ★ 1,066 | 📥 42,198  
  MiniMax 多模态 MoE 模型，支持图像+文本理解，强调 Agent 能力。

- **[Google/DiffusionGemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — google | ★ 980 | 📥 460,173  
  扩散+语言混合架构的视觉语言模型，26B 总参、4B 激活，适合图像对话。

- **[MoonshotAI/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** — moonshotai | ★ 849 | 📥 172,727  
  Kimi 系列代码视觉模型，支持图像+代码联合理解，下载量高。

- **[Ideogram-ai/Ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — ideogram-ai | ★ 569 | 📥 15,477  
  最新文本到图像生成模型，FP8 格式，生成质量领先。

- **[BosonAI/Higgs-Audio-v3-TTS-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — bosonai | ★ 482 | 📥 40,812  
  基于 Qwen3 的 TTS 模型，4B 参数，合成语音自然度高。

- **[Prefeitura-Rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)** — prefeitura-rio | ★ 319 | 📥 189,986  
  Rio 开源 MoE 多模态模型，397B 总参，支持图像对话，下载活跃。

- **[Nex-AGI/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — nex-agi | ★ 317 | 📥 5,579  
  Nex 推出的 MoE 多模态对话模型，强调推理能力。

- **[Zyphra/ZONOS2](https://huggingface.co/Zyphra/ZONOS2)** — Zyphra | ★ 108 | 📥 629  
  Apache-2.0 许可的 TTS 模型，适合商用语音合成。

- **[Zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** — zai-org | ★ 223 | 📥 0  
  图像到视频生成模型，基于 Diffusion 和姿态驱动动画，暂未开放下载。

- **[Nex-AGI/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** — nex-agi | ★ 238 | 📥 9,804  
  Nex-N2 的小型版，兼顾速度与多模态能力。

### 🔧 专用模型（代码、数学、ASR、嵌入）

- **[Yuxinlu1/Gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** — yuxinlu1 | ★ 1,500 | 📥 146,784  
  基于 Gemma-4 微调的代码专用 GGUF 模型，搭配 composer 优化，推理速度快。

- **[NVIDIA/Nemotron-3.5-ASR-Streaming-0.6B](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** — nvidia | ★ 522 | 📥 7,195  
  流式语音识别模型，0.6B 参数，低延迟、缓存感知，适合实时 ASR。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** — WeiboAI | ★ 316 | 📥 1,950  
  (已在 LLM 中列出) 数学推理专用模型，此处不重复。

### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — HauhauCS | ★ 1,940 | 📥 2,876,624  
  Qwen3.6 的激进 uncensored 微调版，GGUF 格式，下载量极高。

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** — DavidAU | ★ 384 | 📥 427,359  
  融合多个模型的夸张微调版，支持 uncensored 和代码推理，GGUF 量化。

- **[Unsloth/Gemma-4-12B-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — unsloth | ★ 646 | 📥 579,224  
  Gemma-4-12B-it 的 GGUF 量化版，unsloth 出品，本地部署首选。

- **[Unsloth/DiffusionGemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)** — unsloth | ★ 299 | 📥 136,634  
  DiffusionGemma 的 GGUF 量化版，降低多模态模型推理门槛。

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — OBLITERATUS | ★ 341 | 📥 78,333  
  对 Gemma-4-12B-it 进行 uncensored 微调并量化，社区热门。

- **[Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)** — Jackrong | ★ 237 | 📥 99,909  
  基于 Qwen3.6 的多模态代码模型，GGUF 量化，兼顾视觉与代码。

- **[Lordx64/Qwable-v1](https://huggingface.co/lordx64/Qwable-v1)** — lordx64 | ★ 104 | 📥 319  
  社区微调的 Qwen3.5 MoE 多模态模型，GGUF 格式。

- **[Unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF)** — unsloth | ★ 128 | 📥 23,956  
  Kimi-K2.7-Code 的 GGUF 量化版，便于本地代码推理。

- **[Unsloth/MiniMax-M3-GGUF](https://huggingface.co/unsloth/MiniMax-M3-GGUF)** — unsloth | ★ 98 | 📥 20,504  
  MiniMax-M3 的 GGUF 量化版，多模态 Agent 模型本地化。

- **[Mia-AiLab/Qwable-3.6-27b](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b)** — Mia-AiLab | ★ 80 | 📥 646  
  Qwen3.6 的社区 GGUF 量化版，27B 参数。

## 🌐 生态信号

- **MoE 占据主流**：本周 Top10 中超过一半采用混合专家架构（Qwen3.6-35B-A3B、DiffusionGemma-26B-A4B、Rio-3.5-397B 等），MoE 在部署效率与性能平衡上优势明显。
- **多模态成为标配**：从 Google、Qwen 到 NVIDIA，几乎所有新模型都内建视觉理解能力，纯文本模型热度相对下降。
- **开源权重继续领先**：DeepSeek-V4-Pro 和 Qwen3.6 权重完全开放，社区二创（uncensored、GGUF）极为活跃，闭源模型（如某些 API 模型）未进热度榜。
- **量化与微调深水区**：Unsloth 霸占多个量化榜，社区 uncensored 微调版下载量惊人（如 HauhauCS 版近 300 万），反映出开发者对“无审查”和本地高速推理的双重偏好。

## 🔭 值得探索

1. **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)** — **效率之王**  
   仅 3B 激活参数达到 355B 模型能力，是研究 MoE 和 Edge 部署的绝佳样本，下载量说明一切。

2. **[NVIDIA/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — **实用定位工具**  
   视觉定位专用模型，精准度领先，可直接用于机器人、图像标注等场景，点赞数（2141）验证其实际价值。

3. **[Google/DiffusionGemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** — **架构创新**  
   Diffusion + Language 混合架构，打开多模态生成的新方向，适合对模型设计感兴趣的研究者。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*