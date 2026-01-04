#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
OUTPUT_DIR="${ROOT_DIR}/dist/go"

mkdir -p "${OUTPUT_DIR}"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"
TAGS="${TAGS:-}"

BIN_NAME="rustdesk-go"
OUT_PATH="${OUTPUT_DIR}/${BIN_NAME}-${GOOS}-${GOARCH}"

if [[ -n "${TAGS}" ]]; then
  echo "Building with tags: ${TAGS}"
  GOOS="${GOOS}" GOARCH="${GOARCH}" go build -tags "${TAGS}" -o "${OUT_PATH}" ./go/cmd/rustdesk
else
  GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o "${OUT_PATH}" ./go/cmd/rustdesk
fi

echo "Built ${OUT_PATH}"
