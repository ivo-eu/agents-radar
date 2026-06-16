# Radar 收藏功能 Phase 3 实现审查

> 审查日期：2026-06-15
> 审查范围：`index.html`、`bridge/cmd/radar-bridge/main.go`、相关 Go/Vitest 测试及 `收藏功能Phase3技术实现.md`
> 结论：未发现 P0/P1 数据破坏或 Bridge 阻断问题；发现 4 个 P2、2 个 P3。建议修完 P2 并补齐前端测试后再部署 Phase 3。

## 1. 总体结论

Phase 3 的四项主体能力已经落地：

- 收藏搜索、来源日期筛选、标签 AND 筛选已经接入收藏视图。
- 标签、备注和 `edited_at` 已进入前端存储、同步水位和 Bridge 数据模型。
- Bridge 已实现 tags/note 的 LWW 合并、输入净化和 Obsidian Markdown 输出。
- 卡片展开/收起、内联编辑、链接标题直达已经实现。

基础质量检查全部通过，但现有自动化测试没有真正覆盖 Phase 3 前端功能。当前主要风险不是数据丢失，而是筛选状态可能让用户看到无法直接恢复的空列表，以及实现与技术文档对收起态交互、LWW 同时间戳和测试覆盖的描述不一致。

## 2. 审查发现

### P2-3.1 跨类型保留的筛选状态可造成“假空列表”，部分场景没有直接清除入口

**位置**

- `index.html:1340-1350`：`favItemMatches` 始终应用全局日期和标签状态。
- `index.html:1362-1378`：先按收藏类型缩小 `items`，再从当前类型生成日期和标签选项。
- `index.html:1384-1388`：当前类型没有任何标签时，整个标签栏不渲染。
- `index.html:1481-1500`：零匹配时只隐藏卡片并把计数改为 0，没有“无匹配收藏”或重置入口。

**问题**

`favSourceDateFilter` 和 `favTagFilter` 在切换“全部/整篇/链接/段落”时不会重置，但可见选项只从当前类型生成。

可复现场景：

1. 在“全部收藏”选中仅链接收藏拥有的标签 `#agent`。
2. 切换到“整篇报告”；该类型有收藏，但没有任何标签。
3. `favTagFilter` 仍含 `agent`，所以所有整篇收藏都被隐藏。
4. 因当前类型没有标签，标签栏和“清空标签”均不渲染；页面只显示“共 0 条”，用户需要切回其他类型才能解除筛选。

日期也有类似问题：旧日期不在当前下拉选项中时，控件视觉上可能显示“全部来源日期”，内部状态却仍是旧日期，继续过滤全部卡片。

这也违反技术文档中“日期选项、标签 chips 来自所有 alive 收藏”和“空结果显示无匹配收藏”的约定。

**建议**

- 日期和标签选项统一从 `aliveItems()` 生成，不要从类型过滤后的 `items` 生成；或者切换类型时清除当前类型不存在的筛选值。
- 无匹配时显示明确空状态，并提供“清除全部筛选”按钮。
- 增加跨类型保留标签、跨类型保留日期、当前类型无标签、零匹配恢复四组测试。

### P2-3.2 卡片收起态与设计不符，来源/时间完全隐藏，整卡也不能展开

**位置**

- `index.html:341-344`：收起态隐藏整个 `.fav-detail`。
- `index.html:1412-1428`：来源、收藏时间和网站全部位于 `.fav-detail` 内。
- `index.html:1449-1453`：只有右侧“展开/收起”按钮绑定切换事件。

**问题**

技术文档要求收起态显示“徽章 + 标题 + 一行来源/时间”，并允许“整卡或右侧 chevron”展开。当前实现收起时完全看不到来源和收藏时间，也只有文字按钮可展开。

这不影响数据正确性，但会显著降低大量收藏的浏览效率，也属于明确的验收项未落地。

**建议**

- 将精简来源/时间摘要移到 `.fav-detail` 外，展开态再显示完整 metadata。
- 给卡片绑定展开事件，并排除标题链接、编辑输入框和操作按钮等交互元素，避免误触。
- 添加收起态可见字段和点击区域测试。

### P2-3.3 `edited_at` 相同时 LWW 结果依赖合并顺序

**位置**

- `index.html:632-636`
- `bridge/cmd/radar-bridge/main.go:1289-1303`
- `bridge/cmd/radar-bridge/main_test.go:281-299`

**问题**

前后端在 `edited_at` 相等时都偏向第二个输入，只要第二个输入有 tags 或 note。直接验证结果：

