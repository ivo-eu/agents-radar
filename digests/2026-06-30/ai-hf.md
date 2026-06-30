# Hugging Face 热门模型日报 2026-06-30

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-30 10:45 UTC

---

# Hugging Face 热门模型日报 (2026-06-30)

## 📌 今日速览

本周 Hugging Face 生态呈现 **“超大规模 MoE + 密集量化”** 双轮驱动：**GLM‑5.2** 凭借其创新的 MoE‑DSA 架构迅速登顶（周赞近 3K），背后带动 unsloth、NVIDIA 等多个量化版本；**DeepSeek V4 系列** 开始向 Pro/Flash 分支扩展，并衍生出安全微调版本 Fable；**Qwen 家族** 继续高歌猛进，AgentWorld、Qwen3.6 等 MoE 模型涌现，并伴随大量 GGUF/NVFP4 量化版；视觉语言模型方面，**NVIDIA LocateAnything‑3B** 以精准目标定位能力成为最大黑马（周赞 2.5K，下载 800K）。社区量化活动异常活跃，unsloth 和 NVIDIA 的量化工具链成为主流。

---

## 🧠 语言模型（LLM / 对话 / 指令微调）

### [zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)  
- 作者: zai-org | 👍 2,980 | ⬇️ 142,547  
- **一句话**: 清华 GLM 最新 MoE‑DSA 架构模型，周赞最高，兼具强对话能力与高效推理。

### [Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)  
- 作者: Qwen | 👍 448 | ⬇️ 28,480  
- **一句话**: 通义千问推出的智能体世界模型，仅激活 3B 参数即可完成复杂 agent 任务。

### [deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)  
- 作者: deepreinforce-ai | 👍 293 | ⬇️ 26,151  
- **一句话**: Ornith 系列 9B 基础版，采用 Qwen3.5 架构，MIT 协议开放权重。

### [deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)  
- 作者: deepreinforce-ai | 👍 251 | ⬇️ 69,048  
- **一句话**: 35B MoE 版本，兼顾性能与部署友好性，成为社区 fine‑tune 热门基座。

### [deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)  
- 作者: deepreinforce-ai | 👍 172 | ⬇️ 2,564  
- **一句话**: 接近 400B 参数的 MoE 巨兽，面向科研与极限推理场景。

### [deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)  
- 作者: deepseek-ai | 👍 238 | ⬇️ 6,939  
- **一句话**: DeepSeek V4 专业版，集成 DSpark 加速技术，性能强悍。

### [deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)  
- 作者: deepseek-ai | 👍 103 | ⬇️ 4,446  
- **一句话**: 轻量 Flash 版本，适合快速推理与资源受限场景。

### [LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)  
- 作者: LiquidAI | 👍 159 | ⬇️ 17,839  
- **一句话**: 230M 参数小模型，适合边缘设备与快速实验。

---

## 🎨 多模态与生成（图像 / 视频 / 音频 / 视觉语言）

### [baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)  
- 作者: baidu | 👍 1,427 | ⬇️ 429,056  
- **一句话**: 百度推出的“无限制”OCR 模型，支持任意文本检测与识别，下载量超 42 万。

### [empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)  
- 作者: empero-ai | 👍 571 | ⬇️ 99,359  
- **一句话**: 基于 Qwen3.5 的视觉语言对话模型，融合 Mythos 数据，多模态推理能力强。

### [krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)  
- 作者: krea | 👍 403 | ⬇️ 45,668  
- **一句话**: Krea 第二代图像生成加速版，在质量与速度间取得平衡。

### [krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)  
- 作者: krea | 👍 246 | ⬇️ 32,008  
- **一句话**: 基础版 Krea‑2，未经过度优化，适合作为 fine‑tune 底模。

### [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)  
- 作者: nvidia | 👍 2,493 | ⬇️ 800,597  
- **一句话**: NVIDIA 发布的通用目标定位模型，3B 参数即可实现任意物体精准定位，周赞第二高。

### [nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)  
- 作者: nvidia | 👍 749 | ⬇️ 87,115  
- **一句话**: 流式语音识别模型，0.6B 参数适合实时 ASR 应用。

### [fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)  
- 作者: fal | 👍 119 | ⬇️ 0  
- **一句话**: 面向 LTX‑2.3 视频模型的 LoRA，增强 3D 真实感，刚发布暂无下载量。

### [ilkerzgi/fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs)  
- 作者: ilkerzgi | 👍 82 | ⬇️ 0  
- **一句话**: Krea‑2 的多风格 LoRA 集合，方便用户快速切换生成效果。

---

## 🔧 专用模型（代码 / 数学 / 安全 / 医疗等）

