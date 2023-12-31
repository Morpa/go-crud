package service

import (
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"github.com/Morpa/go-crud/src/api/model"
	"github.com/Morpa/go-crud/src/api/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(id string) (
		model.UserDomainInterface,
		*rest_err.RestErr,
	)

	FindUserByEmailServices(email string) (
		model.UserDomainInterface,
		*rest_err.RestErr,
	)

	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
