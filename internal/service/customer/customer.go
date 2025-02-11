package customer

import (
	"context"

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
