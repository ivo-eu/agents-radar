# Hugging Face 热门模型日报 2026-07-03

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-03 10:12 UTC

---

# Hugging Face 热门模型日报（2026-07-03）

## 📌 今日速览

国产大模型继续领跑社区热度：**GLM-5.2** 以 3,295 赞登顶，NVIDIA 推出的 **LocateAnything-3B**（2,583 赞）和社区微调的 **gemma-4-12B-coder** GGUF 版本（2,573 赞）紧随其后。**Qwen 3.5/3.6 系列** 涌现多款变体（Qwythos、AgentWorld、Uncensored 版），开源权重格局进一步向 MoE 架构倾斜。量化与微调活动异常活跃，GGUF 版本覆盖几乎所有热门模型，社区“去审查”趋势依然强劲。

---

## 🧠 语言模型（LLM、对话、指令微调）

- **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
  作者：zai-org｜赞：3,295｜下载：191,462  
  采用 MoE-DSA 架构的对话模型，本周热度第一，推理与生成能力受到社区广泛关注。

- **[deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)**  
  作者：deepseek-ai｜赞：316｜下载：9,388  
  DeepSeek V4 系列升级版，支持高效推理，arXiv 论文已公开。

- **[deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)**  
  作者：deepseek-ai｜赞：130｜下载：32,675  
  V4 的轻量快速版，适合本地部署，下载量反超 Pro 版。

- **[LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)**  
  作者：LiquidAI｜赞：194｜下载：29,645  
  仅 230M 参数的高效小模型，适合边缘设备与快速实验。

---

## 🎨 多模态与生成（图像、视频、文本到X）

- **[baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)**  
  作者：baidu｜赞：1,668｜下载：885,040  
  百度推出的通用 OCR 模型，支持任意场景文字提取，下载量极高。

- **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
  作者：nvidia｜赞：2,583｜下载：1,108,586  
  NVIDIA 的多模态定位模型，可精确定位图像中任意物体，本周爆款。

- **[deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)**  
  作者：deepreinforce-ai｜赞：199｜下载：8,079  
  基于 Qwen3.5 MoE 的超大规模多模态模型（397B 参数），旗舰级。

- **[deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)**  
  作者：deepreinforce-ai｜赞：314｜下载：211,406  
  35B 多模态版本，兼具性能与部署友好性。

- **[deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)**  
  作者：deepreinforce-ai｜赞：355｜下载：64,051  
  Ornith 系列最小模型，适合个人开发者快速体验。

- **[Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)**  
  作者：Qwen｜赞：518｜下载：45,455  
  Qwen 官方推出的 Agent 优化版，35B 总参、3B 激活，专为工具调用设计。

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)**  
  作者：empero-ai｜赞：650｜下载：131,323  
  基于 Qwen3.5 的 9B 多模态模型，融合 Claude-Mythos 风格数据。

