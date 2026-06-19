# Hugging Face 热门模型日报 2026-06-19

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-19 12:58 UTC

---

# Hugging Face 热门模型日报 — 2026-06-19

## 📰 今日速览

本周 Hugging Face 社区热度飙升，**DeepSeek-V4-Pro** 以近 5000 点赞登顶，展现 MoE 架构的强劲吸引力。**Gemma-4 系列** 持续爆发，多个微调与量化版本（12B、26B、coder）占据榜单半壁江山，特别是社区通过 GGUF 格式大力推动边缘部署。**多模态模型**全面崛起，视觉 - 语言模型（如 Kimi-K2.7-Code、MiniMax-M3）、流式语音识别（NVIDIA nemotron）以及图像 / 视频生成模型（SCAIL-2、Higgs Audio-TTS）均获得高关注。此外，**代码 / 数学专用模型**（Gemma-4 Coder、VibeThinker-3B）和**无审查微调变体**（Qwen3.6 Unce本Censored）成为社区二次创作的焦点。

## 🧠 语言模型

### deepseek-ai/DeepSeek-V4-Pro
- **作者**: deepseek-ai | **点赞**: 4,959 | **下载**: 3,015,772
- [HF链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)
- **一句话**: DeepSeek 第四代 MoE 旗舰，综合性能对标 GPT-5，本周热度断层第一，下载量突破 300 万。

