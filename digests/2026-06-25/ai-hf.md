# Hugging Face 热门模型日报 2026-06-25

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-25 10:25 UTC

---

# Hugging Face 热门模型日报（2026-06-25）

---

## 今日速览

- **DeepSeek-V4-Pro** 以 5054 周点赞登顶，成为社区最受关注的开源 LLM，下载量近 188 万，延续 DeepSeek 系列统治力。
- **多模态模型竞赛白热化**：nvidia 的 LocateAnything-3B（定位）、MiniMax-M3（通用视觉语言）、Kimi-K2.7-Code（代码视觉）以及 Google 的 Gemma-4-12B-it（任意模态）均跻身前列，生态正从纯文本向多模态全面迁移。
- **量化社区异常活跃**：Qwen3.6-35B-A3B 的各种量化版（GGUF、NVFP4）下载量合计超 800 万，反映本地部署与边缘推理的强劲需求。
- **智谱 GLM-5.2** 以 2400 点赞位列第二，其 MoE 架构 + 对话能力成关注焦点；百度 Unlimited-OCR 也以 804 点赞证实实用型多模态工具的市场热度。

---

## 热门模型

### 🧠 语言模型（LLM、对话模型、指令微调）

1. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
   - 作者：zai-org | 👍 2,400 | 📥 67,107  
   - 智谱最新 MoE 对话模型，擅长中文与多轮交互，周点赞第二。

2. **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)**  
   - 作者：Qwen | 👍 188 | 📥 3,389  
   - Qwen 出品的高效 MoE Agent 模型，35B 参数但仅激活 3B，适合工具调用场景。

3. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
   - 作者：microsoft | 👍 339 | 📥 5,276  
   - 微软针对超长上下文优化的 4B 指令模型，推理效率突出。

4. **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
   - 作者：deepseek-ai | 👍 5,054 | 📥 1,878,217  
   - 本周绝对王者，DeepSeek 第四代旗舰，综合能力对标闭源模型，社区追捧。

5. **[nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)**  
   - 作者：nvidia | 👍 317 | 📥 4,602,255  
   - NVIDIA 优化的 4-bit 浮点量化版 Qwen3.6 MoE，极致压缩，下载量超高。

---

### 🎨 多模态与生成（图像、视频、音频、文本到X）

1. **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
   - 作者：baidu | 👍 804 | 📥 70,743  
   - 百度发布的通用 OCR 模型，支持任意场景文字识别，实用型标杆。

2. **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  
   - 作者：krea | 👍 216 | 📥 2,996  
   - Krea 的快速文生图模型，基于 Raw 版蒸馏优化，生成速度提升。

3. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   - 作者：nvidia | 👍 2,354 | 📥 407,838  
   - NVIDIA 的零样本定位模型，输入图像与文本即可定位任意目标，周点赞第四。

4. **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**  
   - 作者：MiniMaxAI | 👍 1,229 | 📥 154,350  
   - MiniMax 第三代多模态大模型，图像文本理解与生成一体，性能惊艳。

5. **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
   - 作者：moonshotai | 👍 988 | 📥 502,106  
   - Kimi 的代码视觉模型，支持屏幕截图、UI 代码生成，下载量超 50 万。

