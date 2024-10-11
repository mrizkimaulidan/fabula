# Makefile for Fabula project

# Variables
BINARY_NAME=fabula
BIN_DIR=bin
STORIES_DIR=stories
HIGHLIGHTS_DIR=highlights

# Platforms
WINDOWS_AMD64=$(BIN_DIR)/${BINARY_NAME}64-windows.exe
WINDOWS_386=$(BIN_DIR)/${BINARY_NAME}32-windows.exe
MACOS=$(BIN_DIR)/${BINARY_NAME}64-macos
LINUX=$(BIN_DIR)/${BINARY_NAME}64-linux

# Check if Go is installed
check-go:
	@command -v go >/dev/null 2>&1 || { echo "Go is not installed. Please install Go to proceed." && exit 1; }

# Build for multiple platforms
build: check-go
	@echo "Building for Windows 64-bit..."
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_AMD64) *.go
	@echo "Building for Windows 32-bit..."
	GOOS=windows GOARCH=386 go build -o $(WINDOWS_386) *.go
	@echo "Building for MacOS 64-bit..."
	GOOS=darwin GOARCH=amd64 go build -o $(MACOS) *.go
	@echo "Building for Linux 64-bit..."
	GOOS=linux GOARCH=amd64 go build -o $(LINUX) *.go
	@echo "Build completed successfully!"

# Run for downloading stories
story: check-go
	@echo "Downloading stories for user: ${USERNAME}..."
	go build && ./$(BINARY_NAME) -username=$(USERNAME) -option=story

# Run for downloading highlights
highlight: check-go
	@echo "Downloading highlights for user: ${USERNAME}..."
	go build && ./$(BINARY_NAME) -username=$(USERNAME) -option=highlight

# Clean up build files and directories
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	rm -rf $(STORIES_DIR)
	rm -rf $(HIGHLIGHTS_DIR)
	rm -f $(BINARY_NAME)
	@echo "Cleanup completed!"

.PHONY: build story highlight clean check-go
