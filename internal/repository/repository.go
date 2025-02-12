package repository

import "context"

type CoinTransactionRepository interface {
}

type MerchRepository interface {
}

type PurchaseRepository interface {
}

type UserRepository interface {
	CreateUser(ctx context.Context, username, password string) (error)
}
