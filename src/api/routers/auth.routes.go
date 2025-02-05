package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	controller := controllers.GetAuthController()

	r.POST("/auth/login", controller.Login)
	r.POST("/auth/register", controller.Register)
	r.GET("/auth/me", middlewares.AuthorizeUser(), controller.GetMe)
	r.POST("/auth/refresh", controller.RefreshAccessToken)
}
