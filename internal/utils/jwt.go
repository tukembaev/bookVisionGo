package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tukembaev/bookVisionGo/internal/config"
	"github.com/tukembaev/bookVisionGo/internal/models"
)

// Claims - структура JWT claims
type Claims struct {
	UserID   string          `json:"user_id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
	jwt.RegisteredClaims
}

// JWTUtils - утилиты для работы с JWT
type JWTUtils struct {
	secretKey string
	expiresIn int
}

// NewJWTUtils - создание нового JWT Utils
func NewJWTUtils(cfg *config.Config) *JWTUtils {
	return &JWTUtils{
		secretKey: cfg.JWT.SecretKey,
		expiresIn: cfg.JWT.ExpiresIn,
	}
}

// GenerateToken - генерация JWT токена
func (j *JWTUtils) GenerateToken(user *models.User) (string, error) {
	// Создание claims
	fmt.Println("DEBUG: expiresIn value is:", j.expiresIn)

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expiresIn))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bookVisionGo",
			Subject:   user.ID,
		},
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписание токена
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken - валидация JWT токена
func (j *JWTUtils) ValidateToken(tokenString string) (*Claims, error) {
	fmt.Printf("DEBUG: Validating token: %s\n", tokenString)
	fmt.Printf("DEBUG: Using secret key: %s\n", j.secretKey)

	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("DEBUG: Token method: %v\n", token.Method)
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		fmt.Printf("DEBUG: Parse error: %v\n", err)
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Проверка валидности токена
	if !token.Valid {
		fmt.Printf("DEBUG: Token is invalid\n")
		return nil, fmt.Errorf("invalid token")
	}

	// Извлечение claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		fmt.Printf("DEBUG: Invalid claims\n")
		return nil, fmt.Errorf("invalid token claims")
	}

	fmt.Printf("DEBUG: Token validated successfully for user: %s\n", claims.Username)
	return claims, nil
}

// RefreshToken - обновление токена
func (j *JWTUtils) RefreshToken(tokenString string) (string, error) {
	// Валидация текущего токена
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("invalid token for refresh: %w", err)
	}
	// Создание нового токена с тем же пользователем
	newClaims := &Claims{
		UserID:   claims.UserID,
		Username: claims.Username,
		Role:     claims.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expiresIn))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bookVisionGo",
			Subject:   claims.UserID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err = token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}

	return tokenString, nil
}

// ExtractTokenFromHeader - извлечение токена из Authorization header
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is required")
	}
	fmt.Println("Authorization header:", authHeader)
	// Проверка формата "Bearer <token>"
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", fmt.Errorf("authorization header format must be Bearer {token}")
	}

	return authHeader[len(bearerPrefix):], nil
}
