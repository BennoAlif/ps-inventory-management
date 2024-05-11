package productv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	customerrepository "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
	orderrepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
	productusecase "github.com/BennoAlif/ps-cats-social/src/usecase/product"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) Checkout(c echo.Context) (err error) {
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

	uu := productusecase.New(
		productrepository.New(i.DB),
		orderrepository.New(i.DB),
		customerrepository.New(i.DB),
	)

	err = uu.Checkout(u)
	if err != nil {
		if err == productusecase.ErrCustomerNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		if err == productusecase.ErrProductNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		if err == productusecase.ErrTotalPriceNotMatch {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		if err == productusecase.ErrChangeNotMatch {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		if err == productusecase.ErrStockNotAvailable {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "successfully checkout product",
	})
}
