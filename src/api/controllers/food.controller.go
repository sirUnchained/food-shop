package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type foodController struct{}

func GetFoodController() *foodController {
	return &foodController{}
}

func (fc *foodController) GetAll(ctx *gin.Context) {
	fs := services.GetFoodService()

	result := fs.GetFoods(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (fc *foodController) Create(ctx *gin.Context) {
	fs := services.GetFoodService()

	result := fs.CreateFood(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (fc *foodController) Update(ctx *gin.Context) {}

func (fc *foodController) Remove(ctx *gin.Context) {}
