package dto

type RestaurantDTO struct {
	Description string `json:"description" binding:"required,max=255"`
	PostalCode  string `json:"postal_code" binding:"required,numeric"`
	Address     string `json:"address" binding:"required,max=255"`
	CategoryID  uint   `json:"category_id" binding:"required,numeric"`
}

type UpdateRestaurantDTO struct {
	Description string `json:"description" binding:"required,max=255"`
	PostalCode  string `json:"postal_code" binding:"required,numeric"`
	Address     string `json:"address" binding:"required,max=255"`
	CategoryID  uint   `json:"category_id" binding:"required,numeric"`
}
