# Radar 收藏功能 Phase 3 技术实现文档

> 范围：收藏「打磨」三项 + 卡片展开/收回。
> A. 收藏搜索　B. 标签/备注　C. 按来源日期筛选　（外加：收藏卡片可展开/收回）
> 实现位置：`index.html`（主要）+ `bridge/cmd/radar-bridge/main.go`（仅 B：md 渲染 + 合并保留字段）。
> 状态：**已实现**（前端 + Bridge 均完成；`go test`、`pnpm test` 全部通过）。实现纪要见文末 §8。

---

## 0. 背景与编号

原计划的 Phase 3（SecondBrain 统一收藏 md）与「取消收藏跨设备」已在 **Phase 2** 完成。本「Phase 3」对应原计划 **Phase 4「打磨」** 中本次选定的子集：

| 项 | 选中 | 位置 |
|----|------|------|
| A 收藏搜索 | ✅ | 前端 |
| B 标签/备注 | ✅ | 前端 + Bridge |
| C 按来源日期筛选 | ✅ | 前端 |
| 卡片展开/收回 | ✅（本次新增需求） | 前端 |
| D 本地预览也能同步 | ✖ 不做 | — |
| E 墓碑压缩 / F 收藏自动刷新 | ✖ 不做 | — |

**前置依赖**：本阶段建立在 Phase 2 之上。Phase 2 综合审查（`收藏功能Phase2部署验证与综合审查.md`）尚有 P1 待修，其中 **5.1「同步进行中新增的收藏可能被响应覆盖」** 与本阶段强相关——标签/备注编辑会增加同步频次，更容易触发。本阶段把该修复直接纳入设计（见 §2.5）。

当前合并模型已演进为 `max(created_at, revived_at)` vs `deleted_at`（Phase 2 审查 §5.2 已落地，前后端均有 `revived_at`）。本文据此设计。

**现状对齐（开工前已核对当前代码）**：Phase 2 的若干加固已由评审 Agent 落地，本阶段直接复用、无需重做：

- `flushFavSync` 已用 **snapshot + ackWatermark**：`favSyncedAt` 只前进到「快照里最新活动时间」而非「now」，在途新增/编辑不会被确认覆盖（审查 §5.1 已修）。
- 同步响应已走 **`reconcileWithServer(res.items)`**：`服务端 items 先 fold，再 fold 当前本地 items`，用 `mergeItem` 合并而非整体替换（§2.6 因此已无须新做）。
- 已返回并处理 **`vault_written`/`vault_error`**：md 写失败算部分成功、保持 pending 自动重试（审查 §5.5 已修）。

因此 Phase 3 只需**扩展现有合并函数**（`mergeItem`/`mergeTwo` 加 tags/note/edited_at），新字段会自动流过 `reconcileWithServer` 与同步。

---

## 1. 数据模型（仅 B 需要加字段）

`FavoriteItem` 新增三个字段（前端 JS + Go 同步）：

```jsonc
"tags": ["agent", "memory"],   // 标签，整组替换（非并集）
"note": "我的备注…",            // 备注
"edited_at": "2026-06-11T16:00:00.000Z"  // 标签/备注最后编辑时间，用于跨设备 LWW
```

### 1.1 为什么要 `edited_at`

`created_at`/`revived_at`/`deleted_at` 记录的是「收藏/复活/删除」事件，**不能**表达「我在 T 时刻改了备注」。标签和备注是可变字段，跨设备各改一份时需要「谁后改谁赢」。所以加 **item 级 `edited_at`**（任何 tags/note 改动都刷新它）。

> 注意区分：`FavoritesFile.updated_at` 是**文件级**时间戳（已存在），`edited_at` 是**条目级**编辑时间戳（本次新增），两者不冲突。

### 1.2 合并规则增量（前后端镜像）

在现有 `mergeTwo`（Go）/`mergeItem`（JS）基础上加：

```text
edited_at = latest(a.edited_at, b.edited_at)
tags/note LWW，整组替换：
  - edited_at 大者胜；
  - 相等（tie）时取 tagsNoteKey 在 UTF-8 字节序下较大的一方
    （前后端必须逐字节镜像：JS favKeyGreaterEq ↔ Go 字符串比较；
     tagsNoteKey = tags.join("\n") + "\x01" + note）
（其余字段、created/revived/deleted 规则不变）
```

