package postgres

import (
	"Backend-trainee-assignment-winter-2025/internal/models"
	"Backend-trainee-assignment-winter-2025/internal/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) CreateUser(ctx context.Context, username, password string) error {

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

		return err
	}
	user := models.User{Username: username, PasswordHash: string(PasswordHash)}

	result := u.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	result := u.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
