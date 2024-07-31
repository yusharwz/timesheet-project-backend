package response

type WorkResponse struct {
	Id          string `json:"id"`
	Description string `json:"description,omitempty"`
	Fee         int    `json:"fee"`
}
