package main

import (
	"log"

	"github.com/andcto/crud-go/src/controller"
	"github.com/andcto/crud-go/src/controller/routes"
	"github.com/andcto/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro loading .env files")
	}

	//Inicializar dependencias
	service := service.NewUserDomainService()
	userControler := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
