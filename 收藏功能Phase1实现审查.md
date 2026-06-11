# Radar 收藏功能 Phase 1 实现审查

审查日期：2026-06-11

审查范围：

- 对照 `收藏功能重构实现计划.md` 的 Phase 1 要求。
- 重点审查网站端 `index.html` 中的 localStorage 收藏、原文高亮、回跳和左侧“我的收藏”视图。
- 检查前端收藏 ID 与现有 Bridge 规范的兼容性。
- 不要求本轮修改 GitHub 或 SecondBrain 同步逻辑。

## 1. 总体结论

Phase 1 的主要框架已经实现：

- 左侧新增“我的收藏”及按类型筛选入口。
- 支持整篇报告、链接、段落三类收藏。
- 收藏数据写入 localStorage，刷新后仍能恢复。
- 原文存在收藏高亮。
- 收藏列表按 `created_at` 倒序展示。
- 收藏列表支持回到报告或原文。
- 链接收藏支持打开外链。
- 页面会尝试读取公开的 `favorites.json`，文件不存在时按空远端库处理。

但当前仍有 3 个 P1 问题，会直接影响段落收藏和未来 Bridge 同步；另有 1 个 P2 行为问题。建议修复后再将 Phase 1 标记为完成。

## 2. P1：段落收藏保存的是整个块，不是用户选中的文字

### 位置

- `index.html:848-885`
- 关键代码位于 `setupSelectionFavorites()`。

### 当前行为

代码先通过 `window.getSelection()` 判断用户是否进行了文本选择，但生成收藏数据时使用：

```js
const excerpt = cleanExcerpt(block);
```

这里的 `block` 是选区所在的整个 `p`、`li`、`td`、`blockquote` 或标题元素。

因此用户只选择一句话时，实际保存的 `excerpt`、收藏 ID、列表内容和高亮目标仍然基于整个 DOM 块。

### 影响

- 不符合“收藏选中段落/文字”的产品语义。
- 收藏列表可能展示大量未选择内容。
- 同一个长段落中的不同选区无法分别收藏。
- 段落 ID 与用户实际选择内容无关。
- 后续同步到统一收藏库时会保存错误摘录。

### 建议修复

在选区消失前同时保存：

```js
const selectedText = normalizeDisplayText(window.getSelection().toString());
const block = selectionBlock();
```

收藏对象中的 `excerpt` 和稳定 ID 应基于 `selectedText`。`block` 只承担：

- 获取章节标题。
- 提供回跳定位容器。
- 应用原文高亮。

需要注意异步 `sha256` 和点击浮层之间可能导致 Selection 被清空，因此应在 `mouseup` 回调开始时捕获选区文本，而不是点击浮层时重新读取。

### 验收

1. 在一个长段落中只选择其中一句。
2. 点击“收藏段落”。
3. 收藏列表只显示选中的一句。
4. 刷新后仍显示相同摘录。
5. 同一段落选择另一句时能生成另一条收藏。

## 3. P1：标题选区收藏无法恢复高亮或回跳

### 位置

- `index.html:849-857`
- `index.html:942-951`

### 当前行为

`selectionBlock()` 允许以下元素作为收藏容器：

```js
p, blockquote, li, td, h2, h3, h4
```

但 `enhanceFavorites()` 只为以下元素生成 `data-favorite-id`：

```js
p, blockquote, li, td
```

用户可以选择 `h2`、`h3` 或 `h4` 并创建收藏，但重新渲染报告时，这些标题不会获得对应的 `data-favorite-id`。

### 影响

- 标题收藏刷新后不恢复高亮。
- 收藏列表中的“回到原文”找不到目标。
- `applyPendingFocus()` 会静默返回，用户无法知道定位失败。

### 建议修复

不要只通过重新计算“整个 DOM 块文本”的 ID 寻找摘录。建议为 excerpt 收藏保存独立定位信息：

