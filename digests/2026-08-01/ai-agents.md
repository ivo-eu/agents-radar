# OpenClaw 生态日报 2026-08-01

> Issues: 214 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-08-01 00:12 UTC

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

# OpenClaw 项目动态日报 — 2026-08-01

---

## 1. 今日速览

过去 24 小时项目整体活跃度处于**高位**：共产生 214 条 Issue 更新（其中新开/活跃 193 条、已关闭 21 条）和 500 条 PR 更新（待合并 372 条、已合并/关闭 128 条），今日无新版本发布。Issue 侧的讨论热度集中在 Telegram 重复回复、SQLite 快照可靠性、会话上下文膨胀等长期未决问题上；PR 侧则以网关启动性能优化、Ollama 兼容性修复和跨渠道消息投递修复为主。值得关注的是，今日新开的 PR（如 #117118、#117119）标志着维护者正在积极处理大型会话历史的网关阻塞问题，但 372 条待合并 PR 的积压规模也反映出维护者审阅负载仍然较重。

---

## 2. 版本发布

今日无新版本发布（最新 Releases 为空）。上一版本线停留在 2026.7.2-beta.x 系列，当前 Issue 中仍有针对该版本线的未解决问题（如 #116973 文档与配置废弃路径不一致等）。

---

## 3. 项目进展

虽无新版本发布，但今日有 128 条 PR 被合并/关闭，项目在以下方向取得了实质性推进：

