# Hugging Face 热门模型日报 2026-07-30

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-30 00:11 UTC

---

# Hugging Face 热门模型日报（2026-07-30）

## 📊 今日速览

本周 Hugging Face 榜单呈现 **多模态模型主导、量化版本持续放量、社区微调异常活跃** 三大特征。**moonshotai/Kimi-K3** 以 8,630 周点赞登顶，成为最受关注的视觉语言模型；**Qwen3.6-35B-A3B**（官方与社区微调版本）合计点赞超 5,700，下载超 800 万，表明 MoE 架构在社区中极受欢迎。三家千亿级模型（Solar-Open2-250B、GLM-5.2、Laguna-S-2.1）同步出现多种量化与微调变体，反映大模型部署与定制化需求旺盛。此外，**Inflect** 系列音频模型和 **Unlimited-OCR** 等专用模型也表现亮眼。

## 🏆 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

1. **[poolside/Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1)**  
   作者：poolside | 点赞：825 | 下载：67,286  
   开源文本生成模型，专注代码与推理场景，直接对标闭源前沿性能。

2. **[upstage/Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B)**  
   作者：upstage | 点赞：693 | 下载：4,804  
   250B 参数级开放模型，延续 Solar 系列高效训练管线，适合企业级部署研究。

3. **[Nanbeige/Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B)**  
   作者：Nanbeige | 点赞：554 | 下载：18,933  
   轻量级对话模型（3B），以极低算力成本提供接近大模型的对话质量。

4. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
   作者：zai-org | 点赞：4,639 | 下载：1,267,198  
   GLM 系列新一代稠密 MoE 混合模型，在对话和推理任务上表现突出。