- **[InternScience/Agents-A1](https://huggingface.co/InternScience/Agents-A1)**  
  作者：InternScience｜赞：193｜下载：3,530  
  专为 Agent 任务设计的 MoE 多模态模型，强调多步骤推理。

- **[krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)**  
  作者：krea｜赞：467｜下载：84,006  
  Krea 第二代文生图模型的 Turbo 版，生成速度与质量平衡出色。

- **[fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)**  
  作者：fal｜赞：149｜下载：0  
  LTX 视频模型的 3D 真实感 LoRA，专注图像到视频的 3D 风格。

---

## 🔧 专用模型（代码、数学、表格、安全等）

- **[google/tabfm-1.0.0-pytorch](https://huggingface.co/google/tabfm-1.0.0-pytorch)**  
  作者：google｜赞：133｜下载：450  
  Google 表格基础模型，支持零样本分类与回归，适用于结构化数据。

- **[nationaldesignstudio/rampart](https://huggingface.co/nationaldesignstudio/rampart)**  
  作者：nationaldesignstudio｜赞：108｜下载：1,149  
  ONNX 版本的 BERT 模型，专为 PII（个人身份信息）检测优化。

- **[meituan-longcat/LongCat-2.0](https://huggingface.co/meituan-longcat/LongCat-2.0)**  
  作者：meituan-longcat｜赞：167｜下载：0  
  美团发布的长上下文评估基准，专注于超长序列评测。

- **[Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)**  
  作者：Comfy-Org｜赞：235｜下载：10  
  ComfyUI 工作流节点项目，方便在 UI 中使用 Krea-2 模型。

---

## 📦 微调与量化（GGUF、AWQ、社区 LoRA 等）

- **[empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)**  
  作者：empero-ai｜赞：1,305｜下载：1,366,360  
  Qwythos 的 GGUF 量化版，方便 llama.cpp 本地运行。

- **[yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)**  
  作者：yuxinlu1｜赞：2,573｜下载：628,225  
  Gemma-4 编码模型的社区微调版，集成 Fable5 提示合成器，编程利器。

- **[yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)**  
  作者：yuxinlu1｜赞：971｜下载：329,391  
  同一作者对 Gemma-4 的 Agent 向微调，支持终端交互。

- **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
  作者：HauhauCS｜赞：2,409｜下载：3,029,679  
  Qwen3.6 MoE 的“去审查”激进版，下载量超 300 万，社区争议焦点。

- **[deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)**  
  作者：deepreinforce-ai｜赞：667｜下载：322,780  
  Ornith 35B 的 GGUF 版，适用于 llama.cpp 本地推理。

- **[deepreinforce-ai/Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF)**  
  作者：deepreinforce-ai｜赞：401｜下载：287,942  
  Ornith 9B 的 GGUF 版，社区量化首选。

- **[nvidia/Qwen3.6-27B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4)**  
  作者：nvidia｜赞：216｜下载：94,465  
  NVIDIA 使用 ModelOpt 进行 4-bit 浮点量化（NVFP4），兼顾精度与速度。

- **[nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)**  
  作者：nvidia｜赞：210｜下载：189,970  
  GLM-5.2 的 NVFP4 量化版，官方提供的高效推理方案。

- **[huihui-ai/Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF)**  
  作者：huihui-ai｜赞：139｜下载：3,683  
  GLM-5.2 的“去审查”+GGUF 双微调版，社区个性化定制。

- **[Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF)**  
  作者：Jackrong｜赞：120｜下载：44,807  
  Qwen3.6 编码优化的 MTP 版本 GGUF，专注代码多轮修正。

- **[BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6)**  
  作者：BugTraceAI｜赞：121｜下载：11,444  
  基于 Qwen3 的网络安全专用模型，Q6 量化版，针对渗透测试。

- **[ilkerzgi/fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs)**  
  作者：ilkerzgi｜赞：111｜下载：0  
  Krea-2 的风格 LoRA 集合，社区风格适配。

---

## 🌐 生态信号

本周模型生态呈现 **“多极竞争、量化遍地”** 的局面。  
- **家族势头**：GLM-5.2 凭借 MoE 架构和 DSA 技术力压群雄；Qwen 3.5/3.6 系列变体最多（AgentWorld、Qwythos、Uncensored），社区二次创作活跃；DeepSeek V4 双版本并进，论文加持。  
- **开源 vs 闭源**：所有热门模型均为开源权重，且 NVIDIA、百度等巨头积极开源专用模型（LocateAnything、OCR），推动行业标准。  
- **微调与量化**：GGUF 已成社区标准量化格式，几乎所有大模型第一时间被转换为 GGUF；去审查（abliterated/uncensored）微调持续火爆，下载量远超原版。

---

## 🔭 值得探索

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)**  
   - 理由：高赞高下载，NVIDIA 官方出品，零样本物体定位能力突出，适合视觉应用落地。

2. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)**  
   - 理由：本周总赞第一，MoE-DSA 架构可能是未来方向，值得深入对比评测。

3. **[HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)**  
   - 理由：下载量破 300 万，社区对“去审查”的强烈需求可见一斑，也反映了 Qwen3.6 基础模型的强大可塑性。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*