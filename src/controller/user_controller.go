package controller

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/service"
	"github.com/gin-gonic/gin"
)


func NewUserControllerInterface(
	serviceInteface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInteface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}