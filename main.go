package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Songmu/flextime"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// TimeResponse は時刻情報のレスポンス構造体
type TimeResponse struct {
	CurrentTime   string `json:"current_time"`
	CurrentDate   string `json:"current_date"`
	UnixTimestamp int64  `json:"unix_timestamp"`
	TimeZone      string `json:"timezone"`
}

func main() {
	// MCPサーバーの作成
	s := server.NewMCPServer(
		"Time Service 🕒",
		"1.0.0",
	)

	// 現在時刻を取得するツール
	currentTimeTool := mcp.NewTool("get_current_time",
		mcp.WithDescription("Get current time information"),
	)

	// 特定のタイムゾーンの時刻を取得するツール
	timezoneTool := mcp.NewTool("get_time_in_timezone",
		mcp.WithDescription("Get time information in a specific timezone"),
		mcp.WithString("timezone",
			mcp.Required(),
			mcp.Description("Name of the timezone (e.g., 'America/New_York', 'Europe/London', 'Asia/Tokyo')"),
		),
	)

	// ツールハンドラーを登録
	s.AddTool(currentTimeTool, getCurrentTimeHandler)
	s.AddTool(timezoneTool, getTimezoneHandler)

	// stdio形式でサーバーを起動
	fmt.Println("Starting Time Service MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// getCurrentTimeHandler は現在時刻情報を返すツールハンドラー
func getCurrentTimeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 現在時刻を取得（flextime使用）
	now := flextime.Now()
	
	// レスポンス作成
	timeResponse := TimeResponse{
		CurrentTime:   now.Format("15:04:05"),
		CurrentDate:   now.Format("2006-01-02"),
		UnixTimestamp: now.Unix(),
		TimeZone:      now.Location().String(),
	}

	// 結果を整形して返す
	result := fmt.Sprintf(
		"Current Time: %s\nCurrent Date: %s\nUnix Timestamp: %d\nTimezone: %s",
		timeResponse.CurrentTime,
		timeResponse.CurrentDate,
		timeResponse.UnixTimestamp,
		timeResponse.TimeZone,
	)

	return mcp.NewToolResultText(result), nil
}

// getTimezoneHandler は指定されたタイムゾーンでの時刻情報を返すツールハンドラー
func getTimezoneHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// リクエストからタイムゾーンパラメータを取得
	tzName, ok := request.Params.Arguments["timezone"].(string)
	if !ok {
		return nil, errors.New("timezone must be a string")
	}
	
	// タイムゾーンの検証
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid timezone: %s", tzName)), nil
	}
	
	// 現在時刻を取得し、指定されたタイムゾーンに変換
	now := flextime.Now().In(loc)
	
	// レスポンス作成
	timeResponse := TimeResponse{
		CurrentTime:   now.Format("15:04:05"),
		CurrentDate:   now.Format("2006-01-02"),
		UnixTimestamp: now.Unix(),
		TimeZone:      now.Location().String(),
	}

	// 結果を整形して返す
	result := fmt.Sprintf(
		"Time in %s:\nCurrent Time: %s\nCurrent Date: %s\nUnix Timestamp: %d\nTimezone: %s",
		tzName,
		timeResponse.CurrentTime,
		timeResponse.CurrentDate,
		timeResponse.UnixTimestamp,
		timeResponse.TimeZone,
	)

	return mcp.NewToolResultText(result), nil
}
