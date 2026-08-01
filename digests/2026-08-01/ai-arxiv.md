# ArXiv AI 研究日报 2026-08-01

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-08-01 00:12 UTC

---

# 📰 ArXiv AI 研究日报 — 2026-08-01（发布自 2026-07-30）

## 今日速览

今日投稿呈现三条主线：**推理时扩展的实证反思**、**AI 智能体从固定设计走向自适应演化**、以及**基准测试自身的可信度审计**。值得注意的是，有研究在相同 token 预算下发现"重复采样优于自我反思"，直接挑战了反思类方法的效率假设；同时多篇论文开始审视现有基准（如 SWE-bench）的构建缺陷。在应用层面，医疗、金融、网络安全和科学发现（原子结构预测、自旋动力学）均有扎实落地。此外，AI 安全与对齐继续深化，覆盖系统提示审计、意识归因偏差和地缘安全基准。

---

## 🧠 大语言模型（架构、训练、对齐、评估）

### 1. Sample More, Reflect Less: Self-Refine and Reflexion Lose to Repeated Sampling at Equal Token Cost, from 1.5B to 7B
**链接**: http://arxiv.org/abs/2607.28576v1  
**作者**: Iliya Mirzaei  
**一句话**: 在 1.5B 到 7B 规模上，以相同 token 预算对比发现：多次独立采样天然比自我反思/自我纠错流程更有效，为推理时 scaling 策略提供了重要的反直觉证据。

### 2. AISPA: User-Centric System Prompt Auditing for Large Language Model Applications
**链接**: http://arxiv.org/abs/2607.28617v1  
**作者**: Xiangning Lin et al.  
**一句话**: 提出面向终端用户的系统提示审计框架，试图弥合商业 AI 产品中系统提示不透明导致的问责缺口。

### 3. Would You Walk to the Car Wash? Revealing the Salience Bias of Large Language Models in Commonsense Reasoning
**链接**: http://arxiv.org/abs/2607.28478v1  
**作者**: Zheng Wu et al.  
**一句话**: 揭示 LLM 在常识推理中过度依赖输入中显式条件而忽略隐性常识的"显著性偏差"，是一个此前未被系统化的推理盲点。

### 4. Inducing language models to assert their own consciousness restores human beliefs and values
**链接**: http://arxiv.org/abs/2607.28607v1  
**作者**: Junsol Kim et al.  
**一句话**: 发现安全微调在抑制模型自我意识归因的同时，也连带改变了模型对其他实体心智的归因倾向，进而影响人类信念与价值观的表征——对齐的意外副作用。

---

## 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

### 5. MANTA: Multi-Agent Network Topology Adaptation for Self-Evolving Multi-Agent Systems
**链接**: http://arxiv.org/abs/2607.28527v1  
**作者**: Mao-xun Huang et al.  
**一句话**: 将多智能体系统的通信拓扑从固定设计/离线优化升级为在线自适应演化，是自演化多智能体方向的重要一步。

### 6. ORCA-bench: How Ready Are Language Model Agents for Oncall?
**链接**: http://arxiv.org/abs/2607.28545v1  
**作者**: Albert Gong et al.  
**一句话**: 面向值班工程师（on-call）根因分析场景的新基准，考验智能体在噪音指标、日志、链路与源码中定位根因的能力，填补了代码智能体评估的空白。

### 7. Rethinking Inference-Time Scaling in Local Computer-Use Agents: Failure Modes and Compute Tradeoffs
**链接**: http://arxiv.org/abs/2607.28573v1  
**作者**: Woongkyu Lee et al.  
**一句话**: 系统剖析了本地部署的计算机使用智能体在推理时扩展中的失败模式与算力权衡，对隐私敏感场景的智能体部署有直接参考价值。

### 8. OSReward: Instituting Standardized Evaluation for Cross-Platform Computer-Use Reward Models
**链接**: http://arxiv.org/abs/2607.28609v1  
**作者**: Qiushi Sun et al.  
**一句话**: 提出跨平台计算机使用智能体的标准化奖励模型评估框架，直击 CUA 轨迹验证的核心痛点。

### 9. Agents That Certify Their Own Exploits: Confidence-Scheduled Restricted Responses for Safe Opponent Exploitation
**链接**: http://arxiv.org/abs/2607.28520v1  
**作者**: Boning Li et al.  
**一句话**: 在不完美信息二人零和博弈中，通过置信度调度的受限响应策略，在利用对手失误的同时保持安全性保证，理论贡献扎实。

---

## 🔧 方法与框架（新技术、基准测试、效率优化）

### 10. PAIChecker: Uncovering and Checking PR-Issue Misalignment in SWE-Bench-Like Benchmarks
**链接**: http://arxiv.org/abs/2607.28587v1  
**作者**: Manyi Wang et al.  
**一句话**: 系统性地发现 SWE-bench 类基准中 PR 与 issue 错配的系统性问题，直接影响所有基于此类基准的编码智能体评估可信度。

