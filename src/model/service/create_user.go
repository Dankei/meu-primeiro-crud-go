package service

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface,*rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("jorney","createUser"))

	userDomain.EncryptPassword()

	userDomainRepository,err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", 
		err,
		zap.String("jorney","createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
	zap.String("userId", userDomain.GetID()), 
	zap.String("jorney","createUser"))
	return userDomainRepository,nil
}
