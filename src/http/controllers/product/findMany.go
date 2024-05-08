package productv1controller

import (
	"net/http"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
	productusecase "github.com/BennoAlif/ps-cats-social/src/usecase/product"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) Get(c echo.Context) (err error) {
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
	if isAvailableStr := c.QueryParam("isAvailable"); isAvailableStr != "" {
		isAvailable, err := strconv.ParseBool(isAvailableStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'isAvailable'",
			})
		}
		filters.IsAvailable = isAvailable
	}

	// category
	if category := c.QueryParam("category"); category != "" {
		if !isValidCategory(category) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'category'",
			})
		}
		filters.Category = category
	}

	// sku
	if sku := c.QueryParam("sku"); sku != "" {
		filters.SKU = sku
	}

	// price
	if price := c.QueryParam("price"); price != "" {
		if price != "asc" && price != "desc" {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'price'",
			})
		}
		filters.Price = price
	}

	// inStock
	if inStockStr := c.QueryParam("inStock"); inStockStr != "" {
		inStock, err := strconv.ParseBool(inStockStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'inStock'",
			})
		}
		filters.InStock = inStock
	}

	// createdAt
	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt != "asc" && createdAt != "desc" {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'createdAt'",
			})
		}
		filters.CreatedAt = createdAt
	}

	uu := productusecase.New(
		productrepository.New(i.DB),
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

func isValidCategory(category string) bool {
	validCategories := []string{"Clothing", "Accessories", "Footwear", "Beverages"}
	for _, validCategory := range validCategories {
		if category == validCategory {
			return true
		}
	}
	return false
}
