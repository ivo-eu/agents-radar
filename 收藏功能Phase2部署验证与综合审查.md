# Radar 收藏功能 Phase 2 部署验证与综合审查

审查及部署日期：2026-06-11

关联文档：

- `收藏功能Phase2技术实现.md`
- `收藏功能Phase2交接清单.md`
- `收藏功能Phase1实现审查.md`

## 1. 结论

Phase 2 的 Bridge 部署与后端端到端链路已经实际跑通：

```text
浏览器同步载荷
→ 本机 Radar Bridge
→ Bridge 专用 agents-radar clone
→ commit + push favorites.json
→ GitHub Pages 发布 favorites.json
→ SecondBrain/收藏/AI情报收藏.md
```

已验证新增与软删两条路径均返回 `pushed: true`。

但是 Phase 2 目前还不能标记为完整上线：

1. 正式 GitHub Pages 上尚未部署当前工作区的 Phase 2 前端代码。
2. 第 10 节发现的 Vault 自愈和 `--repo-url` 两项问题已经修复，但第三轮复审又发现 1 个新的 P1：日报同步可能覆盖收藏同步刚写入的持久状态，详见第 11 节。

建议实现 Agent 先修复第 11 节问题，再提交并发布 Phase 2 前端与 Bridge 源码。

## 2. 实际部署结果

### 2.1 环境

```text
Go：go1.24.3 darwin/amd64
Vault：/Users/dinsafer/Desktop/ai学习区/SecondBrain
Pages：https://ivo-eu.github.io/agents-radar
Bridge listen：127.0.0.1:4399
专用 clone：
/Users/dinsafer/Library/Application Support/RadarBridge/agents-radar
```

GitHub SSH key 未在 GitHub 账号注册，因此没有使用 SSH remote。

最终使用：

```text
https://github.com/ivo-eu/agents-radar.git
```

虽然 `gh auth status` 中旧 token 已失效，但 macOS Keychain 中的 Git HTTPS 凭证仍然有效，专用 clone 的 `git push --dry-run` 验证通过。

### 2.2 Bridge 安装

新版二进制部署到稳定路径：

```text
/Users/dinsafer/Library/Application Support/RadarBridge/bin/radar-bridge
```

配置文件：

```text
~/Library/Application Support/RadarBridge/config.json
```

配置已包含：

```json
{
  "vault_path": "/Users/dinsafer/Desktop/ai学习区/SecondBrain",
  "pages_url": "https://ivo-eu.github.io/agents-radar",
  "listen": "127.0.0.1:4399",
  "repo_path": "/Users/dinsafer/Library/Application Support/RadarBridge/agents-radar"
}
```

LaunchAgent：

```text
com.ivo.radar-bridge
状态：running
监听：127.0.0.1:4399
```

### 2.3 Git 凭证验证

在 Bridge 专用 clone 中执行：

```bash
git push origin HEAD:master --dry-run
```

结果：

```text
Everything up-to-date
```

说明 HTTPS 凭证助手可供普通 Git 命令和 launchd 下的 Bridge 使用。

## 3. 后端端到端验证

### 3.1 新增收藏

向：

```text
POST http://127.0.0.1:4399/api/favorites/sync
```

发送一条 64 位 hex ID 的链接收藏。

返回：

```json
{
  "alive": 1,
  "pushed": true,
  "total": 1
}
```

GitHub 提交：

```text
bdba45f favorites: sync (1 alive)
```

验证结果：

- 专用 clone 根目录生成 `favorites.json`。
- GitHub 远端 `master` 收到提交。
- GitHub Pages 的 `/favorites.json` 返回 HTTP 200。
- Pages 内容与专用 clone 中的 JSON 一致。
- SecondBrain 生成 `收藏/AI情报收藏.md`。
- Markdown 包含网站、来源、回到 Radar 和原始链接。

### 3.2 软删清理

为避免正式收藏库长期显示测试数据，又发送同 ID 墓碑：

```json
{
  "deleted_at": "2026-06-11T08:05:00Z"
}
```

返回：

```json
{
  "alive": 0,
  "pushed": true,
  "total": 1
}
```

GitHub 提交：

```text
3f208b3 favorites: sync (0 alive)
```

最终结果：

- `favorites.json` 保留一条墓碑，用于跨设备防复活。
- `AI情报收藏.md` 不再展示测试项，只保留文件头。
- Bridge 状态：

```json
{
  "connected": true,
  "favorites_enabled": true,
  "favorites_error": "",
  "syncing": false
}
```

## 4. 测试结果

### 4.1 前端和项目测试

```bash
pnpm typecheck
```

结果：通过。

```bash
pnpm test
```

结果：

```text
13 个测试文件通过
219 个测试通过
```

