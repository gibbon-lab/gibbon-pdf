#!/usr/bin/env bash
set -e

cd "$(dirname "$0")/../"

mkdir -p assets/bin
go build --ldflags '-linkmode external -extldflags "-static"' -o assets/bin/api cmd/gibbon-pdf/main.go
