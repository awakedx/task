package domain

import "github.com/google/uuid"

type Item struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerId    uuid.UUID `json:"sellerId"`
}
