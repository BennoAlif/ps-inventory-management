package v1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (i *V1) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: "OK",
		Data:    "PONG",
	})
}
