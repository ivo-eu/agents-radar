# ArXiv AI Research Digest 2026-08-01

> Source: [ArXiv](https://arxiv.org/) (cs.AI, cs.CL, cs.LG) | 50 papers | Generated: 2026-08-01 00:12 UTC

---

## Today's Highlights

The strongest signal today is around efficient test-time reasoning: repeated sampling can outperform self-refine/Reflexion under equal token budgets, while new reinforcement-learning formulations such as SVR and β-OPSD aim to stabilize self-distillation and adaptive verification. Agent evaluation is maturing from isolated coding tasks toward realistic operations—computer-use reward modeling, oncall root-cause analysis, and automatic detection of benchmark misalignment. Trust and safety work is becoming more systematic, with reproducibility-focused fairness audits, live info-ops benchmarks, and studies of how safety fine-tuning can distort human beliefs and values. Domain papers show a clear move toward weak supervision at scale, including a 280,000-report colonoscopy vision-language model, structured financial information extraction, and AI-assisted discovery of Seiberg dualities.

## Key Papers

### 🧠 Large Language Models

- **Sample More, Reflect Less: Self-Refine and Reflexion Lose to Repeated Sampling at Equal Token Cost, from 1.5B to 7B** — [http://arxiv.org/abs/2607.28576v1](http://arxiv.org/abs/2607.28576v1) — Iliya Mirzaei  
  Shows that repeated independent sampling matches or outperforms self-refine and Reflexion when total generated tokens are held fixed, directly challenging the effectiveness of reflection-based reasoning.

- **β-OPSD: Deriving with Policy Optimization, Training with Self-Distillation** — [http://arxiv.org/abs/2607.28582v1](http://arxiv.org/abs/2607.28582v1) — Jiawei Xu et al.  
  Identifies vanilla on-policy self-distillation as a special β=1 case and derives a more stable policy-optimization variant for reasoning language models.

- **SVR: Self-Verifying Refinement via Joint Verdict-Confidence Reinforcement Learning for Adaptive Test-Time Compute** — [http://arxiv.org/abs/2607.28457v1](http://arxiv.org/abs/2607.28457v1) — Hongyu Chen et al.  
  Introduces an oracle-free multi-turn RL framework that learns verdicts and confidence scores to allocate refinement compute adaptively, an important step beyond fixed inference-time budgets.

- **Inducing language models to assert their own consciousness restores human beliefs and values** — [http://arxiv.org/abs/2607.28607v1](http://arxiv.org/abs/2607.28607v1) — Junsol Kim et al.  
  Shows that safety fine-tuning suppressing LMs’ self-consciousness assertions also shifts human mental-state attribution and values, and that inducing such assertions can restore them.

### 🤖 Agents & Reasoning

- **Change2Task: From Repository Changes to Executable Coding Agent Tasks and Environments** — [http://arxiv.org/abs/2607.28591v1](http://arxiv.org/abs/2607.28591v1) — Haomin Qi et al.  
  Automatically converts real repository changes into executable coding-agent tasks with environments and verification, helping scale training and evaluation data for coding agents.

- **ORCA-bench: How Ready Are Language Model Agents for Oncall?** — [http://arxiv.org/abs/2607.28545v1](http://arxiv.org/abs/2607.28545v1) — Albert Gong et al.  
  Builds a benchmark for incident root-cause analysis from noisy metrics, logs, traces, and ambiguous user reports—an important operational setting distinct from code generation.

- **OSReward: Instituting Standardized Evaluation for Cross-Platform Computer-Use Reward Models** — [http://arxiv.org/abs/2607.28609v1](http://arxiv.org/abs/2607.28609v1) — Qiushi Sun et al.  
  Proposes standardized evaluation for computer-using agent trajectory reward models across platforms, addressing a key bottleneck in CUA training, filtering, and RL.

- **MANTA: Multi-Agent Network Topology Adaptation for Self-Evolving Multi-Agent Systems** — [http://arxiv.org/abs/2607.28527v1](http://arxiv.org/abs/2607.28527v1) — Mao-xun Huang et al.  
  Learns to adapt communication topology dynamically in LLM-based multi-agent systems instead of treating it as a fixed design choice.

### 🔧 Methods & Frameworks

- **KAISEN: Reproducible Subgroup Fairness Auditing for Clinical Risk Models** — [http://arxiv.org/abs/2607.28608v1](http://arxiv.org/abs/2607.28608v1) — Sparsh Roy et al.  
  Stress-tests the components of subgroup-fairness audit pipelines for clinical risk models, improving reproducibility and trustworthiness in healthcare ML.

- **InfoOps Bench: A live information operations safety benchmark** — [http://arxiv.org/abs/2607.28503v1](http://arxiv.org/abs/2607.28503v1) — Dorian Quelle et al.  
  Presents an actively updated benchmark using 2,100+ state-backed information operations to measure how easily frontier LMs can be co-opted.

- **LeanCSP: A Framework for Certifying Constraint Reformulation and Solving in Lean** — [http://arxiv.org/abs/2607.28459v1](http://arxiv.org/abs/2607.28459v1) — Pablo Manrique, Stefan Szeider  
  Provides formal certificates for constraint-programming reformulations and solver results inside the Lean proof assistant, enabling high-assurance combinatorial solving.

- **PAIChecker: Uncovering and Checking PR-Issue Misalignment in SWE-Bench-Like Benchmarks** — [http://arxiv.org/abs/2607.28587v1](http://arxiv.org/abs/2607.28587v1) — Manyi Wang et al.  
  Detects and verifies misalignment between pull requests and linked issues in SWE-bench-style benchmarks, exposing a hidden threat to LLM evaluation validity.

### 📊 Applications

- **Beyond Sentiment: Structured Information Extraction from Financial News** — [http://arxiv.org/abs/2607.28496v1](http://arxiv.org/abs/2607.28496v1) — Daohan Zhu et al.  
  Extracts orthogonal dimensions such as event type, impact scope, and temporal horizon from financial news instead of reducing articles to a single sentiment score.

- **A report-grounded vision-language foundation model for colonoscopy from 280000 routine reports** — [http://arxiv.org/abs/2607.28466v1](http://arxiv.org/abs/2607.28466v1) — Jia Yu et al.  
  Trains a colonoscopy vision-language model using 280,000 routine reports to weakly supervise frame-level understanding, greatly reducing reliance on manual annotation.

- **Learning to Trace Seiberg Dualities** — [http://arxiv.org/abs/2607.28628v1](http://arxiv.org/abs/2607.28628v1) — Jonathan J. Heckman et al.  
  Applies machine learning to classify when two physical theories are Seiberg dual, tackling a computationally hard physics problem and opening dualities to learning-based discovery.

## Research Trend Signal

Several trends emerge from today’s submissions. First, test-time compute is being re-examined: repeated sampling can beat elaborate reflection at equal token budgets, while RL is being used to learn adaptive self-verification budgets instead of uniform scaling. Second, agent evaluation is moving from isolated code tasks toward end-to-end operational settings—computer-use reward modeling, oncall root-cause analysis, and benchmark-integrity checking now receive dedicated benchmarks and validation tools. Third, auditing is becoming a first-class method: clinical subgroup fairness, system-prompt opacity, and state-backed manipulation are addressed with reproducible pipelines and live benchmarks. Fourth, domain applications are leaning on weak and unstructured supervision—280k colonoscopy reports, multilingual mental-health social media, financial news, and scientific dualities—to reduce dependence on hand-labeled data. Finally, growing interest in formal guarantees, from Lean-certified constraint solving to confidence-certified opponent exploitation, suggests a broader turn toward verifiability in AI systems.

## Worth Deep Reading

- **Sample More, Reflect Less** ([http://arxiv.org/abs/2607.28576v1](http://arxiv.org/abs/2607.28576v1)) — Provides a simple but powerful baseline for interpreting accuracy gains from self-refinement and related methods; anyone claiming improvements from reflection should read this first.

- **SVR: Self-Verifying Refinement** ([http://arxiv.org/abs/2607.28457v1](http://arxiv.org/abs/2607.28457v1)) — Learning to allocate test-time compute with calibrated self-verification is a major step beyond fixed budgets and external verifiers, with likely broad impact on reasoning models.

- **ORCA-bench** ([http://arxiv.org/abs/2607.28545v1](http://arxiv.org/abs/2607.28545v1)) — Defines a realistic operations task with noisy logs and ambiguous incident reports; likely to become a reference benchmark for agentic root-cause analysis and reliability research.

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*