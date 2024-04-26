package controller

import (
	"fmt"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s",err.Error()))

		c.JSON(restErr.Code,restErr)
		return
	}

	fmt.Println(UserRequest)

}