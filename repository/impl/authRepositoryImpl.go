package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"github.com/google/uuid"
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

func (a AuthRepository) CreateAccount(account entity.Account) (entity.Account, error) {
	var existingAccount entity.Account

	if err := config.DB.Where("email = ?", account.Email).First(&existingAccount).Error; err == nil {
		return account, errors.New("email already in use")
	}

	newUser := entity.User{
		Base: entity.Base{ID: uuid.NewString()},
	}

	user, err := a.CreateUser(newUser)
	if err != nil {
		return account, err
	}

	account.UserID = user.ID
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
