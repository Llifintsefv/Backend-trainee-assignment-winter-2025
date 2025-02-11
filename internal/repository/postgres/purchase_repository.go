package postgres

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"

	"gorm.io/gorm"
)

type purchaseRepository struct {
	DB *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) repository.PurchaseRepository {
	return &purchaseRepository{DB: db}
}
