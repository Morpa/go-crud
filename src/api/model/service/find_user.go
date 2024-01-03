package service

import (
	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"github.com/Morpa/go-crud/src/api/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init findUserByID services",
		zap.String("journey", "findUserByID"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail services",
		zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init findUserByEmail services",
		zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
