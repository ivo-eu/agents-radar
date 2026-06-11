# Hugging Face 热门模型日报 2026-06-11

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-11 03:32 UTC

---

好的，作为AI模型生态分析师，以下是根据您提供的2026年6月11日数据生成的《Hugging Face热门模型日报》。

---

### **Hugging Face 热门模型日报 | 2026年6月11日**

#### **今日速览**

本周Hugging Face生态呈现出“旗舰模型齐发，多模态与量化并进”的态势。DeepSeek-V4-Pro以绝对优势的点赞和下载量登顶，成为社区焦点。Google的Gemma 4家族持续扩张，衍生出大量量化、微调版本，生态影响力巨大。多模态方面，图像生成、视频生成乃至实时音频生成模型百花齐放，而NVIDIA、Cohere等巨头则展示了其在大规模MoE和专用模型上的深厚积累。社区微调与“去审查”活动依旧活跃，反映了开源社区对模型定制化的强烈需求。

#### **热门模型**

##### 🧠 语言模型（LLM、对话模型、指令微调）

- **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)** — 作者: deepseek-ai | 点赞: 4,762 | 下载: 4,061,006
  一句话说明：DeepSeek的第四代旗舰模型，兼具强大的文本生成与对话能力，凭借其顶级的性能和巨大的参数量，成为本周当之无愧的社区顶流。
- **[nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16)** — 作者: nvidia | 点赞: 189 | 下载: 59,066
  一句话说明：NVIDIA推出的5500亿参数（活跃55B）超大规模MoE模型，代表了当前闭源级开源模型在参数量级上的新高度。
