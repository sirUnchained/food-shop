package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestaurantService struct{}

func GetRestaurantService() *RestaurantService {
	return &RestaurantService{}
}

func (us *RestaurantService) CreateRestaurant(ctx *gin.Context) *helpers.ResultResponse {
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

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user has now a Restaurant.", Data: nil}
}

func (us *RestaurantService) VerifyRestaurant(ctx *gin.Context) *helpers.ResultResponse {
	idStr := ctx.Param("id")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	db := postgres.GetDb()
	restaurant := new(models.Restaurants)
	err = db.Model(&models.Restaurants{}).Where("ID = ?", id).First(restaurant).Error
	if restaurant.ID == 0 {
		if err != nil {
			return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
		}
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	restaurant.IsVerify = true

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "restaurant is verifyed.", Data: nil}
}

func (rc *RestaurantService) Update(ctx *gin.Context) {

}

func (rc *RestaurantService) Remove(ctx *gin.Context) {

}