```text
mergeItem(A, B) -> B
mergeItem(B, A) -> A
```

因此合并不满足交换律。浏览器加载时是“本地后合并远端”，Bridge 同步时是“远端后合并本地”，同一对相同时间戳、不同内容的记录可能在两个位置选择不同结果。

现有“顺序无关”测试只覆盖 `edited_at` 不同的情况，没有覆盖真正的 tie。技术文档 §1.2 写的是 `a.edited_at >= b.edited_at` 时取 a，§8 又声称已验证顺序无关，均与当前代码不一致。

**建议**

- 为相同 `edited_at` 定义确定性 tie-break，例如对规范化后的 tags/note 做稳定字典序比较；前后端必须完全镜像。
- 添加“相同时间戳、不同内容、正反顺序结果一致”的 JS 和 Go 契约测试。

### P2-3.4 Phase 3 前端自动化测试实际缺失，技术文档的测试声明不准确

**位置**

- `src/__tests__/favorites-merge.test.ts:15-27`：测试模型没有 `tags`、`note`、`edited_at`。
- `src/__tests__/favorites-merge.test.ts:76-179`：只覆盖 Phase 2 的 created/delete/revive/in-flight。
- `收藏功能Phase3技术实现.md:166-176`、`230`、`239-246`：声称已覆盖多过滤叠加和前端 tags/note LWW。

**问题**

`pnpm test` 的 229 个测试全部通过，但没有一项测试 Phase 3 的：

- 搜索、日期、标签、类型组合过滤。
- 标签 AND 语义。
- 编辑保存及 `edited_at`。
- 前端 tags/note LWW。
- 卡片展开/收起。
- 零匹配状态和跨类型筛选恢复。

Bridge 的净化测试也只验证了 trim、去重、20 个数量上限和 note trim，没有验证单标签 32 字符、note 2000 字符及 Unicode 边界。

P2-3.1、P2-3.2 和 P2-3.3 正是因为这些路径没有自动化覆盖而未被发现。

**建议**

- 从真实 `index.html` 提取 `favItemMatches`、`parseTagsInput`、`mergeItem`，增加 Phase 3 纯函数测试。
- 使用轻量 DOM 环境覆盖 `renderFavoritesView`、`applyFavFilters` 和编辑保存状态机。
- 补齐 Go 的 32/2000 字符边界和相同时间戳 tie 测试。
- 修正技术文档，只有实际进入测试套件的用例才标记为已完成。

### P3-3.5 前后端“字符长度”定义不一致

**位置**

- `index.html:576`、`1460`：JavaScript `length`/`slice` 按 UTF-16 code unit 截断。
- `bridge/cmd/radar-bridge/main.go:1213-1215`、`1231-1235`：Go 按 Unicode rune 截断。

**问题**

对 emoji 等非 BMP 字符，前端一个字符计为 2，Bridge 计为 1。例如 40 个 emoji 标签在前端会截成 16 个，而 Bridge 的契约允许 32 个。备注也会出现 1000/2000 个 emoji 的差异。

**建议**

前端改用 `Array.from(text)` 或 `[...text]` 按 Unicode code point 截断，并增加中英文、emoji 边界测试。

### P3-3.6 技术文档中的 Obsidian 路径仍是旧值

**位置**

- `收藏功能Phase3技术实现.md:244`：写为 `收藏/AI情报收藏.md`。
- `bridge/cmd/radar-bridge/main.go:1434-1449`：实际写入 `日报/AI情报/收藏/AI情报收藏.md`。

**建议**

统一文档路径，避免部署或人工验收时在旧目录查找文件。

## 3. 自动验证结果

以下检查已通过：

```text
pnpm typecheck
pnpm lint
pnpm format:check
pnpm test
  15 files passed
  229 tests passed

node inline-script syntax check
gofmt -d
go vet ./...
go build ./cmd/radar-bridge
go test -race -cover ./...
  coverage: 45.0%
```

本轮未完成真实浏览器点击验收：Codex 本地浏览器安全策略拒绝访问 `127.0.0.1:4173`，因此没有绕过策略执行。发布前仍需人工验证搜索、筛选恢复、编辑同步、跨浏览器 LWW、卡片展开和 Obsidian 双产物。

## 4. 仓库基线提醒

当前工作区：

```text
master...origin/master [behind 32]
HEAD: cd75af6
origin/master: 2890c1a
```

