package userusecase

import (
	"os"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	"github.com/BennoAlif/ps-cats-social/src/helpers"
)

type (
	ParamsLogin struct {
		PhoneNumber string
		Password    string
	}
	GeneratedToken struct {
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expired_at"`
	}
	ResultLogin struct {
		ID          string `json:"userId"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		AccessToken string `json:"accessToken"`
	}
)

func (i *sUserUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {

	emailMx := helpers.IsValidPhoneNumber(p.PhoneNumber)

	if emailMx != nil {
		return nil, emailMx
	}

	filters := entities.ParamsCreateUser{
		PhoneNumber: p.PhoneNumber,
	}

	user, _ := i.userRepository.FindOne(&filters)

	if user == nil {
		return nil, ErrUserNotFound
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
		ID:          strconv.FormatInt(user.ID, 10),
		Name:        user.Name,
		PhoneNumber: p.PhoneNumber,
		AccessToken: accessToken,
	}, nil
}
