# ArXiv AI 研究日报 2026-07-01

> 数据来源: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 共 50 篇论文 | 生成时间: 2026-07-01 11:36 UTC

---

# ArXiv AI 研究日报 | 2026-07-01

## 今日速览

今日投稿聚焦于语言模型的可解释性与自我认知（元认知反馈、自解释忠实性）、智能体在长程任务中的信用分配与技能组合、以及多模态协作基准的构建。值得注意的是，多项工作关注模型在开放世界中的适应性（自适应潜世界模型、零样本导航），并探索高效推理（视觉令牌跳过、无源域适应）。此外，机器人学与强化学习方向出现了从固定演示向自由形式偏好学习和高效VLA模型训练的转变。

---

## 重点论文

### 🧠 大语言模型（架构、训练、对齐、评估）

**Introspective Coupling: Self-Explanation Training Tracks Behavioral Change Despite Fixed Supervision**  
[ArXiv](http://arxiv.org/abs/2606.32038v1)  
*Guo et al.*  
探究LM生成解释是否真正反映内部决策，而非表面模仿。对可解释性研究至关重要。

**Reinforcement Learning with Metacognitive Feedback Elicits Faithful Uncertainty Expression in LLMs**  
[ArXiv](http://arxiv.org/abs/2606.32032v1)  
*Liu et al.*  
通过元认知反馈训练LLM表达真实不确定性，直接缓解过度自信与幻觉问题。

**When LLMs Read Tables Carelessly: Measuring and Reducing Data Referencing Errors**  
[ArXiv](http://arxiv.org/abs/2606.32029v1)  
*Yang et al.*  
系统化测量LLM在表格任务中错误引用数据的问题，提出减少数据引用错误的方案。

**Surrogate Fidelity: When Can Open LLMs Explain Closed Ones?**  
[ArXiv](http://arxiv.org/abs/2606.32008v1)  
*Chlenski et al.*  
研究仅通过开放模型测量能否可靠解释闭源模型内部机制，对可解释性方法学有启发。

**Self-Study Reconsidered: The Hidden Fragility of Learning from Self-Generated QA**  
[ArXiv](http://arxiv.org/abs/2606.32002v1)  
*Alimaskina et al.*  
揭示模型从自身生成QA对中学习存在隐藏脆弱性，可能引入误导性知识。

### 🤖 智能体与推理（规划、工具使用、多智能体、思维链）

**Freeform Preference Learning for Robotic Manipulation**  
[ArXiv](http://arxiv.org/abs/2606.32027v1)  
*Torne et al.*  
提出自由形式偏好学习，使机器人通过自然语言反馈学习奖励，克服二元偏好信号模糊性。

**Generative Skill Composition for LLM Agents**  
[ArXiv](http://arxiv.org/abs/2606.32025v1)  
*Zhao et al.*  
通过生成式技能组合增强LLM智能体处理复杂任务的能力，封装模块化程序知识。

**TRIAGE: Role-Typed Credit Assignment for Agentic Reinforcement Learning**  
[ArXiv](http://arxiv.org/abs/2606.32017v1)  
*Xu et al.*  
在代理强化学习中引入角色类型化信用分配，区分环境动作与内部推理动作的价值。

**AxDafny: Agentic Verified Code Generation in Dafny**  
[ArXiv](http://arxiv.org/abs/2606.32007v1)  
*Breen et al.*  
结合验证器指导修复，实现同时生成可执行代码与形式化证明的智能体代码生成。

**DigitalCoach: Communication and Grounding Gaps in Human and Agentic Computer Use Coaching**  
[ArXiv](http://arxiv.org/abs/2606.31980v1)  
*Chen et al.*  
多模态数据集揭示智能体在教会人类使用软件时存在的沟通与接地问题，推动可教学智能体研究。

**MECoBench: A Systematic Study of Multimodal Agent Collaboration in Embodied Environments**  
[ArXiv](http://arxiv.org/abs/2606.31966v1)  
*Liu et al.*  
提出多模态具身协作基准，系统评估MLLM智能体的视觉接地合作能力。

### 🔧 方法与框架（新技术、基准测试、效率优化）

**AdaJEPA: An Adaptive Latent World Model**  
[ArXiv](http://arxiv.org/abs/2606.32026v1)  
*Wang et al. (Yann LeCun合作)*  
提出自适应潜世界模型，在测试时动态更新预测以应对分布偏移，提升规划鲁棒性。

**Attend, Transform, or Silence: Operator-Level Visual Skipping for Efficient Multimodal LLM Inference**  
[ArXiv](http://arxiv.org/abs/2606.31903v1)  
*Luo et al.*  
在操作符级别有选择地跳过视觉令牌处理，显著加速多模态LLM推理而不损失准确度。

### 📊 应用（垂直领域、多模态、代码生成）

**Z-1: Efficient Reinforcement Learning for Vision-Language-Action Models**  
[ArXiv](http://arxiv.org/abs/2606.31846v1)  
*Cao et al.*  
提出高效的强化学习方法训练视觉-语言-动作模型，使机器人策略超越行为克隆。

**Real-Time Source-Free Object Detection**  
[ArXiv](http://arxiv.org/abs/2606.31834v1)  
*Sairam et al.*  
在严格延迟和内存约束下实现无源域适应目标检测，适用于自动驾驶与监控。

---

## 研究趋势信号

今日论文凸显几个新兴方向：（1）**元认知与自我反思** —— 多篇工作尝试赋予LLM对自身知识边界和不确定性表达的忠实监控，从“生成答案”转向“知道自己知道什么”。（2）**智能体中的信用分配细化** —— 从简单最终奖励走向角色型、动作级精细信用分配，支撑长程自主任务。（3）**多模态协作基准测试** —— 为具身MLLM智能体提供标准化的合作评估环境。（4）**高效适应与推理** —— 自适应潜世界模型、视觉令牌跳过、无源域目标检测等均强调在资源受限或分布偏移场景下的实用性能。

---

## 值得精读

1. **Introspective Coupling** (2606.32038) —— 深入剖析自解释训练是否真正产生忠实内部映射，是理解LM可解释性本质的关键工作。
2. **Reinforcement Learning with Metacognitive Feedback** (2606.32032) —— 将元认知强化学习引入LLM不确定性表达，方法新颖且直接缓解幻觉问题。
3. **AdaJEPA** (2606.32026) —— 突破传统潜世界模型测试时冻结的限制，实现动态自适应，对开放世界规划有长远影响。

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*