- `excerpt`：用户选中的文本，用于内容和稳定收藏 ID。
- `source.anchor` 或等价字段：所在块的稳定定位 ID。

Phase 1 可以先在 localStorage 中增加该字段，不需要调用 Bridge。回跳时应优先按 anchor 找到容器，再临时闪烁。

若暂时沿用现有模型，至少应让 `selectionBlock()` 与 `enhanceFavorites()` 使用同一套元素选择器，并为 `h2-h4` 生成定位属性。但这只能修复标题场景，不能解决“选区文本 ID 与块定位 ID不同”的根本问题。

### 验收

1. 选择一个 `h2`、`h3` 或 `h4` 标题文字并收藏。
2. 刷新页面，标题仍有收藏状态。
3. 从“我的收藏”点击“回到原文”。
4. 页面滚动到该标题并出现临时闪烁效果。

## 4. P1：前端收藏 ID 与 Bridge 规范不一致

### 位置

- 前端：`index.html:731-756`
- Bridge：`bridge/cmd/radar-bridge/main.go:782-829`

### URL 差异

Bridge 的 `normalizeFavoriteURL()` 会：

- 统一 scheme 和 host 大小写。
- 删除 fragment。
- 删除 `utm_*`、`ref`、`source` 查询参数。
- 使用 `url.Values.Encode()` 对查询参数进行稳定排序。
- 将根路径 `/` 规范为空路径。

前端 `normalizeUrl()` 当前只处理：

- 删除 fragment。
- host 转小写。
- 删除尾部斜杠。

前端没有删除跟踪参数，也没有稳定排序查询参数。

例如以下两个 URL 在 Bridge 中是同一收藏，但前端会生成不同 ID：

```text
https://example.com/post?utm_source=one
https://example.com/post?utm_source=two
```

### 文本差异

Bridge 的 `normalizeText()` 会压缩空白并转为小写；前端只压缩空白。

### 影响

- Phase 1 产生的本地 ID 与 Phase 2 Bridge ID不一致。
- 同一链接可能产生重复收藏。
- 从 `favorites.json` 读取的数据无法匹配原文高亮。
- 本地待同步数据合并时可能被视为不同记录。

### 建议修复

前端规范必须与 Bridge 使用同一套明确规则：

```js
function normalizeFavoriteUrl(raw) {
  const url = new URL(raw, location.href);
  url.protocol = url.protocol.toLowerCase();
  url.hostname = url.hostname.toLowerCase();
  url.hash = '';

  for (const key of [...url.searchParams.keys()]) {
    const lower = key.toLowerCase();
    if (lower.startsWith('utm_') || lower === 'ref' || lower === 'source') {
      url.searchParams.delete(key);
    }
  }

  url.searchParams.sort();
  if (url.pathname === '/') url.pathname = '';
  return url.toString();
}

function normalizeFavoriteText(value) {
  return value.trim().replace(/\s+/g, ' ').toLowerCase();
}
```

需要用固定输入同时验证 Go 和 JS 输出相同 ID。不要只测试规范化结果，还要断言最终 SHA-256 ID。

### 验收

- 不同 `utm_source` 的同一 URL生成相同 ID。
- 查询参数顺序不同的 URL生成相同 ID。
- URL fragment 不影响 ID。
- excerpt 的大小写和连续空白不影响 ID。
- JS 与 Go 对同一组 fixture 生成完全相同的 ID。

## 5. P2：远端收藏的“取消收藏”刷新后会复活

### 位置

- `index.html:519-538`
- `index.html:772-780`
- `index.html:1047-1055`

### 当前行为

页面启动时会把 localStorage 与 `favorites.json` 合并。取消收藏仅从当前 localStorage 状态删除：

```js
favoriteStore.items = favoriteStore.items.filter((it) => it.id !== id);
```

如果该条目来自 `favorites.json`，刷新后它会再次被远端数据合并进来。

### 影响

- UI 显示“取消收藏”成功，但刷新后条目恢复。
- 用户会认为 localStorage 持久化失败。

