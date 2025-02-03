package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func GetOrderController() *OrderController {
	return &OrderController{}
}

func (rc *OrderController) GetAll(ctx *gin.Context)          {}
func (rc *OrderController) Create(ctx *gin.Context)          {}
func (rc *OrderController) DeliveredStatus(ctx *gin.Context) {}
func (rc *OrderController) AddStars(ctx *gin.Context)        {}
func (rc *OrderController) Remove(ctx *gin.Context)          {}
