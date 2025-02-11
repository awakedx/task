package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id         int         `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
	TotalCost  float64     `json:"totalCost"`
	CustomerId uuid.UUID   `json:"customerId"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ItemId   int     `json:"itemId"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
