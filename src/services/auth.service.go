package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/constants"
	"foodshop/data/models"
	"foodshop/data/postgres"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func GetAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(ctx *gin.Context) (*models.Users, *helpers.ResultResponse) {
	var userData dto.LoginDto
	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		if err.Error() != "EOF" {
			return nil, helpers.NewResultResponse(false, 400, err.Error(), nil)
		}
		return nil, helpers.NewResultResponse(false, 400, "validation failed, please full the fields correctly.", nil)
	}
	// find user by username
	db := postgres.GetDb()
	var user models.Users
	db.Model(&models.Users{}).Where("user_name = ?", userData.Username).First(&user)
	if user.ID == 0 {
		return nil, helpers.NewResultResponse(false, 404, "user name or password not found!", nil)
	}
	// Compare passwords
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if compareErr != nil {
		return nil, helpers.NewResultResponse(false, 404, "invalid data", nil)
	}

	return &user, helpers.NewResultResponse(true, 200, "login successful", &user)
}

func (a *AuthService) Register(ctx *gin.Context) (*models.Users, *helpers.ResultResponse) {
	// Bind the JSON payload to the userData struct
	var userData dto.RegisterDto
	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		if err.Error() != "EOF" {
			// Return error response if JSON binding fails
			return nil, helpers.NewResultResponse(false, 400, err.Error(), nil)
		}
		// Return validation error response if fields are not filled correctly
		return nil, helpers.NewResultResponse(false, 400, "validation failed, please full the fields correctly.", nil)
	}

	// Get the database connection
	db := postgres.GetDb()
	var checkUser models.Users

	// Check if a user with the same username, email, or phone already exists
	db.Model(&models.Users{}).Where("user_name = ?", userData.Username).Where("email = ?", userData.Email).Where("phone = ?", userData.Phone).First(&checkUser)
	if checkUser.ID != 0 {
		return nil, helpers.NewResultResponse(false, 400, "these datas are not ok, please chose another.", nil)
	}

	// Hash the user's password
	hashByte, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 15)
	if err != nil {
		return nil, helpers.NewResultResponse(false, 500, "with unknow reason hashing failed", nil)
	}

	// Create a new user with the provided data
	newUser := &models.Users{UserName: userData.Username, Password: string(hashByte), Email: userData.Email, Phone: userData.Phone}
	db.Create(&newUser)
	if newUser.ID == 0 {
		return nil, helpers.NewResultResponse(false, 500, "unknow error to register user", nil)
	}

	// Check if the user table is empty; the first user should be an admin
	docCount := int64(-1)
	db.Model(&models.Users{}).Count(&docCount)
	if docCount == 1 {
		// Assign admin role to the first user
		db.Model(&models.Roles{}).Create(&models.Roles{UserID: newUser.ID, State: string(constants.ADMIN)})
	} else {
		// Assign user role to subsequent users
		db.Model(&models.Roles{}).Create(&models.Roles{UserID: newUser.ID, State: string(constants.USER)})
	}

	return newUser, helpers.NewResultResponse(true, 201, "registration successful", newUser)
}

func (a *AuthService) GetMe() {}
