package repository

import (
	"context"
	"os"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const(

	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"

)



func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr){
	

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)


	collection := ur.databaseConnection.Collection(collection_name)

	 value := converter.ConvertDomainToEntity(userDomain)
	
	result , err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	

	return converter.ConvertEntidyToDomain(*value), nil
}