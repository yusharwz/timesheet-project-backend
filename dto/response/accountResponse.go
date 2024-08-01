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

	AccountUserResponse struct {
		AccountID string `json:"account_id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		IsActive  bool   `json:"is_active"`
		UserID    string `json:"user_id"`
		Name      string `json:"name"`
		Phone     string `json:"phone"`
	}
)
