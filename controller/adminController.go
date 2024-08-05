package controller

import (
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/service"
	"final-project-enigma/service/impl"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

var adminService service.AdminService = impl.NewAdminService()

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
	paging := ctx.DefaultQuery("paging", "1")
	rowsPerPage := ctx.DefaultQuery("rowsPerPage", "10")
	resp, totalRows, totalPage, err := adminService.RetrieveAccountList(paging, rowsPerPage)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccessPaging(ctx, resp, paging, rowsPerPage, totalRows, totalPage)
}

func (AdminController) AccountDetail(ctx *gin.Context) {

	userID := ctx.Param("id")
	resp, err := adminService.DetailAccount(userID)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp)
}

func (AdminController) AccountSoftDelete(ctx *gin.Context) {

	userID := ctx.Param("id")
	err := adminService.SoftDeleteAccount(userID)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, nil)
}
