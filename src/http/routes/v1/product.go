package v1routes

import (
	productv1controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/product"
	"github.com/BennoAlif/ps-cats-social/src/http/middlewares"
)

func (i *V1Routes) MountProduct() {
	g := i.Echo.Group("/product")
	productController := productv1controller.New(&productv1controller.V1Product{
		DB: i.DB,
	})
	authMiddleware := middlewares.Authentication()

	g.POST("", productController.Create, authMiddleware)
	g.GET("", productController.Get, authMiddleware)
	g.PUT("/{id}", productController.Update, authMiddleware)
	g.DELETE("/{id}", productController.Delete, authMiddleware)
}