### Phase 边界

计划中“取消收藏”属于 Phase 4，因此 Phase 1 可以不实现完整的跨端删除。但当前页面已经公开展示取消按钮，所以必须选择一种一致行为：

1. Phase 1 暂时移除取消按钮和已收藏目标的取消文案；或
2. 在 localStorage 增加本地 tombstone/隐藏 ID 集合，使远端条目在当前浏览器内不会重新出现。

不要在本轮修改 Bridge、GitHub 或 SecondBrain 删除逻辑。

### 验收

- 如果保留取消按钮：取消远端条目后刷新，当前浏览器中不会复活。
- 如果移除取消能力：页面上不出现会造成误导的取消操作。

## 6. 已验证通过的功能

本地浏览器手动验证：

- 整篇报告标题旁显示收藏星标。
- 点击后星标变为实心，标题区域高亮。
- 收藏数量即时更新。
- 刷新后整篇报告收藏仍存在。
- 左侧“全部收藏”能够切换到收藏列表。
- 收藏列表能够展示来源和收藏时间。
- 收藏列表能够返回对应报告。
- `favorites.json` 不存在时页面能够正常工作，HTTP 404 被按空远端库处理。

自动检查：

```text
pnpm typecheck
结果：通过

pnpm test
结果：12 个测试文件、203 个测试全部通过

git diff --check -- index.html
结果：通过
```

Bridge 检查：

```text
go version
结果：go1.24.3 darwin/amd64

go test ./...
结果：通过

go test -race ./...
结果：通过

go vet ./...
结果：通过

go test -v -cover ./...
结果：5 个测试通过，语句覆盖率 7.3%
```

现有 Bridge 测试已覆盖：

- URL 跟踪参数规范化。
- 链接收藏 ID 稳定性。
- excerpt 空白和大小写规范化。
- 文件名安全。
- 路径穿越拒绝。

## 7. 测试缺口

当前收藏前端没有自动化测试。建议另一个 Agent 至少补充可独立运行的 JS 纯函数测试：

- `normalizeFavoriteUrl`
- `normalizeFavoriteText`
- `favKeyString`
- SHA-256 稳定 ID fixture
- localStorage 合并和去重
- 远端收藏本地隐藏/取消策略

如果继续把所有 JS 内联在 `index.html`，测试会比较困难。可以将不依赖 DOM 的收藏纯函数提取到独立模块，但不要为了测试大规模重构整个页面。

Bridge 当前覆盖率仅 7.3%，但 Phase 1 不要求扩展 Bridge。Phase 2 开始前应再补：

- `/api/favorites` HTTP handler 测试。
- 收藏合并测试。
- 文件写入失败测试。
- Git 操作失败测试。
- `favorites.json` 渲染和读取测试。

## 8. 建议执行顺序

1. 统一前端和 Bridge 的 URL、文本规范化规则。
2. 分离 excerpt 内容 ID 与原文容器定位 anchor。
3. 修复段落选区，只保存用户选择的文字。
4. 统一标题和普通块的定位增强逻辑。
5. 决定 Phase 1 的远端收藏取消策略。
6. 添加 JS 稳定 ID fixture 测试。
7. 重新进行浏览器手动验收。
8. 运行 `pnpm typecheck`、`pnpm test`、`go test ./...` 和 `go test -race ./...`。

## 9. 不应在本轮修改的内容

为保持 Phase 1 范围，本轮不要：

- 修改 Bridge 收藏 API。
- 写入或提交 `favorites.json`。
- 执行 GitHub commit/push 同步。
- 生成 SecondBrain 统一收藏 Markdown。
- 迁移旧的单条收藏 Markdown。
- 改动日报生成、Provider 或搜索索引逻辑。

修复应主要集中在 `index.html`，只在确有必要时新增前端测试文件或纯函数模块。

## 10. 修复落实记录（2026-06-11）

本轮已按第 8 节顺序完成全部 P1 / P2 修复，范围严格限制在 `index.html`、Bridge 测试与新增前端测试。

