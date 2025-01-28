package dto

type CategoryDTO struct {
	Title string `json:"title" binding:"required,alpha,min=5,lowercase"`
}
