package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

func GetUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetAll() *helpers.ResultResponse {
	users := new([]models.Users)
	db := postgres.GetDb()

	err := db.Model(&models.Users{}).Select("user_name", "email", "phone", "created_at", "updated_at", "deleted_at").Find(users).Error
	if err != nil || len(*users) == 0 {
		message := "no user found."
		if err != nil {
			message = err.Error()
		}
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: message, Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "here you are.", Data: *users}
}

func (us *UserService) Delete(ctx *gin.Context) *helpers.ResultResponse {
	idStr := ctx.Param("id")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	db := postgres.GetDb()
	user := new(models.Users)

	db.Model(&models.Users{}).Where("ID = ?", id).First(user)
	if user.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	err = db.Model(&models.Users{}).Delete(user).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user removec.", Data: nil}
}

func (us *UserService) BecomeChef(ctx *gin.Context) *helpers.ResultResponse {
	idStr := ctx.Param("id")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	resturauntData := new(dto.Restaurant)
	err = ctx.ShouldBindBodyWithJSON(resturauntData)
	if err != nil {
		if err.Error() == "EOF" {
			return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
		}

		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}

	}

	db := postgres.GetDb()
	user := new(models.Users)

	db.Model(&models.Users{}).Where("ID = ?", id).First(user)
	if user.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	resturaunt := new(models.Restaurants)
	db.Model(&models.Restaurants{}).Where("owner = ?", user.ID).First(resturaunt)
	if resturaunt.ID != 0 {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "user has already a resturaunt.", Data: nil}
	}

	resturaunt.Address = resturauntData.Address
	resturaunt.Description = resturauntData.Description
	resturaunt.PostalCode = resturauntData.PostalCode
	resturaunt.IsVerify = false
	resturaunt.Owner = user.ID

	err = db.Model(&models.Restaurants{}).Create(resturaunt).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user removed.", Data: nil}
}
