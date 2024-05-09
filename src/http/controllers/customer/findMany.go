package customerv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	customerrepository "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
	customerusecase "github.com/BennoAlif/ps-cats-social/src/usecase/customer"
	"github.com/labstack/echo/v4"
)

func (i *V1Customer) Get(c echo.Context) (err error) {
	filters := &entities.ParamsCustomer{}

	// id
	if id := c.QueryParam("id"); id != "" {
		filters.ID = id
	}

	// name
	if name := c.QueryParam("name"); name != "" {
		filters.Name = name
	}

	// phoneNumber
	if phoneNumber := c.QueryParam("phoneNumber"); phoneNumber != "" {
		filters.PhoneNumber = phoneNumber
	}

	cu := customerusecase.New(
		customerrepository.New(i.DB),
	)

	customers, err := cu.FindMany(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Product found successfully",
		Data:    customers,
	})
}
