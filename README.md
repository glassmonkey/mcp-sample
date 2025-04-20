# 時計サービス MCP サンプル

Model Context Protocol (MCP) を使用した時計サービスのサンプル実装です。

## サービス機能

このサービスは以下の機能を提供します：

- 現在時刻の取得 (MCP `get_current_time` ツール)
  - 時刻（HH:MM:SS）
  - 日付（YYYY-MM-DD）
  - UNIXタイムスタンプ
  - タイムゾーン情報
- 特定タイムゾーンの時刻取得 (MCP `get_time_in_timezone` ツール)
  - 指定したタイムゾーンでの時刻情報を取得

## インターフェース

### ツール: get_current_time

**説明:** 現在の時刻情報を取得します

**パラメータ:** なし

**レスポンス例:**

```
Current Time: 14:30:45
Current Date: 2023-04-15
Unix Timestamp: 1681539045
Timezone: Asia/Tokyo
```

### ツール: get_time_in_timezone

**説明:** 指定されたタイムゾーンでの時刻情報を取得します

**パラメータ:** 
- `timezone` (必須): タイムゾーン（例: "UTC", "America/New_York", "Asia/Tokyo" など）

**レスポンス例:**

```
Time in Europe/London:
Current Time: 05:30:45
Current Date: 2023-04-15
Unix Timestamp: 1681539045
Timezone: Europe/London
```

## 開発環境

### 前提条件

- Go 1.21+
- Make

### MCP について

このサンプルは [Model Context Protocol (MCP)](https://github.com/mark3labs/mcp-go) を使用しています。MCPは、LLMアプリケーションと外部データソースやツールとのシームレスな統合を可能にするプロトコルです。

### ローカルでの実行

```bash
make run
```

### テスト実行

```bash
make test
```

## 特徴

- **MCPプロトコル対応**: LLMとのシームレスな統合
- **Flextime採用**: テストでの時刻固定対応
- **タイムゾーン対応**: 様々なタイムゾーン指定が可能

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

### MCPサーバーの実行

```bash
make run
```

または直接実行する場合：

```bash
go run main.go
```

MCPサーバーはstdio経由で通信します。これにより、LLMアプリケーションやMCPクライアントとシームレスに連携できます。 