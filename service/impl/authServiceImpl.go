package impl

import (
	"errors"
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
	"timesheet-app/helper"
	"timesheet-app/repository"
	"timesheet-app/repository/impl"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AuthService struct{}

var authRepository repository.AuthRepository = impl.NewAuthRepository()

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) RegisterAccount(req request.RegisterAccountRequest) (resp response.RegisterAccountResponse, err error) {

	code, err := helper.GenerateCode()
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	hashedPassword, err := helper.HashPassword(code)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	role, err := authRepository.GetRoleById(req.RoleId)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	newAccount := entity.Account{
		Base:     entity.Base{ID: uuid.NewString()},
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: false,
		RoleID:   role.ID,
	}

	user := entity.User{
		Base: entity.Base{ID: uuid.NewString()},
		Name: req.Name,
	}

	createdUser, createdAccount, err := authRepository.Register(user, newAccount)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	resp = response.RegisterAccountResponse{
		Email:    createdAccount.Email,
		Name:     createdUser.Name,
		RoleName: role.RoleName,
	}

	err = helper.SendEmailActivatedAccount(resp.Email, code, hashedPassword)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	return resp, nil
}

func (AuthService) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	resp, err = authRepository.Login(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	err = helper.ComparePassword(resp.HashPassword, req.Password)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, errors.New("invalid email or password")
	}

	resp.Token, err = helper.GetTokenJwt(resp.UserId, resp.Name, resp.Email, resp.Role)
	if err != nil {
		log.Error().Msg(err.Error())
		return resp, err
	}

	return resp, err
}

func (AuthService) GetRoleById(id string) (*response.GetRoleResponse, error) {
	result, err := authRepository.GetRoleById(id)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	roleResponse := response.GetRoleResponse{RoleName: result.RoleName}
	return &roleResponse, nil
}
