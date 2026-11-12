# XXL-JOB Console (Crayon) 0.4.5

## 修复 MCP 工具显示为 0/0

- 根因：sidecar 在 `plugin/initialize` 里只声明了 `connections`，没有声明 `mcp`。
- DBX 宿主以 sidecar capability 决定是否加载插件 MCP 工具目录；仅有 `manifest.json` 的 `tools/access` 不够，结果插件中心显示「已开放 0/0 个工具」。
- 对照可正常工作的 DolphinScheduler 插件握手结果为 `["connections","mcp"]`。
- 本次将能力改为 `connections` + `mcp`，保留 0.4.4 已声明的 22 个工具 access。

升级到 0.4.5 后，打开插件 MCP 授权页应能看到工具列表；再在连接的 AI/MCP 面板开启查询工具，并在 DBX MCP 设置中授权连接。
