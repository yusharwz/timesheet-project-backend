package response

type (
	RegisterAccountResponse struct {
		Id       string `json:"-"`
		Email    string `json:"email"`
		Password string `json:"-"`
		IsActive bool   `json:"-"`
		RoleID   string `json:"-"`
		UserID   string `json:"-"`
	}
)
