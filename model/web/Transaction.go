package web

import (
	"time"

	"github.com/google/uuid"
)

type CheckoutRequest struct {
	OriginCity       string `json:"origin_city" binding:"required"`
	DestionationCity string `json:"destination_city" binding:"required"`
	Courier          string `json:"courier" binding:"required"`
	WeightOnGram     string `json:"weight_on_gram" binding:"required,number"`
	Service          string `json:"service" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Price            int    `json:"price" binding:"required,number"`
}

type GetTransactionResponse struct {
	Id                uuid.UUID                        `json:"id"`
	UserId            uuid.UUID                        `json:"-"`
	User              getTransactionUserResponse       `json:"user" gorm:"foreignKey:UserId"`
	StoreId           uuid.UUID                        `json:"-"`
	Store             getTransactionStoreResponse      `json:"store" gorm:"foreignKey:StoreId"`
	Items             []getTransactionItemResponse     `json:"item" gorm:"foreignKey:TransactionId"`
	Expedition        getTransactionExpeditionResponse `json:"expedition" gorm:"foreignKey:TransactionId"`
	TransactionStatus string                           `json:"transaction_status"`
	TotalProductPrice string                           `json:"total_product_price"`
	TotalPrice        string                           `json:"total_price"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}

func (t *GetTransactionResponse) TableName() string {
	return "transactions"
}

type getTransactionUserResponse struct {
	Id          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Address     string        `json:"address"`
	ImageUrl    string        `json:"image_url"`
	CityId      string        `json:"-"`
	Region      getUserRegion `json:"region" gorm:"foreignKey:CityId"`
}

func (t *getTransactionUserResponse) TableName() string {
	return "users"
}

type getTransactionStoreResponse struct {
	Id          uuid.UUID      `json:"id"`
	CityId      string         `json:"-"`
	Region      getStoreRegion `json:"region" gorm:"foreignKey:CityId"`
	StoreName   string         `json:"store_name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	ImageUrl    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (s *getTransactionStoreResponse) TableName() string {
	return "stores"
}

type getTransactionItemResponse struct {
	Id            uuid.UUID                `json:"id"`
	TransactionId uuid.UUID                `json:"-"`
	ProductId     uuid.UUID                `json:"-"`
	Product       cartItemsProductResponse `json:"product" gorm:"foreignKey:ProductId"`
	Amount        string                   `json:"amount"`
}

func (i *getTransactionItemResponse) TableName() string {
	return "transaction_details"
}

type getTransactionExpeditionResponse struct {
	Id              uuid.UUID `json:"id"`
	TransactionId   uuid.UUID `json:"-"`
	OriginCity      string    `json:"origin_city"`
	DestinationCity string    `json:"destination_city"`
	Courier         string    `json:"courier"`
	WeightOnGram    string    `json:"weight_on_gram"`
	Service         string    `json:"service"`
	Description     string    `json:"description"`
	Price           int       `json:"price"`
}

func (e *getTransactionExpeditionResponse) TableName() string {
	return "expeditions"
}
