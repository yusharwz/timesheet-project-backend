package service

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
)

type AuthService interface {
	RegisterAccount(req request.RegisterAccountRequest) (resp response.RegisterAccountResponse, err error)
	Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error)
	GetRoleById(id string) (*response.GetRoleResponse, error)
}
