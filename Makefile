.PHONY: all build test clean fmt lint vet run help notify

# Default target
all: build

# Build the application
build:
	go build -o bin/app
	@make notify msg="ビルドが完了しました"

# Run the application
run: build
	./bin/app
	@make notify msg="アプリケーションの実行が完了しました"

# Clean build artifacts
clean:
	rm -rf bin/
	go clean
	@make notify msg="クリーンアップが完了しました"

# Run tests
test:
	go test -v ./...
	@make notify msg="テストが完了しました"

# Format code
fmt:
	go fmt ./...
	@make notify msg="コードのフォーマットが完了しました"

# Run linter
lint:
	@if command -v golint > /dev/null; then \
		golint ./...; \
	else \
		echo "golint not installed. Run: go install golang.org/x/lint/golint@latest"; \
	fi
	@make notify msg="リント検査が完了しました"

# Run vet
vet:
	go vet ./...
	@make notify msg="vet検査が完了しました"

# Notification
notify:
	@if [ -n "$(msg)" ]; then \
		say "$(msg)"; \
		echo "🔔 $(msg)"; \
	fi

# Help
help:
	@echo "Available targets:"
	@echo "  all    : Build the application (default)"
	@echo "  build  : Build the application"
	@echo "  run    : Run the application"
	@echo "  test   : Run tests"
	@echo "  fmt    : Format code"
	@echo "  lint   : Run linter"
	@echo "  vet    : Run vet"
	@echo "  clean  : Clean build artifacts"
	@echo "  help   : Show this help message" 