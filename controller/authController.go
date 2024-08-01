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

	authGroup := g.Group("/admin")
	{
		authGroup.POST("/register", controller.RegisterAccountRequest)
	}
}

func (AuthController) RegisterAccountRequest(ctx *gin.Context) {
	var req request.RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponBadRequest(ctx, validationError, "bad request", "01", "02")
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

	response.NewResponSucces(ctx, nil, "create account succes, please check your email for activated your account", "01", "01")
}
