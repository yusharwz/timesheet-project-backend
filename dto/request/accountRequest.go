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
		UserID   string `json:"-"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
	}
)
