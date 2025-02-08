package order

import (
	"context"

	"github.com/awakedx/task/internal/repository"
)

type OrderService struct {
	store *repository.Store
}

func NewOrderService(store *repository.Store) *OrderService {
	return &OrderService{
		store: store,
	}
}

func (s OrderService) NewOrder(ctx context.Context, orderD *OrderDetails) (int, error) {

	return 0, nil
}
