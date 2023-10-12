package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	Username string `gorm:"not null"`
	Email string `gorm:"unique;index;not null"`
	Password string `gorm:"index;not null"`
	RefreshToken string `gorm:"index"`
	Role string `gorm:"default:user;index"`
	PhoneNumber string `gorm:"unique"`
	Address string
	ImageUrl string
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}