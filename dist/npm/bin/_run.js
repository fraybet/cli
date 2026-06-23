// Shared launcher: exec the downloaded native binary, forwarding args + stdio.
const path = require("path");
const fs = require("fs");
const { spawnSync } = require("child_process");

module.exports = function run(name) {
  const exe = path.join(
    __dirname,
    process.platform === "win32" ? name + ".exe" : name
  );
  if (!fs.existsSync(exe)) {
    console.error(
      `[fray] the ${name} binary isn't installed. Reinstall with:\n` +
        `  npm rebuild @fraybet/cli\n` +
        `or install from source: go install github.com/fraybet/cli/cmd/${name}@latest`
    );
    process.exit(1);
  }
  const r = spawnSync(exe, process.argv.slice(2), { stdio: "inherit" });
  process.exit(r.status === null ? 1 : r.status);
};
