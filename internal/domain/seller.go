package domain

import (
	"github.com/google/uuid"
)

type Seller struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
}
