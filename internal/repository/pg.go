package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/awakedx/task/internal/config"
	"github.com/awakedx/task/internal/repository/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func Init() (*DB, error) {
	cfg := config.Get()
	if cfg.DbURI == "" {
		return nil, fmt.Errorf("PgURL is empty")
	}
	fmt.Println(cfg.DbURI)
	pgDB, err := pgxpool.New(context.Background(), cfg.DbURI)
	if err != nil {
		fmt.Printf("unable to create connection pool:%v\n", err)
		return nil, fmt.Errorf("unable to create connection pool:%v", err)
	}

	if err = pgDB.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to database")
	}
	err = migrations.MigrationUp()
	slog.Info("Running migrations")
	if err != nil {
		return nil, err
	}
	return &DB{pgDB}, nil
}
