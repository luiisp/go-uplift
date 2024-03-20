package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luiisp/go-uplift/src/configuration/validation"
	"github.com/luiisp/go-uplift/src/controller/model/request"
)

func AddNewUser(c *gin.Context) {

	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		errRest := validation.ValidateUsrError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(UserRequest)
}
