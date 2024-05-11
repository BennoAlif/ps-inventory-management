package productv1controller

import (
	"net/http"

	orderrepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
	productusecase "github.com/BennoAlif/ps-cats-social/src/usecase/product"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) Delete(c echo.Context) (err error) {
	id := c.Param("id")

	uu := productusecase.New(
		productrepository.New(i.DB),
		orderrepository.New(i.DB),
	)

	err = uu.Delete(&id)

	if err != nil {
		if err == productusecase.ErrProductNotFound {
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
		Message: "Product deleted successfully",
		Data:    nil,
	})
}
