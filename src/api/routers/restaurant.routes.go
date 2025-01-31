package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(r *gin.RouterGroup) {

	handler := controllers.GetRestaurantController()

	r.POST("restaurant/user/:userID", middlewares.AuthorizeUser(), handler.Create)
	r.PUT("restaurant/:id", middlewares.AuthorizeUser(), handler.Update)
	r.DELETE("/restaurant/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), handler.Remove)
}
