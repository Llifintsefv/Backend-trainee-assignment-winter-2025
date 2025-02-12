package service

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AuthUser(ctx context.Context, username, password string) (string, error)
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

func (u *userService) AuthUser(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := u.repo.CreateUser(ctx, username, password)
			if err != nil {
				u.logger.ErrorContext(ctx, "failed to create user", "error", err)
				return "", fmt.Errorf("failed to create user: %w", err)
			}
			u.logger.InfoContext(ctx, "user created", "username", username)
		} else {
			u.logger.ErrorContext(ctx, "failed to get user by username", "error", err)
			return "", fmt.Errorf("failed to get user: %w", err)
		}
	}

	valid := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if valid != nil {
		u.logger.ErrorContext(ctx, "invalid password", "error", valid)
		return "", fmt.Errorf("invalid password: %w", valid)
	}

	token, err := u.generateJWToken(username)
	if err != nil {
		u.logger.ErrorContext(ctx, "failed to generate token", "error", err)
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
