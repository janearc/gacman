# Variables
SERVER_BINARY_NAME=gacman
CLIENT_BINARY_NAME=client
SOURCE_DIR=./
BUILD_DIR=./build
GO_FILES=$(shell find $(SOURCE_DIR) -name '*.go')

# Default target: build both the server and client binaries
all: build-server build-client

# Build the server binary
build-server:
	@echo "Building the server..."
	@go build -o $(BUILD_DIR)/$(SERVER_BINARY_NAME) ./server

# Build the client binary
build-client:
	@echo "Building the client..."
	@go build -o $(BUILD_DIR)/$(CLIENT_BINARY_NAME) ./client

# Clean up the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Help
help:
	@echo "Makefile commands:"
	@echo "  build    - Build the Go daemon binary."
	@echo "  run      - Build and run the Go daemon."
	@echo "  clean    - Remove the build directory."
	@echo "  deps     - Install project dependencies."
	@echo "  help     - Display this verreh halpful message."

halp: help
