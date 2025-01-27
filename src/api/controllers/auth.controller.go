package controllers

import (
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type authController struct{}

func GetAuthController() *authController {
	return &authController{}
}

func (a *authController) Login(ctx *gin.Context) {
	au := services.GetAuthService()

	user, err := au.Login(ctx)
	if err != nil {
		helpers.SendResult(false, 500, err.Error(), nil, ctx)
		return
	}

	ts := services.GetTokenService()

	tokenDetail, err := ts.GenerateTokenDetail(user, ctx)
	if err != nil {
		return
	}

	// res, err := ts.GenerateTokenDetail(&token)
	// if err != nil {
	// 	helpers.SendResult(false, 400, err.Error(), nil, ctx)
	// 	return
	// }

	// jwtToken, err := ts.VerifyToken(res.AccessToken)
	// if err != nil {
	// 	helpers.SendResult(false, 400, err.Error(), nil, ctx)
	// 	return
	// }

	// tokenClaim, err := ts.GetTokenClaims(res.AccessToken)
	// if err != nil {
	// 	helpers.SendResult(false, 400, err.Error(), nil, ctx)
	// 	return
	// }
	// fmt.Printf("%+v\n", tokenClaim)

}

func (a *authController) Register(ctx *gin.Context) {
	au := services.GetAuthService()
	result := au.Register(ctx)
	helpers.SendResult(result.Ok, result.Status, result.Message, &result.Data, ctx)
}
