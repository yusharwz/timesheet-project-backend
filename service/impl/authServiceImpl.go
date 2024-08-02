package impl

import (
	"errors"
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

func (AuthService) RegisterAccount(req request.RegisterAccountRequest) (response.RegisterAccountResponse, error) {
	var resp response.RegisterAccountResponse

	newUser := entity.User{
		Base: entity.Base{ID: uuid.NewString()},
	}

	code, err := helper.GenerateCode()
	if err != nil {
		return resp, err
	}

	hashedPassword, err := helper.HashPassword(code)
	if err != nil {
		return resp, err
	}

	newAccount := entity.Account{
		Base:     entity.Base{ID: uuid.NewString()},
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: false,
		RoleID:   req.RoleID,
		UserID:   newUser.ID,
	}

	_, createdAccount, err := authRepository.Register(newUser, newAccount)
	if err != nil {
		return resp, err
	}

	resp = response.RegisterAccountResponse{
		Id:       createdAccount.ID,
		Email:    createdAccount.Email,
		IsActive: createdAccount.IsActive,
		RoleID:   createdAccount.RoleID,
		UserID:   createdAccount.UserID,
	}

	err = helper.SendEmailActivedAccount(resp.Email, code, hashedPassword)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (AuthService) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	resp, err = authRepository.Login(req)
	if err != nil {
		return resp, err
	}

	err = helper.ComparePassword(resp.HashPassword, req.Password)
	if err != nil {
		return resp, errors.New("invalid email or password")
	}

	resp.Token, err = helper.GetTokenJwt(resp.UserId, resp.Name, resp.Email, resp.Role)
	if err != nil {
		return resp, err
	}

	return resp, err
}