其中新增的 `src/__tests__/favorites-id.test.ts` 有 16 个测试，覆盖 JS 与 Go 的 URL、文本和 SHA-256 ID 契约。

### 4.2 Bridge 测试

```bash
go test -race ./...
go vet ./...
go build -o radar-bridge ./cmd/radar-bridge
go test -cover ./...
```

结果全部通过。

Bridge 语句覆盖率：

```text
30.1%
```

当前测试已覆盖：

- 收藏 ID 跨语言一致性。
- 非法 ID、kind、date、report 过滤。
- 合并字段与排序。
- 墓碑和单次复活。
- Markdown 分组及墓碑过滤。
- 临时 Git 仓库中的 commit + push。
- 非 favorites 文件脏工作区拒绝。
- fetch 失败时本地 JSON 和 Markdown 写入。

## 5. 代码审查问题

### 5.1 P1：同步进行中新增的收藏可能被响应覆盖并永久丢失

位置：

- `index.html:612-645`

当前流程：

1. `flushFavSync()` 对当前 `favoriteStore.items` 发起请求。
2. 请求等待期间，用户仍可以新增或取消收藏。
3. 响应返回后，代码直接执行：

```js
favoriteStore.items = res.items;
```

响应只包含请求发出时的旧快照，因此等待期间产生的新本地修改会被覆盖。

随后代码又将：

```js
favSyncedAt = new Date().toISOString();
```

推进到响应时间。即使先合并保住新项目，新项目的 `created_at` 也可能早于这个水位，导致它被误判为已经同步。

影响：

- 用户快速连续收藏时可能静默丢数据。
- 页面不会继续显示“待同步”。
- GitHub 和 SecondBrain 永远收不到被覆盖的项目。

建议：

1. 请求开始时保存 payload 快照和该快照的最大活动时间 `ackWatermark`。
2. 响应回来后，将 `res.items` 与“当前内存中请求后新增的变化”重新合并，而不是覆盖。
3. `favSyncedAt` 只能推进到请求快照的 `ackWatermark`，不能使用响应时刻。
4. 增加可控 Promise 的前端测试，模拟请求在途时新增、删除和复活收藏。

### 5.2 P1：复活事件没有持久化，旧墓碑可在下一次跨设备合并时再次杀死收藏

位置：

- `index.html:555-568`
- `bridge/cmd/radar-bridge/main.go:1041-1061`

规则同时要求：

- `created_at` 永远保留最早时间，用于稳定排序。
- 用最新 `created_at` 判断删除后是否发生重新收藏。

一次合并时可以临时判断复活，但输出又把 `created_at` 写回最早时间，导致“最后一次重新收藏发生在删除之后”这个事件时间丢失。

示例：

```text
原始收藏 created_at = 08:00
删除 deleted_at = 09:00
重新收藏 = 10:00
```

第一次合并会清空墓碑，但落盘后只剩：

```text
created_at = 08:00
deleted_at = 空
```

之后与另一个设备保存的 09:00 墓碑再次合并时，系统只能看到 08:00 的 add 和 09:00 的 delete，收藏会再次被删除。

建议：

- 新增独立事件字段，例如 `updated_at`、`added_at` 或 `revived_at`。
- `created_at` 仅作为首次收藏和排序时间。
- 生死判定使用 `max(added_at/revived_at)` 与 `max(deleted_at)`。
- 增加“三段式合并”测试：add → delete → revive → 再与旧 delete 合并，最终必须保持 alive。

### 5.3 P1：`--repo` 可对任意干净工作区执行 `git reset --hard`

位置：

- `bridge/cmd/radar-bridge/main.go:338-383`
- `bridge/cmd/radar-bridge/main.go:881-909`

CLI 支持：

```bash
radar-bridge install --repo <现有仓库>
```

但同步时无条件执行：

```bash
git reset --hard origin/master
```

脏检查只能保护未提交文件，不能保护：

- 当前位于其他干净分支的用户。
- 本地领先但尚未 push 的提交。
- 分支指针和工作树状态。

这与技术文档中的“绝不 reset --hard 用户工作区”不一致。

本次部署没有触发风险，因为实际使用的是 Bridge 专用 clone。

建议：

- 最安全方案：删除 `--repo`，只允许 `--repo-url` 创建 Bridge 拥有的专用 clone。
- 或在专用 clone 内写 ownership marker，并仅对有 marker 的目录允许 hard reset。
- 若保留 `--repo`，不得执行 hard reset；应使用独立 worktree、临时 clone 或显式 fetch/show 读取远端基线。
- 增加“干净 feature branch”和“本地领先提交”的保护测试。

### 5.4 P1：损坏的 `favorites.json` 会被静默当成空库

位置：

