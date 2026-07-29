# Hugging Face 热门模型日报 2026-07-29

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-29 00:10 UTC

---

# Hugging Face 热门模型日报 — 2026-07-29

## 今日速览

本周 Hugging Face 榜单呈现三大亮点：**Moonshot AI 的 Kimi-K3 多模态模型以 7,951 点赞登顶**，标志着视觉语言模型持续占据社区焦点；**智谱 GLM-5.2（4,602 赞）和阿里 Qwen3.6-35B-A3B（2,567 赞）等国产大模型表现强势**，下载量均超百万；**社区量化与微调活动空前活跃**，Qwen3.6 系列衍生出数十个 GGUF 版本，其中 HauhauCS 的无审查版下载量高达 185 万。此外，百度 Unlimited-OCR 以近 270 万下载成为工具类模型黑马，微软则一口气推出 Mage-Flow 文生图、Fara1.5-27B 计算机控制等多模态新作。

---

## 热门模型

### 🧠 语言模型（LLM、对话、指令微调）

- **[poolside/Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1)**  
  作者：poolside | 点赞：799 | 下载：67,286  
  最新发布的 2.1 版本大型语言模型，周点赞迅速攀升，显示开发者对高效文本生成模型的关注。

- **[upstage/Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B)**  
  作者：upstage | 点赞：645 | 下载：4,804  
  Upstage 开源的 250B 超大参数模型，彰显韩国团队在基础大模型领域的持续投入。

- **[Nanbeige/Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B)**  
  作者：Nanbeige | 点赞：527 | 下载：18,933  
  3B 轻量级 LLM，为资源受限场景提供高性价比选择。

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  作者：zai-org | 点赞：4,602 | 下载：1,267,198  
  智谱 GLM 系列最新迭代，对话能力与推理表现突出，下载量已超 126 万，社区热度极高。

