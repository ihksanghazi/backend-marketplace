package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Stock       string `json:"stock" binding:"required,number"`
	Price       string `json:"price" binding:"required,number"`
	ImageUrl    string `json:"image_url"`
}

type UpdateProductRequest struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Stock       string `json:"stock" binding:"omitempty,number"`
	Price       string `json:"price" binding:"omitempty,number"`
	ImageUrl    string `json:"image_url"`
}

type FindProductResponse struct {
	Id      uuid.UUID `json:"id"`
	StoreId uuid.UUID `json:"-"`

	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Stock       string    `json:"stock"`
	Price       string    `json:"price"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
