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

func (us *RestaurantService) GetAll(ctx *gin.Context) *helpers.ResultResponse {
	var err error

	// Get the limit query parameter, default to 10 if not provided or invalid
	limit_str := ctx.Query("limit")
	var limit int
	if limit, err = strconv.Atoi(limit_str); err != nil {
		limit = 10
	}

	// Get the page query parameter, default to 1 if not provided or invalid
	page_str := ctx.Query("page")
	var page int
	if page, err = strconv.Atoi(page_str); err != nil {
		page = 1
	}

	// Get the database connection
	db := postgres.GetDb()
	var restaurants []models.Restaurants

	// Retrieve the restaurants with pagination
	err = db.Model(&models.Restaurants{}).Offset((page - 1) * limit).Limit(limit).Find(&restaurants).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	// Count the total number of restaurants
	var elemCount int64
	db.Model(&models.Restaurants{}).Count(&elemCount)

	// Prepare the response data
	data := map[string]interface{}{}
	data["result"] = restaurants
	data["count"] = elemCount
	data["limit"] = limit
	data["page"] = page

	// Return success response
	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "here you are.", Data: data}
}

func (us *RestaurantService) CreateRestaurant(ctx *gin.Context) *helpers.ResultResponse {
	// Get the user ID from the URL parameter
	idStr := ctx.Param("userID")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	// Bind the request body to restaurantData
	restaurantData := new(dto.RestaurantDTO)
	err = ctx.ShouldBindBodyWithJSON(restaurantData)
	if err != nil {
		if err.Error() != "EOF" {
			return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
		}

		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	// Get the database connection
	db := postgres.GetDb()
	user := new(models.Users)

	// Find the user by ID
	db.Model(&models.Users{}).Where("ID = ?", id).First(user)
	if user.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "user not found.", Data: nil}
	}

	// Check if the user already has a restaurant
	restaurant := new(models.Restaurants)
	db.Model(&models.Restaurants{}).Where("owner = ?", user.ID).First(restaurant)
	if restaurant.ID != 0 {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "user has already a restaurant.", Data: nil}
	}

	// Create a new restaurant
	restaurant.Address = restaurantData.Address
	restaurant.Description = restaurantData.Description
	restaurant.PostalCode = restaurantData.PostalCode
	restaurant.IsVerify = false
	restaurant.Owner = user.ID

	// Save the new restaurant to the database
	err = db.Model(&models.Restaurants{}).Create(restaurant).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
	}

	// Return success response
	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "user has now a Restaurant.", Data: nil}
}

func (us *RestaurantService) VerifyRestaurant(ctx *gin.Context) *helpers.ResultResponse {
	// get id from param, if it is not valid return invalid response
	idStr := ctx.Param("id")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}
	// try to find restaurant, if could not, return
	db := postgres.GetDb()
	restaurant := new(models.Restaurants)
	err = db.Model(&models.Restaurants{}).Where("ID = ?", id).First(restaurant).Error
	if restaurant.ID == 0 {
		if err != nil {
			return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
		}
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	// verify restaurant
	restaurant.IsVerify = true

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "restaurant is verifyed.", Data: nil}
}

func (rc *RestaurantService) Update(ctx *gin.Context) *helpers.ResultResponse {
	// Get the user from the context
	user, _ := ctx.Get("user")

	// Get the restaurant ID from the URL parameter
	idStr := ctx.Param("id")
	var id int
	var err error
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	// Bind the request body to restaurantDTO
	var restaurantDTO dto.RestaurantDTO
	err = ctx.ShouldBindBodyWithJSON(&restaurantDTO)
	if err != nil {
		if err.Error() != "EOF" {
			return &helpers.ResultResponse{Ok: false, Status: 404, Message: err.Error(), Data: nil}
		}
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "invalid data.", Data: nil}
	}

	// Get the database connection
	db := postgres.GetDb()
	restaurant := new(models.Restaurants)

	// Find the restaurant by ID
	err = db.Model(&models.Restaurants{}).Where("ID = ?", id).First(restaurant).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	} else if restaurant.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	} else if user.(models.Users).ID != restaurant.Owner {
		return &helpers.ResultResponse{Ok: false, Status: 403, Message: "you cannot change this restaurant.", Data: nil}
	}

	// Update the restaurant fields
	restaurant.Address = restaurantDTO.Address
	restaurant.Description = restaurantDTO.Description
	restaurant.PostalCode = restaurantDTO.PostalCode

	// Save the updated restaurant to the database
	db.Save(restaurant)

	// Return success response
	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "updated restaurant.", Data: nil}
}

func (rc *RestaurantService) Remove(ctx *gin.Context) *helpers.ResultResponse {
	// Get the restaurant ID from the URL parameter
	idStr := ctx.Param("id")
	var err error
	var id int
	if id, err = strconv.Atoi(idStr); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	// Get the database connection
	db := postgres.GetDb()
	restaurant := new(models.Restaurants)
	// Find the restaurant by ID
	err = db.Model(&models.Restaurants{}).Where("ID = ?", id).First(restaurant).Error
	if restaurant.ID == 0 {
		if err != nil {
			return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
		}
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	// Delete the restaurant
	err = db.Delete(restaurant).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	// Return success response
	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "restaurant removed.", Data: nil}
}
