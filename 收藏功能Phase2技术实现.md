# Radar 收藏功能 Phase 2 技术实现文档

> 范围：把 Phase 1 的本地收藏，**无感**同步到 ① GitHub 仓库 `favorites.json`（网站读）和
> ② SecondBrain vault 的 `收藏/AI情报收藏.md`（Obsidian 读）—— 一次同步两份产物一起写。
> 实现位置：`index.html`（同步客户端）+ `bridge/cmd/radar-bridge/main.go`（新端点 + git + md 渲染）。
> 状态：**已实现**（Bridge + 前端均完成；`go test`、`pnpm test` 全部通过）。实现纪要见文末 §8。

---

## 0. 已锁定决策

| 决策项 | 结论 |
|--------|------|
| 收藏 ID | **信任浏览器算好的 id**，Bridge 只校验 64 位 hex，不重算 |
| 同步触发 | **防抖自动同步（无感）**：收藏后 5s、连接时、加载时、切后台/关闭前；SYNC 按钮仅兜底 |
| 取消收藏 | **软删墓碑 `deleted_at`**，跨设备生效，不会复活 |
| push | **自动 commit+push**，基于 `origin/master` 合并，只提交 favorites.json 路径，best-effort |
| Obsidian md | **合进本阶段一起做**（即把原 Phase 3 并入），同一次同步写 json + md |
| 安全 | 专用 clone；工作区有非 favorites.json 改动则拒绝；`127.0.0.1` + pairing 不变 |

---

## 1. 数据契约（两端唯一真相）

### 1.1 FavoriteItem（在 Phase 1 基础上加墓碑与 anchor）

```jsonc
{
  "id": "<64-hex>",                 // 浏览器 SHA-256，Bridge 信任
  "kind": "report | link | excerpt",
  "title": "...",
  "url": "https://...",             // 仅 link
  "site": "github.com",             // 仅 link，hostname
  "excerpt": "...",                 // excerpt=段落全文；link=上下文；report=""
  "section": "...",
  "created_at": "2026-06-11T07:42:00.000Z",  // 首次收藏时间，仅作排序键
  "revived_at": "2026-06-11T09:00:00.000Z",  // 可选，最近一次删后复活时间
  "deleted_at": "2026-06-11T08:00:00.000Z",  // 可选，软删墓碑
  "source": {
    "date": "2026-06-10",
    "report": "ai-trending",
    "label": "GitHub AI 趋势",
    "url": "https://ivo-eu.github.io/agents-radar/#2026-06-10/ai-trending",
    "anchor": "<id>"                // 回跳锚点 = data-favorite-id
  }
}
```

### 1.2 favorites.json（仓库根）

```jsonc
{ "version": 1, "updated_at": "ISO", "items": [ /* FavoriteItem[]，含墓碑 */ ] }
```

> 墓碑也写进 json，跨设备才能“删了不复活”。网站与 md 渲染时跳过墓碑。

### 1.3 同步载荷（前端 → Bridge）

`POST /api/favorites/sync`：

```jsonc
{ "items": [ /* 前端 localStorage 全部条目，含墓碑 */ ] }
```

> **发全量**而非增量：幂等、无状态、localStorage 体积小。Bridge 合并后回传权威结果，前端覆盖本地。

### 1.4 合并规则（Bridge，Go；前端 loadFavorites 同规则）

按 id 做并集，同 id 合并：

- `created_at` = 两者中**最早**的非空值（排序主键稳定，**只**用于排序）。
- `revived_at` = 两者中**最晚**的非空值（最近一次删后复活事件时间）。
- 补字段：任一侧有、另一侧缺的字段填上（title/url/site/excerpt/section/source）。
- 墓碑：`delTime = max(deleted_at)`，`aliveTime = max(created_at, revived_at)`
  - 若 `aliveTime > delTime` → **alive**（清空 `deleted_at`，支持删除后再收藏复活）；
  - 否则 `deleted_at = delTime`（墓碑）。
  - > 用 `revived_at` 而非 `created_at` 判生死：`created_at` 合并后写回最早值会丢失复活事件时间，导致旧墓碑跨设备再合并时复活项被二次杀死（审查 §5.2）。`revived_at` 持久化该事件，旧墓碑无法再杀活复活项。
- 输出按 `created_at` 倒序。

---

## 2. 前端改造（`index.html`）

### 2.1 软删墓碑（改 Phase 1 行为）

