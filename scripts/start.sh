#!/usr/bin/env bash
set -euo pipefail

if [ $# -ne 1 ]; then
  echo "Usage: $0 <new-module-path>  (e.g. github.com/abroudoux/caca)"
  exit 1
fi

NEW_MODULE="$1"
OLD_MODULE="github.com/abroudoux/gotemp"
OLD_NAME="gotemp"
NEW_NAME="$(basename "$NEW_MODULE")"

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

FILES=$(find . -type f \
  \( -name '*.go' -o -name 'go.mod' -o -iname 'dockerfile' \
     -o -name '*.yml' -o -name '*.yaml' -o -name '*.toml' -o -name '*.md' \) \
  -not -path './.git/*')

for f in $FILES; do
  perl -pi -e "s/\Q$OLD_MODULE\E/$NEW_MODULE/g; s/\b\Q$OLD_NAME\E\b/$NEW_NAME/g" "$f"
done

go mod tidy

echo "Project renamed: $OLD_MODULE -> $NEW_MODULE ($NEW_NAME)"
