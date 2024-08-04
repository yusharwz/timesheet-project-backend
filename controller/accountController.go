package controller

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/service"
	"final-project-enigma/service/impl"
	"final-project-enigma/utils"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

var accountService service.AccountService = impl.NewAccountService()

func NewAccountController(g *gin.RouterGroup) {
	controller := new(AccountController)

	accountGroup := g.Group("/accounts", middleware.JwtAuthWithRoles("user"))
	{
		accountGroup.GET("/profile", controller.GetAccountDetailByUserID)
		accountGroup.POST("/profile/upload-signature", controller.UploadSignature)
		accountGroup.PUT("/", controller.EditAccount)
		accountGroup.PUT("/change-password", controller.ChangePassword)
	}
	g.GET("accounts/activate", controller.AccountActivation)
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

	response.NewResponseSuccess(ctx, nil)
}

func (AccountController) EditAccount(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	var req request.EditAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError)
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

	response.NewResponseSuccess(ctx, resp)
}

func (AccountController) UploadSignature(ctx *gin.Context) {
	var req request.UploadImagesRequest
	authHeader := ctx.GetHeader("Authorization")
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		response.NewResponseError(ctx, "failed to get file")
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		response.NewResponseError(ctx, "failed to open file")
		return
	}
	req.SignatureImage = file
	resp, err := accountService.UploadSignature(req, authHeader)
	if err != nil {
		response.NewResponseError(ctx, err.Error())
		return
	}
	response.NewResponseSuccess(ctx, resp)

}

func (AccountController) ChangePassword(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	var req request.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := utils.GetValidationError(err)

		if len(validationError) > 0 {
			response.NewResponseBadRequest(ctx, validationError)
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

	response.NewResponseSuccess(ctx, nil)
}

func (AccountController) GetAccountDetailByUserID(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	resp, err := accountService.GetAccountDetail(authHeader)
	if err != nil {
		response.NewResponseForbidden(ctx, err.Error())
		return
	}

	response.NewResponseSuccess(ctx, resp)
}
