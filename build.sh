#!/usr/bin/env bash
# ============================================================
# MergePDF - Build and Run Script for Linux
# Compatible with APT (Ubuntu/Debian) and DNF (Fedora/RHEL)
# ============================================================

set -e

echo "=== MergePDF Build Script (Linux) ==="

# --------------------------
# Step 1: Check if Go is installed
# --------------------------
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go 1.25.1..."

    if command -v apt &> /dev/null; then
        sudo apt update -y
        sudo apt install -y wget
        wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz -O /tmp/go.tar.gz
        sudo rm -rf /usr/local/go
        sudo tar -C /usr/local -xzf /tmp/go.tar.gz
        echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
        source ~/.bashrc
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y wget tar
        wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz -O /tmp/go.tar.gz
        sudo rm -rf /usr/local/go
        sudo tar -C /usr/local -xzf /tmp/go.tar.gz
        echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
        source ~/.bashrc
    else
        echo "Error: Neither APT nor DNF detected. Install Go manually."
        exit 1
    fi
else
    echo "Go detected: $(go version)"
fi

# --------------------------
# Step 2: Check for project files
# --------------------------
if [ ! -f "go.mod" ]; then
    echo "Error: go.mod not found. Run this script from the project root."
    exit 1
fi

# --------------------------
# Step 3: Install dependencies
# --------------------------
echo
echo "Installing dependencies..."
go mod tidy
echo "Dependencies installed successfully."

# --------------------------
# Step 4: Install system libraries (for Fyne GUI)
# --------------------------
echo
echo "Installing system dependencies for GUI..."

if command -v apt &> /dev/null; then
    sudo apt update -y
    sudo apt install -y libx11-dev libxcursor-dev libxrandr-dev \
        libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev \
        zenity
elif command -v dnf &> /dev/null; then
    sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel \
        libXinerama-devel libXi-devel mesa-libGL-devel libXxf86vm-devel \
        zenity
fi

# --------------------------
# Step 5: Build the binary
# --------------------------
echo
echo "Building MergePDF..."
go build -ldflags="-s -w" -o MergePDF .
echo "Build succeeded: ./MergePDF"

# --------------------------
# Step 6: Run the application
# --------------------------
echo
echo "Launching MergePDF..."
./MergePDF &
echo "=== Build & Launch complete ==="
