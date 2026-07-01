# Hugging Face 热门模型日报 2026-07-01

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-01 11:36 UTC

---

# 🤗 Hugging Face 热门模型日报｜2026-07-01

## 今日速览

- **GLM-5.2 系列全面爆发**：zai-org 的 GLM-5.2 以 3,108 点赞位居榜首，其 MoE 架构与全新训练范式引发社区狂热；同时 NVIDIA 和 unsloth 分别推出 NVFP4 量化版和 GGUF 版，形成完整生态链。
- **Qwen 家族持续扩张**：Qwen3.5/3.6 衍生产品占据多个席位，HauhauCS 的无审查版（3,055,962 下载）说明社区对“去限制”模型需求旺盛，NVIDIA 则提供 Qwen3.6 的 NVFP4 优化版。
- **Gemma-4-12B 代码推理模型**：yuxinlu1 的 coder 版 GGUF 收获 2,541 赞，成为本周最热门的本地部署代码模型，与 Qwythos、Ornith 等形成竞争。
- **多模态落地加速**：百度 Unlimited-OCR（1,534 赞）、NVIDIA LocateAnything-3B（2,533 赞）分别代表了文档理解和视觉定位的垂直场景突破。
- **DeepSeek-V4 登场**：DeepSeek-V4-Pro/Fash 双版本及社区微调版（Fable）同时上榜，标志着 DeepSeek 正式进入 V4 时代。

## 热门模型

### 🧠 语言模型（LLM、对话、推理）

