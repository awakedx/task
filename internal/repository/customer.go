package repository

import (
	"context"

	"github.com/awakedx/task/internal/domain"
	"github.com/google/uuid"
)

type CustomerRepo struct {
	db *DB
}

func NewCustomerRepo(db *DB) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}
func (r *CustomerRepo) Create(ctx context.Context, newCustomer *domain.Customer) (uuid.UUID, error) {
	query := `INSERT INTO customers (name,phone) values ($1,$2) RETURNING id`
	var newId uuid.UUID
	err := r.db.QueryRow(ctx, query, newCustomer.Name, newCustomer.Phone).Scan(&newId)
	if err != nil {
		return uuid.Nil, err
	}
	return newId, nil
}
