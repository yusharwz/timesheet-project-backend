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
		accountGroup.GET("/profile", middleware.JwtAuthWithRoles("user"), controller.GetAccountDetailByUserID)
		accountGroup.PUT("/", middleware.JwtAuthWithRoles("user"), controller.EditAccount)
		accountGroup.PUT("/change-password", middleware.JwtAuthWithRoles("user"), controller.ChangePassword)
	}
}
func (AccountController) AccountActivation(ctx *gin.Context) {

	var params request.ActivateAccountRequest

	params.Email = ctx.Query("e")
	params.Password = ctx.Query("unique")

	err := accountService.AccountActivationUrl(params)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, nil, "account has been activated")
}

func (AccountController) EditAccount(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	var req request.EditAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
		response.NewResponseError(ctx, "json request body required")
		return
	}
	resp, err := accountService.EditAccount(req, authHeader)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp, "update account success")
}

func (AccountController) ChangePassword(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	var req request.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
		response.NewResponseError(ctx, "json request body required")
		return
	}
	err := accountService.ChangePassword(req, authHeader)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, nil, "update password success")
}

func (AccountController) GetAccountDetailByUserID(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	resp, err := accountService.GetAccountDetail(authHeader)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp, "get data detail success")
}
