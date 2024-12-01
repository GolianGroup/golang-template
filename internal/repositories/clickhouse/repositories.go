package clickhouse_repositories

import (
	"golang_template/internal/database/clickhouse"
)

type ClickhouseRepository interface {
}

// var (
// ErrGlobal = errors.New("some global error")
// )
type repository struct {
	db clickhouse.ClickhouseDatabase
}

func NewRepository(db clickhouse.ClickhouseDatabase) ClickhouseRepository {
	return &repository{db: db}
}
