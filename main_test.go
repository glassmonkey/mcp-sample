package main

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/mark3labs/mcp-go/mcp"
)

func TestGetCurrentTimeHandler(t *testing.T) {
	// テスト用の固定時刻を設定（2023年1月2日 15:04:05 JST）
	fixedTime := time.Date(2023, 1, 2, 15, 4, 5, 0, time.Local)
	restore := flextime.Fix(fixedTime)
	defer restore() // テスト終了後に元の時刻動作に戻す
	
	// コンテキストとリクエストの作成
	ctx := context.Background()
	request := mcp.CallToolRequest{}
	request.Params.Name = "get_current_time"
	request.Params.Arguments = map[string]interface{}{}
	
	// ハンドラを呼び出し
	result, err := getCurrentTimeHandler(ctx, request)
	
	// エラーがないことを確認
	if err != nil {
		t.Errorf("予期しないエラーが発生しました: %v", err)
	}
	
	// 結果がnilでないことを確認
	if result == nil {
		t.Fatal("ハンドラがnilを返しました")
	}
	
	// エラーでないことを確認
	if result.IsError {
		t.Fatal("ハンドラがエラー結果を返しました")
	}
	
	// コンテンツが存在することを確認
	if len(result.Content) == 0 {
		t.Fatal("コンテンツが空です")
	}
	
	// テキストコンテンツの取得と検証
	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatalf("コンテンツがTextContent型ではありませんでした: %T", result.Content[0])
	}
	
	// 期待する値が含まれていることを確認
	expectedValues := []string{
		"Current Time: 15:04:05",
		"Current Date: 2023-01-02",
		fixedTime.Location().String(),
	}
	
	for _, expected := range expectedValues {
		if !strings.Contains(textContent.Text, expected) {
			t.Errorf("返却されたテキストに期待値が含まれていません: %s", expected)
		}
	}
}

func TestGetTimezoneHandler(t *testing.T) {
	// テスト用の固定時刻を設定（2023年1月2日 15:04:05 JST）
	fixedTime := time.Date(2023, 1, 2, 15, 4, 5, 0, time.Local)
	restore := flextime.Fix(fixedTime)
	defer restore()
	
	// コンテキストとリクエストの作成（UTCタイムゾーン）
	ctx := context.Background()
	request := mcp.CallToolRequest{}
	request.Params.Name = "get_time_in_timezone"
	request.Params.Arguments = map[string]interface{}{
		"timezone": "UTC",
	}
	
	// ハンドラを呼び出し
	result, err := getTimezoneHandler(ctx, request)
	
	// エラーがないことを確認
	if err != nil {
		t.Errorf("予期しないエラーが発生しました: %v", err)
	}
	
	// 結果がnilでないことを確認
	if result == nil {
		t.Fatal("ハンドラがnilを返しました")
	}
	
	// エラーでないことを確認
	if result.IsError {
		t.Fatal("ハンドラがエラー結果を返しました")
	}
	
	// コンテンツが存在することを確認
	if len(result.Content) == 0 {
		t.Fatal("コンテンツが空です")
	}
	
	// テキストコンテンツの取得と検証
	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatalf("コンテンツがTextContent型ではありませんでした: %T", result.Content[0])
	}
	
	// UTCタイムゾーンになっていることを確認
	if !strings.Contains(textContent.Text, "Timezone: UTC") {
		t.Errorf("返却されたテキストにUTCタイムゾーンが含まれていません")
	}
	
	// 期待する文字列が含まれていることを確認
	expectedValues := []string{
		"Time in UTC",
		"Current Time:",
		"Current Date:",
		"Unix Timestamp:",
	}
	
	for _, expected := range expectedValues {
		if !strings.Contains(textContent.Text, expected) {
			t.Errorf("返却されたテキストに期待値が含まれていません: %s", expected)
		}
	}
}

func TestGetTimezoneHandlerInvalidTimezone(t *testing.T) {
	// コンテキストとリクエストの作成（無効なタイムゾーン）
	ctx := context.Background()
	request := mcp.CallToolRequest{}
	request.Params.Name = "get_time_in_timezone"
	request.Params.Arguments = map[string]interface{}{
		"timezone": "InvalidTimezone",
	}
	
	// ハンドラを呼び出し
	result, err := getTimezoneHandler(ctx, request)
	
	// エラーがないことを確認（MCPではエラーは結果のプロパティとして返すため）
	if err != nil {
		t.Errorf("予期しないエラーが発生しました: %v", err)
	}
	
	// 結果がnilでないことを確認
	if result == nil {
		t.Fatal("ハンドラがnilを返しました")
	}
	
	// エラー結果が返されることを確認
	if !result.IsError {
		t.Fatal("無効なタイムゾーンにもかかわらず、エラー結果が返されませんでした")
	}
	
	// コンテンツが存在することを確認
	if len(result.Content) == 0 {
		t.Fatal("コンテンツが空です")
	}
	
	// テキストコンテンツの取得と検証
	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatalf("コンテンツがTextContent型ではありませんでした: %T", result.Content[0])
	}
	
	// エラーメッセージにタイムゾーン名が含まれていることを確認
	if !strings.Contains(textContent.Text, "InvalidTimezone") {
		t.Errorf("エラーメッセージに無効なタイムゾーン名が含まれていません: %s", textContent.Text)
	}
}
