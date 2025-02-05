package services

import (
	"errors"
	"fmt"
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/configs"
	"foodshop/data/models"
	"foodshop/data/postgres"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct{}

func GetTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) GenerateTokenDetail(user *models.Users, ctx *gin.Context) (*dto.TokenDetailDTO, *helpers.ResultResponse) {
	cfg := configs.GetConfigs()
	newTokenDetails := &dto.TokenDetailDTO{}
	var err error

	// call for access token
	newTokenDetails.AccessToken, err = generateAccessToken(cfg, user)
	if err != nil {
		return nil, &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	// call for refresh token
	newTokenDetails.RefreshToken, err = generateRefreshToken(cfg)
	if err != nil {
		return nil, helpers.NewResultResponse(false, 500, "failed to create refresh token.", nil)
	}
	saveNewDeleteOldRefreshToken(newTokenDetails.RefreshToken, user)

	return newTokenDetails, helpers.NewResultResponse(true, 201, "", nil)
}

func (ts *TokenService) RefreshAccessToken(ctx *gin.Context) *helpers.ResultResponse {
	var rtDto dto.RefreshTokenDTO
	cfg := configs.GetConfigs()

	err := ctx.ShouldBindBodyWithJSON(&rtDto)
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 400, Message: "you'r refresh token is no more valid, please sign in again.", Data: nil}
	}

	// search in db try to find access token we got from client, if nothing found then user must login again
	refreshTokenData := new(models.RefreshTokens)
	db := postgres.GetDb()
	db.Model(&models.RefreshTokens{}).Where("token = ?", rtDto.RefreshToken).Preload("User").First(refreshTokenData)
	if refreshTokenData.ID == 0 {
		return &helpers.ResultResponse{Ok: false, Status: 404, Message: "please sign in or sign up first.", Data: nil}
	}

	// if the refresh token found in db, done we can create a new access token and refresh token
	var newTokenDetails dto.TokenDetailDTO
	newTokenDetails.AccessToken, err = generateAccessToken(cfg, &refreshTokenData.User)
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}
	newTokenDetails.RefreshToken, err = generateRefreshToken(cfg)
	if err != nil {
		return &helpers.ResultResponse{Ok: false, Status: 500, Message: err.Error(), Data: nil}
	}

	// save new delete old refresh token
	saveNewDeleteOldRefreshToken(newTokenDetails.RefreshToken, &refreshTokenData.User)
	fmt.Printf("%+v\n", newTokenDetails)

	return &helpers.ResultResponse{Ok: true, Status: 201, Message: "here your refreshed tokens !", Data: newTokenDetails}

}

func (ts *TokenService) VerifyToken(token string) (*jwt.Token, *helpers.ResultResponse) {
	// get the configuration settings
	cfg := configs.GetConfigs()

	// parse the token with a function to return the key for validation
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// check if the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is not valid")
		}
		// return the secret key for validation
		return []byte(cfg.Jwt.AccessSecret), nil
	})
	if err != nil {
		return nil, helpers.NewResultResponse(false, 400, err.Error(), nil)
	}

	return parsedToken, helpers.NewResultResponse(true, 200, "", nil)
}

func (ts *TokenService) GetTokenClaims(token string) (*dto.TokenData, *helpers.ResultResponse) {
	// initialize the result map to store claims
	// claimsResult := map[string]interface{}{}
	claimsResult := &dto.TokenData{}

	// verify the token
	parsedToken, ok := ts.VerifyToken(token)
	if !ok.Ok {
		return nil, helpers.NewResultResponse(false, 400, "token is not valid, please sign in or sign up first", nil)
	}

	// extract claims from the token
	claims, isOk := parsedToken.Claims.(jwt.MapClaims)
	if !isOk && !parsedToken.Valid {
		return nil, helpers.NewResultResponse(false, 500, "failed to parse token claims.", nil)
	}
	// copy claims to the result
	if id, ok := claims["id"].(float64); ok {
		claimsResult.Id = int(id)
	} else {
		return nil, helpers.NewResultResponse(false, 500, "failed to parse token claims.", nil)
	}

	return claimsResult, helpers.NewResultResponse(true, 200, "", nil)

}

func generateAccessToken(cfg *configs.Configs, user *models.Users) (string, error) {
	// create access token claims
	atc := jwt.MapClaims{}
	atc["id"] = user.ID
	// set access token expiration time
	atc["exp"] = time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpiresIn) * time.Minute).Unix()

	// generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	token_str, err := accessToken.SignedString([]byte(cfg.Jwt.AccessSecret))
	if err != nil {
		return "", errors.New("failed to create token.")
	}

	return token_str, nil
}

func generateRefreshToken(cfg *configs.Configs) (string, error) {
	// create refresh token claims
	rtc := jwt.MapClaims{}
	rtc["token"] = rand.Intn(500)
	// set access refresh token expiration time
	rtc["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(cfg.Jwt.RefreshTokenExpiresIn)).Unix()

	// generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	token_str, err := refreshToken.SignedString([]byte(cfg.Jwt.RefreshSecret))
	if err != nil {
		return "", errors.New("failed to generate refresh token.")
	}

	return token_str, nil
}

func saveNewDeleteOldRefreshToken(token string, user *models.Users) {
	var newToken models.RefreshTokens
	newToken.Token = token
	newToken.UserID = user.ID

	var oldToken models.RefreshTokens

	db := postgres.GetDb()
	// delelte old one
	db.Model(&models.RefreshTokens{}).Where("user_id = ?", user.ID).First(&oldToken)
	db.Delete(&oldToken)
	// create new one
	db.Model(&models.RefreshTokens{}).Create(&newToken)
}
