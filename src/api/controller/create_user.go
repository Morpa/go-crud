package controller

import (
	"net/http"

	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/configuration/validation"
	"github.com/Morpa/go-crud/src/api/controller/model/request"
	"github.com/Morpa/go-crud/src/api/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init create_user controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
