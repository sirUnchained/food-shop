package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func GetOrderController() *OrderController {
	return &OrderController{}
}

func (rc *OrderController) GetAllAdmin(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.GetAllAdmin(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)

}

func (rc *OrderController) GetAllUser(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.GetAllUser(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *OrderController) GetOne(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.GetOne(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *OrderController) Create(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.Create(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *OrderController) DeliveredStatus(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.DeliveredStatus(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *OrderController) AddStars(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.AddStars(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}

func (rc *OrderController) Remove(ctx *gin.Context) {
	orderService := services.GetOrderService()

	result := orderService.Remove(ctx)
	if !result.Ok {
		helpers.SendResult(false, result.Status, result.Message, result.Data, ctx)
		return
	}

	helpers.SendResult(true, result.Status, result.Message, result.Data, ctx)
}
