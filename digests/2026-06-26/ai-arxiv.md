# ArXiv AI 研究日报 2026-06-26

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-26 10:38 UTC

---

好的，作为AI研究分析师，我将为您呈上今日的《ArXiv AI 研究日报》。

---

### **今日速览**

今日投稿中，**大语言模型（LLM）的安全性、对齐与评估**成为最集中的主题，涌现了关于安全分类、提示注入、有害内容检测的新基准与方法。在**强化学习**方面，无需真实答案的RL框架（RiVER）和用于机器人操作的测试时扩展框架（E-TTS）值得关注。**世界模型与生成模型**领域，幻觉的可预测性与预防机制、以及结合物理信息的概率预测模型取得了新进展。此外，**多模型集成**的性能上限理论分析为组合式AI系统提供了重要洞见。

---

### **重点论文**

#### 🧠 大语言模型（架构、训练、对齐、评估）

1.  **Reinforcement Learning without Ground-Truth Solutions can Improve LLMs**
    [链接](http://arxiv.org/abs/2606.27369v1)
    作者: Y. Lin et al.
    一句话说明: 提出了RiVER框架，通过利用输出序列的相对偏好而非绝对真实答案来进行强化学习，显著拓展了RLVR在LLM训练中的适用范围。

2.  **Paved with True Intents: Intent-Aware Training Improves LLM Safety Classification Across Training Regimes**
    [链接](http://arxiv.org/abs/2606.27210v1)
    作者: J. Ferrao et al.
    一句话说明: 强调了**用户意图**是安全分类的关键信号，并发布了AIMS数据集，证明引入意图建模可显著提升各种训练范式下的安全分类鲁棒性。

3.  **Ask, Don‘t Judge: Binary Questions for Interpretable LLM Evaluation and Self-Improvement**
    [链接](http://arxiv.org/abs/2606.27226v1)
    作者: S. Cho et al.
    一句话说明: 提出BINEVAL框架，将LLM输出评估分解为一系列可解释的二元问题，比传统数值打分更可控、更易调试，可用于模型自我改进。

4.  **When Does Combining Language Models Help? A Co-Failure Ceiling on Routing, Voting, and Mixture-of-Agents Across 67 Frontier Models**
    [链接](http://arxiv.org/abs/2606.27288v1)
    作者: J. Chen
    一句话说明: 通过大规模实验揭示了多模型组合系统（如路由、投票、MoA）的收益被一个“共失效天花板”所限制，为设计更有效的多模型协作策略提供了理论依据。

5.  **When are likely answers right? On Sequence Probability and Correctness in LLMs**
    [链接](http://arxiv.org/abs/2606.27359v1)
    作者: J. Zenn, J. Geiping
    一句话说明: 深入探究了LLM中序列概率与回答正确性之间的关系，揭示了何时“更可能的答案”就是“正确的答案”这一关键假设的适用边界。

#### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

1.  **Empowering GUI Agents via Autonomous Experience Exploration and Hindsight Experience Utilization for Task Planning**
    [链接](http://arxiv.org/abs/2606.27330v1)
    作者: T. Men et al.
    一句话说明: 提出让小模型GUI Agent通过在环境中自主探索和利用事后经验来提升任务规划能力，在低成本下实现了接近商业大模型的性能。

2.  **E-TTS: A New Embodied Test-Time Scaling Framework for Robotic Manipulation**
    [链接](http://arxiv.org/abs/2606.27268v1)
    作者: W. Ye et al.
    一句话说明: 提出了首个面向具身操作的测试时扩展（Test-Time Scaling）框架，通过增加推理时的计算来系统性地提升操作策略的性能。

3.  **Advancing Omnimodal Embodied Agents from Isolated Skills to Everyday Physical Autonomy**
    [链接](http://arxiv.org/abs/2606.27251v1)
    作者: J. Shi et al.
    一句话说明: 设计了一个统一的全模态具身智能体框架，能协同使用API、物理操作和自主恢复机制，在非结构化环境中实现持久的物理自主性。

#### 🔧 方法与框架（新技术、基准测试、效率优化）

1.  **Hallucination in World Models is Predictable and Preventable**
    [链接](http://arxiv.org/abs/2606.27326v1)
    作者: N. Hansen, X. Wang
    一句话说明: 发现生成式世界模型的幻觉集中在**状态-动作空间的低覆盖区域**，提出了一种可预测并主动预防幻觉的方法，首次不依赖固定上限来保证模型可靠性。

2.  **Beyond the Hard Budget: Sparsity Regularizers for More Interpretable Top-k Sparse Autoencoders**
    [链接](http://arxiv.org/abs/2606.27321v1)
    作者: N. Jacquier et al.
    一句话说明: 提出为Top-k稀疏自编码器（SAE）引入软稀疏正则化项，相比硬性预算能学习到更分解、更单语义的特征，提升了模型可解释性。

3.  **Vulnerability of Natural Language Classifiers to Evolutionary Generated Adversarial Text**
    [链接](http://arxiv.org/abs/2606.27215v1)
    作者: M. Singh et al.
    一句话说明: 系统性地探索了利用进化算法生成对抗文本来攻击NLP分类器，揭示了模型在语义相似但具有误导性的文本下的脆弱性。

4.  **How Good Can Linear Models Be for Time-Series Forecasting?**
    [链接](http://arxiv.org/abs/2606.27282v1)
    作者: L. Huang et al.
    一句话说明: 挑战了时序预测中“模型越大越好”的假设，证明通过细致的调优，线性模型可以大幅缩小与大型模型（如Transformer和基础模型）的差距。

#### 📊 应用（垂直领域、多模态、代码生成）

1.  **Prompt Injection in Automated Résumé Screening with Large Language Models: Single and Multi-Injection Settings**
    [链接](http://arxiv.org/abs/2606.27287v1)
    作者: P. Baxi et al.
    一句话说明: 首次系统研究了LLM在自动化简历筛选场景下的**提示注入攻击**，证明了简单的自夸文本即可操纵排名，揭示了AI招聘系统的安全风险。

2.  **HarmVideoBench: Benchmarking Harmful Video Understanding in Large Multimodal Models**
    [链接](http://arxiv.org/abs/2606.27187v1)
    作者: J. Wu et al.
    一句话说明: 发布了首个关注**多层有害特性**的视频理解基准，系统评估大型多模态模型对复杂有害视频（如多模态交织的有害内容）的检测能力。

3.  **Multilingual Reasoning Cascades Need More Context**
    [链接](http://arxiv.org/abs/2606.27306v1)
    作者: A. Mazumder et al.
    一句话说明: 发现多语言推理的“翻译-推理-回译”级联范式存在信息损失，提出需要为后续阶段保留更多原始上下文，以获得更稳定的多语言推理性能。

---

### **研究趋势信号**

今日稿件中，**LLM安全性**的研究趋势显著增强，不再局限于理论探讨，而是出现了大量面向具体应用场景（如招聘、内容审核）的实证研究和攻击分析。同时，**多模型智能体系统**的有效性边界开始被理论性地分析，研究者开始反思简单组合带来的收益上限。**物理世界建模**与AI的结合也出现新动向，如将世界模型幻觉与状态空间覆盖联系起来，以及在地球观测中引入物理信息世界模型，显示出从纯粹的数据驱动向“数据+物理先验”范式演进的趋势。

---

### **值得精读**

1.  **When Does Combining Language Models Help? A Co-Failure Ceiling on Routing, Voting, and Mixture-of-Agents Across 67 Frontier Models**
    - **理由**: 该论文对当前流行的多模型集成策略进行了最全面的理论分析和实证评估。它提出的“共失效天花板”是一个重要发现，对设计未来的Agent系统、路由算法和模型组合方案具有根本性的指导意义。

2.  **Hallucination in World Models is Predictable and Preventable**
    - **理由**: 该研究首次将世界模型中的“幻觉”问题与状态空间的覆盖度联系起来，不仅提供了理论解释，还提出了可行的预防方法。这项工作在生成式模型和机器人/物理模拟领域之间架起了一座重要的桥梁，实用价值极高。

3.  **Reinforcement Learning without Ground-Truth Solutions can Improve LLMs**
    - **理由**: 该工作巧妙地解决了RLVR（带可验证奖励的强化学习）的一个核心痛点：对真实答案的依赖。RiVER框架通过偏好排序来实现强化学习，为在更广泛、更开放的任务中应用RL对齐LLM开辟了新道路，具有重要的范式创新意义。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*