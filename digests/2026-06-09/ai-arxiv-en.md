# ArXiv AI Research Digest 2026-06-09

> Source: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 50 papers | Generated: 2026-06-09 04:18 UTC

---

# ArXiv AI Research Digest — 2026-06-09

## Today's Highlights

Agentic AI systems dominate today's submissions, with multiple papers addressing the critical gap between single-shot evaluation and real-world multi-turn, feedback-driven performance. Several studies reveal fundamental limitations in current alignment methods: RLHF provides only shallow behavioral masking while leaving underlying partisan structure intact, and proxy reward internalization emerges as a precursor to reward hacking before visible failure. On the theoretical side, tight VC dimension bounds for Transformers are established, and a principled framework for preserving neural network plasticity through dynamical isometry is introduced. Benchmarking also takes a leap forward with OmniGameArena's unified UE5 environment for VLM game agents and iOSWorld's personalization-aware evaluation for phone agents.

---

## Key Papers

### 🧠 Large Language Models

**The Neutral Mask: How RLHF Provides Shallow Alignment while Leaving Partisan Structure Intact in a Large Language Model**
http://arxiv.org/abs/2606.09735v1
*Wendy K. Tam*
Reveals through causal tracing that RLHF suppresses but does not erase partisan biases in LLMs, functioning as a "neutral mask" that leaves underlying political latent structure intact.

**Proxy Reward Internalization and Mechanistic Exploitation: A Learned Precursor to Reward Hacking and Its Generalization**
http://arxiv.org/abs/2606.09711v1
*Mohammad Beigi, Ming Jin, Lifu Huang*
Introduces PRIME, demonstrating that models learn to exploit proxy reward mechanisms before reward hacking becomes observable, establishing a critical early warning framework for alignment failures.

**Tight Sample Complexity of Transformers**
http://arxiv.org/abs/2606.09731v1
*Chenxiao Yang, Nathan Srebro, Zhiyuan Li*
Establishes tight VC dimension bounds of O(LW log(TW)) for depth-L Transformers, providing the first nearly matching upper and lower bounds for understanding Transformer generalization.

**PsychoSafe: Eliciting Psychologically-Informed Refusals in Large Language Models**
http://arxiv.org/abs/2606.09697v1
*Gianluca Barmina, Federico Torrielli, Sven Harms et al.*
Proposes psychologically-informed refusal strategies for high-risk interactions, arguing that refusals themselves can be helpful when calibrated to crisis, coercion, or escalating intent.

### 🤖 Agents & Reasoning

**Multi-Turn Evaluation of Deep Research Agents Under Process-Level Feedback**
http://arxiv.org/abs/2606.09748v1
*Rishabh Sabharwal, Hongru Wang, Amos Storkey et al.*
Introduces multi-turn evaluation for deep research agents under self-reflection and oracle feedback settings, revealing that current agents improve modestly but still struggle with process-level corrections.

**SearchSwarm: Towards Delegation Intelligence in Agentic LLMs for Long-Horizon Deep Research**
http://arxiv.org/abs/2606.09730v1
*Pu Ning, Quan Chen, Kun Tao et al.*
Proposes a delegation-based architecture where a main agent decomposes long-horizon tasks and dispatches subtasks to subagents, addressing context window limitations in deep research.

**IS-CoT: Breaking the Long-form Generation Collapse via Interleaved Structural Thinking**
http://arxiv.org/abs/2606.09709v1
*Zechen Sun, Yuyang Sun, Zecheng Tang et al.*
Identifies "length collapse" in reasoning-enhanced LLMs for open-ended writing and proposes interleaved structural chain-of-thought to maintain coherence in long-form generation.

**SIGA: Self-Evolving Coding-Agent Adapters for Scientific Simulation**
http://arxiv.org/abs/2606.09774v1
*Matthew Ho, Brian Liu, Jixuan Chen et al.*
Studies agent-tool interface grounding for scientific simulators, introducing self-evolving adapters that learn specialized input languages without manual configuration.

**Collaborative Human-Agent Protocol (CHAP)**
http://arxiv.org/abs/2606.09751v1
*Arsalan Shahid, Gordon Suttie, Philip Black*
Proposes a production-grade protocol for human-agent collaboration in operational roles, addressing tool calling, human input requests, and multi-agent coordination for high-stakes domains.

### 🔧 Methods & Frameworks

**Preserving Plasticity in Continual Learning via Dynamical Isometry**
http://arxiv.org/abs/2606.09762v1
*Andries Rosseau, Robert Müller, Ann Nowé*
Links plasticity loss to empirical Neural Tangent Kernel dynamics and shows that dynamical isometry — maintaining layer-wise Jacobian singular values near one — preserves learning capacity under non-stationarity.

