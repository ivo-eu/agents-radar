# Hugging Face 热门模型日报 2026-06-20

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-20 10:17 UTC

---

好的，以下是基于您提供的 2026-06-20 Hugging Face 热门模型数据生成的日报。

---

# 🤗 Hugging Face 热门模型日报 (2026-06-20)

## 今日速览

本周 Hugging Face 生态的核心事件是 **DeepSeek-V4 Pro 以压倒性优势登顶**，点赞数接近 5,000，下载量近 300 万，标志着国产大模型在开源社区的影响力达到新高度。与此同时，**代码与推理模型**成为主流，Gemma-4 系列（尤其是针对代码优化的变体）和基于 Qwen 衍生的各种 MoE 模型密集发布。最引人注目的趋势是 **GGUF 量化版模型的疯狂传播**，下载量远超原始权重模型，其中社区微调版 `Qwen3.6-35B-A3B` 的下载量已突破 380 万，反映出开源社区对本地化高效部署的极端渴求。此外，多模态模型（如图像到文本、TTS）的种类和热度也在稳步上升。

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

- **deepseek-ai/DeepSeek-V4-Pro** ([链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro))
  - 作者: deepseek-ai | 点赞: 4,973 | 下载: 2,797,050
  - 一句话说明：本周绝对的焦点模型，强大的通用对话和推理能力，凭借顶级性能成为社区新宠。

