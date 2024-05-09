package userv1controller

import (
	"net/http"

	userUsecase "github.com/BennoAlif/ps-cats-social/src/usecase/user"

	userRepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	"github.com/labstack/echo/v4"
)

func (i *V1User) Login(c echo.Context) (err error) {
	u := new(loginRequest)

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

	data, err := uu.Login(&userUsecase.ParamsLogin{
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
	})

	if err != nil {
		if err == userUsecase.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}

type (
	loginRequest struct {
		Password    string `json:"password" validate:"required"`
		PhoneNumber string `json:"phoneNumber" validate:"required"`
	}
)
