package customer

import (
	"context"

	common "github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
	"github.com/google/uuid"
)

type CustomerService struct {
	store *repository.Store
}

func NewCustomerService(store *repository.Store) *CustomerService {
	return &CustomerService{
		store: store,
	}
}
func (s *CustomerService) NewCustomer(ctx context.Context, newCustomer *domain.Customer) (uuid.UUID, error) {
	return s.store.Customers.Create(ctx, newCustomer)
}

func (s *CustomerService) GetCustomer(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	return s.store.Customers.Get(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, updateCustomer *common.UpdateCustomer) error {
	return s.store.Customers.Update(ctx, updateCustomer)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	return s.store.Customers.Delete(ctx, id)
}
