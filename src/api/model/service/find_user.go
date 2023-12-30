package service

import (
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"github.com/Morpa/go-crud/src/api/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr,
) {
	return nil, nil
}
