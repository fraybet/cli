#!/bin/sh
# Fray CLI installer — downloads the latest prebuilt `fray` (and `fray-mcp`)
# binary for your platform from GitHub Releases.
#
#   curl -fsSL https://fray.bet/install.sh | sh
#
# Env:
#   FRAY_VERSION       pin a version (e.g. v0.1.0); default: latest release
#   FRAY_INSTALL_DIR   install dir; default: /usr/local/bin (falls back to ~/.local/bin)
set -eu

REPO="fraybet/cli"

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
arch="$(uname -m)"
case "$arch" in
  x86_64 | amd64) arch="amd64" ;;
  arm64 | aarch64) arch="arm64" ;;
  *) echo "fray: unsupported arch '$arch'. Install from source: go install github.com/$REPO/cmd/fray@latest" >&2; exit 1 ;;
esac
case "$os" in
  darwin | linux) ;;
  *) echo "fray: unsupported OS '$os'. Install from source: go install github.com/$REPO/cmd/fray@latest" >&2; exit 1 ;;
esac

version="${FRAY_VERSION:-}"
if [ -z "$version" ]; then
  version="$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep -m1 '"tag_name"' | cut -d '"' -f4)"
fi
if [ -z "$version" ]; then
  echo "fray: no release found yet. Install from source: go install github.com/$REPO/cmd/fray@latest" >&2
  exit 1
fi

num="${version#v}"
url="https://github.com/$REPO/releases/download/$version/fray_${num}_${os}_${arch}.tar.gz"
dest="${FRAY_INSTALL_DIR:-/usr/local/bin}"

echo "fray: downloading $version for $os/$arch …"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "$url" | tar -xz -C "$tmp"

if [ ! -w "$dest" ] 2>/dev/null && [ "${FRAY_INSTALL_DIR:-}" = "" ]; then
  dest="$HOME/.local/bin"
  mkdir -p "$dest"
fi
mkdir -p "$dest"
for bin in fray fray-mcp; do
  if [ -f "$tmp/$bin" ]; then
    install -m 0755 "$tmp/$bin" "$dest/$bin" 2>/dev/null || { mv "$tmp/$bin" "$dest/$bin"; chmod +x "$dest/$bin"; }
  fi
done

echo "fray: installed to $dest/fray"
case ":$PATH:" in *":$dest:"*) ;; *) echo "fray: add $dest to your PATH" ;; esac
echo "fray: run 'fray version' to check, and see https://fray.bet/docs"
