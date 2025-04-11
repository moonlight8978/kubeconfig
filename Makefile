.PHONY: build install clean

# Binary name
BINARY_NAME=kubeconfig

# Build the binary
build:
	go build -o bin/$(BINARY_NAME) bin/main.go

# Install the binary to $GOPATH/bin
install:
	go build -o $(GOPATH)/bin/$(BINARY_NAME) bin/main.go

# Clean build artifacts
clean:
	go clean
	rm -f bin/$(BINARY_NAME)

# Run tests
test:
	go test ./...
