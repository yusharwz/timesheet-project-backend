package router

import (
	"database/sql"
	"final-project-enigma/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB, client *resty.Client) {
	controller.NewAuthController(v1Group)
	controller.NewAccountController(v1Group)
}