**会话状态与消息投递修复**
- [PR #117090](https://github.com/openclaw/openclaw/pull/117090)（已合并）：修复 Control UI 会话侧边栏切换状态过滤器后，过期的 in-flight 子会话响应重新填充新视图的问题。
- [PR #117119](https://github.com/openclaw/openclaw/pull/117119)（新开）：修复 Discord 场景下并行子代理派发时父会话最终合成回复丢失、已完成会话仍显示为活跃的问题。
- [PR #113192](https://github.com/openclaw/openclaw/pull/113192)（已关闭）：使 CLI 会话重置粒度化，避免单个空回复触发 `clearAllCliSessions` 误清理无关会话绑定。

**Ollama 提供商兼容性**
- [PR #116737](https://github.com/openclaw/openclaw/pull/116737)（已合并）：修复原生 Ollama 请求忽略 stop sequences、异步 payload hooks 尚未完成即发起请求的问题。

**会话状态展示优化**
- [PR #116742](https://github.com/openclaw/openclaw/pull/116742)（已合并）：压缩 `session_status(changesSince)` 的模型可见文本，防止内部会话标识符和嵌套 catalog/host 字段绕过脱敏层泄露。

**网关性能方向**
- [PR #117118](https://github.com/openclaw/openclaw/pull/117118)（新开）：在 Gateway prewarm 之前计数大型历史记录，防止事件循环在大历史恢复时被阻塞。

**其他已合并/关闭**
- [PR #116925](https://github.com/openclaw/openclaw/pull/116925)、[PR #115733](https://github.com/openclaw/openclaw/pull/115733) 等仍在等待合入的 PR 也在推进中。

整体来看，项目正在围绕「会话状态一致性」「消息投递可靠性」「网关启动与恢复性能」三条主线密集修复，方向明确。

---

## 4. 社区热点

今日讨论热度最高的 Issue 集中在以下几个长期未决的稳定性问题上：

**[Issue #86519 — Telegram 重复回复回归（14 评论）](https://github.com/openclaw/openclaw/issues/86519)**
自 2026.5.20 更新后，Telegram 上同一用户消息被 agent 重复回复 2-10 次；升级到 5.22 后严重程度有所缓解但仍未彻底修复。该问题从 5 月持续至今仍未关闭，社区关注度高，诉求集中在：希望维护者明确根因、给出可用的 workaround 或合入修复 PR。该 Issue 已被标记为 P1 且带有多个 clawsweeper 流程标签，但尚未有 fix PR 关联。

**[Issue #113306 — SQLite 快照恢复缺乏端到端崩溃与身份保证（13 评论）](https://github.com/openclaw/openclaw/issues/113306)**
SQLite 快照创建/恢复可能报告成功，但未持久链接新建的父目录、绕过最终 target/sidecar 身份守卫，且清理路径仍使用 patched 逻辑。该问题涉及数据安全，被标记为 maintainer 级 P1，目前处于 `needs-maintainer-review` 状态。

**[Issue #67419 — 会话上下文膨胀：bootstrap 文件每轮重复注入浪费 20-30% tokens（11 评论）](https://github.com/openclaw/openclaw/issues/67419)**
这是一个自 4 月持续至今的 P2 问题，用户明确表达了 token 成本浪费的痛点，评论中多次提到希望 bootstrap 注入策略改为「首轮注入 + 后续按需」模式。

**[Issue #115908 — 会话 transcript 投影在持续写入下 livelock（11 评论）](https://github.com/openclaw/openclaw/issues/115908)**
同步重建循环占用 Node 主线程导致所有渠道传输停滞，用户反馈「每次 rebuild 路径都是同步的，事件循环阻塞数十秒」，属于 P1 级别的核心会话稳定性问题。

**[Issue #114137 — 可见渠道轮次间歇性丢失待发送回复（11 评论）](https://github.com/openclaw/openclaw/issues/114137)**
最终文本已持久化到 transcript 但从未投递到渠道——这类「静默消息丢失」类问题在社区中呼声很高，反映了用户对消息投递可靠性的强烈关注。

---

## 5. Bug 与稳定性

今日报告的 Bug/回归问题按严重程度排列如下：

| 严重级 | Issue | 摘要 | 是否有 fix PR |
|--------|-------|------|--------------|
| **P0** | [#112395](https://github.com/openclaw/openclaw/issues/112395) | 从 6.11 升级到 7.1 后启动迁移预检阻塞 gateway，迁移表和租约均为空 | ✅ 有（linked-pr-open） |
| **P1** | [#86519](https://github.com/openclaw/openclaw/issues/86519) | Telegram 重复回复 2-10x（5.20 回归，5.22 部分缓解） | ❌ 无，needs-maintainer-review |
| **P1** | [#113306](https://github.com/openclaw/openclaw/issues/113306) | SQLite 快照恢复缺少端到端崩溃与身份保证 | ❌ 无 |
| **P1** | [#115908](https://github.com/openclaw/openclaw/issues/115908) | 会话 transcript 投影在持续写入下 livelock，阻塞主线程 | ❌ 无 |
| **P1** | [#114137](https://github.com/openclaw/openclaw/issues/114137) | 可见渠道（Signal）消息已持久化但从未投递 | ❌ 无 |
| **P1** | [#87109](https://github.com/openclaw/openclaw/issues/87109) | Gateway 空闲时 heap 增长至 1073MB+，cron 任务静默失败 | ❌ 无 |
| **P1** | [#114211](https://github.com/openclaw/openclaw/issues/114211) | Matrix 房间 agent 可循环于 no-reply 输出、重启恢复与陈旧会话重放 | ❌ 无 |
| **P1** | [#89228](https://github.com/openclaw/openclaw/issues/89228) | 隔离 cron 会话中 exec 间歇性不可用（曾于 4.1 修复后回归） | ❌ 无 |
| **P1** | [#115476](https://github.com/openclaw/openclaw/issues/115476) | 压缩后上下文刷新重放旧 message_id，缺少 Telegram 网关级去重 | ❌ 无 |
| **P1** | [#114654](https://github.com/openclaw/openclaw/issues/114654) | `agents.defaults.compaction.*` 热重载被当作 no-op，配置修改永不生效 | ❌ 无 |
| **P2** | [#115152](https://github.com/openclaw/openclaw/issues/115152) | #95939 修复回归：`bootstrapMaxChars`/`bootstrapTotalMaxChars` 每次重启被删除 | ❌ 无 |
| **P2** | [#114158](https://github.com/openclaw/openclaw/issues/114158) | fs-safe 硬编码 0o600 忽略 umask，破坏共享工作区（NFS/SMB/多用户） | ❌ 无 |
| **P2** | [#114192](https://github.com/openclaw/openclaw/issues/114192) | TUI 会话历史在压缩后消失（底层会话完好） | ❌ 无 |
| **P2** | [#96007](https://github.com/openclaw/openclaw/issues/96007) | Discord 多段回复中内联错误文本导致后续内容被截断 | ✅ 有（linked-pr-open） |

**观察**：P1 级问题数量较多（10 个），且大部分仍处于 `needs-maintainer-review` 或 `needs-product-decision` 状态，仅有 2 个关联了修复 PR。P0 的升级迁移阻塞问题有对应 PR，但尚未合入。

---

## 6. 功能请求与路线图信号

今日活跃的功能请求中，以下几个方向值得关注：

**插件/提供商生态扩展**
- [Issue #87325](https://github.com/openclaw/openclaw/issues/87325)：支持 Azure AI Foundry GPT Realtime Talk via gateway relay——与 [PR #114146](https://github.com/openclaw/openclaw/issues/114146)（为 talk.realtime.providers 添加 baseUrl 配置）形成呼应，表明实时语音的多提供商适配正在推进。
- [Issue #10687](https://github.com/openclaw/openclaw/issues/10687)：完全动态的模型发现（OpenRouter 等快速变化的目录），3 👍，评论中强调当前静态模型目录已无法满足需求。
- [PR #116300](https://github.com/openclaw/openclaw/pull/116300)（新开）：将腾讯提供商源代码外部化，减少 OpenClaw 核心包体积。

**本地化与多语言支持（已有 PR 系列）**
- [PR #111544](https://github.com/openclaw/openclaw/pull/111544)、[PR #111545](https://github.com/openclaw/openclaw/pull/111545)、[PR #111542](https://github.com/openclaw/openclaw/pull/111542)：TUI `/gateway-status`、Gateway 审批错误描述、updater dry-run 预览的本地化实现，均处于 waiting-on-author 状态，有望随 RFC openclaw/rfcs#42 进入下一版本。

**审计与可观测性**
- [PR #117034](https://github.com/openclaw/openclaw/pull/117034)：新增 `openclaw audit --run <id> --explain` 命令，展示不可变执行身份上下文——这是安全审计方向的新能力。
- [Issue #57404](https://github.com/openclaw/openclaw/issues/57404)：在 WebSocket 生命周期事件流中暴露每次运行的 token 用量。

**内存与性能调优**
- [PR #115138](https://github.com/openclaw/openclaw/pull/115138)：为本地文件系统 SQLite 数据库启用 memory-mapped reads，直接回应 #112758 中事件循环阻塞问题。
- [PR #115703](https://github.com/openclaw/openclaw/pull/115703)：将诊断内存压力阈值（RSS 1.5GiB、heap 1GiB 等）暴露为可配置项。

综合来看，下一版本可能包含：本地化基础框架（3 个 XL PR 同时推进）、审计命令、SQLite mmap 性能优化、Ollama 兼容性修复，以及腾讯提供商外部化。

---

## 7. 用户反馈摘要

从今日活跃的 Issue 评论中提炼的用户声音：

**高频痛点**
- 「升级后 Telegram 重复回复 2-10 次，即使降级到 5.22 也只是从 8-10 次降到 2-3 次」——[#86519](https://github.com/openclaw/openclaw/issues/86519) 的用户对升级回归表达了明显不满，且 2 个月未修复加剧了情绪。
- 「Gateway 空闲内存从 558MB 涨到 1073MB+，cron 任务静默失败——无输出、无推送、无错误上报」——[#87109](https://github.com/openclaw/openclaw/issues/87109) 中用户对「静默失败」的不可观测性最为沮丧。
- 「每轮对话都重新注入 bootstrap 文件，浪费 20-30% token」——[#67419](https://github.com/openclaw/openclaw/issues/67419) 用户对 token 成本浪费表达了明确不满（2 👍）。
- 「配置修改后只产生一行 log 说 no-op，实际永不生效」——[#114654](https://github.com/openclaw/openclaw/issues/114654) 用户对配置热重载的言行不一表示困惑。

**使用场景反馈**
- 有用户使用 OpenClaw 助手协助撰写 Issue（如 [#89228](https://github.com/openclaw/openclaw/issues/89228) 由 Ruben 的 OpenClaw assistant 研究并生成），说明项目被实际用于生产环境且用户深度依赖。
- [#115152](https://github.com/openclaw/openclaw/issues/115152) 用户 Lp12138 指出 #95939 虽标记完成但引入了回归，且字段被「删除而非重置」，反映了用户对配置数据安全性的在意。

**积极信号**
- 多个 PR 由社区贡献者提交（如 steipete 一人今日提交了 4 个 PR），且 PR 描述中明确标注「AI-assisted: yes. I reviewed the implementation and validation results.」——说明项目已形成「AI 辅助编码 + 人工审阅」的协作范式。

---

## 8. 待处理积压

以下 Issue/PR 长期未获实质推进，提醒维护者关注：

**长期未闭合的高级别 Issue**
- [Issue #86519](https://github.com/openclaw/openclaw/issues/86519)（P1，5/25 创建）：Telegram 重复回复已持续 2 个月+，14 条评论，仍无修复 PR。
- [Issue #10687](https://github.com/openclaw/openclaw/issues/10687)（P2，2/6 创建）：动态模型发现，3 👍，9 条评论，半年未推进。
- [Issue #47979](https://github.com/openclaw/openclaw/issues/47979)（P2，3/16 创建）：Control UI Dashboard v2 在 Chrome 146 上冻结，Firefox 正常，已带 stale 标签。
- [Issue #53654](https://github.com/openclaw/openclaw/issues/53654)（P2，3/24 创建）：Discord 消息编辑/删除事件支持（edit-to-reprocess / delete-to-cancel），3 👍，社区呼声较高。
- [Issue #8724](https://github.com/openclaw/openclaw/issues/8724)（P2，2/4 创建）：按模型配置独立 generation 超时，5 条评论，未分配。

**等待作者更新的 PR**
- [PR #111544](https://github.com/openclaw/openclaw/pull/111544)、[PR #111545](https://github.com/openclaw/openclaw/pull/111545)、[PR #111542](https://github.com/openclaw/openclaw/pull/111542)：三个本地化 XL PR 均已超过 2 周处于 waiting-on-author 状态，依赖 #112801 的合入，可能形成阻塞链。

**已有关联 PR 但久未合入的高价值修复**
- [Issue #90098](https://github.com/openclaw/openclaw/issues/90098)（P1）：大附件处理导致栈溢出，已有 linked PR，但自 6/4 创建至今仍未合入。
- [PR #115138](https://github.com/openclaw/openclaw/pull/115138)：SQLite mmap 修复已 ready-for-maintainer 状态但尚未合并，该 PR 对大规模部署的性能改善显著。

**值得注意的流程信号**
- 372 条待合并 PR 中，部分已进入「ready for maintainer look」状态但可能受限于维护者带宽。例如 [PR #117034](https://github.com/openclaw/openclaw/pull/117034)（audit 命令 XL PR）和 [PR #115761](https://github.com/openclaw/openclaw/pull/115761)（cron 命令跨热重载保留）均已完成 proof 阶段，等待维护者最终审阅。

---

*报告生成时间：2026-08-01 · 数据来源：[github.com/openclaw/openclaw](https://github.com/openclaw/openclaw)*

---

## 横向生态对比

## 1. 生态全景

过去 24 小时，11/13 个受监测项目有活跃动态，个人 AI 助手/自主智能体开源生态整体处于**高烈度迭代期**：OpenClaw 单日产生 500 条 PR 更新，IronClaw、ZeroClaw、Hermes Agent、CoPaw 等项目的 PR 更新也普遍达到 40–50 条量级。生态竞争焦点集中在**会话可靠性、上下文成本、Provider 兼容性、安全权限边界与运行时部署灵活性**，几乎没有项目停留在纯功能“堆量”阶段。与此同时，多数项目存在 PR 合并积压，说明社区贡献产能已开始超出维护者审阅产能。整体呈现“贡献旺盛、消化偏慢、安全与稳定性优先”的态势。

---

## 2. 各项目活跃度对比

| 项目 | Issues 更新 | PR 更新 | Release | 健康度 / 阶段 |
|---|---|---|---|---|
| **OpenClaw** | 214（新开/活跃 193，关闭 21） | 500（待合并 372，合并/关闭 128） | 无 | 极高活跃；维护者审阅积压严重 |
| **NanoBot** | 4（新开 2，关闭 2） | 17（合并/关闭 7，待合并 10） | 无 | 健康；SQLite 会话存储迁移落地 |
| **Hermes Agent** | 16（活跃 15，关闭 1） | 50（待合并 46，合并/关闭 4） | 无 | 高产出；review 积压明显 |
| **PicoClaw** | 2（新开 2） | 3（待合并 3） | 无 | 中等；PR 合并周期偏长 |
| **NanoClaw** | 8（全部活跃） | 9（合并 3，待合并 6） | 无 | 活跃；安全加固与多运行时并行 |
| **NullClaw** | 0 | 1（待合并 1） | 无 | 低频；等待 grok-cli provider 合并 |
| **IronClaw** | 16（活跃 15，关闭 1） | 50（合并/关闭 32，待合并 18） | 无 | 极高强度；WS1 架构重构三连合并 |
| **LobsterAI** | 0 新增（4 条 stale 自动关闭） | 12（合并/关闭 11，待合并 1） | 无（发布准备 PR 已合并） | 发布冲刺；社区侧偏沉寂 |
| **TinyClaw** | 0 | 0 | 无 | 无活动 |
| **Moltis** | 2 | 7（合并 1，待合并 6） | 无 | 中高活跃；功能与安全加固双线推进 |
| **CoPaw** | 18（活跃 13，关闭 5） | 43（合并/关闭 13，待合并 30） | 无 | 高活跃；Bug 响应快，合并率偏低 |
| **ZeptoClaw** | 0 | 0 | 无 | 无活动 |
| **ZeroClaw** | 27（关闭 2，其余活跃） | 50（合并/关闭 12，待合并 38） | 无 | 高活跃；治理流程成熟，技术债清理系统化 |

---

## 3. OpenClaw 在生态中的定位

OpenClaw 是当前生态中**体量最大、覆盖面最广的核心参照项目**。单日 214 条 Issue 更新、500 条 PR 更新远超同生态其他项目，社区讨论覆盖 Telegram/Discord/Matrix/Signal 等大量渠道，并深入 Ollama 兼容性、SQLite 快照、网关性能等底层稳定性问题。

**技术路线差异**：OpenClaw 更偏向“大而全”的用户驱动式迭代，近期主线集中在会话状态一致性、消息投递可靠性、网关启动与恢复性能；相比之下，IronClaw 走契约提取与模块化的架构重构路线，ZeroClaw 更强调 RFC 治理与安全沙箱，NanoBot/PicoClaw 则选择轻量极简路径。

**社区规模与瓶颈**：OpenClaw 的 Issue/PR 量级是第二梯队（IronClaw、ZeroClaw、Hermes）的 4–10 倍，但 372 条待合并 PR 也意味着维护者审阅带宽严重承压；P1 级问题如 Telegram 重复回复已持续 2 个月未修复，暴露出“高社区热度 + 高维护负债”并存的现实。其生态位仍是多数衍生项目的功能参照系，但正面临来自更专注、更治理化项目的分流压力。

---

## 4. 共同关注的技术方向

**1. 会话状态与消息投递可靠性**
- OpenClaw：Telegram 重复回复（#86519）、Signal 消息已持久化但未投递（#114137）
- NanoBot：微信重新扫码后旧 token 覆盖新 token（#5195）
- CoPaw：微信 cron 推送显示 `success` 但实际从未送达（#6614）；长命令阻塞飞书会话 1.5 小时（#6608）
- ZeroClaw：频道 reload 后 delivery registry 残留（#9591）

**2. 上下文/记忆管理与 Token 成本控制**
- OpenClaw：bootstrap 文件每轮重复注入，浪费 20–30% tokens（#67419）
- LobsterAI：修复 DeepSeek 长会话缓存命中率从 ~57% 恢复到接近 100%（#2413/#2415）
- CoPaw：压缩前未 flush Auto-Memory，导致早间事件从每日记忆永久缺失（#6555/#6592）
- Hermes Agent：子 agent 无工具集限制，系统提示词膨胀数千 tokens（#75737）

**3. Provider 兼容与互操作标准化**
- OpenClaw：Ollama stop sequences 与异步 hooks 修复（#116737）
- NanoBot：DeepSeek Responses API 支持（#5197）
- NullClaw：xAI Grok CLI provider（#981）
- ZeroClaw：OpenAI Chat Completions 兼容端点（#8550/#8486）
- PicoClaw：可配置模型默认 fallback 链（#3200）

**4. 安全、隐私与权限边界**
- IronClaw：P0 级跨用户内存泄漏（#6900）
- Moltis：恶意 zip/模型路径任意文件写入（#1180）、节点配对签名缺失（#1179）
- Hermes Agent：Honcho 插件静默将自托管数据路由到生产云（#43800）
- ZeroClaw：Landlock 沙箱阻断 shell 工具（#8973，已修复）
- NanoClaw：伪造点击可在源校验前篡改 ask_user_question 卡片（#2923）

**5. 运行时与部署灵活性**
- NanoClaw：K8s 受限环境部署困难（#1184）、原生 runner 绕过 Docker（#1732）
- CoPaw：独立 Python 运行环境诉求（#6160）
- ZeroClaw：Wasm-first 插件运行时 RFC（#8135）
- Hermes Agent：多 profile 生命周期管理不稳定（#75598）

---

## 5. 差异化定位分析

| 项目 | 功能侧重 | 典型用户 | 架构/技术特征 |
|---|---|---|---|
| **OpenClaw** | 全功能个人 AI 助手，渠道与 Provider 聚合 | 主流个人用户、社区重度用户 | 大而全，Gateway + 多渠道 + Control UI；迭代快，负债也重 |
| **NanoBot** | 轻量即时聊天助手，微信渠道深度优化 | 个人轻量用户 | 小步快跑；JSONL→SQLite 存储迁移，CI 稳定性投入明显 |
| **Hermes Agent** | 开发者向桌面/CLI Agent，多平台适配 | 开发者和专业用户 | Slack/Discord/WhatsApp 适配、LSP 支持、AI 辅助编码贡献 |
| **PicoClaw** | 极简多协议客户端 | 隐私/极客用户 | IRC/DeltaChat/SimpleX 等协议扩展；资源占用敏感 |
| **NanoClaw** | 极简轻量“Claw” | 轻量部署用户 | Docker 为主，正在向 K8s/原生/Apple Container 扩展 |
| **NullClaw** | 本地 CLI Provider 中立 Hub | 已拥有本地 CLI 工具的用户 | codex/gemini/claude/grok-cli wrapper 模式；尚处早期 |
| **IronClaw** | 平台级 / 多租户 Agent 基础设施 | 企业、平台运维 | Rust + 契约提取（WS1 系列）；MCP 托管、安全加固优先 |
| **LobsterAI** | OpenClaw 衍生，中文桌面与协作场景 | 中文桌面用户 | 深度优化 DeepSeek 缓存；边聊（BTW）隔离；发布收敛期 |
| **Moltis** | 开放协议多 Agent 协作 | 协议爱好者、协作团队 | Nostr/NIP-29；多种记忆后端；可观测性与反馈基础设施 |
| **CoPaw** | AgentScope 生态多 Agent 工作流 | 中文用户、多 Agent 场景 | 飞书/微信集成、ReMe 记忆、多会话 UI；Bug 响应快 |
| **ZeroClaw** | 工程化、安全沙箱化的 Agent 框架 | 工程化团队、安全敏感用户 | RFC 驱动、Landlock、Wasm 插件方向、OpenAI 兼容端点推进 |

---

## 6. 社区热度与成熟度

**快速迭代 / 高活跃层**：OpenClaw、IronClaw、ZeroClaw、Hermes Agent、CoPaw、NanoBot。这些项目单日 PR 更新普遍在 17–500 条之间，且都有清晰的近期主线——OpenClaw 修稳定性，IronClaw 做架构重构，ZeroClaw 做治理与技术债清理，CoPaw 快速修 Bug，NanoBot 落地存储迁移。

**中活跃 / 局部推进层**：NanoClaw、Moltis、LobsterAI、PicoClaw。活跃度中等，但方向明确：NanoClaw 在扩展运行时可选项，Moltis 在安全加固，LobsterAI 在发布前收敛，PicoClaw 在多协议扩展。

**低频 / 萌芽层**：NullClaw、TinyClaw、ZeptoClaw。NullClaw 只有 1 个待合并 PR，TinyClaw 与 ZeptoClaw 完全无活动，仍需观察是否具备持续维护动力。

**成熟度特征**：IronClaw 和 ZeroClaw 已进入“质量巩固”阶段，强调契约、RFC、CI 治理；LobsterAI 处于版本发布前的收敛期；OpenClaw 则仍处于“高速修复与功能堆叠并存”的活跃不稳定期。

---

## 7. 值得关注的趋势信号

**1. 维护者审阅带宽已成为生态瓶颈。**
OpenClaw 372 条、Hermes 46 条、ZeroClaw 38 条、CoPaw 30 条 PR 待合并，说明社区贡献产能已显著大于维护者消化能力。未来的竞争力将部分取决于项目能否建立 AI 辅助审阅、维护者轮值或自动合并机制。

**2. “静默失败”比崩溃更损害用户信任。**
CoPaw 微信推送任务长期显示成功但从未送达，OpenClaw 消息持久化但渠道未投递，ZeroClaw 并发命令静默丢更新，Hermes 配置热重载 no-op。开发者应把失败可观测性作为一等需求。

**3. 上下文与 Token 成本正在成为核心产品指标。**
从 OpenClaw 的 bootstrap 重复注入浪费，到 LobsterAI 将缓存命中率从 57% 修复到接近 100%，再到 Hermes 限制子 agent 工具集以压缩系统提示词——Token 效率已不仅是成本问题，而是体验和留存问题。

**4. 安全与隐私是潜在用户的采纳前提。**
Moltis 外部贡献者表示“想用 Moltis，但需要先合入安全修复”；Hermes 的 Honcho 自托管数据被路由到云；IronClaw 出现 P0 跨用户内存泄漏。安全缺陷已成为阻碍新用户进入的直接门槛。

**5. Docker 硬依赖正在被反噬。**
NanoClaw 用户明确要求 K8s、原生模式、Apple Container；ZeroClaw 用 Wasm 取代插件运行时；CoPaw 用户希望内置 Python 环境。运行时可移植性将从“加分项”变为“必备项”。

**6. 开放协议与互操作是下一阶段竞争焦点。**
ZeroClaw 的 OpenAI Chat Completions 端点、Hermes 的 LSP 支持、NullClaw 的本地 CLI wrapper、Moltis 的 Nostr 群聊支持，共同指向一个方向：Agent 工具链必须能接入标准客户端和既有开发者工作流，而不是自成孤岛。

---

*本报告基于 2026-08-01 各项目 GitHub 公开动态生成，数据范围以各项目日报为准。*

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目动态日报 — 2026-08-01

## 今日速览

NanoBot 在过去 24 小时呈现高活跃度，PR 更新 17 条，其中 7 条已合并/关闭，核心修复集中在微信渠道 session 恢复（#5196、#4223）、WebUI 滚动与滚动归属（#5193）、CI 稳定性（#5145）等方向。值得关注的是，**session 存储从 JSONL 迁移到 SQLite 的大工程（#5173）已于今日合并**，标志着 NanoBot 内部架构的一次重要演进。Issue 侧共 4 条更新，其中 2 条已关闭，2 条新开且尚未有对应修复方案。今日无新版本发布，项目整体处于活跃迭代、快速修 bug 的健康节奏中。

## 版本发布

今日无新版本发布。

## 项目进展

今日合并/关闭的 PR 中，有 3 项值得重点关注：

- **session 存储从 JSONL 迁移至 SQLite（[PR #5173](https://github.com/HKUDS/nanobot/pull/5173)，已合并）** 这是近期最大的架构级变更，使 `sessions.db` 成为唯一运行时 session 存储，首次启动时事务性导入旧 JSONL 文件，并保留 JSONL 作为回滚备份。该变更直接影响 WebUI 会话列表和 Dream 剪枝流程。长期来看将改善会话读写的并发性能与数据一致性，但**需要社区用户注意首次迁移时的启动时间及兼容性**。
- **CI 稳定化与加速（[PR #5145](https://github.com/HKUDS/nanobot/pull/5145)，已合并）** 将依赖安装批次化、用 stdin 门控的握手替代时序敏感的 exec 超时测试，减少 CI 的 flakiness。这反映了项目在工程化上的持续投入。
- **微信渠道 session 过期后的状态恢复（[PR #5196](https://github.com/HKUDS/nanobot/pull/5196)，已合并）** 修复了 #5195：长轮询渠道在遇到 `errcode -14` 进入 60 分钟暂停后，若用户在暂停期间通过扫码刷新了 `account.json`，暂停结束后会自动重新加载 state，而非继续使用内存中的过期 token。同问题也由更早的 [PR #4223](https://github.com/HKUDS/nanobot/pull/4223)（6 月 6 日创建，今日关闭）提出修复思路，最终被 #5196 的实现取代。

此外，Slack 线程作用域（#5192）和 WebUI 近尾部滚动归属（#5193）均已合并，属于用户可感知的体验优化。（数据来源：今日合并/关闭 7 条 PR）

## 社区热点

今日讨论最集中的是 **Issue #5195（[微信重新扫码后 token 被旧值覆盖](https://github.com/HKUDS/nanobot/issues/5195)，2 条评论）**，这也是今日唯一有评论的 Issue。该 bug 的触发路径非常具体：用户在 WebUI 中重新扫码登录微信个人号，但 `stop()` 中的逻辑将新扫码得到的 token 覆盖为旧 token，导致重启后的 channel 实例在首次 `getupdates` 时立即收到 `errcode -14`（session 过期），并被强制暂停 60 分钟。由于微信个人号渠道在 NanoBot 用户群中常用作私人助手入口，这类"越修越坏"的回归对信任打击较大。背后的核心诉求是：**扫码重登必须成为可信赖的一等恢复路径**，而非让用户陷入需要手动删状态文件的困境。

好在该问题已在 24 小时内被响应并修复（#5196），社区响应速度值得肯定。

## Bug 与稳定性

今日报告的 Bug 按严重程度排列如下：

| 严重度 | Issue | 描述 | 状态 |
|---|---|---|---|
| 🔴 高 | [#5195](https://github.com/HKUDS/nanobot/issues/5195) | 微信重新扫码登录后旧 token 覆盖新 token，导致立即 errcode -14 并暂停 60 分钟 | 已关闭，已有修复（[#5196](https://github.com/HKUDS/nanobot/pull/5196)） |
| 🔴 高 | [#5190](https://github.com/HKUDS/nanobot/issues/5190) | Windows 上静态资源（.js）MIME type 被注册表覆盖为 `text/plain`，导致前端模块脚本无法加载，WebUI 完全不可用 | 未关闭，已有修复 PR（[#5191](https://github.com/HKUDS/nanobot/pull/5191)） |
| 🟡 中 | [#5198](https://github.com/HKUDS/nanobot/issues/5198) | 特定 session 内无法切换模型，除非重新配置整个实例；`/model` 命令看似无效 | 未关闭，尚无对应 fix PR |
| 🟢 低 | [#5187](https://github.com/HKUDS/nanobot/issues/5187) | Termux 中因缺少系统时区数据库导致启动失败（无效时区验证错误） | 已关闭，已有修复（[#5189](https://github.com/HKUDS/nanobot/pull/5189)） |

另外，[PR #5200](https://github.com/HKUDS/nanobot/pull/5200) 和 [PR #5201](https://github.com/HKUDS/nanobot/pull/5201) 分别修复了 exec 会话中 `wait_for` 因响应截断丢失目标和 `AutoCompact.prepare_session()` 对损坏的持久化 `_last_summary` 容错不足的问题，均附带回归测试，目前处于待合并状态。

整体来看，今日 Bug 修复效率很高，高严重度问题均已有 PR 响应。

## 功能请求与路线图信号

- **模型切换的交互缺陷（[#5198](https://github.com/HKUDS/nanobot/issues/5198)）** 用户指出 WebUI 中点击模型 blip 无法切换模型，与 Cloud SaaS AI 产品的 UI 惯例不符。这不仅是 bug，也指向多模型工作流中"per-session 模型选择"的需求。目前无对应 PR，但考虑到该 issue 影响核心交互体验，有可能被纳入近期迭代。
- **DeepSeek Responses API 支持（[PR #5197](https://github.com/HKUDS/nanobot/pull/5197)，待合并）** 将 `deepseek-v4-flash` 路由到 DeepSeek 原生 Responses API，复用现有 streaming 与 function-tool 机制。这延续了 NanoBot 对多 provider 的持续接入策略。
- **Quick Chat / Temporary Chat（[PR #5184](https://github.com/HKUDS/nanobot/pull/5184)，待合并）** 为 WebUI 增加持久化的快速会话入口和可选的临时会话（仅内存历史），属于面向日常使用体验的功能扩展，方向与 #5198 中用户对 WebUI 交互的期待一致。
- **长期挂起的功能 PR**：session 导出/导入/搜索/统计命令（[#1565](https://github.com/HKUDS/nanobot/pull/1565)，3 月创建）和 `nanobot skill status` 命令（[#1319](https://github.com/HKUDS/nanobot/pull/1319)，2 月创建）仍未合并，两者目前均带 `conflict` 标记，需要维护者决策或协助解决冲突。

## 用户反馈摘要

- **微信渠道用户**（来自 #5195 的评论与描述）：用户依赖 WebUI 重新扫码登录，但遭遇 session 无限过期死循环（暂停 60 分钟后继续用旧 token 轮询），体验从"可以恢复"退化为"必须手动干预"。这类渠道稳定性问题对真实使用场景影响很大。
- **Termux 用户**（来自 #5187）：用户以探索心态尝试在 Termux 运行 NanoBot，但被时区数据缺失问题直接挡住。这表明项目对"最小 Linux 环境"的兼容性已有关注（相关修复 #5189 已合并），但从 issue 提出到修复仍有约 1 天间隔。
- **Windows WebUI 用户**（来自 #5190）：前端资源因 MIME type 错误而无法加载，属于平台特有的环境问题。修复 PR #5191 已提出，但尚未合并。

## 待处理积压

- **[#5198](https://github.com/HKUDS/nanobot/issues/5198) session 内模型切换不可用**（新开，无 PR）：影响核心使用体验，建议尽快确认设计意图（是 bug 还是尚未支持）。
- **[#5190](https://github.com/HKUDS/nanobot/issues/5190) Windows MIME type 问题**（新开，已有 PR #5191）：虽已有修复，但 PR 仍待审查，考虑到该问题会让 Windows 用户 WebUI 完全不可用，建议优先处理。
- **[PR #1656](https://github.com/HKUDS/nanobot/pull/1656)（schema 验证处理 None）、[PR #1319](https://github.com/HKUDS/nanobot/pull/1319)（skill status）、[PR #1565](https://github.com/HKUDS/nanobot/pull/1565)（session 导入导出/搜索/统计）**：均为 2-3 个月前创建的 PR，至今标记 `conflict` 且无维护者介入。随着 SQLite 迁移（#5173）落地，#1565 的 session 管理功能需要基于新存储层重新适配，建议维护者明确告知社区这些 PR 的去留。

---

*本日报基于 2026-08-01 00:00 UTC 的 GitHub 数据自动生成。数据来源为 HKUDS/nanobot 仓库的公开 Issues 与 Pull Requests。*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent 项目动态日报 — 2026-08-01

## 今日速览

过去 24 小时内项目保持高活跃度：共更新 16 条 Issue（15 条活跃/新增，1 条关闭）与 50 条 PR（46 条待合并，4 条已合并/关闭）。PR 提交量显著高于 Issue 量，说明社区贡献热情高涨，但 46 条待合并 PR 也提示维护者 review 带宽存在一定积压。Issue 侧单日关闭率较低（1/16），多个 P2 级 bug 已在同日收获对应修复 PR，整体项目处于“高产出、待消化”的积极状态。今日无新版本发布，大量修复与功能正汇聚于 0.19.x 之后的管线中。

## 版本发布

今日无新版本发布（最新 Releases 为空）。

## 项目进展

今日无已合并/关闭 PR 的详细清单，但 50 条 PR 更新覆盖了基础设施、平台适配、稳定性、性能等多个方向。其中值得关注的重要改动包括：

- **Slack 平台接入策略修正**（[PR #75739](https://github.com/NousResearch/hermes-agent/pull/75739)）：不再因缺少专用 Slack 工具模型而绝对拒绝 API 访问，改为引导 agent 使用已认证的 Slack skill/CLI，保留操作边界。对 Slack 工作流用户是实质性的能力放开。
- **Discord 适配器解耦**（[PR #75735](https://github.com/NousResearch/hermes-agent/pull/75735)）：将 7,092 行的 `DiscordAdapter` 中的授权逻辑抽取为独立 Mixin，模块从 9,841 行降至约 9,300 行。这是对 "god-file" 持续重构的一部分。
- **上下文长度自定义覆盖修复**（[PR #75738](https://github.com/NousResearch/hermes-agent/pull/75738)）：`custom_providers[].models.<id>.context_length` 此前仅在启动和 `/model` 切换时生效，现扩展到所有调用点。
- **备份健壮性修复**（[PR #75728](https://github.com/NousResearch/hermes-agent/pull/75728)）：修复 `hermes update --backup` 将任意 `.db` 后缀文件当作 SQLite 处理导致备份中止的问题，直接回应用户报告。
- **MCP 可见空白字符告警**（[PR #75736](https://github.com/NousResearch/hermes-agent/pull/75736)）：配置中隐藏的前导/尾随空白（如复制 token 时带入换行）现在会启动时警告并指明具体 key 路径，减少“401 连接失败”类的隐晦问题。

综合来看，今日 PR 主要围绕平台接入（Slack/Discord/WhatsApp）、备份与更新安全、内存/技能作用域隔离、LSP 扩展四大主题推进。

## 社区热点

评论数最高的 Issue：

- **[#75598 [Bug] issue with updates](https://github.com/NousResearch/hermes-agent/issues/75598)**（4 条评论，P2，Windows）：用户报告约一周前开始更新后程序不稳定，多个 gateway 因不同 profile 互相冲突，切换 profile 未停用其他组件。该问题横跨 CLI、桌面端、Windows 平台与安装更新，涉及面广、社区关注度高。
- **[#66084 TUI npm 安装检查误判](https://github.com/NousResearch/hermes-agent/issues/66084)**（3 条评论，P2）：`_tui_need_npm_install()` 用整个 monorepo 的 lockfile 与 workspace 级安装对比，导致 dashboard 每次启动都触发重装。这是一个影响日常开发体验的“慢性”问题。
- **[#43800 Honcho 插件忽略 baseUrl 配置](https://github.com/NousResearch/hermes-agent/issues/43800)**（3 条评论，P3）：自托管用户配置被静默忽略、流量被路由到生产云，涉及数据隐私敏感点，讨论热度持续。

核心诉求集中在：**更新/安装链路稳定性**、**自托管配置可信度**、**多 profile 隔离正确性**等。

## Bug 与稳定性

按严重程度排列（P2 优先；标注修复 PR 状态）：

| 严重度 | Issue | 描述 | 修复状态 |
|---|---|---|---|
| P2 | [#75598](https://github.com/NousResearch/hermes-agent/issues/75598) | 更新后程序不稳定，多 gateway 冲突 | 无 PR；待调查 |
| P2 | [#66084](https://github.com/NousResearch/hermes-agent/issues/66084) | TUI npm 安装判断每次误报 | 无 PR；已讨论 3 周 |
| P2 | [#75724](https://github.com/NousResearch/hermes-agent/issues/75724) | 备份因非 SQLite `.db` 文件中止（Windows） | ✅ [PR #75728](https://github.com/NousResearch/hermes-agent/pull/75728) |
| P2 | [#75727](https://github.com/NousResearch/hermes-agent/issues/75727) | 桌面端新会话使用 localStorage 模型而非配置默认 | 无 PR |
| P2 | [#75708](https://github.com/NousResearch/hermes-agent/issues/75708) | mem0 插件忽略 session key 的用户作用域 | 无 PR；root cause 已定位 |
| P2 | [#75694](https://github.com/NousResearch/hermes-agent/issues/75694) | hermes 执行 chown 后自身 venv 崩溃（权限问题） | 无 PR；needs-repro |
| P2 | [#75695](https://github.com/NousResearch/hermes-agent/issues/75695) | 同源问题：dashboard 报 Frontend not built | 无 PR；needs-repro |
| P3 | [#75725](https://github.com/NousResearch/hermes-agent/issues/75725) | MiniMax-M3 首次工具调用后停止思考 | 无 PR |
| P3 | [#75731](https://github.com/NousResearch/hermes-agent/issues/75731) | 桌面端 `@file` 附件在历史 hydration 后重复渲染 | ✅ [PR #75733](https://github.com/NousResearch/hermes-agent/pull/75733) |
| P3 | [#75710](https://github.com/NousResearch/hermes-agent/issues/75710) | 侧边栏长标题被截断为单行省略号 | 无 PR |

值得注意：kangarooo 用户在 [#75693](https://github.com/NousResearch/hermes-agent/issues/75693) / [#75694](https://github.com/NousResearch/hermes-agent/issues/75694) / [#75695](https://github.com/NousResearch/hermes-agent/issues/75695) 连报三类问题（provider 失败、venv 权限崩坏、dashboard 不可用），核心诱因疑似一次 `chown` 操作，提示工具对敏感 shell 命令的防护仍需加强。

## 功能请求与路线图信号

- **Per-subagent 工具集限制**（[Issue #75737](https://github.com/NousResearch/hermes-agent/issues/75737)）：用户要求 `delegate_task` 支持按子 agent 限制工具集，避免每个子 agent 继承全部 21 个工具集导致系统提示词膨胀数千 tokens。这是对 token 成本敏感的团队的真实痛点，很可能进入下一版本的能力规划。
- **Laravel Blade LSP 支持**（[Issue #75718](https://github.com/NousResearch/hermes-agent/issues/75718)）：已有对应实现 [PR #75720](https://github.com/NousResearch/hermes-agent/pull/75720)，注册 laravel-lsp 处理 `.blade.php` 文件。若 review 通过，将随下个版本落地。
- **Discord god-file 解耦**（[Issue #75734](https://github.com/NousResearch/hermes-agent/issues/75734)）：架构重构请求，已有对应 [PR #75735](https://github.com/NousResearch/hermes-agent/pull/75735)，属于持续代码健康治理。
- **被关闭的功能请求**：[#71045](https://github.com/NousResearch/hermes-agent/issues/71045)（author 主动撤回），此前处于 needs-decision 状态。

## 用户反馈摘要

- **更新信任危机**：用户在 #75598 中反馈“此前所有更新都顺利，约一周前开始出问题”，且“多 gateway 因不同 profile 互相冲突，切换 profile 不关停其他组件”。这指向 **profile 生命周期管理**的不确定性，更新机制的历史稳定性正在被近期回归消耗。
- **备份与升级的安全感**：#75724 用户遭遇备份中途失败，暴露了“全量备份”对非 SQLite `.db` 文件的脆弱假设。好消息是修复 PR 当天即提交，响应迅速。
- **自托管数据边界敏感**：#43800 的 Honcho 插件问题意味着自托管用户可能**在不知情情况下将记忆数据路由到生产云**，属于隐私/信任层面的严重隐患，社区诉求强烈。
- **桌面端体验细节**：#75727 新会话默认模型不符合 profile 配置、#75710 长标题（波斯语 RTL）被截断，反映桌面端在多语言与配置一致性上仍有改善空间。
- **权限操作缺乏 failsafe**：kangarooo 的连环问题（#75693-75695）显示工具执行 `chown` 后直接破坏了自身运行环境，用户明确期望“有 failsafe”。

## 待处理积压

以下 Issue/PR 长期未获解决或合入，建议维护者优先关注：

- **[Issue #43800](https://github.com/NousResearch/hermes-agent/issues/43800)**（2026-06-10 创建，P3）：Honcho 插件静默路由到生产云，已开放 7 周，数据隐私敏感，3 条评论后无后续。
- **[PR #39643](https://github.com/NousResearch/hermes-agent/pull/39643)**（2026-06-05 创建，P2）：API 会话 ID 随机化与 Idempotency-Key 安全修复，带 security 标签，已开放近 2 个月仍待合并。
- **[PR #56833](https://github.com/NousResearch/hermes-agent/pull/56833)**（2026-07-02 创建，P3）：MCP 熔断错误信息措辞软化，防止模型过度遵循指令而自我锁死。行为学细节重要，已开放 1 个月。
- **[Issue #75693](https://github.com/NousResearch/hermes-agent/issues/75693)**（2026-07-31 创建，P2）：provider 失败降级逻辑相关，needs-repro，但与 #75694/#75695 同源，待官方复现确认。

---

**报告生成时间**：2026-08-01  
**数据窗口**：过去 24 小时  
**数据来源**：NousResearch/hermes-agent GitHub 仓库公开数据

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目动态日报 — 2026-08-01

## 今日速览

过去 24 小时内，PicoClaw 共产生 2 条 Issue 更新（新开/活跃 2 条，关闭 0 条）和 3 条 PR 更新（待合并 3 条，合并/关闭 0 条），无新版本发布。整体来看，项目保持稳步迭代节奏：社区讨论集中在 IRC 长消息支持和 UI 性能问题上，而代码侧有 3 个功能型 PR 正在推进中，但尚未有合并动作。活跃度属于中等水平，值得关注的是 PR 和 Issue 均有较长时间的积压，需要维护者加快 review 和合并节奏。

## 版本发布

今日无新版本发布，最新 release 仍为 [PicoClaw v0.3.1](https://github.com/sipeed/picoclaw/releases)（用户反馈中已开始使用该版本）。

## 项目进展

今日无 PR 被合并或关闭，但 3 个待合并 PR 均在近期（7 月 31 日）有更新，说明正处于 review/迭代阶段：

- **[#3222] refactor(deltachat): cleanup implementation, documentation -200LOC** — 由 [trufae](https://github.com/sipeed/picoclaw/pull/3222) 提交，重构 DeltaChat 模块，删除约 200 行代码，移除 legacy 特性和冗余测试，并改用官方 relay list 维护方式。该 PR 若合并将使 DeltaChat 集成更轻量、更易维护。
- **[#3193] Added simplex channel type** — 由 [dim](https://github.com/sipeed/picoclaw/pull/3193) 提交，为 PicoClaw 新增 [SimpleX](https://simplex.chat/) 聊天协议通道支持，属于新功能开发。
- **[#3200] feat(models): add configurable default fallback chain** — 由 [lc6464](https://github.com/sipeed/picoclaw/pull/3200) 提交，允许用户在 Web UI 中配置模型默认降级链（如设置首选模型、按顺序 fallback），并通过后端 API 持久化。

整体来看，项目正在朝 **多协议扩展（SimpleX）、AI 模型配置灵活性提升** 和 **代码精简维护** 三个方向迈进，但目前合并速度偏慢，功能落地进度取决于后续 review 效率。

## 社区热点

- **[#3287] [Feature] Better support long messages in IRC**（👍 2 评论，最受关注）
  链接：https://github.com/sipeed/picoclaw/issues/3287  
  这是目前讨论最活跃的 Issue。用户 [superuser-does](https://github.com/sipeed/picoclaw/issues/3287) 指出 IRC 协议默认限制消息为 512 字节，超过后会被客户端自动拆分，而 PicoClaw 目前无法将这些拆分后的片段识别为同一逻辑消息。该诉求直接指向 **IRCv3 长消息处理能力**，反映了 IRC 渠道使用者对消息完整性的真实需求。

- **[#3292] [stale] CPU usage too high when focus on input box in chat interface**（1 评论）  
  链接：https://github.com/sipeed/picoclaw/issues/3292  
  这是性能类 bug 反馈，用户报告在 Firefox 中聚焦聊天输入框时 CPU 占用过高，虽然带 `[stale]` 标记但仍活跃，说明 Web 界面渲染性能问题已经开始引起社区注意。

## Bug 与稳定性

今日报告 1 个 Bug，无新的合并修复 PR：

| 严重程度 | Issue | 描述 | 状态 |
|---|---|---|---|
| 🟠 中高 | [#3292](https://github.com/sipeed/picoclaw/issues/3292) | 聊天界面输入框聚焦时 CPU 占用过高（环境：Debian/Linux x64 + Firefox，PicoClaw v0.3.1，Go 1.26，使用 deepseek-v4-flash 模型） | 无对应 PR，待排查 |

该 Issue 为前端性能问题，常见于富文本输入框或持续轮询事件处理。由于用户已报告具体版本和浏览器环境，建议维护者优先复现并定位是否与光标闪烁动画或 React 状态更新有关。

## 功能请求与路线图信号

- **[#3287] IRC 长消息整体处理**（https://github.com/sipeed/picoclaw/issues/3287）  
  用户明确希望 PicoClaw 理解 IRCv3 中的长消息（超过 512 字节自动拆分）应被视为单个完整消息。这是一个 **协议层智能化处理** 的需求，当前无对应 PR，但结合已存在的 IRC 相关功能，有可能进入下一个小版本迭代。

- **[#3193] SimpleX 通道支持**（https://github.com/sipeed/picoclaw/pull/3193）  
  虽为 PR 而非 Issue，但其开放已超过 1 个月仍未被合并，说明维护者可能对其存在一些质疑（如加密协议集成复杂度、社区需求验证）。若该 PR 最终合入，将显著拓展 PicoClaw 的隐私保护通道能力。

- **[#3200] 模型默认 fallback 链**（https://github.com/sipeed/picoclaw/pull/3200）  
  这是可配置功能请求，与当前 AI 多模型切换趋势高度符合。该 PR 已实现 Web UI 配置与后端 API 持久化，功能相对完整，属于 **高概率进入下一版本** 的 Feature。

**信号判断**：项目的路线图似乎偏向「扩展通信协议支持（IRC/simplex）＋ 完善 AI 模型管理」，前者由社区需求驱动，后者由开发者主动贡献推进。

## 用户反馈摘要

- **IRC 用户痛点**（[#3287](https://github.com/sipeed/picoclaw/issues/3287)）：当通过 IRC 发送长文本时，由于 IRC 协议限制，消息会被拆分成多条，用户期望 PicoClaw 能智能感知消息边界，将其聚合为一条完整消息后再交给 AI 模型处理。这表明 **协议适配层需要引入消息重组/缓冲机制**，以保证用户在 IRC 环境下获得连贯的对话体验。
- **Web UI 性能反馈**（[#3292](https://github.com/sipeed/picoclaw/issues/3292)）：用户在 Firefox 中聚焦输入框时 CPU 占用异常飙高，直接影响了日常使用体验。这是继功能需求外的 **稳定性负面反馈**，建议维护者重点关注。

## 待处理积压

当前有 3 个 PR 长时间未合并，均已接近或超过 1 个月开放期，建议维护者优先处理：

1. **[#3193] simplex 通道支持**（开放 35 天；创建于 2026-06-27）  
   https://github.com/sipeed/picoclaw/pull/3193  
   ⚠️ 最后一个新通道功能 PR，长时间搁置可能导致代码冲突加剧。

2. **[#3200] 模型默认 fallback 链**（开放 31 天；创建于 2026-07-01）  
   https://github.com/sipeed/picoclaw/pull/3200  
   ⚠️ 功能完整但未合并，可能阻塞后续模型配置相关需求。

3. **[#3222] deltachat 清理/重构**（开放 29 天；创建于 2026-07-03）  
   https://github.com/sipeed/picoclaw/pull/3222  
   ⚠️ 重构类 PR 风险较低，但长期不合并会导致与主干代码的潜在冲突。

另外，Issue [#3287](https://github.com/sipeed/picoclaw/issues/3287) 虽开放不足 10 天，但评论活跃且功能诉求明确，也建议纳入近期规划。**项目健康度总评**：功能开发活跃，但 PR 合并周期偏长，社区反馈的 Bug 尚无人认领，维护响应速度需要提升。

---

*报告生成时间：2026-08-01 | 数据来源：GitHub API*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目动态日报 — 2026-08-01

## 1. 今日速览

今日 NanoClaw 项目保持活跃，过去 24 小时内有 **8 条 Issue 更新**（全部处于开放/活跃状态，无关闭）与 **9 条 PR 更新**（6 条待合并，3 条已合并/关闭），显示社区参与度和维护者响应均维持在健康水平。当前讨论焦点集中在 **多运行时支持（Kubernetes、Apple Container、原生模式）**、**Telegram 配对稳定性**以及**渠道集成扩展（iMessage、Dial）** 三大方向。安全加固类工作持续进行中（日志脱敏、交互响应来源校验）。**无新版本发布**，但有 PR 涉及恢复 v2.1.54 发布路径，预计近期将有版本动作。整体评估：项目处于**活跃迭代期**，社区贡献者参与度高，维护者响应及时。


## 3. 项目进展

过去 24 小时有 **3 个 PR 被合并/关闭**，推进了以下工作：

- **[PR #3163 — fix(release): restore the v2.1.54 release path](https://nanocoai/nanoclaw/pull/3163)**（已关闭，core-team）：修复了 v2.1.54 的发布路径问题，为即将到来的版本发布扫清障碍。
- **[PR #3076 — feat(imessage): unified local+hosted adapter targeting spectrum-ts v11](https://nanocoai/nanoclaw/pull/3076)**（已关闭，Feature skill）：统一 iMessage 本地与托管适配器，目标对齐 spectrum-ts v11，但注意该 PR 侧标注了 superseding #2999 的后续 PR 还在推进中。
- **[PR #1678 — docs(skills): update voice transcription skills for Telegram + Linux](https://nanocoai/nanoclaw/pull/1678)**（已关闭，Docs）：更新语音转录技能文档，移除"仅 WhatsApp"限制，新增 Telegram 与 Linux 支持说明。

> 另注意：PR #3076 虽标记关闭，但与其功能关联的 **PR #3164（Hosted iMessage/Photon 注册流程）** 以"supersede #2999"的方式于同日新开，说明 iMessage 集成功能仍在积极迭代中，尚未完全落地。


## 4. 社区热点

今日讨论热度最高的条目集中在以下三个 Issue（各 3 条评论）：

- **[Issue #1184 — Challenges deploying nanoclaw in restricted K8s environments (Sealos)](https://nanocoai/nanoclaw/issues/1184)**（评论 3，👍 1）：用户认可 NanoClaw 的极简轻量设计，但反馈在受限 K8s 环境（Sealos）中部署遇到困难。背后诉求是**企业用户希望在已有容器平台上运行 NanoClaw**，而非受限于本地 Docker。
- **[Issue #1732 — feat: native runner mode — bypass Docker for host-tool access](https://nanocoai/nanoclaw/issues/1732)**（评论 3）：明确提出**原生运行模式需求**，绕过容器以获得 tmux、有头浏览器、macOS API 等宿主工具访问能力。这是容器隔离优势与宿主集成需求之间的核心矛盾点。
- **[Issue #1225 — Run it without docker](https://nanocoai/nanoclaw/issues/1225)**（评论 2）：老 Issue 但今日仍被更新，用户询问无 Docker 环境（Windows/Linux）下的运行方式。

**分析**：社区当前最集中的诉求是 **"摆脱 Docker 依赖"**，包括直接在宿主机运行、接入 K8s 集群、以及支持 Apple Container。这与 PR #2809（Apple Container runtime）、PR 讨论中的 K8s 方案形成呼应，是明确的路线图信号。


## 5. Bug 与稳定性

按严重程度排列：

- **[Issue #3162 — [Type: Bug, Priority: High] Telegram pairing is silently broken for the whole process lifetime if the boot-time getMe fails](https://nanocoai/nanoclaw/issues/3162)**（新开，0 评论）：**高优先级 Bug**。Telegram 频道启动时若 getMe 调用失败（慢网络/代理抖动/Telegram 临时故障），配对码将**在整个进程生命周期内永久失效**且无任何用户提示。影响：用户被静默锁定在配对功能之外。**目前无对应 fix PR**，需要尽快处理。

- **[Issue #2923 — [Hardening] [security] ask_user_question card can be defaced by a forged click before origin authz](https://nanocoai/nanoclaw/issues/2923)**（0 评论）：**安全加固类 Bug**。伪造的按钮点击可在源校验之前覆盖 ask_user_question 卡片展示文本，属于显示/完整性欺骗。**已有对应 PR #2651**（针对同一边界做来源验证），但尚未合并。

- **[Issue #2589 — Apple Container: host.docker.internal in OneCLI proxy URL doesn't resolve from inside the microVM](https://nanocoai/nanoclaw/issues/2589)**（1 评论）：Apple Container 微虚拟机内无法解析 `host.docker.internal`，且不支持 `--add-host` 注入。影响 Apple Container 环境下 OneCLI 代理功能。

- **[Issue #2588 — skill/apple-container branch is substantially out of sync with mainline](https://nanocoai/nanoclaw/issues/2588)**（1 评论）：`skill/apple-container` 分支与主分支严重失步，引用了不存在的 API/模块、基于 Node+tsc 而非主线的 bun 运行时，导致 `/convert-to-apple-container` skill 立即失败。


## 6. 功能请求与路线图信号

当前开放的功能请求与对应的 PR 推进情况：

| 功能请求 | 状态 | 对应 PR | 纳入下一版本可能性 |
|---------|------|---------|------------------|
| **原生 runner 模式**（绕过 Docker 获取宿主工具，#1732） | 🔥 高热度 | 无直接 PR，但在 #1225（无 Docker 运行）的持续讨论中 | 较高 |
| **Kubernetes 容器运行时**（#2354） | 有 👍 支持 | 无直接 PR | 中高，社区呼声明显 |
| **Apple Container 运行时** | 有专门 PR 推进中 | **[PR #2809](https://nanocoai/nanoclaw/pull/2809)**（Apple Container runtime + 远程 OneCLI gateway，6 月创建，7/31 更新待合并） | 高，但需先解决 #2588/#2589 的兼容性问题 |
| **Dial 频道适配器**（SMS + AI 语音呼叫） | 新功能 | **[PR #3041](https://nanocoai/nanoclaw/pull/3041)**（7/14 创建，待合并） | 高 |
| **Hosted iMessage (Photon) 注册流程** | 核心团队推进 | **[PR #3164](https://nanocoai/nanoclaw/pull/3164)**（今日新开，core-team） | 高 |
| 无 Docker 环境运行（#1225） | 长期存在 | 可能由 #1732 原生模式覆盖 | 待定 |


## 7. 用户反馈摘要

- **对极简设计的好评**（#1184）：用户明确赞赏 NanoClaw "minimalist approach" 和"lightweight, secure alternative to the more bloated agent frameworks"，认为"Using existing code agents to build such a streamlined 'Claw' is brilliant"——产品定位获得了真实用户认可。
- **生产环境部署受阻**（#1184, #1225）：企业/生产环境用户遭遇 Docker 硬依赖带来的部署障碍，尤其在受限 K8s 环境（如 Sealos）中问题更加突出。用户表达"eager to run this in my production environment"但被环境限制挡住。
- **容器隔离 vs 宿主集成两难**（#1732）：用户理解容器隔离是安全优势，但也指出它"becomes a hard blocker for a growing class of agentic use-cases"，需要 tmux、有头浏览器、macOS API 等宿主集成能力。
- **Apple Container 用户面临断裂体验**（#2588）：分支失步导致技能直接失败，说明一些生态分支的维护节奏未能跟上主线迁移（Node→bun），影响用户信任。


## 8. 待处理积压

以下重要项目长期未获解决，建议维护者重点关注：

- **[Issue #2588 — skill/apple-container 分支与主线失步](https://nanocoai/nanoclaw/issues/2588)**（5/22 创建，逾 2 个月）：文档化的迁移技能开箱即失败，直接影响 Apple Container 用户。**与其关联的 PR #2809 已在待合并队列中，建议优先协调合并并同步修复分支失步问题。**

- **[Issue #2589 — Apple Container 中 host.docker.internal 无法解析](https://nanocoai/nanoclaw/issues/2589)**（5/22 创建，逾 2 个月）：阻塞 Apple Container 环境的 OneCLI 代理功能。同上，建议与 PR #2809 合并一并解决。

- **[Issue #1225 — Run it without docker](https://nanocoai/nanoclaw/issues/1225)**（3/18 创建，逾 4 个月）：老 Issue 今日仍有更新，反映用户持续关注无 Docker 运行方案。

- **[PR #2954 — docs(security): add reporting and triage policy](https://nanocoai/nanoclaw/pull/2954)**（7/4 创建，近 1 个月未合并）：安全文档 PR，涉及漏洞报告与分级策略，建议尽快审阅合并以完善社区安全协作机制。

---

**项目健康度总结**：NanoClaw 社区活跃度高、贡献者参与积极、核心团队响应迅速。当前最大的技术债集中在 **Apple Container 分支维护**与 **Telegram 配对可靠性**两块；最大的产品方向机会在于 **扩展容器运行时选项（K8s/原生/Apple Container）** 以解锁更多企业级部署场景。安全加固持续进行中，整体项目处于健康上升期。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报 — 2026-08-01

> 数据窗口：2026-07-31 至 2026-08-01（UTC） | 数据源：github.com/nullclaw/nullclaw

---

## 1. 今日速览

过去 24 小时内，NullClaw 项目整体活跃度较低：**无新 Issue、无 Issue 关闭、无新版本发布**，仅出现 1 条待合并 PR（#981）。该 PR 于 7 月 29 日提出，7 月 31 日有更新，目前仍处于 Open 状态等待审查。虽然数据面较为平静，但 PR #981 的持续活跃表明项目仍在推进第三方 Provider 生态的扩展，属于低频但方向清晰的迭代周期。

---

## 2. 版本发布

今日无新版本发布。

---

## 3. 项目进展

今日无 PR 被合并或关闭，因此没有代码变更正式进入主干。唯一值得关注的进展是 **PR #981 仍处于 Open 且待合并状态**，该 PR 若被合并，将为项目新增一个基于 xAI Grok CLI 的 provider，进一步丰富多渠道接入能力。

- **待合并 PR：**
  - [#981 [OPEN] feat(provider): add grok-cli provider for xAI Grok CLI](https://github.com/nullclaw/nullclaw/pull/981)  
    作者：valonmulolli | 创建：2026-07-29 | 更新：2026-07-31
  - 摘要：新增可选 provider，通过 spawn-per-request 模式调用本地 `grok` CLI，与现有 `codex-cli` / `gemini-cli` / `claude-cli` provider 保持一致。

**整体评估**：项目今日处于“审查与等待”阶段，核心功能开发在 PR 层面已就绪，等待维护者合并。合并后 Provider 矩阵将进一步扩展，提升对 xAI 生态的兼容性。

---

## 4. 社区热点

今日唯一活跃的 PR 为 **#981**，虽然评论数和 👍 数暂未公开，但其存在本身即构成当前社区的焦点：

- [PR #981: feat(provider): add grok-cli provider for xAI Grok CLI](https://github.com/nullclaw/nullclaw/pull/981)

**热点分析**：该 PR 的提出反映了用户对 **xAI Grok CLI 接入本地 AI 助手** 的需求。从实现方式（复用既有 CLI provider 模式）来看，社区希望以一致、可插拔的方式扩展更多本地 CLI 工具作为后端，而非为每个新模型单独开发定制集成。这也印证了 NullClaw 作为个人 AI 助手中立 Hub 的定位。

---

## 5. Bug 与稳定性

今日无新报告的 Bug、崩溃或回归问题。仓库在稳定性方面保持安静状态，无紧急修复需求。

---

## 6. 功能请求与路线图信号

尽管今日无新 Issue，但 **PR #981 本身就是一个明确的路线图信号**：

- **新增 Provider 类型**：基于 CLI 的可选 provider，用户需自行安装并认证 `grok` CLI 后即可使用。这意味着未来版本可能继续沿用“CLI wrapper”模式接入更多 AI 服务。
- **生态扩展方向**：从 codex-cli → gemini-cli → claude-cli → grok-cli，项目正在系统性地覆盖主流 AI 厂商的命令行工具，推测下一版本（如有）将包含此 Provider。

该 PR 若通过审查，预计将进入下一个 minor 或 patch 版本。

---

## 7. 用户反馈摘要

由于今日无 Issue 评论或新增 Issue，以下反馈基于 PR #981 的描述逻辑推演：

- **使用场景**：用户已在本机安装并配置 xAI Grok CLI，希望通过 NullClaw 作为统一入口与其交互。
- **核心诉求**：希望 NullClaw 对本地已有 CLI 工具提供“开箱即用”的适配，而不必等待官方 SDK 集成。
- **潜在痛点**：可选 provider 的依赖（grok CLI 的安装与认证）可能给部分用户带来配置门槛，但现有文档和 codex-cli/claude-cli 的先例已提供了参考路径。

---

## 8. 待处理积压

目前仅有一项待处理事项，需维护者关注：

- **[#981 feat(provider): add grok-cli provider for xAI Grok CLI](https://github.com/nullclaw/nullclaw/pull/981)**  
  状态：Open | 创建：2026-07-29 | 最后更新：2026-07-31  
  提醒：该 PR 已等待 3 天，且代码模式与既有 provider 一致，审查成本应较低。建议维护者尽快进行 review 和合并，避免积压导致冲突，同时也能及时响应用户对 xAI 支持的需求。

---

**项目健康度总结**：无阻塞性 Bug、无回归、无安全事件；1 个待审 PR 功能明确、实现规范；Issue 与 PR 活跃度处于低位但并非停滞。整体健康度良好，处于“等待合并、蓄势待发”的状态。

*报告生成时间：2026-08-01*

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw 项目日报 — 2026-08-01

### 1. 今日速览

过去 24 小时 IronClaw 保持着极高的迭代强度：16 条 Issue 更新（15 条活跃、1 条关闭）、50 条 PR 更新（18 条待合并、32 条已合并/关闭），无新版本发布。核心事件是目标架构 WS1 系列重构迎来三连合并（WS1.1→WS1.3），WS1.4/WS1.5 已入队待合并，架构落地进度过半。与此同时，社区反馈了多项安全与稳定性问题——包括 P0 级跨用户内存泄漏（[#6900](https://github.com/nearai/ironclaw/issues/6900)）、新用户邮箱认证失败（[#6972](https://github.com/nearai/ironclaw/issues/6972)）以及 CI 队列结构性阻塞（[#6978](https://github.com/nearai/ironclaw/issues/6978)），整体活跃度健康，但维护者需在功能推进与稳定性修复之间做出权衡。

### 2. 版本发布

今日无新版本发布。仓库内 [#5598](https://github.com/nearai/ironclaw/pull/5598) release PR 已等待近 30 天，其中包含 `ironclaw_common` 0.5.0 与 `ironclaw_skills` 0.4.0 两个破坏性变更，建议维护者明确合并时间表。

### 3. 项目进展

今日最重要的进展是目标架构 **WS1（契约提取）系列** 在一天内完成三连合并，为后续架构演进奠定了中立契约层基础：

- **[PR #6967](https://github.com/nearai/ironclaw/pull/6967)（WS1.1，已合并）**：完成 `ironclaw_host_api` 中的 turn 词汇表，退役 turns shim。
- **[PR #6975](https://github.com/nearai/ironclaw/pull/6975)（WS1.2，已合并）**：提取 `ironclaw_loop_contracts`，将 `ironclaw_agent_loop` 切换至新契约层。
- **[PR #6977](https://github.com/nearai/ironclaw/pull/6977)（WS1.3，已合并）**：提取 `ironclaw_extension_contracts`，关闭双导入路径。
- 后续 **[PR #6980](https://github.com/nearai/ironclaw/pull/6980)（WS1.4）** 与 **[PR #6981](https://github.com/nearai/ironclaw/pull/6981)（WS1.5，安全敏感）** 已形成堆叠队列，仅差 review。

其他已合并/关闭的重要 PR：

- **[#6930](https://github.com/nearai/ironclaw/pull/6930)**：托管 MCP 服务器注册功能合入（153 文件，+15,002/-1,818），[#6979](https://github.com/nearai/ironclaw/pull/6979) 已将其与目标架构文档对齐。
- **[#4022](https://github.com/nearai/ironclaw/pull/4022)**：修复 [#4014](https://github.com/nearai/ironclaw/issues/4014) 引入的工具 HTTP 响应错误误终止整个 agent 运行的回归。
- **[#3952](https://github.com/nearai/ironclaw/pull/3952)**：通过 openat2/O_NOFOLLOW 对 LocalFilesystem 进行 TOCTOU 加固。
- **[#3942](https://github.com/nearai/ironclaw/pull/3942)**：PilotAllowlist 由字符串匹配重构为 serde 驱动的强类型枚举。
- **[#6932](https://github.com/nearai/ironclaw/pull/6932)**：依赖组批量升级（34 个 crate）。

此外，前置任务 [Issue #6920](https://github.com/nearai/ironclaw/issues/6920) 已在今日关闭，目标架构基线、依赖清理与例外收紧的工作宣告完成。

### 4. 社区热点

- **[Issue #6963](https://github.com/nearai/ironclaw/issues/6963)（5 条评论，CI 门禁追踪）**：Core team 成员 BenKurrek 指出 #6946 未能重写 8 个路径键控的 CI 门禁，并明确"checklist row is weak tracking"。这说明 CI 门禁的路径耦合问题已在开发流程中反复造成阻塞，社区（尤其是核心贡献者）对 CI 可维护性的关注度在上升。
- **[Issue #6940](https://github.com/nearai/ironclaw/issues/6940)（2 条评论，IronHub CTA 404）**：所有技能的 IronHub CTA 按钮都跳转 404，影响面覆盖全部技能，但归属方不明确，评论中正在讨论该缺陷由哪一方认领。

### 5. Bug 与稳定性

按严重程度排列：

**🔴 P0 级**

- **[Issue #6900](https://github.com/nearai/ironclaw/issues/6900)（suggested_P0，安全）**：共享通道的默认主题绑定将所有用户的内存读写坍缩到操作员的命名空间，构成跨用户内存泄漏。当前无关联 fix PR，属于必须优先处置的安全漏洞。

**🟠 P1/P2 级**

- **[Issue #6974](https://github.com/nearai/ironclaw/issues/6974)**：libSQL 工具密集型场景 p95 延迟高达 37–135 秒（post-#6696 回归）。同源修复 [PR #6973](https://github.com/nearai/ironclaw/pull/6973) 正在处理 Postgres 容量恢复，可一并关注。
- **[Issue #6972](https://github.com/nearai/ironclaw/issues/6972)**：新账户通过邮箱注册后无法完成认证与登录，直接影响新用户 onboarding。
- **[Issue #6866](https://github.com/nearai/ironclaw/issues/6866)**：所有用户共享同一 home 目录，工作区互相可见，属于隐私信任问题。
- **[Issue #6940](https://github.com/nearai/ironclaw/issues/6940)**：IronHub 技能 CTA 全量 404，影响所有技能的对外转化入口。

**🟡 CI / 基础设施**

- **[Issue #6978](https://github.com/nearai/ironclaw/issues/6978)**：`reborn-tests.yml` 的 `workflow_dispatch` 运行存在结构性缺陷——`critical-mutation` 在 dispatch 事件下被跳过但又被 roll-up 判定为不允许，导致干净运行也必然红灯。
- **[Issue #6976](https://github.com/nearai/ironclaw/issues/6976)**：`ironclaw service install` 不启用 systemd user lingering，导致 VM/无头服务器无法可靠无人值守运行。

### 6. 功能请求与路线图信号

- **[Issue #6939](https://github.com/nearai/ironclaw/issues/6939)（迁移工具）**：用户反馈从 Hermes/Openclaw 迁移至 IronClaw 的成本过高，无法携带既有配置与记忆。这是留存相关的高价值功能，建议评估纳入 V1 范围。
- **[Issue #6941](https://github.com/nearai/ironclaw/issues/6941)（技能 Epic）**：新 Epic 聚焦"模型可发现、可选择、可使用的技能"及自建技能收益度量，配套 PR [#6938](https://github.com/nearai/ironclaw/pull/6938) 已落地"模型选择技能而非关键字打分器"的改动。
- **[Issue #6971](https://github.com/nearai/ironclaw/issues/6971)（术语统一）**：用户要求澄清 "Tools" 与 "Extensions" 的边界与归属关系，与 [#6854](https://github.com/nearai/ironclaw/issues/6854) 的品牌一致性需求合流，说明产品对外命名和文档体系正进入收敛期。
- **[Issue #6578](https://github.com/nearai/ironclaw/issues/6578)（管理员托管 Agents Epic）**：持续活跃，目标是让租户管理员创建非人类主体（agents、automations、inbound channels），而不破坏私有用隔离模型。

### 7. 用户反馈摘要

- **新用户认证受阻**（[#6972](https://github.com/nearai/ironclaw/issues/6972)）：用户创建账户后无法认证并进入产品，属于最紧急的体验断点。
- **迁移成本高**（[#6939](https://github.com/nearai/ironclaw/issues/6939)）：存量用户明确表示"不愿意从空白状态重新开始"，若无迁移工具可能直接流失。
- **品牌不一致**（[#6854](https://github.com/nearai/ironclaw/issues/6854)）：扩展页描述仍在使用 "Reborn" 品牌，与 "Ironclaw 1.0" 的外部口径不符；[PR #6970](https://github.com/nearai/ironclaw/pull/6970) 已在同步清理文档中的 Reborn 术语。
- **隐私顾虑**（[#6866](https://github.com/nearai/ironclaw/issues/6866)）：用户发现所有用户共享同一 home 目录且工作区互相可见，已明确表达这是隐私担忧。
- **文档脱节**（[#6970](https://github.com/nearai/ironclaw/pull/6970)）：贡献者在处理 #6766 时发现大部分文档与代码库脱节，已顺手完成部分清理。

### 8. 待处理积压

- **[PR #5598](https://github.com/nearai/ironclaw/pull/5598)**：release PR 已开放近 30 天且含破坏性变更，建议给出明确的合并窗口与发布计划。
- **[Issue #6900](https://github.com/nearai/ironclaw/issues/6900)**：P0 级跨用户内存泄漏，无关联修复 PR，应列为最高优先级，优先于一切新功能开发。
- **[PR #6831](https://github.com/nearai/ironclaw/pull/6831)**：标准化消息框架（16 个核心操作 + 12 码错误分类）已开放 4 天，可能被 WS1 系列重构挤压而缺少 review 带宽。
- **[PR #6780](https://github.com/nearai/ironclaw/pull/6780)**：IronHub 深链注册/安装网关（含私有 manifest 源）等待 review。
- **[Issue #6578](https://github.com/nearai/ironclaw/issues/6578)**：管理员托管 Agent Epic 已 9 天无代码进展，仅有评论更新，建议确认排期与负责人。
- **[Issue #6963](https://github.com/nearai/ironclaw/issues/6963)**：8 个路径键控 CI 门禁待重写，且 #6967 的 base CI 存在 6 个仓库级预存失败，CI 队列健康度需要系统性治理。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 — 2026-08-01

## 1. 今日速览

过去 24 小时项目**无新增 Issue**，4 条陈旧 Issue 被自动关闭；PR 侧共 12 条更新，其中 11 条已合并/关闭，**1 条仍待审查**。核心维护者（fisherdaddy）今日集中合并了 3 个 openclaw 稳定性修复（缓存命中率回退、工具协议泄漏）以及 1 个发布准备 PR，表明**项目正处于版本发布冲刺阶段**。但值得警惕的是：今日有一批 2026 年 4 月提交的功能实现 PR 因长期未合并被 stale 机制自动关闭（非 merged），涉及侧边栏拖拽、快捷键提示、骨架屏等明确用户需求，**需维护者确认这些功能是否已以其他形式合入主线，避免功能丢失**。整体活跃度评估：**维护侧高度活跃，社区侧趋于沉寂**。

## 2. 版本发布

今日**无新版本发布**。但合并了发布准备 PR：

- [PR #2416 Release/2026.7.31](https://github.com/netease-youdao/LobsterAI/pull/2416) 已合并，表明 2026.7.31 版本正在收尾，预计近期正式发布。该 PR 模板未填充具体变更明细，建议发布时补充完整 changelog。

## 3. 项目进展

今日合并的 PR 主要围绕 **openclaw 运行时稳定性**与**发布准备**：

- **[PR #2413 fix(openclaw): keep live prompt tool-result history byte-stable across turns](https://github.com/netease-youdao/LobsterAI/pull/2413)**（合并）— 修复 live prompt 组装时重复应用 4x 聚合字符上限导致历史缓存被反复重写的问题，为 DeepSeek 前缀缓存稳定性打基础。
- **[PR #2415 fix(openclaw): drop aggregate cap in live tool-result prompt projection](https://github.com/netease-youdao/LobsterAI/pull/2415)**（合并）— 在 #2413 基础上彻底移除 live 请求中的 aggregate cap，**将 DeepSeek 长会话缓存命中率从 ~57% 恢复到接近 100%**。
- **[PR #2414 fix(cowork): prevent BTW tool protocol leakage](https://github.com/netease-youdao/LobsterAI/pull/2414)**（合并）— 清理边聊（side-chat）结果中的工具调用标记泄漏，保持错误元数据完整。属于数据安全性与协议隔离修复。
- **[PR #2416 Release/2026.7.31](https://github.com/netease-youdao/LobsterAI/pull/2416)**（合并）— 版本发布准备。
- **[PR #2417 fix(sites): add copy success feedback](https://github.com/netease-youdao/LobsterAI/pull/2417)**（合并）— 复用会话复制图标的成功反馈模式，为站点 URL/分享码复制增加成功提示，提升 UI 交互可感知性。

**需要特别关注**：今日被 stale 自动关闭的 6 个 PR（#172、#1308、#1315、#1318、#1320、#1321）均为 2026 年 4 月提交的完整实现，状态为**已关闭而非已合并**，意味着其功能尚未进入主线（详见第 8 节）。

## 4. 社区热点

今日无新增 Issue，社区讨论热度集中在此前提交、今日被 stale 关闭的 4 条功能请求上。虽然它们没有新增评论，但 4 条 Issue 中有 3 条（[#1314](https://github.com/netease-youdao/LobsterAI/issues/1314)、[#1317](https://github.com/netease-youdao/LobsterAI/issues/1317)、[#1319](https://github.com/netease-youdao/LobsterAI/issues/1319)）都来自同一位用户（MaoQianTu），且均有对应实现 PR，构成一组完整的"用户提出诉求 → 开发者实现 → PR 沉淀未合并 → 被 stale 清理"链路：

- **[Issue #1314 支持拖拽调整侧边栏宽度](https://github.com/netease-youdao/LobsterAI/issues/1314)**（评论 2 条）— 用户反映固定 240px 宽度在小屏/大屏/长标题场景下都有体验问题。对应 PR #1315 实现了 180~480px 范围拖拽调整，但未被合并。
- **[Issue #1317 侧边栏按钮显示键盘快捷键 kbd 提示](https://github.com/netease-youdao/LobsterAI/issues/1317)**（评论 2 条）— 用户批评 Ctrl+N / Ctrl+F 等快捷键"发现成本高"，建议以 `<kbd>` 徽标展示。对应 PR #1318 已实现含 macOS 符号转换的完整方案。
- **[Issue #1319 会话列表添加骨架屏加载状态](https://github.com/netease-youdao/LobsterAI/issues/1319)**（评论 2 条）— 启动时空状态闪烁让用户误以为历史记录丢失。对应 PR #1320 已实现 `sessionsLoaded` 标志位 + 骨架屏 UI。

**背后诉求分析**：这三条需求集中于"侧边栏可用性"与"界面反馈可感知性"，反映用户对 AI 助手类应用界面细节的高要求。三者均有完整 PR 但 4 个月未合并，说明项目**近期路线图优先级更偏向核心运行时稳定性而非 UI 打磨**。

## 5. Bug 与稳定性

今日无新 Bug 报告，但有多个既有问题通过 PR 修复：

| 严重程度 | 问题描述 | 状态 |
|---|---|---|
| 🔴 严重 | DeepSeek 长会话缓存命中率从 ~100% 暴跌至 ~57%，源于 live prompt 投影反复改写已缓存历史，破坏前缀缓存稳定性 | 已由 [#2413](https://github.com/netease-youdao/LobsterAI/pull/2413) + [#2415](https://github.com/netease-youdao/LobsterAI/pull/2415) 修复 |
| 🟠 中等 | 边聊（BTW）工具调用协议泄漏：provider tool-call 标记污染侧聊结果，且工具需求无法获得稳定引导响应 | 已由 [#2414](https://github.com/netease-youdao/LobsterAI/pull/2414) 修复 |
| 🟡 一般 | 设置页切换标签后，记忆编辑器/模型连接测试弹窗可能残留为全屏 `absolute inset-0` 覆盖层，界面看似只读 | 修复 PR [#1321](https://github.com/netease-youdao/LobsterAI/pull/1321) **已提交但被 stale 关闭，当前主线可能仍存在该 Bug** |

## 6. 功能请求与路线图信号

- **可能被纳入下一版本的功能**：
  - [Issue #1311 表格内容换行保留原始标签 + 长文本 hover 展示全文](https://github.com/netease-youdao/LobsterAI/issues/1311) — UI 细节改进，4 月提出后无对应 PR 被合并，是否已通过其他方式解决待确认。
  - [Issue #1314 拖拽调整侧边栏宽度](https://github.com/netease-youdao/LobsterAI/issues/1314) / [Issue #1317 快捷键可视化提示](https://github.com/netease-youdao/LobsterAI/issues/1317) / [Issue #1319 会话列表骨架屏](https://github.com/netease-youdao/LobsterAI/issues/1319) — 三个 UI 功能均有完整 PR 实现但未合并，若维护者重新评估，可低成本纳入下一版本。

- **路线图信号**：核心维护者（fisherdaddy）今日主要精力在 openclaw 稳定性与发布准备，这通常意味着**接下来会有一次正式版本发布**。UI 类功能请求目前处于"有实现但未排期"状态，建议维护者明确告知社区这些功能的去向，避免用户重复提交。

## 7. 用户反馈摘要

从今日关闭的 Issue 及对应 PR 描述中可提炼出以下真实用户痛点：

- **信息密度与可读性**：侧边栏固定 240px 导致小屏挤压、大屏浪费，长标题被截断后无法判断内容（[#1314](https://github.com/netease-youdao/LobsterAI/issues/1314)）。表格长文本截断后无全文入口（[#1311](https://github.com/netease-youdao/LobsterAI/issues/1311)）。
- **启动体验焦虑**：会话加载期间显示"暂无历史记录"空状态，用户会误以为数据丢失（[#1319](https://github.com/netease-youdao/LobsterAI/issues/1319)）。
- **快捷键发现成本高**：功能存在但界面无提示，新用户必须进入设置页才能发现（[#1317](https://github.com/netease-youdao/LobsterAI/issues/1317)）。
- **操作反馈缺失**：复制站点 URL/分享码无成功提示（今日已通过 [#2417](https://github.com/netease-youdao/LobsterAI/pull/2417) 修复）。
- **长会话性能敏感**：DeepSeek 缓存命中率骤降直接影响长会话响应速度（来自 [#2415](https://github.com/netease-youdao/LobsterAI/pull/2415) 的开发者报告）。

## 8. 待处理积压

- **[PR #2234 fix(openclaw): cron yield descendant finalization](https://github.com/netease-youdao/LobsterAI/pull/2234)**（**OPEN**）— 当前唯一处于打开状态的 PR，修复 cron 场景下子 agent 完成事件无法驱动父 agent 继续执行的问题，自 2026-06-30 创建至今无人 review，已被标记 stale。**建议维护者优先处理**，该修复涉及三个测试场景，是明确的功能缺陷修复。

- **以下 PR 被 stale 关闭但从未合并**（关闭 ≠ 合并，功能可能已丢失）：
  - [PR #172 feat(oauth): add Antigravity OAuth integration](https://github.com/netease-youdao/LobsterAI/pull/172) — 大块功能：OAuth 子系统 + SQLite 持久化 + 代理兼容，2 月提交 4 个月未审。
  - [PR #1308 feat(cowork): isolate home-screen input draft per agent](https://github.com/netease-youdao/LobsterAI/pull/1308) — 按 agent 隔离首页输入草稿。
  - [PR #1315 拖拽调整侧边栏宽度](https://github.com/netease-youdao/LobsterAI/pull/1315) / [PR #1318 快捷键 kbd 提示](https://github.com/netease-youdao/LobsterAI/pull/1318) / [PR #1320 会话列表骨架屏](https://github.com/netease-youdao/LobsterAI/pull/1320) — 三个 UI 功能实现，对应 Issue #1314 / #1317 / #1319。
  - [PR #1321 fix(settings): dismiss overlays when switching settings tabs](https://github.com/netease-youdao/LobsterAI/pull/1321) — 设置页 overlay 残留 Bug 的修复。

  **建议**：维护者应对上述 PR 做一次"重开 / 合入 / 正式关闭"的人工判定，避免功能请求在社区层面"已实现"、在代码层面"不存在"的割裂状态。

---

**项目健康度总评**：核心运行时稳定性修复效率高、方向正确；但社区功能请求响应链路存在明显断点（实现 4 个月未合并被自动清理），**建议补充"贡献者 PR 评审时效"指标**，避免挫伤社区贡献积极性。整体处于**发布前收敛期**，健康度中等偏上。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目动态日报 — 2026-08-01

## 1. 今日速览

过去24小时内，Moltis 项目保持中高活跃度：共新增/更新 2 条 Issue、7 条 PR，其中 1 条功能相关 PR 已合并关闭，另有 6 条处于待合并状态。值得关注的是，社区提交了 2 个安全加固方向的 PR（节点配对签名验证、模型/zip 路径硬化），并有 1 个新 Bug 报告（#1181）被提交。项目整体处于功能迭代与安全加固并行推进的阶段，活跃度良好，维护者响应及时。


## 2. 版本发布

过去24小时内无新版本发布。


## 3. 项目进展

**已合并/关闭：**

- [#1176 [CLOSED] feat(web): add Markdown copy and session export](https://github.com/moltis-org/moltis/pull/1176) — 该 PR 正式落地了两个关键用户功能：① 在复制助手回复时保留原始 Markdown 格式（含实时/持久化回复，不含模型元数据）；② 新增会话级“另存为 Markdown”操作，可加载完整分页历史并导出用户/助手文本及图片引用。此功能与已关闭的 Issue #1131 直接对应，是用户声量较高的功能请求之一。

**待合并（推动中的主要方向）：**

- [#1179 [OPEN] fix(gateway): verify node pairing signatures](https://github.com/moltis-org/moltis/pull/1179) — 安全修复：将 `node.pair.verify` 绑定到服务器签发的待处理请求，防止调用方自行提供密钥或挑战值，补齐了配对签名验证的安全缺口。
- [#1180 [OPEN] fix(security): harden model and zip paths](https://github.com/moltis-org/moltis/pull/1180) — 安全修复：修复恶意 zip 包或 HuggingFace 仓库导致任意文件写入/代码执行的两类漏洞，涉及 `clawhub.rs` 的 zip 提取路径校验加固。
- [#1168 [OPEN] feat(nostr): add NIP-29 group chat support for Buzz channels](https://github.com/moltis-org/moltis/pull/1168) — 功能扩展：为 Nostr 通道新增 NIP-29 群聊支持，使 Moltis 可以接入 Block 公司的开源工作区产品 Buzz（Nostr relay + NIP-42 认证）。
- [#1158 [OPEN] feat(memory): add zvec vector database memory backend](https://github.com/moltis-org/moltis/pull/1158) — 功能扩展：新增基于 Zvec + redb 的向量数据库记忆后端，通过 `zvec` cargo feature 门控，默认在 `full` 特性中启用。
- [#1174 [OPEN] Add instrumentation and feedback collection infrastructure](https://github.com/moltis-org/moltis/pull/1174) — 基础设施：添加与后端无关的 Agent 插桩、Langfuse v4 导出、OTLP 运维后端，以及终端用户反馈（reaction）收集能力。
- [#1170 [OPEN] fix(channels): gate /sh and privileged tools behind a per-account operators list](https://github.com/moltis-org/moltis/pull/1170) — 安全修复：将访问控制与特权分离，引入 per-account `operators` 列表，避免通过访问白名单的频道用户越权执行特权命令和主机工具。

**整体判断：** 已有 1 个用户可见功能（Markdown 导出）正式合并；另有 3 个安全加固 PR 和 3 个功能/基础设施 PR 正在等待评审合并，项目当前处于“功能交付 + 安全补强”双线推进的活跃阶段。


## 4. 社区热点

- [#1176 feat(web): add Markdown copy and session export](https://github.com/moltis-org/moltis/pull/1176) 是今日唯一合并的 PR，对应社区请求 #1131（Markdown 复制/导出），表明该功能从发起到交付链路完整，社区推动力强。
- [#1180 fix(security): harden model and zip paths](https://github.com/moltis-org/moltis/pull/1180) 与 [#1179 fix(gateway): verify node pairing signatures](https://github.com/moltis-org/moltis/pull/1179) 均来自同一外部贡献者（tsauvajon），提交信息显示其希望先解决安全问题再正式采用 Moltis——“I'd like to use Moltis, but I've got a couple of security fixes I'd like to get in before doing so”，反映出安全加固是潜在用户采纳的关键前提。

**用户诉求背后：** 社区对“数据导出能力”和“安全边界”的需求正在上升，特别是多渠道访问权限、模型路径可信性、配对认证完整性等信任基础设施。


## 5. Bug 与稳定性

**新增 Bug：**

- [#1181 [OPEN] [bug] Issue with GPT 5.6 Luna](https://github.com/moltis-org/moltis/issues/1181) — 用户报告与 GPT 5.6 Luna 相关的问题。严重程度暂未标注，且尚无评论。这是今日唯一新开 Bug，需要维护者跟进复现并评估是否与模型接入层或网关相关。

**安全保障类 PR（未合并，但已提交修复方向）：**

- 高危：[#1180 fix(security): harden model and zip paths](https://github.com/moltis-org/moltis/pull/1180) — 恶意 zip/HuggingFace repo 可覆盖用户信任的文件（配置、凭据、脚本）并导致代码执行，属于任意文件写入漏洞，严重程度高。
- 中高危：[#1179 fix(gateway): verify node pairing signatures](https://github.com/moltis-org/moltis/pull/1179) — 节点配对签名验证缺失，调用方可自行提供密钥/挑战值，可能导致中间人或伪造配对，需尽快评审合并。
- 中危：[#1170 fix(channels): gate /sh and privileged tools behind a per-account operators list](https://github.com/moltis-org/moltis/pull/1170) — 频道发送者可越权执行特权命令，涉及权限边界问题。

综合来看，项目当前存在 3 个已识别/已提交修复方案的安全风险，均处于待合并状态；新建 Bug #1181 目前缺乏足够上下文。


## 6. 功能请求与路线图信号

- **Markdown 复制/导出（#1131 → #1176）**：已合并落地，对应的用户请求已关闭，属于快速兑现。
- **Nostr 群聊 / Buzz 集成（#1168）**：NIP-29 群聊支持指向与 Block 旗下 Buzz 工作区的集成，说明 Moltis 生态正在向多 agent 协作平台扩展，可能成为下一版本的核心能力之一。
- **Zvec 向量记忆后端（#1158）**：新记忆后端的加入意味着 Moltis 在长短期记忆与 embedding 方案上有更多选择空间。
- **可观测性与反馈收集（#1174）**：插桩 + Langfuse 导出 + 用户 reaction 反馈，预示项目路线图将更重视可观测性和基于用户反馈的持续优化。
- **安全与权限模型重塑**：来自外部贡献者的 2 个 Security fix PR（#1179、#1180）及权限隔离 PR（#1170），可能推动下一版本在 Trust & Safety 方面有较大改进。


## 7. 用户反馈摘要

- **安全顾虑阻碍采纳**：贡献者 tsauvajon 在 #1179 中明确表示“想用 Moltis，但需要先合入几个安全修复”，这表明安全问题是部分潜在用户的核心采纳障碍，尤其涉及节点配对和文件路径处理方面。
- **导出功能需求明确**：#1131 关闭标志着“Markdown 复制/导出”这一高👍功能请求正式完成（👍: 1），用户在 #1176 的实现中获得了“保留原始 Markdown + 完整会话导出”的体验。
- **Bug 报告质量意识增强**：#1181 中用户已按照 Preflight Checklist 完成搜索历史和最新版本确认，说明社区使用规范度良好。


## 8. 待处理积压

以下 PR 已存在较长时间且仍未合并，建议维护者优先评估：

- [#1158 feat(memory): add zvec vector database memory backend](https://github.com/moltis-org/moltis/pull/1158) — 创建于 2026-07-17，已持续 15 天，当前仍有更新（07-31），需要评审反馈。
- [#1170 fix(channels): gate /sh and privileged tools behind a per-account operators list](https://github.com/moltis-org/moltis/pull/1170) — 创建于 07-26，涉及安全权限边界，建议尽快处理。
- [#1174 Add instrumentation and feedback collection infrastructure](https://github.com/moltis-org/moltis/pull/1174) — 创建于 07-27，属于较大的基础设施变更，需评估其对现有架构的影响。

另外，安全修复 PR #1179 和 #1180 虽创建时间较短（07-31），但严重程度较高，建议与 #1170 一并纳入快速评审通道，防止漏洞在正式版本中长期暴露。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报（2026-08-01）

## 今日速览

- 过去 24 小时 Issues 更新 18 条（新开/活跃 13，关闭 5），PR 更新 43 条（待合并 30，合并/关闭 13），无新版本发布。
- 社区提交活跃，但 PR 合并率仅约 30%（13/43），待合并 PR 积压明显，维护者 review 压力较大。
- 今日多个严重 Bug 已获得针对性修复 PR（UI 冻结、shell 命令挂起、agent.json 损坏等），修复响应速度较快。
- 关注点集中在稳定性：长命令输出/超时、记忆丢失、会话数据完整性、第三方平台（微信/飞书）集成失效。
- 整体活跃度高，项目处于高频迭代期，但需要加快 PR 合并节奏并跟踪静默失败类问题。

## 版本发布

今日无新版本发布。

## 项目进展

今日关闭/合并的 PR 中，以下工作推进了稳定性、记忆、音频和 UI 方向：

- **#6602 fix(issue 6558) session integrity**：修复多会话 UI 数据完整性问题，保留 Coding/Chat 模式切换时的 in-flight 流，并统一 per-agent TaskTracker 管理。[PR #6602](https://github.com/agentscope-ai/QwenPaw/pull/6602)
- **#6573 fix(audio): restore transcription for channel audio messages**：修复飞书等频道音频消息在 AgentScope 2.0 迁移后静默转写失败的问题。[PR #6573](https://github.com/agentscope-ai/QwenPaw/pull/6573)
- **#6592 fix(memory): flush Auto-Memory before Scroll context eviction**：修复 Scroll 上下文压缩导致早期会话未被 Auto-Memory 写入、最终缺失于每日记忆的问题，同时避免手动 `/compact` 重复提交任务。[PR #6592](https://github.com/agentscope-ai/QwenPaw/pull/6592)
- **#6606 fix(read_file): accept numeric string line ranges**：修复 `read_file` 工具对数字字符串行范围的解析兼容性。[PR #6606](https://github.com/agentscope-ai/QwenPaw/pull/6606)
- **#6604 docs(memory): explain ReMe self-evolving knowledge base**：完善 ReMe 记忆机制文档，明确捕获、索引、整合、召回生命周期。[PR #6604](https://github.com/agentscope-ai/QwenPaw/pull/6604)

整体上，项目在记忆可靠性、渠道消息处理和 Web UI 数据一致性方面有明显向前推进。

## 社区热点

- **#6537 Skill tags disappear on restart (regression of #3270)**：评论 10 条，为今日最活跃 Issue。用户报告 Skill Pool UI 中设置的标签在重启后丢失，API 已成功保存到 `skill_pool/skill.json`，但启动或构建时 manifest 协调会覆盖。背后反映的是配置持久化与内部协调逻辑冲突，用户对“保存成功但重启丢失”的回归问题非常敏感。[Issue #6537](https://github.com/agentscope-ai/QwenPaw/issues/6537)
- **#6601 QwenPaw 不报空响应错误**：评论 5 条。用户指出长会话逼近上下文窗口上限后，模型空响应但 QwenPaw 不报错，导致会话彻底失去响应。这是框架层缺失，直接影响长会话可用性。[Issue #6601](https://github.com/agentscope-ai/QwenPaw/issues/6601)
- **#6588 `spawn_subagent` single-task mode is unusable**：评论 4 条。模型工具 schema 将可选参数 `batch` 暴露为必填，导致前台单任务子代理无法创建。已有 PR #6609 修复。[Issue #6588](https://github.com/agentscope-ai/QwenPaw/issues/6588)

## Bug 与稳定性

按严重程度排列，标注是否已有修复 PR：

1. **严重 — 微信 cron 定时推送静默失败**（#6614）：任务显示 `success` 但微信侧 `ret=-2`，从未真正送达，已消耗约 44M tokens。暂无 fix PR。[Issue #6614](https://github.com/agentscope-ai/QwenPaw/issues/6614)
2. **严重 — 长命令绕过超时并阻塞飞书会话**（#6608）：一个脚本阻塞示例 1.5 小时，后续消息全部排队，取消后遗留孤儿进程，且无 per-channel 总超时。已有 PR #6610 修复 shell 命令挂起。[Issue #6608](https://github.com/agentscope-ai/QwenPaw/issues/6608)
3. **严重 — `execute_shell_command` 大量输出导致 UI 冻结**（#6589）：数万行 stdout 一次性渲染阻塞主线程。已有 PR #6610 覆盖。[Issue #6589](https://github.com/agentscope-ai/QwenPaw/issues/6589)
4. **严重 — QwenPaw 2.0.1 与 agentscope 2.0.4.post1 不兼容**（#6612）：主动/记忆子系统出现 `Msg.content` 类型崩溃和工具权限死锁。已有 PR #6615 修复。[Issue #6612](https://github.com/agentscope-ai/QwenPaw/issues/6612)
5. **严重 — agent.json 系统性损坏**（#6520）：BOM 头、缺失引号、双重编码导致系统完全故障。已有 PR #6528 修复安全读取与写入。[Issue #6520](https://github.com/agentscope-ai/QwenPaw/issues/6520)
6. **中等 — Skill tags 重启消失**（#6537）：回归 #3270，标签在启动协调时丢失。暂无 fix PR。[Issue #6537](https://github.com/agentscope-ai/QwenPaw/issues/6537)
7. **中等 — 空响应错误未提示**（#6601）：长会话静默失去响应。暂无 fix PR。[Issue #6601](https://github.com/agentscope-ai/QwenPaw/issues/6601)
8. **中等 — `spawn_subagent` 单任务模式不可用**（#6588）：schema 必填 `batch`。已有 PR #6609 修复。[Issue #6588](https://github.com/agentscope-ai/QwenPaw/issues/6588)
9. **中等 — Dream/memory 压缩丢失早间事件**（#6555）：上下文压缩后当日记忆永久缺失。已有 PR #6592 合并、#6564 待合并。[Issue #6555](https://github.com/agentscope-ai/QwenPaw/issues/6555)
10. **中等 — ACP `new_session` 响应缺 `models` 字段**（#6529）：客户端无法发现可用模型。暂无 fix PR。[Issue #6529](https://github.com/agentscope-ai/QwenPaw/issues/6529)
11. **中等 — 飞书音频消息转写失败**（#6544）：已由 PR #6573 修复并关闭。[Issue #6544](https://github.com/agentscope-ai/QwenPaw/issues/6544)
12. **中等 — 多会话 UI 数据完整性问题**（#6558）：消息丢失、指令漂移、回复重渲染。已由 PR #6602 修复并关闭。[Issue #6558](https://github.com/agentscope-ai/QwenPaw/issues/6558)
13. **低 — `execute_shell_command` 大输出截断**（#6512）：>30KB 输出被截断或触发 `Internal error`。与 #6589 相关，PR #6610 已缓解超时/冻结，截断机制尚需后续方案。[Issue #6512](https://github.com/agentscope-ai/QwenPaw/issues/6512)
14. **低 — Desktop 输入框被遮挡**（#6549）：Windows 高缩放下 UI 布局问题。暂无 fix PR。[Issue #6549](https://github.com/agentscope-ai/QwenPaw/issues/6549)

## 功能请求与路线图信号

- **工作区产出物快捷访问**（#6083）：希望 Desktop 窗口内一键直达工作区文件夹或下载最近产出物，方便非技术用户。[Issue #6083](https://github.com/agentscope-ai/QwenPaw/issues/6083)
- **独立 Python 运行环境**（#6160）：避免依赖系统全局 Python，内置或复用后端解释器，提升开箱即用体验。[Issue #6160](https://github.com/agentscope-ai/QwenPaw/issues/6160)
- **结果呈现优化**（#6260）：希望折叠思考/工具调用过程，突出最终交付结果，减少过程信息淹没。[Issue #6260](https://github.com/agentscope-ai/QwenPaw/issues/6260)
- **大输出治理**（#6512）：自动写入文件或提供流式读取机制，避免截断和 UI 卡死。[Issue #6512](https://github.com/agentscope-ai/QwenPaw/issues/6512)
- **桌面应用名简化**（#6587）：将 “QwenPaw Desktop” 改为 “QwenPaw”。[Issue #6587](https://github.com/agentscope-ai/QwenPaw/issues/6587)
- **新功能 PR 信号**：全局热键快速输入窗（#6607）、NVIDIA NIM Provider 支持（#6526）、统一 provider 发现/模型元数据/路由/Agent 控制（#6302）。这些 PR 表明路线图正扩展到多 provider 生态和桌面端交互效率。

## 用户反馈摘要

- **配置持久化困惑**：用户 Ra-M497 表示 Skill tags 通过 API 保存成功，但重启后丢失，怀疑是 manifest 协调逻辑覆盖所致（#6537）。
- **长会话稳定性担忧**：用户 rerbin 反馈上下文接近上限时模型空响应，QwenPaw 不报错，会话“彻底失去响应”（#6601）。
- **静默失败成本高**：用户 angelozb 报告微信定时推送从未送达，任务长期显示 success，累计消耗约 44M tokens，严重影响信任（#6614）。
- **阻塞体验差**：用户 feng183043996 描述一个脚本阻塞飞书会话 1.5 小时，后续消息全部排队，取消后仍有孤儿进程（#6608）。
- **非技术用户痛点**：AL-Mint 强调访问工作区文件需要手动打开资源管理器并导航到隐藏目录，流程中断工作流（#6083）。
- **结果呈现优先级**：azear 认为 Agent 思考过程占满屏幕，交付结果被淹没，更希望直接看到最终产出（#6260）。
- **环境适配问题**：xiaobing006 使用 Conda 管理 Python 环境，系统无全局 Python，导致生成的脚本无法执行，希望内置运行时（#6160）。
- **UI 细节**：sinomind 反馈输入框被遮挡，发送按钮需要滚动才能看到，影响日常使用（#6549）。

## 待处理积压

- **#6083 Desktop 工作区产出物快捷访问按钮**（创建 07-14，已 18 天，4 评论）——长期 open 的增强需求，未见对应 PR。[Issue #6083](https://github.com/agentscope-ai/QwenPaw/issues/6083)
- **#6160 独立 Python 运行环境**（创建 07-16，已 16 天，4 评论）——重要易用性需求，仍无明确计划。[Issue #6160](https://github.com/agentscope-ai/QwenPaw/issues/6160)
- **#6260 结果呈现优化**（创建 07-19，已 13 天，2 评论，有 👍1）——UI 体验改进，缺少维护者回复。[Issue #6260](https://github.com/agentscope-ai/QwenPaw/issues/6260)
- **PR #6203 fix(utils): bound and hide the Windows tasklist liveness probe**（创建 07-16，Under Review 16 天）——Windows 稳定性修复，等待 review。[PR #6203](https://github.com/agentscope-ai/QwenPaw/pull/6203)
- **PR #6302 feat: unify provider discovery, model metadata, routing, and agent controls**（创建 07-21，已 11 天）——大型架构性 PR，涉及核心模块，长期未合并可能阻塞其他依赖。[PR #6302](https://github.com/agentscope-ai/QwenPaw/pull/6302)
- **PR #6528 fix: resolve agent.json corruption (BOM etc.)**（创建 07-28，已 4 天）——针对严重问题的修复，应优先 review 并合并。[PR #6528](https://github.com/agentscope-ai/QwenPaw/pull/6528)

---
**项目健康度总结**：社区活跃度高，Bug 修复响应快，但 PR 合并率偏低、待合并积压 30 条，需要维护者加快 review 节奏。多个“静默失败”类问题（微信推送、空响应、标签丢失）对用户体验伤害最大，应作为下一版本优先修复项。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw 项目动态日报 — 2026-08-01

> 数据窗口：过去 24 小时（2026-07-31 ～ 2026-08-01）｜数据来源：GitHub Issues / PR 元数据

---

## 1. 今日速览

ZeroClaw 项目过去 24 小时保持**高度活跃**：PR 侧更新 50 条，其中 12 条已合并/关闭，38 条仍在审查中，代码合入节奏良好；Issue 侧更新 27 条，其中 2 条关闭，并有 10+ 个来自维护者拆分的 follow-up 工作项集中创建（#9591–#9602），说明项目正在系统性地推进技术债清理与架构收敛。安全与稳定性方向的投入显著：Landlock 运行时问题已合入修复（#9114），Dependabot 安全更新链路正在被逐层打通（#9539 → #9601），Wasm 插件体系相关 RFC 持续升温。项目整体健康度良好：Bug 响应快、RFC 讨论有跟踪机制、贡献者分层管理（principal / distinguished / trusted contributor）清晰可见；但 OpenAI 兼容端点、LSP 支持等关键社区诉求仍停留在 PR/RFC 阶段，等待维护者推进决策。

---

## 2. 版本发布

过去 24 小时无新版本发布（最新 Releases 为空），本节省略。

---

## 3. 项目进展

过去 24 小时共有 12 条 PR 被合并或关闭，重点包括：

- **Landlock 沙箱运行时修复已合入**：PR [#9114](https://github.com/zeroclaw-labs/zeroclaw/pull/9114) 允许访问各类设备文件与系统文件，直接修复 Fedora 上 shell 工具不可用的 P1 问题（对应 Issue [#8973](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)）。这是本周最值得关注的安全/兼容性修复合入。
- **agent 历史拆分重构完成（PR-A）**：PR [#9525](https://github.com/zeroclaw-labs/zeroclaw/pull/9525) 将 `provider_messages` 拆分为 `loop_history` 与 `loop_new_messages`，为后续 `before_llm_call` 钩子（PR-B）铺路，属于行为等价重构，风险控制得当。
- **容器构建缓存竞争修复**：PR [#9369](https://github.com/zeroclaw-labs/zeroclaw/pull/9369) 对共享 Cargo 缓存挂载启用 `sharing=locked`，避免多阶段容器构建的缓存损坏。
- **Web 侧边栏交互修复**：PR [#9294](https://github.com/zeroclaw-labs/zeroclaw/pull/9294) 让侧边栏选中状态保持互斥，修复 `/config/agents` 同时选中 Agent 与 Config 父级的问题。
- **robot-kit 子进程挂死修复**：PR [#9087](https://github.com/zeroclaw-labs/zeroclaw/pull/9087) 为 TTS/音频播放子进程加入等待超时，避免在树莓派等硬件环境卡死。
- **测试基础设施隔离**：PR [#9367](https://github.com/zeroclaw-labs/zeroclaw/pull/9367) 为 safety-net agent 构造路径提供独立临时工作区，消除 `/tmp` 共享冲突。
- **SBOM 工具链升级**：PR [#9344](https://github.com/zeroclaw-labs/zeroclaw/pull/9344) 将 `anchore/sbom-action` 升级至 v0.24.0 全 SHA 固定。

**仍在推进中的重量级 PR**（尚未合并，值得关注）：OpenAI Chat Completions 端点（[#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)）、live eval 沙箱执行模式（[#9214](https://github.com/zeroclaw-labs/zeroclaw/pull/9214)）、OTel 跨轮会话关联（[#9352](https://github.com/zeroclaw-labs/zeroclaw/pull/9352)）、Anthropic 存储型 OAuth 支持（[#9420](https://github.com/zeroclaw-labs/zeroclaw/pull/9420)）。这些 PR 一旦合入，将显著提升项目在互操作性、可观测性与认证能力上的完整度。

---

## 4. 社区热点

**讨论最集中的 Issue 与深层诉求：**

- **[#8550 — OpenAI-compatible chat completions endpoint](https://github.com/zeroclaw-labs/zeroclaw/issues/8550)**（6 条评论，P2，risk:high，status:accepted）：社区最强的互操作性诉求。当前 ZeroClaw 只暴露 WebSocket 和渠道专用协议，Open WebUI、LobeChat 等标准客户端无法直接接入。该 Issue 已有对应 PR [#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)，目前处于 `needs-author-action` 状态，需要作者跟进。
- **[#5907 — RFC: Opt-in LSP support](https://github.com/zeroclaw-labs/zeroclaw/issues/5907)**（5 条评论，创建于 2026-04-19，已悬置 3.5 个月）：开发者希望借助 Language Server 降低模型幻觉、提升本地模型生成代码质量。Claude Code 与 OpenCode 都已支持 LSP，社区对其成为 ZeroClaw 编码工作流标配的呼声较高。
- **[#8692 — Maintainer decision queue tracker](https://github.com/zeroclaw-labs/zeroclaw/issues/8692)**（5 条评论，P2）：社区与维护者都意识到 RFC/设计决策积压的问题。该 tracker 的存在本身就是治理优化信号。
- **[#8135 — Wasm-first plugin runtime RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/8135)**（4 条评论，risk:high，P2）：围绕"用 Wasm 取代 Node.js 插件运行时"的架构级讨论，与 [#8187](https://github.com/zeroclaw-labs/zeroclaw/issues/8187) 的 WASI 硬件能力门控形成配套，是项目远期架构的核心方向之一。

**分析**：社区讨论热度集中在三件事——**标准协议互操作**（OpenAI 兼容、LSP）、**插件运行时安全与沙箱化**（Wasm-first）、**决策流程透明化**（maintainer decision queue）。三者都指向项目从"可用"走向"好接、好扩展、好治理"的成长期诉求。

---

## 5. Bug 与稳定性

按严重程度排列（S1 > S2 > S3）：

**S1 / P1 — 工作流阻塞**

- **[#9601 — Dependabot 不产生 transitive Cargo alerts 的 PR](https://github.com/zeroclaw-labs/zeroclaw/issues/9601)**（新开，P1，status:accepted，risk:medium）：仓库级 Dependabot 安全更新已启用，直接 npm 依赖能产生 PR，但可解析的传递性 Cargo 告警无法产生。**尚无 fix PR**，属于安全自动化链路缺口。
- **[#9591 — 频道 reload 移除全部频道后 delivery registry 残留](https://github.com/zeroclaw-labs/zeroclaw/issues/9591)**（新开，P1，status:accepted，follow-up）：`CRON_CHANNEL_REGISTRY` 在频道集合为空时跳过发布，导致旧注册表继续存活。**尚无 fix PR**。

**S2 / P1-P2 — 功能降级**

- **[#8973 — Landlock 阻断 shell 访问 /dev/null（Fedora）](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)**（P1，已关闭 ✅）：用户报告沙箱启用时 `sh` 无法访问 `/dev/null`，shell 工具完全不可用。修复 PR [#9114](https://github.com/zeroclaw-labs/zeroclaw/pull/9114) 已于今日关闭，**已解决**。
- **[#9592 — model-routing 更新后 probe 使用旧配置快照](https://github.com/zeroclaw-labs/zeroclaw/issues/9592)**（P1，status:accepted，follow-up）：保存新 provider alias 后 probe 仍读取更新前运行时快照，导致探测结果不准确。**尚无 fix PR**。
- **[#9593 — 后台委托生命周期状态双写](https://github.com/zeroclaw-labs/zeroclaw/issues/9593)**（P1，type:refactor，status:accepted）：`TaskRecord.status` 与 `BackgroundDelegateResult.status` 双轨维护，存在一致性问题。**尚无 fix PR**。
- **[#9594 — Coding-agent 工具对 action 预算双重计费](https://github.com/zeroclaw-labs/zeroclaw/issues/9594)**（P2，status:accepted）：单次工具调用记录两次 `ToolOperation::Act`。**尚无 fix PR**。
- **[#9546 — updater web-dist 测试依赖宿主安装状态](https://github.com/zeroclaw-labs/zeroclaw/issues/9546)**（P2，type:test）：在开发机上因二进制邻接路径假设不成立而失败。**尚无 fix PR**。

**P3 — 低严重度**

- **[#9590 — models refresh 并发丢缓存条目](https://github.com/zeroclaw-labs/zeroclaw/issues/9590)**（P3，status:accepted）：读-改-写未加跨进程锁，两个并发 `zeroclaw models refresh` 会静默丢失更新。**尚无 fix PR**。

**稳定性观察**：今日新开的 Bug 类 Issue 大多被维护者标记为 `status:accepted` 并附带 `follow-up`，说明这些是从既有 review 中拆出的已知技术债，而非用户突发的未知问题；整体 Bug 治理流程是健康、可控的。

---

## 6. 功能请求与路线图信号

**最可能进入下一版本的功能：**

- **OpenAI Chat Completions 兼容端点**（[#8550](https://github.com/zeroclaw-labs/zeroclaw/issues/8550)）：`status:accepted`，PR [#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486) 已存在，只待作者响应 review 意见。这是当前**离落地最近**的重量级功能。
- **AI-assisted PR pre-review / re-review**（[#9330](https://github.com/zeroclaw-labs/zeroclaw/issues/9330)）：基于现有 CI 结果触发 AI 初审，人类保留最终审批权，符合项目 risk-based 治理风格。
- **PR 风险/尺寸标签自动重算**（[#9345](https://github.com/zeroclaw-labs/zeroclaw/issues/9345)）：`status:accepted`，降低人工维护成本。

**架构级路线图信号（RFC 阶段，尚无对应 PR）：**

- **Wasm-first 插件运行时**（[#8135](https://github.com/zeroclaw-labs/zeroclaw/issues/8135)）+ **capability-gated WASI 硬件函数**（[#8187](https://github.com/zeroclaw-labs/zeroclaw/issues/8187)）：安全沙箱化插件生态是项目避开 Node.js 依赖、实现 signed distribution 的长期方向。
- **统一 package/capability/config/runtime-state catalog 契约**（[#9346](https://github.com/zeroclaw-labs/zeroclaw/issues/9346)）：解决集成、内置项、插件三套清单割裂的问题。
- **SOP capability 权限契约**（[#9598](https://github.com/zeroclaw-labs/zeroclaw/issues/9598)）：在增加更多 side-effecting 能力适配器之前，先定义真实的 `required_permissions` 授权契约。
- **会话持久化契约所有权 tracker**（[#9600](https://github.com/zeroclaw-labs/zeroclaw/issues/9600)）：四个独立工作流同时修改同一契约，需要指定 owner 与合并顺序。
- **执行树迭代预算所有权**（[#9323](https://github.com/zeroclaw-labs/zeroclaw/issues/9323)）：所有生产根节点目前都未设置 `ToolLoop.shared_budget`，需要明确父/子 agent 迭代预算归属。
- **响应缓存策略**（[#8321](https://github.com/zeroclaw-labs/zeroclaw/issues/8321)）：当前日期/时间、记忆召回、频道上下文等易变内容是否可缓存需要统一边界。

---

## 7. 用户反馈摘要

- **Fedora 用户的真实环境闪崩**（[#8973](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)）：Landlock 沙箱导致 `sh` 无法访问 `/dev/null`，shell 工具完全失效。用户对"开启沙箱后基本功能不可用"的体验不满；好消息是修复已合入。这说明安全功能在真实 Linux 发行版上需要更多回归测试。
- **macOS 开发者的测试困惑**（[#9546](https://github.com/zeroclaw-labs/zeroclaw/issues/9546)）：`web-dist` 测试在开发机上依赖二进制邻接路径假设，安装了非预期版本的开发者会遇到误导性失败。反馈指向测试可移植性问题。
- **并发 CLI 使用的可靠性担忧**（[#9590](https://github.com/zeroclaw-labs/zeroclaw/issues/9590)）：用户并行执行 `models refresh` 时缓存条目静默丢失——命令行工具在并发场景下的数据一致性需要加强。
- **Signal/Voice Call 频道的崩溃循环**（PR [#9524](https://github.com/zeroclaw-labs/zeroclaw/pull/9524)）：启用了缺少凭据的 Signal 频道后，监听器无法连接且 supervisor 无限重启，说明频道配置错误时的降级行为应更友好。
- **WebSocket 掉线导致 agent 任务被取消**（PR [#9002](https://github.com/zeroclaw-labs/zeroclaw/pull/9002)）：用户浏览器睡眠或切页后，进行中的 agent 回合被错误取消；这是桌面端与移动端真实使用中的高频痛点，修复方案（将 viewer 与控制权分离）已提交。

---

## 8. 待处理积压

以下重要事项长期未决，需要维护者关注：

- **[#5907 — LSP 支持 RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/5907)**：创建于 2026-04-19，已积压 **3.5 个月**，5 条评论，状态仍未接受/拒绝/延期。这是社区呼声较高的编码工作流增强，建议尽快给出 maintainer 决策。
- **[#8550 / #8486 — OpenAI 兼容端点](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)**：Issue 已 `accepted`，但 PR 处于 `needs-author-action`，作者（REL-mame）需回应 review 意见。若长期无响应，维护者应考虑是否接管或关闭。
- **[#8135 — Wasm-first 插件运行时 RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/8135)**：创建于 2026-06-22，已讨论 1 个月+，仍无范围结论。该 RFC 与 #8187 联动性较强，建议安排一次架构例会集中讨论。
- **[#8078 — zerocode 本地预提交门 RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/8078)**：创建于 2026-06-21，零代码工具链方向的前置投入，需要确认是否纳入近期路线图。
- **多个 `needs-author-action` PR 待跟进**：[#9214](https://github.com/zeroclaw-labs/zeroclaw/pull/9214)（live eval）、[#9002](https://github.com/zeroclaw-labs/zeroclaw/pull/9002)（viewer disconnect）、[#9420](https://github.com/zeroclaw-labs/zeroclaw/pull/9420)（Anthropic OAuth）、[#9313](https://github.com/zeroclaw-labs/zeroclaw/pull/9313)（WeChat cursor）、[#9320](https://github.com/zeroclaw-labs/zeroclaw/pull/9320)（cron 超时）、[#9428](https://github.com/zeroclaw-labs/zeroclaw/pull/9428)（Bluesky/Reddit 权限）、[#9352](https://github.com/zeroclaw-labs/zeroclaw/pull/9352)（OTel 关联）。这些 PR 均为 P1/P2 且有实质改动，建议维护者统一催办或接手。

---

*本日报由 AI 分析师自动生成，基于 ZeroClaw 公开 GitHub 数据。部分 PR 评论数据缺失（undefined），如需完整评论内容请以 GitHub 页面为准。*

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*