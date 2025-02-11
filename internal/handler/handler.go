package handler

import (
	"Backend-trainee-assignment-winter-2025/internal/models"
	"Backend-trainee-assignment-winter-2025/internal/repository"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NewUser(c *gin.Context)
}

type handler struct {
	coinRepo     repository.CoinTransactionRepository
	userRepo     repository.UserRepository
	purchaseRepo repository.PurchaseRepository
	logger       *slog.Logger
}

func NewController(coinRepo repository.CoinTransactionRepository, userRepo repository.UserRepository, purchaseRepo repository.PurchaseRepository, logger *slog.Logger) Handler {
	return &handler{coinRepo: coinRepo, userRepo: userRepo, purchaseRepo: purchaseRepo, logger: logger}
}

func (h *handler) NewUser(c *gin.Context) {
	ctx := c.Request.Context()
	var AuthRequest models.AuthRequest

	if err := c.ShouldBindJSON(&AuthRequest); err != nil {
		h.logger.ErrorContext(ctx, "failed to bind json", "error", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	err := h.userRepo.CreateUser(ctx, AuthRequest.Username, AuthRequest.Password)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{})
}
