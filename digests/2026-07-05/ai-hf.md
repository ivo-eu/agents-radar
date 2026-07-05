# Hugging Face 热门模型日报 2026-07-05

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-05 09:32 UTC

---

# Hugging Face 热门模型日报 (2026-07-05)

## 📰 今日速览

本周 Hugging Face 热度集中在**量化版本的大规模 MoE 语言模型**和**多模态视觉-语言模型**上。微软系社区团队 `deepreinforce-ai` 连续发布 Ornith 系列多个尺度的 GGUF 量化版，引发下载热潮；NVIDIA 的 `LocateAnything-3B` 和智谱的 `GLM-5.2` 分别凭借视觉定位和 MoE 架构登上点赞榜前列。值得注意的是，**“uncensored”与“agentic”导向的微调模型**（如 `HauhauCS/Qwen3.6-35B-A3B-Uncensored`）下载量巨大，反映出社区对于安全限制较宽松且具备推理能力的模型的强烈需求。此外，DeepSeek-V4 系列的双版本发布标志着开源前沿模型竞赛进入新阶段。

---

## 🔥 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

| 模型 | 作者 | 点赞/下载 | 一句话说明 |
|------|------|-----------|------------|
| [zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2) | zai-org | 3,418 / 220,379 | 智谱最新 50B+ MoE 对话模型，凭借高吞吐和强中文能力登顶本周点赞榜。 |
| [deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark) | deepseek-ai | 374 / 12,580 | DeepSeek V4 的 Pro 版本，配合 DSpark 推理优化，面向高性能场景。 |
| [deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark) | deepseek-ai | 158 / 48,696 | V4 的 Flash 轻量版，更注重推理速度，与 Pro 版并行发布。 |
| [mistralai/Leanstral-1.5-119B-A6B](https://huggingface.co/mistralai/Leanstral-1.5-119B-A6B) | mistralai | 104 / 26 | Mistral 新旗舰 119B 参数，仅激活 6B，Apache-2.0 开源但下载量尚小。 |
| [AliesTaha/fable-traces](https://huggingface.co/AliesTaha/fable-traces) | AliesTaha | 141 / 277 | 基于 Qwen3 的指令微调实验模型，探索“fable”式思考链（CoT）。 |

### 🎨 多模态与生成（图像、视频、音频、文本到X）

| 模型 | 作者 | 点赞/下载 | 一句话说明 |
|------|------|-----------|------------|
| [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | 2,607 / 1,247,265 | NVIDIA 的视觉定位模型，支持开放词汇目标检测，下载量超百万。 |
| [HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | 2,463 / 3,018,257 | Qwen3.6 的 MoE 视觉版，去除内容限制并采用激进风格，下载量突破300万。 |
| [empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF) | empero-ai | 1,486 / 1,533,844 | 基于 Qwen3.5 融合 Claude 风格数据的图像-文本模型，GGUF 量化版极受欢迎。 |
| [baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR) | baidu | 1,719 / 1,044,217 | 百度发布的通用 OCR 模型，支持海量场景文字识别，下载超百万。 |
| [Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B) | Qwen | 538 / 55,113 | Qwen 官方推出的 Agent 专用 MoE 模型，含多模态输入，专为工具调用设计。 |
| [krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo) | krea | 502 / 99,049 | Krea-2 的 Turbo 版本，在文本到图像生成上追求更高速度与质量。 |
| [deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B) | deepreinforce-ai | 378 / 76,189 | Ornith 系列 9B 多模态版，支持图像-文本输入，社区微调强。 |
| [deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B) | deepreinforce-ai | 335 / 224,641 | Ornith-35B MoE 多模态版，于量化和下载量均表现突出。 |
| [deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B) | deepreinforce-ai | 210 / 70,585 | Ornith 系列超大杯 397B 参数 MoE，目前权重已开放。 |

### 🔧 专用模型（代码、数学、医疗、嵌入等）

| 模型 | 作者 | 点赞/下载 | 一句话说明 |
|------|------|-----------|------------|
| [google/tabfm-1.0.0-pytorch](https://huggingface.co/google/tabfm-1.0.0-pytorch) | google | 205 / 2,670 | Google 官方发布的表格数据基础模型，支持零样本分类与回归。 |
| [nationaldesignstudio/rampart](https://huggingface.co/nationaldesignstudio/rampart) | nationaldesignstudio | 125 / 2,783 | 基于 BERT 的 PII（个人身份信息）识别模型，支持 ONNX 和 Transformers.js。 |

### 📦 微调与量化（社区微调、GGUF、AWQ）

| 模型 | 作者 | 点赞/下载 | 一句话说明 |
|------|------|-----------|------------|
| [yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) | yuxinlu1 | 2,598 / 651,758 | 基于 Gemma-4-12B 的代码专用微调+GGUF 量化，融合 composer 架构。 |
| [yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF) | yuxinlu1 | 1,015 / 355,871 | 同系列 Agent 版本，强化终端交互和工具使用能力。 |
| [deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF) | deepreinforce-ai | 716 / 394,164 | Ornith-35B 的 GGUF 量化，支持 llama.cpp 本地运行。 |
| [deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF) | deepreinforce-ai | 427 / 352,002 | 9B 版本的 GGUF，适合资源受限环境部署。 |
| [nvidia/Qwen3.6-27B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4) | nvidia | 261 / 297,130 | NVIDIA 优化的 FP4 量化版 Qwen3.6，结合 ModelOpt 工具链。 |
| [nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4) | nvidia | 229 / 280,087 | GLM-5.2 的 FP4 量化版，精度损失小且下载量大。 |
| [Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF) | Jackrong | 134 / 84,951 | Qwen3.6 的 MoE 代码版 + GGUF 量化，支持多任务预测（MTP）。 |
| [huihui-ai/Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF) | huihui-ai | 163 / 5,609 | 社区对 GLM-5.2 进行“ablation”微调并量化，移除部分安全限制。 |
| [BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6) | BugTraceAI | 132 / 12,196 | 面向网络安全攻防的专用模型，28B 参数 Q6 量化。 |
| [DavidAU/Qwen3.5-9B-Claude-4.6-HighIQ-THINKING-HERETIC-UNCENSORED](https://huggingface.co/DavidAU/Qwen3.5-9B-Claude-4.6-HighIQ-THINKING-HERETIC-UNCENSORED) | DavidAU | 145 / 53,962 | 社区融合 Claude 风格的高智商无限制微调版本，主打推理与越狱。 |

---

## 🌐 生态信号

**家族势头**: Qwen3.5/3.6 生态依然最旺盛，衍生出 AgentWorld、Qwythos、Uncensored 等多个变体；**GLM-5.2** 作为智谱新一代 MoE 获得企业和社区双重关注；**Ornith** 系列从 9B 到 397B 全栈覆盖，成为社区最活跃的“开源 MoE 全家桶”。NVIDIA 则以量化工具链（NVFP4）和专用视觉模型切入，**“FP4 量化”正在成为替代 GGUF 的高效新选择**。DeepSeek-V4 双版本发布显示前沿闭源和开源界限模糊，但开源版下载量显著低于社区微调版。同时，“uncensored”类模型下载量超百万，表明 **“高智商 + 低限制”依然是社区最饥渴的需求**，但也带来监管隐忧。

---

## 🔍 值得探索

1. **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
   在 Gemma-4 基础上进行代码与 composer 架构适配的量化版本，代码生成质量出色且可本地运行，是开发者的高效本地编程助手。

2. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   仅 3B 参数实现开放词汇目标定位，在精度与速度间取得极佳平衡，适合视觉巡检、机器人感知等实时应用。

3. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   下载量超 300 万的“真香”模型：尽管争议巨大，但其在推理、创意写作和角色扮演上的表现值得研究，可作为安全对齐实验的对照样本。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*