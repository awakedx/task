package order

import (
	"context"
	"fmt"

	"github.com/awakedx/task/internal/domain"
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

func (s *OrderService) NewOrder(ctx context.Context, orderD *OrderDetails) (int, error) {
	var totalSum float64 = 0
	ItemPrices := make(map[int]float64)
	for _, v := range orderD.Items {
		i, err := s.store.Items.GetById(ctx, v.ItemId)
		if err != nil {
			return 0, err
		}
		if i.Stock < v.Quantity {
			return 0, fmt.Errorf("out of stock %s (requsted: %d, available: %d", i.Name, v.Quantity, i.Stock)
		}
		totalSum += i.Price * float64(v.Quantity)
		ItemPrices[i.Id] = i.Price
	}

	order := domain.Order{
		TotalCost:  totalSum,
		CustomerId: orderD.CustomerId,
		Items:      orderD.Items,
	}
	orderId, err := s.store.Orders.Create(ctx, &order, ItemPrices)
	if err != nil {
		return 0, err
	}
	return orderId, err
}
func (s *OrderService) GetById(ctx context.Context, orderId int) (*domain.Order, error) {
	order, err := s.store.Orders.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}
