package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luiisp/go-uplift/src/configuration/rest_err"
	"github.com/luiisp/go-uplift/src/controller/model/request"
)

func AddNewUser(c *gin.Context) {

	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		rest_err := rest_err.NewBadRequestError(
			fmt.Sprintf("Incorrect Fields, error=%s\n", err.Error()))

		c.JSON(rest_err.Code, rest_err)
		return
	}

	fmt.Println(UserRequest)
}
