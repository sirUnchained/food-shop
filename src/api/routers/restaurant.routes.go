package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(r *gin.RouterGroup) {

	handler := controllers.GetRestaurantController()

	r.GET("/restaurant", handler.GetAll)
	r.GET("/restaurant/:id", handler.GetOne)
	r.POST("/restaurant", middlewares.AuthorizeUser(), handler.Create)
	r.PUT("/restaurant", middlewares.AuthorizeUser(), handler.Update)
	r.PATCH("/restaurant/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), handler.Verify)
	r.DELETE("/restaurant/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), handler.Remove)
}
