package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB, client *resty.Client) {

}