- `bridge/cmd/radar-bridge/main.go:1108-1133`

`parseFavorites()` 在 JSON 解析失败时返回 `nil`，没有返回错误。`readFavoritesItems()` 也会吞掉读取错误。

因此远端 `favorites.json` 一旦损坏或格式异常，下一次同步可能：

1. 把原收藏库当成空数组。
2. 只写当前浏览器传来的项目。
3. commit + push 覆盖原文件。

建议：

- `parseFavorites` 返回 `([]FavoriteItem, error)`。
- 除“文件不存在”外，任何读取和解析错误都必须中止同步。
- API 返回明确错误并保持前端 pending。
- 增加 malformed JSON 和权限错误测试。

### 5.5 P1：Obsidian Markdown 写入失败仍可能返回 `pushed: true`

位置：

- `bridge/cmd/radar-bridge/main.go:899`
- `bridge/cmd/radar-bridge/main.go:913`
- `bridge/cmd/radar-bridge/main.go:1144-1147`

同步使用：

```go
a.writeFavoritesMarkdownBestEffort(merged)
```

该函数只记录日志，不把错误返回给调用方。因此：

- GitHub push 成功。
- Vault 写入失败。
- API 仍返回 `pushed: true`。
- 前端清除待同步状态。
- `favorites_error` 仍为空。

这违反 Phase 2“一次同步两份产物一起写”的已锁定决策。

建议：

- 将 Markdown 写入错误纳入 API 结果。
- 如果允许 GitHub 成功但 Vault 失败，至少返回独立状态：

```json
{
  "pushed": true,
  "vault_written": false,
  "error": "..."
}
```

- 前端必须保留醒目错误提示。
- 增加只读 vault、目录冲突等写入失败测试。

### 5.6 P2：交接清单按原样执行会让 LaunchAgent 指向仓库里的构建产物

位置：

- `bridge/cmd/radar-bridge/main.go:318-323`
- `收藏功能Phase2交接清单.md` 的编译和安装步骤

`install` 使用 `os.Executable()` 直接写入 LaunchAgent。

按交接清单执行：

```bash
cd bridge
go build -o radar-bridge ...
./radar-bridge install ...
```

LaunchAgent 将永久指向：

```text
<工作区>/bridge/radar-bridge
```

以后清理仓库、删除构建产物或移动目录都会使服务启动失败。

本次部署已人工规避：先将二进制复制到 Application Support 的稳定路径，再从该路径执行 install。

建议：

- `install` 自行复制当前二进制到：

```text
~/Library/Application Support/RadarBridge/bin/radar-bridge
```

- LaunchAgent 始终指向稳定副本。
- 使用原子替换升级二进制。
- 更新交接清单，避免依赖人工复制。

### 5.7 P2：Markdown 日期按 UTC 分组，不是技术文档要求的本地日期

位置：

- `bridge/cmd/radar-bridge/main.go:1207-1215`

当前：

```go
return t.UTC().Format("2006-01-02")
```

技术文档写的是按 `created_at` 当地日期分组。上海时区凌晨前后的收藏可能被放到前一天。

建议明确产品语义：

- 若按用户本地日期，使用 `time.Local` 或配置时区。
- 若决定统一 UTC，则同步修改技术文档和 UI 说明。

## 6. 发布阻塞项

当前正式远端 `master` 上没有本工作区的 Phase 2 前端和 Bridge 源码。

已确认远端 `index.html` 不包含：

```text
/api/favorites/sync
FAV_SYNCED_KEY
blockAnchorId
```

目前远端只有 Bridge 实际生成的 `favorites.json` 提交。也就是说：

- 后端 Bridge 真实同步已可用。
- GitHub Pages 已能读取公开 JSON。
- 正式网站还不会执行 Phase 2 自动同步客户端。
- 正式 Pages 的配对、5 秒防抖、待同步提示、取消收藏跨端同步尚未做真实浏览器验收。

在修复第 5 节问题并提交代码后，仍需执行：

1. 等待 Pages 发布新版 `index.html`。
2. 正式 Pages 重新配对本机 Bridge。
3. 连续快速收藏，专门验证同步在途并发场景。
4. 取消收藏并刷新、换浏览器确认不复活。
5. 模拟 Bridge 离线后恢复。
6. 验证 GitHub 与 Obsidian 两份结果一致。

## 7. 建议修复顺序

1. 修复前端在途同步覆盖和错误水位。
2. 修改收藏事件模型，持久化 revive/add 事件时间。
3. 修复损坏 JSON 时的 fail-closed 行为。
4. 让 Vault 写入结果成为同步成功判定的一部分。
5. 收紧 `--repo` 和 hard reset 安全边界。
6. 让 install 自动部署稳定二进制。
7. 明确 Markdown 日期时区。
8. 补自动化测试。
9. 提交并发布 Phase 2 代码。
10. 完成正式 Pages 前端端到端验收。

