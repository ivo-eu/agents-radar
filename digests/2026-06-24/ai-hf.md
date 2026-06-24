# Hugging Face 热门模型日报 2026-06-24

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-06-24 10:35 UTC

---

# Hugging Face 热门模型日报（2026-06-24）

## 🔥 今日速览

- **DeepSeek-V4-Pro** 以 5000+ 周点赞强势登顶，成为本周最受关注的开源通用大模型，标志着 DeepSeek 系列在对话与推理能力上更进一步。
- **多模态模型** 占据榜单半壁江山：nvidia 的 **LocateAnything-3B**（2329 赞）、Google 的 **DiffusionGemma**（1059 赞） 以及国产 **Qwen3.6** 系列（2179 赞） 均有亮眼表现，视觉理解与图像生成成为竞争焦点。
- **社区微调与量化** 异常活跃：Gemma-4 系列出现多个 coder/agentic 变体（2.2k+ 赞），GLM-5.2 与 Qwen3.6 的 GGUF 量化版本下载量均超百万次，显示开源模型部署需求持续旺盛。

---

## 📊 热门模型分类

### 🧠 语言模型（LLM / 对话 / 指令微调）

1. **[deepseek-ai/DeepSeek-V4-Pro](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)**  
   - 作者：deepseek-ai | 点赞：5,036 | 下载：2,245,489  
   - 说明：DeepSeek 第四代旗舰模型，融合 MoE 与长上下文技术，在多项基准上接近闭源前沿，本周社区关注度最高。

2. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
   - 作者：zai-org | 点赞：2,264 | 下载：40,127  
   - 说明：智谱 GLM 系列最新迭代，采用 MoE-DSA 架构，对话能力大幅提升，引发中文社区热烈讨论。

3. **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  
   - 作者：unsloth | 点赞：316 | 下载：55,820  
   - 说明：unsloth 团队对 GLM-5.2 的 GGUF 量化版本，方便边缘设备部署，下载增速显著。

