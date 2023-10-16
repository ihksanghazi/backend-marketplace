package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	StoreId     uuid.UUID `gorm:"foreignkey;not null"`
	ProductName string    `gorm:"not null;index"`
	Description string    `gorm:"not null"`
	Category    string    `gorm:"not null;index"`
	Stock       string    `gorm:"not null;type:numeric"`
	Price       string    `gorm:"not null;type:numeric"`
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Association
	CartDetail []CartDetail `gorm:"foreignKey:ProductId"`
}
