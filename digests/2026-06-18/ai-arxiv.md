# ArXiv AI 研究日报 2026-06-18

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-18 03:43 UTC

---

# ArXiv AI 研究日报 (2026-06-18)

## 今日速览

今日投稿揭示了几个关键趋势：**扩散模型正从图像生成向语言推理领域延伸**，DreamReasoner-8B 首次将块级扩散解码应用于长链思维推理，效果媲美自回归模型。**多智能体决策范式迎来新突破**，基于虚构博弈的 LLM 协作框架显著提升了复杂任务中的决策质量。**模型安全与遗忘机制成为焦点**，MAST 和预训练阶段对齐等研究深入探讨如何精确移除特定能力而不破坏整体性能。此外，**领域专用智能体（药物发现、企业数据、医疗诊断）和评估基准（幻灯片生成、深度伪造检测）密集涌现**，推动 AI 从通用走向务实落地。

## 重点论文

### 🧠 大语言模型（架构、训练、对齐、评估）

- **[DreamReasoner-8B: Block-Size Curriculum Learning for Diffusion Reasoning Models](http://arxiv.org/abs/2606.19257v1)** — Zirui Wu, L. Zheng, J. Ye et al.  
  首个面向长链推理的块扩散语言模型，通过渐进式块尺寸课程学习实现并行解码，推理性能与自回归模型持平。

- **[Rethinking Reward Supervision: Rubric-Conditioned Self-Distillation](http://arxiv.org/abs/2606.19327v1)** — Siyi Gu, J. Chen, S. Zhou et al.  
  提出“评分规则条件自蒸馏”替代昂贵的人工 CoT 标注，用弱监督信号自提升推理模型，大幅降低后训练成本。

- **[Beyond Safe Data: Pretraining-Stage Alignment with Regular Safety Reflection](http://arxiv.org/abs/2606.19168v1)** — Jinhan Li, K. Tang, Y. Xu et al.  
  将安全对齐推进到预训练阶段，通过“安全反思”机制在训练早期注入约束，而非仅靠过滤不安全数据。

- **[Does VLA Even Know the Basics? Measuring Commonsense and World Knowledge Retention in Vision-Language-Action Models](http://arxiv.org/abs/2606.19297v1)** — Nikita Kachaev, A. Moskalenko, M. Skripkin et al.  
  首次系统评估 VLA 模型在机器人微调后遗忘常识与事实知识的程度，揭示多模态模型能力衰减的盲区。

### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

- **[Enhancing Decision-Making with Large Language Models through Multi-Agent Fictitious Play](http://arxiv.org/abs/2606.19308v1)** — Leyang Shen, Y. Zhang, X. Zhao et al.  
  将博弈论中的虚构动态引入多智能体 LLM，使各智能体在迭代中学习最优策略，在复杂决策任务上显著超越传统分工范式。

- **[Data Intelligence Agents: Interpreting, Modeling, and Querying Enterprise Data via Autonomous Coding Agents](http://arxiv.org/abs/2606.19319v1)** — Anoushka Vyas, A. Dhanuka, S. Khoshfetrat Pakazad et al.  
  提出由三个自主编码智能体（解释器、模式生成器、查询器）组成的数据分析系统，自动化企业数据集成，减少人工交接损失。

- **[TxBench-PP: Analyzing AI Agent Performance on Small-Molecule Preclinical Pharmacology](http://arxiv.org/abs/2606.19245v1)** — Hannah Le, R. Ramasamy, A. Urrutia et al.  
  专为药物发现 AI 智能体设计的可验证基准（临床前药理学），覆盖分子属性预测、ADMET 等实际决策任务。

- **[Structured Inference with Large Language Gibbs](http://arxiv.org/abs/2606.19264v1)** — Sanghyeok Choi, H. Gouk, E. S. Whitammer  
  将 LLM 作为吉布斯采样的条件分布，实现对复杂结构化变量（如场景解析、逻辑关系）的概率推理，兼顾知识灵活性与数学严谨性。

### 🔧 方法与框架（新技术、基准测试、效率优化）

- **[A Multi-Domain Benchmark for Detecting AI-Generated Text-Rich Images from GPT-Image-2](http://arxiv.org/abs/2606.19259v1)** — Yijin Wang, S. Wang, W. Zhang et al.  
  首个多领域 AI 生成文本图像检测基准，覆盖票据、海报、截图等高风险场景，推动对抗伪造图像防御。

- **[X+Slides: Benchmarking Audience-Conditioned Slide Generation](http://arxiv.org/abs/2606.19256v1)** — Haodong Chen, X. Zhou, W. Zhou et al.  
  提出面向听众需求的幻灯片生成基准，评估 LLM 能否根据专家 vs 学生等不同受众调整内容和深度。

- **[Explaining Attention with Program Synthesis](http://arxiv.org/abs/2606.19317v1)** — Amiri Hayes, B. Li, J. Andreas  
  用程序合成方法自动为注意力头生成可执行符号描述，首次将注意力机制解释转化为可验证的简单程序。

- **[AGDN: Learning to Solve Traveling Salesman Problem with Anisotropic Graph Diffusion Network](http://arxiv.org/abs/2606.19185v1)** — Bolin Shen, Z. Huang, Z. Cao et al.  
  提出各向异性图扩散网络求解 TSP，通过定向信息传播捕捉图结构局部模式，在规模泛化性上超越现有 GNN 方法。

### 📊 应用（垂直领域、多模态、代码生成）

- **[Trade-offs in Medical LLM Adaptation: An Empirical Study in French QA](http://arxiv.org/abs/2606.19266v1)** — Ikram Belmadani, O. El Khettari, C. Ramisch et al.  
  系统探索法语医疗问答中领域适应与语言适应的权衡，指出单纯微调可能损害语言多样性，为低资源医疗 LLM 提供实用指南。

- **[OneCanvas: 3D Scene Understanding via Panoramic Reprojection](http://arxiv.org/abs/2606.19253v1)** — Bartłomiej Baranowski, D. Z. Chen, M. Nießner  
  用全景重投影将多视图特征汇集到单一全景画布上，使任意 VLM 无需复杂几何编码器即可实现高质量 3D 场景理解。

- **[Forecasting what Matters: Decision-Focused RL for Controlled EV Charging with Unknown Departure Times](http://arxiv.org/abs/2606.19199v1)** — Giuseppe Gabriele, F. Pavirani, S. S. Karimi Madahi et al.  
  将决策聚焦强化学习应用于电动汽车充电调度，在离开时间不确定条件下相比传统方法提升电网稳定性与用户满意度。

## 研究趋势信号

今日投稿中一个值得注意的新兴方向是 **“推理模型的扩散化与遗忘控制”**。DreamReasoner-8B 表明块扩散模型可以胜任链式推理，打破了自回归在推理任务中的垄断。与此同时，MAST 和预训练对齐等工作试图用细粒度的“机制引导”来精准移除或注入特定推理能力，这为模型的可控性打开新思路。另一个趋势是 **“领域智能体评估基准的涌现”**：药物发现（TxBench-PP）、法律（LOCUS）、医学 QA、幻灯片生成等垂直领域的可验证基准密集发布，预示 AI 智能体正在从演示走向可信部署。最后，**评估指标本身被重新审视**（AUC 误导、有效性-区分度权衡），说明社区对评价可靠性的重视上升到一个新高度。

## 值得精读

1. **[DreamReasoner-8B](http://arxiv.org/abs/2606.19257v1)** — 首次将扩散模型成功应用于长链推理，验证了非自回归架构在复杂数学和逻辑任务上的潜力，可能改变未来推理模型的设计范式。

2. **[Explaining Attention with Program Synthesis](http://arxiv.org/abs/2606.19317v1)** — 用可执行程序解释注意力头，跳出了常见的相关性热力图，为深度网络提供真正可验证的符号解释，对可解释性研究具有方法论创新。

3. **[Beyond Safe Data: Pretraining-Stage Alignment with Regular Safety Reflection](http://arxiv.org/abs/2606.19168v1)** — 把安全对齐从微调阶段前移至预训练，不仅思路新颖，而且实验表明能显著减少后训练阶段的安全退火现象，对 LLM 训练流程有直接影响。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*