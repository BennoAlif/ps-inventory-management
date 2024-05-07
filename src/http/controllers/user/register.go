package userv1controller

import (
	"net/http"

	userUsecase "github.com/BennoAlif/ps-cats-social/src/usecase/user"
	"github.com/labstack/echo/v4"

	userRepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

func (i *V1User) Register(c echo.Context) (err error) {
	u := new(createRequest)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	uu := userUsecase.New(
		userRepository.New(i.DB),
	)

	data, err := uu.CreateUser(&userUsecase.ParamsCreateUser{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    data,
	})
}

type (
	createRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)
