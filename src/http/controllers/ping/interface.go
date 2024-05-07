package v1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1 struct {
	DB *sql.DB
}

type iV1 interface {
	Ping(c echo.Context) error
	PingAuth(c echo.Context) error
}

func New(v1 *V1) iV1 {
	return v1
}
