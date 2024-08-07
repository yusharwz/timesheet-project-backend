package repository

import (
	"timesheet-app/dto/request"
	"timesheet-app/dto/response"
	"timesheet-app/entity"
)

type AuthRepository interface {
	Register(user entity.User, account entity.Account) (entity.User, entity.Account, error)
	Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error)
	GetRoleByName(roleName string) (entity.Role, error)
	GetRoleById(id string) (*entity.Role, error)
}
