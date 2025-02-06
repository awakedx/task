package item

import (
	"context"

	"github.com/awakedx/task/internal/common"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
	"github.com/google/uuid"
)

type ItemService struct {
	repo repository.Items
}

func NewItemService(repo repository.Items) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) NewItem(ctx context.Context, itemValues *ItemValues) ([]int, error) {
	ids := make([]int, 0, len(itemValues.Items))
	for _, v := range itemValues.Items {
		item := domain.Item{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			SellerId:    v.SellerId,
		}
		id, err := s.repo.Create(ctx, &item)
		ids = append(ids, id)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

func (s *ItemService) Get(ctx context.Context, itemId int) (*domain.Item, error) {
	item, err := s.repo.GetById(ctx, itemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) Delete(ctx context.Context, itemId int) (int, error) {
	return s.repo.Delete(ctx, itemId)
}

func (s *ItemService) UpdateItem(ctx context.Context, updateItem *common.UpdateItem) error {
	_, err := s.repo.GetById(ctx, *updateItem.Id)
	if err != nil {
		return err
	}
	if err = s.repo.Update(ctx, updateItem); err != nil {
		return err
	}
	return nil
}

type ItemValues struct {
	Items []NewItem `json:"items" validate:"required"`
}

type NewItem struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"desc" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Stock       int       `json:"stock" validate:"required"`
	SellerId    uuid.UUID `json:"sellerId" validate:"required,uuid"`
}
