package impl

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/repository/impl"
)

type AccountService struct{}

var accountRepository = impl.NewAccountRepository()

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (AccountService) AccountActivationUrl(account request.ActivateAccountRequest) error {

	err := accountRepository.AccountActivation(account.Email)
	if err != nil {
		return err
	}

	return nil
}

func (AccountService) EditAccount(req request.EditAccountRequest, authHeader string) (response.AccountDetailResponse, error) {

	userId, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return response.AccountDetailResponse{}, err
	}
	req.UserID = userId

	err = accountRepository.EditAccount(req)
	if err != nil {
		return response.AccountDetailResponse{}, err
	}

	account, user, role, err := accountRepository.DetailAccount(req.UserID)
	if err != nil {
		return response.AccountDetailResponse{}, err
	}

	resp := response.AccountDetailResponse{
		Username: account.Username,
		Name:     user.Name,
		Email:    account.Email,
		Phone:    user.PhoneNumber,
		Role:     role.RoleName,
		IsActive: account.IsActive,
	}

	return resp, nil
}

func (AccountService) ChangePassword(req request.ChangePasswordRequest, authHeader string) error {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return err
	}
	req.NewPassword = userID

	err = accountRepository.ChangePassword(req)
	if err != nil {
		return err
	}

	return nil
}
