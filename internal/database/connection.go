package database

import (
	"context"
	"golang_template/internal/utils"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool   *pgxpool.Pool
	Client *ent.Client
}

func NewDatabase(ctx context.Context, config *utils.DatabaseConfig) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(config.GetDSN())
	if err != nil {
		return nil, err
	}

	// Configure pool
	poolConfig.MaxConns = config.MaxConns
	poolConfig.MinConns = config.MinConns

	// Create pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	// Create ent driver
	driver := sql.OpenDB(dialect.Postgres, sql.OpenDBFromPool(pool))

	// Create ent client
	client := ent.NewClient(ent.Driver(driver))

	return &Database{
		Pool:   pool,
		Client: client,
	}, nil
}

func (db *Database) Close() {
	db.Client.Close()
	db.Pool.Close()
}
