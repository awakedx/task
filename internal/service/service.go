package service

import (
	"context"

	"github.com/awakedx/task/internal/common"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
	"github.com/awakedx/task/internal/service/item"
	"github.com/awakedx/task/internal/service/seller"
	"github.com/google/uuid"
)

type Service struct {
	Items   Items
	Sellers Sellers
}

func NewService(store *repository.Store) *Service {
	return &Service{
		Items:   item.NewItemService(store.Items),
		Sellers: seller.NewSellerService(store.Sellers),
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
