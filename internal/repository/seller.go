package repository

import (
	"context"
	"fmt"

	"github.com/awakedx/task/internal/domain"
	"github.com/google/uuid"
)

type SellerRepo struct {
	db *DB
}

func NewSellerRepo(db *DB) *SellerRepo {
	return &SellerRepo{db: db}
}

func (r *SellerRepo) Create(ctx context.Context, seller *domain.Seller) (uuid.UUID, error) {
	var id uuid.UUID
	query := `
        INSERT INTO sellers (name,phone)
        VALUES ($1,$2)
        RETURNING id
    `
	err := r.db.QueryRow(ctx, query,
		seller.Name,
		seller.Phone,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Failed to create item")
	}
	return id, nil
}
