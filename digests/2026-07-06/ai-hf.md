# Hugging Face 热门模型日报 2026-07-06

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-06 13:05 UTC

---

# Hugging Face 热门模型日报 (2026-07-06)

## 今日速览

本周 Hugging Face 迎来多个重量级发布：百度开源的无限制 OCR 模型 **Unlimited-OCR** 和三款 NVIDIA 新模型（定位、量化、双塔）表现抢眼，Qwen 3.5/3.6 系列及其社区微调版持续霸榜，尤其是 **Qwen3.6-35B-A3B** 系列的 uncensored 和 MoE 变体下载量突破百万。值得关注的是 **GLM-5.2** 系列以 3,505 点赞高居榜首，同时 NVIDIA 推出的 NVFP4 量化格式与 unsloth 的 GGUF 版本竞争激烈，表明开源生态对高效部署的需求仍在升温。此外，**Gemma-4** 与 **Ornith** 系列的代码/智能体量化版也获得大量社区关注。

---

## 热门模型

### 🧠 语言模型（LLM、对话、指令微调）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [**GLM-5.2**](https://huggingface.co/zai-org/GLM-5.2) | zai-org | 3,505 | 231,218 | 智谱 AI 新一代 MoE 对话模型，本周总点赞最高，代表国产开源 LLM 前沿。 |
| [**DeepSeek-V4-Pro-DSpark**](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark) | deepseek-ai | 399 | 14,276 | DeepSeek 第四代 Pro 版本，融合 DSpark 加速技术，推理性能突出。 |
| [**Hy3**](https://huggingface.co/tencent/Hy3) | tencent | 223 | 2 | 腾讯混元 V3 版本，标志性国产大模型，刚发布即上榜。 |
| [**fable-traces**](https://huggingface.co/AliesTaha/fable-traces) | AliesTaha | 172 | 2,903 | 基于 Qwen3 的指令微调模型，专注于长尾对话追踪。 |
| [**Leanstral-1.5-119B-A6B**](https://huggingface.co/mistralai/Leanstral-1.5-119B-A6B) | mistralai | 127 | 106 | Mistral 最新 119B MoE 模型，仅激活 6B 参数，极低推理成本。 |
| [**Nemotron-Labs-TwoTower-30B-A3B-Base-BF16**](https://huggingface.co/nvidia/Nemotron-Labs-TwoTower-30B-A3B-Base-BF16) | nvidia | 124 | 10,766 | NVIDIA 双塔结构 MoE 模型，适合高效批量推理与检索。 |
| [**LongCat-2.0**](https://huggingface.co/meituan-longcat/LongCat-2.0) | meituan-longcat | 91 | 43 | 美团发布的长上下文对话模型，支持超长输入处理。 |

### 🎨 多模态与生成（图像、视频、文本到图像等）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [**Qwythos-9B-Claude-Mythos-5-1M**](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M) | empero-ai | 693 | 149,421 | 基于 Qwen3.5 的视觉语言模型，融合 Claude 风格微调，通用图文理解能力强。 |
| [**Qwen-AgentWorld-35B-A3B**](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B) | Qwen | 551 | 57,835 | 通义千问官方智能体增强版，35B MoE 架构，面向工具调用与环境交互。 |
| [**Krea-2-Turbo**](https://huggingface.co/krea/Krea-2-Turbo) | krea | 522 | 109,470 | 高速文本生成图像模型，基于 Krea-2-Raw 蒸馏，适合实时创作。 |
| [**Ornith-1.0-9B**](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B) | deepreinforce-ai | 386 | 86,136 | Ornith 系列 9B 原版，多模态能力均衡，社区关注度高。 |
| [**Ornith-1.0-35B**](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B) | deepreinforce-ai | 344 | 231,342 | 35B 大参数版本，在多模态推理和视觉问答上表现更佳。 |
| [**Agents-A1**](https://huggingface.co/InternScience/Agents-A1) | InternScience | 325 | 8,766 | 上海 AI 实验室开源的智能体基础模型，支持图文输入与工具规划。 |

### 🔧 专用模型（代码、OCR、表格、定位、安全等）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [**LocateAnything-3B**](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | 2,626 | 1,340,559 | NVIDIA 推出的零样本视觉定位模型，可识别任意物体位置，推理速度快。 |
| [**Unlimited-OCR**](https://huggingface.co/baidu/Unlimited-OCR) | baidu | 1,771 | 1,070,230 | 百度开源的无限制 OCR 模型，支持多种语言和复杂版面，下载量破百万。 |
| [**tabfm-1.0.0-pytorch**](https://huggingface.co/google/tabfm-1.0.0-pytorch) | google | 240 | 7,036 | Google 推出的表格数据基础模型，支持零样本分类与回归，开箱即用。 |
| [**rampart**](https://huggingface.co/nationaldesignstudio/rampart) | nationaldesignstudio | 133 | 3,821 | ONNX 格式的 BERT 模型，专为 PII（个人身份信息）检测与 token 分类设计。 |

### 📦 微调与量化（社区微调、GGUF、NVFP4）

| 模型 | 作者 | 点赞 | 下载 | 一句话说明 |
|------|------|------|------|------------|
| [**gemma-4-12B-coder-fable5-composer2.5-v1-GGUF**](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) | yuxinlu1 | 2,612 | 664,319 | Gemma-4 代码版的高质量 GGUF 量化，针对编程任务优化，点赞极高。 |
| [**Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive-GGUF**](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | 2,511 | 2,910,241 | 去限制的 Qwen3.6 MoE 量化版，下载量逼近 300 万，社区热度顶峰。 |
| [**Qwythos-9B-Claude-Mythos-5-1M-GGUF**](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF) | empero-ai | 1,597 | 1,617,508 | 同系列量化版，满足本地部署需求，下载量持续增长。 |
| [**gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF**](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF) | yuxinlu1 | 1,037 | 370,884 | Gemma-4 智能体版量化，支持终端与工具调用，适合自主代理场景。 |
| [**Qwen3.6-27B-MTP-GGUF**](https://huggingface.co/unsloth/Qwen3.6-27B-MTP-GGUF) | unsloth | 971 | 2,818,499 | unsloth 优化的 Qwen3.6 量化版，MTP 多任务预测，下载量极高。 |
| [**Ornith-1.0-35B-GGUF**](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF) | deepreinforce-ai | 747 | 436,780 | Ornith 35B 的 GGUF 量化，降低部署门槛，MIT 开源协议。 |
| [**Ornith-1.0-9B-GGUF**](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF) | deepreinforce-ai | 437 | 393,142 | 9B 版本量化，轻量灵活，适合个人设备运行。 |
| [**Qwen3.6-27B-NVFP4**](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4) | nvidia | 283 | 430,676 | NVIDIA 发布的 Qwen3.6 四比特浮点量化，兼顾速度与精度。 |
| [**GLM-5.2-NVFP4**](https://huggingface.co/nvidia/GLM-5.2-NVFP4) | nvidia | 242 | 365,499 | 对 GLM-5.2 的 NVFP4 量化版，体现 NVIDIA 与国产模型的生态合作。 |
| [**Huihui-GLM-5.2-abliterated-GGUF**](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF) | huihui-ai | 172 | 6,660 | “去除限制”版 GLM-5.2，社区通过 abliteration 技术解锁对话自由度。 |
| [**Qwen3.5-9B-Claude-4.6-HighIQ-THINKING-HERETIC-UNCENSORED**](https://huggingface.co/DavidAU/Qwen3.5-9B-Claude-4.6-HighIQ-THINKING-HERETIC-UNCENSORED) | DavidAU | 154 | 58,755 | 基于 Qwen3.5 的极端微调，强调高智商与无限制思考，风格独特。 |
| [**Qwopus3.6-35B-A3B-Coder-MTP-GGUF**](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF) | Jackrong | 147 | 126,831 | 面向代码的 Qwen3.6 量化版，MTP 加速，深受开发者喜爱。 |
| [**BugTraceAI-CORE-Ultra-27B-Q6**](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6) | BugTraceAI | 136 | 12,438 | 网络安全专用模型的 Q6 量化，支持渗透测试场景。 |

---

## 生态信号

**Qwen 家族持续领跑，MoE 与量化双轮驱动。** Qwen 3.5/3.6 系列在本周榜单中占据 10 余个席位，其 35B-A3B 的 MoE 结构（约 35B 总参 / 3B 激活）成为社区微调与量化的主力；多款 uncensored / abliterated 版本表明用户对“无限制”能力的强烈需求。**NVIDIA 加速布局模型优化**，通过 NVFP4 量化格式打入 Hugging Face 生态，Qwen3.6 和 GLM-5.2 均获得官方支持，与 unsloth 的 GGUF 形成竞争。**国产模型百花齐放**：GLM-5.2 以点赞数第一彰显智谱 AI 的影响力，百度、腾讯、美团的模型也进入视野，但社区微调版本在下载量上远超大厂原版，反映开源社区的二次创作活力。**量化格式之争趋于白热化**：GGUF 依然占据绝对主流，NVFP4 作为 NVIDIA 硬件原生格式崭露头角，未来可能影响本地推理硬件的选择。

---

## 值得探索

1. **[nvidia/LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B)** — 零样本视觉定位的里程碑，3B 参数实现任意物体检测，下载量 134 万，潜力巨大，适合机器人、自动驾驶等场景快速落地。

2. **[zai-org/GLM-5.2](https://huggingface.co/zai-org/GLM-5.2)** — 本周总点赞最高的国产 MoE LLM，支持对话与推理，且 NVIDIA 已为其推出 NVFP4 量化版，生态成熟度高，值得研究其 MoE 架构设计。

3. **[unsloth/Qwen3.6-27B-MTP-GGUF](https://huggingface.co/unsloth/Qwen3.6-27B-MTP-GGUF)** — 下载量接近 300 万的 Qwen 量化版，unsLoth 的优化技术结合 MTP 多任务预测，是本地部署高性价比之选，适合作为私有部署的基准模型。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*