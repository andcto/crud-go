package controller

import (
	"net/http"

	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
	}

	logger.Info("deleteUser controller executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
