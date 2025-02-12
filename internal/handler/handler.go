package handler

import (
	"Backend-trainee-assignment-winter-2025/internal/models"
	"Backend-trainee-assignment-winter-2025/internal/pkg/validator"
	"Backend-trainee-assignment-winter-2025/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NewUser(c *gin.Context)
	Test(c *gin.Context)
}

type handler struct {
	coinService     service.CoinService
	userService     service.UserService
	purchaseService service.PurchaseService
	logger          *slog.Logger
}

func NewHandler(coinService service.CoinService, userService service.UserService, purchaseService service.PurchaseService, logger *slog.Logger) Handler {
	return &handler{coinService: coinService, userService: userService, purchaseService: purchaseService, logger: logger}
}

func (h *handler) NewUser(c *gin.Context) {
	ctx := c.Request.Context()
	var AuthRequest models.AuthRequest

	if err := c.ShouldBindJSON(&AuthRequest); err != nil {
		h.logger.ErrorContext(ctx, "failed to bind json", "error", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	if err := validator.ValidateStruct(AuthRequest); err != nil {
		h.logger.ErrorContext(ctx, "failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	token, err := h.userService.CreateUser(ctx, AuthRequest.Username, AuthRequest.Password)
	if err != nil {
		h.logger.ErrorContext(ctx, "failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{Token: token})
}

func (h *handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
