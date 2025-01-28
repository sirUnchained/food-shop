package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	cc := controllers.GetCategoryController()

	r.GET("/category", cc.GetAll)
	r.POST("/category", middlewares.AuthorizeUser(), cc.Create)
	r.PATCH("/category/:id", middlewares.AuthorizeUser(), cc.Update)
}
