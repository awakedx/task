package repository

import (
	"context"
	"fmt"
	"strings"

	common "github.com/awakedx/task/internal/common/update"
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

func (r *CustomerRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	query := `SELECT name,phone FROM customers WHERE id=$1`
	var customer domain.Customer

	err := r.db.QueryRow(ctx, query, id).Scan(&customer.Name, &customer.Phone)
	if err != nil {
		return nil, err
	}
	customer.Id = id
	return &customer, nil
}

func (r *CustomerRepo) Update(ctx context.Context, updateCustomer *common.UpdateCustomer) error {
	queryParts := []string{}
	i := 1
	args := make([]any, 0, 3)

	if updateCustomer.Name != nil {
		queryParts = append(queryParts, fmt.Sprintf("name=$%d", i))
		args = append(args, updateCustomer.Name)
		i++
	}

	if updateCustomer.Phone != nil {
		queryParts = append(queryParts, fmt.Sprintf("phone=$%d", i))
		args = append(args, updateCustomer.Phone)
		i++
	}
	if len(args) == 0 {
		return nil
	}

	query := fmt.Sprintf("UPDATE customers SET %s WHERE id=$%d", strings.Join(queryParts, ", "), i)
	args = append(args, updateCustomer.Id)
	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM customers WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
