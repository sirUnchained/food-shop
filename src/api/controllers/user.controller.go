package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

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
	us := services.GetUserService()

	result := us.GetAll()
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
}

func (u *userController) Remove(ctx *gin.Context) {
	us := services.GetUserService()

	result := us.Delete(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
}

func (u *userController) Update(ctx *gin.Context) {
	us := services.GetUserService()

	result := us.UpdateUser(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
}
