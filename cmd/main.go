package main

import (
	"Backend-trainee-assignment-winter-2025/internal/config"
	"Backend-trainee-assignment-winter-2025/internal/handler"
	"Backend-trainee-assignment-winter-2025/internal/repository/postgres"
	"Backend-trainee-assignment-winter-2025/internal/router"
	"Backend-trainee-assignment-winter-2025/internal/service"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	db, err := postgres.NewDB(cfg.DBConnStr)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	userRepo := postgres.NewUserRepository(db)
	coinRepo := postgres.NewCoinTransactionRepository(db)
	purchaseRepo := postgres.NewPurchaseRepository(db)
	merchRepo := postgres.NewMerchRepository(db)

	userService := service.NewUserService(userRepo)
	coinService := service.NewCoinService(coinRepo)
	purchaseService := service.NewPurchaseService(purchaseRepo, coinRepo, merchRepo)

	handler := handler.NewController(userService, coinService, purchaseService,logger)

	app := router.SetupRouter(handler,[]byte(cfg.SecretKey))
	app.Run()

}
