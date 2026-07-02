# Hugging Face 热门模型日报 2026-07-02

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-02 10:17 UTC

---

好的，作为AI模型生态分析师，这是根据您提供的2026年7月2日数据生成的《Hugging Face 热门模型日报》。

---

### 🤗 Hugging Face 热门模型日报 | 2026-07-02

#### **今日速览**

本周Hugging Face生态呈现出“头部大厂与新锐社区并行”的格局。智谱的**GLM-5.2** 系列凭借其MoE架构的强大性能，多个版本（原始版、NVIDIA优化版、社区去审查版）同时冲入榜单，成为本周最受关注的模型家族。DeepSeek发布了**DeepSeek-V4** 的两个变体，将MoE模型的参数规模推向新的高度，引起了广泛讨论。值得注意的是，社区微调与量化活动空前活跃，大量基于GLM-5.2、Qwen3.5/3.6、Gemma-4的高质量GGUF版本占据了榜单的半壁江山，尤其是专注于代码和Agent能力的“fable”系列，证明了开发者对高效、专用化本地部署模型的需求依旧强劲。

#### **热门模型**

##### 🧠 语言模型（LLM、对话模型、指令微调）

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  | 作者: zai-org | 👍 3,202 | ⬇️ 176,154  
  **一句话说明**: 智谱最新一代MoE大型语言模型，支持对话与文本生成，凭借其优异的性能和广泛的社区支持，成为本周趋势榜首。

- **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)**  | 作者: deepseek-ai | 👍 285 | ⬇️ 8,184  
  **一句话说明**: DeepSeek最新的V4系列“Pro”版本，探索更大参数规模下的文本生成能力，代表了模型尺寸的扩张趋势。

- **[deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)**  | 作者: deepseek-ai | 👍 127 | ⬇️ 23,939  
  **一句话说明**: DeepSeek-V4的“Flash”版本，旨在提供更快的推理速度，与Pro版形成互补，满足不同场景需求。

##### 🎨 多模态与生成（图像、视频、音频、文本到X）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  | 作者: baidu | 👍 1,615 | ⬇️ 758,489  
  **一句话说明**: 百度发布的通用OCR模型，能将图像中的文字直接提取为文本，凭借其极高的实用性和准确性，成为本周下载量与点赞双高的明星模型。

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  | 作者: nvidia | 👍 2,557 | ⬇️ 1,006,831  
  **一句话说明**: NVIDIA推出的通用物体定位模型，用户只需文本描述即可在图像中定位任何物体，下载量破百万，显示了强大的市场需求。

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  | 作者: HauhauCS | 👍 2,388 | ⬇️ 3,078,904  
  **一句话说明**: 基于Qwen3.6的去审查(MOE)模型变体，下载量惊人，反映了用户社区对高自由度、无内容限制的模型的强烈需求。

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  | 作者: krea | 👍 446 | ⬇️ 69,788  
  **一句话说明**: 新一代文生图模型Krea-2的Turbo加速版，在保持画质的同时实现了更快的推理速度。

##### 🔧 专用模型（代码、数学、医疗、嵌入）

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  | 作者: yuxinlu1 | 👍 2,557 | ⬇️ 614,069  
  **一句话说明**: 基于Gemma-4的代码专用模型量化版，融合了“fable”和“composer”技术，在编程任务上表现出色，是本周最热门的专用模型之一。

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  | 作者: yuxinlu1 | 👍 931 | ⬇️ 314,374  
  **一句话说明**: 同样是Gemma-4的衍生版本，但专注于增强模型的“Agent”能力，即自主执行任务和操作，体现了AI Agent的探索趋势。

- **[BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)**  | 作者: BugTraceAI | 👍 113 | ⬇️ 8,037  
  **一句话说明**: 专注于网络安全和渗透测试的专用模型，说明AI在特定垂直领域的应用正在细化。

##### 📦 微调与量化（社区微调、GGUF、AWQ）

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  | 作者: empero-ai | 👍 1,193 | ⬇️ 1,250,562  
  **一句话说明**: 基于Qwen3.5的社区微调模型GGUF版，下载量极高，表明社区微调工作正快速以量化形式分发。

- **[unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)**  | 作者: unsloth | 👍 491 | ⬇️ 239,283  
  **一句话说明**: 由知名优化团队Unsloth推出的GLM-5.2 GGUF量化版，降低了大模型本地部署的门槛。

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)**  | 作者: deepreinforce-ai | 👍 627 | ⬇️ 284,585  
  **一句话说明**: Ornith系列模型的GGUF版本，提供了从9B到397B的多种尺寸量化，满足不同硬件条件的用户。

#### **生态信号**

1.  **模型家族化与MoE成为主流**: **GLM-5.2**和**Qwen3.6**（以及其衍生模型）家族正在形成强大的生态。MoE（混合专家）架构已成为大模型的主流选择，榜单上近半数模型都是MoE架构或其量化版本，展现了其在性能和效率间的平衡优势。
2.  **开源社区主导部署，量化是“临门一脚”**: 即便基础模型由大厂发布，其广泛的应用落地仍依赖于社区生态。**GGUF**格式量化模型不仅占据了下载量的绝对优势，而且众多基于同一个基础模型（如Gemma-4、Qwen）的微调版本出现，证明了开源社区的活力是模型影响力扩散的关键。
3.  **应用场景细分化**: 除了通用的聊天模型，专门用于“代码”、“Agent”、“网络安全”甚至“去审查”的模型获得了极高关注。这表明用户的需求正从“拥有一个强大的模型”转向“拥有一个在特定领域能发挥最大价值的模型”。

#### **值得探索**

1.  **探索 [nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**：这是本周下载量最高的模型之一，其“万物可定位”的能力不仅强大且非常实用。可以尝试将其集成到您的图像处理或搜索项目中，体验其零样本定位的精准度。
2.  **对比 [zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2) 与 [nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)**：首先体验GLM-5.2的基础性能，然后尝试NVIDIA的NF4优化版本，感受在同一模型上，通过专业优化带来的推理速度和资源占用率的差异，这代表了模型与硬件协同优化的重要方向。
3.  **微调 [yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**：如果您是一名开发者，这款顶级的代码专用模型值得一试。将其部署到本地，作为您的AI编程助手，特别是在处理复杂逻辑和框架时，其“fable”和“composer”的特性可能会带来惊喜。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*