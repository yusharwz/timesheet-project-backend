package impl

import (
	"errors"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/helper"
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

func (AccountService) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	resp, err = accountRepository.Login(req)
	if err != nil {
		return resp, err
	}

	err = helper.ComparePassword(resp.HashPassword, req.Password)
	if err != nil {
		return resp, errors.New("invalid email or password")
	}

	resp.Token, err = helper.GetTokenJwt(resp.UserId, resp.Username, resp.Email, resp.Role)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (AccountService) RetrieveAccountList() ([]response.ListAccountResponse, error) {

	accounts, users, err := accountRepository.RetrieveAccountList()
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]string)
	for _, user := range users {
		userMap[user.ID] = user.Name
	}

	var resp []response.ListAccountResponse
	for _, account := range accounts {
		var status string
		if account.IsActive {
			status = "Active"
		} else {
			status = "Inactive"
		}

		resp = append(resp, response.ListAccountResponse{
			Username: userMap[account.UserID],
			Email:    account.Email,
			Status:   status,
		})
	}

	return resp, nil
}

func (AccountService) DetailAccount(userID string) (response.AccountDetailResponse, error) {
	var resp response.AccountDetailResponse

	account, user, role, err := accountRepository.DetailAccount(userID)
	if err != nil {
		return resp, err
	}

	resp.Username = account.Username
	resp.Name = user.Name
	resp.Email = account.Email
	resp.Phone = user.PhoneNumber
	resp.Role = role.RoleName
	resp.IsActive = account.IsActive

	return resp, nil
}

func (AccountService) SoftDeleteAccount(userID string) error {
	return accountRepository.SoftDeleteAccount(userID)
}
