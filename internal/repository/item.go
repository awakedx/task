package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
	"github.com/jackc/pgx/v5"
)

type ItemRepo struct {
	db *DB
}

func NewItemRepo(db *DB) *ItemRepo {
	return &ItemRepo{db: db}
}

func (r *ItemRepo) Create(ctx context.Context, item *domain.Item) (int, error) {
	var id int
	query := `
        INSERT INTO items (name,description,price,stock,seller_id)
        VALUES ($1,$2,$3,$4,$5)
        RETURNING id
    `
	err := r.db.QueryRow(ctx, query,
		item.Name,
		item.Description,
		item.Price,
		item.Stock,
		item.SellerId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *ItemRepo) GetById(ctx context.Context, id int) (*domain.Item, error) {
	var itemDB domain.Item
	query := `SELECT id,name,description,price,stock,seller_id FROM items WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(&itemDB.Id, &itemDB.Name, &itemDB.Description, &itemDB.Price, &itemDB.Stock, &itemDB.SellerId)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%w with ID %d", utils.NotFoundError, id)
	}
	if err != nil {
		return nil, fmt.Errorf("%w ,select by id %w", err, utils.InternalError)
	}
	return &itemDB, nil
}

func (r *ItemRepo) Update(ctx context.Context, updateItem *common.UpdateItem) error {
	queryParts := []string{}
	i := 1
	args := make([]any, 0, 4)
	if updateItem.Name != nil {
		queryParts = append(queryParts, fmt.Sprintf("name=$%d", i))
		args = append(args, updateItem.Name)
		i++
	}
	if updateItem.Description != nil {
		queryParts = append(queryParts, fmt.Sprintf("description=$%d", i))
		args = append(args, updateItem.Description)
		i++
	}
	if updateItem.Price != nil {
		queryParts = append(queryParts, fmt.Sprintf("price=$%d", i))
		args = append(args, updateItem.Price)
		i++
	}
	if updateItem.Stock != nil {
		queryParts = append(queryParts, fmt.Sprintf("stock=$%d", i))
		args = append(args, updateItem.Stock)
		i++
	}
	if len(queryParts) == 0 {
		return nil
	}
	query := fmt.Sprintf("UPDATE items SET %s WHERE id=$%d", strings.Join(queryParts, ", "), i)
	args = append(args, updateItem.Id)
	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update item,err:%v", err)
	}
	return nil
}

func (r *ItemRepo) GetAll(ctx context.Context) ([]domain.Item, error) {
	items := make([]domain.Item, 0, 3)
	query := `SELECT * FROM items`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item domain.Item
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Stock,
			&item.SellerId,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemRepo) Delete(ctx context.Context, id int) (int, error) {
	var deletedId int
	query := `DELETE FROM items WHERE id=$1 RETURNING id`
	err := r.db.QueryRow(ctx, query, id).Scan(&deletedId)
	if err != nil && err.Error() == "no rows in result set" {
		return 0, fmt.Errorf("Nothing to delete")
	}
	return deletedId, nil
}
