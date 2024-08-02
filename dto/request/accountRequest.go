package request

type (
	ActivateAccountRequest struct {
		Email    string
		Username string
		Password string
	}

	LoginAccountRequest struct {
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required"`
	}

	EditAccountRequest struct {
		UserID string `json:"-"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Phone  string `json:"phone"`
	}

	ChangePasswordRequest struct {
		UserID      string `json:"user_id"`
		NewPassword string `json:"new_password" binding:"password"`
	}
)
