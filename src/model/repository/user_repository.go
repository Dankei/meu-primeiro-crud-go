package repository

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const(

	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"

)

func NewUserRepository(
	database *mongo.Database,
) UserRepository{
	return &userRepository{
		database,
	}
}

type userRepository struct{
	databaseConnection *mongo.Database

}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}