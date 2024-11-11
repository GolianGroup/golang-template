package entfx

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"master/internal/pkg/config"
	"master/internal/pkg/models"
)

var Module = fx.Provide(initEntClient)

func initEntClient(config *config.Application) (*models.Client, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database, config.DB.SslMode)
	fmt.Println(dsn)
	fmt.Println("Connecting to dsn:", dsn)

	client, err := models.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database, err: %v", err)
	}

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed running schema migration, err: %v", err)
	}

	return client, nil
}
