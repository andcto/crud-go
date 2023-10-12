package main

import (
	"log"

	mongodb "github.com/andcto/crud-go/src/configuration/database"
	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/controller"
	"github.com/andcto/crud-go/src/controller/routes"
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

	mongodb.InitConnection()

	//Inicializar dependencias
	service := service.NewUserDomainService()
	userControler := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
