package productv1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Product struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Product interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Checkout(c echo.Context) error
	History(c echo.Context) error
}

func New(v1Product *V1Product) iV1Product {
	return v1Product
}