1. **GLM-5.2**  
   [https://huggingface.co/zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)  
   作者：zai-org｜点赞 3,108｜下载 159,967  
   ✅ 全新 MoE 架构的语言模型，多项基准超越同规模模型，是本周最受关注的开放权重 LLM。

2. **Qwythos-9B-Claude-Mythos-5-1M**  
   [https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M)  
   作者：empero-ai｜点赞 603｜下载 114,499  
   ✅ 基于 Qwen3.5 的指令微调版本，引入百万级合成对话数据，推理能力突出。

3. **Ornith-1.0-9B / 35B / 397B**  
   [https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B)（9B: 318赞）  
   [https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B)（35B: 274赞）  
   [https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B)（397B: 183赞）  
   ✅ 全系列基于 Qwen3.5 MoE 的 MIT 许可模型，从 9B 到 397B 覆盖不同预算需求，下载量可观。

4. **Qwen-AgentWorld-35B-A3B**  
   [https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B)  
   作者：Qwen｜点赞 482｜下载 34,371  
   ✅ 专为 Agent 世界建模设计的大语言模型，具备多模态理解与规划能力。

5. **DeepSeek-V4-Pro-DSpark / Flash-DSpark**  
   [https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark)（Pro: 264赞）  
   [https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark)（Flash: 119赞）  
   ✅ DeepSeek V4 系列首发，Pro 版专注精度，Flash 版优化速度，配套论文已上 arxiv。

6. **LiquidAI LFM2.5-230M**  
   [https://huggingface.co/LiquidAI/LFM2.5-230M](https://huggingface.co/LiquidAI/LFM2.5-230M)  
   作者：LiquidAI｜点赞 174｜下载 21,935  
   ✅ 极轻量的 230M 参数语言模型，适合边缘部署与快速实验。

7. **Chunjiang-Intelligence/DeepSeek-v4-Fable**  
   [https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable](https://huggingface.co/Chunjiang-Intelligence/DeepSeek-v4-Fable)  
   作者：Chunjiang-Intelligence｜点赞 136｜下载 1,766  
   ✅ 基于 DeepSeek V4 的网络安全专用微调版本，重点关注渗透测试与漏洞分析。

### 🎨 多模态与生成（图像、视频、OCR、视觉定位）

1. **baidu/Unlimited-OCR**  
   [https://huggingface.co/baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR)  
   作者：baidu｜点赞 1,534｜下载 630,246  
   ✅ 百度推出的通用 OCR 模型，支持不限语言、不限场景的文字识别，下载量逼近 63 万。

2. **nvidia/LocateAnything-3B**  
   [https://huggingface.co/nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)  
   作者：nvidia｜点赞 2,533｜下载 896,058  
   ✅ NVIDIA 开源的视觉定位模型，可根据文本描述精准定位图像中任意目标，推理速度快。

3. **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive**  
   [https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive)  
   作者：HauhauCS｜点赞 2,366｜下载 3,055,962  
   ✅ 基于 Qwen3.6 MoE 的无审查微调版，支持多模态输入，下载量突破 300 万，社区高度关注。

4. **krea/Krea-2-Turbo / Krea-2-Raw**  
   [https://huggingface.co/krea/Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo)（Turbo: 431赞）  
   [https://huggingface.co/krea/Krea-2-Raw](https://huggingface.co/krea/Krea-2-Raw)（Raw: 260赞）  
   ✅ Krea 第二代文生图模型，Turbo 版优化速度，Raw 版保留原始输出风格，均基于 Diffusers。

5. **fal/LTX-2.3-3DREAL-LoRA**  
   [https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA)  
   作者：fal｜点赞 132｜下载 0  
   ✅ 面向 LTX 视频模型的 3D 真实风格 LoRA，专为图生视频任务设计。

6. **Comfy-Org/Krea-2**  
   [https://huggingface.co/Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2)  
   作者：Comfy-Org｜点赞 215｜下载 10  
   ✅ ComfyUI 工作流的 Krea-2 集成节点，方便一键调用文生图管线。

### 🔧 专用模型（代码、视觉定位、Agent、安全）

1. **yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**  
   [https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF)  
   作者：yuxinlu1｜点赞 2,541｜下载 597,090  
   ✅ 基于 Gemma-4-12B 的代码推理专用模型，经 Fable 5 + Composer 2.5 微调，GGUF 格式便于本地部署。

2. **yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**  
   [https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF)  
   作者：yuxinlu1｜点赞 902｜下载 288,741  
   ✅ 同一系列的 Agent 版，强化终端交互与工具调用能力，适合自动化任务。

3. **InternScience/Agents-A1**  
   [https://huggingface.co/InternScience/Agents-A1](https://huggingface.co/InternScience/Agents-A1)  
   作者：InternScience｜点赞 111｜下载 511  
   ✅ 基于 Qwen3.5 MoE 的 Agent 模型，专注于多步骤复杂任务完成。

4. **meituan-longcat/LongCat-2.0**  
   [https://huggingface.co/meituan-longcat/LongCat-2.0](https://huggingface.co/meituan-longcat/LongCat-2.0)  
   作者：meituan-longcat｜点赞 117｜下载 0  
   ✅ 美团推出的长文本处理模型（暂无详细标注），标签仅 region:us，值得跟踪。

### 📦 微调与量化（GGUF、NVFP4、社区优化）

1. **empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF**  
   [https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF)  
   作者：empero-ai｜点赞 1,101｜下载 1,113,871  
   ✅ 热门 Qwythos 模型的 GGUF 量化版，下载量超 111 万，社区最爱的本地推理方案。

2. **deepreinforce-ai/Ornith-1.0-35B-GGUF**  
   [https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF)（576赞）  
   **deepreinforce-ai/Ornith-1.0-9B-GGUF**（356赞）  
   ✅ Ornith 系列官方 GGUF，覆盖 9B 和 35B，提供即用型 llama.cpp 推理包。

3. **nvidia/GLM-5.2-NVFP4**  
   [https://huggingface.co/nvidia/GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4)  
   作者：nvidia｜点赞 191｜下载 136,933  
   ✅ NVIDIA 使用 Model Optimizer 对 GLM-5.2 进行 4-bit NVFP4 量化，大幅降低推理显存。

4. **unsloth/GLM-5.2-GGUF**  
   [https://huggingface.co/unsloth/GLM-5.2-GGUF](https://huggingface.co/unsloth/GLM-5.2-GGUF)  
   作者：unsloth｜点赞 487｜下载 212,201  
   ✅ unsloth 出品的高效 GGUF 量化版，兼容 llama.cpp，一键部署 GLM-5.2。

5. **nvidia/Qwen3.6-35B-A3B-NVFP4**  
   [https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-35B-A3B-NVFP4)（397赞）  
   **nvidia/Qwen3.6-27B-NVFP4**（134赞）  
   ✅ NVIDIA 对 Qwen3.6 两大版本进行 NVFP4 量化，下载量高达 576 万，说明企业级量化需求旺盛。

6. **unsloth/Qwen-AgentWorld-35B-A3B-GGUF**  
   [https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF](https://huggingface.co/unsloth/Qwen-AgentWorld-35B-A3B-GGUF)  
   作者：unsloth｜点赞 125｜下载 196,441  
   ✅ AgentWorld 的 GGUF 量化版，方便开发者本地运行 Agent 模型。

7. **huihui-ai/Huihui-GLM-5.2-abliterated-GGUF**  
   [https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF)  
   作者：huihui-ai｜点赞 107｜下载 901  
   ✅ 社区制作的 GLM-5.2 “abliterated” 版（移除知识限制），GGUF 封装，适合创意场景。

## 生态信号

- **MoE 架构成为主流**：本周前 10 名中超过一半模型采用 Mixture-of-Experts（GLM、Qwen3.5/3.6、Ornith），MoE 在保持推理速度的同时提供大容量知识，成为开源社区的新基线。
- **量化工具三足鼎立**：GGUF（llama.cpp 生态）、NVFP4（NVIDIA Model Optimizer）、以及传统 safetensors 共存。NVIDIA 的 NVFP4 在下载量上表现惊人（Qwen3.6 版达 576 万），表明企业级低精度推理正在加速。
- **开源权重完胜闭源**：榜单上 30 个模型全部为开源可下载权重，无任何闭源 API 模型。DeepSeek V4 和 GLM-5.2 均提供完整权重，社区对“真正可控”的模型持续追捧。
- **无审查/“abliterated”需求旺盛**：HauhauCS 的无审查 Qwen3.6 版下载量接近 310 万，huihui-ai 的 GLM ablitated 版也进入榜单，说明用户对模型安全策略的灵活定制有真实需求。

## 值得探索

1. **zai-org/GLM-5.2** – 本周现象级模型，首次公开 MoE 稀疏训练新方法，在 MMLU、GSM8K 等基准上突破同等规模均值，且有多家机构提供量化支持，非常适合作为研究基线或部署测试。
2. **nvidia/LocateAnything-3B** – 视觉定位领域的新标杆，仅 3B 参数即可精准定位任意文本描述的物体，支持图像-文本交叉检索，可集成到机器人、自动驾驶或文档分析管线。
3. **HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive** – 下载量验证了其社区价值，适合需要更自由对话风格的开发者；同时作为 Qwen3.6 MoE 的微调范例，展示了如何利用基础模型进行快速领域 adapt。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*