package converter

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email: 		domain.GetEmail(),
		Password: 	domain.GetPassword(),
		Name: 		domain.GetName(),
		Age:		domain.GetAge(),
	}
}