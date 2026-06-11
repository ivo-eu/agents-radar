# Radar 收藏功能 Phase 3 技术实现文档

> 范围：收藏「打磨」三项 + 卡片展开/收回。
> A. 收藏搜索　B. 标签/备注　C. 按来源日期筛选　（外加：收藏卡片可展开/收回）
> 实现位置：`index.html`（主要）+ `bridge/cmd/radar-bridge/main.go`（仅 B：md 渲染 + 合并保留字段）。
> 状态：**待实现方案**（本文档为开工前的技术设计，供审阅）。

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
若 a.edited_at >= b.edited_at： tags/note 取 a，否则取 b   // LWW，整组替换
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

- 内存态：`favSearchQuery`、`favSourceDateFilter`（''=全部）、`favTagFilter`（''=全部，单选；再点同一个取消）。
- 全部与侧栏 `currentFavFilter`（类型）以 **AND** 组合。

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
- **标签筛选**：头部 chips 来自所有 alive 收藏的 tags 并集，点选过滤（见 §2.1）。

### 2.5 卡片展开/收回

- **收起态（默认）**：徽章 + 标题 + 一行来源/时间；段落摘录截断（CSS clamp，约 2 行）；有标签显示 chips。整卡或右侧 chevron 可点开。
- **展开态**：完整摘录 + 完整 meta（网站等）+ 标签 + 备注 + 编辑入口 + 操作（回到原文/打开外链/取消收藏）。
- 状态：`favExpanded = new Set()`（id 集合，跨重渲染保留）。展开/收回**只切 CSS class**（`.fav-card.expanded`），不整列重渲染；编辑保存才重渲染。
- 现有操作（回到原文/打开外链/取消收藏）移入展开态，避免收起态过挤。

> 备注：「展开/收回」按本文理解为**收藏卡片**级。若你指的是左侧栏「★ 我的收藏」整块像月份组那样可折叠，那是另一种小改动，审阅时说一声即可切换。

### 2.6 同步响应改为「合并」而非「覆盖」（纳入 Phase 2 审查 5.1 修复）

当前 `flushFavSync` 成功后 `favoriteStore.items = res.items`（整体覆盖）。若同步在途时用户又改了标签/收了新条目，会被响应覆盖丢失。改为：

```text
收到 res.items → 与「当前本地 items」逐 id mergeItem（本地为 b=较新方参与合并）
→ 写回 store，而非整体替换
```

这样在途的新增/编辑不丢，且与 Bridge 的合并语义一致（mergeItem 已是 LWW，edited_at/revived_at 会保护较新的本地改动）。

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

## 6. 开工前确认

1. 「展开/收回」= **收藏卡片**级（§2.5），认可？还是要左侧栏「我的收藏」整块可折叠？
2. 标签筛选**单选**（点一个过滤、再点取消）够用，还是要多选 AND？
3. 确认我同时改 `index.html` 与 `main.go`（当前有其它 Agent 在动这两个文件，需排期避免冲突）。
