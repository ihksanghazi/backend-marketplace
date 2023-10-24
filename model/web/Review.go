package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
}

type GetReviewResponse struct {
	Id        uuid.UUID `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *GetReviewResponse) TableName() string {
	return "reviews"
}
