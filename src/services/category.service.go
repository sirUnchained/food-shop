package services

import (
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"strconv"
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

	// get all categories with relation users, ignore some of those users fields.
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
	// get and check current logged admin
	creator, _ := ctx.Get("user")
	// get entered inputs, if there was err then return it
	var input dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "invalid input!", Data: nil}
	}
	// create slug
	slug := strings.Replace(input.Title, " ", "-", -1)
	// check current slug exists
	var checkCategory models.Category
	if err := db.Model(&models.Category{}).Where("slug = ?", slug).First(&checkCategory).Error; err == nil {
		if checkCategory.ID != 0 {
			return &helpers.ResultResponse{Ok: false, Status: 400, Message: "category title is duplicated.", Data: nil}
		}
	}
	// create slug model
	newCategory := models.Category{
		Title:     input.Title,
		Slug:      slug,
		CreatorID: creator.(models.Users).ID,
	}
	// create slug in db, i hop there is no err
	if err := db.Create(&newCategory).Error; err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}
	// return done.
	return &helpers.ResultResponse{Ok: true, Status: 201, Message: "category created,", Data: newCategory}
}

func (cs *categoryService) Update(ctx *gin.Context) *helpers.ResultResponse {
	db := postgres.GetDb()
	// get entered inputs, if there was err then return it
	var input dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "invalid input!", Data: nil}
	}
	// get id from params and check it is a number or not
	categoryID := ctx.Param("id")
	var catID int
	var err error
	if catID, err = strconv.Atoi(categoryID); err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "category not found.", Data: nil}
	}
	// create slug
	slug := strings.Replace(input.Title, " ", "-", -1)
	// check given slug exists
	var checkCategory models.Category
	if err := db.Model(&models.Category{}).Where("slug = ?", slug).First(&checkCategory).Error; err != nil {
		if checkCategory.ID != 0 {
			return &helpers.ResultResponse{Ok: false, Status: 400, Message: "category title is duplicated.", Data: nil}
		}
	}
	// find and check category user want update exist or not
	var category models.Category
	db.Model(&models.Category{}).Where("ID = ?", catID).First(&category)
	if category.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "category not found.", Data: nil}
	}
	// update category
	category.Slug = slug
	category.Title = input.Title
	err = db.Save(&category).Error
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: err.Error(), Data: nil}
	}
	// send result
	return &helpers.ResultResponse{Ok: false, Status: 200, Message: "category updated.", Data: category}
}

// func (cs *categoryService) Remove(ctx *gin.Context) *helpers.ResultResponse {

// }
