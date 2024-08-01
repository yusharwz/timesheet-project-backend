package impl

import (
	"final-project-enigma/dto/response"
	"final-project-enigma/repository/impl"
)

type AdminService struct{}

var adminRepository = impl.NewAdminRepository()

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (AdminService) RetrieveAccountList() ([]response.ListAccountResponse, error) {

	accounts, users, err := adminRepository.RetrieveAccountList()
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]string)
	for _, user := range users {
		userMap[user.ID] = user.Name
	}

	var resp []response.ListAccountResponse
	for _, account := range accounts {
		var status string
		if account.IsActive {
			status = "Active"
		} else {
			status = "Inactive"
		}

		resp = append(resp, response.ListAccountResponse{
			Username: userMap[account.UserID],
			Email:    account.Email,
			Status:   status,
		})
	}

	return resp, nil
}

func (AdminService) DetailAccount(userID string) (response.AccountDetailResponse, error) {
	var resp response.AccountDetailResponse

	account, user, role, err := adminRepository.DetailAccount(userID)
	if err != nil {
		return resp, err
	}

	resp.Username = account.Username
	resp.Name = user.Name
	resp.Email = account.Email
	resp.Phone = user.PhoneNumber
	resp.Role = role.RoleName
	resp.IsActive = account.IsActive

	return resp, nil
}

func (AdminService) SoftDeleteAccount(userID string) error {
	return adminRepository.SoftDeleteAccount(userID)
}
