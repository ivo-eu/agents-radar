# ArXiv AI 研究日报 2026-07-29

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-07-29 00:10 UTC

---

# ArXiv AI 研究日报 | 2026-07-29

---

## 今日速览

今日投稿亮点纷呈：**Kimi K3** 以 2.8T 参数 MoE 架构宣告开放前沿智能，其 1M 上下文窗口与 Delta Attention 设计值得重注；**多轮长周期规划** 方面，一篇论文通过「单/多教师同策略智能体蒸馏」首次系统揭示了规划能力的获取与整合机制；**证据归因与可解释性** 成为多篇论文聚焦点——从医疗视觉语言模型的解释对齐到无坐标的证据定位，再到 SAE 特征与功能的新几何解释，表明社区正从“模型能做什么”转向“模型如何做以及为何可信”。此外，**数据策展** 的粒化、**极端天气预警** 的端到端 LLM 智能体、以及 **幻觉检测** 的隐藏状态谱信号也展现出鲜明的新方向。

---

## 重点论文

### 🧠 大语言模型（架构、训练、对齐、评估）

1. **Kimi K3: Open Frontier Intelligence**  
   [链接](http://arxiv.org/abs/2607.24653v1)  
   作者: Kimi Team et al.  
   **一句话**：2.8T 参数 MoE（104B 激活）、原生视觉与 1M 上下文窗口，凭借 Delta Attention 和 Attention Residuals 实现深度信息流，代表开放的下一代通用模型。

2. **DataOrchestra: Learning to Orchestrate Per-Example Curation of Pretraining Data**  
   [链接](http://arxiv.org/abs/2607.24717v1)  
   作者: Zhen Huang et al.  
   **一句话**：首个逐样本自适应数据策展方法，学习为每个训练例子分配最优处理策略，超越固定域级流水线。

3. **D-Score: A Spectral Hidden-State Signal for Hallucination Detection in Large Language Models**  
   [链接](http://arxiv.org/abs/2607.24586v1)  
   作者: Bianca Raimondi et al.  
   **一句话**：利用隐藏激活的谱几何（D-Score）检测幻觉，无需外部知识库，简单且高效。

4. **MMOE: Modernizing Diffusion Transformers with Efficient Expert Design**  
   [链接](http://arxiv.org/abs/2607.24665v1)  
   作者: Yanhao Jia et al.  
   **一句话**：为扩散 Transformer 引入稀疏 MoE 机制，在保持容量增长同时控制推理与部署成本。

5. **Eviction as Estimation: A Fixed-Lag Smoothing View of Test-Time Memory**  
   [链接](http://arxiv.org/abs/2607.24667v1)  
   作者: Maruthi Vemula et al.  
   **一句话**：将 LLM 工作记忆的逐出问题重新定义为隐藏状态估计，提出比传统累积策略更优的测量式方法。

### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

6. **The Physics of Multi-Turn Long-Horizon Planning: From Pre-training to Post-training via Single- and Multi-Teacher On-Policy Agentic Distillation**  
   [链接](http://arxiv.org/abs/2607.24720v1)  
   作者: Tianyi Men et al.  
   **一句话**：通过可控合成数据与多教师蒸馏，首次系统解构了基础模型多轮规划能力的习得过程，并提出“规划物理”概念。

7. **Reason-Mediated Behavioral Models for Auditing LLM Social Simulators**  
   [链接](http://arxiv.org/abs/2607.24649v1)  
   作者: Atharva Pandey et al.  
   **一句话**：指出仅匹配人群答案不够，需用理由中介的行为模型审计 LLM 社会模拟器，防止“正确错误理由”。

8. **SIREN: Towards End-to-End Extreme-Weather Early Warning with Experience-Grounded LLM Agents**  
   [链接](http://arxiv.org/abs/2607.24588v1)  
   作者: Hang Ni et al.  
   **一句话**：将 LLM 智能体与气象数据融合，实现从预警到行动的全流程自动化，降低人工专家成本。

### 🔧 方法与框架（新技术、基准测试、效率优化）

9. **ClinFusion: A Vision-Centric Multimodal LLM System for Holistic Medical Understanding**  
   [链接](http://arxiv.org/abs/2607.24743v1)  
   作者: Hangjie Yuan et al.  
   **一句话**：多模态医疗 LLM 框架，核心在于异构 2D/3D 医学图像的知识吸收与评估对齐，极具临床潜力。

10. **Evidence Attribution in Visual Document Understanding without Coordinates or Region Labels**  
    [链接](http://arxiv.org/abs/2607.24651v1)  
    作者: Zhuchenyang Liu et al.  
    **一句话**：无需坐标或区域标签，通过隐式推理对答案进行证据归因，推动文档理解的可信度评估。

11. **PIVOT: Efficient Query-Group Indexing for Token-Level Sparse Attention**  
    [链接](http://arxiv.org/abs/2607.24593v1)  
    作者: Hong Liu et al.  
    **一句话**：针对 DeepSeek 稀疏注意力的索引器瓶颈，提出查询组索引，将 token 级 top-k 选择效率大幅提升。

12. **LOCKS: Page-Local Compact Key Summaries for Efficient Long-Context Decoding**  
    [链接](http://arxiv.org/abs/2607.24555v1)  
    作者: Junsung Hwang  
    **一句话**：利用 KV 缓存页面的局部低秩性构造紧凑摘要，显著降低长上下文解码的 KV 读取瓶颈。

13. **CADER: Confidence-Aware Dynamic Evidence Reasoning for Long-Video Understanding**  
    [链接](http://arxiv.org/abs/2607.24582v1)  
    作者: Jinlong Yang et al.  
    **一句话**：根据问题难度动态调用工具辅助推理，避免对所有示例一刀切，在长视频理解中实现效率与精度平衡。

### 📊 应用（垂直领域、多模态、代码生成）

14. **A corrective agentic hybrid RAG and an operations-grounded evaluation for a scientific facility**  
    [链接](http://arxiv.org/abs/2607.24663v1)  
    作者: Rajat Sainju et al.  
    **一句话**：为先进光子源构建 agentic RAG 系统，整合电子日志、文档、聊天记录等异构数据，并提供基于运行真实性的评估框架。

15. **Trust in AI-Assisted Code Review**  
    [链接](http://arxiv.org/abs/2607.24601v1)  
    作者: Zhenhan Gao et al.  
    **一句话**：系统评估可解释 AI 在代码审查中对开发者信任的影响，为 LLM 辅助开发工具的透明性设计提供实证。

16. **ERUnderstand: Evaluating Vision-Language Models on Structured ER Diagrams**  
    [链接](http://arxiv.org/abs/2607.24707v1)  
    作者: Ali Ansari et al.  
    **一句话**：首个结构化 ER 图理解基准，填补了数据库设计中图像到机器可读模式转换的研究空白。

---

## 研究趋势信号

- **证据归因与可解释性的几何化**：SAE 特征与因果效应的下游几何分析（#25）、无坐标文档证据归因（#22）、以及针对医疗场景的解释对齐（#5）共同指向一个趋势——不再满足于激活层面的解释，而是追求因果与空间层面的可审计解释。
- **数据与训练的“粒化”**：从固定域级数据策展走向逐样本自适应（#8），以及针对长尾分布的特殊对齐（#49），表明预训练数据工程正从“大水漫灌”走向“精准滴灌”。
- **智能体系统的可信与安全**：权限策略代数（#26）、理由中介的社会模拟审计（#23）、以及极端天气预警中的端到端经验嵌入（#33），显示智能体研究开始系统处理正确性、安全性与社会可信度。

---

## 值得精读

1. **Kimi K3 (Open Frontier Intelligence)**  
   **理由**：2.8T MoE 与 Delta Attention 的实战验证，开放权重且提供详细技术报告，是当前最接近“下一代开源基础模型”的标杆，架构细节与训练策略对后续百亿级模型有直接参考价值。

2. **The Physics of Multi-Turn Long-Horizon Planning**  
   **理由**：首次提出“规划物理”视角，通过合成数据控制与环境蒸馏分离预训练与后训练阶段的规划能力，方法论原创性强，对理解智能体能力的本质来源至关重要。

3. **Evidence Attribution in Visual Document Understanding without Coordinates or Region Labels**  
   **理由**：摒弃传统的坐标输出范式，通过隐式推理进行证据归因，不仅可迁移性强，而且更贴近人类审查方式，为文档理解系统的可信度评估开辟新路线。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*