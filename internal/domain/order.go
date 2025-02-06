package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id           int
	CustomerId   uuid.UUID
	CreatedAt    time.Time
	OrderDetails []OrderProduct
}

type OrderProduct struct {
	ItemId   int `json:"ItemId"`
	Quantity int `json:"Quantity"`
	Price    int `json:"price"`
}
