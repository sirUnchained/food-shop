package controllers

import "github.com/gin-gonic/gin"

type RestaurantController struct{}

func GetRestaurantController() *RestaurantController {
	return &RestaurantController{}
}

func (rc *RestaurantController) Create(ctx *gin.Context) {

}

func (rc *RestaurantController) Update(ctx *gin.Context) {

}

func (rc *RestaurantController) Remove(ctx *gin.Context) {

}
