# ArXiv AI 研究日报 2026-06-18

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-18 03:18 UTC

---

# ArXiv AI 研究日报（2026-06-18）

## 今日速览

今日投稿聚焦于**推理模型的扩展与对齐**——块扩散语言模型首次被验证可用于长链推理（DreamReasoner-8B），GRPO的熵崩溃问题获得理论分析与改进（STARE）。**神经符号学习**迎来可微分类语义框架（NeSyCat Torch），为统一多种逻辑范式提供基础。**模型安全与遗忘**方面，选择性遗忘RLVR推理行为的方法（MAST）以及预训练阶段对齐（Beyond Safe Data）被提出。此外，**物理增强图神经网络**（P-K-GCN）在高保真超分辨率模拟中展现潜力，而**多智能体虚构博弈**（Multi-Agent Fictitious Play）提升了LLM策略决策能力。

## 重点论文

### 🧠 大语言模型（架构、训练、对齐、评估）

1. **DreamReasoner-8B: Block-Size Curriculum Learning for Diffusion Reasoning Models**  
   [http://arxiv.org/abs/2606.19257v1](http://arxiv.org/abs/2606.19257v1)  
   *Zirui Wu, Lin Zheng, Jiacheng Ye et al.*  
   **开源块扩散推理模型**，通过分块大小课程学习，首次证明块扩散LLM可扩展至长链推理任务，为高效推理提供新范式。

2. **Rethinking Reward Supervision: Rubric-Conditioned Self-Distillation**  
   [http://arxiv.org/abs/2606.19327v1](http://arxiv.org/abs/2606.19327v1)  
   *Siyi Gu, Jialin Chen, Sophia Zhou et al.*  
   **规则条件自蒸馏**：用可得的评分标准（rubric）替代昂贵且可能有噪声的链式思维标注，通过蒸馏提升推理模型性能，减少对人工标注的依赖。

3. **STARE: Surprisal-Guided Token-Level Advantage Reweighting for Policy Entropy Stability**  
   [http://arxiv.org/abs/2606.19236v1](http://arxiv.org/abs/2606.19236v1)  
   *Haipeng Luo, Qingfeng Sun, Songli Wu et al.*  
   **解决GRPO策略熵崩溃**：通过惊奇度（surprisal）引导的Token级优势重加权，稳定策略熵，提升RLVR训练中LLM的推理多样性。

4. **MAST: Mechanism-Guided Selective Unlearning for RLVR-Induced Reasoning**  
   [http://arxiv.org/abs/2606.19222v1](http://arxiv.org/abs/2606.19222v1)  
   *Chenyu Zhou, Qiliang Jiang, Shuning Wu et al.*  
   **机制对齐的选择性遗忘**：针对RLVR诱导的推理能力，仅微调少量参数即可遗忘特定行为，大幅降低对原有能力的附带损害。

5. **Beyond Safe Data: Pretraining-Stage Alignment with Regular Safety Reflection**  
   [http://arxiv.org/abs/2606.19168v1](http://arxiv.org/abs/2606.19168v1)  
   *Jinhan Li, Kexian Tang, Yihan Xu et al.*  
   **预训练阶段安全对齐**：超越过滤不安全数据，引入定期安全反思机制，使LLM在预训练阶段就建立深层安全认知。

### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

6. **Enhancing Decision-Making with Large Language Models through Multi-Agent Fictitious Play**  
   [http://arxiv.org/abs/2606.19308v1](http://arxiv.org/abs/2606.19308v1)  
   *Leyang Shen, Yang Zhang, Xiaoyan Zhao et al.*  
   **多智能体虚构博弈**：将博弈论中的虚构博弈引入LLM多智能体系统，提升在战略决策任务（如谈判、资源分配）中的表现，弥补传统分治范式的不足。

7. **Data Intelligence Agents: Interpreting, Modeling, and Querying Enterprise Data via Autonomous Coding Agents**  
   [http://arxiv.org/abs/2606.19319v1](http://arxiv.org/abs/2606.19319v1)  
   *Anoushka Vyas, Aarushi Dhanuka, Sina Khoshfetrat Pakazad et al.*  
   **数据智能体系统**：三个智能体（解释器、模式创建器、查询器）自主协作，实现企业数据的端到端集成与查询，减少人工重复工作。

### 🔧 方法与框架（新技术、基准测试、效率优化）

8. **NeSyCat Torch: A Differentiable Tensor Implementation of Categorical Semantics for Neurosymbolic Learning**  
   [http://arxiv.org/abs/2606.19279v1](http://arxiv.org/abs/2606.19279v1)  
   *Daniel Romero Schellhorn, Till Mossakowski, Björn Gehrke et al.*  
   **可微范畴语义框架**：将经典、模糊、概率、神经符号等多种语义统一在单一可微的张量实现下，为神经符号学习提供标准工具。

9. **Explaining Attention with Program Synthesis**  
   [http://arxiv.org/abs/2606.19317v1](http://arxiv.org/abs/2606.19317v1)  
   *Amiri Hayes, Belinda Li, Jacob Andreas*  
   **用程序合成解释注意力**：将注意力头近似为可执行程序，首次用符号化描述替代神经网络黑盒，为可解释性提供新路径。

10. **Essential Subspace Merging for Multi-Task Learning**  
    [http://arxiv.org/abs/2606.19164v1](http://arxiv.org/abs/2606.19164v1)  
    *Longhua Li, Lei Qi, Xin Geng et al.*  
    **关键子空间合并**：分析任务特定参数更新的输出偏移，提出仅在必要子空间合并权重，减少多任务干扰，提升模型合并效果。

### 📊 应用（垂直领域、多模态、代码生成）

11. **Native Active Perception as Reasoning for Omni-Modal Understanding**  
    [http://arxiv.org/abs/2606.19341v1](http://arxiv.org/abs/2606.19341v1)  
    *Zhenghao Xing, Ruiyang Xu, Yuxuan Wang et al.*  
    **主动感知推理**：变被动“全看”为主动选择关键帧，显著降低长视频理解的计算成本，同时保持高准确率。

12. **OneCanvas: 3D Scene Understanding via Panoramic Reprojection**  
    [http://arxiv.org/abs/2606.19253v1](http://arxiv.org/abs/2606.19253v1)  
    *Bartłomiej Baranowski, Dave Zhenyu Chen, Matthias Nießner*  
    **全景重投影3D场景理解**：无需复杂几何编码器，将多视角特征聚合到单一全景画布，实现高效且泛化性强的3D VLM。

13. **P-K-GCN: Physics-augmented Koopman-enhanced Graph Convolutional Network for Deep Spatiotemporal Super-resolution**  
    [http://arxiv.org/abs/2606.19303v1](http://arxiv.org/abs/2606.19303v1)  
    *Xizhuo Zhang, Zekai Wang et al.*  
    **物理增强图网络超分辨率**：结合Koopman算子与图卷积，将物理约束融入数据驱动模型，在流体、气候等时空动力学超分辨率任务上显著优于传统方法。

14. **TxBench-PP: Analyzing AI Agent Performance on Small-Molecule Preclinical Pharmacology**  
    [http://arxiv.org/abs/2606.19245v1](http://arxiv.org/abs/2606.19245v1)  
    *Hannah Le, Ramesh Ramasamy, Alex Urrutia et al.*  
    **AI药物发现基准**：首个针对小分子临床前药理学可验证基准，评估AI智能体在高保真药物研发决策中的表现，推动可信AI部署。

## 研究趋势信号

从今日投稿中观察到两个新兴方向：**扩散模型向推理领域拓展**——继图像生成后，块扩散语言模型开始挑战长链推理（DreamReasoner-8B），有望带来推理加速与新型范式；**模型合并与遗忘的“手术刀”式技术**——从关键子空间合并（Essential Subspace Merging）到机制对齐选择性遗忘（MAST），研究者正尝试在保持模型原始能力的前提下进行精准的增删操作，这对大模型持续学习与安全治理至关重要。

## 值得精读

1. **Explaining Attention with Program Synthesis**  
   首次将注意力头翻译为可读程序，为黑箱神经网络提供了真正可理解的符号解释，是AI可解释性领域的重要突破。

2. **NeSyCat Torch**  
   统一了多种神经符号逻辑语义的微分实现，为融合符号推理与深度学习提供基础设施，有望成为该领域标准库。

3. **DreamReasoner-8B**  
   公开了8B规模块扩散推理模型，验证了扩散架构在长链推理上的可行性，可激励更多基于扩散的推理模型研究。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*