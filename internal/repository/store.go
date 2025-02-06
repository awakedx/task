package repository

import (
	"context"

	"github.com/awakedx/task/internal/common"
	"github.com/awakedx/task/internal/domain"
	"github.com/google/uuid"
)

type Store struct {
	Items   Items
	Sellers Sellers
}

func NewStore(db *DB) *Store {
	return &Store{
		Items:   NewItemRepo(db),
		Sellers: NewSellerRepo(db),
	}
}

type Items interface {
	Create(ctx context.Context, item *domain.Item) (int, error)
	GetById(ctx context.Context, id int) (*domain.Item, error)
	Delete(ctx context.Context, id int) (int, error)
	Update(ctx context.Context, updateItem *common.UpdateItem) error
}

type Sellers interface {
	Create(ctx context.Context, seller *domain.Seller) (uuid.UUID, error)
}
