# Hugging Face 热门模型日报 2026-06-18

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-18 12:31 UTC

---

好的，作为 AI 模型生态分析师，这是为您整理的 2026-06-18 Hugging Face 热门模型日报。

---

### 📰 《Hugging Face 热门模型日报》— 2026-06-18

#### **今日速览**

今日 HF 趋势榜呈现三大特征。首先，**DeepSeek-V4-Pro** 以绝对优势登顶，证明了顶级基础模型在社区中无可撼动的号召力。其次，**多模态模型**继续占据主导地位，几乎所有热门模型都具备图像或视频理解能力，“任何模态”输入输出成为新标配。最后，**社区微调与量化**（尤其是 GGUF 格式）异常活跃，围绕 Gemma-4、Kimi-K2.7 等基座模型的二次创作百花齐放，表明用户对模型实用化和本地部署的需求持续高涨。

#### **热门模型**

##### 🧠 语言模型（LLM、对话模型、指令微调）

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** (作者: deepseek-ai, 点赞: 4,935, 下载: 2,948,726)
  - **一句话说明**：DeepSeek 最新旗舰模型，以极佳的对话和推理能力在社区中引起巨大反响，是今日榜单的人气与下载量双料冠军。

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** (作者: zai-org, 点赞: 1,190, 下载: 4,307)
  - **一句话说明**：智谱 AI 的下一代 MoE 大模型，凭借出色的中文对话能力和创新的 DSA 架构，迅速吸引了大量关注。

- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** (作者: CohereLabs, 点赞: 434, 下载: 15,285)
  - **一句话说明**：Cohere 推出的高效代码生成模型，结合了最新的 MoE 架构，旨在以更小的参数量实现优秀的代码理解与生成能力。

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** (作者: microsoft, 点赞: 191, 下载: 957)
  - **一句话说明**：微软出品的专注长上下文处理的 4B 小模型，采用“Explorer SubAgent”技术，旨在高效处理超长序列。

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** (作者: nex-agi, 点赞: 321, 下载: 6,640) / **[nex-agi/Nex-N2-mini](https://huggingface.co/nex-agi/Nex-N2-mini)** (作者: nex-agi, 点赞: 240, 下载: 11,660)
  - **一句话说明**：nex-agi 推出的基于 Qwen3.5 MoE 架构的系列模型，兼顾了多模态理解与对话能力，体现了“以小博大”的参数效率思路。

##### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** (作者: nvidia, 点赞: 2,145, 下载: 183,093)
  - **一句话说明**：NVIDIA 推出的通用定位模型，能根据文本指令精确定位图像中的任何物体，是榜单中技术新颖度极高的亮点。

- **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)** (作者: google, 点赞: 992, 下载: 527,080)
  - **一句话说明**：Google 的旗舰多模态生成模型，将 Gemma 语言模型与扩散模型结合，实现了高效的图文理解和生成，下载量巨大。

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** (作者: MiniMaxAI, 点赞: 1,079, 下载: 56,162)
  - **一句话说明**：MiniMax 的多模态 MoE 大模型，显示出强大的 Agent 能力，是国产模型在多模态领域的又一力作。

- **[prefeitura-rio/Rio-3.5-Open-397B](https://huggingface.co/prefeitura-rio/Rio-3.5-Open-397B)** (作者: prefeitura-rio, 点赞: 322, 下载: 190,501)
  - **一句话说明**：巴西里约市政府开源的超大 MoE 多模态模型，拥有 397B 参数，成为开源社区中规模最大的模型之一，备受关注。

- **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)** (作者: google, 点赞: 1,075, 下载: 1,309,625)
  - **一句话说明**：Google 的多模态旗舰 Gemma-4，“any-to-any”任务标签表明其能处理文本、图像等多种输入并生成多种输出，是新一代通用模型。

- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** (作者: bosonai, 点赞: 484, 下载: 57,380)
  - **一句话说明**：Higgs Audio 最新版语音合成模型，基于 Qwen3 架构，实现了高表现力的 TTS 能力，是音频生成领域的热门选择。

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** (作者: nvidia, 点赞: 528, 下载: 13,033)
  - **一句话说明**：NVIDIA 推出的轻量级流式语音识别模型，专为低延迟实时语音处理场景设计，兼顾了效率与精度。