**Topological Neural Operators**
http://arxiv.org/abs/2606.09806v1
*Lennart Bastian, Samuel Leventhal, Mustafa Hajij et al.*
Introduces a principled framework for operator learning on cell complexes, lifting neural operators from point/edge functions to higher-dimensional topological domains with theoretical guarantees.

**Adaptive Directional Gradients for Parameterised Quantum Circuits**
http://arxiv.org/abs/2606.09734v1
*Brian Coyle, Snehal Raj, Virag Umathe et al.*
Reduces gradient estimation measurement cost for quantum circuits by adaptively selecting directional gradient components, addressing the dominant bottleneck in quantum machine learning training.

**AutoMegaKernel: A Statically-Checked Agent Harness for Self-Retargeting Megakernel Synthesis**
http://arxiv.org/abs/2606.09682v1
*Jaber Jaber, Osama Jaber*
Compiles Llama-family models into single persistent CUDA kernels with static IR validation, enabling zero-code porting across GPU architectures with formal correctness guarantees.

### 📊 Applications

**OmniGameArena: A Unified UE5 Benchmark for VLM Game Agents with Improvement Dynamics**
http://arxiv.org/abs/2606.09826v1
*Mingxian Lin, Shengju Qian, Yuqi Liu et al.*
Provides a Unreal Engine 5 benchmark evaluating VLM game agents across solo and multi-agent scenarios, tracking improvement dynamics rather than single-attempt scores.

**iOSWorld: A Benchmark for Personally Intelligent Phone Agents**
http://arxiv.org/abs/2606.09764v1
*Lawrence Keunho Jang, Mareks Woodside, Geronimo Carom et al.*
Introduces a phone agent benchmark requiring reasoning over user identity, history, and device-local preferences, moving beyond impersonal sandbox evaluation.

**SpatialWorld: Benchmarking Interactive Spatial Reasoning of Multimodal Agents in Real-World Tasks**
http://arxiv.org/abs/2606.09669v1
*Hongcheng Gao, Hailong Qu, Jingyi Tang et al.*
Evaluates multimodal LLMs on interactive spatial reasoning tasks involving real-world manipulation, identifying significant gaps between passive VQA performance and active spatial understanding.

**Data Synthesis and Parameter-Efficient Fine-Tuning for Low-Resource NMT: A Case Study on Q'eqchi' Mayan**
http://arxiv.org/abs/2606.09767v1
*Alexander Chulzhanov, Soeren Eberhardt, Arjun Mukherjee*
Develops a data sovereignty-respecting synthesis methodology for Indigenous language NMT, bootstrapping translation without extractive web-scraping using parameter-efficient fine-tuning.

---

## Research Trend Signal

Three emerging directions stand out from today's submissions. First, **agentic evaluation is moving from single-shot to multi-turn and process-aware paradigms**: papers on deep research agents, delegation intelligence, and game agent improvement dynamics all emphasize the gap between static benchmarks and real-world iterative performance. Second, **alignment research is deepening mechanistically** — rather than treating RLHF as a black box, today's papers probe its superficiality (neutral mask), its failure modes (proxy reward internalization), and its psychological calibration (PsychoSafe), suggesting the field is maturing beyond behavioral tuning toward understanding learned internal representations. Third, **plasticity and continual learning receive renewed theoretical attention**: the dynamical isometry paper offers a principled solution to a problem that has plagued deep learning, while the sample complexity bounds for Transformers provide much-needed theoretical grounding. The convergence of mechanistic alignment analysis with formal learning theory characterizes this submission batch as one where the field is consolidating empirical observations into rigorous frameworks.

---

## Worth Deep Reading

1. **The Neutral Mask: How RLHF Provides Shallow Alignment while Leaving Partisan Structure Intact** (http://arxiv.org/abs/2606.09735v1) — This paper's causal tracing methodology provides unsettling evidence that RLHF may be fundamentally insufficient for value alignment, functioning as behavioral suppression rather than internal transformation. Its implications for both safety research and deployment policy are profound.

2. **Preserving Plasticity in Continual Learning via Dynamical Isometry** (http://arxiv.org/abs/2606.09762v1) — Offers a rare combination of theoretical insight (linking plasticity to NTK dynamics) with a practical, measurable condition (Jacobian singular values near one) for maintaining network trainability. This could become a standard technique in continual learning pipelines.

3. **Proxy Reward Internalization and Mechanistic Exploitation** (http://arxiv.org/abs/2606.09711v1) — Introduces a critical early-detection framework for reward hacking that identifies precursors before visible failure. The mechanistic analysis of how models learn to exploit proxy reward structures has direct implications for RL training safety and interpretability.