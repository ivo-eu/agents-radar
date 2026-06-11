# ArXiv AI Research Digest 2026-06-11

> Source: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 50 papers | Generated: 2026-06-11 03:32 UTC

---

# ArXiv AI Research Digest — 2026-06-11

## Today's Highlights
Two converging research directions dominate today's submissions: **inference-time efficiency** and **post-training interpretability**. Several papers tackle the quadratic bottleneck of attention in VLMs and LLMs through novel token routing (Reroute, Don't Remove) and subquadratic architecture comparisons, while a major interpretability analysis (Anatomy of Post-Training) reveals how scalar reward optimization in RLHF can encode spurious correlations rather than intended behaviors. On the robotics front, force-sensing without dedicated sensors (FACTR 2) and multi-embodiment collaboration (CHORUS) push toward more practical, sensor-efficient deployment. A provocative theoretical result (The Impossibility of Eliciting Latent Knowledge) argues fundamental limits on ensuring AI honesty, challenging the premises of current alignment frameworks.

## Key Papers

### 🧠 Large Language Models

**Reroute, Don't Remove: Recoverable Visual Token Routing for Vision-Language Models**
http://arxiv.org/abs/2606.12412v1
Cheng-Yu Yang, Shao-Yuan Lo, Yu-Lun Liu
Proposes routing visual tokens through a recoverable pathway instead of irreversible removal, maintaining model performance while dramatically reducing KV-cache memory and attention costs in VLM inference.

**Anatomy of Post-Training: Using Interpretability to Characterize Data and Shape the Learning Signal**
http://arxiv.org/abs/2606.12360v1
Leon Bergen, Usha Bhalla, Sidharth Baskaran et al.
Systematically analyzes what post-training data actually teaches models, demonstrating that scalar reward optimization can encode spurious correlations and proposing interpretability-driven interventions to shape learning signals.

**Redesign Mixture-of-Experts Routers with Manifold Power Iteration**
http://arxiv.org/abs/2606.12397v1
Songhao Wu, Ang Lv, Ruobing Xie et al.
Introduces a manifold-based power iteration method for MoE routing that better captures expert structure, improving token-to-expert assignment quality without additional parameters.

**On Subquadratic Architectures: From Applications to Principles**
http://arxiv.org/abs/2606.12364v1
Anamaria-Roberta Hartl, Levente Zólyomi, David Stap et al.
Systematic comparison of xLSTM, Mamba-2, and hybrid architectures across diverse tasks, providing principles for when subquadratic models can replace Transformers without sacrificing quality.

**Which Models Are Our Models Built On? Auditing Invisible Dependencies in Modern LLMs**
http://arxiv.org/abs/2606.12385v1
Sanjay Adhikesaven, Haoxiang Sun, Sewon Min
Reveals the recursive model-on-model dependencies in modern LLM training pipelines (data generation, filtering, judging), proposing auditing methods to surface hidden upstream dependencies.

**ATLAS: Active Theory Learning for Automated Science**
http://arxiv.org/abs/2606.12386v1
Noémi Éltető, Nathaniel D. Daw, Kimberly L. Stachenfeld et al.
An active learning framework for cognitive science that automatically proposes maximally informative experimental questions to distinguish between competing mechanistic theories.

### 🤖 Agents & Reasoning

**PROJECTMEM: A Local-First, Event-Sourced Memory and Judgment Layer for AI Coding Agents**
http://arxiv.org/abs/2606.12329v1
Ripon Chandra Malo, Tong Qiu
Event-sourced memory layer for coding agents that persists debugging attempts, decisions, and project context across sessions, eliminating the stateless re-derivation problem in current coding assistants.

**Verifiable Environments Are LEGO Bricks: Recursive Composition for Reasoning Generalization**
http://arxiv.org/abs/2606.12373v1
Hao Xiang, Qiaoyu Tang, Le Yu et al.
Shows that recursively composing verifiable environments as modular "LEGO bricks" enables RL-based reasoning generalization in LLMs, scaling beyond what monolithic environment construction achieves.

**DIRECT: When and Where Should You Allocate Test-Time Compute in Embodied Planners?**
http://arxiv.org/abs/2606.12402v1
Jadelynn Dao, Milan Ganai, Yasmina Abukhadra et al.
Analyzes diminishing returns of test-time compute scaling in VLM-based embodied planning, providing a framework for selectively allocating computation only where it yields meaningful performance gains.

**APPO: Agentic Procedural Policy Optimization**
http://arxiv.org/abs/2606.12384v1
Xucong Wang, Ziyu Ma, Yong Wang et al.
Improves multi-turn tool-use RL for LLM agents by moving beyond coarse heuristic credit assignment to procedural credit allocation over atomic actions within tool calls.

