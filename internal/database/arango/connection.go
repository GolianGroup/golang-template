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
	Database(ctx context.Context, dbName string) (arangodb.Database, error)
	CreateCollection(ctx context.Context, dbName string, colName string) (arangodb.Collection, error)
	Collection(ctx context.Context, dbName string, colName string) (arangodb.Collection, error)
}

type arangoDB struct {
	client arangodb.Client
}

func NewArangoDB(ctx context.Context, conf *config.ArangoConfig) (ArangoDB, error) {
	connStrs, err := config.GetArangoStrings(conf)
	if err != nil {
		return nil, err
	}
	endpoint := connection.NewRoundRobinEndpoints(connStrs)
	conn := connection.NewHttp2Connection(connection.DefaultHTTP2ConfigurationWrapper(endpoint /*InsecureSkipVerify*/, conf.InsecureSkipVerify))

	auth := connection.NewBasicAuth(conf.User, conf.Pass)
	err = conn.SetAuthentication(auth)
	if err != nil {
		return nil, err
	}

	client := arangodb.NewClient(conn)

	return &arangoDB{
		client: client,
	}, nil
}

func (a *arangoDB) Database(ctx context.Context, dbName string) (arangodb.Database, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dbExists, err := a.client.DatabaseExists(timeoutCtx, dbName)
	if err != nil {
		return nil, err
	}
	if !dbExists {
		return nil, fmt.Errorf("database %s does not exist", dbName)
	}

	db, err := a.client.Database(timeoutCtx, dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (a *arangoDB) CreateCollection(ctx context.Context, dbName string, colName string) (arangodb.Collection, error) {
	db, err := a.Database(ctx, dbName)
	if err != nil {
		return nil, err
	}

	exists, err := db.CollectionExists(ctx, colName)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("collection %s already exists", colName)
	}
	options := arangodb.CreateCollectionOptions{}
	properties := arangodb.CreateCollectionProperties{}

	collection, err := db.CreateCollectionWithOptions(ctx, colName, &properties, &options)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (a *arangoDB) Collection(ctx context.Context, dbName string, colName string) (arangodb.Collection, error) {
	db, err := a.Database(ctx, dbName)
	if err != nil {
		return nil, err
	}

	options := arangodb.GetCollectionOptions{}

	collection, err := db.GetCollection(ctx, colName, &options)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
