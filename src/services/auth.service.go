package services

import (
	"foodshop/api/dto"

	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func (a *AuthService) GetAuthService() *AuthService {
	return &AuthService{}
}

type token struct {
	Id    int
	Phone string
}

func (a *AuthService) Login(ctx *gin.Context) error {
	var userData dto.LoginDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Register(ctx *gin.Context) error {
	var userData dto.RegisterDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		return err
	}

	return nil
}
