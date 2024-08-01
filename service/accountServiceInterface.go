package service

import "final-project-enigma/dto/request"

type AccountService interface {
	AccountActivationUrl(account request.ActivateAccountRequest) error
}
