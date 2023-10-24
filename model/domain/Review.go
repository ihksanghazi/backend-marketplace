package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	Id        uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	UserId    uuid.UUID `gorm:"foreignkey;not null"`
	ProductId uuid.UUID `gorm:"foreignkey;not null"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
