package main

import (
	"context"
	"log"

	mongodb "github.com/andcto/crud-go/src/configuration/database"
	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/controller"
	"github.com/andcto/crud-go/src/controller/routes"
	"github.com/andcto/crud-go/src/model/repository"
	"github.com/andcto/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro loading .env files")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	//Inicializar dependencias
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userControler := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
