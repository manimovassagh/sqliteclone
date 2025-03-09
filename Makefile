# The name of the binary to create
BINARY_NAME = GoSQLLite

# The directory where the compiled binary will be placed
BIN_DIR = bin

# Default target, builds and runs the application
run: build
	$(BIN_DIR)/$(BINARY_NAME)

# Build the application
build:
	# Create the bin/ directory if it doesn't exist
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME)

# Clean the build
clean:
	rm -rf $(BIN_DIR)

# Run the application without building (useful for development)
run_without_build:
	go run .

# Test the application (if you have test cases set up)
test:
	go test -v ./...

# Install the GoSQLLite binary globally (optional)
install:
	go install