package controller

import (
	"fmt"
	"log"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/validation"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil{
		log.Printf("Error trying to marshal object, error=%s\n",err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code,errRest)
		return
	}

	fmt.Println(UserRequest)

}