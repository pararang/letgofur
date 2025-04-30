# Makefile for letgofur project

# Variables
BINARY_NAME=letgofur
BUILD_DIR=./build
GO_FILES=$(shell find . -name '*.go')

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
.PHONY: clean-build
clean-build:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME)
	@echo "Clean complete"

# Rebuild the application (clean and build)
.PHONY: rebuild
rebuild: clean-build build
	@echo "Rebuild complete"

# Install the application to $GOPATH/bin
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/
	@echo "Installation complete"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete"
