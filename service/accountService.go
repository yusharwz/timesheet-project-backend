package service

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
)

type AccountService interface {
	AccountActivationUrl(account request.ActivateAccountRequest) error
	EditAccount(req request.EditAccountRequest, authHeader string) (response.AccountDetailResponse, error)
	ChangePassword(req request.ChangePasswordRequest, authHeader string) error
	GetAccountDetail(authHeader string) (response.AccountUserResponse, error)
}
