package controllers

import (
	"fmt"
	"foodshop/api/helpers"
	"foodshop/data/models"
	"foodshop/services"

	"github.com/gin-gonic/gin"
)

type authController struct{}

func GetAuthController() *authController {
	return &authController{}
}

func (a *authController) Login(ctx *gin.Context) {
	au := services.GetAuthService()

	user, ok := au.Login(ctx)
	if !ok.Ok {
		helpers.SendResult(false, ok.Status, ok.Message, nil, ctx)
		return
	}

	ts := services.GetTokenService()

	tokenDetail, ok := ts.GenerateTokenDetail(user, ctx)
	if !ok.Ok {
		helpers.SendResult(false, ok.Status, ok.Message, nil, ctx)
		return
	}

	fmt.Printf("%+v\n", tokenDetail)

	tokens := map[string]string{}
	tokens["accessToken"] = tokenDetail.AccessToken
	tokens["refreshToken"] = tokenDetail.RefreshToken
	helpers.SendResult(true, 200, "you are now authorized.", tokens, ctx)

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
	user, ok := au.Register(ctx)
	if !ok.Ok {
		helpers.SendResult(false, ok.Status, ok.Message, nil, ctx)
	}

	ts := services.GetTokenService()

	tokenDetail, ok := ts.GenerateTokenDetail(user, ctx)
	if !ok.Ok {
		helpers.SendResult(false, ok.Status, ok.Message, nil, ctx)
	}

	tokens := map[string]string{}
	tokens["accessToken"] = tokenDetail.AccessToken
	tokens["refreshToken"] = tokenDetail.RefreshToken
	helpers.SendResult(true, 200, "you are now authorized.", tokens, ctx)

}

func (a *authController) GetMe(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		helpers.SendResult(false, 500, "somthing went wrong.", nil, ctx)
		return
	}

	fmt.Printf("%+v\n", user)

	result := map[string]interface{}{}
	result["id"] = user.(models.Users).ID
	result["username"] = user.(models.Users).UserName
	result["email"] = user.(models.Users).Email
	result["phone"] = user.(models.Users).Phone
	result["createdAt"] = user.(models.Users).CreatedAt
	result["roles"] = user.(models.Users).Roles

	helpers.SendResult(true, 200, "", result, ctx)
}
