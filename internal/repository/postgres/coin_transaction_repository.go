package postgres

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"

	"gorm.io/gorm"
)

type coinTransactionRepository struct {
	DB *gorm.DB
}

func NewCoinTransactionRepository(db *gorm.DB) repository.CoinTransactionRepository {
	return &coinTransactionRepository{DB: db}
}
