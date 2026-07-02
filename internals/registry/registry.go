package registry

import (
	dbtools "github.com/IceRain26/ToolHub/internals/tools/database"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "query",
		Description: "Execute a read-only SQL SELECT query.",
	}, dbtools.QueryTool)

	mcp.AddTool(server, &mcp.Tool{
		Name: "get_schema",
		Description: "Return the database schema",
	}, dbtools.GetSchemaTool)
}
