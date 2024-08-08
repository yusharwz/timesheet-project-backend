package router

import (
	"final-project-enigma/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup) {

	//auth controller
	controller.NewAuthController(v1Group)

	//admin controller
	controller.NewAdminController(v1Group)

	//user controller
	controller.NewAccountController(v1Group)

	//work controller
	controller.NewWorkController(v1Group)

	//timesheet controller
	controller.NewTimesheetController(v1Group)
}
