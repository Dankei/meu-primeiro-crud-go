package main

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface{
		//Init dependencies
		repo := repository.NewUserRepository(database)
		service := service.NewUserDomainService(repo)
		return controller.NewUserControllerInterface(service)
	
}