package database

import (
	"context"
	"fmt"
	"log"

	"github.com/githubdmn/codego/ent"
	"github.com/githubdmn/codego/internal/config"
)

func NewClient(cfg *config.DBConfig) (*ent.Client, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	client, err := ent.Open(cfg.Driver, connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	// Run the auto-migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	return client, nil
}
