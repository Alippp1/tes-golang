package dto

type CreateItemRequest struct {
	Name  string  `json:"name" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type UpdateItemRequest struct {
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

type ItemResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}
