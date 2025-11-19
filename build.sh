#!/bin/bash

# Build script for System Monitor application
# Builds for Windows and Linux

set -e  # Exit on error

APP_NAME="SystemMonitor"
VERSION="1.0.0"
BUILD_DIR="./build"
APP_ID="com.example.systemmonitor"

echo "================================"
echo "Building $APP_NAME v$VERSION"
echo "================================"

# Create build directory
mkdir -p "$BUILD_DIR"

# Check if fyne is installed
if ! command -v fyne &> /dev/null; then
    echo "Error: fyne command not found"
    echo "Installing fyne CLI tool..."
    go install fyne.io/fyne/v2/cmd/fyne@latest
fi

echo ""
echo "Building for Linux..."
fyne package -os linux -icon icon.png -name "$APP_NAME" -release 
if [ -f "${APP_NAME}.tar.xz" ]; then
    mv "${APP_NAME}.tar.xz" "$BUILD_DIR/${APP_NAME}-linux-v${VERSION}.tar.xz"
    echo "✓ Linux build complete: $BUILD_DIR/${APP_NAME}-linux-v${VERSION}.tar.xz"
else
    echo "✗ Linux build failed"
fi

echo ""
echo "================================"
echo "Build Summary"
echo "================================"
ls -lh "$BUILD_DIR"

echo ""
echo "Build complete! Files are in the '$BUILD_DIR' directory."