package userusecase

import (
	"os"

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

	checkPhoneNumber, _ := i.userRepository.FindByPhoneNumber(&p.PhoneNumber)

	if checkPhoneNumber != nil {
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
		Role:            data.Role,
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
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		AccessToken: accessToken,
	}, nil
}
