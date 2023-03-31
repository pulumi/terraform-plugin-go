#!/usr/bin/env bash

set -euo pipefail

OLD_MODULE="$1"
NEW_MODULE="$2"

find . -type f -name '*.go' \
     -exec sed -i -e "s,$OLD_MODULE,$NEW_MODULE,g" {} \;
