package response

type (
	ListAccountResponse struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Role   string `json:"role"`
		Status string `json:"status"`
	}

	AccountDetailResponse struct {
		Name      string      `json:"name"`
		Email     string      `json:"email"`
		Phone     string      `json:"phone"`
		Role      string      `json:"role"`
		IsActive  bool        `json:"isActive"`
		CreatedAt interface{} `json:"createdAt"`
		UpdatedAt interface{} `json:"updatedAt"`
		DeletedAt interface{} `json:"deletedAt"`
	}

	RoleResponse struct {
		ID       string `json:"id"`
		RoleName string `json:"roleName"`
	}
)
