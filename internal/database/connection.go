package database

import (
	"context"
	"golang_template/internal/ent"
	"golang_template/internal/utils"

	dbsql "database/sql"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Pool   *pgxpool.Pool
	DB     *dbsql.DB
	Client *ent.Client
}

func NewDatabase(ctx context.Context, config *utils.DatabaseConfig) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(utils.GetDSN(config))
	if err != nil {
		return nil, err
	}

	// Configure pool
	poolConfig.MaxConns = int32(config.MaxConns)
	poolConfig.MinConns = int32(config.MinConns)

	// Create pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(*pool.Config().ConnConfig)

	// Create ent driver
	driver := sql.OpenDB(dialect.Postgres, db)

	// Create ent client
	client := ent.NewClient(ent.Driver(driver))

	return &Database{
		Pool:   pool,
		DB:     db,
		Client: client,
	}, nil
}

func (db *Database) Close() {
	db.Client.Close()
	db.Pool.Close()
}
