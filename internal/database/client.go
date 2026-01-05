package database

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/wonjinsin/go-boilerplate/internal/config"
	"github.com/wonjinsin/go-boilerplate/internal/repository/postgres/dao/ent"
)

// NewEntClient creates a new EntGo client with PostgreSQL connection
func NewEntClient(cfg *config.Config) (*ent.Client, error) {
	// Open database connection
	db, err := sql.Open("pgx", cfg.GetDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create ent client
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	return client, nil
}
