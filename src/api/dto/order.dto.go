package dto

type OrderDTO struct {
	Restaurant uint   `json:"restaurant" binding:"required,numeric"`
	Address    string `json:"address" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required,numeric"`
}

type ChangeStarDTO struct {
	Stars uint `json:"stars" binding:"required,min=1,max=5"`
}
