package services

import (
	"errors"
	"foodshop/api/dto"
	"foodshop/api/helpers"
	"foodshop/configs"
	"foodshop/data/models"
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

	// set access token expiration time
	newTokenDetails.AccessTokenExpiresIn = time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpiresIn) * time.Minute).Unix()

	// create access token claims
	atc := jwt.MapClaims{}
	atc["id"] = user.ID

	// generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	var err error
	newTokenDetails.AccessToken, err = accessToken.SignedString([]byte(cfg.Jwt.AccessSecret))
	if err != nil {
		return nil, helpers.NewResultResponse(false, 500, "failed to create access token.", nil)
	}

	// set refresh token expiration time
	newTokenDetails.RefreshTokenExpiresIn = time.Now().Add(time.Duration(cfg.Jwt.RefreshTokenExpiresIn) * time.Minute).Unix()

	// create refresh token claims
	rtc := jwt.MapClaims{}
	rtc["id"] = user.ID

	// generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	newTokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(cfg.Jwt.RefreshSecret))
	if err != nil {
		return nil, helpers.NewResultResponse(false, 500, "failed to create refresh token.", nil)
	}

	return newTokenDetails, helpers.NewResultResponse(true, 201, "", nil)
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
