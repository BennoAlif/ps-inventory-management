package v1routes

import (
	userv1controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/user"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userController := userv1controller.New(&userv1controller.V1User{
		DB: i.DB,
	})

	g.POST("/register", userController.Register)
	g.POST("/login", userController.Login)

}
