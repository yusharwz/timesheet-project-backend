package controller

import (
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/service/impl"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

var adminService = impl.NewAdminService()

func NewAdminController(g *gin.RouterGroup) {
	controller := new(AdminController)

	adminGroup := g.Group("/admin")
	{
		adminGroup.GET("/accounts", middleware.JwtAuthWithRoles("admin"), controller.AccountList)
		adminGroup.GET("/accounts/detail/:id", middleware.JwtAuthWithRoles("admin"), controller.AccountDetail)
		adminGroup.DELETE("/accounts/delete/:id", middleware.JwtAuthWithRoles("admin"), controller.AccountSoftDelete)

	}
}

func (AdminController) AccountList(ctx *gin.Context) {

	resp, err := adminService.RetrieveAccountList()
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp, "success get account list")
}

func (AdminController) AccountDetail(ctx *gin.Context) {

	userID := ctx.Param("id")
	resp, err := adminService.DetailAccount(userID)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp, "success get detail account")
}

func (AdminController) AccountSoftDelete(ctx *gin.Context) {

	userID := ctx.Param("id")
	err := adminService.SoftDeleteAccount(userID)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, nil, "delete account success")
}
