package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"time"
)

type AdminRepository struct{}

func NewAdminRepository() *AccountRepository {
	return &AccountRepository{}
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
