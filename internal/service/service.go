package service

import (
	"context"

	"github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"

	"github.com/awakedx/task/internal/service/customer"
	"github.com/awakedx/task/internal/service/item"
	"github.com/awakedx/task/internal/service/order"
	"github.com/awakedx/task/internal/service/seller"

	"github.com/google/uuid"
)

type Service struct {
	Items     Items
	Sellers   Sellers
	Orders    Orders
	Customers Customers
}

func NewService(store *repository.Store) *Service {
	return &Service{
		Items:     item.NewItemService(store),
		Sellers:   seller.NewSellerService(store),
		Orders:    order.NewOrderService(store),
		Customers: customer.NewCustomerService(store),
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
	Update(ctx context.Context, updateSeller *common.UpdateSeller) error
	Get(ctx context.Context, id uuid.UUID) (*domain.Seller, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Orders interface {
	NewOrder(ctx context.Context, newOrder *order.OrderDetails) (int, error)
	GetById(ctx context.Context, orderId int) (*domain.Order, error)
}

type Customers interface {
	NewCustomer(ctx context.Context, newCustomer *domain.Customer) (uuid.UUID, error)
	GetCustomer(ctx context.Context, id uuid.UUID) (*domain.Customer, error)
	UpdateCustomer(ctx context.Context, updateCustomer *common.UpdateCustomer) error
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
}
