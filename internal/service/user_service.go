package service

import "Backend-trainee-assignment-winter-2025/internal/repository"

type UserService interface{}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
