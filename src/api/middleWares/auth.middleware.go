package middlewares

import (
	"fmt"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"foodshop/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		token := strings.Split(bearerToken, "Bearer ")

		if len(token) < 2 {
			helpers.SendUnAuthorizedResult(ctx)
			ctx.Abort()
			return
		}

		ts := services.GetTokenService()
		tokenClaims, ok := ts.GetTokenClaims(token[1])

		if !ok.Ok {
			helpers.SendResult(false, ok.Status, ok.Message, nil, ctx)
			ctx.Abort()
			return
		}

		var user models.UserModel
		db := postgres.GetDb()
		db.Model(&models.UserModel{}).Where("id = ?", tokenClaims.Id).First(&user)
		if user.ID == 0 {
			fmt.Printf("%+v\n", tokenClaims)
			helpers.SendResult(false, 404, "user not foound.", nil, ctx)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
