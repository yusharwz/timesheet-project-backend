package impl

import (
	"errors"
	"final-project-enigma/dto/response"
	"final-project-enigma/helper"
	"final-project-enigma/repository"
	"final-project-enigma/repository/impl"
	"strconv"
)

type AdminService struct{}

var adminRepository repository.AdminRepository = impl.NewAdminRepository()

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (AdminService) RetrieveAccountList(paging, rowsPerPage string) ([]response.ListAccountResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query for paging")
	}
	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query for rows per page")
	}
	users, totalRows, err := adminRepository.RetrieveAccountList(pagingInt, rowsPerPageInt)
	if err != nil {
		return nil, "0", "0", err
	}

	//userMap := make(map[string]string)
	//for _, user := range users {
	//	userMap[user.ID] = user.Name
	//}

	var resp []response.ListAccountResponse
	for _, user := range users {
		var status string
		if user.Account.IsActive {
			status = "Active"
		} else {
			status = "Inactive"
		}

		resp = append(resp, response.ListAccountResponse{
			Email:  user.Account.Email,
			Status: status,
		})
	}

	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)
	return resp, totalRows, strconv.Itoa(totalPage), nil
}

func (AdminService) DetailAccount(userID string) (response.AccountDetailResponse, error) {
	var resp response.AccountDetailResponse

	account, user, role, err := adminRepository.DetailAccount(userID)
	if err != nil {
		return resp, err
	}

	resp.Name = user.Name
	resp.Email = account.Email
	resp.Phone = user.PhoneNumber
	resp.Role = role.RoleName
	resp.IsActive = account.IsActive
	resp.CreatedAt = account.CreatedAt
	resp.UpdatedAt = account.UpdatedAt
	resp.DeletedAt = account.DeletedAt

	return resp, nil
}

func (AdminService) SoftDeleteAccount(userID string) error {
	return adminRepository.SoftDeleteAccount(userID)
}
