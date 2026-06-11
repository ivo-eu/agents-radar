# Radar 收藏功能 Phase 1 技术实现文档

> 范围：网站端 localStorage 书签系统 —— 三类收藏、原文高亮、回跳、左侧“我的收藏”视图。
> **不含** Bridge / GitHub / SecondBrain 同步（见 Phase 2）。
> 实现位置：单文件 `index.html`（CSS + 内联 `<script>`）。

---

## 1. 目标

把原先“每条链接/段落旁常驻 `☆` 按钮、点击直写 Bridge”的摘录归档，改造成一套**网站内书签系统**：

- 收藏三类对象：整篇报告、链接、段落。
- 收藏只写浏览器 `localStorage`，刷新后保留。
- 原文中已收藏目标高亮。
- 左侧“我的收藏”入口，可按类型筛选，按收藏时间倒序。
- 可从收藏列表回跳到原文位置并临时闪烁高亮；链接还能打开外链。
- 不污染视觉：链接旁不常驻星星，段落不常驻按钮。

为 Phase 2 的多端同步预留接口：ID 规则、数据结构、远端 `favorites.json` 合并逻辑都按最终契约设计。

---

## 2. 数据模型

### 2.1 localStorage

| key | 内容 |
|-----|------|
| `radar-favorites` | `{ version: 1, updated_at: ISO, items: FavoriteItem[] }` |

### 2.2 FavoriteItem

```jsonc
{
  "id": "<64-hex>",                 // SHA-256，见 §3
  "kind": "report | link | excerpt",
  "title": "...",
  "url": "https://...",             // 仅 link
  "site": "github.com",             // 仅 link，hostname
  "excerpt": "...",                 // excerpt=段落全文；link=上下文文本；report=""
  "section": "...",                 // 所在章节标题，可空
  "created_at": "ISO",              // 收藏时间，列表排序主键
  "source": {
    "date": "2026-06-10",           // 来源日报日期
    "report": "ai-trending",        // 报告 id
    "label": "GitHub AI 趋势",
    "url": "https://.../#2026-06-10/ai-trending"
  }
}
```

> Phase 1 的 `removeFavorite` 为**硬删除**（直接从 `items` 移除），无 `deleted_at` 墓碑。
> 墓碑（软删）在 Phase 2 引入，用于跨设备同步防复活。

### 2.3 内存状态

```js
let favoriteStore = { version: 1, updated_at: '', items: [] };
let favoriteIds   = new Set();   // 当前已收藏的 id，供高亮 / toggle / 计数
let currentFavFilter = null;     // null=日报视图；否则 'all'|'report'|'link'|'excerpt'
let pendingFocus  = null;        // 待回跳的目标 id
```

---

## 3. 稳定 ID 生成

ID 用于去重、原文高亮、回跳锚点。**浏览器用 `crypto.subtle` 算 SHA-256 hex**，输入串按类型固定，并与 Phase 2 的 Go Bridge 共用同一契约（Bridge 信任浏览器 id，不重算）。

| kind | 输入串 |
|------|--------|
| report | `report\n<date>\n<report>` |
| link | `link\n<normalizedUrl>` |
| excerpt | `excerpt\n<date>\n<report>\n<normalizedText(excerpt)>` |

归一化规则（`index.html`）：

- `normalizeUrl(href)`：去 `#hash`、host 转小写、去末尾 `/`，**保留 query**。
- `normalizeText(s)`：折叠连续空白为单空格、trim，**保留大小写**。

```js
async function sha256Hex(input) {
  const buf = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(input));
  return [...new Uint8Array(buf)].map(b => b.toString(16).padStart(2, '0')).join('');
}
```

> `crypto.subtle` 需安全上下文：GitHub Pages（https）与 `localhost` 均可；`file://` 直开不可用。
> 因 ID 异步，`enhanceFavorites` 为 `async`，渲染后用 `Promise.all` 批量算 ID。

---

## 4. 三类收藏交互

收藏 UI 在报告渲染后由 `enhanceFavorites(container)` 注入；`data-favorite-id` 写到对应 DOM 元素上，作为高亮与回跳锚点。

### 4.1 整篇报告 — 标题星标（唯一常驻按钮）

- 在 `.md > h1` 旁注入 `☆ / ★` 按钮（`.fav-report-btn`）。
- 已收藏：星标变实心，标题区加 `.fav-report-saved`（左边框 + 浅底高亮）。
- `id = favId('report', {date, report})`。
- 点击 toggle 收藏 / 取消。

### 4.2 链接 — hover 浮层（不常驻星星）

- 移除所有常驻星星。
- 鼠标悬停链接约 **320ms** → 弹出浮层 `.fav-popover`：未收藏显示“收藏链接”，已收藏显示“取消收藏链接”。
- 已收藏链接加 `.fav-link-saved`（浅底 + 虚线下划线）。
- `id = favId('link', {url})`；`title` 取链接文本，`site` 取 hostname，`excerpt` 取所在块上下文。
- 浮层带 hover 保活（移到浮层上不消失），`bindLinkHover()` 管理显隐计时。

### 4.3 段落 — 选区浮层（不常驻按钮）

- 监听全局 `mouseup`：选区落在 `#content .md` 内、选中文本 ≥ 8 字 → 在**鼠标光标右上角**（`clientX+12, clientY-34`）弹出浮层“收藏段落 / 取消收藏段落”。
- **锚定策略**：选区只用于触发；实际收藏锚定到选区所在的**整个段落块**（`p / blockquote / li / td / h2-4`），保存该块全文为 `excerpt`。
  - 理由：`data-favorite-id` 在重渲染后能稳定命中、高亮与回跳可靠，且符合 ID 规则（id = hash(块全文)）。一个段落块对应一条收藏，重复选取不产生重复条目。
