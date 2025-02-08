package service

import (
	"context"

	"github.com/awakedx/task/internal/common/item"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"

	"github.com/awakedx/task/internal/service/item"
	"github.com/awakedx/task/internal/service/order"
	"github.com/awakedx/task/internal/service/seller"

	"github.com/google/uuid"
)

type Service struct {
	Items   Items
	Sellers Sellers
	Orders  Orders
}

func NewService(store *repository.Store) *Service {
	return &Service{
		Items:   item.NewItemService(store),
		Sellers: seller.NewSellerService(store),
		Orders:  order.NewOrderService(store),
	}
}

type Items interface {
	NewItem(ctx context.Context, newItem *item.ItemValues) ([]int, error)
	Get(ctx context.Context, itemId int) (*domain.Item, error)
	Delete(ctx context.Context, itemId int) (int, error)
	UpdateItem(ctx context.Context, updateItem *common.UpdateItem) error
}

type Sellers interface {
	Create(ctx context.Context, newSeller *domain.Seller) (uuid.UUID, error)
}

type Orders interface {
	NewOrder(ctx context.Context, newOrder *order.OrderDetails) (int, error)
}
