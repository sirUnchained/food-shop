package controllers

import (
	"foodshop/api/helpers"

	"github.com/gin-gonic/gin"
)

type foodController struct{}

func GetFoodController() *foodController {
	return &foodController{}
}

func (fc *foodController) GetAll(ctx *gin.Context) *helpers.ResultResponse {}

func (fc *foodController) Create(ctx *gin.Context) *helpers.ResultResponse {}

func (fc *foodController) Update(ctx *gin.Context) *helpers.ResultResponse {}

func (fc *foodController) Removr(ctx *gin.Context) *helpers.ResultResponse {}
