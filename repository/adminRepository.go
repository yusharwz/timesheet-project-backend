package repository

import (
	"timesheet-app/entity"

	"gorm.io/gorm"
)

type AdminRepository interface {
	RetrieveAccountList(spec []func(db *gorm.DB) *gorm.DB) ([]entity.User, string, error)
	DetailAccount(userID string) (entity.Account, entity.User, entity.Role, error)
	SoftDeleteAccount(userID string) error
	GetAllRole() (*[]entity.Role, error)
}
