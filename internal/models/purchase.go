package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	User         User      `gorm:"foreignKey:UserID"`
	MerchItemID  uuid.UUID `gorm:"type:uuid;not null"`
	MerchItem    MerchItem `gorm:"foreignKey:MerchItemID"`
	PurchaseDate time.Time 
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
