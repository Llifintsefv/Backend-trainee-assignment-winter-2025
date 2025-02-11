package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoinTransaction struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	SenderUserID    uuid.UUID `gorm:"type:uuid;"`
	SenderUser      User      `gorm:"foreignKey:SenderUserID"`
	ReceiverUserID  uuid.UUID `gorm:"type:uuid;not null"`
	ReceiverUser    User      `gorm:"foreignKey:ReceiverUserID"`
	Amount          int       `gorm:"not null"`
	TransactionType string    `gorm:"type:varchar(20);not null"`
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (ct *CoinTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	ct.ID = uuid.New()
	return
}
