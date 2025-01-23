package controllers

import "github.com/gin-gonic/gin"

type userController struct{}

func GetUserController() *userController {
	return &userController{}
}

func (u *userController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "get all users"})
}
