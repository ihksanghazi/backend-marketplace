package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	Id        uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	UserId    uuid.UUID `gorm:"foreignkey;not null"`
	StoreId   uuid.UUID `gorm:"foreignkey;not null"`
	Total     string    `gorm:"type:numeric;not null"`
	TotalGram string    `gorm:"type:numeric"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Association
	CartDetail []CartDetail `gorm:"foreignKey:CartId"`
}

type CartDetail struct {
	Id        uuid.UUID `gorm:"primarykey;default:gen_random_uuid()"`
	CartId    uuid.UUID `gorm:"foreignkey;not null"`
	ProductId uuid.UUID `gorm:"foreignkey;not null"`
	Amount    string    `gorm:"type:numeric;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateCart struct {
	TotalPrice int
	TotalGram  int
}
