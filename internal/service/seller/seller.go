package seller

import (
	"context"

	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
	"github.com/google/uuid"
)

type SellerService struct {
	repo repository.Sellers
}

func NewSellerService(repo repository.Sellers) *SellerService {
	return &SellerService{
		repo: repo,
	}
}

func (s *SellerService) Create(ctx context.Context, newSeller *domain.Seller) (uuid.UUID, error) {
	return s.repo.Create(ctx, newSeller)
}