同时存在 workflow、provider、日报产物、搜索索引和收藏功能等大量未提交改动。Phase 3 修复完成后，不应直接把整个工作区一次性提交；应先同步远端并明确收藏相关提交范围，避免把生成产物或其他 Agent 的修改混入 Phase 3。

## 5. 建议修复顺序

1. 修 P2-3.1：筛选选项来源、无匹配状态和一键清除。
2. 修 P2-3.3：定义并测试相同 `edited_at` 的确定性 tie-break。
3. 修 P2-3.2：补齐收起态摘要和整卡展开交互。
4. 修 P2-3.4：补齐真实 Phase 3 前端测试和 Bridge 边界测试。
5. 修 P3-3.5、P3-3.6，并更新技术文档实现纪要。

完成以上修改后，再进行 Pages + 新 Bridge 二进制的真实端到端验收。

## 6. 复审结果（2026-06-15 第二轮）

> 复审范围：针对本报告 P2-3.1 至 P3-3.6 的修复再次检查 `index.html`、Bridge、测试和技术文档。

### 6.1 已确认修复

- P2-3.1 假空列表：已改为从全部 `aliveItems()` 生成日期和标签选项，并加入“清除全部筛选”和零匹配提示。代码位置：`index.html:1386-1404`、`1444-1465`、`1531-1534`。
- P2-3.2 收起态：已新增常驻 `.fav-submeta` 显示来源/收藏时间，整卡收起态可点击展开，并排除链接、按钮、输入框和文字选区。代码位置：`index.html:341-346`、`1423-1424`、`1471-1484`。
- P3-3.5 字符长度：前端标签和备注已改按 Unicode code point 截断，Bridge 仍按 rune 截断。代码位置：`index.html:581-582`、`1491-1492`，`main.go:1213-1215`、`1231-1235`。
- P2-3.4 测试覆盖：已新增 `src/__tests__/favorites-phase3.test.ts`，并给旧的 `favorites-merge/sync` 抽取列表补入 `favTagsNoteKey`。当前测试数量从 229 增至 241。
- P3-3.6 路径文档：实现纪要中的 Markdown 路径已更正为 `日报/AI情报/收藏/AI情报收藏.md`。

### P2-6.2 新发现：LWW tie-break 的“前后端镜像”对非 ASCII 不成立

**位置**

- `index.html:571-572`、`641-646`
- `bridge/cmd/radar-bridge/main.go:1289-1310`
- `src/__tests__/favorites-phase3.test.ts:164-170`
- `bridge/cmd/radar-bridge/main_test.go:303-319`

**问题**

这轮把相同 `edited_at` 的 tie-break 改成 `tagsNoteKey` 字典序较大者。方向是对的，但当前 JS 和 Go 的“字典序”并不等价：

- JavaScript 的 `>=` 比较字符串按 UTF-16 code unit。
- Go 的字符串 `>=` 按 UTF-8 字节序。

因此 ASCII 测试能过，但非 ASCII 特别是 emoji / 非 BMP 字符会分歧。实测：

```text
JS: "😀" < "￿"
Go: "😀" > "￿"
```

如果两个设备在同一个 `edited_at` 上写入不同的非 ASCII 标签/备注，浏览器加载和 Bridge 同步仍可能选择不同赢家，继续违背“JS↔Go 逐字节镜像”的目标。

**建议**

- 让 JS tie-break 显式按 UTF-8 字节比较：用 `TextEncoder().encode(favTagsNoteKey(...))` 做逐字节比较。
- 或者让 Go 改成按 Unicode code point 序列比较，但要和 JS 明确同一规则。
- 补一组非 ASCII 契约测试，例如 tags `["😀"]` vs `["￿"]`，要求 JS 与 Go 选择同一边。
- 顺手把 `index.html:572` 的不可见分隔符字面量改成 `'\x01'`，减少后续误改风险。

### P3-6.3 技术文档 §1.2 仍保留旧 tie 规则

**位置**

- `收藏功能Phase3技术实现.md:57-60`

**问题**

文档前半部分仍写着：

```text
若 a.edited_at >= b.edited_at： tags/note 取 a，否则取 b
```

这已经不是当前实现。后面的 §8.1 写了 `tagsNoteKey`，但前后不一致，交给下一个 Agent 时容易按旧规则回退。

**建议**

把 §1.2 改为当前目标规则：`edited_at` 大者胜；相等时使用前后端一致的确定性 key/tie-break。

### 6.4 当前验证结果

以下检查已通过：

