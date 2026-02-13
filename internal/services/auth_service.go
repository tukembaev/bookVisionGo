package services

import (
	"context"
	"fmt"

	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/repositories/interfaces"
	"github.com/tukembaev/bookVisionGo/internal/utils"
)

// AuthService - сервис аутентификации
type AuthService struct {
	userRepo interfaces.UserRepository
	jwtUtils *utils.JWTUtils
}

// NewAuthService - создание нового AuthService
func NewAuthService(userRepo interfaces.UserRepository, jwtUtils *utils.JWTUtils) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtUtils: jwtUtils,
	}
}

// Register - регистрация нового пользователя
func (s *AuthService) Register(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, string, error) {
	// Проверка существования пользователя
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err == nil && existingUser != nil {
		return nil, "", fmt.Errorf("user with username %s already exists", req.Username)
	}

	// Создание нового пользователя
	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password, // Будет захеширован в репозитории
		Role:         models.UserRoleUser,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	// Генерация JWT токена
	token, err := s.jwtUtils.GenerateToken(user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user.ToResponse(), token, nil
}

// Login - вход пользователя
func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.UserResponse, string, error) {
	// Проверка пароля и получение пользователя
	user, err := s.userRepo.VerifyPassword(ctx, req.Username, req.Password)
	if err != nil {
		return nil, "", fmt.Errorf("invalid credentials: %w", err)
	}

	// Генерация JWT токена
	token, err := s.jwtUtils.GenerateToken(user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user.ToResponse(), token, nil
}

// RefreshToken - обновление токена
func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	// Валидация и обновление токена
	newToken, err := s.jwtUtils.RefreshToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}

	return newToken, nil
}

// ValidateToken - валидация токена
func (s *AuthService) ValidateToken(tokenString string) (*utils.Claims, error) {
	claims, err := s.jwtUtils.ValidateToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}

// GetProfile - получение профиля пользователя
func (s *AuthService) GetProfile(ctx context.Context, userID string) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return user.ToResponse(), nil
}

// UpdateProfile - обновление профиля пользователя
func (s *AuthService) UpdateProfile(ctx context.Context, userID string, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	// Получение текущего пользователя
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Обновление полей если они указаны
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.AvatarURL != nil {
		user.AvatarURL = req.AvatarURL
	}
	if req.Role != nil {
		user.Role = *req.Role
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user.ToResponse(), nil
}
