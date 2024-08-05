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
		IsActive  bool        `json:"isActive"`
		CreatedAt interface{} `json:"createdAt"`
		UpdatedAt interface{} `json:"updatedAt"`
		DeletedAt interface{} `json:"deletedAt"`
	}
)
