# ArXiv AI 研究日报 2026-06-11

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-06-11 03:32 UTC

---

# 📄 ArXiv AI 研究日报 — 2026-06-11

---

## 🔍 今日速览

今日投稿呈现出几大亮点：**推理时计算分配**与**视觉 Token 高效路由**成为 VLM 优化核心；**可验证环境递归组合**与**多智能体协作**推动 LLM 推理泛化能力；**后训练可解释性**与**稀疏子网络发现**为模型理解和压缩提供新视角；此外，**文本到 SQL 的自动化提示优化**和**生物医学多模态缺失学习**等应用方向也有显著突破。整体研究更强调 *“在受限资源下的有效学习”* 及 *“从数据中自动归纳机理”*。

---

## 📌 重点论文

### 🧠 大语言模型（架构、训练、对齐、评估）

1. **Which Models Are Our Models Built On? Auditing Invisible Dependencies in Modern LLMs**  
   [http://arxiv.org/abs/2606.12385v1](http://arxiv.org/abs/2606.12385v1)  
   *Sanjay Adhikesaven, Haoxiang Sun, Sewon Min et al.*  
   → 首次系统审计现代 LLM 训练中的递归模型依赖关系，揭示“模型套模型”的不可见供应链，对可复现性和安全审计意义重大。

2. **Measurement of Epistemic Resilience of LLMs Under Misleading Medical Context**  
   [http://arxiv.org/abs/2606.12291v1](http://arxiv.org/abs/2606.12291v1)  
   *Hongjian Zhou, Xinyu Zou, Jinge Wu et al.*  
   → 证明即使在医学考试中高分，引入误导性上下文后 LLM 判断力显著下降，直接挑战“高准确率≈安全推理”的假设。

3. **Anatomy of Post-Training: Using Interpretability to Characterize Data and Shape the Learning Signal**  
   [http://arxiv.org/abs/2606.12360v1](http://arxiv.org/abs/2606.12360v1)  
   *Leon Bergen, Usha Bhalla, Sidharth Baskaran et al.*  
   → 用可解释性分解后训练数据对模型行为的真实影响，为精细调节学习信号提供新的分析工具箱。

4. **ALIGNBEAM: Inference-Time Alignment Transfer via Cross-Vocabulary Logit Mixing**  
   [http://arxiv.org/abs/2606.12342v1](http://arxiv.org/abs/2606.12342v1)  
   *Chirag Chawla, Pratinav Seth, Vinay Kumar Sankarapu et al.*  
   → 提出跨词表 logit 混合方法，将安全锚模型的对齐隐式迁移到微调后的领域模型，无需共享词表，解决微调后安全性退化问题。

### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

5. **DIRECT: When and Where Should You Allocate Test-Time Compute in Embodied Planners?**  
   [http://arxiv.org/abs/2606.12402v1](http://arxiv.org/abs/2606.12402v1)  
   *Jadelynn Dao, Milan Ganai, Yasmina Abukhadra et al.*  
   → 分析 VLM 体感规划器中扩展测试时计算的收益递减现象，提出按需分配计算位置和数量的原则，减少延迟和 token 浪费。

6. **Verifiable Environments Are LEGO Bricks: Recursive Composition for Reasoning Generalization**  
   [http://arxiv.org/abs/2606.12373v1](http://arxiv.org/abs/2606.12373v1)  
   *Hao Xiang, Qiaoyu Tang, Le Yu et al.*  
   → 将可验证环境像乐高一样递归组合，大幅扩展推理 RL 训练中的环境多样性，显著提升泛化能力，为“可验证 RL”提供新范式。

7. **APPO: Agentic Procedural Policy Optimization**  
   [http://arxiv.org/abs/2606.12384v1](http://arxiv.org/abs/2606.12384v1)  
   *Xucong Wang, Ziyu Ma, Yong Wang et al.*  
   → 针对工具调用的多轮智能体 RL，提出基于“过程单元”的信用分配方法，比传统回合级或工具边界划分更精细，提升复杂任务成功率。

8. **CHORUS: Decentralized Multi-Embodiment Collaboration with One VLA Policy**  
   [http://arxiv.org/abs/2606.12352v1](http://arxiv.org/abs/2606.12352v1)  
   *Ria Doshi, Tian Gao, Annie Chen et al.*  
   → 单个视觉-语言-动作（VLA）策略即可控制异构机器人群体完成协作任务（如搬沙发），无需中心化状态共享，拓展多机器人泛化边界。

### 🔧 方法与框架（新技术、基准测试、效率优化）

9. **On Subquadratic Architectures: From Applications to Principles**  
   [http://arxiv.org/abs/2606.12364v1](http://arxiv.org/abs/2606.12364v1)  
   *Anamaria-Roberta Hartl, Levente Zólyomi, David Stap et al.*  
   → 系统对比 xLSTM、Mamba 等三类亚二次架构在语言建模、图像、音频任务上的表现，试图提炼有效设计原则，指导下一代架构选择。

10. **Redesign Mixture-of-Experts Routers with Manifold Power Iteration**  
    [http://arxiv.org/abs/2606.12397v1](http://arxiv.org/abs/2606.12397v1)  
    *Songhao Wu, Ang Lv, Ruobing Xie et al.*  
    → 用流形幂迭代重新设计 MoE 路由器，使路由器行向量更好地编码专家矩阵的几何结构，提升专家选择和模型性能。

11. **Fourier Features Let Agents Learn High Precision Policies with Imitation Learning**  
    [http://arxiv.org/abs/2606.12334v1](http://arxiv.org/abs/2606.12334v1)  
    *Balázs Gyenes, Emiliyan Gospodinov, Jan Frieling et al.*  
    → 在模仿学习策略输入中加入傅里叶特征，解决 RGB-only 策略的深度模糊问题，显著提升高精度机器人操作的泛化性。

12. **Finding Sparse Subnetworks in One Training Cycle via Progressive Magnitude-Based Pruning**  
    [http://arxiv.org/abs/2606.12278v1](http://arxiv.org/abs/2606.12278v1)  
    *Romana Qureshi, Hafida Benhidour, Said Kerrache et al.*  
    → 提出单次训练周期内基于渐进幅值剪枝的稀疏子网络发现方法，无需迭代重训练，成本远低于 Lottery Ticket 假设。

### 📊 应用（垂直领域、多模态、代码生成）

13. **TAHOE: Text-to-SQL with Automated Hint Optimization from Experience**  
    [http://arxiv.org/abs/2606.12387v1](http://arxiv.org/abs/2606.12387v1)  
    *Zhiyi Chen, Jie Song, Peng Li et al.*  
    → 将 Text-to-SQL 部署问题转化为自动提示优化问题，通过历史经验学习最佳提示模板，避免昂贵的微调并适配差异化 SQL 方言。

14. **Claw-SWE-Bench: A Benchmark for Evaluating OpenClaw-style Agent Harnesses on Coding Tasks**  
    [http://arxiv.org/abs/2606.12344v1](http://arxiv.org/abs/2606.12344v1)  
    *Mengyu Zheng, Kai Han, Boxun Li et al.*  
    → 专为通用型智能体（如 OpenClaw）设计的编码基准，解决 SWE-bench 对非专用 agent 评分困难的问题，推动自主编码agent 的标准化评测。

15. **Latent World Recovery for Multimodal Learning with Missing Modalities**  
    [http://arxiv.org/abs/2606.12362v1](http://arxiv.org/abs/2606.12362v1)  
    *Hui Wang, Tianyu Ren, Joseph Butler et al.*  
    → 提出“潜在世界恢复”框架，通过模态特异性编码和交叉恢复，解决生物医学等场景中模态缺失下的多模态学习，显著提升预测鲁棒性。

---

## 📈 研究趋势信号

- **“可验证”与“可组合”**：从可验证环境递归组合到动作建议知识共享（CCKS），研究者在推进智能体学习从静态任务集向动态、可扩展的推理泛化迁移。  
- **推理时资源感知**：DIRECT、ALIGNBEAM 等工作关注如何在推理阶段动态分配计算（token、logit）、按需调节开销，反映从“更大模型”到“更智能使用”的范式转变。  
- **可解释性驱动训练**：后训练可解释性分析（Anatomy of Post-Training）与可解释性作为训练信号（ECG分类）出现，表明可解释性正从事后分析工具变为训练过程的核心组件。

---

## ⭐ 值得精读

1. **On Subquadratic Architectures: From Applications to Principles**  
   → 全面对比三种主流亚二次架构，试图提炼通用设计原则。对于关注高效序列建模的研究者，此文提供了量化的性能-效率 trade-off 分析，直接指导架构选型。

2. **Verifiable Environments Are LEGO Bricks: Recursive Composition for Reasoning Generalization**  
   → 提出一个简洁而有影响力的框架——将可验证环境递归组合，大幅扩展训练分布。它用极少的手工设计即可获得推理泛化的大幅提升，是“可验证 RL”方向上极具潜力的工作。

3. **Anatomy of Post-Training: Using Interpretability to Characterize Data and Shape the Learning Signal**  
   → 突破了“后训练是一个黑盒标量优化”的认知，用可解释性工具透视数据对模型行为的具体影响，为精细控制模型行为提供了前所未有的粒度。对对齐和安全研究者尤其重要。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*