## 8. 当前可保留的部署状态

以下状态已经正确，可以保留：

- Go 已安装。
- Bridge 稳定二进制已部署。
- LaunchAgent 正常运行。
- Bridge 使用专用 clone。
- HTTPS Git push 凭证有效。
- `config.json` 已启用收藏同步。
- GitHub Pages 已能发布 `favorites.json`。
- SecondBrain 统一收藏 Markdown 已创建。
- 测试数据已软删，不会出现在收藏列表或 Markdown 中。

不要删除专用 clone：

```text
~/Library/Application Support/RadarBridge/agents-radar
```

不要将主开发工作区配置为 Bridge 的 `repo_path`。

## 9. 修复落实记录（2026-06-11）

实现 Agent 已针对第 5 节的 5 个 P1 + 2 个 P2 提交修复，改动集中在 `index.html`、`bridge/cmd/radar-bridge/main.go` 与两端测试；修复后的独立复审结论见第 10 节。

### P1-5.1 在途同步覆盖（`index.html`）

- `flushFavSync` 开始时对 `favoriteStore.items` 做**快照**并发送快照（不再发实时引用）；计算快照内最大活动时间 `ackWatermark`。
- 响应回来用新增的 `reconcileWithServer()`：以服务端 merged 集为权威基底，再 fold 当前 store（含在途新增/删除），用 `mergeItem` 按 id 合并——**不再 `favoriteStore.items = res.items` 直接覆盖**。
- `favSyncedAt` 只 `favLatest(favSyncedAt, ackWatermark)`，**绝不**推进到响应时刻；在途变化活动时间 > 水位 → 仍判 pending，finally 里自动再 flush。
- 新增 `favLastActivity()`（max(created/revived/deleted)）供水位与 pending 判定共用。

### P1-5.2 复活事件持久化（前后端合并模型）

- `FavoriteItem` 新增 `revived_at`：`created_at` 仅作首次收藏 + 排序键；`revived_at` 记最近一次复活。
- 生死判定改为 `aliveTime = max(created_at, revived_at)` 对比 `delTime = max(deleted_at)`（Go `mergeTwo` 与 JS `mergeItem` 同步改）。
- `addFavorite` 复活墓碑时：保留原 `created_at`、`revived_at = now`、清 `deleted_at`。
- 验收：add 08:00 → delete 09:00 → revive 10:00，落盘 `created_at=08:00 / revived_at=10:00`；再与另一设备 09:00 旧墓碑合并仍 **alive**（`TestMergeRevivePersistsAgainstStaleTombstone` + 前端 `favorites-merge.test.ts`）。

### P1-5.4 损坏 favorites.json fail-closed（`main.go`）

- `parseFavorites` 改签名返回 `([]FavoriteItem, error)`：空文件=空库，**畸形 JSON=硬错误**。
- `readFavoritesItems`/`readFavoritesFile` 透传错误（仅"文件不存在"当空库）。
- `syncFavorites` 读基线失败时**中止同步**，不写不推，避免用空集覆盖损坏的远端库（`TestSyncFavoritesFailsClosedOnCorruptBaseline`）。

### P1-5.5 Vault 写入纳入同步结果（`main.go` + `index.html`）

- `syncFavorites` 返回新结构 `favSyncResult{items, pushed, vaultWritten, error}`；API 响应增加 `vault_written`。
- markdown 写入失败不再 best-effort 吞掉：记 `vaultWritten=false` 返回（允许 GitHub 成功但 Vault 失败的部分成功）。
- 前端 `flushFavSync` 在 `pushed && vault_written===false` 时设醒目 `favSyncError`「SecondBrain 写入失败」。
- 删除无用的 `writeFavoritesMarkdownBestEffort`（`TestSyncFavoritesReportsVaultWriteFailure`）。

### P1-5.3 收紧 --repo / hard reset（`main.go`）

- 新增 ownership marker `.git/radar-bridge-owned`：`--repo-url` 创建/复用的专用 clone 会打标。
- 抽出 `gitAlignToOrigin`：**有标** → `reset --hard origin/master`（安全）；**无标**（用户 `--repo` 自有仓库）→ 拒绝 reset，仅在 master 且无领先提交时 `merge --ff-only`，否则报错中止保护用户工作区。
- 验收：未拥有仓库在 feature branch 或有未推送提交时拒绝同步且不丢提交（`TestSyncRefusesHardResetOnUnownedRepo`）。

### P2-5.6 install 部署稳定二进制（`main.go`）