- `removeFavorite(id)`：不再从数组删，改成打 `deleted_at = now`。
- `favoriteIds`：只收 **alive**（无 `deleted_at`）的 id —— 高亮、`toggle`、计数均基于 alive。
- 重新收藏已墓碑的 id：`created_at = now`，清 `deleted_at`。
- `renderFavoritesView` / `renderFavNav`：过滤墓碑。
- `loadFavorites`：合并远端时套用 §1.4（补墓碑逻辑）。

### 2.2 同步客户端（新增）

```js
let favSyncTimer = null;            // 防抖
let favSyncing   = false;
let favSyncError = '';
const FAV_SYNC_DEBOUNCE = 5000;

scheduleFavSync()   // 收藏/取消后：清旧 timer，5s 后 flushFavSync()
flushFavSync(force) // Bridge 在线 → POST /api/favorites/sync（全量）
                    //   成功：用返回 merged items 覆盖 store + persist + 更新 syncedAt
                    //   失败：记 favSyncError，保持 pending，顶部提示
hasPendingFavSync() // 是否有未确认同步的条目
```

**pending 判定**：store 加 `syncedAt`（上次成功同步时间戳）。任一 item 的 `created_at` 或 `deleted_at` > `syncedAt` 即有待同步。比维护独立队列简单且幂等。

**触发点（无感）**：

| 时机 | 行为 |
|------|------|
| add / remove 收藏 | `scheduleFavSync()`（防抖 5s，连点合并为一次） |
| Bridge 连接成功（`refreshBridgeStatus`） | 有 pending → `flushFavSync(true)` |
| 页面加载（Bridge 在线且有 pending） | `flushFavSync(true)` |
| `visibilitychange`→hidden / `beforeunload` | `flushFavSync(true)`（fetch `keepalive`，尽力不丢） |
| 点 SYNC 按钮（兜底） | 现有日报 sync + `flushFavSync(true)` |

### 2.3 顶部“未同步”提示

- `updateBridgeButtons()`：有 pending 时在 bridgeBtn 文案追加 `· N 条待同步`；同步中显示“同步收藏中…”；失败用 title 显示原因。
- Bridge 离线时同样显示待同步条数，连上自动补推。

### 2.4 Pages 构建延迟规避

刚 push 的 favorites.json 约 1 分钟后才在 Pages 生效，所以**同步成功时优先用 Bridge 回传的 merged 结果即时更新**，不等 Pages；`loadFavorites` 的远端读取作为跨设备兜底。

---

## 3. Bridge 改造（`bridge/cmd/radar-bridge/main.go`）

### 3.1 配置

```go
type Config struct {
    VaultPath string                       // SecondBrain（已有）
    PagesURL  string                       // 网站（已有）
    Listen    string                       // 已有
    RepoPath  string `json:"repo_path"`    // 新增：agents-radar 专用 clone
}
```

- `install` 新增 `--repo`：绝对路径、含 `.git`、有 `origin` remote。
- `loadApp` 校验 `RepoPath`（可选；未配置则 `/api/favorites/sync` 返回明确错误，前端保持 pending）。

### 3.2 路由调整

```go
mux.HandleFunc("/api/favorites",      a.auth(a.handleFavoritesGet))   // GET 返回 favorites.json
mux.HandleFunc("/api/favorites/sync", a.auth(a.handleFavoritesSync))  // POST 合并+写+push+md
// 删除：/api/favorites/check、旧 POST /api/favorites（每条写 vault md）
```

### 3.3 新模型 + 删旧代码

- 新增 `FavoriteItem`、`FavoriteSource`、`FavoritesFile`（对齐 §1.1/1.2）。
- **删除**：`FavoriteCandidate`、`favoriteID`、`normalizeFavoriteURL`、旧 `renderFavorite`、`handleFavorite`、`handleFavoriteCheck`。
- `State.Favorites map[string]string`（旧 id→vault 路径）不再使用 → 保留字段以兼容旧 state.json，但不读写。

### 3.4 `handleFavoritesSync` 流程

```
1. 解析 body.items；逐条校验 id 为 64-hex、kind 合法、source.date/report 合法；非法条目跳过
2. 加同步锁（复用 syncMu / flock，避免与日报 sync 抢）
3. RepoPath 未配 → 400「未配置 repo_path」
4. git 安全检查 + 取基线（§3.5）
5. merged = merge(baseline_items, incoming_items)        // §1.4
6. 原子写 RepoPath/favorites.json（稳定排序，updated_at）
7. 渲染并写 VaultPath/收藏/AI情报收藏.md（§3.6，跳过墓碑）
8. git add / commit / push（§3.5）
9. 返回 { items: merged(含墓碑), alive, total, pushed, error }
```

