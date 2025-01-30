package dto

type Restaurant struct {
	Description string `json:"description" binding:"required;max=255;"`
	PostalCode  string `json:"postal_code" binding:"required;numeric;"`
	Address     string `json:"address" binding:"required;max=255;"`
}
