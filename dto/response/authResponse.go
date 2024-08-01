package response

type (
	RegisterAccountResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email" binding:"required,email"`
		Username string `gorm:"unique;not null" json:"username"`
		Password string `json:"password" binding:"required,password"`
		IsActive bool   `json:"isActive"`
		RoleID   string `json:"role"`
		UserID   string `json:"user"`
	}
)
