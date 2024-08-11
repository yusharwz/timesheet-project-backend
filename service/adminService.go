package service

import "timesheet-app/dto/response"

type AdminService interface {
	RetrieveAccountList(paging, rowsPerPage, name string) ([]response.ListAccountResponse, string, string, error)
	DetailAccount(userID string) (response.AccountDetailResponse, error)
	SoftDeleteAccount(userID string) error
	GetAllRole() (*[]response.RoleResponse, error)
}
