package controller

import (
	"net/http"

	"github.com/andcto/crud-go/src/configuration/logger"
	"github.com/andcto/crud-go/src/configuration/validation"
	"github.com/andcto/crud-go/src/controller/request"
	"github.com/andcto/crud-go/src/model"
	"github.com/andcto/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created sucessfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
