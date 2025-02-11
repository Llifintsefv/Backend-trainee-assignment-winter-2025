package postgres

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"

	"gorm.io/gorm"
)

type MerchRepository struct {
	DB *gorm.DB
}

func NewMerchRepository(db *gorm.DB) repository.MerchRepository {
	return &MerchRepository{DB: db}
}
