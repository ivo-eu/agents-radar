# ArXiv AI 研究日报 2026-06-25

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-25 10:25 UTC

---

好的，作为AI研究分析师，以下是根据您提供的2026年6月25日ArXiv论文列表生成的《ArXiv AI研究日报》。

---

### ArXiv AI 研究日报 — 2026-06-25

#### 今日速览

今日投稿揭示了AI安全与对齐研究的深化，从单纯检测“有害行为”转向“模型取证”和“执行时对齐”，并开始审视去中心化AI Agent的信任基础。在推理与智能体方面，多项研究指出了强化学习在复杂任务（如工具使用）中面临的“崩溃”风险，并提出了通过监督信号或价值校准进行稳定训练的方法。此外，关于多模态模型顺序敏感性、语音AI“闻声而不解意”以及AI翻译中的读者偏好等评估性研究，凸显了当前模型在可靠性与细微理解上的关键短板。

---

#### 重点论文

**🧠 大语言模型（架构、训练、对齐、评估）**

1.  **【评估】Same Evidence, Different Answer: Auditing Order Sensitivity in Multimodal Large Language Models**
    - 作者: Paruchuri et al.
    - 链接: http://arxiv.org/abs/2606.26079v1
    - **一句话说明**：引入了Facet-Probe审计工具，揭示了多模态大模型存在严重的“顺序敏感性”，即改变输入内容的顺序可能导致答案改变，挑战了其基本可靠性。

2.  **【评估】SpeechEQ: Benchmarking Emotional Intelligence Quotient in Socially Aware Voice Conversational Models**
    - 作者: Wu et al.
    - 链接: http://arxiv.org/abs/2606.25990v1
    - **一句话说明**：提出SpeechEQ基准，评估语音AI的情商，发现现有领先系统在理解语调和副语言社交线索方面存在严重不足，只是“听见”而非“倾听”。

3.  **【评估】AI translation of literary texts is "fine", but readers still prefer human translations**
    - 作者: Ferstler et al.
    - 链接: http://arxiv.org/abs/2606.26040v1
    - **一句话说明**：通过读者体验实验证实，尽管AI文学翻译在语义上可接受，但读者在沉浸感和文学效果上仍显著偏好人工翻译。

4.  **【训练】On-Policy Self-Distillation with Sampled Demonstrations Reduces Output Diversity**
    - 作者: Nicolicioiu et al.
    - 链接: http://arxiv.org/abs/2606.26091v1
    - **一句话说明**：揭示了on-policy自蒸馏技术在提升模型pass@1准确率的同时，会以牺牲输出多样性（降低pass@k曲线）为代价。

5.  **【对齐】Model Forensics: Investigating Whether Concerning Behavior Reflects Misalignment**
    - 作者: Singh et al.
    - 链接: http://arxiv.org/abs/2606.26071v1
    - **一句话说明**：提出“模型取证”框架，旨在区分模型的有害行为究竟源于“恶意”还是“困惑”，为精准判定模型对齐状态提供了更严谨的方法。

6.  **【预训练】Natural Ungrokking: Asymmetric Control of Which Rules Survive Pretraining**
    - 作者: Li, Sreedhar
    - 链接: http://arxiv.org/abs/2606.26050v1
    - **一句话说明**：观察到语言模型在预训练过程中会“自然遗忘”已学到的简单规则（如代词性别匹配），揭示了预训练规则存活的非对称性，对理解涌现与遗忘机制有启示。

**🤖 智能体与推理（规划、工具使用、多智能体、思维链）**

7.  **【工具使用】Why Multi-Step Tool-Use Reinforcement Learning Collapses and How Supervisory Signals Fix It**
    - 作者: Hao et al.
    - 链接: http://arxiv.org/abs/2606.26027v1
    - **一句话说明**：诊断出将RL应用于多步骤工具使用时模型性能极易“崩溃”的现象，并证明了引入过程监督信号是稳定训练、避免灾难性遗忘的关键。

8.  **【安全】The Unfireable Safety Kernel: Execution-Time AI Alignment for AI Agents and Other Escapable AI Systems**
    - 作者: Dobrin, Chmiel
    - 链接: http://arxiv.org/abs/2606.26057v1
    - **一句话说明**：提出一种在AI Agent运行时环境中执行的、不可绕过的安全内核机制，将安全控制与Agent自身逻辑彻底隔离，以应对Agent越狱或失控的风险。

9.  **【信任】Can Trustless Agents Be Trusted? An Empirical Study of the ERC-8004 Decentralized AI Agent Ecosystem**
    - 作者: Xiong et al.
    - 链接: http://arxiv.org/abs/2606.26028v1
    - **一句话说明**：对首个去中心化AI Agent信任协议ERC-8004进行了实证研究，探讨了在无信任环境中，AI Agent之间的信任建立机制是否有效。

