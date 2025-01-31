package dto

type UpdateUserDTO struct {
	UserName string `json:"user_name" binding:"required,alpha,min=3,max=100,lowercase"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Phone    string `json:"phone" binding:"required,IranMobile,len=11"`
}
