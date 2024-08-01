package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

var accountService = impl.NewAccountService()

func NewAccountController(g *gin.RouterGroup) {
	controller := new(AccountController)

	adminGroup := g.Group("/admin")
	{
		adminGroup.GET("/accounts", middleware.JwtAuthWithRoles("user"), controller.AccountList)
		adminGroup.GET("/accounts/detail/:id", middleware.JwtAuthWithRoles("user"), controller.AccountDetail)

	}

	accountGroup := g.Group("/accounts")
	{
		accountGroup.GET("/activate", controller.AccountActivation)
		accountGroup.POST("/login", controller.AccountLogin)
	}
}
func (AccountController) AccountActivation(ctx *gin.Context) {

	var params request.ActivateAccountRequest

	params.Email = ctx.Query("e")
	params.Username = ctx.Query("un")
	params.Password = ctx.Query("unique")

	err := accountService.AccountActivationUrl(params)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponSucces(ctx, nil, "account has been activated", "01", "01")
}

func (AccountController) AccountLogin(ctx *gin.Context) {

	var req request.LoginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponBadRequest(ctx, validationError, "bad request", "01", "02")
			return
		}
		response.NewResponseError(ctx, "json request body required", "01", "02")
		return
	}
	resp, err := accountService.Login(req)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponSucces(ctx, resp, "logged in", "01", "01")
}

func (AccountController) AccountList(ctx *gin.Context) {

	resp, err := accountService.RetrieveAccountList()
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponSucces(ctx, resp, "success get account list", "01", "01")
}

func (AccountController) AccountDetail(ctx *gin.Context) {

	userID := ctx.Param("id")
	resp, err := accountService.DetailAccount(userID)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponSucces(ctx, resp, "success get detail account", "01", "01")
}
