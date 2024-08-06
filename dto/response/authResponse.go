package response

type (
	RegisterAccountResponse struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required,name"`
		RoleName string `json:"role"`
	}
	GetRoleResponse struct {
		RoleName string `json:"role"`
	}
)
