package main

import (
	"log"

	"github.com/Morpa/go-crud/src/api/configuration/database/mongodb"
	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/controller"
	"github.com/Morpa/go-crud/src/api/controller/routes"
	"github.com/Morpa/go-crud/src/api/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	//Init dependecies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