10. **【智能体】Neglected Free Lunch from Post-training: Progress Advantage for LLM Agents**
    - 作者: Oh et al.
    - 链接: http://arxiv.org/abs/2606.26080v1
    - **一句话说明**：提出利用“过程优势”来训练LLM Agent，即通过判断智能体是否在“进步”而非最终结果来提供奖励信号，克服了长程任务中奖励稀疏和标注困难的问题。

**🔧 方法与框架（新技术、基准测试、效率优化）**

11. **【控制理论】Agentic System as Compressor: Quantifying System Intelligence in Bits**
    - 作者: Qin, Zhang
    - 链接: http://arxiv.org/abs/2606.25960v1
    - **一句话说明**：将“智能即压缩”的理论框架应用于AI智能体系统，通过测量系统在给定任务下的“压缩率”来量化其智能水平，提供了一个理论化的评估新视角。

12. **【优化】Tensorion: A Tensor-Aware Generalization of the Muon Optimizer**
    - 作者: Bogachev et al.
    - 链接: http://arxiv.org/abs/2606.25975v1
    - **一句话说明**：提出了Tensorion优化器，它通过感知参数的张量结构（如矩阵的左右奇异向量），对Muon优化器进行泛化，有望进一步加速神经网络训练。

13. **【效率】Improving Neural Network Training by Decoupling the Magnitude and Direction of Weight Vectors**
    - 作者: Hägele et al.
    - 链接: http://arxiv.org/abs/2606.25971v1
    - **一句话说明**：提出将权重向量的“大小”和“方向”进行解耦训练的方法，并设计新型优化器，为改进现有优化算法提供了新思路。

14. **【压缩】Hierarchical Reinforcement Learning for Neural Network Compression (HiReLC): Pruning and Quantization**
    - 作者: Baghdadi et al.
    - 链接: http://arxiv.org/abs/2606.26002v1
    - **一句话说明**：采用分层强化学习框架，自动、联合地搜索最佳的结构剪枝和量化策略，实现了高效的模型压缩。

**📊 应用（垂直领域、多模态、代码生成）**

15. **【机器人】Learning Action Priors for Cross-embodiment Robot Manipulation**
    - 作者: Jing et al.
    - 链接: http://arxiv.org/abs/2606.26095v1
    - **一句话说明**：探索如何让视觉-语言-动作（VLA）模型学习跨不同机器人形态的动作先验，以提升机器人操作的泛化能力，是机器人基础模型的重要一步。

16. **【机器人】FORCE: Efficient VLA Reinforcement Fine-Tuning via Value-Calibrated Warm-up and Self-Distillation**
    - 作者: Zhang et al.
    - 链接: http://arxiv.org/abs/2606.26006v1
    - **一句话说明**：提出FORCE方法，通过价值校准预热和自蒸馏技术，显著提升VLA模型使用强化学习进行微调的样本效率，解决了RL微调中初始性能崩溃问题。

17. **【软件工程】Helpful or Harmful? Evaluating LLM-Assisted Vulnerability Patching via a Human Study**
    - 作者: Biolo et al.
    - 链接: http://arxiv.org/abs/2606.25973v1
    - **一句话说明**：通过人类研究评估LLM辅助漏洞修复的效果，可能发现AI生成的补丁并非总是有益的，为理解人机协作修复软件漏洞的边界提供了实证。

---

#### 研究趋势信号

今日投稿中涌现出**“AI系统的鲁棒性与可靠性多维度审计”**这一重要趋势。研究者不再满足于模型在干净数据上的平均性能，而是开始系统性地审计其在各种“边缘情况”下的表现，例如：多模态输入顺序改变、语音中的副语言信息理解、AI翻译的读者主观体验以及模型内化规则的遗忘与持久性。这表明领域正在从追求“更高分数”转向追求“更可信、更可解释的行为”。同时，**“从行为反推机理”**的研究范式开始出现，如“模型取证”和通过行为实验逆向工程Agent的策略，为理解和诊断模型内部状态提供了新路径。

---

#### 值得精读

1.  **On-Policy Self-Distillation with Sampled Demonstrations Reduces Output Diversity**
    - **精读理由**: 该论文揭示了一个关键trade-off，对当前广泛使用的自蒸馏训练方法提出了警醒。理解并量化pass@1提升与输出多样性下降之间的内在联系，对于平衡模型的准确性与创造性至关重要。

2.  **Natural Ungrokking: Asymmetric Control of Which Rules Survive Pretraining**
    - **精读理由**: “Ungrokking”现象从预训练角度重新构建，挑战了我们对模型知识习得和遗忘的简单理解。它为研究涌现能力、灾难性遗忘以及模型内部表征的动态演化提供了全新的观察视角。

3.  **The Unfireable Safety Kernel: Execution-Time AI Alignment for AI Agents and Other Escapable AI Systems**
    - **精读理由**: 当AI Agent不再仅仅是对话机器人，而是拥有访问API和工具的自主实体时，传统安全方案（如提示工程）的局限性暴露无遗。这篇论文提出的架构性、不可绕过的安全机制，可能是实现可部署、高能力AI Agent安全的关键技术路径。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*