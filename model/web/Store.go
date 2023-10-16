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

type GetStoreResponse struct {
	Id          uuid.UUID                 `json:"id"`
	Products    []GetStoreProductResponse `json:"products" gorm:"foreignKey:StoreId"`
	StoreName   string                    `json:"store_name"`
	Description string                    `json:"description"`
	Category    string                    `json:"category"`
	ImageUrl    string                    `json:"image_url"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

type GetStoreProductResponse struct {
	Id          uuid.UUID `json:"id"`
	StoreId     uuid.UUID `json:"-"`
	ProductName string    `json:"product_name"`
	Category    string    `json:"category"`
	Stock       string    `json:"stock"`
	Price       string    `json:"price"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *GetStoreProductResponse) TableName() string {
	return "products"
}
