package dto

type CategoryDTO struct {
	Title string `json:"title" binding:"required,min=5,lowercase"`
}
