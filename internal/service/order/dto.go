package order

import "github.com/google/uuid"

type OrderDetails struct {
	CustomerId uuid.UUID   `json:"customerId" validate:"required"`
	Items      []OrderItem `json:"items" validate:"required"`
}

type OrderItem struct {
	ItemId   int `json:"itemId" validtae:"required"`
	Quantity int `json:"quant" validate:"required"`
}
