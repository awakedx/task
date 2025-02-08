package item

import "github.com/google/uuid"

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
