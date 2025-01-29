package middlewares

import (
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"

	"github.com/gin-gonic/gin"
)

func RoleGaurd(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		db := postgres.GetDb()

		userRole := new(models.Roles)
		err := db.Model(&models.Roles{}).Where("user_id = ?", user.(models.Users).ID).First(userRole).Error
		if err != nil {
			helpers.SendResult(false, 500, err.Error(), nil, ctx)
			ctx.Abort()
			return
		}

		if userRole.State == "admin" {
			ctx.Next()
			return
		}

		if userRole.State != role {
			helpers.SendResult(false, 403, "this route is protected and you can't have access.", nil, ctx)
			ctx.Abort()
			return
		}
	}
}
