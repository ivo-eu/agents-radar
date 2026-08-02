# Hugging Face 热门模型日报 2026-08-02

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-08-02 00:13 UTC

---

# Hugging Face 热门模型日报（2026-08-02）

## 今日速览

本日热门榜由月之暗面 `moonshotai/Kimi-K3` 领跑，多模态与生成类模型占据近三分之一席位。DeepSeek-V4-Flash 系列、GLM-5.2 等高权重语言模型持续受关注，下载量领先。Qwen3.6 系列成为社区微调与量化的核心底座，多款 uncensored/角色扮演 GGUF 版本获得极高下载。量化技术呈现多元发展：NVFP4、Ternary 2-bit、BitNet、compressed-tensors 等低比特方案进入主流视野。百度 Unlimited-OCR、微软 VibeVoice-ASR 等专用模型表现突出，垂直任务与边缘部署需求明显上升。

## 热门模型

### 🧠 语言模型

- **deepseek-ai/DeepSeek-V4-Flash-0731**  
  [https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731)  
  作者：deepseek-ai | 点赞：1,417 | 下载：15,366  
  一句话说明：DeepSeek-V4 的最新快照版本，以高性价比推理和对话能力登上趋势榜。

- **zai-org/GLM-5.2**  
  [https://huggingface.co/zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)  
  作者：zai-org | 点赞：4,736 | 下载：1,683,442  
  一句话说明：智谱新一代 GLM 模型，采用 MoE+动态稀疏注意力，对话能力备受关注。

- **Nanbeige/Nanbeige4.2-3B**  
  [https://huggingface.co/Nanbeige/Nanbeige4.2-3B](https://huggingface.co/Nanbeige/Nanbeige4.2-3B)  
  作者：Nanbeige | 点赞：611 | 下载：27,892  
  一句话说明：轻量级 3B 语言模型，适合资源受限场景下的高效部署。

- **poolside/Laguna-S-2.1**  
  [https://huggingface.co/poolside/Laguna-S-2.1](https://huggingface.co/poolside/Laguna-S-2.1)  
  作者：poolside | 点赞：868 | 下载：77,021  
  一句话说明：poolside 自研语言模型新版本，定位企业级 AI 推理。

- **upstage/Solar-Open2-250B**  
  [https://huggingface.co/upstage/Solar-Open2-250B](https://huggingface.co/upstage/Solar-Open2-250B)  
  作者：upstage | 点赞：717 | 下载：13,426  
  一句话说明：Upstage 开源的 250B 超大模型，进一步拉高开源权重上限。

- **deepseek-ai/DeepSeek-V4-Flash**  
  [https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash)  
  作者：deepseek-ai | 点赞：1,947 | 下载：2,814,414  
  一句话说明：DeepSeek-V4 主版本，下载超 280 万，是当前最热门的对话模型之一。

### 🎨 多模态与生成

- **moonshotai/Kimi-K3**  
  [https://huggingface.co/moonshotai/Kimi-K3](https://huggingface.co/moonshotai/Kimi-K3)  
  作者：moonshotai | 点赞：9,476 | 下载：559,924  
  一句话说明：月之暗面最新多模态模型，周点赞近万登顶，采用压缩张量技术。

- **owensong/Inflect-Micro-v2**  
  [https://huggingface.co/owensong/Inflect-Micro-v2](https://huggingface.co/owensong/Inflect-Micro-v2)  
  作者：owensong | 点赞：361 | 下载：1,565  
  一句话说明：面向 CPU/边缘设备的轻量 TTS 模型，主打本地语音合成。

- **thinkingmachines/Inkling-Small**  
  [https://huggingface.co/thinkingmachines/Inkling-Small](https://huggingface.co/thinkingmachines/Inkling-Small)  
  作者：thinkingmachines | 点赞：212 | 下载：3,998  
  一句话说明：Inkling 多模态模型的小型版本，兼顾性能与部署效率。

- **microsoft/Mage-VL**  
  [https://huggingface.co/microsoft/Mage-VL](https://huggingface.co/microsoft/Mage-VL)  
  作者：microsoft | 点赞：171 | 下载：10,525  
  一句话说明：微软视觉语言模型，专注多模态理解任务。

- **Audio8/Audio8-TTS-Preview-0.6b**  
  [https://huggingface.co/Audio8/Audio8-TTS-Preview-0.6b](https://huggingface.co/Audio8/Audio8-TTS-Preview-0.6b)  
  作者：Audio8 | 点赞：166 | 下载：3,254  
  一句话说明：0.6B 参数 TTS 预览版，基于 ArkTTS 架构，提供新的语音合成选择。

- **microsoft/Fara1.5-27B**  
  [https://huggingface.co/microsoft/Fara1.5-27B](https://huggingface.co/microsoft/Fara1.5-27B)  
  作者：microsoft | 点赞：242 | 下载：2,775  
  一句话说明：具备 computer-use 能力的多模态模型，可用于 GUI 代理与视觉交互。

- **XYZAILab/XYZ-Aquila-mini**  
  [https://huggingface.co/XYZAILab/XYZ-Aquila-mini](https://huggingface.co/XYZAILab/XYZ-Aquila-mini)  
  作者：XYZAILab | 点赞：357 | 下载：650  
  一句话说明：基于 Qwen3.5-MoE 的多模态精简模型，适合研究快速迭代。

- **XYZAILab/XYZ-Aquila-pro**  
  [https://huggingface.co/XYZAILab/XYZ-Aquila-pro](https://huggingface.co/XYZAILab/XYZ-Aquila-pro)  
  作者：XYZAILab | 点赞：330 | 下载：923  
  一句话说明：强化 agentic-search 能力的多模态模型，面向智能代理搜索场景。

- **lodestones/Kroma**  
  [https://huggingface.co/lodestones/Kroma](https://huggingface.co/lodestones/Kroma)  
  作者：lodestones | 点赞：93 | 下载：0  
  一句话说明：用于 Krea2 的 LoRA 图像生成模型，兼容 ComfyUI，刚发布即上榜。

- **thinkingmachines/Inkling**  
  [https://huggingface.co/thinkingmachines/Inkling](https://huggingface.co/thinkingmachines/Inkling)  
 作者：thinkingmachines | 点赞：1,672 | 下载：59,076  
  一句话说明：Inkling 多模态模型正式版，凭均衡能力与生态支持获得高关注。

### 🔧 专用模型

- **baidu/Unlimited-OCR**  
  [https://huggingface.co/baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)  
  作者：baidu | 点赞：3,713 | 下载：2,457,387  
  一句话说明：百度推出的通用 OCR 模型，下载超 245 万，是当前最热门的专用视觉模型。

- **Kwaipilot/KAT-Coder-V2.5-Dev**  
  [https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev](https://huggingface.co/Kwaipilot/KAT-Coder-V2.5-Dev)  
  作者：Kwaipilot | 点赞：391 | 下载：10,771  
  一句话说明：基于 Qwen3.5-MoE 的代码生成模型，开发者版本侧重编码能力。

- **microsoft/VibeVoice-ASR-BitNet**  
  [https://huggingface.co/microsoft/VibeVoice-ASR-BitNet](https://huggingface.co/microsoft/VibeVoice-ASR-BitNet)  
  作者：microsoft | 点赞：141 | 下载：5,835  
  一句话说明：微软采用 BitNet 压缩的语音识别模型，为端侧 ASR 提供低比特方案。

### 📦 微调与量化

- **DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF**  
  [https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.6-27B-Fable-Fusion-711-Uncensored-Heretic-NM-DAU-NEO-MAX-MTP-GGUF)  
  作者：DavidAU | 点赞：1,234 | 下载：1,173,001  
  一句话说明：社区微调的 27B uncensored 版本 GGUF，下载量破百万，反映个性化生成需求。

- **unsloth/DeepSeek-V4-Flash-0731-GGUF**  
  [https://huggingface.co/unsloth/DeepSeek-V4-Flash-0731-GGUF](https://huggingface.co/unsloth/DeepSeek-V4-Flash-0731-GGUF)  
  作者：unsloth | 点赞：286 | 下载：4,048  
  一句话说明：unsloth 将 DeepSeek-V4 快照转为 GGUF，便于本地高效推理。

- **unsloth/Kimi-K3-GGUF**  
  [https://huggingface.co/unsloth/Kimi-K3-GGUF](https://huggingface.co/unsloth/Kimi-K3-GGUF)  
  作者：unsloth | 点赞：243 | 下载：41,337  
  一句话说明：Kimi-K3 的 GGUF 量化版，让多模态模型也能在低资源设备上运行。

- **unsloth/Kimi-K3**  
  [https://huggingface.co/unsloth/Kimi-K3](https://huggingface.co/unsloth/Kimi-K3)  
  作者：unsloth | 点赞：220 | 下载：1,072  
  一句话说明：unsloth 适配的 Kimi-K3 压缩张量版本，专注降低显存占用。

- **LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF**  
  [https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF](https://huggingface.co/LuffyTheFox/Qwen3.6-35B-A3B-Uncensored-Genesis-Hermes-V6-GGUF)  
  作者：LuffyTheFox | 点赞：286 | 下载：228,610  
  一句话说明：MoE 架构的 uncensored GGUF，融合 Genesis/Hermes 数据集，社区热度高。

- **nota-ai/Solar-Open2-250B-Nota-NVFP4**  
  [https://huggingface.co/nota-ai/Solar-Open2-250B-Nota-NVFP4](https://huggingface.co/nota-ai/Solar-Open2-250B-Nota-NVFP4)  
  作者：nota-ai | 点赞：151 | 下载：22,396  
  一句话说明：针对 Solar-Open2 的 NVFP4 量化版，配合 vLLM 在低显存下运行超大模型。

- **EschaLabs/Qwen3.6-35B-A3B-Escha-W2**  
  [https://huggingface.co/EschaLabs/Qwen3.6-35B-A3B-Escha-W2](https://huggingface.co/EschaLabs/Qwen3.6-35B-A3B-Escha-W2)  
  作者：EschaLabs | 点赞：112 | 下载：875  
  一句话说明：社区对 Qwen3.6-MoE 的二次微调，面向效率与特定风格输出。

- **DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF**  
  [https://huggingface.co/DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF](https://huggingface.co/DavidAU/Qwen3.5-9B-The-Defiant-Fable-Uncensored-Heretic-NEO-IMATRIX-MAX-MTP-GGUF)  
  作者：DavidAU | 点赞：186 | 下载：267,572  
  一句话说明：9B 尺寸的 uncensored 微调 GGUF，支持 MTP，适用于创意写作和角色扮演。

- **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
  [https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)  
  作者：HauhauCS | 点赞：3,223 | 下载：1,823,436  
  一句话说明：高赞 uncensored GGUF 版本，带视觉能力，下载近 182 万，社区需求极强。

- **prism-ml/Ternary-Bonsai-27B-gguf**  
  [https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf](https://huggingface.co/prism-ml/Ternary-Bonsai-27B-gguf)  
  作者：prism-ml | 点赞：1,133 | 下载：716,341  
  一句话说明：三元量化（2-bit）27B 模型的 GGUF 版，探索极限压缩与 CPU 推理。

- **unsloth/Laguna-S-2.1-GGUF**  
  [https://huggingface.co/unsloth/Laguna-S-2.1-GGUF](https://huggingface.co/unsloth/Laguna-S-2.1-GGUF)  
  作者：unsloth | 点赞：275 | 下载：170,109  
  一句话说明：Laguna-S-2.1 的官方 GGUF 量化，进一步降低本地部署门槛。

## 生态信号

模型家族方面，Kimi、DeepSeek-V4、GLM 和 Qwen3.6/3.5 是当前最强势的几条线，其中 Qwen3.6 已成为社区二次创作最活跃的底座。开源权重继续主导 Hugging Face 热度，头部新模型均开放权重与工具链，未出现闭源模型。量化与微调活动异常活跃：GGUF 覆盖几乎所有热门模型，NVFP4、Ternary 2-bit、BitNet、compressed-tensors 等低比特方案表明算力/内存受限场景成为重要部署目标；大量 uncensored/角色扮演微调版本下载量极高，显示个性化、边界化需求也在推动生态分化。

## 值得探索

- **moonshotai/Kimi-K3**：周点赞近万，采用压缩张量技术，代表新一代多模态模型方向，值得优先研究。  
- **prism-ml/Ternary-Bonsai-27B-gguf**：三元量化到 2-bit，是超低比特量化与 CPU 推理的典型样本，对边缘部署研究很有价值。  
- **microsoft/VibeVoice-ASR-BitNet**：BitNet 与 ASR 结合，展示端侧语音识别的新路径，适合关注高效专用模型的研究者。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*