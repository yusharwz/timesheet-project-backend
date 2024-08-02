package repository

import (
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
)

type AuthRepository interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	CreateUser(user entity.User) (entity.User, error)
	Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error)
	GetRole(roleName string) (entity.Role, error)
}
