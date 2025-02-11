package repository

import (
	"context"

	"github.com/awakedx/task/internal/domain"
	"github.com/jackc/pgx/v5"
)

type OrderRepo struct {
	db *DB
}

func NewOrderRepo(db *DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Create(ctx context.Context, orderD *domain.Order, itemPrices map[int]float64) (int, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, nil
	}

	query := `
        INSERT INTO orders (customer_id, total_cost) VALUES($1,$2)
        RETURNING id
    `
	var orderId int
	err = r.db.QueryRow(ctx, query, orderD.CustomerId, orderD.TotalCost).Scan(&orderId)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	for _, v := range orderD.Items {
		query := `
        INSERT INTO order_items (order_id,item_id,quantity,buy_price)
        VALUES ($1,$2,$3,$4)
        `
		_, err := r.db.Exec(ctx, query, orderId, v.ItemId, v.Quantity, itemPrices[v.ItemId])
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
		query = `
        UPDATE items SET stock=stock-$1 WHERE id=$2
        `
		_, err = r.db.Exec(ctx, query, v.Quantity, v.ItemId)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return 0, nil
	}
	return orderId, nil
}

func (r *OrderRepo) GetById(ctx context.Context, orderId int) (*domain.Order, error) {
	query := `SELECT customer_id,created_at,total_cost FROM orders WHERE id=$1`
	var order domain.Order
	order.Id = orderId
	err := r.db.QueryRow(ctx, query, orderId).Scan(&order.CustomerId, &order.CreatedAt, &order.TotalCost)
	if err != nil {
		return nil, err
	}
	var orderItem domain.OrderItem
	query = `SELECT item_id,quantity,buy_price FROM order_items WHERE order_id=$1`
	resrows, err := r.db.Query(ctx, query, orderId)
	defer resrows.Close()
	for resrows.Next() {
		err := resrows.Scan(&orderItem.ItemId, &orderItem.Quantity, &orderItem.Price)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, orderItem)
	}
	return &order, nil
}
