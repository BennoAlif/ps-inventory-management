package v1routes

import (
	customerv1controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/customer"
	"github.com/BennoAlif/ps-cats-social/src/http/middlewares"
)

func (i *V1Routes) MountCustomer() {
	g := i.Echo.Group("/customer")

	customerController := customerv1controller.New(&customerv1controller.V1Customer{
		DB: i.DB,
	})

	g.GET("", customerController.Get, middlewares.Authentication())
	g.POST("/register", customerController.Create, middlewares.Authentication())
	g.POST("/checkout", customerController.Checkout, middlewares.Authentication())
}
