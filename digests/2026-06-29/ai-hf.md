# Hugging Face 热门模型日报 2026-06-29

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-29 14:39 UTC

---

# Hugging Face 热门模型日报 — 2026-06-29

## 今日速览

本周 Hugging Face 社区呈现出三大趋势：**MoE（混合专家）大模型成为绝对主流**，GLM-5.2、Qwen3.5/3.6 系列及 DeepSeek-V4 的多个变体集中涌现，且量化版本（GGUF / NVFP4）下载量远超完整权重；**“无审查”微调版本持续高热**，HauhauCS 推出的 Qwen3.6 和 Gemma4 未过滤版分别斩获 2326 赞和 309 万下载，反映出社区对内容控制自由的高度需求；**多模态与专用模型并进**，NVIDIA 发布的定位模型 LocateAnything-3B 与百度 Unlimited-OCR 均登上热门，微软 FastContext 则在长上下文方向开辟新线。

---

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

1. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
   - 作者：zai-org | 点赞：2,891 | 下载：133,350  
   - 基于 MoE-DSA 架构的下一代中文对话模型，本周点赞数最高，是 GLM 家族 5.2 的官方原始权重。

2. **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)**  
   - 作者：Qwen | 点赞：423 | 下载：26,223  
   - 通义千问推出的 Agent + World Model 混合模型，35B 参数仅激活 3B，探索语言模型与智能体世界的融合。

3. **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)**  
   - 作者：deepreinforce-ai | 点赞：259 | 下载：19,170  
   - Ornith-1.0 系列的基础 9B 版本，基于 Qwen3.5 微调，以 MIT 许可开放。

4. **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)**  
   - 作者：deepreinforce-ai | 点赞：225 | 下载：38,857  
   - 同系列的 35B 完整版（非 MoE），适合需要更大模型容量的用户。

5. **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)**  
   - 作者：deepreinforce-ai | 点赞：159 | 下载：1,622  
   - 当前 Ornith 系列最大版本，397B 参数的 MoE 巨型模型，面向企业级推理。

6. **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LiquidAI/LFM2.5-230M)**  
   - 作者：LiquidAI | 点赞：147 | 下载：15,463  
   - 仅 2.3 亿参数的小型语言模型，适合边缘端部署，本周因轻量化话题登上趋势。

7. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
   - 作者：microsoft | 点赞：369 | 下载：7,027  
   - 微软推出的长上下文专用模型（4B 参数），通过 SFT 优化超长文本理解，标签含 "Explorer SubAgent"。

8. **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)**  
   - 作者：deepseek-ai | 点赞：204 | 下载：5,460  
   - DeepSeek-V4 的专业推理优化版（DSpark），附有 arxiv 论文，面向高速推理场景。

9. **[deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)**  
   - 作者：deepseek-ai | 点赞：89 | 下载：2,239  
   - V4 的轻量化 "Flash" 版本，进一步降低推理资源需求。

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

1. **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
   - 作者：baidu | 点赞：1,328 | 下载：362,945  
   - 百度开源的通用 OCR 模型，支持不限长度的图像文字识别，本周下载量超高。

2. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
   - 作者：empero-ai | 点赞：545 | 下载：79,540  
   - 结合 Claude 风格数据的 Qwen3.5 微调版本，支持图像+文本输入，推理能力突出。

3. **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  
   - 作者：krea | 点赞：384 | 下载：38,454  
   - Krea 文生图系列的最快推理版本（Turbo），基于 diffusers 框架，生成速度优化。

4. **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)**  
   - 作者：krea | 点赞：241 | 下载：27,464  
   - Krea-2 的基础原始权重，提供给需要自定义微调的创作者。

5. **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)**  
   - 作者：Comfy-Org | 点赞：185 | 下载：10  
   - ComfyUI 官方发布的 Krea-2 工作流，方便在 UI 界面中直接调用 Krea 模型。

6. **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)**  
   - 作者：fal | 点赞：109 | 下载：0  
   - 基于 LTX-2.3 的视频生成 LoRA，专为 3D 真实感场景优化，本周刚发布（下载量尚为零）。

7. **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
   - 作者：nvidia | 点赞：739 | 下载：76,154  
   - NVIDIA 推出的流式语音识别模型，0.6B 参数支持实时转录，NeMo 生态系统最新成果。

---

### 🔧 专用模型（代码、数学、医疗、嵌入）

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   - 作者：nvidia | 点赞：2,460 | 下载：728,320  
   - NVIDIA 的视觉定位模型，可指定任意目标在图像中的位置，本周点赞与下载双高。

