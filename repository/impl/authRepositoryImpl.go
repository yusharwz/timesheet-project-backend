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

func (AuthRepository) Register(user entity.User, account entity.Account) (entity.User, entity.Account, error) {

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingAccount entity.Account
	if err := tx.Where("email = ?", account.Email).First(&existingAccount).Error; err == nil {
		tx.Rollback()
		return user, account, errors.New("email already in use")
	}

	user.Account = account
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return user, account, errors.New("failed to create account")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return user, account, errors.New("transaction commit failed")
	}

	return user, account, nil
}

func (AuthRepository) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
	var account entity.Account
	var user entity.User
	var role entity.Role

	resultAccount := config.DB.Where("email = ?", req.Email).First(&account)
	if resultAccount.Error != nil {
		if errors.Is(resultAccount.Error, gorm.ErrRecordNotFound) {
			return resp, errors.New("invalid email or password")
		}
		return resp, resultAccount.Error
	}

	resultUser := config.DB.Where("id = ?", account.UserID).First(&user)
	if resultUser.Error != nil {
		if errors.Is(resultUser.Error, gorm.ErrRecordNotFound) {
			return resp, errors.New("invalid email or password")
		}
		return resp, resultUser.Error
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
	resp.Name = user.Name
	resp.UserId = account.UserID
	resp.Role = role.RoleName

	return resp, nil
}

func (AuthRepository) GetRole(roleName string) (entity.Role, error) {
	var role entity.Role
	result := config.DB.Where("role_name = ?", roleName).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Role{}, errors.New("invalid role")
		}
		return entity.Role{}, result.Error
	}
	return role, nil
}
