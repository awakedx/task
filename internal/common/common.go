package common

type UpdateItem struct {
	Id          *int
	Name        *string  `json:"name" validate:"omitempty"`
	Description *string  `json:"desc" validate:"omitempty"`
	Price       *float64 `json:"price" validate:"omitempty"`
	Stock       *int     `json:"stock" validate:"omitempty"`
}
