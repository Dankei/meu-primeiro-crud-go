package model

import "github.com/Dankei/meu-primeiro-crud-go.git/src/configuration/rest_err"

func (*UserDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}