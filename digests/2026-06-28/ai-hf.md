# Hugging Face 热门模型日报 2026-06-28

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-28 10:09 UTC

---

好的，以下是 2026-06-28 的《Hugging Face 热门模型日报》。

---

## Hugging Face 热门模型日报 | 2026-06-28

### 📰 今日速览

本周 Hugging Face 生态呈现出“巨头林立、百花齐放”的态势。以 **zai-org** 的 **GLM-5.2** 为首的 MoE 模型家族（包括未量化、GGUF 和 NVIDIA NVFP4 版本）表现抢眼，占据了榜单头部。与此同时，**Qwen** 架构的衍生模型数量最多，从基础大模型到社区微调版本（如 Uncensored 变体）均有超高下载量。多模态趋势依旧强劲，视频理解模型 **nvidia/LocateAnything-3B** 和语音模型 **nvidia/nemotron-3.5-asr** 首次进入视线。社区量化活动频繁，**HauhauCS** 的“Uncensored”系列和 **yuxinlu1** 的“Fable5”系列成为本周微调最大赢家。

### 🏆 热门模型

#### 🧠 语言模型（LLM、对话模型、指令微调）

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** (zai-org, 👍2,716, 📥118,651) — 新一代混合专家模型（MoE），凭借强大的对话和生成能力，成为本周最受关注的通用 LLM。
- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)** (empero-ai, 👍506, 📥52,492) — 基于 Qwen3.5 的融合微调模型，主打推理与多模态理解，是“少量样本达到顶尖推理”口号的实践者。
- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)** (Qwen, 👍369, 📥23,697) — Qwen 官方出品的面向 Agent 世界模型的大规模 MoE 模型，专注于智能体环境模拟与决策。
- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)** (yuxinlu1, 👍747, 📥225,822) — 基于 Gemma 4 的 Agent 专用变体，结合“Fable5”微调范式，支持复杂终端交互与任务规划。
- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)** (deepreinforce-ai, 👍130, 📥1,116) — 当前榜单最大参数量的模型，Qwen3.5 MoE 架构的极致缩放，适合需要极致广度与复杂推理的场景。
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** (HauhauCS, 👍2,285, 📥3,248,724) — 社区最受欢迎的无审查版本，本日下载量冠军，基于 Qwen3.6 的激进风格微调。
- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** (WeiboAI, 👍742, 📥59,337) — 3B 参数的小钢炮，专注数学推理，是低成本部署和学术验证的理想选择。
- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)** (LiquidAI, 👍132, 📥12,384) — 230M 的轻量级模型，属于 Liquid 系列的最新迭代，主打边缘设备上的高效文本生成。
- **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)** (deepseek-ai, 👍153, 📥373) — DeepSeek V4 的专业版，引入“DSpark”架构优化，代表闭源大模型的开源对抗力量。
- **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/chunjiang-Intelligence/DeepSeek-v4-Fable)** (Chunjiang-Intelligence, 👍118, 📥1,409) — 基于 DeepSeek v4 的网络安全专用微调模型，聚焦攻防场景和日志分析。

#### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)** (baidu, 👍1,164, 📥295,064) — 百度的全能 OCR 模型，支持不限字符的高精度识别，下载量极大，是开源 OCR 的首选。
- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** (nvidia, 👍2,418, 📥646,451) — NVIDIA 的目标定位基础模型，能从图像中精准定位任意物体，颠覆了传统视觉定位范式。
- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)** (nvidia, 👍725, 📥67,419) — 600M 参数的流式语音识别模型，支持低延迟的实时转写，是音频领域的明星。
- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)** (krea, 👍329, 📥27,631) — Krea 最新图像生成模型的 Turbo 版本，推理速度极快，基于其基础模型 Krea-2-Raw 的蒸馏优化。
- **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)** (krea, 👍221, 📥22,622) — Krea 第二代图像生成的原始基础模型，提供最高的画质控制上限。
- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** (MiniMaxAI, 👍1,255, 📥188,314) — MiniMax 第三代多模态大模型，统一理解与生成，支持图文混合输入，是目前最强视觉语言模型之一。
- **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)** (Comfy-Org, 👍166, 📥10) — ComfyUI 官方发布的 Krea-2 工作流插件，方便用户在 ComfyUI 中调用 Krea 模型。

