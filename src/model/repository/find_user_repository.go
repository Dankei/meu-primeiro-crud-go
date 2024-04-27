package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository/entity"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr){

	
	logger.Info("Init findUserByEmail repository",
	zap.String("jorney","findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key:"email", Value:email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)
	
	if err != nil{
		if err == mongo.ErrNoDocuments{
			errorMessage := fmt.Sprintf("User not found with this email: %s",email)
			logger.Error(errorMessage,
			err,
			zap.String("jorney","findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("jorney","findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("email", userEntity.Email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("jorney","findUserByEmail"))
	return converter.ConvertEntidyToDomain(*userEntity),nil
}





func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr){

	
	logger.Info("Init findUserByID repository",
	zap.String("jorney","findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key:"_id", Value:id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)
	
	if err != nil{
		if err == mongo.ErrNoDocuments{
			errorMessage := fmt.Sprintf("User not found with this id: %s",id)
			logger.Error(errorMessage,
			err,
			zap.String("jorney","findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("jorney","findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("jorney","findUserByID"))
	return converter.ConvertEntidyToDomain(*userEntity),nil
}