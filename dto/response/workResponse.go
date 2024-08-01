package response

type WorkResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Fee         int    `json:"fee"`
}
