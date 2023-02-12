//go:build wireinject
// +build wireinject

package main

import (
	"github.com/alireza-mf/go-ecommerce/controllers"
	"github.com/alireza-mf/go-ecommerce/repositories"
	"github.com/alireza-mf/go-ecommerce/services"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func initUserAPI(db *mongo.Database) controllers.UserController {
	wire.Build(
		repositories.ProvideUserRepository,
		services.ProvideUserService,
		controllers.ProvideUserController,
	)

	return controllers.UserController{}
}