- **tags 整组替换**：用编辑较新一方的整组 tags，而非并集——否则一端删掉的标签会被另一端旧数据复活。
- `favLastActivity` / `pendingFavCount`（前端）把 `edited_at` 也算进「最近活动」，这样改标签/备注会触发自动同步。

### 1.3 校验与净化（Bridge `validFavorites`）

- tags：trim、丢空、去重、单个 ≤ 32 字符、整组 ≤ 20 个；超出截断。
- note：trim、≤ 2000 字符；超出截断。
- 不影响 id/kind/source 既有校验。

---

## 2. 前端（`index.html`）

收藏视图（`renderFavoritesView`）改造为「头部控件 + 多过滤 + 可展开卡片」。

### 2.1 头部控件（与现有侧栏类型筛选并排叠加）

在「我的收藏」视图顶部一行放三个控件（**不动左侧栏**，不新增标签树）：

```
[ 🔍 搜索收藏…        ]   来源日期 [全部 ▼]
#agent  #memory  #cli      ← 出现过的标签 chips，可点选过滤
```

- 内存态：`favSearchQuery`、`favSourceDateFilter`（''=全部）、`favTagFilter`（**`Set`，多选**；点亮多个标签时取**交集 AND**——条目须同时含全部选中标签；再点已选标签取消）。
- 全部与侧栏 `currentFavFilter`（类型）以 **AND** 组合。
- 多选标签的 chip 高亮已选态；提供「清空标签」入口。

### 2.2 A 收藏搜索

实时按 `标题 / 摘录 / 网站(site) / 来源标签(source.label) / 来源日期 / tags` 子串匹配（小写）。数据量小，输入即过滤（不必防抖）。

### 2.3 C 按来源日期筛选

`来源日期` 下拉的选项 = 当前 `aliveItems()` 里出现过的 `source.date` 去重、倒序；选中后只显示该来源日期的收藏。列表仍按 `created_at` 倒序分组。

### 2.4 B 标签/备注

- **展示**：卡片有标签时显示 chips（`#agent`）；有备注时显示备注块。
- **编辑**：展开态卡片里「编辑」→ 内联展开「标签输入框（空格/逗号分隔）+ 备注 textarea + 保存/取消」。保存：
  - 解析标签（拆分、去 `#`、trim、去空去重）→ `item.tags`
  - `item.note` ← textarea
  - `item.edited_at = now`
  - `persistFavorites()` → `scheduleFavSync()` → 重渲染（保持展开）
- **标签筛选**：头部 chips 来自所有 alive 收藏的 tags 并集，**多选取交集**过滤（见 §2.1）。
- **edited_at 要进 `favLastActivity`**：当前 `favLastActivity = latest(latest(created, revived), deleted)`，**不含 edited_at**。Phase 3 须改为
  `latest(latest(created, revived), latest(deleted, edited_at))`，否则改标签/备注不会触发自动同步、pending 也算不到。

### 2.5 卡片展开/收回

- **收起态（默认）**：徽章 + 标题 + 一行来源/时间；段落摘录截断（CSS clamp，约 2 行）；有标签显示 chips。整卡或右侧 chevron 可点开。
- **展开态**：完整摘录 + 完整 meta（网站等）+ 标签 + 备注 + 编辑入口 + 操作（回到原文/打开外链/取消收藏）。
- **链接收藏标题可直接点开外链**（新增需求）：当 `kind === 'link'`，卡片标题渲染为真实超链接 `<a href={url} target="_blank" rel="noopener noreferrer">`，**一键直达外链**，不必先「回到原文」再点。仅放行 `http(s)`。「回到原文」（跳回日报内位置）与「打开外链」按钮保留，但标题本身已可点，省去多余跳转。收起态标题同样可点。
- 状态：`favExpanded = new Set()`（id 集合，跨重渲染保留）。展开/收回**只切 CSS class**（`.fav-card.expanded`），不整列重渲染；编辑保存才重渲染。
- 现有操作（回到原文/打开外链/取消收藏）移入展开态，避免收起态过挤。

