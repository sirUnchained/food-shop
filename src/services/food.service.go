package services

import (
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type foodService struct{}

func GetFoodService() *foodService {
	return &foodService{}
}

func (fc *foodService) GetFoods(ctx *gin.Context) *helpers.ResultResponse {
	var err error

	limit_str := ctx.Param("limit")
	var limit int
	if limit, err = strconv.Atoi(limit_str); err != nil {
		limit = 10
	}

	page_str := ctx.Param("page")
	var page int
	if page, err = strconv.Atoi(page_str); err != nil {
		page = 1
	}

	// Get the database connection
	db := postgres.GetDb()
	var foods []models.Foods

	// Retrieve the Foods with pagination
	err = db.Model(&models.Foods{}).Offset((page - 1) * limit).Limit(limit).Find(&foods).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	// Count the total number of foods
	var elemCount int64
	db.Model(&models.Foods{}).Count(&elemCount)

	// Prepare the response data
	data := map[string]interface{}{}
	data["result"] = foods
	data["count"] = elemCount
	data["limit"] = limit
	data["page"] = page

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "here you are.", Data: data}

}

func (fc *foodService) CreateFood(ctx *gin.Context) *helpers.ResultResponse {
	user, _ := ctx.Get("user")
	// get db
	db := postgres.GetDb()
	checkRestaurant := new(models.Restaurants)
	// check Restaurant exists
	err := db.Model(&models.Restaurants{}).Where("Owner = ?", user.(models.Users).ID).First(checkRestaurant).Error
	if checkRestaurant.ID == 0 {
		if err != nil {
			return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
		}

		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "restaurant not found.", Data: nil}
	}

	pic, err := ctx.FormFile("pic")
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: err.Error(), Data: nil}
	}
	dst := filepath.Join("public/foods", pic.Filename)
	if err := ctx.SaveUploadedFile(pic, dst); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	description := ctx.PostForm("description")

	food := models.Foods{
		Name:         name,
		Price:        price,
		Description:  description,
		Pic:          dst,
		IsAvailable:  true,
		RestaurantID: checkRestaurant.ID,
	}

	if err := db.Create(&food).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 201, Message: "food created successfully.", Data: food}

}

// func (fc *foodService) UpdateFood(ctx *gin.Context) *helpers.ResultResponse {

// }

// func (fc *foodService) RemoveFood(ctx *gin.Context) *helpers.ResultResponse {

// }
