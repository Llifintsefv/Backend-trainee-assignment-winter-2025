package repository

import (
	"Backend-trainee-assignment-winter-2025/internal/models"
	"context"
)

type CoinTransactionRepository interface {
}

type MerchRepository interface {
}

type PurchaseRepository interface {
}

type UserRepository interface {
	CreateUser(ctx context.Context, username, password string) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
}
