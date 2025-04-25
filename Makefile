# Makefile for Hello World Kubernetes Operator

.PHONY: build run deploy clean test help

# Build the operator
build:
	go build -o bin/hello-world-operator ./cmd

# Run the operator locally
run: build
	./bin/hello-world-operator

# Deploy the operator to the Kubernetes cluster
deploy:
	kubectl apply -f deploy

# Clean up the build artifacts
clean:
	rm -rf bin/hello-world-operator

# Test the operator
test:
	go test ./... -cover

# Help command
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the operator"
	@echo "  run     - Run the operator locally"
	@echo "  deploy   - Deploy the operator to the Kubernetes cluster"
	@echo "  clean   - Clean up build artifacts"
	@echo "  test    - Run tests"
	@echo "  help    - Show this help message"