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

type Database interface {
	Close()
	EntClient() *ent.Client
	DB() *dbsql.DB
}

type database struct {
	pool     *pgxpool.Pool
	database *dbsql.DB
	client   *ent.Client
}

func NewDatabase(ctx context.Context, config *utils.DatabaseConfig) (Database, error) {
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

	return &database{
		pool:     pool,
		database: db,
		client:   client,
	}, nil
}

func (db *database) Close() {
	db.client.Close()
	db.pool.Close()
}

func (db *database) EntClient() *ent.Client {
	return db.client
}

func (db *database) DB() *dbsql.DB {
	return db.database
}
