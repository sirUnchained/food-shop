package services

import (
	"fmt"
	"foodshop/api/dto"
	"foodshop/data/models"
	"foodshop/data/postgres"

	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func GetAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(ctx *gin.Context) error {
	var userData dto.LoginDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		return err
	}

	db := postgres.GetDb()
	result := db.Model(&models.UserModel{}).Where("user_name = ?", userData.Username)
	fmt.Printf("%+v\n", result)

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
