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
	Id          uuid.UUID `json:"id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Stock       string    `json:"stock"`
	Price       string    `json:"price"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetProductResponse struct {
	Id          uuid.UUID               `json:"id"`
	StoreId     uuid.UUID               `json:"-"`
	Store       GetProductStoreResponse `json:"store" gorm:"foreignKey:StoreId"`
	ProductName string                  `json:"product_name"`
	Description string                  `json:"description"`
	Category    string                  `json:"category"`
	Stock       string                  `json:"stock"`
	Price       string                  `json:"price"`
	ImageUrl    string                  `json:"image_url"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
}

type GetProductStoreResponse struct {
	Id        uuid.UUID `json:"id"`
	StoreName string    `json:"store_name"`
	Category  string    `json:"category"`
	ImageUrl  string    `json:"image_url"`
}

func (s *GetProductStoreResponse) TableName() string {
	return "stores"
}
