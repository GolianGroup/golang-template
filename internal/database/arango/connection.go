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
	AsyncDatabase(ctx context.Context, jobID string, dbName string) (arangodb.Database, error)
	AsyncJobResult(ctx context.Context, dbName string) (arangodb.Database, string, error)
}

type arangoDB struct {
	connection connection.Connection
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

	return &arangoDB{
		connection: conn,
	}, nil
}

func (a *arangoDB) syncClient() arangodb.Client {
	return arangodb.NewClient(a.connection)
}

func (a *arangoDB) asyncClient() arangodb.Client {
	conn := connection.NewConnectionAsyncWrapper(a.connection)
	return arangodb.NewClient(conn)
}

func (a *arangoDB) Database(ctx context.Context, dbName string) (arangodb.Database, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := a.syncClient()
	dbExists, err := client.DatabaseExists(timeoutCtx, dbName)
	if err != nil {
		return nil, err
	}
	if !dbExists {
		return nil, fmt.Errorf("database %s does not exist", dbName)
	}

	db, err := client.Database(timeoutCtx, dbName)
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

func (a *arangoDB) AsyncJobResult(ctx context.Context, dbName string) (arangodb.Database, string, error) {
	client := a.asyncClient()

	// Trigger async database check
	dbExists, errWithJobID := client.DatabaseExists(connection.WithAsync(ctx), dbName)
	if !dbExists {
		return nil, "", fmt.Errorf("database %s does not exist", dbName)
	}
	if errWithJobID == nil {
		return nil, "", fmt.Errorf("expected async job ID, got nil")
	}

	// Extract async job ID
	jobID, isAsync := connection.IsAsyncJobInProgress(errWithJobID)
	if !isAsync {
		return nil, "", fmt.Errorf("expected async job ID, got %v", jobID)
	}

	return nil, jobID, nil
}

func (a *arangoDB) AsyncDatabase(ctx context.Context, jobID string, dbName string) (arangodb.Database, error) {
	client := a.asyncClient()

	// Fetch async job result
	dbExists, err := client.DatabaseExists(connection.WithAsyncID(ctx, jobID), dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch async job result: %v", err)
	}
	if !dbExists {
		return nil, fmt.Errorf("database does not exist")
	}

	db, err := client.Database(ctx, dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch database: %v", err)
	}
	return db, nil
}
