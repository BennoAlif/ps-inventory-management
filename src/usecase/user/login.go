package userusecase

import (
	"os"

	"github.com/BennoAlif/ps-cats-social/src/helpers"
)

type (
	ParamsLogin struct {
		Email    string
		Password string
	}
	GeneratedToken struct {
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expired_at"`
	}
	ResultLogin struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		AccessToken string `json:"accessToken"`
	}
)

func (i *sUserUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {

	emailMx := helpers.ValidateMx(p.Email)

	if emailMx != nil {
		return nil, emailMx
	}
	user, _ := i.userRepository.FindByEmail(&p.Email)

	if user == nil {
		return nil, ErrInvalidUser
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          user.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidUser
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTLogin)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &ResultLogin{
		Name:        user.Name,
		Email:       p.Email,
		AccessToken: accessToken,
	}, nil
}
