package model

import (
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.Password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
