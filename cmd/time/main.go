package main

import (
	// "context"

	"github.com/SsrCoder/mcp/pkg/server"
	"github.com/SsrCoder/mcp/pkg/tools"
	// "github.com/mark3labs/mcp-go/mcp"
)

func main() {
	srv := server.NewMCPServer("time", "0.0.1").WithEndpoint("/time")

	srv.AddTool(tools.CurrentTimeTool)
	srv.AddTool(tools.TodayWeekdayTool)

	if err := srv.ListenAndServe(":8080"); err != nil {
		panic(err)
	}
}
