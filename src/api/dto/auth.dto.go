package dto

type LoginDto struct {
	Username string `json:"user_name"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	// Captcha string `json:"captcha" binding:"required"`
}

type RegisterDto struct {
	Username string `json:"user_name" binding:"required,alpha,min=3,max=100,lowercase"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Phone    string `json:"phone" binding:"required,IranMobile,len=11"`
	// Captcha string `json:"captcha" binding:"required"`
}

type TokenDetailDTO struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresIn  int64
	RefreshTokenExpiresIn int64
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenData struct {
	Id int
}
