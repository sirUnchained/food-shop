package controllers

import (
	"fmt"
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type authController struct{}

func GetAuthController() *authController {
	return &authController{}
}

func (a *authController) Login(ctx *gin.Context) {
	ts := &services.TokenService{}
	token := dto.TokenData{Id: 12}

	res, err := ts.GenerateTokenDetail(&token)
	if err != nil {
		helpers.SendResult(false, 400, err.Error(), nil, ctx)
		return
	}

	// jwtToken, err := ts.VerifyToken(res.AccessToken)
	// if err != nil {
	// 	helpers.SendResult(false, 400, err.Error(), nil, ctx)
	// 	return
	// }

	tokenClaim, err := ts.GetTokenClaims(res.AccessToken)
	if err != nil {
		helpers.SendResult(false, 400, err.Error(), nil, ctx)
		return
	}
	fmt.Printf("%+v\n", tokenClaim)

	helpers.SendResult(true, 200, "done", nil, ctx)
}

func (a *authController) Register(ctx *gin.Context) {}
