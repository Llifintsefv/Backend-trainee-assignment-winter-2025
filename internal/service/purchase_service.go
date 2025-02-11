package service

import "Backend-trainee-assignment-winter-2025/internal/repository"

type PurchaseService interface{}

type purchaseService struct {
	purchaseRepo repository.PurchaseRepository
	coinRepo     repository.CoinTransactionRepository
	merchRepo    repository.MerchRepository
}

func NewPurchaseService(purchaseRepo repository.PurchaseRepository, coinRepo repository.CoinTransactionRepository, merchRepo repository.MerchRepository) PurchaseService {
	return &purchaseService{purchaseRepo: purchaseRepo, coinRepo: coinRepo, merchRepo: merchRepo}
}
