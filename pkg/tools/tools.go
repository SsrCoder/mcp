package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	srv "github.com/mark3labs/mcp-go/server"
)

type Tool struct {
	Tool mcp.Tool
	Fn   srv.ToolHandlerFunc
}

func NewTool(name string, opts ...mcp.ToolOption) *Tool {
	return &Tool{
		Tool: mcp.NewTool(name, opts...),
	}
}

func (t *Tool) WithHandler(fn srv.ToolHandlerFunc) *Tool {
	t.Fn = fn
	return t
}
