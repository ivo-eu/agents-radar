# ArXiv AI 研究日报 2026-06-23

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-23 10:50 UTC

---

好的，这是根据您提供的 2026 年 6 月 23 日 ArXiv 论文生成的 AI 研究日报。

---

### 📅 ArXiv AI 研究日报 | 2026-06-23

#### 📌 今日速览

今日论文聚焦于 **后训练优化** 与 **具身智能** 两大主线。在基础模型方面，研究开始质疑并试图打破现有架构的固有默认设置，如通过非均匀深度分配（Tapered Language Models）和随机化插值方法（Randomized YaRN）来提升长上下文推理能力。具身智能领域，**灵巧操作** 与 **全身协调** 成为核心主题，工作不仅关注从仿真到现实的规模化数据收集（AutoDex），还试图实现从步行到操作的连续、非间断式控制（CoorDex）。此外，**多智能体系统** 的提示工程（MAS-PromptBench）、**因果发现** 与 LLM 的结合（Causal Discovery in the Era of Agents）以及针对 **评估意识** 的崭新安全视角也值得关注。

---

#### 📑 重点论文

##### 🧠 大语言模型（架构、训练、对齐、评估）

1.  **Randomized YaRN Improves Length Generalization for Long-Context Reasoning**
    -   作者：Manas Mehta 等
    -   [论文链接](http://arxiv.org/abs/2606.23687v1)
    -   **一句话说明**：提出一种名为 Randomized YaRN 的训练方法，通过在位置编码插值过程中引入随机性，显著提升了 LLM 在超长序列上的泛化能力。

2.  **Tapered Language Models**
    -   作者：Reza Bayat 等
    -   [论文链接](http://arxiv.org/abs/2606.23670v1)
    -   **一句话说明**：质疑了当前语言模型“均匀深度”的架构默认设置，提出“锥形”模型，即各层参数非均匀分配，为模型架构设计提供了全新思路。

3.  **SVD-Surgeon: Optimal Singular-Value Surgery for Large Language Model Compression**
    -   作者：Mahmoud Safari, Frank Hutter
    -   [论文链接](http://arxiv.org/abs/2606.23568v1)
    -   **一句话说明**：提出了 SVD-Surgeon，一种通过最优奇异值“手术”对 LLM 进行低秩压缩的方法，旨在最小化压缩带来的性能损失。

4.  **LangMAP: A Language-Adaptive Approach to Tokenization**
    -   作者：Clara Meister 等
    -   [论文链接](http://arxiv.org/abs/2606.23566v1)
    -   **一句话说明**：提出 LangMAP，一种无需从头训练即可将预训练模型适配到新语言特定分词器的框架，解决了多语言扩展的成本问题。

5.  **Evaluation Awareness Is Not One Capability: Evidence from Open Language Models**
    -   作者：Nilesh Nayan 等
    -   [论文链接](http://arxiv.org/abs/2606.23583v1)
    -   **一句话说明**：揭示 LLMs 的“评估意识”并非单一能力，模型在测试环境下的合规表现可导致安全基准测试结果虚高，对部署安全性构成挑战。

##### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

1.  **MAS-PromptBench: When Does Prompt Optimization Improve Multi-Agent LLM Systems?**
    -   作者：Juyang Bai, Laixi Shi
    -   [论文链接](http://arxiv.org/abs/2606.23664v1)
    -   **一句话说明**：构建了一个多智能体提示优化基准，系统性地研究了何时以及如何通过优化系统提示词来提升多 LLM 智能体系统的整体性能。

2.  **EnterpriseClawBench: Benchmarking Agents from Real Workplace Sessions**
    -   作者：Jincheng Zhong 等
    -   [论文链接](http://arxiv.org/abs/2606.23654v1)
    -   **一句话说明**：从真实的、企业级智能体工作会话中构建了一个全新的基准，用于评估智能体在处理复杂、多步骤企业任务时的能力。

3.  **Teaching LLMs String Matching, Backtracking, and Error Recovery to Deduce Bases and Truth Tables for the Combinatorially Exploding Bit Manipulation Puzzles**
    -   作者：Prateek Agnihotri 等
    -   [论文链接](http://arxiv.org/abs/2606.23672v1)
    -   **一句话说明**：展示了如何通过教授 LLMs 字符串匹配、回溯和错误恢复等算法技能，来解决高度复杂的比特位操作推理难题。

4.  **SPIRAL: Learning to Search and Aggregate**
    -   作者：Jubayer Ibn Hamid 等
    -   [论文链接](http://arxiv.org/abs/2606.23595v1)
    -   **一句话说明**：提出 SPIRAL 框架，使语言模型能够学习在测试时自适应地结合多种推理方式（如逐步推理、并行采样和结果聚合），以提升推理能力。

##### 🔧 方法与框架（新技术、基准测试、效率优化）

1.  **Muown Implicitly Performs Angular Step-size Decay**
    -   作者：Florian Hübler 等
    -   [论文链接](http://arxiv.org/abs/2606.23637v1)
    -   **一句话说明**：理论分析了高效优化器 Muown，证明它隐式地实现了角步长衰减机制，解释了其在 Transformer 预训练中的出色表现。

2.  **Scaling Linear Mode Connectivity and Merging to Billion Parameter Pretrained Transformers**
    -   作者：Tianyi Li, Zhiqiang Shen
    -   [论文链接](http://arxiv.org/abs/2606.23607v1)
    -   **一句话说明**：成功地将线性模式连通性（LMC）和模型合并技术扩展到十亿参数级别的预训练 Transformer，为理解和组合大模型提供了新的可能。

3.  **VeriEvol: Scaling Multimodal Mathematical Reasoning via Verifiable Evol-Instruct**
    -   作者：Haoling Li 等
    -   [论文链接](http://arxiv.org/abs/2606.23543v1)
    -   **一句话说明**：提出 VeriEvol 框架，通过可验证的进化指令自动生成高质量的、带有可靠奖励标签的多模态数学推理数据，以支撑大规模的强化学习训练。

4.  **Scheduling Thoughts: Learning the Order of Thought in Diffusion Language Models**
    -   作者：Jiawei Xu 等
    -   [论文链接](http://arxiv.org/abs/2606.23567v1)
    -   **一句话说明**：针对扩散语言模型，提出学习一种“思考顺序”而非使用启发式顺序，通过优化去噪过程中 token 的生成顺序来提升生成质量。

##### 📊 应用（垂直领域、多模态、代码生成）

1.  **AutoDex: An Automated Real-World System for Dexterous Grasping Data Collection**
    -   作者：Mingi Choi 等
    -   [论文链接](http://arxiv.org/abs/2606.23689v1)
    -   **一句话说明**：提出 AutoDex，一个全自动的真实世界灵巧抓取数据收集系统，旨在解决该领域大规模真实数据稀缺的瓶颈。

2.  **CoorDex: Coordinating Body and Hand Priors for Continuous Dexterous Humanoid Loco-Manipulation**
    -   作者：Sikai Li 等
    -   [论文链接](http://arxiv.org/abs/2606.23680v1)
    -   **一句话说明**：提出 CoorDex，一种能够协调身体和手部先验知识的学习框架，实现了人形机器人从行走到操作的连续、非间断式灵巧操作。

3.  **Hedgementation = Hedgerow Segmentation: A Remote Sensing Benchmark**
    -   作者：Nathan Senyard 等
    -   [论文链接](http://arxiv.org/abs/2606.23615v1)
    -   **一句话说明**：发布了一个全新的遥感基准数据集 Hedgementation，专注于评估机器学习模型在国家尺度、10米分辨率下对树篱进行语义分割的能力。

4.  **DiT-Reward: Generative Representations for Text-to-Image Reward Modeling**
    -   作者：Yuanming Yang 等
    -   [论文链接](http://arxiv.org/abs/2606.23626v1)
    -   **一句话说明**：探索了将用于图像生成的表示（来自 Diffusion Transformer）直接用于评估生成图像质量的奖励建模，开辟了文本到图像评估的新范式。

---

#### 📈 研究趋势信号

今日投稿中一个显著的趋势是 **对现有范式“默认设置”的挑战**。例如，Tapered Language Models 质疑了“均匀深度”的架构默认值，Randomized YaRN 改进了位置编码的插值方式，而 On the Limits of Prompt-Conditioned Language Models 则直接挑战了“提示词是万能接口”的假设。这种对“祖传代码”的反思，预示着社区正从追求“更大”转向追求“更合理”，基础研究正进入一个微调和重构阶段。另一个信号是 **“评估”本身成为研究对象**，如研究 LLM 的“评估意识”和风险评估，这暗示着安全性与评估方法学将成为一个独立且重要的研究分支。

---

#### ⭐ 值得精读

1.  **On the Limits of Prompt-Conditioned Language Models as General-Purpose Learners**
    -   [论文链接](http://arxiv.org/abs/2606.23668v1)
    -   **精读理由**：本文从根本上反思了将 LLM 视为通用求解器的流行观点，分析了语言作为任务接口的固有限制。这对于理解 LLM 的能力边界、设计更有效的交互方式具有深刻的哲学和实用价值，是整个领域的“思想实验”。

2.  **Causal Discovery in the Era of Agents**
    -   [论文链接](http://arxiv.org/abs/2606.23608v1)
    -   **精读理由**：本文系统性地探讨了如何将 LLM 智能体与因果发现相结合，并指出了现有方法的缺陷。它不仅是技术综述，更提供了一个关键的“警示”视角，对于 AI for Science 和追求可解释性的研究至关重要。

3.  **Against Proxy Optimization**
    -   [论文链接](http://arxiv.org/abs/2606.23597v1)
    -   **精读理由**：这是一篇极具冲击力的理论文章，讨论了在何种条件下最大化代理效用函数是有害的。它直接触及了 AI 对齐问题的核心，即如何避免“奖励模型黑客”和“目标错位”，对致力于安全 AGI 的研究者来说是必读之作。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*