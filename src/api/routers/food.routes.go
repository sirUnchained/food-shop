package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.RouterGroup) {
	fc := controllers.GetFoodController()

	r.GET("/foods", fc.GetAll)
	r.POST("/foods", middlewares.AuthorizeUser(), middlewares.RoleGaurd("chef"), fc.Create)
	r.PUT("/foods/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("chef"), fc.Update)
	r.DELETE("/foods/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), fc.Remove)
}
