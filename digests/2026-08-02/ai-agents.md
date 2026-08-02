# OpenClaw 生态日报 2026-08-02

> Issues: 222 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-08-02 00:13 UTC

- [OpenClaw](https://github.com/openclaw/openclaw)
- [NanoBot](https://github.com/HKUDS/nanobot)
- [Hermes Agent](https://github.com/nousresearch/hermes-agent)
- [PicoClaw](https://github.com/sipeed/picoclaw)
- [NanoClaw](https://github.com/qwibitai/nanoclaw)
- [NullClaw](https://github.com/nullclaw/nullclaw)
- [IronClaw](https://github.com/nearai/ironclaw)
- [LobsterAI](https://github.com/netease-youdao/LobsterAI)
- [TinyClaw](https://github.com/TinyAGI/tinyagi)
- [Moltis](https://github.com/moltis-org/moltis)
- [CoPaw](https://github.com/agentscope-ai/CoPaw)
- [ZeptoClaw](https://github.com/qhkm/zeptoclaw)
- [ZeroClaw](https://github.com/zeroclaw-labs/zeroclaw)

---

## OpenClaw 项目深度报告

# OpenClaw 项目动态日报 — 2026-08-02

## 1. 今日速览

过去 24 小时，OpenClaw 项目维持在极高的社区活跃度：新增/活跃 Issue 200 条、关闭 22 条（关闭率约 9.9%）；PR 更新 500 条，其中 103 条已合并/关闭（合并/关闭率约 20.6%）；新发布 1 个预发布版本 `v2026.7.2-beta.6`。从内容信号看，今日最突出的主线是**状态安全与恢复**：新版本聚焦崩溃恢复、数据隔离与 schema 变更保护，而数条 P0/P1 Issue 也集中在数据库损坏、崩溃循环、锁无法释放、消息丢失等数据面问题上。值得注意的是，仓库同时出现了 3 个由活跃维护者提交的、带有生产环境量化数据的性能优化 PR（会话列表、CLI 冷启动、内存搜索），说明项目在稳定性和规模化方面双线推进。需要警示的是，当前仍有 397 条 PR 处于待合并状态，治理积压是项目健康度的主要风险点。

## 2. 版本发布

### v2026.7.2-beta.6（2026.7.2）

- 发布链接：https://github.com/openclaw/openclaw/releases/tag/v2026.7.2-beta.6

本次版本核心议题为 **State safety and recovery（状态安全与恢复）**，包含以下亮点：

1. **隔离存储（Quarantine store）** — 当主数据库受损时，持久化数据先进入隔离存储，避免因主库损坏而直接丢失数据。
2. **崩溃可恢复的 SQLite 快照（crash-recoverable SQLite snapshots）** — 在崩溃后可通过快照恢复会话状态。
3. **崩溃持久的文件系统发布（crash-durable filesystem publication）** — 确保文件发布过程中进程被杀不会留下半写状态。
4. **Schema 升级数据丢失拒绝（schema-upgrade data-loss rejection）** — 升级前检测可能造成数据丢失的 schema 变更并拒绝执行。
5. **回滚写入者快照恢复（rollback-writer snapshot recovery）** — 为写入者回滚场景提供快照级恢复能力。

**迁移注意事项**：这是一个 beta 版本，涉及 SQLite 层的数据保护路径。升级前建议完整备份 `~/.openclaw/state/` 下的 SQLite 文件。若之前用旧版本二进制直接打开了高 schema 版本的数据库目录，新版本会触发新的安全恢复路径（与本日活跃 Issue #115421 所描述的数据丢失场景相关），需特别注意恢复日志。

## 3. 项目进展

过去 24 小时内已合并/关闭 103 条 PR，并有多条高质量 PR 进入维护者审查队列。本次主要推进方向：

### 已关闭 / 自动合并
- **PR #117704（closed, automerge armed）** — `fix(xai): Grok subscription inference fails with HTTP 426 on every request`：修复 Grok 订阅推断在所有请求上返回 HTTP 426 的问题，已进入自动合并闭环。
  https://github.com/openclaw/openclaw/pull/117704

### 性能优化（均附生产环境数据）
- **PR #117707** — `perf(sessions): reuse plugin metadata during session listing`：在生产主机（Hetzner 48 核、Node 26.5、约 150 个已安装插件）上，列出 **344 个会话耗时 31.1–31.3 秒**，且 `--limit 5` 仍需 27 秒——瓶颈在每行独立解析插件元数据。该 PR 通过复用元数据消除该卡顿。
  https://github.com/openclaw/openclaw/pull/117707
- **PR #117705** — `fix(cli): keep gateway-backed agent turns cold`：每个 gateway-backed `openclaw agent` turn 在到达 Gateway 之前需支付约 13 秒的固定启动税。该 PR 将 CLI 侧启动开销大幅削减，适合高频短任务的 gatekeeper 模式。
  https://github.com/openclaw/openclaw/pull/117705
- **PR #117706** — `fix(memory): qualify stale memory search results`：修复 `memory search` 对已存在记忆报 "no results" 的误导性结果（真实复现：事实已写入 `2026-08-01.md` 但搜索无结果）。
  https://github.com/openclaw/openclaw/pull/117706

### 会话与上下文
- **PR #115049** — `fix(codex): reduce active run transcript overhead`：避免 Codex context-engine 在单次 attempt 内重复物化完整 SQLite 转录，大幅降低大会话的解析、上下文构建与图像清理成本。状态：🟡 等待维护者审查（merge-risk: session-state）。
  https://github.com/openclaw/openclaw/pull/115049
- **PR #117509** — `fix(agents): surface sessions_yield waiting status`：修复 `sessions_yield` 携带等待消息时，父 turn 结束却未向用户发送任何内容的问题。状态：需要 proof（merge-risk: message-delivery）。
  https://github.com/openclaw/openclaw/pull/117509
- **PR #116562** — `fix(memory-core): recover primary embedding provider`：修复主 embedding 提供方短暂故障后，系统永久停留在回退提供方的问题。状态：⏳ 等待作者。
  https://github.com/openclaw/openclaw/pull/116562

### 渠道与消息通道
- **PR #114552** — `fix(slack): isolate message ingress by thread`：将 Slack 入站 lane 从“频道级”改为“频道+线程级”，避免某个线程的长回复阻塞同频道其他线程。
  https://github.com/openclaw/openclaw/pull/114552
- **PR #114902** — `fix(qqbot): add heartbeat-ACK liveness watchdog`：为 QQBot 网关 WebSocket 增加 `op:11` 心跳 ACK 检测，避免连接半开时客户端无感知。状态：🟡 ready for maintainer look。
  https://github.com/openclaw/openclaw/pull/114902
- **PR #114580** — `fix(tlon): refresh channel history after monitor restart`：修复监控器重启后复用旧缓存消息的问题。
  https://github.com/openclaw/openclaw/pull/114580
- **PR #117689** — `fix(control-ui): restore sandbox media previews`：恢复 Control UI 中沙箱化媒体路径的解析与预览。
  https://github.com/openclaw/openclaw/pull/117689

### 安全与稳定性
- **PR #117394** — `fix(security): skill scans detect aliased shell execution`：技能代码安全扫描现在可识别 ESM import 别名及 CommonJS 解构别名包装的 `child_process` 执行调用，堵住扫描绕过盲区。状态：🟡 ready for maintainer look（merge-risk: security-boundary）。
  https://github.com/openclaw/openclaw/pull/117394
- **PR #117652** — `fix(channels): contain transport watchdog failures`：防止频道传输看门狗回调抛错导致整个网关进程退出。
  https://github.com/openclaw/openclaw/pull/117652

### 质量与工具链
- **PR #117701** — `test: consolidate OpenClaw E2E fixtures`：集中 E2E 夹具，完整保留 29 个源注册、31 个运行时用例、71 个源断言、92 个运行时匹配调用，并消除重复代码。
- **PR #117698** — `refactor(fal): consolidate image-generation provider test fixtures`：将 Fal provider 测试从 1,186 行重复代码收敛为可维护的共享夹具。
- **PR #117690** — `fix(macos): avoid self-named defaults suite`：修复 macOS 上 `NSUserDefaults(suiteName:)` 使用自身 bundle ID 导致的警告（对应 Issue #88909）。

## 4. 社区热点

1. **Issue #116277 — DeepSeek v4 Flash 静默回复失败（73 评论，P1）**
   https://github.com/openclaw/openclaw/issues/116277
   用户在 Telegram 群组中触发 deepseek-v4-flash，模型静默不生成回复，OpenClaw 仅发布通用 fallback "No reply was generated for this message"。社区焦点在于：**模型失败时必须有明确的错误面，而不是用模糊 fallback 掩盖问题**。该 Issue 同时带 `impact:message-loss` 与 `impact:ux-friction`，最终用户无法区分“模型拒绝”和“系统 bug”。

2. **Issue #116201 — Realtime voice 可无限保留 provider/consult 状态（36 评论，P1）**
   https://github.com/openclaw/openclaw/issues/116201
   实时语音会话在慢速、停滞或突发 provider 行为下，会无限保留被取代的 consult 工作、大尺寸 provider frame、预就绪音频和播放缓存。社区讨论集中在“需要硬性所有权边界（hard ownership bounds）”，而非仅靠 item 数量或取消信号来兜底。

3. **Issue #115326 — 崩溃循环断路器永久抑制 Discord/WhatsApp（24 评论，P1）**
   https://github.com/openclaw/openclaw/issues/115326
   断路器激活后，即使 Gateway 重启成功，Discord 与 WhatsApp 也被“永久抑制”；文档中的恢复路径 `channels.start` 失败并报 WebSocket 1006。用户核心诉求是：**断路器需要手动重置、自动降级或基于存活性探测的自动恢复**。

4. **Issue #99241 — 工具输出渲染为图片附件，agent 不可读（26 评论，P1，今日已关闭）**
   https://github.com/openclaw/openclaw/issues/99241
   长输出/ANSI-heavy 工具结果被折叠成 `(see attached image)`，agent 无法读取原始 stdout/stderr，而这些文本往往是决策证据。该 Issue 今日关闭（修复链路已合入），但评论仍强调“可读性证据链”需求。

5. **PR #110024 — Gateway 频道永久卡死但 /health 保持绿色（长期开放，P1）**
   https://github.com/openclaw/openclaw/pull/110024
   修复 channel provider 断开连接后，通道进入“无重启、无日志、无错误”的永久卡死状态；生产环境曾因此出现两次多小时的宕机（~34h）。这一 PR 被多条 Issue 引用，但至今仍处于 `needs proof`，社区对它的关注度在上升。

## 5. Bug 与稳定性

### P0（数据丢失 / 严重破坏）
- **Issue #115421** — Schema 降级恢复时不得 quarantine/清空状态 DB：旧版二进制打开高版本 schema 的 `openclaw.sqlite` 后，恢复逻辑生成了 `.bak-schema6`、`.moved-schema6` 和全新空库，用户 cron 任务全部丢失。已有 `linked-pr-open`。 ⚠️ 与今日发布的 v2026.7.2-beta.6 恢复机制直接相关。
  https://github.com/openclaw/openclaw/issues/115421
- **Issue #48920** — Live Docs 领先于 release（P0，自 2026-03-17 创建，持续近 5 个月）：文档中的 `IsolatedSessions` 在最新版 2026.3.13 中不存在。4 个 👍，仍无对应修复。
  https://github.com/openclaw/openclaw/issues/48920

### P1（高优先级，大部分已有维护者介入）
- **Issue #115424** — Gateway V8 堆 OOM → 重启恢复变成 7 次核心转储循环：main-session 长 turn 触发 `FATAL ERROR: Reached heap limit`（SIGABRT），restart-recovery 自动热恢复同一会话，形成“一次崩溃变七次崩溃”的雪崩。
  https://github.com/openclaw/openclaw/issues/115424
- **Issue #115326** — Crash-loop breaker 永久抑制 Discord/WhatsApp（详见社区热点）。
- **Issue #115546** — CLI-budget compaction 在 deadline 前 4.9s–50s 超时，大会话 100% 失败，且无重试，形成 wake 死亡螺旋。
  https://github.com/openclaw/openclaw/issues/115546
- **Issue #94939** — 6.x 状态迁移将 channel conversation-store SQLite 落为 0 字节，MS Teams proactive 发送失效。
  https://github.com/openclaw/openclaw/issues/94939
- **Issue #114234** — Usage-cost 刷新锁在容器 PID 复用后永久不可释放，缓存被冻结；已有 `linked-pr-open`（对应 PR #103961）。
  https://github.com/openclaw/openclaw/issues/114234
- **Issue #116488** — 被取代的 reply 操作未从 reply-run registry 释放，会话在 `run:completed` 后仍报告 active work，watchdog 空等 abort 计时器。
  https://github.com/openclaw/openclaw/issues/116488
- **Issue #87763** — SSRF guard 固定 DNS dispatcher 与 Node `autoSelectFamily`（Happy Eyeballs）冲突，导致模型请求 120 秒超时、网关完全无响应；已有 `linked-pr-open`。
  https://github.com/openclaw/openclaw/issues/87763
- **Issue #115476** — 上下文压缩后刷新会重放旧入站 Telegram `message_id`，缺少网关级去重，导致重复消息。
  https://github.com/openclaw/openclaw/issues/115476
- **Issue #114084** — `session_entries` → `session_nodes` 迁移（PR #113071）后，Diagnostic 子系统仍引用旧表名，WhatsApp 消息处理失败。
  https://github.com/openclaw/openclaw/issues/114084
- **Issue #116022** — `/new` 复用稳定 session ID，无法回收已退休的 Codex binding tombstone。
  https://github.com/openclaw/openclaw/issues/116022

### P2（值得关注的回归与缺陷）
- **Issue #116010** — 所有持久化会话被硬编码限制在 128k context，与模型或 `contextTokens` 配置无关。
  https://github.com/openclaw/openclaw/issues/116010
- **Issue #112906** — `richMessages: true` 下 `details` 标签不再渲染为折叠区块，内容平铺且无法折叠；v2026.7.1 回归。
  https://github.com/openclaw/openclaw/issues/112906
- **Issue #98976** — Provider refusal（Anthropic refusal / OpenAI content_filter）不触发 fallback 链，turn 直接以泛化 `LLM request failed.` 结束。已有 `linked-pr-open`。
  https://github.com/openclaw/openclaw/issues/98976
- **Issue #116691** — openai-responses 调用火山引擎长对话时报缺失 `input.status` 参数；中文用户社区受影响。
  https://github.com/openclaw/openclaw/issues/116691

### 今日已关闭的 Bug（修复已验证）
- **Issue #99241** 工具输出图片化不可读（P1）— 已关闭。
- **Issue #34528** Feishu reaction message_id 后缀导致 400（P2）— 已关闭。
- **Issue #106730** exec 工具在 Linux 上将 `|` 当管道运算符（P2）— 已关闭。
- **Issue #90203** 工作区文件存在性检查硬编码 `MEMORY.md` 等固定列表（P3）— 已关闭。
- **Issue #115413** Compaction 报告成功但 summarization 全部失败（P1）— 已关闭。

## 6. 功能请求与路线图信号

### 可能进入下一版本（已有 linked-pr-open 或处于需求收敛阶段）
- **#114146 — `talk.realtime.providers.<id>.baseUrl`（P2）**：为 Realtime voice 增加 OpenAI Realtime 兼容第三方 endpoint 配置。阿里云百炼 Qwen3-ASR-Flash 等已支持兼容协议，社区声音较强。
  https://github.com/openclaw/openclaw/issues/114146
- **#95724 — memory 按 source directory 索引（P2）**：多个 agent 共享同一 workspace 时，目前各自构建独立 vector index，造成大量重复存储。
  https://github.com/openclaw/openclaw/issues/95724
- **#98976 — Provider refusal 触发 fallback 链**：与稳定性直接相关，已有 PR 在推进。

### 社区呼声较高的新能力
- **#113251 — WebChat 文件查看器支持图片浏览（P2，platinum hermit）**：当前只能查看文本类附件。
  https://github.com/openclaw/openclaw/issues/113251
- **#114264 — 基于消息类型的自动模型路由（P2）**：按 text/image/audio/TTS 自动选择模型，适配小米 MiMo V2.5 等多模型用户场景。
  https://github.com/openclaw/openclaw/issues/114264
- **#86983 — 出站 DM 白名单 `dmAllowTo`（P2，security）**：当前 `dmPolicy: allowlist` 只限制入站，无法阻止 agent 主动向任意 Slack 用户发送 DM。
  https://github.com/openclaw/openclaw/issues/86983
- **#95279 — 可信入站装饰契约（+4 👍）**：装饰元数据直接拼进 `content`，消费端只能用可伪造的文本启发式去解析/去重，社区建议提供可信结构化契约。
  https://github.com/openclaw/openclaw/issues/95279
- **#115924 — Idea Shower：Agent 执行时的并行想法收集器（P3）**：支持用户不打断 agent 的前提下并行提交想法，贴近“人类思维并行”的使用习惯。
  https://github.com/openclaw/openclaw/issues/115924
- **#95516 — Skill 生命周期管理（P3）**：自动优化失败技能 + 基于使用率自动退役。
  https://github.com/openclaw/openclaw/issues/95516

### 路线图信号
PR #114173（`feat(gateway): present system-agent setup QR codes`）显示项目在向 **Control UI / 移动端设置体验** 扩展；PR #80228（Codex continuity bridge，长期开放）若能合入，将补全“Codex 高带宽工作台 + OpenClaw 静默操作面”的混合工作流，被多条功能 Issue 视为前置依赖。

## 7. 用户反馈摘要

### 正面反馈
- 无障碍用户（#95601）明确感谢团队在 v2026.6.9 中将剩余用量信息移到模型选择器附近，使 macOS VoiceOver 用户可以在一个键盘可达区域内完成模型与用量查看。
- 多位用户认可 OpenClaw 在长会话、多 agent 编排场景的扩展能力（如 gatekeeper 模式、sessions_send/sessions_yield），但普遍反映“等待语义”和“交付确认”需要补齐。

### 核心痛点
1. **静默失败是最集中的情绪引爆点**：
   - #116277 DeepSeek 静默无回复；#98976 provider refusal 不落 fallback；#115413 compaction 假成功——用户反复遇到“系统提示成功/通用错误，但真实原因不可见”，强烈要求区分模型拒绝、网络错误、配额与内部 bug。
2. **恢复路径不可靠**：
   - #115326 断路器无法手动重置；#114234 锁永久冻结；#116022 `/new` 无法清理 tombstone；#115421 降级直接清库——用户对“文档中的恢复命令实际不可用”表示不满。
3. **长会话/大上下文成为头号性能瓶颈**：
   - #116010 128k 限制与模型无关；#115546 压缩 100% 失败；#115424 OOM 后 7 次循环；#115049 转录重复物化——专业用户大量跑 orchestrator 和长会话，对上下文管理的稳定性极其敏感。
4. **多平台一致性问题**：
   - #114552 Slack 线程互相阻塞、#115476 Telegram 消息重放、#94939 Teams proactive 中断、#117491 心跳投递到错误 Telegram 账号——每接入一个渠道，边界条件就成倍暴露。

## 8. 待处理积压

### 超长期未响应 / 未合并
- **Issue #48920（P0，自 2026-03-17，持续近 5 个月）** — Live Docs 领先于 release。P0 且唯一阻塞是“校准文档与代码版本”，建议维护者立即处理。
  https://github.com/openclaw/openclaw/issues/48920
- **PR #80228（自 2026-05-10）** — Codex continuity bridge，3 个月未合并，状态 `needs-proof` + `needs-pr-context`，但被多条功能 Issue 引用为前置依赖。
  https://github.com/openclaw/openclaw/pull/80228
- **PR #81470（自 2026-05-13）** — WebChat TTS 音频在 `broadcastChatFinal` 中丢失，`merge-risk` 涉及 session-state 与 message-delivery，需重点 review。
  https://github.com/openclaw/openclaw/pull/81470

### 高影响力但等待验证
- **PR #110024（P1）** — 解决频道永久卡死但健康检查为绿的严重问题；生产环境曾有 ~34h 事故。仍处于 `needs proof`。
  https://github.com/openclaw/openclaw/pull/110024
- **PR #103961（P1 关联）** — usage refresh lock 泄漏修复，对应 #114234/#103910 双 P1 issue，自 2026-07-10 起等待验证。
  https://github.com/openclaw/openclaw/pull/103961
- **Issue #87763（P1，自 2026-05-28）** — SSRF guard 与 autoSelectFamily 冲突，2 个月未合入。已有 `linked-pr-open`，建议加速合并优先级。

### 治理积压提示
当前 397 条待合并 PR 中，相当一部分来自 `clawsweeper[bot]` 自动修复和 Dependabot 依赖更新。建议维护者关注自动生成 PR 的合并周期，避免机器人生成速度超过人工 review 吞吐，形成隐性技术债。

---

**项目健康度评估**：整体处于高活跃、高关注、高修复速度的状态，但 P0/P1 问题存量较大（尤其状态持久化与恢复链路），且 397 条待合并 PR 提示治理瓶颈。短期内建议优先合并状态安全相关的修复（#115421、#114234、#94939），并对性能优化 PR（#117707、#117705）做快速验证放行——它们直接影响大规模用户的使用体验。

---

## 横向生态对比

# 个人 AI 助手/自主智能体开源生态横向对比分析报告

**报告日期：2026-08-02 | 数据窗口：2026-08-01 ~ 2026-08-02（UTC）**

---

## 1. 生态全景

个人 AI 助手/自主智能体开源生态正处在**高速分化与密集加固并行**的阶段。以 OpenClaw 为首的活跃项目在保持高吞吐迭代节奏的同时，集体将重心从“功能叠加”转向“生产级稳定性”——状态恢复、数据隔离、静默失败可观测性成为多个项目不约而同的主攻方向。生态内部形成明显分层：通用型运行时（OpenClaw）、桌面优先（Hermes）、渠道接入中心（NanoClaw）、协议原生（Moltis）、轻量嵌入式（PicoClaw）各据其位，但跨项目浮现高度同构的需求痛点，表明行业正从“能做功能”迈向“能可靠运行”。治理瓶颈（PR 积压、文档滞后、stale 自动关闭掩盖真实问题）开始成为制约多个项目健康度的共性因素。

---

## 2. 各项目活跃度对比

| 项目 | Issue 动态 | PR 动态 | Release | 健康度评估 |
|---|---|---|---|---|
| **OpenClaw** | 200 活跃，22 关闭 | 500 更新，103 合并/关闭 | 1 预发布（v2026.7.2-beta.6） | 高活跃；P0/P1 存量较大，397 条 PR 积压 |
| **Hermes Agent** | 11 更新，3 关闭 | 50 更新，20 合并/关闭 | 无 | 健康；安全/更新管道修复已闭环 |
| **NanoBot** | 5 更新，4 关闭 | 25 更新，13 合并/关闭 | 无 | 健康；修复效率高，少量 PR 积压 |
| **IronClaw** | 20 更新 | 25 更新，≥5 合并/关闭 | 无 | 良好；架构重构执行力强，性能问题待解 |
| **NanoClaw** | 2 更新（1 新开 1 关闭） | ~16 动态，6 合并 | v2.1.54（滚动） | 良好；迭代速度快，1 个回归 Bug 未修 |
| **ZeroClaw** | 14 更新，0 关闭 | 50 更新，**0 合并** | 无 | 存在治理瓶颈；S0 数据隔离缺陷无 PR |
| **CoPaw** | 9 活跃 | 12 更新，1 合并 | 无 | 良好；首次贡献者活跃，新功能积累中 |
| **Moltis** | 0 新开 | 4 更新，3 合并/关闭 | 无 | 健康；队列短，安全修复及时 |
| **PicoClaw** | 1 评论更新 | 3 更新，1 合并 | 无 | 中等；核心 bug（Matrix 重连）积压 1 个月 |
| **LobsterAI** | 6 stale 关闭 | 2 待合并 | 无 | 维护期；关键 PR 等待 4 个月未合并 |
| **NullClaw / TinyClaw / ZeptoClaw** | 无活动 | 无活动 | 无 | 停滞 |

---

## 3. OpenClaw 在生态中的定位

OpenClaw 是当前生态的**核心参照基准**，以 500 PR/日、200 Issue/日的量级领先第二名一个数量级。其技术路线最鲜明的特征是**将“状态安全与恢复”作为一级公民**——v2026.7.2-beta.6 引入隔离存储、崩溃可恢复 SQLite 快照、schema 升级数据丢失拒绝等机制，深度押注“进程可能随时被杀”的真实生产场景，而非仅追求功能丰富度。相比之下，Hermes 侧重桌面端配置更新链路安全，NanoClaw 侧重多渠道接入统一，Moltis 聚焦协议原生与企业权限模型，均未在持久化崩溃恢复层面达到 OpenClaw 的纵深。

社区规模上，OpenClaw 的 Issue/PR 讨论深度与生产环境量化数据（如 Hetzner 48 核上的性能基准）远超同类，已形成“事实标准”的吸引力。其主要风险也在规模本身——397 条待合并 PR 的治理积压是当前生态中最突出的健康度警示。

---

## 4. 共同关注的技术方向

| 技术方向 | 涉及项目 | 具体诉求 |
|---|---|---|
| **状态恢复与防数据丢失** | OpenClaw、NanoBot、NanoClaw、ZeroClaw、PicoClaw | SQLite 崩溃快照（OpenClaw）；损坏 session 容错（NanoBot #4801）；outbound.db journal 恢复（NanoClaw #2750）；跨 agent 数据越权 S0（ZeroClaw #9646/#9647）；Matrix 断线静默死亡（PicoClaw #3203） |
| **静默失败与可观测性** | OpenClaw、PicoClaw、NanoClaw、ZeroClaw、Hermes | 模型无回复但系统报通用错误（OpenClaw #116277）；进程存活但完全失能（PicoClaw #3203）；凭据过期误报为文件系统错误（NanoClaw #3167）；超时伪造审计日志（ZeroClaw #9642）；更新失败无提示（Hermes #75598） |
| **配置/密钥/权限隔离** | Hermes、ZeroClaw、Moltis、CoPaw | Profile 间凭据泄露（Hermes #51603）；per-agent 数据所有权（ZeroClaw #9646）；channel 访问权与操作权分离（Moltis #1170）；内置技能越权拦截（CoPaw/类 NanoClaw #3171） |
| **成本与性能规模化** | OpenClaw、IronClaw、NanoBot、ZeroClaw | 会话列表加载 31s 瓶颈（OpenClaw #117707）；libSQL p95 135s（IronClaw #6974）；JSONL 列表加速（NanoBot #5194）；OpenRouter 因无稳定 session_id 导致重复计费（ZeroClaw #9631） |
| **更新/安装管道可靠性** | Hermes、NanoClaw、PicoClaw | EBADENGINE 更新死路（Hermes #76464）；rootless Docker 环境假设失效（NanoClaw #3174）；插件启用 ensurepip 缺失（NanoBot #5205） |
| **模型路由与 fallback 语义** | OpenClaw、IronClaw、CoPaw、ZeroClaw | Provider refusal 不触发 fallback（OpenClaw #98976）；OrcaRouter 接入（IronClaw #7009、CoPaw #6622、PicoClaw #3309）；请求预算上限（Hermes #76458） |

---

## 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 关键架构特征 |
|---|---|---|---|
| **OpenClaw** | 全功能通用 agent 运行时，长会话/多 agent 编排 | 专业开发者、大规模部署 | SQLite 状态层 + Gateway 进程模型，深度崩溃恢复 |
| **Hermes Agent** | 桌面端优先、多 profile 隔离、更新链路安全 | 个人开发者、多租户/多身份场景 | 密钥作用域隔离 + 自管 Node 运行时，重配置管理 |
| **NanoClaw** | 渠道接入中心（iMessage 双后端）、轻量容器化 | 个人用户、Apple 生态 | 双后端（Local Chat SDK + Hosted Photon），滚动发布 |
| **NanoBot** | Python 轻量级、Cron/Dream 自动化、WebUI | 开源社区、快速原型 | Python + uv 工具链，插件生态丰富 |
| **IronClaw** | 架构契约治理、CI 门禁完整性、Wave 式重构 | 企业级开发者、对架构纪律敏感 | product_contracts 边界收敛，叠堆 PR 链式交付 |
| **Moltis** | Nostr 协议原生、企业级权限模型 | 去中心化/隐私敏感用户 | NIP-29 群聊 + per-account operators 权限矩阵 |
| **PicoClaw** | 轻量嵌入式、中文社区、多 provider 路由 | 极简部署、资源受限环境 | 单二进制友好，i18n 覆盖繁中 |
| **ZeroClaw** | 安全/多租户敏感、审批流 | 运营商、安全测试场景 | SSRF 门禁 + 数据所有权隔离 + 审计日志完整性 |
| **LobsterAI** | 桌面应用层、协同时代 IDE 形态 | 中文桌面用户、网易生态 | 前端体验导向，i18n 与 UX 打磨 |

---

## 6. 社区热度与成熟度分层

**第一梯队 — 高速迭代期（日 PR 更新 ≥25）**：OpenClaw、Hermes Agent、NanoBot、IronClaw、ZeroClaw。共同特征：合入速度快、社区贡献者活跃、存在明确的路线图或重构纲领。ZeroClaw 例外——提交量大但零合并，处于“社区热、维护冷”的亚健康状态。

**第二梯队 — 功能扩展与加固并行（日 PR 更新 5~20）**：NanoClaw、CoPaw、Moltis。迭代节奏适中，质量把控较好，首次贡献者融入顺利，无重大阻塞。

**第三梯队 — 稳定维护期（低频更新）**：PicoClaw、LobsterAI。主要工作量集中在存量问题清理与少量功能合入，但关键修复等待时间过长（LobsterAI #1224 积压 4 个月、PicoClaw #3203 一个月无响应），有 contributor 流失风险。

**停滞状态**：NullClaw、TinyClaw、ZeptoClaw 过去 24 小时无任何动态。

---

## 7. 值得关注的趋势信号

1. **“状态安全”正取代“新功能”成为第一优先级**：OpenClaw 新版本全部围绕崩溃恢复与数据隔离，NanoBot/NanoClaw 集中修复 session/日志损坏容错。对开发者的启示：agent 应用已进入“可靠性红利”阶段，能保证数据不丢、进程可恢复是差异化竞争力。

2. **静默失败是最集中的用户情绪引爆点**：从 OpenClaw 的 DeepSeek 无回复、到 PicoClaw 的 Matrix 僵尸进程、再到 NanoClaw 的误导性错误信息，用户对“系统说成功/通用错误但真实原因不可见”的容忍度降至最低。**错误面设计（error surface）必须明确区分模型拒绝、网络错误、配额与内部 bug。**

3. **多 agent / 多 profile 数据隔离从功能变为安全底线**：Hermes 修复跨 profile 密钥泄露、ZeroClaw 出现 S0 级跨 agent 数据越权、Moltis 重构操作权限矩阵——单实例多租户场景的部署比例正在快速上升，隔离能力将成为选型硬指标。

4. **成本控制被前置到架构层**：OpenRouter 缓存 session_id 缺失导致 30-60% 额外成本（ZeroClaw）、每轮 provider 请求预算（Hermes）、token 核算修正（IronClaw）——成本不再只是账单问题，而是架构设计输入。

5. **更新/安装管道可靠性成为用户信任放大器**：Hermes 的 EBADENGINE 修复、NanoBot 的 ensurepip 缺失、NanoClaw 的 rootless Docker 失败——安装配置链路是用户对项目的第一印象，此处出问题会直接摧毁复购意愿。

6. **跨项目治理积压普遍化**：OpenClaw 397 条待合并 PR、ZeroClaw 50 条 PR 零合并、LobsterAI 4 个月无 review——自动化 PR（bots）与人工 review 吞吐的失衡正在成为开源健康度的系统性风险，需要引入更激进的自动合并策略或定期“清仓”机制。

7. **OrcaRouter/模型路由生态信号密集出现**：IronClaw、CoPaw、PicoClaw 三个互不关联的项目在同一天出现 OrcaRouter 接入请求/PR，叠加 OpenClaw 的 provider refusal fallback 讨论，表明“多模型网关 + 智能路由”正成为 agent 基础设施的标配能力。

---

*本报告基于 2026-08-02 各项目 GitHub 公开数据自动生成，仅供技术决策参考。*

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目日报（2026-08-02）

## 1. 今日速览

过去 24 小时 NanoBot 保持高活跃度：共更新 5 条 Issue（4 条关闭），25 条 PR（13 条被合并或关闭，12 条仍待合并），未发布新版本。今日多个 p1 级稳定性修复被收口，覆盖 cron、memory、exec、session 等关键路径，同时有 Quick Chat、跨会话搜索、trusted proxy 认证等新功能 PR 仍在推进。整体看，项目修复效率高、社区贡献者众多，当前处于“稳定加固 + 功能扩展并行”的相对健康阶段。

## 2. 版本发布

本次日报周期内无新版本发布。

## 3. 项目进展

今日被合并/关闭的 PR 主要集中在稳定性修复和防御性改进：

- **Cron / 持久化**：[#5183](https://github.com/HKUDS/nanobot/pull/5183) 修复手动运行 cron 完成状态丢失问题（对应 Issue [#5163](https://github.com/HKUDS/nanobot/issues/5163)）。
- **Memory 健壮性**：[#5153](https://github.com/HKUDS/nanobot/pull/5153) 修复 raw_archive 中非字符串时间戳和缺失 `role` 导致的 KeyError（对应 Issue [#4801](https://github.com/HKUDS/nanobot/issues/4801)）。
- **Dream 游标**：[#5208](https://github.com/HKUDS/nanobot/pull/5208) 修复 Dream cron 在产生持久化修改后不推进历史游标的问题。
- **Exec 输出截断**：[#5200](https://github.com/HKUDS/nanobot/pull/5200) 修复响应截断后 `wait_for` 目标丢失的问题。
- **Session 容错**：[#5201](https://github.com/HKUDS/nanobot/pull/5201) 容忍损坏的持久化 session summary。
- **核心能力**：[#5172](https://github.com/HKUDS/nanobot/pull/5172) 保留 Responses API reasoning 状态并压缩上下文，采纳 OpenAI ARC-AGI-3 报告中强调的能力。
- **频道安全**：[#5108](https://github.com/HKUDS/nanobot/pull/5108) 为渠道适配器增加 per-sender 消息速率限制。
- **Provider 修复**：[#3732](https://github.com/HKUDS/nanobot/pull/3732) 要求先匹配 `api_base`，避免本地 provider 按关键字静默劫持云模型。
- **WebUI / 代码质量**：[#5209](https://github.com/HKUDS/nanobot/pull/5209) 复用侧边栏选中高亮组件；[#5199](https://github.com/HKUDS/nanobot/pull/5199) 缩小 Pyright suppression 范围。

整体来看，项目今日在数据稳定性、cron 可靠性、渠道安全三个方向有实质进展。

## 4. 社区热点

- **[Issue #5185](https://github.com/HKUDS/nanobot/issues/5185)**：今日评论最多（4 条），用户报告“突然在响应中返回 tool call 代码”，最终被标记为 `invalid`/已关闭。这暗示问题更可能来自特定 provider 或配置变更，而非产品本身的确定性 bug。
- **[Issue #5205](https://github.com/HKUDS/nanobot/issues/5205)**（2 条评论）：用户执行 `nanobot plugins enable feishu` 时报错 `No module named ensurepip`。这属于 `uv tool` 安装环境下的 Python 依赖缺失问题，但说明插件启用流程在 Debian + self-contained Python 场景下缺少自动检测与友好提示。
- **[Issue #5198](https://github.com/HKUDS/nanobot/issues/5198)**（1 条评论）：用户无法在特定会话中切换模型，`/model` 命令对另一个模型 ID 似乎无效。该问题虽评论不多，但直接命中日常使用体验，是当前最值得关注的功能缺口。

## 5. Bug 与稳定性

按严重程度 / 优先级排列：

- **高 · 开放中**：[Issue #5198](https://github.com/HKUDS/nanobot/issues/5198) 无法在会话内切换模型，目前没有对应修复 PR。
- **高 · 已修复**：Dream cron 游标不前进导致历史批次重复处理，[PR #5208](https://github.com/HKUDS/nanobot/pull/5208) 今日关闭。
- **中 · 已修复**：
  - [#5163](https://github.com/HKUDS/nanobot/issues/5163) 手动 cron 运行完成状态丢失 → [PR #5183](https://github.com/HKUDS/nanobot/pull/5183)
  - [#4801](https://github.com/HKUDS/nanobot/issues/4801) malformed session 消息导致 KeyError → [PR #5153](https://github.com/HKUDS/nanobot/pull/5153)
- **中 · 环境 / 配置类**：
  - [#5205](https://github.com/HKUDS/nanobot/issues/5205) `ensurepip` 缺失导致飞书插件启用失败（已关闭）
  - [#5185](https://github.com/HKUDS/nanobot/issues/5185) 响应中出现工具调用代码（已关闭，`invalid`）
- **中 · 待观察 / 待合并**：
  - 流式响应日志重复（[PR #5206](https://github.com/HKUDS/nanobot/pull/5206) 开放中）
  - 媒体路径在 session 合并时丢失（[PR #5139](https://github.com/HKUDS/nanobot/pull/5139) 开放中，有 conflict）
  - DeepSeek 对 null/empty 内容敏感导致 400（[PR #3869](https://github.com/HKUDS/nanobot/pull/3869) 长期开放）

## 6. 功能请求与路线图信号

今日没有新增纯 feature request Issue，但从开放 PR 中可以看到清晰的路线图信号：

- **WebUI 交互改进**：
  - [PR #5202](https://github.com/HKUDS/nanobot/pull/5202) 让模型预设切换可发现，直接回应 Issue #5198 的痛点
  - [PR #5184](https://github.com/HKUDS/nanobot/pull/5184) 增加 Quick Chat 和 Temporary Chat
  - [PR #5211](https://github.com/HKUDS/nanobot/pull/5211) 增加跨会话搜索与 @ 提及
- **部署与集成**：
  - [PR #5210](https://github.com/HKUDS/nanobot/pull/5210) 支持 trusted upstream proxy bootstrap auth，适配 Cloudflare Tunnel / Access 场景
  - [PR #5186](https://github.com/HKUDS/nanobot/pull/5186) 支持 skills.sh well-known sources
- **子代理能力**：[PR #5207](https://github.com/HKUDS/nanobot/pull/5207) 为 `spawn` 工具增加模型预设参数
- **性能**：[PR #5194](https://github.com/HKUDS/nanobot/pull/5194) 加速 JSONL session 列表与线程加载

其中 #5184、#5211 可能构成一个较大的 WebUI 会话管理改进批次，预计会是下一个版本的功能亮点。

## 7. 用户反馈摘要

- **配置变更困惑**：[#5185](https://github.com/HKUDS/nanobot/issues/5185) 用户描述“All of a sudden”响应中开始出现工具调用代码，说明当 provider 或工具配置变化时，缺少足够可观测性来定位原因。
- **部署环境摩擦**：[#5205](https://github.com/HKUDS/nanobot/issues/5205) 用户在 Debian 服务器执行插件启用命令时直接失败，报错指向 `uv tool` 安装的 Python 缺少 `ensurepip`。这反映部署流程在特定系统环境下需要更友好的错误捕获和修复指引。
- **WebUI 易用性期待**：[#5198](https://github.com/HKUDS/nanobot/issues/5198) 用户明确希望像商业 SaaS AI 产品一样，点击输入框附近的模型名称即可切换模型，而不是重新配置整个实例。同时 `/model` 命令行为与预期不符，让用户感到困惑。
- **数据容错需求**：[#4801](https://github.com/HKUDS/nanobot/issues/4801) 来自对 malformed session 数据的健壮性要求，已在 [#5153](https://github.com/HKUDS/nanobot/pull/5153) 中修复。

## 8. 待处理积压

以下事项需要维护者重点关注：

- **[PR #3869](https://github.com/HKUDS/nanobot/pull/3869)**（2026-05-16 创建）：DeepSeek 消息加固，已开放近三个月且带有 conflict。若 DeepSeek 用户仍会遇到 400 和空内容注入问题，应尽快决策合入或关闭。
- **[PR #5139](https://github.com/HKUDS/nanobot/pull/5139)**（priority: p1，有 conflict）：修复媒体路径在 session 合并过程中丢失，对应 Issue #5118、#5135，属于数据可靠性问题，建议优先解决冲突。
- **[Issue #5198](https://github.com/HKUDS/nanobot/issues/5198)**：会话内模型切换不可用，是明确且高频的用户痛点，建议结合 [PR #5202](https://github.com/HKUDS/nanobot/pull/5202) 纳入近期版本。
- **[PR #5186](https://github.com/HKUDS/nanobot/pull/5186)**：支持 skills.sh well-known sources，开放时间较长，若 skills.sh 生态在路线图内应加快 review。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent 项目动态日报 — 2026-08-02

## 1. 今日速览

过去 24 小时项目活跃度极高：共产生 **50 条 PR 更新**（其中 20 条已合并/关闭）、**11 条 Issue 更新**（关闭 3 条），未发布新 Release。今日修复主线集中在两条关键路径：**更新/安装管道可靠性**（P1 级 EBADENGINE 故障已由 [#76464](https://github.com/NousResearch/hermes-agent/pull/76464) 合并修复）与**配置/密钥隔离安全加固**（跨配置文件凭据泄露 Issue [#51603](https://github.com/NousResearch/hermes-agent/issues/51603) 随 [#76462](https://github.com/NousResearch/hermes-agent/pull/76462) 合入而关闭）。另有多个 Windows 平台、CLI 与桌面端修复等待合并，整体项目健康度良好，维护者响应迅速。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日合并/关闭的 PR 中，按主题归纳的核心进展如下：

### 安全与密钥隔离
- **[#76462](https://github.com/NousResearch/hermes-agent/pull/76462)（已合并，P2/security）**：完成 profile 密钥范围迁移——`get_env_value`、Anthropic 适配器及 ACP 环境变量清理全部路由到 `agent.secret_scope.get_secret`，并在启动时清理 ACP 路由键继承泄漏。该 PR 直接修复了 [[#51603](https://github.com/NousResearch/hermes-agent/issues/51603)] 与 [[#75141](https://github.com/NousResearch/hermes-agent/issues/75141)] 两个安全 Issue（均已关闭）。

### 更新与安装管道（重点）
- **[#76464](https://github.com/NousResearch/hermes-agent/pull/76464)（已合并，P1）**：`hermes update` 在系统 Node 触发 `engines.npm` 不符（EBADENGINE）时，自动转用 Hermes 自管 Node 运行时，不再以手动处理提示中止更新。这是今日唯一的 P1 级修复。
- **[#75859](https://github.com/NousResearch/hermes-agent/pull/75859)（已合并）**：放宽 npm engine 约束，支持 npm 11.11.0。
- **[#76250](https://github.com/NousResearch/hermes-agent/pull/76250)（已关闭）**：使 npm 恢复在零新提交路径上可重建依赖；后续改进由 [#76467](https://github.com/NousResearch/hermes-agent/pull/76467) 接续。

### 平台与集成功能
- **[#76461](https://github.com/NousResearch/hermes-agent/pull/76461)（已合并）**：新增 `gateway.systemd_unit_name` 配置项，支持自定义 systemd 单元名。
- **[#76465](https://github.com/NousResearch/hermes-agent/pull/76465)（已合并）**：修复 Discord 中继线程重命名未携带父频道区分符的问题。
- **[#7793](https://github.com/NousResearch/hermes-agent/pull/7793)（已关闭）**：Telegram PDF 自动摄取至知识库（KB），含 Docling 管道与 OCR 回退。
- **[#18621](https://github.com/NousResearch/hermes-agent/pull/18621)（已关闭）**：skills 增加 `description_full` 选项，绕过 60 字符描述截断。
- **[#24183](https://github.com/NousResearch/hermes-agent/pull/24183)（已关闭）**：`profile create --clone` 排除 `.archive` 子树，避免废弃技能快照破坏克隆。

整体来看，项目在安全收敛、更新链路稳定性及平台适配三方面同时迈进，且多为直接回应用户反馈的修复型进展。

## 4. 社区热点

今日讨论热度集中在以下 Issue（PR 评论数据未提供）：

- **[#75598 [CLOSED] 更新问题导致程序不稳定](https://github.com/NousResearch/hermes-agent/issues/75598)**（7 条评论，Windows/desktop）：用户报告约一周前开始，更新后出现多个 gateway 配置互相冲突、切换 profile 不解除旧配置，程序整体不稳定。该 Issue 虽已关闭，但其反映的更新可靠性诉求与今日 P1 修复 [#76464](https://github.com/NousResearch/hermes-agent/pull/76464) 一脉相承。

- **[#51603 [CLOSED] 多路复用模式下跨 profile 凭据泄露](https://github.com/NousResearch/hermes-agent/issues/51603)**（5 条评论，security/P2）：详细描述了 `resolve_anthropic_token()` 绕过 profile 密钥作用域、导致 Profile A 密钥可能被 Profile B 请求读取的问题。该 Issue 已由 [#76462](https://github.com/NousResearch/hermes-agent/pull/76462) 修复并关闭，是今日安全方向最受关注的事件。

- **[#43757 [OPEN] Responses API 工具结果在轮次间丢失](https://github.com/NousResearch/hermes-agent/issues/43757)**（3 条评论，P2）：`function_call_output` 在 `input` 数组中未被正确解析，导致多轮工具结果丢失。该问题已开放近两个月，仍无修复 PR，社区关注度持续。

- **[#32887 [OPEN] gateway_state.json 心跳缺失](https://github.com/NousResearch/hermes-agent/issues/32887)**（3 条评论，P3）：空闲超过 2 分钟的 gateway 不更新心跳文件，导致跨容器 WebUI 误判宕机。

**社区诉求分析**：讨论热度最高的问题集中在“更新/升级后的稳定性”与“多 profile 数据隔离”两方面。前者说明更新管道对用户信任影响极大，后者则反映多租户/多 profile 场景已是真实且重要的使用方式，任何边界泄漏都会引发强烈关注。

## 5. Bug 与稳定性

按严重程度排列（含已确认修复）：

| 严重度 | Issue / PR | 状态 | 说明 |
|---|---|---|---|
| P1 | [#76464](https://github.com/NousResearch/hermes-agent/pull/76464) EBADENGINE 导致的更新中断 | ✅ 已合并 | 系统 npm 不满足 engine 约束时自动配置自管 Node 运行时，根治更新死路问题 |
| P2 | [#51603](https://github.com/NousResearch/hermes-agent/issues/51603) 跨 profile 凭据泄露 | ✅ 已修复（[#76462](https://github.com/NousResearch/hermes-agent/pull/76462) 合并） | Anthropic token 绕过 profile 密钥作用域，已在多路复用模式下修复 |
| P2 | [#75141](https://github.com/NousResearch/hermes-agent/issues/75141) profile .env 无法清除继承的 Hermes 环境变量 | ✅ 已关闭 | 相关 ACP/环境变量泄漏已由密钥范围迁移 PR 覆盖 |
| P2 | [#43757](https://github.com/NousResearch/hermes-agent/issues/43757) Responses API 工具结果丢失 | ⚠️ 开放中，无 PR | `function_call_output` 未被正确处理，已影响多轮工具调用 |
| P2 | [#76468](https://github.com/NousResearch/hermes-agent/issues/76468) OmniRoute 503 中止多 agent 回合 | ⚠️ 新开，无 PR | `chat_admission_busy` 被当作普通过载处理，应等待容量恢复而非中止 |
| P2 | [#76469](https://github.com/NousResearch/hermes-agent/issues/76469) Termux 安装找不到 nemo-relay | ⚠️ 新开，无 PR | pip 依赖解析失败，阻塞 Android 端安装 |
| P2 | [#75651](https://github.com/NousResearch/hermes-agent/issues/75651) 安装与同步多个边界问题 | ⚠️ 开放中 | 含 EBADENGINE（已由 #76464 修复）、配置 schema 错误、LXC/macOS Git 同步冻结 |
| P2 | [#76467](https://github.com/NousResearch/hermes-agent/pull/76467) npm 恢复在 no-op 重试时失效 | 🔧 修复 PR 开放中 | 解决失败后重跑 `hermes update` 仍显示 “Already up to date” 但依赖未修复 |
| P2 | [#76459](https://github.com/NousResearch/hermes-agent/pull/76459) Windows 上系统 Node 优先级高于捆绑 Node | 🔧 修复 PR 开放中 | 影响桌面端更新与依赖完整性 |
| P2 | [#76470](https://github.com/NousResearch/hermes-agent/pull/76470) `config set` 把 JSON 数组写成带引号标量 | 🔧 修复 PR 开放中 | 导致 allowlist/telegram env 桥接静默失效 |
| P3 | [#32887](https://github.com/NousResearch/hermes-agent/issues/32887) gateway 心跳缺失 | ⚠️ 开放中，无 PR | 空闲 gateway 被 WebUI 误判为宕机 |
| P3 | [#76064](https://github.com/NousResearch/hermes-agent/issues/76064) 桌面端演示插件默认启用 | ⚠️ 开放中 | 生产构建出现 “clicked N×” 计数器等 UI 噪音 |
| P3 | [#76414](https://github.com/NousResearch/hermes-agent/issues/76414) honcho peers 显示 “(not set)” | ⚠️ 开放中 | host key 使用 `.` 分隔而非 `_`，非默认 profile 全部受影响 |
| P3 | [#76463](https://github.com/NousResearch/hermes-agent/pull/76463) WebGL 图集清除后终端字形错乱 | 🔧 修复 PR 开放中 | 共享图集失效后需刷新兄弟终端 |
| P3 | [#69403](https://github.com/NousResearch/hermes-agent/pull/69403) Windows 下设备/进程读取守卫失效 | 🔧 修复 PR 开放中 | `normpath` 导致路径匹配失败，防护形同虚设 |

**稳定性小结**：今日修复集中在“更新管道”与“密钥隔离”两个高风险区域，且均已合入。当前最需要关注的未修复 P2 是长期存在的 [#43757](https://github.com/NousResearch/hermes-agent/issues/43757)（工具结果丢失）以及两个新报告的安装阻塞（[#76469](https://github.com/NousResearch/hermes-agent/issues/76469)、[#76468](https://github.com/NousResearch/hermes-agent/issues/76468)）。

## 6. 功能请求与路线图信号

- **[#76466 [Feature] TTS 使用端点返回的采样率](https://github.com/NousResearch/hermes-agent/issues/76466)**（新开，P3）：本地 OpenAI 兼容 TTS 返回 44.1 kHz，但 Hermes 硬编码为 24 kHz，导致音频不匹配。这是一个低成本、高兼容性收益的改进，适合纳入后续版本。

- **[#76458 [Feature] 每轮 provider 请求预算](https://github.com/NousResearch/hermes-agent/pull/76458)**（已打上 `duplicate` 标签，仍开放）：提出为每轮用户对话设置物理主模型请求数上限，覆盖重试、fallback、摘要等路径。虽然标记为重复，但“成本控制/请求预算”是明确的功能信号。

- **[#55170 [Feature] Web 端模型 fallback 管理](https://github.com/NousResearch/hermes-agent/pull/55170)**（开放中，P3）：后端 API + Dashboard 面板实现 fallback 提供者增删改与排序。已开放超一个月，等待审查。

- **已合入的功能**：[#7793](https://github.com/NousResearch/hermes-agent/pull/7793)（Telegram PDF 自动摄取 KB）、[#18621](https://github.com/NousResearch/hermes-agent/pull/18621)（`description_full`）、[#76461](https://github.com/NousResearch/hermes-agent/pull/76461)（自定义 systemd 单元名）——这些合入表明项目正在持续强化知识库工作流、技能展示和 Linux 部署运维能力。

## 7. 用户反馈摘要

- **更新稳定性受损（[#75598](https://github.com/NousResearch/hermes-agent/issues/75598)）**：用户 secretgspot 反馈“约一周前开始，更新后整个程序不稳定，多个 gateway 互相冲突，切换 profile 不会停用其他配置”。该用户此前更新一直顺利，说明回归对老用户信任打击较大。

- **演示插件影响生产使用（[#76064](https://github.com/NousResearch/hermes-agent/issues/76064)）**：用户 258692011 指出桌面版默认启用两个 demo/dogfood 插件（“clicked N×”计数器和重复的 gateway 胶囊），在生产构建中显得混乱，获 1 次 👍。

- **TTS 采样率硬编码不匹配（[#76466](https://github.com/NousResearch/hermes-agent/issues/76466)）**：用户 lojack5 使用本地 Echo-TTS 44.1 kHz 端点，但 Hermes 近两周内改为硬编码 24 kHz，造成音频异常。属于近期回归。

- **Android/Termux 安装受阻（[#76469](https://github.com/NousResearch/hermes-agent/issues/76469)）**：用户 zFitness 在 Termux 执行 pip 安装时报 `nemo-relay<0.7,>=0.6.0` 找不到匹配版本，安装被阻断。

- **安全边界关注（[#51603](https://github.com/NousResearch/hermes-agent/issues/51603)、[#75141](https://github.com/NousResearch/hermes-agent/issues/75141)）**：多位用户报告了 profile 间密钥/环境变量隔离不彻底的问题，说明多 profile 部署已较广泛，社区对隔离边界高度敏感。

## 8. 待处理积压

以下问题长期未得到响应或修复，建议维护者重点关注：

- **[#43757](https://github.com/NousResearch/hermes-agent/issues/43757)（P2，开放 53 天）**：Responses API `function_call_output` 丢失，直接影响多轮工具调用，且无关联 PR。作为核心 agent 路径上的缺陷，应优先安排。

- **[#32887](https://github.com/NousResearch/hermes-agent/issues/32887)（P3，开放 67 天）**：`gateway_state.json` 心跳缺失导致 WebUI 误判。虽优先级不高，但已影响 Docker 部署的监控生态，且修复路径清晰。

- **[#55170](https://github.com/NousResearch/hermes-agent/pull/55170)（开放 34 天）**：模型 fallback 管理的 Web 面板 PR 长期未被审查，涉及后端 API 与配置文件写入，建议维护者尽快给出 review 结论。

- **[#69403](https://github.com/NousResearch/hermes-agent/pull/69403)（开放 11 天，P2）**：Windows 上文件读取守卫整体失效的安全加固 PR，等待审查，建议优先合入。

---

**总体评估**：Hermes Agent 今日处于高强度迭代状态，安全修复与更新管道修复均已落地，社区反馈闭环速度快。当前主要风险点集中在两个开放的中级 bug（[#43757](https://github.com/NousResearch/hermes-agent/issues/43757)、[#76468](https://github.com/NousResearch/hermes-agent/issues/76468)）以及一批待审查的修复 PR 上，若保持当前合入节奏，项目健康度有望进一步提升。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目动态日报 — 2026-08-02

> 数据来源：sipeed/picoclaw GitHub 仓库  
> 统计周期：2026-08-01 ~ 2026-08-02（过去24小时）

---

## 1. 今日速览

过去 24 小时内，PicoClaw 项目整体活跃度处于**中等偏上**水平：核心维护动作集中在 PR 合并与代码评审，社区贡献活跃 — 共 3 条 PR 更新，其中 1 条已合并（繁体中文本地化），另有 2 条新功能 PR（Exa 搜索 Provider、OrcaRouter Provider）等待维护者评审。Issue 侧相对平静，仅 1 条存量 Bug (#3203) 获得评论更新，说明维护者与用户在深度讨论该问题的解法方向。无新版本发布，项目处于**功能叠加与稳定性加固并行**的阶段。

> ⚠️ 需要特别关注：#3203 是一个长期未修复的高影响稳定性 bug，已标记 stale，且有持续讨论，建议维护者尽快介入。

---

## 3. 项目进展

今日有 1 个 PR 被合并/关闭，项目本地化能力向前迈进了一步：

- **[#3261 [CLOSED] Add zh-TW locale and Traditional Chinese translations](https://github.com/sipeed/picoclaw/pull/3261)**  
  合并了繁体中文（台湾地区）的界面与文档翻译。该 PR 统一了 WebUI 及文档中的台湾术语习惯，使本地化体验覆盖至初始配置和频道引导流程。  
  **关联推进**：意味着 PicoClaw 的 i18n 基础设施已支持除简体中文外的第二个中文变体，为后续更多语言扩展铺平了路径。

另外 2 个新 PR 处于待合并状态，虽然尚未合入，但已经进入维护者的视野，预计将在下一次版本中贡献重要功能：

- [#3299 Add native Exa web search provider](https://github.com/sipeed/picoclaw/pull/3299)
- [#3309 feat(providers): add OrcaRouter as an OpenAI-compatible provider](https://github.com/sipeed/picoclaw/pull/3309)

这说明项目正在加速扩展外部搜索与模型路由能力，**整体技术方向明确，社区提交质量较高**。

---

## 4. 社区热点

**最受关注：Issue #3203 — Matrix sync loop 无重连逻辑**  
[#3203 [BUG] Matrix sync loop has no reconnection logic — silent death after network/server disruption](https://github.com/sipeed/picoclaw/issues/3203)

- 评论数：7（今日仍在更新）
- 👍 数：2
- 状态：OPEN（已标记 stale）

这是过去 24 小时内讨论热度最高的问题。尽管 issue 创建于 7 月 2 日，但用户在 8 月 1 日仍在追问进展，说明该问题尚未解决且持续影响实际用户。它被标记为 *stale* 却仍在产生评论，本身就释放了一个值得注意的信号——**核心的健壮性问题在长期未获修复的情况下，用户情绪可能会逐渐累积**。

**诉求分析**：用户的核心诉求是生产环境下的**可靠性与可恢复性**。Matrix 是 PicoClaw 的核心通信渠道，一旦房间同步静默失去连接，整个 Bot 处于“僵尸”状态，而 systemd 又由于主进程存活而无法感知故障。用户不是要求新功能，而是要求最基础的“跌倒后能自己爬起来”。

---

## 5. Bug 与稳定性

今日活跃的 Bug 仅有 1 个，但严重程度较高：

| 严重程度 | Issue | 描述 | 修复 PR |
|---------|-------|------|---------|
| 🔴 高 | [#3203](https://github.com/sipeed/picoclaw/issues/3203) | Matrix `/sync` 长轮询在断网或服务器重启后永久死亡，无自动重连；系统因主进程存活而无法触发 systemd 的 `Restart=on-failure`，导致**静默性停机** | ❌ 暂无 |

**影响评估**：此 Bug 涉及“自杀式”故障模式——进程看似存活，实则完全失能，监控手段（systemd、健康检查）均无法感知。对依赖 Matrix 作为唯一通知渠道的用户（如个人助理、家庭自动化），意味着服务中断可能数天而无人察觉。该 issue 从 2026-07-02 创建至今已满一个月，尚无对应修复 PR，**建议维护者将该问题升级为 P1 优先处理**。

此外今日无崩溃、回归或新引入的 Bug 报告。

---

## 6. 功能请求与路线图信号

今日没有提交新的 Feature Request，但有 2 个功能型 PR 等待合入，是路线图的重要信号：

- **[#3299 Add native Exa web search provider](https://github.com/sipeed/picoclaw/pull/3299)**  
  将 Exa 作为原生 `tools.web` / `web_search` 提供方接入。使用 Exa 的 `POST /search` API（`type: "auto"` + `contents.highlights`），通过 `X-Api-Key` 认证，并支持现有 `d`/`w`/`m`/`y` 时间范围过滤。这表明社区希望扩展 **Web 搜索能力的多样性和特定领域的搜索质量**（Exa 在深度/语义搜索方面具有优势）。

- **[#3309 feat(providers): add OrcaRouter as an OpenAI-compatible provider](https://github.com/sipeed/picoclaw/pull/3309)**  
  增加 OrcaRouter 作为“一等公民”的 OpenAI 兼容 Provider（`orcarouter`）。OrcaRouter 是一个多供应商路由器，在 `https://api.orcarouter.ai/v1` 上实现 OpenAI Chat Completions 契约，通过 `vendor/model` ID 寻址上游模型。  
  这延续了 PicoClaw 在聚合多模型路由领域的投入——**与现有 provider 体系无缝衔接，用户可零成本新增一个多模型网关**。

**下一版本预测**：若这两个 PR 在下一周期合并，v0.3.0 或 v0.2.10 将具备以下新能力：① 更多可选的 Web 搜索后端；② 通过 OrcaRouter 接入大量上游模型（可能包括各厂商的最新模型）。整体路线图信号指向 **“更丰富的工具生态 + 更灵活的模型接入”**。

---

## 7. 用户反馈摘要

> 来源：Issue #3203 的评论讨论（共 7 条）

1. **痛点：静默故障最难排查**  
   用户指出，Matrix 同步循环死后“一切看起来都正常”（进程存在、日志无输出），但实际上 Bot 已完全不响应。用户希望在**发生此类情况时有明确的可观测信号**（如日志告警、状态上报）。

2. **痛点：对 systemd 重启机制的失灵感到困扰**  
   用户反馈依赖 `Restart=on-failure` 来维持服务可用性，但由于主进程未退出，该机制完全失效。这暴露了**进程级健康检查对于长连接型任务的无能为力**。

3. **期望：一套稳健的重连退避策略**  
   讨论中用户倾向于支持“指数退避 + 无限重试”的通用重连方案，而不是依赖外部 process manager。这间接反映了用户对**内置韧性的偏爱**——即 Bot 自身应具备自恢复能力，而非依赖运维侧兜底。

4. **对项目整体的印象**  
   虽然本次问题尚未修复，但用户对 PicoClaw 的整体功能感到满意（从参与讨论的语境推断），愿意通过 issue 详细描述环境信息（v0.2.9、Cha...）来帮助复现，说明社区对项目的信任度较高。

---

## 8. 待处理积压

以下为当前值得维护者关注的长期未决事件：

- **[#3203 [stale] Matrix sync loop has no reconnection logic](https://github.com/sipeed/picoclaw/issues/3203)**  
  创建：2026-07-02 | 更新：2026-08-01 | 已 stale  
  **重要性**：高。核心通讯链路的可靠性问题，直接影响生产环境中的服务可用性。已与维护者未交互超过一个月且被 stale-bot 标记，建议手动移除 stale 标签并给出修复计划。

- **[#3299 [OPEN] Add native Exa web search provider](https://github.com/sipeed/picoclaw/pull/3299)**  
  创建：2026-07-26 | 已等待 7 天  
  **重要性**：中高。功能型 PR，代码质量似乎完整，等待维护者 review 或发表修改意见。长时间无响应会降低贡献者积极性。

- **[#3309 [OPEN] feat(providers): add OrcaRouter as an OpenAI-compatible provider](https://github.com/sipeed/picoclaw/pull/3309)**  
  创建：2026-08-01 | 等待 review  
  **重要性**：中高。新 Provider 的扩展对社区价值明确，且创建者 `jinhaosong-source` 尚未有任何互动，需要维护者快速响应以维持健康的协作节奏。

---

## 📌 总结评估

| 维度 | 评级 | 说明 |
|------|------|------|
| 社区活跃度 | ⭐⭐⭐⭐ | PR 提交活跃，issue 讨论深度高 |
| 项目健康度 | ⭐⭐⭐ | 功能进展顺利，但关键 bug 积压超30天 |
| 维护响应速度 | ⭐⭐⭐ | 今日有 PR 合并，但对存量 issue 响应偏慢 |
| 路线图清晰度 | ⭐⭐⭐⭐ | 搜索 + 多 Provider 的方向明确 |

**维护者建议**：优先对 #3203 做出响应（哪怕是声明修复计划），并尽快 review #3299 和 #3309 两个功能 PR——前者将塑造一项核心可靠性功能，后者将直接增加项目的生态适配度和可用性广度。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目动态日报 — 2026-08-02

## 1. 今日速览

过去 24 小时 NanoClaw 处于高强度迭代状态：发布 1 个滚动版本 v2.1.54，合并/关闭 6 个 PR，另有 10 个 PR 处于待合并；Issue 侧 2 条更新，1 新开、1 关闭。核心事件是 **iMessage 通道统一（Local + Hosted 双后端）正式落地 v2.1.54**，同时修复了 setup 流程对非 Claude 用户的误导性引导、新增凭据过期告警。社区贡献活跃，Denver901、petrolette、soren5 等外部贡献者均有新 PR 提交；但项目存在一个已破坏 `migrate-v2` 的回归 Bug（#3166）和一个技能包越权拦截编码请求的问题（#3171），需优先处理。整体看项目推进速度快，合入质量较高，积压 PR 的 review 效率是当前主要瓶颈。

## 2. 版本发布

**v2.1.54**（滚动发布）

- 范围：覆盖 v2.1.18 至 v2.1.54 之间所有合并内容，即自 v2.1.17 tag 以来的全部改动。
- **破坏性变更**：iMessage 统一为单一 `imessage` 通道，采用双后端架构，通过 `/add-imessage` 安装：
  - **Local**：基于 Chat SDK 读取本机 Mac 的 `chat.db`；
  - **Hosted**：基于 [Photon](https://photon.codes) 的原生托管服务。
- **迁移注意**：已有 iMessage 集成的用户需重新运行 `/add-imessage` 完成新通道配置；旧的通道/技能结构将被新架构取代，相关 PR（#2999、#3164）已合并。

## 3. 项目进展

今日合并/关闭的 6 个 PR，对应 4 项实质进展：

- **iMessage 统一通道落地**（[#2999](https://github.com/nanocoai/nanoclaw/pull/2999)、[#3164](https://github.com/nanocoai/nanoclaw/pull/3164)）：#2999 提出将 iMessage 统一为单一通道（Local 基于 Chat SDK + Hosted Photon 双后端）；#3164 作为替代实现，补充了可用的 Photon Hosted 注册流程并标注 supersede #2999。两者构成 v2.1.54 的核心特性，打通了 iMessage 从本地到托管的多场景接入。
- **Setup 失败引导修正**（[#3170](https://github.com/nanocoai/nanoclaw/pull/3170)）：将失败诊断辅助从"总是引导安装 Claude CLI"改为按用户所选 provider 分发，闭环关闭 Issue #3169。
- **凭据过期告警**（[#3167](https://github.com/nanocoai/nanoclaw/pull/3167)）：新增 provider 凭据过期告警。此前 Codex ChatGPT 凭据过期仅表现为容器反复被杀和 `Read-only file system (os error 30)`，运维侧无法定位根因。
- **发布流程安全加固**（[#3168](https://github.com/nanocoai/nanoclaw/pull/3168)）与 **Codex/copilot 变更**（[#3165](https://github.com/nanocoai/nanoclaw/pull/3165)）：前者关闭 release 合并后的安全缺口，后者包含 Codex/Copilot 相关调整。

## 4. 社区热点

今日所有 Issue/PR 的评论数均为 0，公开讨论热度偏低；但以下条目影响面最大、最值得关注：

- **[#3171【OPEN】两个 qodo 技能依赖未配置的集成并拦截正常编码请求](https://github.com/nanocoai/nanoclaw/issues/3171)**：`get-qodo-rules` 和 `qodo-pr-resolver` 读取 `~/.qodo/config.json` 的 API key，但仓库内没有任何机制会创建该配置，导致技能在用户未配置 Qodo 时就被激活并拦截正常请求。核心矛盾是**开箱即用体验与内置技能默认启用之间的冲突**，并已催生移除型修复 PR（#3172）。
- **[#3169【CLOSED】非 Claude 安装的 setup 失败总是引导安装 Claude CLI](https://github.com/nanocoai/nanoclaw/issues/3169)**：用户选择 codex 等非 Claude provider 后，失败提示仍默认预选 "Yes" 引导安装 Claude CLI；若 Claude 已装未登录还会拉起 Anthropic 登录。用户显式选择被忽略，属于流程设计缺陷，已由 #3170 修复。
- **[#3174【OPEN】支持 rootless Docker 运行 agent 容器](https://github.com/nanocoai/nanoclaw/pull/3174)**：外部贡献者 Denver901 报告 agent 容器在 rootless Docker daemon 下完全不可用，且两个独立故障在宿主用户属于 `docker` 组时完全不可见。这类来自真实部署环境的反馈对项目健壮性很有价值。

共同诉求：安装配置流程应尊重用户显式选择，内置技能不应在未配置对应集成时"自作主张"抢占用例。

## 5. Bug 与稳定性

按严重程度排列：

| 严重度 | 问题 | 状态 | 链接 |
|---|---|---|---|
| **高** | `migrate-v2` 完全损坏：`tasks.ts` 静态导入已被删除的 `insertTask`（应为 `insertTaskRow`），迁移步骤直接抛 SyntaxError，无法执行 | 修复 PR [#3166](https://github.com/nanocoai/nanoclaw/pull/3166) 待合并 | [PR #3166](https://github.com/nanocoai/nanoclaw/pull/3166) |
| **中高** | 两个 qodo 技能依赖不存在的 Qodo 集成，且拦截正常编码请求 | Issue [#3171](https://github.com/nanocoai/nanoclaw/issues/3171)；修复 PR [#3172](https://github.com/nanocoai/nanoclaw/pull/3172)（移除两技能）待合并 | [Issue #3171](https://github.com/nanocoai/nanoclaw/issues/3171) |
| **中** | rootless Docker 下 agent 容器不可用，存在两个独立故障，对 docker 组成员不可见 | 修复 PR [#3174](https://github.com/nanocoai/nanoclaw/pull/3174) 待合并 | [PR #3174](https://github.com/nanocoai/nanoclaw/pull/3174) |
| **中** | 凭据过期时错误信息严重误导（表现为 `Read-only file system (os error 30)`），无法定位根因 | 已由 #3167 凭据告警功能修复 | [PR #3167](https://github.com/nanocoai/nanoclaw/pull/3167) |
| **中** | 容器 SIGKILL 后 `outbound.db` 陈旧 journal 未恢复；热 journal 轮询存在竞态（对应 #2516/#2640） | 修复 PR [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) 开放已 7 周 | [PR #2750](https://github.com/nanocoai/nanoclaw/pull/2750) |
| **中** | router 对未受信输入解析不严：`safeParseContent` 对 primitive JSON 直接返回，调用方读 `.text`/`.sender` 得到 `undefined` | 修复 PR [#2801](https://github.com/nanocoai/nanoclaw/pull/2801) 开放已 6 周 | [PR #2801](https://github.com/nanocoai/nanoclaw/pull/2801) |
| **低** | agent 通过 `send_message` 发送后又重复正文，导致消息投递两次 | 修复 PR [#2956](https://github.com/nanocoai/nanoclaw/pull/2956) 待合并 | [PR #2956](https://github.com/nanocoai/nanoclaw/pull/2956) |

其中 **#3166 为回归问题**（`insertTask` 重命名为 `insertTaskRow` 后未同步更新调用方），直接阻断用户升级路径，建议优先合并。

## 6. 功能请求与路线图信号

- **iMessage 统一通道**（[#2999](https://github.com/nanocoai/nanoclaw/pull/2999)/[#3164](https://github.com/nanocoai/nanoclaw/pull/3164)）：已随 v2.1.54 发布，是近期路线图的最大变化。后续版本预计会继续完善 Hosted（Photon）后端的注册流程与稳定性。
- **凭据过期告警**（[#3167](https://github.com/nanocoai/nanoclaw/pull/3167)）：已合入，释放了项目在**可观测性与运维体验**上的投入信号。
- **rootless Docker 支持**（[#3174](https://github.com/nanocoai/nanoclaw/pull/3174)）：若合并，agent 容器的部署环境覆盖面将显著扩大，有望进入 v2.1.55 或 v2.2。
- **qodo 技能移除**（[#3172](https://github.com/nanocoai/nanoclaw/pull/3172)）：若采用"直接移除"方案，说明项目在**收敛内置技能体积**，避免未配置集成干扰用户；未来更可能以显式按需安装的方式回归。
- **reaction 投递 best-effort**（[#3121](https://github.com/nanocoai/nanoclaw/pull/3121)）：让 reaction 投递失败不阻塞主流程，属健壮性改进，可能纳入下一个稳定性版本。

## 7. 用户反馈摘要

> 注：今日 Issue/PR 均无评论，以下反馈提炼自 Issue 正文与 PR 描述中的真实场景。

- **配置流程不尊重用户选择**（[#3169](https://github.com/nanocoai/nanoclaw/issues/3169)）：用户显式选择非 Claude provider，失败引导仍预选安装 Claude CLI，甚至拉起 Anthropic 登录。**核心痛点：诊断辅助应针对当前所选 provider，而非默认退回 Claude。**
- **内置技能"越权"**（[#3171](https://github.com/nanocoai/nanoclaw/issues/3171)）：用户未配置 Qodo，但两个内置技能仍读取 `~/.qodo/config.json` 并拦截正常编码请求。**核心痛点：未配置对应集成的技能不应被默认激活。**
- **错误信息不可读**（[#3167](https://github.com/nanocoai/nanoclaw/pull/3167) 事件描述）：凭据过期表现为容器反复重启和 `Read-only file system (os error 30)`，运维侧无法识别真实原因。**核心痛点：错误传播链路需要补充根因上下文。**
- **隐蔽的环境假设**（[#3174](https://github.com/nanocoai/nanoclaw/pull/3174)）：代码假设宿主用户属于 `docker` 组，不满足时静默失败，直到贡献者刻意以非 docker 组成员身份运行才暴露。**核心痛点：环境前置条件应显式检查或在文档中明确。**

## 8. 待处理积压

以下开放 PR 长期无人评论，建议维护者优先 Review（按开放时长排序）：

- **[#2750](https://github.com/nanocoai/nanoclaw/pull/2750)**：6 月 12 日开放，修复容器 SIGKILL 后 `outbound.db` 陈旧 journal 及热 journal 轮询竞态，关联 #2516/#2640 两个已知 Issue。
- **[#2801](https://github.com/nanocoai/nanoclaw/pull/2801)**：6 月 17 日开放，加固 router 对未受信输入的处理。
- **[#2956](https://github.com/nanocoai/nanoclaw/pull/2956)**：7 月 5 日开放，修复 agent 最终输出与工具发送内容重复导致的双重投递。
- **[#3046](https://github.com/nanocoai/nanoclaw/pull/3046)**：7 月 14 日开放，文档与当前 Telegram 配对状态块对齐。
- **[#3090](https://github.com/nanocoai/nanoclaw/pull/3090)**：7 月 19 日开放，修复模板顶层 context Markdown 处理。
- **[#3121](https://github.com/nanocoai/nanoclaw/pull/3121)**：7 月 23 日开放，reaction 投递改为 best-effort。
- **今日新增待 review**：[#3166](https://github.com/nanocoai/nanoclaw/pull/3166)（migrate-v2 崩溃修复，高优先级）、[#3172](https://github.com/nanocoai/nanoclaw/pull/3172)（移除 qodo 技能）、[#3173](https://github.com/nanocoai/nanoclaw/pull/3173)（egress 更新）、[#3174](https://github.com/nanocoai/nanoclaw/pull/3174)（rootless Docker）。

积压风险提示：#2750、#2801 等老 PR 已开放 6–7 周，与主线代码的冲突风险持续上升，且所涉均为已知 Issue 的修复，建议尽快安排 Review 或明确关闭理由。

---

*本日报基于 GitHub 公开数据自动生成，数据时间窗口为 2026-08-01 00:00 至 2026-08-02 00:00（UTC）。*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw 项目动态日报 — 2026-08-02

## 今日速览

IronClaw 过去 24 小时保持高活跃度：20 条 Issue 更新、25 条 PR 更新，其中至少 5 条已知 PR 完成合并/关闭。WS2 契约反转工程明显提速，[#6998](https://github.com/nearai/ironclaw/pull/6998)（WS2.1）与 [#7002](https://github.com/nearai/ironclaw/pull/7002)（WS5）相继合并，端口边界持续向 `product_contracts` 收敛。另一条主线是 pi-harness 采用计划：一口气提出 7 个 P0/P1 缓存与预算任务（[#6984](https://github.com/nearai/ironclaw/issues/6984)–[#6990](https://github.com/nearai/ironclaw/issues/6990)），并已有 [#6997](https://github.com/nearai/ironclaw/pull/6997)、[#7001](https://github.com/nearai/ironclaw/pull/7001) 两个对应 PR 进入审查。风险面集中在性能与 CI：libSQL 工具重负载场景 p95 最高达 135 秒（[#6974](https://github.com/nearai/ironclaw/issues/6974)），reborn-tests 的 `workflow_dispatch` 存在结构性失败（[#6978](https://github.com/nearai/ironclaw/issues/6978)）。整体交付节奏快、重构纪律强，但需警惕性能回归与超长待办（如 [#5598](https://github.com/nearai/ironclaw/pull/5598)）的累积。

## 版本发布

过去 24 小时无新版本发布。

## 项目进展

已知合并/关闭的 PR：

- **[#6996](https://github.com/nearai/ironclaw/pull/6996)** — 关闭 [#6963](https://github.com/nearai/ironclaw/issues/6963)，以清点驱动方式完成剩余 path-keyed CI 门禁的 fail-closed 改造（六个静默门禁 + 两个 loud-but-flat-keyed 重定向），CI 门禁完整性大幅提升。
- **[#6998](https://github.com/nearai/ironclaw/pull/6998)** — WS2.1：`ironclaw_extension_host` 的产品端口定义从 `ironclaw_product` 反转为 `ironclaw_product_contracts`，行为零变更，是 Wave 2 排序约束最重的半场。
- **[#7002](https://github.com/nearai/ironclaw/pull/7002)** — WS5：webui + openai_compat 传输端口反转至 `product_contracts`，与 #7000 的 union 决议已自动合并进分支。
- **[#6995](https://github.com/nearai/ironclaw/pull/6995)** — Wave 1 真相审计：将 `docs/reborn/target-architecture/` 的决策记录与合并后的 `main`（`a50ad0638`）校准，修正文档漂移。
- **[#6761](https://github.com/nearai/ironclaw/pull/6761)** — 新贡献者 ogarciarevett 贡献的通用出站注册回归测试，覆盖 registry 边界查询，防止 `register_generic_channel_outbound_targets` 被误改成 no-op。

这些合并标志着 Wave 2 重构已有 2/7 个槽位正式落地，叠加 [#6995](https://github.com/nearai/ironclaw/pull/6995) 对 Wave 1 的收尾审计，项目正按 CHECKLIST 有条不紊地向目标架构推进。此外，WS2 后续叠堆链（[#7000](https://github.com/nearai/ironclaw/pull/7000) → [#7003](https://github.com/nearai/ironclaw/pull/7003) → [#7004](https://github.com/nearai/ironclaw/pull/7004) → [#7005](https://github.com/nearai/ironclaw/pull/7005)）均在审查中，契约层重构已进入深水区。

## 社区热点

- **[Issue #6963](https://github.com/nearai/ironclaw/issues/6963)**（7 条评论，已关闭）— Path-keyed CI 门禁追踪问题。讨论围绕 #6946 未覆盖的八处门禁缺陷展开，最终由 [#6996](https://github.com/nearai/ironclaw/pull/6996) 以清单驱动方式全部修复。诉求核心：门禁改写不能留下“静默失效”的盲区，CI 必须是 fail-closed 的。
- **[Issue #6974](https://github.com/nearai/ironclaw/issues/6974)**（2 条评论，OPEN）— libSQL `thread_store_writes` 性能病态：工具密集型场景 p95 37–135 秒，远高于 2.5s 目标。此问题从 [#6973](https://github.com/nearai/ironclaw/pull/6973)（Postgres 容量恢复）中拆分独立追踪，说明维护者对“不同存储后端的性能问题需要独立跟踪”有清晰认知。
- **[Issue #6921](https://github.com/nearai/ironclaw/issues/6921)**（2 条评论，已关闭）— 提取中立 loop/extension/product 契约并密封 evidence minting。与 WS2.1 落地一致，属于架构边界收敛的收尾任务。

PR 侧今日无评论数据，讨论热度集中在 Issue 侧。整体来看，开发者对 CI 门禁可靠性和性能瓶颈最为敏感，这二者是当前影响交付体验的主要因素。

## Bug 与稳定性

按严重程度排列：

**高严重度**

- **[#6974](https://github.com/nearai/ironclaw/issues/6974)**（OPEN）— libSQL 写入路径性能病态：ops 20 / c4 / u50 套件下 p95 37–135 秒。`main` 分支甚至无法在 20 分钟 CI 超时内完成 large-context prefill。无直接 fix PR；[#6973](https://github.com/nearai/ironclaw/pull/6973) 解决的是 Postgres 侧容量恢复，libSQL 需单独跟进。
- **[#6978](https://github.com/nearai/ironclaw/issues/6978)**（OPEN）— `reborn-tests.yml` 的 `workflow_dispatch` 运行结构性失败：`critical-mutation` 的 `if:` 条件排除了手动触发（`reborn-tests.yml:788-793`），但 roll-up 又禁止跳过该作业，导致手动触发永远红色。由 [#6977](https://github.com/nearai/ironclaw/pull/6977) 的 dispatch 运行 30665278857 实证。

**中严重度**

- **[#7006](https://github.com/nearai/ironclaw/issues/7006)**（OPEN）— 变更覆盖率门禁的集成层缺口：steering-queue 切片约 180 行错误路径（fault injection、CAS 冲突、序列化失败等）在 hermetic integration harness 中无法执行，导致 [#5981](https://github.com/nearai/ironclaw/pull/5981) 被门禁卡住。
- **[#7011](https://github.com/nearai/ironclaw/issues/7011)**（OPEN）— `extension_manager` 五个预先存在的问题：false `WriteFilesystem` effect、未测试的锁谓词、两个缺失的 dispatch 测试、六个被丢弃的 cause。全部位于 #7003 字节级平移的代码中，属存量缺陷转移。
- **[#6999](https://github.com/nearai/ironclaw/issues/6999)**（OPEN）— `reborn_dependency_boundaries` 的 server-lifecycle 规则从未覆盖其文档声称的 WebChat v2 路由表面。关闭它需要架构决策而非简单门禁重指向。

**已修复/已关闭**

- **[#6903](https://github.com/nearai/ironclaw/issues/6903)**（CLOSED）— 管理员用户列表超过 100 人时无法加载下一页，后端 `next_cursor` 未被消费。已修复。
- **[#6963](https://github.com/nearai/ironclaw/issues/6963)**（CLOSED）— 八个 CI 门禁缺陷已由 [#6996](https://github.com/nearai/ironclaw/pull/6996) 全部修复。

## 功能请求与路线图信号

- **[#7009](https://github.com/nearai/ironclaw/issues/7009)**（社区贡献）— 请求将 OrcaRouter 加入内置 LLM provider 列表。`providers.json` 已覆盖 OpenRouter、Together、Fireworks 等主流网关，OrcaRouter 的缺失是目前唯一变通使用场景，很可能在下一版本纳入。
- **[#6993](https://github.com/nearai/ironclaw/issues/6993)**（OPEN）— OOBE automation-tasks 原型的后端接线。前端已通过 [#6994](https://github.com/nearai/ironclaw/pull/6994) 以 UI-only 原型落地（基于 mock 数据），后端打通是自然下一步。
- **[#6983](https://github.com/nearai/ironclaw/issues/6983)**（p2, 用户反馈）— 为 `ironhub` CLI 子命令添加 `hub` 别名，以兼容 IronHub dashboard 场景。低风险、收益明确的小改动。
- **pi-harness 采用计划**（[#6984](https://github.com/nearai/ironclaw/issues/6984)–[#6990](https://github.com/nearai/ironclaw/issues/6990)）— 7 个 P0/P1 任务，覆盖显式 cache_control 断点、prompt 前缀字节稳定、工具数组 byte-identical、token 核算修正、上下文预算动态化。其中 [#6984](https://github.com/nearai/ironclaw/issues/6984) 和 [#6985](https://github.com/nearai/ironclaw/issues/6985) 已有对应 PR（[#6997](https://github.com/nearai/ironclaw/pull/6997)、[#7001](https://github.com/nearai/ironclaw/pull/7001)），路线图正在快速转化为实现。

## 用户反馈摘要

- **[#6983](https://github.com/nearai/ironclaw/issues/6983)** — 用户在为 IronHub 编写发布文档时发现 canonical 子命令为 `ironhub`（别名 `iron-hub`），缺少更直观的 `hub` 别名，影响 dashboard 兼容与 CLI 可发现性。
- **[#7009](https://github.com/nearai/ironclaw/issues/7009)** — 用户需要 OrcaRouter 作为内置 provider，当前只能通过通用/变通方式接入，体验割裂。
- **[#6974](https://github.com/nearai/ironclaw/issues/6974)** — 性能痛点显著：在工具密集型会话中，p95 延迟达到 37–135 秒，远超 2.5s 目标，对真实 agent 工作负载影响严重。
- **[#6903](https://github.com/nearai/ironclaw/issues/6903)**（已修复）— 管理员后台用户列表超过 100 人后无法滚动加载，直接阻断大规模租户管理场景。

## 待处理积压

- **[#5598](https://github.com/nearai/ironclaw/pull/5598)** — `chore: release` PR 自 7 月 3 日开放至今已一个月，包含 `ironclaw_common`（0.4.2→0.5.0）和 `ironclaw_skills`（0.3.0→0.4.0）的 breaking changes。发布流程可能已被其他工作阻塞，值得维护者优先处理。
- **[#5981](https://github.com/nearai/ironclaw/pull/5981)** — queued-message steering（XL，7 月 11 日开放）已完成前向移植和 turn-boundary 竞态修复，但被 [#7006](https://github.com/nearai/ironclaw/issues/7006) 的覆盖率门禁卡住，需要门禁规则侧给出豁免或补充测试策略。
- **[#5982](https://github.com/nearai/ironclaw/pull/5982)** — 预算审批-as-blocked-gate PR（XL，7 月 11 日开放）依赖 #5981 先合并，处于排队状态。
- **[#6780](https://github.com/nearai/ironclaw/pull/6780)** — deep-link 注册/安装网关 + 私有 manifest 源（XL，7 月 28 日开放），为 @neo-sky 设计的 re-port，等待审查。
- **WS2 叠堆链** — [#6998](https://github.com/nearai/ironclaw/pull/6998)（已合并） → [#7000](https://github.com/nearai/ironclaw/pull/7000) → [#7003](https://github.com/nearai/ironclaw/pull/7003) → [#7004](https://github.com/nearai/ironclaw/pull/7004) → [#7005](https://github.com/nearai/ironclaw/pull/7005)，中间节点全部 OPEN。叠堆深度已达 4 层，建议维护者按序加速合并以降低冲突风险。

---

**项目健康度小结**：重构执行率高，CI 治理意识强，社区贡献通道顺畅（新贡献者 PR 可正常合入）。主要风险集中在性能回归（libSQL）与长周期 PR 积压（release、queued-message steering），建议下一阶段优先处理这两类问题。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 — 2026-08-02


## 1. 今日速览

过去24小时内，LobsterAI 仓库完成了大规模 stale 清理：6 条超期未更新 Issue 被自动关闭（标记 [stale]），暂无新 Issue 产生，社区活跃度主要来自历史积压项的收尾。PR 侧有 2 条待合并，其中一条为 2026-04-01 提出的 i18n/UX 修复（#1224），等待周期较长；另一条为 7 月中提出的会话重命名失败反馈改进（#2358）。无新版本发布。总体来看，项目当前处于稳定维护期，社区提问热度放缓，但仍有重要修复等待合入，建议维护者尽快处理待合并 PR 以激活项目节奏。


## 3. 项目进展

今日无 PR 被合并或关闭，2 条 PR 仍处于待合并状态：

- **#1224 fix(agent): 修复 i18n 硬编码、Agent 弹窗 Escape 键支持及删除防重复点击**（[netease-youdao/LobsterAI PR #1224](https://github.com/netease-youdao/LobsterAI/pull/1224)）— 2026-04-01 提出，关联 Closes #1223。解决 CoworkPromptInput 硬编码中文「输入文件」导致英文用户提示词混入中文的问题，同时为 Agent 创建/设置弹窗添加 Escape 键关闭和删除按钮防重复点击保护。该 PR 已积压 4 个月，建议维护者安排评审合并。

- **#2358 fix(cowork): show feedback when session rename fails**（[netease-youdao/LobsterAI PR #2358](https://github.com/netease-youdao/LobsterAI/pull/2358)）— 2026-07-18 提出，关联 Fixes #670。为会话重命名失败场景补充本地化错误提示，避免用户无感知标题未保存。等待约 2 周，内容已齐备。


## 4. 社区热点

今日无高热度新讨论，所有更新均为 stale 自动关闭。从评论数据（各 2 条）及 👍 来看，过去 24 小时讨论最集中的仍是被关闭的历史 Issue：

- **#1293 自定义 studio http 的 mcp 无法使用**（[netease-youdao/LobsterAI Issue #1293](https://github.com/netease-youdao/LobsterAI/issues/1293)）— 👍 1，讨论核心：自定义 MCP 未在 OpenClaw 引擎内更新导致无法被调用，仅 SSE 类型可用。反映 MCP 配置在实际使用中与引擎对接存在短板。
- **#1223 CoworkPromptInput 硬编码中文标签 + Agent 弹窗 UX 问题**（[netease-youdao/LobsterAI Issue #1223](https://github.com/netease-youdao/LobsterAI/issues/1223)）— 仍处于 OPEN 且 stale 状态，关联 PR #1224 待合并。用户对国际化与交互细节有持续关注，评论表示已在同一 PR 中修复，等待上游合入。

整体社区热点诉求集中在：**MCP 配置一致性** 与 **国际化/UX 细节打磨**。


## 5. Bug 与稳定性

今日报告的 Bug 均为 stale 自动关闭的历史问题，按严重程度排列如下：

| 严重度 | Issue | 问题描述 | 状态 | 关联 PR |
|--------|-------|----------|------|---------|
| 🔴 高 | #1296 上传长图（3M）解析报错 | 3M 长图解析时页面直接报错，新开任务持续报错，整体功能不可用 | 已关闭（[Issue #1296](https://github.com/netease-youdao/LobsterAI/issues/1296)） | 未发现对应 PR，需关注回归 |
| 🟠 中 | #1298 输入两字提示超出模型限制 | 测试连接通过后，仅输入少量文字即提示「输入内容过长，超出模型限制」，疑似上下文计算偏移 | 已关闭（[Issue #1298](https://github.com/netease-youdao/LobsterAI/issues/1298)） | 未发现对应 PR，需关注回归 |
| 🟠 中 | #1307 关闭编辑面板后无法编辑其他 provider 配置 | 打开再关闭某个模型 provider 配置面板后，切换其他 provider 变为只读 | 已关闭（[Issue #1307](https://github.com/netease-youdao/LobsterAI/issues/1307)） | 未发现对应 PR，需关注回归 |
| 🟡 低 | #1305 定时任务删除后历史标题展示错误 | 定时任务运行成功后删除，历史记录中标题名称展示不对 | 已关闭（[Issue #1305](https://github.com/netease-youdao/LobsterAI/issues/1305)） | 未发现对应 PR |

上述 Issue 均为 stale 自动关闭，不代表已修复。建议维护者排查是否存在未处理的回归风险，必要时重新打开跟踪。


## 6. 功能请求与路线图信号

今日关闭的 Issue 中无新的功能请求提出，但以下待合并 PR 承载了潜在的路线图信号：

- **本地化 / i18n 完善** — PR #1224 将硬编码中文替换为 i18n key，体现项目在多语言支持上的持续投入，预计后续会有更多字符串国际化的推进。
- **交互体验优化** — PR #1224 中 Escape 关闭弹窗、删除防重复点击；PR #2358 中会话重命名失败反馈，均属于「细节体验打磨」方向，若合入将提升桌面端使用舒适度。
- **MCP 生态完善** — Issue #1293 暴露了 http 类型 MCP 与 OpenClaw 引擎的兼容问题，虽已关闭，但 MCP 作为接入外部工具的关键渠道，若用户在真实场景持续遇到，未来可能成为路线图重点。


## 7. 用户反馈摘要

以下为今日被清理的 Issue 及关联评论中提取的真实用户声音：

- **MCP 配置体验受阻**：「自定义的 mcp 实际未在 openclaw 引擎里更新，导致无法被调用。只有 sse 的可以被 openclaw 引擎使用。」（来自 #1293）——用户对于「配置了但未生效」的体验十分挫败，期望引擎与配置层的数据同步能保持一致。
- **长图解析直接不可用**：「让模型解析，页面返回报错，新开任务会一直报错，整体不可用了」（#1296）——高影响场景，直接导致应用不可用，极大影响用户信心。
- **上下文长度计算疑似异常**：「输入两个字的问题，页面直接提示输入内容过长，超出模型限制」（#1298）——用户对于长度限制的判定标准感到困惑，可能是 token 计算或系统提示词注入导致的膨胀。
- **历史记录展示不一致**：「定时任务运行成功后删除，去历史 tab 检查，标题展示不对」（#1305）——删除后历史标题展示不对，影响用户对运行记录的追踪与审计。
- **编辑面板状态残留**：「打开再关闭配置面板后，切换其他 provider 无法编辑」（#1307）——面板状态未重置，交互上的「卡住」感明显。


## 8. 待处理积压

以下 Issue/PR 长期未得到响应或合入，建议维护者重点关注：

| 类型 | 编号 | 标题 | 等待时长 | 说明 |
|------|------|------|----------|------|
| PR | #1224 | fix(agent): i18n 硬编码、Escape 键、防重复点击 | 约 4 个月（2026-04-01 创建） | 有完整实现，Closes #1223，但始终未合并 |
| Issue | #1223 | CoworkPromptInput 硬编码中文 + Agent 弹窗 UX | 约 4 个月（2026-04-01 创建） | 仍为 OPEN 状态，虽已 stale 但修复 PR 未合并 |
| Issue | #1293 | 自定义 studio http 的 mcp 无法使用 | 约 4 个月（2026-04-02 创建） | 已 stale 关闭，但属于 MCP 核心功能缺陷，建议重新评估 |
| Issue | #1296 | 上传长图（3M）解析报错 | 约 4 个月（2026-04-02 创建） | 涉及「整体不可用」的严重问题，关闭前无 fix PR，需确认回归风险 |
| Issue | #1307 | 编辑面板关闭后无法编辑其他 provider 配置 | 约 4 个月（2026-04-02 创建） | 功能状态残留 bug，关闭前无 fix PR，需确认回归风险 |
| PR | #2358 | fix(cowork): session rename 失败反馈 | 约 2 周（2026-07-18 创建） | 实现完整，Fixes #670，等待合入 |

> ⚠️ **项目健康度提示**：今日关闭的 6 条 Issue 中有 5 条属于「有描述、无 fix、stale 自动关闭」状态，且集中在同一时间段创建（2026-04-01/02），存在批量关闭掩盖真实问题的风险。建议对 #1296、#1298、#1307 三条无明显修复关联的 Issue 做一次回归验证，确认是否仍在最新版本中复现，避免用户问题被静默丢弃。

---

*数据来源：[github.com/netease-youdao/LobsterAI](https://github.com/netease-youdao/LobsterAI) | 统计时间窗口：2026-08-01 ~ 2026-08-02*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 — 2026-08-02

## 1. 今日速览

过去 24 小时 Moltis 项目保持稳定的迭代节奏：共 4 条 PR 动态，其中 3 条已合并/关闭、1 条待合并；无新 Issue 提交，亦无新版本发布。合并的 3 个 PR 覆盖了可观测性基础设施、权限安全加固与 Nostr 群聊协议支持三个重要方向，项目在可观测性和企业级安全边界上迈出实质性一步。当前仅 1 个 PR 处于待合并状态，整体队列健康，无阻塞性问题。

## 2. 版本发布

今日无新版本发布。

## 3. 项目进展

今日 3 个 PR 合并/关闭，均为筹备数日的重要功能：

- **[#1174](https://github.com/moltis-org/moltis/pull/1174) 新增插桩与反馈收集基础设施（closed）** — 构建了后端无关的 agent 插桩体系，支持 Langfuse v4 导出、OTLP 运维后端及终端用户反应式反馈。该 PR 同时补齐了流式/非流式输出一致性、provider 故障转移归属、缓存感知 token 用量、推理过程记录等可观测性细节，为项目进入生产化阶段打下基础。
- **[#1170](https://github.com/moltis-org/moltis/pull/1170) 修复 channel 特权工具越权风险（closed）** — 此前通过访问 allowlist 的 channel 发送者仍能触达特权命令与主机工具。本次变更加入 per-account `operators` 列表，将「访问权」与「操作权」彻底分离，并在 commands、callbacks、queue replay、chat execution、external 调用等链路统一施行边界检查。这是一次安全模型的重要收口。
- **[#1168](https://github.com/moltis-org/moltis/pull/1168) 为 Buzz 渠道增加 NIP-29 群聊支持（closed）** — 补齐 `moltis-nostr` 对 NIP-29 群聊协议（基于 NIP-42 认证通道）的适配，使 Moltis 能够原生接入 [Buzz](https://github.com/block/buzz)（Block 开源的 AI/人类协作工作区）。这意味着 Moltis 的 AI agent 可以进入 Buzz 团队频道，与人类成员平等协同。

三条线合在一起看，项目正在从「单机 agent 运行时」向「多协议、可观测、强权限约束的生产级 AI 助手基础设施」演进。

## 4. 社区热点

今日唯一处于开放状态的 PR：

- **[#1182](https://github.com/moltis-org/moltis/pull/1182) [OPEN] fix(sessions): allow deleting and archiving the main session** — 由 shixi-li 于 8 月 1 日提交，直接修复 issue #1132 中「主会话（main session）无法删除/归档」的问题。变更移除了 `delete_impl` 与 `is_archivable_entry` 中对 `main` 会话的 guard，同时保留当前活跃渠道会话的归档限制，以及 `sessions.clear_all` 对 main/channel-bound 会话的保护。

背后诉求很明确：用户希望主会话与其他会话一样受统一规则管理，而不是被特殊对待。该 PR 虽为 bugfix，但涉及会话管理的核心行为边界，取舍上需要兼顾误删保护，是今日社区最值得关注的进展。

## 5. Bug 与稳定性

按严重程度排列：

| 等级 | 问题 | 状态 |
|---|---|---|
| 高 | **特权命令/主机工具越权风险**。channel 发送者通过 allowlist 即可触达特权命令，属安全边界缺失 | [#1170](https://github.com/moltis-org/moltis/pull/1170) 已于今日修复并关闭 |
| 中 | **`main` 会话无法删除/归档**。用户无法像管理其他会话一样管理主会话，影响会话生命周期管理 | issue #1132；修复 PR [#1182](https://github.com/moltis-org/moltis/pull/1182) 已在 8 月 1 日提交，今日仍待合并 |

无崩溃级或数据一致性方面的严重问题上报。项目在安全修复上响应迅速，从发现到合入约一周（7/26 → 8/1）。

## 6. 功能请求与路线图信号

从今日合入的 PR 中可以提取出几个明确的路线图信号：

- **生产级可观测性**：#1174 引入的插桩与反馈收集体系意味着项目准备支持真实生产负载，后续版本大概率会围绕「跟踪、审计、成本监控」提供更多上层功能。
- **权限模型精细化**：#1170 确立的 operators 列表机制，预示着未来会有更多围绕「多用户、多角色、分级管理」的功能展开。
- **企业协作工具集成**：#1168 的 Buzz/NIP-29 支持表明项目有意进入 AI agent 与人类混合协作场景，后续可能适配更多 Nostr 生态之外的协作平台。
- **会话管理操作完备性**：#1182 若顺利合并，意味着会话生命周期管理将趋于完整，为后续高级会话功能（如批量管理、会话归档策略）铺路。

## 7. 用户反馈摘要

今日无新 Issue 评论数据。从已合入 PR 的变更动机可以间接读到两类用户诉求：

- **管理员/重度用户**：需要更细粒度的权限控制，避免「能访问 channel 的人就能用特权工具」的一刀切安全模型（来自 #1170 的修复动机）。
- **日常使用者**：主会话无法删除/归档属于「明明该有却做不了」的操作痛点，说明用户希望会话管理符合直觉、规则统一（来自 #1132 对应 PR #1182 的修复动机）。

## 8. 待处理积压

当前主要待办项：

- **[#1182](https://github.com/moltis-org/moltis/pull/1182) fix(sessions): allow deleting and archiving the main session** — 已提交 24 小时以上仍处 open 状态。该 PR 修复的是 session 管理中的核心体验问题，建议维护者尽快 review 并合入。提交者 shixi-li 在描述中已明确保护性边界（active channel session 与 clear_all 的行为不受影响），技术方案具备较好的向后兼容性。

---

*报告生成时间：2026-08-02 · 数据来源：[Moltis GitHub](https://github.com/moltis-org/moltis)*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报（2026-08-02）

> 数据窗口：2026-08-01 ~ 2026-08-02 | 数据源：github.com/agentscope-ai/CoPaw

---

## 1. 今日速览

过去 24 小时项目活跃度**显著上升**：9 条 Issue 全部处于活跃状态，12 条 PR 中有 11 条待审查/合并，且**多位首次贡献者**集中提交代码（#6623、#6622、#6620、#6618），社区参与度呈积极态势。虽然无新版本发布，但修复覆盖面广：覆盖**记忆自动压缩、ACP 协议竞态、模型 Provider 兼容、时区显示**等关键模块，体现出项目正在经历一波密集的稳定性加固。值得注意的趋势是，用户反馈开始从"能不能用"转向"好不好用"——空间清理、全局快捷键、多智能体协作引导等体验类诉求增多，信号健康。

---

## 2. 版本发布

今日无新版本发布，无相关内容。

---

## 3. 项目进展

今日唯一合并/关闭的 PR 为 **#6598 fix(skills): preserve plugin-sourced skill tags across reconcile cycles**，该修复解决了插件来源技能在重启后标签丢失的问题（对应 Issue #6537）。本质上修复了 `reconcile_pool_manifest()` 与 `reconcile_workspace_manifest()` 在磁盘目录缺失时无条件删除清单条目的缺陷。这是**技能管理系统**的一次重要增强，避免用户重启后需重新配置插件技能。

此外，今日提交的多个修复 PR 均已就绪待合并，一旦合入将直接推进：

- **记忆系统**：#6629 修复自动压缩不触发记忆摘要、#6628 修复压缩占位符角色错误导致 DeepSeek API 400
- **模型兼容**：#6620 修复 Gemini 流式 tool_call 崩溃、#6631 对齐阿里云 coding plan 模型列表
- **ACP 协议**：#6623 修复通知与响应竞态导致文本丢失
- **console 体验**：#6618 修复会话列表强制 UTC 时区显示

大量修复以"一日内 Issue + PR 同出现"的方式完成闭环，项目迭代节奏健康。多个 first-time-contributor 的 PR（#6623、#6622、#6620、#6618）表示出该项目对新贡献者的友好度与 Issue 模板的有效性。

---

## 4. 社区热点

**🔥 #6593 [Feature] 增加统一且专业的清理页面**
链接：https://github.com/agentscope-ai/QwenPaw/issues/6593

今日最受关注的功能请求（2 条评论）。用户 MCQSJ 提出随着 agent 长期使用，自动记忆、工具调用、agent 协作、备份、历史会话等数据持续堆积，而当前缺少统一清理入口，连删除会话都无法联动删除工作区目录。该诉求背后反映的是 **agent 类应用的"数据生命周期管理"问题**，长期看会成为所有本地部署 agent 的刚需。

**🔥 #6621 [Question] 多智能体协作引导缺失**
链接：https://github.com/agentscope-ai/QwenPaw/issues/6621

用户 monicfenga 反馈进行了 50+ 轮多智能体对话才发现 Default Agent 不会自动调用其他 Agent，必须显式在 PROFILE.md 中写入调用指令。用户明确表示"根因是引导缺失，而非未读文档"——这是对**官方文档与产品内引导**的直接批评，值得产品团队高度重视。

**🔥 #6480 [Question] nohup 命令卡住**
链接：https://github.com/agentscope-ai/QwenPaw/issues/6480

创建于 7 月 26 日仍保持活跃，用户报告任何使用 `nohup` 或尾部 `&` 的 shell 命令都导致 agent 永远不能回到 idle。这是一个 shell 进程脱离问题，虽非新 Issue，但在持续获得关注。

**🔥 #6568 [Feature] 全局快捷键唤出浮动快速输入框**
链接：https://github.com/agentscope-ai/QwenPaw/issues/6568

2 条评论。用户希望获得类豆包/Raycast 的轻量唤起体验，当前桌面端必须点托盘图标加载整个 1280×800 主窗口，对于"随手一问"的场景过于笨重。该 PR 还详细核查了 `console/src-tauri/Cargo...` 代码，说明用户深度参与。

---

## 5. Bug 与稳定性

按严重程度排列：

**🔴 严重（崩溃/数据丢失类）**

| Issue | 描述 | 状态 |
|-------|------|------|
| [#6619](https://github.com/agentscope-ai/QwenPaw/issues/6619) | `"ToolCallBlock" object has no field "extra_content"` — Gemini 流式 tool_call 在 `openai_chat_model_compat._parse_stream_response` 崩溃，qwenpaw 2.0.1 + agentscope 2.0.4.post1 下任何带 extra_content 的流式请求都失败 | 已有 fix PR [#6620](https://github.com/agentscope-ai/QwenPaw/pull/6620) |
| [#6625](https://github.com/agentscope-ai/QwenPaw/issues/6625) | ACP `delegate_external_agent` 在 `session/update` 通知与 `session/prompt` 响应同时到达时，偶尔返回 "completed without text output"，尽管 agent 实际产生了文本 | 已有 fix PR [#6623](https://github.com/agentscope-ai/QwenPaw/pull/6623) |

**🟠 中等（功能异常/流程阻断）**

| Issue | 描述 | 状态 |
|-------|------|------|
| [#6624](https://github.com/agentscope-ai/QwenPaw/issues/6624) | Scroll 自动压缩不触发 `summarize_when_compact` 记忆流程，而手动 `/compact` 可触发——上下文自动压缩时记忆丢失 | 已有 fix PR [#6629](https://github.com/agentscope-ai/QwenPaw/pull/6629) |
| [#6626](https://github.com/agentscope-ai/QwenPaw/issues/6626) | CI 的 "Real behavior proof" gate 会完全剥离 fenced Evidence 块（如 ```text / ```bash 终端记录），导致纯代码块证据的 PR 被误拒 | 无 fix PR，建议维护者调整 CI 判定逻辑 |
| [#6480](https://github.com/agentscope-ai/QwenPaw/issues/6480) | 执行 `nohup` 或带 `&` 的 shell 命令后，`execute_shell_command` 永远不返回 idle 状态，agent 无法继续后续操作 | 无 fix PR，**已持续 7 天未解决**，建议提高优先级 |

---

## 6. 功能请求与路线图信号

结合已有 PR 判断，以下功能最有可能进入下一版本：

| 功能请求 | 对应 PR/状态 | 可能性分析 |
|----------|--------------|------------|
| **统一清理页面**（#6593） | 无对应 PR | 需求描述详尽、使用场景清晰，但属于新功能开发，短期大概率不会实现；可作为中远期路线图 |
| **全局快捷键浮动输入框**（#6568） | 无对应 PR | 用户已给出明确场景与方案参考（豆包/Raycast），成本适中，可能被采纳进桌面端 |
| **内置 OrcaRouter Provider** | PR [#6622](https://github.com/agentscope-ai/QwenPaw/pull/6622) 已提交 | 社区贡献者已完成实现，进入审查流程，大概率合入 |
| **空模型响应上报** | PR [#6630](https://github.com/agentscope-ai/QwenPaw/pull/6630) 已提交 | 修复 #6601，提升长会话稳定性，合入概率高 |
| **workspace 快捷方式到侧边栏**（#6083） | PR [#6306](https://github.com/agentscope-ai/QwenPaw/pull/6306) 待合并 | 桌面端体验增强，已存在实现，但等待时间较长（7 月 21 日提交） |
| **Provider 发现/模型元数据/路由统一**（#6167） | PR [#6302](https://github.com/agentscope-ai/QwenPaw/pull/6302) 待合并 | 属于架构级重构，体量大且涉及面广，需要更多 review 时间 |
| **工具图片内联展示** | PR [#5490](https://github.com/agentscope-ai/QwenPaw/pull/5490) 待合并 | 创建于 6 月 24 日的长寿命 PR，功能完善但等待时间过长，建议维护者给出明确结论 |

---

## 7. 用户反馈摘要

**😤 多智能体协作：引导严重缺失**
用户 monicfenga 反馈称其"进行了 50+ 轮多智能体相关对话，直到 8 月 1 日才发现 Default Agent 不会自动调用其他已创建的 Agent"。即便阅读了官方《Multi-Agent》文档，仍未被提示需要显式在 PROFILE.md 中编写调用指令。这种"文档写了但用户找不到"的体验问题，实际消耗了用户大量无效调试时间，属于**产品设计上的隐蔽成本**。

**😟 空间占用焦虑**
#6593 作者提出了一个真实的使用痛点："agent 在运行使用过程中，特别是自动记忆，以及各种工具调用，会产生大量数据，日积月累会越来越臃肿"。并明确指出当前"连删除会话都无法选择删除对应工作区目录"。这不仅是功能缺失，更是用户对**长期运行后系统健康度的担忧**——对于本地部署的 agent 工具，存储空间管理会逐渐成为核心体验的一部分。

**😐 功能隐藏太深**
#6621、#6568 等 Issue 共同指向一个问题：**QwenPaw 功能能力强，但可发现性差**。包括全局快捷键、多智能体协作触发条件、空间管理入口在内，用户期望的功能都存在，但要么无入口、要么需要挖掘文档、要么需要阅读源码才能确认（#6568 用户甚至 pull 了代码检查实现）。

**🙂 积极信号**
多个 Issue 作者表现出了非常高的技术参与度：如 #6619 用户通过 AI 编码 agent 辅助完成了 bug 复现、根因定位和 workaround 验证；#6568 用户自行核查了 `console/src-tauri/Cargo...` 源码。这说明项目吸引了**高技术水平、愿意深度参与**的核心用户群体。

---

## 8. 待处理积压

需维护者特别关注以下长期未响应的 Issue/PR：

| 项目 | 创建时间 | 等待天数 | 说明 |
|------|----------|----------|------|
| **[#6480](https://github.com/agentscope-ai/QwenPaw/issues/6480) nohup/& 命令卡住** | 2026-07-26 | 7 天 | 影响 `execute_shell_command` 工具的基本可用性，属于阻断性问题，且无维护者回复与 fix PR |
| **[#5490](https://github.com/agentscope-ai/QwenPaw/pull/5490) 工具图片内联展示** | 2026-06-24 | 39 天 | 功能完整的 UI 改进 PR，长时间未合并也未关闭，contributor 可能因反馈缺失而流失 |
| **[#6306](https://github.com/agentscope-ai/QwenPaw/pull/6306) workspace 侧边栏快捷方式** | 2026-07-21 | 12 天 | 呼声明确的桌面端体验增强，等待 review |
| **[#6302](https://github.com/agentscope-ai/QwenPaw/pull/6302) Provider 统一架构重构** | 2026-07-21 | 12 天 | 架构级大 PR，建议维护者明确阶段性计划，避免长期悬置 |
| **[#6626](https://github.com/agentscope-ai/QwenPaw/issues/6626) CI gate 剥离 fenced Evidence** | 2026-08-01 | 1 天 | CI 配置缺陷，影响新贡献者 PR 审核效率，建议快速修复 |

---

*以上日报基于 GitHub 公开数据自动生成，仅供项目健康度参考。*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw 项目动态日报 — 2026-08-02

> 数据来源：github.com/zeroclaw-labs/zeroclaw（过去 24 小时）

---

## 1. 今日速览

ZeroClaw 过去 24 小时活跃度处于**高位**：14 条 Issue 更新（全部活跃，无关闭）与 50 条 PR 更新（全部待合并，无合并/关闭）构成密集的协作流量，但 **"零合并"状态暴露出维护管道瓶颈**——社区侧提交热情高涨，维护侧审阅与合并节奏未能跟上。安全与数据隔离是今日绝对主线，4 条 S0/S1 级 Issue 集中于跨 agent 数据越权与审计日志伪造，且多数已进入 in-progress 或已有关联 PR。建议重点关注 #9646 与 #9647 两项数据泄露级缺陷的修复排期。

---

## 2. 版本发布

今日无新版本发布。

---

## 3. 项目进展

**今日无任何 PR 被合并或关闭。** 50 条待合并 PR 中以下关键工作已进入评审尾期或等待作者响应，一旦落地将对项目功能性/安全性产生实质性推进：

- **[PR #9634] fix(channels/telegram): skip unauthorized handler for non-mentioned group messages with mention_only** — 为 Telegram 频道新增 `allow_groups` 配置（含通配符 `*`），修复 mention_only 模式下非提及群消息被错误拒绝的问题。与今日 #9633 Windows null-device 修复共同指向 shell/频道策略完备性。
- **[PR #9362] fix(browser): validate screenshot destination path against workspace policy** — 修复浏览器工具 `screenshot` 动作的任意文件写入逃逸，将写入路径纳入工作区策略校验收口。该 PR 已获 priority:p1 标注。
- **[PR #8713] fix(tools): add allowed_private_hosts opt-in to file_download SSRF gate** — 封堵内部安全审计发现的第三处 SSRF 面，为 `file_download` 工具增加私有主机白名单 opt-in。
- **[PR #9091] feat(computer-use): add native macOS, Linux X11, and Windows drivers** — 跨平台桌面操控能力的规模化功能 PR（size:XL），已获得领域安全评审标记，落地后将显著扩展 agent 桌面自动化边界。
- **[PR #9080] feat(relay): secure transport and browser enrollment frontdoor** — 引入 mTLS 中继传输层与 zerocode 浏览器注册入口，属于远控安全架构的基建级合入。

> 项目整体处于**"功能待命"状态**：大量高价值补丁积压在合并队列，master 分支在过去 24 小时未吸收任何新代码。

---

## 4. 社区热点

**最热 Issue 讨论：**

- **[#8692] [Tracker]: Maintainer decision queue for RFCs and design issues**（7 评论）— 维护者决策队列本身成为本周最活跃讨论线程。该 tracker 汇集了所有待维护者裁决的 RFC、设计议题与发布策略问题，讨论热度反映社区对**决策透明度**的强需求。
- **[#9397] RFC: Treat an empty WhatsApp Web `allowed_groups` as permit-none**（5 评论，priority:p1）— WhatsApp Web 通道默认空 `allowed_groups` 当前会放行所有群组，RFC 提议改为 "permit-none" 语义。5 条评论集中于**向后兼容性**与现有部署的迁移影响，属于安全默认值方向的讨论。

**讨论态势分析：** 社区热点集中在安全默认策略（WhatsApp 空列表语义、OpenRouter 缓存成本）与治理流程透明度（决策队列）两个维度。值得注意 #9642 提出"超时审批被记录为显式拒绝"的**审计日志真实性问题**，已获得 1 条评论且有 follow-up 标记，其讨论深度预计会继续上升。

---

## 5. Bug 与稳定性

以下按严重程度递减排列，全部为过去 24 小时新增或活跃更新：

| 严重度 | Issue | 描述 | 状态 | 关联 PR |
|---|---|---|---|---|
| **S0** | [#9646] | Session/channel 读写工具缺乏 per-agent 所有权隔离，任意 agent 可通过 model-supplied 参数访问其他 agent 的会话 | 新开，无 PR | 暂无 |
| **S0** | [#9647] | 知识图谱为全局共享，无 per-agent 归属，任意 agent 可读写其他 agent 的知识 | 新开，无 PR | 暂无 |
| **P1** | [#9633] | Windows 空设备重定向 `2>nul` 被 shell 安全策略拒绝（策略仅识别 `/dev/null`） | in-progress + accepted | [#9634] 相关（同区域修复） |
| **P1** | [#9642] | Telegram 审批超时被记录为显式操作员拒绝，伪造审计痕迹 | in-progress + follow-up | 暂无 |
| **P1** | [#9631] | OpenRouter 无稳定 session_id，系统提示词与工具 schema 每次请求重放，成本偏高 | needs-maintainer-review | 暂无 |
| **P2** | [#9640] | WhatsApp Web 策略文档注释引用不存在的 V2 配置键 `allowed_numbers` | in-progress + accepted | 暂无 |
| **P2** | [#9628] | 博客缺少 RSS/Atom 订阅源 | accepted | 暂无 |

**特别警示：** #9646 与 #9647 为跨 agent 数据泄露级（S0）缺陷，当前**均无修复 PR**。ZeroClaw 多 agent 架构中若此类越权面扩散，将直接影响平台作为多租户/多角色基础设施的可信度。建议维护团队将这两个 Issue 排入最高优先级队列。

---

## 6. 功能请求与路线图信号

今日新增强类型 Issue/PR 揭示以下路线图信号：

- **[#9631] OpenRouter 稳定 session_id 支持**（enhancement）— 通过传递稳定 `session_id` 使 OpenRouter prompt caching 生效，可降低 chat 场景 30-60% token 成本。该提案兼具性能与成本双重收益，与 OpenRouter 生态日渐紧密的社区诉求吻合，**有望在 v0.9.0 配置重构期间纳入**。
- **[#9632] ACP 默认 agent 选择（--agent 参数）**（enhancement）— 为独立 ACP 场景提供进程级默认 agent。属于 ACP 通道易用性补全，方向与 #8655 及 #9080 的 CLI/通道整合大方向一致。
- **[#9644] RFC: 退役 Lucid memory connector**（risk:high）— 上游项目"集成 4 天后即休眠"，RFC 提议在 v0.9.0 移除。这类"及时止损"的依赖淘汰提议释放出**维护层开始重视依赖健康度**的信号，可能带动后续对休眠 connector/backend 的系统性清理。
- **PR #9385 实现 WhatsApp Web `request_approval`** — 将人工审批能力移植至 WhatsApp Web 传输层，与 v0.9.0 审批/安全强化队列（#7432）直接呼应。

整体路线图信号：**v0.9.0（auth/security/gateway 强化）仍是核心锚点**，今日新增工作大多围绕该里程碑的安全边界议题展开。

---

## 7. 用户反馈摘要

从今日活跃 Issue 评论与 PR 讨论中提炼用户之声：

- **"OpenRouter 成本真的太高了"** — #9631 反馈称单对话产生数十次 LLM 请求，系统提示与工具 schema 每次重复发送，"不必要地昂贵"。这是对**实际使用成本**的直接痛点陈述，非推测性需求。
- **"文档告诉我的配置键根本不存在"** — #9640 报告 V3 配置中不存在 `allowed_numbers`，但文档注释仍引导运营商配置该键。用户对**文档与代码脱节**表示明确不满，属于可快速修复的信任损耗点。
- **"审计日志记录的是系统超时，不是我的决定"** — #9642 指出超时审批被永久记录为操作员显式拒绝，"改变了日志所声称的人类行为"，这一反馈体现了高安全敏感性用户对审计完整性的高度关注。
- **"知识图谱分不开？那谁敢多 agent 共用"** — #9646/#9647 的 S0 报告虽然来自安全测试语境，但留言中隐含对多 agent 数据边界缺乏信心的顾虑。
- **呼吁 RSS feed**（#9628）— 一位用户评论"博客不错，但缺个 RSS 订阅"，反映用户期望以更轻量方式跟踪项目动态，属于低投入高好感度的社区运营改进。

---

## 8. 待处理积压

**长期滞留 PR（stale-candidate 标记）：**

- **[PR #8546] fix(cli): localize status fragments**（6/30 创建，35+ 天未合并）
- **[PR #8576] fix(channels): add env-var fallback for OpenAI STT credentials**（7/1 创建，含被忽略的环境变量配置）
- **[PR #8655] refactor(zerocode): consolidate Code pane, rails, and prompt drafts**（7/3 创建，size:XL，影响面大）
- **[PR #7821] feat(config): add schema struct & risk field**（6/17 创建，v0.9.0 配置改革的一部分）

**积压治理建议：** 上述 stale 候选 PR 均带 `needs-author-action` 标记，需维护者在未来 1-2 个周期内做出"重新请求更新/关停/介入接手"的三选一决策，避免技术债持续膨胀。

**维护者决策积压 — [#8692]** 该决策队列目前收纳了 7 条未决评论所对应的 RFC/设计议题，建议在 #7432（v0.9.0 里程碑）之外为 RFC 评审设立固定时段，以缓解"提交多、决策少"的社区感知失衡。

---

*本日报由 AI 分析师基于 GitHub 公开数据自动生成，仅供参考，不构成对项目健康状况的最终判断。*
*数据窗口：2026-08-01 至 2026-08-02（UTC）*

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*