- 新增 `deployStableBinary()`：原子复制当前二进制到 `~/Library/Application Support/RadarBridge/bin/radar-bridge`，LaunchAgent 指向该稳定副本（不再指向运行时路径/仓库构建产物）。已从稳定路径运行时为幂等 no-op（`TestDeployStableBinary`）。

### P2-5.7 Markdown 本地日期分组（`main.go`）

- `favDay` 由 `t.UTC()` 改为 `t.Local()`，对齐技术文档「`created_at` 当地日期」（`TestFavDayUsesLocalDate`，Asia/Shanghai 跨午夜用例）。

### 验证结果

```text
pnpm typecheck        通过
pnpm test             14 文件 / 226 测试通过（新增前端 7：favorites-merge.test.ts）
pnpm lint             通过
go vet ./...          通过
gofmt -l              干净
go build              成功
go test ./...         22 测试通过（新增 7）
go test -race ./...   通过
```

> 仍待真实浏览器/凭证端到端（与代码无关）：发布新 `index.html` 到 Pages、重新配对、连续快速收藏验证在途并发、取消刷新换浏览器验证不复活、Bridge 离线恢复、GitHub 与 Obsidian 两份一致。见第 6 节发布步骤。

## 10. 修复后复审与真实环境验收（2026-06-11）

### 10.1 复审结论

第 9 节记录的主要修复已经生效：

- 在途请求返回时不再覆盖请求期间产生的本地收藏变化。
- 同步水位只推进到请求快照的 `ackWatermark`。
- `revived_at` 可以保护删除后重新收藏的项目不被旧墓碑再次删除。
- 损坏的 `favorites.json` 会中止同步，不再被当作空库。
- 用户自有 `--repo` 不会被 `reset --hard`。
- install 会将 Bridge 部署到稳定路径。
- Markdown 按本地日期分组。

但本轮复审又发现 1 个 P1 和 1 个 P2，当前仍建议修复后再正式发布。

### 10.2 P1：Vault 写入失败后前端不会真正自动重试

位置：

- `index.html:635-677`
- `bridge/cmd/radar-bridge/main.go:904-923`

当前流程：

1. Bridge 已成功 push GitHub，但写入 SecondBrain Markdown 失败。
2. API 返回：

```json
{
  "pushed": true,
  "vault_written": false
}
```

3. 前端在检查 `vault_written` 之前，已经因为 `res.pushed` 执行：

```js
favSyncedAt = favLatest(favSyncedAt, ackWatermark);
localStorage.setItem(FAV_SYNCED_KEY, favSyncedAt);
```

4. 随后虽然设置了 `favSyncError = "SecondBrain 写入失败"`，并安排 30 秒重试，但下一次进入 `flushFavSync(false)` 时：

```js
if (!showResult && !pendingFavCount()) return;
```

由于同步水位已经推进，`pendingFavCount()` 为 0，重试直接返回，Vault 不会再次写入。

另外：

- `favSyncError` 只存在内存中，刷新页面后警告消失。
- Bridge 的 Markdown 写入错误被转换为 `vaultWritten=false, err=nil`。
- `handleFavoritesSync()` 因 `result.err == nil` 会清空 `FavoritesError`。
- 用户最终可能看到 GitHub 已同步、Bridge 状态无错误，但 SecondBrain 长期缺少该收藏。

影响：

- 违反“一次同步同时产出 GitHub JSON 和 SecondBrain Markdown”的产品约定。
- 部分成功状态无法自愈。
- 页面刷新后没有可见告警。

建议修复：

1. 仅在 `res.pushed && res.vault_written !== false` 时推进 `favSyncedAt`。
2. Vault 失败时保留 pending；30 秒后重新发送全量快照。
3. Bridge 每次同步即使 GitHub 无新 diff，也会先重写 Markdown，因此重试可以修复 Vault。
4. Bridge 应保留实际 Vault 错误文本，例如增加 `vault_error`，或在 `pushed=true` 时仍返回非空 `error`。
5. `FavoritesError` 在 `vault_written=false` 时不得被清空。
6. 手动点击 SYNC 时不得弹出“收藏已同步”，应提示“GitHub 已同步，但 SecondBrain 写入失败”。

必须补充的测试：

- `pushed=true + vault_written=false` 时同步水位不推进。
- Vault 失败后仍有 pending。
- 定时重试会再次调用 `/api/favorites/sync`。
- 第二次 `vault_written=true` 后才清除 pending 和错误。
- 页面刷新或重新读取 Bridge status 后仍能看到未完成状态。

### 10.3 P2：已有专用 clone 时新的 `--repo-url` 会被静默忽略

位置：

- `bridge/cmd/radar-bridge/main.go:363-376`

当前 `resolveRepoPath()` 在目标目录已有 `.git` 时只执行：

