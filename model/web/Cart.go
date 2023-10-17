package web

import (
	"time"

	"github.com/google/uuid"
)

type GetCartResponse struct {
	Id        uuid.UUID                `json:"id"`
	StoreId   uuid.UUID                `json:"-"`
	Store     getCartStoreResponse     `json:"store" gorm:"foreignKey:StoreId"`
	Products  []getCartProductResponse `json:"products" gorm:"many2many:cart_details;foreignKey:Id;joinForeignKey:CartId;References:Id;joinReferences:ProductId"`
	Total     string                   `json:"total"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type getCartStoreResponse struct {
	Id          uuid.UUID `json:"-"`
	StoreName   string    `json:"store_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"image_url"`
}

func (s *getCartStoreResponse) TableName() string {
	return "stores"
}

type getCartProductResponse struct {
	Id          uuid.UUID                    `json:"-"`
	ProductName string                       `json:"product_name"`
	Description string                       `json:"description"`
	Category    string                       `json:"category"`
	Detail      getCartProductDetailResponse `json:"detail" gorm:"foreignKey:ProductId"`
	Price       string                       `json:"price"`
	ImageUrl    string                       `json:"image_url"`
}

func (p *getCartProductResponse) TableName() string {
	return "products"
}

type getCartProductDetailResponse struct {
	ProductId uuid.UUID `json:"-"`
	Amount    string    `json:"amount"`
}

func (d *getCartProductDetailResponse) TableName() string {
	return "cart_details"
}
