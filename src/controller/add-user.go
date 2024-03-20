package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luiisp/go-uplift/src/configuration/logger"
	"github.com/luiisp/go-uplift/src/configuration/validation"
	"github.com/luiisp/go-uplift/src/controller/model/request"
	"go.uber.org/zap"
)

func AddNewUser(c *gin.Context) {
	logger.Info("New user as receive", zap.String("journey", "createuser"))
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error in try validate user infos", err,
			zap.String("journey", "createuser"))
		errRest := validation.ValidateUsrError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created", zap.String("journey", "createuser"))
}
