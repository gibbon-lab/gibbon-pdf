#!/usr/bin/env bash
set -e

cd "$(dirname "$0")/../"

mkdir -p assets/bin
go build -o assets/bin/api cmd/gibbon-pdf/main.go