```text
pnpm typecheck
pnpm lint
pnpm format:check
pnpm test
  16 files passed
  241 tests passed

node inline-script syntax check
gofmt -d
go vet ./...
go build ./cmd/radar-bridge
go test -race -cover ./...
  coverage: 45.5%
```

本轮仍未做真实浏览器点击验收；上次 Codex 本地浏览器策略拒绝访问 `127.0.0.1:4173`，本轮未再次绕过。发布前仍需人工走一遍：筛选恢复、卡片点击展开、编辑标签/备注、跨浏览器同步、Obsidian 双产物。

### 6.5 复审结论

上轮 4 个 P2 + 2 个 P3 中，除 tie-break 的非 ASCII 镜像问题外，其余按代码审查已经基本落实。当前建议先修 **P2-6.2** 和 **P3-6.3**，再进入部署/真实端到端验收。

## 7. 修复落实记录（2026-06-15，第三轮）

针对 §6 复审的 P2-6.2 与 P3-6.3，已完成修复，改动集中在 `index.html`、两份收藏测试与两份文档；Go tie-break 实现本就按 UTF-8 字节序，未改，只让 JS 对齐它。

### P2-6.2 LWW tie-break 非 ASCII 镜像（已修）

- `index.html`：
  - `favTagsNoteKey` 的分隔符由内嵌裸控制字符改为可见转义 `'\x01'`（与 Go `tagsNoteKey` 一致，减少误改风险）。
  - 新增模块级 `favTextEncoder = new TextEncoder()` 与 `favKeyGreaterEq(a, b)`：把两个 key 编码成 UTF-8 字节后**逐字节比较**返回 `a >= b`，镜像 Go 字符串 `>=`（Go 按 UTF-8 字节序）。
  - `mergeItem` 的 tie 分支由 `favTagsNoteKey(a) >= favTagsNoteKey(b)`（JS 默认 UTF-16 code unit 比较）改为 `favKeyGreaterEq(favTagsNoteKey(a), favTagsNoteKey(b))`。
- 实测：JS 与 Go 对 `["😀"]` vs `["￿"]`（同 `edited_at`）均选 `😀`（UTF-8 中 `F0 9F 98 80` > `EF BF BF`），且正反顺序一致；此前 JS 会按 UTF-16（`D83D` < `FFFF`）选 `￿`，与 Go 分歧——根因消除。
- 契约测试：
  - `src/__tests__/favorites-phase3.test.ts` 新增「非 ASCII tie 选 😀 且顺序无关」用例（并把 `favKeyGreaterEq` 加入函数提取列表、补 `favTextEncoder` 常量）。
  - `bridge/cmd/radar-bridge/main_test.go` 的 `TestMergeTagsNoteTieIsDeterministicAndCommutative` 末尾新增 `["😀"]` vs `["￿"]` 双向断言。
  - 顺带修正 `favorites-merge.test.ts`、`favorites-sync.test.ts` 的函数提取列表（补 `favKeyGreaterEq` + `favTextEncoder`），否则它们 grab 的 `mergeItem` 会 `favKeyGreaterEq is not defined`。

### P3-6.3 技术文档 §1.2（已修）

- `收藏功能Phase3技术实现.md` §1.2 的旧规则 `若 a.edited_at >= b.edited_at：tags/note 取 a` 改为当前实现：`edited_at` 大者胜；相等时取 `tagsNoteKey` 在 **UTF-8 字节序**下较大的一方（前后端逐字节镜像：JS `favKeyGreaterEq` ↔ Go 字符串比较）。§8.1 的「字典序」表述同步精确为「UTF-8 字节序」。

### 本轮验证

```text
pnpm typecheck              通过
pnpm lint                   通过
pnpm test                   16 files / 243 tests passed（新增非 ASCII tie 用例）
node inline-script syntax   OK
cd bridge && gofmt -l       干净
cd bridge && go vet ./...   通过
cd bridge && go build ./cmd/radar-bridge   成功
cd bridge && go test -race -cover ./...    coverage 45.5%
```

> 说明：`pnpm format:check`（仅检查 `src`）全量会因 `src/__tests__/generate-manifest.test.ts` 失败——该文件是本会话之外、与 Phase 3 无关的预存未提交改动，未在本轮处理；本轮改动的收藏相关 `src` 测试文件均通过 prettier。`index.html` 不在 `format:check` 范围内（脚本只查 `src`）。
>
> 仍待人工：真实浏览器/Obsidian 端到端（筛选恢复、卡片展开、编辑同步、跨浏览器 LWW、双产物一致）；git commit/push 待用户审核后进行。
