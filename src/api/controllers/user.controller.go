package controllers

import (
	"foodshop/api/helpers"

	"github.com/gin-gonic/gin"
)

type userController struct{}

func GetUserController() *userController {
	return &userController{}
}

func (u *userController) GetAll(ctx *gin.Context) {
	helpers.SendResult(true, 200, "success", nil, ctx)
}
