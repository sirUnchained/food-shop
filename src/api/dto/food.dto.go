package dto

type FoodDTO struct {
	Name        string `json:"name" binding:"required,max=150"`
	Price       string `json:"price" binding:"required,numeric,max=100"`
	Description string `json:"description" binding:"required,max=500"`
}
