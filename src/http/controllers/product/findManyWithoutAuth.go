package productv1controller

import (
	"net/http"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	customerrepository "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
	orderrepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
	productusecase "github.com/BennoAlif/ps-cats-social/src/usecase/product"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) GetWithoutAuth(c echo.Context) (err error) {
	filters := &entities.ProductSearchFilter{}

	// id
	if id := c.QueryParam("id"); id != "" {
		filters.ID = id
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

	// name
	if name := c.QueryParam("name"); name != "" {
		filters.Name = name
	}

	// isAvailable
	filters.IsAvailable = true

	// category
	if category := c.QueryParam("category"); category != "" {
		if isValidCategory(category) {
			filters.Category = category
		}
	}

	// sku
	if sku := c.QueryParam("sku"); sku != "" {
		filters.SKU = sku
	}

	// price
	if price := c.QueryParam("price"); price != "" {
		if price == "asc" || price == "desc" {
			filters.Price = price
		}
	}

	// inStock
	if inStockStr := c.QueryParam("inStock"); inStockStr != "" {
		inStock, err := strconv.ParseBool(inStockStr)
		if err == nil {
			filters.InStock = inStock
		}
	}

	uu := productusecase.New(
		productrepository.New(i.DB),
		orderrepository.New(i.DB),
		customerrepository.New(i.DB),
	)

	data, err := uu.FindMany(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Product found successfully",
		Data:    data,
	})
}
