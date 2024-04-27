package converter

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository/entity"
)

func ConvertEntidyToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain:= model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
	
}