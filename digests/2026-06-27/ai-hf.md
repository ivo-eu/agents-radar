# Hugging Face 热门模型日报 2026-06-27

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-27 09:15 UTC

---

# Hugging Face 热门模型日报（2026-06-27）

## 今日速览

本周 Hugging Face 榜单由 **GLM-5.2**（周赞 2,624）和 **Gemma-4-12B-Coder** 系列（周赞 2,412）领跑，反映出大型 MoE 架构与代码专用模型的热度。多模态模型成为绝对主力：超过半数上榜模型支持图像‑文本交互，其中 **MiniMax-M3**、**Qwen3.6-35B-A3B-Uncensored** 和 **NVIDIA LocateAnything-3B** 均获高赞。量化部署活动空前活跃，**NVFP4** 与 **GGUF** 格式持续扩散，NVIDIA 发布的 Qwen3.6 量化版单周下载量突破 500 万。此外，“无审查”“去对齐”（abliterated）等社区微调模型大量涌现，个性化定制需求强劲。

## 热门模型

### 🧠 语言模型（LLM、对话、指令微调）

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  作者: zai-org | 点赞: 2,624 | 下载: 98,994  
  开源 MoE 大语言模型，融合 DSA 稀疏注意力技术，本周点赞数最高。

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
  作者: yuxinlu1 | 点赞: 699 | 下载: 206,828  
  面向 agent 任务微调的 Gemma-4-12B 量化版，强调终端与编码能力。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
  作者: WeiboAI | 点赞: 736 | 下载: 57,521  
  基于 Qwen2 的 3B 数学推理模型，轻量级专用 LLM，适合数学场景。

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
  作者: unsloth | 点赞: 420 | 下载: 125,230  
  GLM-5.2 的 GGUF 量化版，由 Unsloth 社区提供，便于本地部署。

- **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
  作者: microsoft | 点赞: 359 | 下载: 6,447  
  基于 Qwen3 的 4B 指令微调模型，专为 Explorer SubAgent 场景优化。

- **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)**  
  作者: nvidia | 点赞: 363 | 下载: 5,022,254  
  Qwen3.6-35B 的 NVIDIA FP4 量化版，以极高下载量成为最热门推理格式。

- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)**  
  作者: LiquidAI | 点赞: 120 | 下载: 9,791  
  Liquid AI 推出的 230M 参数高效语言模型，面向边缘部署。

- **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)**  
  作者: Chunjiang-Intelligence | 点赞: 108 | 下载: 1,328  
  DeepSeek v4 的微调版本，专注于网络安全领域。

- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)**  
  作者: nvidia | 点赞: 98 | 下载: 6,464  
  GLM-5.2 的 NVIDIA FP4 量化版，利用 ModelOpt 实现高效推理。

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)**  
  作者: deepreinforce-ai | 点赞: 258 | 下载: 20,266  
  Ornith-1.0-35B 的 GGUF 量化版，采用 MIT 协议，兼容多种推理框架。

- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)**  
  作者: deepreinforce-ai | 点赞: 178 | 下载: 11,034  
  Ornith-1.0-9B 的 GGUF 量化版，轻量级多模态备选（纯文本质询版）。

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  作者: baidu | 点赞: 1,077 | 下载: 212,760  
  百度推出的通用 OCR 模型，支持不限场景的文字识别，下载量极高。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
  作者: empero-ai | 点赞: 622 | 下载: 712,627  
  基于 Qwen3.5 的 9B 视觉‑语言模型 GGUF 版，融合 Claude‑Mythos 风格，推理类应用热门。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
  作者: empero-ai | 点赞: 458 | 下载: 30,298  
  Qwythos 的原始 Safetensors 版本，支持图像‑文本输入输出。

- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)**  
  作者: Qwen | 点赞: 333 | 下载: 18,872  
  通义千问 AgentWorld 系列，35B MoE 架构，面向智能体应用的多模态模型。

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  
  作者: krea | 点赞: 294 | 下载: 17,445  
  Krea-2 的 Turbo 变体，高性能文生图模型，基于 Diffusers 框架。

