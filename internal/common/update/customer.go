package common

import "github.com/google/uuid"

type UpdateCustomer struct {
	Id    *uuid.UUID
	Name  *string `json:"name" validate:"omitempty,alpha"`
	Phone *string `json:"phone" validate:"omitempty,numeric,len=10"`
}
