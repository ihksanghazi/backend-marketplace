package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	Id           uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	UserId			 uuid.UUID `gorm:"foreignkey;unique;not null"`
	StoreName 	 string 	 `gorm:"not null;index"`
	Description  string		 `gorm:"not null"`
	Category		 string		 `gorm:"not null;index"`
	ImageUrl		 string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}