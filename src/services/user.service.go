package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user removed.", Data: nil}
}

func (us *UserService) UpdateUser(ctx *gin.Context) *helpers.ResultResponse {
	newDatas := new(dto.UpdateUserDTO)
	err := ctx.ShouldBindBodyWithJSON(newDatas)
	if err != nil {
		if err.Error() != "EOF" {
			return &helpers.ResultResponse{Ok: false, Status: 400, Message: err.Error(), Data: nil}
		}
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "data is not valid.", Data: nil}
	}

	db := postgres.GetDb()
	user, _ := ctx.Get("user")
	updatedUser := new(models.Users)

	db.Model(&models.Users{}).Where("ID = ?", user.(models.Users).ID).First(updatedUser)
	if updatedUser.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: "normally user should be founded, seems like somthing went wrong.", Data: nil}
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(newDatas.Password), 15)
	updatedUser.Email = newDatas.Email
	updatedUser.Phone = string(hashed)
	updatedUser.Phone = newDatas.Phone
	updatedUser.UserName = newDatas.UserName

	db.Save(updatedUser)

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "updated.", Data: updatedUser}

}

func (us *UserService) Createrestaurant(ctx *gin.Context) *helpers.ResultResponse {
	idStr := ctx.Param("userID")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	restaurantData := new(dto.RestaurantDTO)
	err = ctx.ShouldBindBodyWithJSON(restaurantData)
	if err != nil {
		if err.Error() != "EOF" {
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

	restaurant := new(models.Restaurants)
	db.Model(&models.Restaurants{}).Where("owner = ?", user.ID).First(restaurant)
	if restaurant.ID != 0 {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "user has already a restaurant.", Data: nil}
	}

	restaurant.Address = restaurantData.Address
	restaurant.Description = restaurantData.Description
	restaurant.PostalCode = restaurantData.PostalCode
	restaurant.IsVerify = false
	restaurant.Owner = user.ID

	err = db.Model(&models.Restaurants{}).Create(restaurant).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user removed.", Data: nil}
}
