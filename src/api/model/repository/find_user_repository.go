package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Morpa/go-crud/src/api/configuration/logger"
	"github.com/Morpa/go-crud/src/api/configuration/rest_err"
	"github.com/Morpa/go-crud/src/api/model"
	"github.com/Morpa/go-crud/src/api/model/repository/entity"
	"github.com/Morpa/go-crud/src/api/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByID repository",
		zap.String("journey", "FindUserByID"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByEmailAndPassword repository",
		zap.String("journey", "FindUserByEmailAndPassword"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByEmailAndPassword"),
			)
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "FindUserByEmailAndPassword"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
