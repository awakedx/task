package repository

import (
	"context"

	"github.com/awakedx/task/internal/domain"
)

type OrderRepo struct {
	db *DB
}

func NewOrderRepo(db *DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Create(ctx context.Context, orderD domain.Order) {
	return
}
