package userusecase

import (
	"os"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	"github.com/BennoAlif/ps-cats-social/src/helpers"

	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

type (
	ParamsCreateUser struct {
		PhoneNumber string
		Name        string
		Password    string
	}
)

func (i *sUserUsecase) CreateUser(p *ParamsCreateUser) (*ResultLogin, error) {

	emailMx := helpers.IsValidPhoneNumber(p.PhoneNumber)

	if emailMx != nil {
		return nil, emailMx
	}

	filters := entities.ParamsCreateUser{
		PhoneNumber: p.PhoneNumber,
	}

	checkPhoneNumber, _ := i.userRepository.IsExists(&filters)

	if checkPhoneNumber {
		return nil, ErrPhoneNumberAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)

	data, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		PhoneNumber: p.PhoneNumber,
		Name:        p.Name,
		Password:    hashedPassword,
	})

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          data.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &ResultLogin{
		ID:          strconv.FormatInt(data.ID, 10),
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		AccessToken: accessToken,
	}, nil
}
