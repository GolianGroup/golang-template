package arango

import (
	"context"
	"fmt"
	"golang_template/internal/config"
	"time"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/connection"
)

type ArangoDB interface {
	CreateCollection(ctx context.Context, colName string) (arangodb.Collection, error)
	OpenConnection(ctx context.Context, colName string) (arangodb.Collection, error)
}

type arangoDB struct {
	db arangodb.Database
}

func NewArangoDB(ctx context.Context, conf *config.ArangoConfig) (ArangoDB, error) {
	connStr := config.GetArangoStr(conf)
	endpoint := connection.NewRoundRobinEndpoints([]string{connStr})
	conn := connection.NewHttp2Connection(connection.DefaultHTTP2ConfigurationWrapper(endpoint /*InsecureSkipVerify*/, true))

	auth := connection.NewBasicAuth(conf.User, conf.Pass)
	err := conn.SetAuthentication(auth)
	if err != nil {
		return nil, err
	}

	client := arangodb.NewClient(conn)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dbExists, err := client.DatabaseExists(ctx, conf.DBName)
	if err != nil {
		return nil, err
	}
	if !dbExists {
		return nil, fmt.Errorf("database %s does not exist", conf.DBName)
	}

	database, err := client.Database(timeoutCtx, conf.DBName)
	if err != nil {
		return nil, err
	}

	return &arangoDB{
		db: database,
	}, nil
}

func (a *arangoDB) Database() arangodb.Database {
	return a.db
}

func (a *arangoDB) CreateCollection(ctx context.Context, colName string) (arangodb.Collection, error) {
	exists, err := a.db.CollectionExists(ctx, colName)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("collection %s already exists", colName)
	}
	collection, err := a.db.CreateCollection(ctx, colName, nil)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (a *arangoDB) OpenConnection(ctx context.Context, colName string) (arangodb.Collection, error) {
	connection, err := a.db.Collection(ctx, colName)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
