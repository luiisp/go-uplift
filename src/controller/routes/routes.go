package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luiisp/go-uplift/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userId", controller.FindUserByID)
	r.GET("/userMail/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.AddNewUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)

}
