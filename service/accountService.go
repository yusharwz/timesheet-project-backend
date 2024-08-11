package service

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
)

type AccountService interface {
	AccountActivationUrl(account request.ActivateAccountRequest) error
	UploadSignature(req request.UploadImagesRequest, authHeader string) (resp response.UploadImageResponse, err error)
	EditAccount(req request.EditAccountRequest, authHeader string) (response.AccountDetailResponse, error)
	ChangePassword(req request.ChangePasswordRequest, authHeader string) error
	GetAccountDetail(authHeader string) (*response.AccountUserResponse, error)
	GetAccountByID(id string) (*response.AccountUserResponse, error)
	ForgetPassword(req request.ForgetPasswordRequest) error
}
