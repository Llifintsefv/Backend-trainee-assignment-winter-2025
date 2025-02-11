package service

import "Backend-trainee-assignment-winter-2025/internal/repository"

type CoinService interface {
}

type coinService struct {
	repo repository.CoinTransactionRepository
}

func NewCoinService(repo repository.CoinTransactionRepository) CoinService {
	return &coinService{repo: repo}
}
