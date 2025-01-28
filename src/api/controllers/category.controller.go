package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type categoryController struct{}

func GetCategoryController() *categoryController {
	return &categoryController{}
}

func (cc *categoryController) GetAll(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.GetAll(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (cc *categoryController) Create(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.Create(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (cc *categoryController) Update(ctx *gin.Context) {
	cs := services.GetCategoryService()

	result := cs.Update(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}
