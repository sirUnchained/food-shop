package services

import (
	"errors"
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func GetAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(ctx *gin.Context) (*models.UserModel, error) {
	var userData dto.LoginDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		helpers.SendValidationErrors(400, err.Error(), ctx)
		return nil, errors.New("invalid inputs")
	}

	db := postgres.GetDb()
	var user models.UserModel
	db.Model(&models.UserModel{}).Where("user_name = ?", userData.Username).First(&user)
	if user.ID == 0 {
		helpers.SendResult(false, 404, "user name or password not found !", nil, ctx)
		return nil, errors.New("user name or password not found")
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if compareErr != nil {
		helpers.SendResult(false, 404, "invalid datas", nil, ctx)
		return nil, errors.New("invalid datas")
	}

	return &user, nil
}

func (a *AuthService) Register(ctx *gin.Context) (*models.UserModel, error) {
	var userData dto.RegisterDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		helpers.SendValidationErrors(400, err.Error(), ctx)
		return nil, errors.New("invalid inputs")
	}

	db := postgres.GetDb()
	var checkUser models.UserModel

	db.Model(&models.UserModel{}).Where("user_name = ?", userData.Username).Where("email = ?", userData.Email).Where("phone = ?", userData.Phone).First(&checkUser)
	if checkUser.ID != 0 {
		helpers.SendResult(false, 400, "these datas are not ok, please chose another.", nil, ctx)
		return nil, errors.New("these datas are not ok, please chose another")
	}

	hashByte, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 15)
	if err != nil {
		helpers.SendResult(false, 500, "with unknow reason hashing failed", nil, ctx)
		return nil, errors.New("with unknow reason hashing failed")
	}

	newUser := &models.UserModel{UserName: userData.Username, Password: string(hashByte), Email: userData.Email, Phone: userData.Phone}
	db.Create(&newUser)
	if newUser.ID == 0 {
		helpers.SendResult(false, 500, "unknow error to register user", nil, ctx)
		return nil, errors.New("unknow error to register user")
	}

	return newUser, nil
}

// func hashPass(pass string) (string, error)