- 已收藏块加 `.fav-block-saved`（左边框 + 浅底）。
- `id = favId('excerpt', {date, report, excerpt})`。

> 浮层为单例，`link hover` 与 `selection` 共用 `ensurePopover / showPopover / hidePopover`；`mousedown` 在浮层上 `preventDefault` 以保留选区。

---

## 5. 左侧“我的收藏”视图

### 5.1 侧栏入口

REPORTS 区上方固定区块：

```
★ 我的收藏
  全部收藏   <count>
  整篇报告   <count>
  链接       <count>
  段落       <count>
```

由 `renderFavNav()` 渲染，带各类型计数；点击切换 `location.hash` 到 `favorites` / `favorites/<kind>`。这不是文件夹，是同一收藏库的筛选视图。

### 5.2 列表视图

`renderFavoritesView(kind)` 把列表渲染进 `#content`：

- 仅取对应 `kind`（`all` 取全部），按 `created_at` 倒序。
- **分组**：今天收藏 / 昨天收藏 / 前天收藏 / 更早按“`YYYY-MM-DD` 收藏”（`favTimeGroup()`）。
- 每条卡片（`.fav-card`）：

```
[类型] 标题
（段落额外显示完整摘录）
来源：2026-06-10 · GitHub AI 趋势
收藏时间：2026-06-10 15:42
网站：github.com           （仅链接）
回到原文/回到报告 · 打开外链（仅链接） · 取消收藏
```

- 操作：
  - `回到原文 / 回到报告` → `gotoFavorite(item)`。
  - `打开外链` → `window.open(url)`（仅链接）。
  - `取消收藏` → `removeFavorite(id)` 后重渲染。

---

## 6. 路由与回跳

### 6.1 路由

`getHash` 升级为 `parseRoute()`：

| hash | 路由 |
|------|------|
| `#favorites` / `#favorites/<kind>` | 收藏视图 |
| `#YYYY-MM-DD/<report>` | 日报视图 |

`handleRoute()` 统一分发（init 与 `popstate` 都调用它）；`reportExists()` 校验日报存在，否则回落到首个日报。

### 6.2 回跳 + 闪烁

- `gotoFavorite(item)`：设置 `pendingFocus = id`，`history.pushState` 到 `?focus=<id>#<date>/<report>`，再 `handleRoute()`。
- `readFocus()`：从 `location.search` 读 `focus` 填入 `pendingFocus`（支持外部分享的 `?focus=` 链接）。
- `loadReport` 渲染并 `await enhanceFavorites()`（data-favorite-id 就绪）后调用 `applyPendingFocus()`：
  - 找到 `[data-favorite-id=<id>]` → `scrollIntoView` → 加 `.fav-flash` 动画（1.7s 后移除）。
  - 随后清掉 URL 上的 `?focus`，保留 hash。
- 若目标报告已是当前报告，`handleRoute` 直接 `applyPendingFocus()`（不重载）。

---

## 7. 远端 favorites.json 合并（只读）

`loadFavorites()` 在 init 时：

1. 读本地 `radar-favorites`。
2. 尝试 `fetch('./favorites.json')`（Phase 1 通常不存在 → 当空库；offline 忽略）。
3. 合并：同 id 保留较早 `created_at` 并补缺字段；本地为主。
4. 存入内存 `favoriteStore`，写操作只改 localStorage。

> 这是为 Phase 2 预埋的读取通道：Bridge 推送 favorites.json 后，网站能在另一浏览器/设备读到。Phase 1 不写远端。

---

## 8. 关键样式（CSS class 一览）

| class | 作用 |
|-------|------|
| `.fav-report-btn` / `.fav-report-saved` | 报告星标 / 标题高亮 |
| `.fav-link-saved` | 链接高亮（浅底 + 虚线下划线） |
| `.fav-block-saved` | 段落高亮（左边框 + 浅底） |
| `.fav-flash` | 回跳临时闪烁动画 |
| `.fav-popover` | hover / 选区共用浮层 |
| `.fav-nav-head` / `.fav-nav` / `.fav-nav-btn` / `.fav-count` | 侧栏“我的收藏” |
| `.fav-card` / `.fav-badge` / `.fav-excerpt` / `.fav-meta` / `.fav-actions` / `.fav-action` / `.fav-empty` | 列表视图 |

---

## 9. 已知限制（Phase 1 接受，留待后续）

- **段落收藏锚定整块**：保存的是选区所在段落块全文，而非用户精确选中的子串（见 §4.3，为高亮/回跳稳定性做的取舍）。
- **ID 依赖文本不变**：日报 markdown 若重新生成导致段落文字变化，excerpt 的 id 会变，旧高亮/回跳失效。这是基于内容生成 ID 的固有限制。
- **硬删除无墓碑**：Phase 1 取消收藏直接移除；多端同步下的“删了不复活”由 Phase 2 软删墓碑解决。
- **移动端选区**：长按选择体验受限，浮层在移动端会降级。
- **`file://` 不可用**：需 `localhost` 或 Pages（`crypto.subtle` + `fetch` 要求）。

---

## 10. 验证

- ✅ `node --check` 内联脚本语法通过。
- ✅ ID 生成 / URL 与文本归一化 / 时间分组逻辑用 Node 跑通（大小写、末尾斜杠、hash 归一；query 保留；today/昨天/前天/日期分组正确）。
- ✅ 无残留旧引用（`makeFavoriteButton` / `refreshFavoriteStates` / `/api/favorites` 等已清）。
- ⏳ 真实浏览器交互（hover 时机、选区位置、回跳滚动）需本地 `localhost` 手动确认。
