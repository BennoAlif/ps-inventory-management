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

func (i *V1Product) Update(c echo.Context) (err error) {
	u := new(createRequest)
	id := c.Param("id")

	uu := productusecase.New(
		productrepository.New(i.DB),
		orderrepository.New(i.DB),
		customerrepository.New(i.DB),
	)

	// err = uu.IsExist(&id)

	// if err != nil {
	// 	if err == productusecase.ErrProductNotFound {
	// 		return c.JSON(http.StatusNotFound, ErrorResponse{
	// 			Status:  false,
	// 			Message: err.Error(),
	// 		})
	// 	}
	// }

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if !isValidURL(u.ImageUrl) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid image URL",
		})
	}

	if !ValidateCategory(u.Category) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid Category",
		})
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = uu.Update(&id, &entities.ParamsUpdateProduct{
		Name:        u.Name,
		Sku:         u.Sku,
		Category:    u.Category,
		Notes:       u.Notes,
		ImageUrl:    u.ImageUrl,
		Price:       u.Price,
		Stock:       *u.Stock,
		Location:    u.Location,
		IsAvailable: *u.IsAvailable,
	})

	if err != nil {
		// if err == productusecase.ErrProductNotFound {
		// 	return c.JSON(http.StatusNotFound, ErrorResponse{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// }
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "product updated successfully",
		Data:    nil,
	})
}
