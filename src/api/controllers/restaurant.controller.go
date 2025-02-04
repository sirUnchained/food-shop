package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct{}

func GetRestaurantController() *RestaurantController {
	return &RestaurantController{}
}

func (rc *RestaurantController) GetAll(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.GetAll(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *RestaurantController) GetOne(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.GetOne(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *RestaurantController) Create(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.CreateRestaurant(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *RestaurantController) Update(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.Update(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *RestaurantController) Verify(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.VerifyRestaurant(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *RestaurantController) Remove(ctx *gin.Context) {
	rs := services.GetRestaurantService()

	result := rs.Remove(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}
