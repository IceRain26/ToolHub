package main

import (
	"context"

	"log"

	"github.com/IceRain26/ToolHub/internals/config"
	"github.com/IceRain26/ToolHub/internals/database"
	"github.com/IceRain26/ToolHub/internals/registry"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {

	// Loading Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := database.Connect(cfg.Database); err != nil {
		log.Fatal(err)
	}

	// Creating MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    cfg.Server.Name,
		Version: cfg.Server.Version,
	}, nil)

	registry.Register(server)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}

}
