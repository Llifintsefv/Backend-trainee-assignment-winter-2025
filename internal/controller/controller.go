package controller

import "Backend-trainee-assignment-winter-2025/internal/repository"

type Controller interface {
}

type controller struct {
	coinRepo     repository.CoinTransactionRepository
	userRepo     repository.UserRepository
	purchaseRepo repository.PurchaseRepository
}

func NewController(coinRepo repository.CoinTransactionRepository, userRepo repository.UserRepository, purchaseRepo repository.PurchaseRepository) Controller {
	return &controller{coinRepo: coinRepo, userRepo: userRepo, purchaseRepo: purchaseRepo}
}
