package tools

import (
	"context"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

var CurrentTimeTool = NewTool(
	"current_time",
	mcp.WithDescription("获取当前日期时间"),
).WithHandler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	now := time.Now()
	return mcp.NewToolResultText(now.Local().Format("2006-01-02 15:04:05")), nil
})

var TodayWeekdayTool = NewTool(
	"today_weekday",
	mcp.WithDescription("获取今天星期几"),
).WithHandler(func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	now := time.Now()
	return mcp.NewToolResultText(now.Weekday().String()), nil
})
