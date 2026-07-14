#!/bin/bash
# compile.sh - ELA full one-time build tool
# Usage: ./compile.sh

set -e

# ============================================================
# Colors
# ============================================================
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# ============================================================
# Header
# ============================================================
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  ELA - Efficient Linux Launcher${NC}"
echo -e "${BLUE}  Full One-Time Build Tool${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# ============================================================
# 1. Clean old builds
# ============================================================
echo -e "${YELLOW}[1/7] Cleaning old builds...${NC}"
rm -rf build/ bin/
mkdir -p build bin
echo -e "${GREEN}  Done.${NC}"

# ============================================================
# 2. Detect system info
# ============================================================
echo -e "${YELLOW}[2/7] Detecting system...${NC}"
GO_VERSION=$(go version | awk '{print $3}')
OS=$(go env GOOS)
ARCH=$(go env GOARCH)
echo "  Go: $GO_VERSION"
echo "  OS: $OS"
echo "  Arch: $ARCH"

# ============================================================
# 3. Check dependencies
# ============================================================
echo -e "${YELLOW}[3/7] Checking dependencies...${NC}"
DEPS_OK=1

if ! command -v go &> /dev/null; then
    echo -e "${RED}  Error: Go not found${NC}"
    DEPS_OK=0
else
    echo -e "${GREEN}  Go: OK${NC}"
fi

if ! command -v gcc &> /dev/null; then
    echo -e "${YELLOW}  Warning: gcc not found (C/C++ build will skip)${NC}"
else
    echo -e "${GREEN}  gcc: OK${NC}"
fi

if ! command -v cmake &> /dev/null; then
    echo -e "${YELLOW}  Warning: cmake not found (C/C++ build will skip)${NC}"
else
    echo -e "${GREEN}  cmake: OK${NC}"
fi

if [ $DEPS_OK -eq 0 ]; then
    echo -e "${RED}  Error: Missing required dependencies${NC}"
    exit 1
fi
echo ""

# ============================================================
# 4. Set Go proxy
# ============================================================
echo -e "${YELLOW}[4/7] Setting Go proxy...${NC}"
go env -w GOPROXY=https://goproxy.cn,direct
echo -e "${GREEN}  GOPROXY set to https://goproxy.cn,direct${NC}"

# ============================================================
# 5. Download Go dependencies
# ============================================================
echo -e "${YELLOW}[5/7] Downloading Go dependencies...${NC}"
go mod download
go mod tidy
echo -e "${GREEN}  Dependencies downloaded.${NC}"

# ============================================================
# 6. Build C/C++ components (optional)
# ============================================================
echo -e "${YELLOW}[6/7] Building C/C++ components...${NC}"
if command -v cmake &> /dev/null && command -v gcc &> /dev/null; then
    cd build
    cmake .. -DCMAKE_BUILD_TYPE=Release > /dev/null 2>&1
    make -j$(nproc) > /dev/null 2>&1
    cd ..
    echo -e "${GREEN}  C/C++ build complete.${NC}"
else
    echo -e "${YELLOW}  Skipping C/C++ build (cmake or gcc not found)${NC}"
fi

# ============================================================
# 7. Build Go main program
# ============================================================
echo -e "${YELLOW}[7/7] Building Go main program...${NC}"
CGO_ENABLED=1 go build -o bin/ela-gui cmd/ela/main.go
echo -e "${GREEN}  Go build complete.${NC}"

# ============================================================
# Done
# ============================================================
echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Build complete!${NC}"
echo -e "${GREEN}  Output: bin/ela-gui${NC}"
echo -e "${GREEN}  Run: ./bin/ela-gui${NC}"
echo -e "${GREEN}========================================${NC}"
