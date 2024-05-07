package v1routes

import (
	v1controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/ping"
	"github.com/BennoAlif/ps-cats-social/src/http/middlewares"
)

func (i *V1Routes) MountPing() {
	g := i.Echo.Group("/ping")

	pingController := v1controller.New(&v1controller.V1{
		DB: i.DB,
	})

	g.GET("", pingController.Ping)
	g.GET("/auth", pingController.PingAuth, middlewares.Authentication([]string{}))
	g.GET("/customer", pingController.PingAuth, middlewares.Authentication([]string{"CUSTOMER"}))
	g.GET("/sales", pingController.PingAuth, middlewares.Authentication([]string{"SALES"}))
}
