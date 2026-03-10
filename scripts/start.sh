#!/bin/bash

# Start script - runs the built binary
set -e

# Move to project root (one level up from scripts/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

APP_NAME="go-auth"
BINARY="$PROJECT_ROOT/bin/$APP_NAME"

echo ">>> Starting $APP_NAME..."
cd "$PROJECT_ROOT"

# Check if binary exists
if [ ! -f "$BINARY" ]; then
    echo "Error: Binary not found at bin/$APP_NAME"
    echo "Run scripts/build.sh first"
    exit 1
fi

# Run the binary
"$BINARY"

# Back to original directory
cd - > /dev/null