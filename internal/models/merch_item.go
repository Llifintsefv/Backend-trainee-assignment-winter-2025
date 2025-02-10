package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name      string    `gorm:"uniqueIndex;not null"`
	Price     int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *MerchItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
