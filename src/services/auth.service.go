package services

import (
	"foodshop/api/dto"
	"foodshop/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func (a *AuthService) GetAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(ctx *gin.Context) error {
	var userData dto.LoginDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Register(ctx *gin.Context) error {
	var userData dto.RegisterDto
	err := ctx.ShouldBindBodyWithJSON(&userData)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) GenerateTokenDetail(token *dto.TokenData) (*dto.TokenDetailDTO, error) {
	cfg := configs.GetConfigs()
	newTokenDetails := &dto.TokenDetailDTO{}

	// set access token expiration time
	newTokenDetails.AccessTokenExpiresIn = time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpiresIn) * time.Minute).Unix()

	// create access token claims
	atc := jwt.MapClaims{}
	atc["id"] = token.Id

	// generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, atc)
	var err error
	newTokenDetails.AccessToken, err = accessToken.SignedString([]byte(cfg.Jwt.AccessSecret))
	if err != nil {
		return nil, err
	}

	// set refresh token expiration time
	newTokenDetails.RefreshTokenExpiresIn = time.Now().Add(time.Duration(cfg.Jwt.RefreshTokenExpiresIn) * time.Minute).Unix()

	// create refresh token claims
	rtc := jwt.MapClaims{}
	rtc["id"] = token.Id

	// generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES384, rtc)
	newTokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(cfg.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return newTokenDetails, nil
}