> 备注：「展开/收回」按本文理解为**收藏卡片**级。若你指的是左侧栏「★ 我的收藏」整块像月份组那样可折叠，那是另一种小改动，审阅时说一声即可切换。

### 2.6 同步在途不丢（**已由 Phase 2 加固完成，本阶段无须新做**）

开工前核对：`flushFavSync` 已用 snapshot + ackWatermark，且响应走 `reconcileWithServer`（`mergeItem` 合并而非整体覆盖）——审查 §5.1 已修。Phase 3 只要把 tags/note/edited_at 加进 `mergeItem`，编辑/新增就会**自动**正确流过 reconcile，不会被在途响应覆盖。**不要**改回 `favoriteStore.items = res.items`。

### 2.7 过滤主循环

`renderFavoritesView` 改为：

```text
items = aliveItems()
  → filter 类型(currentFavFilter)
  → filter 搜索(favSearchQuery)
  → filter 来源日期(favSourceDateFilter)
  → filter 标签(favTagFilter)
  → sort created_at desc → 按收藏日期分组渲染
头部渲染搜索框/日期下拉/标签 chips；空结果显示「无匹配收藏」。
```

---

## 3. Bridge（`bridge/cmd/radar-bridge/main.go`）

仅两处，协议不变（仍 `/api/favorites/sync` 发全量）：

### 3.1 合并保留 tags/note/edited_at

`mergeTwo` 增加 §1.2 的 LWW 规则；`FavoriteItem` 加 `Tags []string` `Note string` `EditedAt string`（JSON：`tags` `note` `edited_at,omitempty`）。

### 3.2 md 渲染带标签/备注

`renderFavoritesMarkdown` 每条在 `- 来源` 之后追加：

```md
- 标签：#agent #memory
- 备注：我的备注内容
```

（无标签/无备注则不输出对应行；备注折叠为单行。）

### 3.3 校验净化

`validFavorites` 按 §1.3 净化 tags/note（截断、去空去重）。

---

## 4. 测试

**前端（Node 纯函数）**
- 多过滤叠加（类型 ∧ 搜索 ∧ 日期 ∧ 标签）结果正确。
- `mergeItem`：tags 整组替换、note 按 `edited_at` LWW、与 Go 同用例结果一致。
- 同步响应「合并而非覆盖」：模拟在途新增不丢。

**Bridge（`go test`）**
- `mergeTwo`：tags/note 按 `edited_at` LWW；并集不发生。
- `renderFavoritesMarkdown`：含标签/备注行；无则不输出。
- `validFavorites`：tags/note 截断净化；既有用例不破。
- 既有 8 个 Phase 2 测试 + ID 契约测试保持通过。

**手动**
- 给收藏加标签/备注 → 同步后 favorites.json 带字段、Obsidian md 出现标签/备注行；换浏览器可见。
- 两端各改备注再同步 → 后改的赢，不并集。
- 搜索/日期/标签/类型任意叠加过滤正确；卡片展开收回流畅、状态保持。

---

## 5. 改动量与风险

- `index.html`：约 +150 行（头部控件、多过滤主循环、展开/收回、编辑 UI、合并增量、同步改合并）。
- `main.go`：约 +35 行（FavoriteItem 3 字段、mergeTwo LWW、md 两行、validFavorites 净化）+ 测试。
- 不碰路由 / git 流程 / 安全模型 / 配对。
- 风险：
  - **依赖 Phase 2 审查 5.1 的修复方向**（本文 §2.6 已把「合并而非覆盖」纳入）。若 Phase 2 实现 Agent 也在改 `flushFavSync`，需协调避免双改冲突。
  - 跨设备字段级 LWW 仅靠 `edited_at` 时间戳，单用户场景足够；真多人并发非目标。
  - 与并发改 `index.html`/`main.go` 的其它 Agent 存在撞车可能——开工前确认归属。

---

## 6. 开工前确认（已定）

