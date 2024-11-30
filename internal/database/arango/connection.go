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
	Database(ctx context.Context) arangodb.Database
	VideoCollection(ctx context.Context) (arangodb.Collection, error)
	ReadDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, result interface{}, opts *arangodb.CollectionDocumentReadOptions) (arangodb.DocumentMeta, error)
	CreateDocumentWithOptions(ctx context.Context, collection arangodb.Collection, document interface{}, opts *arangodb.CollectionDocumentCreateOptions) (arangodb.CollectionDocumentCreateResponse, error)
	DeleteDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, opts *arangodb.CollectionDocumentDeleteOptions) error
	UpdateDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, value interface{}, opts *arangodb.CollectionDocumentUpdateOptions) (arangodb.CollectionDocumentUpdateResponse, error)
	Query(ctx context.Context, queryString string, queryOpts arangodb.QueryOptions) (arangodb.Cursor, error)
	CloseCursor(cursor arangodb.Cursor) error
	CursorReadDocument(ctx context.Context, cursor arangodb.Cursor, result interface{}) (arangodb.DocumentMeta, error)
	CursorHasMore(cursor arangodb.Cursor) bool
}

type arangoDB struct {
	database arangodb.Database
	client   arangodb.Client
	config   *config.ArangoConfig
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

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dbExists, err := client.DatabaseExists(timeoutCtx, conf.DBName)
	if err != nil {
		return nil, err
	}
	if !dbExists {
		return nil, fmt.Errorf("database %s does not exist", conf.DBName)
	}

	db, err := client.Database(timeoutCtx, conf.DBName)
	if err != nil {
		return nil, err
	}

	return &arangoDB{
		database: db,
		client:   client,
		config:   conf,
	}, nil
}

func (a *arangoDB) Database(ctx context.Context) arangodb.Database {
	return a.database
}

func (a *arangoDB) VideoCollection(ctx context.Context) (arangodb.Collection, error) {
	options := arangodb.GetCollectionOptions{
		SkipExistCheck: false,
	}

	videoCollection, err := a.database.GetCollection(ctx, "videos_collection", &options)
	if err != nil {
		return nil, err
	}

	return videoCollection, nil
}

func (a *arangoDB) ReadDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, result interface{}, opts *arangodb.CollectionDocumentReadOptions) (arangodb.DocumentMeta, error) {
	return collection.ReadDocumentWithOptions(ctx, key, result, opts)
}

func (a *arangoDB) CreateDocumentWithOptions(ctx context.Context, collection arangodb.Collection, document interface{}, opts *arangodb.CollectionDocumentCreateOptions) (arangodb.CollectionDocumentCreateResponse, error) {
	return collection.CreateDocumentWithOptions(ctx, document, opts)
}

func (a *arangoDB) DeleteDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, opts *arangodb.CollectionDocumentDeleteOptions) error {
	_, err := collection.DeleteDocumentWithOptions(ctx, key, opts)
	if err != nil {
		return err
	}
	return nil
}

func (a *arangoDB) UpdateDocumentWithOptions(ctx context.Context, collection arangodb.Collection, key string, value interface{}, opts *arangodb.CollectionDocumentUpdateOptions) (arangodb.CollectionDocumentUpdateResponse, error) {
	return collection.UpdateDocumentWithOptions(ctx, key, value, opts)
}

func (a *arangoDB) Query(ctx context.Context, queryString string, queryOpts arangodb.QueryOptions) (arangodb.Cursor, error) {
	return a.database.Query(ctx, queryString, &queryOpts)
}

func (a *arangoDB) CloseCursor(cursor arangodb.Cursor) error {
	return cursor.Close()
}

func (a *arangoDB) CursorReadDocument(ctx context.Context, cursor arangodb.Cursor, result interface{}) (arangodb.DocumentMeta, error) {
	return cursor.ReadDocument(ctx, result)
}

func (a *arangoDB) CursorHasMore(cursor arangodb.Cursor) bool {
	return cursor.HasMore()
}
