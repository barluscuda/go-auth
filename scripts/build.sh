#!/bin/bash

# Build script for Go project
set -e

# Move to project root (one level up from scripts/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
SERVER_FILE="cmd/server/main.go"

APP_NAME="go-auth"
OUTPUT_DIR="$PROJECT_ROOT/bin"

echo ">>> Starting build..."
cd "$PROJECT_ROOT"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    exit 1
fi

# Create bin/ output folder if not exists
mkdir -p "$OUTPUT_DIR"

echo ">>> Fetching dependencies..."
go mod tidy

echo ">>> Building..."
go build -o "$OUTPUT_DIR/$APP_NAME" ./$SERVER_FILE

echo ">>> Build complete! Binary at: bin/$APP_NAME"

# Back to original directory
cd - > /dev/null