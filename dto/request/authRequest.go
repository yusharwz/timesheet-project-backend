package request

type (
	RegisterAccountRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `gorm:"unique;not null" json:"username"`
		Password string `json:"password"`
		IsActive bool   `json:"isActive"`
		RoleID   string `json:"role"`
		UserID   string `json:"user"`
	}
)
