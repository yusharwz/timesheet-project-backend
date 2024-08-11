package impl

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/helper"
	"timesheet-app/middleware"
	"timesheet-app/repository"
	"timesheet-app/repository/impl"

	"github.com/rs/zerolog/log"
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
		log.Error().Msg(err.Error())
		return response.AccountDetailResponse{}, err
	}
	req.UserID = userId

	err = accountRepository.EditAccount(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.AccountDetailResponse{}, err
	}

	account, user, role, err := adminRepository.DetailAccount(req.UserID)
	if err != nil {
		log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
		return resp, err
	}
	req.UserID = userID

	resp, err = accountRepository.UserUploadSignatureIMG(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	return resp, nil
}

func (AccountService) ChangePassword(req request.ChangePasswordRequest, authHeader string) error {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	err = accountRepository.ChangePassword(userID, req)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (AccountService) GetAccountDetail(authHeader string) (*response.AccountUserResponse, error) {

	userID, err := middleware.GetIdFromToken(authHeader)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	account, user, err := accountRepository.GetAccountDetailByUserID(userID)
	if err != nil {
		log.Error().Msg(err.Error())
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
		log.Error().Msg(err.Error())
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

func (AccountService) ForgetPassword(req request.ForgetPasswordRequest) error {

	newPassword, err := helper.GenerateCode()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	hashedPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	req.NewPassword = hashedPassword

	if err = accountRepository.ForgetPassword(req); err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if err = helper.SendNewPassword(req.Email, newPassword); err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}