1. 「展开/收回」= **收藏卡片**级（§2.5）。✅ 先按此做，后续可再调整为侧栏整块折叠。
2. 标签筛选 = **多选取交集（AND）**（§2.1）。✅
3. 协调：当前有评审/实现 Agent 在并发改 `index.html` 与 `main.go`（已观察到 `revived_at`、`reconcileWithServer`、`vault_written` 等加固陆续落地）。**开工前需确认这两个文件此刻无人在写**，否则会像 Phase 2 那样写完被重写。建议本阶段收藏功能由单一 Agent 负责，其余 Agent 暂停改 favorites 相关代码。

## 7. 改动清单（精确到函数 / 文件，供实现）

**`index.html`**

| 位置 | 改动 |
|------|------|
| 状态区 | 新增 `favSearchQuery`（''）、`favSourceDateFilter`（''）、`favTagFilter`（`new Set()`）、`favExpanded`（`new Set()`） |
| `mergeItem` | 加 `edited_at = favLatest(a,b)`；tags/note 取 `edited_at` 较大一方（整组替换） |
| `favLastActivity` | 改为含 `edited_at`：`latest(latest(created,revived), latest(deleted, edited_at))` |
| `addFavorite`/`removeFavorite` | 不动（不碰 edited_at；revive 仍走 revived_at） |
| `renderFavoritesView` | 头部加 搜索框/来源日期下拉/标签 chips；过滤主循环 类型∧搜索∧日期∧标签(AND)；卡片改收起/展开，展开态含 完整摘录+标签 chips+备注+编辑面板+操作；**link 卡片标题渲染为 `<a href target=_blank rel=noopener>` 一键开外链（仅 http(s)）** |
| 新增 `parseTagsInput()` | 拆分、去 `#`、trim、去空去重、限长限量 |
| 新增 编辑保存处理 | 写 `tags/note/edited_at` → `persistFavorites` → `scheduleFavSync` → 重渲染（保持展开） |
| 展开/收回 | 用 `favExpanded` + `.fav-card.expanded` class，仅切 class 不整列重渲 |
| CSS | `.fav-card.expanded`、tag chip、note 块、编辑面板、头部控件行、摘录 clamp |
| `reconcileWithServer`/`flushFavSync` | **不动**（已正确合并；新字段自动流过 `mergeItem`） |

**`bridge/cmd/radar-bridge/main.go`**

| 位置 | 改动 |
|------|------|
| `FavoriteItem` | 加 `Tags []string` `json:"tags,omitempty"`、`Note string` `json:"note,omitempty"`、`EditedAt string` `json:"edited_at,omitempty"` |
| `mergeTwo` | 加 `EditedAt = latest`；tags/note 取 `edited_at` 较大一方（整组替换，非并集） |
| `validFavorites` | 净化 tags（trim/去空/去重/单个≤32/整组≤20）、note（≤2000）截断 |
| `renderFavoritesMarkdown` | 每条 `- 来源` 后追加 `- 标签：#a #b`、`- 备注：…`（无则不输出） |

**测试**：前端纯函数（多过滤叠加、mergeItem tags/note LWW 与 Go 对齐）、Go（`mergeTwo` LWW、md 含标签/备注、`validFavorites` 净化、既有用例不破）。

---

## 8. 实现纪要

按本文档实现，几处落地细节：

