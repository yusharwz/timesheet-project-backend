package router

import (
	"final-project-enigma/controller"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitRoute(v1Group *gin.RouterGroup, client *resty.Client) {
	// work
	controller.NewWorkController(v1Group)
}
