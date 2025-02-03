package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup) {
	rc := controllers.GetOrderController()

	r.GET("/orders", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.GetAll)
	r.GET("/orders", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.Create)
	r.PATCH("/orders/deliver-status", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.DeliveredStatus)
	r.PATCH("/orders/star", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.AddStars)
	r.DELETE("/orders/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.Remove)
}
