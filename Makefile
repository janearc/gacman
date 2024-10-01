# Variables
BINARY_NAME=gacman
SOURCE_DIR=./
BUILD_DIR=./build
GO_FILES=$(shell find $(SOURCE_DIR) -name '*.go')

# Default target: build the binary
build: $(GO_FILES)
	@echo "Summoning Gacman..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Run the binary
run: build
	@echo "Running Gacman..."
	@$(BUILD_DIR)/$(BINARY_NAME)

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