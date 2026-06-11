# Phase 2 交接清单 —— 部署与凭证（需 GitHub 账号，代码侧已完成）

> 代码（Bridge Go + 前端 index.html）已写完并测试通过。这份清单是**需要你的 GitHub 凭证**才能跑的部分：
> 让 Bridge 拥有一个能 push 的 agents-radar 专用 clone，并对真实仓库做端到端验证。
> Bridge **不代管 token** —— 用专用 clone 自身的 git 鉴权（SSH key 或凭证助手）。

## 前置

- 一台已配好能 push 到 `agents-radar` 的 git 鉴权的 Mac（SSH key 已加到 GitHub，或 https 凭证助手已登录）。
- 已知你的 Obsidian（SecondBrain）vault 绝对路径。

## 步骤

```bash
# 1) 编译 Bridge
cd bridge
go build -o radar-bridge ./cmd/radar-bridge

# 2) 安装：vault + 自动 clone 一个 Bridge 专用仓库
./radar-bridge install \
  --vault "<你的 Obsidian vault 绝对路径>" \
  --repo-url git@github.com:ivo-eu/agents-radar.git    # 或 https://github.com/ivo-eu/agents-radar.git
# clone 落在 ~/Library/Application Support/RadarBridge/agents-radar
# install 会装好 launchd，serve 常驻；也可手动 ./radar-bridge serve

# 3) 确认凭证 OK（关键！Bridge 不替你配鉴权）
cd ~/Library/Application\ Support/RadarBridge/agents-radar
git push origin HEAD:master --dry-run     # 能通过即 OK；报错就先把 SSH/PAT 配好
```

## 验证（不依赖前端，直接打 Bridge）

```bash
# 取 token：~/Library/Application Support/RadarBridge/state.json 里的 "token"
TOKEN=<token>

curl -s -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" \
  -d '{"items":[{"id":"'"$(printf %064d 1)"'","kind":"link","title":"测试链接","url":"https://github.com/x/y","site":"github.com","created_at":"2026-06-11T00:00:00Z","source":{"date":"2026-06-10","report":"ai-trending","label":"GitHub AI 趋势"}}]}' \
  http://127.0.0.1:4399/api/favorites/sync
```

期望：

- 返回 `{"pushed":true,"alive":1,"total":1,"items":[...]}`。
- 专用 clone（及 GitHub）多一条 `favorites: sync (1 alive)` 提交，根目录有 `favorites.json`。
- vault 出现 `收藏/AI情报收藏.md`，含该条目 + 回到 Radar 链接。
- 等 GitHub Pages 重建（约 1 分钟）后，打开 https://ivo-eu.github.io/agents-radar/ 能读到收藏。

## 前端端到端

1. 打开正式 Pages 站点（**不是** localhost；配对要求 return URL = pages_url）。
2. 点 `本机未连接` → 配对 → 变 `SecondBrain 已连接`。
3. 收藏 2~3 条（报告星标 / hover 链接 / 选中段落）→ 约 5s 后顶部“待同步”消失，GitHub 多一次提交。
4. 取消 1 条 → favorites.json 里该条 `deleted_at` 有值；刷新 / 换浏览器不复活。
5. 断网收藏 → 顶部“N 条待同步”；恢复连接自动补推。

## 排错对照

| 现象 | 原因 / 处理 |
|------|------------|
| `pushed:false` + error 含 git push 报错 | 凭证 / 网络问题；先手动 `git push` 跑通 |
| error 含「仓库工作区有未提交改动」 | `--repo` 指到了你的工作区；改用 `--repo-url` 专用 clone |
| error 含「收藏同步未启用」 | install 时没给 `--repo/--repo-url`，重装 |
| 配对失败 / CORS | 用的是 localhost 预览；改用正式 Pages 站点 |

## 安全说明

- 只监听 `127.0.0.1:4399`，CORS 仅放行 pages_url，pairing token 不变。
- git 只 `add/commit/push` `favorites.json` 一个路径；工作区有其它未提交改动会**拒绝**操作。
- favorites.json 公开，但内容全部来自已公开的日报，不引入新隐私。
- 不读 vault 其它内容、不上传 vault 内容（只写 `收藏/AI情报收藏.md`）。
