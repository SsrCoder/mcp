package server

import (
	"log"

	"github.com/SsrCoder/mcp/pkg/tools"
	srv "github.com/mark3labs/mcp-go/server"
)

type MCPServer struct {
	server   *srv.MCPServer
	endpoint string
}

func NewMCPServer(name string, version string) *MCPServer {
	log.Println("New MCP server:", name, "version:", version)

	return &MCPServer{
		server:   srv.NewMCPServer(name, version),
		endpoint: "/sse",
	}
}

func (s *MCPServer) AddTool(tool *tools.Tool) {
	s.server.AddTool(tool.Tool, tool.Fn)
	log.Println("Added tool", tool.Tool.Name, ":", tool.Tool.Description)
}

func (s *MCPServer) WithEndpoint(ep string) *MCPServer {
	s.endpoint = ep
	return s
}

func (s *MCPServer) ListenAndServe(addr string) error {
	log.Println("Server starting on", addr, "with endpoint:", s.endpoint)
	return srv.NewSSEServer(s.server,
		srv.WithSSEEndpoint(s.endpoint),
		srv.WithMessageEndpoint(s.endpoint+"/message"),
	).Start(addr)
}
