package repository

import "final-project-enigma/entity"

type AuthRepository interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	CreateUser(user entity.User) (entity.User, error)
}