#### 🔧 专用模型（代码、数学、医疗、嵌入）

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** (yuxinlu1, 👍2,439, 📥549,926) — 基于 Gemma 4 的代码生成模型，采用“Fable5”和“Composer”微调策略，在编码任务上表现突出。
- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** (microsoft, 👍368, 📥6,779) — 微软出品的长上下文微调模型，擅长处理超长文本（如整本小说或论文）和复杂信息提取。
- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)** (WeiboAI, 👍742, 📥59,337) — 作为数学专用模型，在此重复出现，强调其在数学推理领域的广泛应用。

#### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)** (empero-ai, 👍699, 📥831,529) — Qwythos 模型的 GGUF 量化版本，极大降低了推理门槛，是本地化部署的热门选择。
- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)** (yuxinlu1, 👍2,439, 📥549,926) — Gemma-4 Coder 的 GGUF 版本，量化后仍保持惊人的编码能力，是编程爱好者的首选量化模型。
- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** (unsloth, 👍429, 📥146,023) — unsloth 出品的高效 GLM-5.2 量化版，推理速度与易用性俱佳。
- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)** (nvidia, 👍143, 📥45,762) — NVIDIA 利用 ModelOpt 对 GLM-5.2 进行 4-bit FP4 量化，在 NVIDIA 硬件上实现极致性能。
- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)** (deepreinforce-ai, 👍350, 📥79,630) — Ornith-35B 的 GGUF 版本，兼顾性能与体积，适合单卡运行。
- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)** (deepreinforce-ai, 👍236, 📥36,846) — Ornith-9B 的精简量化版，10B 以下模型中性价比极高的选择。
- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)** (HauhauCS, 👍2,285, 📥3,248,724) — 社区微调的杰出代表，GGE格式的“Uncensored”系列，因其过滤难度低、风格激进而广受追捧。
- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)** (nvidia, 👍367, 📥5,235,413) — NVIDIA 为 Qwen3.6 定制的 FP4 量化版，是目前下载量最高的量化模型之一，展示了硬件厂商对开源生态的深度支持。
- **[huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated)** (huihui-ai, 👍138, 📥7,210) — “abliterated”系列的最新作品，通过移除安全对齐层来获得无约束的生成能力。
- **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)** (HauhauCS, 👍98, 📥40,820) — 采用QAT（量化感知训练）的Gemma 4无审查版，在保持性能的同时降低了部署延迟。
- **[unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)** (unsloth, 👍84, 📥79,503) — unsloth 与 Qwen 合作带来的 AgentWorld 模型的量化版本，使 Agent 世界模型更易部署。

### 🌐 生态信号

1.  **模型家族“双子星”对决：GLM vs. Qwen**。本周榜单清晰显示出两大模型生态的激烈竞争。GLM-5.2 家族凭借原创 MoE 结构占据高质量点赞榜首，而 Qwen 家族（3.5/3.6）则以海量的社区微调和量化版本（如 HauhauCS 的 Uncensored 系列）统治下载量。**“Uncensored”社区微调**成为最火爆的细分品类，社区对无审查模型的渴望远超预期。
2.  **开源权重持续主导**。所有热门模型均为开源权重，未见纯 API 模型上榜。这标志着开源模型在性能和质量上已足以让社区充满创作热情，闭源模型更多通过论文（如 arxiv:2606.19348）或特定硬件支持（如 NVFP4）间接参与生态。
3.  **量化与硬件绑定加深**。GGUF 仍是本地部署的绝对主流，但 NVIDIA 的 **NVFP4 (ModelOpt)** 异军突起，凭借大规模下载量（Qwen3.6 版超 500 万下载），表明硬件厂商正在通过量化工具链深度绑定自身生态，用户对“硬件优化+模型量化”的组合接受度极高。
4.  **多模态从“识别”走向“定位”**。本周多模态热点从传统的图文理解（MiniMax-M3）转向了更精准的视觉定位（nvidia/LocateAnything-3B）和实时语音（nemotron-3.5-asr），体现模型能力正从“理解是什么”向“定位在哪里”“听懂说什么”升级。

### 🔭 值得探索

1.  **[GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** 及其量化版：作为本周最受瞩目的原创架构，其 MoE 设计在效率和能力平衡上值得深入研究。推荐理由：代表了当前非 Qwen/Llama 架构的最高水平，有很高的学术与工程价值。
2.  **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**：开创了“零样本+开放词汇+任意目标”的定位新范式，是计算机视觉领域的重大突破。推荐理由：3B 参数量便达到如此高的泛化性，极有可能成为未来 AI 视觉 Agent 的标准基础组件。
3.  **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**：在 3B 级别上实现了令人惊叹的数学推理能力。推荐理由：对于资源受限但需要高精度数学推理的应用场景，它可能是最具性价比的选择。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*