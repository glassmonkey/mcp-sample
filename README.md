# MCP Sample Go Application

A simple Go application that prints "Hello, World!".

## Getting Started

### Prerequisites

- Go 1.21+
- Make

## 開発プロセス

このプロジェクトでは以下の開発ルールを採用しています：

1. **日本語での対話**
   - コメントやドキュメントは日本語で記述してください
   - コミットメッセージは日本語を基本としますが、英語も許容します

2. **Makefile経由での操作**
   - 直接`go`コマンドを実行せず、対応するMakeコマンドを使用してください
   - 例: `go build` → `make build`

3. **完了通知**
   - 各Makeコマンドは完了時に音声通知を行います
   - 例: `make build` → 「ビルドが完了しました」

4. **標準作業フロー**
   ```
   git pull
   # コード修正
   make fmt && make lint
   make test
   make build
   git commit -m "変更内容"
   git push
   ```

## Using the Makefile

This project includes a Makefile to simplify common operations.

### Available Commands

```bash
make           # Build the application (default)
make build     # Build the application
make run       # Run the application
make test      # Run tests
make fmt       # Format code
make lint      # Run linter (requires golint)
make vet       # Run go vet
make clean     # Clean build artifacts
make help      # Show help message
```

### Running without Make

If you prefer not to use Make, you can run Go commands directly:

```bash
# Build
go build -o bin/app

# Run
go run main.go

# Test
go test -v ./...
``` 