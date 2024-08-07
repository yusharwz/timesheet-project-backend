package repository

import "final-project-enigma/entity"

type AdminRepository interface {
	RetrieveAccountList(paging, rowsPerPage int) ([]entity.User, string, error)
	DetailAccount(userID string) (entity.Account, entity.User, entity.Role, error)
	SoftDeleteAccount(userID string) error
	GetAllRole() (*[]entity.Role, error)
}
