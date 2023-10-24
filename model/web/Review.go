package web

type CreateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
}
