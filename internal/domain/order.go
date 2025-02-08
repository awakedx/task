package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id         int         `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
	TotalCost  float64     `json:"TotalCost"`
	CustomerId uuid.UUID   `json:"customerId"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ItemId   int `json:"ItemId"`
	Quantity int `json:"Quantity"`
	Price    int `json:"price"`
}