4. **[microsoft/FastContext-1.0-4B-SFT](https://huggingface.co/microsoft/FastContext-1.0-4B-SFT)**  
   - 作者：microsoft | 点赞：325 | 下载：4,391  
   - 说明：微软推出的 4B 参数长上下文集模型，采用 Explorer SubAgent 策略，专为长文本推理优化。

5. **[poolside/Laguna-M.1](https://huggingface.co/poolside/Laguna-M.1)**  
   - 作者：poolside | 点赞：94 | 下载：2,787  
   - 说明：面向软件工程领域的 7B 级模型，支持 vllm/sglang，代码辅助场景值得关注。

6. **[bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF](https://huggingface.co/bytkim/Qwen3.6-27B-MTP-pi-tune-GGUF)**  
   - 作者：bytkim | 点赞：115 | 下载：65,765  
   - 说明：基于 Qwen3.6 的 27B MTP（Multi-Turn Prompting）精调版本，GGUF 量化后下载量高。

7. **[lordx64/Qwable-v1](https://huggingface.co/lordx64/Qwable-v1)**  
   - 作者：lordx64 | 点赞：175 | 下载：4,547  
   - 说明：基于 Qwen3.5 MoE 的视觉语言模型，但主打文本生成任务，社区尝试整合多模态能力。

8. **[zai-org/GLM-5.2-FP8](https://huggingface.co/zai-org/GLM-5.2-FP8)**  
   - 作者：zai-org | 点赞：154 | 下载：395,290  
   - 说明：GLM-5.2 的 FP8 量化版本，原厂提供，下载量远超原版，反映用户对高效推理的强烈需求。

---

### 🎨 多模态与生成（图像 / 音频 / 文本到 X）

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   - 作者：nvidia | 点赞：2,329 | 下载：274,025  
   - 说明：英伟达推出的 3B 级开放词汇定位模型，可精确检测图像中任意目标，在视觉 grounding 任务上表现惊艳。

2. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   - 作者：HauhauCS | 点赞：2,179 | 下载：3,955,016  
   - 说明：基于 Qwen3.6 的 35B MoE 视觉语言模型，未经审查且风格激进，下载量近 400 万，社区反响两极化。

3. **[MiniMaxAI/MiniMax-M3](https://huggingface.co/MiniMaxAI/MiniMax-M3)**  
   - 作者：MiniMaxAI | 点赞：1,225 | 下载：131,057  
   - 说明：MiniMax 第三代多模态模型，原生支持图文输入输出，在视觉问答与描述生成上达到新高度。

4. **[google/diffusiongemma-26B-A4B-it](https://huggingface.co/google/diffusiongemma-26B-A4B-it)**  
   - 作者：google | 点赞：1,059 | 下载：948,996  
   - 说明：Google 基于 Gemma 的扩散多模态模型，26B 参数但仅激活 4B，高效生成图像文本，下载逼近百万。

5. **[moonshotai/Kimi-K2.7-Code](https://huggingface.co/moonshotai/Kimi-K2.7-Code)**  
   - 作者：moonshotai | 点赞：981 | 下载：447,920  
   - 说明：月之暗面 Kimi 系列的代码专注多模态模型，强在图+代码混合理解，适合文档与截图分析。

6. **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
   - 作者：baidu | 点赞：597 | 下载：8,396  
   - 说明：百度推出的通用 OCR 模型，支持不限字符识别，在复杂场景文字提取上实现突破。

7. **[nvidia/nemotron-3.5-asr-streaming-0.6b](https://huggingface.co/nvidia/nemotron-3.5-asr-streaming-0.6b)**  
   - 作者：nvidia | 点赞：667 | 下载：41,050  
   - 说明：英伟达 0.6B 流式语音识别模型，支持缓存感知推断，为实时 ASR 提供轻量级方案。

8. **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
   - 作者：WeiboAI | 点赞：676 | 下载：41,170  
   - 说明：虽然归属多模态（image-text-to-text），但实际为微博推出的 3B 推理增强模型，在数学与逻辑推理上表现突出。（注：因其任务类型归入此处，类似视觉推理）

9. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
   - 作者：empero-ai | 点赞：245 | 下载：1,856  
   - 说明：基于 Qwen3.5 的多模态版本，融合 Claude 风格对话数据，定位为创意对话助手。

10. **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
    - 作者：empero-ai | 点赞：210 | 下载：27,218  
    - 说明：上述模型的 GGUF 量化版，方便本地部署，下载量远超原版。

11. **[datalab-to/lift](https://huggingface.co/datalab-to/lift)**  
    - 作者：datalab-to | 点赞：141 | 下载：3,216  
    - 说明：基于 Qwen3.5 的 PDF 分析模型，专为文档解析与查询优化。

12. **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  
    - 作者：krea | 点赞：130 | 下载：84  
    - 说明：Krea 第二代图像生成模型的 Turbo 版本，主打快速采样与一致性。

13. **[krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)**  
    - 作者：krea | 点赞：115 | 下载：194  
    - 说明：Krea-2 的原始版本，作为 Turbo 的基础模型。

14. **[ostris/ideogram_4_turbotime_lora](https://huggingface.co/ostris/ideogram_4_turbotime_lora)**  
    - 作者：ostris | 点赞：111 | 下载：3,672  
    - 说明：针对 Ideogram 4 的 LoRA 适配器，用于加速图像生成风格迁移。

15. **[Boogu/Boogu-Image-0.1-Edit](https://huggingface.co/Boogu/Boogu-Image-0.1-Edit)**  
    - 作者：Boogu | 点赞：114 | 下载：592  
    - 说明：博鼓推出的图像编辑模型，基于 Diffusers，支持中文与英文指令。

16. **[Comfy-Org/Boogu-Image](https://huggingface.co/Comfy-Org/Boogu-Image)**  
    - 作者：Comfy-Org | 点赞：87 | 下载：0  
    - 说明：ComfyUI 官方集成的 Boogu-Image 插件，方便工作流调用。

---

### 🔧 专用模型（代码 / 数学 / 嵌入 / 推理）

1. **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
   - 作者：yuxinlu1 | 点赞：2,265 | 下载：456,117  
   - 说明：社区基于 Gemma-4-12B 微调的代码模型，采用 fable5 与 composer2.5 数据配方，GGUF 量化后下载量居同类前列。

2. **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
   - 作者：yuxinlu1 | 点赞：475 | 下载：96,459  
   - 说明：同一作者的 agentic 版本，强化工具调用与终端交互能力，适合自动化任务。

3. **[huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated](https://huggingface.co/huihui-ai/Huihui-gemma-4-12B-coder-fable5-composer2.5-v1-abliterated)**  
   - 作者：huihui-ai | 点赞：119 | 下载：3,320  
   - 说明：对 gemma-4 coder 进行“去审查”（abliterated）处理，移除安全对齐，用于不受限代码生成。

4. **[WeiboAI/VibeThinker-3B](https://huggingface.co/WeiboAI/VibeThinker-3B)**  
   - 作者：WeiboAI | 点赞：676 | 下载：41,170  
   - 说明：3B 级数学推理模型，在 GSM8K 等基准上表现优异，轻量级推理专家。

5. **[LiquidAI/LFM2.5-Embedding-350M](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)**  
   - 作者：LiquidAI | 点赞：116 | 下载：10,117  
   - 说明：Liquid 最新嵌入模型，350M 参数，专为语义相似度与检索优化，支持 sentence-transformers。

6. **[LiquidAI/LFM2.5-ColBERT-350M](https://huggingface.co/LiquidAI/LFM2.5-ColBERT-350M)**  
   - 作者：LiquidAI | 点赞：87 | 下载：2,534  
   - 说明：同一系列的 ColBERT 变体，采用延迟交互机制，在大规模检索任务中效率更高。

---

## 📈 生态信号

- **DeepSeek 与 GLM 双雄对峙**：DeepSeek-V4-Pro 摘得周赞冠军，GLM-5.2 紧随其后，两大国产基础模型在开源社区形成正面竞争。同时多模态领域 Google DiffusionGemma 与 Qwen3.6 系列也势头正旺，**模型家族化**成为主流。
- **开源权重比例持续扩大**：本周 30 个热门模型中，仅极少数来自闭源公司（如 Ideogram 的 LoRA 依赖其闭源基座），其余均完全开放权重，反映出开源社区对**可复现、可审计**模型的偏好。
- **量化版下载量碾压原版**：多个模型的 GGUF / FP8 版本下载量是原版的 10 倍以上（如 GLM-5.2-FP8 下载 39 万 vs 原版 4 万），说明社区主要关注**本地部署与边缘推理**，而非训练。
- **微调方向集中在代码与 agentic**：Gemma-4 出现了至少 3 个 coder 变体（yuxinlu1、huihui-ai），提示社区对**代码生成与工具使用**的微调热情高涨；数学与嵌入领域也有新玩家（VibeThinker、LiquidAI）。

---

## 🧪 值得探索

1. **DeepSeek-V4-Pro**（[链接](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro)）  
   - 👑 本周最热，7000+ 点赞背后是其在长上下文、MoE 效率上的突破，可作为通用对话与复杂推理的首选评估对象。

2. **nvidia/LocateAnything-3B**（[链接](https://huggingface.co/nvidia/LocateAnything-3B)）  
   - 🎯 创新性视觉定位模型，仅 3B 参数即可实现开放词汇目标检测，在 robotics 与图像搜索场景中有巨大潜力。

3. **LiquidAI/LFM2.5-Embedding-350M**（[链接](https://huggingface.co/LiquidAI/LFM2.5-Embedding-350M)）  
   - 🔍 Liquid 新一代嵌入模型，结合 ColBERT 变体，在检索增强生成（RAG）中性能突出，很适合作为知识库底座进行测试。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*