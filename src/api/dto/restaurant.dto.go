package dto

type RestaurantDTO struct {
	Title       string `json:"title" binding:"required,max=150,lowercase"`
	Description string `json:"description" binding:"required,max=255"`
	PostalCode  string `json:"postal_code" binding:"required,numeric"`
	Address     string `json:"address" binding:"required,max=255"`
	CategoryID  uint   `json:"category_id" binding:"required,numeric"`
}

type UpdateRestaurantDTO struct {
	Title       string `json:"title" binding:"required,max=150,lowercase"`
	Description string `json:"description" binding:"required,max=255"`
	PostalCode  string `json:"postal_code" binding:"required,numeric"`
	Address     string `json:"address" binding:"required,max=255"`
	CategoryID  uint   `json:"category_id" binding:"required,numeric"`
}
