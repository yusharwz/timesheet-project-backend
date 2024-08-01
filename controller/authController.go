package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authService = impl.NewAuthService()

func NewAuthController(g *gin.RouterGroup) {
	controller := new(AuthController)

	usersGroup := g.Group("/accounts")
	{
		usersGroup.POST("/login", controller.AccountLogin)
	}

	adminroup := g.Group("/admin")
	{
		adminroup.POST("/register", controller.RegisterAccountRequest)
	}
}

func (AuthController) RegisterAccountRequest(ctx *gin.Context) {
	var req request.RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError, "bad request", "01", "02")
			return
		}
		response.NewResponseError(ctx, "json request body required", "01", "02")
		return
	}

	_, err := authService.RegisterAccount(req)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponseSuccess(ctx, nil, "create account succes, please check your email for activated your account", "01", "01")
}

func (AuthController) AccountLogin(ctx *gin.Context) {

	var req request.LoginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError, "bad request", "01", "02")
			return
		}
		response.NewResponseError(ctx, "json request body required", "01", "02")
		return
	}
	resp, err := authService.Login(req)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error(), "01", "01")
		return
	}

	response.NewResponseSuccess(ctx, resp, "logged in", "01", "01")
}
