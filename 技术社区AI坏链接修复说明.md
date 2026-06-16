# 技术社区 AI 坏链接修复说明

> 日期：2026-06-15
> 关联页面：`2026-06-14 / ai-community`
> 关联改动：`index.html`、`src/generate-manifest.ts`、`src/__tests__/generate-manifest.test.ts`
> 目的：为其他 Agent 提供可直接审查的修改背景、实现细节和验证结论

## 1. 问题背景

GitHub Pages 上 `2026-06-14` 的「技术社区 AI」板块中，有一条 Dev.to 文章链接无法打开。

用户提供的页面文案如下：

```md
3. [I Pointed a Skill Linter at a 52k-Star Repo. Here Is What 84/100 Looks Like.]([https://dev.to/sayed_ali_alkamel/i-pointed-a-skill-linter-at-a](https://dev.to/sayed_ali_alkamel/i-pointed-a-skill-linter-at-a) 52k-star-repo-here-is-what-84100-looks-like-28cn)
```

实际正确链接应为：

```text
https://dev.to/sayed_ali_alkamel/i-pointed-a-skill-linter-at-a-52k-star-repo-here-is-what-84100-looks-like-28cn
```

## 2. 根因分析

问题不在 GitHub Pages 部署层，也不在浏览器跳转层，而在于日报 Markdown 源文本中出现了异常的嵌套链接结构：

```md
[标题]([url](url) slug-rest)
```

当前前端直接执行：

```js
marked.parse(raw)
```

因此 `marked` 会把这段非法 Markdown 按原样解析，最终生成错误的 `<a href>`，导致点击失败。

这个问题有两个影响面：

- 页面渲染：GitHub Pages 中该链接不可点击或 href 错误。
- RSS/Feed：`src/generate-manifest.ts` 在生成 `feed.xml` 时同样依赖 Markdown 转 HTML，如果不处理，订阅内容中也会保留坏链接。

## 3. 修改方案

本次没有尝试回写某一天的具体日报源文件，而是加了一个通用兜底修复，专门处理这类被错误拆开的 Markdown 链接。

### 3.1 前端渲染修复

文件：`index.html`

在 `renderMarkdown(raw)` 之前新增两层处理：

1. `joinBrokenMarkdownUrl(head, tail)`
   - 将被拆开的 URL 头部和剩余 slug 重新拼接。
   - 中间按需补 `-`，避免产生错误的双连接符或漏连接符。

2. `normalizeBrokenMarkdownLinks(raw)`
   - 用正则识别如下模式：

```md
[title]([https://example.com/foo](https://example.com/foo) bar-baz)
```

   - 规范化为：

```md
[title](https://example.com/foo-bar-baz)
```

然后再调用：

```js
marked.parse(normalizeBrokenMarkdownLinks(raw))
```

### 3.2 RSS 生成修复

文件：`src/generate-manifest.ts`

新增同名逻辑：

- `joinBrokenMarkdownUrl`
- `normalizeBrokenMarkdownLinks`

并在 `getReportContent()` 中，读取 Markdown 文件后先执行规范化，再交给 `marked.parse()`。

这样可以保证：

- 页面端渲染正确
- RSS/Feed 里的 HTML 内容也正确
- 未来同类坏 Markdown 即使再次进入仓库，也有统一兜底

## 4. 具体改动点

### 4.1 `index.html`

位置：`index.html:451`

新增：

- `joinBrokenMarkdownUrl(head, tail)`
- `normalizeBrokenMarkdownLinks(raw)`

调整：

- `renderMarkdown(raw)` 改为先规范化再 `marked.parse()`

### 4.2 `src/generate-manifest.ts`

位置：`src/generate-manifest.ts:82`

新增：

- `joinBrokenMarkdownUrl(head, tail)`
- `normalizeBrokenMarkdownLinks(markdown)`

调整：

- `getReportContent()` 中读取 Markdown 后先做规范化

### 4.3 `src/__tests__/generate-manifest.test.ts`

新增回归测试：

- 输入用户提供的坏 Markdown
- 断言输出为正确的 Dev.to 完整 URL

## 5. 设计取舍

### 5.1 为什么不只修当天源文件

本地工作区当前没有 `digests/2026-06-14/` 目录，说明线上内容与本地仓库状态并不完全一致。直接修某一天的源文件并不能保证当前工作区能覆盖线上问题。

因此这次选择：

- 修通用渲染入口
- 修通用 feed 生成入口

这样即使坏内容来自：

- 旧日报产物
- 外部生成流程
- 未来再次出现的同类 LLM 输出异常

页面和 feed 也都能自动兜住。

### 5.2 为什么只修这一类坏链接

本次正则仅覆盖已观测到的具体异常模式：

```md
[title]([url](url) tail)
```

这样改动范围更小，风险更可控，不会贸然引入一个“试图修复所有异常 Markdown”的宽泛规则，避免误伤正常链接格式。

## 6. 验证结果

已执行并通过：

```text
pnpm test src/__tests__/generate-manifest.test.ts
pnpm typecheck
pnpm lint
```

另外执行了全量（2026-06-15 复审更新）：

```text
pnpm test
16 files passed
243 tests passed
```

全量测试已全部通过——此前记录的 10 个 favorites/sync 失败（`favKeyGreaterEq is not defined/...`）已随收藏功能修复消除，与本次链接修复无关。

> 当前工作区唯一的遗留红灯曾是 `pnpm format:check` 因 `src/__tests__/generate-manifest.test.ts` 未格式化失败（见 `三项修改综合审查.md` P2-1），已运行 Prettier 修复，`pnpm format:check` 现已通过。

## 7. 审查重点建议

建议其他 Agent 重点看以下几点：

1. 正则范围是否足够精确，是否可能误伤正常 Markdown 链接。
2. `joinBrokenMarkdownUrl()` 的连接规则是否覆盖常见 slug 拼接场景。
3. 前端与 `generate-manifest.ts` 是否保持了相同修复语义，避免页面和 RSS 表现不一致。
4. 回归测试是否还需要补充更多边界，例如：
   - tail 前后有多余空格
   - `head` 以 `/` 结尾
   - `tail` 以 `?`、`#` 等连接符开头

## 8. 当前结论

本次修复是一个低侵入、可回归验证、同时覆盖页面和 RSS 的通用兜底方案。

它不依赖修复某一份具体日报源文件，因此更适合当前“线上内容已出现，但本地未必持有对应日期产物”的场景。
