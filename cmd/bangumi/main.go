package main

import (
	// "context"

	"github.com/SsrCoder/mcp/pkg/server"
	"github.com/SsrCoder/mcp/pkg/tools"
	// "github.com/mark3labs/mcp-go/mcp"
)

func main() {
	srv := server.NewMCPServer("bangumi", "0.0.1").WithEndpoint("/bangumi")

	srv.AddTool(tools.CalendarTool)
	srv.AddTool(tools.CalendarFullTool)

	if err := srv.ListenAndServe(":8080"); err != nil {
		panic(err)
	}
}