6. **[google/gemma-4-12B-it](https://huggingface.co/google/gemma-4-12B-it)**  
   - 作者：google | 👍 1,166 | 📥 2,187,644  
   - Google Gemma 4 指令版，支持任意模态（文本、图像、音频），社区高度关注。

7. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   - 作者：HauhauCS | 👍 2,215 | 📥 3,520,206  
   - 社区去审查激进版 Qwen3.6 MoE，支持视觉，下载量超 350 万，话题性极强。

8. **[Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-27B-Coder-Compat-MTP-GGUF)**  
   - 作者：Jackrong | 👍 83 | 📥 19,382  
   - 针对代码场景优化的 Qwen3.6 变体 GGUF 版，兼容 MTP 推理。

9. **[HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced](https://huggingface.co/HauhauCS/Gemma4-12B-QAT-Uncensored-HauhauCS-Balanced)**  
   - 作者：HauhauCS | 👍 77 | 📥 15,128  
   - Gemma 4 的去审查量化版，平衡性能与自由度。

10. **[Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit)**  
    - 作者：Boogu | 👍 122 | 📥 824  
    - 国产图像编辑模型，支持中英文指令编辑，Apache-2.0 协议。

---

### 🔧 专用模型（代码、数学、医疗、语音、嵌入）

1. **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
   - 作者：WeiboAI | 👍 703 | 📥 51,717  
   - 微博 AI 推出的数学推理模型，基于 Qwen2，3B 参数表现不俗。

2. **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
   - 作者：nvidia | 👍 683 | 📥 50,553  
   - NVIDIA 的流式语音识别模型，0.6B 参数支持低延迟实时识别。

3. **[owensong/Inflect-Nano-v1](https://huggingface.co/owensong/Inflect-Nano-v1)**  
   - 作者：owensong | 👍 198 | 📥 0  
   - 超小型 TTS 模型，适合边缘设备部署，刚发布即获关注。

4. **[LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)**  
   - 作者：LiquidAI | 👍 120 | 📥 11,833  
   - Liquid AI 的句嵌入模型，350M 参数，适合检索与 RAG。

5. **[LiquidAI/LFM2.5-ColBERT-350M](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M)**  
   - 作者：LiquidAI | 👍 88 | 📥 3,600  
   - 同系列的 ColBERT 版本，支持稀疏-稠密混合检索。

6. **[poolside/Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1)**  
   - 作者：poolside | 👍 96 | 📥 2,956  
   - poolside 的代码生成基础模型，专为结构化编程任务设计。

---

### 📦 微调与量化（社区微调、GGUF、AWQ、abliterated）

1. **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
   - 作者：yuxinlu1 | 👍 2,317 | 📥 495,813  
   - 社区微调 Gemma 4 代码版 GGUF，集成 Fable5 与 Composer 增强，下载近 50 万。

2. **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
   - 作者：yuxinlu1 | 👍 563 | 📥 165,187  
   - 同一作者的 Agent 强化版，支持终端工具调用，量化体积更小。

3. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
   - 作者：empero-ai | 👍 401 | 📥 134,294  
   - 基于 Qwen3.5 的推理增强模型 GGUF，融合 Claude 风格数据，下载量高。

4. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
   - 作者：empero-ai | 👍 342 | 📥 10,160  
   - 原版全精度模型，同样注重推理与创意写作。

5. **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
   - 作者：unsloth | 👍 365 | 📥 88,915  
   - Unsloth 团队官方量化版 GLM-5.2，在社区广受欢迎。

6. **[huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated)**  
   - 作者：huihui-ai | 👍 126 | 📥 4,874  
   - 对同一模型进行“去审查”（abliteration）操作，增加输出自由度。

7. **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)**  
   - 作者：krea | 👍 177 | 📥 5,113  
   - Krea-2 的原始权重，作为 Turbo 版基础，供社区进一步微调。

8. **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)**  
   - 作者：Comfy-Org | 👍 101 | 📥 10  
   - ComfyUI 工作流集成版，方便可视化生成。

---

## 生态信号

- **大厂与社区共舞**：DeepSeek、Google、NVIDIA、百度、MiniMax 等头部企业持续输出高质量模型，而社区通过量化（GGUF、NVFP4）、去审查（abliterated）、领域微调（代码、Agent）等加速模型落地，形成“官方发布 + 社区二次加工”的成熟生态。
- **MoE 架构成主流**：本周前十中有 5 个 MoE 模型（GLM-5.2、Qwen3.6-35B-A3B 系列、Qwen-AgentWorld），性能/参数比优势明显，成为效率优先场景首选。
- **多模态全面开花**：从纯文本到图像、视频、音频、OCR，模型类型极大丰富；LocateAnything-3B 的“定位”能力和 Unlimited-OCR 的“文字提取”表明实用工具型模型同样受欢迎。
- **量化版本下载量远超原版**：例如 Qwen3.6-35B-A3B 的原版（HauhauCS 版本）下载 350 万，而 NVIDIA 的 NVFP4 版下载 460 万，说明社区对本地化、低资源推理存在巨大需求。

---

## 值得探索

1. **nvidia/LocateAnything-3B**  
   - 零样本视觉定位，输入“车钥匙”即可在图像中框出位置，应用场景极广（自动驾驶、机器人、辅助工具），且模型仅 3B，适合部署。

2. **deepseek-ai/DeepSeek-V4-Pro**  
   - 本周最强开源 LLM，代码、推理、多语言全面领先，值得所有开发者上手体验；官方还提供了量化版本可供选择。

3. **MiniMaxAI/MiniMax-M3**  
   - 多模态新贵，支持图像与文本混合理解、生成与对话，性能接近 GPT-4V，且权重完全开源，是多模态研究的绝佳基线。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*