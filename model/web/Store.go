package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateStoreRequest struct {
	StoreName   string `json:"store_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"image_url"`
}

type UpdateStoreRequest struct {
	StoreName   string `json:"store_name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
}

type FindStoreResponse struct {
	Id          uuid.UUID `json:"id"`
	StoreName   string    `json:"store_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
