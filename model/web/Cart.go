package web

import (
	"time"

	"github.com/google/uuid"
)

type GetCartResponse struct {
	Id        uuid.UUID            `json:"cart_id"`
	UserId    uuid.UUID            `json:"-"`
	StoreId   uuid.UUID            `json:"-"`
	Store     getCartStoreResponse `json:"store" gorm:"foreignKey:StoreId"`
	Items     []cartItemsResponse  `json:"items" gorm:"foreignKey:CartId"`
	Total     string               `json:"total_price"`
	TotalGram string               `json:"total_gram"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
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

type cartItemsResponse struct {
	Id        uuid.UUID                `json:"id"`
	CartId    uuid.UUID                `json:"-"`
	ProductId uuid.UUID                `json:"-"`
	Amount    string                   `json:"amount"`
	Product   cartItemsProductResponse `json:"product" gorm:"foreignKey:ProductId"`
}

func (i *cartItemsResponse) TableName() string {
	return "cart_details"
}

type cartItemsProductResponse struct {
	Id          uuid.UUID `json:"-"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	WeightGram  int       `json:"weight_on_gram"`
	Price       string    `json:"price"`
	ImageUrl    string    `json:"image_url"`
}

func (p *cartItemsProductResponse) TableName() string {
	return "products"
}