### 3.5 git 操作（安全版，只碰 favorites.json）

```
fetch:   git -C repo fetch origin master           （失败 → pushed=false，仍写本地+md，记 error）
脏检查:  git -C repo status --porcelain
         若存在 favorites.json 以外的改动 → 报错中止（保护用户工作区）
基线:    baseline = git -C repo show origin/master:favorites.json   （无则空库）
写文件 → git -C repo add favorites.json
判空:    git -C repo diff --cached --quiet -- favorites.json
         无变化 → 跳过 commit/push，pushed=true（已最新）
提交:    git -C repo commit -m "favorites: sync (N alive)" -- favorites.json   （仅此路径）
推送:    git -C repo push origin HEAD:master
         non-fast-forward → fetch + 用新基线重新 merge + 重写 + commit --amend + push 重试一次
         再失败 → pushed=false，记 last_error
```

- 全程 `git -C repo`；**绝不**切分支、`reset --hard` 用户工作区、`git add .`。
- `RepoPath` **推荐用 Bridge 专用 clone**（脏检查让主工作区相对安全，但专用 clone 最干净）。
- 推送鉴权用该 clone 现有 git 配置（SSH / 凭证助手）；Bridge 不代管 token。

### 3.6 Obsidian 收藏 md 渲染

写 `VaultPath/收藏/AI情报收藏.md`，跳过墓碑，按收藏日期（`created_at` 当地日期）倒序分组，组内时间倒序：

```md
# AI 情报收藏

> 由 Radar Bridge 自动生成，请勿手改。最后更新：2026-06-11 16:00

## 2026-06-10

### [链接] Cursor 新功能讨论
- 网站：github.com
- 来源：2026-06-10 · GitHub AI 趋势
- 回到 Radar：<https://ivo-eu.github.io/agents-radar/?focus=ID#2026-06-10/ai-trending>
- 原始链接：<https://github.com/...>

> 收藏时上下文：xxx

### [段落] 关于 OpenClaw provider 兼容性的分析
- 来源：2026-06-09 · AI Agents 生态
- 回到 Radar：<https://ivo-eu.github.io/agents-radar/?focus=ID#2026-06-09/ai-agents>

> 收藏的段落全文……

### [整篇] Product Hunt AI 产品日报
- 来源：2026-06-09 · Product Hunt AI
- 回到 Radar：<...>
```

- 链接用 `<...>` 包裹（Obsidian 不渲染成卡片、不误断行）。
- 文件头注明“自动生成勿手改”（避免被覆盖）。
- 旧的单条 `收藏/YYYY/YYYY-MM/*.md` **保留不动**（计划 §6），不再新增。

### 3.7 状态 / 错误

- `status()` 增 `favorites_last_sync`、`favorites_last_error`。
- 同步失败原因透传前端，显示“未同步：<原因>”。
- pending 条数由前端按 `syncedAt` 自算，Bridge 不维护。

---

## 4. 安全复核（对照计划 §4.2）

- ✅ 只监听 `127.0.0.1:4399`；CORS 仅允许 `pages_url` origin；pairing token 不变。
- ✅ 不读 vault 其他内容、不上传 vault 内容（只写 `收藏/AI情报收藏.md`）。
- ✅ favorites.json 公开：内容全部来自**已公开的日报**（标题/链接/段落摘录），不引入新隐私。计划 §3.1 已确认接受公开。
- ✅ git 只动 favorites.json 一个路径，脏检查保护用户工作区。
- ⚠️ 本地预览 `127.0.0.1:4173` 无法配对（return URL 必须等于 pages_url），本地预览收藏只存浏览器、不同步 —— 计划 §4.4 接受。

---

## 5. 测试计划

**Bridge（`go test ./...`）**

- 合并：同 id 取最早 `created_at`；墓碑取最新；复活规则；补字段；倒序。
- id 校验：非 64-hex / 非法 kind / 非法 date 被跳过。
- md 渲染：分组、倒序、墓碑排除、三类格式、特殊字符转义。
- git 失败处理：fetch 失败仍写本地+md 且 `pushed=false`；脏工作区拒绝；non-ff 重试（用临时 git 仓库做夹具）。

**前端（Node 纯函数）**

