package arango

import (
	"context"
	"fmt"
	"golang_template/internal/config"
	"log"
	"testing"
	"time"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/connection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestArangoDB(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "arangodb:latest",
		ExposedPorts: []string{"8529/tcp"},
		Env: map[string]string{
			"ARANGO_ROOT_PASSWORD": "rootpassword",
		},
		WaitingFor: wait.ForHTTP("/").WithPort("8529/tcp").WithStartupTimeout(2 * time.Minute),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start ArangoDB container: %v", err)
	}

	// Ensure the container is terminated after tests
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("Failed to terminate ArangoDB container: %v", err)
		}
	}()

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %v", err)
	}
	mappedPort, err := container.MappedPort(ctx, "8529")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %v", err)
	}

	// Prepare ArangoDB configuration
	arangoConf := &config.ArangoConfig{
		User:               "root",
		Pass:               "rootpassword",
		ConnStrs:           fmt.Sprintf("http://%s:%s", host, mappedPort.Port()),
		DBName:             "test",
		InsecureSkipVerify: true,
	}

	t.Run("Test arango db initialization", func(t *testing.T) {
		conn, err := NewArangoDB(ctx, arangoConf)
		require.NoError(t, err)

		assert.NotNil(t, conn)
	})

	t.Run("Test returning database", func(t *testing.T) {
		ctx := context.Background()
		conn, err := NewArangoDB(ctx, arangoConf)
		require.NoError(t, err)

		dbName := uniqueDatabaseName()
		_, err = createSyncDatabase(ctx, dbName, arangoConf.ConnStrs)
		require.NoError(t, err)

		db, err := conn.Database(ctx, dbName)
		require.NoError(t, err)

		assert.NotNil(t, db)
	})

	t.Run("Test create collection", func(t *testing.T) {
		ctx := context.Background()
		conn, err := NewArangoDB(ctx, arangoConf)
		require.NoError(t, err)

		dbName := uniqueDatabaseName()
		_, err = createSyncDatabase(ctx, dbName, arangoConf.ConnStrs)
		require.NoError(t, err)

		col, err := conn.CreateCollection(ctx, dbName, "testCollection")
		require.NoError(t, err)

		assert.NotNil(t, col)
	})

	t.Run("Test get collection", func(t *testing.T) {
		ctx := context.Background()
		conn, err := NewArangoDB(ctx, arangoConf)
		require.NoError(t, err)

		dbName := uniqueDatabaseName()
		db, err := createSyncDatabase(ctx, dbName, arangoConf.ConnStrs)
		require.NoError(t, err)

		db.CreateCollection(ctx, "testCollection", nil)

		col, err := conn.Collection(ctx, dbName, "testCollection")
		require.NoError(t, err)

		assert.NotNil(t, col)
	})
}

// helper functions
func createSyncDatabase(ctx context.Context, dbName string, connStr string) (arangodb.Database, error) {
	endpoint := connection.NewRoundRobinEndpoints([]string{connStr})
	conn := connection.NewHttp2Connection(connection.DefaultHTTP2ConfigurationWrapper(endpoint /*InsecureSkipVerify*/, true))
	err := conn.SetAuthentication(connection.NewBasicAuth("root", "rootpassword"))
	if err != nil {
		return nil, err
	}
	client := arangodb.NewClient(conn)
	db, err := client.CreateDatabase(ctx, dbName, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func uniqueDatabaseName() string {
	return fmt.Sprintf("test_db_%d", time.Now().UnixNano())
}
