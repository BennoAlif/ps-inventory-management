package v1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1) PingAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: "OK",
		Data:    "PONG AUTH",
	})
}
