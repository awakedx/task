package migrations

import (
	"fmt"

	"github.com/awakedx/task/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationUp() error {
	cfg := config.Get()
	m, err := migrate.New(cfg.MigrationsPath, cfg.DbURI)
	if err != nil {
		fmt.Printf("failed migration init, %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
