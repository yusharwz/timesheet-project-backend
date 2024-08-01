package impl

import (
	"errors"
	"final-project-enigma/config"
	"final-project-enigma/dto/request"
	"final-project-enigma/dto/response"
	"final-project-enigma/entity"
	"time"

	"gorm.io/gorm"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (AccountRepository) AccountActivation(email string) error {

	result := config.DB.Model(&entity.Account{}).Where("email = ?", email).Update("is_active", true)
	if result.Error != nil {
		return errors.New("disini")
	}

	return nil
}

func (AccountRepository) Login(req request.LoginAccountRequest) (resp response.LoginResponse, err error) {
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
	resp.UserId = account.ID
	resp.Username = account.Username
	resp.Role = role.RoleName

	return resp, nil
}

func (AccountRepository) RetrieveAccountList() ([]entity.Account, []entity.User, error) {
	var accounts []entity.Account
	var users []entity.User

	if err := config.DB.Find(&accounts).Error; err != nil {
		return nil, nil, err
	}

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, nil, err
	}

	return accounts, users, nil
}

func (AccountRepository) DetailAccount(userID string) (entity.Account, entity.User, entity.Role, error) {
	var account entity.Account
	var user entity.User
	var role entity.Role

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return account, user, role, err
	}

	if err := config.DB.Where("user_id = ?", user.ID).First(&account).Error; err != nil {
		return account, user, role, err
	}

	if err := config.DB.Where("id = ?", account.RoleID).First(&role).Error; err != nil {
		return account, user, role, err
	}

	return account, user, role, nil
}

func (AccountRepository) SoftDeleteAccount(userID string) error {
	var account entity.Account
	var user entity.User

	if err := config.DB.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	now := time.Now()
	if err := config.DB.Model(&account).Update("deleted_at", &now).Error; err != nil {
		return err
	}
	if err := config.DB.Model(&user).Update("deleted_at", &now).Error; err != nil {
		return err
	}

	return nil
}