- **zai-org/GLM-5.2** ([链接](https://huggingface.co/zai-org/GLM-5.2))
  - 作者: zai-org | 点赞: 1,577 | 下载: 19,683
  - 一句话说明：智谱AI的新一代MoE模型，采用DSA架构，在对话任务上表现优异，受到学术界和爱好者热捧。

- **WeiboAI/VibeThinker-3B** ([链接](https://huggingface.co/WeiboAI/VibeThinker-3B))
  - 作者: WeiboAI | 点赞: 477 | 下载: 16,270
  - 一句话说明：一个专注于数学推理的小型 3B 模型，在社区中凭借高性价比的表现获得了关注。

- **microsoft/FastContext-1.0-4B-SFT** ([链接](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT))
  - 作者: microsoft | 点赞: 235 | 下载: 1,998
  - 一句话说明：微软推出的长上下文 4B 模型，专注于高效处理长文档，探索 Agent 能力。

- **lordx64/Qwable-v1** ([链接](https://huggingface.co/lordx64/Qwable-v1))
  - 作者: lordx64 | 点赞: 133 | 下载: 2,769
  - 一句话说明：基于 Qwen 架构的社区调优模型，主打整合不同技术特点。

- **nex-agi/Nex-N2-Pro** ([链接](https://huggingface.co/nex-agi/Nex-N2-Pro))
  - 作者: nex-agi | 点赞: 337 | 下载: 7,724
  - 一句话说明：基于 Qwen MoE 的高性能模型，结合了视觉与文本能力，是一款“全能型”选手。

- **OBLITERATUS/Gemma-4-12B-OBLITERATED** ([链接](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED))
  - 作者: OBLITERATUS | 点赞: 358 | 下载: 110,450
  - 一句话说明：基于官方 Gemma-4 的社区极限微调版本，名为“抹除”，在风格或特定任务上进行了大胆修改。

### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **MiniMaxAI/MiniMax-M3** ([链接](https://huggingface.co/MiniMaxAI/MiniMax-M3))
  - 作者: MiniMaxAI | 点赞: 1,141 | 下载: 85,771
  - 一句话说明：MiniMax 全新多模态旗舰，支持图文理解生成，代表了端到端多模态的最新水平。

- **moonshotai/Kimi-K2.7-Code** ([链接](https://huggingface.co/moonshotai/Kimi-K2.7-Code))
  - 作者: moonshotai | 点赞: 915 | 下载: 317,963
  - 一句话说明：月之暗面推出的代码理解与生成模型，支持图文交互，具备强大的代码逻辑能力。

- **google/diffusiongemma-26B-A4B-it** ([链接](https://huggingface.co/google/diffusiongemma-26B-A4B-it))
  - 作者: google | 点赞: 1,015 | 下载: 673,464
  - 一句话说明：Google 的 Diffusion Gemma 系列，用扩散模型进行图像理解与对话，是一种创新的多模态范式。

- **bosonai/higgs-audio-v3-tts-4b** ([链接](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b))
  - 作者: bosonai | 点赞: 493 | 下载: 72,225
  - 一句话说明：一个强大的 4B 参数文本转语音模型，在多模态 TTS 领域达到新高度。

- **nvidia/nemotron-3.5-asr-streaming-0.6b** ([链接](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b))
  - 作者: nvidia | 点赞: 568 | 下载: 21,426
  - 一句话说明：NVIDIA 推出的流式语音识别模型，主打低延迟、高精度，适合实时场景。

- **zai-org/SCAIL-2** ([链接](https://huggingface.co/zai-org/SCAIL-2))
  - 作者: zai-org | 点赞: 235 | 下载: 0
  - 一句话说明：智谱 AI 推出的图像到视频模型，专注于姿态驱动的角色动画生成。

- **Zyphra/ZONOS2** ([链接](https://huggingface.co/Zyphra/ZONOS2))
  - 作者: Zyphra | 点赞: 116 | 下载: 763
  - 一句话说明：一款开源的、高质量的自然语音TTS模型，采用 Apache 2.0 许可证，适合商业应用。

- **owensong/Inflect-Nano-v1** ([链接](https://huggingface.co/owensong/Inflect-Nano-v1))
  - 作者: owensong | 点赞: 129 | 下载: 0
  - 一句话说明：超小型的 TTS 模型，主打极致轻量化，适合边缘设备部署。

### 🔧 专用模型（代码、数学、医疗、嵌入）

- **nvidia/LocateAnything-3B** ([链接](https://huggingface.co/nvidia/LocateAnything-3B))
  - 作者: nvidia | 点赞: 2,203 | 下载: 235,606
  - 一句话说明：本周热度第二高的模型！NVIDIA 推出的视觉定位模型，能准确检测和定位图像中的任意目标。

- **CohereLabs/North-Mini-Code-1.0** ([链接](https://huggingface.co/CohereLabs/North-Mini-Code-1.0))
  - 作者: CohereLabs | 点赞: 461 | 下载: 18,783
  - 一句话说明：Cohere 推出的代码专用 MoE 模型，旨在为开发者提供高效、精准的代码生成体验。

### 📦 微调与量化（社区微调、GGUF、AWQ）

- **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF** ([链接](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF))
  - 作者: yuxinlu1 | 点赞: 1,887 | 下载: 312,332
  - 一句话说明：社区对 Gemma-4 代码版的高度优化量化版，专为本地代码生成场景打造。

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** ([链接](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive))
  - 作者: HauhauCS | 点赞: 2,013 | 下载: 3,812,636
  - 一句话说明：下载量冠军！Qwen3.6 MoE 的社区激进微调 + GGUF 量化版，主打“无审查”和个性风格。

- **unsloth/GLM-5.2-GGUF** ([链接](https://huggingface.co/unsloth/GLM-5.2-GGUF))
  - 作者: unsloth | 点赞: 193 | 下载: 22,586
  - 一句话说明：知名量化团队 Unsloth 出品的 GLM-5.2 GGUF 版，确保了最佳的速度与兼容性。

- **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF** ([链接](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF))
  - 作者: yuxinlu1 | 点赞: 126 | 下载: 6,307
  - 一句话说明：针对 Agent（终端调用）场景优化的 Gemma-4 量化版，适合自动化编程任务。

- **unsloth/Kimi-K2.7-Code-GGUF** ([链接](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF))
  - 作者: unsloth | 点赞: 143 | 下载: 37,260
  - 一句话说明：Kimi K2.7 代码模型的官方优化版，降低部署门槛，扩大用户覆盖。

- **zai-org/GLM-5.2-FP8** ([链接](https://huggingface.co/zai-org/GLM-5.2-FP8))
  - 作者: zai-org | 点赞: 108 | 下载: 138,174
  - 一句话说明：GLM-5.2 的高精度 FP8 量化版，在精度与速度之间取得了平衡。

- **Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF** ([链接](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-MTP-GGUF))
  - 作者: Jackrong | 点赞: 263 | 下载: 168,502
  - 一句话说明：Qwen3.6 衍生模型，专注代码生成，并使用 Multi-Token Prediction 技术进行优化。

- **unsloth/diffusiongemma-26B-A4B-it-GGUF** ([链接](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF))
  - 作者: unsloth | 点赞: 318 | 下载: 216,396
  - 一句话说明：Diffusion Gemma 的 GGUF 版，让用户在消费级硬件上体验创新的扩散多模态模型。

- **DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF** ([链接](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF))
  - 作者: DavidAU | 点赞: 406 | 下载: 587,521
  - 一句话说明：名字超长的“缝合怪”社区模型，融合了多种顶尖模型的技术，追求极致的代码与推理能力。

- **bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF** ([链接](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF))
  - 作者: bytkim | 点赞: 89 | 下载: 20,465
  - 一句话说明：基于 Qwen3.6 的微调量化版，结合了 Multi-Token Prediction 技术并进行精细调优。

- **Mia-AiLab/Qwable-3.6-27b** ([链接](https://huggingface.co/Mia-AiLab/Qwable-3.6-27b))
  - 作者: Mia-AiLab | 点赞: 108 | 下载: 17,311
  - 一句话说明：一个基于 Qwen 3.6 的社区版 GGUF 模型，方便用户快速部署。

## 生态信号

- **模型家族“四雄争霸”**：本周生态围绕 **DeepSeek、Qwen、Gemma、GLM** 四大模型家族展开。**Qwen 3.6** 系列是社区微调的最大热门，衍生出大量专用版和量化版；**DeepSeek-V4 Pro** 则以原生强大的性能被公认为“基准”。**MoE 架构已成主流**，榜单中超过一半的大模型都采用了混合专家系统。
- **开源权重生态繁荣**：可以看到，大量高质量的原生权重（如 DeepSeek、GLM）和公司官方版本正在持续输出，而社区基于这些权重进行的 **二次创作、微调、合并** 非常活跃，构成了完整的价值链。
- **GGUF 量化成为“硬通货”**：GGUF 格式模型的数量和下载量是压倒性的。即便如 DeepSeek-V4 Pro 这样的旗舰模型，也急需 GGUF 版本（榜单中虽未显示，但下载量暗示了其量化版可能极其庞大）。社区对于 **“到手即用”的本地本地部署体验**的需求正在塑造整个模型发布的标准流程。

## 值得探索

1.  **deepseek-ai/DeepSeek-V4-Pro** ([链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)): **必试！** 作为本周热度（点赞数）最高的模型，它是测试当前开源大模型性能上限的标杆。如果你有高端硬件，这是理解最新技术水平的首选。
2.  **nvidia/LocateAnything-3B** ([链接](https://huggingface.co/nvidia/LocateAnything-3B)): **计算机视觉爱好者推荐。** 该模型精准地切入“视觉定位”这个细分领域，获得极高热度，表明一个非通用、但实用的视觉模型也能取得巨大成功，值得深入研究其架构。
3.  **Zyphra/ZONOS2** ([链接](https://huggingface.co/Zyphra/ZONOS2)): **文本转语音应用推荐。** 这是榜单中为数不多的高质量、商业友好的开源 TTS 模型（Apache-2.0 许可证）。如果你的项目需要自然逼真的语音合成，这个模型可能提供很好的基础，且社区刚发布，生态支持会逐渐完善。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*