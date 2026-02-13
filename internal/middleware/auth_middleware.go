package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/services"
	"github.com/tukembaev/bookVisionGo/internal/utils"
)

// AuthMiddleware - middleware для проверки JWT токена
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлечение токена из заголовка
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Проверка формата Bearer token
		// const bearerPrefix = "Bearer "
		// if !strings.HasPrefix(authHeader, bearerPrefix) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString := authHeader[len(bearerPrefix):]

		// Валидация токена
		fmt.Println("AuthHeader:", authHeader)
		claims, err := authService.ValidateToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Сохранение информации о пользователе в контексте
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// RequireRole - middleware для проверки роли пользователя
func RequireRole(requiredRole models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		role, ok := userRole.(models.UserRole)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user role type"})
			c.Abort()
			return
		}

		// Проверка роли
		if !hasRequiredRole(role, requiredRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// OptionalAuth - middleware для опциональной аутентификации
func OptionalAuth(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.Next()
			return
		}

		tokenString := authHeader[len(bearerPrefix):]
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		// Сохранение информации о пользователе если токен валиден
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// GetCurrentUser - получение текущего пользователя из контекста
func GetCurrentUser(c *gin.Context) *utils.Claims {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")
	userRole, _ := c.Get("user_role")

	return &utils.Claims{
		UserID:   userID.(string),
		Username: username.(string),
		Role:     userRole.(models.UserRole),
	}
}

// IsAuthenticated - проверка аутентификации пользователя
func IsAuthenticated(c *gin.Context) bool {
	_, exists := c.Get("user_id")
	return exists
}

// hasRequiredRole - проверка необходимых прав доступа
func hasRequiredRole(userRole, requiredRole models.UserRole) bool {
	// Администратор имеет доступ ко всему
	if userRole == models.UserRoleAdmin {
		return true
	}

	// Модератор имеет доступ ко всему кроме админских функций
	if userRole == models.UserRoleModerator && requiredRole != models.UserRoleAdmin {
		return true
	}

	// Обычный пользователь имеет доступ только к пользовательским функциям
	return userRole == requiredRole && requiredRole == models.UserRoleUser
}
