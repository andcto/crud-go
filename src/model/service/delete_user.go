package service

import (
	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init deleteUser model",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return nil
	}

	logger.Info("Init deleteUser model",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
