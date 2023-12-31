package main

import (
	"context"
	"log"

	"github.com/Morpa/go-crud/src/api/configuration/database/mongodb"
	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependecies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
