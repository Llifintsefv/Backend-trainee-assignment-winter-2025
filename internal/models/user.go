package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Coins        int       `gorm:"not null;default:1000"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
