# Project variables
PROJECT_NAME := go-dummy-monitor
BINARY := $(PROJECT_NAME)
GO ?= go
GOLINT ?= golangci-lint
DLV ?= dlv
RM ?= rm -f
RMDIR ?= rm -rf
MKDIR ?= mkdir -p

# Build variables
BUILD_DIR ?= build
DIST_DIR ?= dist
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LDFLAGS ?= -w -s

# Check for required tools
HAS_GOLINT := $(shell command -v $(GOLINT) 2> /dev/null)
HAS_DLV := $(shell command -v $(DLV) 2> /dev/null)

# Declare all phony targets
.PHONY: all clean test lint fmt build build-all debug help

# Default target
all: clean test lint fmt build

clean:
	@echo "Removing the build directory, and all binaries, and logs, and coverage files..."
	@$(RMDIR) $(BUILD_DIR) $(DIST_DIR)
	@$(RM) $(BINARY)
	@$(RM) *.log
	@$(RM) *.out
	@$(RM) coverage.txt

# Clean target - removes build artifacts and binaries
clean-all: clean
	@echo "Deep cleaning the project - removing cache, test cache, and module cache..."
	@$(GO) clean -cache -testcache -modcache

# Test target - runs all tests with coverage
test:
	@echo "Running tests..."
	@$(GO) test -v -coverprofile=coverage.txt -covermode=atomic ./...

# Lint target - runs golangci-lint if available
lint:
	@echo "Running linter..."
ifdef HAS_GOLINT
	@$(GOLINT) run
else
	@echo "Warning: $(GOLINT) not found. Skipping linting."
	@echo "Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
endif

# Format target - formats all Go files
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...

# Build target - builds the binary for the current platform
build:
	@echo "Building for $(GOOS)/$(GOARCH)..."
	@$(GO) mod tidy
	@$(MKDIR) $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY) -ldflags="$(LDFLAGS)" .

# Build all target - builds binaries for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@$(MKDIR) $(DIST_DIR)
	@GOOS=darwin GOARCH=amd64 $(GO) build -o $(DIST_DIR)/$(BINARY)-darwin-amd64 -ldflags="$(LDFLAGS)" .
	@GOOS=darwin GOARCH=arm64 $(GO) build -o $(DIST_DIR)/$(BINARY)-darwin-arm64 -ldflags="$(LDFLAGS)" .
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(DIST_DIR)/$(BINARY)-linux-amd64 -ldflags="$(LDFLAGS)" .
	@GOOS=linux GOARCH=arm64 $(GO) build -o $(DIST_DIR)/$(BINARY)-linux-arm64 -ldflags="$(LDFLAGS)" .
	@GOOS=windows GOARCH=amd64 $(GO) build -o $(DIST_DIR)/$(BINARY)-windows-amd64.exe -ldflags="$(LDFLAGS)" .

# Debug target - runs the binary in the debugger if available
debug:
	@echo "Starting debugger..."
ifdef HAS_DLV
	@$(DLV) debug .
else
	@echo "Error: $(DLV) not found. Cannot start debugger."
	@echo "Install it with: go install github.com/go-delve/delve/cmd/dlv@latest"
	@exit 1
endif

# Help target - shows available targets
help:
	@echo "Available targets:"
	@echo "  all         - Clean, test, lint, format, and build"
	@echo "  clean       - Remove build artifacts and binaries"
	@echo "  test        - Run tests with coverage"
	@echo "  lint        - Run golangci-lint"
	@echo "  fmt         - Format all Go files"
	@echo "  build       - Build for current platform"
	@echo "  build-all   - Build for multiple platforms"
	@echo "  debug       - Run in debugger"
	@echo "  help        - Show this help message"
	@echo
	@echo "Variables that can be overridden:"
	@echo "  GO          - Go compiler (default: go)"
	@echo "  GOLINT      - Go linter (default: golangci-lint)"
	@echo "  DLV         - Debugger (default: dlv)"
	@echo "  BUILD_DIR   - Build directory (default: build)"
	@echo "  DIST_DIR    - Distribution directory (default: dist)"
	@echo "  GOOS        - Target OS (default: current OS)"
	@echo "  GOARCH      - Target architecture (default: current arch)"
	@echo "  LDFLAGS     - Linker flags (default: -w -s)" 