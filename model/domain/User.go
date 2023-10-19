package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	Username     string    `gorm:"index;not null"`
	Email        string    `gorm:"unique;index;not null"`
	Password     string    `gorm:"not null"`
	RefreshToken string
	Role         string `gorm:"default:user;index"`
	PhoneNumber  string `gorm:"unique;default:NULL"`
	Address      string
	CityId       string `gorm:"foreignKey"`
	ImageUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	// Association
	Store Store  `gorm:"foreignKey:UserId"`
	Cart  []Cart `gorm:"foreignKey:UserId"`
}
