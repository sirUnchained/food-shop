package routers

import (
	"foodshop/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	handler := controllers.GetUserController()
	r.GET("/users", handler.GetAll)
}
