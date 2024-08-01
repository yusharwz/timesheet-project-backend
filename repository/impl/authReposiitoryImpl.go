package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"

	"gorm.io/gorm"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (AuthRepository) CreateUser(user entity.User) (entity.User, error) {
	err := config.DB.Create(&user)
	if err.Error != nil {
		return user, errors.New("failed to create user")
	}

	return user, nil
}

func (AuthRepository) CreateAccount(account entity.Account) (entity.Account, error) {
	var existingAccount entity.Account

	if err := config.DB.Where("email = ?", account.Email).First(&existingAccount).Error; err == nil {
		return account, errors.New("email already in use")
	}

	if err := config.DB.Where("username = ?", account.Username).First(&existingAccount).Error; err == nil {
		return account, errors.New("username already in use")
	}

	if err := config.DB.Create(&account).Error; err != nil {
		return account, errors.New("failed to create account")
	}

	return account, nil
}

func (AuthRepository) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	var account entity.Account
	var role entity.Role

	result := config.DB.Where("email = ?", req.Email).First(&account)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return resp, errors.New("invalid email or password")
		}
		return resp, result.Error
	}

	if err := config.DB.Where("id = ?", account.RoleID).First(&role).Error; err != nil {
		return resp, err
	}

	if !account.IsActive {
		return resp, errors.New("account is not active")
	}

	if !account.DeletedAt.Time.IsZero() {
		return resp, errors.New("account has been deleted")
	}

	resp.HashPassword = account.Password
	resp.Email = account.Email
	resp.UserId = account.UserID
	resp.Username = account.Username
	resp.Role = role.RoleName

	return resp, nil
}