- **[zai-org/SCAIL-2](https://huggingface.co/zai-org/SCAIL-2)** (作者: zai-org, 点赞: 225, 下载: 0)
  - **一句话说明**：智谱推出的图像到视频的生成模型，专注于角色动画和姿态驱动，是当日榜单中唯一的视频生成模型，极具探索价值。

##### 🔧 专用模型（代码、数学、医疗、嵌入）

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** (作者: yuxinlu1, 点赞: 1,599, 下载: 211,424)
  - **一句话说明**：基于 Gemma-4 微调的精良代码模型，下载量极高，反映了社区对高性能本地化代码助手的需求旺盛。

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)** (作者: moonshotai, 点赞: 862, 下载: 229,156)
  - **一句话说明**：月之暗面推出的代码专用多模态模型，采用压缩张量技术，在保持代码能力的同时实现了高效的性能。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** (作者: WeiboAI, 点赞: 365, 下载: 6,589)
  - **一句话说明**：微博 AI 推出的 3B 级别数学推理模型，专攻数学问题，小而美的定位吸引了特定领域的用户。

##### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** (作者: HauhauCS, 点赞: 1,953, 下载: 3,420,052)
  - **一句话说明**：基于 Qwen3.6 的激进社区微调版本，主打“无审查”和强烈的个性化风格，下载量惊人，是社区“创意”微调的代表。

- **[DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF](https://huggingface.co/DavidAU/Qwen3.6-40B-Claude-4.6-Opus-Deckard-Heretic-Uncensored-Thinking-NEO-CODE-Di-IMatrix-MAX-GGUF)** (作者: DavidAU, 点赞: 388, 下载: 529,069)
  - **一句话说明**：另一个极端社区微调案例，名称揭示了其混合了多个模型的特质，追求极致性能，体现了社区对模型进行“混合与蒸馏”的探索。

- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** (作者: unsloth, 点赞: 646, 下载: 918,431)
  - **一句话说明**：Unsloth 团队对 Gemma-4 的官方 GGUF 量化版本，极大降低了模型部署门槛，是推动 Gemma-4 流行开来的功臣。

- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** (作者: OBLITERATUS, 点赞: 344, 下载: 96,805)
  - **一句话说明**：另一款社区对 Gemma-4 的独特微调版本，暗示了更极致的性能或风格调整，反映了社区对基座模型的深刻改造需求。

- **其他量化版本：** [unsloth/diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF) (下载: 164k), [unsloth/Kimi-K2.7-Code-GGUF](https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF) (下载: 29k), [zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8) 等，显示出社区量化活动已覆盖了几乎所有顶级新模型。

#### **生态信号**

- **多模态与 MoE 成为绝对主流**：榜单上几乎所有模型都集成了多模态能力（`image-text-to-text` 标签），且大量采用 MoE（混合专家）架构（如 `glm_moe_dsa`, `qwen3_5_moe`, `cohere2_moe`）。这表明行业共识正从纯文本 LLM 转向能理解世界、更高效的通用基础模型。

- **开源生态火爆，社区微调成关键驱动力**：DeepSeek、Google、智谱等公司持续发布顶级开源权重，而 NVIDIA、MiniMax 等也加入开放阵营。同时，以 `HauhauCS`、`DavidAU`、`OBLITERATUS` 为代表的社区创作者异常活跃，其微调版本下载量甚至超过部分官方模型，反映了用户在性能、风格和部署便利性上的多样化需求。

- **量化生态加速模型普及**：`GGUF` 格式的模型占总下载量的半壁江山。Unsloth 团队再次扮演了重要角色，几乎为所有热门基础模型提供了高质量的 GGUF 版本，极大地促进了模型在本地和端侧设备上的部署与使用。

#### **值得探索**

1.  **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**：作为今日热议中心，探索其对话、推理和多语言能力的“天花板”在哪里，是理解当前最强开源模型性能的关键。

2.  **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**：这款“any-to-any”模型是 Google 对下一代多模态交互形式的尝试。研究其如何统一不同类型任务的输入输出，对理解未来 AI Agent 形态有重要意义。

3.  **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**：这是一个非常有趣的专用视觉模型。测试其在真实场景（如机器人操控、图像编辑）中对“物体定位”的精准度和泛化能力，会发现视觉AI落地的新路径。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*