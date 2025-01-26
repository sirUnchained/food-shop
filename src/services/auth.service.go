package services

import (
	"errors"
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

func (a *AuthService) VerifyToken(token string) (*jwt.Token, error) {
	cfg := configs.GetConfigs()

	result, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unknown err in auth service verify token func.")
		}
		return []byte(cfg.Jwt.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AuthService) GetTokenClaims(token string) (claimResult map[string]interface{}, err error) {
	result, err := a.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		for k, v := range claims {
			claimResult[k] = v
		}
		return claimResult, nil
	}

	return nil, errors.New("error in auth service GetTokenClaims func.")
}
