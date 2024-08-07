package repository

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
)

type AuthRepository interface {
	Register(user entity.User, account entity.Account) (entity.User, entity.Account, error)
	Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error)
	GetRoleByName(roleName string) (entity.Role, error)
	GetRoleById(id string) (*entity.Role, error)
}