### zai-org/GLM-5.2
- **作者**: zai-org | **点赞**: 1,445 | **下载**: 11,871
- [HF链接](https://huggingface.co/zai-org/GLM-5.2)
- **一句话**: 智谱最新对话模型，MoE-DSA 架构，官方原版权重开放后迅速跻身前列。

### CohereLabs/North-Mini-Code-1.0
- **作者**: CohereLabs | **点赞**: 454 | **下载**: 17,693
- [HF链接](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)
- **一句话**: Cohere 推出的轻量级代码生成模型，MoE 架构，专注于高效编程辅助。

### WeiboAI/VibeThinker-3B
- **作者**: WeiboAI | **点赞**: 426 | **下载**: 12,148
- [HF链接](https://huggingface.co/WeiboAI/VibeThinker-3B)
- **一句话**: 3B 参数数学推理专家，基于 Qwen2 微调，适合轻量级数学任务。

### microsoft/FastContext-1.0-4B-SFT
- **作者**: microsoft | **点赞**: 213 | **下载**: 1,437
- [HF链接](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)
- **一句话**: 微软出品的长上下文推理模型，4B 参数，支持超长文本理解（SFT 版本）。

## 🎨 多模态与生成

### google/gemma-4-12B-it
- **作者**: google | **点赞**: 1,091 | **下载**: 1,590,882
- [HF链接](https://huggingface.co/google/gemma-4-12B-it)
- **一句话**: 谷歌统一多模态基座，支持文本 / 图像 / 音频任意输入输出，是 Gemma-4 系列原版。

### MiniMaxAI/MiniMax-M3
- **作者**: MiniMaxAI | **点赞**: 1,114 | **下载**: 67,836
- [HF链接](https://huggingface.co/MiniMaxAI/MiniMax-M3)
- **一句话**:  MiniMax 最新视觉语言模型，MoE 架构，擅长图文对话与多模态理解。

### google/diffusiongemma-26B-A4B-it
- **作者**: google | **点赞**: 1,007 | **下载**: 601,208
- [HF链接](https://huggingface.co/google/diffusiongemma-26B-A4B-it)
- **一句话**: Google 推出的扩散 + 语言融合多模态模型，26B 总参 / 4B 激活，适合图像描述与对话。

### moonshotai/Kimi-K2.7-Code
- **作者**: moonshotai | **点赞**: 893 | **下载**: 274,865
- [HF链接](https://huggingface.co/moonshotai/Kimi-K2.7-Code)
- **一句话**: 月之暗面的代码增强视觉模型，支持图像特征提取与代码生成，下载量近 28 万。

### nvidia/nemotron-3.5-asr-streaming-0.6b
- **作者**: nvidia | **点赞**: 556 | **下载**: 18,809
- [HF链接](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)
- **一句话**: NVIDIA 流式语音识别模型，0.6B 参数，支持低延迟实时转写。

### bosonai/higgs-audio-v3-tts-4b
- **作者**: bosonai | **点赞**: 491 | **下载**: 69,143
- [HF链接](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)
- **一句话**: Boson AI 的 4B 文本转语音模型，基于 Qwen3 多模态架构，音质自然。

### prefeitura-rio/Rio-3.5-Open-397B
- **作者**: prefeitura-rio | **点赞**: 325 | **下载**: 190,639
- [HF链接](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)
- **一句话**: 里约市政府开源的 397B MoE 视觉语言模型，适合多语言多模态任务。

### zai-org/SCAIL-2
- **作者**: zai-org | **点赞**: 230 | **下载**: 0（刚发布）
- [HF链接](https://huggingface.co/zai-org/SCAIL-2)
- **一句话**: 基于扩散模型的角色动画视频生成器，支持姿势驱动的人物动画，新模型。

### Zyphra/ZONOS2
- **作者**: Zyphra | **点赞**: 116 | **下载**: 719
- [HF链接](https://huggingface.co/Zyphra/ZONOS2)
- **一句话**: Zyphra 的 TTS 模型，Apache-2.0 许可，适合商用语音合成。

### owensong/Inflect-Nano-v1
- **作者**: owensong | **点赞**: 110 | **下载**: 0
- [HF链接](https://huggingface.co/owensong/Inflect-Nano-v1)
- **一句话**: 超小型 TTS 模型（PyTorch），主打极致轻量，适用于边缘设备。

## 🔧 专用模型

### nvidia/LocateAnything-3B
- **作者**: nvidia | **点赞**: 2,183 | **下载**: 228,669
- [HF链接](https://huggingface.co/nvidia/LocateAnything-3B)
- **一句话**: NVIDIA 的零样本目标定位模型，3B 参数，支持图像中任意物体定位，多模态任务标杆。

## 📦 微调与量化

### yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF
- **作者**: yuxinlu1 | **点赞**: 1,767 | **下载**: 268,102
- [HF链接](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)
- **一句话**: Gemma-4 的代码专用量化版，GGUF 格式，适合本地编程助手，下载超 26 万。

### HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive
- **作者**: HauhauCS | **点赞**: 1,985 | **下载**: 3,730,978
- [HF链接](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)
- **一句话**: 基于 Qwen3.6 的激进无审查 MoE 多模态模型，GGUF 格式，下载量本周第一（373 万）。

### DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF
- **作者**: DavidAU | **点赞**: 403 | **下载**: 588,753
- [HF链接](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)
- **一句话**: 极长名称的无审查微调版，融合 Claude 和 DeepSeek 思想，GGUF 量化，下载近 60 万。

### OBLITERATUS/Gemma-4-12B-OBLITERATED
- **作者**: OBLITERATUS | **点赞**: 351 | **下载**: 106,885
- [HF链接](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)
- **一句话**: Gemma-4 的“抹除”微调版，移除安全限制，safetensors + GGUF 双格式。

### Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF
- **作者**: Jackrong | **点赞**: 256 | **下载**: 148,525
- [HF链接](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF)
- **一句话**: Qwen3.6 代码变体的 GGUF 量化，27B 参数，专注编程辅助。

### nex-agi/Nex-N2-Pro
- **作者**: nex-agi | **点赞**: 332 | **下载**: 7,507
- [HF链接](https://huggingface.co/nex-agi/Nex-N2-Pro)
- **一句话**: 基于 Qwen3.5 MoE 的社区微调版，增强多模态对话能力。

### unsloth/GLM-5.2-GGUF
- **作者**: unsloth | **点赞**: 150 | **下载**: 8,392
- [HF链接](https://huggingface.co/unsloth/GLM-5.2-GGUF)
- **一句话**: GLM-5.2 的官方量化版，unsloth 出品，适合本地部署。

### unsloth/Kimi-K2.7-Code-GGUF
- **作者**: unsloth | **点赞**: 137 | **下载**: 33,667
- [HF链接](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF)
- **一句话**: Kimi-K2.7-Code 的 GGUF 量化版，便于在 llama.cpp 上运行。

### unsloth/diffusiongemma-26B-A4B-it-GGUF
- **作者**: unsloth | **点赞**: 315 | **下载**: 202,867
- [HF链接](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)
- **一句话**: DiffusionGemma 的 GGUF 量化，20 万下载，体现社区对多模态量化的强需求。

### unsloth/MiniMax-M3-GGUF
- **作者**: unsloth | **点赞**: 104 | **下载**: 24,354
- [HF链接](https://huggingface.co/unsloth/MiniMax-M3-GGUF)
- **一句话**: MiniMax-M3 的 GGUF 版，MoE + 多模态，适合 Agent 场景。

### unsloth/gemma-4-12b-it-GGUF
- **作者**: unsloth | **点赞**: 657 | **下载**: 1,150,270
- [HF链接](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)
- **一句话**: Gemma-4-12B-it 的 GGUF 量化版，下载超百万，社区最热量化变体。

### zai-org/GLM-5.2-FP8
- **作者**: zai-org | **点赞**: 96 | **下载**: 93,927
- [HF链接](https://huggingface.co/zai-org/GLM-5.2-FP8)
- **一句话**: GLM-5.2 的 FP8 量化版，兼顾精度与效率，官方提供。

### lordx64/Qwable-v1
- **作者**: lordx64 | **点赞**: 124 | **下载**: 1,865
- [HF链接](https://huggingface.co/lordx64/Qwable-v1)
- **一句话**: 基于 Qwen3.5 MoE 的社区微调视觉语言模型，侧重对话能力。

### Mia-AiLab/Qwable-3.6-27b
- **作者**: Mia-AiLab | **点赞**: 96 | **下载**: 16,105
- [HF链接](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b)
- **一句话**: Qwable 第 3.6 版，27B 参数，GGUF 格式，基于 Qwen3.6。

## 🌿 生态信号

本周生态呈现三大趋势：**1）MoE 架构成为标配**：DeepSeek-V4、GLM-5.2、Gemma-4 系列、Qwen3.6 变体等均采用 MoE，在保持性能的同时降低推理成本。**2）开源权重全面领先**：榜单前 10 中 9 个为开源模型，DeepSeek-V4-Pro 权重开放后下载超 300 万，表明社区对开放可持续的高度信任。**3）量化与微调极度活跃**：GGUF 版本达 13 个，unsolth 成为量化主力；无审查类微调（Uncensored、Aggressive、Heretic）持续涌现，反映用户对内容控制权的追求。此外，语音和视频生成模型（ASR、TTS、SCAIL-2）开始获得稳定关注，多模态生态从静态图像向动态视频 / 音频拓展。

## 🔭 值得探索

1. **deepseek-ai/DeepSeek-V4-Pro** — 本周最强基座模型，MoE 架构、超长上下文、全能对话能力，是评估当前开源 LLM 水平的最优参考。
2. **nvidia/LocateAnything-3B** — 零样本目标定位模型，3B 参数即实现高精度定位，为机器人 / 视觉检测提供轻量级解决方案，值得研究其泛化能力。
3. **zai-org/SCAIL-2** — 角色动画视频生成模型，采用 pose-driven 扩散方法，代表视频生成领域的最新进展，适合创意内容开发。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*