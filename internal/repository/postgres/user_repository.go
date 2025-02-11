package postgres

import (
	"Backend-trainee-assignment-winter-2025/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{DB: db}
}
