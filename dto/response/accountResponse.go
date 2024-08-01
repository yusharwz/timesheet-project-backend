package response

type (
	LoginResponse struct {
		Token        string `json:"token"`
		HashPassword string `json:"-"`
		UserId       string `json:"-"`
		Username     string `json:"-"`
		Email        string `json:"-"`
		Role         string `json:"-"`
	}
)
