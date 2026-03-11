#!/bin/bash

# Run script for Go project
set -e

# Move to project root (one level up from scripts/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
SERVER_FILE="cmd/server/main.go"

APP_NAME="go-auth"

echo ">>> Starting $APP_NAME..."
cd "$PROJECT_ROOT"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    exit 1
fi

echo ">>> Fetching dependencies..."
go mod tidy

echo ">>> Starting $APP_NAME..."
go run ./$SERVER_FILE

# Back to original directory
cd - > /dev/null