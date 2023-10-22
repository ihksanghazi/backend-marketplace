package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id                uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	UserId            uuid.UUID `gorm:"foreignkey;not null"`
	StoreId           uuid.UUID `gorm:"foreignkey;not null"`
	TransactionStatus string    `gorm:"not null"`
	TotalProductPrice string    `gorm:"type:numeric;not null"`
	TotalPrice        string    `gorm:"type:numeric;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	// Association
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId"`
	Expedition        Expedition          `gorm:"foreignKey:TransactionId"`
}

type TransactionDetail struct {
	Id            uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	TransactionId uuid.UUID `gorm:"foreignkey;not null"`
	ProductId     uuid.UUID `gorm:"foreignkey;not null"`
	Amount        string    `gorm:"type:numeric;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Expedition struct {
	Id              uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	TransactionId   uuid.UUID `gorm:"foreignkey;not null"`
	OriginCity      string    `gorm:"not null"`
	DestinationCity string    `gorm:"not null"`
	Courier         string    `gorm:"not null"`
	WeightOnGram    string    `gorm:"type:numeric;not null"`
	Service         string    `gorm:"not null"`
	Description     string    `gorm:"not null"`
	Price           int       `gorm:"not null;type:numeric"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
