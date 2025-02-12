package repository

import (
	"context"

	"github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/google/uuid"
)

type Store struct {
	Items     Items
	Sellers   Sellers
	Orders    Orders
	Customers Customers
}

func NewStore(db *DB) *Store {
	return &Store{
		Items:     NewItemRepo(db),
		Sellers:   NewSellerRepo(db),
		Orders:    NewOrderRepo(db),
		Customers: NewCustomerRepo(db),
	}
}

type Items interface {
	Create(ctx context.Context, item *domain.Item) (int, error)
	GetById(ctx context.Context, id int) (*domain.Item, error)
	Update(ctx context.Context, updateItem *common.UpdateItem) error
	Delete(ctx context.Context, id int) (int, error)
}

type Sellers interface {
	Create(ctx context.Context, seller *domain.Seller) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Seller, error)
	Update(ctx context.Context, updateSeller *common.UpdateSeller) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Orders interface {
	Create(ctx context.Context, order *domain.Order, itemPrices map[int]float64) (int, error)
	GetById(ctx context.Context, orderId int) (*domain.Order, error)
}
type Customers interface {
	Create(ctx context.Context, newCustomer *domain.Customer) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Customer, error)
	Update(ctx context.Context, updateCustomer *common.UpdateCustomer) error
	Delete(ctx context.Context, id uuid.UUID) error
}
