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

	accountGroup := g.Group("/accounts")
	{
		accountGroup.GET("/activate", controller.AccountActivation)
		accountGroup.PUT("/", middleware.JwtAuthWithRoles("user"), controller.EditAccount)
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

	response.NewResponseSuccess(ctx, nil, "account has been activated", "01", "01")
}

func (AccountController) EditAccount(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	var req request.EditAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError, "bad request", "01", "02")
			return
		}
		response.NewResponseError(ctx, "json request body required", "01", "02")
		return
	}
	resp, err := accountService.EditAccount(req, authHeader)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponseSuccess(ctx, resp, "update account success", "01", "01")
}
