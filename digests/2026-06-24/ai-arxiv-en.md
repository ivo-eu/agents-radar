# ArXiv AI Research Digest 2026-06-24

> Source: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 50 papers | Generated: 2026-06-24 10:35 UTC

---

# ArXiv AI Research Digest — 2026-06-24

## Today's Highlights

A strong cluster of papers advances **agentic systems** toward autonomous skill acquisition and structured reasoning, with InSight (paper 1) enabling VLAs to learn new primitives without human intervention and OpenThoughts-Agent (paper 4) providing open training recipes for broadly capable agents. **Code repair and repository-level reasoning** sees a significant breakthrough with SHERLOC (paper 16), which moves from file-level retrieval to actionable diagnostic localization. On the generative side, FLUX3D (paper 3) and OrbitForge (paper 18) push text-to-3D quality through diffusion-aligned sparse representations and video-anchored reconstruction respectively. Finally, theoretical results on **plasticity loss in LLMs** (paper 30) and **scaling laws for distillation** (paper 31) provide crucial guidance for efficient model training and deployment.

---

## Key Papers

### 🧠 Large Language Models

**Can Scale Save Us From Plasticity Loss in Large Language Models?**  
http://arxiv.org/abs/2606.24752v1 — J. Fernando Hernandez-Garcia, Tomás Figliolia, Beren Millidge  
Systematically investigates whether increasing model scale mitigates the loss of plasticity (ability to learn new information after prior training), finding that scale alone is insufficient — a critical result for continual learning in deployed LLMs.

**Scaling Laws for Task-Specific LLM Distillation**  
http://arxiv.org/abs/2606.24747v1 — Lavinia Ghita, Dhruv Desai, Ioana Boier  
Derives empirical scaling laws for domain-specific LLM compression, quantifying the trade-off between student model size, distillation data volume, and downstream task performance — essential for cost-effective deployment.

**Grad Detect: Gradient-Based Hallucination Detection in LLMs**  
http://arxiv.org/abs/2606.24790v1 — Anand Kamat, Daniel Blake, Brent M. Werness  
Proposes a gradient-based method for detecting hallucinations in LLM outputs without requiring external verification, offering a lightweight and practical solution for high-stakes applications.

---

### 🤖 Agents & Reasoning

**InSight: Self-Guided Skill Acquisition via Steerable VLAs**  
http://arxiv.org/abs/2606.24884v1 — Maggie Wang, Lars Osterberg, Stephen Tian et al.  
Unlocks autonomous skill acquisition for vision-language-action models by making them steerable at the primitive-action level, enabling robots to learn new manipulation skills beyond the training data without human demonstration.

**OpenThoughts-Agent: Data Recipes for Agentic Models**  
http://arxiv.org/abs/2606.24855v1 — Negin Raoof, Richard Zhuang, Marianna Nezhurina et al.  
Provides open-source recipes for curating training data for broadly capable agentic models, moving beyond single-benchmark approaches toward general-purpose agent behavior.

**World Models in Pieces: Structural Certification for General Agents**  
http://arxiv.org/abs/2606.24842v1 — Yikai Lu, Yifei Wu, Xinyu Lu et al.  
Formalizes that agents in the "big-world regime" must have specialized world models, and introduces structural certification that distinguishes understanding of critical bottlenecks from irrelevant failures.

**SHERLOC: Structured Diagnostic Localization for Code Repair Agents**  
http://arxiv.org/abs/2606.24820v1 — Hovhannes Tamoyan, Sean Narenthiran, Erik Arakelyan et al.  
Moves code repository localization from file-level retrieval to actionable diagnostic explanations, enabling LLM agents to identify *what* is wrong and *why* rather than just *where* to look.

**Are We Ready For An Agent-Native Memory System?**  
http://arxiv.org/abs/2606.24775v1 — Wei Zhou, Xuanhe Zhou, Shaokun Han et al.  
Argues that LLM agent memory has evolved into a full data management problem, proposing requirements for persistent storage, retrieval, update, consolidation, and lifecycle governance.

**LaGO: Latent Action Guidance for Online Reinforcement Learning**  
http://arxiv.org/abs/2606.24669v1 — Kuan-Yen Liu, Ren-Jyun Huang, Ti-Rong Wu  
Uses LLMs to provide latent action guidance for online RL rather than direct control, achieving more reliable sequential decision-making through guidance in a latent space.

---

### 🔧 Methods & Frameworks

**Real vs. Complex Spectral Bases for Neural Operators: The Role of Green's Function Alignment**  
http://arxiv.org/abs/2606.24851v1 — Jason Sulskis, Sathya Ravi  
Introduces the Hartley Neural Operator, a real-valued alternative to Fourier Neural Operators that eliminates redundant conjugate symmetry while preserving or improving PDE solution operator learning.

**Posterior Refinement: Fast Language Generation via Any-Order Flow Maps**  
http://arxiv.org/abs/2606.24773v1 — Manan Agarwal, Sheel Shah, Chanhyuk Lee et al.  
Proposes a non-autoregressive generation framework that overcomes the limitations of masked diffusion models, enabling iterative refinement through any-order token editing with improved efficiency.

