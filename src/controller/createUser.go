package controller

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	err := rest_err.NewBadRequestError("Você erroou")
	c.JSON(err.Code,err)


}