- **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)**  
  作者: krea | 点赞: 207 | 下载: 17,748  
  Krea-2 的基础版本，提供原始文生图能力，作为 Turbo 的底座模型。

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  作者: HauhauCS | 点赞: 2,268 | 下载: 3,331,475  
  去除审查的 Qwen3.6-35B MoE 多模态模型，采用激进风格，社区极受欢迎。

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
  作者: nvidia | 点赞: 2,391 | 下载: 570,466  
  NVIDIA 推出的 3B 视觉定位模型，可在图像中定位任意目标，技术突破性高。

- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)**  
  作者: deepreinforce-ai | 点赞: 133 | 下载: 7,571  
  Ornith 系列 35B 全精度版，支持图像‑文本多模态推理。

- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)**  
  作者: deepreinforce-ai | 点赞: 130 | 下载: 1,501  
  Ornith-1.0-9B 全精度版，轻量多模态选项。

- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)**  
  作者: deepreinforce-ai | 点赞: 111 | 下载: 463  
  超大规模 MoE 多模态模型（397B），参数规模最大。

- **[huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated)**  
  作者: huihui-ai | 点赞: 136 | 下载: 6,250  
  Gemma-4-12B 的“去对齐”微调版，移除安全限制，保留多模态编码能力。

- **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**  
  作者: MiniMaxAI | 点赞: 1,248 | 下载: 182,714  
  MiniMax 第三代多模态大模型，支持图像‑文本深度交互，性能强劲。

- **[datalab-to/lift](https://huggingface.co/datalab-to/lift)**  
  作者: datalab-to | 点赞: 159 | 下载: 6,676  
  基于 Qwen3.5 的多模态模型，专为 PDF 文档理解优化。

- **[Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF)**  
  作者: Jackrong | 点赞: 95 | 下载: 49,935  
  Qwen3.6-27B 的代码兼容版 GGUF，支持视觉输入与 MTP 推理。

- **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)**  
  作者: HauhauCS | 点赞: 93 | 下载: 32,222  
  Gemma-4-12B 的 QAT 量化且无审查版，强调平衡性与多模态。

- **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
  作者: nvidia | 点赞: 712 | 下载: 61,857  
  NVIDIA 的流式语音识别模型（0.6B），适用于实时 ASR 场景。

---

### 🔧 专用模型（代码、数学、医疗、嵌入等）

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
  作者: yuxinlu1 | 点赞: 2,412 | 下载: 536,130  
  Gemma-4-12B 代码专用微调版 GGUF，专注代码生成与推理，下载极高。

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
  作者: yuxinlu1 | 点赞: 699 | 下载: 206,828  
  面向 Agent 的 Gemma-4 微调量化版，强化编码与终端交互能力。

- **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**（已在语言模型中列出，为保持唯一性，此处不再重复，或可归入）
  注：VibeThinker-3B 亦可归入数学专用模型，已在语言模型部分列出。

---

## 生态信号

**模型家族**：本周 **GLM-5.2**（MoE + DSA）、**Gemma-4-12B**（代码/Agent 微调）、**Qwen 3.5/3.6**（多模态 + MoE）和 **Ornith** 系列构成主力阵营。**DeepSeek v4** 和 **MiniMax-M3** 也展现出强势潜力。**开源权重**继续主导，几乎所有热门模型均开放完整权重（部分通过 GGUF/NVFP4 量化发布）。**量化与微调**成为社区主要动能：NVIDIA 的 NVFP4 格式（Qwen3.6、GLM-5.2）下载量惊人，社区定制（abliterated、uncensored、fable5）层出不穷，反映出用户对可控性、高效率部署和个性化需求的强烈追求。**多模态化**是不可逆趋势，超过半数榜单模型支持图像输入。

## 值得探索

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** (2,391 赞)  
   — 视觉定位领域的突破性模型，参数量仅 3B 即可实现开放词汇目标定位，适合研究图像‑语言对齐与 grounding 技术。

2. **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)** (1,248 赞)  
   — 国产新一代多模态大模型，在图像理解与生成任务上表现优异，代表闭源转开源的重要尝试。

3. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** (2,624 赞)  
   — 最新 MoE 架构 + DSA 稀疏注意力，周赞最高，值得深入测试其推理效率与多轮对话能力。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*