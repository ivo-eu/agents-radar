# Hugging Face 热门模型日报 2026-07-04

> 数据来源: [Hugging Face Hub](https://huggingface.co/) | 共 30 个模型 | 生成时间: 2026-07-04 09:06 UTC

---

# Hugging Face 热门模型日报（2026-07-04）

## 📌 今日速览
本周 Hugging Face 生态呈现三大热点：**Qwen 3.5/3.6 家族衍生模型强势霸榜**，以 Ornith、Qwythos 等社区微调版本及大量 GGUF 量化版本为主；**NVIDIA 推出 NVFP4 优化格式**，针对 Qwen 和 GLM 实现高效部署；**专用模型持续细化**，包括代码 agent、网络安全、PII 检测、表格零样本分类等。此外，多模态定位模型 LocateAnything-3B 下载量突破 119 万，成为本周最大黑马。

## 🧠 语言模型（LLM / 对话 / 推理）
| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [GLM-5.2](https://huggingface.co/zai-org/GLM-5.2) | zai-org | 3,359 / 208,920 | 智谱最新 MoE 对话模型，获本周最高点赞量，体现社区对国产 MoE 的高度关注。 |
| [DeepSeek-V4-Pro-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-DSpark) | deepseek-ai | 353 / 10,306 | DeepSeek 最新推理优化旗舰，搭载分布式稀疏注意力，代表前沿推理架构方向。 |
| [DeepSeek-V4-Flash-DSpark](https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-DSpark) | deepseek-ai | 146 / 40,271 | Pro 版的轻量级版本，兼顾速度与性能，适合资源敏感场景。 |
| [Nemotron-Labs-TwoTower-30B-A3B-Base-BF16](https://huggingface.co/nvidia/Nemotron-Labs-TwoTower-30B-A3B-Base-BF16) | nvidia | 116 / 10,479 | NVIDIA 的 “双塔” MoE 大模型，30B 总参数仅 3B 激活，极低推理成本。 |

## 🎨 多模态与生成（图像 / 视频 / 文本↔图像）
| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR) | baidu | 1,698 / 988,379 | 百度开源通用 OCR 模型，高下载量反映工业级文本识别需求旺盛。 |
| [Ornith-1.0-9B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B) | deepreinforce-ai | 369 / 69,837 | 基于 Qwen3.5 的多模态 MoE，9B 级性价比之选。 |
| [Ornith-1.0-35B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B) | deepreinforce-ai | 325 / 218,657 | Ornith 系列中端主力，35B 参数兼顾深度与速度。 |
| [Ornith-1.0-397B](https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B) | deepreinforce-ai | 203 / 33,268 | 接近 400B 参数的超级 MoE，旨在竞争顶级多模态推理。 |
| [Agents-A1](https://huggingface.co/InternScience/Agents-A1) | InternScience | 216 / 5,456 | 基于 Qwen3.5 MoE 的智能体专用模型，强调工具调用和任务规划。 |
| [Qwen-AgentWorld-35B-A3B](https://huggingface.co/Qwen/Qwen-AgentWorld-35B-A3B) | Qwen | 529 / 50,188 | 阿里官方智能体模型，MoE 稀疏结构，专为复杂环境设计。 |
| [Qwythos-9B-Claude-Mythos-5-1M](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M) | empero-ai | 663 / 135,665 | 融合 Claude 神话风格的 Qwen3.5 微调版，面向角色扮演与创意生成。 |
| [Krea-2-Turbo](https://huggingface.co/krea/Krea-2-Turbo) | krea | 488 / 89,384 | Krea 最新文生图模型，Turbo 版本在速度上显著提升，社区流行。 |
| [LocateAnything-3B](https://huggingface.co/nvidia/LocateAnything-3B) | nvidia | 2,592 / 1,194,542 | NVIDIA 零样本定位模型，3B 参数实现高精度空间理解，下载量惊人。 |
| [Comfy-Org/Krea-2](https://huggingface.co/Comfy-Org/Krea-2) | Comfy-Org | 244 / 10 | Krea-2 的 ComfyUI 工作流封装，便于节点式图像生成。 |

## 🔧 专用模型（代码 / 安全 / 表格 / PII）
| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [gemma-4-12B-agentic…-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-agentic-fable5-composer2.5-v2-3.5x-tau2-GGUF) | yuxinlu1 | 997 / 342,752 | Gemma 4 的 agentic 微调版，擅长终端命令执行和自动化任务。 |
| [gemma-4-12B-coder…-GGUF](https://huggingface.co/yuxinlu1/gemma-4-12B-coder-fable5-composer2.5-v1-GGUF) | yuxinlu1 | 2,589 / 641,260 | Gemma 4 代码专用微调版，高点赞反映开发者对轻量代码模型的需求。 |
| [Google tabfm-1.0.0-pytorch](https://huggingface.co/google/tabfm-1.0.0-pytorch) | google | 156 / 1,177 | Google 开源的表格基础模型，支持零样本分类与回归，开辟新赛道。 |
| [Huihui-GLM-5.2-abliterated-GGUF](https://huggingface.co/huihui-ai/Huihui-GLM-5.2-abliterated-GGUF) | huihui-ai | 149 / 4,701 | 社区对 GLM-5.2 的 “去审查” 微调，追求无约束输出（GGUF 量化）。 |
| [BugTraceAI-CORE-Ultra-27B-Q6](https://huggingface.co/BugTraceAI/BugTraceAI-CORE-Ultra-27B-Q6) | BugTraceAI | 127 / 12,001 | 网络安全专用大模型，支持渗透测试与漏洞分析（GGUF量化）。 |
| [nationaldesignstudio/rampart](https://huggingface.co/nationaldesignstudio/rampart) | nationaldesignstudio | 118 / 1,881 | 基于 BERT 的 PII 检测模型，ONNX 格式，适合隐私合规场景。 |

## 📦 微调与量化（GGUF / NVFP4 / LoRA / 社区优化）
| 模型 | 作者 | 点赞 / 下载 | 一句话说明 |
|------|------|-------------|------------|
| [Qwythos-9B… GGUF](https://huggingface.co/empero-ai/Qwythos-9B-Claude-Mythos-5-1M-GGUF) | empero-ai | 1,399 / 1,464,047 | Qwythos 的 GGUF 量化版，超 146 万下载，本地部署首选。 |
| [Ornith-1.0-35B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-35B-GGUF) | deepreinforce-ai | 692 / 359,659 | Ornith 35B 的 GGUF 量化，兼顾性能与存储。 |
| [Ornith-1.0-9B-GGUF](https://huggingface.co/deepreinforce-ai/Ornith-1.0-9B-GGUF) | deepreinforce-ai | 417 / 320,660 | 9B 级轻量量化，端侧推理热门选择。 |
| [Qwen3.6-27B-MTP-GGUF](https://huggingface.co/unsloth/Qwen3.6-27B-MTP-GGUF) | unsloth | 940 / 2,752,390 | Unsloth 优化的 Qwen3.6 多 token 预测 GGUF，下载量极高。 |
| [Qwen3.6-35B-A3B-Uncensored…](https://huggingface.co/HauhauCS/Qwen3.6-35B-A3B-Uncensored-HauhauCS-Aggressive) | HauhauCS | 2,442 / 2,993,053 | “激进去审查” 版 MoE 多模态模型，社区争议中下载量近 300 万。 |
| [Qwopus3.6-35B-A3B-Coder-MTP-GGUF](https://huggingface.co/Jackrong/Qwopus3.6-35B-A3B-Coder-MTP-GGUF) | Jackrong | 127 / 59,971 | 代码专属 MoE 量化版，支持多 token 预测。 |
| [Qwen3.6-27B-NVFP4](https://huggingface.co/nvidia/Qwen3.6-27B-NVFP4) | nvidia | 234 / 184,521 | NVIDIA 的 NVFP4 精度量化，大幅降低显存占用。 |
| [GLM-5.2-NVFP4](https://huggingface.co/nvidia/GLM-5.2-NVFP4) | nvidia | 216 / 236,501 | GLM-5.2 的 NVFP4 版本，NVIDIA 生态部署优化。 |
| [LTX-2.3-3DREAL-LoRA](https://huggingface.co/fal/LTX-2.3-3DREAL-LoRA) | fal | 152 / 0 | LTX 视频生成模型的 LoRA，专注 3D 写实风格。 |
| [fal-Krea-2-Style-LoRAs](https://huggingface.co/ilkerzgi/fal-Krea-2-Style-LoRAs) | ilkerzgi | 118 / 0 | Krea-2 风格迁移 LoRA 合集，拓展生成多样性。 |

## 🌍 生态信号
- **Qwen 家族持续主导**：Qwen 3.5/3.6 被超过 15 个热门模型用作基座，Ornith、Qwythos、Qwopus 等社区衍生版本百花齐放，MoE 架构成为主流。
- **量化与微调空前活跃**：GGUF 格式占据大量榜单，Unsloth、Imperium AI 等社区推动量化工具链成熟；NVFP4 的出现表明硬件厂商直接参与模型优化。
- **去审查与专用微调双轨并行**：Uncensored、Abliterated 版本下载量极高，反映用户对自由度需求；同时代码、安全、PII、表格等专用模型走向实用化。
- **开源权重持续深化**：DeepSeek V4、GLM-5.2、Google TabFM 等顶级模型开源，闭源与开源的差距进一步缩小。

## 🔍 值得探索
1. **nvidia/LocateAnything-3B** — 零样本定位能力惊艳，可集成到机器人、视觉问答等场景，且仅 3B 参数，非常适合边缘部署。
2. **deepseek-ai/DeepSeek-V4-Pro-DSpark** — 分布式稀疏注意力是推理效率的前沿尝试，关注其学术论文（arXiv:2606.19348）可获得最新技术洞见。
3. **google/tabfm-1.0.0-pytorch** — 表格基础模型零样本表现值得测试，特别适合金融、医疗等结构化数据任务，可能改变表格建模范式。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*