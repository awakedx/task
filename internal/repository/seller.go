package repository

import (
	"context"
	"errors"

	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
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
	var pgErr *pgconn.PgError
	if err != nil && errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return uuid.Nil, &utils.CustomErr{
				Msg:   "Seller with this phone already existing",
				Cause: utils.BadRequestErr,
			}
		}
	}
	return id, nil
}