- 合并规则与 Bridge 对齐（同组用例两端各跑，结果一致）。
- pending 水位判定（`syncedAt` 比较）。

**手动验收**

- 收藏 3 条 → 约 5s 后自动产生 **1 次** commit；favorites.json 含 3 条；Obsidian md 出现 3 条。
- 取消 1 条 → json 该条 `deleted_at` 有值；md 少一条；刷新 / 换浏览器不复活。
- 断网收藏 → 顶部“N 条待同步”；恢复连接自动补推。
- 重复收藏同链接不产生重复条目。

---

## 6. 工作量与风险

- 代码量（估）：`index.html` +约 120 行（软删 + 同步客户端 + 提示）；`main.go` 删约 140 行旧逻辑、加约 250 行（模型 / 合并 / git / md）；`main_test.go` +约 200 行。
- 主要风险：
  - **git push 鉴权** —— 专用 clone 的 SSH / 凭证需预先配好，Bridge 不代管。
  - **Pages 构建延迟** —— 用 Bridge 回传 merged 结果即时更新规避。
  - **与日报 GHA 抢 master** —— non-ff fetch + 重新合并 + amend 重试规避。
- 不在本阶段（留 Phase 4）：收藏搜索、标签、备注、按来源日期筛选、本地预览写入支持。

---

## 7. 开工前待确认（已定）

1. **专用 clone 路径**：`install --repo-url <git地址>` 自动 clone 到 `~/Library/Application Support/RadarBridge/agents-radar`；也支持 `--repo <现成路径>`。
2. **pending 提示位置**：追加在 bridgeBtn 文案后（如 `SecondBrain 已连接 · 3 条待同步`）。
3. **md 文件名**：`收藏/AI情报收藏.md`，照用。

## 8. 实现纪要（与设计的差异 / 落地细节）

开工时发现 `index.html` 已被另一 agent 重写为 1371 行版本（浏览器 ID 用 `encodeGoQuery/goQueryEscape` **逐字节匹配 Go**，两边都有契约测试；已有本地墓碑 `favHidden`；段落收藏存“选中文字”+`blockAnchorId` 锚点）。Phase 2 在此基础上实现，几处落地与初稿不同：

- **ID 契约保留**：除了“Bridge 信任并校验 64-hex id”，还**保留**了那套“浏览器 ID = Go ID”的契约测试（`favorites-id.test.ts` + `TestFavoriteIDMatchesWebsiteContract`），作为归一化漂移的护栏。
- **墓碑模型简化**：不再用单独的 `favHidden` 集合；**墓碑直接是 `favoriteStore.items` 里带 `deleted_at` 的条目**。`removeFavorite` 打 `deleted_at`、`addFavorite` 命中墓碑则复活、视图/计数/高亮按 `aliveItems()` 过滤。旧 `radar-favorites-hidden` 一次性迁移成 epoch 弱墓碑（真实再收藏可覆盖）。
- **前端合并引擎**：`index.html` 内 `mergeItem/mergeSource/favEarliest/favLatest` 是 Go `mergeFavorites` 的镜像，已用 Node 跑通四个场景（最早 created_at、补字段、墓碑、复活、顺序无关），与 Go 行为一致。
- **git 流程改为 reset 对齐**：专用 clone 下 `fetch → 脏检查 → reset --hard origin/master 对齐 → 合并写入 → add/commit(仅 favorites.json)/push HEAD:master`，保证 fast-forward；push 被拒则 fetch 后重跑一轮（再次 reset 重新落在最新 origin/master 上）。比初稿的 `commit --amend` 更干净（amend 不改父提交，仍会 non-ff）。
- **无感同步**：收藏即写本地 + 防抖 5s 自动 flush；连接成功 / 页面加载有 pending / 切后台·关闭前（`keepalive`）均自动 flush；SYNC 按钮为兜底（同时触发日报 sync + 收藏 flush）。push 失败（离线/凭证）保持 pending、顶部显示“N 条待同步”、30s 自动重试。
- **测试**：Bridge 新增 8 个测试（校验、合并墓碑/复活/倒序、md 渲染、**真实 git 夹具**的 commit+push / 脏工作区拒绝 / 离线路径），`go build`/`vet`/`gofmt` 干净；`pnpm test` 219 个全过。

仍需你/其他 agent 完成的（需 GitHub 凭证，代码侧无关）：跑 `install --repo-url`、配好该 clone 的 push 鉴权、对真实仓库做端到端验证 —— 见 `收藏功能Phase2交接清单.md`。