### P1-1 段落只保存选中文字（§2）

- `setupSelectionFavorites()` 在 `mouseup` **同步**捕获 `selectedText = normalizeDisplayText(getSelection().toString())` 与 `block`，再进入异步分支，避免选区被清空。
- `excerpt` 与稳定 ID 现在基于 `selectedText`；`block` 仅用于章节、回跳容器和高亮。
- 同一长段落里选不同句子会生成不同收藏（ID 不同），不再合并成整块。

### P1-2 标题/块定位统一 + 内容 ID 与定位 anchor 分离（§3）

- 新增 `FAV_BLOCK_SELECTOR = 'p, blockquote, li, td, h2, h3, h4'`，`selectionBlock()` 与 `enhanceFavorites()` 共用，`h2-h4` 现在也会拿到定位属性。
- 新增 `blockAnchorId(date, report, blockText)`：基于**整块文本** hash，写入 `data-fav-anchor`，并存入 `item.source.anchor`。
- excerpt 收藏：`id`=选中文字内容 ID（去重/同步用），`source.anchor`=所在块定位 ID（高亮/回跳用），两者解耦。
- `gotoFavorite()` 对 excerpt 用 `source.anchor` 作为 `?focus=` 目标；`applyPendingFocus()` 同时匹配 `[data-favorite-id]`（report/link）与 `[data-fav-anchor]`（excerpt）。
- `blockHasSavedExcerpt(anchor)` 决定块高亮：只要还有任一 excerpt 指向该 anchor 就保持高亮（支持一块多收藏）。

### P1-3 前端 ID 与 Bridge 完全一致（§4）

- `normalizeUrl()` 重写为 Bridge `normalizeFavoriteURL()` 的镜像：小写 scheme/host、去 fragment、删 `utm_*`/`ref`/`source`、`searchParams.sort()`、根路径 `/`→空。
- 新增 `encodeGoQuery` / `goQueryEscape`，逐字节复现 Go `url.Values.Encode()`（空格→`+`，非保留字符按 UTF-8 大写百分号编码），保证 query 序列化一致。
- `normalizeText()` 改为压空白+trim+**小写**（与 Bridge 一致）；显示用途改用新增的 `normalizeDisplayText()`（保留大小写）。
- **JS↔Go 一致性已用 10 组 fixture 实测**（含 utm/ref/source、尾斜杠、根斜杠、空格与 `+`、unicode 路径与 query、端口、excerpt 大小写/空白），两端 SHA-256 ID **逐字节相同**。

### P2 远端取消刷新后复活（§5）

- 采用方案 2（本地 tombstone）：新增 `radar-favorites-hidden` 存隐藏 ID 集合 `favHidden`。
- `removeFavorite()` 写入 tombstone；`loadFavorites()` 合并远端后过滤掉 tombstoned ID；`addFavorite()` 重新收藏时清除对应 tombstone。
- 取消远端来源的收藏后刷新，当前浏览器不再复活，取消按钮保留可用。

### 测试

- 新增 `src/__tests__/favorites-id.test.ts`：直接从 `index.html` 抽取真实函数，断言 URL/文本规范化与 **最终 SHA-256 ID**（fixture 由 Go 实现生成）。16 例通过。
- `bridge/.../main_test.go` 新增 `TestFavoriteIDMatchesWebsiteContract`：用同一组 fixture 在 Go 侧锁定相同 ID，任一端规范化漂移都会让两侧测试失败。

### 验证结果

```text
pnpm typecheck   通过
pnpm test        13 文件 / 219 测试通过（含新增 16）
pnpm lint        通过
git diff --check -- index.html   无空白错误
go test ./...        通过
go test -race ./...  通过
go vet ./...         通过
```

> 仍待人工：真实浏览器中 hover 时机、选区浮层位置、回跳滚动与闪烁（需 `localhost` 或 Pages，`crypto.subtle` 要求安全上下文）。
