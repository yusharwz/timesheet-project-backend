package impl

import (
	"errors"
	"final-project-enigma/dto/response"
	"final-project-enigma/helper"
	"final-project-enigma/repository"
	"final-project-enigma/repository/impl"
	"final-project-enigma/service"
	"strconv"
)

type AdminService struct{}

var adminRepository repository.AdminRepository = impl.NewAdminRepository()
var authService service.AuthService = NewAuthService()

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

	var resp []response.ListAccountResponse
	for _, user := range users {
		var status string
		if user.Account.IsActive {
			status = "active"
		} else {
			status = "inactive"
		}
		result, err := authService.GetRoleById(user.Account.RoleID)
		if err != nil {
			return nil, "0", "0", err
		}
		resp = append(resp, response.ListAccountResponse{
			ID:     user.ID,
			Email:  user.Account.Email,
			Name:   user.Name,
			Role:   result.RoleName,
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

func (AdminService) GetAllRole() (*[]response.RoleResponse, error) {
	roles, err := adminRepository.GetAllRole()
	if err != nil {
		return nil, err
	}
	var responseList = make([]response.RoleResponse, 0)
	for _, role := range *roles {
		responseList = append(responseList, response.RoleResponse{
			ID:       role.ID,
			RoleName: role.RoleName,
		})
	}
	return &responseList, nil
}
