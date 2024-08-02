package response

type (
	RegisterAccountResponse struct {
		Email string `json:"email" binding:"required,email"`
	}
)
