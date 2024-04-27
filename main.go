package main

import (
	"log"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller/routes"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start user application, Traduzindo ta funcionando caralho")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)


	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup,userController)

	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}

}