### [WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)  
- 作者: WeiboAI | 👍 751 | ⬇️ 67,777  
- **一句话**: 专注数学推理的 3B 模型，基于 Qwen2 微调，性能超越同参数级模型。

### [Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)  
- 作者: Chunjiang-Intelligence | 👍 132 | ⬇️ 1,519  
- **一句话**: DeepSeek V4 在网络安全领域的微调版，用于威胁检测与防御。

### [BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)  
- 作者: BugTraceAI | 👍 86 | ⬇️ 253  
- **一句话**: 27B 参数的安全专用模型，Q6 量化，擅长 offensive security 与漏洞分析。

---

## 📦 微调与量化（GGUF / NVFP4 / LoRA / 社区优化）

### [empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)  
- 作者: empero-ai | 👍 1,002 | ⬇️ 970,663  
- **一句话**: 多模态 Qwythos 的量化版，支持 llama.cpp 本地推理，下载近百万。

### [deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)  
- 作者: deepreinforce-ai | 👍 498 | ⬇️ 157,418  
- **一句话**: Ornith 35B 的 GGUF 量化版，极大降低部署门槛。

### [deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)  
- 作者: deepreinforce-ai | 👍 324 | ⬇️ 98,750  
- **一句话**: 9B 量化版，适合个人设备运行。

### [yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)  
- 作者: yuxinlu1 | 👍 2,512 | ⬇️ 575,255  
- **一句话**: 基于 Gemma‑4 的代码专用量化版，周赞第三、下载超 57 万。

### [yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)  
- 作者: yuxinlu1 | 👍 863 | ⬇️ 257,216  
- **一句话**: Gemma‑4 的 agentic 版本量化，强化终端与代理任务能力。

### [unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)  
- 作者: unsloth | 👍 471 | ⬇️ 180,394  
- **一句话**: unsloth 团队出品的 GLM‑5.2 GGUF 版，提供高效推理优化。

### [nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)  
- 作者: nvidia | 👍 176 | ⬇️ 104,746  
- **一句话**: NVIDIA 使用 ModelOpt 工具将 GLM‑5.2 量化为 NVFP4，显存占用大幅降低。

### [unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)  
- 作者: unsloth | 👍 118 | ⬇️ 155,333  
- **一句话**: 智能体世界模型的量化版，便于在本地搭建 agent 系统。

### [nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)  
- 作者: nvidia | 👍 380 | ⬇️ 5,495,402  
- **一句话**: 本榜下载量冠军（近 550 万），Qwen3.6 MoE 的 NVFP4 量化，平衡性能与资源。

### [HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)  
- 作者: HauhauCS | 👍 2,343 | ⬇️ 3,017,678  
- **一句话**: 社区热门的“无审查”视觉 MoE 量化版，下载量超 300 万，适合实验性用途。

### [Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)  
- 作者: Comfy-Org | 👍 195 | ⬇️ 10  
- **一句话**: Krea‑2 的 ComfyUI 节点包，方便在 ComfyUI 工作流中调用（刚上线，下载量尚低）。

---

## 🌐 生态信号

- **MoE 架构全面爆发**：GLM‑5.2（DSA）、Qwen3.6 / AgentWorld、Ornith 系列均采用 MoE 设计，以更低激活参数实现更大规模，成为 2026 年主流范式。  
- **DeepSeek V4 生态初成**：DeepSeek 继 V3 后，V4 推出 Pro / Flash / Fable 分支，并伴随 arXiv 论文，验证开源路线。  
- **量化生态趋近“即用”**：unsloth 与 NVIDIA ModelOpt 两大工具链覆盖 GGUF / NVFP4，几乎所有热门模型均在发布当天被量化，社区对本地部署的渴求强烈。  
- **视觉语言模型下沉**：OCR、目标定位、流式 ASR 等专用多模态模型获得大量下载，表明 AI 应用正从纯文本向感知智能延展。  
- **安全与数学垂直领域活跃**：数学推理（VibeThinker）与网络安全（Fable、BugTrace）小模型获得关注，专用微调正在成为新的生态增长点。

---

## 🔭 值得探索

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   - 极低参数实现通用目标定位，下载量达 80 万，代表“小而专”的具身智能新方向，可与视觉语言模型结合做 grounding 应用。

2. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   - 社区微调的无审查视觉 MoE 模型，周赞 2,343、下载 300 万+，展示了社区对“高自由度”模型的强烈需求，值得研究其微调策略与对齐平衡。

3. **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)**  
   - MIT 许可的 35B MoE，原始版与多个量化版均上榜，证明了其作为基座模型的通用性。适合作为下游 fine‑tune 或量化测试的基准。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*