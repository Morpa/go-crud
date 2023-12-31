package service

import (
	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"github.com/Morpa/go-crud/src/api/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomain.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
