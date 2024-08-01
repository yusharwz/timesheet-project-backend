package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/entity"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (AuthRepository) CreateUser(user entity.User) (entity.User, error) {
	if err := config.DB.Create(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

func (AuthRepository) CreateAccount(account entity.Account) (entity.Account, error) {
	if err := config.DB.Create(&account); err != nil {
		return account, err.Error
	}

	return account, nil
}
