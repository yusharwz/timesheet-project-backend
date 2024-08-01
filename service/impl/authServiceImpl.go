package impl

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"final-project-enigma/repository/impl"

	"github.com/google/uuid"
)

type AuthService struct{}

var authRepository = impl.NewAuthRepository()

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) RegisterAccount(req request.RegisterAccountRequest) (resp response.RegisterAccountResponse, err error) {

	newUser := entity.User{
		Base: entity.Base{ID: uuid.NewString()},
	}

	user, err := authRepository.CreateUser(newUser)
	if err != nil {
		return resp, err
	}

	code, err := helper.GenerateCode()
	if err != nil {
		return resp, err
	}

	hashedPassword, err := helper.HashPassword(code)
	if err != nil {
		return resp, err
	}

	req.Password = hashedPassword
	req.IsActive = false
	req.UserID = user.Base.ID

	newAccount := entity.Account{
		Base:     entity.Base{ID: uuid.NewString()},
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		IsActive: req.IsActive,
		RoleID:   req.RoleID,
		UserID:   req.UserID,
	}

	createdAccount, err := authRepository.CreateAccount(newAccount)
	if err != nil {
		return resp, err
	}

	resp = response.RegisterAccountResponse{
		Id:       createdAccount.ID,
		Email:    createdAccount.Email,
		Username: createdAccount.Username,
		IsActive: createdAccount.IsActive,
		RoleID:   createdAccount.RoleID,
		UserID:   createdAccount.UserID,
	}

	err = helper.SendEmailActivedAccount(resp.Email, resp.Username, code, hashedPassword)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