```text
assertRepoUsable(dest)
git fetch origin
markRepoOwned(dest)
return dest
```

它没有比较当前 `origin` 与本次 install 传入的 `repoURL`。

因此，如果用户执行：

```bash
radar-bridge install --repo-url <新仓库地址>
```

但 Application Support 中已经存在旧专用 clone，Bridge 仍会继续 fetch 和 push 旧 origin，同时安装命令看起来成功。

影响：

- 修改仓库地址、账号或 fork 后可能继续向错误仓库推送。
- 用户难以从 install 输出发现实际 origin 与参数不一致。

建议修复：

- 对 Bridge-owned clone：读取 `git remote get-url origin`，与传入值不一致时明确更新 `origin`，或直接报错要求用户确认。
- 对无法确认 ownership 的目录不得自动修改 remote。
- install 输出实际生效的 origin URL。

必须补充的测试：

- 已有 Bridge-owned clone且 origin 与 `repoURL` 相同：正常复用。
- 已有 Bridge-owned clone但 origin 不同：更新或明确拒绝，不能静默忽略。
- 非 Bridge-owned clone：不得自动改 remote。

### 10.4 自动化验证

本轮实际执行：

```text
pnpm typecheck              通过
pnpm test -- --run          14 文件 / 226 测试通过
go test -race ./...         通过
go vet ./...                通过
go build                    通过
git diff --check            通过
```

补充说明：

- `go test -cover ./...` 报 `cannot find package` 已查明原因：**在仓库根目录 `agents-radar/` 运行**（该目录无 `go.mod`）。Go module 在 `bridge/` 子目录，必须 `cd bridge` 后再跑。非代码回归。
- 复审后已复现确认：`cd bridge && go test -cover ./...` 正常输出，覆盖率随新增测试升至 **40.9%**（原 30.1%）。

### 10.5 浏览器功能验收

本地浏览器已验证：

1. 整篇报告星标收藏。
2. 左侧“我的收藏”计数立即更新。
3. 收藏列表按类型显示。
4. 点击“回到报告”能回跳到原报告。
5. 回跳目标出现闪烁高亮。
6. 未连接 Bridge 时显示“1 条待同步”。
7. 浏览器控制台无 error/warning。

### 10.6 真实 Bridge 端到端验收

最新版 Bridge 已重新构建并安装到：

```text
~/Library/Application Support/RadarBridge/bin/radar-bridge
```

LaunchAgent 正常运行，并使用专用 clone。

通过本机同源预览页面实际执行：

```text
网页收藏整篇报告
→ 顶部显示 1 条待同步
→ 约 5 秒后自动同步
→ GitHub commit + push favorites.json
→ SecondBrain/收藏/AI情报收藏.md 出现收藏
→ 网页取消收藏
→ 再次自动同步
→ favorites.json 写入墓碑
→ Markdown 移除收藏
```

真实提交：

```text
1d414a3 favorites: sync (1 alive)
0fc1343 favorites: sync (0 alive)
```

最终状态：

- 专用 clone `HEAD == origin/master`。
- `favorites_error` 为空。
- `favorites.json` 中测试数据均为墓碑。
- SecondBrain Markdown 当前无测试收藏。

### 10.7 当前人工验收入口

已启动仅监听本机的同源预览代理：

```text
http://127.0.0.1:8766/#2026-06-10/ai-cli
```

该页面：

- 使用当前工作区的 `index.html`。
- 已连接真实 `127.0.0.1:4399` Bridge。
- 收藏操作会真实 push GitHub，并真实写入 SecondBrain。
- 不修改正式 Pages 配置。
- 仅监听本机，不对局域网开放。

注意：正式 GitHub Pages 仍未部署当前 Phase 2 前端源码；本入口用于提交前人工验收。

### 10.8 给实现 Agent 的本轮修复清单

1. 修复 Vault 部分成功后的水位、pending、重试和持久错误状态。
2. 补充该状态机的前端和 Bridge 测试。
3. 修复已有专用 clone 时忽略新 `--repo-url`。
4. 补充 origin 相同、origin 不同、非 owned clone 三类测试。
5. 重新运行 typecheck、前端测试、Go race/vet/build。
6. 复现并确认 `go test -cover ./...` 异常。
7. 修复后再执行一次网页收藏与取消收藏真实端到端。

### 10.9 修复落实记录（2026-06-11，第二轮）

实现 Agent 已处理第 10 节的 1 P1 + 1 P2 + cover 疑点，改动集中在 `index.html`、`bridge/cmd/radar-bridge/main.go` 与两端测试；第三轮独立复审结论见第 11 节。

#### P1-10.2 Vault 部分成功的状态机（`index.html` + `main.go`）

