package main

import (
	"github.com/Morpa/go-crud/src/api/controller"
	"github.com/Morpa/go-crud/src/api/model/repository"
	"github.com/Morpa/go-crud/src/api/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependecies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