- **[fdtn-ai/antares-1b](https://huggingface.co/fdtn-ai/antares-1b)**  
  作者：fdtn-ai | 点赞：222 | 下载：7,666  
  面向安全领域的 1B 小模型，专注垂直场景的轻量化部署。

### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)**  
  作者：moonshotai | 点赞：7,951 | 下载：99,214  
  本周冠军！Moonshot AI 的最新多模态大模型，支持图像与文本联合理解，备受开发者追捧。

- **[microsoft/Mage-Flow](https://huggingface.co/microsoft/Mage-Flow)**  
  作者：microsoft | 点赞：415 | 下载：2,007  
  微软推出的文生图模型，支持编辑功能，标志着文本到图像生成进入实用阶段。

- **[thinkingmachines/Inkling](https://huggingface.co/thinkingmachines/Inkling)**  
  作者：thinkingmachines | 点赞：1,624 | 下载：39,052  
  多模态对话模型，兼顾视觉与文本理解，周点赞超 1,600，社区反馈积极。

- **[microsoft/Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B)**  
  作者：microsoft | 点赞：178 | 下载：1,543  
  支持“计算机使用”（Computer Use）的多模态模型，可执行 GUI 操作，探索 AI Agent 新范式。

- **[baseten/GLM-5.2-Vision-NVFP4](https://huggingface.co/baseten/GLM-5.2-Vision-NVFP4)**  
  作者：baseten | 点赞：130 | 下载：2,756  
  GLM-5.2 的视觉版本，采用 NVFP4 量化优化，在保留性能的同时降低推理成本。

- **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
  作者：moonshotai | 点赞：1,332 | 下载：681,111  
  Kimi 系列代码增强版，可理解多模态代码输入（如截图），下载量已接近 70 万。

- **[microsoft/Mage-Flow-Edit-Turbo](https://huggingface.co/microsoft/Mage-Flow-Edit-Turbo)**  
  作者：microsoft | 点赞：109 | 下载：1,260  
  Mage-Flow 的快速编辑版本，专注于图像到图像的指令式编辑。

- **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)**  
  作者：Qwen | 点赞：2,567 | 下载：6,158,876  
  阿里最新多模态 MoE 模型，35B 激活仅 3B，性能与效率平衡极佳，下载量突破 600 万。

- **[owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)**  
  作者：owensong | 点赞：264 | 下载：645  
  轻量级本地 TTS 模型，主打 CPU 与边缘设备上的语音合成。

- **[owensong/Inflect-Nano-v2](https://huggingface.co/owensong/Inflect-Nano-v2)**  
  作者：owensong | 点赞：102 | 下载：434  
  Inflect 系列更小版本，适合极致边缘场景的语音合成。

### 🔧 专用模型（代码、数学、医疗、嵌入、OCR、ASR 等）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  作者：baidu | 点赞：3,412 | 下载：2,694,935  
  百度推出的“无限制”OCR 模型，可处理任意场景文本识别，下载量接近 270 万，实用价值极高。

- **[Kwaipilot/KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev)**  
  作者：Kwaipilot | 点赞：287 | 下载：6,275  
  基于 Qwen3.5 MoE 的代码生成模型，为开发者提供专业的编程辅助。

- **[ATH-MaaS/OvisOCR2](https://huggingface.co/ATH-MaaS/OvisOCR2)**  
  作者：ATH-MaaS | 点赞：339 | 下载：47,129  
  基于 Qwen3.5 的 OCR 模型，在文档理解与文本提取上表现优秀。

- **[microsoft/VibeVoice-ASR-BitNet](https://huggingface.co/microsoft/VibeVoice-ASR-BitNet)**  
  作者：microsoft | 点赞：88 | 下载：1,754  
  采用 BitNet 架构的语音识别模型，探索低比特量化在 ASR 上的应用。

### 📦 微调与量化（社区微调、GGUF、AWQ、LoRA 等）

- **[DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF)**  
  作者：DavidAU | 点赞：849 | 下载：736,692  
  社区对 Qwen3.6 的极致微调+GGUF 量化版，无审查设计，下载量近 74 万。

- **[unsloth/Laguna-S-2.1-GGUF](https://huggingface.co/unsloth/Laguna-S-2.1-GGUF)**  
  作者：unsloth | 点赞：231 | 下载：129,601  
  Unsloth 团队优化的 Laguna 量化版，兼顾推理速度与模型质量。

- **[prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)**  
  作者：prism-ml | 点赞：1,083 | 下载：665,427  
  采用三值量化（2-bit）的 Bonsai 模型，以极致压缩换取可接受的对话质量。

- **[unsloth/Kimi-K3](https://huggingface.co/unsloth/Kimi-K3)**  
  作者：unsloth | 点赞：146 | 下载：410  
  Kimi-K3 的 Unsloth 量化版，为社区提供更易部署的变体。

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  作者：HauhauCS | 点赞：3,157 | 下载：1,855,505  
  Qwen3.6 MoE 的无审查 GGUF 版，下载量高达 185 万，反映玩家对“无限制”模型的高需求。

- **[poolside/Laguna-S-2.1-NVFP4](https://huggingface.co/poolside/Laguna-S-2.1-NVFP4)**  
  作者：poolside | 点赞：153 | 下载：180,545  
  Laguna 的 NVFP4 量化版，使用 NVIDIA 浮点格式优化推理。

- **[poolside/Laguna-S-2.1-GGUF](https://huggingface.co/poolside/Laguna-S-2.1-GGUF)**  
  作者：poolside | 点赞：160 | 下载：90,106  
  Laguna 的官方 GGUF 版本，方便在 llama.cpp 等框架中运行。

- **[LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF)**  
  作者：LuffyTheFox | 点赞：194 | 下载：99,660  
  另一个 Qwen3.6 无审查 GGUF 变体，融合 Hermes 风格微调。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
  作者：empero-ai | 点赞：2,502 | 下载：1,262,662  
  基于 Claude 数据蒸馏的推理增强模型，GGUF 量化后仍保持高推理能力，下载量超 126 万。

- **[prism-ml/Bonsai-27B-gguf](https://huggingface.co/prism-ml/Bonsai-27B-gguf)**  
  作者：prism-ml | 点赞：674 | 下载：2,339,098  
  1-bit 量化的极致压缩模型，以极小体积（约 3.4GB）保持对话能力，下载量突破 230 万。

- **[conradlocke/krea2-identity-edit](https://huggingface.co/conradlocke/krea2-identity-edit)**  
  作者：conradlocke | 点赞：564 | 下载：0  
  针对 Krea-2 的 LoRA 模型，用于图像身份编辑（如换脸、风格迁移），虽无下载但获赞颇多。

---

## 生态信号

**模型家族格局**：Qwen3.6 系列成为本周最大赢家，不仅原始模型下载超 600 万，社区微调与量化版本也占据了榜单多个席位，形成“母模型+衍生量化”的生态闭环。GLM-5.2 紧随其后，智谱的持续更新维持了高热度。MoonShot AI 的 Kimi 系列则以 K3 和 K2.7-Code 展示多模态与代码结合的潜力。

**开源与闭源趋势**：本周热门几乎全部为开源权重模型，且多数采用 Apache 或其他宽松许可证。微软、阿里、智谱、Upstage 等公司均积极开源，开源模型在社区影响力上已压倒闭源。

**量化与微调活动**：GGUF 格式成为绝对主流，几乎所有大模型版本都提供了 GGUF 量化版。值得注意的是，1-bit（Bonsai）和三值量化（Ternary-Bonsai）创新尝试获得高赞，表明社区对极端压缩不再仅仅是“能用就行”，而是追求在极小体积下保留可用质量。同时，“Uncensored”微调模型持续火爆，反映了部分用户对审查过滤的规避需求。

---

## 值得探索

1. **[prism-ml/Bonsai-27B-gguf](https://huggingface.co/prism-ml/Bonsai-27B-gguf)**  
   **理由**：1-bit 量化的 27B 模型，下载量超 230 万，是验证极端压缩下的语言模型能力边界的最佳案例。适合在极低资源环境下部署 LLM 的开发者研究。

2. **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)**  
   **理由**：阿里最新多模态 MoE 模型，激活参数仅 3B 却达到 35B 模型的效果。如果你需要在保持高性能的同时降低推理成本，这个模型值得深入测试。

3. **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
   **理由**：融合代码理解与视觉输入的独特模型，下载量近 70 万。对于需要从 UI 截图、手绘图或代码文档中提取信息并生成代码的场景，Kimi-K2.7-Code 提供了全新的多模态编程助手范式。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*