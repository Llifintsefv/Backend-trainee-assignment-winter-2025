package service

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"
	"context"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser(ctx context.Context, username, password string) (string, error)
}

type userService struct {
	repo      repository.UserRepository
	secretKey []byte
	logger    *slog.Logger
}

func NewUserService(repo repository.UserRepository, secretKey []byte, logger *slog.Logger) UserService {
	return &userService{repo: repo, secretKey: secretKey, logger: logger}
}

func (u *userService) generateJWToken(username string) (string, error) {
	type Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	claims := Claims{
		Username:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (u *userService) CreateUser(ctx context.Context, username, password string) (string, error) {

	err := u.repo.CreateUser(ctx, username, password)

	if err != nil {
		u.logger.ErrorContext(ctx, "failed to create user", "error", err)
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := u.generateJWToken(username)

	if err != nil {
		u.logger.ErrorContext(ctx, "failed to generate token", "error", err)
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
