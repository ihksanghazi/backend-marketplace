package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	Id          uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	UserId      uuid.UUID `gorm:"foreignkey;unique;not null"`
	StoreName   string    `gorm:"not null;index"`
	Description string    `gorm:"not null"`
	Category    string    `gorm:"not null;index"`
	Address     string
	CityId      string `gorm:"foreignKey"`
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Association
	Products    []Product     `gorm:"foreignKey:StoreId"`
	Cart        []Cart        `gorm:"foreignKey:StoreId"`
	Transaction []Transaction `gorm:"foreignKey:StoreId"`
}