- **Bridge**：`favSyncResult` 增 `vaultErr`；markdown 写入失败时不再吞错，带出错误文本。`handleFavoritesSync` 在 `pushed && !vaultWritten` 时**持久化** `FavoritesError`（不清空），响应增加 `vault_error` 字段。每次同步都会先重写 markdown（在 git diff 之前），因此重试能真正修复 Vault。
- **前端**：`flushFavSync` 区分 `vaultFailed = pushed && vault_written === false`——此时**不推进 `favSyncedAt`**（水位不动 → `pendingFavCount()` 保持非零 → finally 的 30s 重试真正重发），`favSyncError` 设为 `vault_error`；SYNC 手动提示区分「GitHub 已同步，但 SecondBrain 写入失败」。`refreshBridgeStatus` 采纳 Bridge 持久化的 `favorites_error`，刷新页面后告警仍可见。
- 测试：Go `TestHandleFavoritesSyncPersistsVaultError`（响应含 vault_error + state 保留错误）、`TestSyncFavoritesReportsVaultWriteFailure`（扩展断言 `vaultErr` 非空）；前端 `favorites-sync.test.ts` 用 scope-proxy 沙箱驱动**真实 `flushFavSync`**：验证 vault 失败不推水位/保留 pending、重试后二次成功才清 pending、干净同步正常清 pending。

#### P2-10.3 origin 不一致不再静默忽略（`main.go`）

- `resolveRepoPath` 复用已有 clone 时先 `git remote get-url origin` 与传入 `--repo-url` 比对（新增 `sameGitURL` 忽略尾部 `.git`/`/`/空白）：相同→复用；不同且为 Bridge-owned→更新 origin 并打印；不同且非 owned→**报错拒绝**，绝不擅改用户 remote。`install` 末尾打印实际生效的 push origin。
- 测试：`TestSameGitURL`、`TestResolveRepoPathOriginMismatch`（owned 更新 / unowned 拒绝且 origin 不变 / 同 URL 干净复用三类）。

#### cover 疑点（§10.4）

- 复现确认：`cannot find package` 是因为在仓库根（无 `go.mod`）跑；`cd bridge` 后 `go test -cover ./...` 正常，覆盖率 **40.9%**。非代码回归。

#### 本轮验证结果

```text
pnpm typecheck            通过
pnpm test                 15 文件 / 229 测试通过（新增前端 favorites-sync.test.ts 3 个）
pnpm lint                 通过
cd bridge && gofmt -l     干净
cd bridge && go vet       通过
cd bridge && go build     成功
cd bridge && go test -race 通过
cd bridge && go test -cover 覆盖率 40.9%
```

> 仍待（与代码无关）：发布当前 `index.html` 到正式 Pages、重新配对、连续快速收藏验证在途并发、取消刷新换浏览器验证不复活、Bridge 离线恢复、GitHub 与 Obsidian 两份一致；以及在真实只读 vault 场景手动验证 Vault 失败 → 修复后自愈的完整链路。

## 11. 第三轮复审（2026-06-11）

### 11.1 结论

第 10 节的两项修复本身通过复审：

- Vault 写入失败时不再推进前端同步水位，pending 会保留并触发重试。
- Bridge 会返回 `vault_error` 并持久化 `FavoritesError`。
- Vault 恢复后，同一全量同步会重新写 Markdown，再清除错误和 pending。
- 已有专用 clone 的 origin 与新 `--repo-url` 不一致时，不再静默忽略。
- Bridge-owned clone 可以更新 origin；非 owned clone 会拒绝修改。

但是第三轮复审发现 1 个新的 P1，仍不建议直接发布。

### 11.2 P1：日报同步可能覆盖收藏同步刚写入的状态

位置：

- `bridge/cmd/radar-bridge/main.go:636-660`
- `bridge/cmd/radar-bridge/main.go:932-952`
- `index.html:846-849`
- `index.html:1519`

Bridge 有两条可以并发执行的同步链路：

```text
日报同步：syncMu → sync()
收藏同步：favMu  → handleFavoritesSync() / syncFavorites()
```

它们使用不同的互斥锁，但共同读写同一个：

```go
a.state
state.json
```

日报同步开始时会从磁盘读取完整状态，然后整份替换内存状态：

```go
var diskState State
if err := readJSON(a.statePath, &diskState); err == nil {
    a.mu.Lock()
    a.state = diskState
    // ...
    a.mu.Unlock()
}
```

可达的并发顺序：

1. 日报同步从磁盘读到旧 `state.json`，其中 `FavoritesError` 为空。
2. 收藏同步发生 Vault 写入失败。
3. 收藏同步将 `FavoritesError = "SecondBrain 写入失败：..."` 写入内存和磁盘。
4. 日报同步随后执行 `a.state = diskState`，用步骤 1 的旧快照覆盖内存。
5. 日报同步结束时 `saveState()`，又把旧收藏状态写回磁盘。

