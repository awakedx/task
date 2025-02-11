package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	common "github.com/awakedx/task/internal/common/update"
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
	if err != nil {
		return uuid.Nil, &utils.CustomErr{
			Msg:   err.Error(),
			Cause: utils.InternalError,
		}
	}
	return id, nil
}

func (r *SellerRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Seller, error) {
	query := `SELECT id,name,phone FROM sellers WHERE id=$1`
	var seller domain.Seller
	err := r.db.QueryRow(ctx, query, id).Scan(&seller.Id, &seller.Name, &seller.Phone)
	if err != nil {
		return nil, fmt.Errorf("SellerRepo Get,%w", err)
	}
	return &seller, nil
}

func (r *SellerRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM sellers WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("SellerRepo Delete,%w", err)
	}
	return nil
}

func (r *SellerRepo) Update(ctx context.Context, updateSeller *common.UpdateSeller) error {
	queryParts := []string{}
	i := 1
	args := make([]any, 0, 2)
	if updateSeller.Name != nil {
		queryParts = append(queryParts, fmt.Sprintf("name=$%d", i))
		args = append(args, updateSeller.Name)
		i++
	}
	if updateSeller.Phone != nil {
		queryParts = append(queryParts, fmt.Sprintf("phone=$%d", i))
		args = append(args, updateSeller.Phone)
		i++
	}
	if len(queryParts) == 0 {
		return nil
	}
	query := fmt.Sprintf("UPDATE sellers SET %s WHERE id=$%d", strings.Join(queryParts, ", "), i)
	args = append(args, updateSeller.Id)
	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update item,err:%v", err)
	}
	return nil
}
