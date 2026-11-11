# XXL-JOB Console (Crayon) 0.4.4

Independent fork maintained by Crayon-hua, based on Ezreal-byte/dbx-plugin-xxl-job-console and the existing MCP access-declaration fix.

- New plugin ID: `io.github.crayon-hua.xxljob-console`; publisher: `crayon-hua`.
- Explicit access levels for all 22 MCP tools restore external MCP visibility on current DBX hosts.
- Updated connection/workbench/log contribution references and package handshake verifier.
- Corrected backend handshake version to match the package version.
- Preserved upstream permission checks, preview confirmations and MIT license.

This installs alongside the upstream plugin. Existing connections and plugin data are not migrated automatically. Create a connection under the Crayon plugin and authorize it in DBX plus the per-connection AI/MCP panel.

No new claims of live XXL-JOB deployment testing are made by this fork; inherited compatibility reports remain historical upstream evidence. Packages are unsigned candidates until DBX Store review/signing.
