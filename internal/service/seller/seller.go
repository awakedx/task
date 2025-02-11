package seller

import (
	"context"

	common "github.com/awakedx/task/internal/common/update"
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

func (s *SellerService) Update(ctx context.Context, updateSeller *common.UpdateSeller) error {
	return s.store.Sellers.Update(ctx, updateSeller)
}

func (s *SellerService) Get(ctx context.Context, id uuid.UUID) (*domain.Seller, error) {
	return s.store.Sellers.Get(ctx, id)
}

func (s *SellerService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.store.Sellers.Delete(ctx, id)
}
