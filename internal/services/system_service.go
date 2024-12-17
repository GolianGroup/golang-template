package services

import (
	"context"
	"golang_template/internal/database/arango"
	"golang_template/internal/database/clickhouse"
	"golang_template/internal/database/postgres"
	"golang_template/internal/producers"
)

type RepoStatus struct {
	Healthy bool
	Error   error
}

type SystemService interface {
	ReadyCheck(ctx context.Context) map[string]RepoStatus
}

type systemService struct {
	postgresClient 	 postgres.Database
	clickhouseClient clickhouse.ClickhouseDatabase
	arangoClient     arango.ArangoDB
	redisClient     producers.RedisClient
}

func NewSystemService(postgresClient postgres.Database, clickhouseClient clickhouse.ClickhouseDatabase, arangoClient arango.ArangoDB, redisClient producers.RedisClient) SystemService {
	return &systemService{postgresClient: postgresClient, clickhouseClient: clickhouseClient, arangoClient: arangoClient, redisClient: redisClient}
}

func (s *systemService) ReadyCheck(ctx context.Context) map[string]RepoStatus {

	statuses := make(map[string]RepoStatus)

	// Check Postgres
	if err := s.postgresClient.DB().Ping(); err != nil {
		statuses["postgres"] = RepoStatus{Healthy: false, Error: err}
	} else {
		statuses["postgres"] = RepoStatus{Healthy: true, Error: nil}
	}

	// Check Clickhouse
	if err := s.clickhouseClient.Ping(ctx); err != nil {
		statuses["clickhouse"] = RepoStatus{Healthy: false, Error: err}
	} else {
		statuses["clickhouse"] = RepoStatus{Healthy: true, Error: nil}
	}

	// Check ArangoDB
	if err := s.arangoClient.Ping(ctx); err != nil {
		statuses["arango"] = RepoStatus{Healthy: false, Error: err}
	} else {
		statuses["arango"] = RepoStatus{Healthy: true, Error: nil}
	}

	// Check Redis
	if err := s.redisClient.RedisStorage().Conn().Ping(ctx).Err(); err != nil {
		statuses["redis"] = RepoStatus{Healthy: false, Error: err}
	} else {
		statuses["redis"] = RepoStatus{Healthy: true, Error: nil}
	}

	return statuses

}
