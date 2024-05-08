package productv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
	productusecase "github.com/BennoAlif/ps-cats-social/src/usecase/product"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type (
	meValidator struct {
		ID int `mapstructure:"user_id" validate:"required"`
	}
)

func (i *V1Product) Create(c echo.Context) (err error) {
	u := new(createRequest)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	if !ValidateRace(u.Category) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid race",
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
	)

	data, err := uu.Create(&entities.ParamsCreateProduct{
		Name:        u.Name,
		Sku:         u.Sku,
		Category:    u.Category,
		Description: u.Description,
		ImageUrl:    u.ImageUrl,
		Notes:       u.Notes,
		Price:       u.Price,
		Stock:       u.Stock,
		Location:    u.Location,
		IsAvailable: u.IsAvailable,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "success",
		Data:    data,
	})
}

type (
	createRequest struct {
		Name        string `json:"name" validate:"required,min=1,max=30"`
		Sku         string `json:"sku" validate:"required,min=1,max=30"`
		Category    string `json:"category" validate:"required"`
		Description string `json:"description" validate:"required,min=1,max=200"`
		ImageUrl    string `json:"imageUrs" validate:"required,url"`
		Notes       string `json:"notes" validate:"required,min=1,max=200"`
		Price       int    `json:"price" validate:"required,min=1"`
		Stock       int    `json:"stock" validate:"required,min=0,max=100000"`
		Location    string `json:"location" validate:"required,min=1,max=200"`
		IsAvailable bool   `json:"isAvailable" validate:"required,oneof=true false"`
	}
)

func ValidateRace(race string) bool {
	validCategory := map[string]bool{
		"Clothing":    true,
		"Accessories": true,
		"Footwear":    true,
		"Beverages":   true,
	}

	_, ok := validCategory[race]
	return ok
}
