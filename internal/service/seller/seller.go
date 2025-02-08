package seller

import (
	"context"

	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
	"github.com/google/uuid"
)

type SellerService struct {
	store *repository.Store
}

func NewSellerService(store *repository.Store) *SellerService {
	return &SellerService{
		store: store,
	}
}

func (s *SellerService) Create(ctx context.Context, newSeller *domain.Seller) (uuid.UUID, error) {
	return s.store.Sellers.Create(ctx, newSeller)
}
