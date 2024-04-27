package view

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/controller/model/response"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface,) response.UserResponse{
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}