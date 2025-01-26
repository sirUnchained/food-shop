package controllers

import "github.com/gin-gonic/gin"

type authController struct{}

func GetAuthController() *authController {
	return &authController{}
}

func (a *authController) Login(ctx *gin.Context) {}

func (a *authController) Register(ctx *gin.Context) {}
