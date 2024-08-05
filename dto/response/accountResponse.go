package response

type (
	LoginResponse struct {
		Token        string `json:"token"`
		HashPassword string `json:"-"`
		Name         string `json:"-"`
		UserId       string `json:"-"`
		Email        string `json:"-"`
		Role         string `json:"-"`
	}

	AccountUserResponse struct {
		AccountID    string `json:"accountId"`
		Email        string `json:"email"`
		IsActive     bool   `json:"isActive"`
		UserID       string `json:"userId"`
		Name         string `json:"name"`
		Phone        string `json:"phone"`
		SignatureURL string `json:"signatureUrl"`
	}

	UploadImageResponse struct {
		ImageURL string `json:"imageUrl"`
	}
)
