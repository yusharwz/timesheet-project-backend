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
)