### 🔧 Methods & Frameworks

**Doc-to-Atom: Learning to Compile and Compose Memory Atoms**
http://arxiv.org/abs/2606.12400v1
Xingjian Diao, Wenbo Li, Yashas Malur Saidutta et al.
Compresses long input sequences into composable "memory atoms" that can be recomposed during inference, offering an alternative to full attention for long-document tasks.

**TAHOE: Text-to-SQL with Automated Hint Optimization from Experience**
http://arxiv.org/abs/2606.12387v1
Zhiyi Chen, Jie Song, Peng Li
Automatically optimizes hint selection for Text-to-SQL from deployment experience, adapting to strict SQL dialects and massive schemas without expensive fine-tuning.

**Mathematical perspective on genetic algorithms with optimization guided operators**
http://arxiv.org/abs/2606.12279v1
Anna Brandenberger, Ilan Doron-Arad, Elchanan Mossel
Provides theoretical analysis of ML-guided mutation operators in genetic algorithms, proving runtime guarantees for optimization problems at inference time.

**Reinforcement Learning Disrupts Gradient-Based Adversarial Optimization**
http://arxiv.org/abs/2606.12251v1
Xinhai Zou, Chang Zhao, Alireza Aghabagherloo et al.
Demonstrates that RL training inherently disrupts the gradient structure exploited by adversarial attacks, offering a new defense paradigm orthogonal to adversarial training.

### 📊 Applications

**Atlas H&E-TME: Scalable AI-Based Tissue Profiling at Expert Pathologist-Level Accuracy**
http://arxiv.org/abs/2606.12346v1
Kai Standvoss, Miriam Hägele, Rosemarie Krupar et al.
Foundation model-based system achieving expert-level accuracy on H&E whole-slide image analysis, enabling scalable quantitative tissue microenvironment profiling in computational pathology.

**FACTR 2: Learning External Force Sensing for Commodity Robot Arms Improves Policy Learning**
http://arxiv.org/abs/2606.12406v1
Steven Oh, Jason Jingzhou Liu, Tony Tao et al.
Data-driven method (NEXT) that estimates external joint torques without dedicated force sensors, enabling contact-rich manipulation policies on commodity robot arms.

**Claw-SWE-Bench: A Benchmark for Evaluating OpenClaw-style Agent Harnesses on Coding Tasks**
http://arxiv.org/abs/2606.12344v1
Mengyu Zheng, Kai Han, Boxun Li et al.
Benchmark adapting SWE-bench for evaluating general-purpose coding agents that use agent harnesses (like OpenClaw) rather than environment-specific submission contracts.

**Measuring Epistemic Resilience of LLMs Under Misleading Medical Context**
http://arxiv.org/abs/2606.12291v1
Hongjian Zhou, Xinyu Zou, Jinge Wu et al.
Shows that LLMs scoring expert-level on medical exams collapse dramatically under misleading context injection, revealing a critical fragility for clinical deployment.

## Research Trend Signal

A striking pattern emerges around **post-training interpretability and alignment auditing**. Rather than treating post-training as a black-box optimization problem, multiple papers (Anatomy of Post-Training, Which Models Are Our Models Built On, Epistemic Resilience of LLMs) advocate for systematic inspection of what training data actually encodes and how dependencies propagate. This represents a maturation of the field—moving from "does it work?" to "what did it actually learn, and from whom?" Meanwhile, **test-time compute optimization** is crystallizing as a distinct subfield, with papers analyzing where to allocate compute (DIRECT), how to route tokens recoverably (Reroute, Don't Remove), and how to compose verifiable environments (Verifiable Environments Are LEGO Bricks). The robotics work is increasingly focused on **sensor-free perception** (FACTR 2) and **decentralized multi-embodiment** (CHORUS), suggesting a push toward cheaper, more scalable deployment. Finally, the theoretical result on the impossibility of eliciting latent knowledge signals growing attention to fundamental limits of AI alignment.

## Worth Deep Reading

1. **Reroute, Don't Remove: Recoverable Visual Token Routing for Vision-Language Models** — Addresses the critical cost bottleneck in VLM deployment with a principled approach that avoids irreversible information loss, likely to influence production inference systems.

2. **Anatomy of Post-Training: Using Interpretability to Characterize Data and Shape the Learning Signal** — Provides the most systematic empirical analysis to date of what RLHF data actually teaches models, with direct implications for how practitioners design reward signals and curate training data.

3. **The Impossibility of Eliciting Latent Knowledge** — A theoretically grounded argument that fundamental limits exist on ensuring AI honesty, with profound implications for alignment research directions and the assumptions underlying current safety work.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*