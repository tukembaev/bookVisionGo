package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tukembaev/bookVisionGo/internal/middleware"
	"github.com/tukembaev/bookVisionGo/internal/models"
	"github.com/tukembaev/bookVisionGo/internal/services"
	"github.com/tukembaev/bookVisionGo/internal/utils"
)

// AuthHandler - обработчики аутентификации
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler - создание нового AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register - регистрация нового пользователя
// @Summary Регистрация пользователя
// @Description Создание нового аккаунта пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Данные для регистрации"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		if err.Error() == "user with username "+req.Username+" already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

// Login - вход пользователя
// @Summary Вход в систему
// @Description Аутентификация пользователя и получение JWT токена
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Данные для входа"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

// RefreshToken - обновление JWT токена
// @Summary Обновление токена
// @Description Получение нового JWT токена на основе существующего
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString, err := utils.ExtractTokenFromHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	newToken, err := h.authService.RefreshToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": newToken,
	})
}

// GetProfile - получение профиля текущего пользователя
// @Summary Получение профиля
// @Description Получение информации о текущем аутентифицированном пользователе
// @Tags auth
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/auth/profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)

	fmt.Println(currentUser)

	if currentUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := h.authService.GetProfile(c.Request.Context(), currentUser.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateProfile - обновление профиля пользователя
// @Summary Обновление профиля
// @Description Обновление информации о текущем пользователе
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param request body models.UpdateUserRequest true "Данные для обновления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/auth/profile [put]
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.UpdateProfile(c.Request.Context(), currentUser.UserID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Logout - выход пользователя (клиент должен удалить токен)
// @Summary Выход из системы
// @Description Выход пользователя (токен должен быть удален на клиенте)
// @Tags auth
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// В stateless JWT системе logout обрабатывается на клиенте
	// Здесь можно добавить логику для blacklisting токенов если нужно

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})

}
