package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup) {
	rc := controllers.GetOrderController()

	r.GET("/orders/:id", middlewares.AuthorizeUser(), rc.GetOne)
	r.GET("/orders/admin-all", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.GetAllAdmin)
	r.GET("/orders/user-all", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.GetAllUser)
	r.POST("/orders", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.Create)
	r.PATCH("/orders/deliver-status", middlewares.AuthorizeUser(), middlewares.RoleGaurd("chef"), middlewares.RoleGaurd("admin"), rc.DeliveredStatus)
	r.PATCH("/orders/star", middlewares.AuthorizeUser(), rc.AddStars)
	r.DELETE("/orders/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), rc.Remove)
}
