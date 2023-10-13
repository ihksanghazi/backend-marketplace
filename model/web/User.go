package web

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Address  string `json:"address"`
	ImageUrl string `json:"image_url"`
}

type RegisterResponse struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address      string		 `json:"address"`
	ImageUrl     string		 `json:"image_url"`
	CreatedAt 	 time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginRequest struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct{
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address  string `json:"address"`
	ImageUrl string `json:"image_url"`
}