- **数据字段**：`FavoriteItem`（Go + JS）新增 `tags`/`note`/`edited_at`，均 `omitempty`；空值不写盘、不上线，保持 favorites.json 干净。
- **合并 LWW**：`tags/note` 取 `edited_at` 较大一方**整组替换**（Go `mergeTagsNote`、JS `mergeItem`），并集不发生；已用 Node 验证 JS↔Go 行为一致（later-edit-wins / 删标签不复活 / 空字段不出现 / 顺序无关）。`favLastActivity` 已纳入 `edited_at`，改标签/备注会触发自动同步。
- **过滤不丢焦点**：搜索/日期/标签过滤走 `applyFavFilters()`——只对已渲染卡片切 `display` 并隐藏空分组标题、更新计数，**不重建 DOM**，所以搜索框输入不丢焦点。标签 chips 因要刷新选中态用轻量重建。
- **卡片展开/收回**：`favExpanded`（id 集合，跨重渲染保留）+ `.fav-card.expanded`；展开/收起只切 class 不重建。收起态摘录 2 行 clamp、详情区隐藏，但**来源/时间摘要 `.fav-submeta` 常驻可见**，**整卡可点展开**（排除链接/按钮/输入框，且选中文字时不触发）。
- **链接一键直达**：`kind==='link'` 卡片标题渲染为 `<a target=_blank rel=noopener>`（仅 http(s)，`favSafeHttpUrl` 校验），收起/展开态都可点；「回到原文」「打开外链」按钮保留。
- **编辑**：展开态「编辑」→ `.editing` 显示标签输入 + 备注 textarea；保存写 `tags/note/edited_at` → persist → `scheduleFavSync` → 重渲染（保持展开）。标签 ≤20×≤32、备注 ≤2000，前后端双重净化，**前端按 Unicode code point 截断**（`[...t]`，与 Go rune 语义一致，emoji 不再两端不一致）。
- **md 渲染**：`日报/AI情报/收藏/AI情报收藏.md`（实际写入路径，由 `favoritesMarkdownRel` 决定）每条 `原始链接` 后追加 `- 标签：#a #b`、`- 备注：…`（无则不输出）。
- **未改动**：`reconcileWithServer`/`flushFavSync`（已在 Phase 2 加固，新字段自动流过）、路由、git 流程、安全/配对。

### 8.1 Phase 3 实现审查（`收藏功能Phase3实现审查.md`）修复

审查给出 4 P2 + 2 P3，已全部修复：

- **P2-3.1 假空列表**：日期/标签筛选选项改为从 `aliveItems()`（全部 alive）生成，不再从类型过滤后的子集生成；零匹配显示「没有匹配的收藏」+「清除全部筛选」；筛选条还提供「清除全部筛选 / 清空标签」入口。
- **P2-3.2 收起态**：新增常驻 `.fav-submeta`（来源·标签源·收藏时间）；整卡可点展开（排除交互元素与选区）。
- **P2-3.3 LWW 同时间戳**：`edited_at` 相等时按 `tagsNoteKey`（tags `\n` 连接 + `\x01` + note）取 **UTF-8 字节序**较大者，**确定且与顺序无关**；Go 用字符串 `>=`（UTF-8 字节序），JS 用 `favKeyGreaterEq`（`TextEncoder` 编码后逐字节比较，不用默认的 UTF-16 比较），两端逐字节镜像，非 ASCII/emoji 也一致（见审查 P2-6.2）。
- **P3-3.5 字符长度**：前端标签/备注截断改 code point。
- **P3-3.6 文档路径**：md 路径已更正为实际的 `日报/AI情报/收藏/AI情报收藏.md`。

- **测试（P2-3.4 已补齐，并修正本节声明使其与套件一致）**：
  - 新增前端 `src/__tests__/favorites-phase3.test.ts`（12 例）：`parseTagsInput`（拆分/去#/去重/20 上限/32 code point 截断）、`favItemMatches`（搜索/日期/标签 AND/组合）、tags/note LWW（后改赢/无并集/同时间戳顺序无关/空字段不出现）、`favLastActivity` 含 `edited_at`。
  - Go 新增/补齐（共 7 例相关）：tags/note LWW、无并集、**同时间戳 tie 确定且顺序无关**、`sanitize` 32 rune / 2000 rune 边界、`validFavorites` 净化、md 含标签/备注。
  - 同时给已有的 `favorites-merge.test.ts`、`favorites-sync.test.ts` 的函数提取列表补上 `favTagsNoteKey`（否则它们抽取的 `mergeItem` 在 tie 分支会报未定义）。
  - 验证：`pnpm typecheck`/`lint`/`format:check` 通过；`pnpm test` **241 个全过**；Go `go test`（含 race）全过、`gofmt` 干净；`index.html` `node --check` 通过；JS↔Go LWW/tie 行为 Node 实测一致。

仍待你/部署 Agent：真实浏览器端到端（本轮因本地浏览器安全策略未执行）——搜索/筛选恢复、编辑同步、跨浏览器 LWW、卡片展开、Obsidian 双产物；并把工作区前端 + 重新 build 的 Bridge 部署到正式 Pages（参考 `收藏功能Phase2交接清单.md`）。