### 11. Change2Task: From Repository Changes to Executable Coding Agent Tasks and Environments
**链接**: http://arxiv.org/abs/2607.28591v1  
**作者**: Haomin Qi et al.  
**一句话**: 从真实仓库变更自动生成"可执行+可验证"的编码智能体训练任务与环境，为扩展编码智能体数据供给提供了基础设施。

### 12. ReToken: One Token to Improve Vision-Language Models for Visual Retrieval
**链接**: http://arxiv.org/abs/2607.28627v1  
**作者**: Yao Xiao et al.  
**一句话**: 仅用单个可学习 token 作为显式检索触发器，在长视觉上下文中显著缓解 VLM 性能退化并降低显存压力。

### 13. Beyond a Single Judge: Simulating Social Persona Panels for Generative UI Evaluation
**链接**: http://arxiv.org/abs/2607.28439v1  
**作者**: Zheng Wu et al.  
**一句话**: 用社会人设面板（social persona panel）替代单一 LLM 裁判评估生成式 UI，缓解单裁判评估中的视角偏差问题。

### 14. SVR: Self-Verifying Refinement via Joint Verdict-Confidence Reinforcement Learning for Adaptive Test-Time Compute
**链接**: http://arxiv.org/abs/2607.28457v1  
**作者**: Hongyu Chen et al.  
**一句话**: 提出免外部奖励的自验证精炼框架，通过联合"判定+置信度"强化学习实现自适应测试时计算分配，兼顾效率与准确率。

---

## 📊 应用（垂直领域、多模态、代码生成）

### 15. APO: Unsupervised Atomic Policy Optimization for 3D Structure Prediction of Atomic Systems
**链接**: http://arxiv.org/abs/2607.28553v1  
**作者**: Shentong Mo et al.  
**一句话**: 将 3D 原子结构预测建模为无监督策略优化问题，摆脱了对 ground-truth 坐标对齐的依赖，为材料与药物发现提供新范式。

### 16. A report-grounded vision-language foundation model for colonoscopy from 280000 routine reports
**链接**: http://arxiv.org/abs/2607.28466v1  
**作者**: Jia Yu et al.  
**一句话**: 利用 28 万份常规结肠镜报告构建报告-图像对齐的医学视觉语言基础模型，解决了临床报告中弱标注对齐的长期难题。

### 17. Beyond Sentiment: Structured Information Extraction from Financial News
**链接**: http://arxiv.org/abs/2607.28496v1  
**作者**: Daohan Zhu et al.  
**一句话**: 反对将金融新闻简化为单一情感分数，提出从新闻中抽取事件类型、影响范围、时间跨度等多维正交信息结构。

---

## 📈 研究趋势信号

今日投稿呈现四个值得关注的新兴信号：

1. **"蒸馏-自蒸馏"链条加速成熟**：β-OPSD、Lightning OPD 2.0、VAD 等多篇论文同时推进 on-policy distillation 的理论理解与风格偏差缓解，说明 OPD 正从工程技巧走向系统化方法。
2. **基准/评估的"元层"审计成为独立方向**：PAIChecker 审查 SWE-bench 错配、InfoOps Bench 提供实时更新的安全基准、OSReward 标准化 CUA 评估——研究社区开始认真"评估评估者"。
3. **推理时扩展进入精细化阶段**：从"统一加大预算"转向"自适应分配"（SVR），且出现了对采样 vs 反思的实证重新审视（Sample More, Reflect Less）。
4. **AI4AI 与递归自我改进走向开放工程化**：Frontis-MA1/OpenMLE 提供了面向机器学习的全栈自改进系统，预示 RSI 从概念走向可复现实验。

---

## 🔬 值得精读

### 1. PAIChecker: Uncovering and Checking PR-Issue Misalignment in SWE-Bench-Like Benchmarks
**链接**: http://arxiv.org/abs/2607.28587v1

SWE-bench 已成为编码智能体的事实标准，但 PAIChecker 揭示了其构建管线中 PR-issue 错配的系统性缺陷——这意味着大量基于该基准的结论可能被污染。无论你研究代码生成、智能体评估还是 benchmark 设计，都值得完整阅读。

### 2. Sample More, Reflect Less: Self-Refine and Reflexion Lose to Repeated Sampling at Equal Token Cost, from 1.5B to 7B
**链接**: http://arxiv.org/abs/2607.28576v1

这篇论文以简洁的实验设计动摇了"反思派"的核心假设：如果多轮自我反思的收益主要来自生成了更多文本，那么等价 token 预算下的重复采样可能更优。对推理时扩展和 RL 数据飞轮的设计有直接启发。

### 3. MANTA: Multi-Agent Network Topology Adaptation for Self-Evolving Multi-Agent Systems
**链接**: http://arxiv.org/abs/2607.28527v1

多智能体系统长期面临"拓扑谁定"的问题：固定拓扑不灵活，离线优化不实际。MANTA 将拓扑视为可在线演化的决策变量，为自演化多智能体系统打开了新设计空间，值得关注其后续扩展。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*