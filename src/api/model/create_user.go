package model

import (
	"fmt"

	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	hashedPassword, err := ud.EncryptPassword()
	if err != nil {
		logger.Error("Error to encrypt the password", err)
	}

	ud.Password = hashedPassword

	fmt.Println(ud)

	return nil
}
