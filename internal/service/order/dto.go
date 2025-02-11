package order

import (
	"github.com/awakedx/task/internal/domain"
	"github.com/google/uuid"
)

type OrderDetails struct {
	CustomerId uuid.UUID          `json:"customerId" validate:"required,uuid"`
	Items      []domain.OrderItem `json:"items" validate:"required"`
}