可能被覆盖的字段包括：

- `FavoritesError`
- `FavoritesLastSync`
- 收藏链路未来新增的其他状态字段

为什么真实可达：

- 页面初始化调用 `refreshBridgeStatus(true)`。
- 该函数会启动日报 `syncBridge(false)`。
- 如果本地收藏仍 pending，同时还会调用 `flushFavSync(false)`。
- 因此页面启动时两种同步可能同时运行。

影响：

- Vault 失败告警本应刷新后继续显示，但可能被日报同步静默清除。
- `favorites_error` 可能错误显示为空。
- 第 10.2 节修复的“持久错误状态”在并发场景下不可靠。
- 问题不一定被 Go race detector 发现，因为访问都经过 `a.mu`；这是逻辑层面的陈旧快照覆盖。

### 11.3 建议修复

推荐不要在日报同步中整份替换 `a.state`。

方案 A（推荐）：

- 从磁盘只合并日报链路拥有的字段：

```text
Token
Files
LastSync
LastError
```

- 保留当前内存中的：

```text
Favorites
FavoritesLastSync
FavoritesError
```

- 更理想的是明确拆分状态所有权，日报同步与收藏同步各自只修改自己的字段。

方案 B：

- 为 `state.json` 的 read-modify-write 建立一把共享事务锁。
- 日报同步和收藏同步在读取磁盘、合并、写回的完整过程都持有该锁。
- 不能只靠当前 `a.mu`，因为当前代码在读取磁盘和最终写回之间会释放锁并执行较长的网络/文件操作。

不建议：

- 仅让 `syncMu` 与 `favMu` 互相嵌套。
- 仅调整页面调用顺序。
- 仅依赖 race detector。

这些方式容易隐藏问题，但没有建立清晰的状态所有权。

### 11.4 必须补充的测试

增加一个可控并发测试，至少覆盖：

1. 准备磁盘旧状态：`FavoritesError=""`。
2. 让日报同步读完旧状态后暂停。
3. 执行收藏同步，写入 Vault 错误和 `FavoritesLastSync`。
4. 恢复日报同步并让其完成。
5. 断言内存与 `state.json` 中收藏字段仍是新值。

还应覆盖反向恢复路径：

1. 收藏同步成功清除旧 `FavoritesError`。
2. 并发日报同步不得用旧磁盘错误将其恢复。

测试目标不是仅通过 `go test -race`，而是验证最终状态字段不会发生陈旧快照覆盖。

### 11.5 本轮验证结果

```text
pnpm typecheck              通过
pnpm lint                   通过
pnpm test -- --run          15 文件 / 229 测试通过
cd bridge && gofmt -l       干净
cd bridge && go vet ./...   通过
cd bridge && go build       通过
cd bridge && go test -race  通过
cd bridge && go test -cover 覆盖率 40.9%
git diff --check            通过
```

最新版 Bridge 已重新安装并执行真实回归：

```text
新增收藏
→ pushed=true
→ vault_written=true
→ GitHub commit
→ SecondBrain Markdown 出现记录
→ 软删除
→ pushed=true
→ Markdown 移除记录
```

真实回归提交：

```text
b58a1c4 favorites: sync (1 alive)
be47ceb favorites: sync (0 alive)
```

最终：

- 专用 clone `HEAD == origin/master`。
- 测试收藏已软删除。
- SecondBrain Markdown 无测试收藏。
- Bridge `favorites_error` 为空。

### 11.6 GitHub 凭证状态

本轮同时验证了两条 GitHub 鉴权链路：

```text
gh auth status
→ ivo-eu 登录有效
→ scopes: gist, read:org, repo, workflow

gh api user
→ 成功返回 ivo-eu

Bridge 专用 clone:
git push --dry-run origin HEAD:master
→ Everything up-to-date
```

结论：

- 当前 `gh` CLI token 没有失效。
- Bridge 实际使用的 HTTPS + macOS `osxkeychain` Git 凭证也有效。
- 两者是独立检查结果，当前都能正常访问 GitHub。

### 11.7 给实现 Agent 的最新修复清单

1. 修复日报同步整份覆盖 `a.state` 的陈旧快照问题。
2. 明确日报字段与收藏字段的状态所有权。
3. 增加可控并发测试，验证 Vault 错误不会被清除。
4. 增加反向测试，验证已清除的旧错误不会被并发同步恢复。
5. 重新运行前端测试、Go race/vet/build/cover。
6. 修复后再做一次 Vault 失败 → 恢复 → 刷新页面的完整验收。
