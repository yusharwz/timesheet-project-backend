package router

import (
	"final-project-enigma/controller"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitRoute(g *gin.RouterGroup, client *resty.Client) {
	// work
	controller.NewWorkController(g)
}
