package services

import (
	"fmt"
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryService struct{}

func GetCategoryService() *categoryService {
	return &categoryService{}
}

func (cs *categoryService) GetAll(ctx *gin.Context) *helpers.ResultResponse {
	var categories []models.Category

	db := postgres.GetDb()

	err := db.Model(&models.Category{}).Preload("Creator", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_name", "email")
	}).Find(&categories).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 200, Message: "here you are.", Data: categories}

}

func (cs *categoryService) Create(ctx *gin.Context) *helpers.ResultResponse {
	db := postgres.GetDb()

	creator, exist := ctx.Get("user")
	if !exist {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: "something went wrong.", Data: nil}
	}

	fmt.Println(creator)

	var input dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "invalid input!", Data: nil}
	}

	slug := strings.Replace(input.Title, " ", "-", -1)

	var checkCategory models.Category
	if err := db.Model(&models.Category{}).Where("slug = ?", slug).First(&checkCategory).Error; err == nil {
		if checkCategory.ID != 0 {
			return &helpers.ResultResponse{Ok: false, Status: 400, Message: "category title is duplicated.", Data: nil}
		}
	}

	newCategory := models.Category{
		Title:     input.Title,
		Slug:      slug,
		CreatorID: creator.(models.Users).ID,
	}

	if err := db.Create(&newCategory).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	return &helpers.ResultResponse{Ok: true, Status: 201, Message: "category created", Data: newCategory}
}
