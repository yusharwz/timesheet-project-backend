package impl

import (
	"final-project-enigma/config"
	"final-project-enigma/entity"
	"final-project-enigma/helper"
	"gorm.io/gorm"
	"time"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (AdminRepository) RetrieveAccountList(spec []func(db *gorm.DB) *gorm.DB) ([]entity.User, string, error) {
	var users []entity.User

	db := config.DB.Model(entity.User{}).Scopes(spec[1:]...).Preload("Account")
	totalRows := helper.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&users).Error
	return users, totalRows, err
}

func (AdminRepository) DetailAccount(userID string) (entity.Account, entity.User, entity.Role, error) {
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

func (AdminRepository) SoftDeleteAccount(userID string) error {
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

func (AdminRepository) GetAllRole() (*[]entity.Role, error) {
	var roles []entity.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return &roles, nil
}
