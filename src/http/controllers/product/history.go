package productv1controller

import (
	"net/http"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	orderrepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
	orderusecase "github.com/BennoAlif/ps-cats-social/src/usecase/order"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) History(c echo.Context) (err error) {
	filters := &entities.SearchOrderFilter{}

	// customerID
	if customerID := c.QueryParam("customerId"); customerID != "" {
		filters.CustomerID = customerID
	}

	// limit
	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}

	// offset
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}

	// createdAt
	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		filters.CreatedAt = createdAt
	}

	uu := orderusecase.New(
		orderrepository.New(i.DB),
	)

	orders, err := uu.FindMany(filters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Success",
		Data:    orders,
	})
}
