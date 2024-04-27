package service

import (
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/logger"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"
	"github.com/Dankei/meu-primeiro-crud-go.git/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("jorney","createUser"))

	userDomain.EncryptPassword()

	return nil
}
