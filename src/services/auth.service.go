package services

import (
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

func (a *AuthService) Login(ctx *gin.Context) (*models.UserModel, *helpers.ResultResponse) {
	var userData dto.LoginDto
	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		return nil, helpers.NewResultResponse(false, 400, err.Error(), nil)
	}

	db := postgres.GetDb()
	var user models.UserModel
	db.Model(&models.UserModel{}).Where("user_name = ?", userData.Username).First(&user)
	if user.ID == 0 {
		return nil, helpers.NewResultResponse(false, 404, "user name or password not found!", nil)
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if compareErr != nil {
		return nil, helpers.NewResultResponse(false, 404, "invalid data", nil)
	}

	return &user, helpers.NewResultResponse(true, 200, "login successful", &user)
}

func (a *AuthService) Register(ctx *gin.Context) (*models.UserModel, *helpers.ResultResponse) {
	var userData dto.RegisterDto
	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		return nil, helpers.NewResultResponse(false, 400, err.Error(), nil)
	}

	db := postgres.GetDb()
	var checkUser models.UserModel

	db.Model(&models.UserModel{}).Where("user_name = ?", userData.Username).Where("email = ?", userData.Email).Where("phone = ?", userData.Phone).First(&checkUser)
	if checkUser.ID != 0 {
		return nil, helpers.NewResultResponse(false, 400, "these datas are not ok, please chose another.", nil)
	}

	hashByte, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 15)
	if err != nil {
		return nil, helpers.NewResultResponse(false, 500, "with unknow reason hashing failed", nil)
	}

	newUser := &models.UserModel{UserName: userData.Username, Password: string(hashByte), Email: userData.Email, Phone: userData.Phone}
	db.Create(&newUser)
	if newUser.ID == 0 {
		return nil, helpers.NewResultResponse(false, 500, "unknow error to register user", nil)
	}

	return newUser, helpers.NewResultResponse(true, 201, "registration successful", newUser)
}

func (a *AuthService) GetMe() {}
