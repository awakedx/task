package item

import (
	"context"

	"github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/repository"
)

type ItemService struct {
	store *repository.Store
}

func NewItemService(store *repository.Store) *ItemService {
	return &ItemService{
		store: store,
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
		id, err := s.store.Items.Create(ctx, &item)
		ids = append(ids, id)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

func (s *ItemService) Get(ctx context.Context, itemId int) (*domain.Item, error) {
	item, err := s.store.Items.GetById(ctx, itemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) Delete(ctx context.Context, itemId int) (int, error) {
	return s.store.Items.Delete(ctx, itemId)
}

func (s *ItemService) UpdateItem(ctx context.Context, updateItem *common.UpdateItem) error {
	_, err := s.store.Items.GetById(ctx, *updateItem.Id)
	if err != nil {
		return err
	}
	if err = s.store.Items.Update(ctx, updateItem); err != nil {
		return err
	}
	return nil
}
