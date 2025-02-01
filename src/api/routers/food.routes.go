package routers

import (
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.RouterGroup) {
	r.GET("/foods")
	r.POST("/foods", middlewares.AuthorizeUser())
	r.PUT("/foods/:id", middlewares.AuthorizeUser())
	r.DELETE("/foods/:id", middlewares.AuthorizeUser())
}
