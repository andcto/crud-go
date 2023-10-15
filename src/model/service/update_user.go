package service

import (
	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/configuration/rest_err"
	"github.com/andcto/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return nil
	}

	logger.Info("Init updateUser model",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
