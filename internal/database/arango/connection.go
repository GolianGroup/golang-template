package arango

import (
	"context"
	"golang_template/internal/config"
	"log"
	"time"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/connection"
)

type ArangoDB interface {
	Database(ctx context.Context) arangodb.Database
	GetCollection(ctx context.Context, name string) (arangodb.Collection, error)
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

	var db arangodb.Database
	if !dbExists {
		log.Println("Database does not exist, creating...")
		db, err = client.CreateDatabase(timeoutCtx, conf.DBName, nil)
		if err != nil {
			return nil, err
		}
	} else {
		db, err = client.Database(timeoutCtx, conf.DBName)
		if err != nil {
			return nil, err
		}
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

func (a *arangoDB) GetCollection(ctx context.Context, name string) (arangodb.Collection, error) {
	options := arangodb.GetCollectionOptions{
		SkipExistCheck: false,
	}

	return a.database.GetCollection(ctx, name, &options)
}
