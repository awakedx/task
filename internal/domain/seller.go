package domain

import (
	"github.com/google/uuid"
)

type Seller struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name" validate:"required"`
	Phone string    `json:"phone" validate:"required,numeric,len=10"`
}
