package main

import (
	"github.com/andcto/crud-go/src/controller"
	"github.com/andcto/crud-go/src/model/repository"
	"github.com/andcto/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	//Inicializar dependencias
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
