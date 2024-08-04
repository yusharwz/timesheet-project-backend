package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authService = impl.NewAuthService()

func NewAuthController(g *gin.RouterGroup) {
	controller := new(AuthController)

	usersGroup := g.Group("/")
	{
		usersGroup.POST("/login", controller.AccountLogin)
	}

	adminGroup := g.Group("/admin", middleware.JwtAuthWithRoles("admin"))
	{
		adminGroup.POST("/register", controller.RegisterAccountRequest)
	}
}

func (AuthController) RegisterAccountRequest(ctx *gin.Context) {
	var req request.RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError)
			return
		}
		response.NewResponseError(ctx, "json request body required")
		return
	}

	resp, err := authService.RegisterAccount(req)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp)
}

func (AuthController) AccountLogin(ctx *gin.Context) {

	var req request.LoginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError)
			return
		}
		response.NewResponseError(ctx, "json request body required")
		return
	}
	resp, err := authService.Login(req)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp)
}
