package response

type (
	ListAccountResponse struct {
		Email  string `json:"email"`
		Status string `json:"status"`
	}

	AccountDetailResponse struct {
		Name      string      `json:"name"`
		Email     string      `json:"email"`
		Phone     string      `json:"phone"`
		Role      string      `json:"role"`
		IsActive  bool        `json:"is_active"`
		CreatedAt interface{} `json:"created_at"`
		UpdatedAt interface{} `json:"updated_at"`
		DeletedAt interface{} `json:"deleted_at"`
	}
)
