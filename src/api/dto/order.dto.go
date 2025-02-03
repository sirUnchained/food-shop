package dto

type OrderDTO struct {
	Restaurant uint   `json:"restaurant" binding:"required,numeric"`
	Address    string `json:"address" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required,numeric"`
}
