package main

import (
	"context"
	"log"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/database/mongodb"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start user application, Traduzindo ta funcionando caralho")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil{
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}


	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup,userController)

	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}

}