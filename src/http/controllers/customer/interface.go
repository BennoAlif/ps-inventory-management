package customerv1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Customer struct {
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

type iV1Customer interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
}

func New(v1Customer *V1Customer) iV1Customer {
	return v1Customer
}