- **[CohereLabs/North-Mini-Code-1.0](https://huggingface.co/CohereLabs/North-Mini-Code-1.0)** — 作者: CohereLabs | 点赞: 263 | 下载: 1,859
  一句话说明：Cohere发布的代码专用小型MoE模型，专注于编程任务，性能与效率的平衡使其在开发社区中受到关注。
- **[LiquidAI/LFM2.5-8B-A1B](https://huggingface.co/LiquidAI/LiquidAI/LFM2.5-8B-A1B)** — 作者: LiquidAI | 点赞: 584 | 下载: 142,134
  一句话说明：Liquid AI推出的第二代基础模型，8B参数仅激活1B，是高效MoE架构的代表作，吸引了大量关注效率的研究者。
- **[JetBrains/Mellum2-12B-A2.5B-Thinking](https://huggingface.co/JetBrains/Mellum2-12B-A2.5B-Thinking)** — 作者: JetBrains | 点赞: 281 | 下载: 18,273
  一句话说明：JetBrains推出的具备“思考”能力的推理增强模型，旨在提升代码和逻辑任务的准确性，为开发者工具链注入AI新能力。
- **[sapientinc/HRM-Text-1B](https://huggingface.co/sapientinc/HRM-Text-1B)** — 作者: sapientinc | 点赞: 741 | 下载: 134,752
  一句话说明：一款专注于人力资源管理（HRM）领域的1B参数专用模型，其垂直领域的精准定位和高下载量反映了行业模型的强劲需求。

##### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[ideogram-ai/ideogram-4-fp8](https://huggingface.co/ideogram-ai/ideogram-4-fp8)** — 作者: ideogram-ai | 点赞: 474 | 下载: 7,170
  一句话说明：Ideogram最新版图像生成模型的FP8量化版，在保持生成质量的同时显著降低了部署门槛。
- **[stepfun-ai/Step-3.7-Flash](https://huggingface.co/stepfun-ai/Step-3.7-Flash)** — 作者: stepfun-ai | 点赞: 363 | 下载: 50,187
  一句话说明：阶跃星辰推出的视觉语言模型，主打快速高效的图文理解与生成，是中文多模态开源领域的重要竞争者。
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — 作者: nvidia | 点赞: 1,809 | 下载: 131,794
  一句话说明：NVIDIA推出的目标检测与定位基础模型，能通过文本或视觉提示精准定位图像中的任意物体，实用性强，社区反响热烈。
- **[bosonai/higgs-audio-v3-tts-4b](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)** — 作者: bosonai | 点赞: 326 | 下载: 19,948
  一句话说明：博上科技推出的下一代4B参数TTS模型，生成语音自然度和表现力出色，是当前开源TTS领域的有力代表。
- **[google/magenta-realtime-2](https://huggingface.co/google/magenta-realtime-2)** — 作者: google | 点赞: 174 | 下载: 19,806
  一句话说明：Google Magenta团队推出的第二代实时音乐生成模型，专注于text-to-audio任务，为AI音乐创作带来更低延迟和更高品质。
- **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)** — 作者: ByteDance | 点赞: 213 | 下载: 305
  一句话说明：字节跳动发布的文生视频模型，将图像和文本转化为动态视频内容，代表了大厂在视频生成赛道的最新开源尝试。
- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nvidia/nemotron-3.5-asr-streaming-0.6b)** — 作者: nvidia | 点赞: 350 | 下载: 4,965
  一句话说明：NVIDIA推出的流式语音识别（ASR）模型，主打低延迟和缓存感知，适合实时语音交互场景。

##### 🔧 专用模型（代码、数学、医疗、嵌入）

- **[nex-agi/Nex-N2-Pro](https://huggingface.co/nex-agi/Nex-N2-Pro)** — 作者: nex-agi | 点赞: 181 | 下载: 1,185
  一句话说明：基于Qwen3.5-MoE架构的进阶版专业模型，可能针对特定领域进行了优化，虽然下载量不大但热度看涨。
- **[sapientinc/HRM-Text-1B](https://huggingface.co/sapientinc/HRM-Text-1B)** — 作者: sapientinc | 点赞: 741 | 下载: 134,752
  一句话说明：（同上）作为HRM领域的专用模型，其成功上榜印证了行业大模型正在成为独立且重要的赛道。

##### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** — 作者: HauhauCS | 点赞: 1,637 | 下载: 3,057,541
  一句话说明：社区对Qwen3.6进行激进“去审查”微调并量化为GGUF的版本，因其强大的定制化程度和惊人的下载量，成为社区最活跃的微调模型之一。
- **[unsloth/gemma-4-12b-it-GGUF](https://huggingface.co/unsloth/gemma-4-12b-it-GGUF)** — 作者: unsloth | 点赞: 551 | 下载: 711,706
  一句话说明：Unsloth对Google Gemma 4指令版的高效GGUF量化版本，极大降低了其在消费级硬件上的部署门槛，下载量巨大。
- **[huihui-ai/Huihui-gemma-4-12B-it-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-it-abliterated)** — 作者: huihui-ai | 点赞: 136 | 下载: 6,400
  一句话说明：社区对Gemma 4指令版进行“ablation”的版本，移除特定能力或限制，体现了社区对模型进行深度定制的技术实力。
- **[OBLITERATUS/Gemma-4-12B-OBLITERATED](https://huggingface.co/OBLITERATUS/Gemma-4-12B-OBLITERATED)** — 作者: OBLITERATUS | 点赞: 215 | 下载: 14,838
  一句话说明：另一个以“Obliterate”命名的Gemma 4社区微调版，与上述模型类似，代表了对基础模型进行“解禁”或特化处理的一种流行趋势。

#### **生态信号**

1.  **MoE架构持续主导**：本周几乎所有大参数模型，如Nemotron-3 Ultra、Cohere North、Liquid LFM、Nex-N2和Qwen3.6均采用了MoE架构，MoE已成为追求高性能与效率平衡的行业标准。
2.  **Google Gemma 4生态爆发**：Google发布的Gemma 4模型家族迅速成为社区微调与量化的“沃土”。大量基于Gemma 4的GGUF版本（由Unsloth主导）和“去审查”社区版涌现，这表明官方基础模型的质量和许可友好度是催生繁荣生态的关键。
3.  **“去审查”与社区定制化活跃**：带着“Uncensored”、“Abliterated”、“Obliterated”标签的模型频繁上榜，下载量惊人。这反映了开源社区对模型自主控制的强烈需求，旨在突破原始模型的限制，探索更具个性或开放的AI应用。
4.  **多模态“all-in-one”成为新焦点**：以Google Gemma 4和Step-3.7为代表的“image-text-to-text”模型，以及各类TTS、视频生成模型，表明“any-to-any”的跨模态能力不再是少数模型的专利，正向更多开源模型普及。

#### **值得探索**

1.  **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**：作为本周热度与下载量双料冠军，它代表了当前开源语言模型的最高水平之一。无论是进行基准测试、能力评估，还是作为构建上层应用的基座，都值得深入研究。
2.  **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**：这是一个非常实用的多模态模型。它跳出了传统的“看图说话”范式，直接解决“图像中有什么，在哪里”的经典视觉问题。其3B的参数量也使其具备了在边缘设备上部署的潜力。
3.  **[ByteDance/Bernini-R](https://huggingface.co/ByteDance/Bernini-R)**：视频生成是当前AI领域竞争最激烈的赛道之一。作为来自字节跳动的开源尝试，它有望复现或超越某些闭源视频模型的惊艳效果，参与和测试该模型是了解视频生成最新进展的绝佳窗口。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*