2. **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
   - 作者：WeiboAI | 点赞：746 | 下载：63,449  
   - 微博 AI 推出的数学推理专用模型，基于 Qwen2 微调，专注解决数学问题。

3. **[Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)**  
   - 作者：Chunjiang-Intelligence | 点赞：126 | 下载：1,463  
   - DeepSeek-V4 的网络安全专用微调版本，使用 "Fable" 数据集，面向安全分析场景。

---

### 📦 微调与量化（社区微调、GGUF、AWQ）

1. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
   - 作者：empero-ai | 点赞：879 | 下载：907,682  
   - Qwythos-9B 的 GGUF 量化版，下载量接近百万，是本地部署推理的首选。

2. **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
   - 作者：yuxinlu1 | 点赞：828 | 下载：241,409  
   - Gemma4-12B 的 "智能体式" 微调 + GGUF 量化，标签含 terminal、agentic，适合自动化任务。

3. **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)**  
   - 作者：deepreinforce-ai | 点赞：445 | 下载：123,598  
   - Ornith-1.0-35B 的 GGUF 版本，方便在 llama.cpp 中运行大型 MoE 模型。

4. **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
   - 作者：yuxinlu1 | 点赞：2,493 | 下载：561,577  
   - Gemma4 的代码专用微调版 + GGUF 量化，结合 "fable5" 数据集，本周点赞与下载均位居前列。

5. **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)**  
   - 作者：deepreinforce-ai | 点赞：292 | 下载：68,667  
   - Ornith-1.0-9B 的量化版，适合中等算力设备。

6. **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
   - 作者：unsloth | 点赞：456 | 下载：164,180  
   - unsloth 推出的 GLM-5.2 GGUF 版，以其极快的量化效率著称。

7. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   - 作者：HauhauCS | 点赞：2,326 | 下载：3,089,944  
   - Qwen3.6 的无审查“激进”微调 + GGUF 版，本周下载量突破 300 万，社区热度极高。

8. **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)**  
   - 作者：nvidia | 点赞：161 | 下载：81,944  
   - NVIDIA 利用 ModelOpt 工具将 GLM-5.2 量化至 NVFP4 格式，面向 Hopper 架构 GPU 优化。

9. **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)**  
   - 作者：nvidia | 点赞：375 | 下载：5,392,518  
   - 本周下载量最高的模型！NVFP4 量化版 Qwen3.6-35B-MoE，在 NVIDIA GPU 上实现极致推理性能。

10. **[unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)**  
    - 作者：unsloth | 点赞：102 | 下载：116,693  
    - unsloth 量化的 Qwen AgentWorld 模型，让 Agent 模型也能在 llama.cpp 上运行。

11. **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)**  
    - 作者：HauhauCS | 点赞：104 | 下载：46,053  
    - Gemma4-12B 的无审查“平衡”微调 + QAT 量化版，兼顾多模态与推理。

---

## 生态信号

**模型家族势力格局**：本周 Qwen 家族（3.5/3.6/AgentWorld）与 GLM-5.2 形成双主角，两者均有官方、社区微调及多量化版本覆盖。DeepSeek-V4 和 Gemma4 紧随其后，Ornith-1.0 系列以“大小全尺寸+MIT许可”策略吸引开发者。**开源权重的主导地位未变**，30 个热门模型中无一闭源黑箱，但 NVIDIA 的 NVFP4 和 unsloth 的 GGUF 正在将“推理效率”变成新的竞争维度。**量化生态尤其活跃**：NVFP4 独占下载量榜首，GGUF 占据超半数席位，社区对“可部署性”的重视已超越基座模型本身。值得注意的是，“无审查”微调（Uncensored）模型以 300 万+ 下载量成为显性需求，预示未来开源社区在内容控制上的持续博弈。

---

## 值得探索

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — 点赞与下载双高的视觉定位模型，将“指哪看哪”从概念变为开箱即用的工具，适合物体检测、自动驾驶等场景。其 3B 参数兼顾精度与速度，建议开发者优先尝试推理脚本。

2. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)** — 微软在长上下文领域的全新尝试，4B 参数主打“超长文本快速推理”，配合 "Explorer SubAgent" 标签，可能是 Agent 与 RAG 结合的下一代架构。值得关注其论文与基准表现。

3. **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)** — 如果你想体验 GLM-5.2 但资源有限，unsloth 的 GGUF 版是最佳入口。结合其极低的量化时间和 token 损失，可快速部署一个对话或推理助手，并对比与其他量化版本（如 NVFP4）的效果差异。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*