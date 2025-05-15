package main

import (
	"github.com/gin-gonic/gin"
	"github.com/githubdmn/codego/internal/api/route"
	"github.com/githubdmn/codego/internal/config"
	"github.com/githubdmn/codego/pkg/database"
	"log"
)

// @title Go REST API
// @version 1.0
// @description This is a sample REST API server.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Initialize database
	client, err := database.NewClient(&cfg.DB)
	if err != nil {
		log.Fatal("cannot create db client:", err)
	}
	defer client.Close()

	// Initialize Gin
	r := gin.Default()

	// Setup routes
	route.Setup(r, client)

	// Start server
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
