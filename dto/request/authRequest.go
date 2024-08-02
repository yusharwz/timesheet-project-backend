package request

type (
	RegisterAccountRequest struct {
		Email    string `json:"email" binding:"required,email"`
		RoleName string `json:"role"`
	}
)
