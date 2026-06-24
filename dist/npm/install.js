// Postinstall: download the prebuilt fray + fray-mcp binaries for this platform
// from the matching GitHub release into ./bin. The bin/run-*.js launchers exec
// them. Falls back to a clear message (and `go install`) on unsupported setups.
const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

const REPO = "fraybet/cli";
const pkg = require("./package.json");
const version = "v" + pkg.version;

const osMap = { darwin: "darwin", linux: "linux", win32: "windows" };
const archMap = { x64: "amd64", arm64: "arm64" };
const os = osMap[process.platform];
const arch = archMap[process.arch];
const exe = os === "windows" ? ".exe" : "";
const bins = ["fray" + exe, "fray-mcp" + exe];

const binDir = path.join(__dirname, "bin");

if (!os || !arch) {
  console.warn(
    `[fray] no prebuilt binary for ${process.platform}/${process.arch}. ` +
      `Install from source: go install github.com/${REPO}/cmd/fray@latest`
  );
  process.exit(0); // don't fail the whole npm install
}

const url = `https://github.com/${REPO}/releases/download/${version}/fray_${pkg.version}_${os}_${arch}.tar.gz`;

try {
  fs.mkdirSync(binDir, { recursive: true });
  // curl + tar ship with macOS, Linux, and Windows 10+ (1803+). The windows
  // archive holds fray.exe/fray-mcp.exe, so extract by the platform-correct name.
  execSync(`curl -fsSL "${url}" | tar -xz -C "${binDir}" ${bins.join(" ")}`, {
    stdio: "inherit",
  });
  for (const b of bins) {
    const p = path.join(binDir, b);
    if (fs.existsSync(p)) fs.chmodSync(p, 0o755);
  }
  console.log("[fray] installed " + version + " for " + os + "/" + arch);
} catch (e) {
  console.warn(
    `[fray] could not download ${url}\n` +
      `[fray] the release may not be published yet — install from source instead:\n` +
      `       go install github.com/${REPO}/cmd/fray@latest`
  );
  process.exit(0); // soft-fail so `npm i` succeeds; the launcher will warn at runtime
}
