package request

type (
	RegisterAccountRequest struct {
		Email  string `json:"email" binding:"required,email"`
		Name   string `json:"name" binding:"required"`
		RoleId string `json:"roleId" binding:"required"`
	}
)
