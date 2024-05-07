package userusecase

import (
	"os"

	"github.com/BennoAlif/ps-cats-social/src/helpers"

	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

type (
	ParamsCreateUser struct {
		Email    string
		Name     string
		Password string
	}
)

func (i *sUserUsecase) CreateUser(p *ParamsCreateUser) (*ResultLogin, error) {

	checkEmail, _ := i.userRepository.FindByEmail(&p.Email)

	if checkEmail != nil {
		return nil, ErrEmailAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	data, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		Email:    p.Email,
		Name:     p.Name,
		Password: hashedPassword,
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
		Email:       p.Email,
		AccessToken: accessToken,
	}, nil
}
