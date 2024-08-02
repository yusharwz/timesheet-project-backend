package request

type WorkRequest struct {
	Description string `json:"description" binding:"required"`
	Fee         int    `json:"fee" binding:"min=0,required"`
}
