package web

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	CityId      string `json:"city_id" binding:"required,number"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email" binding:"omitempty,email"`
	Password    string `json:"password"`
	CityId      string `json:"city_id" binding:"omitempty,number"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}

type FindUserResponse struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	Id          uuid.UUID            `json:"id"`
	Username    string               `json:"username"`
	Email       string               `json:"email"`
	PhoneNumber string               `json:"phone_number"`
	Address     string               `json:"address"`
	ImageUrl    string               `json:"image_url"`
	Store       getUserStoreResponse `json:"store" gorm:"foreignKey:UserId"`
	Carts       []GetCartResponse    `json:"carts" gorm:"foreignKey:UserId"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type getUserStoreResponse struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"-"`
	StoreName   string    `json:"store_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"image_url"`
}

func (s *getUserStoreResponse) TableName() string {
	return "stores"
}
