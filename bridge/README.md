# Radar Bridge

Radar Bridge is a small, zero-dependency Go service for one local SecondBrain vault.

It:

- downloads `digests/` from the configured agents-radar Pages site using `sync-manifest.json`;
- writes only to `日报/AI情报/` and `收藏/`;
- exposes a loopback-only API at `127.0.0.1:4399`;
- pairs with the Pages UI using a local token;
- never reads or uploads unrelated vault content;
- never runs Git commands or calls an LLM.

## Commands

```bash
radar-bridge install \
  --vault "/absolute/path/to/SecondBrain" \
  --pages-url "https://ivo-eu.github.io/agents-radar"

radar-bridge serve
radar-bridge sync
radar-bridge status
```

`install` creates two user LaunchAgents:

- `com.ivo.radar-bridge`: starts the local API at login and performs a catch-up sync;
- `com.ivo.radar-bridge.sync`: requests a sync every day at 09:15.

Configuration, pairing state and logs live under:

```text
~/Library/Application Support/RadarBridge/
```
