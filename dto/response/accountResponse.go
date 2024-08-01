package response

type (
	LoginResponse struct {
		Token        string `json:"token"`
		HashPassword string `json:"-"`
		UserId       string `json:"-"`
		Username     string `json:"-"`
		Email        string `json:"-"`
		Role         string `json:"-"`
	}

	ListAccountResponse struct {
		Username string `json:"name"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}

	AccountDetailResponse struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Role     string `json:"role"`
		IsActive bool   `json:"is_active"`
	}
)
