package app

import (
	"fmt"
	"golang_template/internal/database"
)

func (a *application) InitDatabase() (*database.Database, error) {
	db, err := database.NewDatabase(a.ctx, &a.config.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}
	return db, nil
}
