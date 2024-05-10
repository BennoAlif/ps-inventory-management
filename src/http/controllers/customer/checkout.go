package customerv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	customerrepository "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
	customerUsecase "github.com/BennoAlif/ps-cats-social/src/usecase/customer"
	"github.com/labstack/echo/v4"
)

func (i *V1Customer) Checkout(c echo.Context) (err error) {
	u := new(entities.ParamsCustomerCheckout)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	_ = customerUsecase.New(
		customerrepository.New(i.DB),
	)

	// data, err := cu.Create(u)

	// if err != nil {
	// 	if err == helpers.ErrBadFormatPhoneNumber {
	// 		return c.JSON(http.StatusBadRequest, ErrorResponse{
	// 			Status:  false,
	// 			Message: err.Error(),
	// 		})
	// 	}
	// 	return c.JSON(http.StatusConflict, ErrorResponse{
	// 		Status:  false,
	// 		Message: err.Error(),
	// 	})
	// }

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "successfully checkout product",
		// Data:    interface{},
	})
}
