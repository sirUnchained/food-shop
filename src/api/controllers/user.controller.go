package controllers

import (
	"fmt"
	"foodshop/api/helpers"

	"github.com/gin-gonic/gin"
)

type userController struct{}

func GetUserController() *userController {
	return &userController{}
}

type RegisterData struct {
	UserName string `json:"user_name" binding:"required,alpha,min=3,max=100"`
	Phone    string `json:"phone" binding:"required,iranMobile,numeric,len=11"`
	Email    string `json:"email" binding:"required,email,max=100"`
}

func (u *userController) GetAll(ctx *gin.Context) {
	var newUser RegisterData

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		helpers.SendValidationErrors(400, err.Error(), ctx)
		return
	}

	helpers.SendResult(true, 200, "success", nil, ctx)
}
