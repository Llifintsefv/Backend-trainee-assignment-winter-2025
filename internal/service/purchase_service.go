package service

import "Backend-trainee-assignment-winter-2025/internal/repository"

type PurchaseService interface{}

type purchaseService struct {
	repo repository.PurchaseRepository
}

func NewPurchaseService(repo repository.PurchaseRepository) PurchaseService {
	return &purchaseService{repo: repo}
}