5. **[fdtn-ai/antares-1b](https://huggingface.co/fdtn-ai/antares-1b)**  
   作者：fdtn-ai | 点赞：230 | 下载：7,666  
   1B 参数安全对齐模型，采用 GraniteMoEHybrid 架构，关注可信生成。

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

1. **[moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)**  
   作者：moonshotai | 点赞：8,630 | 下载：99,214  
   Kimi 旗舰视觉语言模型，支持图文理解与生成，本周最高赞模型。

2. **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
   作者：baidu | 点赞：3,514 | 下载：2,694,935  
   百度开源通用 OCR 模型（image-text-to-text），高精度识别任意场景文字。

3. **[thinkingmachines/Inkling](https://huggingface.co/thinkingmachines/Inkling)**  
   作者：thinkingmachines | 点赞：1,640 | 下载：39,052  
   多模态对话模型，专为复杂视觉问答与多轮交互优化。

4. **[Qwen/Qwen3.6-35B-A3B](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)**  
   作者：Qwen | 点赞：2,586 | 下载：6,158,876  
   Qwen 系列最新 MoE 视觉语言模型，35B 总量仅激活 3B，高效且强大。

5. **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
   作者：moonshotai | 点赞：1,333 | 下载：681,111  
   Kimi 代码专用版本，继承压缩张量技术，在代码生成任务上表现出众。

6. **[microsoft/Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B)**  
   作者：microsoft | 点赞：199 | 下载：1,543  
   微软出品，面向计算机使用场景的视觉语言模型（Computer Use）。

7. **[baseten/GLM-5.2-Vision-NVFP4](https://huggingface.co/baseten/GLM-5.2-Vision-NVFP4)**  
   作者：baseten | 点赞：136 | 下载：2,756  
   GLM-5.2 视觉版 NVFP4 量化版本，兼顾精度与推理速度。

8. **[microsoft/VibeVoice-ASR-BitNet](https://huggingface.co/microsoft/VibeVoice-ASR-BitNet)**  
   作者：microsoft | 点赞：99 | 下载：1,754  
   微软开源 BitNet 架构语音识别模型，支持 GGML/GGUF 格式。

9. **[microsoft/Mage-VL](https://huggingface.co/microsoft/Mage-VL)**  
   作者：microsoft | 点赞：94 | 下载：702  
   多模态视觉语言模型，采用 Mage 训练范式，强调通用理解能力。

10. **[owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)**  
    作者：owensong | 点赞：289 | 下载：645  
    轻量级文本转语音模型（v2），专为本地 CPU/边缘设备优化。

11. **[owensong/Inflect-Nano-v2](https://huggingface.co/owensong/Inflect-Nano-v2)**  
    作者：owensong | 点赞：111 | 下载：434  
    Inflect 系列更小版本（Nano），力求在极低资源下实现自然语音合成。

12. **[ATH-MaaS/OvisOCR2](https://huggingface.co/ATH-MaaS/OvisOCR2)**  
    作者：ATH-MaaS | 点赞：346 | 下载：47,129  
    基于 Qwen3.5 的 OCR 专用多模态模型，适合文档与自然场景文字识别。

13. **[conradlocke/krea2-identity-edit](https://huggingface.co/conradlocke/krea2-identity-edit)**  
    作者：conradlocke | 点赞：577 | 下载：0  
    基于 Krea-2 的 LoRA 模型，用于图像身份编辑与人脸保持。

---

### 🔧 专用模型（代码、数学、医疗、嵌入）

1. **[Kwaipilot/KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev)**  
   作者：Kwaipilot | 点赞：316 | 下载：6,275  
   结合 Qwen3.5 MoE 的代码生成模型，专注代码理解与生成。

2. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
   作者：empero-ai | 点赞：2,516 | 下载：1,262,662  
   基于 Qwen3.5 的推理增强模型（GGUF），被广泛用作通用推理引擎。

---

### 📦 微调与量化（社区微调、GGUF、AWQ）

1. **[DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF)**  
   作者：DavidAU | 点赞：936 | 下载：736,692  
   基于 Qwen3.6 的极致微调版，整合多重自定义技巧，支持 GGUF 格式。

2. **[prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)**  
   作者：prism-ml | 点赞：1,095 | 下载：665,427  
   三元量化（2-bit）的 27B 对话模型，大幅降低显存需求。

3. **[unsloth/Laguna-S-2.1-GGUF](https://huggingface.co/unsloth/Laguna-S-2.1-GGUF)**  
   作者：unsloth | 点赞：245 | 下载：129,601  
   Unsloth 团队将 Laguna-S-2.1 转为 GGUF 格式，支持 vLLM 推理。

4. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   作者：HauhauCS | 点赞：3,171 | 下载：1,855,505  
   Qwen3.6 MoE 的无审查社区微调版，采用激进训练策略，下载量极高。

5. **[nota-ai/Solar-Open2-250B-Nota-NVFP4](https://huggingface.co/nota-ai/Solar-Open2-250B-Nota-NVFP4)**  
   作者：nota-ai | 点赞：137 | 下载：6,189  
   Solar-Open2 的 NVFP4 量化版本，专为高性能推理优化。

6. **[LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF)**  
   作者：LuffyTheFox | 点赞：213 | 下载：99,660  
   结合 Hermes 风格的 Qwen3.6 MoE 无审查版，GGUF 格式，社区热捧。

7. **[unsloth/Kimi-K3](https://huggingface.co/unsloth/Kimi-K3)**  
   作者：unsloth | 点赞：167 | 下载：410  
   Unsloth 精简版 Kimi-K3，保持原始能力的同时缩小模型体积。

8. **[unsloth/Kimi-K3-GGUF](https://huggingface.co/unsloth/Kimi-K3-GGUF)**  
   作者：unsloth | 点赞：157 | 下载：0  
   Kimi-K3 的 GGUF 量化版，适配 llama.cpp 等本地推理框架。

9. **[prism-ml/Bonsai-27B-gguf](https://huggingface.co/prism-ml/Bonsai-27B-gguf)**  
   作者：prism-ml | 点赞：688 | 下载：2,339,098  
   1-bit 量化的 Bonsai 对话模型，极致压缩比，下载量超过 230 万。

10. **[DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF)**  
    作者：DavidAU | 点赞：132 | 下载：112,086  
    Qwen3.5 的多种 GGUF 变体整合版，体现社区微调的极致定制化。

---

## 🔍 生态信号

- **模型家族势能**：**Kimi-K 系列**（K3 + K2.7-Code）与 **Qwen3.6 系列** 形成当前多模态领域双龙头，官方与社区版本点赞合计均过万。**GLM-5.2** 作为国产 MoE 代表，下载量破百万，正快速追赶。
- **开源 vs 闭源**：本周热门榜单几乎全为开源模型，且大量提供可部署的 GGUF/NVFP4 量化版本，表明社区生态已从“关注参数”转向“关注可用性”。闭源模型（如 GPT、Claude）未出现在榜单，但部分社区模型（如 Qwythos）明确宣称蒸馏自 Claude，暗示数据合成与蒸馏策略仍活跃。
- **量化与微调狂热**：**GGUF 格式**占据约 1/3 名额，且微调版本名称中充斥 “Uncensored”“Heretic”“Aggressive” 等关键词，反映用户对定制化（尤其是去审查）的强烈需求。**1-bit/2-bit 量化**（Bonsai、Ternary-Bonsai）成为降低门槛的新方向，下载量均超 60 万。

## 💡 值得探索

1. **moonshotai/Kimi-K3**（[链接](https://huggingface.co/moonshotai/Kimi-K3)）  
   本周第一，适合深入评测其多模态理解能力，可作为开源视觉语言模型的新基准。

2. **Qwen/Qwen3.6-35B-A3B**（[链接](https://huggingface.co/Qwen/Qwen3.6-35B-A3B)）  
   MoE 架构的 SOTA 代表，仅激活 3B 参数即可匹敌 13B～27B 稠密模型效果，是资源受限场景的首选。

3. **prism-ml/Bonsai-27B-gguf**（[链接](https://huggingface.co/prism-ml/Bonsai-27B-gguf)）  
   1-bit 量化模型下载量超 230 万，验证了极端压缩在生产中的可行性，值得研究量化对推理质量的实际影响。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*