**Less is More: Quality-Aware Training Data Selection for Scientific Summarization**  
http://arxiv.org/abs/2606.24828v1 — Maria Nefeli Paraskevopoulou, Tatiana Passali, Grigorios Tsoumakas  
Demonstrates that selective training on high-quality author-written abstracts improves scientific summarization performance more than using all available data — a practical finding for data-constrained domains.

**IV-CoT: Implicit Visual Chain-of-Thought for Structure-Aware Text-to-Image Generation**  
http://arxiv.org/abs/2606.24849v1 — Zixuan Li, Haokun Lin, Yicheng Xiao et al.  
Addresses structure-aware prompt following in text-to-image generation by introducing an implicit visual chain-of-thought that reasons about object counts, spatial relations, and layouts before generation.

---

### 📊 Applications

**FLUX3D: High-Fidelity 3D Gaussian Generation with Diffusion-Aligned Sparse Representation**  
http://arxiv.org/abs/2606.24874v1 — Haorui Ji, Weizhe Liu, Hongdong Li et al.  
Achieves high-fidelity image-to-3D Gaussian Splatting generation by aligning diffusion features with sparse voxel representations, preserving high-frequency visual details that previous methods lose.

**OrbitForge: Text-to-3D Scene Generation via Reconstruction-Anchored Video Synthesis**  
http://arxiv.org/abs/2606.24799v1 — Chenrui Fan, Paolo Favaro  
Leverages text-to-video models as rich scene priors for 3D generation, using reconstruction anchoring to overcome camera control and temporal inconsistency issues in video-to-3D pipelines.

**Solving Inverse Problems of Chaotic Systems with Bidirectional Conditional Flow Matching**  
http://arxiv.org/abs/2606.24824v1 — Peiyan Hu, Jian Zhang, Jiashu Pan et al.  
Tackles the notoriously difficult problem of inferring initial conditions from final states in chaotic systems using bidirectional conditional flow matching, with implications for weather prediction and physical modeling.

**Large-Language-Model Discovery of Quantum LDPC Codes through Structured Concept Evolution**  
http://arxiv.org/abs/2606.24808v1 — Zidu Liu, Florian Marquardt  
Uses LLMs to discover novel quantum low-density parity-check codes by evolving structured concepts, potentially accelerating the path to scalable quantum error correction.

**UniDrive: A Unified Vision-Language and Grounding Framework for Interpretable Risk Understanding in Autonomous Driving**  
http://arxiv.org/abs/2606.24759v1 — Xiaowei Gao, Pengxiang Li, Yitai Cheng et al.  
Bridges the trade-off between temporal reasoning and spatial precision in autonomous driving MLLMs by unifying vision-language understanding with fine-grained grounding.

**The Warrant Gap: Claim-Conditioned Re-scoring for Fact-Checking**  
http://arxiv.org/abs/2606.24627v1 — Arka Ujjal Dey, John Collomosse  
Identifies and addresses the "warrant gap" where LLM fact-checkers output correct verdicts but with cited evidence that does not license the claim, proposing claim-conditioned re-scoring for verifiable reasoning.

---

## Research Trend Signal

A clear emerging direction this week is the **integration of structured, compositional reasoning into generative models** across modalities. Rather than treating generation as a monolithic process, multiple papers decompose it: IV-CoT introduces implicit visual chain-of-thought for text-to-image (paper 7), CineCap structures video captioning around film-language concepts (paper 48), and SHERLOC moves code repair from file retrieval to diagnostic explanations (paper 16). This trend intersects with **agent-native architectures** that treat memory, data, and reasoning as first-class systems rather than add-ons (papers 8, 25). Simultaneously, the community is grappling with **fundamental limits of scaling** — plasticity loss (paper 30), distillation trade-offs (paper 31), and the need for quality over quantity in training data (paper 13) — suggesting a maturing field that increasingly values principled design over brute-force scaling.

---

## Worth Deep Reading

1. **InSight: Self-Guided Skill Acquisition via Steerable VLAs** (paper 1)  
   *Reasoning*: Directly tackles the fundamental limitation of imitation learning — that agents cannot exceed their demonstration data. InSight's approach of making VLAs steerable at the primitive-action level could be a practical path toward open-ended robotic skill acquisition without costly human annotation.

2. **SHERLOC: Structured Diagnostic Localization for Code Repair Agents** (paper 16)  
   *Reasoning*: Elevates code repair from a retrieval problem to a diagnostic reasoning problem. The move from file-level to actionable explanations is a concrete step toward making LLM agents genuinely useful for software maintenance, and the methodology is likely transferable to other tool-use domains.

3. **Solving Inverse Problems of Chaotic Systems with Bidirectional Conditional Flow Matching** (paper 15)  
   *Reasoning*: Addresses a long-standing open problem (chaotic inverse dynamics) with a novel generative approach. The bidirectional flow matching formulation is elegant and has clear high-impact applications in climate science, astrophysics, and engineering where predicting past states from current observations is crucial.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*