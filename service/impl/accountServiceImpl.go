package impl

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/middleware"
	"final-project-enigma/repository"
	"final-project-enigma/repository/impl"
)

type AccountService struct{}

var accountRepository repository.AccountRepository = impl.NewAccountRepository()

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (AccountService) AccountActivationUrl(account request.ActivateAccountRequest) error {

	err := accountRepository.AccountActivation(account.Email, account.Password)
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

	account, user, role, err := adminRepository.DetailAccount(req.UserID)
	if err != nil {
		return response.AccountDetailResponse{}, err
	}

	resp := response.AccountDetailResponse{
		Name:     user.Name,
		Email:    account.Email,
		Phone:    user.PhoneNumber,
		Role:     role.RoleName,
		IsActive: account.IsActive,
	}

	return resp, nil
}

func (AccountService) UploadSignature(req request.UploadImagesRequest, authHeader string) (resp response.UploadImageResponse, err error) {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return resp, err
	}
	req.UserID = userID

	resp, err = accountRepository.UserUploadSignatureIMG(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (AccountService) ChangePassword(req request.ChangePasswordRequest, authHeader string) error {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return err
	}

	err = accountRepository.ChangePassword(userID, req)
	if err != nil {
		return err
	}

	return nil
}

func (AccountService) GetAccountDetail(authHeader string) (*response.AccountUserResponse, error) {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		return nil, err
	}

	account, user, err := accountRepository.GetAccountDetailByUserID(userID)
	if err != nil {
		return nil, err
	}

	accountUserResp := response.AccountUserResponse{
		AccountID:    account.ID,
		Email:        account.Email,
		IsActive:     account.IsActive,
		UserID:       user.ID,
		Name:         user.Name,
		Phone:        user.PhoneNumber,
		SignatureURL: user.Signature,
	}

	return &accountUserResp, nil
}

func (AccountService) GetAccountByID(id string) (*response.AccountUserResponse, error) {
	account, user, err := accountRepository.GetAccountDetailByUserID(id)
	if err != nil {
		return nil, err
	}

	accountUserResp := response.AccountUserResponse{
		AccountID:    account.ID,
		Email:        account.Email,
		IsActive:     account.IsActive,
		UserID:       user.ID,
		Name:         user.Name,
		Phone:        user.PhoneNumber,
		SignatureURL: user.Signature,
	}

	return &accountUserResp, nil
}
